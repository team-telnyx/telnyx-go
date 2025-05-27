# GetPortingOrder200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PortingOrder**](PortingOrder.md) |  | [optional] 
**Meta** | Pointer to [**GetPortingOrder200ResponseMeta**](GetPortingOrder200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetPortingOrder200Response

`func NewGetPortingOrder200Response() *GetPortingOrder200Response`

NewGetPortingOrder200Response instantiates a new GetPortingOrder200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortingOrder200ResponseWithDefaults

`func NewGetPortingOrder200ResponseWithDefaults() *GetPortingOrder200Response`

NewGetPortingOrder200ResponseWithDefaults instantiates a new GetPortingOrder200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPortingOrder200Response) GetData() PortingOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPortingOrder200Response) GetDataOk() (*PortingOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPortingOrder200Response) SetData(v PortingOrder)`

SetData sets Data field to given value.

### HasData

`func (o *GetPortingOrder200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetPortingOrder200Response) GetMeta() GetPortingOrder200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetPortingOrder200Response) GetMetaOk() (*GetPortingOrder200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetPortingOrder200Response) SetMeta(v GetPortingOrder200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetPortingOrder200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


