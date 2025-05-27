# UserAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the user address. | [optional] 
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
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewUserAddress

`func NewUserAddress() *UserAddress`

NewUserAddress instantiates a new UserAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAddressWithDefaults

`func NewUserAddressWithDefaults() *UserAddress`

NewUserAddressWithDefaults instantiates a new UserAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *UserAddress) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *UserAddress) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *UserAddress) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *UserAddress) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCustomerReference

`func (o *UserAddress) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UserAddress) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UserAddress) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UserAddress) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetFirstName

`func (o *UserAddress) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserAddress) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserAddress) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserAddress) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserAddress) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserAddress) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserAddress) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserAddress) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBusinessName

`func (o *UserAddress) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *UserAddress) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *UserAddress) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *UserAddress) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserAddress) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserAddress) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserAddress) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserAddress) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetExtendedAddress

`func (o *UserAddress) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *UserAddress) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *UserAddress) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *UserAddress) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *UserAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *UserAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *UserAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *UserAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *UserAddress) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *UserAddress) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *UserAddress) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *UserAddress) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetNeighborhood

`func (o *UserAddress) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *UserAddress) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *UserAddress) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *UserAddress) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetBorough

`func (o *UserAddress) GetBorough() string`

GetBorough returns the Borough field if non-nil, zero value otherwise.

### GetBoroughOk

`func (o *UserAddress) GetBoroughOk() (*string, bool)`

GetBoroughOk returns a tuple with the Borough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorough

`func (o *UserAddress) SetBorough(v string)`

SetBorough sets Borough field to given value.

### HasBorough

`func (o *UserAddress) HasBorough() bool`

HasBorough returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *UserAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserAddress) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserAddress) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserAddress) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserAddress) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserAddress) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserAddress) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


