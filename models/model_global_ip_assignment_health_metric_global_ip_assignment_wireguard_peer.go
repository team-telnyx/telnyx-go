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

// checks if the GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer{}

// GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer struct for GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer
type GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer struct {
	// A user specified name for the interface.
	Name *string `json:"name,omitempty"`
	// The IP address of the interface.
	IpAddress *string `json:"ip_address,omitempty"`
}

// NewGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer instantiates a new GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer() *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer {
	this := GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer{}
	return &this
}

// NewGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeerWithDefaults instantiates a new GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeerWithDefaults() *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer {
	this := GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) SetName(v string) {
	o.Name = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	return toSerialize, nil
}

type NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer struct {
	value *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer
	isSet bool
}

func (v NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) Get() *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer {
	return v.value
}

func (v *NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) Set(val *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer(val *GlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) *NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer {
	return &NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer{value: val, isSet: true}
}

func (v NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalIpAssignmentHealthMetricGlobalIpAssignmentWireguardPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


