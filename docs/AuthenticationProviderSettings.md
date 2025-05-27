# AuthenticationProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssertionConsumerServiceUrl** | Pointer to **string** | The Assertion Consumer Service URL for the service provider (Telnyx). | [optional] 
**ServiceProviderEntityId** | Pointer to **string** | The Entity ID for the service provider (Telnyx). | [optional] 
**IdpEntityId** | Pointer to **string** | The Entity ID for the identity provider (IdP). | [optional] 
**IdpSsoTargetUrl** | Pointer to **string** | The SSO target url for the identity provider (IdP). | [optional] 
**IdpCertFingerprint** | Pointer to **string** | The certificate fingerprint for the identity provider (IdP) | [optional] 
**IdpCertFingerprintAlgorithm** | Pointer to **string** | The algorithm used to generate the identity provider&#39;s (IdP) certificate fingerprint | [optional] [default to "sha1"]
**NameIdentifierFormat** | Pointer to **string** | The name identifier format associated with the authentication provider. This must be the same for both the Identity Provider (IdP) and the service provider (Telnyx). | [optional] 

## Methods

### NewAuthenticationProviderSettings

`func NewAuthenticationProviderSettings() *AuthenticationProviderSettings`

NewAuthenticationProviderSettings instantiates a new AuthenticationProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderSettingsWithDefaults

`func NewAuthenticationProviderSettingsWithDefaults() *AuthenticationProviderSettings`

NewAuthenticationProviderSettingsWithDefaults instantiates a new AuthenticationProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertionConsumerServiceUrl

`func (o *AuthenticationProviderSettings) GetAssertionConsumerServiceUrl() string`

GetAssertionConsumerServiceUrl returns the AssertionConsumerServiceUrl field if non-nil, zero value otherwise.

### GetAssertionConsumerServiceUrlOk

`func (o *AuthenticationProviderSettings) GetAssertionConsumerServiceUrlOk() (*string, bool)`

GetAssertionConsumerServiceUrlOk returns a tuple with the AssertionConsumerServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerServiceUrl

`func (o *AuthenticationProviderSettings) SetAssertionConsumerServiceUrl(v string)`

SetAssertionConsumerServiceUrl sets AssertionConsumerServiceUrl field to given value.

### HasAssertionConsumerServiceUrl

`func (o *AuthenticationProviderSettings) HasAssertionConsumerServiceUrl() bool`

HasAssertionConsumerServiceUrl returns a boolean if a field has been set.

### GetServiceProviderEntityId

`func (o *AuthenticationProviderSettings) GetServiceProviderEntityId() string`

GetServiceProviderEntityId returns the ServiceProviderEntityId field if non-nil, zero value otherwise.

### GetServiceProviderEntityIdOk

`func (o *AuthenticationProviderSettings) GetServiceProviderEntityIdOk() (*string, bool)`

GetServiceProviderEntityIdOk returns a tuple with the ServiceProviderEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderEntityId

`func (o *AuthenticationProviderSettings) SetServiceProviderEntityId(v string)`

SetServiceProviderEntityId sets ServiceProviderEntityId field to given value.

### HasServiceProviderEntityId

`func (o *AuthenticationProviderSettings) HasServiceProviderEntityId() bool`

HasServiceProviderEntityId returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *AuthenticationProviderSettings) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *AuthenticationProviderSettings) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *AuthenticationProviderSettings) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *AuthenticationProviderSettings) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetIdpSsoTargetUrl

`func (o *AuthenticationProviderSettings) GetIdpSsoTargetUrl() string`

GetIdpSsoTargetUrl returns the IdpSsoTargetUrl field if non-nil, zero value otherwise.

### GetIdpSsoTargetUrlOk

`func (o *AuthenticationProviderSettings) GetIdpSsoTargetUrlOk() (*string, bool)`

GetIdpSsoTargetUrlOk returns a tuple with the IdpSsoTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoTargetUrl

`func (o *AuthenticationProviderSettings) SetIdpSsoTargetUrl(v string)`

SetIdpSsoTargetUrl sets IdpSsoTargetUrl field to given value.

### HasIdpSsoTargetUrl

`func (o *AuthenticationProviderSettings) HasIdpSsoTargetUrl() bool`

HasIdpSsoTargetUrl returns a boolean if a field has been set.

### GetIdpCertFingerprint

`func (o *AuthenticationProviderSettings) GetIdpCertFingerprint() string`

GetIdpCertFingerprint returns the IdpCertFingerprint field if non-nil, zero value otherwise.

### GetIdpCertFingerprintOk

`func (o *AuthenticationProviderSettings) GetIdpCertFingerprintOk() (*string, bool)`

GetIdpCertFingerprintOk returns a tuple with the IdpCertFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertFingerprint

`func (o *AuthenticationProviderSettings) SetIdpCertFingerprint(v string)`

SetIdpCertFingerprint sets IdpCertFingerprint field to given value.

### HasIdpCertFingerprint

`func (o *AuthenticationProviderSettings) HasIdpCertFingerprint() bool`

HasIdpCertFingerprint returns a boolean if a field has been set.

### GetIdpCertFingerprintAlgorithm

`func (o *AuthenticationProviderSettings) GetIdpCertFingerprintAlgorithm() string`

GetIdpCertFingerprintAlgorithm returns the IdpCertFingerprintAlgorithm field if non-nil, zero value otherwise.

### GetIdpCertFingerprintAlgorithmOk

`func (o *AuthenticationProviderSettings) GetIdpCertFingerprintAlgorithmOk() (*string, bool)`

GetIdpCertFingerprintAlgorithmOk returns a tuple with the IdpCertFingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertFingerprintAlgorithm

`func (o *AuthenticationProviderSettings) SetIdpCertFingerprintAlgorithm(v string)`

SetIdpCertFingerprintAlgorithm sets IdpCertFingerprintAlgorithm field to given value.

### HasIdpCertFingerprintAlgorithm

`func (o *AuthenticationProviderSettings) HasIdpCertFingerprintAlgorithm() bool`

HasIdpCertFingerprintAlgorithm returns a boolean if a field has been set.

### GetNameIdentifierFormat

`func (o *AuthenticationProviderSettings) GetNameIdentifierFormat() string`

GetNameIdentifierFormat returns the NameIdentifierFormat field if non-nil, zero value otherwise.

### GetNameIdentifierFormatOk

`func (o *AuthenticationProviderSettings) GetNameIdentifierFormatOk() (*string, bool)`

GetNameIdentifierFormatOk returns a tuple with the NameIdentifierFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdentifierFormat

`func (o *AuthenticationProviderSettings) SetNameIdentifierFormat(v string)`

SetNameIdentifierFormat sets NameIdentifierFormat field to given value.

### HasNameIdentifierFormat

`func (o *AuthenticationProviderSettings) HasNameIdentifierFormat() bool`

HasNameIdentifierFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


