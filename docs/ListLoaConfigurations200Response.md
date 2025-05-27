# ListLoaConfigurations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortingLOAConfiguration**](PortingLOAConfiguration.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListLoaConfigurations200Response

`func NewListLoaConfigurations200Response() *ListLoaConfigurations200Response`

NewListLoaConfigurations200Response instantiates a new ListLoaConfigurations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoaConfigurations200ResponseWithDefaults

`func NewListLoaConfigurations200ResponseWithDefaults() *ListLoaConfigurations200Response`

NewListLoaConfigurations200ResponseWithDefaults instantiates a new ListLoaConfigurations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListLoaConfigurations200Response) GetData() []PortingLOAConfiguration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLoaConfigurations200Response) GetDataOk() (*[]PortingLOAConfiguration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLoaConfigurations200Response) SetData(v []PortingLOAConfiguration)`

SetData sets Data field to given value.

### HasData

`func (o *ListLoaConfigurations200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListLoaConfigurations200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLoaConfigurations200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLoaConfigurations200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLoaConfigurations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


