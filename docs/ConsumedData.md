# ConsumedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **string** |  | [optional] [default to "MB"]
**Amount** | Pointer to **float64** |  | [optional] 

## Methods

### NewConsumedData

`func NewConsumedData() *ConsumedData`

NewConsumedData instantiates a new ConsumedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumedDataWithDefaults

`func NewConsumedDataWithDefaults() *ConsumedData`

NewConsumedDataWithDefaults instantiates a new ConsumedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *ConsumedData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ConsumedData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ConsumedData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *ConsumedData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetAmount

`func (o *ConsumedData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ConsumedData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ConsumedData) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ConsumedData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


