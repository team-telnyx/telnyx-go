// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// UserTagService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserTagService] method instead.
type UserTagService struct {
	Options []option.RequestOption
}

// NewUserTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserTagService(opts ...option.RequestOption) (r UserTagService) {
	r = UserTagService{}
	r.Options = opts
	return
}

// List all user tags.
func (r *UserTagService) List(ctx context.Context, query UserTagListParams, opts ...option.RequestOption) (res *UserTagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user_tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserTagListResponse struct {
	Data UserTagListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTagListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserTagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTagListResponseData struct {
	// A list of your tags on the given resource type. NOTE: The casing of the tags
	// returned will not necessarily match the casing of the tags when they were added
	// to a resource. This is because the tags will have the casing of the first time
	// they were used within the Telnyx system regardless of source.
	NumberTags []string `json:"number_tags"`
	// A list of your tags on the given resource type. NOTE: The casing of the tags
	// returned will not necessarily match the casing of the tags when they were added
	// to a resource. This is because the tags will have the casing of the first time
	// they were used within the Telnyx system regardless of source.
	OutboundProfileTags []string `json:"outbound_profile_tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberTags          respjson.Field
		OutboundProfileTags respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserTagListResponseData) RawJSON() string { return r.JSON.raw }
func (r *UserTagListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserTagListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[starts_with]
	Filter UserTagListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserTagListParams]'s query parameters as `url.Values`.
func (r UserTagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[starts_with]
type UserTagListParamsFilter struct {
	// Filter tags by prefix
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserTagListParamsFilter]'s query parameters as
// `url.Values`.
func (r UserTagListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
