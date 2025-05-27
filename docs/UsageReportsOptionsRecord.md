# UsageReportsOptionsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Telnyx Product | [optional] 
**ProductDimensions** | Pointer to **[]string** | Telnyx Product Dimensions | [optional] 
**ProductMetrics** | Pointer to **[]string** | Telnyx Product Metrics | [optional] 
**RecordTypes** | Pointer to [**[]RecordType**](RecordType.md) | Subproducts if applicable | [optional] 

## Methods

### NewUsageReportsOptionsRecord

`func NewUsageReportsOptionsRecord() *UsageReportsOptionsRecord`

NewUsageReportsOptionsRecord instantiates a new UsageReportsOptionsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportsOptionsRecordWithDefaults

`func NewUsageReportsOptionsRecordWithDefaults() *UsageReportsOptionsRecord`

NewUsageReportsOptionsRecordWithDefaults instantiates a new UsageReportsOptionsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *UsageReportsOptionsRecord) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UsageReportsOptionsRecord) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UsageReportsOptionsRecord) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UsageReportsOptionsRecord) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductDimensions

`func (o *UsageReportsOptionsRecord) GetProductDimensions() []string`

GetProductDimensions returns the ProductDimensions field if non-nil, zero value otherwise.

### GetProductDimensionsOk

`func (o *UsageReportsOptionsRecord) GetProductDimensionsOk() (*[]string, bool)`

GetProductDimensionsOk returns a tuple with the ProductDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDimensions

`func (o *UsageReportsOptionsRecord) SetProductDimensions(v []string)`

SetProductDimensions sets ProductDimensions field to given value.

### HasProductDimensions

`func (o *UsageReportsOptionsRecord) HasProductDimensions() bool`

HasProductDimensions returns a boolean if a field has been set.

### GetProductMetrics

`func (o *UsageReportsOptionsRecord) GetProductMetrics() []string`

GetProductMetrics returns the ProductMetrics field if non-nil, zero value otherwise.

### GetProductMetricsOk

`func (o *UsageReportsOptionsRecord) GetProductMetricsOk() (*[]string, bool)`

GetProductMetricsOk returns a tuple with the ProductMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMetrics

`func (o *UsageReportsOptionsRecord) SetProductMetrics(v []string)`

SetProductMetrics sets ProductMetrics field to given value.

### HasProductMetrics

`func (o *UsageReportsOptionsRecord) HasProductMetrics() bool`

HasProductMetrics returns a boolean if a field has been set.

### GetRecordTypes

`func (o *UsageReportsOptionsRecord) GetRecordTypes() []RecordType`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *UsageReportsOptionsRecord) GetRecordTypesOk() (*[]RecordType, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *UsageReportsOptionsRecord) SetRecordTypes(v []RecordType)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *UsageReportsOptionsRecord) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


