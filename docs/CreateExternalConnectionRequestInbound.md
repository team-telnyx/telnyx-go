# CreateExternalConnectionRequestInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection. | [optional] 

## Methods

### NewCreateExternalConnectionRequestInbound

`func NewCreateExternalConnectionRequestInbound() *CreateExternalConnectionRequestInbound`

NewCreateExternalConnectionRequestInbound instantiates a new CreateExternalConnectionRequestInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalConnectionRequestInboundWithDefaults

`func NewCreateExternalConnectionRequestInboundWithDefaults() *CreateExternalConnectionRequestInbound`

NewCreateExternalConnectionRequestInboundWithDefaults instantiates a new CreateExternalConnectionRequestInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *CreateExternalConnectionRequestInbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CreateExternalConnectionRequestInbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CreateExternalConnectionRequestInbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CreateExternalConnectionRequestInbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


