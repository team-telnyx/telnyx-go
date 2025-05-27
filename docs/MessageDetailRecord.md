# MessageDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier of the message | [optional] 
**UserId** | Pointer to **string** | Identifier of the Telnyx account who owns the message | [optional] 
**CompletedAt** | Pointer to **time.Time** | Message completion time | [optional] 
**CreatedAt** | Pointer to **time.Time** | Message creation time | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Message updated time | [optional] 
**SentAt** | Pointer to **time.Time** | Time when the message was sent | [optional] 
**Carrier** | Pointer to **string** | Country-specific carrier used to send or receive the message | [optional] 
**CarrierFee** | Pointer to **string** | Fee charged by certain carriers in order to deliver certain message types. Telnyx passes this fee on to the customer according to our pricing table | [optional] 
**Cld** | Pointer to **string** | The recipient of the message (to parameter in the Messaging API) | [optional] 
**Cli** | Pointer to **string** | The sender of the message (from parameter in the Messaging API). For Alphanumeric ID messages, this is the sender ID value | [optional] 
**CountryCode** | Pointer to **string** | Two-letter representation of the country of the cld property using the ISO 3166-1 alpha-2 format | [optional] 
**DeliveryStatus** | Pointer to **string** | Final webhook delivery status | [optional] 
**DeliveryStatusFailoverUrl** | Pointer to **string** | Failover customer-provided URL which Telnyx posts delivery status webhooks to | [optional] 
**DeliveryStatusWebhookUrl** | Pointer to **string** | Primary customer-provided URL which Telnyx posts delivery status webhooks to | [optional] 
**Direction** | Pointer to **string** | Logical direction of the message from the Telnyx customer&#39;s perspective. It&#39;s inbound when the Telnyx customer receives the message, or outbound otherwise | [optional] 
**Fteu** | Pointer to **bool** | Indicates whether this is a Free-To-End-User (FTEU) short code message | [optional] 
**Mcc** | Pointer to **string** | Mobile country code. Only available for certain products, such as Global Outbound-Only from Alphanumeric Sender ID | [optional] 
**Mnc** | Pointer to **string** | Mobile network code. Only available for certain products, such as Global Outbound-Only from Alphanumeric Sender ID | [optional] 
**MessageType** | Pointer to **string** | Describes the Messaging service used to send the message. Available services are: Short Message Service (SMS), Multimedia Messaging Service (MMS), and Rich Communication Services (RCS) | [optional] 
**OnNet** | Pointer to **bool** | Indicates whether both sender and recipient numbers are Telnyx-managed | [optional] 
**ProfileId** | Pointer to **string** | Unique identifier of the Messaging Profile used to send or receive the message | [optional] 
**ProfileName** | Pointer to **string** | Name of the Messaging Profile used to send or receive the message | [optional] 
**SourceCountryCode** | Pointer to **string** | Two-letter representation of the country of the cli property using the ISO 3166-1 alpha-2 format | [optional] 
**Status** | Pointer to **string** | Final status of the message after the delivery attempt | [optional] 
**Tags** | Pointer to **string** | Comma-separated tags assigned to the Telnyx number associated with the message | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**Cost** | Pointer to **string** | Amount, in the user currency, for the Telnyx billing cost | [optional] 
**Errors** | Pointer to **[]string** | Telnyx API error codes returned by the Telnyx gateway | [optional] 
**Parts** | Pointer to **int32** | Number of message parts. The message is broken down in multiple parts when its length surpasses the limit of 160 characters | [optional] 
**RecordType** | **string** | Identifies the record schema | [default to "message_detail_record"]

## Methods

### NewMessageDetailRecord

`func NewMessageDetailRecord(recordType string, ) *MessageDetailRecord`

NewMessageDetailRecord instantiates a new MessageDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDetailRecordWithDefaults

`func NewMessageDetailRecordWithDefaults() *MessageDetailRecord`

NewMessageDetailRecordWithDefaults instantiates a new MessageDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MessageDetailRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MessageDetailRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MessageDetailRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MessageDetailRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUserId

