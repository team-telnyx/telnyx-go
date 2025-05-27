# ReferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SipAddress** | **string** | The SIP URI to which the call will be referred to. | 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE. | [optional] 
**SipAuthUsername** | Pointer to **string** | SIP Authentication username used for SIP challenges. | [optional] 
**SipAuthPassword** | Pointer to **string** | SIP Authentication password used for SIP challenges. | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | SIP headers to be added to the request. Currently only User-to-User header is supported. | [optional] 

## Methods

### NewReferRequest

`func NewReferRequest(sipAddress string, ) *ReferRequest`

NewReferRequest instantiates a new ReferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferRequestWithDefaults

`func NewReferRequestWithDefaults() *ReferRequest`

NewReferRequestWithDefaults instantiates a new ReferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSipAddress

`func (o *ReferRequest) GetSipAddress() string`

GetSipAddress returns the SipAddress field if non-nil, zero value otherwise.

### GetSipAddressOk

`func (o *ReferRequest) GetSipAddressOk() (*string, bool)`

GetSipAddressOk returns a tuple with the SipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAddress

`func (o *ReferRequest) SetSipAddress(v string)`

SetSipAddress sets SipAddress field to given value.


### GetClientState

`func (o *ReferRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ReferRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ReferRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *ReferRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *ReferRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *ReferRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *ReferRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *ReferRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *ReferRequest) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *ReferRequest) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *ReferRequest) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *ReferRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *ReferRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *ReferRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *ReferRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *ReferRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *ReferRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *ReferRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *ReferRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *ReferRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetSipHeaders

`func (o *ReferRequest) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *ReferRequest) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *ReferRequest) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *ReferRequest) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


