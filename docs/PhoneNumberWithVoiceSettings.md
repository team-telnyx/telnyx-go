# PhoneNumberWithVoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The phone number in +E164 format. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with this phone number. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**TechPrefixEnabled** | Pointer to **bool** | Controls whether a tech prefix is enabled for this phone number. | [optional] [default to false]
**TranslatedNumber** | Pointer to **string** | This field allows you to rewrite the destination number of an inbound call before the call is routed to you. The value of this field may be any alphanumeric value, and the value will replace the number originally dialed. | [optional] [default to ""]
**CallForwarding** | Pointer to [**CallForwarding**](CallForwarding.md) |  | [optional] 
**CnamListing** | Pointer to [**CnamListing**](CnamListing.md) |  | [optional] 
**Emergency** | Pointer to [**EmergencySettings**](EmergencySettings.md) |  | [optional] 
**UsagePaymentMethod** | Pointer to **string** | Controls whether a number is billed per minute or uses your concurrent channels. | [optional] [default to "pay-per-minute"]
**MediaFeatures** | Pointer to [**MediaFeatures**](MediaFeatures.md) |  | [optional] 
**CallRecording** | Pointer to [**CallRecording**](CallRecording.md) |  | [optional] 
**InboundCallScreening** | Pointer to **string** | The inbound_call_screening setting is a phone number configuration option variable that allows users to configure their settings to block or flag fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This feature has an additional per-number monthly cost associated with it. | [optional] [default to "disabled"]

## Methods

### NewPhoneNumberWithVoiceSettings

`func NewPhoneNumberWithVoiceSettings() *PhoneNumberWithVoiceSettings`

NewPhoneNumberWithVoiceSettings instantiates a new PhoneNumberWithVoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithVoiceSettingsWithDefaults

`func NewPhoneNumberWithVoiceSettingsWithDefaults() *PhoneNumberWithVoiceSettings`

NewPhoneNumberWithVoiceSettingsWithDefaults instantiates a new PhoneNumberWithVoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhoneNumberWithVoiceSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNumberWithVoiceSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNumberWithVoiceSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNumberWithVoiceSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PhoneNumberWithVoiceSettings) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PhoneNumberWithVoiceSettings) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PhoneNumberWithVoiceSettings) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PhoneNumberWithVoiceSettings) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PhoneNumberWithVoiceSettings) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberWithVoiceSettings) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberWithVoiceSettings) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberWithVoiceSettings) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetConnectionId

