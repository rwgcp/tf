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
const ComputeDiskIAMAssetType string = "compute.googleapis.com/Disk"

func resourceConverterComputeDiskIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeDiskIamPolicy,
	}
}

func resourceConverterComputeDiskIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamBindingCaiObject,
		FetchFullResource: FetchComputeDiskIamPolicy,
		MergeCreateUpdate: MergeComputeDiskIamBinding,
		MergeDelete:       MergeComputeDiskIamBindingDelete,
	}
}

func resourceConverterComputeDiskIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeDiskIAMAssetType,
		Convert:           GetComputeDiskIamMemberCaiObject,
		FetchFullResource: FetchComputeDiskIamPolicy,
		MergeCreateUpdate: MergeComputeDiskIamMember,
		MergeDelete:       MergeComputeDiskIamMemberDelete,
	}
}

func GetComputeDiskIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeDiskIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeDiskIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newComputeDiskIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeDiskIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeDiskIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeDiskIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeDiskIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeDiskIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeDiskIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{name}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ComputeDiskIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeDiskIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("zone"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeDiskIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/disks/{{name}}",
		ComputeDiskIAMAssetType,
	)
}
