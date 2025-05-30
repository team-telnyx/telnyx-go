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

// checks if the ConferenceResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceResource{}

// ConferenceResource struct for ConferenceResource
type ConferenceResource struct {
	// The id of the account the resource belongs to.
	AccountSid *string `json:"account_sid,omitempty"`
	// The version of the API that was used to make the request.
	ApiVersion *string `json:"api_version,omitempty"`
	// Caller ID, if present.
	CallSidEndingConference *string `json:"call_sid_ending_conference,omitempty"`
	// The timestamp of when the resource was created.
	DateCreated *string `json:"date_created,omitempty"`
	// The timestamp of when the resource was last updated.
	DateUpdated *string `json:"date_updated,omitempty"`
	// A string that you assigned to describe this conference room.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The reason why a conference ended. When a conference is in progress, will be null.
	ReasonConferenceEnded *string `json:"reason_conference_ended,omitempty"`
	// A string representing the region where the conference is hosted.
	Region *string `json:"region,omitempty"`
	// The unique identifier of the conference.
	Sid *string `json:"sid,omitempty"`
	// The status of this conference.
	Status *string `json:"status,omitempty"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
	// The relative URI for this conference.
	Uri *string `json:"uri,omitempty"`
}

// NewConferenceResource instantiates a new ConferenceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceResource() *ConferenceResource {
	this := ConferenceResource{}
	return &this
}

// NewConferenceResourceWithDefaults instantiates a new ConferenceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceResourceWithDefaults() *ConferenceResource {
	this := ConferenceResource{}
	return &this
}

// GetAccountSid returns the AccountSid field value if set, zero value otherwise.
func (o *ConferenceResource) GetAccountSid() string {
	if o == nil || IsNil(o.AccountSid) {
		var ret string
		return ret
	}
	return *o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetAccountSidOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSid) {
		return nil, false
	}
	return o.AccountSid, true
}

// HasAccountSid returns a boolean if a field has been set.
func (o *ConferenceResource) HasAccountSid() bool {
	if o != nil && !IsNil(o.AccountSid) {
		return true
	}

	return false
}

// SetAccountSid gets a reference to the given string and assigns it to the AccountSid field.
func (o *ConferenceResource) SetAccountSid(v string) {
	o.AccountSid = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ConferenceResource) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ConferenceResource) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ConferenceResource) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCallSidEndingConference returns the CallSidEndingConference field value if set, zero value otherwise.
func (o *ConferenceResource) GetCallSidEndingConference() string {
	if o == nil || IsNil(o.CallSidEndingConference) {
		var ret string
		return ret
	}
	return *o.CallSidEndingConference
}

// GetCallSidEndingConferenceOk returns a tuple with the CallSidEndingConference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetCallSidEndingConferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CallSidEndingConference) {
		return nil, false
	}
	return o.CallSidEndingConference, true
}

// HasCallSidEndingConference returns a boolean if a field has been set.
func (o *ConferenceResource) HasCallSidEndingConference() bool {
	if o != nil && !IsNil(o.CallSidEndingConference) {
		return true
	}

	return false
}

// SetCallSidEndingConference gets a reference to the given string and assigns it to the CallSidEndingConference field.
func (o *ConferenceResource) SetCallSidEndingConference(v string) {
	o.CallSidEndingConference = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ConferenceResource) GetDateCreated() string {
	if o == nil || IsNil(o.DateCreated) {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetDateCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ConferenceResource) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *ConferenceResource) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *ConferenceResource) GetDateUpdated() string {
	if o == nil || IsNil(o.DateUpdated) {
		var ret string
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetDateUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.DateUpdated) {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *ConferenceResource) HasDateUpdated() bool {
	if o != nil && !IsNil(o.DateUpdated) {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given string and assigns it to the DateUpdated field.
func (o *ConferenceResource) SetDateUpdated(v string) {
	o.DateUpdated = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *ConferenceResource) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *ConferenceResource) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *ConferenceResource) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetReasonConferenceEnded returns the ReasonConferenceEnded field value if set, zero value otherwise.
func (o *ConferenceResource) GetReasonConferenceEnded() string {
	if o == nil || IsNil(o.ReasonConferenceEnded) {
		var ret string
		return ret
	}
	return *o.ReasonConferenceEnded
}

// GetReasonConferenceEndedOk returns a tuple with the ReasonConferenceEnded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetReasonConferenceEndedOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonConferenceEnded) {
		return nil, false
	}
	return o.ReasonConferenceEnded, true
}

// HasReasonConferenceEnded returns a boolean if a field has been set.
func (o *ConferenceResource) HasReasonConferenceEnded() bool {
	if o != nil && !IsNil(o.ReasonConferenceEnded) {
		return true
	}

	return false
}

// SetReasonConferenceEnded gets a reference to the given string and assigns it to the ReasonConferenceEnded field.
func (o *ConferenceResource) SetReasonConferenceEnded(v string) {
	o.ReasonConferenceEnded = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ConferenceResource) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ConferenceResource) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ConferenceResource) SetRegion(v string) {
	o.Region = &v
}

// GetSid returns the Sid field value if set, zero value otherwise.
func (o *ConferenceResource) GetSid() string {
	if o == nil || IsNil(o.Sid) {
		var ret string
		return ret
	}
	return *o.Sid
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetSidOk() (*string, bool) {
	if o == nil || IsNil(o.Sid) {
		return nil, false
	}
	return o.Sid, true
}

// HasSid returns a boolean if a field has been set.
func (o *ConferenceResource) HasSid() bool {
	if o != nil && !IsNil(o.Sid) {
		return true
	}

	return false
}

// SetSid gets a reference to the given string and assigns it to the Sid field.
func (o *ConferenceResource) SetSid(v string) {
	o.Sid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConferenceResource) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConferenceResource) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConferenceResource) SetStatus(v string) {
	o.Status = &v
}

// GetSubresourceUris returns the SubresourceUris field value if set, zero value otherwise.
func (o *ConferenceResource) GetSubresourceUris() map[string]interface{} {
	if o == nil || IsNil(o.SubresourceUris) {
		var ret map[string]interface{}
		return ret
	}
	return o.SubresourceUris
}

// GetSubresourceUrisOk returns a tuple with the SubresourceUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetSubresourceUrisOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SubresourceUris) {
		return map[string]interface{}{}, false
	}
	return o.SubresourceUris, true
}

// HasSubresourceUris returns a boolean if a field has been set.
func (o *ConferenceResource) HasSubresourceUris() bool {
	if o != nil && !IsNil(o.SubresourceUris) {
		return true
	}

	return false
}

// SetSubresourceUris gets a reference to the given map[string]interface{} and assigns it to the SubresourceUris field.
func (o *ConferenceResource) SetSubresourceUris(v map[string]interface{}) {
	o.SubresourceUris = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ConferenceResource) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceResource) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ConferenceResource) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ConferenceResource) SetUri(v string) {
	o.Uri = &v
}

func (o ConferenceResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountSid) {
		toSerialize["account_sid"] = o.AccountSid
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !IsNil(o.CallSidEndingConference) {
		toSerialize["call_sid_ending_conference"] = o.CallSidEndingConference
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.DateUpdated) {
		toSerialize["date_updated"] = o.DateUpdated
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendly_name"] = o.FriendlyName
	}
	if !IsNil(o.ReasonConferenceEnded) {
		toSerialize["reason_conference_ended"] = o.ReasonConferenceEnded
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Sid) {
		toSerialize["sid"] = o.Sid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubresourceUris) {
		toSerialize["subresource_uris"] = o.SubresourceUris
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableConferenceResource struct {
	value *ConferenceResource
	isSet bool
}

func (v NullableConferenceResource) Get() *ConferenceResource {
	return v.value
}

func (v *NullableConferenceResource) Set(val *ConferenceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceResource(val *ConferenceResource) *NullableConferenceResource {
	return &NullableConferenceResource{value: val, isSet: true}
}

func (v NullableConferenceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


