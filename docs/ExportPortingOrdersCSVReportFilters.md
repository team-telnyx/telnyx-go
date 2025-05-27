# ExportPortingOrdersCSVReportFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusIn** | Pointer to **[]string** | The status of the porting orders to include in the report. | [optional] 
**CustomerReferenceIn** | Pointer to **[]string** | The customer reference of the porting orders to include in the report. | [optional] 
**CreatedAtLt** | Pointer to **time.Time** | The date and time the porting order was created before. | [optional] 
**CreatedAtGt** | Pointer to **time.Time** | The date and time the porting order was created after. | [optional] 

## Methods

### NewExportPortingOrdersCSVReportFilters

`func NewExportPortingOrdersCSVReportFilters() *ExportPortingOrdersCSVReportFilters`

NewExportPortingOrdersCSVReportFilters instantiates a new ExportPortingOrdersCSVReportFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportPortingOrdersCSVReportFiltersWithDefaults

`func NewExportPortingOrdersCSVReportFiltersWithDefaults() *ExportPortingOrdersCSVReportFilters`

NewExportPortingOrdersCSVReportFiltersWithDefaults instantiates a new ExportPortingOrdersCSVReportFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusIn

`func (o *ExportPortingOrdersCSVReportFilters) GetStatusIn() []string`

GetStatusIn returns the StatusIn field if non-nil, zero value otherwise.

### GetStatusInOk

`func (o *ExportPortingOrdersCSVReportFilters) GetStatusInOk() (*[]string, bool)`

GetStatusInOk returns a tuple with the StatusIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIn

`func (o *ExportPortingOrdersCSVReportFilters) SetStatusIn(v []string)`

SetStatusIn sets StatusIn field to given value.

### HasStatusIn

`func (o *ExportPortingOrdersCSVReportFilters) HasStatusIn() bool`

HasStatusIn returns a boolean if a field has been set.

### GetCustomerReferenceIn

`func (o *ExportPortingOrdersCSVReportFilters) GetCustomerReferenceIn() []string`

GetCustomerReferenceIn returns the CustomerReferenceIn field if non-nil, zero value otherwise.

### GetCustomerReferenceInOk

`func (o *ExportPortingOrdersCSVReportFilters) GetCustomerReferenceInOk() (*[]string, bool)`

GetCustomerReferenceInOk returns a tuple with the CustomerReferenceIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceIn

`func (o *ExportPortingOrdersCSVReportFilters) SetCustomerReferenceIn(v []string)`

SetCustomerReferenceIn sets CustomerReferenceIn field to given value.

### HasCustomerReferenceIn

`func (o *ExportPortingOrdersCSVReportFilters) HasCustomerReferenceIn() bool`

HasCustomerReferenceIn returns a boolean if a field has been set.

### GetCreatedAtLt

`func (o *ExportPortingOrdersCSVReportFilters) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *ExportPortingOrdersCSVReportFilters) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *ExportPortingOrdersCSVReportFilters) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *ExportPortingOrdersCSVReportFilters) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### GetCreatedAtGt

`func (o *ExportPortingOrdersCSVReportFilters) GetCreatedAtGt() time.Time`

GetCreatedAtGt returns the CreatedAtGt field if non-nil, zero value otherwise.

### GetCreatedAtGtOk

`func (o *ExportPortingOrdersCSVReportFilters) GetCreatedAtGtOk() (*time.Time, bool)`

GetCreatedAtGtOk returns a tuple with the CreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGt

`func (o *ExportPortingOrdersCSVReportFilters) SetCreatedAtGt(v time.Time)`

SetCreatedAtGt sets CreatedAtGt field to given value.

### HasCreatedAtGt

`func (o *ExportPortingOrdersCSVReportFilters) HasCreatedAtGt() bool`

HasCreatedAtGt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


