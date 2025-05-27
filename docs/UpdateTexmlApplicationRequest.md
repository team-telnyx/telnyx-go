# UpdateTexmlApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | **string** | A user-assigned name to help manage the application. | 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**FirstCommandTimeout** | Pointer to **bool** | Specifies whether calls to phone numbers associated with this connection should hangup after timing out. | [optional] [default to false]
**FirstCommandTimeoutSecs** | Pointer to **int32** | Specifies how many seconds to wait before timing out a dial command. | [optional] [default to 30]
**VoiceUrl** | **string** | URL to which Telnyx will deliver your XML Translator webhooks. | 
**VoiceFallbackUrl** | Pointer to **string** | URL to which Telnyx will deliver your XML Translator webhooks if we get an error response from your voice_url. | [optional] 
**VoiceMethod** | Pointer to **string** | HTTP request method Telnyx will use to interact with your XML Translator webhooks. Either &#39;get&#39; or &#39;post&#39;. | [optional] [default to "post"]
**StatusCallback** | Pointer to **string** | URL for Telnyx to send requests to containing information about call progress events. | [optional] 
**StatusCallbackMethod** | Pointer to **string** | HTTP request method Telnyx should use when requesting the status_callback URL. | [optional] [default to "post"]
**Tags** | Pointer to **[]string** | Tags associated with the Texml Application. | [optional] 
**Inbound** | Pointer to [**CreateTexmlApplicationRequestInbound**](CreateTexmlApplicationRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateTexmlApplicationRequestOutbound**](CreateTexmlApplicationRequestOutbound.md) |  | [optional] 

## Methods

### NewUpdateTexmlApplicationRequest

`func NewUpdateTexmlApplicationRequest(friendlyName string, voiceUrl string, ) *UpdateTexmlApplicationRequest`

NewUpdateTexmlApplicationRequest instantiates a new UpdateTexmlApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTexmlApplicationRequestWithDefaults

`func NewUpdateTexmlApplicationRequestWithDefaults() *UpdateTexmlApplicationRequest`

NewUpdateTexmlApplicationRequestWithDefaults instantiates a new UpdateTexmlApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *UpdateTexmlApplicationRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *UpdateTexmlApplicationRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *UpdateTexmlApplicationRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.


### GetActive

`func (o *UpdateTexmlApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateTexmlApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateTexmlApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateTexmlApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *UpdateTexmlApplicationRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *UpdateTexmlApplicationRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *UpdateTexmlApplicationRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *UpdateTexmlApplicationRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetDtmfType

`func (o *UpdateTexmlApplicationRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *UpdateTexmlApplicationRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *UpdateTexmlApplicationRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *UpdateTexmlApplicationRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetFirstCommandTimeout

`func (o *UpdateTexmlApplicationRequest) GetFirstCommandTimeout() bool`

GetFirstCommandTimeout returns the FirstCommandTimeout field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutOk

`func (o *UpdateTexmlApplicationRequest) GetFirstCommandTimeoutOk() (*bool, bool)`

GetFirstCommandTimeoutOk returns a tuple with the FirstCommandTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeout

`func (o *UpdateTexmlApplicationRequest) SetFirstCommandTimeout(v bool)`

SetFirstCommandTimeout sets FirstCommandTimeout field to given value.

### HasFirstCommandTimeout

`func (o *UpdateTexmlApplicationRequest) HasFirstCommandTimeout() bool`

HasFirstCommandTimeout returns a boolean if a field has been set.

### GetFirstCommandTimeoutSecs

`func (o *UpdateTexmlApplicationRequest) GetFirstCommandTimeoutSecs() int32`

GetFirstCommandTimeoutSecs returns the FirstCommandTimeoutSecs field if non-nil, zero value otherwise.

### GetFirstCommandTimeoutSecsOk

`func (o *UpdateTexmlApplicationRequest) GetFirstCommandTimeoutSecsOk() (*int32, bool)`

GetFirstCommandTimeoutSecsOk returns a tuple with the FirstCommandTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCommandTimeoutSecs

`func (o *UpdateTexmlApplicationRequest) SetFirstCommandTimeoutSecs(v int32)`

SetFirstCommandTimeoutSecs sets FirstCommandTimeoutSecs field to given value.

### HasFirstCommandTimeoutSecs

`func (o *UpdateTexmlApplicationRequest) HasFirstCommandTimeoutSecs() bool`

HasFirstCommandTimeoutSecs returns a boolean if a field has been set.

### GetVoiceUrl

`func (o *UpdateTexmlApplicationRequest) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *UpdateTexmlApplicationRequest) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *UpdateTexmlApplicationRequest) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.


### GetVoiceFallbackUrl

`func (o *UpdateTexmlApplicationRequest) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *UpdateTexmlApplicationRequest) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *UpdateTexmlApplicationRequest) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *UpdateTexmlApplicationRequest) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### GetVoiceMethod

`func (o *UpdateTexmlApplicationRequest) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *UpdateTexmlApplicationRequest) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *UpdateTexmlApplicationRequest) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *UpdateTexmlApplicationRequest) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### GetStatusCallback

`func (o *UpdateTexmlApplicationRequest) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *UpdateTexmlApplicationRequest) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *UpdateTexmlApplicationRequest) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *UpdateTexmlApplicationRequest) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *UpdateTexmlApplicationRequest) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *UpdateTexmlApplicationRequest) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *UpdateTexmlApplicationRequest) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *UpdateTexmlApplicationRequest) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetTags

`func (o *UpdateTexmlApplicationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateTexmlApplicationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateTexmlApplicationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateTexmlApplicationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInbound

`func (o *UpdateTexmlApplicationRequest) GetInbound() CreateTexmlApplicationRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateTexmlApplicationRequest) GetInboundOk() (*CreateTexmlApplicationRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateTexmlApplicationRequest) SetInbound(v CreateTexmlApplicationRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateTexmlApplicationRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateTexmlApplicationRequest) GetOutbound() CreateTexmlApplicationRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateTexmlApplicationRequest) GetOutboundOk() (*CreateTexmlApplicationRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateTexmlApplicationRequest) SetOutbound(v CreateTexmlApplicationRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateTexmlApplicationRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


