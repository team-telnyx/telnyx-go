# DetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier of the message | [optional] 
**UserId** | Pointer to **string** | User id | [optional] 
**CompletedAt** | Pointer to **time.Time** | Message completion time | [optional] 
**CreatedAt** | Pointer to **time.Time** | Event creation time | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**SentAt** | Pointer to **time.Time** | Time when the message was sent | [optional] 
**Carrier** | Pointer to **string** | Country-specific carrier used to send or receive the message | [optional] 
**CarrierFee** | Pointer to **string** | Fee charged by certain carriers in order to deliver certain message types. Telnyx passes this fee on to the customer according to our pricing table | [optional] 
**Cld** | Pointer to **string** | The recipient of the message (to parameter in the Messaging API) | [optional] 
**Cli** | Pointer to **string** | The sender of the message (from parameter in the Messaging API). For Alphanumeric ID messages, this is the sender ID value | [optional] 
**CountryCode** | Pointer to **string** | Two-letter representation of the country of the cld property using the ISO 3166-1 alpha-2 format | [optional] 
**DeliveryStatus** | Pointer to **string** |  | [optional] 
**DeliveryStatusFailoverUrl** | Pointer to **string** | Failover customer-provided URL which Telnyx posts delivery status webhooks to | [optional] 
**DeliveryStatusWebhookUrl** | Pointer to **string** | Primary customer-provided URL which Telnyx posts delivery status webhooks to | [optional] 
**Direction** | Pointer to **string** | Logical direction of the message from the Telnyx customer&#39;s perspective. It&#39;s inbound when the Telnyx customer receives the message, or outbound otherwise | [optional] 
**Fteu** | Pointer to **bool** | Indicates whether this is a Free-To-End-User (FTEU) short code message | [optional] 
**Mcc** | Pointer to **string** | Mobile country code | [optional] 
**Mnc** | Pointer to **string** | Mobile network code | [optional] 
**MessageType** | Pointer to **string** | Describes the Messaging service used to send the message. Available services are: Short Message Service (SMS), Multimedia Messaging Service (MMS), and Rich Communication Services (RCS) | [optional] 
**OnNet** | Pointer to **bool** | Indicates whether both sender and recipient numbers are Telnyx-managed | [optional] 
**ProfileId** | Pointer to **string** | Unique identifier of the Messaging Profile used to send or receive the message | [optional] 
**ProfileName** | Pointer to **string** | Name of the Messaging Profile used to send or receive the message | [optional] 
**SourceCountryCode** | Pointer to **string** | Two-letter representation of the country of the cli property using the ISO 3166-1 alpha-2 format | [optional] 
**Status** | Pointer to **string** | Request status | [optional] 
**Tags** | Pointer to **string** | User-provided tags | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**Cost** | Pointer to **string** | Currency amount for Telnyx billing cost | [optional] 
**Errors** | Pointer to **[]string** | Telnyx API error codes returned by the Telnyx gateway | [optional] 
**Parts** | Pointer to **int32** | Number of message parts. The message is broken down in multiple parts when its length surpasses the limit of 160 characters | [optional] 
**RecordType** | **string** |  | [default to "media_storage"]
**Id** | Pointer to **string** | Unique identifier for the Media Storage Event | [optional] 
**Name** | Pointer to **string** | Conference name | [optional] 
**StartedAt** | Pointer to **time.Time** | Conference start time | [optional] 
**EndedAt** | Pointer to **time.Time** | Conference end time | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Conference expiry time | [optional] 
**Region** | Pointer to **string** | Region where the conference is hosted | [optional] 
**CallLegId** | Pointer to **string** | Telnyx UUID that identifies the related call leg | [optional] 
**CallSessionId** | Pointer to **string** | Telnyx UUID that identifies the related call session | [optional] 
**ConnectionId** | Pointer to **string** | Connection id | [optional] 
**CallSec** | Pointer to **int32** | Duration of the conference call in seconds | [optional] 
**ParticipantCount** | Pointer to **int32** | Number of participants that joined the conference call | [optional] 
**ParticipantCallSec** | Pointer to **int32** | Sum of the conference call duration for all participants in seconds | [optional] 
**IsTelnyxBillable** | Pointer to **bool** | Indicates whether Telnyx billing charges might be applicable | [optional] 
**ConferenceId** | Pointer to **string** | Conference id | [optional] 
**JoinedAt** | Pointer to **time.Time** | Participant join time | [optional] 
**LeftAt** | Pointer to **time.Time** | Participant leave time | [optional] 
**DestinationNumber** | Pointer to **string** | Number called by the participant to join the conference | [optional] 
**OriginatingNumber** | Pointer to **string** | Participant origin number used in the conference call | [optional] 
**BilledSec** | Pointer to **int32** | Duration of the conference call for billing purposes | [optional] 
**RateMeasuredIn** | Pointer to **string** | Billing unit used to calculate the Telnyx billing cost | [optional] 
**InvokedAt** | Pointer to **time.Time** | Feature invocation time | [optional] 
**Feature** | Pointer to **string** | Feature name | [optional] 
**BillingGroupId** | Pointer to **string** | Billing Group id | [optional] 
**BillingGroupName** | Pointer to **string** | Name of the Billing Group specified in billing_group_id | [optional] 
**ConnectionName** | Pointer to **string** | Connection name | [optional] 
**VerifyProfileId** | Pointer to **string** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**DestinationPhoneNumber** | Pointer to **string** | E.164 formatted phone number | [optional] 
**VerifyChannelType** | Pointer to **string** | Depending on the type of verification, the &#x60;verify_channel_id&#x60; points to one of the following channel ids; --- verify_channel_type | verify_channel_id ------------------- | ----------------- sms, psd2           | messaging_id call, flashcall     | call_control_id ---  | [optional] 
**VerifyChannelId** | Pointer to **string** |  | [optional] 
**VerifyUsageFee** | Pointer to **string** | Currency amount for Verify Usage Fee | [optional] 
**ClosedAt** | Pointer to **time.Time** | Event close time | [optional] 
**IpAddress** | Pointer to **string** | Ip address that generated the event | [optional] 
**DownlinkData** | Pointer to **float32** | Number of megabytes downloaded | [optional] 
**Imsi** | Pointer to **string** | International Mobile Subscriber Identity | [optional] 
**DataUnit** | Pointer to **string** | Unit of wireless link consumption | [optional] 
**DataRate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**SimGroupName** | Pointer to **string** | Sim group name for sim card | [optional] 
**SimCardId** | Pointer to **string** | Unique identifier for SIM card | [optional] 
**SimGroupId** | Pointer to **string** | Unique identifier for SIM group | [optional] 
**SimCardTags** | Pointer to **string** | User-provided tags | [optional] 
**PhoneNumber** | Pointer to **string** | Telephone number associated to SIM card | [optional] 
**UplinkData** | Pointer to **float32** | Number of megabytes uploaded | [optional] 
**DataCost** | Pointer to **float32** | Data cost | [optional] 
**AssetId** | Pointer to **string** | Asset id | [optional] 
**OrgId** | Pointer to **string** | Organization owner id | [optional] 
**ActionType** | Pointer to **string** | Type of action performed against the Media Storage API | [optional] 
**LinkChannelType** | Pointer to **string** | Link channel type | [optional] 
**LinkChannelId** | Pointer to **string** | Link channel id | [optional] 
**WebhookId** | Pointer to **string** | Webhook id | [optional] 

