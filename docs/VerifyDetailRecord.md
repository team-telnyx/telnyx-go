# VerifyDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the verification | [optional] 
**VerifyProfileId** | Pointer to **string** |  | [optional] 
**DeliveryStatus** | Pointer to **string** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**DestinationPhoneNumber** | Pointer to **string** | E.164 formatted phone number | [optional] 
**VerifyChannelType** | Pointer to **string** | Depending on the type of verification, the &#x60;verify_channel_id&#x60; points to one of the following channel ids; --- verify_channel_type | verify_channel_id ------------------- | ----------------- sms, psd2           | messaging_id call, flashcall     | call_control_id ---  | [optional] 
**VerifyChannelId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing costs | [optional] 
**RateMeasuredIn** | Pointer to **string** | Billing unit used to calculate the Telnyx billing costs | [optional] 
**VerifyUsageFee** | Pointer to **string** | Currency amount for Verify Usage Fee | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing costs | [optional] 
**RecordType** | **string** |  | [default to "verification_detail_record"]

## Methods

### NewVerifyDetailRecord

`func NewVerifyDetailRecord(recordType string, ) *VerifyDetailRecord`

NewVerifyDetailRecord instantiates a new VerifyDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyDetailRecordWithDefaults

`func NewVerifyDetailRecordWithDefaults() *VerifyDetailRecord`

NewVerifyDetailRecordWithDefaults instantiates a new VerifyDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVerifyProfileId

`func (o *VerifyDetailRecord) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *VerifyDetailRecord) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *VerifyDetailRecord) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.

### HasVerifyProfileId

`func (o *VerifyDetailRecord) HasVerifyProfileId() bool`

HasVerifyProfileId returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *VerifyDetailRecord) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *VerifyDetailRecord) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *VerifyDetailRecord) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *VerifyDetailRecord) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *VerifyDetailRecord) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerifyDetailRecord) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerifyDetailRecord) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *VerifyDetailRecord) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetDestinationPhoneNumber

`func (o *VerifyDetailRecord) GetDestinationPhoneNumber() string`

GetDestinationPhoneNumber returns the DestinationPhoneNumber field if non-nil, zero value otherwise.

### GetDestinationPhoneNumberOk

`func (o *VerifyDetailRecord) GetDestinationPhoneNumberOk() (*string, bool)`

GetDestinationPhoneNumberOk returns a tuple with the DestinationPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPhoneNumber

`func (o *VerifyDetailRecord) SetDestinationPhoneNumber(v string)`

SetDestinationPhoneNumber sets DestinationPhoneNumber field to given value.

### HasDestinationPhoneNumber

`func (o *VerifyDetailRecord) HasDestinationPhoneNumber() bool`

HasDestinationPhoneNumber returns a boolean if a field has been set.

### GetVerifyChannelType

`func (o *VerifyDetailRecord) GetVerifyChannelType() string`

GetVerifyChannelType returns the VerifyChannelType field if non-nil, zero value otherwise.

### GetVerifyChannelTypeOk

`func (o *VerifyDetailRecord) GetVerifyChannelTypeOk() (*string, bool)`

GetVerifyChannelTypeOk returns a tuple with the VerifyChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyChannelType

`func (o *VerifyDetailRecord) SetVerifyChannelType(v string)`

SetVerifyChannelType sets VerifyChannelType field to given value.

### HasVerifyChannelType

`func (o *VerifyDetailRecord) HasVerifyChannelType() bool`

HasVerifyChannelType returns a boolean if a field has been set.

### GetVerifyChannelId

`func (o *VerifyDetailRecord) GetVerifyChannelId() string`

GetVerifyChannelId returns the VerifyChannelId field if non-nil, zero value otherwise.

### GetVerifyChannelIdOk

`func (o *VerifyDetailRecord) GetVerifyChannelIdOk() (*string, bool)`

GetVerifyChannelIdOk returns a tuple with the VerifyChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyChannelId

`func (o *VerifyDetailRecord) SetVerifyChannelId(v string)`

SetVerifyChannelId sets VerifyChannelId field to given value.

### HasVerifyChannelId

`func (o *VerifyDetailRecord) HasVerifyChannelId() bool`

HasVerifyChannelId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerifyDetailRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerifyDetailRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerifyDetailRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VerifyDetailRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VerifyDetailRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VerifyDetailRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VerifyDetailRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VerifyDetailRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRate

`func (o *VerifyDetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *VerifyDetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *VerifyDetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *VerifyDetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateMeasuredIn

`func (o *VerifyDetailRecord) GetRateMeasuredIn() string`

GetRateMeasuredIn returns the RateMeasuredIn field if non-nil, zero value otherwise.

### GetRateMeasuredInOk

`func (o *VerifyDetailRecord) GetRateMeasuredInOk() (*string, bool)`

GetRateMeasuredInOk returns a tuple with the RateMeasuredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMeasuredIn

`func (o *VerifyDetailRecord) SetRateMeasuredIn(v string)`

SetRateMeasuredIn sets RateMeasuredIn field to given value.

### HasRateMeasuredIn

`func (o *VerifyDetailRecord) HasRateMeasuredIn() bool`

HasRateMeasuredIn returns a boolean if a field has been set.

### GetVerifyUsageFee

`func (o *VerifyDetailRecord) GetVerifyUsageFee() string`

GetVerifyUsageFee returns the VerifyUsageFee field if non-nil, zero value otherwise.

### GetVerifyUsageFeeOk

`func (o *VerifyDetailRecord) GetVerifyUsageFeeOk() (*string, bool)`

GetVerifyUsageFeeOk returns a tuple with the VerifyUsageFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUsageFee

`func (o *VerifyDetailRecord) SetVerifyUsageFee(v string)`

SetVerifyUsageFee sets VerifyUsageFee field to given value.

### HasVerifyUsageFee

`func (o *VerifyDetailRecord) HasVerifyUsageFee() bool`

HasVerifyUsageFee returns a boolean if a field has been set.

### GetCurrency

`func (o *VerifyDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VerifyDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VerifyDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VerifyDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecordType

`func (o *VerifyDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VerifyDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VerifyDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


