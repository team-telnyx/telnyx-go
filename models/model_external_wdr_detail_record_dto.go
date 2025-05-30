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

// checks if the ExternalWdrDetailRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalWdrDetailRecordDto{}

// ExternalWdrDetailRecordDto struct for ExternalWdrDetailRecordDto
type ExternalWdrDetailRecordDto struct {
	// WDR id
	Id *string `json:"id,omitempty"`
	// Record created time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Cost *WirelessCost `json:"cost,omitempty"`
	// Mobile country code.
	Mcc *string `json:"mcc,omitempty"`
	// Mobile network code.
	Mnc *string `json:"mnc,omitempty"`
	DownlinkData *DownlinkData `json:"downlink_data,omitempty"`
	// Session duration in seconds.
	DurationSeconds *float32 `json:"duration_seconds,omitempty"`
	// International mobile subscriber identity.
	Imsi *string `json:"imsi,omitempty"`
	Rate *WirelessRate `json:"rate,omitempty"`
	// Defined sim group name
	SimGroupName *string `json:"sim_group_name,omitempty"`
	// Sim group unique identifier
	SimGroupId *string `json:"sim_group_id,omitempty"`
	// Sim card unique identifier
	SimCardId *string `json:"sim_card_id,omitempty"`
	// Phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	UplinkData *UplinkData `json:"uplink_data,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
}

// NewExternalWdrDetailRecordDto instantiates a new ExternalWdrDetailRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalWdrDetailRecordDto() *ExternalWdrDetailRecordDto {
	this := ExternalWdrDetailRecordDto{}
	return &this
}

// NewExternalWdrDetailRecordDtoWithDefaults instantiates a new ExternalWdrDetailRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalWdrDetailRecordDtoWithDefaults() *ExternalWdrDetailRecordDto {
	this := ExternalWdrDetailRecordDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalWdrDetailRecordDto) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ExternalWdrDetailRecordDto) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetCost() WirelessCost {
	if o == nil || IsNil(o.Cost) {
		var ret WirelessCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetCostOk() (*WirelessCost, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given WirelessCost and assigns it to the Cost field.
func (o *ExternalWdrDetailRecordDto) SetCost(v WirelessCost) {
	o.Cost = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *ExternalWdrDetailRecordDto) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetMnc() string {
	if o == nil || IsNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetMncOk() (*string, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *ExternalWdrDetailRecordDto) SetMnc(v string) {
	o.Mnc = &v
}

// GetDownlinkData returns the DownlinkData field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetDownlinkData() DownlinkData {
	if o == nil || IsNil(o.DownlinkData) {
		var ret DownlinkData
		return ret
	}
	return *o.DownlinkData
}

// GetDownlinkDataOk returns a tuple with the DownlinkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetDownlinkDataOk() (*DownlinkData, bool) {
	if o == nil || IsNil(o.DownlinkData) {
		return nil, false
	}
	return o.DownlinkData, true
}

// HasDownlinkData returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasDownlinkData() bool {
	if o != nil && !IsNil(o.DownlinkData) {
		return true
	}

	return false
}

// SetDownlinkData gets a reference to the given DownlinkData and assigns it to the DownlinkData field.
func (o *ExternalWdrDetailRecordDto) SetDownlinkData(v DownlinkData) {
	o.DownlinkData = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetDurationSeconds() float32 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret float32
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetDurationSecondsOk() (*float32, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given float32 and assigns it to the DurationSeconds field.
func (o *ExternalWdrDetailRecordDto) SetDurationSeconds(v float32) {
	o.DurationSeconds = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetImsi() string {
	if o == nil || IsNil(o.Imsi) {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetImsiOk() (*string, bool) {
	if o == nil || IsNil(o.Imsi) {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasImsi() bool {
	if o != nil && !IsNil(o.Imsi) {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *ExternalWdrDetailRecordDto) SetImsi(v string) {
	o.Imsi = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetRate() WirelessRate {
	if o == nil || IsNil(o.Rate) {
		var ret WirelessRate
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetRateOk() (*WirelessRate, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given WirelessRate and assigns it to the Rate field.
func (o *ExternalWdrDetailRecordDto) SetRate(v WirelessRate) {
	o.Rate = &v
}

// GetSimGroupName returns the SimGroupName field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetSimGroupName() string {
	if o == nil || IsNil(o.SimGroupName) {
		var ret string
		return ret
	}
	return *o.SimGroupName
}

// GetSimGroupNameOk returns a tuple with the SimGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetSimGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.SimGroupName) {
		return nil, false
	}
	return o.SimGroupName, true
}

// HasSimGroupName returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasSimGroupName() bool {
	if o != nil && !IsNil(o.SimGroupName) {
		return true
	}

	return false
}

// SetSimGroupName gets a reference to the given string and assigns it to the SimGroupName field.
func (o *ExternalWdrDetailRecordDto) SetSimGroupName(v string) {
	o.SimGroupName = &v
}

// GetSimGroupId returns the SimGroupId field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetSimGroupId() string {
	if o == nil || IsNil(o.SimGroupId) {
		var ret string
		return ret
	}
	return *o.SimGroupId
}

// GetSimGroupIdOk returns a tuple with the SimGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetSimGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SimGroupId) {
		return nil, false
	}
	return o.SimGroupId, true
}

// HasSimGroupId returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasSimGroupId() bool {
	if o != nil && !IsNil(o.SimGroupId) {
		return true
	}

	return false
}

// SetSimGroupId gets a reference to the given string and assigns it to the SimGroupId field.
func (o *ExternalWdrDetailRecordDto) SetSimGroupId(v string) {
	o.SimGroupId = &v
}

// GetSimCardId returns the SimCardId field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetSimCardId() string {
	if o == nil || IsNil(o.SimCardId) {
		var ret string
		return ret
	}
	return *o.SimCardId
}

// GetSimCardIdOk returns a tuple with the SimCardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetSimCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.SimCardId) {
		return nil, false
	}
	return o.SimCardId, true
}

// HasSimCardId returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasSimCardId() bool {
	if o != nil && !IsNil(o.SimCardId) {
		return true
	}

	return false
}

// SetSimCardId gets a reference to the given string and assigns it to the SimCardId field.
func (o *ExternalWdrDetailRecordDto) SetSimCardId(v string) {
	o.SimCardId = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ExternalWdrDetailRecordDto) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetUplinkData returns the UplinkData field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetUplinkData() UplinkData {
	if o == nil || IsNil(o.UplinkData) {
		var ret UplinkData
		return ret
	}
	return *o.UplinkData
}

// GetUplinkDataOk returns a tuple with the UplinkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetUplinkDataOk() (*UplinkData, bool) {
	if o == nil || IsNil(o.UplinkData) {
		return nil, false
	}
	return o.UplinkData, true
}