## Methods

### NewDetailRecord

`func NewDetailRecord(recordType string, ) *DetailRecord`

NewDetailRecord instantiates a new DetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailRecordWithDefaults

`func NewDetailRecordWithDefaults() *DetailRecord`

NewDetailRecordWithDefaults instantiates a new DetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DetailRecord) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DetailRecord) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DetailRecord) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DetailRecord) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUserId

`func (o *DetailRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DetailRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DetailRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DetailRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetCompletedAt

`func (o *DetailRecord) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DetailRecord) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DetailRecord) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DetailRecord) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DetailRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DetailRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DetailRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DetailRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DetailRecord) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DetailRecord) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DetailRecord) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DetailRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSentAt

`func (o *DetailRecord) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *DetailRecord) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *DetailRecord) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *DetailRecord) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetCarrier

`func (o *DetailRecord) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *DetailRecord) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *DetailRecord) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *DetailRecord) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetCarrierFee

`func (o *DetailRecord) GetCarrierFee() string`

GetCarrierFee returns the CarrierFee field if non-nil, zero value otherwise.

### GetCarrierFeeOk

`func (o *DetailRecord) GetCarrierFeeOk() (*string, bool)`

GetCarrierFeeOk returns a tuple with the CarrierFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierFee

`func (o *DetailRecord) SetCarrierFee(v string)`

