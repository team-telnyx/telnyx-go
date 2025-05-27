# PatchRoomRequest

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

### NewPatchRoomRequest

`func NewPatchRoomRequest() *PatchRoomRequest`

NewPatchRoomRequest instantiates a new PatchRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRoomRequestWithDefaults

`func NewPatchRoomRequestWithDefaults() *PatchRoomRequest`

NewPatchRoomRequestWithDefaults instantiates a new PatchRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueName

`func (o *PatchRoomRequest) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *PatchRoomRequest) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *PatchRoomRequest) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *PatchRoomRequest) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetMaxParticipants

`func (o *PatchRoomRequest) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *PatchRoomRequest) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *PatchRoomRequest) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *PatchRoomRequest) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### GetEnableRecording

`func (o *PatchRoomRequest) GetEnableRecording() bool`

GetEnableRecording returns the EnableRecording field if non-nil, zero value otherwise.

### GetEnableRecordingOk

`func (o *PatchRoomRequest) GetEnableRecordingOk() (*bool, bool)`

GetEnableRecordingOk returns a tuple with the EnableRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecording

`func (o *PatchRoomRequest) SetEnableRecording(v bool)`

SetEnableRecording sets EnableRecording field to given value.

### HasEnableRecording

`func (o *PatchRoomRequest) HasEnableRecording() bool`

HasEnableRecording returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *PatchRoomRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *PatchRoomRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *PatchRoomRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *PatchRoomRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *PatchRoomRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *PatchRoomRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *PatchRoomRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *PatchRoomRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *PatchRoomRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *PatchRoomRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *PatchRoomRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *PatchRoomRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *PatchRoomRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *PatchRoomRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *PatchRoomRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *PatchRoomRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


