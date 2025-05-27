# FindNotificationsEventsConditions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]NotificationEventCondition**](NotificationEventCondition.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewFindNotificationsEventsConditions200Response

`func NewFindNotificationsEventsConditions200Response() *FindNotificationsEventsConditions200Response`

NewFindNotificationsEventsConditions200Response instantiates a new FindNotificationsEventsConditions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindNotificationsEventsConditions200ResponseWithDefaults

`func NewFindNotificationsEventsConditions200ResponseWithDefaults() *FindNotificationsEventsConditions200Response`

NewFindNotificationsEventsConditions200ResponseWithDefaults instantiates a new FindNotificationsEventsConditions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindNotificationsEventsConditions200Response) GetData() []NotificationEventCondition`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindNotificationsEventsConditions200Response) GetDataOk() (*[]NotificationEventCondition, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindNotificationsEventsConditions200Response) SetData(v []NotificationEventCondition)`

SetData sets Data field to given value.

### HasData

`func (o *FindNotificationsEventsConditions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *FindNotificationsEventsConditions200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindNotificationsEventsConditions200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindNotificationsEventsConditions200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindNotificationsEventsConditions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


