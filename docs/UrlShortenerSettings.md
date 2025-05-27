# UrlShortenerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | One of the domains provided by the Telnyx URL shortener service.  | 
**Prefix** | Pointer to **string** | Optional prefix that can be used to identify your brand, and will appear in the Telnyx generated URLs after the domain name.  | [optional] 
**ReplaceBlacklistOnly** | Pointer to **bool** | Use the link replacement tool only for links that are specifically blacklisted by Telnyx.  | [optional] 
**SendWebhooks** | Pointer to **bool** | Receive webhooks for when your replaced links are clicked. Webhooks are sent to the webhooks on the messaging profile.  | [optional] 

## Methods

### NewUrlShortenerSettings

`func NewUrlShortenerSettings(domain string, ) *UrlShortenerSettings`

NewUrlShortenerSettings instantiates a new UrlShortenerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlShortenerSettingsWithDefaults

`func NewUrlShortenerSettingsWithDefaults() *UrlShortenerSettings`

NewUrlShortenerSettingsWithDefaults instantiates a new UrlShortenerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *UrlShortenerSettings) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UrlShortenerSettings) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UrlShortenerSettings) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPrefix

`func (o *UrlShortenerSettings) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UrlShortenerSettings) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UrlShortenerSettings) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UrlShortenerSettings) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReplaceBlacklistOnly

`func (o *UrlShortenerSettings) GetReplaceBlacklistOnly() bool`

GetReplaceBlacklistOnly returns the ReplaceBlacklistOnly field if non-nil, zero value otherwise.

### GetReplaceBlacklistOnlyOk

`func (o *UrlShortenerSettings) GetReplaceBlacklistOnlyOk() (*bool, bool)`

GetReplaceBlacklistOnlyOk returns a tuple with the ReplaceBlacklistOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceBlacklistOnly

`func (o *UrlShortenerSettings) SetReplaceBlacklistOnly(v bool)`

SetReplaceBlacklistOnly sets ReplaceBlacklistOnly field to given value.

### HasReplaceBlacklistOnly

`func (o *UrlShortenerSettings) HasReplaceBlacklistOnly() bool`

HasReplaceBlacklistOnly returns a boolean if a field has been set.

### GetSendWebhooks

`func (o *UrlShortenerSettings) GetSendWebhooks() bool`

GetSendWebhooks returns the SendWebhooks field if non-nil, zero value otherwise.

### GetSendWebhooksOk

`func (o *UrlShortenerSettings) GetSendWebhooksOk() (*bool, bool)`

GetSendWebhooksOk returns a tuple with the SendWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWebhooks

`func (o *UrlShortenerSettings) SetSendWebhooks(v bool)`

SetSendWebhooks sets SendWebhooks field to given value.

### HasSendWebhooks

`func (o *UrlShortenerSettings) HasSendWebhooks() bool`

HasSendWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


