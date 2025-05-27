# LogMessageMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TelephoneNumber** | Pointer to **string** | The telephone number the log message is associated with, if any. | [optional] 
**ExternalConnectionId** | Pointer to **string** | The external connection the log message is associated with, if any. | [optional] 
**TicketId** | Pointer to **string** | The ticket ID for an operation that generated the log message, if any. | [optional] 

## Methods

### NewLogMessageMeta

`func NewLogMessageMeta() *LogMessageMeta`

NewLogMessageMeta instantiates a new LogMessageMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogMessageMetaWithDefaults

`func NewLogMessageMetaWithDefaults() *LogMessageMeta`

NewLogMessageMetaWithDefaults instantiates a new LogMessageMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelephoneNumber

`func (o *LogMessageMeta) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *LogMessageMeta) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *LogMessageMeta) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *LogMessageMeta) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetExternalConnectionId

`func (o *LogMessageMeta) GetExternalConnectionId() string`

GetExternalConnectionId returns the ExternalConnectionId field if non-nil, zero value otherwise.

### GetExternalConnectionIdOk

`func (o *LogMessageMeta) GetExternalConnectionIdOk() (*string, bool)`

GetExternalConnectionIdOk returns a tuple with the ExternalConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnectionId

`func (o *LogMessageMeta) SetExternalConnectionId(v string)`

SetExternalConnectionId sets ExternalConnectionId field to given value.

### HasExternalConnectionId

`func (o *LogMessageMeta) HasExternalConnectionId() bool`

HasExternalConnectionId returns a boolean if a field has been set.

### GetTicketId

`func (o *LogMessageMeta) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *LogMessageMeta) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *LogMessageMeta) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *LogMessageMeta) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


