# CampaignDeletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **float32** |  | 
**RecordType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignDeletionResponse

`func NewCampaignDeletionResponse(time float32, ) *CampaignDeletionResponse`

NewCampaignDeletionResponse instantiates a new CampaignDeletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDeletionResponseWithDefaults

`func NewCampaignDeletionResponseWithDefaults() *CampaignDeletionResponse`

NewCampaignDeletionResponseWithDefaults instantiates a new CampaignDeletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *CampaignDeletionResponse) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CampaignDeletionResponse) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CampaignDeletionResponse) SetTime(v float32)`

SetTime sets Time field to given value.


### GetRecordType

`func (o *CampaignDeletionResponse) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CampaignDeletionResponse) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CampaignDeletionResponse) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CampaignDeletionResponse) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetMessage

`func (o *CampaignDeletionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CampaignDeletionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CampaignDeletionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CampaignDeletionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


