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
)

// AIConversationInsightGroupService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationInsightGroupService] method instead.
type AIConversationInsightGroupService struct {
	Options  []option.RequestOption
	Insights AIConversationInsightGroupInsightService
}

// NewAIConversationInsightGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIConversationInsightGroupService(opts ...option.RequestOption) (r AIConversationInsightGroupService) {
	r = AIConversationInsightGroupService{}
	r.Options = opts
	r.Insights = NewAIConversationInsightGroupInsightService(opts...)
	return
}

// Get insight group by ID
func (r *AIConversationInsightGroupService) Get(ctx context.Context, groupID string, opts ...option.RequestOption) (res *InsightTemplateGroupDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insight-groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an insight template group
func (r *AIConversationInsightGroupService) Update(ctx context.Context, groupID string, body AIConversationInsightGroupUpdateParams, opts ...option.RequestOption) (res *InsightTemplateGroupDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insight-groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete insight group by ID
func (r *AIConversationInsightGroupService) Delete(ctx context.Context, groupID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insight-groups/%s", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Create a new insight group
func (r *AIConversationInsightGroupService) InsightGroups(ctx context.Context, body AIConversationInsightGroupInsightGroupsParams, opts ...option.RequestOption) (res *InsightTemplateGroupDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversations/insight-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get all insight groups
func (r *AIConversationInsightGroupService) GetInsightGroups(ctx context.Context, query AIConversationInsightGroupGetInsightGroupsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[InsightTemplateGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/conversations/insight-groups"
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

// Get all insight groups
func (r *AIConversationInsightGroupService) GetInsightGroupsAutoPaging(ctx context.Context, query AIConversationInsightGroupGetInsightGroupsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[InsightTemplateGroup] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.GetInsightGroups(ctx, query, opts...))
}

type InsightTemplateGroup struct {
	ID          string            `json:"id,required" format:"uuid"`
	CreatedAt   time.Time         `json:"created_at,required" format:"date-time"`
	Name        string            `json:"name,required"`
	Description string            `json:"description"`
	Insights    []InsightTemplate `json:"insights"`
	Webhook     string            `json:"webhook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Description respjson.Field
		Insights    respjson.Field
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTemplateGroup) RawJSON() string { return r.JSON.raw }
func (r *InsightTemplateGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightTemplateGroupDetail struct {
	Data InsightTemplateGroup `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTemplateGroupDetail) RawJSON() string { return r.JSON.raw }
func (r *InsightTemplateGroupDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightGroupUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Webhook     param.Opt[string] `json:"webhook,omitzero"`
	paramObj
}

func (r AIConversationInsightGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightGroupInsightGroupsParams struct {
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Webhook     param.Opt[string] `json:"webhook,omitzero"`
	paramObj
}

func (r AIConversationInsightGroupInsightGroupsParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightGroupInsightGroupsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightGroupInsightGroupsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightGroupGetInsightGroupsParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationInsightGroupGetInsightGroupsParams]'s query
// parameters as `url.Values`.
func (r AIConversationInsightGroupGetInsightGroupsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
