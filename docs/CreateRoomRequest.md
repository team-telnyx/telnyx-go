# CreateRoomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueName** | Pointer to **string** | The unique (within the Telnyx account scope) name of the room. | [optional] 
**MaxParticipants** | Pointer to **int32** | The maximum amount of participants allowed in a room. If new participants try to join after that limit is reached, their request will be rejected. | [optional] [default to 10]
**EnableRecording** | Pointer to **bool** | Enable or disable recording for that room. | [optional] [default to false]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this room will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this room will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 

## Methods

### NewCreateRoomRequest

`func NewCreateRoomRequest() *CreateRoomRequest`

NewCreateRoomRequest instantiates a new CreateRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoomRequestWithDefaults

`func NewCreateRoomRequestWithDefaults() *CreateRoomRequest`

NewCreateRoomRequestWithDefaults instantiates a new CreateRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueName

`func (o *CreateRoomRequest) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *CreateRoomRequest) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *CreateRoomRequest) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *CreateRoomRequest) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetMaxParticipants

`func (o *CreateRoomRequest) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *CreateRoomRequest) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *CreateRoomRequest) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *CreateRoomRequest) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### GetEnableRecording

`func (o *CreateRoomRequest) GetEnableRecording() bool`

GetEnableRecording returns the EnableRecording field if non-nil, zero value otherwise.

### GetEnableRecordingOk

`func (o *CreateRoomRequest) GetEnableRecordingOk() (*bool, bool)`

GetEnableRecordingOk returns a tuple with the EnableRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecording

`func (o *CreateRoomRequest) SetEnableRecording(v bool)`

SetEnableRecording sets EnableRecording field to given value.

### HasEnableRecording

`func (o *CreateRoomRequest) HasEnableRecording() bool`

HasEnableRecording returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *CreateRoomRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CreateRoomRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CreateRoomRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CreateRoomRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CreateRoomRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CreateRoomRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CreateRoomRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CreateRoomRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CreateRoomRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CreateRoomRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *CreateRoomRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CreateRoomRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CreateRoomRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CreateRoomRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CreateRoomRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CreateRoomRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


