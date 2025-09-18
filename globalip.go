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
	opts = append(r.Options[:], opts...)
	path := "global_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Global IP.
func (r *GlobalIPService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Global IPs.
func (r *GlobalIPService) List(ctx context.Context, query GlobalIPListParams, opts ...option.RequestOption) (res *GlobalIPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Global IP.
func (r *GlobalIPService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalIPDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("global_ips/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type GlobalIPNewResponse struct {
	Data GlobalIPNewResponseData `json:"data"`
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

type GlobalIPNewResponseData struct {
	// A user specified description for the address.
	Description string `json:"description"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// A user specified name for the address.
	Name string `json:"name"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		IPAddress   respjson.Field
		Name        respjson.Field
		Ports       respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPGetResponse struct {
	Data GlobalIPGetResponseData `json:"data"`
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

type GlobalIPGetResponseData struct {
	// A user specified description for the address.
	Description string `json:"description"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// A user specified name for the address.
	Name string `json:"name"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		IPAddress   respjson.Field
		Name        respjson.Field
		Ports       respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPListResponse struct {
	Data []GlobalIPListResponseData `json:"data"`
	Meta PaginationMeta             `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPListResponseData struct {
	// A user specified description for the address.
	Description string `json:"description"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// A user specified name for the address.
	Name string `json:"name"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		IPAddress   respjson.Field
		Name        respjson.Field
		Ports       respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPListResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPDeleteResponse struct {
	Data GlobalIPDeleteResponseData `json:"data"`
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

type GlobalIPDeleteResponseData struct {
	// A user specified description for the address.
	Description string `json:"description"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// A user specified name for the address.
	Name string `json:"name"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		IPAddress   respjson.Field
		Name        respjson.Field
		Ports       respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Record
}

// Returns the unmodified JSON received from the API
func (r GlobalIPDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPNewParams struct {
	// A user specified description for the address.
	Description param.Opt[string] `json:"description,omitzero"`
	// A user specified name for the address.
	Name param.Opt[string] `json:"name,omitzero"`
	// A Global IP ports grouped by protocol code.
	Ports map[string]any `json:"ports,omitzero"`
	paramObj
}

func (r GlobalIPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalIPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page GlobalIPListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPListParams]'s query parameters as `url.Values`.
func (r GlobalIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type GlobalIPListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPListParamsPage]'s query parameters as `url.Values`.
func (r GlobalIPListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
