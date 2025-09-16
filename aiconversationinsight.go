// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// AIConversationInsightService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationInsightService] method instead.
type AIConversationInsightService struct {
	Options []option.RequestOption
}

// NewAIConversationInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIConversationInsightService(opts ...option.RequestOption) (r AIConversationInsightService) {
	r = AIConversationInsightService{}
	r.Options = opts
	return
}

// Create a new insight
func (r *AIConversationInsightService) New(ctx context.Context, body AIConversationInsightNewParams, opts ...option.RequestOption) (res *InsightTemplateDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/conversations/insights"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get insight by ID
func (r *AIConversationInsightService) Get(ctx context.Context, insightID string, opts ...option.RequestOption) (res *InsightTemplateDetail, err error) {
	opts = append(r.Options[:], opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insights/%s", insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an insight template
func (r *AIConversationInsightService) Update(ctx context.Context, insightID string, body AIConversationInsightUpdateParams, opts ...option.RequestOption) (res *InsightTemplateDetail, err error) {
	opts = append(r.Options[:], opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insights/%s", insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get all insights
func (r *AIConversationInsightService) List(ctx context.Context, query AIConversationInsightListParams, opts ...option.RequestOption) (res *AIConversationInsightListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ai/conversations/insights"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete insight by ID
func (r *AIConversationInsightService) Delete(ctx context.Context, insightID string, opts ...option.RequestOption) (res *AIConversationInsightDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insights/%s", insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type InsightTemplate struct {
	ID           string    `json:"id,required" format:"uuid"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Instructions string    `json:"instructions,required"`
	// Any of "custom", "default".
	InsightType InsightTemplateInsightType `json:"insight_type"`
	// If specified, the output will follow the JSON schema.
	JsonSchema any    `json:"json_schema"`
	Name       string `json:"name"`
	Webhook    string `json:"webhook"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Instructions respjson.Field
		InsightType  respjson.Field
		JsonSchema   respjson.Field
		Name         respjson.Field
		Webhook      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTemplate) RawJSON() string { return r.JSON.raw }
func (r *InsightTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightTemplateInsightType string

const (
	InsightTemplateInsightTypeCustom  InsightTemplateInsightType = "custom"
	InsightTemplateInsightTypeDefault InsightTemplateInsightType = "default"
)

type InsightTemplateDetail struct {
	Data InsightTemplate `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightTemplateDetail) RawJSON() string { return r.JSON.raw }
func (r *InsightTemplateDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightListResponse struct {
	Data []InsightTemplate `json:"data,required"`
	Meta Meta              `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationInsightListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationInsightListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightDeleteResponse = any

type AIConversationInsightNewParams struct {
	Instructions string            `json:"instructions,required"`
	Name         string            `json:"name,required"`
	Webhook      param.Opt[string] `json:"webhook,omitzero"`
	// If specified, the output will follow the JSON schema.
	JsonSchema any `json:"json_schema,omitzero"`
	paramObj
}

func (r AIConversationInsightNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightUpdateParams struct {
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	Webhook      param.Opt[string] `json:"webhook,omitzero"`
	JsonSchema   any               `json:"json_schema,omitzero"`
	paramObj
}

func (r AIConversationInsightUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationInsightListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page AIConversationInsightListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationInsightListParams]'s query parameters as
// `url.Values`.
func (r AIConversationInsightListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type AIConversationInsightListParamsPage struct {
	// Page number (0-based)
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Number of items per page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationInsightListParamsPage]'s query parameters as
// `url.Values`.
func (r AIConversationInsightListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
