# DTMFTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**SendDtmf** | **map[string]interface{}** |  | 

## Methods

### NewDTMFTool

`func NewDTMFTool(type_ string, sendDtmf map[string]interface{}, ) *DTMFTool`

NewDTMFTool instantiates a new DTMFTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTMFToolWithDefaults

`func NewDTMFToolWithDefaults() *DTMFTool`

NewDTMFToolWithDefaults instantiates a new DTMFTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DTMFTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DTMFTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DTMFTool) SetType(v string)`

SetType sets Type field to given value.


### GetSendDtmf

`func (o *DTMFTool) GetSendDtmf() map[string]interface{}`

GetSendDtmf returns the SendDtmf field if non-nil, zero value otherwise.

### GetSendDtmfOk

`func (o *DTMFTool) GetSendDtmfOk() (*map[string]interface{}, bool)`

GetSendDtmfOk returns a tuple with the SendDtmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDtmf

`func (o *DTMFTool) SetSendDtmf(v map[string]interface{})`

SetSendDtmf sets SendDtmf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


