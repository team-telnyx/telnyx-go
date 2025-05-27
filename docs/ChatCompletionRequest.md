# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]ChatCompletionSystemMessageParam**](ChatCompletionSystemMessageParam.md) | A list of the previous chat messages for context. | 
**Model** | Pointer to **string** | The language model to chat with. If you are optimizing for speed + price, try &#x60;meta-llama/Meta-Llama-3.1-8B-Instruct&#x60;. For quality, try &#x60;meta-llama/Meta-Llama-3.1-70B-Instruct&#x60;. Or explore our [LLM Library](https://telnyx.com/products/llm-library). | [optional] [default to "meta-llama/Meta-Llama-3.1-8B-Instruct"]
**ApiKeyRef** | Pointer to **string** | If you are using an external inference provider like xAI or OpenAI, this field allows you to pass along a reference to your API key. After creating an [integration secret](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for you API key, pass the secret&#39;s &#x60;identifier&#x60; in this field. | [optional] 
**Stream** | Pointer to **bool** | Whether or not to stream data-only server-sent events as they become available. | [optional] [default to false]
**Temperature** | Pointer to **float32** | Adjusts the \&quot;creativity\&quot; of the model. Lower values make the model more deterministic and repetitive, while higher values make the model more random and creative. | [optional] [default to 0.1]
**MaxTokens** | Pointer to **int32** | Maximum number of completion tokens the model should generate. | [optional] 
**Tools** | Pointer to [**[]ChatCompletionRequestToolsInner**](ChatCompletionRequestToolsInner.md) | The &#x60;function&#x60; tool type follows the same schema as the [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat). The &#x60;retrieval&#x60; tool type is unique to Telnyx. You may pass a list of [embedded storage buckets](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding) for retrieval-augmented generation. | [optional] 
**ToolChoice** | Pointer to **string** |  | [optional] 
**ResponseFormat** | Pointer to [**ChatCompletionResponseFormatParam**](ChatCompletionResponseFormatParam.md) |  | [optional] 
**GuidedJson** | Pointer to **map[string]interface{}** | Must be a valid JSON schema. If specified, the output will follow the JSON schema. | [optional] 
**GuidedRegex** | Pointer to **string** | If specified, the output will follow the regex pattern. | [optional] 
**GuidedChoice** | Pointer to **[]string** | If specified, the output will be exactly one of the choices. | [optional] 
**MinP** | Pointer to **float32** | This is an alternative to &#x60;top_p&#x60; that [many prefer](https://github.com/huggingface/transformers/issues/27670). Must be in [0, 1]. | [optional] 
**N** | Pointer to **float32** | This will return multiple choices for you instead of a single chat completion. | [optional] 
**UseBeamSearch** | Pointer to **bool** | Setting this to &#x60;true&#x60; will allow the model to [explore more completion options](https://huggingface.co/blog/how-to-generate#beam-search). This is not supported by OpenAI. | [optional] [default to false]
**BestOf** | Pointer to **int32** | This is used with &#x60;use_beam_search&#x60; to determine how many candidate beams to explore. | [optional] 
**LengthPenalty** | Pointer to **float32** | This is used with &#x60;use_beam_search&#x60; to prefer shorter or longer completions. | [optional] [default to 1]
**EarlyStopping** | Pointer to **bool** | This is used with &#x60;use_beam_search&#x60;. If &#x60;true&#x60;, generation stops as soon as there are &#x60;best_of&#x60; complete candidates; if &#x60;false&#x60;, a heuristic is applied and the generation stops when is it very unlikely to find better candidates. | [optional] [default to false]
**Logprobs** | Pointer to **bool** | Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the &#x60;content&#x60; of &#x60;message&#x60;. | [optional] [default to false]
**TopLogprobs** | Pointer to **int32** | This is used with &#x60;logprobs&#x60;. An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability. | [optional] 
**FrequencyPenalty** | Pointer to **float32** | Higher values will penalize the model from repeating the same output tokens. | [optional] [default to 0]
**PresencePenalty** | Pointer to **float32** | Higher values will penalize the model from repeating the same output tokens. | [optional] [default to 0]
**TopP** | Pointer to **float32** | An alternative or complement to &#x60;temperature&#x60;. This adjusts how many of the top possibilities to consider. | [optional] 

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(messages []ChatCompletionSystemMessageParam, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []ChatCompletionSystemMessageParam`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]ChatCompletionSystemMessageParam, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []ChatCompletionSystemMessageParam)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *ChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatCompletionRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetApiKeyRef

`func (o *ChatCompletionRequest) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *ChatCompletionRequest) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *ChatCompletionRequest) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.

