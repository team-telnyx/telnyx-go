# ManagedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** | Identifies the type of the resource. | 
**Id** | **string** | Uniquely identifies the managed account. | 
**Email** | **string** | The managed account&#39;s email. | 
**ApiKey** | **string** | The managed account&#39;s V2 API access key | 
**ApiUser** | **string** | The manager account&#39;s email, which serves as the V1 API user identifier | 
**ApiToken** | **string** | The managed account&#39;s V1 API token | 
**OrganizationName** | Pointer to **string** | The organization the managed account is associated with. | [optional] 
**ManagerAccountId** | **string** | The ID of the manager account associated with the managed account. | 
**Balance** | Pointer to [**ManagedAccountBalance**](ManagedAccountBalance.md) |  | [optional] 
**CreatedAt** | **string** | ISO 8601 formatted date indicating when the resource was created. | 
**UpdatedAt** | **string** | ISO 8601 formatted date indicating when the resource was updated. | 
**ManagedAccountAllowCustomPricing** | Pointer to **bool** | Boolean value that indicates if the managed account is able to have custom pricing set for it or not. If false, uses the pricing of the manager account. Defaults to false. There may be time lag between when the value is changed and pricing changes take effect. | [optional] 
**RollupBilling** | Pointer to **bool** | Boolean value that indicates if the billing information and charges to the managed account \&quot;roll up\&quot; to the manager account. If true, the managed account will not have its own balance and will use the shared balance with the manager account. This value cannot be changed after account creation without going through Telnyx support as changes require manual updates to the account ledger. Defaults to false. | [optional] 

## Methods

### NewManagedAccount

`func NewManagedAccount(recordType string, id string, email string, apiKey string, apiUser string, apiToken string, managerAccountId string, createdAt string, updatedAt string, ) *ManagedAccount`

NewManagedAccount instantiates a new ManagedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountWithDefaults

`func NewManagedAccountWithDefaults() *ManagedAccount`

NewManagedAccountWithDefaults instantiates a new ManagedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ManagedAccount) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ManagedAccount) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ManagedAccount) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *ManagedAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccount) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *ManagedAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ManagedAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ManagedAccount) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetApiKey

`func (o *ManagedAccount) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ManagedAccount) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ManagedAccount) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiUser

`func (o *ManagedAccount) GetApiUser() string`

GetApiUser returns the ApiUser field if non-nil, zero value otherwise.

### GetApiUserOk

`func (o *ManagedAccount) GetApiUserOk() (*string, bool)`

GetApiUserOk returns a tuple with the ApiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUser

`func (o *ManagedAccount) SetApiUser(v string)`

SetApiUser sets ApiUser field to given value.


### GetApiToken

`func (o *ManagedAccount) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ManagedAccount) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ManagedAccount) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetOrganizationName

`func (o *ManagedAccount) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ManagedAccount) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ManagedAccount) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ManagedAccount) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetManagerAccountId

`func (o *ManagedAccount) GetManagerAccountId() string`

GetManagerAccountId returns the ManagerAccountId field if non-nil, zero value otherwise.

### GetManagerAccountIdOk

`func (o *ManagedAccount) GetManagerAccountIdOk() (*string, bool)`

GetManagerAccountIdOk returns a tuple with the ManagerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAccountId

`func (o *ManagedAccount) SetManagerAccountId(v string)`

SetManagerAccountId sets ManagerAccountId field to given value.


### GetBalance

`func (o *ManagedAccount) GetBalance() ManagedAccountBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ManagedAccount) GetBalanceOk() (*ManagedAccountBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ManagedAccount) SetBalance(v ManagedAccountBalance)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ManagedAccount) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManagedAccount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedAccount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedAccount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManagedAccount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedAccount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedAccount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetManagedAccountAllowCustomPricing

`func (o *ManagedAccount) GetManagedAccountAllowCustomPricing() bool`

GetManagedAccountAllowCustomPricing returns the ManagedAccountAllowCustomPricing field if non-nil, zero value otherwise.

### GetManagedAccountAllowCustomPricingOk

`func (o *ManagedAccount) GetManagedAccountAllowCustomPricingOk() (*bool, bool)`

GetManagedAccountAllowCustomPricingOk returns a tuple with the ManagedAccountAllowCustomPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAccountAllowCustomPricing

`func (o *ManagedAccount) SetManagedAccountAllowCustomPricing(v bool)`

SetManagedAccountAllowCustomPricing sets ManagedAccountAllowCustomPricing field to given value.

### HasManagedAccountAllowCustomPricing

`func (o *ManagedAccount) HasManagedAccountAllowCustomPricing() bool`

HasManagedAccountAllowCustomPricing returns a boolean if a field has been set.

### GetRollupBilling

`func (o *ManagedAccount) GetRollupBilling() bool`

GetRollupBilling returns the RollupBilling field if non-nil, zero value otherwise.

### GetRollupBillingOk

`func (o *ManagedAccount) GetRollupBillingOk() (*bool, bool)`

GetRollupBillingOk returns a tuple with the RollupBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupBilling

`func (o *ManagedAccount) SetRollupBilling(v bool)`

SetRollupBilling sets RollupBilling field to given value.

### HasRollupBilling

`func (o *ManagedAccount) HasRollupBilling() bool`

HasRollupBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


