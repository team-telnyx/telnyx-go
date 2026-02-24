// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// MessagingOptoutService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingOptoutService] method instead.
type MessagingOptoutService struct {
	Options []option.RequestOption
}

// NewMessagingOptoutService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingOptoutService(opts ...option.RequestOption) (r MessagingOptoutService) {
	r = MessagingOptoutService{}
	r.Options = opts
	return
}

// Retrieve a list of opt-out blocks.
func (r *MessagingOptoutService) List(ctx context.Context, query MessagingOptoutListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[MessagingOptoutListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging_optouts"
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

// Retrieve a list of opt-out blocks.
func (r *MessagingOptoutService) ListAutoPaging(ctx context.Context, query MessagingOptoutListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[MessagingOptoutListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type MessagingOptoutListResponse struct {
	// The timestamp when the opt-out was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	From string `json:"from"`
	// The keyword that triggered the opt-out.
	Keyword string `json:"keyword" api:"nullable"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id" api:"nullable"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		From               respjson.Field
		Keyword            respjson.Field
		MessagingProfileID respjson.Field
		To                 respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingOptoutListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingOptoutListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingOptoutListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// If receiving address (+E.164 formatted phone number) should be redacted
	RedactionEnabled param.Opt[string] `query:"redaction_enabled,omitzero" json:"-"`
	// Consolidated created_at parameter (deepObject style). Originally:
	// created_at[gte], created_at[lte]
	CreatedAt MessagingOptoutListParamsCreatedAt `query:"created_at,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[messaging_profile_id], filter[from]
	Filter MessagingOptoutListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingOptoutListParams]'s query parameters as
// `url.Values`.
func (r MessagingOptoutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated created_at parameter (deepObject style). Originally:
// created_at[gte], created_at[lte]
type MessagingOptoutListParamsCreatedAt struct {
	// Filter opt-outs created after this date (ISO-8601 format)
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter opt-outs created before this date (ISO-8601 format)
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingOptoutListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r MessagingOptoutListParamsCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[messaging_profile_id], filter[from]
type MessagingOptoutListParamsFilter struct {
	// The sending address (+E.164 formatted phone number, alphanumeric sender ID, or
	// short code) to retrieve opt-outs for
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// The ID of the messaging profile to retrieve opt-outs for
	MessagingProfileID param.Opt[string] `query:"messaging_profile_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingOptoutListParamsFilter]'s query parameters as
// `url.Values`.
func (r MessagingOptoutListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
