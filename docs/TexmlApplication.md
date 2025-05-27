# TexmlApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**FriendlyName** | Pointer to **string** | A user-assigned name to help manage the application. | [optional] 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**FirstCommandTimeout** | Pointer to **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | Pointer to **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**VoiceUrl** | Pointer to **string** | URL to which Telnyx will deliver your XML Translator webhooks. | [optional] 
**VoiceFallbackUrl** | Pointer to **string** | URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url. | [optional] 
**VoiceMethod** | Pointer to **string** | HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either &#39;get&#39; or &#39;post&#39;. | [optional] [default to "post"]
**StatusCallback** | Pointer to **string** | URL for Telnyx to send requests to containing information about call progress events. | [optional] 
**StatusCallbackMethod** | Pointer to **string** | HTTP request method Telnyx should use when requesting the status_callback URL. | [optional] [default to "post"]
**Tags** | Pointer to **[]string** | Tags associated with the Texml Application. | [optional] 
**Inbound** | Pointer to [**CreateTexmlApplicationRequestInbound**](CreateTexmlApplicationRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateTexmlApplicationRequestOutbound**](CreateTexmlApplicationRequestOutbound.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewTexmlApplication

`func NewTexmlApplication() *TexmlApplication`

NewTexmlApplication instantiates a new TexmlApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTexmlApplicationWithDefaults

`func NewTexmlApplicationWithDefaults() *TexmlApplication`

NewTexmlApplicationWithDefaults instantiates a new TexmlApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TexmlApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TexmlApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TexmlApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TexmlApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *TexmlApplication) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *TexmlApplication) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *TexmlApplication) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *TexmlApplication) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetFriendlyName

`func (o *TexmlApplication) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TexmlApplication) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TexmlApplication) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TexmlApplication) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetActive

`func (o *TexmlApplication) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TexmlApplication) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TexmlApplication) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TexmlApplication) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *TexmlApplication) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *TexmlApplication) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *TexmlApplication) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *TexmlApplication) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetDtmfType

`func (o *TexmlApplication) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *TexmlApplication) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *TexmlApplication) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *TexmlApplication) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetFirstCommandTimeout

`func (o *TexmlApplication) GetFirstCommandTimeout() bool`

GetFirstCommandTimeout returns the FirstCommandTimeout field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutOk

`func (o *TexmlApplication) GetFirstCommandTimeoutOk() (*bool, bool)`

GetFirstCommandTimeoutOk returns a tuple with the FirstCommandTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeout

`func (o *TexmlApplication) SetFirstCommandTimeout(v bool)`

SetFirstCommandTimeout sets FirstCommandTimeout field to given value.

### HasFirstCommandTimeout

`func (o *TexmlApplication) HasFirstCommandTimeout() bool`

HasFirstCommandTimeout returns a boolean if a field has been set.

### GetFirstCommandTimeoutSecs

`func (o *TexmlApplication) GetFirstCommandTimeoutSecs() int32`

GetFirstCommandTimeoutSecs returns the FirstCommandTimeoutSecs field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutSecsOk

`func (o *TexmlApplication) GetFirstCommandTimeoutSecsOk() (*int32, bool)`

GetFirstCommandTimeoutSecsOk returns a tuple with the FirstCommandTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeoutSecs

`func (o *TexmlApplication) SetFirstCommandTimeoutSecs(v int32)`

SetFirstCommandTimeoutSecs sets FirstCommandTimeoutSecs field to given value.

### HasFirstCommandTimeoutSecs

`func (o *TexmlApplication) HasFirstCommandTimeoutSecs() bool`

HasFirstCommandTimeoutSecs returns a boolean if a field has been set.

### GetVoiceUrl

`func (o *TexmlApplication) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *TexmlApplication) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *TexmlApplication) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *TexmlApplication) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### GetVoiceFallbackUrl

`func (o *TexmlApplication) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *TexmlApplication) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *TexmlApplication) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *TexmlApplication) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### GetVoiceMethod

`func (o *TexmlApplication) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *TexmlApplication) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *TexmlApplication) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *TexmlApplication) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### GetStatusCallback

`func (o *TexmlApplication) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *TexmlApplication) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *TexmlApplication) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *TexmlApplication) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *TexmlApplication) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *TexmlApplication) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *TexmlApplication) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *TexmlApplication) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetTags

`func (o *TexmlApplication) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TexmlApplication) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TexmlApplication) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TexmlApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInbound

`func (o *TexmlApplication) GetInbound() CreateTexmlApplicationRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *TexmlApplication) GetInboundOk() (*CreateTexmlApplicationRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *TexmlApplication) SetInbound(v CreateTexmlApplicationRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *TexmlApplication) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *TexmlApplication) GetOutbound() CreateTexmlApplicationRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *TexmlApplication) GetOutboundOk() (*CreateTexmlApplicationRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *TexmlApplication) SetOutbound(v CreateTexmlApplicationRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *TexmlApplication) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TexmlApplication) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TexmlApplication) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TexmlApplication) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TexmlApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TexmlApplication) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TexmlApplication) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TexmlApplication) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TexmlApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


