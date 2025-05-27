# BillingBundleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Bundle&#39;s ID, this is used to identify the bundle in the API. | 
**Name** | **string** | Bundle&#39;s name, this is used to identify the bundle in the UI. | 
**Slug** | Pointer to **string** | Slugified version of the bundle&#39;s name. | [optional] 
**CostCode** | **string** | Bundle&#39;s cost code, this is used to identify the bundle in the billing system. | 
**Active** | **bool** | If that bundle is active or not. | 
**IsPublic** | **bool** | Available to all customers or only to specific customers. | 
**CreatedAt** | **string** | Date the bundle was created. | 
**BundleLimits** | [**[]BundleLimitSchema**](BundleLimitSchema.md) |  | 

## Methods

### NewBillingBundleSchema

`func NewBillingBundleSchema(id string, name string, costCode string, active bool, isPublic bool, createdAt string, bundleLimits []BundleLimitSchema, ) *BillingBundleSchema`

NewBillingBundleSchema instantiates a new BillingBundleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingBundleSchemaWithDefaults

`func NewBillingBundleSchemaWithDefaults() *BillingBundleSchema`

NewBillingBundleSchemaWithDefaults instantiates a new BillingBundleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingBundleSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingBundleSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingBundleSchema) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BillingBundleSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingBundleSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingBundleSchema) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BillingBundleSchema) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BillingBundleSchema) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BillingBundleSchema) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BillingBundleSchema) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetCostCode

`func (o *BillingBundleSchema) GetCostCode() string`

GetCostCode returns the CostCode field if non-nil, zero value otherwise.

### GetCostCodeOk

`func (o *BillingBundleSchema) GetCostCodeOk() (*string, bool)`

GetCostCodeOk returns a tuple with the CostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCode

`func (o *BillingBundleSchema) SetCostCode(v string)`

SetCostCode sets CostCode field to given value.


### GetActive

`func (o *BillingBundleSchema) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BillingBundleSchema) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BillingBundleSchema) SetActive(v bool)`

SetActive sets Active field to given value.


### GetIsPublic

`func (o *BillingBundleSchema) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *BillingBundleSchema) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *BillingBundleSchema) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetCreatedAt

`func (o *BillingBundleSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingBundleSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingBundleSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetBundleLimits

`func (o *BillingBundleSchema) GetBundleLimits() []BundleLimitSchema`

GetBundleLimits returns the BundleLimits field if non-nil, zero value otherwise.

### GetBundleLimitsOk

`func (o *BillingBundleSchema) GetBundleLimitsOk() (*[]BundleLimitSchema, bool)`

GetBundleLimitsOk returns a tuple with the BundleLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleLimits

`func (o *BillingBundleSchema) SetBundleLimits(v []BundleLimitSchema)`

SetBundleLimits sets BundleLimits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


