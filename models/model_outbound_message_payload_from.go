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

// checks if the OutboundMessagePayloadFrom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutboundMessagePayloadFrom{}

// OutboundMessagePayloadFrom struct for OutboundMessagePayloadFrom
type OutboundMessagePayloadFrom struct {
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The carrier of the receiver.
	Carrier *string `json:"carrier,omitempty"`
	// The line-type of the receiver.
	LineType *string `json:"line_type,omitempty"`
}

// NewOutboundMessagePayloadFrom instantiates a new OutboundMessagePayloadFrom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundMessagePayloadFrom() *OutboundMessagePayloadFrom {
	this := OutboundMessagePayloadFrom{}
	return &this
}

// NewOutboundMessagePayloadFromWithDefaults instantiates a new OutboundMessagePayloadFrom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundMessagePayloadFromWithDefaults() *OutboundMessagePayloadFrom {
	this := OutboundMessagePayloadFrom{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OutboundMessagePayloadFrom) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundMessagePayloadFrom) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OutboundMessagePayloadFrom) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OutboundMessagePayloadFrom) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *OutboundMessagePayloadFrom) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundMessagePayloadFrom) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *OutboundMessagePayloadFrom) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *OutboundMessagePayloadFrom) SetCarrier(v string) {
	o.Carrier = &v
}

// GetLineType returns the LineType field value if set, zero value otherwise.
func (o *OutboundMessagePayloadFrom) GetLineType() string {
	if o == nil || IsNil(o.LineType) {
		var ret string
		return ret
	}
	return *o.LineType
}

// GetLineTypeOk returns a tuple with the LineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundMessagePayloadFrom) GetLineTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LineType) {
		return nil, false
	}
	return o.LineType, true
}

// HasLineType returns a boolean if a field has been set.
func (o *OutboundMessagePayloadFrom) HasLineType() bool {
	if o != nil && !IsNil(o.LineType) {
		return true
	}

	return false
}

// SetLineType gets a reference to the given string and assigns it to the LineType field.
func (o *OutboundMessagePayloadFrom) SetLineType(v string) {
	o.LineType = &v
}

func (o OutboundMessagePayloadFrom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundMessagePayloadFrom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.LineType) {
		toSerialize["line_type"] = o.LineType
	}
	return toSerialize, nil
}

type NullableOutboundMessagePayloadFrom struct {
	value *OutboundMessagePayloadFrom
	isSet bool
}

func (v NullableOutboundMessagePayloadFrom) Get() *OutboundMessagePayloadFrom {
	return v.value
}

func (v *NullableOutboundMessagePayloadFrom) Set(val *OutboundMessagePayloadFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundMessagePayloadFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundMessagePayloadFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundMessagePayloadFrom(val *OutboundMessagePayloadFrom) *NullableOutboundMessagePayloadFrom {
	return &NullableOutboundMessagePayloadFrom{value: val, isSet: true}
}

func (v NullableOutboundMessagePayloadFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundMessagePayloadFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


