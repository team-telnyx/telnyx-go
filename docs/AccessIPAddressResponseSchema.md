# AccessIPAddressResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**IpAddress** | **string** |  | 
**Source** | **string** |  | 
**Status** | [**CloudflareSyncStatus**](CloudflareSyncStatus.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**UserId** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccessIPAddressResponseSchema

`func NewAccessIPAddressResponseSchema(id string, ipAddress string, source string, status CloudflareSyncStatus, userId string, ) *AccessIPAddressResponseSchema`

NewAccessIPAddressResponseSchema instantiates a new AccessIPAddressResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessIPAddressResponseSchemaWithDefaults

`func NewAccessIPAddressResponseSchemaWithDefaults() *AccessIPAddressResponseSchema`

NewAccessIPAddressResponseSchemaWithDefaults instantiates a new AccessIPAddressResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessIPAddressResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessIPAddressResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessIPAddressResponseSchema) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *AccessIPAddressResponseSchema) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AccessIPAddressResponseSchema) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AccessIPAddressResponseSchema) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetSource

`func (o *AccessIPAddressResponseSchema) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AccessIPAddressResponseSchema) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AccessIPAddressResponseSchema) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *AccessIPAddressResponseSchema) GetStatus() CloudflareSyncStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessIPAddressResponseSchema) GetStatusOk() (*CloudflareSyncStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessIPAddressResponseSchema) SetStatus(v CloudflareSyncStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *AccessIPAddressResponseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessIPAddressResponseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessIPAddressResponseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessIPAddressResponseSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserId

`func (o *AccessIPAddressResponseSchema) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessIPAddressResponseSchema) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessIPAddressResponseSchema) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *AccessIPAddressResponseSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessIPAddressResponseSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessIPAddressResponseSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessIPAddressResponseSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccessIPAddressResponseSchema) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccessIPAddressResponseSchema) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccessIPAddressResponseSchema) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccessIPAddressResponseSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


