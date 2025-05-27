# CdrUsageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Connections** | Pointer to **[]int64** |  | [optional] 
**AggregationType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**ProductBreakdown** | Pointer to **string** |  | [optional] 

## Methods

### NewCdrUsageReportResponse

`func NewCdrUsageReportResponse() *CdrUsageReportResponse`

NewCdrUsageReportResponse instantiates a new CdrUsageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdrUsageReportResponseWithDefaults

`func NewCdrUsageReportResponseWithDefaults() *CdrUsageReportResponse`

NewCdrUsageReportResponseWithDefaults instantiates a new CdrUsageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdrUsageReportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdrUsageReportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdrUsageReportResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdrUsageReportResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *CdrUsageReportResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CdrUsageReportResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CdrUsageReportResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CdrUsageReportResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CdrUsageReportResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CdrUsageReportResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CdrUsageReportResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CdrUsageReportResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetConnections

`func (o *CdrUsageReportResponse) GetConnections() []int64`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CdrUsageReportResponse) GetConnectionsOk() (*[]int64, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CdrUsageReportResponse) SetConnections(v []int64)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *CdrUsageReportResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetAggregationType

`func (o *CdrUsageReportResponse) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *CdrUsageReportResponse) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *CdrUsageReportResponse) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *CdrUsageReportResponse) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetStatus

`func (o *CdrUsageReportResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdrUsageReportResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdrUsageReportResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdrUsageReportResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReportUrl

`func (o *CdrUsageReportResponse) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *CdrUsageReportResponse) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *CdrUsageReportResponse) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *CdrUsageReportResponse) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetResult

`func (o *CdrUsageReportResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CdrUsageReportResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CdrUsageReportResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *CdrUsageReportResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CdrUsageReportResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CdrUsageReportResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CdrUsageReportResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CdrUsageReportResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CdrUsageReportResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CdrUsageReportResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CdrUsageReportResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CdrUsageReportResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRecordType

`func (o *CdrUsageReportResponse) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CdrUsageReportResponse) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CdrUsageReportResponse) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CdrUsageReportResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetProductBreakdown

`func (o *CdrUsageReportResponse) GetProductBreakdown() string`

GetProductBreakdown returns the ProductBreakdown field if non-nil, zero value otherwise.

### GetProductBreakdownOk

`func (o *CdrUsageReportResponse) GetProductBreakdownOk() (*string, bool)`

GetProductBreakdownOk returns a tuple with the ProductBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBreakdown

`func (o *CdrUsageReportResponse) SetProductBreakdown(v string)`

SetProductBreakdown sets ProductBreakdown field to given value.

### HasProductBreakdown

`func (o *CdrUsageReportResponse) HasProductBreakdown() bool`

HasProductBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


