# MdrUsageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Connections** | Pointer to **[]int64** |  | [optional] 
**AggregationType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]MdrUsageRecord**](MdrUsageRecord.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Profiles** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewMdrUsageReportResponse

`func NewMdrUsageReportResponse() *MdrUsageReportResponse`

NewMdrUsageReportResponse instantiates a new MdrUsageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdrUsageReportResponseWithDefaults

`func NewMdrUsageReportResponseWithDefaults() *MdrUsageReportResponse`

NewMdrUsageReportResponseWithDefaults instantiates a new MdrUsageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MdrUsageReportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MdrUsageReportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MdrUsageReportResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MdrUsageReportResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartDate

`func (o *MdrUsageReportResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MdrUsageReportResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MdrUsageReportResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *MdrUsageReportResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *MdrUsageReportResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MdrUsageReportResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MdrUsageReportResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *MdrUsageReportResponse) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetConnections

`func (o *MdrUsageReportResponse) GetConnections() []int64`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *MdrUsageReportResponse) GetConnectionsOk() (*[]int64, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *MdrUsageReportResponse) SetConnections(v []int64)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *MdrUsageReportResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetAggregationType

`func (o *MdrUsageReportResponse) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *MdrUsageReportResponse) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *MdrUsageReportResponse) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *MdrUsageReportResponse) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetStatus

`func (o *MdrUsageReportResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MdrUsageReportResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MdrUsageReportResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MdrUsageReportResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReportUrl

`func (o *MdrUsageReportResponse) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *MdrUsageReportResponse) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *MdrUsageReportResponse) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *MdrUsageReportResponse) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetResult

`func (o *MdrUsageReportResponse) GetResult() []MdrUsageRecord`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MdrUsageReportResponse) GetResultOk() (*[]MdrUsageRecord, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MdrUsageReportResponse) SetResult(v []MdrUsageRecord)`

SetResult sets Result field to given value.

### HasResult

`func (o *MdrUsageReportResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MdrUsageReportResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MdrUsageReportResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MdrUsageReportResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MdrUsageReportResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MdrUsageReportResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MdrUsageReportResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MdrUsageReportResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MdrUsageReportResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProfiles

`func (o *MdrUsageReportResponse) GetProfiles() string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MdrUsageReportResponse) GetProfilesOk() (*string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MdrUsageReportResponse) SetProfiles(v string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MdrUsageReportResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetRecordType

`func (o *MdrUsageReportResponse) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MdrUsageReportResponse) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MdrUsageReportResponse) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MdrUsageReportResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


