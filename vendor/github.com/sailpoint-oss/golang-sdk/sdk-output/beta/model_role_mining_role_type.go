/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package beta

import (
	"encoding/json"
	"fmt"
)

// RoleMiningRoleType Role type
type RoleMiningRoleType string

// List of RoleMiningRoleType
const (
	ROLEMININGROLETYPE_SPECIALIZED RoleMiningRoleType = "SPECIALIZED"
	ROLEMININGROLETYPE_COMMON RoleMiningRoleType = "COMMON"
)

// All allowed values of RoleMiningRoleType enum
var AllowedRoleMiningRoleTypeEnumValues = []RoleMiningRoleType{
	"SPECIALIZED",
	"COMMON",
}

func (v *RoleMiningRoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleMiningRoleType(value)
	for _, existing := range AllowedRoleMiningRoleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleMiningRoleType", value)
}

// NewRoleMiningRoleTypeFromValue returns a pointer to a valid RoleMiningRoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleMiningRoleTypeFromValue(v string) (*RoleMiningRoleType, error) {
	ev := RoleMiningRoleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleMiningRoleType: valid values are %v", v, AllowedRoleMiningRoleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleMiningRoleType) IsValid() bool {
	for _, existing := range AllowedRoleMiningRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleMiningRoleType value
func (v RoleMiningRoleType) Ptr() *RoleMiningRoleType {
	return &v
}

type NullableRoleMiningRoleType struct {
	value *RoleMiningRoleType
	isSet bool
}

func (v NullableRoleMiningRoleType) Get() *RoleMiningRoleType {
	return v.value
}

func (v *NullableRoleMiningRoleType) Set(val *RoleMiningRoleType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMiningRoleType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMiningRoleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMiningRoleType(val *RoleMiningRoleType) *NullableRoleMiningRoleType {
	return &NullableRoleMiningRoleType{value: val, isSet: true}
}

func (v NullableRoleMiningRoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMiningRoleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

