# UpdateMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaUrl** | Pointer to **string** | The URL where the media to be stored in Telnyx network is currently hosted. The maximum allowed size is 20 MB. | [optional] 
**TtlSecs** | Pointer to **int32** | The number of seconds after which the media resource will be deleted, defaults to 2 days. The maximum allowed vale is 630720000, which translates to 20 years. | [optional] 

## Methods

### NewUpdateMediaRequest

`func NewUpdateMediaRequest() *UpdateMediaRequest`

NewUpdateMediaRequest instantiates a new UpdateMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMediaRequestWithDefaults

`func NewUpdateMediaRequestWithDefaults() *UpdateMediaRequest`

NewUpdateMediaRequestWithDefaults instantiates a new UpdateMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaUrl

`func (o *UpdateMediaRequest) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *UpdateMediaRequest) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *UpdateMediaRequest) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *UpdateMediaRequest) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetTtlSecs

`func (o *UpdateMediaRequest) GetTtlSecs() int32`

GetTtlSecs returns the TtlSecs field if non-nil, zero value otherwise.

### GetTtlSecsOk

`func (o *UpdateMediaRequest) GetTtlSecsOk() (*int32, bool)`

GetTtlSecsOk returns a tuple with the TtlSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecs

`func (o *UpdateMediaRequest) SetTtlSecs(v int32)`

SetTtlSecs sets TtlSecs field to given value.

### HasTtlSecs

`func (o *UpdateMediaRequest) HasTtlSecs() bool`

HasTtlSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


