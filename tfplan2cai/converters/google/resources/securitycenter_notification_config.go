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

const SecurityCenterNotificationConfigAssetType string = "securitycenter.googleapis.com/NotificationConfig"

func resourceConverterSecurityCenterNotificationConfig() ResourceConverter {
	return ResourceConverter{
		AssetType: SecurityCenterNotificationConfigAssetType,
		Convert:   GetSecurityCenterNotificationConfigCaiObject,
	}
}

func GetSecurityCenterNotificationConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//securitycenter.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetSecurityCenterNotificationConfigApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: SecurityCenterNotificationConfigAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "NotificationConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetSecurityCenterNotificationConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterNotificationConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	pubsubTopicProp, err := expandSecurityCenterNotificationConfigPubsubTopic(d.Get("pubsub_topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(pubsubTopicProp)) && (ok || !reflect.DeepEqual(v, pubsubTopicProp)) {
		obj["pubsubTopic"] = pubsubTopicProp
	}
	streamingConfigProp, err := expandSecurityCenterNotificationConfigStreamingConfig(d.Get("streaming_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("streaming_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(streamingConfigProp)) && (ok || !reflect.DeepEqual(v, streamingConfigProp)) {
		obj["streamingConfig"] = streamingConfigProp
	}

	return obj, nil
}

func expandSecurityCenterNotificationConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterNotificationConfigStreamingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFilter, err := expandSecurityCenterNotificationConfigStreamingConfigFilter(original["filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["filter"] = transformedFilter
	}

	return transformed, nil
}

func expandSecurityCenterNotificationConfigStreamingConfigFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
