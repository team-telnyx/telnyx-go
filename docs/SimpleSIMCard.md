# SimpleSIMCard

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
**DataLimit** | Pointer to [**SimpleSIMCardDataLimit**](SimpleSIMCardDataLimit.md) |  | [optional] 
**CurrentBillingPeriodConsumedData** | Pointer to [**SIMCardCurrentBillingPeriodConsumedData**](SIMCardCurrentBillingPeriodConsumedData.md) |  | [optional] 
**ActionsInProgress** | Pointer to **bool** | Indicate whether the SIM card has any pending (in-progress) actions. | [optional] [readonly] [default to false]
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSimpleSIMCard

`func NewSimpleSIMCard() *SimpleSIMCard`

NewSimpleSIMCard instantiates a new SimpleSIMCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSIMCardWithDefaults

`func NewSimpleSIMCardWithDefaults() *SimpleSIMCard`

NewSimpleSIMCardWithDefaults instantiates a new SimpleSIMCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleSIMCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleSIMCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleSIMCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleSIMCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SimpleSIMCard) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SimpleSIMCard) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SimpleSIMCard) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SimpleSIMCard) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleSIMCard) GetStatus() SIMCardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleSIMCard) GetStatusOk() (*SIMCardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleSIMCard) SetStatus(v SIMCardStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleSIMCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SimpleSIMCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleSIMCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleSIMCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleSIMCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIccid

`func (o *SimpleSIMCard) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *SimpleSIMCard) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *SimpleSIMCard) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *SimpleSIMCard) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### GetImsi

`func (o *SimpleSIMCard) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *SimpleSIMCard) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *SimpleSIMCard) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *SimpleSIMCard) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetMsisdn

`func (o *SimpleSIMCard) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *SimpleSIMCard) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *SimpleSIMCard) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *SimpleSIMCard) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetSimCardGroupId

`func (o *SimpleSIMCard) GetSimCardGroupId() string`

GetSimCardGroupId returns the SimCardGroupId field if non-nil, zero value otherwise.

### GetSimCardGroupIdOk

`func (o *SimpleSIMCard) GetSimCardGroupIdOk() (*string, bool)`

GetSimCardGroupIdOk returns a tuple with the SimCardGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardGroupId

`func (o *SimpleSIMCard) SetSimCardGroupId(v string)`

SetSimCardGroupId sets SimCardGroupId field to given value.

### HasSimCardGroupId

`func (o *SimpleSIMCard) HasSimCardGroupId() bool`

HasSimCardGroupId returns a boolean if a field has been set.

### GetTags

`func (o *SimpleSIMCard) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SimpleSIMCard) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SimpleSIMCard) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SimpleSIMCard) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataLimit

`func (o *SimpleSIMCard) GetDataLimit() SimpleSIMCardDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SimpleSIMCard) GetDataLimitOk() (*SimpleSIMCardDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SimpleSIMCard) SetDataLimit(v SimpleSIMCardDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SimpleSIMCard) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### GetCurrentBillingPeriodConsumedData

`func (o *SimpleSIMCard) GetCurrentBillingPeriodConsumedData() SIMCardCurrentBillingPeriodConsumedData`

GetCurrentBillingPeriodConsumedData returns the CurrentBillingPeriodConsumedData field if non-nil, zero value otherwise.

### GetCurrentBillingPeriodConsumedDataOk

`func (o *SimpleSIMCard) GetCurrentBillingPeriodConsumedDataOk() (*SIMCardCurrentBillingPeriodConsumedData, bool)`

GetCurrentBillingPeriodConsumedDataOk returns a tuple with the CurrentBillingPeriodConsumedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBillingPeriodConsumedData

`func (o *SimpleSIMCard) SetCurrentBillingPeriodConsumedData(v SIMCardCurrentBillingPeriodConsumedData)`

SetCurrentBillingPeriodConsumedData sets CurrentBillingPeriodConsumedData field to given value.

### HasCurrentBillingPeriodConsumedData

`func (o *SimpleSIMCard) HasCurrentBillingPeriodConsumedData() bool`

HasCurrentBillingPeriodConsumedData returns a boolean if a field has been set.

### GetActionsInProgress

`func (o *SimpleSIMCard) GetActionsInProgress() bool`

GetActionsInProgress returns the ActionsInProgress field if non-nil, zero value otherwise.

### GetActionsInProgressOk

`func (o *SimpleSIMCard) GetActionsInProgressOk() (*bool, bool)`

GetActionsInProgressOk returns a tuple with the ActionsInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsInProgress

`func (o *SimpleSIMCard) SetActionsInProgress(v bool)`

SetActionsInProgress sets ActionsInProgress field to given value.

### HasActionsInProgress

`func (o *SimpleSIMCard) HasActionsInProgress() bool`

HasActionsInProgress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimpleSIMCard) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimpleSIMCard) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimpleSIMCard) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimpleSIMCard) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimpleSIMCard) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimpleSIMCard) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimpleSIMCard) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SimpleSIMCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


