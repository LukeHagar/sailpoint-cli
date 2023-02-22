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

// ErrorResponseDto struct for ErrorResponseDto
type ErrorResponseDto struct {
	// Fine-grained error code providing more detail of the error.
	DetailCode *string `json:"detailCode,omitempty"`
	// Unique tracking id for the error.
	TrackingId *string `json:"trackingId,omitempty"`
	// Generic localized reason for error
	Messages []ErrorMessageDto `json:"messages,omitempty"`
	// Plain-text descriptive reasons to provide additional detail to the text provided in the messages field
	Causes []ErrorMessageDto `json:"causes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorResponseDto ErrorResponseDto

// NewErrorResponseDto instantiates a new ErrorResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseDto() *ErrorResponseDto {
	this := ErrorResponseDto{}
	return &this
}

// NewErrorResponseDtoWithDefaults instantiates a new ErrorResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseDtoWithDefaults() *ErrorResponseDto {
	this := ErrorResponseDto{}
	return &this
}

// GetDetailCode returns the DetailCode field value if set, zero value otherwise.
func (o *ErrorResponseDto) GetDetailCode() string {
	if o == nil || isNil(o.DetailCode) {
		var ret string
		return ret
	}
	return *o.DetailCode
}

// GetDetailCodeOk returns a tuple with the DetailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDto) GetDetailCodeOk() (*string, bool) {
	if o == nil || isNil(o.DetailCode) {
		return nil, false
	}
	return o.DetailCode, true
}

// HasDetailCode returns a boolean if a field has been set.
func (o *ErrorResponseDto) HasDetailCode() bool {
	if o != nil && !isNil(o.DetailCode) {
		return true
	}

	return false
}

// SetDetailCode gets a reference to the given string and assigns it to the DetailCode field.
func (o *ErrorResponseDto) SetDetailCode(v string) {
	o.DetailCode = &v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *ErrorResponseDto) GetTrackingId() string {
	if o == nil || isNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDto) GetTrackingIdOk() (*string, bool) {
	if o == nil || isNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *ErrorResponseDto) HasTrackingId() bool {
	if o != nil && !isNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *ErrorResponseDto) SetTrackingId(v string) {
	o.TrackingId = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ErrorResponseDto) GetMessages() []ErrorMessageDto {
	if o == nil || isNil(o.Messages) {
		var ret []ErrorMessageDto
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDto) GetMessagesOk() ([]ErrorMessageDto, bool) {
	if o == nil || isNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ErrorResponseDto) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ErrorMessageDto and assigns it to the Messages field.
func (o *ErrorResponseDto) SetMessages(v []ErrorMessageDto) {
	o.Messages = v
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *ErrorResponseDto) GetCauses() []ErrorMessageDto {
	if o == nil || isNil(o.Causes) {
		var ret []ErrorMessageDto
		return ret
	}
	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseDto) GetCausesOk() ([]ErrorMessageDto, bool) {
	if o == nil || isNil(o.Causes) {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *ErrorResponseDto) HasCauses() bool {
	if o != nil && !isNil(o.Causes) {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []ErrorMessageDto and assigns it to the Causes field.
func (o *ErrorResponseDto) SetCauses(v []ErrorMessageDto) {
	o.Causes = v
}

func (o ErrorResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DetailCode) {
		toSerialize["detailCode"] = o.DetailCode
	}
	if !isNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.Causes) {
		toSerialize["causes"] = o.Causes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorResponseDto) UnmarshalJSON(bytes []byte) (err error) {
	varErrorResponseDto := _ErrorResponseDto{}

	if err = json.Unmarshal(bytes, &varErrorResponseDto); err == nil {
		*o = ErrorResponseDto(varErrorResponseDto)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "detailCode")
		delete(additionalProperties, "trackingId")
		delete(additionalProperties, "messages")
		delete(additionalProperties, "causes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorResponseDto struct {
	value *ErrorResponseDto
	isSet bool
}

func (v NullableErrorResponseDto) Get() *ErrorResponseDto {
	return v.value
}

func (v *NullableErrorResponseDto) Set(val *ErrorResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseDto(val *ErrorResponseDto) *NullableErrorResponseDto {
	return &NullableErrorResponseDto{value: val, isSet: true}
}

func (v NullableErrorResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


