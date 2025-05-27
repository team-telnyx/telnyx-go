# ListPortoutRequest200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortoutDetails**](PortoutDetails.md) |  | [optional] 
**Meta** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewListPortoutRequest200Response

`func NewListPortoutRequest200Response() *ListPortoutRequest200Response`

NewListPortoutRequest200Response instantiates a new ListPortoutRequest200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPortoutRequest200ResponseWithDefaults

`func NewListPortoutRequest200ResponseWithDefaults() *ListPortoutRequest200Response`

NewListPortoutRequest200ResponseWithDefaults instantiates a new ListPortoutRequest200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPortoutRequest200Response) GetData() []PortoutDetails`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPortoutRequest200Response) GetDataOk() (*[]PortoutDetails, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPortoutRequest200Response) SetData(v []PortoutDetails)`

SetData sets Data field to given value.

### HasData

`func (o *ListPortoutRequest200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPortoutRequest200Response) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPortoutRequest200Response) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPortoutRequest200Response) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPortoutRequest200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


