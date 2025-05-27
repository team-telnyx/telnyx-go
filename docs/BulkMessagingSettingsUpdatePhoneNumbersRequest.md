# BulkMessagingSettingsUpdatePhoneNumbersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProfileId** | **string** | Configure the messaging profile these phone numbers are assigned to:  * Set this field to &#x60;\&quot;\&quot;&#x60; to unassign each number from their respective messaging profile * Set this field to a quoted UUID of a messaging profile to assign these numbers to that messaging profile | 
**Numbers** | **[]string** | The list of phone numbers to update. | 

## Methods

### NewBulkMessagingSettingsUpdatePhoneNumbersRequest

`func NewBulkMessagingSettingsUpdatePhoneNumbersRequest(messagingProfileId string, numbers []string, ) *BulkMessagingSettingsUpdatePhoneNumbersRequest`

NewBulkMessagingSettingsUpdatePhoneNumbersRequest instantiates a new BulkMessagingSettingsUpdatePhoneNumbersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMessagingSettingsUpdatePhoneNumbersRequestWithDefaults

`func NewBulkMessagingSettingsUpdatePhoneNumbersRequestWithDefaults() *BulkMessagingSettingsUpdatePhoneNumbersRequest`

NewBulkMessagingSettingsUpdatePhoneNumbersRequestWithDefaults instantiates a new BulkMessagingSettingsUpdatePhoneNumbersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProfileId

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.


### GetNumbers

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) GetNumbers() []string`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) GetNumbersOk() (*[]string, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *BulkMessagingSettingsUpdatePhoneNumbersRequest) SetNumbers(v []string)`

SetNumbers sets Numbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


