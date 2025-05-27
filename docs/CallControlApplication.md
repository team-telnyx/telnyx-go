# CallControlApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to **string** | &#x60;Latency&#x60; directs Telnyx to route media through the site with the lowest round-trip time to the user&#39;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.  | [optional] [default to "\"Latency\""]
**ApplicationName** | Pointer to **string** | A user-assigned name to help manage the application. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was created | [optional] 
**DtmfType** | Pointer to **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to "RFC 2833"]
**FirstCommandTimeout** | Pointer to **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | Pointer to **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**Tags** | Pointer to **[]string** | Tags assigned to the Call Control Application. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Inbound** | Pointer to [**CallControlApplicationInbound**](CallControlApplicationInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CallControlApplicationOutbound**](CallControlApplicationOutbound.md) |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [default to "call_control_application"]
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the resource was last updated | [optional] 
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#x60;https&#x60;. | [optional] [default to ""]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#x60;https&#x60;. | [optional] 
**WebhookTimeoutSecs** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCallControlApplication

`func NewCallControlApplication() *CallControlApplication`

NewCallControlApplication instantiates a new CallControlApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallControlApplicationWithDefaults

`func NewCallControlApplicationWithDefaults() *CallControlApplication`

NewCallControlApplicationWithDefaults instantiates a new CallControlApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CallControlApplication) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CallControlApplication) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CallControlApplication) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CallControlApplication) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *CallControlApplication) GetAnchorsiteOverride() string`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *CallControlApplication) GetAnchorsiteOverrideOk() (*string, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *CallControlApplication) SetAnchorsiteOverride(v string)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *CallControlApplication) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetApplicationName

`func (o *CallControlApplication) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *CallControlApplication) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *CallControlApplication) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *CallControlApplication) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CallControlApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CallControlApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CallControlApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CallControlApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDtmfType

`func (o *CallControlApplication) GetDtmfType() string`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *CallControlApplication) GetDtmfTypeOk() (*string, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *CallControlApplication) SetDtmfType(v string)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *CallControlApplication) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetFirstCommandTimeout

`func (o *CallControlApplication) GetFirstCommandTimeout() bool`

GetFirstCommandTimeout returns the FirstCommandTimeout field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutOk

`func (o *CallControlApplication) GetFirstCommandTimeoutOk() (*bool, bool)`

GetFirstCommandTimeoutOk returns a tuple with the FirstCommandTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeout

`func (o *CallControlApplication) SetFirstCommandTimeout(v bool)`

SetFirstCommandTimeout sets FirstCommandTimeout field to given value.

### HasFirstCommandTimeout

`func (o *CallControlApplication) HasFirstCommandTimeout() bool`

HasFirstCommandTimeout returns a boolean if a field has been set.

### GetFirstCommandTimeoutSecs

`func (o *CallControlApplication) GetFirstCommandTimeoutSecs() int32`

GetFirstCommandTimeoutSecs returns the FirstCommandTimeoutSecs field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutSecsOk

`func (o *CallControlApplication) GetFirstCommandTimeoutSecsOk() (*int32, bool)`

GetFirstCommandTimeoutSecsOk returns a tuple with the FirstCommandTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeoutSecs

`func (o *CallControlApplication) SetFirstCommandTimeoutSecs(v int32)`

SetFirstCommandTimeoutSecs sets FirstCommandTimeoutSecs field to given value.

### HasFirstCommandTimeoutSecs

`func (o *CallControlApplication) HasFirstCommandTimeoutSecs() bool`

HasFirstCommandTimeoutSecs returns a boolean if a field has been set.

### GetTags

`func (o *CallControlApplication) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CallControlApplication) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CallControlApplication) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CallControlApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *CallControlApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallControlApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallControlApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallControlApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInbound

`func (o *CallControlApplication) GetInbound() CallControlApplicationInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CallControlApplication) GetInboundOk() (*CallControlApplicationInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CallControlApplication) SetInbound(v CallControlApplicationInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CallControlApplication) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CallControlApplication) GetOutbound() CallControlApplicationOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CallControlApplication) GetOutboundOk() (*CallControlApplicationOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CallControlApplication) SetOutbound(v CallControlApplicationOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CallControlApplication) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetRecordType

`func (o *CallControlApplication) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallControlApplication) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallControlApplication) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallControlApplication) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CallControlApplication) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CallControlApplication) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CallControlApplication) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CallControlApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookApiVersion

`func (o *CallControlApplication) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *CallControlApplication) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *CallControlApplication) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *CallControlApplication) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CallControlApplication) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CallControlApplication) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CallControlApplication) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CallControlApplication) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CallControlApplication) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CallControlApplication) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookEventUrl

`func (o *CallControlApplication) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CallControlApplication) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CallControlApplication) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CallControlApplication) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *CallControlApplication) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CallControlApplication) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CallControlApplication) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CallControlApplication) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CallControlApplication) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CallControlApplication) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


