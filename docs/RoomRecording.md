# RoomRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the room recording. | [optional] 
**RoomId** | Pointer to **string** | Identify the room associated with the room recording. | [optional] 
**SessionId** | Pointer to **string** | Identify the room session associated with the room recording. | [optional] 
**ParticipantId** | Pointer to **string** | Identify the room participant associated with the room recording. | [optional] 
**Status** | Pointer to **string** | Shows the room recording status. | [optional] 
**Type** | Pointer to **string** | Shows the room recording type. | [optional] 
**SizeMb** | Pointer to **float32** | Shows the room recording size in MB. | [optional] 
**DownloadUrl** | Pointer to **string** | Url to download the recording. | [optional] 
**Codec** | Pointer to **string** | Shows the codec used for the room recording. | [optional] 
**DurationSecs** | Pointer to **int32** | Shows the room recording duration in seconds. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 timestamp when the room recording was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when the room recording was updated. | [optional] 
**EndedAt** | Pointer to **string** | ISO 8601 timestamp when the room recording has ended. | [optional] 
**StartedAt** | Pointer to **string** | ISO 8601 timestamp when the room recording has stated. | [optional] 
**CompletedAt** | Pointer to **string** | ISO 8601 timestamp when the room recording has completed. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoomRecording

`func NewRoomRecording() *RoomRecording`

NewRoomRecording instantiates a new RoomRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomRecordingWithDefaults

`func NewRoomRecordingWithDefaults() *RoomRecording`

NewRoomRecordingWithDefaults instantiates a new RoomRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoomRecording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomRecording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomRecording) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoomRecording) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoomId

`func (o *RoomRecording) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *RoomRecording) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *RoomRecording) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *RoomRecording) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetSessionId

`func (o *RoomRecording) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *RoomRecording) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *RoomRecording) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *RoomRecording) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetParticipantId

`func (o *RoomRecording) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *RoomRecording) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *RoomRecording) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *RoomRecording) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetStatus

`func (o *RoomRecording) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoomRecording) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoomRecording) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoomRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *RoomRecording) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoomRecording) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoomRecording) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoomRecording) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSizeMb

`func (o *RoomRecording) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *RoomRecording) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *RoomRecording) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.

### HasSizeMb

`func (o *RoomRecording) HasSizeMb() bool`

HasSizeMb returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *RoomRecording) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *RoomRecording) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *RoomRecording) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *RoomRecording) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetCodec

`func (o *RoomRecording) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *RoomRecording) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *RoomRecording) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *RoomRecording) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetDurationSecs

`func (o *RoomRecording) GetDurationSecs() int32`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *RoomRecording) GetDurationSecsOk() (*int32, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *RoomRecording) SetDurationSecs(v int32)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *RoomRecording) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoomRecording) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoomRecording) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoomRecording) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoomRecording) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoomRecording) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoomRecording) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoomRecording) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoomRecording) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *RoomRecording) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RoomRecording) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RoomRecording) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RoomRecording) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RoomRecording) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RoomRecording) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RoomRecording) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RoomRecording) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *RoomRecording) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *RoomRecording) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *RoomRecording) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *RoomRecording) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetRecordType

`func (o *RoomRecording) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RoomRecording) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RoomRecording) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RoomRecording) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


