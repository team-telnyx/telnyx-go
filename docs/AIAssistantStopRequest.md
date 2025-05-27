# AIAssistantStopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewAIAssistantStopRequest

`func NewAIAssistantStopRequest() *AIAssistantStopRequest`

NewAIAssistantStopRequest instantiates a new AIAssistantStopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAssistantStopRequestWithDefaults

`func NewAIAssistantStopRequestWithDefaults() *AIAssistantStopRequest`

NewAIAssistantStopRequestWithDefaults instantiates a new AIAssistantStopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientState

`func (o *AIAssistantStopRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *AIAssistantStopRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *AIAssistantStopRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *AIAssistantStopRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *AIAssistantStopRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *AIAssistantStopRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *AIAssistantStopRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *AIAssistantStopRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


