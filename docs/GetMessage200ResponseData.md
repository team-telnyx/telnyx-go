# GetMessage200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Direction** | Pointer to **string** | The direction of the message. Inbound messages are sent to you whereas outbound messages are sent from you. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Type** | Pointer to **string** | The type of message. | [optional] 
**MessagingProfileId** | Pointer to **string** | Unique identifier for a messaging profile. | [optional] 
**OrganizationId** | Pointer to **string** | The id of the organization the messaging profile belongs to. | [optional] 
**From** | Pointer to [**InboundMessagePayloadFrom**](InboundMessagePayloadFrom.md) |  | [optional] 
**To** | Pointer to [**[]InboundMessagePayloadToInner**](InboundMessagePayloadToInner.md) |  | [optional] 
**Text** | Pointer to **string** | Message body (i.e., content) as a non-empty string.  **Required for SMS** | [optional] 
**Subject** | Pointer to **NullableString** | Subject of multimedia message | [optional] 
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
**Cc** | Pointer to [**[]InboundMessagePayloadCcInner**](InboundMessagePayloadCcInner.md) |  | [optional] 

## Methods

### NewGetMessage200ResponseData

`func NewGetMessage200ResponseData() *GetMessage200ResponseData`

NewGetMessage200ResponseData instantiates a new GetMessage200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessage200ResponseDataWithDefaults

`func NewGetMessage200ResponseDataWithDefaults() *GetMessage200ResponseData`

NewGetMessage200ResponseDataWithDefaults instantiates a new GetMessage200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *GetMessage200ResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GetMessage200ResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GetMessage200ResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GetMessage200ResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDirection

`func (o *GetMessage200ResponseData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetMessage200ResponseData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetMessage200ResponseData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetMessage200ResponseData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *GetMessage200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMessage200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMessage200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetMessage200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMessage200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMessage200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMessage200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMessage200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *GetMessage200ResponseData) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *GetMessage200ResponseData) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *GetMessage200ResponseData) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *GetMessage200ResponseData) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetMessage200ResponseData) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetMessage200ResponseData) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetMessage200ResponseData) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetMessage200ResponseData) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetFrom

`func (o *GetMessage200ResponseData) GetFrom() InboundMessagePayloadFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetMessage200ResponseData) GetFromOk() (*InboundMessagePayloadFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetMessage200ResponseData) SetFrom(v InboundMessagePayloadFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetMessage200ResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetMessage200ResponseData) GetTo() []InboundMessagePayloadToInner`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetMessage200ResponseData) GetToOk() (*[]InboundMessagePayloadToInner, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetMessage200ResponseData) SetTo(v []InboundMessagePayloadToInner)`

SetTo sets To field to given value.

### HasTo

`func (o *GetMessage200ResponseData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetText

`func (o *GetMessage200ResponseData) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GetMessage200ResponseData) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GetMessage200ResponseData) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GetMessage200ResponseData) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSubject

