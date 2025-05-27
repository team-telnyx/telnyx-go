# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpEntityId** | **string** | The Entity ID for the identity provider (IdP). | 
**IdpSsoTargetUrl** | **string** | The SSO target url for the identity provider (IdP). | 
**IdpCertFingerprint** | **string** | The certificate fingerprint for the identity provider (IdP) | 
**IdpCertFingerprintAlgorithm** | Pointer to **string** | The algorithm used to generate the identity provider&#39;s (IdP) certificate fingerprint | [optional] [default to "sha1"]

## Methods

### NewSettings

`func NewSettings(idpEntityId string, idpSsoTargetUrl string, idpCertFingerprint string, ) *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpEntityId

`func (o *Settings) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *Settings) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *Settings) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.


### GetIdpSsoTargetUrl

`func (o *Settings) GetIdpSsoTargetUrl() string`

GetIdpSsoTargetUrl returns the IdpSsoTargetUrl field if non-nil, zero value otherwise.

### GetIdpSsoTargetUrlOk

`func (o *Settings) GetIdpSsoTargetUrlOk() (*string, bool)`

GetIdpSsoTargetUrlOk returns a tuple with the IdpSsoTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoTargetUrl

`func (o *Settings) SetIdpSsoTargetUrl(v string)`

SetIdpSsoTargetUrl sets IdpSsoTargetUrl field to given value.


### GetIdpCertFingerprint

`func (o *Settings) GetIdpCertFingerprint() string`

GetIdpCertFingerprint returns the IdpCertFingerprint field if non-nil, zero value otherwise.

### GetIdpCertFingerprintOk

`func (o *Settings) GetIdpCertFingerprintOk() (*string, bool)`

GetIdpCertFingerprintOk returns a tuple with the IdpCertFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertFingerprint

`func (o *Settings) SetIdpCertFingerprint(v string)`

SetIdpCertFingerprint sets IdpCertFingerprint field to given value.


### GetIdpCertFingerprintAlgorithm

`func (o *Settings) GetIdpCertFingerprintAlgorithm() string`

GetIdpCertFingerprintAlgorithm returns the IdpCertFingerprintAlgorithm field if non-nil, zero value otherwise.

### GetIdpCertFingerprintAlgorithmOk

`func (o *Settings) GetIdpCertFingerprintAlgorithmOk() (*string, bool)`

GetIdpCertFingerprintAlgorithmOk returns a tuple with the IdpCertFingerprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertFingerprintAlgorithm

`func (o *Settings) SetIdpCertFingerprintAlgorithm(v string)`

SetIdpCertFingerprintAlgorithm sets IdpCertFingerprintAlgorithm field to given value.

### HasIdpCertFingerprintAlgorithm

`func (o *Settings) HasIdpCertFingerprintAlgorithm() bool`

HasIdpCertFingerprintAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


