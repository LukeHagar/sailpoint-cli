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

// CompletionStatus The status after completion.
type CompletionStatus string

// List of CompletionStatus
const (
	COMPLETIONSTATUS_SUCCESS CompletionStatus = "SUCCESS"
	COMPLETIONSTATUS_FAILURE CompletionStatus = "FAILURE"
	COMPLETIONSTATUS_INCOMPLETE CompletionStatus = "INCOMPLETE"
	COMPLETIONSTATUS_PENDING CompletionStatus = "PENDING"
)

// All allowed values of CompletionStatus enum
var AllowedCompletionStatusEnumValues = []CompletionStatus{
	"SUCCESS",
	"FAILURE",
	"INCOMPLETE",
	"PENDING",
}

func (v *CompletionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompletionStatus(value)
	for _, existing := range AllowedCompletionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompletionStatus", value)
}

// NewCompletionStatusFromValue returns a pointer to a valid CompletionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompletionStatusFromValue(v string) (*CompletionStatus, error) {
	ev := CompletionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompletionStatus: valid values are %v", v, AllowedCompletionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompletionStatus) IsValid() bool {
	for _, existing := range AllowedCompletionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompletionStatus value
func (v CompletionStatus) Ptr() *CompletionStatus {
	return &v
}

type NullableCompletionStatus struct {
	value *CompletionStatus
	isSet bool
}

func (v NullableCompletionStatus) Get() *CompletionStatus {
	return v.value
}

func (v *NullableCompletionStatus) Set(val *CompletionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionStatus(val *CompletionStatus) *NullableCompletionStatus {
	return &NullableCompletionStatus{value: val, isSet: true}
}

func (v NullableCompletionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

