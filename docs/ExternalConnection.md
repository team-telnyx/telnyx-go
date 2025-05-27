# ExternalConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**CredentialActive** | Pointer to **bool** | If the credential associated with this service is active. | [optional] [default to false]
**ExternalSipConnection** | Pointer to [**ExternalSipConnection**](ExternalSipConnection.md) |  | [optional] [default to ZOOM]
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Inbound** | Pointer to [**CreateExternalConnectionRequestInbound**](CreateExternalConnectionRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateExternalConnectionRequestOutbound**](CreateExternalConnectionRequestOutbound.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewExternalConnection

`func NewExternalConnection() *ExternalConnection`

NewExternalConnection instantiates a new ExternalConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConnectionWithDefaults

`func NewExternalConnectionWithDefaults() *ExternalConnection`

NewExternalConnectionWithDefaults instantiates a new ExternalConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *ExternalConnection) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ExternalConnection) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ExternalConnection) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ExternalConnection) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActive

`func (o *ExternalConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ExternalConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ExternalConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ExternalConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCredentialActive

`func (o *ExternalConnection) GetCredentialActive() bool`

GetCredentialActive returns the CredentialActive field if non-nil, zero value otherwise.

### GetCredentialActiveOk

`func (o *ExternalConnection) GetCredentialActiveOk() (*bool, bool)`

GetCredentialActiveOk returns a tuple with the CredentialActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialActive

`func (o *ExternalConnection) SetCredentialActive(v bool)`

SetCredentialActive sets CredentialActive field to given value.

### HasCredentialActive

`func (o *ExternalConnection) HasCredentialActive() bool`

HasCredentialActive returns a boolean if a field has been set.

### GetExternalSipConnection

`func (o *ExternalConnection) GetExternalSipConnection() ExternalSipConnection`

GetExternalSipConnection returns the ExternalSipConnection field if non-nil, zero value otherwise.

### GetExternalSipConnectionOk

`func (o *ExternalConnection) GetExternalSipConnectionOk() (*ExternalSipConnection, bool)`

GetExternalSipConnectionOk returns a tuple with the ExternalSipConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSipConnection

`func (o *ExternalConnection) SetExternalSipConnection(v ExternalSipConnection)`

SetExternalSipConnection sets ExternalSipConnection field to given value.

### HasExternalSipConnection

`func (o *ExternalConnection) HasExternalSipConnection() bool`

HasExternalSipConnection returns a boolean if a field has been set.

### GetTags

`func (o *ExternalConnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExternalConnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExternalConnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExternalConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *ExternalConnection) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *ExternalConnection) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *ExternalConnection) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *ExternalConnection) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *ExternalConnection) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *ExternalConnection) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *ExternalConnection) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *ExternalConnection) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *ExternalConnection) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *ExternalConnection) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *ExternalConnection) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *ExternalConnection) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *ExternalConnection) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *ExternalConnection) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *ExternalConnection) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *ExternalConnection) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *ExternalConnection) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *ExternalConnection) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *ExternalConnection) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *ExternalConnection) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetInbound

`func (o *ExternalConnection) GetInbound() CreateExternalConnectionRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *ExternalConnection) GetInboundOk() (*CreateExternalConnectionRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *ExternalConnection) SetInbound(v CreateExternalConnectionRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *ExternalConnection) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *ExternalConnection) GetOutbound() CreateExternalConnectionRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *ExternalConnection) GetOutboundOk() (*CreateExternalConnectionRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *ExternalConnection) SetOutbound(v CreateExternalConnectionRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *ExternalConnection) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExternalConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExternalConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExternalConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExternalConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExternalConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExternalConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


