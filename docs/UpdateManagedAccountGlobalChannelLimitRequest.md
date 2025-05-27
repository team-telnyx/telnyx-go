# UpdateManagedAccountGlobalChannelLimitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | Integer value that indicates the number of allocatable global outbound channels that should be allocated to the managed account. Must be 0 or more. If the value is 0 then the account will have no usable channels and will not be able to perform outbound calling. | [optional] 

## Methods

### NewUpdateManagedAccountGlobalChannelLimitRequest

`func NewUpdateManagedAccountGlobalChannelLimitRequest() *UpdateManagedAccountGlobalChannelLimitRequest`

NewUpdateManagedAccountGlobalChannelLimitRequest instantiates a new UpdateManagedAccountGlobalChannelLimitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateManagedAccountGlobalChannelLimitRequestWithDefaults

`func NewUpdateManagedAccountGlobalChannelLimitRequestWithDefaults() *UpdateManagedAccountGlobalChannelLimitRequest`

NewUpdateManagedAccountGlobalChannelLimitRequestWithDefaults instantiates a new UpdateManagedAccountGlobalChannelLimitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *UpdateManagedAccountGlobalChannelLimitRequest) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *UpdateManagedAccountGlobalChannelLimitRequest) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *UpdateManagedAccountGlobalChannelLimitRequest) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *UpdateManagedAccountGlobalChannelLimitRequest) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


