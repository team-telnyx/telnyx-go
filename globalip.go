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
// GlobalIPService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPService] method instead.
type GlobalIPService struct {
	Options []option.RequestOption
}

// NewGlobalIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGlobalIPService(opts ...option.RequestOption) (r GlobalIPService) {
	r = GlobalIPService{}
	r.Options = opts
	return
}

// Create a Global IP.
func (r *GlobalIPService) New(ctx context.Context, body GlobalIPNewParams, opts ...option.RequestOption) (res *GlobalIPNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a Global IP.
func (r *GlobalIPService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("global_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all Global IPs.
func (r *GlobalIPService) List(ctx context.Context, query GlobalIPListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[GlobalIP], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "global_ips"
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

// List all Global IPs.
func (r *GlobalIPService) ListAutoPaging(ctx context.Context, query GlobalIPListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[GlobalIP] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Global IP.
func (r *GlobalIPService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("global_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GlobalIP struct {
	// A user specified description for the address.
	Description string `json:"description"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// A user specified name for the address.
	Name string `json:"name"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		IPAddress   respjson.Field
		Name        respjson.Field
		Ports       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIP) RawJSON() string { return r.JSON.raw }
func (r *GlobalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GlobalIP to a GlobalIPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GlobalIPParam.Overrides()
func (r GlobalIP) ToParam() GlobalIPParam {
	return param.Override[GlobalIPParam](json.RawMessage(r.RawJSON()))
}

type GlobalIPParam struct {
	// A user specified description for the address.
	Description param.Opt[string] `json:"description,omitzero"`
	// A user specified name for the address.
	Name param.Opt[string] `json:"name,omitzero"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports,omitzero"`
	RecordParam
}

func (r GlobalIPParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*GlobalIPParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type GlobalIPNewResponse struct {
	Data GlobalIP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPGetResponse struct {
	Data GlobalIP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPDeleteResponse struct {
	Data GlobalIP `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPNewParams struct {
	GlobalIP GlobalIPParam
	paramObj
}

func (r GlobalIPNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GlobalIP)
}
func (r *GlobalIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPListParams]'s query parameters as `url.Values`.
func (r GlobalIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
