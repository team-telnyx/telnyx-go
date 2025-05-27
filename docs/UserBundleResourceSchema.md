# UserBundleResourceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Resource&#39;s ID. | 
**Resource** | **string** | The resource itself (usually a phone number). | 
**ResourceType** | **string** | The type of the resource (usually &#39;number&#39;). | 
**CreatedAt** | **string** | Date the resource was created. | 
**UpdatedAt** | Pointer to **NullableString** | Date the resource was last updated. | [optional] 

## Methods

### NewUserBundleResourceSchema

`func NewUserBundleResourceSchema(id string, resource string, resourceType string, createdAt string, ) *UserBundleResourceSchema`

NewUserBundleResourceSchema instantiates a new UserBundleResourceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBundleResourceSchemaWithDefaults

`func NewUserBundleResourceSchemaWithDefaults() *UserBundleResourceSchema`

NewUserBundleResourceSchemaWithDefaults instantiates a new UserBundleResourceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserBundleResourceSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserBundleResourceSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserBundleResourceSchema) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *UserBundleResourceSchema) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UserBundleResourceSchema) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UserBundleResourceSchema) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResourceType

`func (o *UserBundleResourceSchema) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *UserBundleResourceSchema) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *UserBundleResourceSchema) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetCreatedAt

`func (o *UserBundleResourceSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserBundleResourceSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserBundleResourceSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserBundleResourceSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserBundleResourceSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserBundleResourceSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserBundleResourceSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *UserBundleResourceSchema) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *UserBundleResourceSchema) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


