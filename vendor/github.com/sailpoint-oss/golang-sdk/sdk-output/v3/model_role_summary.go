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

// RoleSummary Role
type RoleSummary struct {
	// The unique ID of the referenced object.
	Id *string `json:"id,omitempty"`
	// The human readable name of the referenced object.
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Type *DtoType `json:"type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Owner *DisplayReference `json:"owner,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Revocable *bool `json:"revocable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleSummary RoleSummary

// NewRoleSummary instantiates a new RoleSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleSummary() *RoleSummary {
	this := RoleSummary{}
	return &this
}

// NewRoleSummaryWithDefaults instantiates a new RoleSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleSummaryWithDefaults() *RoleSummary {
	this := RoleSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleSummary) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RoleSummary) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RoleSummary) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RoleSummary) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleSummary) GetType() DtoType {
	if o == nil || isNil(o.Type) {
		var ret DtoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetTypeOk() (*DtoType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleSummary) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DtoType and assigns it to the Type field.
func (o *RoleSummary) SetType(v DtoType) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleSummary) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleSummary) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RoleSummary) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RoleSummary) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RoleSummary) UnsetDescription() {
	o.Description.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *RoleSummary) GetOwner() DisplayReference {
	if o == nil || isNil(o.Owner) {
		var ret DisplayReference
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetOwnerOk() (*DisplayReference, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *RoleSummary) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given DisplayReference and assigns it to the Owner field.
func (o *RoleSummary) SetOwner(v DisplayReference) {
	o.Owner = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *RoleSummary) GetDisabled() bool {
	if o == nil || isNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *RoleSummary) HasDisabled() bool {
	if o != nil && !isNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *RoleSummary) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetRevocable returns the Revocable field value if set, zero value otherwise.
func (o *RoleSummary) GetRevocable() bool {
	if o == nil || isNil(o.Revocable) {
		var ret bool
		return ret
	}
	return *o.Revocable
}

// GetRevocableOk returns a tuple with the Revocable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSummary) GetRevocableOk() (*bool, bool) {
	if o == nil || isNil(o.Revocable) {
		return nil, false
	}
	return o.Revocable, true
}

// HasRevocable returns a boolean if a field has been set.
func (o *RoleSummary) HasRevocable() bool {
	if o != nil && !isNil(o.Revocable) {
		return true
	}

	return false
}

// SetRevocable gets a reference to the given bool and assigns it to the Revocable field.
func (o *RoleSummary) SetRevocable(v bool) {
	o.Revocable = &v
}

func (o RoleSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !isNil(o.Revocable) {
		toSerialize["revocable"] = o.Revocable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleSummary) UnmarshalJSON(bytes []byte) (err error) {
	varRoleSummary := _RoleSummary{}

	if err = json.Unmarshal(bytes, &varRoleSummary); err == nil {
		*o = RoleSummary(varRoleSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "revocable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleSummary struct {
	value *RoleSummary
	isSet bool
}

func (v NullableRoleSummary) Get() *RoleSummary {
	return v.value
}

func (v *NullableRoleSummary) Set(val *RoleSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleSummary(val *RoleSummary) *NullableRoleSummary {
	return &NullableRoleSummary{value: val, isSet: true}
}

func (v NullableRoleSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


