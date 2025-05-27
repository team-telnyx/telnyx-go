# UnusedUserBundlesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingBundle** | [**BillingBundleSummary**](BillingBundleSummary.md) |  | 
**UserBundleIds** | **[]string** | List of user bundle IDs for given bundle. | 

## Methods

### NewUnusedUserBundlesSchema

`func NewUnusedUserBundlesSchema(billingBundle BillingBundleSummary, userBundleIds []string, ) *UnusedUserBundlesSchema`

NewUnusedUserBundlesSchema instantiates a new UnusedUserBundlesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnusedUserBundlesSchemaWithDefaults

`func NewUnusedUserBundlesSchemaWithDefaults() *UnusedUserBundlesSchema`

NewUnusedUserBundlesSchemaWithDefaults instantiates a new UnusedUserBundlesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingBundle

`func (o *UnusedUserBundlesSchema) GetBillingBundle() BillingBundleSummary`

GetBillingBundle returns the BillingBundle field if non-nil, zero value otherwise.

### GetBillingBundleOk

`func (o *UnusedUserBundlesSchema) GetBillingBundleOk() (*BillingBundleSummary, bool)`

GetBillingBundleOk returns a tuple with the BillingBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingBundle

`func (o *UnusedUserBundlesSchema) SetBillingBundle(v BillingBundleSummary)`

SetBillingBundle sets BillingBundle field to given value.


### GetUserBundleIds

`func (o *UnusedUserBundlesSchema) GetUserBundleIds() []string`

GetUserBundleIds returns the UserBundleIds field if non-nil, zero value otherwise.

### GetUserBundleIdsOk

`func (o *UnusedUserBundlesSchema) GetUserBundleIdsOk() (*[]string, bool)`

GetUserBundleIdsOk returns a tuple with the UserBundleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBundleIds

`func (o *UnusedUserBundlesSchema) SetUserBundleIds(v []string)`

SetUserBundleIds sets UserBundleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


