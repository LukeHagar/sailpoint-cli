/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// ProvisioningConfigPlanInitializerScript This is a reference to a plan initializer script.
type ProvisioningConfigPlanInitializerScript struct {
	// This is a Rule that allows provisioning instruction changes.
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConfigPlanInitializerScript ProvisioningConfigPlanInitializerScript

// NewProvisioningConfigPlanInitializerScript instantiates a new ProvisioningConfigPlanInitializerScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConfigPlanInitializerScript() *ProvisioningConfigPlanInitializerScript {
	this := ProvisioningConfigPlanInitializerScript{}
	return &this
}

// NewProvisioningConfigPlanInitializerScriptWithDefaults instantiates a new ProvisioningConfigPlanInitializerScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConfigPlanInitializerScriptWithDefaults() *ProvisioningConfigPlanInitializerScript {
	this := ProvisioningConfigPlanInitializerScript{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProvisioningConfigPlanInitializerScript) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConfigPlanInitializerScript) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProvisioningConfigPlanInitializerScript) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProvisioningConfigPlanInitializerScript) SetSource(v string) {
	o.Source = &v
}

func (o ProvisioningConfigPlanInitializerScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConfigPlanInitializerScript) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConfigPlanInitializerScript := _ProvisioningConfigPlanInitializerScript{}

	if err = json.Unmarshal(bytes, &varProvisioningConfigPlanInitializerScript); err == nil {
		*o = ProvisioningConfigPlanInitializerScript(varProvisioningConfigPlanInitializerScript)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConfigPlanInitializerScript struct {
	value *ProvisioningConfigPlanInitializerScript
	isSet bool
}

func (v NullableProvisioningConfigPlanInitializerScript) Get() *ProvisioningConfigPlanInitializerScript {
	return v.value
}

func (v *NullableProvisioningConfigPlanInitializerScript) Set(val *ProvisioningConfigPlanInitializerScript) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConfigPlanInitializerScript) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConfigPlanInitializerScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConfigPlanInitializerScript(val *ProvisioningConfigPlanInitializerScript) *NullableProvisioningConfigPlanInitializerScript {
	return &NullableProvisioningConfigPlanInitializerScript{value: val, isSet: true}
}

func (v NullableProvisioningConfigPlanInitializerScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConfigPlanInitializerScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


