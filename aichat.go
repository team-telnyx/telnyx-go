// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"encoding/json"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
	"github.com/tidwall/gjson"
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
	// The service tier to use for this request. Supported values vary by model; use
	// `GET /v2/ai/openai/models` and inspect the model's `service_tiers` field. If
	// omitted, Telnyx-hosted models use `default`.
	ServiceTier param.Opt[string] `json:"service_tier,omitzero"`
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
	// How strictly `region` is applied. `preferred` (the default when `region` is set)
	// tries that region first and falls back to another when the model cannot be
	// served there, so a request that would have succeeded still succeeds. `strict`
	// pins the request: it is served from that region or it fails with a 422, never
	// redirected to another region. Requires `region`.
	//
	// Any of "preferred", "strict".
	Mode ChatCompletionRequestMode `json:"mode,omitzero"`
	// Controls the reasoning effort for models that support it. When set, the model
	// spends more or less compute on internal reasoning before generating its
	// response. Supported values: none, minimal, low, medium, high, xhigh, max. Not
	// all models support all values; unsupported values are rejected with a 400 error.
	// When omitted, reasoning models use their default effort level.
	//
	// Any of "none", "minimal", "low", "medium", "high", "xhigh", "max".
	ReasoningEffort ChatCompletionRequestReasoningEffort `json:"reasoning_effort,omitzero"`
	// Optional data-residency region the request should be served from, using the same
	// vocabulary as your account's Data Locality setting. Behavior depends on `mode`.
	// Supported for Telnyx-hosted models only: a request routed to an external
	// provider never passes through Telnyx model routing, so a region cannot be
	// enforced for it. Omit for today's latency-based routing.
	//
	// Any of "USA", "EU", "AUS", "UAE".
	Region ChatCompletionRequestRegion `json:"region,omitzero"`
	// Controls the format of the model output. `json_object` guarantees valid JSON
	// output without defining a schema; `json_schema` constrains the output to the
	// JSON schema you supply via the `json_schema` property and is the supported way
	// to get guaranteed structured output on Telnyx-hosted models.
	ResponseFormat ChatCompletionRequestResponseFormatUnionParam `json:"response_format,omitzero"`
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

// How strictly `region` is applied. `preferred` (the default when `region` is set)
// tries that region first and falls back to another when the model cannot be
// served there, so a request that would have succeeded still succeeds. `strict`
// pins the request: it is served from that region or it fails with a 422, never
// redirected to another region. Requires `region`.
type ChatCompletionRequestMode string

const (
	ChatCompletionRequestModePreferred ChatCompletionRequestMode = "preferred"
	ChatCompletionRequestModeStrict    ChatCompletionRequestMode = "strict"
)

// Controls the reasoning effort for models that support it. When set, the model
// spends more or less compute on internal reasoning before generating its
// response. Supported values: none, minimal, low, medium, high, xhigh, max. Not
// all models support all values; unsupported values are rejected with a 400 error.
// When omitted, reasoning models use their default effort level.
type ChatCompletionRequestReasoningEffort string

const (
	ChatCompletionRequestReasoningEffortNone    ChatCompletionRequestReasoningEffort = "none"
	ChatCompletionRequestReasoningEffortMinimal ChatCompletionRequestReasoningEffort = "minimal"
	ChatCompletionRequestReasoningEffortLow     ChatCompletionRequestReasoningEffort = "low"
	ChatCompletionRequestReasoningEffortMedium  ChatCompletionRequestReasoningEffort = "medium"
	ChatCompletionRequestReasoningEffortHigh    ChatCompletionRequestReasoningEffort = "high"
	ChatCompletionRequestReasoningEffortXhigh   ChatCompletionRequestReasoningEffort = "xhigh"
	ChatCompletionRequestReasoningEffortMax     ChatCompletionRequestReasoningEffort = "max"
)

// Optional data-residency region the request should be served from, using the same
// vocabulary as your account's Data Locality setting. Behavior depends on `mode`.
// Supported for Telnyx-hosted models only: a request routed to an external
// provider never passes through Telnyx model routing, so a region cannot be
// enforced for it. Omit for today's latency-based routing.
type ChatCompletionRequestRegion string

