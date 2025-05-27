# ListPhoneNumbersResponse1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SlimPhoneNumberDetailed**](SlimPhoneNumberDetailed.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPhoneNumbersResponse1

`func NewListPhoneNumbersResponse1() *ListPhoneNumbersResponse1`

NewListPhoneNumbersResponse1 instantiates a new ListPhoneNumbersResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPhoneNumbersResponse1WithDefaults

`func NewListPhoneNumbersResponse1WithDefaults() *ListPhoneNumbersResponse1`

NewListPhoneNumbersResponse1WithDefaults instantiates a new ListPhoneNumbersResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPhoneNumbersResponse1) GetData() []SlimPhoneNumberDetailed`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPhoneNumbersResponse1) GetDataOk() (*[]SlimPhoneNumberDetailed, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPhoneNumbersResponse1) SetData(v []SlimPhoneNumberDetailed)`

SetData sets Data field to given value.

### HasData

`func (o *ListPhoneNumbersResponse1) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPhoneNumbersResponse1) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPhoneNumbersResponse1) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPhoneNumbersResponse1) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPhoneNumbersResponse1) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


