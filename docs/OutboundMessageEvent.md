# OutboundMessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**OutboundMessage**](OutboundMessage.md) |  | [optional] 
**Meta** | Pointer to [**OutboundMessageEventMeta**](OutboundMessageEventMeta.md) |  | [optional] 

## Methods

### NewOutboundMessageEvent

`func NewOutboundMessageEvent() *OutboundMessageEvent`

NewOutboundMessageEvent instantiates a new OutboundMessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessageEventWithDefaults

`func NewOutboundMessageEventWithDefaults() *OutboundMessageEvent`

NewOutboundMessageEventWithDefaults instantiates a new OutboundMessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OutboundMessageEvent) GetData() OutboundMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundMessageEvent) GetDataOk() (*OutboundMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundMessageEvent) SetData(v OutboundMessage)`

SetData sets Data field to given value.

### HasData

`func (o *OutboundMessageEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *OutboundMessageEvent) GetMeta() OutboundMessageEventMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OutboundMessageEvent) GetMetaOk() (*OutboundMessageEventMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OutboundMessageEvent) SetMeta(v OutboundMessageEventMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OutboundMessageEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


