# SIMCardGroupPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user friendly name for the SIM card group. | [optional] 
**DataLimit** | Pointer to [**SIMCardGroupDataLimit**](SIMCardGroupDataLimit.md) |  | [optional] 

## Methods

### NewSIMCardGroupPatch

`func NewSIMCardGroupPatch() *SIMCardGroupPatch`

NewSIMCardGroupPatch instantiates a new SIMCardGroupPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardGroupPatchWithDefaults

`func NewSIMCardGroupPatchWithDefaults() *SIMCardGroupPatch`

NewSIMCardGroupPatchWithDefaults instantiates a new SIMCardGroupPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SIMCardGroupPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SIMCardGroupPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SIMCardGroupPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SIMCardGroupPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataLimit

`func (o *SIMCardGroupPatch) GetDataLimit() SIMCardGroupDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SIMCardGroupPatch) GetDataLimitOk() (*SIMCardGroupDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SIMCardGroupPatch) SetDataLimit(v SIMCardGroupDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SIMCardGroupPatch) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


