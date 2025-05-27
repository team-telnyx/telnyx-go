# CreateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code).  **Required if sending with a phone number, short code, or alphanumeric sender ID.**  | [optional] 
**MessagingProfileId** | Pointer to **string** | Unique identifier for a messaging profile.  **Required if sending via number pool or with an alphanumeric sender ID.**  | [optional] 
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

### NewCreateMessageRequest

`func NewCreateMessageRequest(to string, ) *CreateMessageRequest`

NewCreateMessageRequest instantiates a new CreateMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessageRequestWithDefaults

`func NewCreateMessageRequestWithDefaults() *CreateMessageRequest`

NewCreateMessageRequestWithDefaults instantiates a new CreateMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreateMessageRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateMessageRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateMessageRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreateMessageRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *CreateMessageRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *CreateMessageRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *CreateMessageRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *CreateMessageRequest) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetTo

`func (o *CreateMessageRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateMessageRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateMessageRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetText

`func (o *CreateMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSubject

`func (o *CreateMessageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateMessageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateMessageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateMessageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMediaUrls

`func (o *CreateMessageRequest) GetMediaUrls() []string`

GetMediaUrls returns the MediaUrls field if non-nil, zero value otherwise.

### GetMediaUrlsOk

`func (o *CreateMessageRequest) GetMediaUrlsOk() (*[]string, bool)`

GetMediaUrlsOk returns a tuple with the MediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrls

`func (o *CreateMessageRequest) SetMediaUrls(v []string)`

SetMediaUrls sets MediaUrls field to given value.

### HasMediaUrls

`func (o *CreateMessageRequest) HasMediaUrls() bool`

HasMediaUrls returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateMessageRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateMessageRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateMessageRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateMessageRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *CreateMessageRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateMessageRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateMessageRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateMessageRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetUseProfileWebhooks

`func (o *CreateMessageRequest) GetUseProfileWebhooks() bool`

GetUseProfileWebhooks returns the UseProfileWebhooks field if non-nil, zero value otherwise.

### GetUseProfileWebhooksOk

`func (o *CreateMessageRequest) GetUseProfileWebhooksOk() (*bool, bool)`

GetUseProfileWebhooksOk returns a tuple with the UseProfileWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProfileWebhooks

`func (o *CreateMessageRequest) SetUseProfileWebhooks(v bool)`

SetUseProfileWebhooks sets UseProfileWebhooks field to given value.

### HasUseProfileWebhooks

`func (o *CreateMessageRequest) HasUseProfileWebhooks() bool`

HasUseProfileWebhooks returns a boolean if a field has been set.

### GetType

`func (o *CreateMessageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMessageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMessageRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateMessageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoDetect

`func (o *CreateMessageRequest) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *CreateMessageRequest) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *CreateMessageRequest) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *CreateMessageRequest) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


