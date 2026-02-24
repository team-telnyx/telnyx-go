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
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
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
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversations/insights"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get insight by ID
func (r *AIConversationInsightService) Get(ctx context.Context, insightID string, opts ...option.RequestOption) (res *InsightTemplateDetail, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insights/%s", insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get all insights
func (r *AIConversationInsightService) List(ctx context.Context, query AIConversationInsightListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[InsightTemplate], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/conversations/insights"
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

// Get all insights
func (r *AIConversationInsightService) ListAutoPaging(ctx context.Context, query AIConversationInsightListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[InsightTemplate] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete insight by ID
func (r *AIConversationInsightService) Delete(ctx context.Context, insightID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insights/%s", insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type InsightTemplate struct {
	ID           string    `json:"id" api:"required" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" api:"required" format:"date-time"`
	Instructions string    `json:"instructions" api:"required"`
	// Any of "custom", "default".
	InsightType InsightTemplateInsightType `json:"insight_type"`
	// If specified, the output will follow the JSON schema.
	JsonSchema InsightTemplateJsonSchemaUnion `json:"json_schema"`
	Name       string                         `json:"name"`
	Webhook    string                         `json:"webhook"`
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

// InsightTemplateJsonSchemaUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInsightTemplateJsonSchemaJsonSchemaObjectItem]
type InsightTemplateJsonSchemaUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfInsightTemplateJsonSchemaJsonSchemaObjectItem any `json:",inline"`
	JSON                                            struct {
		OfString                                        respjson.Field
		OfInsightTemplateJsonSchemaJsonSchemaObjectItem respjson.Field
		raw                                             string
	} `json:"-"`
}

func (u InsightTemplateJsonSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightTemplateJsonSchemaUnion) AsJsonSchemaObject() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightTemplateJsonSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightTemplateJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightTemplateDetail struct {
	Data InsightTemplate `json:"data" api:"required"`
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

type AIConversationInsightNewParams struct {
	Instructions string            `json:"instructions" api:"required"`
	Name         string            `json:"name" api:"required"`
	Webhook      param.Opt[string] `json:"webhook,omitzero"`
	// If specified, the output will follow the JSON schema.
	JsonSchema AIConversationInsightNewParamsJsonSchemaUnion `json:"json_schema,omitzero"`
	paramObj
}

func (r AIConversationInsightNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIConversationInsightNewParamsJsonSchemaUnion struct {
	OfString           param.Opt[string] `json:",omitzero,inline"`
	OfJsonSchemaObject map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u AIConversationInsightNewParamsJsonSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfJsonSchemaObject)
}
func (u *AIConversationInsightNewParamsJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIConversationInsightNewParamsJsonSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfJsonSchemaObject) {
		return &u.OfJsonSchemaObject
	}
	return nil
}

type AIConversationInsightUpdateParams struct {
	Instructions param.Opt[string]                                `json:"instructions,omitzero"`
	Name         param.Opt[string]                                `json:"name,omitzero"`
	Webhook      param.Opt[string]                                `json:"webhook,omitzero"`
	JsonSchema   AIConversationInsightUpdateParamsJsonSchemaUnion `json:"json_schema,omitzero"`
	paramObj
}

func (r AIConversationInsightUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationInsightUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationInsightUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIConversationInsightUpdateParamsJsonSchemaUnion struct {
	OfString           param.Opt[string] `json:",omitzero,inline"`
	OfJsonSchemaObject map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u AIConversationInsightUpdateParamsJsonSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfJsonSchemaObject)
}
func (u *AIConversationInsightUpdateParamsJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIConversationInsightUpdateParamsJsonSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfJsonSchemaObject) {
		return &u.OfJsonSchemaObject
	}
	return nil
}

type AIConversationInsightListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
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
