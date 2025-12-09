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
)

// Number10dlcPartnerCampaignService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcPartnerCampaignService] method instead.
type Number10dlcPartnerCampaignService struct {
	Options []option.RequestOption
}

// NewNumber10dlcPartnerCampaignService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNumber10dlcPartnerCampaignService(opts ...option.RequestOption) (r Number10dlcPartnerCampaignService) {
	r = Number10dlcPartnerCampaignService{}
	r.Options = opts
	return
}

// Retrieve campaign details by `campaignId`.
func (r *Number10dlcPartnerCampaignService) Get(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *TelnyxDownstreamCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/partner_campaigns/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update campaign details by `campaignId`. **Please note:** Only webhook urls are
// editable.
func (r *Number10dlcPartnerCampaignService) Update(ctx context.Context, campaignID string, body Number10dlcPartnerCampaignUpdateParams, opts ...option.RequestOption) (res *TelnyxDownstreamCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/partner_campaigns/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve all partner campaigns you have shared to Telnyx in a paginated fashion.
//
// This endpoint is currently limited to only returning shared campaigns that
// Telnyx has accepted. In other words, shared but pending campaigns are currently
// omitted from the response from this endpoint.
func (r *Number10dlcPartnerCampaignService) List(ctx context.Context, query Number10dlcPartnerCampaignListParams, opts ...option.RequestOption) (res *Number10dlcPartnerCampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/partner_campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Number10dlcPartnerCampaignListResponse struct {
	Page         int64                      `json:"page"`
	Records      []TelnyxDownstreamCampaign `json:"records"`
	TotalRecords int64                      `json:"totalRecords"`
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
func (r Number10dlcPartnerCampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcPartnerCampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPartnerCampaignUpdateParams struct {
	// Webhook failover to which campaign status updates are sent.
	WebhookFailoverURL param.Opt[string] `json:"webhookFailoverURL,omitzero"`
	// Webhook to which campaign status updates are sent.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	paramObj
}

func (r Number10dlcPartnerCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcPartnerCampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcPartnerCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPartnerCampaignListParams struct {
	// The 1-indexed page number to get. The default value is `1`.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The amount of records per page, limited to between 1 and 500 inclusive. The
	// default value is `10`.
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// createdAt in descending order.
	//
	// Any of "assignedPhoneNumbersCount", "-assignedPhoneNumbersCount",
	// "brandDisplayName", "-brandDisplayName", "tcrBrandId", "-tcrBranId",
	// "tcrCampaignId", "-tcrCampaignId", "createdAt", "-createdAt", "campaignStatus",
	// "-campaignStatus".
	Sort Number10dlcPartnerCampaignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcPartnerCampaignListParams]'s query parameters as
// `url.Values`.
func (r Number10dlcPartnerCampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type Number10dlcPartnerCampaignListParamsSort string

const (
	Number10dlcPartnerCampaignListParamsSortAssignedPhoneNumbersCount     Number10dlcPartnerCampaignListParamsSort = "assignedPhoneNumbersCount"
	Number10dlcPartnerCampaignListParamsSortAssignedPhoneNumbersCountDesc Number10dlcPartnerCampaignListParamsSort = "-assignedPhoneNumbersCount"
	Number10dlcPartnerCampaignListParamsSortBrandDisplayName              Number10dlcPartnerCampaignListParamsSort = "brandDisplayName"
	Number10dlcPartnerCampaignListParamsSortBrandDisplayNameDesc          Number10dlcPartnerCampaignListParamsSort = "-brandDisplayName"
	Number10dlcPartnerCampaignListParamsSortTcrBrandID                    Number10dlcPartnerCampaignListParamsSort = "tcrBrandId"
	Number10dlcPartnerCampaignListParamsSortTcrBranIDDesc                 Number10dlcPartnerCampaignListParamsSort = "-tcrBranId"
	Number10dlcPartnerCampaignListParamsSortTcrCampaignID                 Number10dlcPartnerCampaignListParamsSort = "tcrCampaignId"
	Number10dlcPartnerCampaignListParamsSortTcrCampaignIDDesc             Number10dlcPartnerCampaignListParamsSort = "-tcrCampaignId"
	Number10dlcPartnerCampaignListParamsSortCreatedAt                     Number10dlcPartnerCampaignListParamsSort = "createdAt"
	Number10dlcPartnerCampaignListParamsSortCreatedAtDesc                 Number10dlcPartnerCampaignListParamsSort = "-createdAt"
	Number10dlcPartnerCampaignListParamsSortCampaignStatus                Number10dlcPartnerCampaignListParamsSort = "campaignStatus"
	Number10dlcPartnerCampaignListParamsSortCampaignStatusDesc            Number10dlcPartnerCampaignListParamsSort = "-campaignStatus"
)
