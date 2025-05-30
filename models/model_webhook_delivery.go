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
	"time"
)

// checks if the WebhookDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookDelivery{}

// WebhookDelivery Record of all attempts to deliver a webhook.
type WebhookDelivery struct {
	// Uniquely identifies the webhook_delivery record.
	Id *string `json:"id,omitempty"`
	// Uniquely identifies the user that owns the webhook_delivery record.
	UserId *string `json:"user_id,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// Delivery status: 'delivered' when successfuly delivered or 'failed' if all attempts have failed.
	Status *string `json:"status,omitempty"`
	Webhook *WebhookDeliveryWebhook `json:"webhook,omitempty"`
	// ISO 8601 timestamp indicating when the first request attempt was initiated.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// ISO 8601 timestamp indicating when the last webhook response has been received.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// Detailed delivery attempts, ordered by most recent.
	Attempts []Attempt `json:"attempts,omitempty"`
}

// NewWebhookDelivery instantiates a new WebhookDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDelivery() *WebhookDelivery {
	this := WebhookDelivery{}
	return &this
}

// NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryWithDefaults() *WebhookDelivery {
	this := WebhookDelivery{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebhookDelivery) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebhookDelivery) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebhookDelivery) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WebhookDelivery) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WebhookDelivery) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *WebhookDelivery) SetUserId(v string) {
	o.UserId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *WebhookDelivery) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *WebhookDelivery) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *WebhookDelivery) SetRecordType(v string) {
	o.RecordType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhookDelivery) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhookDelivery) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhookDelivery) SetStatus(v string) {
	o.Status = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *WebhookDelivery) GetWebhook() WebhookDeliveryWebhook {
	if o == nil || IsNil(o.Webhook) {
		var ret WebhookDeliveryWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetWebhookOk() (*WebhookDeliveryWebhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *WebhookDelivery) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhookDeliveryWebhook and assigns it to the Webhook field.
func (o *WebhookDelivery) SetWebhook(v WebhookDeliveryWebhook) {
	o.Webhook = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *WebhookDelivery) GetStartedAt() time.Time {
	if o == nil || IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *WebhookDelivery) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *WebhookDelivery) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *WebhookDelivery) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *WebhookDelivery) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *WebhookDelivery) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *WebhookDelivery) GetAttempts() []Attempt {
	if o == nil || IsNil(o.Attempts) {
		var ret []Attempt
		return ret
	}
	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetAttemptsOk() ([]Attempt, bool) {
	if o == nil || IsNil(o.Attempts) {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *WebhookDelivery) HasAttempts() bool {
	if o != nil && !IsNil(o.Attempts) {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given []Attempt and assigns it to the Attempts field.
func (o *WebhookDelivery) SetAttempts(v []Attempt) {
	o.Attempts = v
}

func (o WebhookDelivery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if !IsNil(o.Attempts) {
		toSerialize["attempts"] = o.Attempts
	}
	return toSerialize, nil
}

type NullableWebhookDelivery struct {
	value *WebhookDelivery
	isSet bool
}

func (v NullableWebhookDelivery) Get() *WebhookDelivery {
	return v.value
}

func (v *NullableWebhookDelivery) Set(val *WebhookDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDelivery(val *WebhookDelivery) *NullableWebhookDelivery {
	return &NullableWebhookDelivery{value: val, isSet: true}
}

func (v NullableWebhookDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


