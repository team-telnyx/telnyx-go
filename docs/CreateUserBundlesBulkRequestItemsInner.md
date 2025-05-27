# CreateUserBundlesBulkRequestItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingBundleId** | **string** | Quantity of user bundles to order. | 
**Quantity** | **int32** | Quantity of user bundles to order. | 

## Methods

### NewCreateUserBundlesBulkRequestItemsInner

`func NewCreateUserBundlesBulkRequestItemsInner(billingBundleId string, quantity int32, ) *CreateUserBundlesBulkRequestItemsInner`

NewCreateUserBundlesBulkRequestItemsInner instantiates a new CreateUserBundlesBulkRequestItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserBundlesBulkRequestItemsInnerWithDefaults

`func NewCreateUserBundlesBulkRequestItemsInnerWithDefaults() *CreateUserBundlesBulkRequestItemsInner`

NewCreateUserBundlesBulkRequestItemsInnerWithDefaults instantiates a new CreateUserBundlesBulkRequestItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingBundleId

`func (o *CreateUserBundlesBulkRequestItemsInner) GetBillingBundleId() string`

GetBillingBundleId returns the BillingBundleId field if non-nil, zero value otherwise.

### GetBillingBundleIdOk

`func (o *CreateUserBundlesBulkRequestItemsInner) GetBillingBundleIdOk() (*string, bool)`

GetBillingBundleIdOk returns a tuple with the BillingBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBundleId

`func (o *CreateUserBundlesBulkRequestItemsInner) SetBillingBundleId(v string)`

SetBillingBundleId sets BillingBundleId field to given value.


### GetQuantity

`func (o *CreateUserBundlesBulkRequestItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateUserBundlesBulkRequestItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateUserBundlesBulkRequestItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


