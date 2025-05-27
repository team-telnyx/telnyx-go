# ListAvailablePhoneNumbersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AvailablePhoneNumber**](AvailablePhoneNumber.md) |  | [optional] 
**Meta** | Pointer to [**AvailablePhoneNumbersMetadata**](AvailablePhoneNumbersMetadata.md) |  | [optional] 

## Methods

### NewListAvailablePhoneNumbersResponse

`func NewListAvailablePhoneNumbersResponse() *ListAvailablePhoneNumbersResponse`

NewListAvailablePhoneNumbersResponse instantiates a new ListAvailablePhoneNumbersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailablePhoneNumbersResponseWithDefaults

`func NewListAvailablePhoneNumbersResponseWithDefaults() *ListAvailablePhoneNumbersResponse`

NewListAvailablePhoneNumbersResponseWithDefaults instantiates a new ListAvailablePhoneNumbersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAvailablePhoneNumbersResponse) GetData() []AvailablePhoneNumber`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAvailablePhoneNumbersResponse) GetDataOk() (*[]AvailablePhoneNumber, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAvailablePhoneNumbersResponse) SetData(v []AvailablePhoneNumber)`

SetData sets Data field to given value.

### HasData

`func (o *ListAvailablePhoneNumbersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListAvailablePhoneNumbersResponse) GetMeta() AvailablePhoneNumbersMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAvailablePhoneNumbersResponse) GetMetaOk() (*AvailablePhoneNumbersMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAvailablePhoneNumbersResponse) SetMeta(v AvailablePhoneNumbersMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAvailablePhoneNumbersResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


