# SIMCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**SIMCardStatus**](SIMCardStatus.md) |  | [optional] 
**Type** | Pointer to **string** | The type of SIM card | [optional] [readonly] 
**Iccid** | Pointer to **string** | The ICCID is the identifier of the specific SIM card/chip. Each SIM is internationally identified by its integrated circuit card identifier (ICCID). ICCIDs are stored in the SIM card&#39;s memory and are also engraved or printed on the SIM card body during a process called personalization.  | [optional] [readonly] 
**Imsi** | Pointer to **string** | SIM cards are identified on their individual network operators by a unique International Mobile Subscriber Identity (IMSI). &lt;br/&gt; Mobile network operators connect mobile phone calls and communicate with their market SIM cards using their IMSIs. The IMSI is stored in the Subscriber  Identity Module (SIM) inside the device and is sent by the device to the appropriate network. It is used to acquire the details of the device in the Home  Location Register (HLR) or the Visitor Location Register (VLR).  | [optional] [readonly] 
**Msisdn** | Pointer to **string** | Mobile Station International Subscriber Directory Number (MSISDN) is a number used to identify a mobile phone number internationally. &lt;br/&gt; MSISDN is defined by the E.164 numbering plan. It includes a country code and a National Destination Code which identifies the subscriber&#39;s operator.  | [optional] [readonly] 
**SimCardGroupId** | Pointer to **string** | The group SIMCardGroup identification. This attribute can be &lt;code&gt;null&lt;/code&gt; when it&#39;s present in an associated resource. | [optional] 
**Tags** | Pointer to **[]string** | Searchable tags associated with the SIM card | [optional] 
**AuthorizedImeis** | Pointer to **[]string** | List of IMEIs authorized to use a given SIM card. | [optional] 
**CurrentImei** | Pointer to **string** | IMEI of the device where a given SIM card is currently being used. | [optional] [readonly] 
**DataLimit** | Pointer to [**SIMCardDataLimit**](SIMCardDataLimit.md) |  | [optional] 
**CurrentBillingPeriodConsumedData** | Pointer to [**SIMCardCurrentBillingPeriodConsumedData**](SIMCardCurrentBillingPeriodConsumedData.md) |  | [optional] 
**ActionsInProgress** | Pointer to **bool** | Indicate whether the SIM card has any pending (in-progress) actions. | [optional] [readonly] [default to false]
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**Ipv4** | Pointer to **string** | The SIM&#39;s address in the currently connected network. This IPv4 address is usually obtained dynamically, so it may vary according to the location or new connections.  | [optional] [readonly] 
**Ipv6** | Pointer to **string** | The SIM&#39;s address in the currently connected network. This IPv6 address is usually obtained dynamically, so it may vary according to the location or new connections.  | [optional] [readonly] 
**CurrentDeviceLocation** | Pointer to [**SIMCardCurrentDeviceLocation**](SIMCardCurrentDeviceLocation.md) |  | [optional] 
**CurrentMnc** | Pointer to **string** | Mobile Network Code of the current network to which the SIM card is connected. It&#39;s a two to three decimal digits that identify a network.&lt;br/&gt;&lt;br/&gt;  This code is commonly seen joined with a Mobile Country Code (MCC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] [readonly] 
**CurrentMcc** | Pointer to **string** | Mobile Country Code of the current network to which the SIM card is connected. It&#39;s a three decimal digit that identifies a country.&lt;br/&gt;&lt;br/&gt; This code is commonly seen joined with a Mobile Network Code (MNC) in a tuple that allows identifying a carrier known as PLMN (Public Land Mobile Network) code. | [optional] [readonly] 
**LiveDataSession** | Pointer to **string** | Indicates whether the device is actively connected to a network and able to run data. | [optional] [readonly] 

## Methods

### NewSIMCard

`func NewSIMCard() *SIMCard`

NewSIMCard instantiates a new SIMCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardWithDefaults

