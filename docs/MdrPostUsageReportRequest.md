# MdrPostUsageReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 
**AggregationType** | **string** |  | 
**Profiles** | Pointer to **string** |  | [optional] 

## Methods

### NewMdrPostUsageReportRequest

`func NewMdrPostUsageReportRequest(startDate time.Time, endDate time.Time, aggregationType string, ) *MdrPostUsageReportRequest`

NewMdrPostUsageReportRequest instantiates a new MdrPostUsageReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdrPostUsageReportRequestWithDefaults

`func NewMdrPostUsageReportRequestWithDefaults() *MdrPostUsageReportRequest`

NewMdrPostUsageReportRequestWithDefaults instantiates a new MdrPostUsageReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *MdrPostUsageReportRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MdrPostUsageReportRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MdrPostUsageReportRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *MdrPostUsageReportRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MdrPostUsageReportRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MdrPostUsageReportRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetAggregationType

`func (o *MdrPostUsageReportRequest) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *MdrPostUsageReportRequest) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *MdrPostUsageReportRequest) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.


### GetProfiles

`func (o *MdrPostUsageReportRequest) GetProfiles() string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *MdrPostUsageReportRequest) GetProfilesOk() (*string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *MdrPostUsageReportRequest) SetProfiles(v string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *MdrPostUsageReportRequest) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


