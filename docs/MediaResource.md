# MediaResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaName** | Pointer to **string** | Uniquely identifies a media resource. | [optional] 
**ExpiresAt** | Pointer to **string** | ISO 8601 formatted date of when the media resource will expire and be deleted. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the media resource was created | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date of when the media resource was last updated | [optional] 
**ContentType** | Pointer to **string** | Content type of the file | [optional] 

## Methods

### NewMediaResource

`func NewMediaResource() *MediaResource`

NewMediaResource instantiates a new MediaResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaResourceWithDefaults

`func NewMediaResourceWithDefaults() *MediaResource`

NewMediaResourceWithDefaults instantiates a new MediaResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaName

`func (o *MediaResource) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *MediaResource) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *MediaResource) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *MediaResource) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *MediaResource) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MediaResource) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MediaResource) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *MediaResource) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaResource) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaResource) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaResource) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MediaResource) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaResource) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaResource) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MediaResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetContentType

`func (o *MediaResource) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MediaResource) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MediaResource) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MediaResource) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


