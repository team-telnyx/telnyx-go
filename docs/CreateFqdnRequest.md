# CreateFqdnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | ID of the FQDN connection to which this IP should be attached. | 
**Fqdn** | **string** | FQDN represented by this resource. | 
**Port** | Pointer to **NullableInt32** | Port to use when connecting to this FQDN. | [optional] [default to 5060]
**DnsRecordType** | **string** | The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be &#39;srv&#39;. For cases where a port is set, the DNS record type must be &#39;a&#39;. If the DNS record type is &#39;a&#39; and a port is not specified, 5060 will be used. | 

## Methods

### NewCreateFqdnRequest

`func NewCreateFqdnRequest(connectionId string, fqdn string, dnsRecordType string, ) *CreateFqdnRequest`

NewCreateFqdnRequest instantiates a new CreateFqdnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFqdnRequestWithDefaults

`func NewCreateFqdnRequestWithDefaults() *CreateFqdnRequest`

NewCreateFqdnRequestWithDefaults instantiates a new CreateFqdnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *CreateFqdnRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateFqdnRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateFqdnRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFqdn

`func (o *CreateFqdnRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateFqdnRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateFqdnRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetPort

`func (o *CreateFqdnRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateFqdnRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateFqdnRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateFqdnRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CreateFqdnRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CreateFqdnRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDnsRecordType

`func (o *CreateFqdnRequest) GetDnsRecordType() string`

GetDnsRecordType returns the DnsRecordType field if non-nil, zero value otherwise.

### GetDnsRecordTypeOk

`func (o *CreateFqdnRequest) GetDnsRecordTypeOk() (*string, bool)`

GetDnsRecordTypeOk returns a tuple with the DnsRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecordType

`func (o *CreateFqdnRequest) SetDnsRecordType(v string)`

SetDnsRecordType sets DnsRecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


