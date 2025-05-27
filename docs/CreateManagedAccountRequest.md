# CreateManagedAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address for the managed account. If not provided, the email address will be generated based on the email address of the manager account. | [optional] 
**Password** | Pointer to **string** | Password for the managed account. If a password is not supplied, the account will not be able to be signed into directly. (A password reset may still be performed later to enable sign-in via password.) | [optional] 
**BusinessName** | **string** | The name of the business for which the new managed account is being created, that will be used as the managed accounts&#39;s organization&#39;s name. | 
**ManagedAccountAllowCustomPricing** | Pointer to **bool** | Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. This value may be changed after creation, but there may be time lag between when the value is changed and pricing changes take effect. | [optional] 
**RollupBilling** | Pointer to **bool** | Boolean value that indicates if the billing information and charges to the managed account \&quot;roll up\&quot; to the manager account. If true, the managed account will not have its own balance and will use the shared balance with the manager account. This value cannot be changed after account creation without going through Telnyx support as changes require manual updates to the account ledger. Defaults to false. | [optional] 

## Methods

### NewCreateManagedAccountRequest

`func NewCreateManagedAccountRequest(businessName string, ) *CreateManagedAccountRequest`

NewCreateManagedAccountRequest instantiates a new CreateManagedAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManagedAccountRequestWithDefaults

`func NewCreateManagedAccountRequestWithDefaults() *CreateManagedAccountRequest`

NewCreateManagedAccountRequestWithDefaults instantiates a new CreateManagedAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateManagedAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateManagedAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateManagedAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateManagedAccountRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *CreateManagedAccountRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateManagedAccountRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateManagedAccountRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateManagedAccountRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetBusinessName

`func (o *CreateManagedAccountRequest) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *CreateManagedAccountRequest) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *CreateManagedAccountRequest) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetManagedAccountAllowCustomPricing

`func (o *CreateManagedAccountRequest) GetManagedAccountAllowCustomPricing() bool`

GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field if non-nil, zero value otherwise.

### GetManagedAccountAllowCustomPricingOk

`func (o *CreateManagedAccountRequest) GetManagedAccountAllowCustomPricingOk() (*bool, bool)`

GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccountAllowCustomPricing

`func (o *CreateManagedAccountRequest) SetManagedAccountAllowCustomPricing(v bool)`

SetManagedAccountAllowCustomPricing sets ManagedAccountAllowCustomPricing field to given value.

### HasManagedAccountAllowCustomPricing

`func (o *CreateManagedAccountRequest) HasManagedAccountAllowCustomPricing() bool`

HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.

### GetRollupBilling

`func (o *CreateManagedAccountRequest) GetRollupBilling() bool`

GetRollupBilling returns the RollupBilling field if non-nil, zero value otherwise.

### GetRollupBillingOk

`func (o *CreateManagedAccountRequest) GetRollupBillingOk() (*bool, bool)`

GetRollupBillingOk returns a tuple with the RollupBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupBilling

`func (o *CreateManagedAccountRequest) SetRollupBilling(v bool)`

SetRollupBilling sets RollupBilling field to given value.

### HasRollupBilling

`func (o *CreateManagedAccountRequest) HasRollupBilling() bool`

HasRollupBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


