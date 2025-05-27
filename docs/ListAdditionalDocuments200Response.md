# ListAdditionalDocuments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortingAdditionalDocument**](PortingAdditionalDocument.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListAdditionalDocuments200Response

`func NewListAdditionalDocuments200Response() *ListAdditionalDocuments200Response`

NewListAdditionalDocuments200Response instantiates a new ListAdditionalDocuments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAdditionalDocuments200ResponseWithDefaults

`func NewListAdditionalDocuments200ResponseWithDefaults() *ListAdditionalDocuments200Response`

NewListAdditionalDocuments200ResponseWithDefaults instantiates a new ListAdditionalDocuments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAdditionalDocuments200Response) GetData() []PortingAdditionalDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAdditionalDocuments200Response) GetDataOk() (*[]PortingAdditionalDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAdditionalDocuments200Response) SetData(v []PortingAdditionalDocument)`

SetData sets Data field to given value.

### HasData

`func (o *ListAdditionalDocuments200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListAdditionalDocuments200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAdditionalDocuments200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAdditionalDocuments200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAdditionalDocuments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


