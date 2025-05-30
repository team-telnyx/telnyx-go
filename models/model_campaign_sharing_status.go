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

// checks if the CampaignSharingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignSharingStatus{}

// CampaignSharingStatus struct for CampaignSharingStatus
type CampaignSharingStatus struct {
	DownstreamCnpId *string `json:"downstreamCnpId,omitempty"`
	SharedDate *string `json:"sharedDate,omitempty"`
	SharingStatus *string `json:"sharingStatus,omitempty"`
	StatusDate *string `json:"statusDate,omitempty"`
	UpstreamCnpId *string `json:"upstreamCnpId,omitempty"`
}

// NewCampaignSharingStatus instantiates a new CampaignSharingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSharingStatus() *CampaignSharingStatus {
	this := CampaignSharingStatus{}
	return &this
}

// NewCampaignSharingStatusWithDefaults instantiates a new CampaignSharingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSharingStatusWithDefaults() *CampaignSharingStatus {
	this := CampaignSharingStatus{}
	return &this
}

// GetDownstreamCnpId returns the DownstreamCnpId field value if set, zero value otherwise.
func (o *CampaignSharingStatus) GetDownstreamCnpId() string {
	if o == nil || IsNil(o.DownstreamCnpId) {
		var ret string
		return ret
	}
	return *o.DownstreamCnpId
}

// GetDownstreamCnpIdOk returns a tuple with the DownstreamCnpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSharingStatus) GetDownstreamCnpIdOk() (*string, bool) {
	if o == nil || IsNil(o.DownstreamCnpId) {
		return nil, false
	}
	return o.DownstreamCnpId, true
}

// HasDownstreamCnpId returns a boolean if a field has been set.
func (o *CampaignSharingStatus) HasDownstreamCnpId() bool {
	if o != nil && !IsNil(o.DownstreamCnpId) {
		return true
	}

	return false
}

// SetDownstreamCnpId gets a reference to the given string and assigns it to the DownstreamCnpId field.
func (o *CampaignSharingStatus) SetDownstreamCnpId(v string) {
	o.DownstreamCnpId = &v
}

// GetSharedDate returns the SharedDate field value if set, zero value otherwise.
func (o *CampaignSharingStatus) GetSharedDate() string {
	if o == nil || IsNil(o.SharedDate) {
		var ret string
		return ret
	}
	return *o.SharedDate
}

// GetSharedDateOk returns a tuple with the SharedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSharingStatus) GetSharedDateOk() (*string, bool) {
	if o == nil || IsNil(o.SharedDate) {
		return nil, false
	}
	return o.SharedDate, true
}

// HasSharedDate returns a boolean if a field has been set.
func (o *CampaignSharingStatus) HasSharedDate() bool {
	if o != nil && !IsNil(o.SharedDate) {
		return true
	}

	return false
}

// SetSharedDate gets a reference to the given string and assigns it to the SharedDate field.
func (o *CampaignSharingStatus) SetSharedDate(v string) {
	o.SharedDate = &v
}

// GetSharingStatus returns the SharingStatus field value if set, zero value otherwise.
func (o *CampaignSharingStatus) GetSharingStatus() string {
	if o == nil || IsNil(o.SharingStatus) {
		var ret string
		return ret
	}
	return *o.SharingStatus
}

// GetSharingStatusOk returns a tuple with the SharingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSharingStatus) GetSharingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SharingStatus) {
		return nil, false
	}
	return o.SharingStatus, true
}

// HasSharingStatus returns a boolean if a field has been set.
func (o *CampaignSharingStatus) HasSharingStatus() bool {
	if o != nil && !IsNil(o.SharingStatus) {
		return true
	}

	return false
}

// SetSharingStatus gets a reference to the given string and assigns it to the SharingStatus field.
func (o *CampaignSharingStatus) SetSharingStatus(v string) {
	o.SharingStatus = &v
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise.
func (o *CampaignSharingStatus) GetStatusDate() string {
	if o == nil || IsNil(o.StatusDate) {
		var ret string
		return ret
	}
	return *o.StatusDate
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSharingStatus) GetStatusDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDate) {
		return nil, false
	}
	return o.StatusDate, true
}

// HasStatusDate returns a boolean if a field has been set.
func (o *CampaignSharingStatus) HasStatusDate() bool {
	if o != nil && !IsNil(o.StatusDate) {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given string and assigns it to the StatusDate field.
func (o *CampaignSharingStatus) SetStatusDate(v string) {
	o.StatusDate = &v
}

// GetUpstreamCnpId returns the UpstreamCnpId field value if set, zero value otherwise.
func (o *CampaignSharingStatus) GetUpstreamCnpId() string {
	if o == nil || IsNil(o.UpstreamCnpId) {
		var ret string
		return ret
	}
	return *o.UpstreamCnpId
}

// GetUpstreamCnpIdOk returns a tuple with the UpstreamCnpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSharingStatus) GetUpstreamCnpIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamCnpId) {
		return nil, false
	}
	return o.UpstreamCnpId, true
}

// HasUpstreamCnpId returns a boolean if a field has been set.
func (o *CampaignSharingStatus) HasUpstreamCnpId() bool {
	if o != nil && !IsNil(o.UpstreamCnpId) {
		return true
	}

	return false
}

// SetUpstreamCnpId gets a reference to the given string and assigns it to the UpstreamCnpId field.
func (o *CampaignSharingStatus) SetUpstreamCnpId(v string) {
	o.UpstreamCnpId = &v
}

func (o CampaignSharingStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignSharingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownstreamCnpId) {
		toSerialize["downstreamCnpId"] = o.DownstreamCnpId
	}
	if !IsNil(o.SharedDate) {
		toSerialize["sharedDate"] = o.SharedDate
	}
	if !IsNil(o.SharingStatus) {
		toSerialize["sharingStatus"] = o.SharingStatus
	}
	if !IsNil(o.StatusDate) {
		toSerialize["statusDate"] = o.StatusDate
	}
	if !IsNil(o.UpstreamCnpId) {
		toSerialize["upstreamCnpId"] = o.UpstreamCnpId
	}
	return toSerialize, nil
}

type NullableCampaignSharingStatus struct {
	value *CampaignSharingStatus
	isSet bool
}

func (v NullableCampaignSharingStatus) Get() *CampaignSharingStatus {
	return v.value
}

func (v *NullableCampaignSharingStatus) Set(val *CampaignSharingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSharingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSharingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSharingStatus(val *CampaignSharingStatus) *NullableCampaignSharingStatus {
	return &NullableCampaignSharingStatus{value: val, isSet: true}
}

func (v NullableCampaignSharingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSharingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


