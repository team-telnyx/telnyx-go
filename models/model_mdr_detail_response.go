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
	"time"
)

// checks if the MdrDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MdrDetailResponse{}

// MdrDetailResponse struct for MdrDetailResponse
type MdrDetailResponse struct {
	// Message sent time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Configured profile name. New profiles can be created and configured on Telnyx portal
	ProfileName *string `json:"profile_name,omitempty"`
	// Direction of message - inbound or outbound.
	Direction *string `json:"direction,omitempty"`
	// Number of parts this message has. Max number of character is 160. If message contains more characters then that it will be broken down in multiple parts
	Parts *float32 `json:"parts,omitempty"`
	// Message status
	Status *string `json:"status,omitempty"`
	// The destination number for a call, or the callee
	Cld *string `json:"cld,omitempty"`
	// The number associated with the person initiating the call, or the caller
	Cli *string `json:"cli,omitempty"`
	// Rate applied to the message
	Rate *string `json:"rate,omitempty"`
	// Final cost. Cost is calculated as rate * parts
	Cost *string `json:"cost,omitempty"`
	// Currency of the rate and cost
	Currency *string `json:"currency,omitempty"`
	// Id of message detail record
	Id *string `json:"id,omitempty"`
	// Type of message
	MessageType *string `json:"message_type,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
}

// NewMdrDetailResponse instantiates a new MdrDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMdrDetailResponse() *MdrDetailResponse {
	this := MdrDetailResponse{}
	return &this
}

// NewMdrDetailResponseWithDefaults instantiates a new MdrDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMdrDetailResponseWithDefaults() *MdrDetailResponse {
	this := MdrDetailResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MdrDetailResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetProfileName() string {
	if o == nil || IsNil(o.ProfileName) {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileName) {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasProfileName() bool {
	if o != nil && !IsNil(o.ProfileName) {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *MdrDetailResponse) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MdrDetailResponse) SetDirection(v string) {
	o.Direction = &v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetParts() float32 {
	if o == nil || IsNil(o.Parts) {
		var ret float32
		return ret
	}
	return *o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetPartsOk() (*float32, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given float32 and assigns it to the Parts field.
func (o *MdrDetailResponse) SetParts(v float32) {
	o.Parts = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MdrDetailResponse) SetStatus(v string) {
	o.Status = &v
}

// GetCld returns the Cld field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetCld() string {
	if o == nil || IsNil(o.Cld) {
		var ret string
		return ret
	}
	return *o.Cld
}

// GetCldOk returns a tuple with the Cld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetCldOk() (*string, bool) {
	if o == nil || IsNil(o.Cld) {
		return nil, false
	}
	return o.Cld, true
}

// HasCld returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasCld() bool {
	if o != nil && !IsNil(o.Cld) {
		return true
	}

	return false
}

// SetCld gets a reference to the given string and assigns it to the Cld field.
func (o *MdrDetailResponse) SetCld(v string) {
	o.Cld = &v
}

// GetCli returns the Cli field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetCli() string {
	if o == nil || IsNil(o.Cli) {
		var ret string
		return ret
	}
	return *o.Cli
}

// GetCliOk returns a tuple with the Cli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetCliOk() (*string, bool) {
	if o == nil || IsNil(o.Cli) {
		return nil, false
	}
	return o.Cli, true
}

// HasCli returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasCli() bool {
	if o != nil && !IsNil(o.Cli) {
		return true
	}

	return false
}

// SetCli gets a reference to the given string and assigns it to the Cli field.
func (o *MdrDetailResponse) SetCli(v string) {
	o.Cli = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *MdrDetailResponse) SetRate(v string) {
	o.Rate = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetCost() string {
	if o == nil || IsNil(o.Cost) {
		var ret string
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetCostOk() (*string, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given string and assigns it to the Cost field.
func (o *MdrDetailResponse) SetCost(v string) {
	o.Cost = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MdrDetailResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MdrDetailResponse) SetId(v string) {
	o.Id = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetMessageType() string {
	if o == nil || IsNil(o.MessageType) {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetMessageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *MdrDetailResponse) SetMessageType(v string) {
	o.MessageType = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *MdrDetailResponse) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MdrDetailResponse) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *MdrDetailResponse) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *MdrDetailResponse) SetRecordType(v string) {
	o.RecordType = &v
}

func (o MdrDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MdrDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ProfileName) {
		toSerialize["profile_name"] = o.ProfileName
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Parts) {
		toSerialize["parts"] = o.Parts
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Cld) {
		toSerialize["cld"] = o.Cld
	}
	if !IsNil(o.Cli) {
		toSerialize["cli"] = o.Cli
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MessageType) {
		toSerialize["message_type"] = o.MessageType
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	return toSerialize, nil
}

type NullableMdrDetailResponse struct {
	value *MdrDetailResponse
	isSet bool
}

func (v NullableMdrDetailResponse) Get() *MdrDetailResponse {
	return v.value
}

func (v *NullableMdrDetailResponse) Set(val *MdrDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMdrDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMdrDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdrDetailResponse(val *MdrDetailResponse) *NullableMdrDetailResponse {
	return &NullableMdrDetailResponse{value: val, isSet: true}
}

func (v NullableMdrDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdrDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


