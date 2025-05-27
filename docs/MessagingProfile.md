# MessagingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**MmsFallBackToSms** | Pointer to **bool** | enables SMS fallback for MMS messages. | [optional] [default to false]
**MmsTranscoding** | Pointer to **bool** | enables automated resizing of MMS media. | [optional] [default to false]
**Name** | Pointer to **string** | A user friendly name for the messaging profile. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the messaging profile is enabled or not. | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL where webhooks related to this messaging profile will be sent. | [optional] 
**WebhookFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this messaging profile will be sent if sending to the primary URL fails. | [optional] 
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1, v2, or a legacy 2010-04-01 format. | [optional] 
**WhitelistedDestinations** | Pointer to **[]string** | Destinations to which the messaging profile is allowed to send. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 
**V1Secret** | Pointer to **string** | Secret used to authenticate with v1 endpoints. | [optional] 
**NumberPoolSettings** | Pointer to [**NullableNumberPoolSettings**](NumberPoolSettings.md) |  | [optional] 
**UrlShortenerSettings** | Pointer to [**NullableUrlShortenerSettings**](UrlShortenerSettings.md) |  | [optional] 
**AlphaSender** | Pointer to **NullableString** | The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID. | [optional] 
**DailySpendLimit** | Pointer to **string** | The maximum amount of money (in USD) that can be spent by this profile before midnight UTC. | [optional] 
**DailySpendLimitEnabled** | Pointer to **bool** | Whether to enforce the value configured by &#x60;daily_spend_limit&#x60;. | [optional] 

## Methods

### NewMessagingProfile

`func NewMessagingProfile() *MessagingProfile`

NewMessagingProfile instantiates a new MessagingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileWithDefaults

`func NewMessagingProfileWithDefaults() *MessagingProfile`

NewMessagingProfileWithDefaults instantiates a new MessagingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *MessagingProfile) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MessagingProfile) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MessagingProfile) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MessagingProfile) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *MessagingProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessagingProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessagingProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessagingProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMmsFallBackToSms

`func (o *MessagingProfile) GetMmsFallBackToSms() bool`

GetMmsFallBackToSms returns the MmsFallBackToSms field if non-nil, zero value otherwise.

### GetMmsFallBackToSmsOk

`func (o *MessagingProfile) GetMmsFallBackToSmsOk() (*bool, bool)`

GetMmsFallBackToSmsOk returns a tuple with the MmsFallBackToSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsFallBackToSms

`func (o *MessagingProfile) SetMmsFallBackToSms(v bool)`

SetMmsFallBackToSms sets MmsFallBackToSms field to given value.

### HasMmsFallBackToSms

`func (o *MessagingProfile) HasMmsFallBackToSms() bool`

HasMmsFallBackToSms returns a boolean if a field has been set.

### GetMmsTranscoding

`func (o *MessagingProfile) GetMmsTranscoding() bool`

GetMmsTranscoding returns the MmsTranscoding field if non-nil, zero value otherwise.

### GetMmsTranscodingOk

`func (o *MessagingProfile) GetMmsTranscodingOk() (*bool, bool)`

GetMmsTranscodingOk returns a tuple with the MmsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsTranscoding

`func (o *MessagingProfile) SetMmsTranscoding(v bool)`

SetMmsTranscoding sets MmsTranscoding field to given value.

### HasMmsTranscoding

`func (o *MessagingProfile) HasMmsTranscoding() bool`

HasMmsTranscoding returns a boolean if a field has been set.

### GetName

`func (o *MessagingProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessagingProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessagingProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessagingProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *MessagingProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MessagingProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MessagingProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MessagingProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *MessagingProfile) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *MessagingProfile) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *MessagingProfile) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *MessagingProfile) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *MessagingProfile) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *MessagingProfile) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookFailoverUrl

`func (o *MessagingProfile) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *MessagingProfile) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *MessagingProfile) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *MessagingProfile) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### SetWebhookFailoverUrlNil

`func (o *MessagingProfile) SetWebhookFailoverUrlNil(b bool)`

 SetWebhookFailoverUrlNil sets the value for WebhookFailoverUrl to be an explicit nil

### UnsetWebhookFailoverUrl
`func (o *MessagingProfile) UnsetWebhookFailoverUrl()`

UnsetWebhookFailoverUrl ensures that no value is present for WebhookFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *MessagingProfile) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *MessagingProfile) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *MessagingProfile) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *MessagingProfile) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *MessagingProfile) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *MessagingProfile) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *MessagingProfile) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *MessagingProfile) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessagingProfile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessagingProfile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessagingProfile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessagingProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MessagingProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessagingProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessagingProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessagingProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetV1Secret

`func (o *MessagingProfile) GetV1Secret() string`

GetV1Secret returns the V1Secret field if non-nil, zero value otherwise.

### GetV1SecretOk

`func (o *MessagingProfile) GetV1SecretOk() (*string, bool)`

GetV1SecretOk returns a tuple with the V1Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1Secret

`func (o *MessagingProfile) SetV1Secret(v string)`

SetV1Secret sets V1Secret field to given value.

