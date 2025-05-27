# UpdatePhoneNumberVoiceSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TechPrefixEnabled** | Pointer to **bool** | Controls whether a tech prefix is enabled for this phone number. | [optional] [default to false]
**TranslatedNumber** | Pointer to **string** | This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed. | [optional] 
**CallerIdNameEnabled** | Pointer to **bool** | Controls whether the caller ID name is enabled for this phone number. | [optional] [default to false]
**CallForwarding** | Pointer to [**CallForwarding**](CallForwarding.md) |  | [optional] 
**CnamListing** | Pointer to [**CnamListing**](CnamListing.md) |  | [optional] 
**UsagePaymentMethod** | Pointer to **string** | Controls whether a number is billed per minute or uses your concurrent channels. | [optional] [default to "pay-per-minute"]
**MediaFeatures** | Pointer to [**MediaFeatures**](MediaFeatures.md) |  | [optional] 
**CallRecording** | Pointer to [**CallRecording**](CallRecording.md) |  | [optional] 
**InboundCallScreening** | Pointer to **string** | The inbound_call_screening setting is a phone number configuration option variable that allows users to configure their settings to block or flag fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This feature has an additional per-number monthly cost associated with it. | [optional] [default to "disabled"]

## Methods

### NewUpdatePhoneNumberVoiceSettingsRequest

`func NewUpdatePhoneNumberVoiceSettingsRequest() *UpdatePhoneNumberVoiceSettingsRequest`

NewUpdatePhoneNumberVoiceSettingsRequest instantiates a new UpdatePhoneNumberVoiceSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePhoneNumberVoiceSettingsRequestWithDefaults

`func NewUpdatePhoneNumberVoiceSettingsRequestWithDefaults() *UpdatePhoneNumberVoiceSettingsRequest`

NewUpdatePhoneNumberVoiceSettingsRequestWithDefaults instantiates a new UpdatePhoneNumberVoiceSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTechPrefixEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetTechPrefixEnabled() bool`

GetTechPrefixEnabled returns the TechPrefixEnabled field if non-nil, zero value otherwise.

### GetTechPrefixEnabledOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetTechPrefixEnabledOk() (*bool, bool)`

GetTechPrefixEnabledOk returns a tuple with the TechPrefixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechPrefixEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetTechPrefixEnabled(v bool)`

SetTechPrefixEnabled sets TechPrefixEnabled field to given value.

### HasTechPrefixEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasTechPrefixEnabled() bool`

HasTechPrefixEnabled returns a boolean if a field has been set.

### GetTranslatedNumber

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetTranslatedNumber() string`

GetTranslatedNumber returns the TranslatedNumber field if non-nil, zero value otherwise.

### GetTranslatedNumberOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetTranslatedNumberOk() (*string, bool)`

GetTranslatedNumberOk returns a tuple with the TranslatedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedNumber

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetTranslatedNumber(v string)`

SetTranslatedNumber sets TranslatedNumber field to given value.

### HasTranslatedNumber

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasTranslatedNumber() bool`

HasTranslatedNumber returns a boolean if a field has been set.

### GetCallerIdNameEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallerIdNameEnabled() bool`

GetCallerIdNameEnabled returns the CallerIdNameEnabled field if non-nil, zero value otherwise.

### GetCallerIdNameEnabledOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallerIdNameEnabledOk() (*bool, bool)`

GetCallerIdNameEnabledOk returns a tuple with the CallerIdNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNameEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetCallerIdNameEnabled(v bool)`

SetCallerIdNameEnabled sets CallerIdNameEnabled field to given value.

### HasCallerIdNameEnabled

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasCallerIdNameEnabled() bool`

HasCallerIdNameEnabled returns a boolean if a field has been set.

### GetCallForwarding

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallForwarding() CallForwarding`

GetCallForwarding returns the CallForwarding field if non-nil, zero value otherwise.

### GetCallForwardingOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallForwardingOk() (*CallForwarding, bool)`

GetCallForwardingOk returns a tuple with the CallForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwarding

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetCallForwarding(v CallForwarding)`

SetCallForwarding sets CallForwarding field to given value.

### HasCallForwarding

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasCallForwarding() bool`

HasCallForwarding returns a boolean if a field has been set.

### GetCnamListing

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCnamListing() CnamListing`

GetCnamListing returns the CnamListing field if non-nil, zero value otherwise.

### GetCnamListingOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCnamListingOk() (*CnamListing, bool)`

GetCnamListingOk returns a tuple with the CnamListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListing

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetCnamListing(v CnamListing)`

SetCnamListing sets CnamListing field to given value.

### HasCnamListing

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasCnamListing() bool`

HasCnamListing returns a boolean if a field has been set.

### GetUsagePaymentMethod

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetUsagePaymentMethod() string`

GetUsagePaymentMethod returns the UsagePaymentMethod field if non-nil, zero value otherwise.

### GetUsagePaymentMethodOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetUsagePaymentMethodOk() (*string, bool)`

GetUsagePaymentMethodOk returns a tuple with the UsagePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePaymentMethod

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetUsagePaymentMethod(v string)`

SetUsagePaymentMethod sets UsagePaymentMethod field to given value.

### HasUsagePaymentMethod

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasUsagePaymentMethod() bool`

HasUsagePaymentMethod returns a boolean if a field has been set.

### GetMediaFeatures

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetMediaFeatures() MediaFeatures`

GetMediaFeatures returns the MediaFeatures field if non-nil, zero value otherwise.

### GetMediaFeaturesOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetMediaFeaturesOk() (*MediaFeatures, bool)`

GetMediaFeaturesOk returns a tuple with the MediaFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaFeatures

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetMediaFeatures(v MediaFeatures)`

SetMediaFeatures sets MediaFeatures field to given value.

### HasMediaFeatures

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasMediaFeatures() bool`

HasMediaFeatures returns a boolean if a field has been set.

### GetCallRecording

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallRecording() CallRecording`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetCallRecordingOk() (*CallRecording, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetCallRecording(v CallRecording)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetInboundCallScreening

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetInboundCallScreening() string`

GetInboundCallScreening returns the InboundCallScreening field if non-nil, zero value otherwise.

### GetInboundCallScreeningOk

`func (o *UpdatePhoneNumberVoiceSettingsRequest) GetInboundCallScreeningOk() (*string, bool)`

GetInboundCallScreeningOk returns a tuple with the InboundCallScreening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallScreening

`func (o *UpdatePhoneNumberVoiceSettingsRequest) SetInboundCallScreening(v string)`

SetInboundCallScreening sets InboundCallScreening field to given value.

### HasInboundCallScreening

`func (o *UpdatePhoneNumberVoiceSettingsRequest) HasInboundCallScreening() bool`

HasInboundCallScreening returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


