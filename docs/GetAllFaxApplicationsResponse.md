# GetAllFaxApplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FaxApplication**](FaxApplication.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetAllFaxApplicationsResponse

`func NewGetAllFaxApplicationsResponse() *GetAllFaxApplicationsResponse`

NewGetAllFaxApplicationsResponse instantiates a new GetAllFaxApplicationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllFaxApplicationsResponseWithDefaults

`func NewGetAllFaxApplicationsResponseWithDefaults() *GetAllFaxApplicationsResponse`

NewGetAllFaxApplicationsResponseWithDefaults instantiates a new GetAllFaxApplicationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllFaxApplicationsResponse) GetData() []FaxApplication`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllFaxApplicationsResponse) GetDataOk() (*[]FaxApplication, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllFaxApplicationsResponse) SetData(v []FaxApplication)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllFaxApplicationsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetAllFaxApplicationsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllFaxApplicationsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllFaxApplicationsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAllFaxApplicationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


