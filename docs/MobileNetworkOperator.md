# MobileNetworkOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Name** | Pointer to **string** | The network operator name. | [optional] 
**Tadig** | Pointer to **string** | TADIG stands for Transferred Account Data Interchange Group. The TADIG code is a unique identifier for network operators in GSM mobile networks. | [optional] 
**CountryCode** | Pointer to **string** | The mobile operator two-character (ISO 3166-1 alpha-2) origin country code. | [optional] 
**Mcc** | Pointer to **string** | MCC stands for Mobile Country Code. It&#39;s a three decimal digit that identifies a country.&lt;br/&gt;&lt;br/&gt; This code is commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] 
**Mnc** | Pointer to **string** | MNC stands for Mobile Network Code. It&#39;s a two to three decimal digits that identify a network.&lt;br/&gt;&lt;br/&gt;  This code is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] 
**NetworkPreferencesEnabled** | Pointer to **bool** | Indicate whether the mobile network operator can be set as preferred in the Network Preferences API. | [optional] [readonly] 

## Methods

### NewMobileNetworkOperator

`func NewMobileNetworkOperator() *MobileNetworkOperator`

NewMobileNetworkOperator instantiates a new MobileNetworkOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileNetworkOperatorWithDefaults

`func NewMobileNetworkOperatorWithDefaults() *MobileNetworkOperator`

NewMobileNetworkOperatorWithDefaults instantiates a new MobileNetworkOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobileNetworkOperator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobileNetworkOperator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobileNetworkOperator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MobileNetworkOperator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *MobileNetworkOperator) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MobileNetworkOperator) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MobileNetworkOperator) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MobileNetworkOperator) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetName

`func (o *MobileNetworkOperator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileNetworkOperator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileNetworkOperator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileNetworkOperator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTadig

`func (o *MobileNetworkOperator) GetTadig() string`

GetTadig returns the Tadig field if non-nil, zero value otherwise.

### GetTadigOk

`func (o *MobileNetworkOperator) GetTadigOk() (*string, bool)`

GetTadigOk returns a tuple with the Tadig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTadig

`func (o *MobileNetworkOperator) SetTadig(v string)`

SetTadig sets Tadig field to given value.

### HasTadig

`func (o *MobileNetworkOperator) HasTadig() bool`

HasTadig returns a boolean if a field has been set.

### GetCountryCode

`func (o *MobileNetworkOperator) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MobileNetworkOperator) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MobileNetworkOperator) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MobileNetworkOperator) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetMcc

`func (o *MobileNetworkOperator) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *MobileNetworkOperator) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *MobileNetworkOperator) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *MobileNetworkOperator) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *MobileNetworkOperator) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *MobileNetworkOperator) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *MobileNetworkOperator) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *MobileNetworkOperator) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetNetworkPreferencesEnabled

`func (o *MobileNetworkOperator) GetNetworkPreferencesEnabled() bool`

GetNetworkPreferencesEnabled returns the NetworkPreferencesEnabled field if non-nil, zero value otherwise.

### GetNetworkPreferencesEnabledOk

`func (o *MobileNetworkOperator) GetNetworkPreferencesEnabledOk() (*bool, bool)`

GetNetworkPreferencesEnabledOk returns a tuple with the NetworkPreferencesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPreferencesEnabled

`func (o *MobileNetworkOperator) SetNetworkPreferencesEnabled(v bool)`

SetNetworkPreferencesEnabled sets NetworkPreferencesEnabled field to given value.

### HasNetworkPreferencesEnabled

`func (o *MobileNetworkOperator) HasNetworkPreferencesEnabled() bool`

HasNetworkPreferencesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


