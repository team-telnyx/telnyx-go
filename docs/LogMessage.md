# LogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Title** | **string** |  | 
**Detail** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**LogMessageSource**](LogMessageSource.md) |  | [optional] 
**Meta** | Pointer to [**LogMessageMeta**](LogMessageMeta.md) |  | [optional] 

## Methods

### NewLogMessage

`func NewLogMessage(code int32, title string, ) *LogMessage`

NewLogMessage instantiates a new LogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMessageWithDefaults

`func NewLogMessageWithDefaults() *LogMessage`

NewLogMessageWithDefaults instantiates a new LogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LogMessage) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LogMessage) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LogMessage) SetCode(v int32)`

SetCode sets Code field to given value.


### GetTitle

`func (o *LogMessage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogMessage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogMessage) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *LogMessage) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LogMessage) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LogMessage) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LogMessage) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *LogMessage) GetSource() LogMessageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LogMessage) GetSourceOk() (*LogMessageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LogMessage) SetSource(v LogMessageSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *LogMessage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetMeta

`func (o *LogMessage) GetMeta() LogMessageMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LogMessage) GetMetaOk() (*LogMessageMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LogMessage) SetMeta(v LogMessageMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LogMessage) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


