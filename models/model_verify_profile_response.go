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

// checks if the VerifyProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyProfileResponse{}

// VerifyProfileResponse struct for VerifyProfileResponse
type VerifyProfileResponse struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	WebhookUrl *string `json:"webhook_url,omitempty"`
	WebhookFailoverUrl *string `json:"webhook_failover_url,omitempty"`
	RecordType *VerificationProfileRecordType `json:"record_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Sms VerifyProfileSMSResponse `json:"sms,omitempty"`
	Call VerifyProfileCallResponse `json:"call,omitempty"`
	Flashcall VerifyProfileFlashcallResponse `json:"flashcall,omitempty"`
	Language *string `json:"language,omitempty"`
}

// NewVerifyProfileResponse instantiates a new VerifyProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyProfileResponse() *VerifyProfileResponse {
	this := VerifyProfileResponse{}
	return &this
}

// NewVerifyProfileResponseWithDefaults instantiates a new VerifyProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyProfileResponseWithDefaults() *VerifyProfileResponse {
	this := VerifyProfileResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VerifyProfileResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VerifyProfileResponse) SetName(v string) {
	o.Name = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *VerifyProfileResponse) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetWebhookFailoverUrl returns the WebhookFailoverUrl field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetWebhookFailoverUrl() string {
	if o == nil || IsNil(o.WebhookFailoverUrl) {
		var ret string
		return ret
	}
	return *o.WebhookFailoverUrl
}

// GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetWebhookFailoverUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookFailoverUrl) {
		return nil, false
	}
	return o.WebhookFailoverUrl, true
}

// HasWebhookFailoverUrl returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasWebhookFailoverUrl() bool {
	if o != nil && !IsNil(o.WebhookFailoverUrl) {
		return true
	}

	return false
}

// SetWebhookFailoverUrl gets a reference to the given string and assigns it to the WebhookFailoverUrl field.
func (o *VerifyProfileResponse) SetWebhookFailoverUrl(v string) {
	o.WebhookFailoverUrl = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetRecordType() VerificationProfileRecordType {
	if o == nil || IsNil(o.RecordType) {
		var ret VerificationProfileRecordType
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetRecordTypeOk() (*VerificationProfileRecordType, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given VerificationProfileRecordType and assigns it to the RecordType field.
func (o *VerifyProfileResponse) SetRecordType(v VerificationProfileRecordType) {
	o.RecordType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *VerifyProfileResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *VerifyProfileResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetSms() VerifyProfileSMSResponse {
	if o == nil || IsNil(o.Sms) {
		var ret VerifyProfileSMSResponse
		return ret
	}
	return o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetSmsOk() (VerifyProfileSMSResponse, bool) {
	if o == nil || IsNil(o.Sms) {
		return VerifyProfileSMSResponse{}, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasSms() bool {
	if o != nil && !IsNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given VerifyProfileSMSResponse and assigns it to the Sms field.
func (o *VerifyProfileResponse) SetSms(v VerifyProfileSMSResponse) {
	o.Sms = v
}

// GetCall returns the Call field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetCall() VerifyProfileCallResponse {
	if o == nil || IsNil(o.Call) {
		var ret VerifyProfileCallResponse
		return ret
	}
	return o.Call
}

// GetCallOk returns a tuple with the Call field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetCallOk() (VerifyProfileCallResponse, bool) {
	if o == nil || IsNil(o.Call) {
		return VerifyProfileCallResponse{}, false
	}
	return o.Call, true
}

// HasCall returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasCall() bool {
	if o != nil && !IsNil(o.Call) {
		return true
	}

	return false
}

// SetCall gets a reference to the given VerifyProfileCallResponse and assigns it to the Call field.
func (o *VerifyProfileResponse) SetCall(v VerifyProfileCallResponse) {
	o.Call = v
}

// GetFlashcall returns the Flashcall field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetFlashcall() VerifyProfileFlashcallResponse {
	if o == nil || IsNil(o.Flashcall) {
		var ret VerifyProfileFlashcallResponse
		return ret
	}
	return o.Flashcall
}

// GetFlashcallOk returns a tuple with the Flashcall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetFlashcallOk() (VerifyProfileFlashcallResponse, bool) {
	if o == nil || IsNil(o.Flashcall) {
		return VerifyProfileFlashcallResponse{}, false
	}
	return o.Flashcall, true
}

// HasFlashcall returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasFlashcall() bool {
	if o != nil && !IsNil(o.Flashcall) {
		return true
	}

	return false
}

// SetFlashcall gets a reference to the given VerifyProfileFlashcallResponse and assigns it to the Flashcall field.
func (o *VerifyProfileResponse) SetFlashcall(v VerifyProfileFlashcallResponse) {
	o.Flashcall = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *VerifyProfileResponse) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyProfileResponse) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *VerifyProfileResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *VerifyProfileResponse) SetLanguage(v string) {
	o.Language = &v
}

func (o VerifyProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if !IsNil(o.WebhookFailoverUrl) {
		toSerialize["webhook_failover_url"] = o.WebhookFailoverUrl
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}
	if !IsNil(o.Call) {
		toSerialize["call"] = o.Call
	}
	if !IsNil(o.Flashcall) {
		toSerialize["flashcall"] = o.Flashcall
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableVerifyProfileResponse struct {
	value *VerifyProfileResponse
	isSet bool
}

func (v NullableVerifyProfileResponse) Get() *VerifyProfileResponse {
	return v.value
}

func (v *NullableVerifyProfileResponse) Set(val *VerifyProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyProfileResponse(val *VerifyProfileResponse) *NullableVerifyProfileResponse {
	return &NullableVerifyProfileResponse{value: val, isSet: true}
}

func (v NullableVerifyProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


