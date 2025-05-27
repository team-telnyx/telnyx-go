# ListObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to [**[]ListObjectsResponseContentsInner**](ListObjectsResponseContentsInner.md) |  | [optional] 

## Methods

### NewListObjectsResponse

`func NewListObjectsResponse() *ListObjectsResponse`

NewListObjectsResponse instantiates a new ListObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListObjectsResponseWithDefaults

`func NewListObjectsResponseWithDefaults() *ListObjectsResponse`

NewListObjectsResponseWithDefaults instantiates a new ListObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListObjectsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListObjectsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListObjectsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListObjectsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContents

`func (o *ListObjectsResponse) GetContents() []ListObjectsResponseContentsInner`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ListObjectsResponse) GetContentsOk() (*[]ListObjectsResponseContentsInner, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ListObjectsResponse) SetContents(v []ListObjectsResponseContentsInner)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ListObjectsResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


