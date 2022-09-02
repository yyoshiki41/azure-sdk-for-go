// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
	"github.com/stretchr/testify/require"
)

func Test_InMemoryCheckpointStore_Checkpoints(t *testing.T) {
	store := newCheckpointStoreForTest()

	checkpoints, err := store.ListCheckpoints(context.Background(), "ns", "eh", "cg", nil)
	require.NoError(t, err)
	require.Empty(t, checkpoints)

	for i := int64(0); i < 5; i++ {
		err = store.UpdateCheckpoint(context.Background(), Checkpoint{
			CheckpointStoreAddress: CheckpointStoreAddress{
				FullyQualifiedNamespace: "ns",
				EventHubName:            "eh",
				ConsumerGroup:           "cg",
				PartitionID:             "100",
			},
			CheckpointData: CheckpointData{
				Offset:         i,
				SequenceNumber: i + 1,
			},
		}, nil)
		require.NoError(t, err)

		checkpoints, err = store.ListCheckpoints(context.Background(), "ns", "eh", "cg", nil)
		require.NoError(t, err)

		require.Equal(t, []Checkpoint{
			{
				CheckpointStoreAddress: CheckpointStoreAddress{
					FullyQualifiedNamespace: "ns",
					EventHubName:            "eh",
					ConsumerGroup:           "cg",
					PartitionID:             "100",
				},
				CheckpointData: CheckpointData{
					Offset:         i,
					SequenceNumber: i + 1,
				},
			},
		}, checkpoints)
	}
}

func Test_InMemoryCheckpointStore_Ownership(t *testing.T) {
	store := newCheckpointStoreForTest()

	ownerships, err := store.ListOwnership(context.Background(), "ns", "eh", "cg", nil)
	require.NoError(t, err)
	require.Empty(t, ownerships)

	previousETag := ""

	for i := int64(0); i < 5; i++ {
		ownerships, err = store.ClaimOwnership(context.Background(), []Ownership{
			{
				CheckpointStoreAddress: CheckpointStoreAddress{
					FullyQualifiedNamespace: "ns",
					EventHubName:            "eh",
					ConsumerGroup:           "cg",
					PartitionID:             "100",
				},
				OwnershipData: OwnershipData{
					OwnerID:          "owner-id",
					LastModifiedTime: time.Time{},
					ETag:             previousETag,
				},
			}}, nil)
		require.NoError(t, err)

		expectedOwnership := Ownership{
			CheckpointStoreAddress: CheckpointStoreAddress{
				FullyQualifiedNamespace: "ns",
				EventHubName:            "eh",
				ConsumerGroup:           "cg",
				PartitionID:             "100",
			},
			OwnershipData: OwnershipData{
				OwnerID: "owner-id",
				// these fields are dynamically generated, so we just make sure
				// they do get filled out
				LastModifiedTime: ownerships[0].LastModifiedTime,
				ETag:             ownerships[0].ETag,
			},
		}

		require.NotEqual(t, previousETag, ownerships[0].ETag)
		require.NotZero(t, ownerships[0].LastModifiedTime)
		require.Equal(t, []Ownership{expectedOwnership}, ownerships)

		ownerships, err = store.ListOwnership(context.Background(), "ns", "eh", "cg", nil)
		require.NoError(t, err)

		require.NotEqual(t, previousETag, ownerships[0].ETag)
		require.NotZero(t, ownerships[0].LastModifiedTime)
		require.Equal(t, []Ownership{expectedOwnership}, ownerships)

		previousETag = ownerships[0].ETag
	}
}

func Test_InMemoryCheckpointStore_OwnershipLoss(t *testing.T) {
	store := newCheckpointStoreForTest()

	ownerships, err := store.ListOwnership(context.Background(), "ns", "eh", "cg", nil)
	require.NoError(t, err)
	require.Empty(t, ownerships)

	// If you don't specify an etag (ie, it's blank) then you always win ownership.
	ownerships, err = store.ClaimOwnership(context.Background(), []Ownership{
		{
			CheckpointStoreAddress: CheckpointStoreAddress{
				FullyQualifiedNamespace: "ns",
				EventHubName:            "eh",
				ConsumerGroup:           "cg",
				PartitionID:             "100",
			},
			OwnershipData: OwnershipData{
				OwnerID:          "owner-id",
				LastModifiedTime: time.Time{},
			},
		}}, nil)
	require.NoError(t, err)

	previousETag := ownerships[0].ETag

	// now let's try to claim the partition, but use an etag that doesn't match
	// the current one.
	ownerships, err = store.ClaimOwnership(context.Background(), []Ownership{
		{
			CheckpointStoreAddress: CheckpointStoreAddress{
				FullyQualifiedNamespace: "ns",
				EventHubName:            "eh",
				ConsumerGroup:           "cg",
				PartitionID:             "100",
			},
			OwnershipData: OwnershipData{
				OwnerID:          "new-owner-id",
				LastModifiedTime: time.Time{},
				ETag:             "non-matching-etag",
			},
		}}, nil)
	require.NoError(t, err)
	require.Empty(t, ownerships, "we weren't able to claim any partitions because our etag didn't match")

	ownerships, err = store.ListOwnership(context.Background(), "ns", "eh", "cg", nil)
	require.NoError(t, err)

	// note that the owner didn't change since our etag didn't match
	// this is expected to happen if we're fighting over ownership - someone will update
	// the ownership blob before us, and they're considered the owner from that point.
	require.Equal(t, "owner-id", ownerships[0].OwnerID)

	// okay, let's claim the partition properly (with a matching etag)
	ownerships, err = store.ClaimOwnership(context.Background(), []Ownership{
		{
			CheckpointStoreAddress: CheckpointStoreAddress{
				FullyQualifiedNamespace: "ns",
				EventHubName:            "eh",
				ConsumerGroup:           "cg",
				PartitionID:             "100",
			},
			OwnershipData: OwnershipData{
				OwnerID:          "new-owner-id",
				LastModifiedTime: time.Time{},
				ETag:             previousETag,
			},
		}}, nil)
	require.NoError(t, err)
	require.Equal(t, "new-owner-id", ownerships[0].OwnerID)

	ownerships, err = store.ListOwnership(context.Background(), "ns", "eh", "cg", nil)
	require.NoError(t, err)
	require.Equal(t, "new-owner-id", ownerships[0].OwnerID)
}

