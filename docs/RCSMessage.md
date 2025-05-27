# RCSMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | RCS Agent ID | 
**To** | **string** | Phone number in +E.164 format | 
**MessagingProfileId** | **string** | A valid messaging profile ID | 
**Type** | Pointer to **string** | Message type - must be set to \&quot;RCS\&quot; | [optional] 
**WebhookUrl** | Pointer to **string** | The URL where webhooks related to this message will be sent. | [optional] 
**AgentMessage** | [**RCSAgentMessage**](RCSAgentMessage.md) |  | 
**SmsFallback** | Pointer to [**SMSFallback**](SMSFallback.md) |  | [optional] 

## Methods

### NewRCSMessage

`func NewRCSMessage(agentId string, to string, messagingProfileId string, agentMessage RCSAgentMessage, ) *RCSMessage`

NewRCSMessage instantiates a new RCSMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSMessageWithDefaults

`func NewRCSMessageWithDefaults() *RCSMessage`

NewRCSMessageWithDefaults instantiates a new RCSMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *RCSMessage) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *RCSMessage) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *RCSMessage) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.


### GetTo

`func (o *RCSMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RCSMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RCSMessage) SetTo(v string)`

SetTo sets To field to given value.


### GetMessagingProfileId

`func (o *RCSMessage) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *RCSMessage) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *RCSMessage) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.


### GetType

`func (o *RCSMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RCSMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RCSMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RCSMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *RCSMessage) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *RCSMessage) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *RCSMessage) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *RCSMessage) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetAgentMessage

`func (o *RCSMessage) GetAgentMessage() RCSAgentMessage`

GetAgentMessage returns the AgentMessage field if non-nil, zero value otherwise.

### GetAgentMessageOk

`func (o *RCSMessage) GetAgentMessageOk() (*RCSAgentMessage, bool)`

GetAgentMessageOk returns a tuple with the AgentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMessage

`func (o *RCSMessage) SetAgentMessage(v RCSAgentMessage)`

SetAgentMessage sets AgentMessage field to given value.


### GetSmsFallback

`func (o *RCSMessage) GetSmsFallback() SMSFallback`

GetSmsFallback returns the SmsFallback field if non-nil, zero value otherwise.

### GetSmsFallbackOk

`func (o *RCSMessage) GetSmsFallbackOk() (*SMSFallback, bool)`

GetSmsFallbackOk returns a tuple with the SmsFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallback

`func (o *RCSMessage) SetSmsFallback(v SMSFallback)`

SetSmsFallback sets SmsFallback field to given value.

### HasSmsFallback

`func (o *RCSMessage) HasSmsFallback() bool`

HasSmsFallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


