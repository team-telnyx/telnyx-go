// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// RequirementTypeService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequirementTypeService] method instead.
type RequirementTypeService struct {
	Options []option.RequestOption
}

// NewRequirementTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequirementTypeService(opts ...option.RequestOption) (r RequirementTypeService) {
	r = RequirementTypeService{}
	r.Options = opts
	return
}

// Retrieve a requirement type by id
func (r *RequirementTypeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RequirementTypeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirement_types/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all requirement types ordered by created_at descending
func (r *RequirementTypeService) List(ctx context.Context, query RequirementTypeListParams, opts ...option.RequestOption) (res *RequirementTypeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "requirement_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RequirementTypeGetResponse struct {
	Data shared.DocReqsRequirementType `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementTypeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RequirementTypeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementTypeListResponse struct {
	Data []shared.DocReqsRequirementType `json:"data"`
	Meta PaginationMeta                  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *RequirementTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementTypeListParams struct {
	// Consolidated filter parameter for requirement types (deepObject style).
	// Originally: filter[name]
	Filter RequirementTypeListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated sort parameter for requirement types (deepObject style).
	// Originally: sort[]
	//
	// Any of "name", "created_at", "updated_at", "-name", "-created_at",
	// "-updated_at".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementTypeListParams]'s query parameters as
// `url.Values`.
func (r RequirementTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for requirement types (deepObject style).
// Originally: filter[name]
type RequirementTypeListParamsFilter struct {
	Name RequirementTypeListParamsFilterName `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementTypeListParamsFilter]'s query parameters as
// `url.Values`.
func (r RequirementTypeListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RequirementTypeListParamsFilterName struct {
	// Filters requirement types to those whose name contains a certain string.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementTypeListParamsFilterName]'s query parameters as
// `url.Values`.
func (r RequirementTypeListParamsFilterName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
