# GenericError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**UpdateOutboundChannelsDefaultResponseErrorsInnerSource**](UpdateOutboundChannelsDefaultResponseErrorsInnerSource.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGenericError

`func NewGenericError() *GenericError`

NewGenericError instantiates a new GenericError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericErrorWithDefaults

`func NewGenericErrorWithDefaults() *GenericError`

NewGenericErrorWithDefaults instantiates a new GenericError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GenericError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenericError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenericError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GenericError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *GenericError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GenericError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GenericError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GenericError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *GenericError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GenericError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GenericError) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GenericError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *GenericError) GetSource() UpdateOutboundChannelsDefaultResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GenericError) GetSourceOk() (*UpdateOutboundChannelsDefaultResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GenericError) SetSource(v UpdateOutboundChannelsDefaultResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GenericError) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *GenericError) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GenericError) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GenericError) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GenericError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


