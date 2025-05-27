# ListMigrationSources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MigrationSourceParams**](MigrationSourceParams.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMetaSimple**](PaginationMetaSimple.md) |  | [optional] 

## Methods

### NewListMigrationSources200Response

`func NewListMigrationSources200Response() *ListMigrationSources200Response`

NewListMigrationSources200Response instantiates a new ListMigrationSources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMigrationSources200ResponseWithDefaults

`func NewListMigrationSources200ResponseWithDefaults() *ListMigrationSources200Response`

NewListMigrationSources200ResponseWithDefaults instantiates a new ListMigrationSources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMigrationSources200Response) GetData() []MigrationSourceParams`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMigrationSources200Response) GetDataOk() (*[]MigrationSourceParams, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMigrationSources200Response) SetData(v []MigrationSourceParams)`

SetData sets Data field to given value.

### HasData

`func (o *ListMigrationSources200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListMigrationSources200Response) GetMeta() PaginationMetaSimple`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMigrationSources200Response) GetMetaOk() (*PaginationMetaSimple, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMigrationSources200Response) SetMeta(v PaginationMetaSimple)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMigrationSources200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


