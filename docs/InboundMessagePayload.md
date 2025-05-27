# InboundMessagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Direction** | Pointer to **string** | The direction of the message. Inbound messages are sent to you whereas outbound messages are sent from you. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Type** | Pointer to **string** | The type of message. This value can be either &#39;sms&#39; or &#39;mms&#39;. | [optional] 
**MessagingProfileId** | Pointer to **string** | Unique identifier for a messaging profile. | [optional] 
**To** | Pointer to [**[]InboundMessagePayloadToInner**](InboundMessagePayloadToInner.md) |  | [optional] 
**Cc** | Pointer to [**[]InboundMessagePayloadCcInner**](InboundMessagePayloadCcInner.md) |  | [optional] 
**From** | Pointer to [**InboundMessagePayloadFrom**](InboundMessagePayloadFrom.md) |  | [optional] 
**Text** | Pointer to **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] 
**Media** | Pointer to [**[]InboundMessagePayloadMediaInner**](InboundMessagePayloadMediaInner.md) |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL where webhooks related to this message will be sent. | [optional] 
**WebhookFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this message will be sent if sending to the primary URL fails. | [optional] 
**Encoding** | Pointer to **string** | Encoding scheme used for the message body. | [optional] 
**Parts** | Pointer to **int32** | Number of parts into which the message&#39;s body must be split. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the resource. | [optional] 
**Cost** | Pointer to [**NullableInboundMessagePayloadCost**](InboundMessagePayloadCost.md) |  | [optional] 
**CostBreakdown** | Pointer to [**NullableInboundMessagePayloadCostBreakdown**](InboundMessagePayloadCostBreakdown.md) |  | [optional] 
**TcrCampaignId** | Pointer to **NullableString** | The Campaign Registry (TCR) campaign ID associated with the message. | [optional] 
**TcrCampaignBillable** | Pointer to **bool** | Indicates whether the TCR campaign is billable. | [optional] 
**TcrCampaignRegistered** | Pointer to **NullableString** | The registration status of the TCR campaign. | [optional] 
**ReceivedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the message request was received. | [optional] 
**SentAt** | Pointer to **NullableTime** | Not used for inbound messages. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | Not used for inbound messages. | [optional] 
**ValidUntil** | Pointer to **NullableTime** | Not used for inbound messages. | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | These errors may point at addressees when referring to unsuccessful/unconfirmed delivery statuses. | [optional] 

## Methods

### NewInboundMessagePayload

`func NewInboundMessagePayload() *InboundMessagePayload`

NewInboundMessagePayload instantiates a new InboundMessagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundMessagePayloadWithDefaults

`func NewInboundMessagePayloadWithDefaults() *InboundMessagePayload`

NewInboundMessagePayloadWithDefaults instantiates a new InboundMessagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *InboundMessagePayload) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *InboundMessagePayload) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *InboundMessagePayload) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *InboundMessagePayload) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDirection

`func (o *InboundMessagePayload) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InboundMessagePayload) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InboundMessagePayload) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InboundMessagePayload) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *InboundMessagePayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboundMessagePayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboundMessagePayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InboundMessagePayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InboundMessagePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InboundMessagePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InboundMessagePayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InboundMessagePayload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *InboundMessagePayload) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *InboundMessagePayload) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *InboundMessagePayload) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *InboundMessagePayload) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetTo

`func (o *InboundMessagePayload) GetTo() []InboundMessagePayloadToInner`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InboundMessagePayload) GetToOk() (*[]InboundMessagePayloadToInner, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InboundMessagePayload) SetTo(v []InboundMessagePayloadToInner)`

SetTo sets To field to given value.

### HasTo

`func (o *InboundMessagePayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *InboundMessagePayload) GetCc() []InboundMessagePayloadCcInner`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *InboundMessagePayload) GetCcOk() (*[]InboundMessagePayloadCcInner, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *InboundMessagePayload) SetCc(v []InboundMessagePayloadCcInner)`

SetCc sets Cc field to given value.

### HasCc

`func (o *InboundMessagePayload) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetFrom

