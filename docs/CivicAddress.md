# CivicAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**CityOrTown** | Pointer to **string** |  | [optional] 
**CityOrTownAlias** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryOrDistrict** | Pointer to **string** |  | [optional] 
**DefaultLocationId** | Pointer to **string** | Identifies what is the default location in the list of locations. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HouseNumber** | Pointer to **string** |  | [optional] 
**HouseNumberSuffix** | Pointer to **string** |  | [optional] 
**PostalOrZipCode** | Pointer to **string** |  | [optional] 
**StateOrProvince** | Pointer to **string** |  | [optional] 
**StreetName** | Pointer to **string** |  | [optional] 
**StreetSuffix** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 

## Methods

### NewCivicAddress

`func NewCivicAddress() *CivicAddress`

NewCivicAddress instantiates a new CivicAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCivicAddressWithDefaults

`func NewCivicAddressWithDefaults() *CivicAddress`

NewCivicAddressWithDefaults instantiates a new CivicAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CivicAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CivicAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CivicAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CivicAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *CivicAddress) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CivicAddress) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CivicAddress) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CivicAddress) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCityOrTown

`func (o *CivicAddress) GetCityOrTown() string`

GetCityOrTown returns the CityOrTown field if non-nil, zero value otherwise.

### GetCityOrTownOk

`func (o *CivicAddress) GetCityOrTownOk() (*string, bool)`

GetCityOrTownOk returns a tuple with the CityOrTown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityOrTown

`func (o *CivicAddress) SetCityOrTown(v string)`

SetCityOrTown sets CityOrTown field to given value.

### HasCityOrTown

`func (o *CivicAddress) HasCityOrTown() bool`

HasCityOrTown returns a boolean if a field has been set.

### GetCityOrTownAlias

`func (o *CivicAddress) GetCityOrTownAlias() string`

GetCityOrTownAlias returns the CityOrTownAlias field if non-nil, zero value otherwise.

### GetCityOrTownAliasOk

`func (o *CivicAddress) GetCityOrTownAliasOk() (*string, bool)`

GetCityOrTownAliasOk returns a tuple with the CityOrTownAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityOrTownAlias

`func (o *CivicAddress) SetCityOrTownAlias(v string)`

SetCityOrTownAlias sets CityOrTownAlias field to given value.

### HasCityOrTownAlias

`func (o *CivicAddress) HasCityOrTownAlias() bool`

HasCityOrTownAlias returns a boolean if a field has been set.

### GetCompanyName

`func (o *CivicAddress) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CivicAddress) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CivicAddress) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CivicAddress) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *CivicAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CivicAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CivicAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CivicAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryOrDistrict

`func (o *CivicAddress) GetCountryOrDistrict() string`

GetCountryOrDistrict returns the CountryOrDistrict field if non-nil, zero value otherwise.

### GetCountryOrDistrictOk

`func (o *CivicAddress) GetCountryOrDistrictOk() (*string, bool)`

GetCountryOrDistrictOk returns a tuple with the CountryOrDistrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOrDistrict

`func (o *CivicAddress) SetCountryOrDistrict(v string)`

SetCountryOrDistrict sets CountryOrDistrict field to given value.

### HasCountryOrDistrict

`func (o *CivicAddress) HasCountryOrDistrict() bool`

HasCountryOrDistrict returns a boolean if a field has been set.

### GetDefaultLocationId

`func (o *CivicAddress) GetDefaultLocationId() string`

GetDefaultLocationId returns the DefaultLocationId field if non-nil, zero value otherwise.

### GetDefaultLocationIdOk

`func (o *CivicAddress) GetDefaultLocationIdOk() (*string, bool)`

GetDefaultLocationIdOk returns a tuple with the DefaultLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocationId

`func (o *CivicAddress) SetDefaultLocationId(v string)`

SetDefaultLocationId sets DefaultLocationId field to given value.

### HasDefaultLocationId

`func (o *CivicAddress) HasDefaultLocationId() bool`

HasDefaultLocationId returns a boolean if a field has been set.

### GetDescription

`func (o *CivicAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CivicAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CivicAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CivicAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHouseNumber

`func (o *CivicAddress) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *CivicAddress) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *CivicAddress) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *CivicAddress) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetHouseNumberSuffix

`func (o *CivicAddress) GetHouseNumberSuffix() string`

GetHouseNumberSuffix returns the HouseNumberSuffix field if non-nil, zero value otherwise.

### GetHouseNumberSuffixOk

`func (o *CivicAddress) GetHouseNumberSuffixOk() (*string, bool)`

GetHouseNumberSuffixOk returns a tuple with the HouseNumberSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumberSuffix

`func (o *CivicAddress) SetHouseNumberSuffix(v string)`

SetHouseNumberSuffix sets HouseNumberSuffix field to given value.

### HasHouseNumberSuffix

`func (o *CivicAddress) HasHouseNumberSuffix() bool`

HasHouseNumberSuffix returns a boolean if a field has been set.

### GetPostalOrZipCode

`func (o *CivicAddress) GetPostalOrZipCode() string`

GetPostalOrZipCode returns the PostalOrZipCode field if non-nil, zero value otherwise.

### GetPostalOrZipCodeOk

`func (o *CivicAddress) GetPostalOrZipCodeOk() (*string, bool)`

GetPostalOrZipCodeOk returns a tuple with the PostalOrZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalOrZipCode

`func (o *CivicAddress) SetPostalOrZipCode(v string)`

SetPostalOrZipCode sets PostalOrZipCode field to given value.

### HasPostalOrZipCode

`func (o *CivicAddress) HasPostalOrZipCode() bool`

HasPostalOrZipCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *CivicAddress) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *CivicAddress) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *CivicAddress) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *CivicAddress) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### GetStreetName

`func (o *CivicAddress) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *CivicAddress) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *CivicAddress) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *CivicAddress) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### GetStreetSuffix

`func (o *CivicAddress) GetStreetSuffix() string`

GetStreetSuffix returns the StreetSuffix field if non-nil, zero value otherwise.

### GetStreetSuffixOk

`func (o *CivicAddress) GetStreetSuffixOk() (*string, bool)`

GetStreetSuffixOk returns a tuple with the StreetSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetSuffix

`func (o *CivicAddress) SetStreetSuffix(v string)`

SetStreetSuffix sets StreetSuffix field to given value.

### HasStreetSuffix

`func (o *CivicAddress) HasStreetSuffix() bool`

HasStreetSuffix returns a boolean if a field has been set.

### GetLocations

`func (o *CivicAddress) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *CivicAddress) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *CivicAddress) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *CivicAddress) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


