# BrandFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | **string** | ID of the brand being queried about | 
**Category** | [**[]BrandFeedbackCategory**](BrandFeedbackCategory.md) | A list of reasons why brand creation/revetting didn&#39;t go as planned | 

## Methods

### NewBrandFeedback

`func NewBrandFeedback(brandId string, category []BrandFeedbackCategory, ) *BrandFeedback`

NewBrandFeedback instantiates a new BrandFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandFeedbackWithDefaults

`func NewBrandFeedbackWithDefaults() *BrandFeedback`

NewBrandFeedbackWithDefaults instantiates a new BrandFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *BrandFeedback) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *BrandFeedback) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *BrandFeedback) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.


### GetCategory

`func (o *BrandFeedback) GetCategory() []BrandFeedbackCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BrandFeedback) GetCategoryOk() (*[]BrandFeedbackCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BrandFeedback) SetCategory(v []BrandFeedbackCategory)`

SetCategory sets Category field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


