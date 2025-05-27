# CreateInventoryCoverage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]InventoryCoverage**](InventoryCoverage.md) |  | [optional] 
**Meta** | Pointer to [**InventoryCoverageMetadata**](InventoryCoverageMetadata.md) |  | [optional] 

## Methods

### NewCreateInventoryCoverage200Response

`func NewCreateInventoryCoverage200Response() *CreateInventoryCoverage200Response`

NewCreateInventoryCoverage200Response instantiates a new CreateInventoryCoverage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryCoverage200ResponseWithDefaults

`func NewCreateInventoryCoverage200ResponseWithDefaults() *CreateInventoryCoverage200Response`

NewCreateInventoryCoverage200ResponseWithDefaults instantiates a new CreateInventoryCoverage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateInventoryCoverage200Response) GetData() []InventoryCoverage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateInventoryCoverage200Response) GetDataOk() (*[]InventoryCoverage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateInventoryCoverage200Response) SetData(v []InventoryCoverage)`

SetData sets Data field to given value.

### HasData

`func (o *CreateInventoryCoverage200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *CreateInventoryCoverage200Response) GetMeta() InventoryCoverageMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateInventoryCoverage200Response) GetMetaOk() (*InventoryCoverageMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateInventoryCoverage200Response) SetMeta(v InventoryCoverageMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CreateInventoryCoverage200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


