# CreateTexmlApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | **string** | A user-assigned name to help manage the application. | 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**FirstCommandTimeout** | Pointer to **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | Pointer to **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**Tags** | Pointer to **[]string** | Tags associated with the Texml Application. | [optional] 
**VoiceUrl** | **string** | URL to which Telnyx will deliver your XML Translator webhooks. | 
**VoiceFallbackUrl** | Pointer to **string** | URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url. | [optional] 
**VoiceMethod** | Pointer to **string** | HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either &#39;get&#39; or &#39;post&#39;. | [optional] [default to "post"]
**StatusCallback** | Pointer to **string** | URL for Telnyx to send requests to containing information about call progress events. | [optional] 
**StatusCallbackMethod** | Pointer to **string** | HTTP request method Telnyx should use when requesting the status_callback URL. | [optional] [default to "post"]
**Inbound** | Pointer to [**CreateTexmlApplicationRequestInbound**](CreateTexmlApplicationRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateTexmlApplicationRequestOutbound**](CreateTexmlApplicationRequestOutbound.md) |  | [optional] 

## Methods

### NewCreateTexmlApplicationRequest

`func NewCreateTexmlApplicationRequest(friendlyName string, voiceUrl string, ) *CreateTexmlApplicationRequest`

NewCreateTexmlApplicationRequest instantiates a new CreateTexmlApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTexmlApplicationRequestWithDefaults

`func NewCreateTexmlApplicationRequestWithDefaults() *CreateTexmlApplicationRequest`

NewCreateTexmlApplicationRequestWithDefaults instantiates a new CreateTexmlApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *CreateTexmlApplicationRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *CreateTexmlApplicationRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *CreateTexmlApplicationRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.


### GetActive

`func (o *CreateTexmlApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateTexmlApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateTexmlApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateTexmlApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *CreateTexmlApplicationRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *CreateTexmlApplicationRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *CreateTexmlApplicationRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *CreateTexmlApplicationRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetDtmfType

`func (o *CreateTexmlApplicationRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *CreateTexmlApplicationRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *CreateTexmlApplicationRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *CreateTexmlApplicationRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetFirstCommandTimeout

`func (o *CreateTexmlApplicationRequest) GetFirstCommandTimeout() bool`

GetFirstCommandTimeout returns the FirstCommandTimeout field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutOk

`func (o *CreateTexmlApplicationRequest) GetFirstCommandTimeoutOk() (*bool, bool)`

GetFirstCommandTimeoutOk returns a tuple with the FirstCommandTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeout

`func (o *CreateTexmlApplicationRequest) SetFirstCommandTimeout(v bool)`

SetFirstCommandTimeout sets FirstCommandTimeout field to given value.

### HasFirstCommandTimeout

`func (o *CreateTexmlApplicationRequest) HasFirstCommandTimeout() bool`

HasFirstCommandTimeout returns a boolean if a field has been set.

### GetFirstCommandTimeoutSecs

`func (o *CreateTexmlApplicationRequest) GetFirstCommandTimeoutSecs() int32`

GetFirstCommandTimeoutSecs returns the FirstCommandTimeoutSecs field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutSecsOk

`func (o *CreateTexmlApplicationRequest) GetFirstCommandTimeoutSecsOk() (*int32, bool)`

GetFirstCommandTimeoutSecsOk returns a tuple with the FirstCommandTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeoutSecs

`func (o *CreateTexmlApplicationRequest) SetFirstCommandTimeoutSecs(v int32)`

SetFirstCommandTimeoutSecs sets FirstCommandTimeoutSecs field to given value.

### HasFirstCommandTimeoutSecs

`func (o *CreateTexmlApplicationRequest) HasFirstCommandTimeoutSecs() bool`

HasFirstCommandTimeoutSecs returns a boolean if a field has been set.

### GetTags

`func (o *CreateTexmlApplicationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateTexmlApplicationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateTexmlApplicationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateTexmlApplicationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVoiceUrl

`func (o *CreateTexmlApplicationRequest) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *CreateTexmlApplicationRequest) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *CreateTexmlApplicationRequest) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.


### GetVoiceFallbackUrl

`func (o *CreateTexmlApplicationRequest) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *CreateTexmlApplicationRequest) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *CreateTexmlApplicationRequest) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *CreateTexmlApplicationRequest) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### GetVoiceMethod

`func (o *CreateTexmlApplicationRequest) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *CreateTexmlApplicationRequest) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *CreateTexmlApplicationRequest) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *CreateTexmlApplicationRequest) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### GetStatusCallback

`func (o *CreateTexmlApplicationRequest) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *CreateTexmlApplicationRequest) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *CreateTexmlApplicationRequest) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *CreateTexmlApplicationRequest) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *CreateTexmlApplicationRequest) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *CreateTexmlApplicationRequest) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *CreateTexmlApplicationRequest) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *CreateTexmlApplicationRequest) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetInbound

`func (o *CreateTexmlApplicationRequest) GetInbound() CreateTexmlApplicationRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CreateTexmlApplicationRequest) GetInboundOk() (*CreateTexmlApplicationRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CreateTexmlApplicationRequest) SetInbound(v CreateTexmlApplicationRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CreateTexmlApplicationRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CreateTexmlApplicationRequest) GetOutbound() CreateTexmlApplicationRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CreateTexmlApplicationRequest) GetOutboundOk() (*CreateTexmlApplicationRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CreateTexmlApplicationRequest) SetOutbound(v CreateTexmlApplicationRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CreateTexmlApplicationRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


