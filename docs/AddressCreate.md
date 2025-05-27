# AddressCreate

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
**AddressBook** | Pointer to **bool** | Indicates whether or not the address should be considered part of your list of addresses that appear for regular use. | [optional] [default to true]
**ValidateAddress** | Pointer to **bool** | Indicates whether or not the address should be validated for emergency use upon creation or not. This should be left with the default value of &#x60;true&#x60; unless you have used the &#x60;/addresses/actions/validate&#x60; endpoint to validate the address separately prior to creation. If an address is not validated for emergency use upon creation and it is not valid, it will not be able to be used for emergency services. | [optional] [default to true]

## Methods

### NewAddressCreate

`func NewAddressCreate(firstName string, lastName string, businessName string, streetAddress string, locality string, countryCode string, ) *AddressCreate`

NewAddressCreate instantiates a new AddressCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressCreateWithDefaults

`func NewAddressCreateWithDefaults() *AddressCreate`

NewAddressCreateWithDefaults instantiates a new AddressCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerReference

`func (o *AddressCreate) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *AddressCreate) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *AddressCreate) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *AddressCreate) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetFirstName

`func (o *AddressCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddressCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddressCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AddressCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddressCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddressCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetBusinessName

`func (o *AddressCreate) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *AddressCreate) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *AddressCreate) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetPhoneNumber

`func (o *AddressCreate) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AddressCreate) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AddressCreate) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AddressCreate) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *AddressCreate) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *AddressCreate) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *AddressCreate) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### GetExtendedAddress

`func (o *AddressCreate) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *AddressCreate) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *AddressCreate) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *AddressCreate) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *AddressCreate) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *AddressCreate) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *AddressCreate) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetAdministrativeArea

`func (o *AddressCreate) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *AddressCreate) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *AddressCreate) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *AddressCreate) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetNeighborhood

`func (o *AddressCreate) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *AddressCreate) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *AddressCreate) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *AddressCreate) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetBorough

`func (o *AddressCreate) GetBorough() string`

GetBorough returns the Borough field if non-nil, zero value otherwise.

### GetBoroughOk

`func (o *AddressCreate) GetBoroughOk() (*string, bool)`

GetBoroughOk returns a tuple with the Borough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorough

`func (o *AddressCreate) SetBorough(v string)`

SetBorough sets Borough field to given value.

### HasBorough

`func (o *AddressCreate) HasBorough() bool`

HasBorough returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressCreate) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressCreate) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressCreate) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressCreate) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressCreate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressCreate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressCreate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetAddressBook

`func (o *AddressCreate) GetAddressBook() bool`

GetAddressBook returns the AddressBook field if non-nil, zero value otherwise.

### GetAddressBookOk

`func (o *AddressCreate) GetAddressBookOk() (*bool, bool)`

GetAddressBookOk returns a tuple with the AddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBook

`func (o *AddressCreate) SetAddressBook(v bool)`

SetAddressBook sets AddressBook field to given value.

### HasAddressBook

`func (o *AddressCreate) HasAddressBook() bool`

HasAddressBook returns a boolean if a field has been set.

### GetValidateAddress

`func (o *AddressCreate) GetValidateAddress() bool`

GetValidateAddress returns the ValidateAddress field if non-nil, zero value otherwise.

### GetValidateAddressOk

`func (o *AddressCreate) GetValidateAddressOk() (*bool, bool)`

GetValidateAddressOk returns a tuple with the ValidateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAddress

`func (o *AddressCreate) SetValidateAddress(v bool)`

SetValidateAddress sets ValidateAddress field to given value.

### HasValidateAddress

`func (o *AddressCreate) HasValidateAddress() bool`

HasValidateAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


