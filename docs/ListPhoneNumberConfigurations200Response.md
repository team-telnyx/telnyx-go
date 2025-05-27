# ListPhoneNumberConfigurations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortingPhoneNumberConfiguration**](PortingPhoneNumberConfiguration.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPhoneNumberConfigurations200Response

`func NewListPhoneNumberConfigurations200Response() *ListPhoneNumberConfigurations200Response`

NewListPhoneNumberConfigurations200Response instantiates a new ListPhoneNumberConfigurations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPhoneNumberConfigurations200ResponseWithDefaults

`func NewListPhoneNumberConfigurations200ResponseWithDefaults() *ListPhoneNumberConfigurations200Response`

NewListPhoneNumberConfigurations200ResponseWithDefaults instantiates a new ListPhoneNumberConfigurations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPhoneNumberConfigurations200Response) GetData() []PortingPhoneNumberConfiguration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPhoneNumberConfigurations200Response) GetDataOk() (*[]PortingPhoneNumberConfiguration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPhoneNumberConfigurations200Response) SetData(v []PortingPhoneNumberConfiguration)`

SetData sets Data field to given value.

### HasData

`func (o *ListPhoneNumberConfigurations200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPhoneNumberConfigurations200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPhoneNumberConfigurations200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPhoneNumberConfigurations200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPhoneNumberConfigurations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