`func (o *MessageDetailRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MessageDetailRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MessageDetailRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MessageDetailRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MessageDetailRecord) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MessageDetailRecord) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MessageDetailRecord) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MessageDetailRecord) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageDetailRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageDetailRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageDetailRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageDetailRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MessageDetailRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageDetailRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageDetailRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MessageDetailRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSentAt

`func (o *MessageDetailRecord) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *MessageDetailRecord) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *MessageDetailRecord) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *MessageDetailRecord) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetCarrier

`func (o *MessageDetailRecord) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *MessageDetailRecord) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *MessageDetailRecord) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *MessageDetailRecord) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierFee

`func (o *MessageDetailRecord) GetCarrierFee() string`

GetCarrierFee returns the CarrierFee field if non-nil, zero value otherwise.

### GetCarrierFeeOk

`func (o *MessageDetailRecord) GetCarrierFeeOk() (*string, bool)`

GetCarrierFeeOk returns a tuple with the CarrierFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierFee

`func (o *MessageDetailRecord) SetCarrierFee(v string)`

SetCarrierFee sets CarrierFee field to given value.

### HasCarrierFee

`func (o *MessageDetailRecord) HasCarrierFee() bool`

HasCarrierFee returns a boolean if a field has been set.

### GetCld

`func (o *MessageDetailRecord) GetCld() string`

GetCld returns the Cld field if non-nil, zero value otherwise.

### GetCldOk

`func (o *MessageDetailRecord) GetCldOk() (*string, bool)`

GetCldOk returns a tuple with the Cld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCld

`func (o *MessageDetailRecord) SetCld(v string)`

SetCld sets Cld field to given value.

### HasCld

`func (o *MessageDetailRecord) HasCld() bool`

HasCld returns a boolean if a field has been set.

### GetCli

`func (o *MessageDetailRecord) GetCli() string`

GetCli returns the Cli field if non-nil, zero value otherwise.

### GetCliOk

`func (o *MessageDetailRecord) GetCliOk() (*string, bool)`

GetCliOk returns a tuple with the Cli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCli

`func (o *MessageDetailRecord) SetCli(v string)`

SetCli sets Cli field to given value.

### HasCli

`func (o *MessageDetailRecord) HasCli() bool`

HasCli returns a boolean if a field has been set.

### GetCountryCode

`func (o *MessageDetailRecord) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *MessageDetailRecord) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *MessageDetailRecord) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *MessageDetailRecord) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *MessageDetailRecord) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *MessageDetailRecord) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *MessageDetailRecord) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *MessageDetailRecord) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetDeliveryStatusFailoverUrl

`func (o *MessageDetailRecord) GetDeliveryStatusFailoverUrl() string`

GetDeliveryStatusFailoverUrl returns the DeliveryStatusFailoverUrl field if non-nil, zero value otherwise.

### GetDeliveryStatusFailoverUrlOk

`func (o *MessageDetailRecord) GetDeliveryStatusFailoverUrlOk() (*string, bool)`

GetDeliveryStatusFailoverUrlOk returns a tuple with the DeliveryStatusFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatusFailoverUrl

`func (o *MessageDetailRecord) SetDeliveryStatusFailoverUrl(v string)`

SetDeliveryStatusFailoverUrl sets DeliveryStatusFailoverUrl field to given value.

### HasDeliveryStatusFailoverUrl

`func (o *MessageDetailRecord) HasDeliveryStatusFailoverUrl() bool`

HasDeliveryStatusFailoverUrl returns a boolean if a field has been set.

### GetDeliveryStatusWebhookUrl

`func (o *MessageDetailRecord) GetDeliveryStatusWebhookUrl() string`

GetDeliveryStatusWebhookUrl returns the DeliveryStatusWebhookUrl field if non-nil, zero value otherwise.

### GetDeliveryStatusWebhookUrlOk

`func (o *MessageDetailRecord) GetDeliveryStatusWebhookUrlOk() (*string, bool)`

