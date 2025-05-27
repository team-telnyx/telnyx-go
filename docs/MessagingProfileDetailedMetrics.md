# MessagingProfileDetailedMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overview** | Pointer to [**MessagingProfileHighLevelMetrics**](MessagingProfileHighLevelMetrics.md) |  | [optional] 
**Detailed** | Pointer to [**[]MessagingProfileDetailedMetric**](MessagingProfileDetailedMetric.md) |  | [optional] 

## Methods

### NewMessagingProfileDetailedMetrics

`func NewMessagingProfileDetailedMetrics() *MessagingProfileDetailedMetrics`

NewMessagingProfileDetailedMetrics instantiates a new MessagingProfileDetailedMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileDetailedMetricsWithDefaults

`func NewMessagingProfileDetailedMetricsWithDefaults() *MessagingProfileDetailedMetrics`

NewMessagingProfileDetailedMetricsWithDefaults instantiates a new MessagingProfileDetailedMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverview

`func (o *MessagingProfileDetailedMetrics) GetOverview() MessagingProfileHighLevelMetrics`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MessagingProfileDetailedMetrics) GetOverviewOk() (*MessagingProfileHighLevelMetrics, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MessagingProfileDetailedMetrics) SetOverview(v MessagingProfileHighLevelMetrics)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *MessagingProfileDetailedMetrics) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetDetailed

`func (o *MessagingProfileDetailedMetrics) GetDetailed() []MessagingProfileDetailedMetric`

GetDetailed returns the Detailed field if non-nil, zero value otherwise.

### GetDetailedOk

`func (o *MessagingProfileDetailedMetrics) GetDetailedOk() (*[]MessagingProfileDetailedMetric, bool)`

GetDetailedOk returns a tuple with the Detailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailed

`func (o *MessagingProfileDetailedMetrics) SetDetailed(v []MessagingProfileDetailedMetric)`

SetDetailed sets Detailed field to given value.

### HasDetailed

`func (o *MessagingProfileDetailedMetrics) HasDetailed() bool`

HasDetailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


