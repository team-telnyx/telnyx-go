# CreateGroupMMSMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Phone number, in +E.164 format, used to send the message. | 
**To** | **[]string** | A list of destinations. No more than 8 destinations are allowed. | 
**Text** | Pointer to **string** | Message body (i.e., content) as a non-empty string. | [optional] 
**Subject** | Pointer to **string** | Subject of multimedia message | [optional] 
**MediaUrls** | Pointer to **[]string** | A list of media URLs. The total media size must be less than 1 MB. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL where webhooks related to this message will be sent. | [optional] 
**WebhookFailoverUrl** | Pointer to **string** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] 
**UseProfileWebhooks** | Pointer to **bool** | If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile. | [optional] [default to true]

## Methods

### NewCreateGroupMMSMessageRequest

`func NewCreateGroupMMSMessageRequest(from string, to []string, ) *CreateGroupMMSMessageRequest`

NewCreateGroupMMSMessageRequest instantiates a new CreateGroupMMSMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupMMSMessageRequestWithDefaults

`func NewCreateGroupMMSMessageRequestWithDefaults() *CreateGroupMMSMessageRequest`

NewCreateGroupMMSMessageRequestWithDefaults instantiates a new CreateGroupMMSMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreateGroupMMSMessageRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateGroupMMSMessageRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateGroupMMSMessageRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CreateGroupMMSMessageRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateGroupMMSMessageRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateGroupMMSMessageRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetText

`func (o *CreateGroupMMSMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateGroupMMSMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateGroupMMSMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateGroupMMSMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSubject

`func (o *CreateGroupMMSMessageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateGroupMMSMessageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateGroupMMSMessageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateGroupMMSMessageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMediaUrls

`func (o *CreateGroupMMSMessageRequest) GetMediaUrls() []string`

GetMediaUrls returns the MediaUrls field if non-nil, zero value otherwise.

### GetMediaUrlsOk

`func (o *CreateGroupMMSMessageRequest) GetMediaUrlsOk() (*[]string, bool)`

GetMediaUrlsOk returns a tuple with the MediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrls

`func (o *CreateGroupMMSMessageRequest) SetMediaUrls(v []string)`

SetMediaUrls sets MediaUrls field to given value.

### HasMediaUrls

`func (o *CreateGroupMMSMessageRequest) HasMediaUrls() bool`

HasMediaUrls returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateGroupMMSMessageRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateGroupMMSMessageRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateGroupMMSMessageRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateGroupMMSMessageRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *CreateGroupMMSMessageRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateGroupMMSMessageRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateGroupMMSMessageRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateGroupMMSMessageRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetUseProfileWebhooks

`func (o *CreateGroupMMSMessageRequest) GetUseProfileWebhooks() bool`

GetUseProfileWebhooks returns the UseProfileWebhooks field if non-nil, zero value otherwise.

### GetUseProfileWebhooksOk

`func (o *CreateGroupMMSMessageRequest) GetUseProfileWebhooksOk() (*bool, bool)`

GetUseProfileWebhooksOk returns a tuple with the UseProfileWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProfileWebhooks

`func (o *CreateGroupMMSMessageRequest) SetUseProfileWebhooks(v bool)`

SetUseProfileWebhooks sets UseProfileWebhooks field to given value.

### HasUseProfileWebhooks

`func (o *CreateGroupMMSMessageRequest) HasUseProfileWebhooks() bool`

HasUseProfileWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


