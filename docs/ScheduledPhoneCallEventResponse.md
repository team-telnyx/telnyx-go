# ScheduledPhoneCallEventResponse

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

## Methods

### NewScheduledPhoneCallEventResponse

`func NewScheduledPhoneCallEventResponse(telnyxConversationChannel ConversationChannelType, telnyxEndUserTarget string, telnyxAgentTarget string, scheduledAtFixedDatetime time.Time, assistantId string, ) *ScheduledPhoneCallEventResponse`

NewScheduledPhoneCallEventResponse instantiates a new ScheduledPhoneCallEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledPhoneCallEventResponseWithDefaults

`func NewScheduledPhoneCallEventResponseWithDefaults() *ScheduledPhoneCallEventResponse`

NewScheduledPhoneCallEventResponseWithDefaults instantiates a new ScheduledPhoneCallEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelnyxConversationChannel

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxConversationChannel() ConversationChannelType`

GetTelnyxConversationChannel returns the TelnyxConversationChannel field if non-nil, zero value otherwise.

### GetTelnyxConversationChannelOk

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxConversationChannelOk() (*ConversationChannelType, bool)`

GetTelnyxConversationChannelOk returns a tuple with the TelnyxConversationChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxConversationChannel

`func (o *ScheduledPhoneCallEventResponse) SetTelnyxConversationChannel(v ConversationChannelType)`

SetTelnyxConversationChannel sets TelnyxConversationChannel field to given value.


### GetTelnyxEndUserTarget

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxEndUserTarget() string`

GetTelnyxEndUserTarget returns the TelnyxEndUserTarget field if non-nil, zero value otherwise.

### GetTelnyxEndUserTargetOk

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxEndUserTargetOk() (*string, bool)`

GetTelnyxEndUserTargetOk returns a tuple with the TelnyxEndUserTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxEndUserTarget

`func (o *ScheduledPhoneCallEventResponse) SetTelnyxEndUserTarget(v string)`

SetTelnyxEndUserTarget sets TelnyxEndUserTarget field to given value.


### GetTelnyxAgentTarget

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxAgentTarget() string`

GetTelnyxAgentTarget returns the TelnyxAgentTarget field if non-nil, zero value otherwise.

### GetTelnyxAgentTargetOk

`func (o *ScheduledPhoneCallEventResponse) GetTelnyxAgentTargetOk() (*string, bool)`

GetTelnyxAgentTargetOk returns a tuple with the TelnyxAgentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxAgentTarget

`func (o *ScheduledPhoneCallEventResponse) SetTelnyxAgentTarget(v string)`

SetTelnyxAgentTarget sets TelnyxAgentTarget field to given value.


### GetScheduledAtFixedDatetime

`func (o *ScheduledPhoneCallEventResponse) GetScheduledAtFixedDatetime() time.Time`

GetScheduledAtFixedDatetime returns the ScheduledAtFixedDatetime field if non-nil, zero value otherwise.

### GetScheduledAtFixedDatetimeOk

`func (o *ScheduledPhoneCallEventResponse) GetScheduledAtFixedDatetimeOk() (*time.Time, bool)`

GetScheduledAtFixedDatetimeOk returns a tuple with the ScheduledAtFixedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAtFixedDatetime

`func (o *ScheduledPhoneCallEventResponse) SetScheduledAtFixedDatetime(v time.Time)`

SetScheduledAtFixedDatetime sets ScheduledAtFixedDatetime field to given value.


### GetAssistantId

`func (o *ScheduledPhoneCallEventResponse) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *ScheduledPhoneCallEventResponse) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *ScheduledPhoneCallEventResponse) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.


### GetRetryCount

`func (o *ScheduledPhoneCallEventResponse) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ScheduledPhoneCallEventResponse) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ScheduledPhoneCallEventResponse) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ScheduledPhoneCallEventResponse) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryAttempts

`func (o *ScheduledPhoneCallEventResponse) GetRetryAttempts() int32`

GetRetryAttempts returns the RetryAttempts field if non-nil, zero value otherwise.

### GetRetryAttemptsOk

`func (o *ScheduledPhoneCallEventResponse) GetRetryAttemptsOk() (*int32, bool)`

GetRetryAttemptsOk returns a tuple with the RetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttempts

`func (o *ScheduledPhoneCallEventResponse) SetRetryAttempts(v int32)`

SetRetryAttempts sets RetryAttempts field to given value.

### HasRetryAttempts

`func (o *ScheduledPhoneCallEventResponse) HasRetryAttempts() bool`

HasRetryAttempts returns a boolean if a field has been set.

### GetScheduledEventId

`func (o *ScheduledPhoneCallEventResponse) GetScheduledEventId() string`

GetScheduledEventId returns the ScheduledEventId field if non-nil, zero value otherwise.

### GetScheduledEventIdOk

`func (o *ScheduledPhoneCallEventResponse) GetScheduledEventIdOk() (*string, bool)`

GetScheduledEventIdOk returns a tuple with the ScheduledEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEventId

`func (o *ScheduledPhoneCallEventResponse) SetScheduledEventId(v string)`

SetScheduledEventId sets ScheduledEventId field to given value.

### HasScheduledEventId

`func (o *ScheduledPhoneCallEventResponse) HasScheduledEventId() bool`

HasScheduledEventId returns a boolean if a field has been set.

### GetConversationId

`func (o *ScheduledPhoneCallEventResponse) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ScheduledPhoneCallEventResponse) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ScheduledPhoneCallEventResponse) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ScheduledPhoneCallEventResponse) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScheduledPhoneCallEventResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledPhoneCallEventResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledPhoneCallEventResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduledPhoneCallEventResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduledPhoneCallEventResponse) GetStatus() EventStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledPhoneCallEventResponse) GetStatusOk() (*EventStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledPhoneCallEventResponse) SetStatus(v EventStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledPhoneCallEventResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConversationMetadata

`func (o *ScheduledPhoneCallEventResponse) GetConversationMetadata() map[string]ConversationMetadataValue`

GetConversationMetadata returns the ConversationMetadata field if non-nil, zero value otherwise.

### GetConversationMetadataOk

`func (o *ScheduledPhoneCallEventResponse) GetConversationMetadataOk() (*map[string]ConversationMetadataValue, bool)`

GetConversationMetadataOk returns a tuple with the ConversationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationMetadata

`func (o *ScheduledPhoneCallEventResponse) SetConversationMetadata(v map[string]ConversationMetadataValue)`

SetConversationMetadata sets ConversationMetadata field to given value.

### HasConversationMetadata

`func (o *ScheduledPhoneCallEventResponse) HasConversationMetadata() bool`

HasConversationMetadata returns a boolean if a field has been set.

### GetErrors

`func (o *ScheduledPhoneCallEventResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ScheduledPhoneCallEventResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ScheduledPhoneCallEventResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ScheduledPhoneCallEventResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