`func NewSIMCardWithDefaults() *SIMCard`

NewSIMCardWithDefaults instantiates a new SIMCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCard) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCard) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCard) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCard) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStatus

`func (o *SIMCard) GetStatus() SIMCardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SIMCard) GetStatusOk() (*SIMCardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SIMCard) SetStatus(v SIMCardStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SIMCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SIMCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SIMCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SIMCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SIMCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIccid

`func (o *SIMCard) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *SIMCard) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *SIMCard) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *SIMCard) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImsi

`func (o *SIMCard) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SIMCard) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SIMCard) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SIMCard) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetMsisdn

`func (o *SIMCard) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *SIMCard) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *SIMCard) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *SIMCard) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetSimCardGroupId

`func (o *SIMCard) GetSimCardGroupId() string`

GetSimCardGroupId returns the SimCardGroupId field if non-nil, zero value otherwise.

### GetSimCardGroupIdOk

`func (o *SIMCard) GetSimCardGroupIdOk() (*string, bool)`

GetSimCardGroupIdOk returns a tuple with the SimCardGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardGroupId

`func (o *SIMCard) SetSimCardGroupId(v string)`

SetSimCardGroupId sets SimCardGroupId field to given value.

### HasSimCardGroupId

`func (o *SIMCard) HasSimCardGroupId() bool`

HasSimCardGroupId returns a boolean if a field has been set.

### GetTags

`func (o *SIMCard) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SIMCard) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SIMCard) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SIMCard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAuthorizedImeis

`func (o *SIMCard) GetAuthorizedImeis() []string`

GetAuthorizedImeis returns the AuthorizedImeis field if non-nil, zero value otherwise.

### GetAuthorizedImeisOk

`func (o *SIMCard) GetAuthorizedImeisOk() (*[]string, bool)`

GetAuthorizedImeisOk returns a tuple with the AuthorizedImeis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedImeis

`func (o *SIMCard) SetAuthorizedImeis(v []string)`

SetAuthorizedImeis sets AuthorizedImeis field to given value.

### HasAuthorizedImeis

`func (o *SIMCard) HasAuthorizedImeis() bool`

HasAuthorizedImeis returns a boolean if a field has been set.

### GetCurrentImei

`func (o *SIMCard) GetCurrentImei() string`

GetCurrentImei returns the CurrentImei field if non-nil, zero value otherwise.

### GetCurrentImeiOk

`func (o *SIMCard) GetCurrentImeiOk() (*string, bool)`

GetCurrentImeiOk returns a tuple with the CurrentImei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentImei

`func (o *SIMCard) SetCurrentImei(v string)`

SetCurrentImei sets CurrentImei field to given value.

### HasCurrentImei

`func (o *SIMCard) HasCurrentImei() bool`

HasCurrentImei returns a boolean if a field has been set.

### GetDataLimit

`func (o *SIMCard) GetDataLimit() SIMCardDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SIMCard) GetDataLimitOk() (*SIMCardDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SIMCard) SetDataLimit(v SIMCardDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SIMCard) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### GetCurrentBillingPeriodConsumedData

`func (o *SIMCard) GetCurrentBillingPeriodConsumedData() SIMCardCurrentBillingPeriodConsumedData`

GetCurrentBillingPeriodConsumedData returns the CurrentBillingPeriodConsumedData field if non-nil, zero value otherwise.

### GetCurrentBillingPeriodConsumedDataOk

`func (o *SIMCard) GetCurrentBillingPeriodConsumedDataOk() (*SIMCardCurrentBillingPeriodConsumedData, bool)`

GetCurrentBillingPeriodConsumedDataOk returns a tuple with the CurrentBillingPeriodConsumedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBillingPeriodConsumedData

`func (o *SIMCard) SetCurrentBillingPeriodConsumedData(v SIMCardCurrentBillingPeriodConsumedData)`

