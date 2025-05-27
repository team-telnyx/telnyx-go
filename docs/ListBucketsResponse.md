# ListBucketsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**[]ListBucketsResponseBucketsInner**](ListBucketsResponseBucketsInner.md) |  | [optional] 

## Methods

### NewListBucketsResponse

`func NewListBucketsResponse() *ListBucketsResponse`

NewListBucketsResponse instantiates a new ListBucketsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBucketsResponseWithDefaults

`func NewListBucketsResponseWithDefaults() *ListBucketsResponse`

NewListBucketsResponseWithDefaults instantiates a new ListBucketsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ListBucketsResponse) GetBuckets() []ListBucketsResponseBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ListBucketsResponse) GetBucketsOk() (*[]ListBucketsResponseBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ListBucketsResponse) SetBuckets(v []ListBucketsResponseBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ListBucketsResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


