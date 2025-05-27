# RoomSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the room session. | [optional] 
**RoomId** | Pointer to **string** | Identify the room hosting that room session. | [optional] 
**Active** | Pointer to **bool** | Shows if the room session is active or not. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 timestamp when the room session was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when the room session was updated. | [optional] 
**EndedAt** | Pointer to **string** | ISO 8601 timestamp when the room session has ended. | [optional] 
**Participants** | Pointer to [**[]RoomParticipant**](RoomParticipant.md) |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoomSession

`func NewRoomSession() *RoomSession`

NewRoomSession instantiates a new RoomSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomSessionWithDefaults

`func NewRoomSessionWithDefaults() *RoomSession`

NewRoomSessionWithDefaults instantiates a new RoomSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoomSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoomSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoomId

`func (o *RoomSession) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *RoomSession) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *RoomSession) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *RoomSession) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetActive

`func (o *RoomSession) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RoomSession) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RoomSession) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RoomSession) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoomSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoomSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoomSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoomSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoomSession) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoomSession) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoomSession) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoomSession) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *RoomSession) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RoomSession) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RoomSession) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RoomSession) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetParticipants

`func (o *RoomSession) GetParticipants() []RoomParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *RoomSession) GetParticipantsOk() (*[]RoomParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *RoomSession) SetParticipants(v []RoomParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *RoomSession) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetRecordType

`func (o *RoomSession) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RoomSession) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RoomSession) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RoomSession) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