type testCheckpointStore struct {
	checkpointsMu sync.RWMutex
	checkpoints   map[string]Checkpoint

	ownershipMu sync.RWMutex
	ownerships  map[string]Ownership
}

func newCheckpointStoreForTest() *testCheckpointStore {
	return &testCheckpointStore{
		checkpoints: map[string]Checkpoint{},
		ownerships:  map[string]Ownership{},
	}
}

func (cps *testCheckpointStore) ExpireOwnership(o Ownership) {
	key := strings.Join([]string{o.CheckpointStoreAddress.FullyQualifiedNamespace, o.CheckpointStoreAddress.EventHubName, o.CheckpointStoreAddress.ConsumerGroup, o.CheckpointStoreAddress.PartitionID}, "/")

	cps.ownershipMu.Lock()
	defer cps.ownershipMu.Unlock()

	oldO := cps.ownerships[key]
	oldO.LastModifiedTime = time.Now().UTC().Add(-2 * time.Hour)
	cps.ownerships[key] = oldO
}

func (cps *testCheckpointStore) ClaimOwnership(ctx context.Context, partitionOwnership []Ownership, options *ClaimOwnershipOptions) ([]Ownership, error) {
	var owned []Ownership

	for _, po := range partitionOwnership {
		ownership, err := func(po Ownership) (*Ownership, error) {
			cps.ownershipMu.Lock()
			defer cps.ownershipMu.Unlock()

			if po.CheckpointStoreAddress.ConsumerGroup == "" ||
				po.CheckpointStoreAddress.EventHubName == "" ||
				po.CheckpointStoreAddress.FullyQualifiedNamespace == "" ||
				po.CheckpointStoreAddress.PartitionID == "" {
				panic("bad test, not all required fields were filled out for ownership data")
			}

			key := strings.Join([]string{po.CheckpointStoreAddress.FullyQualifiedNamespace, po.CheckpointStoreAddress.EventHubName, po.CheckpointStoreAddress.ConsumerGroup, po.CheckpointStoreAddress.PartitionID}, "/")

			current, exists := cps.ownerships[key]

			if exists && po.ETag != "" && current.ETag != po.ETag {
				// can't own it, didn't have the expected etag
				return nil, nil
			}

			newOwnership := po

			uuid, err := uuid.New()

			if err != nil {
				return nil, err
			}

			newOwnership.OwnershipData.ETag = uuid.String()
			newOwnership.OwnershipData.LastModifiedTime = time.Now().UTC()
			cps.ownerships[key] = newOwnership

			return &newOwnership, nil
		}(po)

		if err != nil {
			return nil, err
		}

		if ownership != nil {
			owned = append(owned, *ownership)
		}
	}

	return owned, nil
}

func (cps *testCheckpointStore) ListCheckpoints(ctx context.Context, fullyQualifiedNamespace string, eventHubName string, consumerGroup string, options *ListCheckpointsOptions) ([]Checkpoint, error) {
	cps.checkpointsMu.RLock()
	defer cps.checkpointsMu.RUnlock()

	var checkpoints []Checkpoint

	for _, v := range cps.checkpoints {
		checkpoints = append(checkpoints, v)
	}

	return checkpoints, nil
}

func (cps *testCheckpointStore) ListOwnership(ctx context.Context, fullyQualifiedNamespace string, eventHubName string, consumerGroup string, options *ListOwnershipOptions) ([]Ownership, error) {
	cps.ownershipMu.RLock()
	defer cps.ownershipMu.RUnlock()

	var ownerships []Ownership

	for _, v := range cps.ownerships {
		ownerships = append(ownerships, v)
	}

	return ownerships, nil
}

func (cps *testCheckpointStore) UpdateCheckpoint(ctx context.Context, checkpoint Checkpoint, options *UpdateCheckpointOptions) error {
	cps.checkpointsMu.Lock()
	defer cps.checkpointsMu.Unlock()

	if checkpoint.CheckpointStoreAddress.ConsumerGroup == "" ||
		checkpoint.CheckpointStoreAddress.EventHubName == "" ||
		checkpoint.CheckpointStoreAddress.FullyQualifiedNamespace == "" ||
		checkpoint.CheckpointStoreAddress.PartitionID == "" {
		panic("bad test, not all required fields were filled out for checkpoint data")
	}

	key := toInMemoryKey(checkpoint.CheckpointStoreAddress)
	cps.checkpoints[key] = checkpoint

	return nil
}

func toInMemoryKey(a CheckpointStoreAddress) string {
	return strings.Join([]string{a.FullyQualifiedNamespace, a.EventHubName, a.ConsumerGroup, a.PartitionID}, "/")
}
