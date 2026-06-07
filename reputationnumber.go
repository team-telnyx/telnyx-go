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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// ReputationNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReputationNumberService] method instead.
type ReputationNumberService struct {
	Options []option.RequestOption
}

// NewReputationNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReputationNumberService(opts ...option.RequestOption) (r ReputationNumberService) {
	r = ReputationNumberService{}
	r.Options = opts
	return
}

// Convenience alias for
// `GET /v2/enterprises/{enterprise_id}/reputation/numbers/{phone_number}`.
func (r *ReputationNumberService) Get(ctx context.Context, phoneNumber string, query ReputationNumberGetParams, opts ...option.RequestOption) (res *ReputationNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("reputation/numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Convenience alias for `GET /v2/enterprises/{enterprise_id}/reputation/numbers`
// that returns numbers across every enterprise you own. Useful when you don't want
// to look up the enterprise id first.
func (r *ReputationNumberService) List(ctx context.Context, query ReputationNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ReputationNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "reputation/numbers"
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

// Convenience alias for `GET /v2/enterprises/{enterprise_id}/reputation/numbers`
// that returns numbers across every enterprise you own. Useful when you don't want
// to look up the enterprise id first.
func (r *ReputationNumberService) ListAutoPaging(ctx context.Context, query ReputationNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ReputationNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Convenience alias for
// `DELETE /v2/enterprises/{enterprise_id}/reputation/numbers/{phone_number}`.
func (r *ReputationNumberService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("reputation/numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ReputationNumberGetResponse struct {
	Data ReputationNumberGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberGetResponseData struct {
	ID           string    `json:"id" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// `null` until the first refresh has been collected for this number.
	ReputationData shared.ReputationData `json:"reputation_data" api:"nullable"`
	UpdatedAt      time.Time             `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumber    respjson.Field
		ReputationData respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberListResponse struct {
	ID           string    `json:"id" format:"uuid"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	EnterpriseID string    `json:"enterprise_id" format:"uuid"`
	// E.164 with leading `+`.
	PhoneNumber string `json:"phone_number"`
	// `null` until the first refresh has been collected for this number.
	ReputationData shared.ReputationData `json:"reputation_data" api:"nullable"`
	UpdatedAt      time.Time             `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumber    respjson.Field
		ReputationData respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberGetParams struct {
	// When true, fetches fresh reputation data (incurs API cost). When false
	// (default), returns cached data.
	Fresh param.Opt[bool] `query:"fresh,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReputationNumberGetParams]'s query parameters as
// `url.Values`.
func (r ReputationNumberGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReputationNumberListParams struct {
	// Filter by enterprise ID.
	FilterEnterpriseID param.Opt[string] `query:"filter[enterprise_id],omitzero" format:"uuid" json:"-"`
	// Partial match on phone number. Must contain at least 5 digits.
	FilterPhoneNumberContains param.Opt[string] `query:"filter[phone_number][contains],omitzero" json:"-"`
	// Exact phone-number match (E.164).
	FilterPhoneNumberEq param.Opt[string] `query:"filter[phone_number][eq],omitzero" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Maximum 250; values above are clamped to 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by specific phone number (E.164 format).
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReputationNumberListParams]'s query parameters as
// `url.Values`.
func (r ReputationNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
