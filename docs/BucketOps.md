# BucketOps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesSent** | Pointer to **int32** | The number of bytes sent | [optional] 
**BytesReceived** | Pointer to **int32** | The number of bytes received | [optional] 
**Ops** | Pointer to **int32** | The number of operations | [optional] 
**SuccessfulOps** | Pointer to **int32** | The number of successful operations | [optional] 
**Category** | Pointer to **string** | The category of the bucket operation | [optional] 

## Methods

### NewBucketOps

`func NewBucketOps() *BucketOps`

NewBucketOps instantiates a new BucketOps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketOpsWithDefaults

`func NewBucketOpsWithDefaults() *BucketOps`

NewBucketOpsWithDefaults instantiates a new BucketOps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesSent

`func (o *BucketOps) GetBytesSent() int32`

GetBytesSent returns the BytesSent field if non-nil, zero value otherwise.

### GetBytesSentOk

`func (o *BucketOps) GetBytesSentOk() (*int32, bool)`

GetBytesSentOk returns a tuple with the BytesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSent

`func (o *BucketOps) SetBytesSent(v int32)`

SetBytesSent sets BytesSent field to given value.

### HasBytesSent

`func (o *BucketOps) HasBytesSent() bool`

HasBytesSent returns a boolean if a field has been set.

### GetBytesReceived

`func (o *BucketOps) GetBytesReceived() int32`

GetBytesReceived returns the BytesReceived field if non-nil, zero value otherwise.

### GetBytesReceivedOk

`func (o *BucketOps) GetBytesReceivedOk() (*int32, bool)`

GetBytesReceivedOk returns a tuple with the BytesReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesReceived

`func (o *BucketOps) SetBytesReceived(v int32)`

SetBytesReceived sets BytesReceived field to given value.

### HasBytesReceived

`func (o *BucketOps) HasBytesReceived() bool`

HasBytesReceived returns a boolean if a field has been set.

### GetOps

`func (o *BucketOps) GetOps() int32`

GetOps returns the Ops field if non-nil, zero value otherwise.

### GetOpsOk

`func (o *BucketOps) GetOpsOk() (*int32, bool)`

GetOpsOk returns a tuple with the Ops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOps

`func (o *BucketOps) SetOps(v int32)`

SetOps sets Ops field to given value.

### HasOps

`func (o *BucketOps) HasOps() bool`

HasOps returns a boolean if a field has been set.

### GetSuccessfulOps

`func (o *BucketOps) GetSuccessfulOps() int32`

GetSuccessfulOps returns the SuccessfulOps field if non-nil, zero value otherwise.

### GetSuccessfulOpsOk

`func (o *BucketOps) GetSuccessfulOpsOk() (*int32, bool)`

GetSuccessfulOpsOk returns a tuple with the SuccessfulOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulOps

`func (o *BucketOps) SetSuccessfulOps(v int32)`

SetSuccessfulOps sets SuccessfulOps field to given value.

### HasSuccessfulOps

`func (o *BucketOps) HasSuccessfulOps() bool`

HasSuccessfulOps returns a boolean if a field has been set.

### GetCategory

`func (o *BucketOps) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BucketOps) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BucketOps) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BucketOps) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


