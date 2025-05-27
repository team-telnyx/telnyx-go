# CreatePortingPhoneNumberExtensionRequestActivationRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | **int32** | Specifies the start of the activation range. Must be greater or equal the start of the extension range. | 
**EndAt** | **int32** | Specifies the end of the activation range. It must be no more than the end of the extension range. | 

## Methods

### NewCreatePortingPhoneNumberExtensionRequestActivationRangesInner

`func NewCreatePortingPhoneNumberExtensionRequestActivationRangesInner(startAt int32, endAt int32, ) *CreatePortingPhoneNumberExtensionRequestActivationRangesInner`

NewCreatePortingPhoneNumberExtensionRequestActivationRangesInner instantiates a new CreatePortingPhoneNumberExtensionRequestActivationRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingPhoneNumberExtensionRequestActivationRangesInnerWithDefaults

`func NewCreatePortingPhoneNumberExtensionRequestActivationRangesInnerWithDefaults() *CreatePortingPhoneNumberExtensionRequestActivationRangesInner`

NewCreatePortingPhoneNumberExtensionRequestActivationRangesInnerWithDefaults instantiates a new CreatePortingPhoneNumberExtensionRequestActivationRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) GetEndAt() int32`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) GetEndAtOk() (*int32, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreatePortingPhoneNumberExtensionRequestActivationRangesInner) SetEndAt(v int32)`

SetEndAt sets EndAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


