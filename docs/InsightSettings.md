# InsightSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsightGroupId** | Pointer to **string** | Reference to an Insight Group. Insights in this group will be run automatically for all the assistant&#39;s conversations. | [optional] 

## Methods

### NewInsightSettings

`func NewInsightSettings() *InsightSettings`

NewInsightSettings instantiates a new InsightSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightSettingsWithDefaults

`func NewInsightSettingsWithDefaults() *InsightSettings`

NewInsightSettingsWithDefaults instantiates a new InsightSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsightGroupId

`func (o *InsightSettings) GetInsightGroupId() string`

GetInsightGroupId returns the InsightGroupId field if non-nil, zero value otherwise.

### GetInsightGroupIdOk

`func (o *InsightSettings) GetInsightGroupIdOk() (*string, bool)`

GetInsightGroupIdOk returns a tuple with the InsightGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightGroupId

`func (o *InsightSettings) SetInsightGroupId(v string)`

SetInsightGroupId sets InsightGroupId field to given value.

### HasInsightGroupId

`func (o *InsightSettings) HasInsightGroupId() bool`

HasInsightGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


