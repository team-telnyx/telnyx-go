# GetBucketUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BucketUsage**](BucketUsage.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMetaSimple**](PaginationMetaSimple.md) |  | [optional] 

## Methods

### NewGetBucketUsage200Response

`func NewGetBucketUsage200Response() *GetBucketUsage200Response`

NewGetBucketUsage200Response instantiates a new GetBucketUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBucketUsage200ResponseWithDefaults

`func NewGetBucketUsage200ResponseWithDefaults() *GetBucketUsage200Response`

NewGetBucketUsage200ResponseWithDefaults instantiates a new GetBucketUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBucketUsage200Response) GetData() []BucketUsage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBucketUsage200Response) GetDataOk() (*[]BucketUsage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBucketUsage200Response) SetData(v []BucketUsage)`

SetData sets Data field to given value.

### HasData

`func (o *GetBucketUsage200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetBucketUsage200Response) GetMeta() PaginationMetaSimple`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBucketUsage200Response) GetMetaOk() (*PaginationMetaSimple, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBucketUsage200Response) SetMeta(v PaginationMetaSimple)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetBucketUsage200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


