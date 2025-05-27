# CreatePortingReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportType** | **string** | Identifies the type of report | 
**Params** | [**ExportPortingOrdersCSVReport**](ExportPortingOrdersCSVReport.md) |  | 

## Methods

### NewCreatePortingReportRequest

`func NewCreatePortingReportRequest(reportType string, params ExportPortingOrdersCSVReport, ) *CreatePortingReportRequest`

NewCreatePortingReportRequest instantiates a new CreatePortingReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingReportRequestWithDefaults

`func NewCreatePortingReportRequestWithDefaults() *CreatePortingReportRequest`

NewCreatePortingReportRequestWithDefaults instantiates a new CreatePortingReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportType

`func (o *CreatePortingReportRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *CreatePortingReportRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *CreatePortingReportRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetParams

`func (o *CreatePortingReportRequest) GetParams() ExportPortingOrdersCSVReport`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreatePortingReportRequest) GetParamsOk() (*ExportPortingOrdersCSVReport, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreatePortingReportRequest) SetParams(v ExportPortingOrdersCSVReport)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


