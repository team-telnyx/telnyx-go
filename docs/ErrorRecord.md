# ErrorRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | an application-specific error code, expressed as a stringified 5-digit integer | [optional] 
**Title** | Pointer to **string** | a short, human-readable summary of the problem, with NO punctuation, that SHOULD NOT CHANGE from occurrence to occurrence of the problem, except for purposes of localization | [optional] 
**Detail** | Pointer to **string** | a human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized | [optional] 
**Source** | Pointer to [**SourceResponse**](SourceResponse.md) |  | [optional] 
**Meta** | Pointer to [**MetaResponse**](MetaResponse.md) |  | [optional] 

## Methods

### NewErrorRecord

`func NewErrorRecord() *ErrorRecord`

NewErrorRecord instantiates a new ErrorRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorRecordWithDefaults

`func NewErrorRecordWithDefaults() *ErrorRecord`

NewErrorRecordWithDefaults instantiates a new ErrorRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorRecord) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorRecord) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorRecord) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorRecord) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *ErrorRecord) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorRecord) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorRecord) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorRecord) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorRecord) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorRecord) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorRecord) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorRecord) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *ErrorRecord) GetSource() SourceResponse`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ErrorRecord) GetSourceOk() (*SourceResponse, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ErrorRecord) SetSource(v SourceResponse)`

SetSource sets Source field to given value.

### HasSource

`func (o *ErrorRecord) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorRecord) GetMeta() MetaResponse`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorRecord) GetMetaOk() (*MetaResponse, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorRecord) SetMeta(v MetaResponse)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


