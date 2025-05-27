# CreateRoomCompositionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** | The desired format of the room composition. | [optional] [default to "mp4"]
**Resolution** | Pointer to **NullableString** | The desired resolution (width/height in pixels) of the resulting video of the room composition. Both width and height are required to be between 16 and 1280; and width * height should not exceed 1280 * 720 | [optional] [default to "1280x720"]
**SessionId** | Pointer to **NullableString** | id of the room session associated with the room composition. | [optional] 
**VideoLayout** | Pointer to [**map[string]VideoRegion**](VideoRegion.md) | Describes the video layout of the room composition in terms of regions. | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this room composition will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this room composition will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 

## Methods

### NewCreateRoomCompositionRequest

`func NewCreateRoomCompositionRequest() *CreateRoomCompositionRequest`

NewCreateRoomCompositionRequest instantiates a new CreateRoomCompositionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoomCompositionRequestWithDefaults

`func NewCreateRoomCompositionRequestWithDefaults() *CreateRoomCompositionRequest`

NewCreateRoomCompositionRequestWithDefaults instantiates a new CreateRoomCompositionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *CreateRoomCompositionRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateRoomCompositionRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateRoomCompositionRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateRoomCompositionRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *CreateRoomCompositionRequest) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *CreateRoomCompositionRequest) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetResolution

`func (o *CreateRoomCompositionRequest) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *CreateRoomCompositionRequest) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *CreateRoomCompositionRequest) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *CreateRoomCompositionRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *CreateRoomCompositionRequest) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *CreateRoomCompositionRequest) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSessionId

`func (o *CreateRoomCompositionRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CreateRoomCompositionRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CreateRoomCompositionRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CreateRoomCompositionRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *CreateRoomCompositionRequest) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *CreateRoomCompositionRequest) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetVideoLayout

`func (o *CreateRoomCompositionRequest) GetVideoLayout() map[string]VideoRegion`

GetVideoLayout returns the VideoLayout field if non-nil, zero value otherwise.

### GetVideoLayoutOk

`func (o *CreateRoomCompositionRequest) GetVideoLayoutOk() (*map[string]VideoRegion, bool)`

GetVideoLayoutOk returns a tuple with the VideoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLayout

`func (o *CreateRoomCompositionRequest) SetVideoLayout(v map[string]VideoRegion)`

SetVideoLayout sets VideoLayout field to given value.

### HasVideoLayout

`func (o *CreateRoomCompositionRequest) HasVideoLayout() bool`

HasVideoLayout returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *CreateRoomCompositionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CreateRoomCompositionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CreateRoomCompositionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CreateRoomCompositionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CreateRoomCompositionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CreateRoomCompositionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CreateRoomCompositionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CreateRoomCompositionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CreateRoomCompositionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CreateRoomCompositionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *CreateRoomCompositionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CreateRoomCompositionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CreateRoomCompositionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CreateRoomCompositionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CreateRoomCompositionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CreateRoomCompositionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


