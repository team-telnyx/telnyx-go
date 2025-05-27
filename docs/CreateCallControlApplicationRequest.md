# CreateCallControlApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | **string** | A user-assigned name to help manage the application. | 
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to **string** | &lt;code&gt;Latency&lt;/code&gt; directs Telnyx to route media through the site with the lowest round-trip time to the user&#39;s connection. Telnyx calculates this time using ICMP ping messages. This can be disabled by specifying a site to handle all media.  | [optional] [default to "\"Latency\""]
**DtmfType** | Pointer to **string** | Sets the type of DTMF digits sent from Telnyx to this Connection. Note that DTMF digits sent to Telnyx will be accepted in all formats. | [optional] [default to "RFC 2833"]
**FirstCommandTimeout** | Pointer to **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | Pointer to **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**Inbound** | Pointer to [**CallControlApplicationInbound**](CallControlApplicationInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CallControlApplicationOutbound**](CallControlApplicationOutbound.md) |  | [optional] 
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 

## Methods

### NewCreateCallControlApplicationRequest

`func NewCreateCallControlApplicationRequest(applicationName string, webhookEventUrl string, ) *CreateCallControlApplicationRequest`

NewCreateCallControlApplicationRequest instantiates a new CreateCallControlApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCallControlApplicationRequestWithDefaults

`func NewCreateCallControlApplicationRequestWithDefaults() *CreateCallControlApplicationRequest`

NewCreateCallControlApplicationRequestWithDefaults instantiates a new CreateCallControlApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *CreateCallControlApplicationRequest) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *CreateCallControlApplicationRequest) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *CreateCallControlApplicationRequest) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetWebhookEventUrl

`func (o *CreateCallControlApplicationRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CreateCallControlApplicationRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CreateCallControlApplicationRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.


### GetActive

`func (o *CreateCallControlApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateCallControlApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateCallControlApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateCallControlApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *CreateCallControlApplicationRequest) GetAnchorsiteOverride() string`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *CreateCallControlApplicationRequest) GetAnchorsiteOverrideOk() (*string, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *CreateCallControlApplicationRequest) SetAnchorsiteOverride(v string)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *CreateCallControlApplicationRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetDtmfType

`func (o *CreateCallControlApplicationRequest) GetDtmfType() string`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *CreateCallControlApplicationRequest) GetDtmfTypeOk() (*string, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *CreateCallControlApplicationRequest) SetDtmfType(v string)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *CreateCallControlApplicationRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetFirstCommandTimeout

`func (o *CreateCallControlApplicationRequest) GetFirstCommandTimeout() bool`

GetFirstCommandTimeout returns the FirstCommandTimeout field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutOk

`func (o *CreateCallControlApplicationRequest) GetFirstCommandTimeoutOk() (*bool, bool)`

GetFirstCommandTimeoutOk returns a tuple with the FirstCommandTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeout

`func (o *CreateCallControlApplicationRequest) SetFirstCommandTimeout(v bool)`

SetFirstCommandTimeout sets FirstCommandTimeout field to given value.

### HasFirstCommandTimeout

`func (o *CreateCallControlApplicationRequest) HasFirstCommandTimeout() bool`

HasFirstCommandTimeout returns a boolean if a field has been set.

### GetFirstCommandTimeoutSecs

`func (o *CreateCallControlApplicationRequest) GetFirstCommandTimeoutSecs() int32`

GetFirstCommandTimeoutSecs returns the FirstCommandTimeoutSecs field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutSecsOk

`func (o *CreateCallControlApplicationRequest) GetFirstCommandTimeoutSecsOk() (*int32, bool)`

GetFirstCommandTimeoutSecsOk returns a tuple with the FirstCommandTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeoutSecs

`func (o *CreateCallControlApplicationRequest) SetFirstCommandTimeoutSecs(v int32)`

SetFirstCommandTimeoutSecs sets FirstCommandTimeoutSecs field to given value.

### HasFirstCommandTimeoutSecs

`func (o *CreateCallControlApplicationRequest) HasFirstCommandTimeoutSecs() bool`

HasFirstCommandTimeoutSecs returns a boolean if a field has been set.

### GetInbound

`func (o *CreateCallControlApplicationRequest) GetInbound() CallControlApplicationInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CreateCallControlApplicationRequest) GetInboundOk() (*CallControlApplicationInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CreateCallControlApplicationRequest) SetInbound(v CallControlApplicationInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CreateCallControlApplicationRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CreateCallControlApplicationRequest) GetOutbound() CallControlApplicationOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CreateCallControlApplicationRequest) GetOutboundOk() (*CallControlApplicationOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CreateCallControlApplicationRequest) SetOutbound(v CallControlApplicationOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CreateCallControlApplicationRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetWebhookApiVersion

`func (o *CreateCallControlApplicationRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *CreateCallControlApplicationRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *CreateCallControlApplicationRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *CreateCallControlApplicationRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CreateCallControlApplicationRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CreateCallControlApplicationRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CreateCallControlApplicationRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CreateCallControlApplicationRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CreateCallControlApplicationRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CreateCallControlApplicationRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *CreateCallControlApplicationRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CreateCallControlApplicationRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CreateCallControlApplicationRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CreateCallControlApplicationRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CreateCallControlApplicationRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CreateCallControlApplicationRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


