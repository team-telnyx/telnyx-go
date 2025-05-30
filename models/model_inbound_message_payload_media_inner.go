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

// checks if the InboundMessagePayloadMediaInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundMessagePayloadMediaInner{}

// InboundMessagePayloadMediaInner struct for InboundMessagePayloadMediaInner
type InboundMessagePayloadMediaInner struct {
	// The url of the media requested to be sent.
	Url *string `json:"url,omitempty"`
	// The MIME type of the requested media.
	ContentType *string `json:"content_type,omitempty"`
	// The size of the requested media.
	Size *int32 `json:"size,omitempty"`
	// The SHA256 hash of the requested media.
	HashSha256 *string `json:"hash_sha256,omitempty"`
}

// NewInboundMessagePayloadMediaInner instantiates a new InboundMessagePayloadMediaInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundMessagePayloadMediaInner() *InboundMessagePayloadMediaInner {
	this := InboundMessagePayloadMediaInner{}
	return &this
}

// NewInboundMessagePayloadMediaInnerWithDefaults instantiates a new InboundMessagePayloadMediaInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundMessagePayloadMediaInnerWithDefaults() *InboundMessagePayloadMediaInner {
	this := InboundMessagePayloadMediaInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InboundMessagePayloadMediaInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessagePayloadMediaInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InboundMessagePayloadMediaInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InboundMessagePayloadMediaInner) SetUrl(v string) {
	o.Url = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *InboundMessagePayloadMediaInner) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessagePayloadMediaInner) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *InboundMessagePayloadMediaInner) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *InboundMessagePayloadMediaInner) SetContentType(v string) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InboundMessagePayloadMediaInner) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessagePayloadMediaInner) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InboundMessagePayloadMediaInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *InboundMessagePayloadMediaInner) SetSize(v int32) {
	o.Size = &v
}

// GetHashSha256 returns the HashSha256 field value if set, zero value otherwise.
func (o *InboundMessagePayloadMediaInner) GetHashSha256() string {
	if o == nil || IsNil(o.HashSha256) {
		var ret string
		return ret
	}
	return *o.HashSha256
}

// GetHashSha256Ok returns a tuple with the HashSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundMessagePayloadMediaInner) GetHashSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.HashSha256) {
		return nil, false
	}
	return o.HashSha256, true
}

// HasHashSha256 returns a boolean if a field has been set.
func (o *InboundMessagePayloadMediaInner) HasHashSha256() bool {
	if o != nil && !IsNil(o.HashSha256) {
		return true
	}

	return false
}

// SetHashSha256 gets a reference to the given string and assigns it to the HashSha256 field.
func (o *InboundMessagePayloadMediaInner) SetHashSha256(v string) {
	o.HashSha256 = &v
}

func (o InboundMessagePayloadMediaInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundMessagePayloadMediaInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.HashSha256) {
		toSerialize["hash_sha256"] = o.HashSha256
	}
	return toSerialize, nil
}

type NullableInboundMessagePayloadMediaInner struct {
	value *InboundMessagePayloadMediaInner
	isSet bool
}

func (v NullableInboundMessagePayloadMediaInner) Get() *InboundMessagePayloadMediaInner {
	return v.value
}

func (v *NullableInboundMessagePayloadMediaInner) Set(val *InboundMessagePayloadMediaInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundMessagePayloadMediaInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundMessagePayloadMediaInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundMessagePayloadMediaInner(val *InboundMessagePayloadMediaInner) *NullableInboundMessagePayloadMediaInner {
	return &NullableInboundMessagePayloadMediaInner{value: val, isSet: true}
}

func (v NullableInboundMessagePayloadMediaInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundMessagePayloadMediaInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


