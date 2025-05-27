# RoomParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the room participant. | [optional] 
**SessionId** | Pointer to **string** | Identify the room session that participant is part of. | [optional] 
**Context** | Pointer to **string** | Context provided to the given participant through the client SDK | [optional] 
**JoinedAt** | Pointer to **string** | ISO 8601 timestamp when the participant joined the session. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when the participant was updated. | [optional] 
**LeftAt** | Pointer to **string** | ISO 8601 timestamp when the participant left the session. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoomParticipant

`func NewRoomParticipant() *RoomParticipant`

NewRoomParticipant instantiates a new RoomParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomParticipantWithDefaults

`func NewRoomParticipantWithDefaults() *RoomParticipant`

NewRoomParticipantWithDefaults instantiates a new RoomParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoomParticipant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomParticipant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomParticipant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoomParticipant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSessionId

`func (o *RoomParticipant) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *RoomParticipant) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *RoomParticipant) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *RoomParticipant) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetContext

`func (o *RoomParticipant) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RoomParticipant) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RoomParticipant) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *RoomParticipant) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetJoinedAt

`func (o *RoomParticipant) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *RoomParticipant) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *RoomParticipant) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *RoomParticipant) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoomParticipant) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoomParticipant) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoomParticipant) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoomParticipant) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLeftAt

`func (o *RoomParticipant) GetLeftAt() string`

GetLeftAt returns the LeftAt field if non-nil, zero value otherwise.

### GetLeftAtOk

`func (o *RoomParticipant) GetLeftAtOk() (*string, bool)`

GetLeftAtOk returns a tuple with the LeftAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAt

`func (o *RoomParticipant) SetLeftAt(v string)`

SetLeftAt sets LeftAt field to given value.

### HasLeftAt

`func (o *RoomParticipant) HasLeftAt() bool`

HasLeftAt returns a boolean if a field has been set.

### GetRecordType

`func (o *RoomParticipant) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RoomParticipant) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RoomParticipant) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RoomParticipant) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


