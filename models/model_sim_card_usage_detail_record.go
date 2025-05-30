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
	"bytes"
	"fmt"
)

// checks if the SimCardUsageDetailRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimCardUsageDetailRecord{}

// SimCardUsageDetailRecord struct for SimCardUsageDetailRecord
type SimCardUsageDetailRecord struct {
	// Unique identifier for this SIM Card Usage
	Id *string `json:"id,omitempty"`
	// Event creation time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Event close time
	ClosedAt *time.Time `json:"closed_at,omitempty"`
	// Ip address that generated the event
	IpAddress *string `json:"ip_address,omitempty"`
	// Number of megabytes downloaded
	DownlinkData *float32 `json:"downlink_data,omitempty"`
	// International Mobile Subscriber Identity
	Imsi *string `json:"imsi,omitempty"`
	// Mobile country code
	Mcc *string `json:"mcc,omitempty"`
	// Mobile network code
	Mnc *string `json:"mnc,omitempty"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency *string `json:"currency,omitempty"`
	// Unit of wireless link consumption
	DataUnit *string `json:"data_unit,omitempty"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	DataRate *string `json:"data_rate,omitempty"`
	// Sim group name for sim card
	SimGroupName *string `json:"sim_group_name,omitempty"`
	// Unique identifier for SIM card
	SimCardId *string `json:"sim_card_id,omitempty"`
	// Unique identifier for SIM group
	SimGroupId *string `json:"sim_group_id,omitempty"`
	// User-provided tags
	SimCardTags *string `json:"sim_card_tags,omitempty"`
	// Telephone number associated to SIM card
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Number of megabytes uploaded
	UplinkData *float32 `json:"uplink_data,omitempty"`
	// Data cost
	DataCost *float32 `json:"data_cost,omitempty"`
	RecordType string `json:"record_type"`
}

type _SimCardUsageDetailRecord SimCardUsageDetailRecord

// NewSimCardUsageDetailRecord instantiates a new SimCardUsageDetailRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimCardUsageDetailRecord(recordType string) *SimCardUsageDetailRecord {
	this := SimCardUsageDetailRecord{}
	this.RecordType = recordType
	return &this
}

// NewSimCardUsageDetailRecordWithDefaults instantiates a new SimCardUsageDetailRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimCardUsageDetailRecordWithDefaults() *SimCardUsageDetailRecord {
	this := SimCardUsageDetailRecord{}
	var recordType string = "sim_card_usage"
	this.RecordType = recordType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimCardUsageDetailRecord) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SimCardUsageDetailRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetClosedAt() time.Time {
	if o == nil || IsNil(o.ClosedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedAt) {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasClosedAt() bool {
	if o != nil && !IsNil(o.ClosedAt) {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *SimCardUsageDetailRecord) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SimCardUsageDetailRecord) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetDownlinkData returns the DownlinkData field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetDownlinkData() float32 {
	if o == nil || IsNil(o.DownlinkData) {
		var ret float32
		return ret
	}
	return *o.DownlinkData
}

// GetDownlinkDataOk returns a tuple with the DownlinkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetDownlinkDataOk() (*float32, bool) {
	if o == nil || IsNil(o.DownlinkData) {
		return nil, false
	}
	return o.DownlinkData, true
}

// HasDownlinkData returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasDownlinkData() bool {
	if o != nil && !IsNil(o.DownlinkData) {
		return true
	}

	return false
}

// SetDownlinkData gets a reference to the given float32 and assigns it to the DownlinkData field.
func (o *SimCardUsageDetailRecord) SetDownlinkData(v float32) {
	o.DownlinkData = &v
}

// GetImsi returns the Imsi field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetImsi() string {
	if o == nil || IsNil(o.Imsi) {
		var ret string
		return ret
	}
	return *o.Imsi
}

// GetImsiOk returns a tuple with the Imsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetImsiOk() (*string, bool) {
	if o == nil || IsNil(o.Imsi) {
		return nil, false
	}
	return o.Imsi, true
}

// HasImsi returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasImsi() bool {
	if o != nil && !IsNil(o.Imsi) {
		return true
	}

	return false
}

// SetImsi gets a reference to the given string and assigns it to the Imsi field.
func (o *SimCardUsageDetailRecord) SetImsi(v string) {
	o.Imsi = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *SimCardUsageDetailRecord) SetMcc(v string) {
	o.Mcc = &v
}

// GetMnc returns the Mnc field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetMnc() string {
	if o == nil || IsNil(o.Mnc) {
		var ret string
		return ret
	}
	return *o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetMncOk() (*string, bool) {
	if o == nil || IsNil(o.Mnc) {
		return nil, false
	}
	return o.Mnc, true
}

