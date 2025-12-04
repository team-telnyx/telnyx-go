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
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// ShortCodeService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShortCodeService] method instead.
type ShortCodeService struct {
	Options []option.RequestOption
}

// NewShortCodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewShortCodeService(opts ...option.RequestOption) (r ShortCodeService) {
	r = ShortCodeService{}
	r.Options = opts
	return
}

// Retrieve a short code
func (r *ShortCodeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ShortCodeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("short_codes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the settings for a specific short code. To unbind a short code from a
// profile, set the `messaging_profile_id` to `null` or an empty string. To add or
// update tags, include the tags field as an array of strings.
func (r *ShortCodeService) Update(ctx context.Context, id string, body ShortCodeUpdateParams, opts ...option.RequestOption) (res *ShortCodeUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("short_codes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List short codes
func (r *ShortCodeService) List(ctx context.Context, query ShortCodeListParams, opts ...option.RequestOption) (res *ShortCodeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "short_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ShortCodeGetResponse struct {
	Data shared.ShortCode `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortCodeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortCodeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortCodeUpdateResponse struct {
	Data shared.ShortCode `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortCodeUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortCodeUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortCodeListResponse struct {
	Data []shared.ShortCode `json:"data"`
	Meta PaginationMeta     `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortCodeListResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortCodeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortCodeUpdateParams struct {
	// Unique identifier for a messaging profile.
	MessagingProfileID string   `json:"messaging_profile_id,required"`
	Tags               []string `json:"tags,omitzero"`
	paramObj
}

func (r ShortCodeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ShortCodeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortCodeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortCodeListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[messaging_profile_id]
	Filter ShortCodeListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page ShortCodeListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShortCodeListParams]'s query parameters as `url.Values`.
func (r ShortCodeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[messaging_profile_id]
type ShortCodeListParamsFilter struct {
	// Filter by Messaging Profile ID. Use the string `null` for phone numbers without
	// assigned profiles. A synonym for the `/messaging_profiles/{id}/short_codes`
	// endpoint when querying about an extant profile.
	MessagingProfileID param.Opt[string] `query:"messaging_profile_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShortCodeListParamsFilter]'s query parameters as
// `url.Values`.
func (r ShortCodeListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type ShortCodeListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShortCodeListParamsPage]'s query parameters as
// `url.Values`.
func (r ShortCodeListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
