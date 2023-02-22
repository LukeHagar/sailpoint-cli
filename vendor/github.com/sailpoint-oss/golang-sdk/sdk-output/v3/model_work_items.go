/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"time"
)

// WorkItems struct for WorkItems
type WorkItems struct {
	// ID of the work item
	Id *string `json:"id,omitempty"`
	// ID of the requester
	RequesterId *string `json:"requesterId,omitempty"`
	// The displayname of the requester
	RequesterDisplayName *string `json:"requesterDisplayName,omitempty"`
	// The ID of the owner
	OwnerId *string `json:"ownerId,omitempty"`
	// The name of the owner
	OwnerName *string `json:"ownerName,omitempty"`
	// Time when the work item was created
	Created *time.Time `json:"created,omitempty"`
	// Time when the work item was last updated
	Modified *time.Time `json:"modified,omitempty"`
	// The description of the work item
	Description *string `json:"description,omitempty"`
	State *WorkItemState `json:"state,omitempty"`
	Type *WorkItemType `json:"type,omitempty"`
	RemediationItems *RemediationItemDetails `json:"remediationItems,omitempty"`
	ApprovalItems *ApprovalItemDetails `json:"approvalItems,omitempty"`
	// The work item name
	Name *string `json:"name,omitempty"`
	// The time at which the work item completed
	Completed *time.Time `json:"completed,omitempty"`
	// The number of items in the work item
	NumItems *int32 `json:"numItems,omitempty"`
	Form *FormDetails `json:"form,omitempty"`
	// An array of errors that ocurred during the work item
	Errors []string `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkItems WorkItems

// NewWorkItems instantiates a new WorkItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkItems() *WorkItems {
	this := WorkItems{}
	return &this
}

// NewWorkItemsWithDefaults instantiates a new WorkItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkItemsWithDefaults() *WorkItems {
	this := WorkItems{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkItems) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkItems) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkItems) SetId(v string) {
	o.Id = &v
}

// GetRequesterId returns the RequesterId field value if set, zero value otherwise.
func (o *WorkItems) GetRequesterId() string {
	if o == nil || isNil(o.RequesterId) {
		var ret string
		return ret
	}
	return *o.RequesterId
}

// GetRequesterIdOk returns a tuple with the RequesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetRequesterIdOk() (*string, bool) {
	if o == nil || isNil(o.RequesterId) {
		return nil, false
	}
	return o.RequesterId, true
}

// HasRequesterId returns a boolean if a field has been set.
func (o *WorkItems) HasRequesterId() bool {
	if o != nil && !isNil(o.RequesterId) {
		return true
	}

	return false
}

// SetRequesterId gets a reference to the given string and assigns it to the RequesterId field.
func (o *WorkItems) SetRequesterId(v string) {
	o.RequesterId = &v
}

// GetRequesterDisplayName returns the RequesterDisplayName field value if set, zero value otherwise.
func (o *WorkItems) GetRequesterDisplayName() string {
	if o == nil || isNil(o.RequesterDisplayName) {
		var ret string
		return ret
	}
	return *o.RequesterDisplayName
}

// GetRequesterDisplayNameOk returns a tuple with the RequesterDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetRequesterDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.RequesterDisplayName) {
		return nil, false
	}
	return o.RequesterDisplayName, true
}

// HasRequesterDisplayName returns a boolean if a field has been set.
func (o *WorkItems) HasRequesterDisplayName() bool {
	if o != nil && !isNil(o.RequesterDisplayName) {
		return true
	}

	return false
}

// SetRequesterDisplayName gets a reference to the given string and assigns it to the RequesterDisplayName field.
func (o *WorkItems) SetRequesterDisplayName(v string) {
	o.RequesterDisplayName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *WorkItems) GetOwnerId() string {
	if o == nil || isNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetOwnerIdOk() (*string, bool) {
	if o == nil || isNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *WorkItems) HasOwnerId() bool {
	if o != nil && !isNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *WorkItems) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *WorkItems) GetOwnerName() string {
	if o == nil || isNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetOwnerNameOk() (*string, bool) {
	if o == nil || isNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *WorkItems) HasOwnerName() bool {
	if o != nil && !isNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *WorkItems) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WorkItems) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WorkItems) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WorkItems) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *WorkItems) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *WorkItems) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *WorkItems) SetModified(v time.Time) {
	o.Modified = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkItems) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkItems) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkItems) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkItems) GetState() WorkItemState {
	if o == nil || isNil(o.State) {
		var ret WorkItemState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetStateOk() (*WorkItemState, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkItems) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given WorkItemState and assigns it to the State field.
func (o *WorkItems) SetState(v WorkItemState) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkItems) GetType() WorkItemType {
	if o == nil || isNil(o.Type) {
		var ret WorkItemType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetTypeOk() (*WorkItemType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkItems) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given WorkItemType and assigns it to the Type field.
func (o *WorkItems) SetType(v WorkItemType) {
	o.Type = &v
}

// GetRemediationItems returns the RemediationItems field value if set, zero value otherwise.
func (o *WorkItems) GetRemediationItems() RemediationItemDetails {
	if o == nil || isNil(o.RemediationItems) {
		var ret RemediationItemDetails
		return ret
	}
	return *o.RemediationItems
}

// GetRemediationItemsOk returns a tuple with the RemediationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetRemediationItemsOk() (*RemediationItemDetails, bool) {
	if o == nil || isNil(o.RemediationItems) {
		return nil, false
	}
	return o.RemediationItems, true
}

// HasRemediationItems returns a boolean if a field has been set.
func (o *WorkItems) HasRemediationItems() bool {
	if o != nil && !isNil(o.RemediationItems) {
		return true
	}

	return false
}

// SetRemediationItems gets a reference to the given RemediationItemDetails and assigns it to the RemediationItems field.
func (o *WorkItems) SetRemediationItems(v RemediationItemDetails) {
	o.RemediationItems = &v
}

// GetApprovalItems returns the ApprovalItems field value if set, zero value otherwise.
func (o *WorkItems) GetApprovalItems() ApprovalItemDetails {
	if o == nil || isNil(o.ApprovalItems) {
		var ret ApprovalItemDetails
		return ret
	}
	return *o.ApprovalItems
}

// GetApprovalItemsOk returns a tuple with the ApprovalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetApprovalItemsOk() (*ApprovalItemDetails, bool) {
	if o == nil || isNil(o.ApprovalItems) {
		return nil, false
	}
	return o.ApprovalItems, true
}

// HasApprovalItems returns a boolean if a field has been set.
func (o *WorkItems) HasApprovalItems() bool {
	if o != nil && !isNil(o.ApprovalItems) {
		return true
	}

	return false
}

// SetApprovalItems gets a reference to the given ApprovalItemDetails and assigns it to the ApprovalItems field.
func (o *WorkItems) SetApprovalItems(v ApprovalItemDetails) {
	o.ApprovalItems = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkItems) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkItems) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkItems) SetName(v string) {
	o.Name = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *WorkItems) GetCompleted() time.Time {
	if o == nil || isNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetCompletedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *WorkItems) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *WorkItems) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetNumItems returns the NumItems field value if set, zero value otherwise.
func (o *WorkItems) GetNumItems() int32 {
	if o == nil || isNil(o.NumItems) {
		var ret int32
		return ret
	}
	return *o.NumItems
}

// GetNumItemsOk returns a tuple with the NumItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetNumItemsOk() (*int32, bool) {
	if o == nil || isNil(o.NumItems) {
		return nil, false
	}
	return o.NumItems, true
}

// HasNumItems returns a boolean if a field has been set.
func (o *WorkItems) HasNumItems() bool {
	if o != nil && !isNil(o.NumItems) {
		return true
	}

	return false
}

// SetNumItems gets a reference to the given int32 and assigns it to the NumItems field.
func (o *WorkItems) SetNumItems(v int32) {
	o.NumItems = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *WorkItems) GetForm() FormDetails {
	if o == nil || isNil(o.Form) {
		var ret FormDetails
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetFormOk() (*FormDetails, bool) {
	if o == nil || isNil(o.Form) {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *WorkItems) HasForm() bool {
	if o != nil && !isNil(o.Form) {
		return true
	}

	return false
}

// SetForm gets a reference to the given FormDetails and assigns it to the Form field.
func (o *WorkItems) SetForm(v FormDetails) {
	o.Form = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *WorkItems) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkItems) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *WorkItems) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *WorkItems) SetErrors(v []string) {
	o.Errors = v
}

func (o WorkItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.RequesterId) {
		toSerialize["requesterId"] = o.RequesterId
	}
	if !isNil(o.RequesterDisplayName) {
		toSerialize["requesterDisplayName"] = o.RequesterDisplayName
	}
	if !isNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !isNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.RemediationItems) {
		toSerialize["remediationItems"] = o.RemediationItems
	}
	if !isNil(o.ApprovalItems) {
		toSerialize["approvalItems"] = o.ApprovalItems
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !isNil(o.NumItems) {
		toSerialize["numItems"] = o.NumItems
	}
	if !isNil(o.Form) {
		toSerialize["form"] = o.Form
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkItems) UnmarshalJSON(bytes []byte) (err error) {
	varWorkItems := _WorkItems{}

	if err = json.Unmarshal(bytes, &varWorkItems); err == nil {
		*o = WorkItems(varWorkItems)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requesterId")
		delete(additionalProperties, "requesterDisplayName")
		delete(additionalProperties, "ownerId")
		delete(additionalProperties, "ownerName")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "type")
		delete(additionalProperties, "remediationItems")
		delete(additionalProperties, "approvalItems")
		delete(additionalProperties, "name")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "numItems")
		delete(additionalProperties, "form")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkItems struct {
	value *WorkItems
	isSet bool
}

func (v NullableWorkItems) Get() *WorkItems {
	return v.value
}

func (v *NullableWorkItems) Set(val *WorkItems) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkItems) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkItems(val *WorkItems) *NullableWorkItems {
	return &NullableWorkItems{value: val, isSet: true}
}

func (v NullableWorkItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


