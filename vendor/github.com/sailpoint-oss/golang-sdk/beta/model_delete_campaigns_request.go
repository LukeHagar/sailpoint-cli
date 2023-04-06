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

// DeleteCampaignsRequest struct for DeleteCampaignsRequest
type DeleteCampaignsRequest struct {
	// The ids of the campaigns to delete
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteCampaignsRequest DeleteCampaignsRequest

// NewDeleteCampaignsRequest instantiates a new DeleteCampaignsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCampaignsRequest() *DeleteCampaignsRequest {
	this := DeleteCampaignsRequest{}
	return &this
}

// NewDeleteCampaignsRequestWithDefaults instantiates a new DeleteCampaignsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCampaignsRequestWithDefaults() *DeleteCampaignsRequest {
	this := DeleteCampaignsRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteCampaignsRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCampaignsRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteCampaignsRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteCampaignsRequest) SetIds(v []string) {
	o.Ids = v
}

func (o DeleteCampaignsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteCampaignsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteCampaignsRequest := _DeleteCampaignsRequest{}

	if err = json.Unmarshal(bytes, &varDeleteCampaignsRequest); err == nil {
		*o = DeleteCampaignsRequest(varDeleteCampaignsRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteCampaignsRequest struct {
	value *DeleteCampaignsRequest
	isSet bool
}

func (v NullableDeleteCampaignsRequest) Get() *DeleteCampaignsRequest {
	return v.value
}

func (v *NullableDeleteCampaignsRequest) Set(val *DeleteCampaignsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCampaignsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCampaignsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCampaignsRequest(val *DeleteCampaignsRequest) *NullableDeleteCampaignsRequest {
	return &NullableDeleteCampaignsRequest{value: val, isSet: true}
}

func (v NullableDeleteCampaignsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCampaignsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

