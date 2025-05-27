# CreateLongCodeMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | Phone number, in +E.164 format, used to send the message. | 
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

### NewCreateLongCodeMessageRequest

`func NewCreateLongCodeMessageRequest(from string, to string, ) *CreateLongCodeMessageRequest`

NewCreateLongCodeMessageRequest instantiates a new CreateLongCodeMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLongCodeMessageRequestWithDefaults

`func NewCreateLongCodeMessageRequestWithDefaults() *CreateLongCodeMessageRequest`

NewCreateLongCodeMessageRequestWithDefaults instantiates a new CreateLongCodeMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *CreateLongCodeMessageRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateLongCodeMessageRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateLongCodeMessageRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *CreateLongCodeMessageRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateLongCodeMessageRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateLongCodeMessageRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetText

`func (o *CreateLongCodeMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateLongCodeMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateLongCodeMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateLongCodeMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSubject

`func (o *CreateLongCodeMessageRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateLongCodeMessageRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateLongCodeMessageRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateLongCodeMessageRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMediaUrls

`func (o *CreateLongCodeMessageRequest) GetMediaUrls() []string`

GetMediaUrls returns the MediaUrls field if non-nil, zero value otherwise.

### GetMediaUrlsOk

`func (o *CreateLongCodeMessageRequest) GetMediaUrlsOk() (*[]string, bool)`

GetMediaUrlsOk returns a tuple with the MediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrls

`func (o *CreateLongCodeMessageRequest) SetMediaUrls(v []string)`

SetMediaUrls sets MediaUrls field to given value.

### HasMediaUrls

`func (o *CreateLongCodeMessageRequest) HasMediaUrls() bool`

HasMediaUrls returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateLongCodeMessageRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateLongCodeMessageRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateLongCodeMessageRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateLongCodeMessageRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookFailoverUrl

`func (o *CreateLongCodeMessageRequest) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *CreateLongCodeMessageRequest) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *CreateLongCodeMessageRequest) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *CreateLongCodeMessageRequest) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### GetUseProfileWebhooks

`func (o *CreateLongCodeMessageRequest) GetUseProfileWebhooks() bool`

GetUseProfileWebhooks returns the UseProfileWebhooks field if non-nil, zero value otherwise.

### GetUseProfileWebhooksOk

`func (o *CreateLongCodeMessageRequest) GetUseProfileWebhooksOk() (*bool, bool)`

GetUseProfileWebhooksOk returns a tuple with the UseProfileWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProfileWebhooks

`func (o *CreateLongCodeMessageRequest) SetUseProfileWebhooks(v bool)`

SetUseProfileWebhooks sets UseProfileWebhooks field to given value.

### HasUseProfileWebhooks

`func (o *CreateLongCodeMessageRequest) HasUseProfileWebhooks() bool`

HasUseProfileWebhooks returns a boolean if a field has been set.

### GetType

`func (o *CreateLongCodeMessageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLongCodeMessageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLongCodeMessageRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateLongCodeMessageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAutoDetect

`func (o *CreateLongCodeMessageRequest) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *CreateLongCodeMessageRequest) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *CreateLongCodeMessageRequest) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *CreateLongCodeMessageRequest) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


