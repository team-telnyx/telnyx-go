# PresignedObjectUrlContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | The token for the object | [optional] 
**PresignedUrl** | Pointer to **string** | The presigned URL for the object | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The expiration time of the token | [optional] 

## Methods

### NewPresignedObjectUrlContent

`func NewPresignedObjectUrlContent() *PresignedObjectUrlContent`

NewPresignedObjectUrlContent instantiates a new PresignedObjectUrlContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignedObjectUrlContentWithDefaults

`func NewPresignedObjectUrlContentWithDefaults() *PresignedObjectUrlContent`

NewPresignedObjectUrlContentWithDefaults instantiates a new PresignedObjectUrlContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PresignedObjectUrlContent) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PresignedObjectUrlContent) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PresignedObjectUrlContent) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PresignedObjectUrlContent) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetPresignedUrl

`func (o *PresignedObjectUrlContent) GetPresignedUrl() string`

GetPresignedUrl returns the PresignedUrl field if non-nil, zero value otherwise.

### GetPresignedUrlOk

`func (o *PresignedObjectUrlContent) GetPresignedUrlOk() (*string, bool)`

GetPresignedUrlOk returns a tuple with the PresignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresignedUrl

`func (o *PresignedObjectUrlContent) SetPresignedUrl(v string)`

SetPresignedUrl sets PresignedUrl field to given value.

### HasPresignedUrl

`func (o *PresignedObjectUrlContent) HasPresignedUrl() bool`

HasPresignedUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PresignedObjectUrlContent) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PresignedObjectUrlContent) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PresignedObjectUrlContent) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PresignedObjectUrlContent) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


