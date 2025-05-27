# ListPortingReports200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortingReport**](PortingReport.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPortingReports200Response

`func NewListPortingReports200Response() *ListPortingReports200Response`

NewListPortingReports200Response instantiates a new ListPortingReports200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPortingReports200ResponseWithDefaults

`func NewListPortingReports200ResponseWithDefaults() *ListPortingReports200Response`

NewListPortingReports200ResponseWithDefaults instantiates a new ListPortingReports200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPortingReports200Response) GetData() []PortingReport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPortingReports200Response) GetDataOk() (*[]PortingReport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPortingReports200Response) SetData(v []PortingReport)`

SetData sets Data field to given value.

### HasData

`func (o *ListPortingReports200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPortingReports200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPortingReports200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPortingReports200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPortingReports200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