`func (o *PhoneNumberWithVoiceSettings) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PhoneNumberWithVoiceSettings) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PhoneNumberWithVoiceSettings) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *PhoneNumberWithVoiceSettings) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PhoneNumberWithVoiceSettings) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PhoneNumberWithVoiceSettings) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PhoneNumberWithVoiceSettings) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PhoneNumberWithVoiceSettings) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetTechPrefixEnabled

`func (o *PhoneNumberWithVoiceSettings) GetTechPrefixEnabled() bool`

GetTechPrefixEnabled returns the TechPrefixEnabled field if non-nil, zero value otherwise.

### GetTechPrefixEnabledOk

`func (o *PhoneNumberWithVoiceSettings) GetTechPrefixEnabledOk() (*bool, bool)`

GetTechPrefixEnabledOk returns a tuple with the TechPrefixEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechPrefixEnabled

`func (o *PhoneNumberWithVoiceSettings) SetTechPrefixEnabled(v bool)`

SetTechPrefixEnabled sets TechPrefixEnabled field to given value.

### HasTechPrefixEnabled

`func (o *PhoneNumberWithVoiceSettings) HasTechPrefixEnabled() bool`

HasTechPrefixEnabled returns a boolean if a field has been set.

### GetTranslatedNumber

`func (o *PhoneNumberWithVoiceSettings) GetTranslatedNumber() string`

GetTranslatedNumber returns the TranslatedNumber field if non-nil, zero value otherwise.

### GetTranslatedNumberOk

`func (o *PhoneNumberWithVoiceSettings) GetTranslatedNumberOk() (*string, bool)`

GetTranslatedNumberOk returns a tuple with the TranslatedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedNumber

`func (o *PhoneNumberWithVoiceSettings) SetTranslatedNumber(v string)`

SetTranslatedNumber sets TranslatedNumber field to given value.

### HasTranslatedNumber

`func (o *PhoneNumberWithVoiceSettings) HasTranslatedNumber() bool`

HasTranslatedNumber returns a boolean if a field has been set.

### GetCallForwarding

`func (o *PhoneNumberWithVoiceSettings) GetCallForwarding() CallForwarding`

GetCallForwarding returns the CallForwarding field if non-nil, zero value otherwise.

### GetCallForwardingOk

`func (o *PhoneNumberWithVoiceSettings) GetCallForwardingOk() (*CallForwarding, bool)`

GetCallForwardingOk returns a tuple with the CallForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwarding

`func (o *PhoneNumberWithVoiceSettings) SetCallForwarding(v CallForwarding)`

SetCallForwarding sets CallForwarding field to given value.

### HasCallForwarding

`func (o *PhoneNumberWithVoiceSettings) HasCallForwarding() bool`

HasCallForwarding returns a boolean if a field has been set.

### GetCnamListing

`func (o *PhoneNumberWithVoiceSettings) GetCnamListing() CnamListing`

GetCnamListing returns the CnamListing field if non-nil, zero value otherwise.

### GetCnamListingOk

`func (o *PhoneNumberWithVoiceSettings) GetCnamListingOk() (*CnamListing, bool)`

GetCnamListingOk returns a tuple with the CnamListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListing

`func (o *PhoneNumberWithVoiceSettings) SetCnamListing(v CnamListing)`

SetCnamListing sets CnamListing field to given value.

### HasCnamListing

`func (o *PhoneNumberWithVoiceSettings) HasCnamListing() bool`

HasCnamListing returns a boolean if a field has been set.

### GetEmergency

`func (o *PhoneNumberWithVoiceSettings) GetEmergency() EmergencySettings`

GetEmergency returns the Emergency field if non-nil, zero value otherwise.

### GetEmergencyOk

`func (o *PhoneNumberWithVoiceSettings) GetEmergencyOk() (*EmergencySettings, bool)`

GetEmergencyOk returns a tuple with the Emergency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergency

`func (o *PhoneNumberWithVoiceSettings) SetEmergency(v EmergencySettings)`

SetEmergency sets Emergency field to given value.

### HasEmergency

`func (o *PhoneNumberWithVoiceSettings) HasEmergency() bool`

HasEmergency returns a boolean if a field has been set.

### GetUsagePaymentMethod

`func (o *PhoneNumberWithVoiceSettings) GetUsagePaymentMethod() string`

GetUsagePaymentMethod returns the UsagePaymentMethod field if non-nil, zero value otherwise.

### GetUsagePaymentMethodOk

`func (o *PhoneNumberWithVoiceSettings) GetUsagePaymentMethodOk() (*string, bool)`

GetUsagePaymentMethodOk returns a tuple with the UsagePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePaymentMethod

`func (o *PhoneNumberWithVoiceSettings) SetUsagePaymentMethod(v string)`

SetUsagePaymentMethod sets UsagePaymentMethod field to given value.

### HasUsagePaymentMethod

`func (o *PhoneNumberWithVoiceSettings) HasUsagePaymentMethod() bool`

HasUsagePaymentMethod returns a boolean if a field has been set.

### GetMediaFeatures

`func (o *PhoneNumberWithVoiceSettings) GetMediaFeatures() MediaFeatures`

GetMediaFeatures returns the MediaFeatures field if non-nil, zero value otherwise.

### GetMediaFeaturesOk

`func (o *PhoneNumberWithVoiceSettings) GetMediaFeaturesOk() (*MediaFeatures, bool)`

GetMediaFeaturesOk returns a tuple with the MediaFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaFeatures

`func (o *PhoneNumberWithVoiceSettings) SetMediaFeatures(v MediaFeatures)`

SetMediaFeatures sets MediaFeatures field to given value.

### HasMediaFeatures

`func (o *PhoneNumberWithVoiceSettings) HasMediaFeatures() bool`

HasMediaFeatures returns a boolean if a field has been set.

### GetCallRecording

`func (o *PhoneNumberWithVoiceSettings) GetCallRecording() CallRecording`

GetCallRecording returns the CallRecording field if non-nil, zero value otherwise.

### GetCallRecordingOk

`func (o *PhoneNumberWithVoiceSettings) GetCallRecordingOk() (*CallRecording, bool)`

GetCallRecordingOk returns a tuple with the CallRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecording

`func (o *PhoneNumberWithVoiceSettings) SetCallRecording(v CallRecording)`

SetCallRecording sets CallRecording field to given value.

### HasCallRecording

`func (o *PhoneNumberWithVoiceSettings) HasCallRecording() bool`

HasCallRecording returns a boolean if a field has been set.

### GetInboundCallScreening

`func (o *PhoneNumberWithVoiceSettings) GetInboundCallScreening() string`

GetInboundCallScreening returns the InboundCallScreening field if non-nil, zero value otherwise.

### GetInboundCallScreeningOk

`func (o *PhoneNumberWithVoiceSettings) GetInboundCallScreeningOk() (*string, bool)`

GetInboundCallScreeningOk returns a tuple with the InboundCallScreening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallScreening

`func (o *PhoneNumberWithVoiceSettings) SetInboundCallScreening(v string)`

SetInboundCallScreening sets InboundCallScreening field to given value.

### HasInboundCallScreening

`func (o *PhoneNumberWithVoiceSettings) HasInboundCallScreening() bool`

HasInboundCallScreening returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


