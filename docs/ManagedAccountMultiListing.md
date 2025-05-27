# ManagedAccountMultiListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** | Identifies the type of the resource. | 
**Id** | **string** | Uniquely identifies the managed account. | 
**Email** | **string** | The managed account&#39;s email. | 
**ApiUser** | **string** | The manager account&#39;s email, which serves as the V1 API user identifier | 
**OrganizationName** | Pointer to **string** | The organization the managed account is associated with. | [optional] 
**ManagerAccountId** | **string** | The ID of the manager account associated with the managed account. | 
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | 
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | 
**ManagedAccountAllowCustomPricing** | Pointer to **bool** | Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. There may be time lag between when the value is changed and pricing changes take effect. | [optional] 
**RollupBilling** | Pointer to **bool** | Boolean value that indicates if the billing information and charges to the managed account \&quot;roll up\&quot; to the manager account. If true, the managed account will not have its own balance and will use the shared balance with the manager account. This value cannot be changed after account creation without going through Telnyx support as changes require manual updates to the account ledger. Defaults to false. | [optional] 

## Methods

### NewManagedAccountMultiListing

`func NewManagedAccountMultiListing(recordType string, id string, email string, apiUser string, managerAccountId string, createdAt string, updatedAt string, ) *ManagedAccountMultiListing`

NewManagedAccountMultiListing instantiates a new ManagedAccountMultiListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountMultiListingWithDefaults

`func NewManagedAccountMultiListingWithDefaults() *ManagedAccountMultiListing`

NewManagedAccountMultiListingWithDefaults instantiates a new ManagedAccountMultiListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ManagedAccountMultiListing) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ManagedAccountMultiListing) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ManagedAccountMultiListing) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *ManagedAccountMultiListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccountMultiListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccountMultiListing) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *ManagedAccountMultiListing) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ManagedAccountMultiListing) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ManagedAccountMultiListing) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiUser

`func (o *ManagedAccountMultiListing) GetApiUser() string`

GetApiUser returns the ApiUser field if non-nil, zero value otherwise.

### GetApiUserOk

`func (o *ManagedAccountMultiListing) GetApiUserOk() (*string, bool)`

GetApiUserOk returns a tuple with the ApiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUser

`func (o *ManagedAccountMultiListing) SetApiUser(v string)`

SetApiUser sets ApiUser field to given value.


### GetOrganizationName

`func (o *ManagedAccountMultiListing) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ManagedAccountMultiListing) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ManagedAccountMultiListing) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ManagedAccountMultiListing) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetManagerAccountId

`func (o *ManagedAccountMultiListing) GetManagerAccountId() string`

GetManagerAccountId returns the ManagerAccountId field if non-nil, zero value otherwise.

### GetManagerAccountIdOk

`func (o *ManagedAccountMultiListing) GetManagerAccountIdOk() (*string, bool)`

GetManagerAccountIdOk returns a tuple with the ManagerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAccountId

`func (o *ManagedAccountMultiListing) SetManagerAccountId(v string)`

SetManagerAccountId sets ManagerAccountId field to given value.


### GetCreatedAt

`func (o *ManagedAccountMultiListing) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedAccountMultiListing) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedAccountMultiListing) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManagedAccountMultiListing) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedAccountMultiListing) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedAccountMultiListing) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetManagedAccountAllowCustomPricing

`func (o *ManagedAccountMultiListing) GetManagedAccountAllowCustomPricing() bool`

GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field if non-nil, zero value otherwise.

### GetManagedAccountAllowCustomPricingOk

`func (o *ManagedAccountMultiListing) GetManagedAccountAllowCustomPricingOk() (*bool, bool)`

GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccountAllowCustomPricing

`func (o *ManagedAccountMultiListing) SetManagedAccountAllowCustomPricing(v bool)`

SetManagedAccountAllowCustomPricing sets ManagedAccountAllowCustomPricing field to given value.

### HasManagedAccountAllowCustomPricing

`func (o *ManagedAccountMultiListing) HasManagedAccountAllowCustomPricing() bool`

HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.

### GetRollupBilling

`func (o *ManagedAccountMultiListing) GetRollupBilling() bool`

GetRollupBilling returns the RollupBilling field if non-nil, zero value otherwise.

### GetRollupBillingOk

`func (o *ManagedAccountMultiListing) GetRollupBillingOk() (*bool, bool)`

GetRollupBillingOk returns a tuple with the RollupBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupBilling

`func (o *ManagedAccountMultiListing) SetRollupBilling(v bool)`

SetRollupBilling sets RollupBilling field to given value.

### HasRollupBilling

`func (o *ManagedAccountMultiListing) HasRollupBilling() bool`

HasRollupBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


