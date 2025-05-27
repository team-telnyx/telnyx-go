# CreateUserBundlesBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | Pointer to **string** | Idempotency key for the request. Can be any UUID, but should always be unique for each request. | [optional] 
**Items** | Pointer to [**[]CreateUserBundlesBulkRequestItemsInner**](CreateUserBundlesBulkRequestItemsInner.md) |  | [optional] 

## Methods

### NewCreateUserBundlesBulkRequest

`func NewCreateUserBundlesBulkRequest() *CreateUserBundlesBulkRequest`

NewCreateUserBundlesBulkRequest instantiates a new CreateUserBundlesBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserBundlesBulkRequestWithDefaults

`func NewCreateUserBundlesBulkRequestWithDefaults() *CreateUserBundlesBulkRequest`

NewCreateUserBundlesBulkRequestWithDefaults instantiates a new CreateUserBundlesBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdempotencyKey

`func (o *CreateUserBundlesBulkRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CreateUserBundlesBulkRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CreateUserBundlesBulkRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CreateUserBundlesBulkRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetItems

`func (o *CreateUserBundlesBulkRequest) GetItems() []CreateUserBundlesBulkRequestItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CreateUserBundlesBulkRequest) GetItemsOk() (*[]CreateUserBundlesBulkRequestItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CreateUserBundlesBulkRequest) SetItems(v []CreateUserBundlesBulkRequestItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CreateUserBundlesBulkRequest) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


