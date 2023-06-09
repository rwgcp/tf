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
const PubsubTopicIAMAssetType string = "pubsub.googleapis.com/Topic"

func resourceConverterPubsubTopicIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamPolicyCaiObject,
		MergeCreateUpdate: MergePubsubTopicIamPolicy,
	}
}

func resourceConverterPubsubTopicIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamBindingCaiObject,
		FetchFullResource: FetchPubsubTopicIamPolicy,
		MergeCreateUpdate: MergePubsubTopicIamBinding,
		MergeDelete:       MergePubsubTopicIamBindingDelete,
	}
}

func resourceConverterPubsubTopicIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         PubsubTopicIAMAssetType,
		Convert:           GetPubsubTopicIamMemberCaiObject,
		FetchFullResource: FetchPubsubTopicIamPolicy,
		MergeCreateUpdate: MergePubsubTopicIamMember,
		MergeDelete:       MergePubsubTopicIamMemberDelete,
	}
}

func GetPubsubTopicIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPubsubTopicIamAsset(d, config, expandIamPolicyBindings)
}

func GetPubsubTopicIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPubsubTopicIamAsset(d, config, expandIamRoleBindings)
}

func GetPubsubTopicIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPubsubTopicIamAsset(d, config, expandIamMemberBindings)
}

func MergePubsubTopicIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergePubsubTopicIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergePubsubTopicIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergePubsubTopicIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergePubsubTopicIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newPubsubTopicIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/topics/{{topic}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: PubsubTopicIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchPubsubTopicIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("topic"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		PubsubTopicIamUpdaterProducer,
		d,
		config,
		"//pubsub.googleapis.com/projects/{{project}}/topics/{{topic}}",
		PubsubTopicIAMAssetType,
	)
}
