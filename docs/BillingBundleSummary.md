# BillingBundleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Bundle&#39;s ID, this is used to identify the bundle in the API. | 
**Name** | **string** | Bundle&#39;s name, this is used to identify the bundle in the UI. | 
**Slug** | Pointer to **string** | Slugified version of the bundle&#39;s name. | [optional] 
**CostCode** | **string** | Bundle&#39;s cost code, this is used to identify the bundle in the billing system. | 
**IsPublic** | **bool** | Available to all customers or only to specific customers. | 
**CreatedAt** | **string** | Date the bundle was created. | 
**MrcPrice** | Pointer to **float32** | Monthly recurring charge price. | [optional] 
**Currency** | Pointer to **string** | Bundle&#39;s currency code. | [optional] 
**Specs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBillingBundleSummary

`func NewBillingBundleSummary(id string, name string, costCode string, isPublic bool, createdAt string, ) *BillingBundleSummary`

NewBillingBundleSummary instantiates a new BillingBundleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBundleSummaryWithDefaults

`func NewBillingBundleSummaryWithDefaults() *BillingBundleSummary`

NewBillingBundleSummaryWithDefaults instantiates a new BillingBundleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingBundleSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingBundleSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingBundleSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BillingBundleSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingBundleSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingBundleSummary) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BillingBundleSummary) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BillingBundleSummary) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BillingBundleSummary) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BillingBundleSummary) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCostCode

`func (o *BillingBundleSummary) GetCostCode() string`

GetCostCode returns the CostCode field if non-nil, zero value otherwise.

### GetCostCodeOk

`func (o *BillingBundleSummary) GetCostCodeOk() (*string, bool)`

GetCostCodeOk returns a tuple with the CostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCode

`func (o *BillingBundleSummary) SetCostCode(v string)`

SetCostCode sets CostCode field to given value.


### GetIsPublic

`func (o *BillingBundleSummary) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BillingBundleSummary) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BillingBundleSummary) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetCreatedAt

`func (o *BillingBundleSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingBundleSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingBundleSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMrcPrice

`func (o *BillingBundleSummary) GetMrcPrice() float32`

GetMrcPrice returns the MrcPrice field if non-nil, zero value otherwise.

### GetMrcPriceOk

`func (o *BillingBundleSummary) GetMrcPriceOk() (*float32, bool)`

GetMrcPriceOk returns a tuple with the MrcPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcPrice

`func (o *BillingBundleSummary) SetMrcPrice(v float32)`

SetMrcPrice sets MrcPrice field to given value.

### HasMrcPrice

`func (o *BillingBundleSummary) HasMrcPrice() bool`

HasMrcPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingBundleSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingBundleSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingBundleSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingBundleSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSpecs

`func (o *BillingBundleSummary) GetSpecs() []string`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *BillingBundleSummary) GetSpecsOk() (*[]string, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *BillingBundleSummary) SetSpecs(v []string)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *BillingBundleSummary) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


