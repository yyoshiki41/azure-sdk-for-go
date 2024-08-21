//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetHDInsightCapabilities.json
func ExampleLocationsClient_GetCapabilities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().GetCapabilities(ctx, "West US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CapabilitiesResult = armhdinsight.CapabilitiesResult{
	// 	Features: []*string{
	// 		to.Ptr("ALLOW_GATEWAY_AUTH_BYPASS"),
	// 		to.Ptr("AUX_IAAS_FUNCTIONALITIES"),
	// 		to.Ptr("BLOCK_APPS_FOR_SECURE_CLUSTER"),
	// 		to.Ptr("CLUSTER_RESIZE"),
	// 		to.Ptr("CLUSTERS_CONTRACT_1"),
	// 		to.Ptr("CLUSTERS_CONTRACT_1_SDK"),
	// 		to.Ptr("CLUSTERS_CONTRACT_2_SDK"),
	// 		to.Ptr("CLUSTERS_CONTRACT_VERSION_3_SDK"),
	// 		to.Ptr("COLLECTANONYMIZEDLOGS"),
	// 		to.Ptr("CONTAINER_LOG_COLLECTOR"),
	// 		to.Ptr("CONTAINER_RESOURCE_V3"),
	// 		to.Ptr("CUSTOM_ACTIONS_V2"),
	// 		to.Ptr("DATALAKE"),
	// 		to.Ptr("DATALAKE_3_2"),
	// 		to.Ptr("DATALAKE_3_4"),
	// 		to.Ptr("DATALAKE_3_5"),
	// 		to.Ptr("DATALAKE_DEFAULTFS_3_5"),
	// 		to.Ptr("DEFAULT_CONTAINER_IDENTIFIER_AND_STORAGEFQDN_ALREADY_USED"),
	// 		to.Ptr("DOCUMENT_VALIDATION_IN_API"),
	// 		to.Ptr("ENABLEDATALAKE"),
	// 		to.Ptr("ENABLEGENEVAANALYTICS"),
	// 		to.Ptr("HADOOP_VIRTUAL_NETWORK_ENABLED"),
	// 		to.Ptr("HS2_ZK_ROUTER_INTERACTIVEHIVE"),
	// 		to.Ptr("IAAS_ALLOW_CUSTOM_DNS"),
	// 		to.Ptr("IAAS_AMBARI_APP_TIMELINE_SERVER_HA_SERVICE"),
	// 		to.Ptr("IAAS_AMBARI_DEPENDENCY_BASED_DEPLOYMENT"),
	// 		to.Ptr("IAAS_AMBARI_HA_SERVICES"),
	// 		to.Ptr("IAAS_AMBARI_HISTORYSERVER_HA_SERVICE"),
	// 		to.Ptr("IAAS_AMBARI_LOWER_LATENCY"),
	// 		to.Ptr("IAAS_AMBARI_SKIP_COMPONENTS_INSTALL"),
	// 		to.Ptr("IAAS_ARM_PROVISIONING"),
	// 		to.Ptr("IAAS_CLOSE_HEAD_HTTPS_END_POINT_AFTER_CLUSTER_CREATE"),
	// 		to.Ptr("IAAS_CLUSTER_APPLICATION_ALLOW_MULTIPLE_ROLE_INSTANCES"),
	// 		to.Ptr("IAAS_CLUSTER_APPLICATION_REMOVE"),
	// 		to.Ptr("IAAS_CLUSTER_CONTAINER_PREUPLOAD_SUBMIT_WAIT"),
	// 		to.Ptr("IAAS_CLUSTER_RSERVER"),
	// 		to.Ptr("IAAS_DB_CREATION_IN_PARALLEL_TO_VM"),
	// 		to.Ptr("IAAS_DELETE_LEAKED_RESOURCES"),
	// 		to.Ptr("IAAS_DEPLOYMENTS"),
	// 		to.Ptr("IAAS_DO_NOT_CREATE_WASB_TABLES_IN_CUSTOMER_STORAGE"),
	// 		to.Ptr("IAAS_ENABLE_CLUSTER_CONFIG_OVERRIDES"),
	// 		to.Ptr("IAAS_ENABLE_OFFLINE_CLEANUP"),
	// 		to.Ptr("IAAS_ENABLE_OFFLINE_CLEANUP_FOR_DELETING_VMS"),
	// 		to.Ptr("IAAS_INCLUDE_STORAGE_IN_SUBSCRIPTION_SELECTION"),
	// 		to.Ptr("IAAS_PARALLEL_DB_CREATE"),
	// 		to.Ptr("IAAS_PREPROVISION_METASTORES_SCHEMAS"),
	// 		to.Ptr("IAAS_SCRIPTACTIONS_DELETE_VMS_CRUD_FAILURES"),
	// 		to.Ptr("IAAS_SCRIPTACTIONS_RUNNING"),
	// 		to.Ptr("IAAS_SHORT_VM_NAME"),
	// 		to.Ptr("IAAS_SUBMIT_AMBARI_REQUEST_ONCE_LINUX_VM_ARE_AVAILABLE"),
	// 		to.Ptr("IAAS_TEZ_ATS_V15"),
	// 		to.Ptr("IAAS_USE_UNATTENDED_UPGRADES_FOR_PATCHING"),
	// 		to.Ptr("IAAS_VALIDATE_CUSTOM_VNET"),
	// 		to.Ptr("IAAS_VALIDATE_NSG"),
	// 		to.Ptr("IAAS_WAIT_FOR_CLOSING_HEAD_HTTPS_END_POINT_AFTER_CLUSTER_CREATE"),
	// 		to.Ptr("IAAS_YARN_HDINSIGHT_SQL_TIMELINE_STORE"),
	// 		to.Ptr("IAAS_YARN_HDINSIGHT_TIMELINE_STORE"),
	// 		to.Ptr("INTERACTIVEHIVE"),
	// 		to.Ptr("MDSCENTRALLOGGING"),
	// 		to.Ptr("NODE_SETUP_POLLER_ENABLED"),
	// 		to.Ptr("OVERPROVISION_HOSTGROUP_edgenode"),
	// 		to.Ptr("OVERPROVISION_HOSTGROUP_Gateway"),
	// 		to.Ptr("OVERPROVISION_HOSTGROUP_Workernode"),
	// 		to.Ptr("OVERPROVISION_HOSTGROUP_zookeepernode"),
	// 		to.Ptr("PERF_OPTIMIZED_RESOURCE_LOCATION_FETCH"),
	// 		to.Ptr("PORTALAPPINSTALL"),
	// 		to.Ptr("POWERSHELL_SCRIPT_ACTION"),
	// 		to.Ptr("POWERSHELL_SCRIPT_ACTION_SDK"),
	// 		to.Ptr("PREMIUM_TIER_PREVIEW"),
	// 		to.Ptr("PROVISIONING_AGENT"),
	// 		to.Ptr("RMHA"),
	// 		to.Ptr("RSERVER_CLUSTERTYPE_3_5_ENABLED"),
	// 		to.Ptr("RSERVER_CLUSTERTYPE_ENABLED"),
	// 		to.Ptr("SHOW_HUMBOLDT_GA"),
	// 		to.Ptr("SHOW_IBIZA_CREATE"),
	// 		to.Ptr("SPARK_2_1"),
	// 		to.Ptr("SPARK_EXPERIMENTAL"),
	// 		to.Ptr("STORM_PREVIEW"),
	// 		to.Ptr("UI_CREATE_WIZARD_V2"),
	// 		to.Ptr("VIRTUAL_NETWORK_ENABLED"),
	// 		to.Ptr("VMSIZES_AUX")},
	// 		Quota: &armhdinsight.QuotaCapability{
	// 			CoresUsed: to.Ptr[int64](0),
	// 			MaxCoresAllowed: to.Ptr[int64](3000),
	// 			RegionalQuotas: []*armhdinsight.RegionalQuotaCapability{
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Australia East"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Australia Southeast"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Brazil South"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Canada Central"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Canada East"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Central India"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](24),
	// 					RegionName: to.Ptr("Central US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Central US EUAP"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](24),
	// 					RegionName: to.Ptr("East Asia"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](24),
	// 					RegionName: to.Ptr("East US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("East US 2"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Japan East"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Japan West"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("North Central US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("North Europe"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("South Central US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("Southeast Asia"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("West Central US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("West Europe"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](132),
	// 					RegionName: to.Ptr("West US"),
	// 				},
	// 				{
	// 					CoresAvailable: to.Ptr[int64](1000),
	// 					CoresUsed: to.Ptr[int64](0),
	// 					RegionName: to.Ptr("West US 2"),
	// 			}},
	// 		},
	// 		Regions: map[string]*armhdinsight.RegionsCapability{
	// 			"iaas": &armhdinsight.RegionsCapability{
	// 				Available: []*string{
	// 					to.Ptr("AUSTRALIA EAST"),
	// 					to.Ptr("AUSTRALIA SOUTHEAST"),
	// 					to.Ptr("BRAZIL SOUTH"),
	// 					to.Ptr("CANADA CENTRAL"),
	// 					to.Ptr("CANADA EAST"),
	// 					to.Ptr("CENTRAL INDIA"),
	// 					to.Ptr("CENTRAL US"),
	// 					to.Ptr("CENTRAL US EUAP"),
	// 					to.Ptr("EAST ASIA"),
	// 					to.Ptr("EAST US"),
	// 					to.Ptr("EAST US 2"),
	// 					to.Ptr("JAPAN EAST"),
	// 					to.Ptr("JAPAN WEST"),
	// 					to.Ptr("NORTH CENTRAL US"),
	// 					to.Ptr("NORTH EUROPE"),
	// 					to.Ptr("SOUTH CENTRAL US"),
	// 					to.Ptr("SOUTHEAST ASIA"),
	// 					to.Ptr("WEST CENTRAL US"),
	// 					to.Ptr("WEST EUROPE"),
	// 					to.Ptr("WEST US"),
	// 					to.Ptr("WEST US 2")},
	// 				},
	// 				"paas": &armhdinsight.RegionsCapability{
	// 					Available: []*string{
	// 						to.Ptr("AUSTRALIA EAST"),
	// 						to.Ptr("AUSTRALIA SOUTHEAST"),
	// 						to.Ptr("BRAZIL SOUTH"),
	// 						to.Ptr("CENTRAL INDIA"),
	// 						to.Ptr("CENTRAL US"),
	// 						to.Ptr("EAST ASIA"),
	// 						to.Ptr("EAST US"),
	// 						to.Ptr("EAST US 2"),
	// 						to.Ptr("JAPAN EAST"),
	// 						to.Ptr("JAPAN WEST"),
	// 						to.Ptr("NORTH CENTRAL US"),
	// 						to.Ptr("NORTH EUROPE"),
	// 						to.Ptr("SOUTH CENTRAL US"),
	// 						to.Ptr("SOUTHEAST ASIA"),
	// 						to.Ptr("WEST EUROPE"),
	// 						to.Ptr("WEST US"),
	// 						to.Ptr("WEST US 2")},
	// 					},
	// 				},
	// 				Versions: map[string]*armhdinsight.VersionsCapability{
	// 					"iaas": &armhdinsight.VersionsCapability{
	// 						Available: []*armhdinsight.VersionSpec{
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("0.98.4"),
	// 									"HDP": to.Ptr("2.2"),
	// 									"Hadoop": to.Ptr("2.6.0"),
	// 									"Storm": to.Ptr("0.9.3"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.2.1000.0.8840373"),
	// 								FriendlyName: to.Ptr("3.2"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("1.1.2"),
	// 									"HDP": to.Ptr("2.3"),
	// 									"Hadoop": to.Ptr("2.7.0"),
	// 									"Spark": to.Ptr("1.5.2"),
	// 									"Storm": to.Ptr("0.10.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.3.1000.0.9776961"),
	// 								FriendlyName: to.Ptr("3.3"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("1.1.2"),
	// 									"HDP": to.Ptr("2.4"),
	// 									"Hadoop": to.Ptr("2.7.1"),
	// 									"RServer": to.Ptr("8.0"),
	// 									"Spark": to.Ptr("1.6.2"),
	// 									"Storm": to.Ptr("0.10.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.4.1000.0.9719475"),
	// 								FriendlyName: to.Ptr("3.4"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("1.1.2"),
	// 									"HDP": to.Ptr("2.5"),
	// 									"Hadoop": to.Ptr("2.7.3"),
	// 									"InteractiveHive": to.Ptr("2.1.0"),
	// 									"RServer": to.Ptr("9.0"),
	// 									"Spark": to.Ptr("1.6.3,2.0.2"),
	// 									"Storm": to.Ptr("1.0.1"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.5.1000.0.9732704"),
	// 								FriendlyName: to.Ptr("3.5"),
	// 								IsDefault: to.Ptr(true),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HDP": to.Ptr("2.6"),
	// 									"Spark": to.Ptr("2.1.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.6.1000.0.9503998"),
	// 								FriendlyName: to.Ptr("3.6"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								DisplayName: to.Ptr("Version 99.152.1000.0.6943836"),
	// 								FriendlyName: to.Ptr("99.152"),
	// 								IsDefault: to.Ptr(false),
	// 						}},
	// 					},
	// 					"paas": &armhdinsight.VersionsCapability{
	// 						Available: []*armhdinsight.VersionSpec{
	// 							{
	// 								DisplayName: to.Ptr("HdInsight version 1.6.1.0.335536"),
	// 								FriendlyName: to.Ptr("1.6"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HDP": to.Ptr("1.3"),
	// 									"Hadoop": to.Ptr("1.2.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 2.1.9.406.1221105"),
	// 								FriendlyName: to.Ptr("2.1"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HDP": to.Ptr("2.0"),
	// 									"Hadoop": to.Ptr("2.2.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.0.6.989.2441725"),
	// 								FriendlyName: to.Ptr("3.0"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("0.98"),
	// 									"HDP": to.Ptr("2.1.7"),
	// 									"Hadoop": to.Ptr("2.4.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.1.4.989.2441725"),
	// 								FriendlyName: to.Ptr("3.1"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("0.98.4"),
	// 									"HDP": to.Ptr("2.2"),
	// 									"Hadoop": to.Ptr("2.6.0"),
	// 									"Storm": to.Ptr("0.9.3"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.2.7.989.2441725"),
	// 								FriendlyName: to.Ptr("3.2"),
	// 								IsDefault: to.Ptr(false),
	// 							},
	// 							{
	// 								ComponentVersions: map[string]*string{
	// 									"HBase": to.Ptr("1.1.2"),
	// 									"HDP": to.Ptr("2.3"),
	// 									"Hadoop": to.Ptr("2.7.0"),
	// 									"Storm": to.Ptr("0.10.0"),
	// 								},
	// 								DisplayName: to.Ptr("Version 3.3.0.989.2441725"),
	// 								FriendlyName: to.Ptr("3.3"),
	// 								IsDefault: to.Ptr(true),
	// 						}},
	// 					},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetHDInsightUsages.json
func ExampleLocationsClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().ListUsages(ctx, "West US", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UsagesListResult = armhdinsight.UsagesListResult{
	// 	Value: []*armhdinsight.Usage{
	// 		{
	// 			Name: &armhdinsight.LocalizedName{
	// 				LocalizedValue: to.Ptr("Cores"),
	// 				Value: to.Ptr("cores"),
	// 			},
	// 			CurrentValue: to.Ptr[int64](0),
	// 			Limit: to.Ptr[int64](5000),
	// 			Unit: to.Ptr("Count"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Locations_ListBillingSpecs.json
func ExampleLocationsClient_ListBillingSpecs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().ListBillingSpecs(ctx, "East US 2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BillingResponseListResult = armhdinsight.BillingResponseListResult{
	// 	BillingResources: []*armhdinsight.BillingResources{
	// 		{
	// 			BillingMeters: []*armhdinsight.BillingMeters{
	// 			},
	// 			DiskBillingMeters: []*armhdinsight.DiskBillingMeters{
	// 			},
	// 			Region: to.Ptr("East US 2"),
	// 		},
	// 		{
	// 			BillingMeters: []*armhdinsight.BillingMeters{
	// 				{
	// 					Meter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					MeterParameter: to.Ptr("default"),
	// 					Unit: to.Ptr("CoreHours"),
	// 				},
	// 				{
	// 					Meter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					MeterParameter: to.Ptr("Kafka"),
	// 					Unit: to.Ptr("CoreHours"),
	// 			}},
	// 			DiskBillingMeters: []*armhdinsight.DiskBillingMeters{
	// 				{
	// 					DiskRpMeter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					SKU: to.Ptr("All"),
	// 					Tier: to.Ptr(armhdinsight.TierStandard),
	// 				},
	// 				{
	// 					DiskRpMeter: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					SKU: to.Ptr("All"),
	// 					Tier: to.Ptr(armhdinsight.TierStandard),
	// 			}},
	// 			Region: to.Ptr("default"),
	// 	}},
	// 	VMSizeFilters: []*armhdinsight.VMSizeCompatibilityFilterV2{
	// 		{
	// 			FilterMode: to.Ptr(armhdinsight.FilterModeExclude),
	// 	}},
	// 	VMSizes: []*string{
	// 		to.Ptr("A5"),
	// 		to.Ptr("A6"),
	// 		to.Ptr("A7")},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Locations_GetAsyncOperationStatus.json
func ExampleLocationsClient_GetAzureAsyncOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().GetAzureAsyncOperationStatus(ctx, "East US 2", "8a0348f4-8a85-4ec2-abe0-03b26104a9a0-0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AsyncOperationResult = armhdinsight.AsyncOperationResult{
	// 	Status: to.Ptr(armhdinsight.AsyncOperationStateSucceeded),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Locations_CheckClusterNameAvailability.json
func ExampleLocationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().CheckNameAvailability(ctx, "westus", armhdinsight.NameAvailabilityCheckRequestParameters{
		Name: to.Ptr("test123"),
		Type: to.Ptr("clusters"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NameAvailabilityCheckResult = armhdinsight.NameAvailabilityCheckResult{
	// 	Message: to.Ptr("Cluster name 'test123' is unavailable"),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("AlreadyExists"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/aec83a5f0ed56da4fd16fa027b9fa27edfa8988b/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/HDI_Locations_ValidateClusterCreateRequest.json
func ExampleLocationsClient_ValidateClusterCreateRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocationsClient().ValidateClusterCreateRequest(ctx, "southcentralus", armhdinsight.ClusterCreateRequestValidationParameters{
		Location: to.Ptr("southcentralus"),
		Properties: &armhdinsight.ClusterCreateProperties{
			ClusterDefinition: &armhdinsight.ClusterDefinition{
				ComponentVersion: map[string]*string{
					"Spark": to.Ptr("2.4"),
				},
				Configurations: map[string]any{
					"gateway": map[string]any{
						"restAuthCredential.isEnabled": true,
						"restAuthCredential.password":  "**********",
						"restAuthCredential.username":  "admin",
					},
				},
				Kind: to.Ptr("spark"),
			},
			ClusterVersion: to.Ptr("4.0"),
			ComputeProfile: &armhdinsight.ComputeProfile{
				Roles: []*armhdinsight.Role{
					{
						Name: to.Ptr("headnode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_E8_V3"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](2),
					},
					{
						Name: to.Ptr("workernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_E8_V3"),
						},
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](4),
					},
					{
						Name: to.Ptr("zookeepernode"),
						HardwareProfile: &armhdinsight.HardwareProfile{
							VMSize: to.Ptr("Standard_D13_V2"),
						},
						MinInstanceCount: to.Ptr[int32](1),
						OSProfile: &armhdinsight.OsProfile{
							LinuxOperatingSystemProfile: &armhdinsight.LinuxOperatingSystemProfile{
								Password: to.Ptr("**********"),
								Username: to.Ptr("sshuser"),
							},
						},
						ScriptActions:       []*armhdinsight.ScriptAction{},
						TargetInstanceCount: to.Ptr[int32](3),
					}},
			},
			MinSupportedTLSVersion: to.Ptr("1.2"),
			OSType:                 to.Ptr(armhdinsight.OSTypeLinux),
			StorageProfile: &armhdinsight.StorageProfile{
				Storageaccounts: []*armhdinsight.StorageAccount{
					{
						Name:                to.Ptr("storagename.blob.core.windows.net"),
						Container:           to.Ptr("contianername"),
						EnableSecureChannel: to.Ptr(true),
						IsDefault:           to.Ptr(true),
						Key:                 to.Ptr("*******"),
						ResourceID:          to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/storagename"),
					}},
			},
			Tier: to.Ptr(armhdinsight.TierStandard),
		},
		Tags:               map[string]*string{},
		Name:               to.Ptr("testclustername"),
		Type:               to.Ptr("Microsoft.HDInsight/clusters"),
		FetchAaddsResource: to.Ptr(false),
		TenantID:           to.Ptr("00000000-0000-0000-0000-000000000000"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterCreateValidationResult = armhdinsight.ClusterCreateValidationResult{
	// 	EstimatedCreationDuration: to.Ptr("PT20M"),
	// 	ValidationErrors: []*armhdinsight.ValidationErrorInfo{
	// 	},
	// }
}
