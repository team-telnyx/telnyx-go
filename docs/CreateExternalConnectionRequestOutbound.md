# CreateExternalConnectionRequestOutbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | When set, this will limit the number of concurrent outbound calls to phone numbers associated with this connection. | [optional] 
**OutboundVoiceProfileId** | Pointer to **string** | Identifies the associated outbound voice profile. | [optional] 

## Methods

### NewCreateExternalConnectionRequestOutbound

`func NewCreateExternalConnectionRequestOutbound() *CreateExternalConnectionRequestOutbound`

NewCreateExternalConnectionRequestOutbound instantiates a new CreateExternalConnectionRequestOutbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalConnectionRequestOutboundWithDefaults

`func NewCreateExternalConnectionRequestOutboundWithDefaults() *CreateExternalConnectionRequestOutbound`

NewCreateExternalConnectionRequestOutboundWithDefaults instantiates a new CreateExternalConnectionRequestOutbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *CreateExternalConnectionRequestOutbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CreateExternalConnectionRequestOutbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CreateExternalConnectionRequestOutbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CreateExternalConnectionRequestOutbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetOutboundVoiceProfileId

`func (o *CreateExternalConnectionRequestOutbound) GetOutboundVoiceProfileId() string`

GetOutboundVoiceProfileId returns the OutboundVoiceProfileId field if non-nil, zero value otherwise.

### GetOutboundVoiceProfileIdOk

`func (o *CreateExternalConnectionRequestOutbound) GetOutboundVoiceProfileIdOk() (*string, bool)`

GetOutboundVoiceProfileIdOk returns a tuple with the OutboundVoiceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundVoiceProfileId

`func (o *CreateExternalConnectionRequestOutbound) SetOutboundVoiceProfileId(v string)`

SetOutboundVoiceProfileId sets OutboundVoiceProfileId field to given value.

### HasOutboundVoiceProfileId

`func (o *CreateExternalConnectionRequestOutbound) HasOutboundVoiceProfileId() bool`

HasOutboundVoiceProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


