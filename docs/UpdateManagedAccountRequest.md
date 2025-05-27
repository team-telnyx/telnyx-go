# UpdateManagedAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedAccountAllowCustomPricing** | Pointer to **bool** | Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. This value may be changed, but there may be time lag between when the value is changed and pricing changes take effect. | [optional] 

## Methods

### NewUpdateManagedAccountRequest

`func NewUpdateManagedAccountRequest() *UpdateManagedAccountRequest`

NewUpdateManagedAccountRequest instantiates a new UpdateManagedAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateManagedAccountRequestWithDefaults

`func NewUpdateManagedAccountRequestWithDefaults() *UpdateManagedAccountRequest`

NewUpdateManagedAccountRequestWithDefaults instantiates a new UpdateManagedAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedAccountAllowCustomPricing

`func (o *UpdateManagedAccountRequest) GetManagedAccountAllowCustomPricing() bool`

GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field if non-nil, zero value otherwise.

### GetManagedAccountAllowCustomPricingOk

`func (o *UpdateManagedAccountRequest) GetManagedAccountAllowCustomPricingOk() (*bool, bool)`

GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccountAllowCustomPricing

`func (o *UpdateManagedAccountRequest) SetManagedAccountAllowCustomPricing(v bool)`

SetManagedAccountAllowCustomPricing sets ManagedAccountAllowCustomPricing field to given value.

### HasManagedAccountAllowCustomPricing

`func (o *UpdateManagedAccountRequest) HasManagedAccountAllowCustomPricing() bool`

HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


