# RecursiveCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**ClusterHeader** | Pointer to **string** |  | [optional] 
**ClusterSummary** | **string** |  | 
**Nodes** | Pointer to [**[]ClusterNode**](ClusterNode.md) |  | [optional] 
**TotalNumberOfNodes** | **int32** |  | 
**Subclusters** | Pointer to [**[]RecursiveCluster**](RecursiveCluster.md) |  | [optional] 

## Methods

### NewRecursiveCluster

`func NewRecursiveCluster(clusterId string, clusterSummary string, totalNumberOfNodes int32, ) *RecursiveCluster`

NewRecursiveCluster instantiates a new RecursiveCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecursiveClusterWithDefaults

`func NewRecursiveClusterWithDefaults() *RecursiveCluster`

NewRecursiveClusterWithDefaults instantiates a new RecursiveCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *RecursiveCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RecursiveCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RecursiveCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterHeader

`func (o *RecursiveCluster) GetClusterHeader() string`

GetClusterHeader returns the ClusterHeader field if non-nil, zero value otherwise.

### GetClusterHeaderOk

`func (o *RecursiveCluster) GetClusterHeaderOk() (*string, bool)`

GetClusterHeaderOk returns a tuple with the ClusterHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHeader

`func (o *RecursiveCluster) SetClusterHeader(v string)`

SetClusterHeader sets ClusterHeader field to given value.

### HasClusterHeader

`func (o *RecursiveCluster) HasClusterHeader() bool`

HasClusterHeader returns a boolean if a field has been set.

### GetClusterSummary

`func (o *RecursiveCluster) GetClusterSummary() string`

GetClusterSummary returns the ClusterSummary field if non-nil, zero value otherwise.

### GetClusterSummaryOk

`func (o *RecursiveCluster) GetClusterSummaryOk() (*string, bool)`

GetClusterSummaryOk returns a tuple with the ClusterSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSummary

`func (o *RecursiveCluster) SetClusterSummary(v string)`

SetClusterSummary sets ClusterSummary field to given value.


### GetNodes

`func (o *RecursiveCluster) GetNodes() []ClusterNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *RecursiveCluster) GetNodesOk() (*[]ClusterNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *RecursiveCluster) SetNodes(v []ClusterNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *RecursiveCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetTotalNumberOfNodes

`func (o *RecursiveCluster) GetTotalNumberOfNodes() int32`

GetTotalNumberOfNodes returns the TotalNumberOfNodes field if non-nil, zero value otherwise.

### GetTotalNumberOfNodesOk

`func (o *RecursiveCluster) GetTotalNumberOfNodesOk() (*int32, bool)`

GetTotalNumberOfNodesOk returns a tuple with the TotalNumberOfNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfNodes

`func (o *RecursiveCluster) SetTotalNumberOfNodes(v int32)`

SetTotalNumberOfNodes sets TotalNumberOfNodes field to given value.


### GetSubclusters

`func (o *RecursiveCluster) GetSubclusters() []RecursiveCluster`

GetSubclusters returns the Subclusters field if non-nil, zero value otherwise.

### GetSubclustersOk

`func (o *RecursiveCluster) GetSubclustersOk() (*[]RecursiveCluster, bool)`

GetSubclustersOk returns a tuple with the Subclusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubclusters

`func (o *RecursiveCluster) SetSubclusters(v []RecursiveCluster)`

SetSubclusters sets Subclusters field to given value.

### HasSubclusters

`func (o *RecursiveCluster) HasSubclusters() bool`

HasSubclusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


