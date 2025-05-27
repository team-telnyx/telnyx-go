# CreateFaxApplicationRequestInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimit** | Pointer to **int32** | When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection. | [optional] 
**SipSubdomain** | Pointer to **string** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] 
**SipSubdomainReceiveSettings** | Pointer to **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] [default to "from_anyone"]

## Methods

### NewCreateFaxApplicationRequestInbound

`func NewCreateFaxApplicationRequestInbound() *CreateFaxApplicationRequestInbound`

NewCreateFaxApplicationRequestInbound instantiates a new CreateFaxApplicationRequestInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFaxApplicationRequestInboundWithDefaults

`func NewCreateFaxApplicationRequestInboundWithDefaults() *CreateFaxApplicationRequestInbound`

NewCreateFaxApplicationRequestInboundWithDefaults instantiates a new CreateFaxApplicationRequestInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimit

`func (o *CreateFaxApplicationRequestInbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CreateFaxApplicationRequestInbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CreateFaxApplicationRequestInbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CreateFaxApplicationRequestInbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetSipSubdomain

`func (o *CreateFaxApplicationRequestInbound) GetSipSubdomain() string`

GetSipSubdomain returns the SipSubdomain field if non-nil, zero value otherwise.

### GetSipSubdomainOk

`func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainOk() (*string, bool)`

GetSipSubdomainOk returns a tuple with the SipSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomain

`func (o *CreateFaxApplicationRequestInbound) SetSipSubdomain(v string)`

SetSipSubdomain sets SipSubdomain field to given value.

### HasSipSubdomain

`func (o *CreateFaxApplicationRequestInbound) HasSipSubdomain() bool`

HasSipSubdomain returns a boolean if a field has been set.

### GetSipSubdomainReceiveSettings

`func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainReceiveSettings() string`

GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field if non-nil, zero value otherwise.

### GetSipSubdomainReceiveSettingsOk

`func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainReceiveSettingsOk() (*string, bool)`

GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomainReceiveSettings

`func (o *CreateFaxApplicationRequestInbound) SetSipSubdomainReceiveSettings(v string)`

SetSipSubdomainReceiveSettings sets SipSubdomainReceiveSettings field to given value.

### HasSipSubdomainReceiveSettings

`func (o *CreateFaxApplicationRequestInbound) HasSipSubdomainReceiveSettings() bool`

HasSipSubdomainReceiveSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


