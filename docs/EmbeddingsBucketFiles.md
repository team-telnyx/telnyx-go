# EmbeddingsBucketFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**Status** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**LastEmbeddedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorReason** | Pointer to **string** |  | [optional] 

## Methods

### NewEmbeddingsBucketFiles

`func NewEmbeddingsBucketFiles(filename string, status string, createdAt time.Time, ) *EmbeddingsBucketFiles`

NewEmbeddingsBucketFiles instantiates a new EmbeddingsBucketFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingsBucketFilesWithDefaults

`func NewEmbeddingsBucketFilesWithDefaults() *EmbeddingsBucketFiles`

NewEmbeddingsBucketFilesWithDefaults instantiates a new EmbeddingsBucketFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *EmbeddingsBucketFiles) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EmbeddingsBucketFiles) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EmbeddingsBucketFiles) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetStatus

`func (o *EmbeddingsBucketFiles) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmbeddingsBucketFiles) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmbeddingsBucketFiles) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *EmbeddingsBucketFiles) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmbeddingsBucketFiles) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmbeddingsBucketFiles) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EmbeddingsBucketFiles) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmbeddingsBucketFiles) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmbeddingsBucketFiles) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EmbeddingsBucketFiles) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastEmbeddedAt

`func (o *EmbeddingsBucketFiles) GetLastEmbeddedAt() time.Time`

GetLastEmbeddedAt returns the LastEmbeddedAt field if non-nil, zero value otherwise.

### GetLastEmbeddedAtOk

`func (o *EmbeddingsBucketFiles) GetLastEmbeddedAtOk() (*time.Time, bool)`

GetLastEmbeddedAtOk returns a tuple with the LastEmbeddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEmbeddedAt

`func (o *EmbeddingsBucketFiles) SetLastEmbeddedAt(v time.Time)`

SetLastEmbeddedAt sets LastEmbeddedAt field to given value.

### HasLastEmbeddedAt

`func (o *EmbeddingsBucketFiles) HasLastEmbeddedAt() bool`

HasLastEmbeddedAt returns a boolean if a field has been set.

### GetErrorReason

`func (o *EmbeddingsBucketFiles) GetErrorReason() string`

GetErrorReason returns the ErrorReason field if non-nil, zero value otherwise.

### GetErrorReasonOk

`func (o *EmbeddingsBucketFiles) GetErrorReasonOk() (*string, bool)`

GetErrorReasonOk returns a tuple with the ErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorReason

`func (o *EmbeddingsBucketFiles) SetErrorReason(v string)`

SetErrorReason sets ErrorReason field to given value.

### HasErrorReason

`func (o *EmbeddingsBucketFiles) HasErrorReason() bool`

HasErrorReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


