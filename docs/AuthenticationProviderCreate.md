# AuthenticationProviderCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name associated with the authentication provider. | 
**ShortName** | **string** | The short name associated with the authentication provider. This must be unique and URL-friendly, as it&#39;s going to be part of the login URL. | 
**Active** | Pointer to **bool** | The active status of the authentication provider | [optional] [default to true]
**Settings** | [**Settings**](Settings.md) |  | 
**SettingsUrl** | Pointer to **string** | The URL for the identity provider metadata file to populate the settings automatically. If the settings attribute is provided, that will be used instead. | [optional] 

## Methods

### NewAuthenticationProviderCreate

`func NewAuthenticationProviderCreate(name string, shortName string, settings Settings, ) *AuthenticationProviderCreate`

NewAuthenticationProviderCreate instantiates a new AuthenticationProviderCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderCreateWithDefaults

`func NewAuthenticationProviderCreateWithDefaults() *AuthenticationProviderCreate`

NewAuthenticationProviderCreateWithDefaults instantiates a new AuthenticationProviderCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticationProviderCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationProviderCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationProviderCreate) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *AuthenticationProviderCreate) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *AuthenticationProviderCreate) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *AuthenticationProviderCreate) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetActive

`func (o *AuthenticationProviderCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthenticationProviderCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthenticationProviderCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthenticationProviderCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSettings

`func (o *AuthenticationProviderCreate) GetSettings() Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AuthenticationProviderCreate) GetSettingsOk() (*Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AuthenticationProviderCreate) SetSettings(v Settings)`

SetSettings sets Settings field to given value.


### GetSettingsUrl

`func (o *AuthenticationProviderCreate) GetSettingsUrl() string`

GetSettingsUrl returns the SettingsUrl field if non-nil, zero value otherwise.

### GetSettingsUrlOk

`func (o *AuthenticationProviderCreate) GetSettingsUrlOk() (*string, bool)`

GetSettingsUrlOk returns a tuple with the SettingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsUrl

`func (o *AuthenticationProviderCreate) SetSettingsUrl(v string)`

SetSettingsUrl sets SettingsUrl field to given value.

### HasSettingsUrl

`func (o *AuthenticationProviderCreate) HasSettingsUrl() bool`

HasSettingsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


