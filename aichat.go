// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/shared/constant"
)

// AIChatService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIChatService] method instead.
type AIChatService struct {
	Options []option.RequestOption
}

// NewAIChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIChatService(opts ...option.RequestOption) (r AIChatService) {
	r = AIChatService{}
	r.Options = opts
	return
}

// Chat with a language model. This endpoint is consistent with the
// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat)
// and may be used with the OpenAI JS or Python SDK.
func (r *AIChatService) NewCompletion(ctx context.Context, body AIChatNewCompletionParams, opts ...option.RequestOption) (res *AIChatNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIChatNewCompletionResponse = any

type AIChatNewCompletionParams struct {
	// A list of the previous chat messages for context.
	Messages []AIChatNewCompletionParamsMessage `json:"messages,omitzero,required"`
	// If you are using an external inference provider like xAI or OpenAI, this field
	// allows you to pass along a reference to your API key. After creating an
	// [integration secret](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// for you API key, pass the secret's `identifier` in this field.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// This is used with `use_beam_search` to determine how many candidate beams to
	// explore.
	BestOf param.Opt[int64] `json:"best_of,omitzero"`
	// This is used with `use_beam_search`. If `true`, generation stops as soon as
	// there are `best_of` complete candidates; if `false`, a heuristic is applied and
	// the generation stops when is it very unlikely to find better candidates.
	EarlyStopping param.Opt[bool] `json:"early_stopping,omitzero"`
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
	ResponseFormat AIChatNewCompletionParamsResponseFormat `json:"response_format,omitzero"`
	// Any of "none", "auto", "required".
	ToolChoice AIChatNewCompletionParamsToolChoice `json:"tool_choice,omitzero"`
	// The `function` tool type follows the same schema as the
	// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat).
	// The `retrieval` tool type is unique to Telnyx. You may pass a list of
	// [embedded storage buckets](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding)
	// for retrieval-augmented generation.
	Tools []AIChatNewCompletionParamsToolUnion `json:"tools,omitzero"`
	paramObj
}

func (r AIChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type AIChatNewCompletionParamsMessage struct {
	Content AIChatNewCompletionParamsMessageContentUnion `json:"content,omitzero,required"`
	// Any of "system", "user", "assistant", "tool".
	Role string `json:"role,omitzero,required"`
	paramObj
}

func (r AIChatNewCompletionParamsMessage) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIChatNewCompletionParamsMessage](
		"role", "system", "user", "assistant", "tool",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIChatNewCompletionParamsMessageContentUnion struct {
	OfString            param.Opt[string]                                              `json:",omitzero,inline"`
	OfTextAndImageArray []AIChatNewCompletionParamsMessageContentTextAndImageArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u AIChatNewCompletionParamsMessageContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfTextAndImageArray)
}
func (u *AIChatNewCompletionParamsMessageContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIChatNewCompletionParamsMessageContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfTextAndImageArray) {
		return &u.OfTextAndImageArray
	}
	return nil
}

// The property Type is required.
type AIChatNewCompletionParamsMessageContentTextAndImageArrayItem struct {
	// Any of "text", "image_url".
	Type     string            `json:"type,omitzero,required"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	Text     param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r AIChatNewCompletionParamsMessageContentTextAndImageArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsMessageContentTextAndImageArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsMessageContentTextAndImageArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIChatNewCompletionParamsMessageContentTextAndImageArrayItem](
		"type", "text", "image_url",
	)
}

// Use this is you want to guarantee a JSON output without defining a schema. For
// control over the schema, use `guided_json`.
//
// The property Type is required.
type AIChatNewCompletionParamsResponseFormat struct {
	// Any of "text", "json_object".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r AIChatNewCompletionParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsResponseFormat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AIChatNewCompletionParamsResponseFormat](
		"type", "text", "json_object",
	)
}

type AIChatNewCompletionParamsToolChoice string

const (
	AIChatNewCompletionParamsToolChoiceNone     AIChatNewCompletionParamsToolChoice = "none"
	AIChatNewCompletionParamsToolChoiceAuto     AIChatNewCompletionParamsToolChoice = "auto"
	AIChatNewCompletionParamsToolChoiceRequired AIChatNewCompletionParamsToolChoice = "required"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIChatNewCompletionParamsToolUnion struct {
	OfFunction  *AIChatNewCompletionParamsToolFunction  `json:",omitzero,inline"`
	OfRetrieval *AIChatNewCompletionParamsToolRetrieval `json:",omitzero,inline"`
	paramUnion
}

func (u AIChatNewCompletionParamsToolUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunction, u.OfRetrieval)
}
func (u *AIChatNewCompletionParamsToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIChatNewCompletionParamsToolUnion) asAny() any {
	if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIChatNewCompletionParamsToolUnion) GetFunction() *AIChatNewCompletionParamsToolFunctionFunction {
	if vt := u.OfFunction; vt != nil {
		return &vt.Function
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIChatNewCompletionParamsToolUnion) GetRetrieval() *InferenceEmbeddingBucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AIChatNewCompletionParamsToolUnion) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AIChatNewCompletionParamsToolUnion](
		"type",
		apijson.Discriminator[AIChatNewCompletionParamsToolFunction]("function"),
		apijson.Discriminator[AIChatNewCompletionParamsToolRetrieval]("retrieval"),
	)
}

// The properties Function, Type are required.
type AIChatNewCompletionParamsToolFunction struct {
	Function AIChatNewCompletionParamsToolFunctionFunction `json:"function,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r AIChatNewCompletionParamsToolFunction) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsToolFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsToolFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type AIChatNewCompletionParamsToolFunctionFunction struct {
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Parameters  map[string]any    `json:"parameters,omitzero"`
	paramObj
}

func (r AIChatNewCompletionParamsToolFunctionFunction) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsToolFunctionFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsToolFunctionFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retrieval, Type are required.
type AIChatNewCompletionParamsToolRetrieval struct {
	Retrieval InferenceEmbeddingBucketIDsParam `json:"retrieval,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "retrieval".
	Type constant.Retrieval `json:"type,required"`
	paramObj
}

func (r AIChatNewCompletionParamsToolRetrieval) MarshalJSON() (data []byte, err error) {
	type shadow AIChatNewCompletionParamsToolRetrieval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIChatNewCompletionParamsToolRetrieval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
