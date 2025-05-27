# RCSMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **string** | The height of the media within a rich card with a vertical layout. For a standalone card with horizontal layout, height is not customizable, and this field is ignored. | [optional] 
**ContentInfo** | Pointer to [**RCSContentInfo**](RCSContentInfo.md) |  | [optional] 

## Methods

### NewRCSMedia

`func NewRCSMedia() *RCSMedia`

NewRCSMedia instantiates a new RCSMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSMediaWithDefaults

`func NewRCSMediaWithDefaults() *RCSMedia`

NewRCSMediaWithDefaults instantiates a new RCSMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *RCSMedia) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *RCSMedia) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *RCSMedia) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *RCSMedia) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetContentInfo

`func (o *RCSMedia) GetContentInfo() RCSContentInfo`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *RCSMedia) GetContentInfoOk() (*RCSContentInfo, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentInfo

`func (o *RCSMedia) SetContentInfo(v RCSContentInfo)`

SetContentInfo sets ContentInfo field to given value.

### HasContentInfo

`func (o *RCSMedia) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


