# Fqdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**ConnectionId** | Pointer to **string** | ID of the FQDN connection to which this FQDN is attached. | [optional] 
**Fqdn** | Pointer to **string** | FQDN represented by this resource. | [optional] 
**Port** | Pointer to **int32** | Port to use when connecting to this FQDN. | [optional] [default to 5060]
**DnsRecordType** | Pointer to **string** | The DNS record type for the FQDN. For cases where a port is not set, the DNS record type must be &#39;srv&#39;. For cases where a port is set, the DNS record type must be &#39;a&#39;. If the DNS record type is &#39;a&#39; and a port is not specified, 5060 will be used. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewFqdn

`func NewFqdn() *Fqdn`

NewFqdn instantiates a new Fqdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFqdnWithDefaults

`func NewFqdnWithDefaults() *Fqdn`

NewFqdnWithDefaults instantiates a new Fqdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fqdn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fqdn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fqdn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fqdn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *Fqdn) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Fqdn) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Fqdn) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Fqdn) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetConnectionId

`func (o *Fqdn) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Fqdn) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Fqdn) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Fqdn) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetFqdn

`func (o *Fqdn) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Fqdn) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Fqdn) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Fqdn) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetPort

`func (o *Fqdn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Fqdn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Fqdn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Fqdn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDnsRecordType

`func (o *Fqdn) GetDnsRecordType() string`

GetDnsRecordType returns the DnsRecordType field if non-nil, zero value otherwise.

### GetDnsRecordTypeOk

`func (o *Fqdn) GetDnsRecordTypeOk() (*string, bool)`

GetDnsRecordTypeOk returns a tuple with the DnsRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecordType

`func (o *Fqdn) SetDnsRecordType(v string)`

SetDnsRecordType sets DnsRecordType field to given value.

### HasDnsRecordType

`func (o *Fqdn) HasDnsRecordType() bool`

HasDnsRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Fqdn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Fqdn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Fqdn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Fqdn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Fqdn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Fqdn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Fqdn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Fqdn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


