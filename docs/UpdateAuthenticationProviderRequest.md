# UpdateAuthenticationProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name associated with the authentication provider. | [optional] 
**ShortName** | Pointer to **string** | The short name associated with the authentication provider. This must be unique and URL-friendly, as it&#39;s going to be part of the login URL. | [optional] 
**Active** | Pointer to **bool** | The active status of the authentication provider | [optional] [default to true]
**Settings** | Pointer to [**Settings**](Settings.md) |  | [optional] 
**SettingsUrl** | Pointer to **string** | The URL for the identity provider metadata file to populate the settings automatically. If the settings attribute is provided, that will be used instead. | [optional] 

## Methods

### NewUpdateAuthenticationProviderRequest

`func NewUpdateAuthenticationProviderRequest() *UpdateAuthenticationProviderRequest`

NewUpdateAuthenticationProviderRequest instantiates a new UpdateAuthenticationProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthenticationProviderRequestWithDefaults

`func NewUpdateAuthenticationProviderRequestWithDefaults() *UpdateAuthenticationProviderRequest`

NewUpdateAuthenticationProviderRequestWithDefaults instantiates a new UpdateAuthenticationProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAuthenticationProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAuthenticationProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAuthenticationProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAuthenticationProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *UpdateAuthenticationProviderRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *UpdateAuthenticationProviderRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *UpdateAuthenticationProviderRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *UpdateAuthenticationProviderRequest) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetActive

`func (o *UpdateAuthenticationProviderRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateAuthenticationProviderRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateAuthenticationProviderRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateAuthenticationProviderRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateAuthenticationProviderRequest) GetSettings() Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateAuthenticationProviderRequest) GetSettingsOk() (*Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateAuthenticationProviderRequest) SetSettings(v Settings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateAuthenticationProviderRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSettingsUrl

`func (o *UpdateAuthenticationProviderRequest) GetSettingsUrl() string`

GetSettingsUrl returns the SettingsUrl field if non-nil, zero value otherwise.

### GetSettingsUrlOk

`func (o *UpdateAuthenticationProviderRequest) GetSettingsUrlOk() (*string, bool)`

GetSettingsUrlOk returns a tuple with the SettingsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsUrl

`func (o *UpdateAuthenticationProviderRequest) SetSettingsUrl(v string)`

SetSettingsUrl sets SettingsUrl field to given value.

### HasSettingsUrl

`func (o *UpdateAuthenticationProviderRequest) HasSettingsUrl() bool`

HasSettingsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


