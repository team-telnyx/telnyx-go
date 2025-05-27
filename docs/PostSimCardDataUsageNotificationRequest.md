# PostSimCardDataUsageNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimCardId** | **string** | The identification UUID of the related SIM card resource. | 
**Threshold** | [**PostSimCardDataUsageNotificationRequestThreshold**](PostSimCardDataUsageNotificationRequestThreshold.md) |  | 

## Methods

### NewPostSimCardDataUsageNotificationRequest

`func NewPostSimCardDataUsageNotificationRequest(simCardId string, threshold PostSimCardDataUsageNotificationRequestThreshold, ) *PostSimCardDataUsageNotificationRequest`

NewPostSimCardDataUsageNotificationRequest instantiates a new PostSimCardDataUsageNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSimCardDataUsageNotificationRequestWithDefaults

`func NewPostSimCardDataUsageNotificationRequestWithDefaults() *PostSimCardDataUsageNotificationRequest`

NewPostSimCardDataUsageNotificationRequestWithDefaults instantiates a new PostSimCardDataUsageNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimCardId

`func (o *PostSimCardDataUsageNotificationRequest) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *PostSimCardDataUsageNotificationRequest) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *PostSimCardDataUsageNotificationRequest) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.


### GetThreshold

`func (o *PostSimCardDataUsageNotificationRequest) GetThreshold() PostSimCardDataUsageNotificationRequestThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *PostSimCardDataUsageNotificationRequest) GetThresholdOk() (*PostSimCardDataUsageNotificationRequestThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *PostSimCardDataUsageNotificationRequest) SetThreshold(v PostSimCardDataUsageNotificationRequestThreshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


