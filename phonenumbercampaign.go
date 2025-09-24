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

// PhoneNumberCampaignService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberCampaignService] method instead.
type PhoneNumberCampaignService struct {
	Options []option.RequestOption
}

// NewPhoneNumberCampaignService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberCampaignService(opts ...option.RequestOption) (r PhoneNumberCampaignService) {
	r = PhoneNumberCampaignService{}
	r.Options = opts
	return
}

// Create New Phone Number Campaign
func (r *PhoneNumberCampaignService) New(ctx context.Context, body PhoneNumberCampaignNewParams, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_number_campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an individual phone number/campaign assignment by `phoneNumber`.
func (r *PhoneNumberCampaignService) Get(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create New Phone Number Campaign
func (r *PhoneNumberCampaignService) Update(ctx context.Context, phoneNumber string, body PhoneNumberCampaignUpdateParams, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve All Phone Number Campaigns
func (r *PhoneNumberCampaignService) List(ctx context.Context, query PhoneNumberCampaignListParams, opts ...option.RequestOption) (res *PhoneNumberCampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_number_campaigns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint allows you to remove a campaign assignment from the supplied
// `phoneNumber`.
func (r *PhoneNumberCampaignService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *PhoneNumberCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phoneNumber parameter")
		return
	}
	path := fmt.Sprintf("phone_number_campaigns/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PhoneNumberCampaign struct {
	// For shared campaigns, this is the TCR campaign ID, otherwise it is the campaign
	// ID
	CampaignID  string `json:"campaignId,required"`
	CreatedAt   string `json:"createdAt,required"`
	PhoneNumber string `json:"phoneNumber,required"`
	UpdatedAt   string `json:"updatedAt,required"`
	// The assignment status of the number.
	//
	// Any of "FAILED_ASSIGNMENT", "PENDING_ASSIGNMENT", "ASSIGNED",
	// "PENDING_UNASSIGNMENT", "FAILED_UNASSIGNMENT".
	AssignmentStatus PhoneNumberCampaignAssignmentStatus `json:"assignmentStatus"`
	// Brand ID. Empty if the number is associated to a shared campaign.
	BrandID string `json:"brandId"`
	// Extra info about a failure to assign/unassign a number. Relevant only if the
	// assignmentStatus is either FAILED_ASSIGNMENT or FAILED_UNASSIGNMENT
	FailureReasons string `json:"failureReasons"`
	// TCR's alphanumeric ID for the brand.
	TcrBrandID string `json:"tcrBrandId"`
	// TCR's alphanumeric ID for the campaign.
	TcrCampaignID string `json:"tcrCampaignId"`
	// Campaign ID. Empty if the number is associated to a shared campaign.
	TelnyxCampaignID string `json:"telnyxCampaignId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CampaignID       respjson.Field
		CreatedAt        respjson.Field
		PhoneNumber      respjson.Field
		UpdatedAt        respjson.Field
		AssignmentStatus respjson.Field
		BrandID          respjson.Field
		FailureReasons   respjson.Field
		TcrBrandID       respjson.Field
		TcrCampaignID    respjson.Field
		TelnyxCampaignID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCampaign) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The assignment status of the number.
type PhoneNumberCampaignAssignmentStatus string

const (
	PhoneNumberCampaignAssignmentStatusFailedAssignment    PhoneNumberCampaignAssignmentStatus = "FAILED_ASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusPendingAssignment   PhoneNumberCampaignAssignmentStatus = "PENDING_ASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusAssigned            PhoneNumberCampaignAssignmentStatus = "ASSIGNED"
	PhoneNumberCampaignAssignmentStatusPendingUnassignment PhoneNumberCampaignAssignmentStatus = "PENDING_UNASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusFailedUnassignment  PhoneNumberCampaignAssignmentStatus = "FAILED_UNASSIGNMENT"
)

// The properties CampaignID, PhoneNumber are required.
type PhoneNumberCampaignCreateParam struct {
	// The ID of the campaign you want to link to the specified phone number.
	CampaignID string `json:"campaignId,required"`
	// The phone number you want to link to a specified campaign.
	PhoneNumber string `json:"phoneNumber,required"`
	paramObj
}

func (r PhoneNumberCampaignCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberCampaignCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberCampaignCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCampaignListResponse struct {
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
func (r PhoneNumberCampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCampaignNewParams struct {
	PhoneNumberCampaignCreate PhoneNumberCampaignCreateParam
	paramObj
}

func (r PhoneNumberCampaignNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PhoneNumberCampaignCreate)
}
func (r *PhoneNumberCampaignNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PhoneNumberCampaignCreate)
}

type PhoneNumberCampaignUpdateParams struct {
	PhoneNumberCampaignCreate PhoneNumberCampaignCreateParam
	paramObj
}

func (r PhoneNumberCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PhoneNumberCampaignCreate)
}
func (r *PhoneNumberCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PhoneNumberCampaignCreate)
}

type PhoneNumberCampaignListParams struct {
	Page           param.Opt[int64] `query:"page,omitzero" json:"-"`
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[telnyx_campaign_id], filter[telnyx_brand_id], filter[tcr_campaign_id],
	// filter[tcr_brand_id]
	Filter PhoneNumberCampaignListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// createdAt in descending order.
	//
	// Any of "assignmentStatus", "-assignmentStatus", "createdAt", "-createdAt",
	// "phoneNumber", "-phoneNumber".
	Sort PhoneNumberCampaignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCampaignListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberCampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[telnyx_campaign_id], filter[telnyx_brand_id], filter[tcr_campaign_id],
// filter[tcr_brand_id]
type PhoneNumberCampaignListParamsFilter struct {
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

// URLQuery serializes [PhoneNumberCampaignListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberCampaignListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type PhoneNumberCampaignListParamsSort string

const (
	PhoneNumberCampaignListParamsSortAssignmentStatus     PhoneNumberCampaignListParamsSort = "assignmentStatus"
	PhoneNumberCampaignListParamsSortAssignmentStatusDesc PhoneNumberCampaignListParamsSort = "-assignmentStatus"
	PhoneNumberCampaignListParamsSortCreatedAt            PhoneNumberCampaignListParamsSort = "createdAt"
	PhoneNumberCampaignListParamsSortCreatedAtDesc        PhoneNumberCampaignListParamsSort = "-createdAt"
	PhoneNumberCampaignListParamsSortPhoneNumber          PhoneNumberCampaignListParamsSort = "phoneNumber"
	PhoneNumberCampaignListParamsSortPhoneNumberDesc      PhoneNumberCampaignListParamsSort = "-phoneNumber"
)
