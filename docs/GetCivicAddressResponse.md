# GetCivicAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 

## Methods

### NewGetCivicAddressResponse

`func NewGetCivicAddressResponse() *GetCivicAddressResponse`

NewGetCivicAddressResponse instantiates a new GetCivicAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCivicAddressResponseWithDefaults

`func NewGetCivicAddressResponseWithDefaults() *GetCivicAddressResponse`

NewGetCivicAddressResponseWithDefaults instantiates a new GetCivicAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCivicAddressResponse) GetData() CivicAddress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCivicAddressResponse) GetDataOk() (*CivicAddress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCivicAddressResponse) SetData(v CivicAddress)`

SetData sets Data field to given value.

### HasData

`func (o *GetCivicAddressResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