`func (o *InboundMessagePayload) GetFrom() InboundMessagePayloadFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InboundMessagePayload) GetFromOk() (*InboundMessagePayloadFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InboundMessagePayload) SetFrom(v InboundMessagePayloadFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InboundMessagePayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetText

`func (o *InboundMessagePayload) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InboundMessagePayload) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InboundMessagePayload) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *InboundMessagePayload) HasText() bool`

HasText returns a boolean if a field has been set.

### GetMedia

`func (o *InboundMessagePayload) GetMedia() []InboundMessagePayloadMediaInner`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *InboundMessagePayload) GetMediaOk() (*[]InboundMessagePayloadMediaInner, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *InboundMessagePayload) SetMedia(v []InboundMessagePayloadMediaInner)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *InboundMessagePayload) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *InboundMessagePayload) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *InboundMessagePayload) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *InboundMessagePayload) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *InboundMessagePayload) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *InboundMessagePayload) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *InboundMessagePayload) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookFailoverUrl

`func (o *InboundMessagePayload) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *InboundMessagePayload) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *InboundMessagePayload) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *InboundMessagePayload) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### SetWebhookFailoverUrlNil

`func (o *InboundMessagePayload) SetWebhookFailoverUrlNil(b bool)`

 SetWebhookFailoverUrlNil sets the value for WebhookFailoverUrl to be an explicit nil

### UnsetWebhookFailoverUrl
`func (o *InboundMessagePayload) UnsetWebhookFailoverUrl()`

UnsetWebhookFailoverUrl ensures that no value is present for WebhookFailoverUrl, not even an explicit nil
### GetEncoding

`func (o *InboundMessagePayload) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *InboundMessagePayload) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *InboundMessagePayload) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *InboundMessagePayload) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetParts

`func (o *InboundMessagePayload) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *InboundMessagePayload) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *InboundMessagePayload) SetParts(v int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *InboundMessagePayload) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetTags

`func (o *InboundMessagePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InboundMessagePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InboundMessagePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InboundMessagePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCost

`func (o *InboundMessagePayload) GetCost() InboundMessagePayloadCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *InboundMessagePayload) GetCostOk() (*InboundMessagePayloadCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *InboundMessagePayload) SetCost(v InboundMessagePayloadCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *InboundMessagePayload) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *InboundMessagePayload) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *InboundMessagePayload) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostBreakdown

`func (o *InboundMessagePayload) GetCostBreakdown() InboundMessagePayloadCostBreakdown`

GetCostBreakdown returns the CostBreakdown field if non-nil, zero value otherwise.

### GetCostBreakdownOk

`func (o *InboundMessagePayload) GetCostBreakdownOk() (*InboundMessagePayloadCostBreakdown, bool)`

GetCostBreakdownOk returns a tuple with the CostBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBreakdown

`func (o *InboundMessagePayload) SetCostBreakdown(v InboundMessagePayloadCostBreakdown)`

SetCostBreakdown sets CostBreakdown field to given value.

### HasCostBreakdown

`func (o *InboundMessagePayload) HasCostBreakdown() bool`

HasCostBreakdown returns a boolean if a field has been set.

### SetCostBreakdownNil

`func (o *InboundMessagePayload) SetCostBreakdownNil(b bool)`

 SetCostBreakdownNil sets the value for CostBreakdown to be an explicit nil

### UnsetCostBreakdown
`func (o *InboundMessagePayload) UnsetCostBreakdown()`

UnsetCostBreakdown ensures that no value is present for CostBreakdown, not even an explicit nil
### GetTcrCampaignId

`func (o *InboundMessagePayload) GetTcrCampaignId() string`

GetTcrCampaignId returns the TcrCampaignId field if non-nil, zero value otherwise.

### GetTcrCampaignIdOk

