# VerifyProfileCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingTemplateId** | Pointer to **string** | The message template identifier selected from /verify_profiles/templates | [optional] 
**AppName** | Pointer to **string** | The name that identifies the application requesting 2fa in the verification message. | [optional] 
**CodeLength** | Pointer to **int32** | The length of the verify code to generate. | [optional] [default to 5]
**WhitelistedDestinations** | Pointer to **[]string** | Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed. | [optional] 
**DefaultVerificationTimeoutSecs** | Pointer to **int32** | For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity. | [optional] [default to 300]

## Methods

### NewVerifyProfileCallResponse

`func NewVerifyProfileCallResponse() *VerifyProfileCallResponse`

NewVerifyProfileCallResponse instantiates a new VerifyProfileCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyProfileCallResponseWithDefaults

`func NewVerifyProfileCallResponseWithDefaults() *VerifyProfileCallResponse`

NewVerifyProfileCallResponseWithDefaults instantiates a new VerifyProfileCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingTemplateId

`func (o *VerifyProfileCallResponse) GetMessagingTemplateId() string`

GetMessagingTemplateId returns the MessagingTemplateId field if non-nil, zero value otherwise.

### GetMessagingTemplateIdOk

`func (o *VerifyProfileCallResponse) GetMessagingTemplateIdOk() (*string, bool)`

GetMessagingTemplateIdOk returns a tuple with the MessagingTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingTemplateId

`func (o *VerifyProfileCallResponse) SetMessagingTemplateId(v string)`

SetMessagingTemplateId sets MessagingTemplateId field to given value.

### HasMessagingTemplateId

`func (o *VerifyProfileCallResponse) HasMessagingTemplateId() bool`

HasMessagingTemplateId returns a boolean if a field has been set.

### GetAppName

`func (o *VerifyProfileCallResponse) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *VerifyProfileCallResponse) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *VerifyProfileCallResponse) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *VerifyProfileCallResponse) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCodeLength

`func (o *VerifyProfileCallResponse) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *VerifyProfileCallResponse) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *VerifyProfileCallResponse) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.

### HasCodeLength

`func (o *VerifyProfileCallResponse) HasCodeLength() bool`

HasCodeLength returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *VerifyProfileCallResponse) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *VerifyProfileCallResponse) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *VerifyProfileCallResponse) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *VerifyProfileCallResponse) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetDefaultVerificationTimeoutSecs

`func (o *VerifyProfileCallResponse) GetDefaultVerificationTimeoutSecs() int32`

GetDefaultVerificationTimeoutSecs returns the DefaultVerificationTimeoutSecs field if non-nil, zero value otherwise.

### GetDefaultVerificationTimeoutSecsOk

`func (o *VerifyProfileCallResponse) GetDefaultVerificationTimeoutSecsOk() (*int32, bool)`

GetDefaultVerificationTimeoutSecsOk returns a tuple with the DefaultVerificationTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVerificationTimeoutSecs

`func (o *VerifyProfileCallResponse) SetDefaultVerificationTimeoutSecs(v int32)`

SetDefaultVerificationTimeoutSecs sets DefaultVerificationTimeoutSecs field to given value.

### HasDefaultVerificationTimeoutSecs

`func (o *VerifyProfileCallResponse) HasDefaultVerificationTimeoutSecs() bool`

HasDefaultVerificationTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


