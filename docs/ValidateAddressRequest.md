# ValidateAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreetAddress** | **string** | The primary street address information about the address. | 
**ExtendedAddress** | Pointer to **string** | Additional street address information about the address such as, but not limited to, unit number or apartment number. | [optional] 
**Locality** | Pointer to **string** | The locality of the address. For US addresses, this corresponds to the city of the address. | [optional] 
**AdministrativeArea** | Pointer to **string** | The locality of the address. For US addresses, this corresponds to the state of the address. | [optional] 
**PostalCode** | **string** | The postal code of the address. | 
**CountryCode** | **string** | The two-character (ISO 3166-1 alpha-2) country code of the address. | 

## Methods

### NewValidateAddressRequest

`func NewValidateAddressRequest(streetAddress string, postalCode string, countryCode string, ) *ValidateAddressRequest`

NewValidateAddressRequest instantiates a new ValidateAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddressRequestWithDefaults

`func NewValidateAddressRequestWithDefaults() *ValidateAddressRequest`

NewValidateAddressRequestWithDefaults instantiates a new ValidateAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreetAddress

`func (o *ValidateAddressRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ValidateAddressRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ValidateAddressRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### GetExtendedAddress

`func (o *ValidateAddressRequest) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *ValidateAddressRequest) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *ValidateAddressRequest) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *ValidateAddressRequest) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *ValidateAddressRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ValidateAddressRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ValidateAddressRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ValidateAddressRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *ValidateAddressRequest) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *ValidateAddressRequest) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *ValidateAddressRequest) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *ValidateAddressRequest) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetPostalCode

`func (o *ValidateAddressRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ValidateAddressRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ValidateAddressRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountryCode

`func (o *ValidateAddressRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ValidateAddressRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ValidateAddressRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


