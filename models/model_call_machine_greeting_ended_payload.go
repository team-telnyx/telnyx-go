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

// checks if the CallMachineGreetingEndedPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallMachineGreetingEndedPayload{}

// CallMachineGreetingEndedPayload struct for CallMachineGreetingEndedPayload
type CallMachineGreetingEndedPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlId *string `json:"call_control_id,omitempty"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegId *string `json:"call_leg_id,omitempty"`
	// ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionId *string `json:"call_session_id,omitempty"`
	// State received from a command.
	ClientState *string `json:"client_state,omitempty"`
	// Number or SIP URI placing the call.
	From *string `json:"from,omitempty"`
	// Destination number or SIP URI of the call.
	To *string `json:"to,omitempty"`
	// Answering machine greeting ended result.
	Result *string `json:"result,omitempty"`
}

// NewCallMachineGreetingEndedPayload instantiates a new CallMachineGreetingEndedPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallMachineGreetingEndedPayload() *CallMachineGreetingEndedPayload {
	this := CallMachineGreetingEndedPayload{}
	return &this
}

// NewCallMachineGreetingEndedPayloadWithDefaults instantiates a new CallMachineGreetingEndedPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallMachineGreetingEndedPayloadWithDefaults() *CallMachineGreetingEndedPayload {
	this := CallMachineGreetingEndedPayload{}
	return &this
}

// GetCallControlId returns the CallControlId field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetCallControlId() string {
	if o == nil || IsNil(o.CallControlId) {
		var ret string
		return ret
	}
	return *o.CallControlId
}

// GetCallControlIdOk returns a tuple with the CallControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetCallControlIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallControlId) {
		return nil, false
	}
	return o.CallControlId, true
}

// HasCallControlId returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasCallControlId() bool {
	if o != nil && !IsNil(o.CallControlId) {
		return true
	}

	return false
}

// SetCallControlId gets a reference to the given string and assigns it to the CallControlId field.
func (o *CallMachineGreetingEndedPayload) SetCallControlId(v string) {
	o.CallControlId = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *CallMachineGreetingEndedPayload) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetCallLegId returns the CallLegId field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetCallLegId() string {
	if o == nil || IsNil(o.CallLegId) {
		var ret string
		return ret
	}
	return *o.CallLegId
}

// GetCallLegIdOk returns a tuple with the CallLegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetCallLegIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallLegId) {
		return nil, false
	}
	return o.CallLegId, true
}

// HasCallLegId returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasCallLegId() bool {
	if o != nil && !IsNil(o.CallLegId) {
		return true
	}

	return false
}

// SetCallLegId gets a reference to the given string and assigns it to the CallLegId field.
func (o *CallMachineGreetingEndedPayload) SetCallLegId(v string) {
	o.CallLegId = &v
}

// GetCallSessionId returns the CallSessionId field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetCallSessionId() string {
	if o == nil || IsNil(o.CallSessionId) {
		var ret string
		return ret
	}
	return *o.CallSessionId
}

// GetCallSessionIdOk returns a tuple with the CallSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetCallSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallSessionId) {
		return nil, false
	}
	return o.CallSessionId, true
}

// HasCallSessionId returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasCallSessionId() bool {
	if o != nil && !IsNil(o.CallSessionId) {
		return true
	}

	return false
}

// SetCallSessionId gets a reference to the given string and assigns it to the CallSessionId field.
func (o *CallMachineGreetingEndedPayload) SetCallSessionId(v string) {
	o.CallSessionId = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *CallMachineGreetingEndedPayload) SetClientState(v string) {
	o.ClientState = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *CallMachineGreetingEndedPayload) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *CallMachineGreetingEndedPayload) SetTo(v string) {
	o.To = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CallMachineGreetingEndedPayload) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallMachineGreetingEndedPayload) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CallMachineGreetingEndedPayload) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *CallMachineGreetingEndedPayload) SetResult(v string) {
	o.Result = &v
}

func (o CallMachineGreetingEndedPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallMachineGreetingEndedPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallControlId) {
		toSerialize["call_control_id"] = o.CallControlId
	}
	if !IsNil(o.ConnectionId) {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if !IsNil(o.CallLegId) {
		toSerialize["call_leg_id"] = o.CallLegId
	}
	if !IsNil(o.CallSessionId) {
		toSerialize["call_session_id"] = o.CallSessionId
	}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCallMachineGreetingEndedPayload struct {
	value *CallMachineGreetingEndedPayload
	isSet bool
}

func (v NullableCallMachineGreetingEndedPayload) Get() *CallMachineGreetingEndedPayload {
	return v.value
}

func (v *NullableCallMachineGreetingEndedPayload) Set(val *CallMachineGreetingEndedPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCallMachineGreetingEndedPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCallMachineGreetingEndedPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallMachineGreetingEndedPayload(val *CallMachineGreetingEndedPayload) *NullableCallMachineGreetingEndedPayload {
	return &NullableCallMachineGreetingEndedPayload{value: val, isSet: true}
}

func (v NullableCallMachineGreetingEndedPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallMachineGreetingEndedPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


