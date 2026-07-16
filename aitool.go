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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Configure AI assistant specifications
//
// AIToolService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIToolService] method instead.
type AIToolService struct {
	Options []option.RequestOption
}

// NewAIToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIToolService(opts ...option.RequestOption) (r AIToolService) {
	r = AIToolService{}
	r.Options = opts
	return
}

// Create Tool
func (r *AIToolService) New(ctx context.Context, body AIToolNewParams, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Tool
func (r *AIToolService) Get(ctx context.Context, toolID string, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Tool
func (r *AIToolService) Update(ctx context.Context, toolID string, body AIToolUpdateParams, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Tools
func (r *AIToolService) List(ctx context.Context, query AIToolListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SharedToolResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/tools"
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

// List Tools
func (r *AIToolService) ListAutoPaging(ctx context.Context, query AIToolListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SharedToolResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete Tool
func (r *AIToolService) Delete(ctx context.Context, toolID string, opts ...option.RequestOption) (res *AIToolDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type PayToolParamsResp struct {
	// The name of the pay connector configured in the Telnyx API. Must reference an
	// existing pay connector for this organization.
	ConnectorName string `json:"connector_name" api:"required"`
	// Default currency for payments processed by this tool.
	Currency string `json:"currency"`
	// Optional description of the pay tool that will be passed to the assistant.
	Description string `json:"description" api:"nullable"`
	// Default payment method for payments processed by this tool.
	PaymentMethod string `json:"payment_method"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName respjson.Field
		Currency      respjson.Field
		Description   respjson.Field
		PaymentMethod respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *PayToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PayToolParamsResp to a PayToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PayToolParams.Overrides()
func (r PayToolParamsResp) ToParam() PayToolParams {
	return param.Override[PayToolParams](json.RawMessage(r.RawJSON()))
}

// The property ConnectorName is required.
type PayToolParams struct {
	// The name of the pay connector configured in the Telnyx API. Must reference an
	// existing pay connector for this organization.
	ConnectorName string `json:"connector_name" api:"required"`
	// Optional description of the pay tool that will be passed to the assistant.
	Description param.Opt[string] `json:"description,omitzero"`
	// Default currency for payments processed by this tool.
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Default payment method for payments processed by this tool.
	PaymentMethod param.Opt[string] `json:"payment_method,omitzero"`
	paramObj
}

func (r PayToolParams) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SharedToolResponse struct {
	ID             string         `json:"id" api:"required"`
	ToolDefinition map[string]any `json:"tool_definition" api:"required"`
	Type           string         `json:"type" api:"required"`
	CreatedAt      string         `json:"created_at"`
	DisplayName    string         `json:"display_name"`
	TimeoutMs      int64          `json:"timeout_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ToolDefinition respjson.Field
		Type           respjson.Field
		CreatedAt      respjson.Field
		DisplayName    respjson.Field
		TimeoutMs      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SharedToolResponse) RawJSON() string { return r.JSON.raw }
func (r *SharedToolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolDeleteResponse = any

type AIToolNewParams struct {
	DisplayName    string           `json:"display_name" api:"required"`
	Type           string           `json:"type" api:"required"`
	TimeoutMs      param.Opt[int64] `json:"timeout_ms,omitzero"`
	ClientSideTool map[string]any   `json:"client_side_tool,omitzero"`
	Function       map[string]any   `json:"function,omitzero"`
	Handoff        map[string]any   `json:"handoff,omitzero"`
	Invite         map[string]any   `json:"invite,omitzero"`
	Pay            PayToolParams    `json:"pay,omitzero"`
	Retrieval      map[string]any   `json:"retrieval,omitzero"`
	Webhook        map[string]any   `json:"webhook,omitzero"`
	paramObj
}

func (r AIToolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIToolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIToolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolUpdateParams struct {
	DisplayName    param.Opt[string] `json:"display_name,omitzero"`
	TimeoutMs      param.Opt[int64]  `json:"timeout_ms,omitzero"`
	Type           param.Opt[string] `json:"type,omitzero"`
	ClientSideTool map[string]any    `json:"client_side_tool,omitzero"`
	Function       map[string]any    `json:"function,omitzero"`
	Handoff        map[string]any    `json:"handoff,omitzero"`
	Invite         map[string]any    `json:"invite,omitzero"`
	Pay            PayToolParams     `json:"pay,omitzero"`
	Retrieval      map[string]any    `json:"retrieval,omitzero"`
	Webhook        map[string]any    `json:"webhook,omitzero"`
	paramObj
}

func (r AIToolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIToolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIToolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolListParams struct {
	// Filter results by filter name.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Filter results by filter type.
	FilterType param.Opt[string] `query:"filter[type],omitzero" json:"-"`
	// Page number to retrieve (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items to return per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIToolListParams]'s query parameters as `url.Values`.
func (r AIToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