// HasUplinkData returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasUplinkData() bool {
	if o != nil && !IsNil(o.UplinkData) {
		return true
	}

	return false
}

// SetUplinkData gets a reference to the given UplinkData and assigns it to the UplinkData field.
func (o *ExternalWdrDetailRecordDto) SetUplinkData(v UplinkData) {
	o.UplinkData = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ExternalWdrDetailRecordDto) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalWdrDetailRecordDto) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ExternalWdrDetailRecordDto) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *ExternalWdrDetailRecordDto) SetRecordType(v string) {
	o.RecordType = &v
}

func (o ExternalWdrDetailRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalWdrDetailRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !IsNil(o.DownlinkData) {
		toSerialize["downlink_data"] = o.DownlinkData
	}
	if !IsNil(o.DurationSeconds) {
		toSerialize["duration_seconds"] = o.DurationSeconds
	}
	if !IsNil(o.Imsi) {
		toSerialize["imsi"] = o.Imsi
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.SimGroupName) {
		toSerialize["sim_group_name"] = o.SimGroupName
	}
	if !IsNil(o.SimGroupId) {
		toSerialize["sim_group_id"] = o.SimGroupId
	}
	if !IsNil(o.SimCardId) {
		toSerialize["sim_card_id"] = o.SimCardId
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.UplinkData) {
		toSerialize["uplink_data"] = o.UplinkData
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	return toSerialize, nil
}

type NullableExternalWdrDetailRecordDto struct {
	value *ExternalWdrDetailRecordDto
	isSet bool
}

func (v NullableExternalWdrDetailRecordDto) Get() *ExternalWdrDetailRecordDto {
	return v.value
}

func (v *NullableExternalWdrDetailRecordDto) Set(val *ExternalWdrDetailRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalWdrDetailRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalWdrDetailRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalWdrDetailRecordDto(val *ExternalWdrDetailRecordDto) *NullableExternalWdrDetailRecordDto {
	return &NullableExternalWdrDetailRecordDto{value: val, isSet: true}
}

func (v NullableExternalWdrDetailRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalWdrDetailRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


