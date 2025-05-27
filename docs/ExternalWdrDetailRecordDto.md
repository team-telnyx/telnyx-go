# ExternalWdrDetailRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | WDR id | [optional] 
**CreatedAt** | Pointer to **time.Time** | Record created time | [optional] 
**Cost** | Pointer to [**WirelessCost**](WirelessCost.md) |  | [optional] 
**Mcc** | Pointer to **string** | Mobile country code. | [optional] 
**Mnc** | Pointer to **string** | Mobile network code. | [optional] 
**DownlinkData** | Pointer to [**DownlinkData**](DownlinkData.md) |  | [optional] 
**DurationSeconds** | Pointer to **float32** | Session duration in seconds. | [optional] 
**Imsi** | Pointer to **string** | International mobile subscriber identity. | [optional] 
**Rate** | Pointer to [**WirelessRate**](WirelessRate.md) |  | [optional] 
**SimGroupName** | Pointer to **string** | Defined sim group name | [optional] 
**SimGroupId** | Pointer to **string** | Sim group unique identifier | [optional] 
**SimCardId** | Pointer to **string** | Sim card unique identifier | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number | [optional] 
**UplinkData** | Pointer to [**UplinkData**](UplinkData.md) |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalWdrDetailRecordDto

`func NewExternalWdrDetailRecordDto() *ExternalWdrDetailRecordDto`

NewExternalWdrDetailRecordDto instantiates a new ExternalWdrDetailRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalWdrDetailRecordDtoWithDefaults

`func NewExternalWdrDetailRecordDtoWithDefaults() *ExternalWdrDetailRecordDto`

NewExternalWdrDetailRecordDtoWithDefaults instantiates a new ExternalWdrDetailRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalWdrDetailRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalWdrDetailRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalWdrDetailRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalWdrDetailRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExternalWdrDetailRecordDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalWdrDetailRecordDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalWdrDetailRecordDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalWdrDetailRecordDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCost

`func (o *ExternalWdrDetailRecordDto) GetCost() WirelessCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ExternalWdrDetailRecordDto) GetCostOk() (*WirelessCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ExternalWdrDetailRecordDto) SetCost(v WirelessCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ExternalWdrDetailRecordDto) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetMcc

`func (o *ExternalWdrDetailRecordDto) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *ExternalWdrDetailRecordDto) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *ExternalWdrDetailRecordDto) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *ExternalWdrDetailRecordDto) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *ExternalWdrDetailRecordDto) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *ExternalWdrDetailRecordDto) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *ExternalWdrDetailRecordDto) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *ExternalWdrDetailRecordDto) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetDownlinkData

`func (o *ExternalWdrDetailRecordDto) GetDownlinkData() DownlinkData`

GetDownlinkData returns the DownlinkData field if non-nil, zero value otherwise.

### GetDownlinkDataOk

`func (o *ExternalWdrDetailRecordDto) GetDownlinkDataOk() (*DownlinkData, bool)`

GetDownlinkDataOk returns a tuple with the DownlinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkData

`func (o *ExternalWdrDetailRecordDto) SetDownlinkData(v DownlinkData)`

SetDownlinkData sets DownlinkData field to given value.

### HasDownlinkData

`func (o *ExternalWdrDetailRecordDto) HasDownlinkData() bool`

HasDownlinkData returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *ExternalWdrDetailRecordDto) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ExternalWdrDetailRecordDto) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ExternalWdrDetailRecordDto) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *ExternalWdrDetailRecordDto) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetImsi

`func (o *ExternalWdrDetailRecordDto) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *ExternalWdrDetailRecordDto) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *ExternalWdrDetailRecordDto) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *ExternalWdrDetailRecordDto) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetRate

`func (o *ExternalWdrDetailRecordDto) GetRate() WirelessRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ExternalWdrDetailRecordDto) GetRateOk() (*WirelessRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ExternalWdrDetailRecordDto) SetRate(v WirelessRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ExternalWdrDetailRecordDto) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSimGroupName

`func (o *ExternalWdrDetailRecordDto) GetSimGroupName() string`

GetSimGroupName returns the SimGroupName field if non-nil, zero value otherwise.

### GetSimGroupNameOk

`func (o *ExternalWdrDetailRecordDto) GetSimGroupNameOk() (*string, bool)`

GetSimGroupNameOk returns a tuple with the SimGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupName

`func (o *ExternalWdrDetailRecordDto) SetSimGroupName(v string)`

SetSimGroupName sets SimGroupName field to given value.

### HasSimGroupName

`func (o *ExternalWdrDetailRecordDto) HasSimGroupName() bool`

HasSimGroupName returns a boolean if a field has been set.

### GetSimGroupId

`func (o *ExternalWdrDetailRecordDto) GetSimGroupId() string`

GetSimGroupId returns the SimGroupId field if non-nil, zero value otherwise.

### GetSimGroupIdOk

`func (o *ExternalWdrDetailRecordDto) GetSimGroupIdOk() (*string, bool)`

GetSimGroupIdOk returns a tuple with the SimGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupId

`func (o *ExternalWdrDetailRecordDto) SetSimGroupId(v string)`

SetSimGroupId sets SimGroupId field to given value.

### HasSimGroupId

`func (o *ExternalWdrDetailRecordDto) HasSimGroupId() bool`

HasSimGroupId returns a boolean if a field has been set.

### GetSimCardId

`func (o *ExternalWdrDetailRecordDto) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *ExternalWdrDetailRecordDto) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *ExternalWdrDetailRecordDto) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *ExternalWdrDetailRecordDto) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ExternalWdrDetailRecordDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ExternalWdrDetailRecordDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ExternalWdrDetailRecordDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ExternalWdrDetailRecordDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUplinkData

`func (o *ExternalWdrDetailRecordDto) GetUplinkData() UplinkData`

GetUplinkData returns the UplinkData field if non-nil, zero value otherwise.

### GetUplinkDataOk

`func (o *ExternalWdrDetailRecordDto) GetUplinkDataOk() (*UplinkData, bool)`

GetUplinkDataOk returns a tuple with the UplinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkData

`func (o *ExternalWdrDetailRecordDto) SetUplinkData(v UplinkData)`

SetUplinkData sets UplinkData field to given value.

### HasUplinkData

`func (o *ExternalWdrDetailRecordDto) HasUplinkData() bool`

HasUplinkData returns a boolean if a field has been set.

### GetRecordType

`func (o *ExternalWdrDetailRecordDto) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ExternalWdrDetailRecordDto) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ExternalWdrDetailRecordDto) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ExternalWdrDetailRecordDto) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


