# BucketNotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to [**[]NotFoundError**](NotFoundError.md) |  | [optional] 

## Methods

### NewBucketNotFoundError

`func NewBucketNotFoundError() *BucketNotFoundError`

NewBucketNotFoundError instantiates a new BucketNotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketNotFoundErrorWithDefaults

`func NewBucketNotFoundErrorWithDefaults() *BucketNotFoundError`

NewBucketNotFoundErrorWithDefaults instantiates a new BucketNotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *BucketNotFoundError) GetDetail() []NotFoundError`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BucketNotFoundError) GetDetailOk() (*[]NotFoundError, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BucketNotFoundError) SetDetail(v []NotFoundError)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BucketNotFoundError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


