# PurchaseESim202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SimpleSIMCard**](SimpleSIMCard.md) | Successfully registered SIM cards. | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewPurchaseESim202Response

`func NewPurchaseESim202Response() *PurchaseESim202Response`

NewPurchaseESim202Response instantiates a new PurchaseESim202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseESim202ResponseWithDefaults

`func NewPurchaseESim202ResponseWithDefaults() *PurchaseESim202Response`

NewPurchaseESim202ResponseWithDefaults instantiates a new PurchaseESim202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PurchaseESim202Response) GetData() []SimpleSIMCard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PurchaseESim202Response) GetDataOk() (*[]SimpleSIMCard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PurchaseESim202Response) SetData(v []SimpleSIMCard)`

SetData sets Data field to given value.

### HasData

`func (o *PurchaseESim202Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrors

`func (o *PurchaseESim202Response) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PurchaseESim202Response) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PurchaseESim202Response) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PurchaseESim202Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


