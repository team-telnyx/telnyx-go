# CreatePortingPhoneNumberBlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumberRange** | [**CreatePortingPhoneNumberBlockRequestPhoneNumberRange**](CreatePortingPhoneNumberBlockRequestPhoneNumberRange.md) |  | 
**ActivationRanges** | [**[]CreatePortingPhoneNumberBlockRequestActivationRangesInner**](CreatePortingPhoneNumberBlockRequestActivationRangesInner.md) | Specifies the activation ranges for this porting phone number block. The activation range must be within the block range and should not overlap with other activation ranges. | 

## Methods

### NewCreatePortingPhoneNumberBlockRequest

`func NewCreatePortingPhoneNumberBlockRequest(phoneNumberRange CreatePortingPhoneNumberBlockRequestPhoneNumberRange, activationRanges []CreatePortingPhoneNumberBlockRequestActivationRangesInner, ) *CreatePortingPhoneNumberBlockRequest`

NewCreatePortingPhoneNumberBlockRequest instantiates a new CreatePortingPhoneNumberBlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingPhoneNumberBlockRequestWithDefaults

`func NewCreatePortingPhoneNumberBlockRequestWithDefaults() *CreatePortingPhoneNumberBlockRequest`

NewCreatePortingPhoneNumberBlockRequestWithDefaults instantiates a new CreatePortingPhoneNumberBlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumberRange

`func (o *CreatePortingPhoneNumberBlockRequest) GetPhoneNumberRange() CreatePortingPhoneNumberBlockRequestPhoneNumberRange`

GetPhoneNumberRange returns the PhoneNumberRange field if non-nil, zero value otherwise.

### GetPhoneNumberRangeOk

`func (o *CreatePortingPhoneNumberBlockRequest) GetPhoneNumberRangeOk() (*CreatePortingPhoneNumberBlockRequestPhoneNumberRange, bool)`

GetPhoneNumberRangeOk returns a tuple with the PhoneNumberRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberRange

`func (o *CreatePortingPhoneNumberBlockRequest) SetPhoneNumberRange(v CreatePortingPhoneNumberBlockRequestPhoneNumberRange)`

SetPhoneNumberRange sets PhoneNumberRange field to given value.


### GetActivationRanges

`func (o *CreatePortingPhoneNumberBlockRequest) GetActivationRanges() []CreatePortingPhoneNumberBlockRequestActivationRangesInner`

GetActivationRanges returns the ActivationRanges field if non-nil, zero value otherwise.

### GetActivationRangesOk

`func (o *CreatePortingPhoneNumberBlockRequest) GetActivationRangesOk() (*[]CreatePortingPhoneNumberBlockRequestActivationRangesInner, bool)`

GetActivationRangesOk returns a tuple with the ActivationRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRanges

`func (o *CreatePortingPhoneNumberBlockRequest) SetActivationRanges(v []CreatePortingPhoneNumberBlockRequestActivationRangesInner)`

SetActivationRanges sets ActivationRanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


