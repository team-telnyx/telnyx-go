// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Checks an instance runs on model and tool requests.
//
// AgentGatewayGuardrailService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayGuardrailService] method instead.
type AgentGatewayGuardrailService struct {
	Options []option.RequestOption
}

// NewAgentGatewayGuardrailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentGatewayGuardrailService(opts ...option.RequestOption) (r AgentGatewayGuardrailService) {
	r = AgentGatewayGuardrailService{}
	r.Options = opts
	return
}

// Adds a check the instance runs on model and tool requests. It runs on every
// request when `always_on`, and otherwise on the requests that name it in
// `guardrails`. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayGuardrailService) New(ctx context.Context, instanceID string, body AgentGatewayGuardrailNewParams, opts ...option.RequestOption) (res *AgentGatewayGuardrailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/guardrails", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes a guardrail's `name`, `modes`, `always_on`, or `settings`. `settings` is
// replaced whole, secrets included. The `type` cannot change. The instance must be
// `ready`; otherwise the request answers `409`.
func (r *AgentGatewayGuardrailService) Update(ctx context.Context, guardrailID string, params AgentGatewayGuardrailUpdateParams, opts ...option.RequestOption) (res *AgentGatewayGuardrailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if guardrailID == "" {
		err = errors.New("missing required guardrail_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/guardrails/%s", params.InstanceID, guardrailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the instance's guardrails, without their secrets. The instance must be
// `ready`; otherwise the request answers `409`.
func (r *AgentGatewayGuardrailService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayGuardrailListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/guardrails", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes a guardrail and returns it as it was. The instance must be `ready`;
// otherwise the request answers `409`.
func (r *AgentGatewayGuardrailService) Delete(ctx context.Context, guardrailID string, body AgentGatewayGuardrailDeleteParams, opts ...option.RequestOption) (res *AgentGatewayGuardrailResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if guardrailID == "" {
		err = errors.New("missing required guardrail_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/guardrails/%s", body.InstanceID, guardrailID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AgentGatewayGuardrail struct {
	// Identifies the guardrail.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether it runs on every request.
	AlwaysOn  bool      `json:"always_on" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the guardrail runs.
	//
	// Any of "before_request", "during_request", "after_request", "logging_only",
	// "before_tool_request", "during_tool_request", "after_tool_request".
	Modes []string `json:"modes" api:"required"`
	// What a request names in its `guardrails` field.
	Name string `json:"name" api:"required"`
	// The type's settings. `content_filter`: `patterns` (prebuilt patterns by name,
	// such as `us_ssn` or `email`), `regexes`, `keywords`, and `categories` (such as
	// `prompt_injection`), each with an `action` of `block` or `mask`.
	// `tool_permission`: `rules` (`id`, `tool_name`, `tool_type`, `decision`,
	// `argument_patterns`), a `default` decision, and `on_denied` (`block` or
	// `rewrite`). `model_judge`: a `model` on the instance and weighted `criteria`
	// (`name`, `description`, `weight`, the weights summing to 100). `custom_code`:
	// `code` defining `apply_guardrail(inputs, request_data, input_type)`. `lakera`:
	// `api_key`, `project_id`, `api_base`. `presidio`: `analyzer_url`,
	// `anonymizer_url`, `entities` (entity name to `block` or `mask`), `language`.
	// `bedrock`: `guardrail_id`, `version`, `region`, and `access_key_id` with
	// `secret_access_key`, or `role_name`. `azure_content_safety`: `api_key`,
	// `endpoint`, `severity` (0 to 7). Provider secrets are write-only. Responses
	// leave the secrets out.
	Settings map[string]any `json:"settings" api:"required"`
	// Any of "content_filter", "tool_permission", "model_judge", "custom_code",
	// "lakera", "presidio", "bedrock", "azure_content_safety".
	Type      AgentGatewayGuardrailType `json:"type" api:"required"`
	UpdatedAt time.Time                 `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AlwaysOn    respjson.Field
		CreatedAt   respjson.Field
		Modes       respjson.Field
		Name        respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGuardrail) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGuardrail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGuardrailType string

const (
	AgentGatewayGuardrailTypeContentFilter      AgentGatewayGuardrailType = "content_filter"
	AgentGatewayGuardrailTypeToolPermission     AgentGatewayGuardrailType = "tool_permission"
	AgentGatewayGuardrailTypeModelJudge         AgentGatewayGuardrailType = "model_judge"
	AgentGatewayGuardrailTypeCustomCode         AgentGatewayGuardrailType = "custom_code"
	AgentGatewayGuardrailTypeLakera             AgentGatewayGuardrailType = "lakera"
	AgentGatewayGuardrailTypePresidio           AgentGatewayGuardrailType = "presidio"
	AgentGatewayGuardrailTypeBedrock            AgentGatewayGuardrailType = "bedrock"
	AgentGatewayGuardrailTypeAzureContentSafety AgentGatewayGuardrailType = "azure_content_safety"
)

type AgentGatewayGuardrailResponse struct {
	Data AgentGatewayGuardrail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGuardrailResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGuardrailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGuardrailListResponse struct {
	Data []AgentGatewayGuardrail `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayGuardrailListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayGuardrailListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGuardrailNewParams struct {
	// When the guardrail runs. Each type runs in some modes only.
	//
	// Any of "before_request", "during_request", "after_request", "logging_only",
	// "before_tool_request", "during_tool_request", "after_tool_request".
	Modes []string `json:"modes,omitzero" api:"required"`
	// A name, unique on the instance. A request names it in its `guardrails` field.
	Name string `json:"name" api:"required"`
	// The kind of check. It cannot change after create.
	//
	// Any of "content_filter", "tool_permission", "model_judge", "custom_code",
	// "lakera", "presidio", "bedrock", "azure_content_safety".
	Type AgentGatewayGuardrailNewParamsType `json:"type,omitzero" api:"required"`
	// `true` runs the guardrail on every request, and a request cannot turn it off.
	// `false` runs it on the requests that name it.
	AlwaysOn param.Opt[bool] `json:"always_on,omitzero"`
	// The type's settings. `content_filter`: `patterns` (prebuilt patterns by `name`,
	// such as `us_ssn` or `email`), `regexes` (each a `name`, which tags masked
	// matches, and a `pattern`), `keywords` (each a `keyword`), and `categories` (each
	// a `name`), each entry with an `action` of `block` or `mask`. `tool_permission`:
	// `rules` (`id`, `tool_name`, `tool_type`, `decision`, `argument_patterns`), a
	// `default` decision, and `on_denied` (`block` or `rewrite`). `model_judge`: a
	// `model` on the instance and weighted `criteria` (`name`, `description`,
	// `weight`, the weights summing to 100). `custom_code`: `code` defining
	// `apply_guardrail(inputs, request_data, input_type)`. `lakera`: `api_key`,
	// `project_id`, `api_base`. `presidio`: `analyzer_url`, `anonymizer_url`,
	// `entities` (entity name to `block` or `mask`), `language`. `bedrock`:
	// `guardrail_id`, `version`, `region`, and `access_key_id` with
	// `secret_access_key`, or `role_name`. `azure_content_safety`: `api_key`,
	// `endpoint`, `severity` (0 to 7). Provider secrets are write-only.
	Settings map[string]any `json:"settings,omitzero"`
	paramObj
}

func (r AgentGatewayGuardrailNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayGuardrailNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayGuardrailNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of check. It cannot change after create.
type AgentGatewayGuardrailNewParamsType string

const (
	AgentGatewayGuardrailNewParamsTypeContentFilter      AgentGatewayGuardrailNewParamsType = "content_filter"
	AgentGatewayGuardrailNewParamsTypeToolPermission     AgentGatewayGuardrailNewParamsType = "tool_permission"
	AgentGatewayGuardrailNewParamsTypeModelJudge         AgentGatewayGuardrailNewParamsType = "model_judge"
	AgentGatewayGuardrailNewParamsTypeCustomCode         AgentGatewayGuardrailNewParamsType = "custom_code"
	AgentGatewayGuardrailNewParamsTypeLakera             AgentGatewayGuardrailNewParamsType = "lakera"
	AgentGatewayGuardrailNewParamsTypePresidio           AgentGatewayGuardrailNewParamsType = "presidio"
	AgentGatewayGuardrailNewParamsTypeBedrock            AgentGatewayGuardrailNewParamsType = "bedrock"
	AgentGatewayGuardrailNewParamsTypeAzureContentSafety AgentGatewayGuardrailNewParamsType = "azure_content_safety"
)

type AgentGatewayGuardrailUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// `true` runs the guardrail on every request, and a request cannot turn it off.
	// `false` runs it on the requests that name it.
	AlwaysOn param.Opt[bool] `json:"always_on,omitzero"`
	// A name, unique on the instance. A request names it in its `guardrails` field.
	Name param.Opt[string] `json:"name,omitzero"`
	// When the guardrail runs. Each type runs in some modes only.
	//
	// Any of "before_request", "during_request", "after_request", "logging_only",
	// "before_tool_request", "during_tool_request", "after_tool_request".
	Modes []string `json:"modes,omitzero"`
	// The type's settings. `content_filter`: `patterns` (prebuilt patterns by name,
	// such as `us_ssn` or `email`), `regexes`, `keywords`, and `categories` (such as
	// `prompt_injection`), each with an `action` of `block` or `mask`.
	// `tool_permission`: `rules` (`id`, `tool_name`, `tool_type`, `decision`,
	// `argument_patterns`), a `default` decision, and `on_denied` (`block` or
	// `rewrite`). `model_judge`: a `model` on the instance and weighted `criteria`
	// (`name`, `description`, `weight`, the weights summing to 100). `custom_code`:
	// `code` defining `apply_guardrail(inputs, request_data, input_type)`. `lakera`:
	// `api_key`, `project_id`, `api_base`. `presidio`: `analyzer_url`,
	// `anonymizer_url`, `entities` (entity name to `block` or `mask`), `language`.
	// `bedrock`: `guardrail_id`, `version`, `region`, and `access_key_id` with
	// `secret_access_key`, or `role_name`. `azure_content_safety`: `api_key`,
	// `endpoint`, `severity` (0 to 7). Provider secrets are write-only.
	Settings map[string]any `json:"settings,omitzero"`
	paramObj
}

func (r AgentGatewayGuardrailUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayGuardrailUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayGuardrailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayGuardrailDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
