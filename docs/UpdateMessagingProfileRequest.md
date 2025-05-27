# UpdateMessagingProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**Name** | Pointer to **string** | A user friendly name for the messaging profile. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the messaging profile is enabled or not. | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL where webhooks related to this messaging profile will be sent. | [optional] 
**WebhookFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this messaging profile will be sent if sending to the primary URL fails. | [optional] 
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1, v2, or a legacy 2010-04-01 format. | [optional] 
**WhitelistedDestinations** | Pointer to **[]string** | Destinations to which the messaging profile is allowed to send. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed.  This field is required if the messaging profile doesn&#39;t have it defined yet. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 
**V1Secret** | Pointer to **string** | Secret used to authenticate with v1 endpoints. | [optional] 
**NumberPoolSettings** | Pointer to [**NullableNumberPoolSettings**](NumberPoolSettings.md) |  | [optional] 
**UrlShortenerSettings** | Pointer to [**NullableUrlShortenerSettings**](UrlShortenerSettings.md) |  | [optional] 
**AlphaSender** | Pointer to **NullableString** | The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID. | [optional] 
**DailySpendLimit** | Pointer to **string** | The maximum amount of money (in USD) that can be spent by this profile before midnight UTC. | [optional] 
**DailySpendLimitEnabled** | Pointer to **bool** | Whether to enforce the value configured by &#x60;daily_spend_limit&#x60;. | [optional] 
**MmsFallBackToSms** | Pointer to **bool** | enables SMS fallback for MMS messages. | [optional] [default to false]
**MmsTranscoding** | Pointer to **bool** | enables automated resizing of MMS media. | [optional] [default to false]

## Methods

### NewUpdateMessagingProfileRequest

`func NewUpdateMessagingProfileRequest() *UpdateMessagingProfileRequest`

NewUpdateMessagingProfileRequest instantiates a new UpdateMessagingProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMessagingProfileRequestWithDefaults

`func NewUpdateMessagingProfileRequestWithDefaults() *UpdateMessagingProfileRequest`

NewUpdateMessagingProfileRequestWithDefaults instantiates a new UpdateMessagingProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *UpdateMessagingProfileRequest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *UpdateMessagingProfileRequest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *UpdateMessagingProfileRequest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *UpdateMessagingProfileRequest) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *UpdateMessagingProfileRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMessagingProfileRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMessagingProfileRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMessagingProfileRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateMessagingProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMessagingProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMessagingProfileRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMessagingProfileRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateMessagingProfileRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMessagingProfileRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMessagingProfileRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateMessagingProfileRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UpdateMessagingProfileRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateMessagingProfileRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateMessagingProfileRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateMessagingProfileRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *UpdateMessagingProfileRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *UpdateMessagingProfileRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookFailoverUrl

