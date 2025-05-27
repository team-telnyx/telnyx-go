# BucketUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | The size of the bucket in bytes | [optional] 
**SizeKb** | Pointer to **int32** | The size of the bucket in kilobytes | [optional] 
**NumObjects** | Pointer to **int32** | The number of objects in the bucket | [optional] 
**Timestamp** | Pointer to **time.Time** | The time the snapshot was taken | [optional] 

## Methods

### NewBucketUsage

`func NewBucketUsage() *BucketUsage`

NewBucketUsage instantiates a new BucketUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketUsageWithDefaults

`func NewBucketUsageWithDefaults() *BucketUsage`

NewBucketUsageWithDefaults instantiates a new BucketUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *BucketUsage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BucketUsage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BucketUsage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BucketUsage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeKb

`func (o *BucketUsage) GetSizeKb() int32`

GetSizeKb returns the SizeKb field if non-nil, zero value otherwise.

### GetSizeKbOk

`func (o *BucketUsage) GetSizeKbOk() (*int32, bool)`

GetSizeKbOk returns a tuple with the SizeKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeKb

`func (o *BucketUsage) SetSizeKb(v int32)`

SetSizeKb sets SizeKb field to given value.

### HasSizeKb

`func (o *BucketUsage) HasSizeKb() bool`

HasSizeKb returns a boolean if a field has been set.

### GetNumObjects

`func (o *BucketUsage) GetNumObjects() int32`

GetNumObjects returns the NumObjects field if non-nil, zero value otherwise.

### GetNumObjectsOk

`func (o *BucketUsage) GetNumObjectsOk() (*int32, bool)`

GetNumObjectsOk returns a tuple with the NumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjects

`func (o *BucketUsage) SetNumObjects(v int32)`

SetNumObjects sets NumObjects field to given value.

### HasNumObjects

`func (o *BucketUsage) HasNumObjects() bool`

HasNumObjects returns a boolean if a field has been set.

### GetTimestamp

`func (o *BucketUsage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BucketUsage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BucketUsage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BucketUsage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


