# TransferTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Transfer** | [**TransferToolParams**](TransferToolParams.md) |  | 

## Methods

### NewTransferTool

`func NewTransferTool(type_ string, transfer TransferToolParams, ) *TransferTool`

NewTransferTool instantiates a new TransferTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToolWithDefaults

`func NewTransferToolWithDefaults() *TransferTool`

NewTransferToolWithDefaults instantiates a new TransferTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransferTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferTool) SetType(v string)`

SetType sets Type field to given value.


### GetTransfer

`func (o *TransferTool) GetTransfer() TransferToolParams`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *TransferTool) GetTransferOk() (*TransferToolParams, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *TransferTool) SetTransfer(v TransferToolParams)`

SetTransfer sets Transfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


