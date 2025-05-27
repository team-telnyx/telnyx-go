# AccessIPRangeResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CidrBlock** | **string** |  | 
**Status** | [**CloudflareSyncStatus**](CloudflareSyncStatus.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccessIPRangeResponseSchema

`func NewAccessIPRangeResponseSchema(id string, cidrBlock string, status CloudflareSyncStatus, userId string, ) *AccessIPRangeResponseSchema`

NewAccessIPRangeResponseSchema instantiates a new AccessIPRangeResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessIPRangeResponseSchemaWithDefaults

`func NewAccessIPRangeResponseSchemaWithDefaults() *AccessIPRangeResponseSchema`

NewAccessIPRangeResponseSchemaWithDefaults instantiates a new AccessIPRangeResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessIPRangeResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessIPRangeResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessIPRangeResponseSchema) SetId(v string)`

SetId sets Id field to given value.


### GetCidrBlock

`func (o *AccessIPRangeResponseSchema) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AccessIPRangeResponseSchema) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AccessIPRangeResponseSchema) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetStatus

`func (o *AccessIPRangeResponseSchema) GetStatus() CloudflareSyncStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessIPRangeResponseSchema) GetStatusOk() (*CloudflareSyncStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessIPRangeResponseSchema) SetStatus(v CloudflareSyncStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *AccessIPRangeResponseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessIPRangeResponseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessIPRangeResponseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessIPRangeResponseSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserId

`func (o *AccessIPRangeResponseSchema) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessIPRangeResponseSchema) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessIPRangeResponseSchema) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *AccessIPRangeResponseSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessIPRangeResponseSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessIPRangeResponseSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessIPRangeResponseSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccessIPRangeResponseSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccessIPRangeResponseSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccessIPRangeResponseSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccessIPRangeResponseSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


