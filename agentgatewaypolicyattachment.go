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

// Where an instance applies each policy.
//
// AgentGatewayPolicyAttachmentService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayPolicyAttachmentService] method instead.
type AgentGatewayPolicyAttachmentService struct {
	Options []option.RequestOption
}

// NewAgentGatewayPolicyAttachmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAgentGatewayPolicyAttachmentService(opts ...option.RequestOption) (r AgentGatewayPolicyAttachmentService) {
	r = AgentGatewayPolicyAttachmentService{}
	r.Options = opts
	return
}

// Applies a policy to every request, or to the requests for the models named. It
// applies from the next request on.
func (r *AgentGatewayPolicyAttachmentService) New(ctx context.Context, instanceID string, body AgentGatewayPolicyAttachmentNewParams, opts ...option.RequestOption) (res *AgentGatewayPolicyAttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policy_attachments", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists where the instance applies each policy.
func (r *AgentGatewayPolicyAttachmentService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayPolicyAttachmentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policy_attachments", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes an attachment and returns it as it was. The policy stays.
func (r *AgentGatewayPolicyAttachmentService) Delete(ctx context.Context, attachmentID string, body AgentGatewayPolicyAttachmentDeleteParams, opts ...option.RequestOption) (res *AgentGatewayPolicyAttachmentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if attachmentID == "" {
		err = errors.New("missing required attachment_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/policy_attachments/%s", body.InstanceID, attachmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Where an instance applies a policy.
type AgentGatewayPolicyAttachment struct {
	// The attachment's ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the attachment was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// `true` applies the policy only to the requests no other attachment matches.
	Default bool `json:"default" api:"required"`
	// The models whose requests the policy applies to, by name. A name can use `*` as
	// a wildcard. Empty applies the policy to every request.
	Models []string `json:"models" api:"required"`
	// The policy to apply, by name.
	Policy string `json:"policy" api:"required"`
	// An attachment with a priority applies before the others.
	Priority int64 `json:"priority" api:"required"`
	// When the attachment last changed.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Default     respjson.Field
		Models      respjson.Field
		Policy      respjson.Field
		Priority    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyAttachment) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyAttachmentResponse struct {
	// Where an instance applies a policy.
	Data AgentGatewayPolicyAttachment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyAttachmentResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyAttachmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyAttachmentListResponse struct {
	Data []AgentGatewayPolicyAttachment `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayPolicyAttachmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayPolicyAttachmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyAttachmentNewParams struct {
	// The policy to apply, by name.
	Policy string `json:"policy" api:"required"`
	// An attachment with a priority applies before the others.
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// `true` applies the policy only to the requests no other attachment matches.
	Default param.Opt[bool] `json:"default,omitzero"`
	// The models whose requests the policy applies to, by name. A name can use `*` as
	// a wildcard. Empty applies the policy to every request.
	Models []string `json:"models,omitzero"`
	paramObj
}

func (r AgentGatewayPolicyAttachmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPolicyAttachmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPolicyAttachmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPolicyAttachmentDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
