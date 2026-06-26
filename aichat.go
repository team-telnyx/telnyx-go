// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Generate text with LLMs
//
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

// **Deprecated**: Use `POST /v2/ai/openai/chat/completions` instead. Chat with a
// language model. This endpoint is consistent with the
// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat)
// and may be used with the OpenAI JS or Python SDK.
//
// Deprecated: deprecated
func (r *AIChatService) NewCompletion(ctx context.Context, body AIChatNewCompletionParams, opts ...option.RequestOption) (res *AIChatNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type BucketIDs struct {
	// List of
	// [embedded storage buckets](https://developers.telnyx.com/api-reference/embeddings/embed-documents)
	// to use for retrieval-augmented generation.
	BucketIDs []string `json:"bucket_ids" api:"required"`
	// The maximum number of results to retrieve as context for the language model.
	MaxNumResults int64 `json:"max_num_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketIDs     respjson.Field
		MaxNumResults respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketIDs) RawJSON() string { return r.JSON.raw }
func (r *BucketIDs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BucketIDs to a BucketIDsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BucketIDsParam.Overrides()
func (r BucketIDs) ToParam() BucketIDsParam {
	return param.Override[BucketIDsParam](json.RawMessage(r.RawJSON()))
}

// The property BucketIDs is required.
type BucketIDsParam struct {
	// List of
	// [embedded storage buckets](https://developers.telnyx.com/api-reference/embeddings/embed-documents)
	// to use for retrieval-augmented generation.
	BucketIDs []string `json:"bucket_ids,omitzero" api:"required"`
	// The maximum number of results to retrieve as context for the language model.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r BucketIDsParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketIDsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketIDsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Messages is required.
type ChatCompletionRequestParam struct {
	// A list of the previous chat messages for context.
	Messages []ChatCompletionRequestMessageParam `json:"messages,omitzero" api:"required"`
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
	ResponseFormat ChatCompletionRequestResponseFormatParam `json:"response_format,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop ChatCompletionRequestStopUnionParam `json:"stop,omitzero"`
	// Any of "none", "auto", "required".
	ToolChoice ChatCompletionRequestToolChoice `json:"tool_choice,omitzero"`
	// The `function` tool type follows the same schema as the
	// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat).
	// The `retrieval` tool type is unique to Telnyx. You may pass a list of
	// [embedded storage buckets](https://developers.telnyx.com/api-reference/embeddings/embed-documents)
	// for retrieval-augmented generation.
	Tools []ChatCompletionRequestToolsUnionParam `json:"tools,omitzero"`
	paramObj
}

func (r ChatCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type ChatCompletionRequestMessageParam struct {
	Content ChatCompletionRequestMessagesContentUnionParam `json:"content,omitzero" api:"required"`
	// Any of "system", "user", "assistant", "tool".
	Role string `json:"role,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionRequestMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionRequestMessageParam](
		"role", "system", "user", "assistant", "tool",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionRequestMessagesContentUnionParam struct {
	OfString            param.Opt[string]                                                `json:",omitzero,inline"`
	OfTextAndImageArray []ChatCompletionRequestMessagesContentTextAndImageArrayItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionRequestMessagesContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfTextAndImageArray)
}
func (u *ChatCompletionRequestMessagesContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionRequestMessagesContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfTextAndImageArray) {
		return &u.OfTextAndImageArray
	}
	return nil
}

// The property Type is required.
type ChatCompletionRequestMessagesContentTextAndImageArrayItemParam struct {
	// Any of "text", "image_url".
	Type     string            `json:"type,omitzero" api:"required"`
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	Text     param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r ChatCompletionRequestMessagesContentTextAndImageArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestMessagesContentTextAndImageArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestMessagesContentTextAndImageArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionRequestMessagesContentTextAndImageArrayItemParam](
		"type", "text", "image_url",
	)
}

// Use this is you want to guarantee a JSON output without defining a schema. For
// control over the schema, use `guided_json`.
//
// The property Type is required.
type ChatCompletionRequestResponseFormatParam struct {
	// Any of "text", "json_object".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatCompletionRequestResponseFormatParam](
		"type", "text", "json_object",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionRequestStopUnionParam struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionRequestStopUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ChatCompletionRequestStopUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionRequestStopUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type ChatCompletionRequestToolChoice string

const (
	ChatCompletionRequestToolChoiceNone     ChatCompletionRequestToolChoice = "none"
	ChatCompletionRequestToolChoiceAuto     ChatCompletionRequestToolChoice = "auto"
	ChatCompletionRequestToolChoiceRequired ChatCompletionRequestToolChoice = "required"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionRequestToolsUnionParam struct {
	OfFunction  *ChatCompletionRequestToolsFunctionParam  `json:",omitzero,inline"`
	OfRetrieval *ChatCompletionRequestToolsRetrievalParam `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionRequestToolsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunction, u.OfRetrieval)
}
func (u *ChatCompletionRequestToolsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionRequestToolsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFunction) {
		return u.OfFunction
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionRequestToolsUnionParam) GetFunction() *ChatCompletionRequestToolsFunctionFunctionParam {
	if vt := u.OfFunction; vt != nil {
		return &vt.Function
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionRequestToolsUnionParam) GetRetrieval() *BucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionRequestToolsUnionParam) GetType() *string {
	if vt := u.OfFunction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionRequestToolsUnionParam](
		"type",
		apijson.Discriminator[ChatCompletionRequestToolsFunctionParam]("function"),
		apijson.Discriminator[ChatCompletionRequestToolsRetrievalParam]("retrieval"),
	)
}

// The properties Function, Type are required.
type ChatCompletionRequestToolsFunctionParam struct {
	Function ChatCompletionRequestToolsFunctionFunctionParam `json:"function,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type" default:"function"`
	paramObj
}

func (r ChatCompletionRequestToolsFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestToolsFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestToolsFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ChatCompletionRequestToolsFunctionFunctionParam struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Parameters  map[string]any    `json:"parameters,omitzero"`
	paramObj
}

func (r ChatCompletionRequestToolsFunctionFunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestToolsFunctionFunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestToolsFunctionFunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retrieval, Type are required.
type ChatCompletionRequestToolsRetrievalParam struct {
	Retrieval BucketIDsParam `json:"retrieval,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "retrieval".
	Type constant.Retrieval `json:"type" default:"retrieval"`
	paramObj
}

func (r ChatCompletionRequestToolsRetrievalParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestToolsRetrievalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestToolsRetrievalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIChatNewCompletionResponse map[string]any

type AIChatNewCompletionParams struct {
	ChatCompletionRequest ChatCompletionRequestParam
	paramObj
}

func (r AIChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChatCompletionRequest)
}
func (r *AIChatNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
