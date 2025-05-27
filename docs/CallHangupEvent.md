# CallHangupEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallHangup**](CallHangup.md) |  | [optional] 

## Methods

### NewCallHangupEvent

`func NewCallHangupEvent() *CallHangupEvent`

NewCallHangupEvent instantiates a new CallHangupEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallHangupEventWithDefaults

`func NewCallHangupEventWithDefaults() *CallHangupEvent`

NewCallHangupEventWithDefaults instantiates a new CallHangupEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CallHangupEvent) GetData() CallHangup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallHangupEvent) GetDataOk() (*CallHangup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallHangupEvent) SetData(v CallHangup)`

SetData sets Data field to given value.

### HasData

`func (o *CallHangupEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