`func (o *InboundMessagePayload) GetTcrCampaignIdOk() (*string, bool)`

GetTcrCampaignIdOk returns a tuple with the TcrCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignId

`func (o *InboundMessagePayload) SetTcrCampaignId(v string)`

SetTcrCampaignId sets TcrCampaignId field to given value.

### HasTcrCampaignId

`func (o *InboundMessagePayload) HasTcrCampaignId() bool`

HasTcrCampaignId returns a boolean if a field has been set.

### SetTcrCampaignIdNil

`func (o *InboundMessagePayload) SetTcrCampaignIdNil(b bool)`

 SetTcrCampaignIdNil sets the value for TcrCampaignId to be an explicit nil

### UnsetTcrCampaignId
`func (o *InboundMessagePayload) UnsetTcrCampaignId()`

UnsetTcrCampaignId ensures that no value is present for TcrCampaignId, not even an explicit nil
### GetTcrCampaignBillable

`func (o *InboundMessagePayload) GetTcrCampaignBillable() bool`

GetTcrCampaignBillable returns the TcrCampaignBillable field if non-nil, zero value otherwise.

### GetTcrCampaignBillableOk

`func (o *InboundMessagePayload) GetTcrCampaignBillableOk() (*bool, bool)`

GetTcrCampaignBillableOk returns a tuple with the TcrCampaignBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignBillable

`func (o *InboundMessagePayload) SetTcrCampaignBillable(v bool)`

SetTcrCampaignBillable sets TcrCampaignBillable field to given value.

### HasTcrCampaignBillable

`func (o *InboundMessagePayload) HasTcrCampaignBillable() bool`

HasTcrCampaignBillable returns a boolean if a field has been set.

### GetTcrCampaignRegistered

`func (o *InboundMessagePayload) GetTcrCampaignRegistered() string`

GetTcrCampaignRegistered returns the TcrCampaignRegistered field if non-nil, zero value otherwise.

### GetTcrCampaignRegisteredOk

`func (o *InboundMessagePayload) GetTcrCampaignRegisteredOk() (*string, bool)`

GetTcrCampaignRegisteredOk returns a tuple with the TcrCampaignRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignRegistered

`func (o *InboundMessagePayload) SetTcrCampaignRegistered(v string)`

SetTcrCampaignRegistered sets TcrCampaignRegistered field to given value.

### HasTcrCampaignRegistered

`func (o *InboundMessagePayload) HasTcrCampaignRegistered() bool`

HasTcrCampaignRegistered returns a boolean if a field has been set.

### SetTcrCampaignRegisteredNil

`func (o *InboundMessagePayload) SetTcrCampaignRegisteredNil(b bool)`

 SetTcrCampaignRegisteredNil sets the value for TcrCampaignRegistered to be an explicit nil

### UnsetTcrCampaignRegistered
`func (o *InboundMessagePayload) UnsetTcrCampaignRegistered()`

UnsetTcrCampaignRegistered ensures that no value is present for TcrCampaignRegistered, not even an explicit nil
### GetReceivedAt

`func (o *InboundMessagePayload) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *InboundMessagePayload) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *InboundMessagePayload) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *InboundMessagePayload) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetSentAt

`func (o *InboundMessagePayload) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InboundMessagePayload) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InboundMessagePayload) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *InboundMessagePayload) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *InboundMessagePayload) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *InboundMessagePayload) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetCompletedAt

`func (o *InboundMessagePayload) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InboundMessagePayload) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InboundMessagePayload) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InboundMessagePayload) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *InboundMessagePayload) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *InboundMessagePayload) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetValidUntil

`func (o *InboundMessagePayload) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *InboundMessagePayload) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *InboundMessagePayload) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *InboundMessagePayload) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *InboundMessagePayload) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *InboundMessagePayload) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetErrors

`func (o *InboundMessagePayload) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InboundMessagePayload) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InboundMessagePayload) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InboundMessagePayload) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


