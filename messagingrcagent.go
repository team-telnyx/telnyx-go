// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// Send RCS messages
//
// MessagingRcAgentService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingRcAgentService] method instead.
type MessagingRcAgentService struct {
	Options []option.RequestOption
}

// NewMessagingRcAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingRcAgentService(opts ...option.RequestOption) (r MessagingRcAgentService) {
	r = MessagingRcAgentService{}
	r.Options = opts
	return
}

// Retrieve an RCS agent
func (r *MessagingRcAgentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RcsAgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging/rcs/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify an RCS agent
func (r *MessagingRcAgentService) Update(ctx context.Context, id string, body MessagingRcAgentUpdateParams, opts ...option.RequestOption) (res *RcsAgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging/rcs/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all RCS agents
func (r *MessagingRcAgentService) List(ctx context.Context, query MessagingRcAgentListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[RcsAgent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging/rcs/agents"
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

// List all RCS agents
func (r *MessagingRcAgentService) ListAutoPaging(ctx context.Context, query MessagingRcAgentListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[RcsAgent] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type MessagingRcAgentUpdateParams struct {
	// Messaging profile ID associated with the RCS Agent
	ProfileID param.Opt[string] `json:"profile_id,omitzero" format:"uuid"`
	// Failover URL to receive RCS events
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// URL to receive RCS events
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	paramObj
}

func (r MessagingRcAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingRcAgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingRcAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcAgentListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingRcAgentListParams]'s query parameters as
// `url.Values`.
func (r MessagingRcAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
