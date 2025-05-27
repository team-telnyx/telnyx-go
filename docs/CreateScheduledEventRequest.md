# CreateScheduledEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TelnyxConversationChannel** | [**ConversationChannelType**](ConversationChannelType.md) |  | 
**TelnyxEndUserTarget** | **string** | The phone number, SIP URI, to schedule the call or text to. | 
**TelnyxAgentTarget** | **string** | The phone number, SIP URI, to schedule the call or text from. | 
**ScheduledAtFixedDatetime** | **time.Time** | The datetime at which the event should be scheduled. Formatted as ISO 8601. | 
**Text** | Pointer to **string** | Required for sms scheduled events. The text to be sent to the end user. | [optional] 
**ConversationMetadata** | Pointer to [**map[string]ConversationMetadataValue**](ConversationMetadataValue.md) | Metadata associated with the conversation. Telnyx provides several pieces of metadata, but customers can also add their own. | [optional] 

## Methods

### NewCreateScheduledEventRequest

`func NewCreateScheduledEventRequest(telnyxConversationChannel ConversationChannelType, telnyxEndUserTarget string, telnyxAgentTarget string, scheduledAtFixedDatetime time.Time, ) *CreateScheduledEventRequest`

NewCreateScheduledEventRequest instantiates a new CreateScheduledEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduledEventRequestWithDefaults

`func NewCreateScheduledEventRequestWithDefaults() *CreateScheduledEventRequest`

NewCreateScheduledEventRequestWithDefaults instantiates a new CreateScheduledEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelnyxConversationChannel

`func (o *CreateScheduledEventRequest) GetTelnyxConversationChannel() ConversationChannelType`

GetTelnyxConversationChannel returns the TelnyxConversationChannel field if non-nil, zero value otherwise.

### GetTelnyxConversationChannelOk

`func (o *CreateScheduledEventRequest) GetTelnyxConversationChannelOk() (*ConversationChannelType, bool)`

GetTelnyxConversationChannelOk returns a tuple with the TelnyxConversationChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxConversationChannel

`func (o *CreateScheduledEventRequest) SetTelnyxConversationChannel(v ConversationChannelType)`

SetTelnyxConversationChannel sets TelnyxConversationChannel field to given value.


### GetTelnyxEndUserTarget

`func (o *CreateScheduledEventRequest) GetTelnyxEndUserTarget() string`

GetTelnyxEndUserTarget returns the TelnyxEndUserTarget field if non-nil, zero value otherwise.

### GetTelnyxEndUserTargetOk

`func (o *CreateScheduledEventRequest) GetTelnyxEndUserTargetOk() (*string, bool)`

GetTelnyxEndUserTargetOk returns a tuple with the TelnyxEndUserTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxEndUserTarget

`func (o *CreateScheduledEventRequest) SetTelnyxEndUserTarget(v string)`

SetTelnyxEndUserTarget sets TelnyxEndUserTarget field to given value.


### GetTelnyxAgentTarget

`func (o *CreateScheduledEventRequest) GetTelnyxAgentTarget() string`

GetTelnyxAgentTarget returns the TelnyxAgentTarget field if non-nil, zero value otherwise.

### GetTelnyxAgentTargetOk

`func (o *CreateScheduledEventRequest) GetTelnyxAgentTargetOk() (*string, bool)`

GetTelnyxAgentTargetOk returns a tuple with the TelnyxAgentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxAgentTarget

`func (o *CreateScheduledEventRequest) SetTelnyxAgentTarget(v string)`

SetTelnyxAgentTarget sets TelnyxAgentTarget field to given value.


### GetScheduledAtFixedDatetime

`func (o *CreateScheduledEventRequest) GetScheduledAtFixedDatetime() time.Time`

GetScheduledAtFixedDatetime returns the ScheduledAtFixedDatetime field if non-nil, zero value otherwise.

### GetScheduledAtFixedDatetimeOk

`func (o *CreateScheduledEventRequest) GetScheduledAtFixedDatetimeOk() (*time.Time, bool)`

GetScheduledAtFixedDatetimeOk returns a tuple with the ScheduledAtFixedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAtFixedDatetime

`func (o *CreateScheduledEventRequest) SetScheduledAtFixedDatetime(v time.Time)`

SetScheduledAtFixedDatetime sets ScheduledAtFixedDatetime field to given value.


### GetText

`func (o *CreateScheduledEventRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateScheduledEventRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateScheduledEventRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateScheduledEventRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetConversationMetadata

`func (o *CreateScheduledEventRequest) GetConversationMetadata() map[string]ConversationMetadataValue`

GetConversationMetadata returns the ConversationMetadata field if non-nil, zero value otherwise.

### GetConversationMetadataOk

`func (o *CreateScheduledEventRequest) GetConversationMetadataOk() (*map[string]ConversationMetadataValue, bool)`

GetConversationMetadataOk returns a tuple with the ConversationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationMetadata

`func (o *CreateScheduledEventRequest) SetConversationMetadata(v map[string]ConversationMetadataValue)`

SetConversationMetadata sets ConversationMetadata field to given value.

### HasConversationMetadata

`func (o *CreateScheduledEventRequest) HasConversationMetadata() bool`

HasConversationMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