### HasV1Secret

`func (o *MessagingProfile) HasV1Secret() bool`

HasV1Secret returns a boolean if a field has been set.

### GetNumberPoolSettings

`func (o *MessagingProfile) GetNumberPoolSettings() NumberPoolSettings`

GetNumberPoolSettings returns the NumberPoolSettings field if non-nil, zero value otherwise.

### GetNumberPoolSettingsOk

`func (o *MessagingProfile) GetNumberPoolSettingsOk() (*NumberPoolSettings, bool)`

GetNumberPoolSettingsOk returns a tuple with the NumberPoolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberPoolSettings

`func (o *MessagingProfile) SetNumberPoolSettings(v NumberPoolSettings)`

SetNumberPoolSettings sets NumberPoolSettings field to given value.

### HasNumberPoolSettings

`func (o *MessagingProfile) HasNumberPoolSettings() bool`

HasNumberPoolSettings returns a boolean if a field has been set.

### SetNumberPoolSettingsNil

`func (o *MessagingProfile) SetNumberPoolSettingsNil(b bool)`

 SetNumberPoolSettingsNil sets the value for NumberPoolSettings to be an explicit nil

### UnsetNumberPoolSettings
`func (o *MessagingProfile) UnsetNumberPoolSettings()`

UnsetNumberPoolSettings ensures that no value is present for NumberPoolSettings, not even an explicit nil
### GetUrlShortenerSettings

`func (o *MessagingProfile) GetUrlShortenerSettings() UrlShortenerSettings`

GetUrlShortenerSettings returns the UrlShortenerSettings field if non-nil, zero value otherwise.

### GetUrlShortenerSettingsOk

`func (o *MessagingProfile) GetUrlShortenerSettingsOk() (*UrlShortenerSettings, bool)`

GetUrlShortenerSettingsOk returns a tuple with the UrlShortenerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlShortenerSettings

`func (o *MessagingProfile) SetUrlShortenerSettings(v UrlShortenerSettings)`

SetUrlShortenerSettings sets UrlShortenerSettings field to given value.

### HasUrlShortenerSettings

`func (o *MessagingProfile) HasUrlShortenerSettings() bool`

HasUrlShortenerSettings returns a boolean if a field has been set.

### SetUrlShortenerSettingsNil

`func (o *MessagingProfile) SetUrlShortenerSettingsNil(b bool)`

 SetUrlShortenerSettingsNil sets the value for UrlShortenerSettings to be an explicit nil

### UnsetUrlShortenerSettings
`func (o *MessagingProfile) UnsetUrlShortenerSettings()`

UnsetUrlShortenerSettings ensures that no value is present for UrlShortenerSettings, not even an explicit nil
### GetAlphaSender

`func (o *MessagingProfile) GetAlphaSender() string`

GetAlphaSender returns the AlphaSender field if non-nil, zero value otherwise.

### GetAlphaSenderOk

`func (o *MessagingProfile) GetAlphaSenderOk() (*string, bool)`

GetAlphaSenderOk returns a tuple with the AlphaSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaSender

`func (o *MessagingProfile) SetAlphaSender(v string)`

SetAlphaSender sets AlphaSender field to given value.

### HasAlphaSender

`func (o *MessagingProfile) HasAlphaSender() bool`

HasAlphaSender returns a boolean if a field has been set.

### SetAlphaSenderNil

`func (o *MessagingProfile) SetAlphaSenderNil(b bool)`

 SetAlphaSenderNil sets the value for AlphaSender to be an explicit nil

### UnsetAlphaSender
`func (o *MessagingProfile) UnsetAlphaSender()`

UnsetAlphaSender ensures that no value is present for AlphaSender, not even an explicit nil
### GetDailySpendLimit

`func (o *MessagingProfile) GetDailySpendLimit() string`

GetDailySpendLimit returns the DailySpendLimit field if non-nil, zero value otherwise.

### GetDailySpendLimitOk

`func (o *MessagingProfile) GetDailySpendLimitOk() (*string, bool)`

GetDailySpendLimitOk returns a tuple with the DailySpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimit

`func (o *MessagingProfile) SetDailySpendLimit(v string)`

SetDailySpendLimit sets DailySpendLimit field to given value.

### HasDailySpendLimit

`func (o *MessagingProfile) HasDailySpendLimit() bool`

HasDailySpendLimit returns a boolean if a field has been set.

### GetDailySpendLimitEnabled

`func (o *MessagingProfile) GetDailySpendLimitEnabled() bool`

GetDailySpendLimitEnabled returns the DailySpendLimitEnabled field if non-nil, zero value otherwise.

### GetDailySpendLimitEnabledOk

`func (o *MessagingProfile) GetDailySpendLimitEnabledOk() (*bool, bool)`

GetDailySpendLimitEnabledOk returns a tuple with the DailySpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimitEnabled

`func (o *MessagingProfile) SetDailySpendLimitEnabled(v bool)`

SetDailySpendLimitEnabled sets DailySpendLimitEnabled field to given value.

### HasDailySpendLimitEnabled

`func (o *MessagingProfile) HasDailySpendLimitEnabled() bool`

HasDailySpendLimitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


