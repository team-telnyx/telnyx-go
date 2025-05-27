# SingleManagedAccountGlobalOutboundChannels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | Integer value that indicates the number of allocatable global outbound channels that are allocated to the managed account. If the value is 0 then the account will have no usable channels and will not be able to perform outbound calling. | [optional] 
**Email** | Pointer to **string** | The email of the managed account. | [optional] 
**Id** | Pointer to **string** | The user ID of the managed account. | [optional] 
**ManagerAccountId** | Pointer to **string** | The user ID of the manager of the account. | [optional] 
**RecordType** | Pointer to **string** | The name of the type of data in the response. | [optional] 

## Methods

### NewSingleManagedAccountGlobalOutboundChannels

`func NewSingleManagedAccountGlobalOutboundChannels() *SingleManagedAccountGlobalOutboundChannels`

NewSingleManagedAccountGlobalOutboundChannels instantiates a new SingleManagedAccountGlobalOutboundChannels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleManagedAccountGlobalOutboundChannelsWithDefaults

`func NewSingleManagedAccountGlobalOutboundChannelsWithDefaults() *SingleManagedAccountGlobalOutboundChannels`

NewSingleManagedAccountGlobalOutboundChannelsWithDefaults instantiates a new SingleManagedAccountGlobalOutboundChannels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *SingleManagedAccountGlobalOutboundChannels) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *SingleManagedAccountGlobalOutboundChannels) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *SingleManagedAccountGlobalOutboundChannels) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *SingleManagedAccountGlobalOutboundChannels) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetEmail

`func (o *SingleManagedAccountGlobalOutboundChannels) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SingleManagedAccountGlobalOutboundChannels) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SingleManagedAccountGlobalOutboundChannels) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SingleManagedAccountGlobalOutboundChannels) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *SingleManagedAccountGlobalOutboundChannels) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SingleManagedAccountGlobalOutboundChannels) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SingleManagedAccountGlobalOutboundChannels) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SingleManagedAccountGlobalOutboundChannels) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagerAccountId

`func (o *SingleManagedAccountGlobalOutboundChannels) GetManagerAccountId() string`

GetManagerAccountId returns the ManagerAccountId field if non-nil, zero value otherwise.

### GetManagerAccountIdOk

`func (o *SingleManagedAccountGlobalOutboundChannels) GetManagerAccountIdOk() (*string, bool)`

GetManagerAccountIdOk returns a tuple with the ManagerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAccountId

`func (o *SingleManagedAccountGlobalOutboundChannels) SetManagerAccountId(v string)`

SetManagerAccountId sets ManagerAccountId field to given value.

### HasManagerAccountId

`func (o *SingleManagedAccountGlobalOutboundChannels) HasManagerAccountId() bool`

HasManagerAccountId returns a boolean if a field has been set.

### GetRecordType

`func (o *SingleManagedAccountGlobalOutboundChannels) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SingleManagedAccountGlobalOutboundChannels) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SingleManagedAccountGlobalOutboundChannels) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SingleManagedAccountGlobalOutboundChannels) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


