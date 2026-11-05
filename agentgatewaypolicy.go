// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Named sets of guardrails an instance applies to requests.
//
// AgentGatewayPolicyService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayPolicyService] method instead.
type AgentGatewayPolicyService struct {
	Options []option.RequestOption
}

// NewAgentGatewayPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentGatewayPolicyService(opts ...option.RequestOption) (r AgentGatewayPolicyService) {
	r = AgentGatewayPolicyService{}
	r.Options = opts
	return
}

// Creates a named set of guardrails. It applies from the next request on, to the
// requests that name it in `policies` and to those its
// [attachments](https://developers.telnyx.com/api-reference/agent-gateway-policy-attachments/attach-a-policy)
// match.
func (r *AgentGatewayPolicyService) New(ctx context.Context, instanceID string, body AgentGatewayPolicyNewParams, opts ...option.RequestOption) (res *AgentGatewayPolicyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policies", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns one policy, in its current version.
func (r *AgentGatewayPolicyService) Get(ctx context.Context, name string, query AgentGatewayPolicyGetParams, opts ...option.RequestOption) (res *AgentGatewayPolicyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policies/%s", query.InstanceID, url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Changes a policy, which publishes a new version of it. A field left out stays as
// it is; `guardrails` and `pipeline` are replaced whole, and `model: null` removes
// the condition.
func (r *AgentGatewayPolicyService) Update(ctx context.Context, name string, params AgentGatewayPolicyUpdateParams, opts ...option.RequestOption) (res *AgentGatewayPolicyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policies/%s", params.InstanceID, url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the instance's policies, each in its current version.
func (r *AgentGatewayPolicyService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayPolicyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policies", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a policy and every version of it, and returns it as it was. A policy
// that an attachment names, or that another policy builds on, cannot be deleted.
func (r *AgentGatewayPolicyService) Delete(ctx context.Context, name string, body AgentGatewayPolicyDeleteParams, opts ...option.RequestOption) (res *AgentGatewayPolicyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policies/%s", body.InstanceID, url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A named set of guardrails.
type AgentGatewayPolicy struct {
	// A policy this one builds on, by name. The policy starts with the base's
	// guardrails.
	Base string `json:"base" api:"required"`
	// When the policy's current version was created. An update creates a new version.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// What the policy is for.
	Description string `json:"description" api:"required"`
	// The guardrails a policy adds to its base's, and the ones it removes from them,
	// by name.
	Guardrails AgentGatewayPolicyGuardrails `json:"guardrails" api:"required"`
	// A condition on the model, by name or regular expression. The policy's guardrails
	// run only on requests for a matching model. `null` matches every model.
	Model string `json:"model" api:"required"`
	// A name, unique on the instance. A request names the policy in its `policies`
	// field.
	Name string `json:"name" api:"required"`
	// Ordered steps that run in place of the policy's guardrails running
	// independently.
	Pipeline AgentGatewayPolicyPipeline `json:"pipeline" api:"required"`
	// When the policy last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base        respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Guardrails  respjson.Field
		Model       respjson.Field
		Name        respjson.Field
		Pipeline    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicy) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The guardrails a policy adds to its base's, and the ones it removes from them,
// by name.
type AgentGatewayPolicyGuardrails struct {
	// Guardrails the policy runs in addition to its base's.
	Add []string `json:"add"`
	// Guardrails from its base that the policy does not run.
	Remove []string `json:"remove"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Add         respjson.Field
		Remove      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyGuardrails) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyGuardrails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentGatewayPolicyGuardrails to a
// AgentGatewayPolicyGuardrailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentGatewayPolicyGuardrailsParam.Overrides()
func (r AgentGatewayPolicyGuardrails) ToParam() AgentGatewayPolicyGuardrailsParam {
	return param.Override[AgentGatewayPolicyGuardrailsParam](json.RawMessage(r.RawJSON()))
}

// The guardrails a policy adds to its base's, and the ones it removes from them,
// by name.
type AgentGatewayPolicyGuardrailsParam struct {
	// Guardrails the policy runs in addition to its base's.
	Add []string `json:"add,omitzero"`
	// Guardrails from its base that the policy does not run.
	Remove []string `json:"remove,omitzero"`
	paramObj
}

func (r AgentGatewayPolicyGuardrailsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyGuardrailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyGuardrailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ordered steps that run in place of the policy's guardrails running
// independently.
type AgentGatewayPolicyPipeline struct {
	// `before_request` runs the steps on the request, and `after_request` on the
	// model's response.
	//
	// Any of "before_request", "after_request".
	Hook  AgentGatewayPolicyPipelineHook   `json:"hook" api:"required"`
	Steps []AgentGatewayPolicyPipelineStep `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hook        respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyPipeline) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyPipeline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentGatewayPolicyPipeline to a
// AgentGatewayPolicyPipelineParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentGatewayPolicyPipelineParam.Overrides()
func (r AgentGatewayPolicyPipeline) ToParam() AgentGatewayPolicyPipelineParam {
	return param.Override[AgentGatewayPolicyPipelineParam](json.RawMessage(r.RawJSON()))
}

// `before_request` runs the steps on the request, and `after_request` on the
// model's response.
type AgentGatewayPolicyPipelineHook string

const (
	AgentGatewayPolicyPipelineHookBeforeRequest AgentGatewayPolicyPipelineHook = "before_request"
	AgentGatewayPolicyPipelineHookAfterRequest  AgentGatewayPolicyPipelineHook = "after_request"
)

// One pipeline step: a guardrail, and what happens when it passes, fails, or
// errors.
type AgentGatewayPolicyPipelineStep struct {
	// A guardrail on the instance, by name.
	Guardrail string `json:"guardrail" api:"required"`
	// The response that replaces the model's when an outcome is `respond`. Required
	// with `respond`.
	Message string `json:"message" api:"nullable"`
	// What happens when the guardrail errors. Left out, the same as `on_fail`.
	//
	// Any of "next", "allow", "block", "respond".
	OnError string `json:"on_error" api:"nullable"`
	// What happens when the guardrail fails.
	//
	// Any of "next", "allow", "block", "respond".
	OnFail string `json:"on_fail"`
	// What happens when the guardrail passes. `next` goes on to the next step, `allow`
	// allows the request, `block` blocks it, and `respond` answers with `message`.
	//
	// Any of "next", "allow", "block", "respond".
	OnPass string `json:"on_pass"`
	// `true` hands the step's modified request, such as masked text, to the next step.
	PassData bool `json:"pass_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guardrail   respjson.Field
		Message     respjson.Field
		OnError     respjson.Field
		OnFail      respjson.Field
		OnPass      respjson.Field
		PassData    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyPipelineStep) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyPipelineStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ordered steps that run in place of the policy's guardrails running
// independently.
//
// The properties Hook, Steps are required.
type AgentGatewayPolicyPipelineParam struct {
	// `before_request` runs the steps on the request, and `after_request` on the
	// model's response.
	//
	// Any of "before_request", "after_request".
	Hook  AgentGatewayPolicyPipelineHook        `json:"hook,omitzero" api:"required"`
	Steps []AgentGatewayPolicyPipelineStepParam `json:"steps,omitzero" api:"required"`
	paramObj
}

func (r AgentGatewayPolicyPipelineParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyPipelineParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyPipelineParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One pipeline step: a guardrail, and what happens when it passes, fails, or
// errors.
//
// The property Guardrail is required.
type AgentGatewayPolicyPipelineStepParam struct {
	// A guardrail on the instance, by name.
	Guardrail string `json:"guardrail" api:"required"`
	// The response that replaces the model's when an outcome is `respond`. Required
	// with `respond`.
	Message param.Opt[string] `json:"message,omitzero"`
	// `true` hands the step's modified request, such as masked text, to the next step.
	PassData param.Opt[bool] `json:"pass_data,omitzero"`
	// What happens when the guardrail errors. Left out, the same as `on_fail`.
	//
	// Any of "next", "allow", "block", "respond".
	OnError string `json:"on_error,omitzero"`
	// What happens when the guardrail fails.
	//
	// Any of "next", "allow", "block", "respond".
	OnFail string `json:"on_fail,omitzero"`
	// What happens when the guardrail passes. `next` goes on to the next step, `allow`
	// allows the request, `block` blocks it, and `respond` answers with `message`.
	//
	// Any of "next", "allow", "block", "respond".
	OnPass string `json:"on_pass,omitzero"`
	paramObj
}

func (r AgentGatewayPolicyPipelineStepParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyPipelineStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyPipelineStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentGatewayPolicyPipelineStepParam](
		"on_error", "next", "allow", "block", "respond",
	)
	apijson.RegisterFieldValidator[AgentGatewayPolicyPipelineStepParam](
		"on_fail", "next", "allow", "block", "respond",
	)
	apijson.RegisterFieldValidator[AgentGatewayPolicyPipelineStepParam](
		"on_pass", "next", "allow", "block", "respond",
	)
}

type AgentGatewayPolicyResponse struct {
	// A named set of guardrails.
	Data AgentGatewayPolicy `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyListResponse struct {
	Data []AgentGatewayPolicy `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyNewParams struct {
	// A name, unique on the instance. A request names the policy in its `policies`
	// field.
	Name string `json:"name" api:"required"`
	// A policy this one builds on, by name. The policy starts with the base's
	// guardrails.
	Base param.Opt[string] `json:"base,omitzero"`
	// What the policy is for.
	Description param.Opt[string] `json:"description,omitzero"`
	// A condition on the model, by name or regular expression. The policy's guardrails
	// run only on requests for a matching model. `null` matches every model.
	Model param.Opt[string] `json:"model,omitzero"`
	// The guardrails a policy adds to its base's, and the ones it removes from them,
	// by name.
	Guardrails AgentGatewayPolicyGuardrailsParam `json:"guardrails,omitzero"`
	// Ordered steps that run in place of the policy's guardrails running
	// independently.
	Pipeline AgentGatewayPolicyPipelineParam `json:"pipeline,omitzero"`
	paramObj
}

func (r AgentGatewayPolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyGetParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AgentGatewayPolicyUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// A policy this one builds on, by name. The policy starts with the base's
	// guardrails.
	Base param.Opt[string] `json:"base,omitzero"`
	// What the policy is for.
	Description param.Opt[string] `json:"description,omitzero"`
	// A condition on the model, by name or regular expression. The policy's guardrails
	// run only on requests for a matching model. `null` matches every model.
	Model param.Opt[string] `json:"model,omitzero"`
	// The guardrails a policy adds to its base's, and the ones it removes from them,
	// by name.
	Guardrails AgentGatewayPolicyGuardrailsParam `json:"guardrails,omitzero"`
	// Ordered steps that run in place of the policy's guardrails running
	// independently.
	Pipeline AgentGatewayPolicyPipelineParam `json:"pipeline,omitzero"`
	paramObj
}

func (r AgentGatewayPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
