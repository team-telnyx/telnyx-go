# StartSiprecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorName** | Pointer to **string** | Name of configured SIPREC connector to be used. | [optional] 
**SipTransport** | Pointer to **string** | Specifies SIP transport protocol. | [optional] [default to "udp"]
**SiprecTrack** | Pointer to **string** | Specifies which track should be sent on siprec session. | [optional] [default to "both_tracks"]
**IncludeMetadataCustomHeaders** | Pointer to **bool** | When set, custom parameters will be added as metadata (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip headers. | [optional] 
**Secure** | Pointer to **bool** | Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set you need to configure SRS port in your connector to 5061. | [optional] 
**SessionTimeoutSecs** | Pointer to **int32** | Sets &#x60;Session-Expires&#x60; header to the INVITE. A reinvite is sent every half the value set. Usefull for session keep alive. Minimum value is 90, set to 0 to disable. | [optional] [default to 1800]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 

## Methods

### NewStartSiprecRequest

`func NewStartSiprecRequest() *StartSiprecRequest`

NewStartSiprecRequest instantiates a new StartSiprecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartSiprecRequestWithDefaults

`func NewStartSiprecRequestWithDefaults() *StartSiprecRequest`

NewStartSiprecRequestWithDefaults instantiates a new StartSiprecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorName

`func (o *StartSiprecRequest) GetConnectorName() string`

GetConnectorName returns the ConnectorName field if non-nil, zero value otherwise.

### GetConnectorNameOk

`func (o *StartSiprecRequest) GetConnectorNameOk() (*string, bool)`

GetConnectorNameOk returns a tuple with the ConnectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorName

`func (o *StartSiprecRequest) SetConnectorName(v string)`

SetConnectorName sets ConnectorName field to given value.

### HasConnectorName

`func (o *StartSiprecRequest) HasConnectorName() bool`

HasConnectorName returns a boolean if a field has been set.

### GetSipTransport

`func (o *StartSiprecRequest) GetSipTransport() string`

GetSipTransport returns the SipTransport field if non-nil, zero value otherwise.

### GetSipTransportOk

`func (o *StartSiprecRequest) GetSipTransportOk() (*string, bool)`

GetSipTransportOk returns a tuple with the SipTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTransport

`func (o *StartSiprecRequest) SetSipTransport(v string)`

SetSipTransport sets SipTransport field to given value.

### HasSipTransport

`func (o *StartSiprecRequest) HasSipTransport() bool`

HasSipTransport returns a boolean if a field has been set.

### GetSiprecTrack

`func (o *StartSiprecRequest) GetSiprecTrack() string`

GetSiprecTrack returns the SiprecTrack field if non-nil, zero value otherwise.

### GetSiprecTrackOk

`func (o *StartSiprecRequest) GetSiprecTrackOk() (*string, bool)`

GetSiprecTrackOk returns a tuple with the SiprecTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiprecTrack

`func (o *StartSiprecRequest) SetSiprecTrack(v string)`

SetSiprecTrack sets SiprecTrack field to given value.

### HasSiprecTrack

`func (o *StartSiprecRequest) HasSiprecTrack() bool`

HasSiprecTrack returns a boolean if a field has been set.

### GetIncludeMetadataCustomHeaders

`func (o *StartSiprecRequest) GetIncludeMetadataCustomHeaders() bool`

GetIncludeMetadataCustomHeaders returns the IncludeMetadataCustomHeaders field if non-nil, zero value otherwise.

### GetIncludeMetadataCustomHeadersOk

`func (o *StartSiprecRequest) GetIncludeMetadataCustomHeadersOk() (*bool, bool)`

GetIncludeMetadataCustomHeadersOk returns a tuple with the IncludeMetadataCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetadataCustomHeaders

`func (o *StartSiprecRequest) SetIncludeMetadataCustomHeaders(v bool)`

SetIncludeMetadataCustomHeaders sets IncludeMetadataCustomHeaders field to given value.

### HasIncludeMetadataCustomHeaders

`func (o *StartSiprecRequest) HasIncludeMetadataCustomHeaders() bool`

HasIncludeMetadataCustomHeaders returns a boolean if a field has been set.

### GetSecure

`func (o *StartSiprecRequest) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *StartSiprecRequest) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *StartSiprecRequest) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *StartSiprecRequest) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetSessionTimeoutSecs

`func (o *StartSiprecRequest) GetSessionTimeoutSecs() int32`

GetSessionTimeoutSecs returns the SessionTimeoutSecs field if non-nil, zero value otherwise.

### GetSessionTimeoutSecsOk

`func (o *StartSiprecRequest) GetSessionTimeoutSecsOk() (*int32, bool)`

GetSessionTimeoutSecsOk returns a tuple with the SessionTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeoutSecs

`func (o *StartSiprecRequest) SetSessionTimeoutSecs(v int32)`

SetSessionTimeoutSecs sets SessionTimeoutSecs field to given value.

### HasSessionTimeoutSecs

`func (o *StartSiprecRequest) HasSessionTimeoutSecs() bool`

HasSessionTimeoutSecs returns a boolean if a field has been set.

### GetClientState

`func (o *StartSiprecRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StartSiprecRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StartSiprecRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StartSiprecRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