### HasApiKeyRef

`func (o *ChatCompletionRequest) HasApiKeyRef() bool`

HasApiKeyRef returns a boolean if a field has been set.

### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetTemperature

`func (o *ChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetMaxTokens

`func (o *ChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetTools

`func (o *ChatCompletionRequest) GetTools() []ChatCompletionRequestToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatCompletionRequest) GetToolsOk() (*[]ChatCompletionRequestToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatCompletionRequest) SetTools(v []ChatCompletionRequestToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *ChatCompletionRequest) GetToolChoice() string`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ChatCompletionRequest) GetToolChoiceOk() (*string, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ChatCompletionRequest) SetToolChoice(v string)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetResponseFormat

`func (o *ChatCompletionRequest) GetResponseFormat() ChatCompletionResponseFormatParam`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatCompletionRequest) GetResponseFormatOk() (*ChatCompletionResponseFormatParam, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatCompletionRequest) SetResponseFormat(v ChatCompletionResponseFormatParam)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetGuidedJson

`func (o *ChatCompletionRequest) GetGuidedJson() map[string]interface{}`

GetGuidedJson returns the GuidedJson field if non-nil, zero value otherwise.

### GetGuidedJsonOk

`func (o *ChatCompletionRequest) GetGuidedJsonOk() (*map[string]interface{}, bool)`

GetGuidedJsonOk returns a tuple with the GuidedJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidedJson

`func (o *ChatCompletionRequest) SetGuidedJson(v map[string]interface{})`

SetGuidedJson sets GuidedJson field to given value.

### HasGuidedJson

`func (o *ChatCompletionRequest) HasGuidedJson() bool`

HasGuidedJson returns a boolean if a field has been set.

### GetGuidedRegex

`func (o *ChatCompletionRequest) GetGuidedRegex() string`

GetGuidedRegex returns the GuidedRegex field if non-nil, zero value otherwise.

### GetGuidedRegexOk

`func (o *ChatCompletionRequest) GetGuidedRegexOk() (*string, bool)`

GetGuidedRegexOk returns a tuple with the GuidedRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidedRegex

`func (o *ChatCompletionRequest) SetGuidedRegex(v string)`

SetGuidedRegex sets GuidedRegex field to given value.

### HasGuidedRegex

`func (o *ChatCompletionRequest) HasGuidedRegex() bool`

HasGuidedRegex returns a boolean if a field has been set.

### GetGuidedChoice

`func (o *ChatCompletionRequest) GetGuidedChoice() []string`

GetGuidedChoice returns the GuidedChoice field if non-nil, zero value otherwise.

### GetGuidedChoiceOk

`func (o *ChatCompletionRequest) GetGuidedChoiceOk() (*[]string, bool)`

GetGuidedChoiceOk returns a tuple with the GuidedChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidedChoice

`func (o *ChatCompletionRequest) SetGuidedChoice(v []string)`

SetGuidedChoice sets GuidedChoice field to given value.

### HasGuidedChoice

`func (o *ChatCompletionRequest) HasGuidedChoice() bool`

HasGuidedChoice returns a boolean if a field has been set.

### GetMinP

`func (o *ChatCompletionRequest) GetMinP() float32`

