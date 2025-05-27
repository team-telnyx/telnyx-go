# PortoutReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the report. | [optional] 
**ReportType** | Pointer to **string** | Identifies the type of report | [optional] 
**Status** | Pointer to **string** | The current status of the report generation. | [optional] 
**Params** | Pointer to [**ExportPortoutsCSVReport**](ExportPortoutsCSVReport.md) |  | [optional] 
**DocumentId** | Pointer to **string** | Identifies the document that was uploaded when report was generated. This field is only populated when the report is under completed status. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortoutReport

`func NewPortoutReport() *PortoutReport`

NewPortoutReport instantiates a new PortoutReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortoutReportWithDefaults

`func NewPortoutReportWithDefaults() *PortoutReport`

NewPortoutReportWithDefaults instantiates a new PortoutReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortoutReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortoutReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortoutReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortoutReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReportType

`func (o *PortoutReport) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *PortoutReport) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *PortoutReport) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *PortoutReport) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetStatus

`func (o *PortoutReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortoutReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortoutReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortoutReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParams

`func (o *PortoutReport) GetParams() ExportPortoutsCSVReport`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PortoutReport) GetParamsOk() (*ExportPortoutsCSVReport, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PortoutReport) SetParams(v ExportPortoutsCSVReport)`

SetParams sets Params field to given value.

### HasParams

`func (o *PortoutReport) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetDocumentId

`func (o *PortoutReport) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *PortoutReport) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *PortoutReport) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *PortoutReport) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetRecordType

`func (o *PortoutReport) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortoutReport) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortoutReport) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortoutReport) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortoutReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortoutReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortoutReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortoutReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortoutReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortoutReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortoutReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortoutReport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


