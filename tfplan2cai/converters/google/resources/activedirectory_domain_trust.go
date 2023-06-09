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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const ActiveDirectoryDomainTrustAssetType string = "managedidentities.googleapis.com/DomainTrust"

func resourceConverterActiveDirectoryDomainTrust() ResourceConverter {
	return ResourceConverter{
		AssetType: ActiveDirectoryDomainTrustAssetType,
		Convert:   GetActiveDirectoryDomainTrustCaiObject,
	}
}

func GetActiveDirectoryDomainTrustCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//managedidentities.googleapis.com/projects/{{project}}/locations/global/domains/{{domain}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetActiveDirectoryDomainTrustApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ActiveDirectoryDomainTrustAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/managedidentities/v1/rest",
				DiscoveryName:        "DomainTrust",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetActiveDirectoryDomainTrustApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	targetDomainNameProp, err := expandActiveDirectoryDomainTrustTargetDomainName(d.Get("target_domain_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_domain_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetDomainNameProp)) && (ok || !reflect.DeepEqual(v, targetDomainNameProp)) {
		obj["targetDomainName"] = targetDomainNameProp
	}
	trustTypeProp, err := expandActiveDirectoryDomainTrustTrustType(d.Get("trust_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trust_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(trustTypeProp)) && (ok || !reflect.DeepEqual(v, trustTypeProp)) {
		obj["trustType"] = trustTypeProp
	}
	trustDirectionProp, err := expandActiveDirectoryDomainTrustTrustDirection(d.Get("trust_direction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trust_direction"); !tpgresource.IsEmptyValue(reflect.ValueOf(trustDirectionProp)) && (ok || !reflect.DeepEqual(v, trustDirectionProp)) {
		obj["trustDirection"] = trustDirectionProp
	}
	selectiveAuthenticationProp, err := expandActiveDirectoryDomainTrustSelectiveAuthentication(d.Get("selective_authentication"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("selective_authentication"); !tpgresource.IsEmptyValue(reflect.ValueOf(selectiveAuthenticationProp)) && (ok || !reflect.DeepEqual(v, selectiveAuthenticationProp)) {
		obj["selectiveAuthentication"] = selectiveAuthenticationProp
	}
	targetDnsIpAddressesProp, err := expandActiveDirectoryDomainTrustTargetDnsIpAddresses(d.Get("target_dns_ip_addresses"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_dns_ip_addresses"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetDnsIpAddressesProp)) && (ok || !reflect.DeepEqual(v, targetDnsIpAddressesProp)) {
		obj["targetDnsIpAddresses"] = targetDnsIpAddressesProp
	}
	trustHandshakeSecretProp, err := expandActiveDirectoryDomainTrustTrustHandshakeSecret(d.Get("trust_handshake_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trust_handshake_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(trustHandshakeSecretProp)) && (ok || !reflect.DeepEqual(v, trustHandshakeSecretProp)) {
		obj["trustHandshakeSecret"] = trustHandshakeSecretProp
	}

	return resourceActiveDirectoryDomainTrustEncoder(d, config, obj)
}

func resourceActiveDirectoryDomainTrustEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	wrappedReq := map[string]interface{}{
		"trust": obj,
	}
	return wrappedReq, nil
}

func expandActiveDirectoryDomainTrustTargetDomainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainTrustTrustType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainTrustTrustDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainTrustSelectiveAuthentication(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandActiveDirectoryDomainTrustTargetDnsIpAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandActiveDirectoryDomainTrustTrustHandshakeSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
