// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// AgentGatewayService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayService] method instead.
type AgentGatewayService struct {
	Options []option.RequestOption
	// REST APIs on an instance, served to agents as tools.
	APIs AgentGatewayAPIService
	// Checks an instance runs on model and tool requests.
	Guardrails AgentGatewayGuardrailService
	// Keys on an instance.
	Keys AgentGatewayKeyService
	// MCP servers on an instance, served to agents at one MCP endpoint.
	McpServers AgentGatewayMcpServerService
	// Named sets of guardrails an instance applies to requests.
	Policies AgentGatewayPolicyService
	// Where an instance applies each policy.
	PolicyAttachments AgentGatewayPolicyAttachmentService
	// Teams on an instance: shared model lists, budgets, and rate limits.
	Teams AgentGatewayTeamService
}

// NewAgentGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentGatewayService(opts ...option.RequestOption) (r AgentGatewayService) {
	r = AgentGatewayService{}
	r.Options = opts
	r.APIs = NewAgentGatewayAPIService(opts...)
	r.Guardrails = NewAgentGatewayGuardrailService(opts...)
	r.Keys = NewAgentGatewayKeyService(opts...)
	r.McpServers = NewAgentGatewayMcpServerService(opts...)
	r.Policies = NewAgentGatewayPolicyService(opts...)
	r.PolicyAttachments = NewAgentGatewayPolicyAttachmentService(opts...)
	r.Teams = NewAgentGatewayTeamService(opts...)
	return
}

