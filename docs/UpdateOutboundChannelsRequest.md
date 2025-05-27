# UpdateOutboundChannelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **int32** | The new number of concurrent channels for the account | 

## Methods

### NewUpdateOutboundChannelsRequest

`func NewUpdateOutboundChannelsRequest(channels int32, ) *UpdateOutboundChannelsRequest`

NewUpdateOutboundChannelsRequest instantiates a new UpdateOutboundChannelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOutboundChannelsRequestWithDefaults

`func NewUpdateOutboundChannelsRequestWithDefaults() *UpdateOutboundChannelsRequest`

NewUpdateOutboundChannelsRequestWithDefaults instantiates a new UpdateOutboundChannelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *UpdateOutboundChannelsRequest) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *UpdateOutboundChannelsRequest) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *UpdateOutboundChannelsRequest) SetChannels(v int32)`

SetChannels sets Channels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


