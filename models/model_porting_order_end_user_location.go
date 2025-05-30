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

// checks if the PortingOrderEndUserLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortingOrderEndUserLocation{}

// PortingOrderEndUserLocation struct for PortingOrderEndUserLocation
type PortingOrderEndUserLocation struct {
	// First line of billing address
	StreetAddress *string `json:"street_address,omitempty"`
	// Second line of billing address
	ExtendedAddress *string `json:"extended_address,omitempty"`
	// City or municipality of billing address
	Locality *string `json:"locality,omitempty"`
	// State, province, or similar of billing address
	AdministrativeArea *string `json:"administrative_area,omitempty"`
	// Postal Code of billing address
	PostalCode *string `json:"postal_code,omitempty"`
	// ISO3166-1 alpha-2 country code of billing address
	CountryCode *string `json:"country_code,omitempty"`
}

// NewPortingOrderEndUserLocation instantiates a new PortingOrderEndUserLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortingOrderEndUserLocation() *PortingOrderEndUserLocation {
	this := PortingOrderEndUserLocation{}
	return &this
}

// NewPortingOrderEndUserLocationWithDefaults instantiates a new PortingOrderEndUserLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortingOrderEndUserLocationWithDefaults() *PortingOrderEndUserLocation {
	this := PortingOrderEndUserLocation{}
	return &this
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *PortingOrderEndUserLocation) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetExtendedAddress returns the ExtendedAddress field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetExtendedAddress() string {
	if o == nil || IsNil(o.ExtendedAddress) {
		var ret string
		return ret
	}
	return *o.ExtendedAddress
}

// GetExtendedAddressOk returns a tuple with the ExtendedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetExtendedAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedAddress) {
		return nil, false
	}
	return o.ExtendedAddress, true
}

// HasExtendedAddress returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasExtendedAddress() bool {
	if o != nil && !IsNil(o.ExtendedAddress) {
		return true
	}

	return false
}

// SetExtendedAddress gets a reference to the given string and assigns it to the ExtendedAddress field.
func (o *PortingOrderEndUserLocation) SetExtendedAddress(v string) {
	o.ExtendedAddress = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *PortingOrderEndUserLocation) SetLocality(v string) {
	o.Locality = &v
}

// GetAdministrativeArea returns the AdministrativeArea field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetAdministrativeArea() string {
	if o == nil || IsNil(o.AdministrativeArea) {
		var ret string
		return ret
	}
	return *o.AdministrativeArea
}

// GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetAdministrativeAreaOk() (*string, bool) {
	if o == nil || IsNil(o.AdministrativeArea) {
		return nil, false
	}
	return o.AdministrativeArea, true
}

// HasAdministrativeArea returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasAdministrativeArea() bool {
	if o != nil && !IsNil(o.AdministrativeArea) {
		return true
	}

	return false
}

// SetAdministrativeArea gets a reference to the given string and assigns it to the AdministrativeArea field.
func (o *PortingOrderEndUserLocation) SetAdministrativeArea(v string) {
	o.AdministrativeArea = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PortingOrderEndUserLocation) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PortingOrderEndUserLocation) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrderEndUserLocation) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PortingOrderEndUserLocation) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PortingOrderEndUserLocation) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o PortingOrderEndUserLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortingOrderEndUserLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.ExtendedAddress) {
		toSerialize["extended_address"] = o.ExtendedAddress
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.AdministrativeArea) {
		toSerialize["administrative_area"] = o.AdministrativeArea
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullablePortingOrderEndUserLocation struct {
	value *PortingOrderEndUserLocation
	isSet bool
}

func (v NullablePortingOrderEndUserLocation) Get() *PortingOrderEndUserLocation {
	return v.value
}

func (v *NullablePortingOrderEndUserLocation) Set(val *PortingOrderEndUserLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePortingOrderEndUserLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePortingOrderEndUserLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortingOrderEndUserLocation(val *PortingOrderEndUserLocation) *NullablePortingOrderEndUserLocation {
	return &NullablePortingOrderEndUserLocation{value: val, isSet: true}
}

func (v NullablePortingOrderEndUserLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortingOrderEndUserLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


