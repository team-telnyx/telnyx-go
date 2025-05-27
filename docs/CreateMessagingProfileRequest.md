# CreateMessagingProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user friendly name for the messaging profile. | 
**WhitelistedDestinations** | **[]string** | Destinations to which the messaging profile is allowed to send. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60; all destinations will be allowed. | 
**Enabled** | Pointer to **bool** | Specifies whether the messaging profile is enabled or not. | [optional] [default to true]
**WebhookUrl** | Pointer to **NullableString** | The URL where webhooks related to this messaging profile will be sent. | [optional] [default to ""]
**WebhookFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this messaging profile will be sent if sending to the primary URL fails. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1, v2, or a legacy 2010-04-01 format. | [optional] [default to "2"]
**NumberPoolSettings** | Pointer to [**NullableNumberPoolSettings**](NumberPoolSettings.md) |  | [optional] 
**UrlShortenerSettings** | Pointer to [**NullableUrlShortenerSettings**](UrlShortenerSettings.md) |  | [optional] 
**AlphaSender** | Pointer to **NullableString** | The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID. | [optional] 
**DailySpendLimit** | Pointer to **string** | The maximum amount of money (in USD) that can be spent by this profile before midnight UTC. | [optional] 
**DailySpendLimitEnabled** | Pointer to **bool** | Whether to enforce the value configured by &#x60;daily_spend_limit&#x60;. | [optional] 
**MmsFallBackToSms** | Pointer to **bool** | enables SMS fallback for MMS messages. | [optional] [default to false]
**MmsTranscoding** | Pointer to **bool** | enables automated resizing of MMS media. | [optional] [default to false]

## Methods

### NewCreateMessagingProfileRequest

`func NewCreateMessagingProfileRequest(name string, whitelistedDestinations []string, ) *CreateMessagingProfileRequest`

NewCreateMessagingProfileRequest instantiates a new CreateMessagingProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessagingProfileRequestWithDefaults

`func NewCreateMessagingProfileRequestWithDefaults() *CreateMessagingProfileRequest`

NewCreateMessagingProfileRequestWithDefaults instantiates a new CreateMessagingProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMessagingProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMessagingProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMessagingProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWhitelistedDestinations

