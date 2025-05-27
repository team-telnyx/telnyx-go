# CursorPaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursors** | Pointer to [**Cursor**](Cursor.md) |  | [optional] 
**Next** | Pointer to **string** | Path to next page. | [optional] 
**Previous** | Pointer to **string** | Path to previous page. | [optional] 

## Methods

### NewCursorPaginationMeta

`func NewCursorPaginationMeta() *CursorPaginationMeta`

NewCursorPaginationMeta instantiates a new CursorPaginationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPaginationMetaWithDefaults

`func NewCursorPaginationMetaWithDefaults() *CursorPaginationMeta`

NewCursorPaginationMetaWithDefaults instantiates a new CursorPaginationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursors

`func (o *CursorPaginationMeta) GetCursors() Cursor`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *CursorPaginationMeta) GetCursorsOk() (*Cursor, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *CursorPaginationMeta) SetCursors(v Cursor)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *CursorPaginationMeta) HasCursors() bool`

HasCursors returns a boolean if a field has been set.

### GetNext

`func (o *CursorPaginationMeta) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CursorPaginationMeta) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CursorPaginationMeta) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CursorPaginationMeta) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *CursorPaginationMeta) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *CursorPaginationMeta) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *CursorPaginationMeta) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *CursorPaginationMeta) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


