// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// GlobalIPHealthCheckService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPHealthCheckService] method instead.
type GlobalIPHealthCheckService struct {
	Options []option.RequestOption
}

// NewGlobalIPHealthCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalIPHealthCheckService(opts ...option.RequestOption) (r GlobalIPHealthCheckService) {
	r = GlobalIPHealthCheckService{}
	r.Options = opts
	return
}

// Create a Global IP health check.
func (r *GlobalIPHealthCheckService) New(ctx context.Context, body GlobalIPHealthCheckNewParams, opts ...option.RequestOption) (res *GlobalIPHealthCheckNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_health_checks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Global IP health check.
func (r *GlobalIPHealthCheckService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPHealthCheckGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ip_health_checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Global IP health checks.
func (r *GlobalIPHealthCheckService) List(ctx context.Context, query GlobalIPHealthCheckListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[GlobalIPHealthCheckListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "global_ip_health_checks"
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

// List all Global IP health checks.
func (r *GlobalIPHealthCheckService) ListAutoPaging(ctx context.Context, query GlobalIPHealthCheckListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[GlobalIPHealthCheckListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Global IP health check.
func (r *GlobalIPHealthCheckService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPHealthCheckDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ip_health_checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type GlobalIPHealthCheckNewResponse struct {
	Data GlobalIPHealthCheckNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckNewResponseData struct {
	// Global IP ID.
	GlobalIPID string `json:"global_ip_id" format:"uuid"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params"`
	// The Global IP health check type.
	HealthCheckType string `json:"health_check_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIPID        respjson.Field
		HealthCheckParams respjson.Field
		HealthCheckType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckGetResponse struct {
	Data GlobalIPHealthCheckGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckGetResponseData struct {
	// Global IP ID.
	GlobalIPID string `json:"global_ip_id" format:"uuid"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params"`
	// The Global IP health check type.
	HealthCheckType string `json:"health_check_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIPID        respjson.Field
		HealthCheckParams respjson.Field
		HealthCheckType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckListResponse struct {
	// Global IP ID.
	GlobalIPID string `json:"global_ip_id" format:"uuid"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params"`
	// The Global IP health check type.
	HealthCheckType string `json:"health_check_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIPID        respjson.Field
		HealthCheckParams respjson.Field
		HealthCheckType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckDeleteResponse struct {
	Data GlobalIPHealthCheckDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckDeleteResponseData struct {
	// Global IP ID.
	GlobalIPID string `json:"global_ip_id" format:"uuid"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params"`
	// The Global IP health check type.
	HealthCheckType string `json:"health_check_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIPID        respjson.Field
		HealthCheckParams respjson.Field
		HealthCheckType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPHealthCheckDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheckDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckNewParams struct {
	// Global IP ID.
	GlobalIPID param.Opt[string] `json:"global_ip_id,omitzero" format:"uuid"`
	// The Global IP health check type.
	HealthCheckType param.Opt[string] `json:"health_check_type,omitzero"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params,omitzero"`
	paramObj
}

func (r GlobalIPHealthCheckNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalIPHealthCheckNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalIPHealthCheckNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page GlobalIPHealthCheckListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPHealthCheckListParams]'s query parameters as
// `url.Values`.
func (r GlobalIPHealthCheckListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type GlobalIPHealthCheckListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPHealthCheckListParamsPage]'s query parameters as
// `url.Values`.
func (r GlobalIPHealthCheckListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
