# MigrationSourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the data migration source. | [optional] [readonly] 
**Provider** | **string** | Cloud provider from which to migrate data. Use &#39;telnyx&#39; if you want to migrate data from one Telnyx bucket to another. | 
**SourceRegion** | Pointer to **string** | For intra-Telnyx buckets migration, specify the source bucket region in this field. | [optional] 
**ProviderAuth** | [**MigrationSourceParamsProviderAuth**](MigrationSourceParamsProviderAuth.md) |  | 
**BucketName** | **string** | Bucket name to migrate the data from. | 

## Methods

### NewMigrationSourceParams

`func NewMigrationSourceParams(provider string, providerAuth MigrationSourceParamsProviderAuth, bucketName string, ) *MigrationSourceParams`

NewMigrationSourceParams instantiates a new MigrationSourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationSourceParamsWithDefaults

`func NewMigrationSourceParamsWithDefaults() *MigrationSourceParams`

NewMigrationSourceParamsWithDefaults instantiates a new MigrationSourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationSourceParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationSourceParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationSourceParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MigrationSourceParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *MigrationSourceParams) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MigrationSourceParams) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MigrationSourceParams) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetSourceRegion

`func (o *MigrationSourceParams) GetSourceRegion() string`

GetSourceRegion returns the SourceRegion field if non-nil, zero value otherwise.

### GetSourceRegionOk

`func (o *MigrationSourceParams) GetSourceRegionOk() (*string, bool)`

GetSourceRegionOk returns a tuple with the SourceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegion

`func (o *MigrationSourceParams) SetSourceRegion(v string)`

SetSourceRegion sets SourceRegion field to given value.

### HasSourceRegion

`func (o *MigrationSourceParams) HasSourceRegion() bool`

HasSourceRegion returns a boolean if a field has been set.

### GetProviderAuth

`func (o *MigrationSourceParams) GetProviderAuth() MigrationSourceParamsProviderAuth`

GetProviderAuth returns the ProviderAuth field if non-nil, zero value otherwise.

### GetProviderAuthOk

`func (o *MigrationSourceParams) GetProviderAuthOk() (*MigrationSourceParamsProviderAuth, bool)`

GetProviderAuthOk returns a tuple with the ProviderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAuth

`func (o *MigrationSourceParams) SetProviderAuth(v MigrationSourceParamsProviderAuth)`

SetProviderAuth sets ProviderAuth field to given value.


### GetBucketName

`func (o *MigrationSourceParams) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *MigrationSourceParams) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *MigrationSourceParams) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


