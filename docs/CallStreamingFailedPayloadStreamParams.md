# CallStreamingFailedPayloadStreamParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamUrl** | Pointer to **string** | The destination WebSocket address where the stream is going to be delivered. | [optional] 
**Track** | Pointer to **string** | Specifies which track should be streamed. | [optional] [default to "inbound_track"]

## Methods

### NewCallStreamingFailedPayloadStreamParams

`func NewCallStreamingFailedPayloadStreamParams() *CallStreamingFailedPayloadStreamParams`

NewCallStreamingFailedPayloadStreamParams instantiates a new CallStreamingFailedPayloadStreamParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStreamingFailedPayloadStreamParamsWithDefaults

`func NewCallStreamingFailedPayloadStreamParamsWithDefaults() *CallStreamingFailedPayloadStreamParams`

NewCallStreamingFailedPayloadStreamParamsWithDefaults instantiates a new CallStreamingFailedPayloadStreamParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamUrl

`func (o *CallStreamingFailedPayloadStreamParams) GetStreamUrl() string`

GetStreamUrl returns the StreamUrl field if non-nil, zero value otherwise.

### GetStreamUrlOk

`func (o *CallStreamingFailedPayloadStreamParams) GetStreamUrlOk() (*string, bool)`

GetStreamUrlOk returns a tuple with the StreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamUrl

`func (o *CallStreamingFailedPayloadStreamParams) SetStreamUrl(v string)`

SetStreamUrl sets StreamUrl field to given value.

### HasStreamUrl

`func (o *CallStreamingFailedPayloadStreamParams) HasStreamUrl() bool`

HasStreamUrl returns a boolean if a field has been set.

### GetTrack

`func (o *CallStreamingFailedPayloadStreamParams) GetTrack() string`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *CallStreamingFailedPayloadStreamParams) GetTrackOk() (*string, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *CallStreamingFailedPayloadStreamParams) SetTrack(v string)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *CallStreamingFailedPayloadStreamParams) HasTrack() bool`

HasTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