GetMinP returns the MinP field if non-nil, zero value otherwise.

### GetMinPOk

`func (o *ChatCompletionRequest) GetMinPOk() (*float32, bool)`

GetMinPOk returns a tuple with the MinP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinP

`func (o *ChatCompletionRequest) SetMinP(v float32)`

SetMinP sets MinP field to given value.

### HasMinP

`func (o *ChatCompletionRequest) HasMinP() bool`

HasMinP returns a boolean if a field has been set.

### GetN

`func (o *ChatCompletionRequest) GetN() float32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ChatCompletionRequest) GetNOk() (*float32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ChatCompletionRequest) SetN(v float32)`

SetN sets N field to given value.

### HasN

`func (o *ChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUseBeamSearch

`func (o *ChatCompletionRequest) GetUseBeamSearch() bool`

GetUseBeamSearch returns the UseBeamSearch field if non-nil, zero value otherwise.

### GetUseBeamSearchOk

`func (o *ChatCompletionRequest) GetUseBeamSearchOk() (*bool, bool)`

GetUseBeamSearchOk returns a tuple with the UseBeamSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBeamSearch

`func (o *ChatCompletionRequest) SetUseBeamSearch(v bool)`

SetUseBeamSearch sets UseBeamSearch field to given value.

### HasUseBeamSearch

`func (o *ChatCompletionRequest) HasUseBeamSearch() bool`

HasUseBeamSearch returns a boolean if a field has been set.

### GetBestOf

`func (o *ChatCompletionRequest) GetBestOf() int32`

GetBestOf returns the BestOf field if non-nil, zero value otherwise.

### GetBestOfOk

`func (o *ChatCompletionRequest) GetBestOfOk() (*int32, bool)`

GetBestOfOk returns a tuple with the BestOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOf

`func (o *ChatCompletionRequest) SetBestOf(v int32)`

SetBestOf sets BestOf field to given value.

### HasBestOf

`func (o *ChatCompletionRequest) HasBestOf() bool`

HasBestOf returns a boolean if a field has been set.

### GetLengthPenalty

`func (o *ChatCompletionRequest) GetLengthPenalty() float32`

GetLengthPenalty returns the LengthPenalty field if non-nil, zero value otherwise.

### GetLengthPenaltyOk

`func (o *ChatCompletionRequest) GetLengthPenaltyOk() (*float32, bool)`

GetLengthPenaltyOk returns a tuple with the LengthPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthPenalty

`func (o *ChatCompletionRequest) SetLengthPenalty(v float32)`

SetLengthPenalty sets LengthPenalty field to given value.

### HasLengthPenalty

`func (o *ChatCompletionRequest) HasLengthPenalty() bool`

HasLengthPenalty returns a boolean if a field has been set.

### GetEarlyStopping

`func (o *ChatCompletionRequest) GetEarlyStopping() bool`

GetEarlyStopping returns the EarlyStopping field if non-nil, zero value otherwise.

### GetEarlyStoppingOk

`func (o *ChatCompletionRequest) GetEarlyStoppingOk() (*bool, bool)`

GetEarlyStoppingOk returns a tuple with the EarlyStopping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyStopping

`func (o *ChatCompletionRequest) SetEarlyStopping(v bool)`

SetEarlyStopping sets EarlyStopping field to given value.

### HasEarlyStopping

`func (o *ChatCompletionRequest) HasEarlyStopping() bool`

HasEarlyStopping returns a boolean if a field has been set.

### GetLogprobs

`func (o *ChatCompletionRequest) GetLogprobs() bool`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ChatCompletionRequest) GetLogprobsOk() (*bool, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ChatCompletionRequest) SetLogprobs(v bool)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ChatCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### GetTopLogprobs

`func (o *ChatCompletionRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ChatCompletionRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ChatCompletionRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ChatCompletionRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *ChatCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ChatCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ChatCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *ChatCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ChatCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ChatCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### GetTopP

`func (o *ChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


