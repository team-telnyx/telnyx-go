# CountryCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Country ISO code | [optional] 
**Numbers** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to **[]string** | Set of features supported | [optional] 
**PhoneNumberType** | Pointer to **[]string** | Phone number type | [optional] 
**Reservable** | Pointer to **bool** | Supports reservable | [optional] 
**Quickship** | Pointer to **bool** | Supports quickship | [optional] 
**InternationalSms** | Pointer to **bool** |  | [optional] 
**P2p** | Pointer to **bool** |  | [optional] 
**Local** | Pointer to [**CountryCoverageLocal**](CountryCoverageLocal.md) |  | [optional] 
**TollFree** | Pointer to [**CountryCoverageLocal**](CountryCoverageLocal.md) |  | [optional] 
**Mobile** | Pointer to **map[string]interface{}** |  | [optional] 
**National** | Pointer to **map[string]interface{}** |  | [optional] 
**InventoryCoverage** | Pointer to **bool** | Indicates whether country can be queried with inventory coverage endpoint | [optional] 
**SharedCost** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCountryCoverage

`func NewCountryCoverage() *CountryCoverage`

NewCountryCoverage instantiates a new CountryCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryCoverageWithDefaults

`func NewCountryCoverageWithDefaults() *CountryCoverage`

NewCountryCoverageWithDefaults instantiates a new CountryCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CountryCoverage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryCoverage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryCoverage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CountryCoverage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetNumbers

`func (o *CountryCoverage) GetNumbers() bool`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *CountryCoverage) GetNumbersOk() (*bool, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *CountryCoverage) SetNumbers(v bool)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *CountryCoverage) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetFeatures

`func (o *CountryCoverage) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CountryCoverage) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CountryCoverage) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CountryCoverage) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *CountryCoverage) GetPhoneNumberType() []string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *CountryCoverage) GetPhoneNumberTypeOk() (*[]string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *CountryCoverage) SetPhoneNumberType(v []string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *CountryCoverage) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetReservable

`func (o *CountryCoverage) GetReservable() bool`

GetReservable returns the Reservable field if non-nil, zero value otherwise.

### GetReservableOk

`func (o *CountryCoverage) GetReservableOk() (*bool, bool)`

GetReservableOk returns a tuple with the Reservable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservable

`func (o *CountryCoverage) SetReservable(v bool)`

SetReservable sets Reservable field to given value.

### HasReservable

`func (o *CountryCoverage) HasReservable() bool`

HasReservable returns a boolean if a field has been set.

### GetQuickship

`func (o *CountryCoverage) GetQuickship() bool`

GetQuickship returns the Quickship field if non-nil, zero value otherwise.

### GetQuickshipOk

`func (o *CountryCoverage) GetQuickshipOk() (*bool, bool)`

GetQuickshipOk returns a tuple with the Quickship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickship

`func (o *CountryCoverage) SetQuickship(v bool)`

SetQuickship sets Quickship field to given value.

### HasQuickship

`func (o *CountryCoverage) HasQuickship() bool`

HasQuickship returns a boolean if a field has been set.

### GetInternationalSms

`func (o *CountryCoverage) GetInternationalSms() bool`

GetInternationalSms returns the InternationalSms field if non-nil, zero value otherwise.

### GetInternationalSmsOk

`func (o *CountryCoverage) GetInternationalSmsOk() (*bool, bool)`

GetInternationalSmsOk returns a tuple with the InternationalSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalSms

`func (o *CountryCoverage) SetInternationalSms(v bool)`

SetInternationalSms sets InternationalSms field to given value.

### HasInternationalSms

`func (o *CountryCoverage) HasInternationalSms() bool`

HasInternationalSms returns a boolean if a field has been set.

### GetP2p

`func (o *CountryCoverage) GetP2p() bool`

GetP2p returns the P2p field if non-nil, zero value otherwise.

### GetP2pOk

`func (o *CountryCoverage) GetP2pOk() (*bool, bool)`

GetP2pOk returns a tuple with the P2p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2p

`func (o *CountryCoverage) SetP2p(v bool)`

SetP2p sets P2p field to given value.

### HasP2p

`func (o *CountryCoverage) HasP2p() bool`

HasP2p returns a boolean if a field has been set.

### GetLocal

`func (o *CountryCoverage) GetLocal() CountryCoverageLocal`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *CountryCoverage) GetLocalOk() (*CountryCoverageLocal, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *CountryCoverage) SetLocal(v CountryCoverageLocal)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *CountryCoverage) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetTollFree

`func (o *CountryCoverage) GetTollFree() CountryCoverageLocal`

GetTollFree returns the TollFree field if non-nil, zero value otherwise.

### GetTollFreeOk

`func (o *CountryCoverage) GetTollFreeOk() (*CountryCoverageLocal, bool)`

GetTollFreeOk returns a tuple with the TollFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollFree

`func (o *CountryCoverage) SetTollFree(v CountryCoverageLocal)`

SetTollFree sets TollFree field to given value.

### HasTollFree

`func (o *CountryCoverage) HasTollFree() bool`

HasTollFree returns a boolean if a field has been set.

### GetMobile

`func (o *CountryCoverage) GetMobile() map[string]interface{}`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CountryCoverage) GetMobileOk() (*map[string]interface{}, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CountryCoverage) SetMobile(v map[string]interface{})`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *CountryCoverage) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetNational

`func (o *CountryCoverage) GetNational() map[string]interface{}`

GetNational returns the National field if non-nil, zero value otherwise.

### GetNationalOk

`func (o *CountryCoverage) GetNationalOk() (*map[string]interface{}, bool)`

GetNationalOk returns a tuple with the National field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNational

`func (o *CountryCoverage) SetNational(v map[string]interface{})`

SetNational sets National field to given value.

### HasNational

`func (o *CountryCoverage) HasNational() bool`

HasNational returns a boolean if a field has been set.

### GetInventoryCoverage

`func (o *CountryCoverage) GetInventoryCoverage() bool`

GetInventoryCoverage returns the InventoryCoverage field if non-nil, zero value otherwise.

### GetInventoryCoverageOk

`func (o *CountryCoverage) GetInventoryCoverageOk() (*bool, bool)`

GetInventoryCoverageOk returns a tuple with the InventoryCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryCoverage

`func (o *CountryCoverage) SetInventoryCoverage(v bool)`

SetInventoryCoverage sets InventoryCoverage field to given value.

### HasInventoryCoverage

`func (o *CountryCoverage) HasInventoryCoverage() bool`

HasInventoryCoverage returns a boolean if a field has been set.

### GetSharedCost

`func (o *CountryCoverage) GetSharedCost() map[string]interface{}`

GetSharedCost returns the SharedCost field if non-nil, zero value otherwise.

### GetSharedCostOk

`func (o *CountryCoverage) GetSharedCostOk() (*map[string]interface{}, bool)`

GetSharedCostOk returns a tuple with the SharedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedCost

`func (o *CountryCoverage) SetSharedCost(v map[string]interface{})`

SetSharedCost sets SharedCost field to given value.

### HasSharedCost

`func (o *CountryCoverage) HasSharedCost() bool`

HasSharedCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