`func (o *GetMessage200ResponseData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetMessage200ResponseData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetMessage200ResponseData) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetMessage200ResponseData) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *GetMessage200ResponseData) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *GetMessage200ResponseData) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMedia

`func (o *GetMessage200ResponseData) GetMedia() []InboundMessagePayloadMediaInner`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *GetMessage200ResponseData) GetMediaOk() (*[]InboundMessagePayloadMediaInner, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *GetMessage200ResponseData) SetMedia(v []InboundMessagePayloadMediaInner)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *GetMessage200ResponseData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *GetMessage200ResponseData) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GetMessage200ResponseData) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GetMessage200ResponseData) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *GetMessage200ResponseData) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *GetMessage200ResponseData) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *GetMessage200ResponseData) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhookFailoverUrl

`func (o *GetMessage200ResponseData) GetWebhookFailoverUrl() string`

GetWebhookFailoverUrl returns the WebhookFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookFailoverUrlOk

`func (o *GetMessage200ResponseData) GetWebhookFailoverUrlOk() (*string, bool)`

GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverUrl

`func (o *GetMessage200ResponseData) SetWebhookFailoverUrl(v string)`

SetWebhookFailoverUrl sets WebhookFailoverUrl field to given value.

### HasWebhookFailoverUrl

`func (o *GetMessage200ResponseData) HasWebhookFailoverUrl() bool`

HasWebhookFailoverUrl returns a boolean if a field has been set.

### SetWebhookFailoverUrlNil

`func (o *GetMessage200ResponseData) SetWebhookFailoverUrlNil(b bool)`

 SetWebhookFailoverUrlNil sets the value for WebhookFailoverUrl to be an explicit nil

### UnsetWebhookFailoverUrl
`func (o *GetMessage200ResponseData) UnsetWebhookFailoverUrl()`

UnsetWebhookFailoverUrl ensures that no value is present for WebhookFailoverUrl, not even an explicit nil
### GetEncoding

`func (o *GetMessage200ResponseData) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GetMessage200ResponseData) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GetMessage200ResponseData) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GetMessage200ResponseData) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetParts

`func (o *GetMessage200ResponseData) GetParts() int32`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *GetMessage200ResponseData) GetPartsOk() (*int32, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *GetMessage200ResponseData) SetParts(v int32)`

SetParts sets Parts field to given value.

### HasParts

`func (o *GetMessage200ResponseData) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetTags

`func (o *GetMessage200ResponseData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetMessage200ResponseData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetMessage200ResponseData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetMessage200ResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCost

`func (o *GetMessage200ResponseData) GetCost() InboundMessagePayloadCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetMessage200ResponseData) GetCostOk() (*InboundMessagePayloadCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetMessage200ResponseData) SetCost(v InboundMessagePayloadCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetMessage200ResponseData) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GetMessage200ResponseData) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GetMessage200ResponseData) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostBreakdown

`func (o *GetMessage200ResponseData) GetCostBreakdown() InboundMessagePayloadCostBreakdown`

GetCostBreakdown returns the CostBreakdown field if non-nil, zero value otherwise.

### GetCostBreakdownOk

`func (o *GetMessage200ResponseData) GetCostBreakdownOk() (*InboundMessagePayloadCostBreakdown, bool)`

GetCostBreakdownOk returns a tuple with the CostBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBreakdown

`func (o *GetMessage200ResponseData) SetCostBreakdown(v InboundMessagePayloadCostBreakdown)`

SetCostBreakdown sets CostBreakdown field to given value.

### HasCostBreakdown

`func (o *GetMessage200ResponseData) HasCostBreakdown() bool`

HasCostBreakdown returns a boolean if a field has been set.

### SetCostBreakdownNil

`func (o *GetMessage200ResponseData) SetCostBreakdownNil(b bool)`

 SetCostBreakdownNil sets the value for CostBreakdown to be an explicit nil

### UnsetCostBreakdown
`func (o *GetMessage200ResponseData) UnsetCostBreakdown()`

UnsetCostBreakdown ensures that no value is present for CostBreakdown, not even an explicit nil
### GetTcrCampaignId

`func (o *GetMessage200ResponseData) GetTcrCampaignId() string`

GetTcrCampaignId returns the TcrCampaignId field if non-nil, zero value otherwise.

### GetTcrCampaignIdOk

`func (o *GetMessage200ResponseData) GetTcrCampaignIdOk() (*string, bool)`

GetTcrCampaignIdOk returns a tuple with the TcrCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignId

`func (o *GetMessage200ResponseData) SetTcrCampaignId(v string)`

SetTcrCampaignId sets TcrCampaignId field to given value.

### HasTcrCampaignId

`func (o *GetMessage200ResponseData) HasTcrCampaignId() bool`

HasTcrCampaignId returns a boolean if a field has been set.

### SetTcrCampaignIdNil