SetCurrentBillingPeriodConsumedData sets CurrentBillingPeriodConsumedData field to given value.

### HasCurrentBillingPeriodConsumedData

`func (o *SIMCard) HasCurrentBillingPeriodConsumedData() bool`

HasCurrentBillingPeriodConsumedData returns a boolean if a field has been set.

### GetActionsInProgress

`func (o *SIMCard) GetActionsInProgress() bool`

GetActionsInProgress returns the ActionsInProgress field if non-nil, zero value otherwise.

### GetActionsInProgressOk

`func (o *SIMCard) GetActionsInProgressOk() (*bool, bool)`

GetActionsInProgressOk returns a tuple with the ActionsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsInProgress

`func (o *SIMCard) SetActionsInProgress(v bool)`

SetActionsInProgress sets ActionsInProgress field to given value.

### HasActionsInProgress

`func (o *SIMCard) HasActionsInProgress() bool`

HasActionsInProgress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SIMCard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCard) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCard) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCard) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIpv4

`func (o *SIMCard) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *SIMCard) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *SIMCard) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *SIMCard) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *SIMCard) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *SIMCard) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *SIMCard) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *SIMCard) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetCurrentDeviceLocation

`func (o *SIMCard) GetCurrentDeviceLocation() SIMCardCurrentDeviceLocation`

GetCurrentDeviceLocation returns the CurrentDeviceLocation field if non-nil, zero value otherwise.

### GetCurrentDeviceLocationOk

`func (o *SIMCard) GetCurrentDeviceLocationOk() (*SIMCardCurrentDeviceLocation, bool)`

GetCurrentDeviceLocationOk returns a tuple with the CurrentDeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeviceLocation

`func (o *SIMCard) SetCurrentDeviceLocation(v SIMCardCurrentDeviceLocation)`

SetCurrentDeviceLocation sets CurrentDeviceLocation field to given value.

### HasCurrentDeviceLocation

`func (o *SIMCard) HasCurrentDeviceLocation() bool`

HasCurrentDeviceLocation returns a boolean if a field has been set.

### GetCurrentMnc

`func (o *SIMCard) GetCurrentMnc() string`

GetCurrentMnc returns the CurrentMnc field if non-nil, zero value otherwise.

### GetCurrentMncOk

`func (o *SIMCard) GetCurrentMncOk() (*string, bool)`

GetCurrentMncOk returns a tuple with the CurrentMnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMnc

`func (o *SIMCard) SetCurrentMnc(v string)`

SetCurrentMnc sets CurrentMnc field to given value.

### HasCurrentMnc

`func (o *SIMCard) HasCurrentMnc() bool`

HasCurrentMnc returns a boolean if a field has been set.

### GetCurrentMcc

`func (o *SIMCard) GetCurrentMcc() string`

GetCurrentMcc returns the CurrentMcc field if non-nil, zero value otherwise.

### GetCurrentMccOk

`func (o *SIMCard) GetCurrentMccOk() (*string, bool)`

GetCurrentMccOk returns a tuple with the CurrentMcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMcc

`func (o *SIMCard) SetCurrentMcc(v string)`

SetCurrentMcc sets CurrentMcc field to given value.

### HasCurrentMcc

`func (o *SIMCard) HasCurrentMcc() bool`

HasCurrentMcc returns a boolean if a field has been set.

### GetLiveDataSession

`func (o *SIMCard) GetLiveDataSession() string`

GetLiveDataSession returns the LiveDataSession field if non-nil, zero value otherwise.

### GetLiveDataSessionOk

`func (o *SIMCard) GetLiveDataSessionOk() (*string, bool)`

GetLiveDataSessionOk returns a tuple with the LiveDataSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDataSession

`func (o *SIMCard) SetLiveDataSession(v string)`

SetLiveDataSession sets LiveDataSession field to given value.

### HasLiveDataSession

`func (o *SIMCard) HasLiveDataSession() bool`

HasLiveDataSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


