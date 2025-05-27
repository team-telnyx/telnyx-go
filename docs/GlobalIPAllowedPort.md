# GlobalIPAllowedPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**ProtocolCode** | Pointer to **string** | The Global IP Protocol code. | [optional] [readonly] 
**Name** | Pointer to **string** | A name for the Global IP ports range. | [optional] 
**FirstPort** | Pointer to **int32** | First port of a range. | [optional] 
**LastPort** | Pointer to **int32** | Last port of a range. | [optional] 

## Methods

### NewGlobalIPAllowedPort

`func NewGlobalIPAllowedPort() *GlobalIPAllowedPort`

NewGlobalIPAllowedPort instantiates a new GlobalIPAllowedPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIPAllowedPortWithDefaults

`func NewGlobalIPAllowedPortWithDefaults() *GlobalIPAllowedPort`

NewGlobalIPAllowedPortWithDefaults instantiates a new GlobalIPAllowedPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalIPAllowedPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalIPAllowedPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalIPAllowedPort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalIPAllowedPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *GlobalIPAllowedPort) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GlobalIPAllowedPort) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GlobalIPAllowedPort) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GlobalIPAllowedPort) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetProtocolCode

`func (o *GlobalIPAllowedPort) GetProtocolCode() string`

GetProtocolCode returns the ProtocolCode field if non-nil, zero value otherwise.

### GetProtocolCodeOk

`func (o *GlobalIPAllowedPort) GetProtocolCodeOk() (*string, bool)`

GetProtocolCodeOk returns a tuple with the ProtocolCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolCode

`func (o *GlobalIPAllowedPort) SetProtocolCode(v string)`

SetProtocolCode sets ProtocolCode field to given value.

### HasProtocolCode

`func (o *GlobalIPAllowedPort) HasProtocolCode() bool`

HasProtocolCode returns a boolean if a field has been set.

### GetName

`func (o *GlobalIPAllowedPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalIPAllowedPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalIPAllowedPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalIPAllowedPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFirstPort

`func (o *GlobalIPAllowedPort) GetFirstPort() int32`

GetFirstPort returns the FirstPort field if non-nil, zero value otherwise.

### GetFirstPortOk

`func (o *GlobalIPAllowedPort) GetFirstPortOk() (*int32, bool)`

GetFirstPortOk returns a tuple with the FirstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPort

`func (o *GlobalIPAllowedPort) SetFirstPort(v int32)`

SetFirstPort sets FirstPort field to given value.

### HasFirstPort

`func (o *GlobalIPAllowedPort) HasFirstPort() bool`

HasFirstPort returns a boolean if a field has been set.

### GetLastPort

`func (o *GlobalIPAllowedPort) GetLastPort() int32`

GetLastPort returns the LastPort field if non-nil, zero value otherwise.

### GetLastPortOk

`func (o *GlobalIPAllowedPort) GetLastPortOk() (*int32, bool)`

GetLastPortOk returns a tuple with the LastPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPort

`func (o *GlobalIPAllowedPort) SetLastPort(v int32)`

SetLastPort sets LastPort field to given value.

### HasLastPort

`func (o *GlobalIPAllowedPort) HasLastPort() bool`

HasLastPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


