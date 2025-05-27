# ListNotificationChannels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]NotificationChannel**](NotificationChannel.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListNotificationChannels200Response

`func NewListNotificationChannels200Response() *ListNotificationChannels200Response`

NewListNotificationChannels200Response instantiates a new ListNotificationChannels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationChannels200ResponseWithDefaults

`func NewListNotificationChannels200ResponseWithDefaults() *ListNotificationChannels200Response`

NewListNotificationChannels200ResponseWithDefaults instantiates a new ListNotificationChannels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListNotificationChannels200Response) GetData() []NotificationChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNotificationChannels200Response) GetDataOk() (*[]NotificationChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNotificationChannels200Response) SetData(v []NotificationChannel)`

SetData sets Data field to given value.

### HasData

`func (o *ListNotificationChannels200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListNotificationChannels200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNotificationChannels200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNotificationChannels200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNotificationChannels200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


