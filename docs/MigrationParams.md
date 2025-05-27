# MigrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the data migration. | [optional] [readonly] 
**SourceId** | **string** | ID of the Migration Source from which to migrate data. | 
**TargetBucketName** | **string** | Bucket name to migrate the data into. Will default to the same name as the &#x60;source_bucket_name&#x60;. | 
**TargetRegion** | **string** | Telnyx Cloud Storage region to migrate the data to. | 
**Refresh** | Pointer to **bool** | If true, will continue to poll the source bucket to ensure new data is continually migrated over. | [optional] 
**LastCopy** | Pointer to **time.Time** | Time when data migration was last copied from the source. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the migration. | [optional] [readonly] 
**BytesToMigrate** | Pointer to **int32** | Total amount of data found in source bucket to migrate. | [optional] [readonly] 
**BytesMigrated** | Pointer to **int32** | Total amount of data that has been succesfully migrated. | [optional] [readonly] 
**Speed** | Pointer to **int32** | Current speed of the migration. | [optional] [readonly] 
**Eta** | Pointer to **time.Time** | Estimated time the migration will complete. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Time when data migration was created | [optional] [readonly] 

## Methods

### NewMigrationParams

`func NewMigrationParams(sourceId string, targetBucketName string, targetRegion string, ) *MigrationParams`

NewMigrationParams instantiates a new MigrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationParamsWithDefaults

`func NewMigrationParamsWithDefaults() *MigrationParams`

NewMigrationParamsWithDefaults instantiates a new MigrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MigrationParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MigrationParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MigrationParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MigrationParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *MigrationParams) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MigrationParams) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MigrationParams) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetBucketName

`func (o *MigrationParams) GetTargetBucketName() string`

GetTargetBucketName returns the TargetBucketName field if non-nil, zero value otherwise.

### GetTargetBucketNameOk

`func (o *MigrationParams) GetTargetBucketNameOk() (*string, bool)`

GetTargetBucketNameOk returns a tuple with the TargetBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBucketName

`func (o *MigrationParams) SetTargetBucketName(v string)`

SetTargetBucketName sets TargetBucketName field to given value.


### GetTargetRegion

`func (o *MigrationParams) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### GetTargetRegionOk

`func (o *MigrationParams) GetTargetRegionOk() (*string, bool)`

GetTargetRegionOk returns a tuple with the TargetRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRegion

`func (o *MigrationParams) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


### GetRefresh

`func (o *MigrationParams) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *MigrationParams) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *MigrationParams) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *MigrationParams) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetLastCopy

`func (o *MigrationParams) GetLastCopy() time.Time`

GetLastCopy returns the LastCopy field if non-nil, zero value otherwise.

### GetLastCopyOk

`func (o *MigrationParams) GetLastCopyOk() (*time.Time, bool)`

GetLastCopyOk returns a tuple with the LastCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCopy

`func (o *MigrationParams) SetLastCopy(v time.Time)`

SetLastCopy sets LastCopy field to given value.

### HasLastCopy

`func (o *MigrationParams) HasLastCopy() bool`

HasLastCopy returns a boolean if a field has been set.

### GetStatus

`func (o *MigrationParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MigrationParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MigrationParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MigrationParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBytesToMigrate

`func (o *MigrationParams) GetBytesToMigrate() int32`

GetBytesToMigrate returns the BytesToMigrate field if non-nil, zero value otherwise.

### GetBytesToMigrateOk

`func (o *MigrationParams) GetBytesToMigrateOk() (*int32, bool)`

GetBytesToMigrateOk returns a tuple with the BytesToMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesToMigrate

`func (o *MigrationParams) SetBytesToMigrate(v int32)`

SetBytesToMigrate sets BytesToMigrate field to given value.

### HasBytesToMigrate

`func (o *MigrationParams) HasBytesToMigrate() bool`

HasBytesToMigrate returns a boolean if a field has been set.

### GetBytesMigrated

`func (o *MigrationParams) GetBytesMigrated() int32`

GetBytesMigrated returns the BytesMigrated field if non-nil, zero value otherwise.

### GetBytesMigratedOk

`func (o *MigrationParams) GetBytesMigratedOk() (*int32, bool)`

GetBytesMigratedOk returns a tuple with the BytesMigrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesMigrated

`func (o *MigrationParams) SetBytesMigrated(v int32)`

SetBytesMigrated sets BytesMigrated field to given value.

### HasBytesMigrated

`func (o *MigrationParams) HasBytesMigrated() bool`

HasBytesMigrated returns a boolean if a field has been set.

### GetSpeed

`func (o *MigrationParams) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MigrationParams) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MigrationParams) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MigrationParams) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetEta

`func (o *MigrationParams) GetEta() time.Time`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *MigrationParams) GetEtaOk() (*time.Time, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *MigrationParams) SetEta(v time.Time)`

SetEta sets Eta field to given value.

### HasEta

`func (o *MigrationParams) HasEta() bool`

HasEta returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MigrationParams) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MigrationParams) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MigrationParams) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MigrationParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


