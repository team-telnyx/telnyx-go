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
	"bytes"
	"fmt"
)

// checks if the CreateNumberBlockOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNumberBlockOrderRequest{}

// CreateNumberBlockOrderRequest struct for CreateNumberBlockOrderRequest
type CreateNumberBlockOrderRequest struct {
	Id *string `json:"id,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	// Starting phone number block
	StartingNumber string `json:"starting_number"`
	// The phone number range included in the block.
	Range int32 `json:"range"`
	// The count of phone numbers in the number order.
	PhoneNumbersCount *int32 `json:"phone_numbers_count,omitempty"`
	// Identifies the connection associated with this phone number.
	ConnectionId *string `json:"connection_id,omitempty"`
	// Identifies the messaging profile associated with the phone number.
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
	// Errors the reservation could happen upon
	Errors *string `json:"errors,omitempty"`
}

type _CreateNumberBlockOrderRequest CreateNumberBlockOrderRequest

// NewCreateNumberBlockOrderRequest instantiates a new CreateNumberBlockOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNumberBlockOrderRequest(startingNumber string, range_ int32) *CreateNumberBlockOrderRequest {
	this := CreateNumberBlockOrderRequest{}
	this.StartingNumber = startingNumber
	this.Range = range_
	return &this
}

// NewCreateNumberBlockOrderRequestWithDefaults instantiates a new CreateNumberBlockOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNumberBlockOrderRequestWithDefaults() *CreateNumberBlockOrderRequest {
	this := CreateNumberBlockOrderRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateNumberBlockOrderRequest) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *CreateNumberBlockOrderRequest) SetRecordType(v string) {
	o.RecordType = &v
}

// GetStartingNumber returns the StartingNumber field value
func (o *CreateNumberBlockOrderRequest) GetStartingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartingNumber
}

// GetStartingNumberOk returns a tuple with the StartingNumber field value
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetStartingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartingNumber, true
}

// SetStartingNumber sets field value
func (o *CreateNumberBlockOrderRequest) SetStartingNumber(v string) {
	o.StartingNumber = v
}

// GetRange returns the Range field value
func (o *CreateNumberBlockOrderRequest) GetRange() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetRangeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *CreateNumberBlockOrderRequest) SetRange(v int32) {
	o.Range = v
}

// GetPhoneNumbersCount returns the PhoneNumbersCount field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetPhoneNumbersCount() int32 {
	if o == nil || IsNil(o.PhoneNumbersCount) {
		var ret int32
		return ret
	}
	return *o.PhoneNumbersCount
}

// GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetPhoneNumbersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PhoneNumbersCount) {
		return nil, false
	}
	return o.PhoneNumbersCount, true
}

// HasPhoneNumbersCount returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasPhoneNumbersCount() bool {
	if o != nil && !IsNil(o.PhoneNumbersCount) {
		return true
	}

	return false
}

// SetPhoneNumbersCount gets a reference to the given int32 and assigns it to the PhoneNumbersCount field.
func (o *CreateNumberBlockOrderRequest) SetPhoneNumbersCount(v int32) {
	o.PhoneNumbersCount = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *CreateNumberBlockOrderRequest) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetMessagingProfileId returns the MessagingProfileId field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetMessagingProfileId() string {
	if o == nil || IsNil(o.MessagingProfileId) {
		var ret string
		return ret
	}
	return *o.MessagingProfileId
}

// GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetMessagingProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessagingProfileId) {
		return nil, false
	}
	return o.MessagingProfileId, true
}

// HasMessagingProfileId returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasMessagingProfileId() bool {
	if o != nil && !IsNil(o.MessagingProfileId) {
		return true
	}

	return false
}

// SetMessagingProfileId gets a reference to the given string and assigns it to the MessagingProfileId field.
func (o *CreateNumberBlockOrderRequest) SetMessagingProfileId(v string) {
	o.MessagingProfileId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateNumberBlockOrderRequest) SetStatus(v string) {
	o.Status = &v
}

// GetCustomerReference returns the CustomerReference field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetCustomerReference() string {
	if o == nil || IsNil(o.CustomerReference) {
		var ret string
		return ret
	}
	return *o.CustomerReference
}

// GetCustomerReferenceOk returns a tuple with the CustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetCustomerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerReference) {
		return nil, false
	}
	return o.CustomerReference, true
}

// HasCustomerReference returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasCustomerReference() bool {
	if o != nil && !IsNil(o.CustomerReference) {
		return true
	}

	return false
}

// SetCustomerReference gets a reference to the given string and assigns it to the CustomerReference field.
func (o *CreateNumberBlockOrderRequest) SetCustomerReference(v string) {
	o.CustomerReference = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CreateNumberBlockOrderRequest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CreateNumberBlockOrderRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetRequirementsMet returns the RequirementsMet field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetRequirementsMet() bool {
	if o == nil || IsNil(o.RequirementsMet) {
		var ret bool
		return ret
	}
	return *o.RequirementsMet
}

// GetRequirementsMetOk returns a tuple with the RequirementsMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetRequirementsMetOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirementsMet) {
		return nil, false
	}
	return o.RequirementsMet, true
}

// HasRequirementsMet returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasRequirementsMet() bool {
	if o != nil && !IsNil(o.RequirementsMet) {
		return true
	}

	return false
}

// SetRequirementsMet gets a reference to the given bool and assigns it to the RequirementsMet field.
func (o *CreateNumberBlockOrderRequest) SetRequirementsMet(v bool) {
	o.RequirementsMet = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CreateNumberBlockOrderRequest) GetErrors() string {
	if o == nil || IsNil(o.Errors) {
		var ret string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberBlockOrderRequest) GetErrorsOk() (*string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CreateNumberBlockOrderRequest) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given string and assigns it to the Errors field.
func (o *CreateNumberBlockOrderRequest) SetErrors(v string) {
	o.Errors = &v
}

func (o CreateNumberBlockOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNumberBlockOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	toSerialize["starting_number"] = o.StartingNumber
	toSerialize["range"] = o.Range
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
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

func (o *CreateNumberBlockOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"starting_number",
		"range",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNumberBlockOrderRequest := _CreateNumberBlockOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNumberBlockOrderRequest)

	if err != nil {
		return err
	}

	*o = CreateNumberBlockOrderRequest(varCreateNumberBlockOrderRequest)

	return err
}

type NullableCreateNumberBlockOrderRequest struct {
	value *CreateNumberBlockOrderRequest
	isSet bool
}

func (v NullableCreateNumberBlockOrderRequest) Get() *CreateNumberBlockOrderRequest {
	return v.value
}

func (v *NullableCreateNumberBlockOrderRequest) Set(val *CreateNumberBlockOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNumberBlockOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNumberBlockOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNumberBlockOrderRequest(val *CreateNumberBlockOrderRequest) *NullableCreateNumberBlockOrderRequest {
	return &NullableCreateNumberBlockOrderRequest{value: val, isSet: true}
}

func (v NullableCreateNumberBlockOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNumberBlockOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


