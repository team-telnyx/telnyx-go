# ListLogMessagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogMessages** | Pointer to [**[]LogMessage**](LogMessage.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListLogMessagesResponse

`func NewListLogMessagesResponse() *ListLogMessagesResponse`

NewListLogMessagesResponse instantiates a new ListLogMessagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogMessagesResponseWithDefaults

`func NewListLogMessagesResponseWithDefaults() *ListLogMessagesResponse`

NewListLogMessagesResponseWithDefaults instantiates a new ListLogMessagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogMessages

`func (o *ListLogMessagesResponse) GetLogMessages() []LogMessage`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *ListLogMessagesResponse) GetLogMessagesOk() (*[]LogMessage, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *ListLogMessagesResponse) SetLogMessages(v []LogMessage)`

SetLogMessages sets LogMessages field to given value.

### HasLogMessages

`func (o *ListLogMessagesResponse) HasLogMessages() bool`

HasLogMessages returns a boolean if a field has been set.

### GetMeta

`func (o *ListLogMessagesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLogMessagesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLogMessagesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLogMessagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


