# CallStreamingStartedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallStreamingStarted**](CallStreamingStarted.md) |  | [optional] 

## Methods

### NewCallStreamingStartedEvent

`func NewCallStreamingStartedEvent() *CallStreamingStartedEvent`

NewCallStreamingStartedEvent instantiates a new CallStreamingStartedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStreamingStartedEventWithDefaults

`func NewCallStreamingStartedEventWithDefaults() *CallStreamingStartedEvent`

NewCallStreamingStartedEventWithDefaults instantiates a new CallStreamingStartedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CallStreamingStartedEvent) GetData() CallStreamingStarted`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallStreamingStartedEvent) GetDataOk() (*CallStreamingStarted, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallStreamingStartedEvent) SetData(v CallStreamingStarted)`

SetData sets Data field to given value.

### HasData

`func (o *CallStreamingStartedEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


