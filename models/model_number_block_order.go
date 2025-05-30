/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"encoding/json"
)

// checks if the NumberBlockOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NumberBlockOrder{}

// NumberBlockOrder struct for NumberBlockOrder
type NumberBlockOrder struct {
	Id *string `json:"id,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	// Starting phone number block
	StartingNumber *string `json:"starting_number,omitempty"`
	// The phone number range included in the block.
	Range *int32 `json:"range,omitempty"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount *int32 `json:"phone_numbers_count,omitempty"`
	// Identifies the connection associated to all numbers in the phone number block.
	ConnectionId *string `json:"connection_id,omitempty"`
	// Identifies the messaging profile associated to all numbers in the phone number block.
	MessagingProfileId *string `json:"messaging_profile_id,omitempty"`
	// The status of the order.
	Status *string `json:"status,omitempty"`
	// A customer reference string for customer look ups.
	CustomerReference *string `json:"customer_reference,omitempty"`
	// An ISO 8901 datetime string denoting when the number order was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// An ISO 8901 datetime string for when the number order was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// True if all requirements are met for every phone number, false otherwise.
	RequirementsMet *bool `json:"requirements_met,omitempty"`
}

// NewNumberBlockOrder instantiates a new NumberBlockOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberBlockOrder() *NumberBlockOrder {
	this := NumberBlockOrder{}
	return &this
}

// NewNumberBlockOrderWithDefaults instantiates a new NumberBlockOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberBlockOrderWithDefaults() *NumberBlockOrder {
	this := NumberBlockOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NumberBlockOrder) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NumberBlockOrder) SetRecordType(v string) {
	o.RecordType = &v
}

// GetStartingNumber returns the StartingNumber field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetStartingNumber() string {
	if o == nil || IsNil(o.StartingNumber) {
		var ret string
		return ret
	}
	return *o.StartingNumber
}

// GetStartingNumberOk returns a tuple with the StartingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetStartingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.StartingNumber) {
		return nil, false
	}
	return o.StartingNumber, true
}

// HasStartingNumber returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasStartingNumber() bool {
	if o != nil && !IsNil(o.StartingNumber) {
		return true
	}

	return false
}

// SetStartingNumber gets a reference to the given string and assigns it to the StartingNumber field.
func (o *NumberBlockOrder) SetStartingNumber(v string) {
	o.StartingNumber = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetRange() int32 {
	if o == nil || IsNil(o.Range) {
		var ret int32
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.Range) {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasRange() bool {
	if o != nil && !IsNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given int32 and assigns it to the Range field.
func (o *NumberBlockOrder) SetRange(v int32) {
	o.Range = &v
}

// GetPhoneNumbersCount returns the PhoneNumbersCount field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetPhoneNumbersCount() int32 {
	if o == nil || IsNil(o.PhoneNumbersCount) {
		var ret int32
		return ret
	}
	return *o.PhoneNumbersCount
}

// GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetPhoneNumbersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PhoneNumbersCount) {
		return nil, false
	}
	return o.PhoneNumbersCount, true
}

// HasPhoneNumbersCount returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasPhoneNumbersCount() bool {
	if o != nil && !IsNil(o.PhoneNumbersCount) {
		return true
	}

	return false
}

// SetPhoneNumbersCount gets a reference to the given int32 and assigns it to the PhoneNumbersCount field.
func (o *NumberBlockOrder) SetPhoneNumbersCount(v int32) {
	o.PhoneNumbersCount = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *NumberBlockOrder) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetMessagingProfileId returns the MessagingProfileId field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetMessagingProfileId() string {
	if o == nil || IsNil(o.MessagingProfileId) {
		var ret string
		return ret
	}
	return *o.MessagingProfileId
}

// GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetMessagingProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessagingProfileId) {
		return nil, false
	}
	return o.MessagingProfileId, true
}

// HasMessagingProfileId returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasMessagingProfileId() bool {
	if o != nil && !IsNil(o.MessagingProfileId) {
		return true
	}

	return false
}

// SetMessagingProfileId gets a reference to the given string and assigns it to the MessagingProfileId field.
func (o *NumberBlockOrder) SetMessagingProfileId(v string) {
	o.MessagingProfileId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NumberBlockOrder) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerReference returns the CustomerReference field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetCustomerReference() string {
	if o == nil || IsNil(o.CustomerReference) {
		var ret string
		return ret
	}
	return *o.CustomerReference
}

// GetCustomerReferenceOk returns a tuple with the CustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetCustomerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerReference) {
		return nil, false
	}
	return o.CustomerReference, true
}

// HasCustomerReference returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasCustomerReference() bool {
	if o != nil && !IsNil(o.CustomerReference) {
		return true
	}

	return false
}

// SetCustomerReference gets a reference to the given string and assigns it to the CustomerReference field.
func (o *NumberBlockOrder) SetCustomerReference(v string) {
	o.CustomerReference = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *NumberBlockOrder) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *NumberBlockOrder) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetRequirementsMet returns the RequirementsMet field value if set, zero value otherwise.
func (o *NumberBlockOrder) GetRequirementsMet() bool {
	if o == nil || IsNil(o.RequirementsMet) {
		var ret bool
		return ret
	}
	return *o.RequirementsMet
}

// GetRequirementsMetOk returns a tuple with the RequirementsMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberBlockOrder) GetRequirementsMetOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirementsMet) {
		return nil, false
	}
	return o.RequirementsMet, true
}

// HasRequirementsMet returns a boolean if a field has been set.
func (o *NumberBlockOrder) HasRequirementsMet() bool {
	if o != nil && !IsNil(o.RequirementsMet) {
		return true
	}

	return false
}

// SetRequirementsMet gets a reference to the given bool and assigns it to the RequirementsMet field.
func (o *NumberBlockOrder) SetRequirementsMet(v bool) {
	o.RequirementsMet = &v
}

func (o NumberBlockOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberBlockOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.StartingNumber) {
		toSerialize["starting_number"] = o.StartingNumber
	}
	if !IsNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !IsNil(o.PhoneNumbersCount) {
		toSerialize["phone_numbers_count"] = o.PhoneNumbersCount
	}
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.MessagingProfileId) {
		toSerialize["messaging_profile_id"] = o.MessagingProfileId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CustomerReference) {
		toSerialize["customer_reference"] = o.CustomerReference
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.RequirementsMet) {
		toSerialize["requirements_met"] = o.RequirementsMet
	}
	return toSerialize, nil
}

type NullableNumberBlockOrder struct {
	value *NumberBlockOrder
	isSet bool
}

func (v NullableNumberBlockOrder) Get() *NumberBlockOrder {
	return v.value
}

func (v *NullableNumberBlockOrder) Set(val *NumberBlockOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberBlockOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberBlockOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberBlockOrder(val *NumberBlockOrder) *NullableNumberBlockOrder {
	return &NullableNumberBlockOrder{value: val, isSet: true}
}

func (v NullableNumberBlockOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberBlockOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


