# CreateUploadRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Describes wether or not the operation was successful | [optional] 
**TicketId** | Pointer to **string** | Ticket id of the upload request | [optional] 

## Methods

### NewCreateUploadRequestResponse

`func NewCreateUploadRequestResponse() *CreateUploadRequestResponse`

NewCreateUploadRequestResponse instantiates a new CreateUploadRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUploadRequestResponseWithDefaults

`func NewCreateUploadRequestResponseWithDefaults() *CreateUploadRequestResponse`

NewCreateUploadRequestResponseWithDefaults instantiates a new CreateUploadRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateUploadRequestResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateUploadRequestResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateUploadRequestResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateUploadRequestResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTicketId

`func (o *CreateUploadRequestResponse) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *CreateUploadRequestResponse) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *CreateUploadRequestResponse) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *CreateUploadRequestResponse) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


