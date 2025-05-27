# SIMCardOrderPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCost** | Pointer to [**SIMCardOrderPreviewTotalCost**](SIMCardOrderPreviewTotalCost.md) |  | [optional] 
**ShippingCost** | Pointer to [**SIMCardOrderPreviewTotalCost**](SIMCardOrderPreviewTotalCost.md) |  | [optional] 
**SimCardsCost** | Pointer to [**SIMCardOrderPreviewTotalCost**](SIMCardOrderPreviewTotalCost.md) |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Quantity** | Pointer to **int32** | The amount of SIM cards requested in the SIM card order. | [optional] 

## Methods

### NewSIMCardOrderPreview

`func NewSIMCardOrderPreview() *SIMCardOrderPreview`

NewSIMCardOrderPreview instantiates a new SIMCardOrderPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardOrderPreviewWithDefaults

`func NewSIMCardOrderPreviewWithDefaults() *SIMCardOrderPreview`

NewSIMCardOrderPreviewWithDefaults instantiates a new SIMCardOrderPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCost

`func (o *SIMCardOrderPreview) GetTotalCost() SIMCardOrderPreviewTotalCost`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SIMCardOrderPreview) GetTotalCostOk() (*SIMCardOrderPreviewTotalCost, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SIMCardOrderPreview) SetTotalCost(v SIMCardOrderPreviewTotalCost)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *SIMCardOrderPreview) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetShippingCost

`func (o *SIMCardOrderPreview) GetShippingCost() SIMCardOrderPreviewTotalCost`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *SIMCardOrderPreview) GetShippingCostOk() (*SIMCardOrderPreviewTotalCost, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *SIMCardOrderPreview) SetShippingCost(v SIMCardOrderPreviewTotalCost)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *SIMCardOrderPreview) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetSimCardsCost

`func (o *SIMCardOrderPreview) GetSimCardsCost() SIMCardOrderPreviewTotalCost`

GetSimCardsCost returns the SimCardsCost field if non-nil, zero value otherwise.

### GetSimCardsCostOk

`func (o *SIMCardOrderPreview) GetSimCardsCostOk() (*SIMCardOrderPreviewTotalCost, bool)`

GetSimCardsCostOk returns a tuple with the SimCardsCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardsCost

`func (o *SIMCardOrderPreview) SetSimCardsCost(v SIMCardOrderPreviewTotalCost)`

SetSimCardsCost sets SimCardsCost field to given value.

### HasSimCardsCost

`func (o *SIMCardOrderPreview) HasSimCardsCost() bool`

HasSimCardsCost returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCardOrderPreview) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardOrderPreview) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardOrderPreview) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardOrderPreview) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetQuantity

`func (o *SIMCardOrderPreview) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SIMCardOrderPreview) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SIMCardOrderPreview) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SIMCardOrderPreview) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


