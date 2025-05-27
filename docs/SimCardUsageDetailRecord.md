# SimCardUsageDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this SIM Card Usage | [optional] 
**CreatedAt** | Pointer to **time.Time** | Event creation time | [optional] 
**ClosedAt** | Pointer to **time.Time** | Event close time | [optional] 
**IpAddress** | Pointer to **string** | Ip address that generated the event | [optional] 
**DownlinkData** | Pointer to **float32** | Number of megabytes downloaded | [optional] 
**Imsi** | Pointer to **string** | International Mobile Subscriber Identity | [optional] 
**Mcc** | Pointer to **string** | Mobile country code | [optional] 
**Mnc** | Pointer to **string** | Mobile network code | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**DataUnit** | Pointer to **string** | Unit of wireless link consumption | [optional] 
**DataRate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**SimGroupName** | Pointer to **string** | Sim group name for sim card | [optional] 
**SimCardId** | Pointer to **string** | Unique identifier for SIM card | [optional] 
**SimGroupId** | Pointer to **string** | Unique identifier for SIM group | [optional] 
**SimCardTags** | Pointer to **string** | User-provided tags | [optional] 
**PhoneNumber** | Pointer to **string** | Telephone number associated to SIM card | [optional] 
**UplinkData** | Pointer to **float32** | Number of megabytes uploaded | [optional] 
**DataCost** | Pointer to **float32** | Data cost | [optional] 
**RecordType** | **string** |  | [default to "sim_card_usage"]

## Methods

### NewSimCardUsageDetailRecord

`func NewSimCardUsageDetailRecord(recordType string, ) *SimCardUsageDetailRecord`

NewSimCardUsageDetailRecord instantiates a new SimCardUsageDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimCardUsageDetailRecordWithDefaults

`func NewSimCardUsageDetailRecordWithDefaults() *SimCardUsageDetailRecord`

NewSimCardUsageDetailRecordWithDefaults instantiates a new SimCardUsageDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimCardUsageDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimCardUsageDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimCardUsageDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimCardUsageDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimCardUsageDetailRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimCardUsageDetailRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimCardUsageDetailRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimCardUsageDetailRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *SimCardUsageDetailRecord) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *SimCardUsageDetailRecord) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *SimCardUsageDetailRecord) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *SimCardUsageDetailRecord) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetIpAddress

`func (o *SimCardUsageDetailRecord) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SimCardUsageDetailRecord) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SimCardUsageDetailRecord) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SimCardUsageDetailRecord) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDownlinkData

`func (o *SimCardUsageDetailRecord) GetDownlinkData() float32`

GetDownlinkData returns the DownlinkData field if non-nil, zero value otherwise.

### GetDownlinkDataOk

`func (o *SimCardUsageDetailRecord) GetDownlinkDataOk() (*float32, bool)`

GetDownlinkDataOk returns a tuple with the DownlinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkData

`func (o *SimCardUsageDetailRecord) SetDownlinkData(v float32)`

SetDownlinkData sets DownlinkData field to given value.

### HasDownlinkData

`func (o *SimCardUsageDetailRecord) HasDownlinkData() bool`

HasDownlinkData returns a boolean if a field has been set.

### GetImsi

`func (o *SimCardUsageDetailRecord) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SimCardUsageDetailRecord) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SimCardUsageDetailRecord) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SimCardUsageDetailRecord) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetMcc

