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

// TriggerInputAccountAggregationCompletedStats Overall statistics about the account aggregation.
type TriggerInputAccountAggregationCompletedStats struct {
	// The number of accounts which were scanned / iterated over.
	Scanned int32 `json:"scanned"`
	// The number of accounts which existed before, but had no changes.
	Unchanged int32 `json:"unchanged"`
	// The number of accounts which existed before, but had changes.
	Changed int32 `json:"changed"`
	// The number of accounts which are new - have not existed before.
	Added int32 `json:"added"`
	// The number accounts which existed before, but no longer exist (thus getting removed).
	Removed int32 `json:"removed"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputAccountAggregationCompletedStats TriggerInputAccountAggregationCompletedStats

// NewTriggerInputAccountAggregationCompletedStats instantiates a new TriggerInputAccountAggregationCompletedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputAccountAggregationCompletedStats(scanned int32, unchanged int32, changed int32, added int32, removed int32) *TriggerInputAccountAggregationCompletedStats {
	this := TriggerInputAccountAggregationCompletedStats{}
	this.Scanned = scanned
	this.Unchanged = unchanged
	this.Changed = changed
	this.Added = added
	this.Removed = removed
	return &this
}

// NewTriggerInputAccountAggregationCompletedStatsWithDefaults instantiates a new TriggerInputAccountAggregationCompletedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputAccountAggregationCompletedStatsWithDefaults() *TriggerInputAccountAggregationCompletedStats {
	this := TriggerInputAccountAggregationCompletedStats{}
	return &this
}

// GetScanned returns the Scanned field value
func (o *TriggerInputAccountAggregationCompletedStats) GetScanned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Scanned
}

// GetScannedOk returns a tuple with the Scanned field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAggregationCompletedStats) GetScannedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scanned, true
}

// SetScanned sets field value
func (o *TriggerInputAccountAggregationCompletedStats) SetScanned(v int32) {
	o.Scanned = v
}

// GetUnchanged returns the Unchanged field value
func (o *TriggerInputAccountAggregationCompletedStats) GetUnchanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Unchanged
}

// GetUnchangedOk returns a tuple with the Unchanged field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAggregationCompletedStats) GetUnchangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unchanged, true
}

// SetUnchanged sets field value
func (o *TriggerInputAccountAggregationCompletedStats) SetUnchanged(v int32) {
	o.Unchanged = v
}

// GetChanged returns the Changed field value
func (o *TriggerInputAccountAggregationCompletedStats) GetChanged() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Changed
}

// GetChangedOk returns a tuple with the Changed field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAggregationCompletedStats) GetChangedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Changed, true
}

// SetChanged sets field value
func (o *TriggerInputAccountAggregationCompletedStats) SetChanged(v int32) {
	o.Changed = v
}

// GetAdded returns the Added field value
func (o *TriggerInputAccountAggregationCompletedStats) GetAdded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAggregationCompletedStats) GetAddedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *TriggerInputAccountAggregationCompletedStats) SetAdded(v int32) {
	o.Added = v
}

// GetRemoved returns the Removed field value
func (o *TriggerInputAccountAggregationCompletedStats) GetRemoved() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *TriggerInputAccountAggregationCompletedStats) GetRemovedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *TriggerInputAccountAggregationCompletedStats) SetRemoved(v int32) {
	o.Removed = v
}

func (o TriggerInputAccountAggregationCompletedStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scanned"] = o.Scanned
	}
	if true {
		toSerialize["unchanged"] = o.Unchanged
	}
	if true {
		toSerialize["changed"] = o.Changed
	}
	if true {
		toSerialize["added"] = o.Added
	}
	if true {
		toSerialize["removed"] = o.Removed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputAccountAggregationCompletedStats) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputAccountAggregationCompletedStats := _TriggerInputAccountAggregationCompletedStats{}

	if err = json.Unmarshal(bytes, &varTriggerInputAccountAggregationCompletedStats); err == nil {
		*o = TriggerInputAccountAggregationCompletedStats(varTriggerInputAccountAggregationCompletedStats)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "scanned")
		delete(additionalProperties, "unchanged")
		delete(additionalProperties, "changed")
		delete(additionalProperties, "added")
		delete(additionalProperties, "removed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputAccountAggregationCompletedStats struct {
	value *TriggerInputAccountAggregationCompletedStats
	isSet bool
}

func (v NullableTriggerInputAccountAggregationCompletedStats) Get() *TriggerInputAccountAggregationCompletedStats {
	return v.value
}

func (v *NullableTriggerInputAccountAggregationCompletedStats) Set(val *TriggerInputAccountAggregationCompletedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputAccountAggregationCompletedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputAccountAggregationCompletedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputAccountAggregationCompletedStats(val *TriggerInputAccountAggregationCompletedStats) *NullableTriggerInputAccountAggregationCompletedStats {
	return &NullableTriggerInputAccountAggregationCompletedStats{value: val, isSet: true}
}

func (v NullableTriggerInputAccountAggregationCompletedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputAccountAggregationCompletedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


