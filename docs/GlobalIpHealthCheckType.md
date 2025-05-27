# GlobalIpHealthCheckType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**HealthCheckType** | Pointer to **string** | Global IP Health check type. | [optional] 
**HealthCheckParams** | Pointer to **map[string]interface{}** | Global IP Health check params. | [optional] 

## Methods

### NewGlobalIpHealthCheckType

`func NewGlobalIpHealthCheckType() *GlobalIpHealthCheckType`

NewGlobalIpHealthCheckType instantiates a new GlobalIpHealthCheckType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpHealthCheckTypeWithDefaults

`func NewGlobalIpHealthCheckTypeWithDefaults() *GlobalIpHealthCheckType`

NewGlobalIpHealthCheckTypeWithDefaults instantiates a new GlobalIpHealthCheckType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *GlobalIpHealthCheckType) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GlobalIpHealthCheckType) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GlobalIpHealthCheckType) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GlobalIpHealthCheckType) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetHealthCheckType

`func (o *GlobalIpHealthCheckType) GetHealthCheckType() string`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *GlobalIpHealthCheckType) GetHealthCheckTypeOk() (*string, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *GlobalIpHealthCheckType) SetHealthCheckType(v string)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *GlobalIpHealthCheckType) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetHealthCheckParams

`func (o *GlobalIpHealthCheckType) GetHealthCheckParams() map[string]interface{}`

GetHealthCheckParams returns the HealthCheckParams field if non-nil, zero value otherwise.

### GetHealthCheckParamsOk

`func (o *GlobalIpHealthCheckType) GetHealthCheckParamsOk() (*map[string]interface{}, bool)`

GetHealthCheckParamsOk returns a tuple with the HealthCheckParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckParams

`func (o *GlobalIpHealthCheckType) SetHealthCheckParams(v map[string]interface{})`

SetHealthCheckParams sets HealthCheckParams field to given value.

### HasHealthCheckParams

`func (o *GlobalIpHealthCheckType) HasHealthCheckParams() bool`

HasHealthCheckParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


