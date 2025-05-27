# UserBundleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Active** | **bool** |  | 
**UserId** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**BillingBundle** | [**BillingBundleSummary**](BillingBundleSummary.md) |  | 
**Resources** | [**[]UserBundleResourceSchema**](UserBundleResourceSchema.md) |  | 

## Methods

### NewUserBundleSummary

`func NewUserBundleSummary(id string, active bool, userId string, createdAt string, billingBundle BillingBundleSummary, resources []UserBundleResourceSchema, ) *UserBundleSummary`

NewUserBundleSummary instantiates a new UserBundleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBundleSummaryWithDefaults

`func NewUserBundleSummaryWithDefaults() *UserBundleSummary`

NewUserBundleSummaryWithDefaults instantiates a new UserBundleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBundleSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBundleSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBundleSummary) SetId(v string)`

SetId sets Id field to given value.


### GetActive

`func (o *UserBundleSummary) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserBundleSummary) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserBundleSummary) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUserId

`func (o *UserBundleSummary) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserBundleSummary) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserBundleSummary) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *UserBundleSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserBundleSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserBundleSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserBundleSummary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserBundleSummary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserBundleSummary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserBundleSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetBillingBundle

`func (o *UserBundleSummary) GetBillingBundle() BillingBundleSummary`

GetBillingBundle returns the BillingBundle field if non-nil, zero value otherwise.

### GetBillingBundleOk

`func (o *UserBundleSummary) GetBillingBundleOk() (*BillingBundleSummary, bool)`

GetBillingBundleOk returns a tuple with the BillingBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBundle

`func (o *UserBundleSummary) SetBillingBundle(v BillingBundleSummary)`

SetBillingBundle sets BillingBundle field to given value.


### GetResources

`func (o *UserBundleSummary) GetResources() []UserBundleResourceSchema`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserBundleSummary) GetResourcesOk() (*[]UserBundleResourceSchema, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserBundleSummary) SetResources(v []UserBundleResourceSchema)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


