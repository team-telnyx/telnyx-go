# Comment

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

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *Comment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Comment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Comment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Comment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCommenter

`func (o *Comment) GetCommenter() string`

GetCommenter returns the Commenter field if non-nil, zero value otherwise.

### GetCommenterOk

`func (o *Comment) GetCommenterOk() (*string, bool)`

GetCommenterOk returns a tuple with the Commenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenter

`func (o *Comment) SetCommenter(v string)`

SetCommenter sets Commenter field to given value.

### HasCommenter

`func (o *Comment) HasCommenter() bool`

HasCommenter returns a boolean if a field has been set.

### GetCommenterType

`func (o *Comment) GetCommenterType() string`

GetCommenterType returns the CommenterType field if non-nil, zero value otherwise.

### GetCommenterTypeOk

`func (o *Comment) GetCommenterTypeOk() (*string, bool)`

GetCommenterTypeOk returns a tuple with the CommenterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterType

`func (o *Comment) SetCommenterType(v string)`

SetCommenterType sets CommenterType field to given value.

### HasCommenterType

`func (o *Comment) HasCommenterType() bool`

HasCommenterType returns a boolean if a field has been set.

### GetCommentRecordType

`func (o *Comment) GetCommentRecordType() string`

GetCommentRecordType returns the CommentRecordType field if non-nil, zero value otherwise.

### GetCommentRecordTypeOk

`func (o *Comment) GetCommentRecordTypeOk() (*string, bool)`

GetCommentRecordTypeOk returns a tuple with the CommentRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentRecordType

`func (o *Comment) SetCommentRecordType(v string)`

SetCommentRecordType sets CommentRecordType field to given value.

### HasCommentRecordType

`func (o *Comment) HasCommentRecordType() bool`

HasCommentRecordType returns a boolean if a field has been set.

### GetCommentRecordId

`func (o *Comment) GetCommentRecordId() string`

GetCommentRecordId returns the CommentRecordId field if non-nil, zero value otherwise.

### GetCommentRecordIdOk

`func (o *Comment) GetCommentRecordIdOk() (*string, bool)`

GetCommentRecordIdOk returns a tuple with the CommentRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentRecordId

`func (o *Comment) SetCommentRecordId(v string)`

SetCommentRecordId sets CommentRecordId field to given value.

### HasCommentRecordId

`func (o *Comment) HasCommentRecordId() bool`

HasCommentRecordId returns a boolean if a field has been set.

### GetReadAt

`func (o *Comment) GetReadAt() string`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *Comment) GetReadAtOk() (*string, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *Comment) SetReadAt(v string)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *Comment) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Comment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Comment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Comment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Comment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Comment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Comment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


