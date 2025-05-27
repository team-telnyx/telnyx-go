# TransferToolParamsCustomHeadersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | The value of the header. Note that we support mustache templating for the value. For example you can use &#x60;{{#integration_secret}}test-secret{{/integration_secret}}&#x60; to pass the value of the integration secret. | [optional] 

## Methods

### NewTransferToolParamsCustomHeadersInner

`func NewTransferToolParamsCustomHeadersInner() *TransferToolParamsCustomHeadersInner`

NewTransferToolParamsCustomHeadersInner instantiates a new TransferToolParamsCustomHeadersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToolParamsCustomHeadersInnerWithDefaults

`func NewTransferToolParamsCustomHeadersInnerWithDefaults() *TransferToolParamsCustomHeadersInner`

NewTransferToolParamsCustomHeadersInnerWithDefaults instantiates a new TransferToolParamsCustomHeadersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TransferToolParamsCustomHeadersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransferToolParamsCustomHeadersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransferToolParamsCustomHeadersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransferToolParamsCustomHeadersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *TransferToolParamsCustomHeadersInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransferToolParamsCustomHeadersInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransferToolParamsCustomHeadersInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransferToolParamsCustomHeadersInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


