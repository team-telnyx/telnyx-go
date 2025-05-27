# CreateNumberPoolMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProfileId** | **string** | Unique identifier for a messaging profile. | 
**To** | **string** | Receiving address (+E.164 formatted phone number or short code). | 
**Text** | Pointer to **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] 
**Subject** | Pointer to **string** | Subject of multimedia message | [optional] 
**MediaUrls** | Pointer to **[]string** | A list of media URLs. The total media size must be less than 1 MB.  **Required for MMS** | [optional] 
**WebhookUrl** | Pointer to **string** | The URL where webhooks related to this message will be sent. | [optional] 
**WebhookFailoverUrl** | Pointer to **string** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] 
**UseProfileWebhooks** | Pointer to **bool** | If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile. | [optional] [default to true]
**Type** | Pointer to **string** | The protocol for sending the message, either SMS or MMS. | [optional] 
**AutoDetect** | Pointer to **bool** | Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts. | [optional] [default to false]

## Methods

### NewCreateNumberPoolMessageRequest

`func NewCreateNumberPoolMessageRequest(messagingProfileId string, to string, ) *CreateNumberPoolMessageRequest`

NewCreateNumberPoolMessageRequest instantiates a new CreateNumberPoolMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNumberPoolMessageRequestWithDefaults

`func NewCreateNumberPoolMessageRequestWithDefaults() *CreateNumberPoolMessageRequest`

NewCreateNumberPoolMessageRequestWithDefaults instantiates a new CreateNumberPoolMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProfileId

`func (o *CreateNumberPoolMessageRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *CreateNumberPoolMessageRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *CreateNumberPoolMessageRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.


### GetTo

`func (o *CreateNumberPoolMessageRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateNumberPoolMessageRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateNumberPoolMessageRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetText

`func (o *CreateNumberPoolMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateNumberPoolMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateNumberPoolMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateNumberPoolMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSubject

`func (o *CreateNumberPoolMessageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateNumberPoolMessageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateNumberPoolMessageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateNumberPoolMessageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMediaUrls

`func (o *CreateNumberPoolMessageRequest) GetMediaUrls() []string`

GetMediaUrls returns the MediaUrls field if non-nil, zero value otherwise.

### GetMediaUrlsOk

`func (o *CreateNumberPoolMessageRequest) GetMediaUrlsOk() (*[]string, bool)`

GetMediaUrlsOk returns a tuple with the MediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrls

`func (o *CreateNumberPoolMessageRequest) SetMediaUrls(v []string)`

SetMediaUrls sets MediaUrls field to given value.

### HasMediaUrls

`func (o *CreateNumberPoolMessageRequest) HasMediaUrls() bool`

HasMediaUrls returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateNumberPoolMessageRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateNumberPoolMessageRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateNumberPoolMessageRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateNumberPoolMessageRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *CreateNumberPoolMessageRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateNumberPoolMessageRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateNumberPoolMessageRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateNumberPoolMessageRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetUseProfileWebhooks

`func (o *CreateNumberPoolMessageRequest) GetUseProfileWebhooks() bool`

GetUseProfileWebhooks returns the UseProfileWebhooks field if non-nil, zero value otherwise.

### GetUseProfileWebhooksOk

`func (o *CreateNumberPoolMessageRequest) GetUseProfileWebhooksOk() (*bool, bool)`

GetUseProfileWebhooksOk returns a tuple with the UseProfileWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProfileWebhooks

`func (o *CreateNumberPoolMessageRequest) SetUseProfileWebhooks(v bool)`

SetUseProfileWebhooks sets UseProfileWebhooks field to given value.

### HasUseProfileWebhooks

`func (o *CreateNumberPoolMessageRequest) HasUseProfileWebhooks() bool`

HasUseProfileWebhooks returns a boolean if a field has been set.

### GetType

`func (o *CreateNumberPoolMessageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNumberPoolMessageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNumberPoolMessageRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNumberPoolMessageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoDetect

`func (o *CreateNumberPoolMessageRequest) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *CreateNumberPoolMessageRequest) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *CreateNumberPoolMessageRequest) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *CreateNumberPoolMessageRequest) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


