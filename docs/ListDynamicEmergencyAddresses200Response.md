# ListDynamicEmergencyAddresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DynamicEmergencyAddress**](DynamicEmergencyAddress.md) |  | [optional] 
**Meta** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 

## Methods

### NewListDynamicEmergencyAddresses200Response

`func NewListDynamicEmergencyAddresses200Response() *ListDynamicEmergencyAddresses200Response`

NewListDynamicEmergencyAddresses200Response instantiates a new ListDynamicEmergencyAddresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDynamicEmergencyAddresses200ResponseWithDefaults

`func NewListDynamicEmergencyAddresses200ResponseWithDefaults() *ListDynamicEmergencyAddresses200Response`

NewListDynamicEmergencyAddresses200ResponseWithDefaults instantiates a new ListDynamicEmergencyAddresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDynamicEmergencyAddresses200Response) GetData() []DynamicEmergencyAddress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDynamicEmergencyAddresses200Response) GetDataOk() (*[]DynamicEmergencyAddress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDynamicEmergencyAddresses200Response) SetData(v []DynamicEmergencyAddress)`

SetData sets Data field to given value.

### HasData

`func (o *ListDynamicEmergencyAddresses200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListDynamicEmergencyAddresses200Response) GetMeta() Metadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDynamicEmergencyAddresses200Response) GetMetaOk() (*Metadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDynamicEmergencyAddresses200Response) SetMeta(v Metadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDynamicEmergencyAddresses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


