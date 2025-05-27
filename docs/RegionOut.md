# RegionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | Pointer to **string** | The region interface is deployed to. | [optional] 
**Region** | Pointer to [**RegionOutRegion**](RegionOutRegion.md) |  | [optional] 

## Methods

### NewRegionOut

`func NewRegionOut() *RegionOut`

NewRegionOut instantiates a new RegionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionOutWithDefaults

`func NewRegionOutWithDefaults() *RegionOut`

NewRegionOutWithDefaults instantiates a new RegionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionCode

`func (o *RegionOut) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *RegionOut) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *RegionOut) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *RegionOut) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegion

`func (o *RegionOut) GetRegion() RegionOutRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionOut) GetRegionOk() (*RegionOutRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionOut) SetRegion(v RegionOutRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RegionOut) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