// HasMnc returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasMnc() bool {
	if o != nil && !IsNil(o.Mnc) {
		return true
	}

	return false
}

// SetMnc gets a reference to the given string and assigns it to the Mnc field.
func (o *SimCardUsageDetailRecord) SetMnc(v string) {
	o.Mnc = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SimCardUsageDetailRecord) SetCurrency(v string) {
	o.Currency = &v
}

// GetDataUnit returns the DataUnit field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetDataUnit() string {
	if o == nil || IsNil(o.DataUnit) {
		var ret string
		return ret
	}
	return *o.DataUnit
}

// GetDataUnitOk returns a tuple with the DataUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetDataUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DataUnit) {
		return nil, false
	}
	return o.DataUnit, true
}

// HasDataUnit returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasDataUnit() bool {
	if o != nil && !IsNil(o.DataUnit) {
		return true
	}

	return false
}

// SetDataUnit gets a reference to the given string and assigns it to the DataUnit field.
func (o *SimCardUsageDetailRecord) SetDataUnit(v string) {
	o.DataUnit = &v
}

// GetDataRate returns the DataRate field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetDataRate() string {
	if o == nil || IsNil(o.DataRate) {
		var ret string
		return ret
	}
	return *o.DataRate
}

// GetDataRateOk returns a tuple with the DataRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetDataRateOk() (*string, bool) {
	if o == nil || IsNil(o.DataRate) {
		return nil, false
	}
	return o.DataRate, true
}

// HasDataRate returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasDataRate() bool {
	if o != nil && !IsNil(o.DataRate) {
		return true
	}

	return false
}

// SetDataRate gets a reference to the given string and assigns it to the DataRate field.
func (o *SimCardUsageDetailRecord) SetDataRate(v string) {
	o.DataRate = &v
}

// GetSimGroupName returns the SimGroupName field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetSimGroupName() string {
	if o == nil || IsNil(o.SimGroupName) {
		var ret string
		return ret
	}
	return *o.SimGroupName
}

// GetSimGroupNameOk returns a tuple with the SimGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetSimGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.SimGroupName) {
		return nil, false
	}
	return o.SimGroupName, true
}

// HasSimGroupName returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasSimGroupName() bool {
	if o != nil && !IsNil(o.SimGroupName) {
		return true
	}

	return false
}

// SetSimGroupName gets a reference to the given string and assigns it to the SimGroupName field.
func (o *SimCardUsageDetailRecord) SetSimGroupName(v string) {
	o.SimGroupName = &v
}

// GetSimCardId returns the SimCardId field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetSimCardId() string {
	if o == nil || IsNil(o.SimCardId) {
		var ret string
		return ret
	}
	return *o.SimCardId
}

// GetSimCardIdOk returns a tuple with the SimCardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetSimCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.SimCardId) {
		return nil, false
	}
	return o.SimCardId, true
}

// HasSimCardId returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasSimCardId() bool {
	if o != nil && !IsNil(o.SimCardId) {
		return true
	}

	return false
}

// SetSimCardId gets a reference to the given string and assigns it to the SimCardId field.
func (o *SimCardUsageDetailRecord) SetSimCardId(v string) {
	o.SimCardId = &v
}

// GetSimGroupId returns the SimGroupId field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetSimGroupId() string {
	if o == nil || IsNil(o.SimGroupId) {
		var ret string
		return ret
	}
	return *o.SimGroupId
}

// GetSimGroupIdOk returns a tuple with the SimGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetSimGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SimGroupId) {
		return nil, false
	}
	return o.SimGroupId, true
}

// HasSimGroupId returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasSimGroupId() bool {
	if o != nil && !IsNil(o.SimGroupId) {
		return true
	}

	return false
}

// SetSimGroupId gets a reference to the given string and assigns it to the SimGroupId field.
func (o *SimCardUsageDetailRecord) SetSimGroupId(v string) {
	o.SimGroupId = &v
}

// GetSimCardTags returns the SimCardTags field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetSimCardTags() string {
	if o == nil || IsNil(o.SimCardTags) {
		var ret string
		return ret
	}
	return *o.SimCardTags
}

// GetSimCardTagsOk returns a tuple with the SimCardTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetSimCardTagsOk() (*string, bool) {
	if o == nil || IsNil(o.SimCardTags) {
		return nil, false
	}
	return o.SimCardTags, true
}

