# GlobalIPHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**HealthCheckType** | Pointer to **string** | The Global IP health check type. | [optional] 
**HealthCheckParams** | Pointer to **map[string]interface{}** | A Global IP health check params. | [optional] 
**GlobalIpId** | Pointer to **string** | Global IP ID. | [optional] 

## Methods

### NewGlobalIPHealthCheck

`func NewGlobalIPHealthCheck() *GlobalIPHealthCheck`

NewGlobalIPHealthCheck instantiates a new GlobalIPHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIPHealthCheckWithDefaults

`func NewGlobalIPHealthCheckWithDefaults() *GlobalIPHealthCheck`

NewGlobalIPHealthCheckWithDefaults instantiates a new GlobalIPHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalIPHealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalIPHealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalIPHealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalIPHealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *GlobalIPHealthCheck) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GlobalIPHealthCheck) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GlobalIPHealthCheck) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GlobalIPHealthCheck) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalIPHealthCheck) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalIPHealthCheck) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalIPHealthCheck) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalIPHealthCheck) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalIPHealthCheck) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalIPHealthCheck) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalIPHealthCheck) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalIPHealthCheck) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHealthCheckType

`func (o *GlobalIPHealthCheck) GetHealthCheckType() string`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *GlobalIPHealthCheck) GetHealthCheckTypeOk() (*string, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *GlobalIPHealthCheck) SetHealthCheckType(v string)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *GlobalIPHealthCheck) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetHealthCheckParams

`func (o *GlobalIPHealthCheck) GetHealthCheckParams() map[string]interface{}`

GetHealthCheckParams returns the HealthCheckParams field if non-nil, zero value otherwise.

### GetHealthCheckParamsOk

`func (o *GlobalIPHealthCheck) GetHealthCheckParamsOk() (*map[string]interface{}, bool)`

GetHealthCheckParamsOk returns a tuple with the HealthCheckParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckParams

`func (o *GlobalIPHealthCheck) SetHealthCheckParams(v map[string]interface{})`

SetHealthCheckParams sets HealthCheckParams field to given value.

### HasHealthCheckParams

`func (o *GlobalIPHealthCheck) HasHealthCheckParams() bool`

HasHealthCheckParams returns a boolean if a field has been set.

### GetGlobalIpId

`func (o *GlobalIPHealthCheck) GetGlobalIpId() string`

GetGlobalIpId returns the GlobalIpId field if non-nil, zero value otherwise.

### GetGlobalIpIdOk

`func (o *GlobalIPHealthCheck) GetGlobalIpIdOk() (*string, bool)`

GetGlobalIpIdOk returns a tuple with the GlobalIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIpId

`func (o *GlobalIPHealthCheck) SetGlobalIpId(v string)`

SetGlobalIpId sets GlobalIpId field to given value.

### HasGlobalIpId

`func (o *GlobalIPHealthCheck) HasGlobalIpId() bool`

HasGlobalIpId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


