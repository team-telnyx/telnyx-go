# CallInitiatedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**ConnectionCodecs** | Pointer to **string** | The list of comma-separated codecs enabled for the connection. | [optional] 
**OfferedCodecs** | Pointer to **string** | The list of comma-separated codecs offered by caller. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers from sip invite | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | User-to-User and Diversion headers from sip invite. | [optional] 
**ShakenStirAttestation** | Pointer to **string** | SHAKEN/STIR attestation level. | [optional] 
**ShakenStirValidated** | Pointer to **bool** | Whether attestation was successfully validated or not. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CallerIdName** | Pointer to **string** | Caller id. | [optional] 
**CallScreeningResult** | Pointer to **string** | Call screening result. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 
**Direction** | Pointer to **string** | Whether the call is &#x60;incoming&#x60; or &#x60;outgoing&#x60;. | [optional] 
**State** | Pointer to **string** | State received from a command. | [optional] 
**StartTime** | Pointer to **time.Time** | ISO 8601 datetime of when the call started. | [optional] 
**Tags** | Pointer to **[]string** | Array of tags associated to number. | [optional] 

## Methods

### NewCallInitiatedPayload

`func NewCallInitiatedPayload() *CallInitiatedPayload`

NewCallInitiatedPayload instantiates a new CallInitiatedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallInitiatedPayloadWithDefaults

`func NewCallInitiatedPayloadWithDefaults() *CallInitiatedPayload`

NewCallInitiatedPayloadWithDefaults instantiates a new CallInitiatedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallInitiatedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallInitiatedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallInitiatedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallInitiatedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallInitiatedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallInitiatedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallInitiatedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallInitiatedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConnectionCodecs

`func (o *CallInitiatedPayload) GetConnectionCodecs() string`

GetConnectionCodecs returns the ConnectionCodecs field if non-nil, zero value otherwise.

### GetConnectionCodecsOk

`func (o *CallInitiatedPayload) GetConnectionCodecsOk() (*string, bool)`

GetConnectionCodecsOk returns a tuple with the ConnectionCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCodecs

`func (o *CallInitiatedPayload) SetConnectionCodecs(v string)`

SetConnectionCodecs sets ConnectionCodecs field to given value.

### HasConnectionCodecs

`func (o *CallInitiatedPayload) HasConnectionCodecs() bool`

HasConnectionCodecs returns a boolean if a field has been set.

### GetOfferedCodecs

`func (o *CallInitiatedPayload) GetOfferedCodecs() string`

GetOfferedCodecs returns the OfferedCodecs field if non-nil, zero value otherwise.

### GetOfferedCodecsOk

`func (o *CallInitiatedPayload) GetOfferedCodecsOk() (*string, bool)`

GetOfferedCodecsOk returns a tuple with the OfferedCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferedCodecs

`func (o *CallInitiatedPayload) SetOfferedCodecs(v string)`

SetOfferedCodecs sets OfferedCodecs field to given value.

### HasOfferedCodecs

`func (o *CallInitiatedPayload) HasOfferedCodecs() bool`

HasOfferedCodecs returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallInitiatedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallInitiatedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallInitiatedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallInitiatedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CallInitiatedPayload) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CallInitiatedPayload) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CallInitiatedPayload) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CallInitiatedPayload) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetSipHeaders

`func (o *CallInitiatedPayload) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *CallInitiatedPayload) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *CallInitiatedPayload) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *CallInitiatedPayload) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetShakenStirAttestation

`func (o *CallInitiatedPayload) GetShakenStirAttestation() string`

GetShakenStirAttestation returns the ShakenStirAttestation field if non-nil, zero value otherwise.

### GetShakenStirAttestationOk

`func (o *CallInitiatedPayload) GetShakenStirAttestationOk() (*string, bool)`

GetShakenStirAttestationOk returns a tuple with the ShakenStirAttestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirAttestation

`func (o *CallInitiatedPayload) SetShakenStirAttestation(v string)`

SetShakenStirAttestation sets ShakenStirAttestation field to given value.

### HasShakenStirAttestation

`func (o *CallInitiatedPayload) HasShakenStirAttestation() bool`

HasShakenStirAttestation returns a boolean if a field has been set.

### GetShakenStirValidated

`func (o *CallInitiatedPayload) GetShakenStirValidated() bool`

GetShakenStirValidated returns the ShakenStirValidated field if non-nil, zero value otherwise.

### GetShakenStirValidatedOk

`func (o *CallInitiatedPayload) GetShakenStirValidatedOk() (*bool, bool)`

GetShakenStirValidatedOk returns a tuple with the ShakenStirValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirValidated

`func (o *CallInitiatedPayload) SetShakenStirValidated(v bool)`

SetShakenStirValidated sets ShakenStirValidated field to given value.

### HasShakenStirValidated

`func (o *CallInitiatedPayload) HasShakenStirValidated() bool`

HasShakenStirValidated returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallInitiatedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallInitiatedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallInitiatedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallInitiatedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallInitiatedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallInitiatedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallInitiatedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallInitiatedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCallerIdName

`func (o *CallInitiatedPayload) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *CallInitiatedPayload) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *CallInitiatedPayload) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *CallInitiatedPayload) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallScreeningResult

`func (o *CallInitiatedPayload) GetCallScreeningResult() string`

GetCallScreeningResult returns the CallScreeningResult field if non-nil, zero value otherwise.

### GetCallScreeningResultOk

`func (o *CallInitiatedPayload) GetCallScreeningResultOk() (*string, bool)`

GetCallScreeningResultOk returns a tuple with the CallScreeningResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallScreeningResult

`func (o *CallInitiatedPayload) SetCallScreeningResult(v string)`

SetCallScreeningResult sets CallScreeningResult field to given value.

### HasCallScreeningResult

`func (o *CallInitiatedPayload) HasCallScreeningResult() bool`

HasCallScreeningResult returns a boolean if a field has been set.

### GetFrom

`func (o *CallInitiatedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallInitiatedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallInitiatedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallInitiatedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallInitiatedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallInitiatedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallInitiatedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallInitiatedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetDirection

`func (o *CallInitiatedPayload) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CallInitiatedPayload) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CallInitiatedPayload) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CallInitiatedPayload) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetState

`func (o *CallInitiatedPayload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CallInitiatedPayload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CallInitiatedPayload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CallInitiatedPayload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartTime

`func (o *CallInitiatedPayload) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallInitiatedPayload) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallInitiatedPayload) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallInitiatedPayload) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTags

`func (o *CallInitiatedPayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CallInitiatedPayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CallInitiatedPayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CallInitiatedPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


