# PortingOrdersComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** | Body of comment | [optional] 
**PortingOrderId** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** | Indicates whether this comment was created by a Telnyx Admin, user, or system | [optional] 
**UserId** | Pointer to **string** | The ID of the user who created this comment | [optional] 
**UserEmail** | Pointer to **string** | The email address of the user who created this comment | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 

## Methods

### NewPortingOrdersComment

`func NewPortingOrdersComment() *PortingOrdersComment`

NewPortingOrdersComment instantiates a new PortingOrdersComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrdersCommentWithDefaults

`func NewPortingOrdersCommentWithDefaults() *PortingOrdersComment`

NewPortingOrdersCommentWithDefaults instantiates a new PortingOrdersComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingOrdersComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingOrdersComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingOrdersComment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingOrdersComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *PortingOrdersComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PortingOrdersComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PortingOrdersComment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PortingOrdersComment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingOrdersComment) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingOrdersComment) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingOrdersComment) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingOrdersComment) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetUserType

`func (o *PortingOrdersComment) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *PortingOrdersComment) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *PortingOrdersComment) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *PortingOrdersComment) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUserId

`func (o *PortingOrdersComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PortingOrdersComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PortingOrdersComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PortingOrdersComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserEmail

`func (o *PortingOrdersComment) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *PortingOrdersComment) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *PortingOrdersComment) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *PortingOrdersComment) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrdersComment) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrdersComment) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrdersComment) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrdersComment) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingOrdersComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingOrdersComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingOrdersComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingOrdersComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


