# HangupTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Hangup** | [**HangupToolParams**](HangupToolParams.md) |  | 

## Methods

### NewHangupTool

`func NewHangupTool(type_ string, hangup HangupToolParams, ) *HangupTool`

NewHangupTool instantiates a new HangupTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHangupToolWithDefaults

`func NewHangupToolWithDefaults() *HangupTool`

NewHangupToolWithDefaults instantiates a new HangupTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HangupTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HangupTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HangupTool) SetType(v string)`

SetType sets Type field to given value.


### GetHangup

`func (o *HangupTool) GetHangup() HangupToolParams`

GetHangup returns the Hangup field if non-nil, zero value otherwise.

### GetHangupOk

`func (o *HangupTool) GetHangupOk() (*HangupToolParams, bool)`

GetHangupOk returns a tuple with the Hangup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangup

`func (o *HangupTool) SetHangup(v HangupToolParams)`

SetHangup sets Hangup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


