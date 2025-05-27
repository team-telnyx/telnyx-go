# PortoutComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Body** | **string** | Comment body | 
**PortoutId** | Pointer to **string** | Identifies the associated port request | [optional] 
**UserId** | **string** | Identifies the user who created the comment. Will be null if created by Telnyx Admin | 
**CreatedAt** | **string** | Comment creation timestamp in ISO 8601 format | 

## Methods

### NewPortoutComment

`func NewPortoutComment(id string, body string, userId string, createdAt string, ) *PortoutComment`

NewPortoutComment instantiates a new PortoutComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortoutCommentWithDefaults

`func NewPortoutCommentWithDefaults() *PortoutComment`

NewPortoutCommentWithDefaults instantiates a new PortoutComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortoutComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortoutComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortoutComment) SetId(v string)`

SetId sets Id field to given value.


### GetRecordType

`func (o *PortoutComment) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortoutComment) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortoutComment) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortoutComment) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetBody

`func (o *PortoutComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PortoutComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PortoutComment) SetBody(v string)`

SetBody sets Body field to given value.


### GetPortoutId

`func (o *PortoutComment) GetPortoutId() string`

GetPortoutId returns the PortoutId field if non-nil, zero value otherwise.

### GetPortoutIdOk

`func (o *PortoutComment) GetPortoutIdOk() (*string, bool)`

GetPortoutIdOk returns a tuple with the PortoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortoutId

`func (o *PortoutComment) SetPortoutId(v string)`

SetPortoutId sets PortoutId field to given value.

### HasPortoutId

`func (o *PortoutComment) HasPortoutId() bool`

HasPortoutId returns a boolean if a field has been set.

### GetUserId

`func (o *PortoutComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PortoutComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PortoutComment) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *PortoutComment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortoutComment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortoutComment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


