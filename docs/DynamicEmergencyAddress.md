# DynamicEmergencyAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**SipGeolocationId** | Pointer to **string** | Unique location reference string to be used in SIP INVITE from / p-asserted headers. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of dynamic emergency address | [optional] [readonly] 
**HouseNumber** | **string** |  | 
**HouseSuffix** | Pointer to **string** |  | [optional] 
**StreetPreDirectional** | Pointer to **string** |  | [optional] 
**StreetName** | **string** |  | 
**StreetSuffix** | Pointer to **string** |  | [optional] 
**StreetPostDirectional** | Pointer to **string** |  | [optional] 
**ExtendedAddress** | Pointer to **string** |  | [optional] 
**Locality** | **string** |  | 
**AdministrativeArea** | **string** |  | 
**PostalCode** | **string** |  | 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was created | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was last updated | [optional] [readonly] 

## Methods

### NewDynamicEmergencyAddress

`func NewDynamicEmergencyAddress(houseNumber string, streetName string, locality string, administrativeArea string, postalCode string, ) *DynamicEmergencyAddress`

NewDynamicEmergencyAddress instantiates a new DynamicEmergencyAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicEmergencyAddressWithDefaults

`func NewDynamicEmergencyAddressWithDefaults() *DynamicEmergencyAddress`

NewDynamicEmergencyAddressWithDefaults instantiates a new DynamicEmergencyAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DynamicEmergencyAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DynamicEmergencyAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DynamicEmergencyAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DynamicEmergencyAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *DynamicEmergencyAddress) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DynamicEmergencyAddress) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DynamicEmergencyAddress) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DynamicEmergencyAddress) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSipGeolocationId

`func (o *DynamicEmergencyAddress) GetSipGeolocationId() string`

GetSipGeolocationId returns the SipGeolocationId field if non-nil, zero value otherwise.

### GetSipGeolocationIdOk

`func (o *DynamicEmergencyAddress) GetSipGeolocationIdOk() (*string, bool)`

GetSipGeolocationIdOk returns a tuple with the SipGeolocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipGeolocationId

`func (o *DynamicEmergencyAddress) SetSipGeolocationId(v string)`

SetSipGeolocationId sets SipGeolocationId field to given value.

### HasSipGeolocationId

`func (o *DynamicEmergencyAddress) HasSipGeolocationId() bool`

HasSipGeolocationId returns a boolean if a field has been set.

### GetStatus

`func (o *DynamicEmergencyAddress) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DynamicEmergencyAddress) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DynamicEmergencyAddress) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DynamicEmergencyAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHouseNumber

`func (o *DynamicEmergencyAddress) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *DynamicEmergencyAddress) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *DynamicEmergencyAddress) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.


### GetHouseSuffix

`func (o *DynamicEmergencyAddress) GetHouseSuffix() string`

GetHouseSuffix returns the HouseSuffix field if non-nil, zero value otherwise.

### GetHouseSuffixOk

`func (o *DynamicEmergencyAddress) GetHouseSuffixOk() (*string, bool)`

GetHouseSuffixOk returns a tuple with the HouseSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseSuffix

`func (o *DynamicEmergencyAddress) SetHouseSuffix(v string)`

SetHouseSuffix sets HouseSuffix field to given value.

### HasHouseSuffix

`func (o *DynamicEmergencyAddress) HasHouseSuffix() bool`

HasHouseSuffix returns a boolean if a field has been set.

### GetStreetPreDirectional

`func (o *DynamicEmergencyAddress) GetStreetPreDirectional() string`

GetStreetPreDirectional returns the StreetPreDirectional field if non-nil, zero value otherwise.

### GetStreetPreDirectionalOk

`func (o *DynamicEmergencyAddress) GetStreetPreDirectionalOk() (*string, bool)`

GetStreetPreDirectionalOk returns a tuple with the StreetPreDirectional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetPreDirectional

`func (o *DynamicEmergencyAddress) SetStreetPreDirectional(v string)`

SetStreetPreDirectional sets StreetPreDirectional field to given value.

### HasStreetPreDirectional

`func (o *DynamicEmergencyAddress) HasStreetPreDirectional() bool`

HasStreetPreDirectional returns a boolean if a field has been set.

### GetStreetName

`func (o *DynamicEmergencyAddress) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *DynamicEmergencyAddress) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *DynamicEmergencyAddress) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.


### GetStreetSuffix

`func (o *DynamicEmergencyAddress) GetStreetSuffix() string`

GetStreetSuffix returns the StreetSuffix field if non-nil, zero value otherwise.

### GetStreetSuffixOk

`func (o *DynamicEmergencyAddress) GetStreetSuffixOk() (*string, bool)`

GetStreetSuffixOk returns a tuple with the StreetSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetSuffix

`func (o *DynamicEmergencyAddress) SetStreetSuffix(v string)`

SetStreetSuffix sets StreetSuffix field to given value.

### HasStreetSuffix

`func (o *DynamicEmergencyAddress) HasStreetSuffix() bool`

HasStreetSuffix returns a boolean if a field has been set.

### GetStreetPostDirectional

`func (o *DynamicEmergencyAddress) GetStreetPostDirectional() string`

GetStreetPostDirectional returns the StreetPostDirectional field if non-nil, zero value otherwise.

### GetStreetPostDirectionalOk

`func (o *DynamicEmergencyAddress) GetStreetPostDirectionalOk() (*string, bool)`

GetStreetPostDirectionalOk returns a tuple with the StreetPostDirectional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetPostDirectional

`func (o *DynamicEmergencyAddress) SetStreetPostDirectional(v string)`

SetStreetPostDirectional sets StreetPostDirectional field to given value.

### HasStreetPostDirectional

`func (o *DynamicEmergencyAddress) HasStreetPostDirectional() bool`

HasStreetPostDirectional returns a boolean if a field has been set.

### GetExtendedAddress

`func (o *DynamicEmergencyAddress) GetExtendedAddress() string`

GetExtendedAddress returns the ExtendedAddress field if non-nil, zero value otherwise.

### GetExtendedAddressOk

`func (o *DynamicEmergencyAddress) GetExtendedAddressOk() (*string, bool)`

GetExtendedAddressOk returns a tuple with the ExtendedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAddress

`func (o *DynamicEmergencyAddress) SetExtendedAddress(v string)`

SetExtendedAddress sets ExtendedAddress field to given value.

### HasExtendedAddress

`func (o *DynamicEmergencyAddress) HasExtendedAddress() bool`

HasExtendedAddress returns a boolean if a field has been set.

### GetLocality

`func (o *DynamicEmergencyAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *DynamicEmergencyAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *DynamicEmergencyAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetAdministrativeArea

`func (o *DynamicEmergencyAddress) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *DynamicEmergencyAddress) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *DynamicEmergencyAddress) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.


### GetPostalCode

`func (o *DynamicEmergencyAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *DynamicEmergencyAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *DynamicEmergencyAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountryCode

`func (o *DynamicEmergencyAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DynamicEmergencyAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DynamicEmergencyAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DynamicEmergencyAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DynamicEmergencyAddress) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DynamicEmergencyAddress) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DynamicEmergencyAddress) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DynamicEmergencyAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DynamicEmergencyAddress) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DynamicEmergencyAddress) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DynamicEmergencyAddress) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DynamicEmergencyAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


