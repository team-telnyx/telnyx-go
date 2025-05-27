# CallPlaybackEndedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallPlaybackEnded**](CallPlaybackEnded.md) |  | [optional] 

## Methods

### NewCallPlaybackEndedEvent

`func NewCallPlaybackEndedEvent() *CallPlaybackEndedEvent`

NewCallPlaybackEndedEvent instantiates a new CallPlaybackEndedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPlaybackEndedEventWithDefaults

`func NewCallPlaybackEndedEventWithDefaults() *CallPlaybackEndedEvent`

NewCallPlaybackEndedEventWithDefaults instantiates a new CallPlaybackEndedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CallPlaybackEndedEvent) GetData() CallPlaybackEnded`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallPlaybackEndedEvent) GetDataOk() (*CallPlaybackEnded, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallPlaybackEndedEvent) SetData(v CallPlaybackEnded)`

SetData sets Data field to given value.

### HasData

`func (o *CallPlaybackEndedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