SetCarrierFee sets CarrierFee field to given value.

### HasCarrierFee

`func (o *DetailRecord) HasCarrierFee() bool`

HasCarrierFee returns a boolean if a field has been set.

### GetCld

`func (o *DetailRecord) GetCld() string`

GetCld returns the Cld field if non-nil, zero value otherwise.

### GetCldOk

`func (o *DetailRecord) GetCldOk() (*string, bool)`

GetCldOk returns a tuple with the Cld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCld

`func (o *DetailRecord) SetCld(v string)`

SetCld sets Cld field to given value.

### HasCld

`func (o *DetailRecord) HasCld() bool`

HasCld returns a boolean if a field has been set.

### GetCli

`func (o *DetailRecord) GetCli() string`

GetCli returns the Cli field if non-nil, zero value otherwise.

### GetCliOk

`func (o *DetailRecord) GetCliOk() (*string, bool)`

GetCliOk returns a tuple with the Cli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCli

`func (o *DetailRecord) SetCli(v string)`

SetCli sets Cli field to given value.

### HasCli

`func (o *DetailRecord) HasCli() bool`

HasCli returns a boolean if a field has been set.

### GetCountryCode

`func (o *DetailRecord) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DetailRecord) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DetailRecord) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DetailRecord) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *DetailRecord) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *DetailRecord) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *DetailRecord) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *DetailRecord) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetDeliveryStatusFailoverUrl

`func (o *DetailRecord) GetDeliveryStatusFailoverUrl() string`

GetDeliveryStatusFailoverUrl returns the DeliveryStatusFailoverUrl field if non-nil, zero value otherwise.

### GetDeliveryStatusFailoverUrlOk

`func (o *DetailRecord) GetDeliveryStatusFailoverUrlOk() (*string, bool)`

GetDeliveryStatusFailoverUrlOk returns a tuple with the DeliveryStatusFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatusFailoverUrl

`func (o *DetailRecord) SetDeliveryStatusFailoverUrl(v string)`

SetDeliveryStatusFailoverUrl sets DeliveryStatusFailoverUrl field to given value.

### HasDeliveryStatusFailoverUrl

`func (o *DetailRecord) HasDeliveryStatusFailoverUrl() bool`

HasDeliveryStatusFailoverUrl returns a boolean if a field has been set.

### GetDeliveryStatusWebhookUrl

`func (o *DetailRecord) GetDeliveryStatusWebhookUrl() string`

GetDeliveryStatusWebhookUrl returns the DeliveryStatusWebhookUrl field if non-nil, zero value otherwise.

### GetDeliveryStatusWebhookUrlOk

`func (o *DetailRecord) GetDeliveryStatusWebhookUrlOk() (*string, bool)`

GetDeliveryStatusWebhookUrlOk returns a tuple with the DeliveryStatusWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatusWebhookUrl

`func (o *DetailRecord) SetDeliveryStatusWebhookUrl(v string)`

SetDeliveryStatusWebhookUrl sets DeliveryStatusWebhookUrl field to given value.

### HasDeliveryStatusWebhookUrl

`func (o *DetailRecord) HasDeliveryStatusWebhookUrl() bool`

HasDeliveryStatusWebhookUrl returns a boolean if a field has been set.

### GetDirection

`func (o *DetailRecord) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *DetailRecord) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *DetailRecord) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *DetailRecord) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFteu

`func (o *DetailRecord) GetFteu() bool`

GetFteu returns the Fteu field if non-nil, zero value otherwise.

### GetFteuOk

`func (o *DetailRecord) GetFteuOk() (*bool, bool)`

GetFteuOk returns a tuple with the Fteu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFteu

`func (o *DetailRecord) SetFteu(v bool)`

SetFteu sets Fteu field to given value.

### HasFteu

`func (o *DetailRecord) HasFteu() bool`

