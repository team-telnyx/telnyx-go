# GetLogMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogMessages** | Pointer to [**[]LogMessage**](LogMessage.md) |  | [optional] 

## Methods

### NewGetLogMessageResponse

`func NewGetLogMessageResponse() *GetLogMessageResponse`

NewGetLogMessageResponse instantiates a new GetLogMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogMessageResponseWithDefaults

`func NewGetLogMessageResponseWithDefaults() *GetLogMessageResponse`

NewGetLogMessageResponseWithDefaults instantiates a new GetLogMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogMessages

`func (o *GetLogMessageResponse) GetLogMessages() []LogMessage`

GetLogMessages returns the LogMessages field if non-nil, zero value otherwise.

### GetLogMessagesOk

`func (o *GetLogMessageResponse) GetLogMessagesOk() (*[]LogMessage, bool)`

GetLogMessagesOk returns a tuple with the LogMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMessages

`func (o *GetLogMessageResponse) SetLogMessages(v []LogMessage)`

SetLogMessages sets LogMessages field to given value.

### HasLogMessages

`func (o *GetLogMessageResponse) HasLogMessages() bool`

HasLogMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


