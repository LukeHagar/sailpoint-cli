/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
)

// ConnectorRuleValidationResponse ConnectorRuleValidationResponse
type ConnectorRuleValidationResponse struct {
	State string `json:"state"`
	Details []ConnectorRuleValidationResponseDetailsInner `json:"details"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorRuleValidationResponse ConnectorRuleValidationResponse

// NewConnectorRuleValidationResponse instantiates a new ConnectorRuleValidationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRuleValidationResponse(state string, details []ConnectorRuleValidationResponseDetailsInner) *ConnectorRuleValidationResponse {
	this := ConnectorRuleValidationResponse{}
	this.State = state
	this.Details = details
	return &this
}

// NewConnectorRuleValidationResponseWithDefaults instantiates a new ConnectorRuleValidationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRuleValidationResponseWithDefaults() *ConnectorRuleValidationResponse {
	this := ConnectorRuleValidationResponse{}
	return &this
}

// GetState returns the State field value
func (o *ConnectorRuleValidationResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleValidationResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ConnectorRuleValidationResponse) SetState(v string) {
	o.State = v
}

// GetDetails returns the Details field value
func (o *ConnectorRuleValidationResponse) GetDetails() []ConnectorRuleValidationResponseDetailsInner {
	if o == nil {
		var ret []ConnectorRuleValidationResponseDetailsInner
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *ConnectorRuleValidationResponse) GetDetailsOk() ([]ConnectorRuleValidationResponseDetailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *ConnectorRuleValidationResponse) SetDetails(v []ConnectorRuleValidationResponseDetailsInner) {
	o.Details = v
}

func (o ConnectorRuleValidationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorRuleValidationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorRuleValidationResponse := _ConnectorRuleValidationResponse{}

	if err = json.Unmarshal(bytes, &varConnectorRuleValidationResponse); err == nil {
		*o = ConnectorRuleValidationResponse(varConnectorRuleValidationResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorRuleValidationResponse struct {
	value *ConnectorRuleValidationResponse
	isSet bool
}

func (v NullableConnectorRuleValidationResponse) Get() *ConnectorRuleValidationResponse {
	return v.value
}

func (v *NullableConnectorRuleValidationResponse) Set(val *ConnectorRuleValidationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRuleValidationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRuleValidationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRuleValidationResponse(val *ConnectorRuleValidationResponse) *NullableConnectorRuleValidationResponse {
	return &NullableConnectorRuleValidationResponse{value: val, isSet: true}
}

func (v NullableConnectorRuleValidationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRuleValidationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