HasFteu returns a boolean if a field has been set.

### GetMcc

`func (o *DetailRecord) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *DetailRecord) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *DetailRecord) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *DetailRecord) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMnc

`func (o *DetailRecord) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *DetailRecord) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *DetailRecord) SetMnc(v string)`

SetMnc sets Mnc field to given value.

### HasMnc

`func (o *DetailRecord) HasMnc() bool`

HasMnc returns a boolean if a field has been set.

### GetMessageType

`func (o *DetailRecord) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *DetailRecord) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *DetailRecord) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *DetailRecord) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetOnNet

`func (o *DetailRecord) GetOnNet() bool`

GetOnNet returns the OnNet field if non-nil, zero value otherwise.

### GetOnNetOk

`func (o *DetailRecord) GetOnNetOk() (*bool, bool)`

GetOnNetOk returns a tuple with the OnNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnNet

`func (o *DetailRecord) SetOnNet(v bool)`

SetOnNet sets OnNet field to given value.

### HasOnNet

`func (o *DetailRecord) HasOnNet() bool`

HasOnNet returns a boolean if a field has been set.

### GetProfileId

`func (o *DetailRecord) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *DetailRecord) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *DetailRecord) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *DetailRecord) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *DetailRecord) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *DetailRecord) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *DetailRecord) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *DetailRecord) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetSourceCountryCode

`func (o *DetailRecord) GetSourceCountryCode() string`

GetSourceCountryCode returns the SourceCountryCode field if non-nil, zero value otherwise.

### GetSourceCountryCodeOk

`func (o *DetailRecord) GetSourceCountryCodeOk() (*string, bool)`

GetSourceCountryCodeOk returns a tuple with the SourceCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCountryCode

`func (o *DetailRecord) SetSourceCountryCode(v string)`

SetSourceCountryCode sets SourceCountryCode field to given value.

### HasSourceCountryCode

`func (o *DetailRecord) HasSourceCountryCode() bool`

HasSourceCountryCode returns a boolean if a field has been set.

### GetStatus

`func (o *DetailRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetailRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetailRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetailRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *DetailRecord) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DetailRecord) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DetailRecord) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DetailRecord) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRate

`func (o *DetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *DetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *DetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *DetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCurrency

`func (o *DetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCost

`func (o *DetailRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *DetailRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *DetailRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *DetailRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetErrors

`func (o *DetailRecord) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DetailRecord) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DetailRecord) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DetailRecord) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetParts

`func (o *DetailRecord) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *DetailRecord) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *DetailRecord) SetParts(v int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *DetailRecord) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetRecordType

`func (o *DetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *DetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DetailRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartedAt

`func (o *DetailRecord) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DetailRecord) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DetailRecord) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DetailRecord) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *DetailRecord) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *DetailRecord) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *DetailRecord) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *DetailRecord) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DetailRecord) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DetailRecord) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DetailRecord) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DetailRecord) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRegion

`func (o *DetailRecord) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DetailRecord) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DetailRecord) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DetailRecord) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCallLegId

`func (o *DetailRecord) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *DetailRecord) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *DetailRecord) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *DetailRecord) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *DetailRecord) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *DetailRecord) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *DetailRecord) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *DetailRecord) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetConnectionId

`func (o *DetailRecord) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DetailRecord) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DetailRecord) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *DetailRecord) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallSec

`func (o *DetailRecord) GetCallSec() int32`

GetCallSec returns the CallSec field if non-nil, zero value otherwise.

### GetCallSecOk

`func (o *DetailRecord) GetCallSecOk() (*int32, bool)`

GetCallSecOk returns a tuple with the CallSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSec

`func (o *DetailRecord) SetCallSec(v int32)`

SetCallSec sets CallSec field to given value.

### HasCallSec

`func (o *DetailRecord) HasCallSec() bool`

HasCallSec returns a boolean if a field has been set.

### GetParticipantCount

