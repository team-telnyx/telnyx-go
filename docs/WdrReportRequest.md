# WdrReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **string** | ISO 8601 formatted date-time indicating the start time. | [optional] 
**EndTime** | Pointer to **string** | ISO 8601 formatted date-time indicating the end time. | [optional] 

## Methods

### NewWdrReportRequest

`func NewWdrReportRequest() *WdrReportRequest`

NewWdrReportRequest instantiates a new WdrReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWdrReportRequestWithDefaults

`func NewWdrReportRequestWithDefaults() *WdrReportRequest`

NewWdrReportRequestWithDefaults instantiates a new WdrReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *WdrReportRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WdrReportRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WdrReportRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WdrReportRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *WdrReportRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WdrReportRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WdrReportRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *WdrReportRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


