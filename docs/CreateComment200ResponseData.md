# CreateComment200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Body** | Pointer to **string** |  | [optional] 
**Commenter** | Pointer to **string** |  | [optional] [readonly] 
**CommenterType** | Pointer to **string** |  | [optional] [readonly] 
**CommentRecordType** | Pointer to **string** |  | [optional] 
**CommentRecordId** | Pointer to **string** |  | [optional] 
**ReadAt** | Pointer to **string** | An ISO 8901 datetime string for when the comment was read. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the comment was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the comment was updated. | [optional] [readonly] 

## Methods

### NewCreateComment200ResponseData

`func NewCreateComment200ResponseData() *CreateComment200ResponseData`

NewCreateComment200ResponseData instantiates a new CreateComment200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComment200ResponseDataWithDefaults

`func NewCreateComment200ResponseDataWithDefaults() *CreateComment200ResponseData`

NewCreateComment200ResponseDataWithDefaults instantiates a new CreateComment200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateComment200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateComment200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateComment200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateComment200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *CreateComment200ResponseData) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateComment200ResponseData) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateComment200ResponseData) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateComment200ResponseData) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCommenter

`func (o *CreateComment200ResponseData) GetCommenter() string`

GetCommenter returns the Commenter field if non-nil, zero value otherwise.

### GetCommenterOk

`func (o *CreateComment200ResponseData) GetCommenterOk() (*string, bool)`

GetCommenterOk returns a tuple with the Commenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenter

`func (o *CreateComment200ResponseData) SetCommenter(v string)`

SetCommenter sets Commenter field to given value.

### HasCommenter

`func (o *CreateComment200ResponseData) HasCommenter() bool`

HasCommenter returns a boolean if a field has been set.

### GetCommenterType

`func (o *CreateComment200ResponseData) GetCommenterType() string`

GetCommenterType returns the CommenterType field if non-nil, zero value otherwise.

### GetCommenterTypeOk

`func (o *CreateComment200ResponseData) GetCommenterTypeOk() (*string, bool)`

GetCommenterTypeOk returns a tuple with the CommenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterType

`func (o *CreateComment200ResponseData) SetCommenterType(v string)`

SetCommenterType sets CommenterType field to given value.

### HasCommenterType

`func (o *CreateComment200ResponseData) HasCommenterType() bool`

HasCommenterType returns a boolean if a field has been set.

### GetCommentRecordType

`func (o *CreateComment200ResponseData) GetCommentRecordType() string`

GetCommentRecordType returns the CommentRecordType field if non-nil, zero value otherwise.

### GetCommentRecordTypeOk

`func (o *CreateComment200ResponseData) GetCommentRecordTypeOk() (*string, bool)`

GetCommentRecordTypeOk returns a tuple with the CommentRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentRecordType

`func (o *CreateComment200ResponseData) SetCommentRecordType(v string)`

SetCommentRecordType sets CommentRecordType field to given value.

### HasCommentRecordType

`func (o *CreateComment200ResponseData) HasCommentRecordType() bool`

HasCommentRecordType returns a boolean if a field has been set.

### GetCommentRecordId

`func (o *CreateComment200ResponseData) GetCommentRecordId() string`

GetCommentRecordId returns the CommentRecordId field if non-nil, zero value otherwise.

### GetCommentRecordIdOk

`func (o *CreateComment200ResponseData) GetCommentRecordIdOk() (*string, bool)`

GetCommentRecordIdOk returns a tuple with the CommentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentRecordId

`func (o *CreateComment200ResponseData) SetCommentRecordId(v string)`

SetCommentRecordId sets CommentRecordId field to given value.

### HasCommentRecordId

`func (o *CreateComment200ResponseData) HasCommentRecordId() bool`

HasCommentRecordId returns a boolean if a field has been set.

### GetReadAt

`func (o *CreateComment200ResponseData) GetReadAt() string`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *CreateComment200ResponseData) GetReadAtOk() (*string, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *CreateComment200ResponseData) SetReadAt(v string)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *CreateComment200ResponseData) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateComment200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateComment200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateComment200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateComment200ResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateComment200ResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateComment200ResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateComment200ResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateComment200ResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


