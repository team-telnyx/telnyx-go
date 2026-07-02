// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/shared/constant"
)

// AIOpenAIChatService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIChatService] method instead.
type AIOpenAIChatService struct {
	Options []option.RequestOption
}

// NewAIOpenAIChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIChatService(opts ...option.RequestOption) (r AIOpenAIChatService) {
	r = AIOpenAIChatService{}
	r.Options = opts
	return
}

// Chat with a language model. This endpoint is consistent with the
// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat)
// and may be used with the OpenAI JS or Python SDK by setting the base URL to
// `https://api.telnyx.com/v2/ai/openai`.
func (r *AIOpenAIChatService) NewCompletion(ctx context.Context, body AIOpenAIChatNewCompletionParams, opts ...option.RequestOption) (res *AIOpenAIChatNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AIOpenAIChatNewCompletionResponse map[string]any

type AIOpenAIChatNewCompletionParams struct {
	// A list of the previous chat messages for context.
	Messages []AIOpenAIChatNewCompletionParamsMessage `json:"messages,omitzero" api:"required"`
	// If you are using an external inference provider like xAI or OpenAI, this field
	// allows you to pass along a reference to your API key. After creating an
	// [integration secret](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for you API key, pass the secret's `identifier` in this field.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// This is used with `use_beam_search` to determine how many candidate beams to
	// explore.
	BestOf param.Opt[int64] `json:"best_of,omitzero"`
	// This is used with `use_beam_search`. If `true`, generation stops as soon as
	// there are `best_of` complete candidates; if `false`, a heuristic is applied and
	// the generation stops when is it very unlikely to find better candidates.
	EarlyStopping param.Opt[bool] `json:"early_stopping,omitzero"`
	// Whether to enable the thinking/reasoning phase for models that support it (e.g.,
	// QwQ, Qwen3). When set to false, the model will skip the internal reasoning step
	// and respond directly, which can reduce latency. Defaults to true.
	EnableThinking param.Opt[bool] `json:"enable_thinking,omitzero"`
	// Higher values will penalize the model from repeating the same output tokens.
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	// If specified, the output will follow the regex pattern.
	GuidedRegex param.Opt[string] `json:"guided_regex,omitzero"`
	// This is used with `use_beam_search` to prefer shorter or longer completions.
	LengthPenalty param.Opt[float64] `json:"length_penalty,omitzero"`
	// Whether to return log probabilities of the output tokens or not. If true,
	// returns the log probabilities of each output token returned in the `content` of
	// `message`.
	Logprobs param.Opt[bool] `json:"logprobs,omitzero"`
	// Maximum number of completion tokens the model should generate.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// This is an alternative to `top_p` that
	// [many prefer](https://github.com/huggingface/transformers/issues/27670). Must be
	// in [0, 1].
	MinP param.Opt[float64] `json:"min_p,omitzero"`
	// The language model to chat with.
	Model param.Opt[string] `json:"model,omitzero"`
	// This will return multiple choices for you instead of a single chat completion.
	N param.Opt[float64] `json:"n,omitzero"`
	// Higher values will penalize the model from repeating the same output tokens.
	PresencePenalty param.Opt[float64] `json:"presence_penalty,omitzero"`
	// If specified, the system will make a best effort to sample deterministically,
	// such that repeated requests with the same `seed` and parameters should return
	// the same result.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Whether or not to stream data-only server-sent events as they become available.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Adjusts the "creativity" of the model. Lower values make the model more
	// deterministic and repetitive, while higher values make the model more random and
	// creative.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// This is used with `logprobs`. An integer between 0 and 20 specifying the number
	// of most likely tokens to return at each token position, each with an associated
	// log probability.
	TopLogprobs param.Opt[int64] `json:"top_logprobs,omitzero"`
	// An alternative or complement to `temperature`. This adjusts how many of the top
	// possibilities to consider.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Setting this to `true` will allow the model to
	// [explore more completion options](https://huggingface.co/blog/how-to-generate#beam-search).
	// This is not supported by OpenAI.
	UseBeamSearch param.Opt[bool] `json:"use_beam_search,omitzero"`
	// If specified, the output will be exactly one of the choices.
	GuidedChoice []string `json:"guided_choice,omitzero"`
	// Must be a valid JSON schema. If specified, the output will follow the JSON
	// schema.
	GuidedJson map[string]any `json:"guided_json,omitzero"`
	// Use this is you want to guarantee a JSON output without defining a schema. For
	// control over the schema, use `guided_json`.
	ResponseFormat AIOpenAIChatNewCompletionParamsResponseFormat `json:"response_format,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop AIOpenAIChatNewCompletionParamsStopUnion `json:"stop,omitzero"`
	// Any of "none", "auto", "required".
	ToolChoice AIOpenAIChatNewCompletionParamsToolChoice `json:"tool_choice,omitzero"`
	// The `function` tool type follows the same schema as the
	// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat).
	// The `retrieval` tool type is unique to Telnyx. You may pass a list of
	// [embedded storage buckets](https://developers.telnyx.com/api-reference/embeddings/embed-documents)
	// for retrieval-augmented generation.
	Tools []AIOpenAIChatNewCompletionParamsToolUnion `json:"tools,omitzero"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type AIOpenAIChatNewCompletionParamsMessage struct {
	Content AIOpenAIChatNewCompletionParamsMessageContentUnion `json:"content,omitzero" api:"required"`
	// Any of "system", "user", "assistant", "tool".
	Role string `json:"role,omitzero" api:"required"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIOpenAIChatNewCompletionParamsMessage](
		"role", "system", "user", "assistant", "tool",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIOpenAIChatNewCompletionParamsMessageContentUnion struct {
	OfString            param.Opt[string]                                                    `json:",omitzero,inline"`
	OfTextAndImageArray []AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u AIOpenAIChatNewCompletionParamsMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfTextAndImageArray)
}
func (u *AIOpenAIChatNewCompletionParamsMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIOpenAIChatNewCompletionParamsMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfTextAndImageArray) {
		return &u.OfTextAndImageArray
	}
	return nil
}

// The property Type is required.
type AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem struct {
	// Any of "text", "image_url".
	Type     string            `json:"type,omitzero" api:"required"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	Text     param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIOpenAIChatNewCompletionParamsMessageContentTextAndImageArrayItem](
		"type", "text", "image_url",
	)
}

// Use this is you want to guarantee a JSON output without defining a schema. For
// control over the schema, use `guided_json`.
//
// The property Type is required.
type AIOpenAIChatNewCompletionParamsResponseFormat struct {
	// Any of "text", "json_object".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsResponseFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIOpenAIChatNewCompletionParamsResponseFormat](
		"type", "text", "json_object",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIOpenAIChatNewCompletionParamsStopUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u AIOpenAIChatNewCompletionParamsStopUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *AIOpenAIChatNewCompletionParamsStopUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIOpenAIChatNewCompletionParamsStopUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type AIOpenAIChatNewCompletionParamsToolChoice string

const (
	AIOpenAIChatNewCompletionParamsToolChoiceNone     AIOpenAIChatNewCompletionParamsToolChoice = "none"
	AIOpenAIChatNewCompletionParamsToolChoiceAuto     AIOpenAIChatNewCompletionParamsToolChoice = "auto"
	AIOpenAIChatNewCompletionParamsToolChoiceRequired AIOpenAIChatNewCompletionParamsToolChoice = "required"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIOpenAIChatNewCompletionParamsToolUnion struct {
	OfFunction  *AIOpenAIChatNewCompletionParamsToolFunction  `json:",omitzero,inline"`
	OfRetrieval *AIOpenAIChatNewCompletionParamsToolRetrieval `json:",omitzero,inline"`
	paramUnion
}

func (u AIOpenAIChatNewCompletionParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunction, u.OfRetrieval)
}
func (u *AIOpenAIChatNewCompletionParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIOpenAIChatNewCompletionParamsToolUnion) asAny() any {
	if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIOpenAIChatNewCompletionParamsToolUnion) GetFunction() *AIOpenAIChatNewCompletionParamsToolFunctionFunction {
	if vt := u.OfFunction; vt != nil {
		return &vt.Function
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIOpenAIChatNewCompletionParamsToolUnion) GetRetrieval() *BucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIOpenAIChatNewCompletionParamsToolUnion) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AIOpenAIChatNewCompletionParamsToolUnion](
		"type",
		apijson.Discriminator[AIOpenAIChatNewCompletionParamsToolFunction]("function"),
		apijson.Discriminator[AIOpenAIChatNewCompletionParamsToolRetrieval]("retrieval"),
	)
}

// The properties Function, Type are required.
type AIOpenAIChatNewCompletionParamsToolFunction struct {
	Function AIOpenAIChatNewCompletionParamsToolFunctionFunction `json:"function,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type" default:"function"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsToolFunction) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsToolFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type AIOpenAIChatNewCompletionParamsToolFunctionFunction struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Parameters  map[string]any    `json:"parameters,omitzero"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsToolFunctionFunction) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsToolFunctionFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsToolFunctionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retrieval, Type are required.
type AIOpenAIChatNewCompletionParamsToolRetrieval struct {
	Retrieval BucketIDsParam `json:"retrieval,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "retrieval".
	Type constant.Retrieval `json:"type" default:"retrieval"`
	paramObj
}

func (r AIOpenAIChatNewCompletionParamsToolRetrieval) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAIChatNewCompletionParamsToolRetrieval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAIChatNewCompletionParamsToolRetrieval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
