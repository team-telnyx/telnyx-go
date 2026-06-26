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
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Global IPs
//
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
	return res, err
}

// Retrieve a Global IP health check.
func (r *GlobalIPHealthCheckService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPHealthCheckGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("global_ip_health_checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all Global IP health checks.
func (r *GlobalIPHealthCheckService) List(ctx context.Context, query GlobalIPHealthCheckListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[GlobalIPHealthCheck], err error) {
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
func (r *GlobalIPHealthCheckService) ListAutoPaging(ctx context.Context, query GlobalIPHealthCheckListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[GlobalIPHealthCheck] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Global IP health check.
func (r *GlobalIPHealthCheckService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPHealthCheckDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("global_ip_health_checks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GlobalIPHealthCheck struct {
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
func (r GlobalIPHealthCheck) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPHealthCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GlobalIPHealthCheck to a GlobalIPHealthCheckParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GlobalIPHealthCheckParam.Overrides()
func (r GlobalIPHealthCheck) ToParam() GlobalIPHealthCheckParam {
	return param.Override[GlobalIPHealthCheckParam](json.RawMessage(r.RawJSON()))
}

type GlobalIPHealthCheckParam struct {
	// Global IP ID.
	GlobalIPID param.Opt[string] `json:"global_ip_id,omitzero" format:"uuid"`
	// A Global IP health check params.
	HealthCheckParams map[string]any `json:"health_check_params,omitzero"`
	// The Global IP health check type.
	HealthCheckType param.Opt[string] `json:"health_check_type,omitzero"`
	RecordParam
}

func (r GlobalIPHealthCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*GlobalIPHealthCheckParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type GlobalIPHealthCheckNewResponse struct {
	Data GlobalIPHealthCheck `json:"data"`
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

type GlobalIPHealthCheckGetResponse struct {
	Data GlobalIPHealthCheck `json:"data"`
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

type GlobalIPHealthCheckDeleteResponse struct {
	Data GlobalIPHealthCheck `json:"data"`
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

type GlobalIPHealthCheckNewParams struct {
	GlobalIPHealthCheck GlobalIPHealthCheckParam
	paramObj
}

func (r GlobalIPHealthCheckNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GlobalIPHealthCheck)
}
func (r *GlobalIPHealthCheckNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPHealthCheckListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
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