`func (o *SimCardUsageDetailRecord) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *SimCardUsageDetailRecord) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *SimCardUsageDetailRecord) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *SimCardUsageDetailRecord) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *SimCardUsageDetailRecord) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *SimCardUsageDetailRecord) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *SimCardUsageDetailRecord) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *SimCardUsageDetailRecord) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetCurrency

`func (o *SimCardUsageDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SimCardUsageDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SimCardUsageDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SimCardUsageDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDataUnit

`func (o *SimCardUsageDetailRecord) GetDataUnit() string`

GetDataUnit returns the DataUnit field if non-nil, zero value otherwise.

### GetDataUnitOk

`func (o *SimCardUsageDetailRecord) GetDataUnitOk() (*string, bool)`

GetDataUnitOk returns a tuple with the DataUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUnit

`func (o *SimCardUsageDetailRecord) SetDataUnit(v string)`

SetDataUnit sets DataUnit field to given value.

### HasDataUnit

`func (o *SimCardUsageDetailRecord) HasDataUnit() bool`

HasDataUnit returns a boolean if a field has been set.

### GetDataRate

`func (o *SimCardUsageDetailRecord) GetDataRate() string`

GetDataRate returns the DataRate field if non-nil, zero value otherwise.

### GetDataRateOk

`func (o *SimCardUsageDetailRecord) GetDataRateOk() (*string, bool)`

GetDataRateOk returns a tuple with the DataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRate

`func (o *SimCardUsageDetailRecord) SetDataRate(v string)`

SetDataRate sets DataRate field to given value.

### HasDataRate

`func (o *SimCardUsageDetailRecord) HasDataRate() bool`

HasDataRate returns a boolean if a field has been set.

### GetSimGroupName

`func (o *SimCardUsageDetailRecord) GetSimGroupName() string`

GetSimGroupName returns the SimGroupName field if non-nil, zero value otherwise.

### GetSimGroupNameOk

`func (o *SimCardUsageDetailRecord) GetSimGroupNameOk() (*string, bool)`

GetSimGroupNameOk returns a tuple with the SimGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupName

`func (o *SimCardUsageDetailRecord) SetSimGroupName(v string)`

SetSimGroupName sets SimGroupName field to given value.

### HasSimGroupName

`func (o *SimCardUsageDetailRecord) HasSimGroupName() bool`

HasSimGroupName returns a boolean if a field has been set.

### GetSimCardId

`func (o *SimCardUsageDetailRecord) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *SimCardUsageDetailRecord) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *SimCardUsageDetailRecord) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *SimCardUsageDetailRecord) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetSimGroupId

`func (o *SimCardUsageDetailRecord) GetSimGroupId() string`

GetSimGroupId returns the SimGroupId field if non-nil, zero value otherwise.

### GetSimGroupIdOk

`func (o *SimCardUsageDetailRecord) GetSimGroupIdOk() (*string, bool)`

GetSimGroupIdOk returns a tuple with the SimGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupId

`func (o *SimCardUsageDetailRecord) SetSimGroupId(v string)`

SetSimGroupId sets SimGroupId field to given value.

### HasSimGroupId

`func (o *SimCardUsageDetailRecord) HasSimGroupId() bool`

HasSimGroupId returns a boolean if a field has been set.

### GetSimCardTags

`func (o *SimCardUsageDetailRecord) GetSimCardTags() string`

GetSimCardTags returns the SimCardTags field if non-nil, zero value otherwise.

### GetSimCardTagsOk

`func (o *SimCardUsageDetailRecord) GetSimCardTagsOk() (*string, bool)`

GetSimCardTagsOk returns a tuple with the SimCardTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardTags

`func (o *SimCardUsageDetailRecord) SetSimCardTags(v string)`

SetSimCardTags sets SimCardTags field to given value.

### HasSimCardTags

`func (o *SimCardUsageDetailRecord) HasSimCardTags() bool`

HasSimCardTags returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *SimCardUsageDetailRecord) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SimCardUsageDetailRecord) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SimCardUsageDetailRecord) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SimCardUsageDetailRecord) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUplinkData

`func (o *SimCardUsageDetailRecord) GetUplinkData() float32`

GetUplinkData returns the UplinkData field if non-nil, zero value otherwise.

### GetUplinkDataOk

`func (o *SimCardUsageDetailRecord) GetUplinkDataOk() (*float32, bool)`

GetUplinkDataOk returns a tuple with the UplinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkData

`func (o *SimCardUsageDetailRecord) SetUplinkData(v float32)`

SetUplinkData sets UplinkData field to given value.

### HasUplinkData

`func (o *SimCardUsageDetailRecord) HasUplinkData() bool`

HasUplinkData returns a boolean if a field has been set.

### GetDataCost

`func (o *SimCardUsageDetailRecord) GetDataCost() float32`

GetDataCost returns the DataCost field if non-nil, zero value otherwise.

### GetDataCostOk

`func (o *SimCardUsageDetailRecord) GetDataCostOk() (*float32, bool)`

GetDataCostOk returns a tuple with the DataCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCost

`func (o *SimCardUsageDetailRecord) SetDataCost(v float32)`

SetDataCost sets DataCost field to given value.

### HasDataCost

`func (o *SimCardUsageDetailRecord) HasDataCost() bool`

HasDataCost returns a boolean if a field has been set.

### GetRecordType

`func (o *SimCardUsageDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SimCardUsageDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SimCardUsageDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


