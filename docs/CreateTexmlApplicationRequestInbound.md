# CreateTexmlApplicationRequestInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] 
**ShakenStirEnabled** | Pointer to **bool** | When enabled Telnyx will include Shaken/Stir data in the Webhook for new inbound calls. | [optional] [default to false]
**SipSubdomain** | Pointer to **string** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] 
**SipSubdomainReceiveSettings** | Pointer to **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] [default to "from_anyone"]

## Methods

### NewCreateTexmlApplicationRequestInbound

`func NewCreateTexmlApplicationRequestInbound() *CreateTexmlApplicationRequestInbound`

NewCreateTexmlApplicationRequestInbound instantiates a new CreateTexmlApplicationRequestInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTexmlApplicationRequestInboundWithDefaults

`func NewCreateTexmlApplicationRequestInboundWithDefaults() *CreateTexmlApplicationRequestInbound`

NewCreateTexmlApplicationRequestInboundWithDefaults instantiates a new CreateTexmlApplicationRequestInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *CreateTexmlApplicationRequestInbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CreateTexmlApplicationRequestInbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CreateTexmlApplicationRequestInbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CreateTexmlApplicationRequestInbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetShakenStirEnabled

`func (o *CreateTexmlApplicationRequestInbound) GetShakenStirEnabled() bool`

GetShakenStirEnabled returns the ShakenStirEnabled field if non-nil, zero value otherwise.

### GetShakenStirEnabledOk

`func (o *CreateTexmlApplicationRequestInbound) GetShakenStirEnabledOk() (*bool, bool)`

GetShakenStirEnabledOk returns a tuple with the ShakenStirEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirEnabled

`func (o *CreateTexmlApplicationRequestInbound) SetShakenStirEnabled(v bool)`

SetShakenStirEnabled sets ShakenStirEnabled field to given value.

### HasShakenStirEnabled

`func (o *CreateTexmlApplicationRequestInbound) HasShakenStirEnabled() bool`

HasShakenStirEnabled returns a boolean if a field has been set.

### GetSipSubdomain

`func (o *CreateTexmlApplicationRequestInbound) GetSipSubdomain() string`

GetSipSubdomain returns the SipSubdomain field if non-nil, zero value otherwise.

### GetSipSubdomainOk

`func (o *CreateTexmlApplicationRequestInbound) GetSipSubdomainOk() (*string, bool)`

GetSipSubdomainOk returns a tuple with the SipSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomain

`func (o *CreateTexmlApplicationRequestInbound) SetSipSubdomain(v string)`

SetSipSubdomain sets SipSubdomain field to given value.

### HasSipSubdomain

`func (o *CreateTexmlApplicationRequestInbound) HasSipSubdomain() bool`

HasSipSubdomain returns a boolean if a field has been set.

### GetSipSubdomainReceiveSettings

`func (o *CreateTexmlApplicationRequestInbound) GetSipSubdomainReceiveSettings() string`

GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field if non-nil, zero value otherwise.

### GetSipSubdomainReceiveSettingsOk

`func (o *CreateTexmlApplicationRequestInbound) GetSipSubdomainReceiveSettingsOk() (*string, bool)`

GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomainReceiveSettings

`func (o *CreateTexmlApplicationRequestInbound) SetSipSubdomainReceiveSettings(v string)`

SetSipSubdomainReceiveSettings sets SipSubdomainReceiveSettings field to given value.

### HasSipSubdomainReceiveSettings

`func (o *CreateTexmlApplicationRequestInbound) HasSipSubdomainReceiveSettings() bool`

HasSipSubdomainReceiveSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


