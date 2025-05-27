# SIMCardOrderOrderAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the address for the order. | [optional] 
**FirstName** | Pointer to **string** | The first name of the shipping recipient. | [optional] 
**LastName** | Pointer to **string** | The last name of the shipping recipient. | [optional] 
**BusinessName** | Pointer to **string** | The name of the business where the address is located. | [optional] 
**StreetAddress** | Pointer to **string** | The name of the street where the address is located. | [optional] 
**ExtendedAddress** | Pointer to **string** | Supplemental field for address information. | [optional] 
**Locality** | Pointer to **string** | The name of the city where the address is located. | [optional] 
**AdministrativeArea** | Pointer to **string** | State or province where the address is located. | [optional] 
**CountryCode** | Pointer to **string** | The mobile operator two-character (ISO 3166-1 alpha-2) origin country code. | [optional] 
**PostalCode** | Pointer to **string** | Postal code for the address. | [optional] 

## Methods

### NewSIMCardOrderOrderAddress

`func NewSIMCardOrderOrderAddress() *SIMCardOrderOrderAddress`

NewSIMCardOrderOrderAddress instantiates a new SIMCardOrderOrderAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardOrderOrderAddressWithDefaults

`func NewSIMCardOrderOrderAddressWithDefaults() *SIMCardOrderOrderAddress`

NewSIMCardOrderOrderAddressWithDefaults instantiates a new SIMCardOrderOrderAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCardOrderOrderAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCardOrderOrderAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCardOrderOrderAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCardOrderOrderAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *SIMCardOrderOrderAddress) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SIMCardOrderOrderAddress) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SIMCardOrderOrderAddress) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SIMCardOrderOrderAddress) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SIMCardOrderOrderAddress) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SIMCardOrderOrderAddress) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SIMCardOrderOrderAddress) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SIMCardOrderOrderAddress) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBusinessName

`func (o *SIMCardOrderOrderAddress) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *SIMCardOrderOrderAddress) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *SIMCardOrderOrderAddress) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *SIMCardOrderOrderAddress) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetStreetAddress

`func (o *SIMCardOrderOrderAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *SIMCardOrderOrderAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *SIMCardOrderOrderAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *SIMCardOrderOrderAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetExtendedAddress

`func (o *SIMCardOrderOrderAddress) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *SIMCardOrderOrderAddress) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *SIMCardOrderOrderAddress) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *SIMCardOrderOrderAddress) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *SIMCardOrderOrderAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *SIMCardOrderOrderAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *SIMCardOrderOrderAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *SIMCardOrderOrderAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *SIMCardOrderOrderAddress) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *SIMCardOrderOrderAddress) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *SIMCardOrderOrderAddress) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *SIMCardOrderOrderAddress) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetCountryCode

`func (o *SIMCardOrderOrderAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SIMCardOrderOrderAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SIMCardOrderOrderAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SIMCardOrderOrderAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPostalCode

`func (o *SIMCardOrderOrderAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SIMCardOrderOrderAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SIMCardOrderOrderAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SIMCardOrderOrderAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


