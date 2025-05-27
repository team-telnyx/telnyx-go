# FindNotificationsProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]NotificationProfile**](NotificationProfile.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewFindNotificationsProfiles200Response

`func NewFindNotificationsProfiles200Response() *FindNotificationsProfiles200Response`

NewFindNotificationsProfiles200Response instantiates a new FindNotificationsProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindNotificationsProfiles200ResponseWithDefaults

`func NewFindNotificationsProfiles200ResponseWithDefaults() *FindNotificationsProfiles200Response`

NewFindNotificationsProfiles200ResponseWithDefaults instantiates a new FindNotificationsProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindNotificationsProfiles200Response) GetData() []NotificationProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindNotificationsProfiles200Response) GetDataOk() (*[]NotificationProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindNotificationsProfiles200Response) SetData(v []NotificationProfile)`

SetData sets Data field to given value.

### HasData

`func (o *FindNotificationsProfiles200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *FindNotificationsProfiles200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindNotificationsProfiles200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindNotificationsProfiles200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindNotificationsProfiles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


