# ImportExternalVetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpId** | **string** | External vetting provider ID for the brand. | 
**VettingId** | **string** | Unique ID that identifies a vetting transaction performed by a vetting provider. This ID is provided by the vetting provider at time of vetting. | 
**VettingToken** | Pointer to **string** | Required by some providers for vetting record confirmation. | [optional] 

## Methods

### NewImportExternalVetting

`func NewImportExternalVetting(evpId string, vettingId string, ) *ImportExternalVetting`

NewImportExternalVetting instantiates a new ImportExternalVetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportExternalVettingWithDefaults

`func NewImportExternalVettingWithDefaults() *ImportExternalVetting`

NewImportExternalVettingWithDefaults instantiates a new ImportExternalVetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpId

`func (o *ImportExternalVetting) GetEvpId() string`

GetEvpId returns the EvpId field if non-nil, zero value otherwise.

### GetEvpIdOk

`func (o *ImportExternalVetting) GetEvpIdOk() (*string, bool)`

GetEvpIdOk returns a tuple with the EvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpId

`func (o *ImportExternalVetting) SetEvpId(v string)`

SetEvpId sets EvpId field to given value.


### GetVettingId

`func (o *ImportExternalVetting) GetVettingId() string`

GetVettingId returns the VettingId field if non-nil, zero value otherwise.

### GetVettingIdOk

`func (o *ImportExternalVetting) GetVettingIdOk() (*string, bool)`

GetVettingIdOk returns a tuple with the VettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingId

`func (o *ImportExternalVetting) SetVettingId(v string)`

SetVettingId sets VettingId field to given value.


### GetVettingToken

`func (o *ImportExternalVetting) GetVettingToken() string`

GetVettingToken returns the VettingToken field if non-nil, zero value otherwise.

### GetVettingTokenOk

`func (o *ImportExternalVetting) GetVettingTokenOk() (*string, bool)`

GetVettingTokenOk returns a tuple with the VettingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingToken

`func (o *ImportExternalVetting) SetVettingToken(v string)`

SetVettingToken sets VettingToken field to given value.

### HasVettingToken

`func (o *ImportExternalVetting) HasVettingToken() bool`

HasVettingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


