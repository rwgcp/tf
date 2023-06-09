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
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const KMSKeyRingImportJobAssetType string = "cloudkms.googleapis.com/KeyRingImportJob"

func resourceConverterKMSKeyRingImportJob() ResourceConverter {
	return ResourceConverter{
		AssetType: KMSKeyRingImportJobAssetType,
		Convert:   GetKMSKeyRingImportJobCaiObject,
	}
}

func GetKMSKeyRingImportJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudkms.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetKMSKeyRingImportJobApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: KMSKeyRingImportJobAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "KeyRingImportJob",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetKMSKeyRingImportJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	importMethodProp, err := expandKMSKeyRingImportJobImportMethod(d.Get("import_method"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_method"); !tpgresource.IsEmptyValue(reflect.ValueOf(importMethodProp)) && (ok || !reflect.DeepEqual(v, importMethodProp)) {
		obj["importMethod"] = importMethodProp
	}
	protectionLevelProp, err := expandKMSKeyRingImportJobProtectionLevel(d.Get("protection_level"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protection_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(protectionLevelProp)) && (ok || !reflect.DeepEqual(v, protectionLevelProp)) {
		obj["protectionLevel"] = protectionLevelProp
	}

	return obj, nil
}

func expandKMSKeyRingImportJobImportMethod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyRingImportJobProtectionLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
