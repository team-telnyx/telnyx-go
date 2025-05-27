# CreatePortoutReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** | Identifies the type of report | 
**Params** | [**ExportPortoutsCSVReport**](ExportPortoutsCSVReport.md) |  | 

## Methods

### NewCreatePortoutReportRequest

`func NewCreatePortoutReportRequest(reportType string, params ExportPortoutsCSVReport, ) *CreatePortoutReportRequest`

NewCreatePortoutReportRequest instantiates a new CreatePortoutReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortoutReportRequestWithDefaults

`func NewCreatePortoutReportRequestWithDefaults() *CreatePortoutReportRequest`

NewCreatePortoutReportRequestWithDefaults instantiates a new CreatePortoutReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *CreatePortoutReportRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *CreatePortoutReportRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *CreatePortoutReportRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetParams

`func (o *CreatePortoutReportRequest) GetParams() ExportPortoutsCSVReport`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreatePortoutReportRequest) GetParamsOk() (*ExportPortoutsCSVReport, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreatePortoutReportRequest) SetParams(v ExportPortoutsCSVReport)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


