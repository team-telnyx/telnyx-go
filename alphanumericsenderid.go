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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AlphanumericSenderIDService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphanumericSenderIDService] method instead.
type AlphanumericSenderIDService struct {
	Options []option.RequestOption
}

// NewAlphanumericSenderIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlphanumericSenderIDService(opts ...option.RequestOption) (r AlphanumericSenderIDService) {
	r = AlphanumericSenderIDService{}
	r.Options = opts
	return
}

// Create a new alphanumeric sender ID associated with a messaging profile.
func (r *AlphanumericSenderIDService) New(ctx context.Context, body AlphanumericSenderIDNewParams, opts ...option.RequestOption) (res *AlphanumericSenderIDNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "alphanumeric_sender_ids"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific alphanumeric sender ID.
func (r *AlphanumericSenderIDService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AlphanumericSenderIDGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("alphanumeric_sender_ids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all alphanumeric sender IDs for the authenticated user.
func (r *AlphanumericSenderIDService) List(ctx context.Context, query AlphanumericSenderIDListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AlphanumericSenderIDListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "alphanumeric_sender_ids"
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

// List all alphanumeric sender IDs for the authenticated user.
func (r *AlphanumericSenderIDService) ListAutoPaging(ctx context.Context, query AlphanumericSenderIDListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AlphanumericSenderIDListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete an alphanumeric sender ID and disassociate it from its messaging profile.
func (r *AlphanumericSenderIDService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AlphanumericSenderIDDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("alphanumeric_sender_ids/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AlphanumericSenderIDNewResponse struct {
	Data AlphanumericSenderIDNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDNewResponseData struct {
	// Uniquely identifies the alphanumeric sender ID resource.
	ID string `json:"id" format:"uuid"`
	// The alphanumeric sender ID string.
	AlphanumericSenderID string `json:"alphanumeric_sender_id"`
	// The messaging profile this sender ID belongs to.
	MessagingProfileID string `json:"messaging_profile_id" format:"uuid"`
	// The organization that owns this sender ID.
	OrganizationID string `json:"organization_id"`
	// Any of "alphanumeric_sender_id".
	RecordType string `json:"record_type"`
	// A US long code number to use as fallback when sending to US destinations.
	UsLongCodeFallback string `json:"us_long_code_fallback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlphanumericSenderID respjson.Field
		MessagingProfileID   respjson.Field
		OrganizationID       respjson.Field
		RecordType           respjson.Field
		UsLongCodeFallback   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDGetResponse struct {
	Data AlphanumericSenderIDGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDGetResponseData struct {
	// Uniquely identifies the alphanumeric sender ID resource.
	ID string `json:"id" format:"uuid"`
	// The alphanumeric sender ID string.
	AlphanumericSenderID string `json:"alphanumeric_sender_id"`
	// The messaging profile this sender ID belongs to.
	MessagingProfileID string `json:"messaging_profile_id" format:"uuid"`
	// The organization that owns this sender ID.
	OrganizationID string `json:"organization_id"`
	// Any of "alphanumeric_sender_id".
	RecordType string `json:"record_type"`
	// A US long code number to use as fallback when sending to US destinations.
	UsLongCodeFallback string `json:"us_long_code_fallback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlphanumericSenderID respjson.Field
		MessagingProfileID   respjson.Field
		OrganizationID       respjson.Field
		RecordType           respjson.Field
		UsLongCodeFallback   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDListResponse struct {
	// Uniquely identifies the alphanumeric sender ID resource.
	ID string `json:"id" format:"uuid"`
	// The alphanumeric sender ID string.
	AlphanumericSenderID string `json:"alphanumeric_sender_id"`
	// The messaging profile this sender ID belongs to.
	MessagingProfileID string `json:"messaging_profile_id" format:"uuid"`
	// The organization that owns this sender ID.
	OrganizationID string `json:"organization_id"`
	// Any of "alphanumeric_sender_id".
	RecordType AlphanumericSenderIDListResponseRecordType `json:"record_type"`
	// A US long code number to use as fallback when sending to US destinations.
	UsLongCodeFallback string `json:"us_long_code_fallback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlphanumericSenderID respjson.Field
		MessagingProfileID   respjson.Field
		OrganizationID       respjson.Field
		RecordType           respjson.Field
		UsLongCodeFallback   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDListResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDListResponseRecordType string

const (
	AlphanumericSenderIDListResponseRecordTypeAlphanumericSenderID AlphanumericSenderIDListResponseRecordType = "alphanumeric_sender_id"
)

type AlphanumericSenderIDDeleteResponse struct {
	Data AlphanumericSenderIDDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDDeleteResponseData struct {
	// Uniquely identifies the alphanumeric sender ID resource.
	ID string `json:"id" format:"uuid"`
	// The alphanumeric sender ID string.
	AlphanumericSenderID string `json:"alphanumeric_sender_id"`
	// The messaging profile this sender ID belongs to.
	MessagingProfileID string `json:"messaging_profile_id" format:"uuid"`
	// The organization that owns this sender ID.
	OrganizationID string `json:"organization_id"`
	// Any of "alphanumeric_sender_id".
	RecordType string `json:"record_type"`
	// A US long code number to use as fallback when sending to US destinations.
	UsLongCodeFallback string `json:"us_long_code_fallback"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AlphanumericSenderID respjson.Field
		MessagingProfileID   respjson.Field
		OrganizationID       respjson.Field
		RecordType           respjson.Field
		UsLongCodeFallback   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphanumericSenderIDDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *AlphanumericSenderIDDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDNewParams struct {
	// The alphanumeric sender ID string.
	AlphanumericSenderID string `json:"alphanumeric_sender_id,required"`
	// The messaging profile to associate the sender ID with.
	MessagingProfileID string `json:"messaging_profile_id,required" format:"uuid"`
	// A US long code number to use as fallback when sending to US destinations.
	UsLongCodeFallback param.Opt[string] `json:"us_long_code_fallback,omitzero"`
	paramObj
}

func (r AlphanumericSenderIDNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphanumericSenderIDNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphanumericSenderIDNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphanumericSenderIDListParams struct {
	// Filter by messaging profile ID.
	FilterMessagingProfileID param.Opt[string] `query:"filter[messaging_profile_id],omitzero" format:"uuid" json:"-"`
	// Page number.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Page size.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphanumericSenderIDListParams]'s query parameters as
// `url.Values`.
func (r AlphanumericSenderIDListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
