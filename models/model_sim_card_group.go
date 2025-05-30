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

// checks if the SIMCardGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIMCardGroup{}

// SIMCardGroup struct for SIMCardGroup
type SIMCardGroup struct {
	// Identifies the resource.
	Id *string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// Indicates whether the SIM card group is the users default group.<br/>The default group is created for the user and can not be removed.
	Default *bool `json:"default,omitempty"`
	// A user friendly name for the SIM card group.
	Name *string `json:"name,omitempty"`
	DataLimit *SIMCardGroupDataLimit `json:"data_limit,omitempty"`
	ConsumedData *ConsumedData `json:"consumed_data,omitempty"`
	// The identification of the related Private Wireless Gateway resource.
	PrivateWirelessGatewayId *string `json:"private_wireless_gateway_id,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewSIMCardGroup instantiates a new SIMCardGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIMCardGroup() *SIMCardGroup {
	this := SIMCardGroup{}
	return &this
}

// NewSIMCardGroupWithDefaults instantiates a new SIMCardGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIMCardGroupWithDefaults() *SIMCardGroup {
	this := SIMCardGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SIMCardGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SIMCardGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SIMCardGroup) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *SIMCardGroup) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *SIMCardGroup) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *SIMCardGroup) SetRecordType(v string) {
	o.RecordType = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *SIMCardGroup) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *SIMCardGroup) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *SIMCardGroup) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SIMCardGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SIMCardGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SIMCardGroup) SetName(v string) {
	o.Name = &v
}

// GetDataLimit returns the DataLimit field value if set, zero value otherwise.
func (o *SIMCardGroup) GetDataLimit() SIMCardGroupDataLimit {
	if o == nil || IsNil(o.DataLimit) {
		var ret SIMCardGroupDataLimit
		return ret
	}
	return *o.DataLimit
}

// GetDataLimitOk returns a tuple with the DataLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetDataLimitOk() (*SIMCardGroupDataLimit, bool) {
	if o == nil || IsNil(o.DataLimit) {
		return nil, false
	}
	return o.DataLimit, true
}

// HasDataLimit returns a boolean if a field has been set.
func (o *SIMCardGroup) HasDataLimit() bool {
	if o != nil && !IsNil(o.DataLimit) {
		return true
	}

	return false
}

// SetDataLimit gets a reference to the given SIMCardGroupDataLimit and assigns it to the DataLimit field.
func (o *SIMCardGroup) SetDataLimit(v SIMCardGroupDataLimit) {
	o.DataLimit = &v
}

// GetConsumedData returns the ConsumedData field value if set, zero value otherwise.
func (o *SIMCardGroup) GetConsumedData() ConsumedData {
	if o == nil || IsNil(o.ConsumedData) {
		var ret ConsumedData
		return ret
	}
	return *o.ConsumedData
}

// GetConsumedDataOk returns a tuple with the ConsumedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetConsumedDataOk() (*ConsumedData, bool) {
	if o == nil || IsNil(o.ConsumedData) {
		return nil, false
	}
	return o.ConsumedData, true
}

// HasConsumedData returns a boolean if a field has been set.
func (o *SIMCardGroup) HasConsumedData() bool {
	if o != nil && !IsNil(o.ConsumedData) {
		return true
	}

	return false
}

// SetConsumedData gets a reference to the given ConsumedData and assigns it to the ConsumedData field.
func (o *SIMCardGroup) SetConsumedData(v ConsumedData) {
	o.ConsumedData = &v
}

// GetPrivateWirelessGatewayId returns the PrivateWirelessGatewayId field value if set, zero value otherwise.
func (o *SIMCardGroup) GetPrivateWirelessGatewayId() string {
	if o == nil || IsNil(o.PrivateWirelessGatewayId) {
		var ret string
		return ret
	}
	return *o.PrivateWirelessGatewayId
}

// GetPrivateWirelessGatewayIdOk returns a tuple with the PrivateWirelessGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetPrivateWirelessGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateWirelessGatewayId) {
		return nil, false
	}
	return o.PrivateWirelessGatewayId, true
}

// HasPrivateWirelessGatewayId returns a boolean if a field has been set.
func (o *SIMCardGroup) HasPrivateWirelessGatewayId() bool {
	if o != nil && !IsNil(o.PrivateWirelessGatewayId) {
		return true
	}

	return false
}

// SetPrivateWirelessGatewayId gets a reference to the given string and assigns it to the PrivateWirelessGatewayId field.
func (o *SIMCardGroup) SetPrivateWirelessGatewayId(v string) {
	o.PrivateWirelessGatewayId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SIMCardGroup) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SIMCardGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SIMCardGroup) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SIMCardGroup) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroup) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SIMCardGroup) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SIMCardGroup) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o SIMCardGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIMCardGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DataLimit) {
		toSerialize["data_limit"] = o.DataLimit
	}
	if !IsNil(o.ConsumedData) {
		toSerialize["consumed_data"] = o.ConsumedData
	}
	if !IsNil(o.PrivateWirelessGatewayId) {
		toSerialize["private_wireless_gateway_id"] = o.PrivateWirelessGatewayId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableSIMCardGroup struct {
	value *SIMCardGroup
	isSet bool
}

func (v NullableSIMCardGroup) Get() *SIMCardGroup {
	return v.value
}

func (v *NullableSIMCardGroup) Set(val *SIMCardGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSIMCardGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSIMCardGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIMCardGroup(val *SIMCardGroup) *NullableSIMCardGroup {
	return &NullableSIMCardGroup{value: val, isSet: true}
}

func (v NullableSIMCardGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIMCardGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


