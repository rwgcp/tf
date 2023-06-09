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

const OSLoginSSHPublicKeyAssetType string = "oslogin.googleapis.com/SSHPublicKey"

func resourceConverterOSLoginSSHPublicKey() ResourceConverter {
	return ResourceConverter{
		AssetType: OSLoginSSHPublicKeyAssetType,
		Convert:   GetOSLoginSSHPublicKeyCaiObject,
	}
}

func GetOSLoginSSHPublicKeyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//oslogin.googleapis.com/users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetOSLoginSSHPublicKeyApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: OSLoginSSHPublicKeyAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/oslogin/v1/rest",
				DiscoveryName:        "SSHPublicKey",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetOSLoginSSHPublicKeyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	keyProp, err := expandOSLoginSSHPublicKeyKey(d.Get("key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyProp)) && (ok || !reflect.DeepEqual(v, keyProp)) {
		obj["key"] = keyProp
	}
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !tpgresource.IsEmptyValue(reflect.ValueOf(expirationTimeUsecProp)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	return obj, nil
}

func expandOSLoginSSHPublicKeyKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
