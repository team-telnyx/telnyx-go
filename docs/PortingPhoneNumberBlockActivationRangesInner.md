# PortingPhoneNumberBlockActivationRangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | Pointer to **string** | Specifies the start of the activation range. Must be greater or equal the start of the phone number range. | [optional] 
**EndAt** | Pointer to **string** | Specifies the end of the activation range. It must be no more than the end of the phone number range. | [optional] 

## Methods

### NewPortingPhoneNumberBlockActivationRangesInner

`func NewPortingPhoneNumberBlockActivationRangesInner() *PortingPhoneNumberBlockActivationRangesInner`

NewPortingPhoneNumberBlockActivationRangesInner instantiates a new PortingPhoneNumberBlockActivationRangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingPhoneNumberBlockActivationRangesInnerWithDefaults

`func NewPortingPhoneNumberBlockActivationRangesInnerWithDefaults() *PortingPhoneNumberBlockActivationRangesInner`

NewPortingPhoneNumberBlockActivationRangesInnerWithDefaults instantiates a new PortingPhoneNumberBlockActivationRangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PortingPhoneNumberBlockActivationRangesInner) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetEndAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *PortingPhoneNumberBlockActivationRangesInner) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *PortingPhoneNumberBlockActivationRangesInner) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


