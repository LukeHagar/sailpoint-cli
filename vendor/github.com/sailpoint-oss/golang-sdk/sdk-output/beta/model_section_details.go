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

// SectionDetails struct for SectionDetails
type SectionDetails struct {
	// Name of the FormItem
	Name *string `json:"name,omitempty"`
	// Label of the section
	Label *string `json:"label,omitempty"`
	// List of FormItems. FormItems can be SectionDetails and/or FieldDetails
	FormItems []map[string]interface{} `json:"formItems,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SectionDetails SectionDetails

// NewSectionDetails instantiates a new SectionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionDetails() *SectionDetails {
	this := SectionDetails{}
	return &this
}

// NewSectionDetailsWithDefaults instantiates a new SectionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionDetailsWithDefaults() *SectionDetails {
	this := SectionDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SectionDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SectionDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SectionDetails) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SectionDetails) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionDetails) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SectionDetails) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SectionDetails) SetLabel(v string) {
	o.Label = &v
}

// GetFormItems returns the FormItems field value if set, zero value otherwise.
func (o *SectionDetails) GetFormItems() []map[string]interface{} {
	if o == nil || isNil(o.FormItems) {
		var ret []map[string]interface{}
		return ret
	}
	return o.FormItems
}

// GetFormItemsOk returns a tuple with the FormItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionDetails) GetFormItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.FormItems) {
		return nil, false
	}
	return o.FormItems, true
}

// HasFormItems returns a boolean if a field has been set.
func (o *SectionDetails) HasFormItems() bool {
	if o != nil && !isNil(o.FormItems) {
		return true
	}

	return false
}

// SetFormItems gets a reference to the given []map[string]interface{} and assigns it to the FormItems field.
func (o *SectionDetails) SetFormItems(v []map[string]interface{}) {
	o.FormItems = v
}

func (o SectionDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.FormItems) {
		toSerialize["formItems"] = o.FormItems
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SectionDetails) UnmarshalJSON(bytes []byte) (err error) {
	varSectionDetails := _SectionDetails{}

	if err = json.Unmarshal(bytes, &varSectionDetails); err == nil {
		*o = SectionDetails(varSectionDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "formItems")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSectionDetails struct {
	value *SectionDetails
	isSet bool
}

func (v NullableSectionDetails) Get() *SectionDetails {
	return v.value
}

func (v *NullableSectionDetails) Set(val *SectionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionDetails(val *SectionDetails) *NullableSectionDetails {
	return &NullableSectionDetails{value: val, isSet: true}
}

func (v NullableSectionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

