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

// checks if the CallReferFailedPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallReferFailedPayload{}

// CallReferFailedPayload struct for CallReferFailedPayload
type CallReferFailedPayload struct {
	// Unique ID for controlling the call.
	CallControlId *string `json:"call_control_id,omitempty"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegId *string `json:"call_leg_id,omitempty"`
	// ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionId *string `json:"call_session_id,omitempty"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// State received from a command.
	ClientState *string `json:"client_state,omitempty"`
	// Number or SIP URI placing the call.
	From *string `json:"from,omitempty"`
	// SIP NOTIFY event status for tracking the REFER attempt.
	SipNotifyResponse *int32 `json:"sip_notify_response,omitempty"`
	// Destination number or SIP URI of the call.
	To *string `json:"to,omitempty"`
}

// NewCallReferFailedPayload instantiates a new CallReferFailedPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallReferFailedPayload() *CallReferFailedPayload {
	this := CallReferFailedPayload{}
	return &this
}

// NewCallReferFailedPayloadWithDefaults instantiates a new CallReferFailedPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallReferFailedPayloadWithDefaults() *CallReferFailedPayload {
	this := CallReferFailedPayload{}
	return &this
}

// GetCallControlId returns the CallControlId field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetCallControlId() string {
	if o == nil || IsNil(o.CallControlId) {
		var ret string
		return ret
	}
	return *o.CallControlId
}

// GetCallControlIdOk returns a tuple with the CallControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetCallControlIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallControlId) {
		return nil, false
	}
	return o.CallControlId, true
}

// HasCallControlId returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasCallControlId() bool {
	if o != nil && !IsNil(o.CallControlId) {
		return true
	}

	return false
}

// SetCallControlId gets a reference to the given string and assigns it to the CallControlId field.
func (o *CallReferFailedPayload) SetCallControlId(v string) {
	o.CallControlId = &v
}

// GetCallLegId returns the CallLegId field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetCallLegId() string {
	if o == nil || IsNil(o.CallLegId) {
		var ret string
		return ret
	}
	return *o.CallLegId
}

// GetCallLegIdOk returns a tuple with the CallLegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetCallLegIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallLegId) {
		return nil, false
	}
	return o.CallLegId, true
}

// HasCallLegId returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasCallLegId() bool {
	if o != nil && !IsNil(o.CallLegId) {
		return true
	}

	return false
}

// SetCallLegId gets a reference to the given string and assigns it to the CallLegId field.
func (o *CallReferFailedPayload) SetCallLegId(v string) {
	o.CallLegId = &v
}

// GetCallSessionId returns the CallSessionId field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetCallSessionId() string {
	if o == nil || IsNil(o.CallSessionId) {
		var ret string
		return ret
	}
	return *o.CallSessionId
}

// GetCallSessionIdOk returns a tuple with the CallSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetCallSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallSessionId) {
		return nil, false
	}
	return o.CallSessionId, true
}

// HasCallSessionId returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasCallSessionId() bool {
	if o != nil && !IsNil(o.CallSessionId) {
		return true
	}

	return false
}

// SetCallSessionId gets a reference to the given string and assigns it to the CallSessionId field.
func (o *CallReferFailedPayload) SetCallSessionId(v string) {
	o.CallSessionId = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *CallReferFailedPayload) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *CallReferFailedPayload) SetClientState(v string) {
	o.ClientState = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CallReferFailedPayload) SetFrom(v string) {
	o.From = &v
}

// GetSipNotifyResponse returns the SipNotifyResponse field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetSipNotifyResponse() int32 {
	if o == nil || IsNil(o.SipNotifyResponse) {
		var ret int32
		return ret
	}
	return *o.SipNotifyResponse
}

// GetSipNotifyResponseOk returns a tuple with the SipNotifyResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetSipNotifyResponseOk() (*int32, bool) {
	if o == nil || IsNil(o.SipNotifyResponse) {
		return nil, false
	}
	return o.SipNotifyResponse, true
}

// HasSipNotifyResponse returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasSipNotifyResponse() bool {
	if o != nil && !IsNil(o.SipNotifyResponse) {
		return true
	}

	return false
}

// SetSipNotifyResponse gets a reference to the given int32 and assigns it to the SipNotifyResponse field.
func (o *CallReferFailedPayload) SetSipNotifyResponse(v int32) {
	o.SipNotifyResponse = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CallReferFailedPayload) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedPayload) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CallReferFailedPayload) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CallReferFailedPayload) SetTo(v string) {
	o.To = &v
}

func (o CallReferFailedPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallReferFailedPayload) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.SipNotifyResponse) {
		toSerialize["sip_notify_response"] = o.SipNotifyResponse
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableCallReferFailedPayload struct {
	value *CallReferFailedPayload
	isSet bool
}

func (v NullableCallReferFailedPayload) Get() *CallReferFailedPayload {
	return v.value
}

func (v *NullableCallReferFailedPayload) Set(val *CallReferFailedPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCallReferFailedPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCallReferFailedPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallReferFailedPayload(val *CallReferFailedPayload) *NullableCallReferFailedPayload {
	return &NullableCallReferFailedPayload{value: val, isSet: true}
}

func (v NullableCallReferFailedPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallReferFailedPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