const (
	ChatCompletionRequestRegionUsa ChatCompletionRequestRegion = "USA"
	ChatCompletionRequestRegionEu  ChatCompletionRequestRegion = "EU"
	ChatCompletionRequestRegionAus ChatCompletionRequestRegion = "AUS"
	ChatCompletionRequestRegionUae ChatCompletionRequestRegion = "UAE"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatCompletionRequestResponseFormatUnionParam struct {
	OfResponseFormatText       *ChatCompletionRequestResponseFormatResponseFormatTextParam            `json:",omitzero,inline"`
	OfResponseFormatJsonObject *ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam      `json:",omitzero,inline"`
	OfResponseFormatJsonSchema *ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam `json:",omitzero,inline"`
	paramUnion
}

func (u ChatCompletionRequestResponseFormatUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfResponseFormatText, u.OfResponseFormatJsonObject, u.OfResponseFormatJsonSchema)
}
func (u *ChatCompletionRequestResponseFormatUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatCompletionRequestResponseFormatUnionParam) asAny() any {
	if !param.IsOmitted(u.OfResponseFormatText) {
		return u.OfResponseFormatText
	} else if !param.IsOmitted(u.OfResponseFormatJsonObject) {
		return u.OfResponseFormatJsonObject
	} else if !param.IsOmitted(u.OfResponseFormatJsonSchema) {
		return u.OfResponseFormatJsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionRequestResponseFormatUnionParam) GetJsonSchema() *ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam {
	if vt := u.OfResponseFormatJsonSchema; vt != nil {
		return &vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ChatCompletionRequestResponseFormatUnionParam) GetType() *string {
	if vt := u.OfResponseFormatText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfResponseFormatJsonObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfResponseFormatJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ChatCompletionRequestResponseFormatUnionParam](
		"",
		apijson.Variant[ChatCompletionRequestResponseFormatResponseFormatTextParam](gjson.JSON),
		apijson.Variant[ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam](gjson.JSON),
		apijson.Variant[ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam](gjson.JSON),
	)
}

func NewChatCompletionRequestResponseFormatResponseFormatTextParam() ChatCompletionRequestResponseFormatResponseFormatTextParam {
	return ChatCompletionRequestResponseFormatResponseFormatTextParam{
		Type: "text",
	}
}

// Plain text output.
//
// This struct has a constant value, construct it with
// [NewChatCompletionRequestResponseFormatResponseFormatTextParam].
type ChatCompletionRequestResponseFormatResponseFormatTextParam struct {
	Type constant.Text `json:"type" default:"text"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatResponseFormatTextParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatResponseFormatTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatResponseFormatTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewChatCompletionRequestResponseFormatResponseFormatJsonObjectParam() ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam {
	return ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam{
		Type: "json_object",
	}
}

// JSON mode: the model output is valid JSON, without a schema.
//
// This struct has a constant value, construct it with
// [NewChatCompletionRequestResponseFormatResponseFormatJsonObjectParam].
type ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam struct {
	Type constant.JsonObject `json:"type" default:"json_object"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatResponseFormatJsonObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured output: the model output is constrained to the JSON schema supplied
// in `json_schema`.
//
// The properties JsonSchema, Type are required.
type ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam struct {
	// The JSON schema configuration, required when `type` is `json_schema`. Matches
	// the
	// [OpenAI structured outputs](https://platform.openai.com/docs/guides/structured-outputs)
	// `json_schema` response format.
	JsonSchema ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam `json:"json_schema,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type" default:"json_schema"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The JSON schema configuration, required when `type` is `json_schema`. Matches
// the
// [OpenAI structured outputs](https://platform.openai.com/docs/guides/structured-outputs)
// `json_schema` response format.
//
// The property Name is required.
type ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam struct {
	// The name of the response format. Used for clarity only.
	Name string `json:"name" api:"required"`
	// A description of what the response format is for, typically used to guide the
	// model.
	Description param.Opt[string] `json:"description,omitzero"`
	// Enables strict schema adherence when supported by the model. If the generated
	// output does not match the provided schema, the request fails instead of
	// returning non-conformant output.
	Strict param.Opt[bool] `json:"strict,omitzero"`
	// The JSON schema the model output must conform to. A valid
	// [JSON Schema](https://json-schema.org) object, e.g. a Pydantic
	// `model_json_schema()` export.
	Schema map[string]any `json:"schema,omitzero"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatResponseFormatJsonSchemaParamJsonSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
