# IntegrationSecretsListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IntegrationSecret**](IntegrationSecret.md) |  | 
**Meta** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewIntegrationSecretsListData

`func NewIntegrationSecretsListData(data []IntegrationSecret, meta Metadata, ) *IntegrationSecretsListData`

NewIntegrationSecretsListData instantiates a new IntegrationSecretsListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSecretsListDataWithDefaults

`func NewIntegrationSecretsListDataWithDefaults() *IntegrationSecretsListData`

NewIntegrationSecretsListDataWithDefaults instantiates a new IntegrationSecretsListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IntegrationSecretsListData) GetData() []IntegrationSecret`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IntegrationSecretsListData) GetDataOk() (*[]IntegrationSecret, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IntegrationSecretsListData) SetData(v []IntegrationSecret)`

SetData sets Data field to given value.


### GetMeta

`func (o *IntegrationSecretsListData) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IntegrationSecretsListData) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IntegrationSecretsListData) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


