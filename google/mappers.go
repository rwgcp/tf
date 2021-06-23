// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"sort"
)

type convertFunc func(d TerraformResourceData, config *Config) ([]Asset, error)

// fetchFunc allows initial data for a resource to be fetched from the API and merged
// with the planned changes. This is useful for resources that are only partially managed
// by Terraform, like IAM policies managed with member/binding resources.
type fetchFunc func(d TerraformResourceData, config *Config) (Asset, error)

// mergeFunc combines multiple terraform resources into a single CAI asset.
// The incoming asset will either be an asset that was created/updated or deleted.
type mergeFunc func(existing, incoming Asset) Asset

// mapper pairs related conversion/merging functions.
type Mapper struct {
	convert           convertFunc // required
	fetch             fetchFunc   // optional
	mergeCreateUpdate mergeFunc   // optional
	mergeDelete       mergeFunc   // optional
}

// mappers maps terraform resource types (i.e. `google_project`) into
// a slice of mapperFuncs.
//
// Modelling of relationships:
// terraform resources to CAI assets as []mapperFuncs:
// 1:1 = [mapper{convert: convertAbc}]                  (len=1)
// 1:N = [mapper{convert: convertAbc}, ...]             (len=N)
// N:1 = [mapper{convert: convertAbc, merge: mergeAbc}] (len=1)
func Mappers() map[string][]Mapper {
	return map[string][]Mapper{
		"google_compute_firewall":               {{convert: GetComputeFirewallCaiObject}},
		"google_compute_disk":                   {{convert: GetComputeDiskCaiObject}},
		"google_compute_forwarding_rule":        {{convert: GetComputeForwardingRuleCaiObject}},
		"google_compute_global_forwarding_rule": {{convert: GetComputeGlobalForwardingRuleCaiObject}},
		"google_compute_instance":               {{convert: GetComputeInstanceCaiObject}},
		"google_compute_network":                {{convert: GetComputeNetworkCaiObject}},
		"google_compute_subnetwork":             {{convert: GetComputeSubnetworkCaiObject}},
		"google_storage_bucket":                 {{convert: GetStorageBucketCaiObject}},
		"google_sql_database_instance":          {{convert: GetSQLDatabaseInstanceCaiObject}},
		"google_container_cluster":              {{convert: GetContainerClusterCaiObject}},
		"google_container_node_pool":            {{convert: GetContainerNodePoolCaiObject}},
		"google_bigquery_dataset":               {{convert: GetBigQueryDatasetCaiObject}},
		"google_spanner_instance":               {{convert: GetSpannerInstanceCaiObject}},
		"google_project_service":                {{convert: GetServiceUsageCaiObject}},
		"google_pubsub_subscription":            {{convert: GetPubsubSubscriptionCaiObject}},
		"google_pubsub_topic":                   {{convert: GetPubsubTopicCaiObject}},
		"google_kms_crypto_key":                 {{convert: GetKMSCryptoKeyCaiObject}},
		"google_kms_key_ring":                   {{convert: GetKMSKeyRingCaiObject}},
		"google_filestore_instance":             {{convert: GetFilestoreInstanceCaiObject}},
		"google_bigquery_table_iam_policy": {
			{
				convert:           GetBigQueryTableIamPolicyCaiObject,
				mergeCreateUpdate: MergeBigQueryTableIamPolicy,
			},
		},
		"google_bigquery_table_iam_binding": {
			{
				convert:           GetBigQueryTableIamBindingCaiObject,
				mergeCreateUpdate: MergeBigQueryTableIamBinding,
				mergeDelete:       MergeBigQueryTableIamBindingDelete,
				fetch:             FetchBigQueryTableIamPolicy,
			},
		},
		"google_bigquery_table_iam_member": {
			{
				convert:           GetBigQueryTableIamMemberCaiObject,
				mergeCreateUpdate: MergeBigQueryTableIamMember,
				mergeDelete:       MergeBigQueryTableIamMemberDelete,
				fetch:             FetchBigQueryTableIamPolicy,
			},
		},
		"google_binary_authorization_attestor_iam_policy": {
			{
				convert:           GetBinaryAuthorizationAttestorIamPolicyCaiObject,
				mergeCreateUpdate: MergeBinaryAuthorizationAttestorIamPolicy,
			},
		},
		"google_binary_authorization_attestor_iam_binding": {
			{
				convert:           GetBinaryAuthorizationAttestorIamBindingCaiObject,
				mergeCreateUpdate: MergeBinaryAuthorizationAttestorIamBinding,
				mergeDelete:       MergeBinaryAuthorizationAttestorIamBindingDelete,
				fetch:             FetchBinaryAuthorizationAttestorIamPolicy,
			},
		},
		"google_binary_authorization_attestor_iam_member": {
			{
				convert:           GetBinaryAuthorizationAttestorIamMemberCaiObject,
				mergeCreateUpdate: MergeBinaryAuthorizationAttestorIamMember,
				mergeDelete:       MergeBinaryAuthorizationAttestorIamMemberDelete,
				fetch:             FetchBinaryAuthorizationAttestorIamPolicy,
			},
		},
		"google_cloudfunctions_function_iam_policy": {
			{
				convert:           GetCloudFunctionsCloudFunctionIamPolicyCaiObject,
				mergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamPolicy,
			},
		},
		"google_cloudfunctions_function_iam_binding": {
			{
				convert:           GetCloudFunctionsCloudFunctionIamBindingCaiObject,
				mergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamBinding,
				mergeDelete:       MergeCloudFunctionsCloudFunctionIamBindingDelete,
				fetch:             FetchCloudFunctionsCloudFunctionIamPolicy,
			},
		},
		"google_cloudfunctions_function_iam_member": {
			{
				convert:           GetCloudFunctionsCloudFunctionIamMemberCaiObject,
				mergeCreateUpdate: MergeCloudFunctionsCloudFunctionIamMember,
				mergeDelete:       MergeCloudFunctionsCloudFunctionIamMemberDelete,
				fetch:             FetchCloudFunctionsCloudFunctionIamPolicy,
			},
		},
		"google_cloud_run_service_iam_policy": {
			{
				convert:           GetCloudRunServiceIamPolicyCaiObject,
				mergeCreateUpdate: MergeCloudRunServiceIamPolicy,
			},
		},
		"google_cloud_run_service_iam_binding": {
			{
				convert:           GetCloudRunServiceIamBindingCaiObject,
				mergeCreateUpdate: MergeCloudRunServiceIamBinding,
				mergeDelete:       MergeCloudRunServiceIamBindingDelete,
				fetch:             FetchCloudRunServiceIamPolicy,
			},
		},
		"google_cloud_run_service_iam_member": {
			{
				convert:           GetCloudRunServiceIamMemberCaiObject,
				mergeCreateUpdate: MergeCloudRunServiceIamMember,
				mergeDelete:       MergeCloudRunServiceIamMemberDelete,
				fetch:             FetchCloudRunServiceIamPolicy,
			},
		},
		"google_compute_disk_iam_policy": {
			{
				convert:           GetComputeDiskIamPolicyCaiObject,
				mergeCreateUpdate: MergeComputeDiskIamPolicy,
			},
		},
		"google_compute_disk_iam_binding": {
			{
				convert:           GetComputeDiskIamBindingCaiObject,
				mergeCreateUpdate: MergeComputeDiskIamBinding,
				mergeDelete:       MergeComputeDiskIamBindingDelete,
				fetch:             FetchComputeDiskIamPolicy,
			},
		},
		"google_compute_disk_iam_member": {
			{
				convert:           GetComputeDiskIamMemberCaiObject,
				mergeCreateUpdate: MergeComputeDiskIamMember,
				mergeDelete:       MergeComputeDiskIamMemberDelete,
				fetch:             FetchComputeDiskIamPolicy,
			},
		},
		"google_compute_image_iam_policy": {
			{
				convert:           GetComputeImageIamPolicyCaiObject,
				mergeCreateUpdate: MergeComputeImageIamPolicy,
			},
		},
		"google_compute_image_iam_binding": {
			{
				convert:           GetComputeImageIamBindingCaiObject,
				mergeCreateUpdate: MergeComputeImageIamBinding,
				mergeDelete:       MergeComputeImageIamBindingDelete,
				fetch:             FetchComputeImageIamPolicy,
			},
		},
		"google_compute_image_iam_member": {
			{
				convert:           GetComputeImageIamMemberCaiObject,
				mergeCreateUpdate: MergeComputeImageIamMember,
				mergeDelete:       MergeComputeImageIamMemberDelete,
				fetch:             FetchComputeImageIamPolicy,
			},
		},
		"google_compute_instance_iam_policy": {
			{
				convert:           GetComputeInstanceIamPolicyCaiObject,
				mergeCreateUpdate: MergeComputeInstanceIamPolicy,
			},
		},
		"google_compute_instance_iam_binding": {
			{
				convert:           GetComputeInstanceIamBindingCaiObject,
				mergeCreateUpdate: MergeComputeInstanceIamBinding,
				mergeDelete:       MergeComputeInstanceIamBindingDelete,
				fetch:             FetchComputeInstanceIamPolicy,
			},
		},
		"google_compute_instance_iam_member": {
			{
				convert:           GetComputeInstanceIamMemberCaiObject,
				mergeCreateUpdate: MergeComputeInstanceIamMember,
				mergeDelete:       MergeComputeInstanceIamMemberDelete,
				fetch:             FetchComputeInstanceIamPolicy,
			},
		},
		"google_compute_region_disk_iam_policy": {
			{
				convert:           GetComputeRegionDiskIamPolicyCaiObject,
				mergeCreateUpdate: MergeComputeRegionDiskIamPolicy,
			},
		},
		"google_compute_region_disk_iam_binding": {
			{
				convert:           GetComputeRegionDiskIamBindingCaiObject,
				mergeCreateUpdate: MergeComputeRegionDiskIamBinding,
				mergeDelete:       MergeComputeRegionDiskIamBindingDelete,
				fetch:             FetchComputeRegionDiskIamPolicy,
			},
		},
		"google_compute_region_disk_iam_member": {
			{
				convert:           GetComputeRegionDiskIamMemberCaiObject,
				mergeCreateUpdate: MergeComputeRegionDiskIamMember,
				mergeDelete:       MergeComputeRegionDiskIamMemberDelete,
				fetch:             FetchComputeRegionDiskIamPolicy,
			},
		},
		"google_compute_subnetwork_iam_policy": {
			{
				convert:           GetComputeSubnetworkIamPolicyCaiObject,
				mergeCreateUpdate: MergeComputeSubnetworkIamPolicy,
			},
		},
		"google_compute_subnetwork_iam_binding": {
			{
				convert:           GetComputeSubnetworkIamBindingCaiObject,
				mergeCreateUpdate: MergeComputeSubnetworkIamBinding,
				mergeDelete:       MergeComputeSubnetworkIamBindingDelete,
				fetch:             FetchComputeSubnetworkIamPolicy,
			},
		},
		"google_compute_subnetwork_iam_member": {
			{
				convert:           GetComputeSubnetworkIamMemberCaiObject,
				mergeCreateUpdate: MergeComputeSubnetworkIamMember,
				mergeDelete:       MergeComputeSubnetworkIamMemberDelete,
				fetch:             FetchComputeSubnetworkIamPolicy,
			},
		},
		"google_data_catalog_entry_group_iam_policy": {
			{
				convert:           GetDataCatalogEntryGroupIamPolicyCaiObject,
				mergeCreateUpdate: MergeDataCatalogEntryGroupIamPolicy,
			},
		},
		"google_data_catalog_entry_group_iam_binding": {
			{
				convert:           GetDataCatalogEntryGroupIamBindingCaiObject,
				mergeCreateUpdate: MergeDataCatalogEntryGroupIamBinding,
				mergeDelete:       MergeDataCatalogEntryGroupIamBindingDelete,
				fetch:             FetchDataCatalogEntryGroupIamPolicy,
			},
		},
		"google_data_catalog_entry_group_iam_member": {
			{
				convert:           GetDataCatalogEntryGroupIamMemberCaiObject,
				mergeCreateUpdate: MergeDataCatalogEntryGroupIamMember,
				mergeDelete:       MergeDataCatalogEntryGroupIamMemberDelete,
				fetch:             FetchDataCatalogEntryGroupIamPolicy,
			},
		},
		"google_data_catalog_tag_template_iam_policy": {
			{
				convert:           GetDataCatalogTagTemplateIamPolicyCaiObject,
				mergeCreateUpdate: MergeDataCatalogTagTemplateIamPolicy,
			},
		},
		"google_data_catalog_tag_template_iam_binding": {
			{
				convert:           GetDataCatalogTagTemplateIamBindingCaiObject,
				mergeCreateUpdate: MergeDataCatalogTagTemplateIamBinding,
				mergeDelete:       MergeDataCatalogTagTemplateIamBindingDelete,
				fetch:             FetchDataCatalogTagTemplateIamPolicy,
			},
		},
		"google_data_catalog_tag_template_iam_member": {
			{
				convert:           GetDataCatalogTagTemplateIamMemberCaiObject,
				mergeCreateUpdate: MergeDataCatalogTagTemplateIamMember,
				mergeDelete:       MergeDataCatalogTagTemplateIamMemberDelete,
				fetch:             FetchDataCatalogTagTemplateIamPolicy,
			},
		},
		"google_healthcare_consent_store_iam_policy": {
			{
				convert:           GetHealthcareConsentStoreIamPolicyCaiObject,
				mergeCreateUpdate: MergeHealthcareConsentStoreIamPolicy,
			},
		},
		"google_healthcare_consent_store_iam_binding": {
			{
				convert:           GetHealthcareConsentStoreIamBindingCaiObject,
				mergeCreateUpdate: MergeHealthcareConsentStoreIamBinding,
				mergeDelete:       MergeHealthcareConsentStoreIamBindingDelete,
				fetch:             FetchHealthcareConsentStoreIamPolicy,
			},
		},
		"google_healthcare_consent_store_iam_member": {
			{
				convert:           GetHealthcareConsentStoreIamMemberCaiObject,
				mergeCreateUpdate: MergeHealthcareConsentStoreIamMember,
				mergeDelete:       MergeHealthcareConsentStoreIamMemberDelete,
				fetch:             FetchHealthcareConsentStoreIamPolicy,
			},
		},
		"google_iap_web_iam_policy": {
			{
				convert:           GetIapWebIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapWebIamPolicy,
			},
		},
		"google_iap_web_iam_binding": {
			{
				convert:           GetIapWebIamBindingCaiObject,
				mergeCreateUpdate: MergeIapWebIamBinding,
				mergeDelete:       MergeIapWebIamBindingDelete,
				fetch:             FetchIapWebIamPolicy,
			},
		},
		"google_iap_web_iam_member": {
			{
				convert:           GetIapWebIamMemberCaiObject,
				mergeCreateUpdate: MergeIapWebIamMember,
				mergeDelete:       MergeIapWebIamMemberDelete,
				fetch:             FetchIapWebIamPolicy,
			},
		},
		"google_iap_web_type_compute_iam_policy": {
			{
				convert:           GetIapWebTypeComputeIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapWebTypeComputeIamPolicy,
			},
		},
		"google_iap_web_type_compute_iam_binding": {
			{
				convert:           GetIapWebTypeComputeIamBindingCaiObject,
				mergeCreateUpdate: MergeIapWebTypeComputeIamBinding,
				mergeDelete:       MergeIapWebTypeComputeIamBindingDelete,
				fetch:             FetchIapWebTypeComputeIamPolicy,
			},
		},
		"google_iap_web_type_compute_iam_member": {
			{
				convert:           GetIapWebTypeComputeIamMemberCaiObject,
				mergeCreateUpdate: MergeIapWebTypeComputeIamMember,
				mergeDelete:       MergeIapWebTypeComputeIamMemberDelete,
				fetch:             FetchIapWebTypeComputeIamPolicy,
			},
		},
		"google_iap_web_type_app_engine_iam_policy": {
			{
				convert:           GetIapWebTypeAppEngineIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapWebTypeAppEngineIamPolicy,
			},
		},
		"google_iap_web_type_app_engine_iam_binding": {
			{
				convert:           GetIapWebTypeAppEngineIamBindingCaiObject,
				mergeCreateUpdate: MergeIapWebTypeAppEngineIamBinding,
				mergeDelete:       MergeIapWebTypeAppEngineIamBindingDelete,
				fetch:             FetchIapWebTypeAppEngineIamPolicy,
			},
		},
		"google_iap_web_type_app_engine_iam_member": {
			{
				convert:           GetIapWebTypeAppEngineIamMemberCaiObject,
				mergeCreateUpdate: MergeIapWebTypeAppEngineIamMember,
				mergeDelete:       MergeIapWebTypeAppEngineIamMemberDelete,
				fetch:             FetchIapWebTypeAppEngineIamPolicy,
			},
		},
		"google_iap_app_engine_version_iam_policy": {
			{
				convert:           GetIapAppEngineVersionIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapAppEngineVersionIamPolicy,
			},
		},
		"google_iap_app_engine_version_iam_binding": {
			{
				convert:           GetIapAppEngineVersionIamBindingCaiObject,
				mergeCreateUpdate: MergeIapAppEngineVersionIamBinding,
				mergeDelete:       MergeIapAppEngineVersionIamBindingDelete,
				fetch:             FetchIapAppEngineVersionIamPolicy,
			},
		},
		"google_iap_app_engine_version_iam_member": {
			{
				convert:           GetIapAppEngineVersionIamMemberCaiObject,
				mergeCreateUpdate: MergeIapAppEngineVersionIamMember,
				mergeDelete:       MergeIapAppEngineVersionIamMemberDelete,
				fetch:             FetchIapAppEngineVersionIamPolicy,
			},
		},
		"google_iap_app_engine_service_iam_policy": {
			{
				convert:           GetIapAppEngineServiceIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapAppEngineServiceIamPolicy,
			},
		},
		"google_iap_app_engine_service_iam_binding": {
			{
				convert:           GetIapAppEngineServiceIamBindingCaiObject,
				mergeCreateUpdate: MergeIapAppEngineServiceIamBinding,
				mergeDelete:       MergeIapAppEngineServiceIamBindingDelete,
				fetch:             FetchIapAppEngineServiceIamPolicy,
			},
		},
		"google_iap_app_engine_service_iam_member": {
			{
				convert:           GetIapAppEngineServiceIamMemberCaiObject,
				mergeCreateUpdate: MergeIapAppEngineServiceIamMember,
				mergeDelete:       MergeIapAppEngineServiceIamMemberDelete,
				fetch:             FetchIapAppEngineServiceIamPolicy,
			},
		},
		"google_iap_web_backend_service_iam_policy": {
			{
				convert:           GetIapWebBackendServiceIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapWebBackendServiceIamPolicy,
			},
		},
		"google_iap_web_backend_service_iam_binding": {
			{
				convert:           GetIapWebBackendServiceIamBindingCaiObject,
				mergeCreateUpdate: MergeIapWebBackendServiceIamBinding,
				mergeDelete:       MergeIapWebBackendServiceIamBindingDelete,
				fetch:             FetchIapWebBackendServiceIamPolicy,
			},
		},
		"google_iap_web_backend_service_iam_member": {
			{
				convert:           GetIapWebBackendServiceIamMemberCaiObject,
				mergeCreateUpdate: MergeIapWebBackendServiceIamMember,
				mergeDelete:       MergeIapWebBackendServiceIamMemberDelete,
				fetch:             FetchIapWebBackendServiceIamPolicy,
			},
		},
		"google_iap_tunnel_instance_iam_policy": {
			{
				convert:           GetIapTunnelInstanceIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapTunnelInstanceIamPolicy,
			},
		},
		"google_iap_tunnel_instance_iam_binding": {
			{
				convert:           GetIapTunnelInstanceIamBindingCaiObject,
				mergeCreateUpdate: MergeIapTunnelInstanceIamBinding,
				mergeDelete:       MergeIapTunnelInstanceIamBindingDelete,
				fetch:             FetchIapTunnelInstanceIamPolicy,
			},
		},
		"google_iap_tunnel_instance_iam_member": {
			{
				convert:           GetIapTunnelInstanceIamMemberCaiObject,
				mergeCreateUpdate: MergeIapTunnelInstanceIamMember,
				mergeDelete:       MergeIapTunnelInstanceIamMemberDelete,
				fetch:             FetchIapTunnelInstanceIamPolicy,
			},
		},
		"google_iap_tunnel_iam_policy": {
			{
				convert:           GetIapTunnelIamPolicyCaiObject,
				mergeCreateUpdate: MergeIapTunnelIamPolicy,
			},
		},
		"google_iap_tunnel_iam_binding": {
			{
				convert:           GetIapTunnelIamBindingCaiObject,
				mergeCreateUpdate: MergeIapTunnelIamBinding,
				mergeDelete:       MergeIapTunnelIamBindingDelete,
				fetch:             FetchIapTunnelIamPolicy,
			},
		},
		"google_iap_tunnel_iam_member": {
			{
				convert:           GetIapTunnelIamMemberCaiObject,
				mergeCreateUpdate: MergeIapTunnelIamMember,
				mergeDelete:       MergeIapTunnelIamMemberDelete,
				fetch:             FetchIapTunnelIamPolicy,
			},
		},
		"google_notebooks_instance_iam_policy": {
			{
				convert:           GetNotebooksInstanceIamPolicyCaiObject,
				mergeCreateUpdate: MergeNotebooksInstanceIamPolicy,
			},
		},
		"google_notebooks_instance_iam_binding": {
			{
				convert:           GetNotebooksInstanceIamBindingCaiObject,
				mergeCreateUpdate: MergeNotebooksInstanceIamBinding,
				mergeDelete:       MergeNotebooksInstanceIamBindingDelete,
				fetch:             FetchNotebooksInstanceIamPolicy,
			},
		},
		"google_notebooks_instance_iam_member": {
			{
				convert:           GetNotebooksInstanceIamMemberCaiObject,
				mergeCreateUpdate: MergeNotebooksInstanceIamMember,
				mergeDelete:       MergeNotebooksInstanceIamMemberDelete,
				fetch:             FetchNotebooksInstanceIamPolicy,
			},
		},
		"google_pubsub_topic_iam_policy": {
			{
				convert:           GetPubsubTopicIamPolicyCaiObject,
				mergeCreateUpdate: MergePubsubTopicIamPolicy,
			},
		},
		"google_pubsub_topic_iam_binding": {
			{
				convert:           GetPubsubTopicIamBindingCaiObject,
				mergeCreateUpdate: MergePubsubTopicIamBinding,
				mergeDelete:       MergePubsubTopicIamBindingDelete,
				fetch:             FetchPubsubTopicIamPolicy,
			},
		},
		"google_pubsub_topic_iam_member": {
			{
				convert:           GetPubsubTopicIamMemberCaiObject,
				mergeCreateUpdate: MergePubsubTopicIamMember,
				mergeDelete:       MergePubsubTopicIamMemberDelete,
				fetch:             FetchPubsubTopicIamPolicy,
			},
		},
		"google_runtimeconfig_config_iam_policy": {
			{
				convert:           GetRuntimeConfigConfigIamPolicyCaiObject,
				mergeCreateUpdate: MergeRuntimeConfigConfigIamPolicy,
			},
		},
		"google_runtimeconfig_config_iam_binding": {
			{
				convert:           GetRuntimeConfigConfigIamBindingCaiObject,
				mergeCreateUpdate: MergeRuntimeConfigConfigIamBinding,
				mergeDelete:       MergeRuntimeConfigConfigIamBindingDelete,
				fetch:             FetchRuntimeConfigConfigIamPolicy,
			},
		},
		"google_runtimeconfig_config_iam_member": {
			{
				convert:           GetRuntimeConfigConfigIamMemberCaiObject,
				mergeCreateUpdate: MergeRuntimeConfigConfigIamMember,
				mergeDelete:       MergeRuntimeConfigConfigIamMemberDelete,
				fetch:             FetchRuntimeConfigConfigIamPolicy,
			},
		},
		"google_secret_manager_secret_iam_policy": {
			{
				convert:           GetSecretManagerSecretIamPolicyCaiObject,
				mergeCreateUpdate: MergeSecretManagerSecretIamPolicy,
			},
		},
		"google_secret_manager_secret_iam_binding": {
			{
				convert:           GetSecretManagerSecretIamBindingCaiObject,
				mergeCreateUpdate: MergeSecretManagerSecretIamBinding,
				mergeDelete:       MergeSecretManagerSecretIamBindingDelete,
				fetch:             FetchSecretManagerSecretIamPolicy,
			},
		},
		"google_secret_manager_secret_iam_member": {
			{
				convert:           GetSecretManagerSecretIamMemberCaiObject,
				mergeCreateUpdate: MergeSecretManagerSecretIamMember,
				mergeDelete:       MergeSecretManagerSecretIamMemberDelete,
				fetch:             FetchSecretManagerSecretIamPolicy,
			},
		},
		"google_endpoints_service_iam_policy": {
			{
				convert:           GetServiceManagementServiceIamPolicyCaiObject,
				mergeCreateUpdate: MergeServiceManagementServiceIamPolicy,
			},
		},
		"google_endpoints_service_iam_binding": {
			{
				convert:           GetServiceManagementServiceIamBindingCaiObject,
				mergeCreateUpdate: MergeServiceManagementServiceIamBinding,
				mergeDelete:       MergeServiceManagementServiceIamBindingDelete,
				fetch:             FetchServiceManagementServiceIamPolicy,
			},
		},
		"google_endpoints_service_iam_member": {
			{
				convert:           GetServiceManagementServiceIamMemberCaiObject,
				mergeCreateUpdate: MergeServiceManagementServiceIamMember,
				mergeDelete:       MergeServiceManagementServiceIamMemberDelete,
				fetch:             FetchServiceManagementServiceIamPolicy,
			},
		},
		"google_sourcerepo_repository_iam_policy": {
			{
				convert:           GetSourceRepoRepositoryIamPolicyCaiObject,
				mergeCreateUpdate: MergeSourceRepoRepositoryIamPolicy,
			},
		},
		"google_sourcerepo_repository_iam_binding": {
			{
				convert:           GetSourceRepoRepositoryIamBindingCaiObject,
				mergeCreateUpdate: MergeSourceRepoRepositoryIamBinding,
				mergeDelete:       MergeSourceRepoRepositoryIamBindingDelete,
				fetch:             FetchSourceRepoRepositoryIamPolicy,
			},
		},
		"google_sourcerepo_repository_iam_member": {
			{
				convert:           GetSourceRepoRepositoryIamMemberCaiObject,
				mergeCreateUpdate: MergeSourceRepoRepositoryIamMember,
				mergeDelete:       MergeSourceRepoRepositoryIamMemberDelete,
				fetch:             FetchSourceRepoRepositoryIamPolicy,
			},
		},
		"google_storage_bucket_iam_policy": {
			{
				convert:           GetStorageBucketIamPolicyCaiObject,
				mergeCreateUpdate: MergeStorageBucketIamPolicy,
			},
		},
		"google_storage_bucket_iam_binding": {
			{
				convert:           GetStorageBucketIamBindingCaiObject,
				mergeCreateUpdate: MergeStorageBucketIamBinding,
				mergeDelete:       MergeStorageBucketIamBindingDelete,
				fetch:             FetchStorageBucketIamPolicy,
			},
		},
		"google_storage_bucket_iam_member": {
			{
				convert:           GetStorageBucketIamMemberCaiObject,
				mergeCreateUpdate: MergeStorageBucketIamMember,
				mergeDelete:       MergeStorageBucketIamMemberDelete,
				fetch:             FetchStorageBucketIamPolicy,
			},
		},
		"google_tags_tag_key_iam_policy": {
			{
				convert:           GetTagsTagKeyIamPolicyCaiObject,
				mergeCreateUpdate: MergeTagsTagKeyIamPolicy,
			},
		},
		"google_tags_tag_key_iam_binding": {
			{
				convert:           GetTagsTagKeyIamBindingCaiObject,
				mergeCreateUpdate: MergeTagsTagKeyIamBinding,
				mergeDelete:       MergeTagsTagKeyIamBindingDelete,
				fetch:             FetchTagsTagKeyIamPolicy,
			},
		},
		"google_tags_tag_key_iam_member": {
			{
				convert:           GetTagsTagKeyIamMemberCaiObject,
				mergeCreateUpdate: MergeTagsTagKeyIamMember,
				mergeDelete:       MergeTagsTagKeyIamMemberDelete,
				fetch:             FetchTagsTagKeyIamPolicy,
			},
		},
		"google_tags_tag_value_iam_policy": {
			{
				convert:           GetTagsTagValueIamPolicyCaiObject,
				mergeCreateUpdate: MergeTagsTagValueIamPolicy,
			},
		},
		"google_tags_tag_value_iam_binding": {
			{
				convert:           GetTagsTagValueIamBindingCaiObject,
				mergeCreateUpdate: MergeTagsTagValueIamBinding,
				mergeDelete:       MergeTagsTagValueIamBindingDelete,
				fetch:             FetchTagsTagValueIamPolicy,
			},
		},
		"google_tags_tag_value_iam_member": {
			{
				convert:           GetTagsTagValueIamMemberCaiObject,
				mergeCreateUpdate: MergeTagsTagValueIamMember,
				mergeDelete:       MergeTagsTagValueIamMemberDelete,
				fetch:             FetchTagsTagValueIamPolicy,
			},
		},
		"google_project": {
			{
				convert:           GetProjectCaiObject,
				mergeCreateUpdate: MergeProject,
			},
			{convert: GetProjectBillingInfoCaiObject},
		},
		"google_bigtable_instance": {
			{
				convert: GetBigtableInstanceCaiObject,
			},
			{
				convert: GetBigtableClusterCaiObject,
			},
		},
		"google_organization_iam_policy": {
			{
				convert:           GetOrganizationIamPolicyCaiObject,
				mergeCreateUpdate: MergeOrganizationIamPolicy,
			},
		},
		"google_project_organization_policy": {
			{
				convert:           GetProjectOrgPolicyCaiObject,
				mergeCreateUpdate: MergeProjectOrgPolicy,
			},
		},
		"google_organization_iam_binding": {
			{
				convert:           GetOrganizationIamBindingCaiObject,
				mergeCreateUpdate: MergeOrganizationIamBinding,
				mergeDelete:       MergeOrganizationIamBindingDelete,
				fetch:             FetchOrganizationIamPolicy,
			},
		},
		"google_organization_iam_member": {
			{
				convert:           GetOrganizationIamMemberCaiObject,
				mergeCreateUpdate: MergeOrganizationIamMember,
				mergeDelete:       MergeOrganizationIamMemberDelete,
				fetch:             FetchOrganizationIamPolicy,
			},
		},
		"google_folder_iam_policy": {
			{
				convert:           GetFolderIamPolicyCaiObject,
				mergeCreateUpdate: MergeFolderIamPolicy,
			},
		},
		"google_folder_iam_binding": {
			{
				convert:           GetFolderIamBindingCaiObject,
				mergeCreateUpdate: MergeFolderIamBinding,
				mergeDelete:       MergeFolderIamBindingDelete,
				fetch:             FetchFolderIamPolicy,
			},
		},
		"google_folder_iam_member": {
			{
				convert:           GetFolderIamMemberCaiObject,
				mergeCreateUpdate: MergeFolderIamMember,
				mergeDelete:       MergeFolderIamMemberDelete,
				fetch:             FetchFolderIamPolicy,
			},
		},
		"google_project_iam_policy": {
			{
				convert:           GetProjectIamPolicyCaiObject,
				mergeCreateUpdate: MergeProjectIamPolicy,
			},
		},
		"google_project_iam_binding": {
			{
				convert:           GetProjectIamBindingCaiObject,
				mergeCreateUpdate: MergeProjectIamBinding,
				mergeDelete:       MergeProjectIamBindingDelete,
				fetch:             FetchProjectIamPolicy,
			},
		},
		"google_project_iam_member": {
			{
				convert:           GetProjectIamMemberCaiObject,
				mergeCreateUpdate: MergeProjectIamMember,
				mergeDelete:       MergeProjectIamMemberDelete,
				fetch:             FetchProjectIamPolicy,
			},
		},
	}
}

func MergeProject(existing, incoming Asset) Asset {
	existing.Resource = incoming.Resource
	return existing
}

// SupportedResources returns a sorted list of terraform resource names.
func SupportedTerraformResources() []string {
	fns := Mappers()
	list := make([]string, 0, len(fns))
	for k := range fns {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}