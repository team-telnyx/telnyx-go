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

// checks if the RecordingResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordingResponseData{}

// RecordingResponseData struct for RecordingResponseData
type RecordingResponseData struct {
	// Unique identifier and token for controlling the call.
	CallControlId *string `json:"call_control_id,omitempty"`
	// ID unique to the call leg (used to correlate webhook events).
	CallLegId *string `json:"call_leg_id,omitempty"`
	// ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionId *string `json:"call_session_id,omitempty"`
	// When `dual`, the final audio file has the first leg on channel A, and the rest on channel B.
	Channels *string `json:"channels,omitempty"`
	// Uniquely identifies the conference.
	ConferenceId *string `json:"conference_id,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	DownloadUrls *RecordingResponseDataDownloadUrls `json:"download_urls,omitempty"`
	// The duration of the recording in milliseconds.
	DurationMillis *int32 `json:"duration_millis,omitempty"`
	// Uniquely identifies the recording.
	Id *string `json:"id,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	// ISO 8601 formatted date of when the recording started.
	RecordingStartedAt *string `json:"recording_started_at,omitempty"`
	// ISO 8601 formatted date of when the recording ended.
	RecordingEndedAt *string `json:"recording_ended_at,omitempty"`
	// The kind of event that led to this recording being created.
	Source *string `json:"source,omitempty"`
	// The status of the recording. Only `completed` recordings are currently supported.
	Status *string `json:"status,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewRecordingResponseData instantiates a new RecordingResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordingResponseData() *RecordingResponseData {
	this := RecordingResponseData{}
	return &this
}

// NewRecordingResponseDataWithDefaults instantiates a new RecordingResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordingResponseDataWithDefaults() *RecordingResponseData {
	this := RecordingResponseData{}
	return &this
}

// GetCallControlId returns the CallControlId field value if set, zero value otherwise.
func (o *RecordingResponseData) GetCallControlId() string {
	if o == nil || IsNil(o.CallControlId) {
		var ret string
		return ret
	}
	return *o.CallControlId
}

// GetCallControlIdOk returns a tuple with the CallControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetCallControlIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallControlId) {
		return nil, false
	}
	return o.CallControlId, true
}

// HasCallControlId returns a boolean if a field has been set.
func (o *RecordingResponseData) HasCallControlId() bool {
	if o != nil && !IsNil(o.CallControlId) {
		return true
	}

	return false
}

// SetCallControlId gets a reference to the given string and assigns it to the CallControlId field.
func (o *RecordingResponseData) SetCallControlId(v string) {
	o.CallControlId = &v
}

// GetCallLegId returns the CallLegId field value if set, zero value otherwise.
func (o *RecordingResponseData) GetCallLegId() string {
	if o == nil || IsNil(o.CallLegId) {
		var ret string
		return ret
	}
	return *o.CallLegId
}

// GetCallLegIdOk returns a tuple with the CallLegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetCallLegIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallLegId) {
		return nil, false
	}
	return o.CallLegId, true
}

// HasCallLegId returns a boolean if a field has been set.
func (o *RecordingResponseData) HasCallLegId() bool {
	if o != nil && !IsNil(o.CallLegId) {
		return true
	}

	return false
}

// SetCallLegId gets a reference to the given string and assigns it to the CallLegId field.
func (o *RecordingResponseData) SetCallLegId(v string) {
	o.CallLegId = &v
}

// GetCallSessionId returns the CallSessionId field value if set, zero value otherwise.
func (o *RecordingResponseData) GetCallSessionId() string {
	if o == nil || IsNil(o.CallSessionId) {
		var ret string
		return ret
	}
	return *o.CallSessionId
}

// GetCallSessionIdOk returns a tuple with the CallSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetCallSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallSessionId) {
		return nil, false
	}
	return o.CallSessionId, true
}

// HasCallSessionId returns a boolean if a field has been set.
func (o *RecordingResponseData) HasCallSessionId() bool {
	if o != nil && !IsNil(o.CallSessionId) {
		return true
	}

	return false
}

// SetCallSessionId gets a reference to the given string and assigns it to the CallSessionId field.
func (o *RecordingResponseData) SetCallSessionId(v string) {
	o.CallSessionId = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *RecordingResponseData) GetChannels() string {
	if o == nil || IsNil(o.Channels) {
		var ret string
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetChannelsOk() (*string, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *RecordingResponseData) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given string and assigns it to the Channels field.
func (o *RecordingResponseData) SetChannels(v string) {
	o.Channels = &v
}

// GetConferenceId returns the ConferenceId field value if set, zero value otherwise.
func (o *RecordingResponseData) GetConferenceId() string {
	if o == nil || IsNil(o.ConferenceId) {
		var ret string
		return ret
	}
	return *o.ConferenceId
}

// GetConferenceIdOk returns a tuple with the ConferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetConferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConferenceId) {
		return nil, false
	}
	return o.ConferenceId, true
}

// HasConferenceId returns a boolean if a field has been set.
func (o *RecordingResponseData) HasConferenceId() bool {
	if o != nil && !IsNil(o.ConferenceId) {
		return true
	}

	return false
}

// SetConferenceId gets a reference to the given string and assigns it to the ConferenceId field.
func (o *RecordingResponseData) SetConferenceId(v string) {
	o.ConferenceId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RecordingResponseData) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RecordingResponseData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RecordingResponseData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDownloadUrls returns the DownloadUrls field value if set, zero value otherwise.
func (o *RecordingResponseData) GetDownloadUrls() RecordingResponseDataDownloadUrls {
	if o == nil || IsNil(o.DownloadUrls) {
		var ret RecordingResponseDataDownloadUrls
		return ret
	}
	return *o.DownloadUrls
}

// GetDownloadUrlsOk returns a tuple with the DownloadUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetDownloadUrlsOk() (*RecordingResponseDataDownloadUrls, bool) {
	if o == nil || IsNil(o.DownloadUrls) {
		return nil, false
	}
	return o.DownloadUrls, true
}

// HasDownloadUrls returns a boolean if a field has been set.
func (o *RecordingResponseData) HasDownloadUrls() bool {
	if o != nil && !IsNil(o.DownloadUrls) {
		return true
	}

	return false
}

// SetDownloadUrls gets a reference to the given RecordingResponseDataDownloadUrls and assigns it to the DownloadUrls field.
func (o *RecordingResponseData) SetDownloadUrls(v RecordingResponseDataDownloadUrls) {
	o.DownloadUrls = &v
}

// GetDurationMillis returns the DurationMillis field value if set, zero value otherwise.
func (o *RecordingResponseData) GetDurationMillis() int32 {
	if o == nil || IsNil(o.DurationMillis) {
		var ret int32
		return ret
	}
	return *o.DurationMillis
}

// GetDurationMillisOk returns a tuple with the DurationMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetDurationMillisOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMillis) {
		return nil, false
	}
	return o.DurationMillis, true
}

// HasDurationMillis returns a boolean if a field has been set.
func (o *RecordingResponseData) HasDurationMillis() bool {
	if o != nil && !IsNil(o.DurationMillis) {
		return true
	}

	return false
}

// SetDurationMillis gets a reference to the given int32 and assigns it to the DurationMillis field.
func (o *RecordingResponseData) SetDurationMillis(v int32) {
	o.DurationMillis = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecordingResponseData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecordingResponseData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecordingResponseData) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *RecordingResponseData) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *RecordingResponseData) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *RecordingResponseData) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordingStartedAt returns the RecordingStartedAt field value if set, zero value otherwise.
func (o *RecordingResponseData) GetRecordingStartedAt() string {
	if o == nil || IsNil(o.RecordingStartedAt) {
		var ret string
		return ret
	}
	return *o.RecordingStartedAt
}

// GetRecordingStartedAtOk returns a tuple with the RecordingStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetRecordingStartedAtOk() (*string, bool) {
	if o == nil || IsNil(o.RecordingStartedAt) {
		return nil, false
	}
	return o.RecordingStartedAt, true
}

// HasRecordingStartedAt returns a boolean if a field has been set.
func (o *RecordingResponseData) HasRecordingStartedAt() bool {
	if o != nil && !IsNil(o.RecordingStartedAt) {
		return true
	}

	return false
}

// SetRecordingStartedAt gets a reference to the given string and assigns it to the RecordingStartedAt field.
func (o *RecordingResponseData) SetRecordingStartedAt(v string) {
	o.RecordingStartedAt = &v
}

// GetRecordingEndedAt returns the RecordingEndedAt field value if set, zero value otherwise.
func (o *RecordingResponseData) GetRecordingEndedAt() string {
	if o == nil || IsNil(o.RecordingEndedAt) {
		var ret string
		return ret
	}
	return *o.RecordingEndedAt
}

// GetRecordingEndedAtOk returns a tuple with the RecordingEndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetRecordingEndedAtOk() (*string, bool) {
	if o == nil || IsNil(o.RecordingEndedAt) {
		return nil, false
	}
	return o.RecordingEndedAt, true
}

// HasRecordingEndedAt returns a boolean if a field has been set.
func (o *RecordingResponseData) HasRecordingEndedAt() bool {
	if o != nil && !IsNil(o.RecordingEndedAt) {
		return true
	}

	return false
}

// SetRecordingEndedAt gets a reference to the given string and assigns it to the RecordingEndedAt field.
func (o *RecordingResponseData) SetRecordingEndedAt(v string) {
	o.RecordingEndedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RecordingResponseData) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RecordingResponseData) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *RecordingResponseData) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecordingResponseData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecordingResponseData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecordingResponseData) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RecordingResponseData) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordingResponseData) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RecordingResponseData) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *RecordingResponseData) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o RecordingResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordingResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallControlId) {
		toSerialize["call_control_id"] = o.CallControlId
	}
	if !IsNil(o.CallLegId) {
		toSerialize["call_leg_id"] = o.CallLegId
	}
	if !IsNil(o.CallSessionId) {
		toSerialize["call_session_id"] = o.CallSessionId
	}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.ConferenceId) {
		toSerialize["conference_id"] = o.ConferenceId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DownloadUrls) {
		toSerialize["download_urls"] = o.DownloadUrls
	}
	if !IsNil(o.DurationMillis) {
		toSerialize["duration_millis"] = o.DurationMillis
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.RecordingStartedAt) {
		toSerialize["recording_started_at"] = o.RecordingStartedAt
	}
	if !IsNil(o.RecordingEndedAt) {
		toSerialize["recording_ended_at"] = o.RecordingEndedAt
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableRecordingResponseData struct {
	value *RecordingResponseData
	isSet bool
}

func (v NullableRecordingResponseData) Get() *RecordingResponseData {
	return v.value
}

func (v *NullableRecordingResponseData) Set(val *RecordingResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingResponseData(val *RecordingResponseData) *NullableRecordingResponseData {
	return &NullableRecordingResponseData{value: val, isSet: true}
}

func (v NullableRecordingResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


