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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// AIMcpServerService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMcpServerService] method instead.
type AIMcpServerService struct {
	Options []option.RequestOption
}

// NewAIMcpServerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMcpServerService(opts ...option.RequestOption) (r AIMcpServerService) {
	r = AIMcpServerService{}
	r.Options = opts
	return
}

// Create a new MCP server.
func (r *AIMcpServerService) New(ctx context.Context, body AIMcpServerNewParams, opts ...option.RequestOption) (res *AIMcpServerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/mcp_servers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details for a specific MCP server.
func (r *AIMcpServerService) Get(ctx context.Context, mcpServerID string, opts ...option.RequestOption) (res *AIMcpServerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("ai/mcp_servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing MCP server.
func (r *AIMcpServerService) Update(ctx context.Context, mcpServerID string, body AIMcpServerUpdateParams, opts ...option.RequestOption) (res *AIMcpServerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("ai/mcp_servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of MCP servers.
func (r *AIMcpServerService) List(ctx context.Context, query AIMcpServerListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPaginationTopLevelArray[AIMcpServerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/mcp_servers"
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

// Retrieve a list of MCP servers.
func (r *AIMcpServerService) ListAutoPaging(ctx context.Context, query AIMcpServerListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationTopLevelArrayAutoPager[AIMcpServerListResponse] {
	return pagination.NewDefaultFlatPaginationTopLevelArrayAutoPager(r.List(ctx, query, opts...))
}

// Delete a specific MCP server.
func (r *AIMcpServerService) Delete(ctx context.Context, mcpServerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("ai/mcp_servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AIMcpServerNewResponse struct {
	ID           string    `json:"id,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Name         string    `json:"name,required"`
	Type         string    `json:"type,required"`
	URL          string    `json:"url,required"`
	AllowedTools []string  `json:"allowed_tools,nullable"`
	APIKeyRef    string    `json:"api_key_ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		AllowedTools respjson.Field
		APIKeyRef    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMcpServerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMcpServerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerGetResponse struct {
	ID           string    `json:"id,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Name         string    `json:"name,required"`
	Type         string    `json:"type,required"`
	URL          string    `json:"url,required"`
	AllowedTools []string  `json:"allowed_tools,nullable"`
	APIKeyRef    string    `json:"api_key_ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		AllowedTools respjson.Field
		APIKeyRef    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMcpServerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMcpServerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerUpdateResponse struct {
	ID           string    `json:"id,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Name         string    `json:"name,required"`
	Type         string    `json:"type,required"`
	URL          string    `json:"url,required"`
	AllowedTools []string  `json:"allowed_tools,nullable"`
	APIKeyRef    string    `json:"api_key_ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		AllowedTools respjson.Field
		APIKeyRef    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMcpServerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMcpServerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerListResponse struct {
	ID           string    `json:"id,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	Name         string    `json:"name,required"`
	Type         string    `json:"type,required"`
	URL          string    `json:"url,required"`
	AllowedTools []string  `json:"allowed_tools,nullable"`
	APIKeyRef    string    `json:"api_key_ref,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		URL          respjson.Field
		AllowedTools respjson.Field
		APIKeyRef    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMcpServerListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMcpServerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerNewParams struct {
	Name         string            `json:"name,required"`
	Type         string            `json:"type,required"`
	URL          string            `json:"url,required"`
	APIKeyRef    param.Opt[string] `json:"api_key_ref,omitzero"`
	AllowedTools []string          `json:"allowed_tools,omitzero"`
	paramObj
}

func (r AIMcpServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMcpServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMcpServerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerUpdateParams struct {
	APIKeyRef    param.Opt[string]    `json:"api_key_ref,omitzero"`
	ID           param.Opt[string]    `json:"id,omitzero"`
	CreatedAt    param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	Name         param.Opt[string]    `json:"name,omitzero"`
	Type         param.Opt[string]    `json:"type,omitzero"`
	URL          param.Opt[string]    `json:"url,omitzero"`
	AllowedTools []string             `json:"allowed_tools,omitzero"`
	paramObj
}

func (r AIMcpServerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMcpServerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMcpServerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMcpServerListParams struct {
	PageNumber param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	Type       param.Opt[string] `query:"type,omitzero" json:"-"`
	URL        param.Opt[string] `query:"url,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMcpServerListParams]'s query parameters as `url.Values`.
func (r AIMcpServerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
