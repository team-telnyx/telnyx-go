# CreateVerifyProfileSMSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingTemplateId** | Pointer to **string** | The message template identifier selected from /verify_profiles/templates | [optional] 
**AppName** | Pointer to **string** | The name that identifies the application requesting 2fa in the verification message. | [optional] 
**AlphaSender** | Pointer to **NullableString** | The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID. | [optional] [default to "Telnyx"]
**CodeLength** | Pointer to **int32** | The length of the verify code to generate. | [optional] [default to 5]
**WhitelistedDestinations** | **[]string** | Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to &#x60;[\&quot;*\&quot;]&#x60;, all destinations will be allowed. | 
**DefaultVerificationTimeoutSecs** | Pointer to **int32** | For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity. | [optional] [default to 300]

## Methods

### NewCreateVerifyProfileSMSRequest

`func NewCreateVerifyProfileSMSRequest(whitelistedDestinations []string, ) *CreateVerifyProfileSMSRequest`

NewCreateVerifyProfileSMSRequest instantiates a new CreateVerifyProfileSMSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerifyProfileSMSRequestWithDefaults

`func NewCreateVerifyProfileSMSRequestWithDefaults() *CreateVerifyProfileSMSRequest`

NewCreateVerifyProfileSMSRequestWithDefaults instantiates a new CreateVerifyProfileSMSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingTemplateId

`func (o *CreateVerifyProfileSMSRequest) GetMessagingTemplateId() string`

GetMessagingTemplateId returns the MessagingTemplateId field if non-nil, zero value otherwise.

### GetMessagingTemplateIdOk

`func (o *CreateVerifyProfileSMSRequest) GetMessagingTemplateIdOk() (*string, bool)`

GetMessagingTemplateIdOk returns a tuple with the MessagingTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingTemplateId

`func (o *CreateVerifyProfileSMSRequest) SetMessagingTemplateId(v string)`

SetMessagingTemplateId sets MessagingTemplateId field to given value.

### HasMessagingTemplateId

`func (o *CreateVerifyProfileSMSRequest) HasMessagingTemplateId() bool`

HasMessagingTemplateId returns a boolean if a field has been set.

### GetAppName

`func (o *CreateVerifyProfileSMSRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *CreateVerifyProfileSMSRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *CreateVerifyProfileSMSRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *CreateVerifyProfileSMSRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetAlphaSender

`func (o *CreateVerifyProfileSMSRequest) GetAlphaSender() string`

GetAlphaSender returns the AlphaSender field if non-nil, zero value otherwise.

### GetAlphaSenderOk

`func (o *CreateVerifyProfileSMSRequest) GetAlphaSenderOk() (*string, bool)`

GetAlphaSenderOk returns a tuple with the AlphaSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaSender

`func (o *CreateVerifyProfileSMSRequest) SetAlphaSender(v string)`

SetAlphaSender sets AlphaSender field to given value.

### HasAlphaSender

`func (o *CreateVerifyProfileSMSRequest) HasAlphaSender() bool`

HasAlphaSender returns a boolean if a field has been set.

### SetAlphaSenderNil

`func (o *CreateVerifyProfileSMSRequest) SetAlphaSenderNil(b bool)`

 SetAlphaSenderNil sets the value for AlphaSender to be an explicit nil

### UnsetAlphaSender
`func (o *CreateVerifyProfileSMSRequest) UnsetAlphaSender()`

UnsetAlphaSender ensures that no value is present for AlphaSender, not even an explicit nil
### GetCodeLength

`func (o *CreateVerifyProfileSMSRequest) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *CreateVerifyProfileSMSRequest) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *CreateVerifyProfileSMSRequest) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.

### HasCodeLength

`func (o *CreateVerifyProfileSMSRequest) HasCodeLength() bool`

HasCodeLength returns a boolean if a field has been set.

### GetWhitelistedDestinations

`func (o *CreateVerifyProfileSMSRequest) GetWhitelistedDestinations() []string`

GetWhitelistedDestinations returns the WhitelistedDestinations field if non-nil, zero value otherwise.

### GetWhitelistedDestinationsOk

`func (o *CreateVerifyProfileSMSRequest) GetWhitelistedDestinationsOk() (*[]string, bool)`

GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedDestinations

`func (o *CreateVerifyProfileSMSRequest) SetWhitelistedDestinations(v []string)`

SetWhitelistedDestinations sets WhitelistedDestinations field to given value.


### GetDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileSMSRequest) GetDefaultVerificationTimeoutSecs() int32`

GetDefaultVerificationTimeoutSecs returns the DefaultVerificationTimeoutSecs field if non-nil, zero value otherwise.

### GetDefaultVerificationTimeoutSecsOk

`func (o *CreateVerifyProfileSMSRequest) GetDefaultVerificationTimeoutSecsOk() (*int32, bool)`

GetDefaultVerificationTimeoutSecsOk returns a tuple with the DefaultVerificationTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileSMSRequest) SetDefaultVerificationTimeoutSecs(v int32)`

SetDefaultVerificationTimeoutSecs sets DefaultVerificationTimeoutSecs field to given value.

### HasDefaultVerificationTimeoutSecs

`func (o *CreateVerifyProfileSMSRequest) HasDefaultVerificationTimeoutSecs() bool`

HasDefaultVerificationTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


