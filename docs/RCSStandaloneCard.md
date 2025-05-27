# RCSStandaloneCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardOrientation** | **string** | Orientation of the card. | 
**ThumbnailImageAlignment** | **string** | Image preview alignment for standalone cards with horizontal layout. | 
**CardContent** | [**RCSCardContent**](RCSCardContent.md) |  | 

## Methods

### NewRCSStandaloneCard

`func NewRCSStandaloneCard(cardOrientation string, thumbnailImageAlignment string, cardContent RCSCardContent, ) *RCSStandaloneCard`

NewRCSStandaloneCard instantiates a new RCSStandaloneCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSStandaloneCardWithDefaults

`func NewRCSStandaloneCardWithDefaults() *RCSStandaloneCard`

NewRCSStandaloneCardWithDefaults instantiates a new RCSStandaloneCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardOrientation

`func (o *RCSStandaloneCard) GetCardOrientation() string`

GetCardOrientation returns the CardOrientation field if non-nil, zero value otherwise.

### GetCardOrientationOk

`func (o *RCSStandaloneCard) GetCardOrientationOk() (*string, bool)`

GetCardOrientationOk returns a tuple with the CardOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardOrientation

`func (o *RCSStandaloneCard) SetCardOrientation(v string)`

SetCardOrientation sets CardOrientation field to given value.


### GetThumbnailImageAlignment

`func (o *RCSStandaloneCard) GetThumbnailImageAlignment() string`

GetThumbnailImageAlignment returns the ThumbnailImageAlignment field if non-nil, zero value otherwise.

### GetThumbnailImageAlignmentOk

`func (o *RCSStandaloneCard) GetThumbnailImageAlignmentOk() (*string, bool)`

GetThumbnailImageAlignmentOk returns a tuple with the ThumbnailImageAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageAlignment

`func (o *RCSStandaloneCard) SetThumbnailImageAlignment(v string)`

SetThumbnailImageAlignment sets ThumbnailImageAlignment field to given value.


### GetCardContent

`func (o *RCSStandaloneCard) GetCardContent() RCSCardContent`

GetCardContent returns the CardContent field if non-nil, zero value otherwise.

### GetCardContentOk

`func (o *RCSStandaloneCard) GetCardContentOk() (*RCSCardContent, bool)`

GetCardContentOk returns a tuple with the CardContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardContent

`func (o *RCSStandaloneCard) SetCardContent(v RCSCardContent)`

SetCardContent sets CardContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


