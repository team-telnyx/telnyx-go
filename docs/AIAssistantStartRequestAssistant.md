# AIAssistantStartRequestAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the AI assistant to use | [optional] 
**OpenaiApiKeyRef** | Pointer to **string** | Reference to the OpenAI API key. Required only when using OpenAI models | [optional] 
**Instructions** | Pointer to **string** | The system instructions that the voice assistant uses during the start assistant command. This will overwrite the instructions set in the assistant configuration. | [optional] 

## Methods

### NewAIAssistantStartRequestAssistant

`func NewAIAssistantStartRequestAssistant() *AIAssistantStartRequestAssistant`

NewAIAssistantStartRequestAssistant instantiates a new AIAssistantStartRequestAssistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAssistantStartRequestAssistantWithDefaults

`func NewAIAssistantStartRequestAssistantWithDefaults() *AIAssistantStartRequestAssistant`

NewAIAssistantStartRequestAssistantWithDefaults instantiates a new AIAssistantStartRequestAssistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AIAssistantStartRequestAssistant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIAssistantStartRequestAssistant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIAssistantStartRequestAssistant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AIAssistantStartRequestAssistant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpenaiApiKeyRef

`func (o *AIAssistantStartRequestAssistant) GetOpenaiApiKeyRef() string`

GetOpenaiApiKeyRef returns the OpenaiApiKeyRef field if non-nil, zero value otherwise.

### GetOpenaiApiKeyRefOk

`func (o *AIAssistantStartRequestAssistant) GetOpenaiApiKeyRefOk() (*string, bool)`

GetOpenaiApiKeyRefOk returns a tuple with the OpenaiApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenaiApiKeyRef

`func (o *AIAssistantStartRequestAssistant) SetOpenaiApiKeyRef(v string)`

SetOpenaiApiKeyRef sets OpenaiApiKeyRef field to given value.

### HasOpenaiApiKeyRef

`func (o *AIAssistantStartRequestAssistant) HasOpenaiApiKeyRef() bool`

HasOpenaiApiKeyRef returns a boolean if a field has been set.

### GetInstructions

`func (o *AIAssistantStartRequestAssistant) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AIAssistantStartRequestAssistant) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AIAssistantStartRequestAssistant) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AIAssistantStartRequestAssistant) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


