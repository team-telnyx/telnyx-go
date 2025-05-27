# PortingOrderMisc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PortingOrderType**](PortingOrderType.md) |  | [optional] 
**RemainingNumbersAction** | Pointer to **string** | Remaining numbers can be either kept with their current service provider or disconnected. &#39;new_billing_telephone_number&#39; is required when &#39;remaining_numbers_action&#39; is &#39;keep&#39;. | [optional] 
**NewBillingPhoneNumber** | Pointer to **string** | New billing phone number for the remaining numbers. Used in case the current billing phone number is being ported to Telnyx. This will be set on your account with your current service provider and should be one of the numbers remaining on that account. | [optional] 

## Methods

### NewPortingOrderMisc

`func NewPortingOrderMisc() *PortingOrderMisc`

NewPortingOrderMisc instantiates a new PortingOrderMisc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderMiscWithDefaults

`func NewPortingOrderMiscWithDefaults() *PortingOrderMisc`

NewPortingOrderMiscWithDefaults instantiates a new PortingOrderMisc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PortingOrderMisc) GetType() PortingOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortingOrderMisc) GetTypeOk() (*PortingOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortingOrderMisc) SetType(v PortingOrderType)`

SetType sets Type field to given value.

### HasType

`func (o *PortingOrderMisc) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRemainingNumbersAction

`func (o *PortingOrderMisc) GetRemainingNumbersAction() string`

GetRemainingNumbersAction returns the RemainingNumbersAction field if non-nil, zero value otherwise.

### GetRemainingNumbersActionOk

`func (o *PortingOrderMisc) GetRemainingNumbersActionOk() (*string, bool)`

GetRemainingNumbersActionOk returns a tuple with the RemainingNumbersAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingNumbersAction

`func (o *PortingOrderMisc) SetRemainingNumbersAction(v string)`

SetRemainingNumbersAction sets RemainingNumbersAction field to given value.

### HasRemainingNumbersAction

`func (o *PortingOrderMisc) HasRemainingNumbersAction() bool`

HasRemainingNumbersAction returns a boolean if a field has been set.

### GetNewBillingPhoneNumber

`func (o *PortingOrderMisc) GetNewBillingPhoneNumber() string`

GetNewBillingPhoneNumber returns the NewBillingPhoneNumber field if non-nil, zero value otherwise.

### GetNewBillingPhoneNumberOk

`func (o *PortingOrderMisc) GetNewBillingPhoneNumberOk() (*string, bool)`

GetNewBillingPhoneNumberOk returns a tuple with the NewBillingPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewBillingPhoneNumber

`func (o *PortingOrderMisc) SetNewBillingPhoneNumber(v string)`

SetNewBillingPhoneNumber sets NewBillingPhoneNumber field to given value.

### HasNewBillingPhoneNumber

`func (o *PortingOrderMisc) HasNewBillingPhoneNumber() bool`

HasNewBillingPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