// Creates an Agent Gateway instance on one of the account's WireGuard Interfaces.
// The instance starts `pending`, moves to `provisioning`, and then to `ready` or
// `failed`. It joins the WireGuard Interface as a peer named `agw-<instance id>`,
// and its `gateway_url` is set when it is `ready`.
func (r *AgentGatewayService) New(ctx context.Context, body AgentGatewayNewParams, opts ...option.RequestOption) (res *AgentGatewayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agent_gateways"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns one of the account's Agent Gateway instances.
func (r *AgentGatewayService) Get(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the account's Agent Gateway instances, newest first, one page at a time.
func (r *AgentGatewayService) List(ctx context.Context, query AgentGatewayListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AgentGateway], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "agent_gateways"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the account's Agent Gateway instances, newest first, one page at a time.
func (r *AgentGatewayService) ListAutoPaging(ctx context.Context, query AgentGatewayListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AgentGateway] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an Agent Gateway instance in any status. The response shows it
// `deleting`; repeating the request answers the same until the instance is gone,
// and then `404`. The instance's engine and everything stored on it are removed.
func (r *AgentGatewayService) Delete(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists the models the instance serves. Each has the name a request uses, the
// Telnyx Inference model behind it, and its price per million tokens.
func (r *AgentGatewayService) GetModels(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayGetModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/models", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the requests the instance served with its keys, newest first, one page at
// a time. The window defaults to the last 24 hours and spans at most 31 days. A
// row appears a few seconds after its request. The instance must be `ready`;
// otherwise the request answers `409`.
func (r *AgentGatewayService) GetSpendLogs(ctx context.Context, instanceID string, query AgentGatewayGetSpendLogsParams, opts ...option.RequestOption) (res *AgentGatewayGetSpendLogsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/spend_logs", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// An Agent Gateway instance on a WireGuard Interface.
type AgentGateway struct {
	// Identifies the instance.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the instance was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The gateway engine.
	Engine constant.Litellm `json:"engine" default:"litellm"`
	// The engine version the instance runs.
	EngineVersion string `json:"engine_version" api:"required"`
	// The base URL of the instance's API, `http://<address>:4000`, where the address
	// is the instance's peer address on the WireGuard Interface. Reachable only from
	// hosts on the instance's _Network_, and every request to it carries a key from
	// this API. `null` until the instance is `ready`.
	GatewayURL string `json:"gateway_url" api:"required"`
	// A label for people; `null` when none was given.
	Name string `json:"name" api:"required"`
	// The [_Network_](https://developers.telnyx.com/docs/network/overview) the
	// WireGuard Interface belongs to.
	NetworkID string `json:"network_id" api:"required" format:"uuid"`
	// Where the instance is in its lifecycle. `pending`: created and waiting to start.
	// `provisioning`: starting. `ready`: running, with its `gateway_url` set.
	// `failed`: did not start in time; delete it. `deleting`: being removed.
	//
	// Any of "pending", "provisioning", "ready", "failed", "deleting".
	Status AgentGatewayStatus `json:"status" api:"required"`
	// The WireGuard Interface the instance joins as a peer.
	WireguardInterfaceID string `json:"wireguard_interface_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Engine               respjson.Field
		EngineVersion        respjson.Field
		GatewayURL           respjson.Field
		Name                 respjson.Field
		NetworkID            respjson.Field
		Status               respjson.Field
		WireguardInterfaceID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGateway) RawJSON() string { return r.JSON.raw }
func (r *AgentGateway) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the instance is in its lifecycle. `pending`: created and waiting to start.
// `provisioning`: starting. `ready`: running, with its `gateway_url` set.
// `failed`: did not start in time; delete it. `deleting`: being removed.
type AgentGatewayStatus string

const (
	AgentGatewayStatusPending      AgentGatewayStatus = "pending"
	AgentGatewayStatusProvisioning AgentGatewayStatus = "provisioning"
	AgentGatewayStatusReady        AgentGatewayStatus = "ready"
	AgentGatewayStatusFailed       AgentGatewayStatus = "failed"
	AgentGatewayStatusDeleting     AgentGatewayStatus = "deleting"
)

type AgentGatewayResponse struct {
	// An Agent Gateway instance on a WireGuard Interface.
	Data AgentGateway `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGetModelsResponse struct {
	Data []AgentGatewayGetModelsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model an instance serves.
type AgentGatewayGetModelsResponseData struct {
	// The model's ID on the instance.
	ID string `json:"id" api:"required" format:"uuid"`
	// The Telnyx Inference model behind it, such as `zai-org/GLM-5.3-Flash`.
	Model string `json:"model" api:"required"`
	// The name a request uses, such as `GLM-5.3-Flash`.
	Name string `json:"name" api:"required"`
	// When the price was last set.
	PricedAt time.Time `json:"priced_at" api:"required" format:"date-time"`
	// A model's price in USD per million tokens.
	Pricing AgentGatewayGetModelsResponseDataPricing `json:"pricing" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		PricedAt    respjson.Field
		Pricing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model's price in USD per million tokens.
type AgentGatewayGetModelsResponseDataPricing struct {
	// USD per million cached input tokens.
	CachedInput float64 `json:"cached_input" api:"required"`
	// Always `USD`.
	Currency string `json:"currency" api:"required"`
	// USD per million input tokens.
	Input float64 `json:"input" api:"required"`
	// USD per million output tokens.
	Output float64 `json:"output" api:"required"`
	// Always `1M_tokens`.
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedInput respjson.Field
		Currency    respjson.Field
		Input       respjson.Field
		Output      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetModelsResponseDataPricing) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetModelsResponseDataPricing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGetSpendLogsResponse struct {
	Data []AgentGatewayGetSpendLogsResponseData `json:"data" api:"required"`
	Meta AgentGatewayGetSpendLogsResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetSpendLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetSpendLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One request the instance served.
type AgentGatewayGetSpendLogsResponseData struct {
	// The request's ID.
	ID string `json:"id" api:"required"`
	// Output tokens, reasoning included.
	CompletionTokens int64 `json:"completion_tokens" api:"required"`
	// How long the request took, in milliseconds.
	DurationMs int64 `json:"duration_ms" api:"required"`
	// When the request ended.
	EndedAt time.Time `json:"ended_at" api:"required" format:"date-time"`
	// The key the request used.
	KeyID string `json:"key_id" api:"required"`
	// The Telnyx Inference model that answered, such as `zai-org/GLM-5.3-Flash`. For
	// an MCP tool request, `MCP: ` and the tool's name.
	Model string `json:"model" api:"required"`
	// The request type. `acompletion` is a chat completion, `aresponses` a response,
	// `anthropic_messages` a message, `call_mcp_tool` an MCP tool request, and
	// `list_mcp_tools` an MCP tool list.
	Operation string `json:"operation" api:"required"`
	// Input tokens.
	PromptTokens int64 `json:"prompt_tokens" api:"required"`
	// What the request cost, in USD, with its tokens at Telnyx Inference prices.
	Spend float64 `json:"spend" api:"required"`
	// When the request started.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// `success` or `failure`.
	Status string `json:"status" api:"required"`
	// The key's team.
	TeamID string `json:"team_id" api:"required"`
	// Input and output tokens.
	TotalTokens int64 `json:"total_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletionTokens respjson.Field
		DurationMs       respjson.Field
		EndedAt          respjson.Field
		KeyID            respjson.Field
		Model            respjson.Field
		Operation        respjson.Field
		PromptTokens     respjson.Field
		Spend            respjson.Field
		StartedAt        respjson.Field
		Status           respjson.Field
		TeamID           respjson.Field
		TotalTokens      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetSpendLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetSpendLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGetSpendLogsResponseMeta struct {
	// The page returned, counting from 1.
	PageNumber int64 `json:"page_number" api:"required"`
	// Results per page.
	PageSize int64 `json:"page_size" api:"required"`
	// Pages that can be requested; 0 when nothing matched.
	TotalPages int64 `json:"total_pages" api:"required"`
	// Results across every page.
	TotalResults int64 `json:"total_results" api:"required"`
	// `true` when the window holds more than 10,000 rows. The two totals are then
	// lower bounds: narrow the window, or page until a page comes back empty.
	TotalResultsCapped bool `json:"total_results_capped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber         respjson.Field
		PageSize           respjson.Field
		TotalPages         respjson.Field
		TotalResults       respjson.Field
		TotalResultsCapped respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGetSpendLogsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGetSpendLogsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayNewParams struct {
	// One of the account's WireGuard Interfaces. The instance joins it as a peer, and
	// the interface's other peers reach the instance.
	WireguardInterfaceID string `json:"wireguard_interface_id" api:"required" format:"uuid"`
	// A label for people.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AgentGatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayListParams struct {
	// The page to return, counting from 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Results per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentGatewayListParams]'s query parameters as `url.Values`.
func (r AgentGatewayListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentGatewayGetSpendLogsParams struct {
	// The window's end, an RFC 3339 timestamp with an offset. Defaults to now.
	End param.Opt[time.Time] `query:"end,omitzero" format:"date-time" json:"-"`
	// Only the requests made with this key.
	KeyID param.Opt[string] `query:"key_id,omitzero" json:"-"`
	// The page to return, counting from 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Results per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// The window's start, an RFC 3339 timestamp with an offset. Defaults to 24 hours
	// before `end`.
	Start param.Opt[time.Time] `query:"start,omitzero" format:"date-time" json:"-"`
	// Only the requests made with this team's keys.
	TeamID param.Opt[string] `query:"team_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AgentGatewayGetSpendLogsParams]'s query parameters as
// `url.Values`.
func (r AgentGatewayGetSpendLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
