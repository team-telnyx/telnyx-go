# CustomStorageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | **string** |  | 
**Configuration** | [**CustomStorageConfigurationConfiguration**](CustomStorageConfigurationConfiguration.md) |  | 

## Methods

### NewCustomStorageConfiguration

`func NewCustomStorageConfiguration(backend string, configuration CustomStorageConfigurationConfiguration, ) *CustomStorageConfiguration`

NewCustomStorageConfiguration instantiates a new CustomStorageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomStorageConfigurationWithDefaults

`func NewCustomStorageConfigurationWithDefaults() *CustomStorageConfiguration`

NewCustomStorageConfigurationWithDefaults instantiates a new CustomStorageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *CustomStorageConfiguration) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *CustomStorageConfiguration) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *CustomStorageConfiguration) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetConfiguration

`func (o *CustomStorageConfiguration) GetConfiguration() CustomStorageConfigurationConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CustomStorageConfiguration) GetConfigurationOk() (*CustomStorageConfigurationConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CustomStorageConfiguration) SetConfiguration(v CustomStorageConfigurationConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