`func (o *DetailRecord) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *DetailRecord) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *DetailRecord) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *DetailRecord) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetParticipantCallSec

`func (o *DetailRecord) GetParticipantCallSec() int32`

GetParticipantCallSec returns the ParticipantCallSec field if non-nil, zero value otherwise.

### GetParticipantCallSecOk

`func (o *DetailRecord) GetParticipantCallSecOk() (*int32, bool)`

GetParticipantCallSecOk returns a tuple with the ParticipantCallSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCallSec

`func (o *DetailRecord) SetParticipantCallSec(v int32)`

SetParticipantCallSec sets ParticipantCallSec field to given value.

### HasParticipantCallSec

`func (o *DetailRecord) HasParticipantCallSec() bool`

HasParticipantCallSec returns a boolean if a field has been set.

### GetIsTelnyxBillable

`func (o *DetailRecord) GetIsTelnyxBillable() bool`

GetIsTelnyxBillable returns the IsTelnyxBillable field if non-nil, zero value otherwise.

### GetIsTelnyxBillableOk

`func (o *DetailRecord) GetIsTelnyxBillableOk() (*bool, bool)`

GetIsTelnyxBillableOk returns a tuple with the IsTelnyxBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTelnyxBillable

`func (o *DetailRecord) SetIsTelnyxBillable(v bool)`

SetIsTelnyxBillable sets IsTelnyxBillable field to given value.

### HasIsTelnyxBillable

`func (o *DetailRecord) HasIsTelnyxBillable() bool`

HasIsTelnyxBillable returns a boolean if a field has been set.

### GetConferenceId

`func (o *DetailRecord) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *DetailRecord) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *DetailRecord) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *DetailRecord) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetJoinedAt

`func (o *DetailRecord) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *DetailRecord) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *DetailRecord) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *DetailRecord) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLeftAt

`func (o *DetailRecord) GetLeftAt() time.Time`

GetLeftAt returns the LeftAt field if non-nil, zero value otherwise.

### GetLeftAtOk

`func (o *DetailRecord) GetLeftAtOk() (*time.Time, bool)`

GetLeftAtOk returns a tuple with the LeftAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAt

`func (o *DetailRecord) SetLeftAt(v time.Time)`

SetLeftAt sets LeftAt field to given value.

### HasLeftAt

`func (o *DetailRecord) HasLeftAt() bool`

HasLeftAt returns a boolean if a field has been set.

### GetDestinationNumber

`func (o *DetailRecord) GetDestinationNumber() string`

GetDestinationNumber returns the DestinationNumber field if non-nil, zero value otherwise.

### GetDestinationNumberOk

`func (o *DetailRecord) GetDestinationNumberOk() (*string, bool)`

GetDestinationNumberOk returns a tuple with the DestinationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNumber

`func (o *DetailRecord) SetDestinationNumber(v string)`

SetDestinationNumber sets DestinationNumber field to given value.

### HasDestinationNumber

`func (o *DetailRecord) HasDestinationNumber() bool`

HasDestinationNumber returns a boolean if a field has been set.

### GetOriginatingNumber

`func (o *DetailRecord) GetOriginatingNumber() string`

GetOriginatingNumber returns the OriginatingNumber field if non-nil, zero value otherwise.

### GetOriginatingNumberOk

`func (o *DetailRecord) GetOriginatingNumberOk() (*string, bool)`

GetOriginatingNumberOk returns a tuple with the OriginatingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingNumber

`func (o *DetailRecord) SetOriginatingNumber(v string)`

SetOriginatingNumber sets OriginatingNumber field to given value.

### HasOriginatingNumber

`func (o *DetailRecord) HasOriginatingNumber() bool`

HasOriginatingNumber returns a boolean if a field has been set.

### GetBilledSec

`func (o *DetailRecord) GetBilledSec() int32`

GetBilledSec returns the BilledSec field if non-nil, zero value otherwise.

### GetBilledSecOk

`func (o *DetailRecord) GetBilledSecOk() (*int32, bool)`

GetBilledSecOk returns a tuple with the BilledSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledSec

`func (o *DetailRecord) SetBilledSec(v int32)`

SetBilledSec sets BilledSec field to given value.

### HasBilledSec

`func (o *DetailRecord) HasBilledSec() bool`

HasBilledSec returns a boolean if a field has been set.

### GetRateMeasuredIn

`func (o *DetailRecord) GetRateMeasuredIn() string`

GetRateMeasuredIn returns the RateMeasuredIn field if non-nil, zero value otherwise.

### GetRateMeasuredInOk

`func (o *DetailRecord) GetRateMeasuredInOk() (*string, bool)`

GetRateMeasuredInOk returns a tuple with the RateMeasuredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMeasuredIn

`func (o *DetailRecord) SetRateMeasuredIn(v string)`

SetRateMeasuredIn sets RateMeasuredIn field to given value.

### HasRateMeasuredIn

`func (o *DetailRecord) HasRateMeasuredIn() bool`

HasRateMeasuredIn returns a boolean if a field has been set.

### GetInvokedAt

`func (o *DetailRecord) GetInvokedAt() time.Time`

GetInvokedAt returns the InvokedAt field if non-nil, zero value otherwise.

### GetInvokedAtOk

`func (o *DetailRecord) GetInvokedAtOk() (*time.Time, bool)`

GetInvokedAtOk returns a tuple with the InvokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokedAt

`func (o *DetailRecord) SetInvokedAt(v time.Time)`

SetInvokedAt sets InvokedAt field to given value.

### HasInvokedAt

`func (o *DetailRecord) HasInvokedAt() bool`

HasInvokedAt returns a boolean if a field has been set.

### GetFeature

`func (o *DetailRecord) GetFeature() string`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DetailRecord) GetFeatureOk() (*string, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DetailRecord) SetFeature(v string)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DetailRecord) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *DetailRecord) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *DetailRecord) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *DetailRecord) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *DetailRecord) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetBillingGroupName

`func (o *DetailRecord) GetBillingGroupName() string`

GetBillingGroupName returns the BillingGroupName field if non-nil, zero value otherwise.

### GetBillingGroupNameOk

`func (o *DetailRecord) GetBillingGroupNameOk() (*string, bool)`

GetBillingGroupNameOk returns a tuple with the BillingGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupName

`func (o *DetailRecord) SetBillingGroupName(v string)`

SetBillingGroupName sets BillingGroupName field to given value.

### HasBillingGroupName

`func (o *DetailRecord) HasBillingGroupName() bool`

HasBillingGroupName returns a boolean if a field has been set.

### GetConnectionName

`func (o *DetailRecord) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *DetailRecord) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *DetailRecord) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *DetailRecord) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetVerifyProfileId

`func (o *DetailRecord) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *DetailRecord) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *DetailRecord) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.

### HasVerifyProfileId

`func (o *DetailRecord) HasVerifyProfileId() bool`

HasVerifyProfileId returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *DetailRecord) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *DetailRecord) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *DetailRecord) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *DetailRecord) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetDestinationPhoneNumber

`func (o *DetailRecord) GetDestinationPhoneNumber() string`

GetDestinationPhoneNumber returns the DestinationPhoneNumber field if non-nil, zero value otherwise.

### GetDestinationPhoneNumberOk

`func (o *DetailRecord) GetDestinationPhoneNumberOk() (*string, bool)`

GetDestinationPhoneNumberOk returns a tuple with the DestinationPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPhoneNumber

`func (o *DetailRecord) SetDestinationPhoneNumber(v string)`

SetDestinationPhoneNumber sets DestinationPhoneNumber field to given value.

### HasDestinationPhoneNumber

`func (o *DetailRecord) HasDestinationPhoneNumber() bool`

HasDestinationPhoneNumber returns a boolean if a field has been set.

### GetVerifyChannelType

`func (o *DetailRecord) GetVerifyChannelType() string`

GetVerifyChannelType returns the VerifyChannelType field if non-nil, zero value otherwise.

### GetVerifyChannelTypeOk

`func (o *DetailRecord) GetVerifyChannelTypeOk() (*string, bool)`

GetVerifyChannelTypeOk returns a tuple with the VerifyChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyChannelType

`func (o *DetailRecord) SetVerifyChannelType(v string)`

SetVerifyChannelType sets VerifyChannelType field to given value.

### HasVerifyChannelType

`func (o *DetailRecord) HasVerifyChannelType() bool`

HasVerifyChannelType returns a boolean if a field has been set.

### GetVerifyChannelId

`func (o *DetailRecord) GetVerifyChannelId() string`

GetVerifyChannelId returns the VerifyChannelId field if non-nil, zero value otherwise.

### GetVerifyChannelIdOk

`func (o *DetailRecord) GetVerifyChannelIdOk() (*string, bool)`

GetVerifyChannelIdOk returns a tuple with the VerifyChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyChannelId

`func (o *DetailRecord) SetVerifyChannelId(v string)`

SetVerifyChannelId sets VerifyChannelId field to given value.

### HasVerifyChannelId

`func (o *DetailRecord) HasVerifyChannelId() bool`

HasVerifyChannelId returns a boolean if a field has been set.

### GetVerifyUsageFee

`func (o *DetailRecord) GetVerifyUsageFee() string`

GetVerifyUsageFee returns the VerifyUsageFee field if non-nil, zero value otherwise.

### GetVerifyUsageFeeOk

`func (o *DetailRecord) GetVerifyUsageFeeOk() (*string, bool)`

GetVerifyUsageFeeOk returns a tuple with the VerifyUsageFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUsageFee

`func (o *DetailRecord) SetVerifyUsageFee(v string)`

SetVerifyUsageFee sets VerifyUsageFee field to given value.

### HasVerifyUsageFee

`func (o *DetailRecord) HasVerifyUsageFee() bool`

HasVerifyUsageFee returns a boolean if a field has been set.

### GetClosedAt

`func (o *DetailRecord) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *DetailRecord) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *DetailRecord) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *DetailRecord) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetIpAddress

