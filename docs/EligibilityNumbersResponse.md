# EligibilityNumbersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to [**[]EligibilityNumberResponse**](EligibilityNumberResponse.md) | List of phone numbers with their eligibility status. | [optional] 

## Methods

### NewEligibilityNumbersResponse

`func NewEligibilityNumbersResponse() *EligibilityNumbersResponse`

NewEligibilityNumbersResponse instantiates a new EligibilityNumbersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEligibilityNumbersResponseWithDefaults

`func NewEligibilityNumbersResponseWithDefaults() *EligibilityNumbersResponse`

NewEligibilityNumbersResponseWithDefaults instantiates a new EligibilityNumbersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *EligibilityNumbersResponse) GetPhoneNumbers() []EligibilityNumberResponse`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *EligibilityNumbersResponse) GetPhoneNumbersOk() (*[]EligibilityNumberResponse, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *EligibilityNumbersResponse) SetPhoneNumbers(v []EligibilityNumberResponse)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *EligibilityNumbersResponse) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


