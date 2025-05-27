# ESimPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimCardGroupId** | Pointer to **string** | The group SIMCardGroup identification. This attribute can be &lt;code&gt;null&lt;/code&gt; when it&#39;s present in an associated resource. | [optional] 
**Tags** | Pointer to **[]string** | Searchable tags associated with the SIM cards | [optional] 
**Product** | Pointer to **string** | Type of product to be purchased, specify \&quot;whitelabel\&quot; to use a custom SPN | [optional] 
**WhitelabelName** | Pointer to **string** | Service Provider Name (SPN) for the Whitelabel eSIM product. It will be displayed as the mobile service name by operating systems of smartphones. This parameter must only contain letters, numbers and whitespaces. | [optional] 
**Amount** | **int32** | The amount of eSIMs to be purchased. | 
**Status** | Pointer to **string** | Status on which the SIM cards will be set after being successfully registered. | [optional] [default to "enabled"]

## Methods

### NewESimPurchase

`func NewESimPurchase(amount int32, ) *ESimPurchase`

NewESimPurchase instantiates a new ESimPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESimPurchaseWithDefaults

`func NewESimPurchaseWithDefaults() *ESimPurchase`

NewESimPurchaseWithDefaults instantiates a new ESimPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimCardGroupId

`func (o *ESimPurchase) GetSimCardGroupId() string`

GetSimCardGroupId returns the SimCardGroupId field if non-nil, zero value otherwise.

### GetSimCardGroupIdOk

`func (o *ESimPurchase) GetSimCardGroupIdOk() (*string, bool)`

GetSimCardGroupIdOk returns a tuple with the SimCardGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardGroupId

`func (o *ESimPurchase) SetSimCardGroupId(v string)`

SetSimCardGroupId sets SimCardGroupId field to given value.

### HasSimCardGroupId

`func (o *ESimPurchase) HasSimCardGroupId() bool`

HasSimCardGroupId returns a boolean if a field has been set.

### GetTags

`func (o *ESimPurchase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ESimPurchase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ESimPurchase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ESimPurchase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProduct

`func (o *ESimPurchase) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ESimPurchase) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ESimPurchase) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ESimPurchase) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetWhitelabelName

`func (o *ESimPurchase) GetWhitelabelName() string`

GetWhitelabelName returns the WhitelabelName field if non-nil, zero value otherwise.

### GetWhitelabelNameOk

`func (o *ESimPurchase) GetWhitelabelNameOk() (*string, bool)`

GetWhitelabelNameOk returns a tuple with the WhitelabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelabelName

`func (o *ESimPurchase) SetWhitelabelName(v string)`

SetWhitelabelName sets WhitelabelName field to given value.

### HasWhitelabelName

`func (o *ESimPurchase) HasWhitelabelName() bool`

HasWhitelabelName returns a boolean if a field has been set.

### GetAmount

`func (o *ESimPurchase) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ESimPurchase) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ESimPurchase) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *ESimPurchase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ESimPurchase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ESimPurchase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ESimPurchase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


