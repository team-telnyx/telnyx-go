# UpdateIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | ID of the IP Connection to which this IP should be attached. | [optional] 
**IpAddress** | **string** | IP adddress represented by this resource. | 
**Port** | Pointer to **int32** | Port to use when connecting to this IP. | [optional] [default to 5060]

## Methods

### NewUpdateIpRequest

`func NewUpdateIpRequest(ipAddress string, ) *UpdateIpRequest`

NewUpdateIpRequest instantiates a new UpdateIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpRequestWithDefaults

`func NewUpdateIpRequestWithDefaults() *UpdateIpRequest`

NewUpdateIpRequestWithDefaults instantiates a new UpdateIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *UpdateIpRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UpdateIpRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UpdateIpRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *UpdateIpRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetIpAddress

`func (o *UpdateIpRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateIpRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateIpRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *UpdateIpRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateIpRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateIpRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateIpRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


