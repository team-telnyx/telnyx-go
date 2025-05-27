# Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the room. | [optional] 
**MaxParticipants** | Pointer to **int32** | Maximum participants allowed in the room. | [optional] 
**UniqueName** | Pointer to **string** | The unique (within the Telnyx account scope) name of the room. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 timestamp when the room was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when the room was updated. | [optional] 
**ActiveSessionId** | Pointer to **string** | The identifier of the active room session if any. | [optional] 
**Sessions** | Pointer to [**[]RoomSession**](RoomSession.md) |  | [optional] 
**EnableRecording** | Pointer to **bool** | Enable or disable recording for that room. | [optional] [default to false]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this room will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this room will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoom

`func NewRoom() *Room`

NewRoom instantiates a new Room object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomWithDefaults

`func NewRoomWithDefaults() *Room`

NewRoomWithDefaults instantiates a new Room object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Room) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Room) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Room) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Room) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxParticipants

`func (o *Room) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *Room) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *Room) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *Room) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### GetUniqueName

`func (o *Room) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *Room) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *Room) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *Room) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Room) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Room) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Room) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Room) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Room) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Room) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Room) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Room) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetActiveSessionId

`func (o *Room) GetActiveSessionId() string`

GetActiveSessionId returns the ActiveSessionId field if non-nil, zero value otherwise.

### GetActiveSessionIdOk

`func (o *Room) GetActiveSessionIdOk() (*string, bool)`

GetActiveSessionIdOk returns a tuple with the ActiveSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessionId

`func (o *Room) SetActiveSessionId(v string)`

SetActiveSessionId sets ActiveSessionId field to given value.

### HasActiveSessionId

`func (o *Room) HasActiveSessionId() bool`

HasActiveSessionId returns a boolean if a field has been set.

### GetSessions

`func (o *Room) GetSessions() []RoomSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *Room) GetSessionsOk() (*[]RoomSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *Room) SetSessions(v []RoomSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *Room) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetEnableRecording

`func (o *Room) GetEnableRecording() bool`

GetEnableRecording returns the EnableRecording field if non-nil, zero value otherwise.

### GetEnableRecordingOk

`func (o *Room) GetEnableRecordingOk() (*bool, bool)`

GetEnableRecordingOk returns a tuple with the EnableRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecording

`func (o *Room) SetEnableRecording(v bool)`

SetEnableRecording sets EnableRecording field to given value.

### HasEnableRecording

`func (o *Room) HasEnableRecording() bool`

HasEnableRecording returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *Room) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *Room) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *Room) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *Room) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *Room) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *Room) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *Room) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *Room) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *Room) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *Room) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *Room) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *Room) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *Room) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *Room) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *Room) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *Room) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetRecordType

`func (o *Room) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Room) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Room) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Room) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


