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

// checks if the GlobalIPAllowedPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalIPAllowedPort{}

// GlobalIPAllowedPort struct for GlobalIPAllowedPort
type GlobalIPAllowedPort struct {
	// Identifies the resource.
	Id *string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// The Global IP Protocol code.
	ProtocolCode *string `json:"protocol_code,omitempty"`
	// A name for the Global IP ports range.
	Name *string `json:"name,omitempty"`
	// First port of a range.
	FirstPort *int32 `json:"first_port,omitempty"`
	// Last port of a range.
	LastPort *int32 `json:"last_port,omitempty"`
}

// NewGlobalIPAllowedPort instantiates a new GlobalIPAllowedPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalIPAllowedPort() *GlobalIPAllowedPort {
	this := GlobalIPAllowedPort{}
	return &this
}

// NewGlobalIPAllowedPortWithDefaults instantiates a new GlobalIPAllowedPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalIPAllowedPortWithDefaults() *GlobalIPAllowedPort {
	this := GlobalIPAllowedPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GlobalIPAllowedPort) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *GlobalIPAllowedPort) SetRecordType(v string) {
	o.RecordType = &v
}

// GetProtocolCode returns the ProtocolCode field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetProtocolCode() string {
	if o == nil || IsNil(o.ProtocolCode) {
		var ret string
		return ret
	}
	return *o.ProtocolCode
}

// GetProtocolCodeOk returns a tuple with the ProtocolCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetProtocolCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolCode) {
		return nil, false
	}
	return o.ProtocolCode, true
}

// HasProtocolCode returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasProtocolCode() bool {
	if o != nil && !IsNil(o.ProtocolCode) {
		return true
	}

	return false
}

// SetProtocolCode gets a reference to the given string and assigns it to the ProtocolCode field.
func (o *GlobalIPAllowedPort) SetProtocolCode(v string) {
	o.ProtocolCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalIPAllowedPort) SetName(v string) {
	o.Name = &v
}

// GetFirstPort returns the FirstPort field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetFirstPort() int32 {
	if o == nil || IsNil(o.FirstPort) {
		var ret int32
		return ret
	}
	return *o.FirstPort
}

// GetFirstPortOk returns a tuple with the FirstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetFirstPortOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstPort) {
		return nil, false
	}
	return o.FirstPort, true
}

// HasFirstPort returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasFirstPort() bool {
	if o != nil && !IsNil(o.FirstPort) {
		return true
	}

	return false
}

// SetFirstPort gets a reference to the given int32 and assigns it to the FirstPort field.
func (o *GlobalIPAllowedPort) SetFirstPort(v int32) {
	o.FirstPort = &v
}

// GetLastPort returns the LastPort field value if set, zero value otherwise.
func (o *GlobalIPAllowedPort) GetLastPort() int32 {
	if o == nil || IsNil(o.LastPort) {
		var ret int32
		return ret
	}
	return *o.LastPort
}

// GetLastPortOk returns a tuple with the LastPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIPAllowedPort) GetLastPortOk() (*int32, bool) {
	if o == nil || IsNil(o.LastPort) {
		return nil, false
	}
	return o.LastPort, true
}

// HasLastPort returns a boolean if a field has been set.
func (o *GlobalIPAllowedPort) HasLastPort() bool {
	if o != nil && !IsNil(o.LastPort) {
		return true
	}

	return false
}

// SetLastPort gets a reference to the given int32 and assigns it to the LastPort field.
func (o *GlobalIPAllowedPort) SetLastPort(v int32) {
	o.LastPort = &v
}

func (o GlobalIPAllowedPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalIPAllowedPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.ProtocolCode) {
		toSerialize["protocol_code"] = o.ProtocolCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FirstPort) {
		toSerialize["first_port"] = o.FirstPort
	}
	if !IsNil(o.LastPort) {
		toSerialize["last_port"] = o.LastPort
	}
	return toSerialize, nil
}

type NullableGlobalIPAllowedPort struct {
	value *GlobalIPAllowedPort
	isSet bool
}

func (v NullableGlobalIPAllowedPort) Get() *GlobalIPAllowedPort {
	return v.value
}

func (v *NullableGlobalIPAllowedPort) Set(val *GlobalIPAllowedPort) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalIPAllowedPort) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalIPAllowedPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalIPAllowedPort(val *GlobalIPAllowedPort) *NullableGlobalIPAllowedPort {
	return &NullableGlobalIPAllowedPort{value: val, isSet: true}
}

func (v NullableGlobalIPAllowedPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalIPAllowedPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


