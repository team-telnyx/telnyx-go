# CallRecordingSavedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**RecordingStartedAt** | Pointer to **time.Time** | ISO 8601 datetime of when recording started. | [optional] 
**RecordingEndedAt** | Pointer to **time.Time** | ISO 8601 datetime of when recording ended. | [optional] 
**Channels** | Pointer to **string** | Whether recording was recorded in &#x60;single&#x60; or &#x60;dual&#x60; channel. | [optional] 
**RecordingUrls** | Pointer to [**CallRecordingSavedPayloadRecordingUrls**](CallRecordingSavedPayloadRecordingUrls.md) |  | [optional] 
**PublicRecordingUrls** | Pointer to [**CallRecordingSavedPayloadPublicRecordingUrls**](CallRecordingSavedPayloadPublicRecordingUrls.md) |  | [optional] 

## Methods

### NewCallRecordingSavedPayload

`func NewCallRecordingSavedPayload() *CallRecordingSavedPayload`

NewCallRecordingSavedPayload instantiates a new CallRecordingSavedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingSavedPayloadWithDefaults

`func NewCallRecordingSavedPayloadWithDefaults() *CallRecordingSavedPayload`

NewCallRecordingSavedPayloadWithDefaults instantiates a new CallRecordingSavedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallLegId

`func (o *CallRecordingSavedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallRecordingSavedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallRecordingSavedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallRecordingSavedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallRecordingSavedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallRecordingSavedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallRecordingSavedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallRecordingSavedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallRecordingSavedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallRecordingSavedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallRecordingSavedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallRecordingSavedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallRecordingSavedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallRecordingSavedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallRecordingSavedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallRecordingSavedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetRecordingStartedAt

`func (o *CallRecordingSavedPayload) GetRecordingStartedAt() time.Time`

GetRecordingStartedAt returns the RecordingStartedAt field if non-nil, zero value otherwise.

### GetRecordingStartedAtOk

`func (o *CallRecordingSavedPayload) GetRecordingStartedAtOk() (*time.Time, bool)`

GetRecordingStartedAtOk returns a tuple with the RecordingStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStartedAt

`func (o *CallRecordingSavedPayload) SetRecordingStartedAt(v time.Time)`

SetRecordingStartedAt sets RecordingStartedAt field to given value.

### HasRecordingStartedAt

`func (o *CallRecordingSavedPayload) HasRecordingStartedAt() bool`

HasRecordingStartedAt returns a boolean if a field has been set.

### GetRecordingEndedAt

`func (o *CallRecordingSavedPayload) GetRecordingEndedAt() time.Time`

GetRecordingEndedAt returns the RecordingEndedAt field if non-nil, zero value otherwise.

### GetRecordingEndedAtOk

`func (o *CallRecordingSavedPayload) GetRecordingEndedAtOk() (*time.Time, bool)`

GetRecordingEndedAtOk returns a tuple with the RecordingEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEndedAt

`func (o *CallRecordingSavedPayload) SetRecordingEndedAt(v time.Time)`

SetRecordingEndedAt sets RecordingEndedAt field to given value.

### HasRecordingEndedAt

`func (o *CallRecordingSavedPayload) HasRecordingEndedAt() bool`

HasRecordingEndedAt returns a boolean if a field has been set.

### GetChannels

`func (o *CallRecordingSavedPayload) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CallRecordingSavedPayload) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CallRecordingSavedPayload) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CallRecordingSavedPayload) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetRecordingUrls

`func (o *CallRecordingSavedPayload) GetRecordingUrls() CallRecordingSavedPayloadRecordingUrls`

GetRecordingUrls returns the RecordingUrls field if non-nil, zero value otherwise.

### GetRecordingUrlsOk

`func (o *CallRecordingSavedPayload) GetRecordingUrlsOk() (*CallRecordingSavedPayloadRecordingUrls, bool)`

GetRecordingUrlsOk returns a tuple with the RecordingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingUrls

`func (o *CallRecordingSavedPayload) SetRecordingUrls(v CallRecordingSavedPayloadRecordingUrls)`

SetRecordingUrls sets RecordingUrls field to given value.

### HasRecordingUrls

`func (o *CallRecordingSavedPayload) HasRecordingUrls() bool`

HasRecordingUrls returns a boolean if a field has been set.

### GetPublicRecordingUrls

`func (o *CallRecordingSavedPayload) GetPublicRecordingUrls() CallRecordingSavedPayloadPublicRecordingUrls`

GetPublicRecordingUrls returns the PublicRecordingUrls field if non-nil, zero value otherwise.

### GetPublicRecordingUrlsOk

`func (o *CallRecordingSavedPayload) GetPublicRecordingUrlsOk() (*CallRecordingSavedPayloadPublicRecordingUrls, bool)`

GetPublicRecordingUrlsOk returns a tuple with the PublicRecordingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRecordingUrls

`func (o *CallRecordingSavedPayload) SetPublicRecordingUrls(v CallRecordingSavedPayloadPublicRecordingUrls)`

SetPublicRecordingUrls sets PublicRecordingUrls field to given value.

### HasPublicRecordingUrls

`func (o *CallRecordingSavedPayload) HasPublicRecordingUrls() bool`

HasPublicRecordingUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


