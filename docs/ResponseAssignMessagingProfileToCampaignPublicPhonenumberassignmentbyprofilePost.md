# ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProfileId** | **string** | The ID of the messaging profile that you want to link to the specified campaign. | 
**TcrCampaignId** | Pointer to **string** | The TCR ID of the shared campaign you want to link to the specified messaging profile (for campaigns not created using Telnyx 10DLC services only). If you supply this ID in the request, do not also include a campaignId. | [optional] 
**CampaignId** | Pointer to **string** | The ID of the campaign you want to link to the specified messaging profile. If you supply this ID in the request, do not also include a tcrCampaignId. | [optional] 
**TaskId** | **string** | The ID of the task associated with assigning a messaging profile to a campaign. | 
**Message** | **string** |  | 

## Methods

### NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost

`func NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost(messagingProfileId string, taskId string, message string, ) *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost`

NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost instantiates a new ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePostWithDefaults

`func NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePostWithDefaults() *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost`

NewResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePostWithDefaults instantiates a new ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProfileId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.


### GetTcrCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetTcrCampaignId() string`

GetTcrCampaignId returns the TcrCampaignId field if non-nil, zero value otherwise.

### GetTcrCampaignIdOk

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetTcrCampaignIdOk() (*string, bool)`

GetTcrCampaignIdOk returns a tuple with the TcrCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) SetTcrCampaignId(v string)`

SetTcrCampaignId sets TcrCampaignId field to given value.

### HasTcrCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) HasTcrCampaignId() bool`

HasTcrCampaignId returns a boolean if a field has been set.

### GetCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetTaskId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetMessage

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseAssignMessagingProfileToCampaignPublicPhonenumberassignmentbyprofilePost) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


