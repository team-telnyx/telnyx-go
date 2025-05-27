# ExternalVetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvpId** | Pointer to **string** | External vetting provider ID for the brand. | [optional] 
**VettingId** | Pointer to **string** | Unique ID that identifies a vetting transaction performed by a vetting provider. This ID is provided by the vetting provider at time of vetting. | [optional] 
**VettingToken** | Pointer to **string** | Required by some providers for vetting record confirmation. | [optional] 
**VettingScore** | Pointer to **int32** | Vetting score ranging from 0-100. | [optional] 
**VettingClass** | Pointer to **string** | Identifies the vetting classification. | [optional] 
**VettedDate** | Pointer to **string** | Vetting effective date. This is the date when vetting was completed, or the starting effective date in ISO 8601 format. If this date is missing, then the vetting was not complete or not valid. | [optional] 
**CreateDate** | Pointer to **string** | Vetting submission date. This is the date when the vetting request is generated in ISO 8601 format. | [optional] 

## Methods

### NewExternalVetting

`func NewExternalVetting() *ExternalVetting`

NewExternalVetting instantiates a new ExternalVetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalVettingWithDefaults

`func NewExternalVettingWithDefaults() *ExternalVetting`

NewExternalVettingWithDefaults instantiates a new ExternalVetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvpId

`func (o *ExternalVetting) GetEvpId() string`

GetEvpId returns the EvpId field if non-nil, zero value otherwise.

### GetEvpIdOk

`func (o *ExternalVetting) GetEvpIdOk() (*string, bool)`

GetEvpIdOk returns a tuple with the EvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvpId

`func (o *ExternalVetting) SetEvpId(v string)`

SetEvpId sets EvpId field to given value.

### HasEvpId

`func (o *ExternalVetting) HasEvpId() bool`

HasEvpId returns a boolean if a field has been set.

### GetVettingId

`func (o *ExternalVetting) GetVettingId() string`

GetVettingId returns the VettingId field if non-nil, zero value otherwise.

### GetVettingIdOk

`func (o *ExternalVetting) GetVettingIdOk() (*string, bool)`

GetVettingIdOk returns a tuple with the VettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingId

`func (o *ExternalVetting) SetVettingId(v string)`

SetVettingId sets VettingId field to given value.

### HasVettingId

`func (o *ExternalVetting) HasVettingId() bool`

HasVettingId returns a boolean if a field has been set.

### GetVettingToken

`func (o *ExternalVetting) GetVettingToken() string`

GetVettingToken returns the VettingToken field if non-nil, zero value otherwise.

### GetVettingTokenOk

`func (o *ExternalVetting) GetVettingTokenOk() (*string, bool)`

GetVettingTokenOk returns a tuple with the VettingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingToken

`func (o *ExternalVetting) SetVettingToken(v string)`

SetVettingToken sets VettingToken field to given value.

### HasVettingToken

`func (o *ExternalVetting) HasVettingToken() bool`

HasVettingToken returns a boolean if a field has been set.

### GetVettingScore

`func (o *ExternalVetting) GetVettingScore() int32`

GetVettingScore returns the VettingScore field if non-nil, zero value otherwise.

### GetVettingScoreOk

`func (o *ExternalVetting) GetVettingScoreOk() (*int32, bool)`

GetVettingScoreOk returns a tuple with the VettingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingScore

`func (o *ExternalVetting) SetVettingScore(v int32)`

SetVettingScore sets VettingScore field to given value.

### HasVettingScore

`func (o *ExternalVetting) HasVettingScore() bool`

HasVettingScore returns a boolean if a field has been set.

### GetVettingClass

`func (o *ExternalVetting) GetVettingClass() string`

GetVettingClass returns the VettingClass field if non-nil, zero value otherwise.

### GetVettingClassOk

`func (o *ExternalVetting) GetVettingClassOk() (*string, bool)`

GetVettingClassOk returns a tuple with the VettingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettingClass

`func (o *ExternalVetting) SetVettingClass(v string)`

SetVettingClass sets VettingClass field to given value.

### HasVettingClass

`func (o *ExternalVetting) HasVettingClass() bool`

HasVettingClass returns a boolean if a field has been set.

### GetVettedDate

`func (o *ExternalVetting) GetVettedDate() string`

GetVettedDate returns the VettedDate field if non-nil, zero value otherwise.

### GetVettedDateOk

`func (o *ExternalVetting) GetVettedDateOk() (*string, bool)`

GetVettedDateOk returns a tuple with the VettedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVettedDate

`func (o *ExternalVetting) SetVettedDate(v string)`

SetVettedDate sets VettedDate field to given value.

### HasVettedDate

`func (o *ExternalVetting) HasVettedDate() bool`

HasVettedDate returns a boolean if a field has been set.

### GetCreateDate

`func (o *ExternalVetting) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ExternalVetting) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ExternalVetting) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ExternalVetting) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


