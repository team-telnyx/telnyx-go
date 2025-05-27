# ListRooms200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Room**](Room.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListRooms200Response

`func NewListRooms200Response() *ListRooms200Response`

NewListRooms200Response instantiates a new ListRooms200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRooms200ResponseWithDefaults

`func NewListRooms200ResponseWithDefaults() *ListRooms200Response`

NewListRooms200ResponseWithDefaults instantiates a new ListRooms200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRooms200Response) GetData() []Room`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRooms200Response) GetDataOk() (*[]Room, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRooms200Response) SetData(v []Room)`

SetData sets Data field to given value.

### HasData

`func (o *ListRooms200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListRooms200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRooms200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRooms200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRooms200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


