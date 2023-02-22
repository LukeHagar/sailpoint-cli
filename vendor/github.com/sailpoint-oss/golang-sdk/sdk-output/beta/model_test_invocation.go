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

// TestInvocation struct for TestInvocation
type TestInvocation struct {
	// Trigger ID
	TriggerId string `json:"triggerId"`
	// Mock input to use for test invocation.  This must adhere to the input schema defined in the trigger being invoked.  If this property is omitted, then the default trigger sample payload will be sent.
	Input map[string]interface{} `json:"input,omitempty"`
	// JSON map of invocation metadata.
	ContentJson map[string]interface{} `json:"contentJson"`
	// Only send the test event to the subscription IDs listed.  If omitted, the test event will be sent to all subscribers.
	SubscriptionIds []string `json:"subscriptionIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestInvocation TestInvocation

// NewTestInvocation instantiates a new TestInvocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInvocation(triggerId string, contentJson map[string]interface{}) *TestInvocation {
	this := TestInvocation{}
	this.TriggerId = triggerId
	this.ContentJson = contentJson
	return &this
}

// NewTestInvocationWithDefaults instantiates a new TestInvocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInvocationWithDefaults() *TestInvocation {
	this := TestInvocation{}
	return &this
}

// GetTriggerId returns the TriggerId field value
func (o *TestInvocation) GetTriggerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value
// and a boolean to check if the value has been set.
func (o *TestInvocation) GetTriggerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerId, true
}

// SetTriggerId sets field value
func (o *TestInvocation) SetTriggerId(v string) {
	o.TriggerId = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *TestInvocation) GetInput() map[string]interface{} {
	if o == nil || isNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInvocation) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *TestInvocation) HasInput() bool {
	if o != nil && !isNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *TestInvocation) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetContentJson returns the ContentJson field value
func (o *TestInvocation) GetContentJson() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ContentJson
}

// GetContentJsonOk returns a tuple with the ContentJson field value
// and a boolean to check if the value has been set.
func (o *TestInvocation) GetContentJsonOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.ContentJson, true
}

// SetContentJson sets field value
func (o *TestInvocation) SetContentJson(v map[string]interface{}) {
	o.ContentJson = v
}

// GetSubscriptionIds returns the SubscriptionIds field value if set, zero value otherwise.
func (o *TestInvocation) GetSubscriptionIds() []string {
	if o == nil || isNil(o.SubscriptionIds) {
		var ret []string
		return ret
	}
	return o.SubscriptionIds
}

// GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInvocation) GetSubscriptionIdsOk() ([]string, bool) {
	if o == nil || isNil(o.SubscriptionIds) {
		return nil, false
	}
	return o.SubscriptionIds, true
}

// HasSubscriptionIds returns a boolean if a field has been set.
func (o *TestInvocation) HasSubscriptionIds() bool {
	if o != nil && !isNil(o.SubscriptionIds) {
		return true
	}

	return false
}

// SetSubscriptionIds gets a reference to the given []string and assigns it to the SubscriptionIds field.
func (o *TestInvocation) SetSubscriptionIds(v []string) {
	o.SubscriptionIds = v
}

func (o TestInvocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["triggerId"] = o.TriggerId
	}
	if !isNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if true {
		toSerialize["contentJson"] = o.ContentJson
	}
	if !isNil(o.SubscriptionIds) {
		toSerialize["subscriptionIds"] = o.SubscriptionIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestInvocation) UnmarshalJSON(bytes []byte) (err error) {
	varTestInvocation := _TestInvocation{}

	if err = json.Unmarshal(bytes, &varTestInvocation); err == nil {
		*o = TestInvocation(varTestInvocation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "triggerId")
		delete(additionalProperties, "input")
		delete(additionalProperties, "contentJson")
		delete(additionalProperties, "subscriptionIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInvocation struct {
	value *TestInvocation
	isSet bool
}

func (v NullableTestInvocation) Get() *TestInvocation {
	return v.value
}

func (v *NullableTestInvocation) Set(val *TestInvocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInvocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInvocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInvocation(val *TestInvocation) *NullableTestInvocation {
	return &NullableTestInvocation{value: val, isSet: true}
}

func (v NullableTestInvocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInvocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


