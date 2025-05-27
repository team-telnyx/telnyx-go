# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the address. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**FirstName** | Pointer to **string** | The first name associated with the address. An address must have either a first last name or a business name. | [optional] 
**LastName** | Pointer to **string** | The last name associated with the address. An address must have either a first last name or a business name. | [optional] 
**BusinessName** | Pointer to **string** | The business name associated with the address. An address must have either a first last name or a business name. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number associated with the address. | [optional] 
**StreetAddress** | Pointer to **string** | The primary street address information about the address. | [optional] 
**ExtendedAddress** | Pointer to **string** | Additional street address information about the address such as, but not limited to, unit number or apartment number. | [optional] 
**Locality** | Pointer to **string** | The locality of the address. For US addresses, this corresponds to the city of the address. | [optional] 
**AdministrativeArea** | Pointer to **string** | The locality of the address. For US addresses, this corresponds to the state of the address. | [optional] 
**Neighborhood** | Pointer to **string** | The neighborhood of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] 
**Borough** | Pointer to **string** | The borough of the address. This field is not used for addresses in the US but is used for some international addresses. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the address. | [optional] 
**CountryCode** | Pointer to **string** | The two-character (ISO 3166-1 alpha-2) country code of the address. | [optional] 
**AddressBook** | Pointer to **bool** | Indicates whether or not the address should be considered part of your list of addresses that appear for regular use. | [optional] [default to true]
**ValidateAddress** | Pointer to **bool** | Indicates whether or not the address should be validated for emergency use upon creation or not. This should be left with the default value of &#x60;true&#x60; unless you have used the &#x60;/addresses/actions/validate&#x60; endpoint to validate the address separately prior to creation. If an address is not validated for emergency use upon creation and it is not valid, it will not be able to be used for emergency services. | [optional] [default to true]
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Address) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *Address) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Address) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Address) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Address) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCustomerReference

`func (o *Address) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *Address) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *Address) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *Address) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetFirstName

`func (o *Address) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Address) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Address) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Address) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Address) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Address) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Address) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Address) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBusinessName

`func (o *Address) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *Address) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *Address) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *Address) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Address) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Address) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Address) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Address) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *Address) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *Address) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *Address) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *Address) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetExtendedAddress

`func (o *Address) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *Address) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *Address) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *Address) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *Address) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *Address) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *Address) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *Address) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *Address) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *Address) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *Address) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *Address) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetNeighborhood

`func (o *Address) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *Address) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *Address) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *Address) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetBorough

`func (o *Address) GetBorough() string`

GetBorough returns the Borough field if non-nil, zero value otherwise.

### GetBoroughOk

`func (o *Address) GetBoroughOk() (*string, bool)`

GetBoroughOk returns a tuple with the Borough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorough

`func (o *Address) SetBorough(v string)`

SetBorough sets Borough field to given value.

### HasBorough

`func (o *Address) HasBorough() bool`

HasBorough returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *Address) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Address) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Address) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Address) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetAddressBook

`func (o *Address) GetAddressBook() bool`

GetAddressBook returns the AddressBook field if non-nil, zero value otherwise.

### GetAddressBookOk

`func (o *Address) GetAddressBookOk() (*bool, bool)`

GetAddressBookOk returns a tuple with the AddressBook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBook

`func (o *Address) SetAddressBook(v bool)`

SetAddressBook sets AddressBook field to given value.

### HasAddressBook

`func (o *Address) HasAddressBook() bool`

HasAddressBook returns a boolean if a field has been set.

### GetValidateAddress

`func (o *Address) GetValidateAddress() bool`

GetValidateAddress returns the ValidateAddress field if non-nil, zero value otherwise.

### GetValidateAddressOk

`func (o *Address) GetValidateAddressOk() (*bool, bool)`

GetValidateAddressOk returns a tuple with the ValidateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAddress

`func (o *Address) SetValidateAddress(v bool)`

SetValidateAddress sets ValidateAddress field to given value.

### HasValidateAddress

`func (o *Address) HasValidateAddress() bool`

HasValidateAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Address) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Address) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Address) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Address) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Address) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Address) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Address) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Address) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


