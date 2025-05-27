# ConferenceRecordingSavedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Participant&#39;s call ID used to issue commands via Call Control API. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**Channels** | Pointer to **string** | Whether recording was recorded in &#x60;single&#x60; or &#x60;dual&#x60; channel. | [optional] 
**ConferenceId** | Pointer to **string** | ID of the conference that is being recorded. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**Format** | Pointer to **string** | The audio file format used when storing the call recording. Can be either &#x60;mp3&#x60; or &#x60;wav&#x60;. | [optional] 
**PublicRecordingUrls** | Pointer to [**CallRecordingSavedPayloadPublicRecordingUrls**](CallRecordingSavedPayloadPublicRecordingUrls.md) |  | [optional] 
**RecordingEndedAt** | Pointer to **time.Time** | ISO 8601 datetime of when recording ended. | [optional] 
**RecordingId** | Pointer to **string** | ID of the conference recording. | [optional] 
**RecordingStartedAt** | Pointer to **time.Time** | ISO 8601 datetime of when recording started. | [optional] 
**RecordingUrls** | Pointer to [**CallRecordingSavedPayloadRecordingUrls**](CallRecordingSavedPayloadRecordingUrls.md) |  | [optional] 

## Methods

### NewConferenceRecordingSavedPayload

`func NewConferenceRecordingSavedPayload() *ConferenceRecordingSavedPayload`

NewConferenceRecordingSavedPayload instantiates a new ConferenceRecordingSavedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceRecordingSavedPayloadWithDefaults

`func NewConferenceRecordingSavedPayloadWithDefaults() *ConferenceRecordingSavedPayload`

NewConferenceRecordingSavedPayloadWithDefaults instantiates a new ConferenceRecordingSavedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *ConferenceRecordingSavedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *ConferenceRecordingSavedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *ConferenceRecordingSavedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *ConferenceRecordingSavedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceRecordingSavedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceRecordingSavedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceRecordingSavedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceRecordingSavedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *ConferenceRecordingSavedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ConferenceRecordingSavedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ConferenceRecordingSavedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *ConferenceRecordingSavedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetChannels

`func (o *ConferenceRecordingSavedPayload) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ConferenceRecordingSavedPayload) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ConferenceRecordingSavedPayload) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ConferenceRecordingSavedPayload) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceRecordingSavedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceRecordingSavedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceRecordingSavedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceRecordingSavedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConferenceRecordingSavedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceRecordingSavedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceRecordingSavedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceRecordingSavedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetFormat

`func (o *ConferenceRecordingSavedPayload) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ConferenceRecordingSavedPayload) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ConferenceRecordingSavedPayload) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ConferenceRecordingSavedPayload) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPublicRecordingUrls

`func (o *ConferenceRecordingSavedPayload) GetPublicRecordingUrls() CallRecordingSavedPayloadPublicRecordingUrls`

GetPublicRecordingUrls returns the PublicRecordingUrls field if non-nil, zero value otherwise.

### GetPublicRecordingUrlsOk

`func (o *ConferenceRecordingSavedPayload) GetPublicRecordingUrlsOk() (*CallRecordingSavedPayloadPublicRecordingUrls, bool)`

GetPublicRecordingUrlsOk returns a tuple with the PublicRecordingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRecordingUrls

`func (o *ConferenceRecordingSavedPayload) SetPublicRecordingUrls(v CallRecordingSavedPayloadPublicRecordingUrls)`

SetPublicRecordingUrls sets PublicRecordingUrls field to given value.

### HasPublicRecordingUrls

`func (o *ConferenceRecordingSavedPayload) HasPublicRecordingUrls() bool`

HasPublicRecordingUrls returns a boolean if a field has been set.

### GetRecordingEndedAt

`func (o *ConferenceRecordingSavedPayload) GetRecordingEndedAt() time.Time`

GetRecordingEndedAt returns the RecordingEndedAt field if non-nil, zero value otherwise.

### GetRecordingEndedAtOk

`func (o *ConferenceRecordingSavedPayload) GetRecordingEndedAtOk() (*time.Time, bool)`

GetRecordingEndedAtOk returns a tuple with the RecordingEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEndedAt

`func (o *ConferenceRecordingSavedPayload) SetRecordingEndedAt(v time.Time)`

SetRecordingEndedAt sets RecordingEndedAt field to given value.

### HasRecordingEndedAt

`func (o *ConferenceRecordingSavedPayload) HasRecordingEndedAt() bool`

HasRecordingEndedAt returns a boolean if a field has been set.

### GetRecordingId

`func (o *ConferenceRecordingSavedPayload) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *ConferenceRecordingSavedPayload) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *ConferenceRecordingSavedPayload) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *ConferenceRecordingSavedPayload) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.

### GetRecordingStartedAt

`func (o *ConferenceRecordingSavedPayload) GetRecordingStartedAt() time.Time`

GetRecordingStartedAt returns the RecordingStartedAt field if non-nil, zero value otherwise.

### GetRecordingStartedAtOk

`func (o *ConferenceRecordingSavedPayload) GetRecordingStartedAtOk() (*time.Time, bool)`

GetRecordingStartedAtOk returns a tuple with the RecordingStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStartedAt

`func (o *ConferenceRecordingSavedPayload) SetRecordingStartedAt(v time.Time)`

SetRecordingStartedAt sets RecordingStartedAt field to given value.

### HasRecordingStartedAt

`func (o *ConferenceRecordingSavedPayload) HasRecordingStartedAt() bool`

HasRecordingStartedAt returns a boolean if a field has been set.

### GetRecordingUrls

`func (o *ConferenceRecordingSavedPayload) GetRecordingUrls() CallRecordingSavedPayloadRecordingUrls`

GetRecordingUrls returns the RecordingUrls field if non-nil, zero value otherwise.

### GetRecordingUrlsOk

`func (o *ConferenceRecordingSavedPayload) GetRecordingUrlsOk() (*CallRecordingSavedPayloadRecordingUrls, bool)`

GetRecordingUrlsOk returns a tuple with the RecordingUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingUrls

`func (o *ConferenceRecordingSavedPayload) SetRecordingUrls(v CallRecordingSavedPayloadRecordingUrls)`

SetRecordingUrls sets RecordingUrls field to given value.

### HasRecordingUrls

`func (o *ConferenceRecordingSavedPayload) HasRecordingUrls() bool`

HasRecordingUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


