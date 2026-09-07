// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// AIAnthropicV1Service contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAnthropicV1Service] method instead.
type AIAnthropicV1Service struct {
	Options []option.RequestOption
}

// NewAIAnthropicV1Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAnthropicV1Service(opts ...option.RequestOption) (r AIAnthropicV1Service) {
	r = AIAnthropicV1Service{}
	r.Options = opts
	return
}

// Send a message to a language model using the Anthropic Messages API format. This
// endpoint is compatible with the
// [Anthropic Messages API](https://docs.anthropic.com/en/api/messages) and may be
// used with the Anthropic JS or Python SDK by setting the base URL to
// `https://api.telnyx.com/v2/ai/anthropic`.
//
// The endpoint translates Anthropic-format requests into Telnyx's inference
// internals, then translates the response back to the Anthropic message shape.
// Streaming responses use Anthropic SSE event types (`message_start`,
// `content_block_start`, `content_block_delta`, `content_block_stop`,
// `message_delta`, `message_stop`).
func (r *AIAnthropicV1Service) Messages(ctx context.Context, body AIAnthropicV1MessagesParams, opts ...option.RequestOption) (res *AIAnthropicV1MessagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/anthropic/v1/messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AIAnthropicV1MessagesResponse map[string]any

type AIAnthropicV1MessagesParams struct {
	// The maximum number of tokens to generate in the response.
	MaxTokens int64 `json:"max_tokens" api:"required"`
	// The messages to send to the model, following the
	// [Anthropic Messages API](https://docs.anthropic.com/en/api/messages) format.
	Messages []map[string]any `json:"messages,omitzero" api:"required"`
	// The model to use for generating the response, for example
	// `zai-org/GLM-5.3-Flash` or another model available from the Telnyx models
	// endpoint.
	Model string `json:"model" api:"required"`
	// If you are using an external inference provider, this field allows you to pass
	// along a reference to your API key. After creating an
	// [integration secret](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for your API key, pass the secret's `identifier` in this field.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// The billing group ID to associate with this request.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero" format:"uuid"`
	// Maximum number of retries for the request.
	MaxRetries param.Opt[int64] `json:"max_retries,omitzero"`
	// The service tier to use for this request. Supported values vary by model; use
	// the Telnyx models endpoint and inspect the model's `service_tiers` field. If
	// omitted, Telnyx-hosted models use `default`.
	ServiceTier param.Opt[string] `json:"service_tier,omitzero"`
	// Whether to stream the response as Anthropic-format Server-Sent Events.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Amount of randomness injected into the response. Ranges from 0 to 1.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Request timeout in seconds.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Top-k sampling parameter. Only sample from the top K options for each subsequent
	// token.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Nucleus sampling parameter. Use temperature or top_p, but not both.
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Configuration for model fallback behavior when the primary model is unavailable.
	FallbackConfig map[string]any `json:"fallback_config,omitzero"`
	// List of MCP (Model Context Protocol) servers to make available to the model.
	McpServers []map[string]any `json:"mcp_servers,omitzero"`
	// An object describing metadata about the request.
	Metadata map[string]any `json:"metadata,omitzero"`
	// How strictly `region` is applied. `preferred` (the default when `region` is set)
	// tries that region first and falls back to another when the model cannot be
	// served there, so a request that would have succeeded still succeeds. `strict`
	// pins the request: it is served from that region or it fails with a 422, never
	// redirected to another region. Requires `region`.
	//
	// Any of "preferred", "strict".
	Mode AIAnthropicV1MessagesParamsMode `json:"mode,omitzero"`
	// Optional data-residency region the request should be served from, using the same
	// vocabulary as your account's Data Locality setting. Behavior depends on `mode`.
	// Supported for Telnyx-hosted models only: a request routed to an external
	// provider never passes through Telnyx model routing, so a region cannot be
	// enforced for it. Omit for today's latency-based routing.
	//
	// Any of "USA", "EU", "AUS", "UAE".
	Region AIAnthropicV1MessagesParamsRegion `json:"region,omitzero"`
	// Custom sequences that will cause the model to stop generating.
	StopSequences []string `json:"stop_sequences,omitzero"`
	// System prompt. Can be a string or an array of content blocks following the
	// Anthropic API format.
	System AIAnthropicV1MessagesParamsSystemUnion `json:"system,omitzero"`
	// Extended thinking configuration for models that support it. Set `type` to
	// `enabled` to turn on extended thinking.
	Thinking map[string]any `json:"thinking,omitzero"`
	// Controls how the model uses tools, following the Anthropic API format.
	ToolChoice map[string]any `json:"tool_choice,omitzero"`
	// Definitions of tools that the model may use, following the Anthropic API format.
	Tools []map[string]any `json:"tools,omitzero"`
	paramObj
}

func (r AIAnthropicV1MessagesParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAnthropicV1MessagesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAnthropicV1MessagesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How strictly `region` is applied. `preferred` (the default when `region` is set)
// tries that region first and falls back to another when the model cannot be
// served there, so a request that would have succeeded still succeeds. `strict`
// pins the request: it is served from that region or it fails with a 422, never
// redirected to another region. Requires `region`.
type AIAnthropicV1MessagesParamsMode string

const (
	AIAnthropicV1MessagesParamsModePreferred AIAnthropicV1MessagesParamsMode = "preferred"
	AIAnthropicV1MessagesParamsModeStrict    AIAnthropicV1MessagesParamsMode = "strict"
)

// Optional data-residency region the request should be served from, using the same
// vocabulary as your account's Data Locality setting. Behavior depends on `mode`.
// Supported for Telnyx-hosted models only: a request routed to an external
// provider never passes through Telnyx model routing, so a region cannot be
// enforced for it. Omit for today's latency-based routing.
type AIAnthropicV1MessagesParamsRegion string

const (
	AIAnthropicV1MessagesParamsRegionUsa AIAnthropicV1MessagesParamsRegion = "USA"
	AIAnthropicV1MessagesParamsRegionEu  AIAnthropicV1MessagesParamsRegion = "EU"
	AIAnthropicV1MessagesParamsRegionAus AIAnthropicV1MessagesParamsRegion = "AUS"
	AIAnthropicV1MessagesParamsRegionUae AIAnthropicV1MessagesParamsRegion = "UAE"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIAnthropicV1MessagesParamsSystemUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfMapOfAnyMap []map[string]any  `json:",omitzero,inline"`
	paramUnion
}

func (u AIAnthropicV1MessagesParamsSystemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfMapOfAnyMap)
}
func (u *AIAnthropicV1MessagesParamsSystemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIAnthropicV1MessagesParamsSystemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	}
	return nil
}