GetDeliveryStatusWebhookUrlOk returns a tuple with the DeliveryStatusWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatusWebhookUrl

`func (o *MessageDetailRecord) SetDeliveryStatusWebhookUrl(v string)`

SetDeliveryStatusWebhookUrl sets DeliveryStatusWebhookUrl field to given value.

### HasDeliveryStatusWebhookUrl

`func (o *MessageDetailRecord) HasDeliveryStatusWebhookUrl() bool`

HasDeliveryStatusWebhookUrl returns a boolean if a field has been set.

### GetDirection

`func (o *MessageDetailRecord) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MessageDetailRecord) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MessageDetailRecord) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MessageDetailRecord) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFteu

`func (o *MessageDetailRecord) GetFteu() bool`

GetFteu returns the Fteu field if non-nil, zero value otherwise.

### GetFteuOk

`func (o *MessageDetailRecord) GetFteuOk() (*bool, bool)`

GetFteuOk returns a tuple with the Fteu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFteu

`func (o *MessageDetailRecord) SetFteu(v bool)`

SetFteu sets Fteu field to given value.

### HasFteu

`func (o *MessageDetailRecord) HasFteu() bool`

HasFteu returns a boolean if a field has been set.

### GetMcc

`func (o *MessageDetailRecord) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *MessageDetailRecord) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *MessageDetailRecord) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *MessageDetailRecord) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *MessageDetailRecord) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *MessageDetailRecord) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *MessageDetailRecord) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *MessageDetailRecord) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetMessageType

`func (o *MessageDetailRecord) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *MessageDetailRecord) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *MessageDetailRecord) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *MessageDetailRecord) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetOnNet

`func (o *MessageDetailRecord) GetOnNet() bool`

GetOnNet returns the OnNet field if non-nil, zero value otherwise.

### GetOnNetOk

`func (o *MessageDetailRecord) GetOnNetOk() (*bool, bool)`

GetOnNetOk returns a tuple with the OnNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnNet

`func (o *MessageDetailRecord) SetOnNet(v bool)`

SetOnNet sets OnNet field to given value.

### HasOnNet

`func (o *MessageDetailRecord) HasOnNet() bool`

HasOnNet returns a boolean if a field has been set.

### GetProfileId

`func (o *MessageDetailRecord) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MessageDetailRecord) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MessageDetailRecord) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MessageDetailRecord) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *MessageDetailRecord) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *MessageDetailRecord) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *MessageDetailRecord) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *MessageDetailRecord) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetSourceCountryCode

`func (o *MessageDetailRecord) GetSourceCountryCode() string`

GetSourceCountryCode returns the SourceCountryCode field if non-nil, zero value otherwise.

### GetSourceCountryCodeOk

`func (o *MessageDetailRecord) GetSourceCountryCodeOk() (*string, bool)`

GetSourceCountryCodeOk returns a tuple with the SourceCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCountryCode

`func (o *MessageDetailRecord) SetSourceCountryCode(v string)`

SetSourceCountryCode sets SourceCountryCode field to given value.

### HasSourceCountryCode

`func (o *MessageDetailRecord) HasSourceCountryCode() bool`

HasSourceCountryCode returns a boolean if a field has been set.

### GetStatus

`func (o *MessageDetailRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageDetailRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageDetailRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageDetailRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MessageDetailRecord) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MessageDetailRecord) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MessageDetailRecord) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MessageDetailRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRate

`func (o *MessageDetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MessageDetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MessageDetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MessageDetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCurrency

`func (o *MessageDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MessageDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MessageDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MessageDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCost

`func (o *MessageDetailRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MessageDetailRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MessageDetailRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MessageDetailRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetErrors

`func (o *MessageDetailRecord) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MessageDetailRecord) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MessageDetailRecord) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MessageDetailRecord) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetParts

`func (o *MessageDetailRecord) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *MessageDetailRecord) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *MessageDetailRecord) SetParts(v int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *MessageDetailRecord) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetRecordType

`func (o *MessageDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MessageDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MessageDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


