// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DNSManagedZoneIAMAssetType string = "dns.googleapis.com/ManagedZone"

func resourceConverterDNSManagedZoneIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamPolicyCaiObject,
		MergeCreateUpdate: MergeDNSManagedZoneIamPolicy,
	}
}

func resourceConverterDNSManagedZoneIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamBindingCaiObject,
		FetchFullResource: FetchDNSManagedZoneIamPolicy,
		MergeCreateUpdate: MergeDNSManagedZoneIamBinding,
		MergeDelete:       MergeDNSManagedZoneIamBindingDelete,
	}
}

func resourceConverterDNSManagedZoneIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         DNSManagedZoneIAMAssetType,
		Convert:           GetDNSManagedZoneIamMemberCaiObject,
		FetchFullResource: FetchDNSManagedZoneIamPolicy,
		MergeCreateUpdate: MergeDNSManagedZoneIamMember,
		MergeDelete:       MergeDNSManagedZoneIamMemberDelete,
	}
}

func GetDNSManagedZoneIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, expandIamPolicyBindings)
}

func GetDNSManagedZoneIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, expandIamRoleBindings)
}

func GetDNSManagedZoneIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDNSManagedZoneIamAsset(d, config, expandIamMemberBindings)
}

func MergeDNSManagedZoneIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDNSManagedZoneIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeDNSManagedZoneIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeDNSManagedZoneIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeDNSManagedZoneIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newDNSManagedZoneIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//dns.googleapis.com/projects/{{project}}/managedZones/{{managed_zone}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: DNSManagedZoneIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDNSManagedZoneIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("managed_zone"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		DNSManagedZoneIamUpdaterProducer,
		d,
		config,
		"//dns.googleapis.com/projects/{{project}}/managedZones/{{managed_zone}}",
		DNSManagedZoneIAMAssetType,
	)
}
