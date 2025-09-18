// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging/rcs/agents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all RCS agents
func (r *MessagingRcAgentService) List(ctx context.Context, query MessagingRcAgentListParams, opts ...option.RequestOption) (res *MessagingRcAgentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "messaging/rcs/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MessagingRcAgentListResponse struct {
	Data []RcsAgent     `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingRcAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingRcAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingRcAgentListParamsPage `query:"page,omitzero" json:"-"`
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

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingRcAgentListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingRcAgentListParamsPage]'s query parameters as
// `url.Values`.
func (r MessagingRcAgentListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
