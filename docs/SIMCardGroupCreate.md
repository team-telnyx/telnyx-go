# SIMCardGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A user friendly name for the SIM card group. | 
**DataLimit** | Pointer to [**SIMCardGroupDataLimit**](SIMCardGroupDataLimit.md) |  | [optional] 

## Methods

### NewSIMCardGroupCreate

`func NewSIMCardGroupCreate(name string, ) *SIMCardGroupCreate`

NewSIMCardGroupCreate instantiates a new SIMCardGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardGroupCreateWithDefaults

`func NewSIMCardGroupCreateWithDefaults() *SIMCardGroupCreate`

NewSIMCardGroupCreateWithDefaults instantiates a new SIMCardGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SIMCardGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SIMCardGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SIMCardGroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDataLimit

`func (o *SIMCardGroupCreate) GetDataLimit() SIMCardGroupDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SIMCardGroupCreate) GetDataLimitOk() (*SIMCardGroupDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SIMCardGroupCreate) SetDataLimit(v SIMCardGroupDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SIMCardGroupCreate) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


