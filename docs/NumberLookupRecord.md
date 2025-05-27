# NumberLookupRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of record | [optional] 
**CountryCode** | Pointer to **string** | Region code that matches the specific country calling code | [optional] 
**NationalFormat** | Pointer to **string** | Hyphen-separated national number, preceded by the national destination code (NDC), with a 0 prefix, if an NDC is found | [optional] 
**PhoneNumber** | Pointer to **string** | E164-formatted phone number | [optional] 
**Fraud** | Pointer to **string** | Unused | [optional] 
**Carrier** | Pointer to [**Carrier**](Carrier.md) |  | [optional] 
**CallerName** | Pointer to [**CallerName**](CallerName.md) |  | [optional] 
**Portability** | Pointer to [**Portability**](Portability.md) |  | [optional] 

## Methods

### NewNumberLookupRecord

`func NewNumberLookupRecord() *NumberLookupRecord`

NewNumberLookupRecord instantiates a new NumberLookupRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberLookupRecordWithDefaults

`func NewNumberLookupRecordWithDefaults() *NumberLookupRecord`

NewNumberLookupRecordWithDefaults instantiates a new NumberLookupRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *NumberLookupRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberLookupRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberLookupRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberLookupRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCountryCode

`func (o *NumberLookupRecord) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *NumberLookupRecord) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *NumberLookupRecord) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *NumberLookupRecord) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetNationalFormat

`func (o *NumberLookupRecord) GetNationalFormat() string`

GetNationalFormat returns the NationalFormat field if non-nil, zero value otherwise.

### GetNationalFormatOk

`func (o *NumberLookupRecord) GetNationalFormatOk() (*string, bool)`

GetNationalFormatOk returns a tuple with the NationalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalFormat

`func (o *NumberLookupRecord) SetNationalFormat(v string)`

SetNationalFormat sets NationalFormat field to given value.

### HasNationalFormat

`func (o *NumberLookupRecord) HasNationalFormat() bool`

HasNationalFormat returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *NumberLookupRecord) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *NumberLookupRecord) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *NumberLookupRecord) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *NumberLookupRecord) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetFraud

`func (o *NumberLookupRecord) GetFraud() string`

GetFraud returns the Fraud field if non-nil, zero value otherwise.

### GetFraudOk

`func (o *NumberLookupRecord) GetFraudOk() (*string, bool)`

GetFraudOk returns a tuple with the Fraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraud

`func (o *NumberLookupRecord) SetFraud(v string)`

SetFraud sets Fraud field to given value.

### HasFraud

`func (o *NumberLookupRecord) HasFraud() bool`

HasFraud returns a boolean if a field has been set.

### GetCarrier

`func (o *NumberLookupRecord) GetCarrier() Carrier`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *NumberLookupRecord) GetCarrierOk() (*Carrier, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *NumberLookupRecord) SetCarrier(v Carrier)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *NumberLookupRecord) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCallerName

`func (o *NumberLookupRecord) GetCallerName() CallerName`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *NumberLookupRecord) GetCallerNameOk() (*CallerName, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *NumberLookupRecord) SetCallerName(v CallerName)`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *NumberLookupRecord) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### GetPortability

`func (o *NumberLookupRecord) GetPortability() Portability`

GetPortability returns the Portability field if non-nil, zero value otherwise.

### GetPortabilityOk

`func (o *NumberLookupRecord) GetPortabilityOk() (*Portability, bool)`

GetPortabilityOk returns a tuple with the Portability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortability

`func (o *NumberLookupRecord) SetPortability(v Portability)`

SetPortability sets Portability field to given value.

### HasPortability

`func (o *NumberLookupRecord) HasPortability() bool`

HasPortability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