`func (o *UpdateMessagingProfileRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *UpdateMessagingProfileRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *UpdateMessagingProfileRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *UpdateMessagingProfileRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### SetWebhookFailoverUrlNil

`func (o *UpdateMessagingProfileRequest) SetWebhookFailoverUrlNil(b bool)`

 SetWebhookFailoverUrlNil sets the value for WebhookFailoverUrl to be an explicit nil

### UnsetWebhookFailoverUrl
`func (o *UpdateMessagingProfileRequest) UnsetWebhookFailoverUrl()`

UnsetWebhookFailoverUrl ensures that no value is present for WebhookFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *UpdateMessagingProfileRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *UpdateMessagingProfileRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *UpdateMessagingProfileRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *UpdateMessagingProfileRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *UpdateMessagingProfileRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *UpdateMessagingProfileRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *UpdateMessagingProfileRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *UpdateMessagingProfileRequest) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateMessagingProfileRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateMessagingProfileRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateMessagingProfileRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateMessagingProfileRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UpdateMessagingProfileRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateMessagingProfileRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateMessagingProfileRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateMessagingProfileRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetV1Secret

`func (o *UpdateMessagingProfileRequest) GetV1Secret() string`

GetV1Secret returns the V1Secret field if non-nil, zero value otherwise.

### GetV1SecretOk

`func (o *UpdateMessagingProfileRequest) GetV1SecretOk() (*string, bool)`

GetV1SecretOk returns a tuple with the V1Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1Secret

`func (o *UpdateMessagingProfileRequest) SetV1Secret(v string)`

SetV1Secret sets V1Secret field to given value.

### HasV1Secret

`func (o *UpdateMessagingProfileRequest) HasV1Secret() bool`

HasV1Secret returns a boolean if a field has been set.

### GetNumberPoolSettings

`func (o *UpdateMessagingProfileRequest) GetNumberPoolSettings() NumberPoolSettings`

GetNumberPoolSettings returns the NumberPoolSettings field if non-nil, zero value otherwise.

### GetNumberPoolSettingsOk

`func (o *UpdateMessagingProfileRequest) GetNumberPoolSettingsOk() (*NumberPoolSettings, bool)`

GetNumberPoolSettingsOk returns a tuple with the NumberPoolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberPoolSettings

`func (o *UpdateMessagingProfileRequest) SetNumberPoolSettings(v NumberPoolSettings)`

SetNumberPoolSettings sets NumberPoolSettings field to given value.

### HasNumberPoolSettings

`func (o *UpdateMessagingProfileRequest) HasNumberPoolSettings() bool`

HasNumberPoolSettings returns a boolean if a field has been set.

### SetNumberPoolSettingsNil

`func (o *UpdateMessagingProfileRequest) SetNumberPoolSettingsNil(b bool)`

 SetNumberPoolSettingsNil sets the value for NumberPoolSettings to be an explicit nil

### UnsetNumberPoolSettings
`func (o *UpdateMessagingProfileRequest) UnsetNumberPoolSettings()`

UnsetNumberPoolSettings ensures that no value is present for NumberPoolSettings, not even an explicit nil
### GetUrlShortenerSettings

`func (o *UpdateMessagingProfileRequest) GetUrlShortenerSettings() UrlShortenerSettings`

GetUrlShortenerSettings returns the UrlShortenerSettings field if non-nil, zero value otherwise.

### GetUrlShortenerSettingsOk

`func (o *UpdateMessagingProfileRequest) GetUrlShortenerSettingsOk() (*UrlShortenerSettings, bool)`

GetUrlShortenerSettingsOk returns a tuple with the UrlShortenerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlShortenerSettings

`func (o *UpdateMessagingProfileRequest) SetUrlShortenerSettings(v UrlShortenerSettings)`

SetUrlShortenerSettings sets UrlShortenerSettings field to given value.

### HasUrlShortenerSettings

`func (o *UpdateMessagingProfileRequest) HasUrlShortenerSettings() bool`

HasUrlShortenerSettings returns a boolean if a field has been set.

### SetUrlShortenerSettingsNil

`func (o *UpdateMessagingProfileRequest) SetUrlShortenerSettingsNil(b bool)`

 SetUrlShortenerSettingsNil sets the value for UrlShortenerSettings to be an explicit nil

### UnsetUrlShortenerSettings
`func (o *UpdateMessagingProfileRequest) UnsetUrlShortenerSettings()`

UnsetUrlShortenerSettings ensures that no value is present for UrlShortenerSettings, not even an explicit nil
### GetAlphaSender

`func (o *UpdateMessagingProfileRequest) GetAlphaSender() string`

GetAlphaSender returns the AlphaSender field if non-nil, zero value otherwise.

### GetAlphaSenderOk

`func (o *UpdateMessagingProfileRequest) GetAlphaSenderOk() (*string, bool)`

GetAlphaSenderOk returns a tuple with the AlphaSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaSender

`func (o *UpdateMessagingProfileRequest) SetAlphaSender(v string)`

SetAlphaSender sets AlphaSender field to given value.

### HasAlphaSender

`func (o *UpdateMessagingProfileRequest) HasAlphaSender() bool`

HasAlphaSender returns a boolean if a field has been set.

### SetAlphaSenderNil

`func (o *UpdateMessagingProfileRequest) SetAlphaSenderNil(b bool)`

 SetAlphaSenderNil sets the value for AlphaSender to be an explicit nil

### UnsetAlphaSender
`func (o *UpdateMessagingProfileRequest) UnsetAlphaSender()`

UnsetAlphaSender ensures that no value is present for AlphaSender, not even an explicit nil
### GetDailySpendLimit

`func (o *UpdateMessagingProfileRequest) GetDailySpendLimit() string`

GetDailySpendLimit returns the DailySpendLimit field if non-nil, zero value otherwise.

### GetDailySpendLimitOk

`func (o *UpdateMessagingProfileRequest) GetDailySpendLimitOk() (*string, bool)`

GetDailySpendLimitOk returns a tuple with the DailySpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimit

`func (o *UpdateMessagingProfileRequest) SetDailySpendLimit(v string)`

SetDailySpendLimit sets DailySpendLimit field to given value.

### HasDailySpendLimit

`func (o *UpdateMessagingProfileRequest) HasDailySpendLimit() bool`

HasDailySpendLimit returns a boolean if a field has been set.

### GetDailySpendLimitEnabled

`func (o *UpdateMessagingProfileRequest) GetDailySpendLimitEnabled() bool`

GetDailySpendLimitEnabled returns the DailySpendLimitEnabled field if non-nil, zero value otherwise.

### GetDailySpendLimitEnabledOk

`func (o *UpdateMessagingProfileRequest) GetDailySpendLimitEnabledOk() (*bool, bool)`

GetDailySpendLimitEnabledOk returns a tuple with the DailySpendLimitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySpendLimitEnabled

`func (o *UpdateMessagingProfileRequest) SetDailySpendLimitEnabled(v bool)`

SetDailySpendLimitEnabled sets DailySpendLimitEnabled field to given value.

### HasDailySpendLimitEnabled

`func (o *UpdateMessagingProfileRequest) HasDailySpendLimitEnabled() bool`

HasDailySpendLimitEnabled returns a boolean if a field has been set.

### GetMmsFallBackToSms

`func (o *UpdateMessagingProfileRequest) GetMmsFallBackToSms() bool`

GetMmsFallBackToSms returns the MmsFallBackToSms field if non-nil, zero value otherwise.

### GetMmsFallBackToSmsOk

`func (o *UpdateMessagingProfileRequest) GetMmsFallBackToSmsOk() (*bool, bool)`

GetMmsFallBackToSmsOk returns a tuple with the MmsFallBackToSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsFallBackToSms

`func (o *UpdateMessagingProfileRequest) SetMmsFallBackToSms(v bool)`

SetMmsFallBackToSms sets MmsFallBackToSms field to given value.

### HasMmsFallBackToSms

`func (o *UpdateMessagingProfileRequest) HasMmsFallBackToSms() bool`

HasMmsFallBackToSms returns a boolean if a field has been set.

### GetMmsTranscoding

`func (o *UpdateMessagingProfileRequest) GetMmsTranscoding() bool`

GetMmsTranscoding returns the MmsTranscoding field if non-nil, zero value otherwise.

### GetMmsTranscodingOk

`func (o *UpdateMessagingProfileRequest) GetMmsTranscodingOk() (*bool, bool)`

GetMmsTranscodingOk returns a tuple with the MmsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsTranscoding

`func (o *UpdateMessagingProfileRequest) SetMmsTranscoding(v bool)`

SetMmsTranscoding sets MmsTranscoding field to given value.

### HasMmsTranscoding

`func (o *UpdateMessagingProfileRequest) HasMmsTranscoding() bool`

HasMmsTranscoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


