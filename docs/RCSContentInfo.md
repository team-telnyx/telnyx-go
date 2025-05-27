# RCSContentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileUrl** | **string** | Publicly reachable URL of the file. | 
**ThumbnailUrl** | Pointer to **string** | Publicly reachable URL of the thumbnail. Maximum size of 100 kB. | [optional] 
**ForceRefresh** | Pointer to **bool** | If set the URL content will not be cached. | [optional] 

## Methods

### NewRCSContentInfo

`func NewRCSContentInfo(fileUrl string, ) *RCSContentInfo`

NewRCSContentInfo instantiates a new RCSContentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSContentInfoWithDefaults

`func NewRCSContentInfoWithDefaults() *RCSContentInfo`

NewRCSContentInfoWithDefaults instantiates a new RCSContentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileUrl

`func (o *RCSContentInfo) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *RCSContentInfo) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *RCSContentInfo) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetThumbnailUrl

`func (o *RCSContentInfo) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *RCSContentInfo) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *RCSContentInfo) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *RCSContentInfo) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetForceRefresh

`func (o *RCSContentInfo) GetForceRefresh() bool`

GetForceRefresh returns the ForceRefresh field if non-nil, zero value otherwise.

### GetForceRefreshOk

`func (o *RCSContentInfo) GetForceRefreshOk() (*bool, bool)`

GetForceRefreshOk returns a tuple with the ForceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRefresh

`func (o *RCSContentInfo) SetForceRefresh(v bool)`

SetForceRefresh sets ForceRefresh field to given value.

### HasForceRefresh

`func (o *RCSContentInfo) HasForceRefresh() bool`

HasForceRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


