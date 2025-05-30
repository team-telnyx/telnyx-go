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
	"bytes"
	"fmt"
)

// checks if the SendFaxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendFaxRequest{}

// SendFaxRequest struct for SendFaxRequest
type SendFaxRequest struct {
	// The connection ID to send the fax with.
	ConnectionId string `json:"connection_id"`
	// The URL (or list of URLs) to the PDF used for the fax's media. media_url and media_name/contents can't be submitted together.
	MediaUrl *string `json:"media_url,omitempty"`
	// The media_name used for the fax's media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can't be submitted together.
	MediaName *string `json:"media_name,omitempty"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The `from_display_name` string to be used as the caller id name (SIP From Display Name) presented to the destination (`to` number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the `from` field.
	FromDisplayName *string `json:"from_display_name,omitempty"`
	Quality *Quality `json:"quality,omitempty"`
	// The flag to disable the T.38 protocol.
	T38Enabled *bool `json:"t38_enabled,omitempty"`
	// The flag to enable monochrome, true black and white fax results.
	Monochrome *bool `json:"monochrome,omitempty"`
	// Should fax media be stored on temporary URL. It does not support media_name, they can't be submitted together.
	StoreMedia *bool `json:"store_media,omitempty"`
	// Should fax preview be stored on temporary URL.
	StorePreview *bool `json:"store_preview,omitempty"`
	// The format for the preview file in case the `store_preview` is `true`.
	PreviewFormat *string `json:"preview_format,omitempty"`
	// Use this field to override the URL to which Telnyx will send subsequent webhooks for this fax.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
}

type _SendFaxRequest SendFaxRequest

// NewSendFaxRequest instantiates a new SendFaxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendFaxRequest(connectionId string, to string, from string) *SendFaxRequest {
	this := SendFaxRequest{}
	this.ConnectionId = connectionId
	this.To = to
	this.From = from
	var quality Quality = HIGH
	this.Quality = &quality
	var t38Enabled bool = true
	this.T38Enabled = &t38Enabled
	var monochrome bool = false
	this.Monochrome = &monochrome
	var storeMedia bool = false
	this.StoreMedia = &storeMedia
	var storePreview bool = false
	this.StorePreview = &storePreview
	var previewFormat string = "tiff"
	this.PreviewFormat = &previewFormat
	return &this
}

// NewSendFaxRequestWithDefaults instantiates a new SendFaxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendFaxRequestWithDefaults() *SendFaxRequest {
	this := SendFaxRequest{}
	var quality Quality = HIGH
	this.Quality = &quality
	var t38Enabled bool = true
	this.T38Enabled = &t38Enabled
	var monochrome bool = false
	this.Monochrome = &monochrome
	var storeMedia bool = false
	this.StoreMedia = &storeMedia
	var storePreview bool = false
	this.StorePreview = &storePreview
	var previewFormat string = "tiff"
	this.PreviewFormat = &previewFormat
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *SendFaxRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *SendFaxRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetMediaUrl returns the MediaUrl field value if set, zero value otherwise.
func (o *SendFaxRequest) GetMediaUrl() string {
	if o == nil || IsNil(o.MediaUrl) {
		var ret string
		return ret
	}
	return *o.MediaUrl
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetMediaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MediaUrl) {
		return nil, false
	}
	return o.MediaUrl, true
}

// HasMediaUrl returns a boolean if a field has been set.
func (o *SendFaxRequest) HasMediaUrl() bool {
	if o != nil && !IsNil(o.MediaUrl) {
		return true
	}

	return false
}

// SetMediaUrl gets a reference to the given string and assigns it to the MediaUrl field.
func (o *SendFaxRequest) SetMediaUrl(v string) {
	o.MediaUrl = &v
}

// GetMediaName returns the MediaName field value if set, zero value otherwise.
func (o *SendFaxRequest) GetMediaName() string {
	if o == nil || IsNil(o.MediaName) {
		var ret string
		return ret
	}
	return *o.MediaName
}

// GetMediaNameOk returns a tuple with the MediaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetMediaNameOk() (*string, bool) {
	if o == nil || IsNil(o.MediaName) {
		return nil, false
	}
	return o.MediaName, true
}

// HasMediaName returns a boolean if a field has been set.
func (o *SendFaxRequest) HasMediaName() bool {
	if o != nil && !IsNil(o.MediaName) {
		return true
	}

	return false
}

// SetMediaName gets a reference to the given string and assigns it to the MediaName field.
func (o *SendFaxRequest) SetMediaName(v string) {
	o.MediaName = &v
}

// GetTo returns the To field value
func (o *SendFaxRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *SendFaxRequest) SetTo(v string) {
	o.To = v
}

// GetFrom returns the From field value
func (o *SendFaxRequest) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *SendFaxRequest) SetFrom(v string) {
	o.From = v
}

// GetFromDisplayName returns the FromDisplayName field value if set, zero value otherwise.
func (o *SendFaxRequest) GetFromDisplayName() string {
	if o == nil || IsNil(o.FromDisplayName) {
		var ret string
		return ret
	}
	return *o.FromDisplayName
}

// GetFromDisplayNameOk returns a tuple with the FromDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetFromDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.FromDisplayName) {
		return nil, false
	}
	return o.FromDisplayName, true
}

// HasFromDisplayName returns a boolean if a field has been set.
func (o *SendFaxRequest) HasFromDisplayName() bool {
	if o != nil && !IsNil(o.FromDisplayName) {
		return true
	}

	return false
}

// SetFromDisplayName gets a reference to the given string and assigns it to the FromDisplayName field.
func (o *SendFaxRequest) SetFromDisplayName(v string) {
	o.FromDisplayName = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *SendFaxRequest) GetQuality() Quality {
	if o == nil || IsNil(o.Quality) {
		var ret Quality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetQualityOk() (*Quality, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *SendFaxRequest) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given Quality and assigns it to the Quality field.
func (o *SendFaxRequest) SetQuality(v Quality) {
	o.Quality = &v
}

// GetT38Enabled returns the T38Enabled field value if set, zero value otherwise.
func (o *SendFaxRequest) GetT38Enabled() bool {
	if o == nil || IsNil(o.T38Enabled) {
		var ret bool
		return ret
	}
	return *o.T38Enabled
}

// GetT38EnabledOk returns a tuple with the T38Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetT38EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.T38Enabled) {
		return nil, false
	}
	return o.T38Enabled, true
}

// HasT38Enabled returns a boolean if a field has been set.
func (o *SendFaxRequest) HasT38Enabled() bool {
	if o != nil && !IsNil(o.T38Enabled) {
		return true
	}

	return false
}

// SetT38Enabled gets a reference to the given bool and assigns it to the T38Enabled field.
func (o *SendFaxRequest) SetT38Enabled(v bool) {
	o.T38Enabled = &v
}

// GetMonochrome returns the Monochrome field value if set, zero value otherwise.
func (o *SendFaxRequest) GetMonochrome() bool {
	if o == nil || IsNil(o.Monochrome) {
		var ret bool
		return ret
	}
	return *o.Monochrome
}

// GetMonochromeOk returns a tuple with the Monochrome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetMonochromeOk() (*bool, bool) {
	if o == nil || IsNil(o.Monochrome) {
		return nil, false
	}
	return o.Monochrome, true
}

// HasMonochrome returns a boolean if a field has been set.
func (o *SendFaxRequest) HasMonochrome() bool {
	if o != nil && !IsNil(o.Monochrome) {
		return true
	}

	return false
}

// SetMonochrome gets a reference to the given bool and assigns it to the Monochrome field.
func (o *SendFaxRequest) SetMonochrome(v bool) {
	o.Monochrome = &v
}

// GetStoreMedia returns the StoreMedia field value if set, zero value otherwise.
func (o *SendFaxRequest) GetStoreMedia() bool {
	if o == nil || IsNil(o.StoreMedia) {
		var ret bool
		return ret
	}
	return *o.StoreMedia
}

// GetStoreMediaOk returns a tuple with the StoreMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetStoreMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreMedia) {
		return nil, false
	}
	return o.StoreMedia, true
}

// HasStoreMedia returns a boolean if a field has been set.
func (o *SendFaxRequest) HasStoreMedia() bool {
	if o != nil && !IsNil(o.StoreMedia) {
		return true
	}

	return false
}

// SetStoreMedia gets a reference to the given bool and assigns it to the StoreMedia field.
func (o *SendFaxRequest) SetStoreMedia(v bool) {
	o.StoreMedia = &v
}

// GetStorePreview returns the StorePreview field value if set, zero value otherwise.
func (o *SendFaxRequest) GetStorePreview() bool {
	if o == nil || IsNil(o.StorePreview) {
		var ret bool
		return ret
	}
	return *o.StorePreview
}

// GetStorePreviewOk returns a tuple with the StorePreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetStorePreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.StorePreview) {
		return nil, false
	}
	return o.StorePreview, true
}

// HasStorePreview returns a boolean if a field has been set.
func (o *SendFaxRequest) HasStorePreview() bool {
	if o != nil && !IsNil(o.StorePreview) {
		return true
	}

	return false
}

// SetStorePreview gets a reference to the given bool and assigns it to the StorePreview field.
func (o *SendFaxRequest) SetStorePreview(v bool) {
	o.StorePreview = &v
}

// GetPreviewFormat returns the PreviewFormat field value if set, zero value otherwise.
func (o *SendFaxRequest) GetPreviewFormat() string {
	if o == nil || IsNil(o.PreviewFormat) {
		var ret string
		return ret
	}
	return *o.PreviewFormat
}

// GetPreviewFormatOk returns a tuple with the PreviewFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetPreviewFormatOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFormat) {
		return nil, false
	}
	return o.PreviewFormat, true
}

// HasPreviewFormat returns a boolean if a field has been set.
func (o *SendFaxRequest) HasPreviewFormat() bool {
	if o != nil && !IsNil(o.PreviewFormat) {
		return true
	}

	return false
}

// SetPreviewFormat gets a reference to the given string and assigns it to the PreviewFormat field.
func (o *SendFaxRequest) SetPreviewFormat(v string) {
	o.PreviewFormat = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *SendFaxRequest) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *SendFaxRequest) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *SendFaxRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *SendFaxRequest) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxRequest) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *SendFaxRequest) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *SendFaxRequest) SetClientState(v string) {
	o.ClientState = &v
}

func (o SendFaxRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendFaxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection_id"] = o.ConnectionId
	if !IsNil(o.MediaUrl) {
		toSerialize["media_url"] = o.MediaUrl
	}
	if !IsNil(o.MediaName) {
		toSerialize["media_name"] = o.MediaName
	}
	toSerialize["to"] = o.To
	toSerialize["from"] = o.From
	if !IsNil(o.FromDisplayName) {
		toSerialize["from_display_name"] = o.FromDisplayName
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !IsNil(o.T38Enabled) {
		toSerialize["t38_enabled"] = o.T38Enabled
	}
	if !IsNil(o.Monochrome) {
		toSerialize["monochrome"] = o.Monochrome
	}
	if !IsNil(o.StoreMedia) {
		toSerialize["store_media"] = o.StoreMedia
	}
	if !IsNil(o.StorePreview) {
		toSerialize["store_preview"] = o.StorePreview
	}
	if !IsNil(o.PreviewFormat) {
		toSerialize["preview_format"] = o.PreviewFormat
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	return toSerialize, nil
}

func (o *SendFaxRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connection_id",
		"to",
		"from",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSendFaxRequest := _SendFaxRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSendFaxRequest)

	if err != nil {
		return err
	}

	*o = SendFaxRequest(varSendFaxRequest)

	return err
}

type NullableSendFaxRequest struct {
	value *SendFaxRequest
	isSet bool
}

func (v NullableSendFaxRequest) Get() *SendFaxRequest {
	return v.value
}

func (v *NullableSendFaxRequest) Set(val *SendFaxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendFaxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendFaxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendFaxRequest(val *SendFaxRequest) *NullableSendFaxRequest {
	return &NullableSendFaxRequest{value: val, isSet: true}
}

func (v NullableSendFaxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendFaxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


