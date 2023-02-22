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

// PostExternalExecuteWorkflow200Response struct for PostExternalExecuteWorkflow200Response
type PostExternalExecuteWorkflow200Response struct {
	// The workflow execution id
	WorkflowExecutionId *string `json:"workflowExecutionId,omitempty"`
	// An error message if any errors occurred
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostExternalExecuteWorkflow200Response PostExternalExecuteWorkflow200Response

// NewPostExternalExecuteWorkflow200Response instantiates a new PostExternalExecuteWorkflow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostExternalExecuteWorkflow200Response() *PostExternalExecuteWorkflow200Response {
	this := PostExternalExecuteWorkflow200Response{}
	return &this
}

// NewPostExternalExecuteWorkflow200ResponseWithDefaults instantiates a new PostExternalExecuteWorkflow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostExternalExecuteWorkflow200ResponseWithDefaults() *PostExternalExecuteWorkflow200Response {
	this := PostExternalExecuteWorkflow200Response{}
	return &this
}

// GetWorkflowExecutionId returns the WorkflowExecutionId field value if set, zero value otherwise.
func (o *PostExternalExecuteWorkflow200Response) GetWorkflowExecutionId() string {
	if o == nil || isNil(o.WorkflowExecutionId) {
		var ret string
		return ret
	}
	return *o.WorkflowExecutionId
}

// GetWorkflowExecutionIdOk returns a tuple with the WorkflowExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExternalExecuteWorkflow200Response) GetWorkflowExecutionIdOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowExecutionId) {
		return nil, false
	}
	return o.WorkflowExecutionId, true
}

// HasWorkflowExecutionId returns a boolean if a field has been set.
func (o *PostExternalExecuteWorkflow200Response) HasWorkflowExecutionId() bool {
	if o != nil && !isNil(o.WorkflowExecutionId) {
		return true
	}

	return false
}

// SetWorkflowExecutionId gets a reference to the given string and assigns it to the WorkflowExecutionId field.
func (o *PostExternalExecuteWorkflow200Response) SetWorkflowExecutionId(v string) {
	o.WorkflowExecutionId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostExternalExecuteWorkflow200Response) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostExternalExecuteWorkflow200Response) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostExternalExecuteWorkflow200Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostExternalExecuteWorkflow200Response) SetMessage(v string) {
	o.Message = &v
}

func (o PostExternalExecuteWorkflow200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WorkflowExecutionId) {
		toSerialize["workflowExecutionId"] = o.WorkflowExecutionId
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PostExternalExecuteWorkflow200Response) UnmarshalJSON(bytes []byte) (err error) {
	varPostExternalExecuteWorkflow200Response := _PostExternalExecuteWorkflow200Response{}

	if err = json.Unmarshal(bytes, &varPostExternalExecuteWorkflow200Response); err == nil {
		*o = PostExternalExecuteWorkflow200Response(varPostExternalExecuteWorkflow200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workflowExecutionId")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostExternalExecuteWorkflow200Response struct {
	value *PostExternalExecuteWorkflow200Response
	isSet bool
}

func (v NullablePostExternalExecuteWorkflow200Response) Get() *PostExternalExecuteWorkflow200Response {
	return v.value
}

func (v *NullablePostExternalExecuteWorkflow200Response) Set(val *PostExternalExecuteWorkflow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostExternalExecuteWorkflow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostExternalExecuteWorkflow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostExternalExecuteWorkflow200Response(val *PostExternalExecuteWorkflow200Response) *NullablePostExternalExecuteWorkflow200Response {
	return &NullablePostExternalExecuteWorkflow200Response{value: val, isSet: true}
}

func (v NullablePostExternalExecuteWorkflow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostExternalExecuteWorkflow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