// HasSimCardTags returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasSimCardTags() bool {
	if o != nil && !IsNil(o.SimCardTags) {
		return true
	}

	return false
}

// SetSimCardTags gets a reference to the given string and assigns it to the SimCardTags field.
func (o *SimCardUsageDetailRecord) SetSimCardTags(v string) {
	o.SimCardTags = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *SimCardUsageDetailRecord) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetUplinkData returns the UplinkData field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetUplinkData() float32 {
	if o == nil || IsNil(o.UplinkData) {
		var ret float32
		return ret
	}
	return *o.UplinkData
}

// GetUplinkDataOk returns a tuple with the UplinkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetUplinkDataOk() (*float32, bool) {
	if o == nil || IsNil(o.UplinkData) {
		return nil, false
	}
	return o.UplinkData, true
}

// HasUplinkData returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasUplinkData() bool {
	if o != nil && !IsNil(o.UplinkData) {
		return true
	}

	return false
}

// SetUplinkData gets a reference to the given float32 and assigns it to the UplinkData field.
func (o *SimCardUsageDetailRecord) SetUplinkData(v float32) {
	o.UplinkData = &v
}

// GetDataCost returns the DataCost field value if set, zero value otherwise.
func (o *SimCardUsageDetailRecord) GetDataCost() float32 {
	if o == nil || IsNil(o.DataCost) {
		var ret float32
		return ret
	}
	return *o.DataCost
}

// GetDataCostOk returns a tuple with the DataCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetDataCostOk() (*float32, bool) {
	if o == nil || IsNil(o.DataCost) {
		return nil, false
	}
	return o.DataCost, true
}

// HasDataCost returns a boolean if a field has been set.
func (o *SimCardUsageDetailRecord) HasDataCost() bool {
	if o != nil && !IsNil(o.DataCost) {
		return true
	}

	return false
}

// SetDataCost gets a reference to the given float32 and assigns it to the DataCost field.
func (o *SimCardUsageDetailRecord) SetDataCost(v float32) {
	o.DataCost = &v
}

// GetRecordType returns the RecordType field value
func (o *SimCardUsageDetailRecord) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *SimCardUsageDetailRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *SimCardUsageDetailRecord) SetRecordType(v string) {
	o.RecordType = v
}

func (o SimCardUsageDetailRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimCardUsageDetailRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ClosedAt) {
		toSerialize["closed_at"] = o.ClosedAt
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.DownlinkData) {
		toSerialize["downlink_data"] = o.DownlinkData
	}
	if !IsNil(o.Imsi) {
		toSerialize["imsi"] = o.Imsi
	}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.Mnc) {
		toSerialize["mnc"] = o.Mnc
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DataUnit) {
		toSerialize["data_unit"] = o.DataUnit
	}
	if !IsNil(o.DataRate) {
		toSerialize["data_rate"] = o.DataRate
	}
	if !IsNil(o.SimGroupName) {
		toSerialize["sim_group_name"] = o.SimGroupName
	}
	if !IsNil(o.SimCardId) {
		toSerialize["sim_card_id"] = o.SimCardId
	}
	if !IsNil(o.SimGroupId) {
		toSerialize["sim_group_id"] = o.SimGroupId
	}
	if !IsNil(o.SimCardTags) {
		toSerialize["sim_card_tags"] = o.SimCardTags
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.UplinkData) {
		toSerialize["uplink_data"] = o.UplinkData
	}
	if !IsNil(o.DataCost) {
		toSerialize["data_cost"] = o.DataCost
	}
	toSerialize["record_type"] = o.RecordType
	return toSerialize, nil
}

func (o *SimCardUsageDetailRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"record_type",
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

	varSimCardUsageDetailRecord := _SimCardUsageDetailRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimCardUsageDetailRecord)

	if err != nil {
		return err
	}

	*o = SimCardUsageDetailRecord(varSimCardUsageDetailRecord)

	return err
}

type NullableSimCardUsageDetailRecord struct {
	value *SimCardUsageDetailRecord
	isSet bool
}

func (v NullableSimCardUsageDetailRecord) Get() *SimCardUsageDetailRecord {
	return v.value
}

func (v *NullableSimCardUsageDetailRecord) Set(val *SimCardUsageDetailRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSimCardUsageDetailRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSimCardUsageDetailRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimCardUsageDetailRecord(val *SimCardUsageDetailRecord) *NullableSimCardUsageDetailRecord {
	return &NullableSimCardUsageDetailRecord{value: val, isSet: true}
}

func (v NullableSimCardUsageDetailRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimCardUsageDetailRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


