# CreateVerifyProfileCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingTemplateId** | Pointer to **string** | The message template identifier selected from /verify_profiles/templates | [optional] 
**AppName** | Pointer to **string** | The name that identifies the application requesting 2fa in the verification message. | [optional] 
**CodeLength** | Pointer to **int32** | The length of the verify code to generate. | [optional] [default to 5]
**WhitelistedDestinations** | Pointer to **[]string** | Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed. | [optional] 
**DefaultVerificationTimeoutSecs** | Pointer to **int32** | For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity. | [optional] [default to 300]

## Methods

### NewCreateVerifyProfileCallRequest

`func NewCreateVerifyProfileCallRequest() *CreateVerifyProfileCallRequest`

NewCreateVerifyProfileCallRequest instantiates a new CreateVerifyProfileCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerifyProfileCallRequestWithDefaults

`func NewCreateVerifyProfileCallRequestWithDefaults() *CreateVerifyProfileCallRequest`

NewCreateVerifyProfileCallRequestWithDefaults instantiates a new CreateVerifyProfileCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingTemplateId

`func (o *CreateVerifyProfileCallRequest) GetMessagingTemplateId() string`

GetMessagingTemplateId returns the MessagingTemplateId field if non-nil, zero value otherwise.

### GetMessagingTemplateIdOk

`func (o *CreateVerifyProfileCallRequest) GetMessagingTemplateIdOk() (*string, bool)`

GetMessagingTemplateIdOk returns a tuple with the MessagingTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingTemplateId

`func (o *CreateVerifyProfileCallRequest) SetMessagingTemplateId(v string)`

SetMessagingTemplateId sets MessagingTemplateId field to given value.

### HasMessagingTemplateId

`func (o *CreateVerifyProfileCallRequest) HasMessagingTemplateId() bool`

HasMessagingTemplateId returns a boolean if a field has been set.

### GetAppName

`func (o *CreateVerifyProfileCallRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CreateVerifyProfileCallRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CreateVerifyProfileCallRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *CreateVerifyProfileCallRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetCodeLength

`func (o *CreateVerifyProfileCallRequest) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *CreateVerifyProfileCallRequest) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *CreateVerifyProfileCallRequest) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.

### HasCodeLength

`func (o *CreateVerifyProfileCallRequest) HasCodeLength() bool`

HasCodeLength returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *CreateVerifyProfileCallRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *CreateVerifyProfileCallRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *CreateVerifyProfileCallRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.

### HasWhitelistedDestinations

`func (o *CreateVerifyProfileCallRequest) HasWhitelistedDestinations() bool`

HasWhitelistedDestinations returns a boolean if a field has been set.

### GetDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileCallRequest) GetDefaultVerificationTimeoutSecs() int32`

GetDefaultVerificationTimeoutSecs returns the DefaultVerificationTimeoutSecs field if non-nil, zero value otherwise.

### GetDefaultVerificationTimeoutSecsOk

`func (o *CreateVerifyProfileCallRequest) GetDefaultVerificationTimeoutSecsOk() (*int32, bool)`

GetDefaultVerificationTimeoutSecsOk returns a tuple with the DefaultVerificationTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileCallRequest) SetDefaultVerificationTimeoutSecs(v int32)`

SetDefaultVerificationTimeoutSecs sets DefaultVerificationTimeoutSecs field to given value.

### HasDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileCallRequest) HasDefaultVerificationTimeoutSecs() bool`

HasDefaultVerificationTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


