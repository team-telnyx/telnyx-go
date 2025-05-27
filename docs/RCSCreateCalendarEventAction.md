# RCSCreateCalendarEventAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **string** | Event title. Maximum 100 characters. | [optional] 
**Description** | Pointer to **string** | Event description. Maximum 500 characters. | [optional] 

## Methods

### NewRCSCreateCalendarEventAction

`func NewRCSCreateCalendarEventAction() *RCSCreateCalendarEventAction`

NewRCSCreateCalendarEventAction instantiates a new RCSCreateCalendarEventAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSCreateCalendarEventActionWithDefaults

`func NewRCSCreateCalendarEventActionWithDefaults() *RCSCreateCalendarEventAction`

NewRCSCreateCalendarEventActionWithDefaults instantiates a new RCSCreateCalendarEventAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *RCSCreateCalendarEventAction) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RCSCreateCalendarEventAction) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RCSCreateCalendarEventAction) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RCSCreateCalendarEventAction) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *RCSCreateCalendarEventAction) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RCSCreateCalendarEventAction) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RCSCreateCalendarEventAction) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RCSCreateCalendarEventAction) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTitle

`func (o *RCSCreateCalendarEventAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RCSCreateCalendarEventAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RCSCreateCalendarEventAction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RCSCreateCalendarEventAction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RCSCreateCalendarEventAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RCSCreateCalendarEventAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RCSCreateCalendarEventAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RCSCreateCalendarEventAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


