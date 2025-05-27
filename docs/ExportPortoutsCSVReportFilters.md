# ExportPortoutsCSVReportFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusIn** | Pointer to **[]string** | The status of the port-outs to include in the report. | [optional] 
**CustomerReferenceIn** | Pointer to **[]string** | The customer reference of the port-outs to include in the report. | [optional] 
**EndUserName** | Pointer to **string** | The end user name of the port-outs to include in the report. | [optional] 
**PhoneNumbersOverlaps** | Pointer to **[]string** | A list of phone numbers that the port-outs phone numbers must overlap with. | [optional] 
**CreatedAtLt** | Pointer to **time.Time** | The date and time the port-out was created before. | [optional] 
**CreatedAtGt** | Pointer to **time.Time** | The date and time the port-out was created after. | [optional] 

## Methods

### NewExportPortoutsCSVReportFilters

`func NewExportPortoutsCSVReportFilters() *ExportPortoutsCSVReportFilters`

NewExportPortoutsCSVReportFilters instantiates a new ExportPortoutsCSVReportFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportPortoutsCSVReportFiltersWithDefaults

`func NewExportPortoutsCSVReportFiltersWithDefaults() *ExportPortoutsCSVReportFilters`

NewExportPortoutsCSVReportFiltersWithDefaults instantiates a new ExportPortoutsCSVReportFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusIn

`func (o *ExportPortoutsCSVReportFilters) GetStatusIn() []string`

GetStatusIn returns the StatusIn field if non-nil, zero value otherwise.

### GetStatusInOk

`func (o *ExportPortoutsCSVReportFilters) GetStatusInOk() (*[]string, bool)`

GetStatusInOk returns a tuple with the StatusIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusIn

`func (o *ExportPortoutsCSVReportFilters) SetStatusIn(v []string)`

SetStatusIn sets StatusIn field to given value.

### HasStatusIn

`func (o *ExportPortoutsCSVReportFilters) HasStatusIn() bool`

HasStatusIn returns a boolean if a field has been set.

### GetCustomerReferenceIn

`func (o *ExportPortoutsCSVReportFilters) GetCustomerReferenceIn() []string`

GetCustomerReferenceIn returns the CustomerReferenceIn field if non-nil, zero value otherwise.

### GetCustomerReferenceInOk

`func (o *ExportPortoutsCSVReportFilters) GetCustomerReferenceInOk() (*[]string, bool)`

GetCustomerReferenceInOk returns a tuple with the CustomerReferenceIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceIn

`func (o *ExportPortoutsCSVReportFilters) SetCustomerReferenceIn(v []string)`

SetCustomerReferenceIn sets CustomerReferenceIn field to given value.

### HasCustomerReferenceIn

`func (o *ExportPortoutsCSVReportFilters) HasCustomerReferenceIn() bool`

HasCustomerReferenceIn returns a boolean if a field has been set.

### GetEndUserName

`func (o *ExportPortoutsCSVReportFilters) GetEndUserName() string`

GetEndUserName returns the EndUserName field if non-nil, zero value otherwise.

### GetEndUserNameOk

`func (o *ExportPortoutsCSVReportFilters) GetEndUserNameOk() (*string, bool)`

GetEndUserNameOk returns a tuple with the EndUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserName

`func (o *ExportPortoutsCSVReportFilters) SetEndUserName(v string)`

SetEndUserName sets EndUserName field to given value.

### HasEndUserName

`func (o *ExportPortoutsCSVReportFilters) HasEndUserName() bool`

HasEndUserName returns a boolean if a field has been set.

### GetPhoneNumbersOverlaps

`func (o *ExportPortoutsCSVReportFilters) GetPhoneNumbersOverlaps() []string`

GetPhoneNumbersOverlaps returns the PhoneNumbersOverlaps field if non-nil, zero value otherwise.

### GetPhoneNumbersOverlapsOk

`func (o *ExportPortoutsCSVReportFilters) GetPhoneNumbersOverlapsOk() (*[]string, bool)`

GetPhoneNumbersOverlapsOk returns a tuple with the PhoneNumbersOverlaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersOverlaps

`func (o *ExportPortoutsCSVReportFilters) SetPhoneNumbersOverlaps(v []string)`

SetPhoneNumbersOverlaps sets PhoneNumbersOverlaps field to given value.

### HasPhoneNumbersOverlaps

`func (o *ExportPortoutsCSVReportFilters) HasPhoneNumbersOverlaps() bool`

HasPhoneNumbersOverlaps returns a boolean if a field has been set.

### GetCreatedAtLt

`func (o *ExportPortoutsCSVReportFilters) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *ExportPortoutsCSVReportFilters) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *ExportPortoutsCSVReportFilters) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *ExportPortoutsCSVReportFilters) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### GetCreatedAtGt

`func (o *ExportPortoutsCSVReportFilters) GetCreatedAtGt() time.Time`

GetCreatedAtGt returns the CreatedAtGt field if non-nil, zero value otherwise.

### GetCreatedAtGtOk

`func (o *ExportPortoutsCSVReportFilters) GetCreatedAtGtOk() (*time.Time, bool)`

GetCreatedAtGtOk returns a tuple with the CreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGt

`func (o *ExportPortoutsCSVReportFilters) SetCreatedAtGt(v time.Time)`

SetCreatedAtGt sets CreatedAtGt field to given value.

### HasCreatedAtGt

`func (o *ExportPortoutsCSVReportFilters) HasCreatedAtGt() bool`

HasCreatedAtGt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


