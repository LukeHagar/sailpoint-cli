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

// NonEmployeeSourceLiteWithSchemaAttributes struct for NonEmployeeSourceLiteWithSchemaAttributes
type NonEmployeeSourceLiteWithSchemaAttributes struct {
	// Non-Employee source id.
	Id *string `json:"id,omitempty"`
	// Source Id associated with this non-employee source.
	SourceId *string `json:"sourceId,omitempty"`
	// Source name associated with this non-employee source.
	Name *string `json:"name,omitempty"`
	// Source description associated with this non-employee source.
	Description *string `json:"description,omitempty"`
	// List of schema attributes associated with this non-employee source.
	SchemaAttributes []NonEmployeeSchemaAttribute `json:"schemaAttributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NonEmployeeSourceLiteWithSchemaAttributes NonEmployeeSourceLiteWithSchemaAttributes

// NewNonEmployeeSourceLiteWithSchemaAttributes instantiates a new NonEmployeeSourceLiteWithSchemaAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonEmployeeSourceLiteWithSchemaAttributes() *NonEmployeeSourceLiteWithSchemaAttributes {
	this := NonEmployeeSourceLiteWithSchemaAttributes{}
	return &this
}

// NewNonEmployeeSourceLiteWithSchemaAttributesWithDefaults instantiates a new NonEmployeeSourceLiteWithSchemaAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonEmployeeSourceLiteWithSchemaAttributesWithDefaults() *NonEmployeeSourceLiteWithSchemaAttributes {
	this := NonEmployeeSourceLiteWithSchemaAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) SetSourceId(v string) {
	o.SourceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetSchemaAttributes returns the SchemaAttributes field value if set, zero value otherwise.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetSchemaAttributes() []NonEmployeeSchemaAttribute {
	if o == nil || isNil(o.SchemaAttributes) {
		var ret []NonEmployeeSchemaAttribute
		return ret
	}
	return o.SchemaAttributes
}

// GetSchemaAttributesOk returns a tuple with the SchemaAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) GetSchemaAttributesOk() ([]NonEmployeeSchemaAttribute, bool) {
	if o == nil || isNil(o.SchemaAttributes) {
		return nil, false
	}
	return o.SchemaAttributes, true
}

// HasSchemaAttributes returns a boolean if a field has been set.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) HasSchemaAttributes() bool {
	if o != nil && !isNil(o.SchemaAttributes) {
		return true
	}

	return false
}

// SetSchemaAttributes gets a reference to the given []NonEmployeeSchemaAttribute and assigns it to the SchemaAttributes field.
func (o *NonEmployeeSourceLiteWithSchemaAttributes) SetSchemaAttributes(v []NonEmployeeSchemaAttribute) {
	o.SchemaAttributes = v
}

func (o NonEmployeeSourceLiteWithSchemaAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.SchemaAttributes) {
		toSerialize["schemaAttributes"] = o.SchemaAttributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NonEmployeeSourceLiteWithSchemaAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varNonEmployeeSourceLiteWithSchemaAttributes := _NonEmployeeSourceLiteWithSchemaAttributes{}

	if err = json.Unmarshal(bytes, &varNonEmployeeSourceLiteWithSchemaAttributes); err == nil {
		*o = NonEmployeeSourceLiteWithSchemaAttributes(varNonEmployeeSourceLiteWithSchemaAttributes)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "sourceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "schemaAttributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNonEmployeeSourceLiteWithSchemaAttributes struct {
	value *NonEmployeeSourceLiteWithSchemaAttributes
	isSet bool
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributes) Get() *NonEmployeeSourceLiteWithSchemaAttributes {
	return v.value
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributes) Set(val *NonEmployeeSourceLiteWithSchemaAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonEmployeeSourceLiteWithSchemaAttributes(val *NonEmployeeSourceLiteWithSchemaAttributes) *NullableNonEmployeeSourceLiteWithSchemaAttributes {
	return &NullableNonEmployeeSourceLiteWithSchemaAttributes{value: val, isSet: true}
}

func (v NullableNonEmployeeSourceLiteWithSchemaAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonEmployeeSourceLiteWithSchemaAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


