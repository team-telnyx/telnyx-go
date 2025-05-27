# CreateExternalConnectionUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberIds** | Pointer to **[]string** |  | [optional] 
**Usage** | Pointer to **string** | The use case of the upload request. NOTE: &#x60;calling_user_assignment&#x60; is not supported for toll free numbers. | [optional] 
**AdditionalUsages** | Pointer to **[]string** |  | [optional] 
**LocationId** | Pointer to **string** | Identifies the location to assign all phone numbers to. | [optional] 
**CivicAddressId** | Pointer to **string** | Identifies the civic address to assign all phone numbers to. | [optional] 

## Methods

### NewCreateExternalConnectionUploadRequest

`func NewCreateExternalConnectionUploadRequest() *CreateExternalConnectionUploadRequest`

NewCreateExternalConnectionUploadRequest instantiates a new CreateExternalConnectionUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalConnectionUploadRequestWithDefaults

`func NewCreateExternalConnectionUploadRequestWithDefaults() *CreateExternalConnectionUploadRequest`

NewCreateExternalConnectionUploadRequestWithDefaults instantiates a new CreateExternalConnectionUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberIds

`func (o *CreateExternalConnectionUploadRequest) GetNumberIds() []string`

GetNumberIds returns the NumberIds field if non-nil, zero value otherwise.

### GetNumberIdsOk

`func (o *CreateExternalConnectionUploadRequest) GetNumberIdsOk() (*[]string, bool)`

GetNumberIdsOk returns a tuple with the NumberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberIds

`func (o *CreateExternalConnectionUploadRequest) SetNumberIds(v []string)`

SetNumberIds sets NumberIds field to given value.

### HasNumberIds

`func (o *CreateExternalConnectionUploadRequest) HasNumberIds() bool`

HasNumberIds returns a boolean if a field has been set.

### GetUsage

`func (o *CreateExternalConnectionUploadRequest) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateExternalConnectionUploadRequest) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateExternalConnectionUploadRequest) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CreateExternalConnectionUploadRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetAdditionalUsages

`func (o *CreateExternalConnectionUploadRequest) GetAdditionalUsages() []string`

GetAdditionalUsages returns the AdditionalUsages field if non-nil, zero value otherwise.

### GetAdditionalUsagesOk

`func (o *CreateExternalConnectionUploadRequest) GetAdditionalUsagesOk() (*[]string, bool)`

GetAdditionalUsagesOk returns a tuple with the AdditionalUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsages

`func (o *CreateExternalConnectionUploadRequest) SetAdditionalUsages(v []string)`

SetAdditionalUsages sets AdditionalUsages field to given value.

### HasAdditionalUsages

`func (o *CreateExternalConnectionUploadRequest) HasAdditionalUsages() bool`

HasAdditionalUsages returns a boolean if a field has been set.

### GetLocationId

`func (o *CreateExternalConnectionUploadRequest) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CreateExternalConnectionUploadRequest) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CreateExternalConnectionUploadRequest) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *CreateExternalConnectionUploadRequest) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetCivicAddressId

`func (o *CreateExternalConnectionUploadRequest) GetCivicAddressId() string`

GetCivicAddressId returns the CivicAddressId field if non-nil, zero value otherwise.

### GetCivicAddressIdOk

`func (o *CreateExternalConnectionUploadRequest) GetCivicAddressIdOk() (*string, bool)`

GetCivicAddressIdOk returns a tuple with the CivicAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddressId

`func (o *CreateExternalConnectionUploadRequest) SetCivicAddressId(v string)`

SetCivicAddressId sets CivicAddressId field to given value.

### HasCivicAddressId

`func (o *CreateExternalConnectionUploadRequest) HasCivicAddressId() bool`

HasCivicAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


