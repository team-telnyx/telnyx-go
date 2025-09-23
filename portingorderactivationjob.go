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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingOrderActivationJobService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderActivationJobService] method instead.
type PortingOrderActivationJobService struct {
	Options []option.RequestOption
}

// NewPortingOrderActivationJobService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderActivationJobService(opts ...option.RequestOption) (r PortingOrderActivationJobService) {
	r = PortingOrderActivationJobService{}
	r.Options = opts
	return
}

// Returns a porting activation job.
func (r *PortingOrderActivationJobService) Get(ctx context.Context, activationJobID string, query PortingOrderActivationJobGetParams, opts ...option.RequestOption) (res *PortingOrderActivationJobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if activationJobID == "" {
		err = errors.New("missing required activationJobId parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/activation_jobs/%s", query.ID, activationJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the activation time of a porting activation job.
func (r *PortingOrderActivationJobService) Update(ctx context.Context, activationJobID string, params PortingOrderActivationJobUpdateParams, opts ...option.RequestOption) (res *PortingOrderActivationJobUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if activationJobID == "" {
		err = errors.New("missing required activationJobId parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/activation_jobs/%s", params.ID, activationJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of your porting activation jobs.
func (r *PortingOrderActivationJobService) List(ctx context.Context, id string, query PortingOrderActivationJobListParams, opts ...option.RequestOption) (res *PortingOrderActivationJobListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/activation_jobs", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PortingOrderActivationJobGetResponse struct {
	Data PortingOrdersActivationJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActivationJobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActivationJobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActivationJobUpdateResponse struct {
	Data PortingOrdersActivationJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActivationJobUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActivationJobUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActivationJobListResponse struct {
	Data []PortingOrdersActivationJob `json:"data"`
	Meta PaginationMeta               `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActivationJobListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActivationJobListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActivationJobGetParams struct {
	ID string `path:"id,required" format:"uuid" json:"-"`
	paramObj
}

type PortingOrderActivationJobUpdateParams struct {
	ID string `path:"id,required" format:"uuid" json:"-"`
	// The desired activation time. The activation time should be between any of the
	// activation windows.
	ActivateAt param.Opt[time.Time] `json:"activate_at,omitzero" format:"date-time"`
	paramObj
}

func (r PortingOrderActivationJobUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderActivationJobUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderActivationJobUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActivationJobListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderActivationJobListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActivationJobListParams]'s query parameters as
// `url.Values`.
func (r PortingOrderActivationJobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderActivationJobListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderActivationJobListParamsPage]'s query parameters
// as `url.Values`.
func (r PortingOrderActivationJobListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
