# EligibilityNumberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** | Detailed information about the eligibility status. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number in e164 format. | [optional] 
**Eligible** | Pointer to **bool** | Whether the phone number is eligible for hosted messaging. | [optional] 
**EligibleStatus** | Pointer to **string** | The eligibility status of the phone number. | [optional] 

## Methods

### NewEligibilityNumberResponse

`func NewEligibilityNumberResponse() *EligibilityNumberResponse`

NewEligibilityNumberResponse instantiates a new EligibilityNumberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityNumberResponseWithDefaults

`func NewEligibilityNumberResponseWithDefaults() *EligibilityNumberResponse`

NewEligibilityNumberResponseWithDefaults instantiates a new EligibilityNumberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *EligibilityNumberResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EligibilityNumberResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EligibilityNumberResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EligibilityNumberResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *EligibilityNumberResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *EligibilityNumberResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *EligibilityNumberResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *EligibilityNumberResponse) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEligible

`func (o *EligibilityNumberResponse) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *EligibilityNumberResponse) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *EligibilityNumberResponse) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *EligibilityNumberResponse) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetEligibleStatus

`func (o *EligibilityNumberResponse) GetEligibleStatus() string`

GetEligibleStatus returns the EligibleStatus field if non-nil, zero value otherwise.

### GetEligibleStatusOk

`func (o *EligibilityNumberResponse) GetEligibleStatusOk() (*string, bool)`

GetEligibleStatusOk returns a tuple with the EligibleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleStatus

`func (o *EligibilityNumberResponse) SetEligibleStatus(v string)`

SetEligibleStatus sets EligibleStatus field to given value.

### HasEligibleStatus

`func (o *EligibilityNumberResponse) HasEligibleStatus() bool`

HasEligibleStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


