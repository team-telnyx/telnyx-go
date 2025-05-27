# StartForkingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rx** | Pointer to **string** | The network target, &lt;udp:ip_address:port&gt;, where the call&#39;s incoming RTP media packets should be forwarded. | [optional] 
**StreamType** | Pointer to **string** | Optionally specify a media type to stream. If &#x60;decrypted&#x60; selected, Telnyx will decrypt incoming SIP media before forking to the target. &#x60;rx&#x60; and &#x60;tx&#x60; are required fields if &#x60;decrypted&#x60; selected. | [optional] [default to "decrypted"]
**Tx** | Pointer to **string** | The network target, &lt;udp:ip_address:port&gt;, where the call&#39;s outgoing RTP media packets should be forwarded. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewStartForkingRequest

`func NewStartForkingRequest() *StartForkingRequest`

NewStartForkingRequest instantiates a new StartForkingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartForkingRequestWithDefaults

`func NewStartForkingRequestWithDefaults() *StartForkingRequest`

NewStartForkingRequestWithDefaults instantiates a new StartForkingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRx

`func (o *StartForkingRequest) GetRx() string`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *StartForkingRequest) GetRxOk() (*string, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *StartForkingRequest) SetRx(v string)`

SetRx sets Rx field to given value.

### HasRx

`func (o *StartForkingRequest) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetStreamType

`func (o *StartForkingRequest) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *StartForkingRequest) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *StartForkingRequest) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *StartForkingRequest) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetTx

`func (o *StartForkingRequest) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *StartForkingRequest) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *StartForkingRequest) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *StartForkingRequest) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetClientState

`func (o *StartForkingRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StartForkingRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StartForkingRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StartForkingRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *StartForkingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StartForkingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StartForkingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StartForkingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


