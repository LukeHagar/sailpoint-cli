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

// SavedSearchComplete struct for SavedSearchComplete
type SavedSearchComplete struct {
	// A name for the report file.
	FileName string `json:"fileName"`
	// The email address of the identity that owns the saved search.
	OwnerEmail string `json:"ownerEmail"`
	// The name of the identity that owns the saved search.
	OwnerName string `json:"ownerName"`
	// The search query that was used to generate the report.
	Query string `json:"query"`
	// The name of the saved search.
	SearchName string `json:"searchName"`
	SearchResults TriggerInputSavedSearchCompleteSearchResults `json:"searchResults"`
	// The Amazon S3 URL to download the report from.
	SignedS3Url string `json:"signedS3Url"`
	AdditionalProperties map[string]interface{}
}

type _SavedSearchComplete SavedSearchComplete

// NewSavedSearchComplete instantiates a new SavedSearchComplete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchComplete(fileName string, ownerEmail string, ownerName string, query string, searchName string, searchResults TriggerInputSavedSearchCompleteSearchResults, signedS3Url string) *SavedSearchComplete {
	this := SavedSearchComplete{}
	this.FileName = fileName
	this.OwnerEmail = ownerEmail
	this.OwnerName = ownerName
	this.Query = query
	this.SearchName = searchName
	this.SearchResults = searchResults
	this.SignedS3Url = signedS3Url
	return &this
}

// NewSavedSearchCompleteWithDefaults instantiates a new SavedSearchComplete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchCompleteWithDefaults() *SavedSearchComplete {
	this := SavedSearchComplete{}
	return &this
}

// GetFileName returns the FileName field value
func (o *SavedSearchComplete) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *SavedSearchComplete) SetFileName(v string) {
	o.FileName = v
}

// GetOwnerEmail returns the OwnerEmail field value
func (o *SavedSearchComplete) GetOwnerEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetOwnerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerEmail, true
}

// SetOwnerEmail sets field value
func (o *SavedSearchComplete) SetOwnerEmail(v string) {
	o.OwnerEmail = v
}

// GetOwnerName returns the OwnerName field value
func (o *SavedSearchComplete) GetOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerName, true
}

// SetOwnerName sets field value
func (o *SavedSearchComplete) SetOwnerName(v string) {
	o.OwnerName = v
}

// GetQuery returns the Query field value
func (o *SavedSearchComplete) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *SavedSearchComplete) SetQuery(v string) {
	o.Query = v
}

// GetSearchName returns the SearchName field value
func (o *SavedSearchComplete) GetSearchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchName
}

// GetSearchNameOk returns a tuple with the SearchName field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetSearchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchName, true
}

// SetSearchName sets field value
func (o *SavedSearchComplete) SetSearchName(v string) {
	o.SearchName = v
}

// GetSearchResults returns the SearchResults field value
func (o *SavedSearchComplete) GetSearchResults() TriggerInputSavedSearchCompleteSearchResults {
	if o == nil {
		var ret TriggerInputSavedSearchCompleteSearchResults
		return ret
	}

	return o.SearchResults
}

// GetSearchResultsOk returns a tuple with the SearchResults field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetSearchResultsOk() (*TriggerInputSavedSearchCompleteSearchResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchResults, true
}

// SetSearchResults sets field value
func (o *SavedSearchComplete) SetSearchResults(v TriggerInputSavedSearchCompleteSearchResults) {
	o.SearchResults = v
}

// GetSignedS3Url returns the SignedS3Url field value
func (o *SavedSearchComplete) GetSignedS3Url() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedS3Url
}

// GetSignedS3UrlOk returns a tuple with the SignedS3Url field value
// and a boolean to check if the value has been set.
func (o *SavedSearchComplete) GetSignedS3UrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedS3Url, true
}

// SetSignedS3Url sets field value
func (o *SavedSearchComplete) SetSignedS3Url(v string) {
	o.SignedS3Url = v
}

func (o SavedSearchComplete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fileName"] = o.FileName
	}
	if true {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	if true {
		toSerialize["ownerName"] = o.OwnerName
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["searchName"] = o.SearchName
	}
	if true {
		toSerialize["searchResults"] = o.SearchResults
	}
	if true {
		toSerialize["signedS3Url"] = o.SignedS3Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SavedSearchComplete) UnmarshalJSON(bytes []byte) (err error) {
	varSavedSearchComplete := _SavedSearchComplete{}

	if err = json.Unmarshal(bytes, &varSavedSearchComplete); err == nil {
		*o = SavedSearchComplete(varSavedSearchComplete)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "ownerEmail")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "query")
		delete(additionalProperties, "searchName")
		delete(additionalProperties, "searchResults")
		delete(additionalProperties, "signedS3Url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSavedSearchComplete struct {
	value *SavedSearchComplete
	isSet bool
}

func (v NullableSavedSearchComplete) Get() *SavedSearchComplete {
	return v.value
}

func (v *NullableSavedSearchComplete) Set(val *SavedSearchComplete) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchComplete) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchComplete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchComplete(val *SavedSearchComplete) *NullableSavedSearchComplete {
	return &NullableSavedSearchComplete{value: val, isSet: true}
}

func (v NullableSavedSearchComplete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchComplete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


