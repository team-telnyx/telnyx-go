# RoomComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the room composition. | [optional] 
**RoomId** | Pointer to **string** | Identify the room associated with the room composition. | [optional] 
**SessionId** | Pointer to **string** | Identify the room session associated with the room composition. | [optional] 
**UserId** | Pointer to **string** | Identify the user associated with the room composition. | [optional] 
**Status** | Pointer to **string** | Shows the room composition status. | [optional] 
**SizeMb** | Pointer to **float32** | Shows the room composition size in MB. | [optional] 
**DownloadUrl** | Pointer to **string** | Url to download the composition. | [optional] 
**DurationSecs** | Pointer to **int32** | Shows the room composition duration in seconds. | [optional] 
**Format** | Pointer to **string** | Shows format of the room composition. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 timestamp when the room composition was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 timestamp when the room composition was updated. | [optional] 
**EndedAt** | Pointer to **string** | ISO 8601 timestamp when the room composition has ended. | [optional] 
**StartedAt** | Pointer to **string** | ISO 8601 timestamp when the room composition has stated. | [optional] 
**CompletedAt** | Pointer to **string** | ISO 8601 timestamp when the room composition has completed. | [optional] 
**VideoLayout** | Pointer to [**map[string]VideoRegion**](VideoRegion.md) | Describes the video layout of the room composition in terms of regions. Limited to 2 regions. | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this room composition will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this room composition will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewRoomComposition

`func NewRoomComposition() *RoomComposition`

NewRoomComposition instantiates a new RoomComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomCompositionWithDefaults

`func NewRoomCompositionWithDefaults() *RoomComposition`

NewRoomCompositionWithDefaults instantiates a new RoomComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoomComposition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomComposition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomComposition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoomComposition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRoomId

`func (o *RoomComposition) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *RoomComposition) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *RoomComposition) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *RoomComposition) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetSessionId

`func (o *RoomComposition) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *RoomComposition) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *RoomComposition) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *RoomComposition) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUserId

`func (o *RoomComposition) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RoomComposition) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RoomComposition) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RoomComposition) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *RoomComposition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoomComposition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoomComposition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoomComposition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSizeMb

`func (o *RoomComposition) GetSizeMb() float32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *RoomComposition) GetSizeMbOk() (*float32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *RoomComposition) SetSizeMb(v float32)`

SetSizeMb sets SizeMb field to given value.

### HasSizeMb

`func (o *RoomComposition) HasSizeMb() bool`

HasSizeMb returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *RoomComposition) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *RoomComposition) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *RoomComposition) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *RoomComposition) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetDurationSecs

`func (o *RoomComposition) GetDurationSecs() int32`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *RoomComposition) GetDurationSecsOk() (*int32, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *RoomComposition) SetDurationSecs(v int32)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *RoomComposition) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetFormat

`func (o *RoomComposition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RoomComposition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RoomComposition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RoomComposition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RoomComposition) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoomComposition) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoomComposition) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RoomComposition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoomComposition) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoomComposition) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoomComposition) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoomComposition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *RoomComposition) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RoomComposition) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RoomComposition) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RoomComposition) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RoomComposition) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RoomComposition) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RoomComposition) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RoomComposition) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *RoomComposition) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *RoomComposition) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *RoomComposition) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *RoomComposition) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetVideoLayout

`func (o *RoomComposition) GetVideoLayout() map[string]VideoRegion`

GetVideoLayout returns the VideoLayout field if non-nil, zero value otherwise.

### GetVideoLayoutOk

`func (o *RoomComposition) GetVideoLayoutOk() (*map[string]VideoRegion, bool)`

GetVideoLayoutOk returns a tuple with the VideoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLayout

`func (o *RoomComposition) SetVideoLayout(v map[string]VideoRegion)`

SetVideoLayout sets VideoLayout field to given value.

### HasVideoLayout

`func (o *RoomComposition) HasVideoLayout() bool`

HasVideoLayout returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *RoomComposition) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *RoomComposition) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *RoomComposition) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *RoomComposition) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *RoomComposition) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *RoomComposition) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *RoomComposition) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *RoomComposition) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *RoomComposition) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *RoomComposition) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *RoomComposition) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *RoomComposition) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *RoomComposition) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *RoomComposition) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *RoomComposition) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *RoomComposition) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetRecordType

`func (o *RoomComposition) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RoomComposition) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RoomComposition) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RoomComposition) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