`func (o *GetMessage200ResponseData) SetTcrCampaignIdNil(b bool)`

 SetTcrCampaignIdNil sets the value for TcrCampaignId to be an explicit nil

### UnsetTcrCampaignId
`func (o *GetMessage200ResponseData) UnsetTcrCampaignId()`

UnsetTcrCampaignId ensures that no value is present for TcrCampaignId, not even an explicit nil
### GetTcrCampaignBillable

`func (o *GetMessage200ResponseData) GetTcrCampaignBillable() bool`

GetTcrCampaignBillable returns the TcrCampaignBillable field if non-nil, zero value otherwise.

### GetTcrCampaignBillableOk

`func (o *GetMessage200ResponseData) GetTcrCampaignBillableOk() (*bool, bool)`

GetTcrCampaignBillableOk returns a tuple with the TcrCampaignBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignBillable

`func (o *GetMessage200ResponseData) SetTcrCampaignBillable(v bool)`

SetTcrCampaignBillable sets TcrCampaignBillable field to given value.

### HasTcrCampaignBillable

`func (o *GetMessage200ResponseData) HasTcrCampaignBillable() bool`

HasTcrCampaignBillable returns a boolean if a field has been set.

### GetTcrCampaignRegistered

`func (o *GetMessage200ResponseData) GetTcrCampaignRegistered() string`

GetTcrCampaignRegistered returns the TcrCampaignRegistered field if non-nil, zero value otherwise.

### GetTcrCampaignRegisteredOk

`func (o *GetMessage200ResponseData) GetTcrCampaignRegisteredOk() (*string, bool)`

GetTcrCampaignRegisteredOk returns a tuple with the TcrCampaignRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrCampaignRegistered

`func (o *GetMessage200ResponseData) SetTcrCampaignRegistered(v string)`

SetTcrCampaignRegistered sets TcrCampaignRegistered field to given value.

### HasTcrCampaignRegistered

`func (o *GetMessage200ResponseData) HasTcrCampaignRegistered() bool`

HasTcrCampaignRegistered returns a boolean if a field has been set.

### SetTcrCampaignRegisteredNil

`func (o *GetMessage200ResponseData) SetTcrCampaignRegisteredNil(b bool)`

 SetTcrCampaignRegisteredNil sets the value for TcrCampaignRegistered to be an explicit nil

### UnsetTcrCampaignRegistered
`func (o *GetMessage200ResponseData) UnsetTcrCampaignRegistered()`

UnsetTcrCampaignRegistered ensures that no value is present for TcrCampaignRegistered, not even an explicit nil
### GetReceivedAt

`func (o *GetMessage200ResponseData) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *GetMessage200ResponseData) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *GetMessage200ResponseData) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *GetMessage200ResponseData) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetSentAt

`func (o *GetMessage200ResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *GetMessage200ResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *GetMessage200ResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *GetMessage200ResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *GetMessage200ResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *GetMessage200ResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetCompletedAt

`func (o *GetMessage200ResponseData) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetMessage200ResponseData) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetMessage200ResponseData) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetMessage200ResponseData) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetMessage200ResponseData) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetMessage200ResponseData) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetValidUntil

`func (o *GetMessage200ResponseData) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *GetMessage200ResponseData) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *GetMessage200ResponseData) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *GetMessage200ResponseData) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *GetMessage200ResponseData) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *GetMessage200ResponseData) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetErrors

`func (o *GetMessage200ResponseData) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetMessage200ResponseData) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetMessage200ResponseData) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetMessage200ResponseData) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetCc

`func (o *GetMessage200ResponseData) GetCc() []InboundMessagePayloadCcInner`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *GetMessage200ResponseData) GetCcOk() (*[]InboundMessagePayloadCcInner, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *GetMessage200ResponseData) SetCc(v []InboundMessagePayloadCcInner)`

SetCc sets Cc field to given value.

### HasCc

`func (o *GetMessage200ResponseData) HasCc() bool`

HasCc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


