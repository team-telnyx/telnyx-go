# UplinkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Uplink data | [optional] 
**Unit** | Pointer to **string** | Transmission unit | [optional] [default to "MB"]

## Methods

### NewUplinkData

`func NewUplinkData() *UplinkData`

NewUplinkData instantiates a new UplinkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplinkDataWithDefaults

`func NewUplinkDataWithDefaults() *UplinkData`

NewUplinkDataWithDefaults instantiates a new UplinkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UplinkData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UplinkData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UplinkData) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UplinkData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUnit

`func (o *UplinkData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UplinkData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UplinkData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UplinkData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


