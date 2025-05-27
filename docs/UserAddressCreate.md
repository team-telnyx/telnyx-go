# UserAddressCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**FirstName** | **string** | The first name associated with the address. An address must have either a first last name or a business name. | 
**LastName** | **string** | The last name associated with the address. An address must have either a first last name or a business name. | 
**BusinessName** | **string** | The business name associated with the address. An address must have either a first last name or a business name. | 
**PhoneNumber** | Pointer to **string** | The phone number associated with the address. | [optional] 
**StreetAddress** | **string** | The primary street address information about the address. | 
**ExtendedAddress** | Pointer to **string** | Additional street address information about the address such as, but not limited to, unit number or apartment number. | [optional] 
**Locality** | **string** | The locality of the address. For US addresses, this corresponds to the city of the address. | 
**AdministrativeArea** | Pointer to **string** | The locality of the address. For US addresses, this corresponds to the state of the address. | [optional] 
**Neighborhood** | Pointer to **string** | The neighborhood of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] 
**Borough** | Pointer to **string** | The borough of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the address. | [optional] 
**CountryCode** | **string** | The two-character (ISO 3166-1 alpha-2) country code of the address. | 
**SkipAddressVerification** | Pointer to **string** | An optional boolean value specifying if verification of the address should be skipped or not. UserAddresses are generally used for shipping addresses, and failure to validate your shipping address will likely result in a failure to deliver SIM cards or other items ordered from Telnyx. Do not use this parameter unless you are sure that the address is correct even though it cannot be validated. If this is set to any value other than true, verification of the address will be attempted, and the user address will not be allowed if verification fails. If verification fails but suggested values are available that might make the address correct, they will be present in the response as well. If this value is set to true, then the verification will not be attempted. Defaults to false (verification will be performed). | [optional] [default to "false"]

## Methods

### NewUserAddressCreate

`func NewUserAddressCreate(firstName string, lastName string, businessName string, streetAddress string, locality string, countryCode string, ) *UserAddressCreate`

NewUserAddressCreate instantiates a new UserAddressCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAddressCreateWithDefaults

`func NewUserAddressCreateWithDefaults() *UserAddressCreate`

NewUserAddressCreateWithDefaults instantiates a new UserAddressCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerReference

`func (o *UserAddressCreate) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UserAddressCreate) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UserAddressCreate) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UserAddressCreate) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetFirstName

`func (o *UserAddressCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserAddressCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserAddressCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UserAddressCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserAddressCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserAddressCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetBusinessName

`func (o *UserAddressCreate) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *UserAddressCreate) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *UserAddressCreate) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetPhoneNumber

`func (o *UserAddressCreate) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserAddressCreate) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserAddressCreate) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserAddressCreate) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserAddressCreate) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserAddressCreate) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserAddressCreate) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### GetExtendedAddress

`func (o *UserAddressCreate) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *UserAddressCreate) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *UserAddressCreate) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *UserAddressCreate) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *UserAddressCreate) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *UserAddressCreate) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *UserAddressCreate) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetAdministrativeArea

`func (o *UserAddressCreate) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *UserAddressCreate) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *UserAddressCreate) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *UserAddressCreate) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetNeighborhood

`func (o *UserAddressCreate) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *UserAddressCreate) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *UserAddressCreate) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *UserAddressCreate) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetBorough

`func (o *UserAddressCreate) GetBorough() string`

GetBorough returns the Borough field if non-nil, zero value otherwise.

### GetBoroughOk

`func (o *UserAddressCreate) GetBoroughOk() (*string, bool)`

GetBoroughOk returns a tuple with the Borough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorough

`func (o *UserAddressCreate) SetBorough(v string)`

SetBorough sets Borough field to given value.

### HasBorough

`func (o *UserAddressCreate) HasBorough() bool`

HasBorough returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserAddressCreate) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserAddressCreate) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserAddressCreate) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserAddressCreate) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *UserAddressCreate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserAddressCreate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserAddressCreate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetSkipAddressVerification

`func (o *UserAddressCreate) GetSkipAddressVerification() string`

GetSkipAddressVerification returns the SkipAddressVerification field if non-nil, zero value otherwise.

### GetSkipAddressVerificationOk

`func (o *UserAddressCreate) GetSkipAddressVerificationOk() (*string, bool)`

GetSkipAddressVerificationOk returns a tuple with the SkipAddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipAddressVerification

`func (o *UserAddressCreate) SetSkipAddressVerification(v string)`

SetSkipAddressVerification sets SkipAddressVerification field to given value.

### HasSkipAddressVerification

`func (o *UserAddressCreate) HasSkipAddressVerification() bool`

HasSkipAddressVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


