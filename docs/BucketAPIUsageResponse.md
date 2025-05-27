# BucketAPIUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]BucketOps**](BucketOps.md) |  | [optional] 
**Total** | Pointer to [**BucketOpsTotal**](BucketOpsTotal.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** | The time the usage was recorded | [optional] 

## Methods

### NewBucketAPIUsageResponse

`func NewBucketAPIUsageResponse() *BucketAPIUsageResponse`

NewBucketAPIUsageResponse instantiates a new BucketAPIUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketAPIUsageResponseWithDefaults

`func NewBucketAPIUsageResponseWithDefaults() *BucketAPIUsageResponse`

NewBucketAPIUsageResponseWithDefaults instantiates a new BucketAPIUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *BucketAPIUsageResponse) GetCategories() []BucketOps`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BucketAPIUsageResponse) GetCategoriesOk() (*[]BucketOps, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BucketAPIUsageResponse) SetCategories(v []BucketOps)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *BucketAPIUsageResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTotal

`func (o *BucketAPIUsageResponse) GetTotal() BucketOpsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BucketAPIUsageResponse) GetTotalOk() (*BucketOpsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BucketAPIUsageResponse) SetTotal(v BucketOpsTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BucketAPIUsageResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTimestamp

`func (o *BucketAPIUsageResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BucketAPIUsageResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BucketAPIUsageResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BucketAPIUsageResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


