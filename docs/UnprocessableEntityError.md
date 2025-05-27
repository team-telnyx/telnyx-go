# UnprocessableEntityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Title** | **string** |  | 
**Detail** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**UpdateOutboundChannelsDefaultResponseErrorsInnerSource**](UpdateOutboundChannelsDefaultResponseErrorsInnerSource.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUnprocessableEntityError

`func NewUnprocessableEntityError(code string, title string, ) *UnprocessableEntityError`

NewUnprocessableEntityError instantiates a new UnprocessableEntityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprocessableEntityErrorWithDefaults

`func NewUnprocessableEntityErrorWithDefaults() *UnprocessableEntityError`

NewUnprocessableEntityErrorWithDefaults instantiates a new UnprocessableEntityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnprocessableEntityError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnprocessableEntityError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnprocessableEntityError) SetCode(v string)`

SetCode sets Code field to given value.


### GetTitle

`func (o *UnprocessableEntityError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UnprocessableEntityError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UnprocessableEntityError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *UnprocessableEntityError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UnprocessableEntityError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UnprocessableEntityError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *UnprocessableEntityError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *UnprocessableEntityError) GetSource() UpdateOutboundChannelsDefaultResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UnprocessableEntityError) GetSourceOk() (*UpdateOutboundChannelsDefaultResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UnprocessableEntityError) SetSource(v UpdateOutboundChannelsDefaultResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *UnprocessableEntityError) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *UnprocessableEntityError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnprocessableEntityError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnprocessableEntityError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnprocessableEntityError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


