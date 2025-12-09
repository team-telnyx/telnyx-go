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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcPhoneNumberCampaignService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcPhoneNumberCampaignService] method instead.
type Number10dlcPhoneNumberCampaignService struct {
	Options []option.RequestOption
}

// NewNumber10dlcPhoneNumberCampaignService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNumber10dlcPhoneNumberCampaignService(opts ...option.RequestOption) (r Number10dlcPhoneNumberCampaignService) {
	r = Number10dlcPhoneNumberCampaignService{}
	r.Options = opts
	return
}

// Create New Phone Number Campaign
func (r *Number10dlcPhoneNumberCampaignService) New(ctx context.Context, body Number10dlcPhoneNumberCampaignNewParams, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/phone_number_campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an individual phone number/campaign assignment by `phoneNumber`.
func (r *Number10dlcPhoneNumberCampaignService) Get(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("10dlc/phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create New Phone Number Campaign
func (r *Number10dlcPhoneNumberCampaignService) Update(ctx context.Context, phoneNumber string, body Number10dlcPhoneNumberCampaignUpdateParams, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("10dlc/phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve All Phone Number Campaigns
func (r *Number10dlcPhoneNumberCampaignService) List(ctx context.Context, query Number10dlcPhoneNumberCampaignListParams, opts ...option.RequestOption) (res *Number10dlcPhoneNumberCampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/phone_number_campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint allows you to remove a campaign assignment from the supplied
// `phoneNumber`.
func (r *Number10dlcPhoneNumberCampaignService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("10dlc/phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Number10dlcPhoneNumberCampaignListResponse struct {
	Page         int64                 `json:"page,required"`
	Records      []PhoneNumberCampaign `json:"records,required"`
	TotalRecords int64                 `json:"totalRecords,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page         respjson.Field
		Records      respjson.Field
		TotalRecords respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcPhoneNumberCampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcPhoneNumberCampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPhoneNumberCampaignNewParams struct {
	PhoneNumberCampaignCreate PhoneNumberCampaignCreateParam
	paramObj
}

func (r Number10dlcPhoneNumberCampaignNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PhoneNumberCampaignCreate)
}
func (r *Number10dlcPhoneNumberCampaignNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PhoneNumberCampaignCreate)
}

type Number10dlcPhoneNumberCampaignUpdateParams struct {
	PhoneNumberCampaignCreate PhoneNumberCampaignCreateParam
	paramObj
}

func (r Number10dlcPhoneNumberCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PhoneNumberCampaignCreate)
}
func (r *Number10dlcPhoneNumberCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PhoneNumberCampaignCreate)
}

type Number10dlcPhoneNumberCampaignListParams struct {
	Page           param.Opt[int64] `query:"page,omitzero" json:"-"`
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[telnyx_campaign_id], filter[telnyx_brand_id], filter[tcr_campaign_id],
	// filter[tcr_brand_id]
	Filter Number10dlcPhoneNumberCampaignListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// createdAt in descending order.
	//
	// Any of "assignmentStatus", "-assignmentStatus", "createdAt", "-createdAt",
	// "phoneNumber", "-phoneNumber".
	Sort Number10dlcPhoneNumberCampaignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcPhoneNumberCampaignListParams]'s query
// parameters as `url.Values`.
func (r Number10dlcPhoneNumberCampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[telnyx_campaign_id], filter[telnyx_brand_id], filter[tcr_campaign_id],
// filter[tcr_brand_id]
type Number10dlcPhoneNumberCampaignListParamsFilter struct {
	// Filter results by the TCR Brand id
	TcrBrandID param.Opt[string] `query:"tcr_brand_id,omitzero" json:"-"`
	// Filter results by the TCR Campaign id
	TcrCampaignID param.Opt[string] `query:"tcr_campaign_id,omitzero" json:"-"`
	// Filter results by the Telnyx Brand id
	TelnyxBrandID param.Opt[string] `query:"telnyx_brand_id,omitzero" format:"uuid" json:"-"`
	// Filter results by the Telnyx Campaign id
	TelnyxCampaignID param.Opt[string] `query:"telnyx_campaign_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcPhoneNumberCampaignListParamsFilter]'s query
// parameters as `url.Values`.
func (r Number10dlcPhoneNumberCampaignListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type Number10dlcPhoneNumberCampaignListParamsSort string

const (
	Number10dlcPhoneNumberCampaignListParamsSortAssignmentStatus     Number10dlcPhoneNumberCampaignListParamsSort = "assignmentStatus"
	Number10dlcPhoneNumberCampaignListParamsSortAssignmentStatusDesc Number10dlcPhoneNumberCampaignListParamsSort = "-assignmentStatus"
	Number10dlcPhoneNumberCampaignListParamsSortCreatedAt            Number10dlcPhoneNumberCampaignListParamsSort = "createdAt"
	Number10dlcPhoneNumberCampaignListParamsSortCreatedAtDesc        Number10dlcPhoneNumberCampaignListParamsSort = "-createdAt"
	Number10dlcPhoneNumberCampaignListParamsSortPhoneNumber          Number10dlcPhoneNumberCampaignListParamsSort = "phoneNumber"
	Number10dlcPhoneNumberCampaignListParamsSortPhoneNumberDesc      Number10dlcPhoneNumberCampaignListParamsSort = "-phoneNumber"
)
