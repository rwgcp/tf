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

const FirestoreDatabaseAssetType string = "firestore.googleapis.com/Database"

func resourceConverterFirestoreDatabase() ResourceConverter {
	return ResourceConverter{
		AssetType: FirestoreDatabaseAssetType,
		Convert:   GetFirestoreDatabaseCaiObject,
	}
}

func GetFirestoreDatabaseCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//firestore.googleapis.com/projects/{{project}}/databases/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetFirestoreDatabaseApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: FirestoreDatabaseAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "Database",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetFirestoreDatabaseApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandFirestoreDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationIdProp, err := expandFirestoreDatabaseLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyModeProp)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineIntegrationModeProp)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	return obj, nil
}

func expandFirestoreDatabaseName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
}

func expandFirestoreDatabaseLocationId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseConcurrencyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
