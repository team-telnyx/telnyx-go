# BrandFeedbackCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | One of &#x60;TAX_ID&#x60;, &#x60;STOCK_SYMBOL&#x60;, &#x60;GOVERNMENT_ENTITY&#x60;, &#x60;NONPROFIT&#x60;, and &#x60;OTHERS&#x60; | 
**DisplayName** | **string** | Human-readable version of the &#x60;id&#x60; field | 
**Description** | **string** | Long-form description of the feedback with additional information | 
**Fields** | **[]string** | List of relevant fields in the originally-submitted brand json | 

## Methods

### NewBrandFeedbackCategory

`func NewBrandFeedbackCategory(id string, displayName string, description string, fields []string, ) *BrandFeedbackCategory`

NewBrandFeedbackCategory instantiates a new BrandFeedbackCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandFeedbackCategoryWithDefaults

`func NewBrandFeedbackCategoryWithDefaults() *BrandFeedbackCategory`

NewBrandFeedbackCategoryWithDefaults instantiates a new BrandFeedbackCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BrandFeedbackCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandFeedbackCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandFeedbackCategory) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *BrandFeedbackCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BrandFeedbackCategory) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BrandFeedbackCategory) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *BrandFeedbackCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BrandFeedbackCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BrandFeedbackCategory) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFields

`func (o *BrandFeedbackCategory) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BrandFeedbackCategory) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BrandFeedbackCategory) SetFields(v []string)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


