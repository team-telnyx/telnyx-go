# CredentialOutbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallParkingEnabled** | Pointer to **NullableBool** | Forces all SIP calls originated on this connection to be \&quot;parked\&quot; instead of \&quot;bridged\&quot; to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next. | [optional] [default to false]
**AniOverride** | Pointer to **string** | Set a phone number as the ani_override value to override caller id number on outbound calls. | [optional] [default to ""]
**AniOverrideType** | Pointer to **string** | Specifies when we apply your ani_override setting. Only applies when ani_override is not blank. | [optional] [default to "always"]
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of outbound calls to phone numbers associated with this connection. | [optional] 
**InstantRingbackEnabled** | Pointer to **bool** | When set, ringback will not wait for indication before sending ringback tone to calling party. | [optional] [default to true]
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**Localization** | Pointer to **string** | A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to &#x60;US&#x60; then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default. | [optional] [default to "US"]
**T38ReinviteSource** | Pointer to **string** | This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite.&lt;br/&gt;&lt;br/&gt; By default, Telnyx will send the re-invite. If set to &#x60;customer&#x60;, the caller is expected to send the t.38 reinvite. | [optional] [default to "telnyx"]
**OutboundVoiceProfileId** | Pointer to **string** | Identifies the associated outbound voice profile. | [optional] 

## Methods

### NewCredentialOutbound

`func NewCredentialOutbound() *CredentialOutbound`

NewCredentialOutbound instantiates a new CredentialOutbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialOutboundWithDefaults

`func NewCredentialOutboundWithDefaults() *CredentialOutbound`

NewCredentialOutboundWithDefaults instantiates a new CredentialOutbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallParkingEnabled

`func (o *CredentialOutbound) GetCallParkingEnabled() bool`

GetCallParkingEnabled returns the CallParkingEnabled field if non-nil, zero value otherwise.

### GetCallParkingEnabledOk

`func (o *CredentialOutbound) GetCallParkingEnabledOk() (*bool, bool)`

GetCallParkingEnabledOk returns a tuple with the CallParkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallParkingEnabled

`func (o *CredentialOutbound) SetCallParkingEnabled(v bool)`

SetCallParkingEnabled sets CallParkingEnabled field to given value.

### HasCallParkingEnabled

`func (o *CredentialOutbound) HasCallParkingEnabled() bool`

HasCallParkingEnabled returns a boolean if a field has been set.

### SetCallParkingEnabledNil

`func (o *CredentialOutbound) SetCallParkingEnabledNil(b bool)`

 SetCallParkingEnabledNil sets the value for CallParkingEnabled to be an explicit nil

### UnsetCallParkingEnabled
`func (o *CredentialOutbound) UnsetCallParkingEnabled()`

UnsetCallParkingEnabled ensures that no value is present for CallParkingEnabled, not even an explicit nil
### GetAniOverride

`func (o *CredentialOutbound) GetAniOverride() string`

GetAniOverride returns the AniOverride field if non-nil, zero value otherwise.

### GetAniOverrideOk

`func (o *CredentialOutbound) GetAniOverrideOk() (*string, bool)`

GetAniOverrideOk returns a tuple with the AniOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverride

`func (o *CredentialOutbound) SetAniOverride(v string)`

SetAniOverride sets AniOverride field to given value.

### HasAniOverride

`func (o *CredentialOutbound) HasAniOverride() bool`

HasAniOverride returns a boolean if a field has been set.

### GetAniOverrideType

`func (o *CredentialOutbound) GetAniOverrideType() string`

GetAniOverrideType returns the AniOverrideType field if non-nil, zero value otherwise.

### GetAniOverrideTypeOk

`func (o *CredentialOutbound) GetAniOverrideTypeOk() (*string, bool)`

GetAniOverrideTypeOk returns a tuple with the AniOverrideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverrideType

`func (o *CredentialOutbound) SetAniOverrideType(v string)`

SetAniOverrideType sets AniOverrideType field to given value.

### HasAniOverrideType

`func (o *CredentialOutbound) HasAniOverrideType() bool`

HasAniOverrideType returns a boolean if a field has been set.

### GetChannelLimit

`func (o *CredentialOutbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CredentialOutbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CredentialOutbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CredentialOutbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetInstantRingbackEnabled

`func (o *CredentialOutbound) GetInstantRingbackEnabled() bool`

GetInstantRingbackEnabled returns the InstantRingbackEnabled field if non-nil, zero value otherwise.

### GetInstantRingbackEnabledOk

`func (o *CredentialOutbound) GetInstantRingbackEnabledOk() (*bool, bool)`

GetInstantRingbackEnabledOk returns a tuple with the InstantRingbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRingbackEnabled

`func (o *CredentialOutbound) SetInstantRingbackEnabled(v bool)`

SetInstantRingbackEnabled sets InstantRingbackEnabled field to given value.

### HasInstantRingbackEnabled

`func (o *CredentialOutbound) HasInstantRingbackEnabled() bool`

HasInstantRingbackEnabled returns a boolean if a field has been set.

### GetGenerateRingbackTone

`func (o *CredentialOutbound) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *CredentialOutbound) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *CredentialOutbound) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *CredentialOutbound) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetLocalization

`func (o *CredentialOutbound) GetLocalization() string`

GetLocalization returns the Localization field if non-nil, zero value otherwise.

### GetLocalizationOk

`func (o *CredentialOutbound) GetLocalizationOk() (*string, bool)`

GetLocalizationOk returns a tuple with the Localization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalization

`func (o *CredentialOutbound) SetLocalization(v string)`

SetLocalization sets Localization field to given value.

### HasLocalization

`func (o *CredentialOutbound) HasLocalization() bool`

HasLocalization returns a boolean if a field has been set.

### GetT38ReinviteSource

`func (o *CredentialOutbound) GetT38ReinviteSource() string`

GetT38ReinviteSource returns the T38ReinviteSource field if non-nil, zero value otherwise.

### GetT38ReinviteSourceOk

`func (o *CredentialOutbound) GetT38ReinviteSourceOk() (*string, bool)`

GetT38ReinviteSourceOk returns a tuple with the T38ReinviteSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38ReinviteSource

`func (o *CredentialOutbound) SetT38ReinviteSource(v string)`

SetT38ReinviteSource sets T38ReinviteSource field to given value.

### HasT38ReinviteSource

`func (o *CredentialOutbound) HasT38ReinviteSource() bool`

HasT38ReinviteSource returns a boolean if a field has been set.

### GetOutboundVoiceProfileId

`func (o *CredentialOutbound) GetOutboundVoiceProfileId() string`

GetOutboundVoiceProfileId returns the OutboundVoiceProfileId field if non-nil, zero value otherwise.

### GetOutboundVoiceProfileIdOk

`func (o *CredentialOutbound) GetOutboundVoiceProfileIdOk() (*string, bool)`

GetOutboundVoiceProfileIdOk returns a tuple with the OutboundVoiceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundVoiceProfileId

`func (o *CredentialOutbound) SetOutboundVoiceProfileId(v string)`

SetOutboundVoiceProfileId sets OutboundVoiceProfileId field to given value.

### HasOutboundVoiceProfileId

`func (o *CredentialOutbound) HasOutboundVoiceProfileId() bool`

HasOutboundVoiceProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


