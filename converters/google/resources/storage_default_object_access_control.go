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

import "reflect"

const StorageDefaultObjectAccessControlAssetType string = "storage.googleapis.com/DefaultObjectAccessControl"

func resourceConverterStorageDefaultObjectAccessControl() ResourceConverter {
	return ResourceConverter{
		AssetType: StorageDefaultObjectAccessControlAssetType,
		Convert:   GetStorageDefaultObjectAccessControlCaiObject,
	}
}

func GetStorageDefaultObjectAccessControlCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//storage.googleapis.com/b/{{bucket}}/defaultObjectAcl/{{entity}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetStorageDefaultObjectAccessControlApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: StorageDefaultObjectAccessControlAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "DefaultObjectAccessControl",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetStorageDefaultObjectAccessControlApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketProp, err := expandStorageDefaultObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageDefaultObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageDefaultObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(objectProp)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageDefaultObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandStorageDefaultObjectAccessControlBucket(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlEntity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlObject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlRole(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
