# AuthenticationProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the authentication provider. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Name** | Pointer to **string** | The name associated with the authentication provider. | [optional] 
**ShortName** | Pointer to **string** | The short name associated with the authentication provider. This must be unique and URL-friendly, as it&#39;s going to be part of the login URL. | [optional] 
**OrganizationId** | Pointer to **string** | The id from the Organization the authentication provider belongs to. | [optional] 
**Active** | Pointer to **bool** | The active status of the authentication provider | [optional] [default to true]
**Settings** | Pointer to [**AuthenticationProviderSettings**](AuthenticationProviderSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewAuthenticationProvider

`func NewAuthenticationProvider() *AuthenticationProvider`

NewAuthenticationProvider instantiates a new AuthenticationProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationProviderWithDefaults

`func NewAuthenticationProviderWithDefaults() *AuthenticationProvider`

NewAuthenticationProviderWithDefaults instantiates a new AuthenticationProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *AuthenticationProvider) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AuthenticationProvider) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AuthenticationProvider) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AuthenticationProvider) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *AuthenticationProvider) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *AuthenticationProvider) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *AuthenticationProvider) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *AuthenticationProvider) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AuthenticationProvider) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AuthenticationProvider) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AuthenticationProvider) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AuthenticationProvider) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetActive

`func (o *AuthenticationProvider) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthenticationProvider) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthenticationProvider) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthenticationProvider) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSettings

`func (o *AuthenticationProvider) GetSettings() AuthenticationProviderSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AuthenticationProvider) GetSettingsOk() (*AuthenticationProviderSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AuthenticationProvider) SetSettings(v AuthenticationProviderSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AuthenticationProvider) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthenticationProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthenticationProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthenticationProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthenticationProvider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthenticationProvider) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthenticationProvider) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthenticationProvider) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthenticationProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