`func (o *CreateMessagingProfileRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *CreateMessagingProfileRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *CreateMessagingProfileRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.


### GetEnabled

`func (o *CreateMessagingProfileRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateMessagingProfileRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateMessagingProfileRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateMessagingProfileRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateMessagingProfileRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateMessagingProfileRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateMessagingProfileRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateMessagingProfileRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *CreateMessagingProfileRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *CreateMessagingProfileRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookFailoverUrl

`func (o *CreateMessagingProfileRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateMessagingProfileRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateMessagingProfileRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateMessagingProfileRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### SetWebhookFailoverUrlNil

`func (o *CreateMessagingProfileRequest) SetWebhookFailoverUrlNil(b bool)`

 SetWebhookFailoverUrlNil sets the value for WebhookFailoverUrl to be an explicit nil

### UnsetWebhookFailoverUrl
`func (o *CreateMessagingProfileRequest) UnsetWebhookFailoverUrl()`

UnsetWebhookFailoverUrl ensures that no value is present for WebhookFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *CreateMessagingProfileRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *CreateMessagingProfileRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *CreateMessagingProfileRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *CreateMessagingProfileRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetNumberPoolSettings

`func (o *CreateMessagingProfileRequest) GetNumberPoolSettings() NumberPoolSettings`

GetNumberPoolSettings returns the NumberPoolSettings field if non-nil, zero value otherwise.

### GetNumberPoolSettingsOk

`func (o *CreateMessagingProfileRequest) GetNumberPoolSettingsOk() (*NumberPoolSettings, bool)`

GetNumberPoolSettingsOk returns a tuple with the NumberPoolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberPoolSettings

`func (o *CreateMessagingProfileRequest) SetNumberPoolSettings(v NumberPoolSettings)`

SetNumberPoolSettings sets NumberPoolSettings field to given value.

### HasNumberPoolSettings

`func (o *CreateMessagingProfileRequest) HasNumberPoolSettings() bool`

HasNumberPoolSettings returns a boolean if a field has been set.

### SetNumberPoolSettingsNil

`func (o *CreateMessagingProfileRequest) SetNumberPoolSettingsNil(b bool)`

 SetNumberPoolSettingsNil sets the value for NumberPoolSettings to be an explicit nil

### UnsetNumberPoolSettings
`func (o *CreateMessagingProfileRequest) UnsetNumberPoolSettings()`

UnsetNumberPoolSettings ensures that no value is present for NumberPoolSettings, not even an explicit nil
### GetUrlShortenerSettings

`func (o *CreateMessagingProfileRequest) GetUrlShortenerSettings() UrlShortenerSettings`

GetUrlShortenerSettings returns the UrlShortenerSettings field if non-nil, zero value otherwise.

### GetUrlShortenerSettingsOk

`func (o *CreateMessagingProfileRequest) GetUrlShortenerSettingsOk() (*UrlShortenerSettings, bool)`

GetUrlShortenerSettingsOk returns a tuple with the UrlShortenerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlShortenerSettings

`func (o *CreateMessagingProfileRequest) SetUrlShortenerSettings(v UrlShortenerSettings)`

SetUrlShortenerSettings sets UrlShortenerSettings field to given value.

### HasUrlShortenerSettings

`func (o *CreateMessagingProfileRequest) HasUrlShortenerSettings() bool`

HasUrlShortenerSettings returns a boolean if a field has been set.

### SetUrlShortenerSettingsNil

`func (o *CreateMessagingProfileRequest) SetUrlShortenerSettingsNil(b bool)`

 SetUrlShortenerSettingsNil sets the value for UrlShortenerSettings to be an explicit nil

### UnsetUrlShortenerSettings
`func (o *CreateMessagingProfileRequest) UnsetUrlShortenerSettings()`

UnsetUrlShortenerSettings ensures that no value is present for UrlShortenerSettings, not even an explicit nil
### GetAlphaSender

`func (o *CreateMessagingProfileRequest) GetAlphaSender() string`

GetAlphaSender returns the AlphaSender field if non-nil, zero value otherwise.

### GetAlphaSenderOk

`func (o *CreateMessagingProfileRequest) GetAlphaSenderOk() (*string, bool)`

GetAlphaSenderOk returns a tuple with the AlphaSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaSender

`func (o *CreateMessagingProfileRequest) SetAlphaSender(v string)`

SetAlphaSender sets AlphaSender field to given value.

### HasAlphaSender

`func (o *CreateMessagingProfileRequest) HasAlphaSender() bool`

HasAlphaSender returns a boolean if a field has been set.

### SetAlphaSenderNil

`func (o *CreateMessagingProfileRequest) SetAlphaSenderNil(b bool)`

 SetAlphaSenderNil sets the value for AlphaSender to be an explicit nil

### UnsetAlphaSender
`func (o *CreateMessagingProfileRequest) UnsetAlphaSender()`

UnsetAlphaSender ensures that no value is present for AlphaSender, not even an explicit nil
### GetDailySpendLimit

`func (o *CreateMessagingProfileRequest) GetDailySpendLimit() string`

GetDailySpendLimit returns the DailySpendLimit field if non-nil, zero value otherwise.

### GetDailySpendLimitOk

`func (o *CreateMessagingProfileRequest) GetDailySpendLimitOk() (*string, bool)`

GetDailySpendLimitOk returns a tuple with the DailySpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimit

`func (o *CreateMessagingProfileRequest) SetDailySpendLimit(v string)`

SetDailySpendLimit sets DailySpendLimit field to given value.

### HasDailySpendLimit

`func (o *CreateMessagingProfileRequest) HasDailySpendLimit() bool`

HasDailySpendLimit returns a boolean if a field has been set.

### GetDailySpendLimitEnabled

`func (o *CreateMessagingProfileRequest) GetDailySpendLimitEnabled() bool`

GetDailySpendLimitEnabled returns the DailySpendLimitEnabled field if non-nil, zero value otherwise.

### GetDailySpendLimitEnabledOk

`func (o *CreateMessagingProfileRequest) GetDailySpendLimitEnabledOk() (*bool, bool)`

GetDailySpendLimitEnabledOk returns a tuple with the DailySpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimitEnabled

`func (o *CreateMessagingProfileRequest) SetDailySpendLimitEnabled(v bool)`

SetDailySpendLimitEnabled sets DailySpendLimitEnabled field to given value.

### HasDailySpendLimitEnabled

`func (o *CreateMessagingProfileRequest) HasDailySpendLimitEnabled() bool`

HasDailySpendLimitEnabled returns a boolean if a field has been set.

### GetMmsFallBackToSms

`func (o *CreateMessagingProfileRequest) GetMmsFallBackToSms() bool`

GetMmsFallBackToSms returns the MmsFallBackToSms field if non-nil, zero value otherwise.

### GetMmsFallBackToSmsOk

`func (o *CreateMessagingProfileRequest) GetMmsFallBackToSmsOk() (*bool, bool)`

GetMmsFallBackToSmsOk returns a tuple with the MmsFallBackToSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsFallBackToSms

`func (o *CreateMessagingProfileRequest) SetMmsFallBackToSms(v bool)`

SetMmsFallBackToSms sets MmsFallBackToSms field to given value.

### HasMmsFallBackToSms

`func (o *CreateMessagingProfileRequest) HasMmsFallBackToSms() bool`

HasMmsFallBackToSms returns a boolean if a field has been set.

### GetMmsTranscoding

`func (o *CreateMessagingProfileRequest) GetMmsTranscoding() bool`

GetMmsTranscoding returns the MmsTranscoding field if non-nil, zero value otherwise.

### GetMmsTranscodingOk

`func (o *CreateMessagingProfileRequest) GetMmsTranscodingOk() (*bool, bool)`

GetMmsTranscodingOk returns a tuple with the MmsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsTranscoding

`func (o *CreateMessagingProfileRequest) SetMmsTranscoding(v bool)`

SetMmsTranscoding sets MmsTranscoding field to given value.

### HasMmsTranscoding

`func (o *CreateMessagingProfileRequest) HasMmsTranscoding() bool`

HasMmsTranscoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