`func (o *DetailRecord) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DetailRecord) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DetailRecord) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DetailRecord) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDownlinkData

`func (o *DetailRecord) GetDownlinkData() float32`

GetDownlinkData returns the DownlinkData field if non-nil, zero value otherwise.

### GetDownlinkDataOk

`func (o *DetailRecord) GetDownlinkDataOk() (*float32, bool)`

GetDownlinkDataOk returns a tuple with the DownlinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkData

`func (o *DetailRecord) SetDownlinkData(v float32)`

SetDownlinkData sets DownlinkData field to given value.

### HasDownlinkData

`func (o *DetailRecord) HasDownlinkData() bool`

HasDownlinkData returns a boolean if a field has been set.

### GetImsi

`func (o *DetailRecord) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *DetailRecord) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *DetailRecord) SetImsi(v string)`

SetImsi sets Imsi field to given value.

### HasImsi

`func (o *DetailRecord) HasImsi() bool`

HasImsi returns a boolean if a field has been set.

### GetDataUnit

`func (o *DetailRecord) GetDataUnit() string`

GetDataUnit returns the DataUnit field if non-nil, zero value otherwise.

### GetDataUnitOk

`func (o *DetailRecord) GetDataUnitOk() (*string, bool)`

GetDataUnitOk returns a tuple with the DataUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUnit

`func (o *DetailRecord) SetDataUnit(v string)`

SetDataUnit sets DataUnit field to given value.

### HasDataUnit

`func (o *DetailRecord) HasDataUnit() bool`

HasDataUnit returns a boolean if a field has been set.

### GetDataRate

`func (o *DetailRecord) GetDataRate() string`

GetDataRate returns the DataRate field if non-nil, zero value otherwise.

### GetDataRateOk

`func (o *DetailRecord) GetDataRateOk() (*string, bool)`

GetDataRateOk returns a tuple with the DataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRate

`func (o *DetailRecord) SetDataRate(v string)`

SetDataRate sets DataRate field to given value.

### HasDataRate

`func (o *DetailRecord) HasDataRate() bool`

HasDataRate returns a boolean if a field has been set.

### GetSimGroupName

`func (o *DetailRecord) GetSimGroupName() string`

GetSimGroupName returns the SimGroupName field if non-nil, zero value otherwise.

### GetSimGroupNameOk

`func (o *DetailRecord) GetSimGroupNameOk() (*string, bool)`

GetSimGroupNameOk returns a tuple with the SimGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupName

`func (o *DetailRecord) SetSimGroupName(v string)`

SetSimGroupName sets SimGroupName field to given value.

### HasSimGroupName

`func (o *DetailRecord) HasSimGroupName() bool`

HasSimGroupName returns a boolean if a field has been set.

### GetSimCardId

`func (o *DetailRecord) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *DetailRecord) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *DetailRecord) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *DetailRecord) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetSimGroupId

