# UploadMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaUrl** | **string** | The URL where the media to be stored in Telnyx network is currently hosted. The maximum allowed size is 20 MB. | 
**TtlSecs** | Pointer to **int32** | The number of seconds after which the media resource will be deleted, defaults to 2 days. The maximum allowed vale is 630720000, which translates to 20 years. | [optional] 
**MediaName** | Pointer to **string** | The unique identifier of a file. | [optional] 

## Methods

### NewUploadMediaRequest

`func NewUploadMediaRequest(mediaUrl string, ) *UploadMediaRequest`

NewUploadMediaRequest instantiates a new UploadMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadMediaRequestWithDefaults

`func NewUploadMediaRequestWithDefaults() *UploadMediaRequest`

NewUploadMediaRequestWithDefaults instantiates a new UploadMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaUrl

`func (o *UploadMediaRequest) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *UploadMediaRequest) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *UploadMediaRequest) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.


### GetTtlSecs

`func (o *UploadMediaRequest) GetTtlSecs() int32`

GetTtlSecs returns the TtlSecs field if non-nil, zero value otherwise.

### GetTtlSecsOk

`func (o *UploadMediaRequest) GetTtlSecsOk() (*int32, bool)`

GetTtlSecsOk returns a tuple with the TtlSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecs

`func (o *UploadMediaRequest) SetTtlSecs(v int32)`

SetTtlSecs sets TtlSecs field to given value.

### HasTtlSecs

`func (o *UploadMediaRequest) HasTtlSecs() bool`

HasTtlSecs returns a boolean if a field has been set.

### GetMediaName

`func (o *UploadMediaRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *UploadMediaRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *UploadMediaRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *UploadMediaRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


