# DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TelnyxConversationChannel** | [**ConversationChannelType**](ConversationChannelType.md) |  | 
**TelnyxEndUserTarget** | **string** |  | 
**TelnyxAgentTarget** | **string** |  | 
**ScheduledAtFixedDatetime** | **time.Time** |  | 
**AssistantId** | **string** |  | 
**RetryCount** | Pointer to **int32** |  | [optional] [default to 0]
**RetryAttempts** | Pointer to **int32** |  | [optional] 
**ScheduledEventId** | Pointer to **string** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**EventStatus**](EventStatus.md) |  | [optional] [default to PENDING]
**ConversationMetadata** | Pointer to [**map[string]ConversationMetadataValue**](ConversationMetadataValue.md) |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewDataInner

`func NewDataInner(telnyxConversationChannel ConversationChannelType, telnyxEndUserTarget string, telnyxAgentTarget string, scheduledAtFixedDatetime time.Time, assistantId string, text string, ) *DataInner`

NewDataInner instantiates a new DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerWithDefaults

`func NewDataInnerWithDefaults() *DataInner`

NewDataInnerWithDefaults instantiates a new DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelnyxConversationChannel

`func (o *DataInner) GetTelnyxConversationChannel() ConversationChannelType`

GetTelnyxConversationChannel returns the TelnyxConversationChannel field if non-nil, zero value otherwise.

### GetTelnyxConversationChannelOk

`func (o *DataInner) GetTelnyxConversationChannelOk() (*ConversationChannelType, bool)`

GetTelnyxConversationChannelOk returns a tuple with the TelnyxConversationChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxConversationChannel

`func (o *DataInner) SetTelnyxConversationChannel(v ConversationChannelType)`

SetTelnyxConversationChannel sets TelnyxConversationChannel field to given value.


### GetTelnyxEndUserTarget

`func (o *DataInner) GetTelnyxEndUserTarget() string`

GetTelnyxEndUserTarget returns the TelnyxEndUserTarget field if non-nil, zero value otherwise.

### GetTelnyxEndUserTargetOk

`func (o *DataInner) GetTelnyxEndUserTargetOk() (*string, bool)`

GetTelnyxEndUserTargetOk returns a tuple with the TelnyxEndUserTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxEndUserTarget

`func (o *DataInner) SetTelnyxEndUserTarget(v string)`

SetTelnyxEndUserTarget sets TelnyxEndUserTarget field to given value.


### GetTelnyxAgentTarget

`func (o *DataInner) GetTelnyxAgentTarget() string`

GetTelnyxAgentTarget returns the TelnyxAgentTarget field if non-nil, zero value otherwise.

### GetTelnyxAgentTargetOk

`func (o *DataInner) GetTelnyxAgentTargetOk() (*string, bool)`

GetTelnyxAgentTargetOk returns a tuple with the TelnyxAgentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxAgentTarget

`func (o *DataInner) SetTelnyxAgentTarget(v string)`

SetTelnyxAgentTarget sets TelnyxAgentTarget field to given value.


### GetScheduledAtFixedDatetime

`func (o *DataInner) GetScheduledAtFixedDatetime() time.Time`

GetScheduledAtFixedDatetime returns the ScheduledAtFixedDatetime field if non-nil, zero value otherwise.

### GetScheduledAtFixedDatetimeOk

`func (o *DataInner) GetScheduledAtFixedDatetimeOk() (*time.Time, bool)`

GetScheduledAtFixedDatetimeOk returns a tuple with the ScheduledAtFixedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAtFixedDatetime

`func (o *DataInner) SetScheduledAtFixedDatetime(v time.Time)`

SetScheduledAtFixedDatetime sets ScheduledAtFixedDatetime field to given value.


### GetAssistantId

`func (o *DataInner) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *DataInner) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *DataInner) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.


### GetRetryCount

`func (o *DataInner) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *DataInner) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *DataInner) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *DataInner) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryAttempts

`func (o *DataInner) GetRetryAttempts() int32`

GetRetryAttempts returns the RetryAttempts field if non-nil, zero value otherwise.

### GetRetryAttemptsOk

`func (o *DataInner) GetRetryAttemptsOk() (*int32, bool)`

GetRetryAttemptsOk returns a tuple with the RetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttempts

`func (o *DataInner) SetRetryAttempts(v int32)`

SetRetryAttempts sets RetryAttempts field to given value.

### HasRetryAttempts

`func (o *DataInner) HasRetryAttempts() bool`

HasRetryAttempts returns a boolean if a field has been set.

### GetScheduledEventId

`func (o *DataInner) GetScheduledEventId() string`

GetScheduledEventId returns the ScheduledEventId field if non-nil, zero value otherwise.

### GetScheduledEventIdOk

`func (o *DataInner) GetScheduledEventIdOk() (*string, bool)`

GetScheduledEventIdOk returns a tuple with the ScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEventId

`func (o *DataInner) SetScheduledEventId(v string)`

SetScheduledEventId sets ScheduledEventId field to given value.

### HasScheduledEventId

`func (o *DataInner) HasScheduledEventId() bool`

HasScheduledEventId returns a boolean if a field has been set.

### GetConversationId

`func (o *DataInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *DataInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *DataInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *DataInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DataInner) GetStatus() EventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataInner) GetStatusOk() (*EventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataInner) SetStatus(v EventStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversationMetadata

`func (o *DataInner) GetConversationMetadata() map[string]ConversationMetadataValue`

GetConversationMetadata returns the ConversationMetadata field if non-nil, zero value otherwise.

### GetConversationMetadataOk

`func (o *DataInner) GetConversationMetadataOk() (*map[string]ConversationMetadataValue, bool)`

GetConversationMetadataOk returns a tuple with the ConversationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationMetadata

`func (o *DataInner) SetConversationMetadata(v map[string]ConversationMetadataValue)`

SetConversationMetadata sets ConversationMetadata field to given value.

### HasConversationMetadata

`func (o *DataInner) HasConversationMetadata() bool`

HasConversationMetadata returns a boolean if a field has been set.

### GetErrors

`func (o *DataInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DataInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DataInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DataInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetText

`func (o *DataInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DataInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DataInner) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


