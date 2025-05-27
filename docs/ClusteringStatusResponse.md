# ClusteringStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**Bucket** | **string** |  | 
**Clusters** | [**[]RecursiveCluster**](RecursiveCluster.md) |  | 

## Methods

### NewClusteringStatusResponse

`func NewClusteringStatusResponse(status TaskStatus, bucket string, clusters []RecursiveCluster, ) *ClusteringStatusResponse`

NewClusteringStatusResponse instantiates a new ClusteringStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusteringStatusResponseWithDefaults

`func NewClusteringStatusResponseWithDefaults() *ClusteringStatusResponse`

NewClusteringStatusResponseWithDefaults instantiates a new ClusteringStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ClusteringStatusResponse) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusteringStatusResponse) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusteringStatusResponse) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetBucket

`func (o *ClusteringStatusResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ClusteringStatusResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ClusteringStatusResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetClusters

`func (o *ClusteringStatusResponse) GetClusters() []RecursiveCluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ClusteringStatusResponse) GetClustersOk() (*[]RecursiveCluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ClusteringStatusResponse) SetClusters(v []RecursiveCluster)`

SetClusters sets Clusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


