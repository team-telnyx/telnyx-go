# PhoneNumberCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** |  | 
**BrandId** | Pointer to **string** | Brand ID. Empty if the number is associated to a shared campaign. | [optional] 
**TcrBrandId** | Pointer to **string** | TCR&#39;s alphanumeric ID for the brand. | [optional] 
**CampaignId** | **string** | For shared campaigns, this is the TCR campaign ID, otherwise it is the campaign ID  | 
**TcrCampaignId** | Pointer to **string** | TCR&#39;s alphanumeric ID for the campaign. | [optional] 
**TelnyxCampaignId** | Pointer to **string** | Campaign ID. Empty if the number is associated to a shared campaign. | [optional] 
**AssignmentStatus** | Pointer to **string** | The assignment status of the number. | [optional] 
**FailureReasons** | Pointer to **interface{}** | Extra info about a failure to assign/unassign a number. Relevant only if the assignmentStatus is either FAILED_ASSIGNMENT or FAILED_UNASSIGNMENT | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewPhoneNumberCampaign

`func NewPhoneNumberCampaign(phoneNumber string, campaignId string, createdAt string, updatedAt string, ) *PhoneNumberCampaign`

NewPhoneNumberCampaign instantiates a new PhoneNumberCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberCampaignWithDefaults

`func NewPhoneNumberCampaignWithDefaults() *PhoneNumberCampaign`

NewPhoneNumberCampaignWithDefaults instantiates a new PhoneNumberCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *PhoneNumberCampaign) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberCampaign) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberCampaign) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetBrandId

`func (o *PhoneNumberCampaign) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *PhoneNumberCampaign) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *PhoneNumberCampaign) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *PhoneNumberCampaign) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetTcrBrandId

`func (o *PhoneNumberCampaign) GetTcrBrandId() string`

GetTcrBrandId returns the TcrBrandId field if non-nil, zero value otherwise.

### GetTcrBrandIdOk

`func (o *PhoneNumberCampaign) GetTcrBrandIdOk() (*string, bool)`

GetTcrBrandIdOk returns a tuple with the TcrBrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrBrandId

`func (o *PhoneNumberCampaign) SetTcrBrandId(v string)`

SetTcrBrandId sets TcrBrandId field to given value.

### HasTcrBrandId

`func (o *PhoneNumberCampaign) HasTcrBrandId() bool`

HasTcrBrandId returns a boolean if a field has been set.

### GetCampaignId

`func (o *PhoneNumberCampaign) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *PhoneNumberCampaign) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *PhoneNumberCampaign) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetTcrCampaignId

`func (o *PhoneNumberCampaign) GetTcrCampaignId() string`

GetTcrCampaignId returns the TcrCampaignId field if non-nil, zero value otherwise.

### GetTcrCampaignIdOk

`func (o *PhoneNumberCampaign) GetTcrCampaignIdOk() (*string, bool)`

GetTcrCampaignIdOk returns a tuple with the TcrCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignId

`func (o *PhoneNumberCampaign) SetTcrCampaignId(v string)`

SetTcrCampaignId sets TcrCampaignId field to given value.

### HasTcrCampaignId

`func (o *PhoneNumberCampaign) HasTcrCampaignId() bool`

HasTcrCampaignId returns a boolean if a field has been set.

### GetTelnyxCampaignId

`func (o *PhoneNumberCampaign) GetTelnyxCampaignId() string`

GetTelnyxCampaignId returns the TelnyxCampaignId field if non-nil, zero value otherwise.

### GetTelnyxCampaignIdOk

`func (o *PhoneNumberCampaign) GetTelnyxCampaignIdOk() (*string, bool)`

GetTelnyxCampaignIdOk returns a tuple with the TelnyxCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxCampaignId

`func (o *PhoneNumberCampaign) SetTelnyxCampaignId(v string)`

SetTelnyxCampaignId sets TelnyxCampaignId field to given value.

### HasTelnyxCampaignId

`func (o *PhoneNumberCampaign) HasTelnyxCampaignId() bool`

HasTelnyxCampaignId returns a boolean if a field has been set.

### GetAssignmentStatus

`func (o *PhoneNumberCampaign) GetAssignmentStatus() string`

GetAssignmentStatus returns the AssignmentStatus field if non-nil, zero value otherwise.

### GetAssignmentStatusOk

`func (o *PhoneNumberCampaign) GetAssignmentStatusOk() (*string, bool)`

GetAssignmentStatusOk returns a tuple with the AssignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentStatus

`func (o *PhoneNumberCampaign) SetAssignmentStatus(v string)`

SetAssignmentStatus sets AssignmentStatus field to given value.

### HasAssignmentStatus

`func (o *PhoneNumberCampaign) HasAssignmentStatus() bool`

HasAssignmentStatus returns a boolean if a field has been set.

### GetFailureReasons

`func (o *PhoneNumberCampaign) GetFailureReasons() interface{}`

GetFailureReasons returns the FailureReasons field if non-nil, zero value otherwise.

### GetFailureReasonsOk

`func (o *PhoneNumberCampaign) GetFailureReasonsOk() (*interface{}, bool)`

GetFailureReasonsOk returns a tuple with the FailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasons

`func (o *PhoneNumberCampaign) SetFailureReasons(v interface{})`

SetFailureReasons sets FailureReasons field to given value.

### HasFailureReasons

`func (o *PhoneNumberCampaign) HasFailureReasons() bool`

HasFailureReasons returns a boolean if a field has been set.

### SetFailureReasonsNil

`func (o *PhoneNumberCampaign) SetFailureReasonsNil(b bool)`

 SetFailureReasonsNil sets the value for FailureReasons to be an explicit nil

### UnsetFailureReasons
`func (o *PhoneNumberCampaign) UnsetFailureReasons()`

UnsetFailureReasons ensures that no value is present for FailureReasons, not even an explicit nil
### GetCreatedAt

`func (o *PhoneNumberCampaign) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneNumberCampaign) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneNumberCampaign) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PhoneNumberCampaign) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PhoneNumberCampaign) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PhoneNumberCampaign) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


