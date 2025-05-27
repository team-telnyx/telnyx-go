# ListRoomParticipants200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RoomParticipant**](RoomParticipant.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListRoomParticipants200Response

`func NewListRoomParticipants200Response() *ListRoomParticipants200Response`

NewListRoomParticipants200Response instantiates a new ListRoomParticipants200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomParticipants200ResponseWithDefaults

`func NewListRoomParticipants200ResponseWithDefaults() *ListRoomParticipants200Response`

NewListRoomParticipants200ResponseWithDefaults instantiates a new ListRoomParticipants200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRoomParticipants200Response) GetData() []RoomParticipant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRoomParticipants200Response) GetDataOk() (*[]RoomParticipant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRoomParticipants200Response) SetData(v []RoomParticipant)`

SetData sets Data field to given value.

### HasData

`func (o *ListRoomParticipants200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListRoomParticipants200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomParticipants200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomParticipants200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomParticipants200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


