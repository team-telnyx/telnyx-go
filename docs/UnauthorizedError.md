# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **interface{}** |  | [optional] 
**Detail** | Pointer to **interface{}** |  | [optional] 
**Source** | Pointer to [**UpdateOutboundChannelsDefaultResponseErrorsInnerSource**](UpdateOutboundChannelsDefaultResponseErrorsInnerSource.md) |  | [optional] 
**Meta** | Pointer to [**UnauthorizedErrorAllOfMeta**](UnauthorizedErrorAllOfMeta.md) |  | [optional] 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError() *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnauthorizedError) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedError) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedError) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *UnauthorizedError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UnauthorizedError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UnauthorizedError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTitle

`func (o *UnauthorizedError) GetTitle() interface{}`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UnauthorizedError) GetTitleOk() (*interface{}, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UnauthorizedError) SetTitle(v interface{})`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UnauthorizedError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UnauthorizedError) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UnauthorizedError) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDetail

`func (o *UnauthorizedError) GetDetail() interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UnauthorizedError) GetDetailOk() (*interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UnauthorizedError) SetDetail(v interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *UnauthorizedError) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *UnauthorizedError) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *UnauthorizedError) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetSource

`func (o *UnauthorizedError) GetSource() UpdateOutboundChannelsDefaultResponseErrorsInnerSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UnauthorizedError) GetSourceOk() (*UpdateOutboundChannelsDefaultResponseErrorsInnerSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UnauthorizedError) SetSource(v UpdateOutboundChannelsDefaultResponseErrorsInnerSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *UnauthorizedError) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *UnauthorizedError) GetMeta() UnauthorizedErrorAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnauthorizedError) GetMetaOk() (*UnauthorizedErrorAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnauthorizedError) SetMeta(v UnauthorizedErrorAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnauthorizedError) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


