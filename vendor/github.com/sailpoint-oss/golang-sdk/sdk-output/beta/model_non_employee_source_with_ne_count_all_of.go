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

// NonEmployeeSourceWithNECountAllOf struct for NonEmployeeSourceWithNECountAllOf
type NonEmployeeSourceWithNECountAllOf struct {
	// Number of non-employee records associated with this source.
	NonEmployeeCount *int32 `json:"nonEmployeeCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceWithNECountAllOf NonEmployeeSourceWithNECountAllOf

// NewNonEmployeeSourceWithNECountAllOf instantiates a new NonEmployeeSourceWithNECountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceWithNECountAllOf() *NonEmployeeSourceWithNECountAllOf {
	this := NonEmployeeSourceWithNECountAllOf{}
	return &this
}

// NewNonEmployeeSourceWithNECountAllOfWithDefaults instantiates a new NonEmployeeSourceWithNECountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceWithNECountAllOfWithDefaults() *NonEmployeeSourceWithNECountAllOf {
	this := NonEmployeeSourceWithNECountAllOf{}
	return &this
}

// GetNonEmployeeCount returns the NonEmployeeCount field value if set, zero value otherwise.
func (o *NonEmployeeSourceWithNECountAllOf) GetNonEmployeeCount() int32 {
	if o == nil || isNil(o.NonEmployeeCount) {
		var ret int32
		return ret
	}
	return *o.NonEmployeeCount
}

// GetNonEmployeeCountOk returns a tuple with the NonEmployeeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceWithNECountAllOf) GetNonEmployeeCountOk() (*int32, bool) {
	if o == nil || isNil(o.NonEmployeeCount) {
		return nil, false
	}
	return o.NonEmployeeCount, true
}

// HasNonEmployeeCount returns a boolean if a field has been set.
func (o *NonEmployeeSourceWithNECountAllOf) HasNonEmployeeCount() bool {
	if o != nil && !isNil(o.NonEmployeeCount) {
		return true
	}

	return false
}

// SetNonEmployeeCount gets a reference to the given int32 and assigns it to the NonEmployeeCount field.
func (o *NonEmployeeSourceWithNECountAllOf) SetNonEmployeeCount(v int32) {
	o.NonEmployeeCount = &v
}

func (o NonEmployeeSourceWithNECountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NonEmployeeCount) {
		toSerialize["nonEmployeeCount"] = o.NonEmployeeCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeSourceWithNECountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeSourceWithNECountAllOf := _NonEmployeeSourceWithNECountAllOf{}

	if err = json.Unmarshal(bytes, &varNonEmployeeSourceWithNECountAllOf); err == nil {
		*o = NonEmployeeSourceWithNECountAllOf(varNonEmployeeSourceWithNECountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nonEmployeeCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceWithNECountAllOf struct {
	value *NonEmployeeSourceWithNECountAllOf
	isSet bool
}

func (v NullableNonEmployeeSourceWithNECountAllOf) Get() *NonEmployeeSourceWithNECountAllOf {
	return v.value
}

func (v *NullableNonEmployeeSourceWithNECountAllOf) Set(val *NonEmployeeSourceWithNECountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceWithNECountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceWithNECountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceWithNECountAllOf(val *NonEmployeeSourceWithNECountAllOf) *NullableNonEmployeeSourceWithNECountAllOf {
	return &NullableNonEmployeeSourceWithNECountAllOf{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceWithNECountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceWithNECountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


