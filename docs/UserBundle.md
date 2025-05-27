# UserBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User bundle&#39;s ID, this is used to identify the user bundle in the API. | 
**Active** | **bool** | Status of the user bundle. | 
**UserId** | **string** | The customer&#39;s ID that owns this user bundle. | 
**CreatedAt** | **string** | Date the user bundle was created. | 
**UpdatedAt** | Pointer to **NullableString** | Date the user bundle was last updated. | [optional] 
**BillingBundle** | [**BillingBundleSummary**](BillingBundleSummary.md) |  | 
**Resources** | [**[]UserBundleResourceSchema**](UserBundleResourceSchema.md) |  | 

## Methods

### NewUserBundle

`func NewUserBundle(id string, active bool, userId string, createdAt string, billingBundle BillingBundleSummary, resources []UserBundleResourceSchema, ) *UserBundle`

NewUserBundle instantiates a new UserBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBundleWithDefaults

`func NewUserBundleWithDefaults() *UserBundle`

NewUserBundleWithDefaults instantiates a new UserBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBundle) SetId(v string)`

SetId sets Id field to given value.


### GetActive

`func (o *UserBundle) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserBundle) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserBundle) SetActive(v bool)`

SetActive sets Active field to given value.


### GetUserId

`func (o *UserBundle) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserBundle) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserBundle) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *UserBundle) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserBundle) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserBundle) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserBundle) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserBundle) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserBundle) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserBundle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *UserBundle) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UserBundle) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetBillingBundle

`func (o *UserBundle) GetBillingBundle() BillingBundleSummary`

GetBillingBundle returns the BillingBundle field if non-nil, zero value otherwise.

### GetBillingBundleOk

`func (o *UserBundle) GetBillingBundleOk() (*BillingBundleSummary, bool)`

GetBillingBundleOk returns a tuple with the BillingBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBundle

`func (o *UserBundle) SetBillingBundle(v BillingBundleSummary)`

SetBillingBundle sets BillingBundle field to given value.


### GetResources

`func (o *UserBundle) GetResources() []UserBundleResourceSchema`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserBundle) GetResourcesOk() (*[]UserBundleResourceSchema, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserBundle) SetResources(v []UserBundleResourceSchema)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


