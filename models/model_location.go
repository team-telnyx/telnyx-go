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

// checks if the Location type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Location{}

// Location struct for Location
type Location struct {
	// Identifies the geographical region of location.
	Region *string `json:"region,omitempty"`
	// Site of location.
	Site *string `json:"site,omitempty"`
	// Point of presence of location.
	Pop *string `json:"pop,omitempty"`
	// Location code.
	Code *string `json:"code,omitempty"`
	// Human readable name of location.
	Name *string `json:"name,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Location) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Location) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Location) SetRegion(v string) {
	o.Region = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *Location) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *Location) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *Location) SetSite(v string) {
	o.Site = &v
}

// GetPop returns the Pop field value if set, zero value otherwise.
func (o *Location) GetPop() string {
	if o == nil || IsNil(o.Pop) {
		var ret string
		return ret
	}
	return *o.Pop
}

// GetPopOk returns a tuple with the Pop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetPopOk() (*string, bool) {
	if o == nil || IsNil(o.Pop) {
		return nil, false
	}
	return o.Pop, true
}

// HasPop returns a boolean if a field has been set.
func (o *Location) HasPop() bool {
	if o != nil && !IsNil(o.Pop) {
		return true
	}

	return false
}

// SetPop gets a reference to the given string and assigns it to the Pop field.
func (o *Location) SetPop(v string) {
	o.Pop = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Location) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Location) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Location) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Location) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Location) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Location) SetName(v string) {
	o.Name = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Location) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Pop) {
		toSerialize["pop"] = o.Pop
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