`func (o *DetailRecord) GetSimGroupId() string`

GetSimGroupId returns the SimGroupId field if non-nil, zero value otherwise.

### GetSimGroupIdOk

`func (o *DetailRecord) GetSimGroupIdOk() (*string, bool)`

GetSimGroupIdOk returns a tuple with the SimGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimGroupId

`func (o *DetailRecord) SetSimGroupId(v string)`

SetSimGroupId sets SimGroupId field to given value.

### HasSimGroupId

`func (o *DetailRecord) HasSimGroupId() bool`

HasSimGroupId returns a boolean if a field has been set.

### GetSimCardTags

`func (o *DetailRecord) GetSimCardTags() string`

GetSimCardTags returns the SimCardTags field if non-nil, zero value otherwise.

### GetSimCardTagsOk

`func (o *DetailRecord) GetSimCardTagsOk() (*string, bool)`

GetSimCardTagsOk returns a tuple with the SimCardTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardTags

`func (o *DetailRecord) SetSimCardTags(v string)`

SetSimCardTags sets SimCardTags field to given value.

### HasSimCardTags

`func (o *DetailRecord) HasSimCardTags() bool`

HasSimCardTags returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *DetailRecord) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *DetailRecord) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *DetailRecord) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *DetailRecord) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUplinkData

`func (o *DetailRecord) GetUplinkData() float32`

GetUplinkData returns the UplinkData field if non-nil, zero value otherwise.

### GetUplinkDataOk

`func (o *DetailRecord) GetUplinkDataOk() (*float32, bool)`

GetUplinkDataOk returns a tuple with the UplinkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkData

`func (o *DetailRecord) SetUplinkData(v float32)`

SetUplinkData sets UplinkData field to given value.

### HasUplinkData

`func (o *DetailRecord) HasUplinkData() bool`

HasUplinkData returns a boolean if a field has been set.

### GetDataCost

`func (o *DetailRecord) GetDataCost() float32`

GetDataCost returns the DataCost field if non-nil, zero value otherwise.

### GetDataCostOk

`func (o *DetailRecord) GetDataCostOk() (*float32, bool)`

GetDataCostOk returns a tuple with the DataCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCost

`func (o *DetailRecord) SetDataCost(v float32)`

SetDataCost sets DataCost field to given value.

### HasDataCost

`func (o *DetailRecord) HasDataCost() bool`

HasDataCost returns a boolean if a field has been set.

### GetAssetId

`func (o *DetailRecord) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *DetailRecord) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *DetailRecord) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *DetailRecord) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetOrgId

`func (o *DetailRecord) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DetailRecord) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DetailRecord) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DetailRecord) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetActionType

`func (o *DetailRecord) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *DetailRecord) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *DetailRecord) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *DetailRecord) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetLinkChannelType

`func (o *DetailRecord) GetLinkChannelType() string`

GetLinkChannelType returns the LinkChannelType field if non-nil, zero value otherwise.

### GetLinkChannelTypeOk

`func (o *DetailRecord) GetLinkChannelTypeOk() (*string, bool)`

GetLinkChannelTypeOk returns a tuple with the LinkChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkChannelType

`func (o *DetailRecord) SetLinkChannelType(v string)`

SetLinkChannelType sets LinkChannelType field to given value.

### HasLinkChannelType

`func (o *DetailRecord) HasLinkChannelType() bool`

HasLinkChannelType returns a boolean if a field has been set.

### GetLinkChannelId

`func (o *DetailRecord) GetLinkChannelId() string`

GetLinkChannelId returns the LinkChannelId field if non-nil, zero value otherwise.

### GetLinkChannelIdOk

`func (o *DetailRecord) GetLinkChannelIdOk() (*string, bool)`

GetLinkChannelIdOk returns a tuple with the LinkChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkChannelId

`func (o *DetailRecord) SetLinkChannelId(v string)`

SetLinkChannelId sets LinkChannelId field to given value.

### HasLinkChannelId

`func (o *DetailRecord) HasLinkChannelId() bool`

HasLinkChannelId returns a boolean if a field has been set.

### GetWebhookId

`func (o *DetailRecord) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DetailRecord) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DetailRecord) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *DetailRecord) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


