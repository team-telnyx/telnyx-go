# CreatePortingPhoneNumberBlockRequestActivationRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | **string** | Specifies the start of the activation range. Must be greater or equal the start of the extension range. | 
**EndAt** | **string** | Specifies the end of the activation range. It must be no more than the end of the extension range. | 

## Methods

### NewCreatePortingPhoneNumberBlockRequestActivationRangesInner

`func NewCreatePortingPhoneNumberBlockRequestActivationRangesInner(startAt string, endAt string, ) *CreatePortingPhoneNumberBlockRequestActivationRangesInner`

NewCreatePortingPhoneNumberBlockRequestActivationRangesInner instantiates a new CreatePortingPhoneNumberBlockRequestActivationRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingPhoneNumberBlockRequestActivationRangesInnerWithDefaults

`func NewCreatePortingPhoneNumberBlockRequestActivationRangesInnerWithDefaults() *CreatePortingPhoneNumberBlockRequestActivationRangesInner`

NewCreatePortingPhoneNumberBlockRequestActivationRangesInnerWithDefaults instantiates a new CreatePortingPhoneNumberBlockRequestActivationRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *CreatePortingPhoneNumberBlockRequestActivationRangesInner) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


