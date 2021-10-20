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
	"reflect"
)

const ComputeGlobalAddressAssetType string = "compute.googleapis.com/GlobalAddress"

func resourceConverterComputeGlobalAddress() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeGlobalAddressAssetType,
		Convert:   GetComputeGlobalAddressCaiObject,
	}
}

func GetComputeGlobalAddressCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/addresses/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeGlobalAddressApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeGlobalAddressAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "GlobalAddress",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeGlobalAddressApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	addressProp, err := expandComputeGlobalAddressAddress(d.Get("address"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("address"); !isEmptyValue(reflect.ValueOf(addressProp)) && (ok || !reflect.DeepEqual(v, addressProp)) {
		obj["address"] = addressProp
	}
	descriptionProp, err := expandComputeGlobalAddressDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeGlobalAddressName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	ipVersionProp, err := expandComputeGlobalAddressIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	prefixLengthProp, err := expandComputeGlobalAddressPrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("prefix_length"); !isEmptyValue(reflect.ValueOf(prefixLengthProp)) && (ok || !reflect.DeepEqual(v, prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	addressTypeProp, err := expandComputeGlobalAddressAddressType(d.Get("address_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("address_type"); !isEmptyValue(reflect.ValueOf(addressTypeProp)) && (ok || !reflect.DeepEqual(v, addressTypeProp)) {
		obj["addressType"] = addressTypeProp
	}
	purposeProp, err := expandComputeGlobalAddressPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !isEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	networkProp, err := expandComputeGlobalAddressNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}

	return obj, nil
}

func expandComputeGlobalAddressAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressIpVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPrefixLength(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressAddressType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressPurpose(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalAddressNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}
