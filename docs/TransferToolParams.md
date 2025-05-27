# TransferToolParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | [**[]TransferToolParamsTargetsInner**](TransferToolParamsTargetsInner.md) | The different possible targets of the transfer. The assistant will be able to choose one of the targets to transfer the call to. | 
**From** | **string** | Number or SIP URI placing the call. | 
**CustomHeaders** | Pointer to [**[]TransferToolParamsCustomHeadersInner**](TransferToolParamsCustomHeadersInner.md) | Custom headers to be added to the SIP INVITE for the transfer command. | [optional] 

## Methods

### NewTransferToolParams

`func NewTransferToolParams(targets []TransferToolParamsTargetsInner, from string, ) *TransferToolParams`

NewTransferToolParams instantiates a new TransferToolParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToolParamsWithDefaults

`func NewTransferToolParamsWithDefaults() *TransferToolParams`

NewTransferToolParamsWithDefaults instantiates a new TransferToolParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *TransferToolParams) GetTargets() []TransferToolParamsTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TransferToolParams) GetTargetsOk() (*[]TransferToolParamsTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TransferToolParams) SetTargets(v []TransferToolParamsTargetsInner)`

SetTargets sets Targets field to given value.


### GetFrom

`func (o *TransferToolParams) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransferToolParams) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransferToolParams) SetFrom(v string)`

SetFrom sets From field to given value.


### GetCustomHeaders

`func (o *TransferToolParams) GetCustomHeaders() []TransferToolParamsCustomHeadersInner`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *TransferToolParams) GetCustomHeadersOk() (*[]TransferToolParamsCustomHeadersInner, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *TransferToolParams) SetCustomHeaders(v []TransferToolParamsCustomHeadersInner)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *TransferToolParams) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


