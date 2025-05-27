# BucketIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketIds** | **[]string** | List of [embedded storage buckets](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding) to use for retrieval-augmented generation. | 
**MaxNumResults** | Pointer to **int32** | The maximum number of results to retrieve as context for the language model. | [optional] 

## Methods

### NewBucketIds

`func NewBucketIds(bucketIds []string, ) *BucketIds`

NewBucketIds instantiates a new BucketIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketIdsWithDefaults

`func NewBucketIdsWithDefaults() *BucketIds`

NewBucketIdsWithDefaults instantiates a new BucketIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketIds

`func (o *BucketIds) GetBucketIds() []string`

GetBucketIds returns the BucketIds field if non-nil, zero value otherwise.

### GetBucketIdsOk

`func (o *BucketIds) GetBucketIdsOk() (*[]string, bool)`

GetBucketIdsOk returns a tuple with the BucketIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketIds

`func (o *BucketIds) SetBucketIds(v []string)`

SetBucketIds sets BucketIds field to given value.


### GetMaxNumResults

`func (o *BucketIds) GetMaxNumResults() int32`

GetMaxNumResults returns the MaxNumResults field if non-nil, zero value otherwise.

### GetMaxNumResultsOk

`func (o *BucketIds) GetMaxNumResultsOk() (*int32, bool)`

GetMaxNumResultsOk returns a tuple with the MaxNumResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumResults

`func (o *BucketIds) SetMaxNumResults(v int32)`

SetMaxNumResults sets MaxNumResults field to given value.

### HasMaxNumResults

`func (o *BucketIds) HasMaxNumResults() bool`

HasMaxNumResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


