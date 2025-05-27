# CallStreamingStoppedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallStreamingStopped**](CallStreamingStopped.md) |  | [optional] 

## Methods

### NewCallStreamingStoppedEvent

`func NewCallStreamingStoppedEvent() *CallStreamingStoppedEvent`

NewCallStreamingStoppedEvent instantiates a new CallStreamingStoppedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStreamingStoppedEventWithDefaults

`func NewCallStreamingStoppedEventWithDefaults() *CallStreamingStoppedEvent`

NewCallStreamingStoppedEventWithDefaults instantiates a new CallStreamingStoppedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CallStreamingStoppedEvent) GetData() CallStreamingStopped`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallStreamingStoppedEvent) GetDataOk() (*CallStreamingStopped, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallStreamingStoppedEvent) SetData(v CallStreamingStopped)`

SetData sets Data field to given value.

### HasData

`func (o *CallStreamingStoppedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


