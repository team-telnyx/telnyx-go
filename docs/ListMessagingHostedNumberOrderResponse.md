# ListMessagingHostedNumberOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MessagingHostedNumberOrder**](MessagingHostedNumberOrder.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListMessagingHostedNumberOrderResponse

`func NewListMessagingHostedNumberOrderResponse() *ListMessagingHostedNumberOrderResponse`

NewListMessagingHostedNumberOrderResponse instantiates a new ListMessagingHostedNumberOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessagingHostedNumberOrderResponseWithDefaults

`func NewListMessagingHostedNumberOrderResponseWithDefaults() *ListMessagingHostedNumberOrderResponse`

NewListMessagingHostedNumberOrderResponseWithDefaults instantiates a new ListMessagingHostedNumberOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMessagingHostedNumberOrderResponse) GetData() []MessagingHostedNumberOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMessagingHostedNumberOrderResponse) GetDataOk() (*[]MessagingHostedNumberOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMessagingHostedNumberOrderResponse) SetData(v []MessagingHostedNumberOrder)`

SetData sets Data field to given value.

### HasData

`func (o *ListMessagingHostedNumberOrderResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListMessagingHostedNumberOrderResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMessagingHostedNumberOrderResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMessagingHostedNumberOrderResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMessagingHostedNumberOrderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


