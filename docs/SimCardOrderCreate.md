# SimCardOrderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressId** | **string** | Uniquely identifies the address for the order. | 
**Quantity** | **int32** | The amount of SIM cards to order. | 

## Methods

### NewSimCardOrderCreate

`func NewSimCardOrderCreate(addressId string, quantity int32, ) *SimCardOrderCreate`

NewSimCardOrderCreate instantiates a new SimCardOrderCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimCardOrderCreateWithDefaults

`func NewSimCardOrderCreateWithDefaults() *SimCardOrderCreate`

NewSimCardOrderCreateWithDefaults instantiates a new SimCardOrderCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressId

`func (o *SimCardOrderCreate) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *SimCardOrderCreate) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *SimCardOrderCreate) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.


### GetQuantity

`func (o *SimCardOrderCreate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SimCardOrderCreate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SimCardOrderCreate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


