# UpdateVerifyProfileFlashcallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WhitelistedDestinations** | Pointer to **[]string** | Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed. | [optional] 
**DefaultVerificationTimeoutSecs** | Pointer to **int32** | For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity. | [optional] [default to 300]

## Methods

### NewUpdateVerifyProfileFlashcallRequest

`func NewUpdateVerifyProfileFlashcallRequest() *UpdateVerifyProfileFlashcallRequest`

NewUpdateVerifyProfileFlashcallRequest instantiates a new UpdateVerifyProfileFlashcallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVerifyProfileFlashcallRequestWithDefaults

`func NewUpdateVerifyProfileFlashcallRequestWithDefaults() *UpdateVerifyProfileFlashcallRequest`

NewUpdateVerifyProfileFlashcallRequestWithDefaults instantiates a new UpdateVerifyProfileFlashcallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhitelistedDestinations

`func (o *UpdateVerifyProfileFlashcallRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *UpdateVerifyProfileFlashcallRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *UpdateVerifyProfileFlashcallRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *UpdateVerifyProfileFlashcallRequest) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetDefaultVerificationTimeoutSecs

`func (o *UpdateVerifyProfileFlashcallRequest) GetDefaultVerificationTimeoutSecs() int32`

GetDefaultVerificationTimeoutSecs returns the DefaultVerificationTimeoutSecs field if non-nil, zero value otherwise.

### GetDefaultVerificationTimeoutSecsOk

`func (o *UpdateVerifyProfileFlashcallRequest) GetDefaultVerificationTimeoutSecsOk() (*int32, bool)`

GetDefaultVerificationTimeoutSecsOk returns a tuple with the DefaultVerificationTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVerificationTimeoutSecs

`func (o *UpdateVerifyProfileFlashcallRequest) SetDefaultVerificationTimeoutSecs(v int32)`

SetDefaultVerificationTimeoutSecs sets DefaultVerificationTimeoutSecs field to given value.

### HasDefaultVerificationTimeoutSecs

`func (o *UpdateVerifyProfileFlashcallRequest) HasDefaultVerificationTimeoutSecs() bool`

HasDefaultVerificationTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


