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
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Messaging10dlcPartnerCampaignService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessaging10dlcPartnerCampaignService] method instead.
type Messaging10dlcPartnerCampaignService struct {
	Options []option.RequestOption
}

// NewMessaging10dlcPartnerCampaignService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessaging10dlcPartnerCampaignService(opts ...option.RequestOption) (r Messaging10dlcPartnerCampaignService) {
	r = Messaging10dlcPartnerCampaignService{}
	r.Options = opts
	return
}

// Retrieve campaign details by `campaignId`.
func (r *Messaging10dlcPartnerCampaignService) Get(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *TelnyxDownstreamCampaign, err error) {
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
func (r *Messaging10dlcPartnerCampaignService) Update(ctx context.Context, campaignID string, body Messaging10dlcPartnerCampaignUpdateParams, opts ...option.RequestOption) (res *TelnyxDownstreamCampaign, err error) {
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
func (r *Messaging10dlcPartnerCampaignService) List(ctx context.Context, query Messaging10dlcPartnerCampaignListParams, opts ...option.RequestOption) (res *pagination.PerPagePaginationV2[TelnyxDownstreamCampaign], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "10dlc/partner_campaigns"
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

// Retrieve all partner campaigns you have shared to Telnyx in a paginated fashion.
//
// This endpoint is currently limited to only returning shared campaigns that
// Telnyx has accepted. In other words, shared but pending campaigns are currently
// omitted from the response from this endpoint.
func (r *Messaging10dlcPartnerCampaignService) ListAutoPaging(ctx context.Context, query Messaging10dlcPartnerCampaignListParams, opts ...option.RequestOption) *pagination.PerPagePaginationV2AutoPager[TelnyxDownstreamCampaign] {
	return pagination.NewPerPagePaginationV2AutoPager(r.List(ctx, query, opts...))
}

// Get all partner campaigns you have shared to Telnyx in a paginated fashion
//
// This endpoint is currently limited to only returning shared campaigns that
// Telnyx has accepted. In other words, shared but pending campaigns are currently
// omitted from the response from this endpoint.
func (r *Messaging10dlcPartnerCampaignService) ListSharedByMe(ctx context.Context, query Messaging10dlcPartnerCampaignListSharedByMeParams, opts ...option.RequestOption) (res *pagination.PerPagePaginationV2[Messaging10dlcPartnerCampaignListSharedByMeResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "10dlc/partnerCampaign/sharedByMe"
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

// Get all partner campaigns you have shared to Telnyx in a paginated fashion
//
// This endpoint is currently limited to only returning shared campaigns that
// Telnyx has accepted. In other words, shared but pending campaigns are currently
// omitted from the response from this endpoint.
func (r *Messaging10dlcPartnerCampaignService) ListSharedByMeAutoPaging(ctx context.Context, query Messaging10dlcPartnerCampaignListSharedByMeParams, opts ...option.RequestOption) *pagination.PerPagePaginationV2AutoPager[Messaging10dlcPartnerCampaignListSharedByMeResponse] {
	return pagination.NewPerPagePaginationV2AutoPager(r.ListSharedByMe(ctx, query, opts...))
}

// Get Sharing Status
func (r *Messaging10dlcPartnerCampaignService) GetSharingStatus(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Messaging10dlcPartnerCampaignGetSharingStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/partnerCampaign/%s/sharing", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Campaign is generated by the 10DLC registry once the corresponding campaign
// request is approved. Each campaign is assigned a unique identifier -
// **campaignId**. Once a campaign is activated, limited information is published
// to the NetNumber OSR service for consumption by members of the ecosystem. When a
// campaign is suspended(reversible) or expired(non-reversible), campaign data is
// deleted from the OSR service.
type TelnyxDownstreamCampaign struct {
	// Unique identifier assigned to the brand by the registry.
	TcrBrandID string `json:"tcrBrandId,required"`
	// Unique identifier assigned to the campaign by the registry.
	TcrCampaignID string `json:"tcrCampaignId,required"`
	// Age gated content in campaign.
	AgeGated bool `json:"ageGated"`
	// Number of phone numbers associated with the campaign
	AssignedPhoneNumbersCount float64 `json:"assignedPhoneNumbersCount"`
	// Display or marketing name of the brand.
	BrandDisplayName string `json:"brandDisplayName"`
	// Campaign status
	//
	// Any of "TCR_PENDING", "TCR_SUSPENDED", "TCR_EXPIRED", "TCR_ACCEPTED",
	// "TCR_FAILED", "TELNYX_ACCEPTED", "TELNYX_FAILED", "MNO_PENDING", "MNO_ACCEPTED",
	// "MNO_REJECTED", "MNO_PROVISIONED", "MNO_PROVISIONING_FAILED".
	CampaignStatus TelnyxDownstreamCampaignCampaignStatus `json:"campaignStatus"`
	// Date and time that the brand was created at.
	CreatedAt string `json:"createdAt"`
	// Summary description of this campaign.
	Description string `json:"description"`
	// Direct lending or loan arrangement.
	DirectLending bool `json:"directLending"`
	// Does message generated by the campaign include URL link in SMS?
	EmbeddedLink bool `json:"embeddedLink"`
	// Sample of an embedded link that will be sent to subscribers.
	EmbeddedLinkSample string `json:"embeddedLinkSample"`
	// Does message generated by the campaign include phone number in SMS?
	EmbeddedPhone bool `json:"embeddedPhone"`
	// Failure reasons if campaign submission failed
	FailureReasons string `json:"failureReasons"`
	// Subscriber help keywords. Multiple keywords are comma separated without space.
	HelpKeywords string `json:"helpKeywords"`
	// Help message of the campaign.
	HelpMessage string `json:"helpMessage"`
	// Indicates whether the campaign has a T-Mobile number pool ID associated with it.
	IsNumberPoolingEnabled bool `json:"isNumberPoolingEnabled"`
	// Message flow description.
	MessageFlow string `json:"messageFlow"`
	// Does campaign utilize pool of phone numbers?
	NumberPool bool `json:"numberPool"`
	// Subscriber opt-in keywords. Multiple keywords are comma separated without space.
	OptinKeywords string `json:"optinKeywords"`
	// Subscriber opt-in message.
	OptinMessage string `json:"optinMessage"`
	// Subscriber opt-out keywords. Multiple keywords are comma separated without
	// space.
	OptoutKeywords string `json:"optoutKeywords"`
	// Subscriber opt-out message.
	OptoutMessage string `json:"optoutMessage"`
	// Link to the campaign's privacy policy.
	PrivacyPolicyLink string `json:"privacyPolicyLink"`
	// Message sample. Some campaign tiers require 1 or more message samples.
	Sample1 string `json:"sample1"`
	// Message sample. Some campaign tiers require 2 or more message samples.
	Sample2 string `json:"sample2"`
	// Message sample. Some campaign tiers require 3 or more message samples.
	Sample3 string `json:"sample3"`
	// Message sample. Some campaign tiers require 4 or more message samples.
	Sample4 string `json:"sample4"`
	// Message sample. Some campaign tiers require 5 or more message samples.
	Sample5 string `json:"sample5"`
	// Does campaign require subscriber to opt-in before SMS is sent to subscriber?
	SubscriberOptin bool `json:"subscriberOptin"`
	// Does campaign support subscriber opt-out keyword(s)?
	SubscriberOptout bool `json:"subscriberOptout"`
	// Campaign sub-usecases. Must be of defined valid sub-usecase types. Use
	// `/10dlc/enum/usecase` operation to retrieve list of valid sub-usecases
	SubUsecases []string `json:"subUsecases"`
	// Is terms & conditions accepted?
	TermsAndConditions bool `json:"termsAndConditions"`
	// Link to the campaign's terms and conditions.
	TermsAndConditionsLink string `json:"termsAndConditionsLink"`
	// Date and time that the brand was last updated at.
	UpdatedAt string `json:"updatedAt"`
	// Campaign usecase. Must be of defined valid types. Use `/10dlc/enum/usecase`
	// operation to retrieve usecases available for given brand.
	Usecase string `json:"usecase"`
	// Failover webhook to which campaign status updates are sent.
	WebhookFailoverURL string `json:"webhookFailoverURL"`
	// Webhook to which campaign status updates are sent.
	WebhookURL string `json:"webhookURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TcrBrandID                respjson.Field
		TcrCampaignID             respjson.Field
		AgeGated                  respjson.Field
		AssignedPhoneNumbersCount respjson.Field
		BrandDisplayName          respjson.Field
		CampaignStatus            respjson.Field
		CreatedAt                 respjson.Field
		Description               respjson.Field
		DirectLending             respjson.Field
		EmbeddedLink              respjson.Field
		EmbeddedLinkSample        respjson.Field
		EmbeddedPhone             respjson.Field
		FailureReasons            respjson.Field
		HelpKeywords              respjson.Field
		HelpMessage               respjson.Field
		IsNumberPoolingEnabled    respjson.Field
		MessageFlow               respjson.Field
		NumberPool                respjson.Field
		OptinKeywords             respjson.Field
		OptinMessage              respjson.Field
		OptoutKeywords            respjson.Field
		OptoutMessage             respjson.Field
		PrivacyPolicyLink         respjson.Field
		Sample1                   respjson.Field
		Sample2                   respjson.Field
		Sample3                   respjson.Field
		Sample4                   respjson.Field
		Sample5                   respjson.Field
		SubscriberOptin           respjson.Field
		SubscriberOptout          respjson.Field
		SubUsecases               respjson.Field
		TermsAndConditions        respjson.Field
		TermsAndConditionsLink    respjson.Field
		UpdatedAt                 respjson.Field
		Usecase                   respjson.Field
		WebhookFailoverURL        respjson.Field
		WebhookURL                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelnyxDownstreamCampaign) RawJSON() string { return r.JSON.raw }
func (r *TelnyxDownstreamCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign status
type TelnyxDownstreamCampaignCampaignStatus string

const (
	TelnyxDownstreamCampaignCampaignStatusTcrPending            TelnyxDownstreamCampaignCampaignStatus = "TCR_PENDING"
	TelnyxDownstreamCampaignCampaignStatusTcrSuspended          TelnyxDownstreamCampaignCampaignStatus = "TCR_SUSPENDED"
	TelnyxDownstreamCampaignCampaignStatusTcrExpired            TelnyxDownstreamCampaignCampaignStatus = "TCR_EXPIRED"
	TelnyxDownstreamCampaignCampaignStatusTcrAccepted           TelnyxDownstreamCampaignCampaignStatus = "TCR_ACCEPTED"
	TelnyxDownstreamCampaignCampaignStatusTcrFailed             TelnyxDownstreamCampaignCampaignStatus = "TCR_FAILED"
	TelnyxDownstreamCampaignCampaignStatusTelnyxAccepted        TelnyxDownstreamCampaignCampaignStatus = "TELNYX_ACCEPTED"
	TelnyxDownstreamCampaignCampaignStatusTelnyxFailed          TelnyxDownstreamCampaignCampaignStatus = "TELNYX_FAILED"
	TelnyxDownstreamCampaignCampaignStatusMnoPending            TelnyxDownstreamCampaignCampaignStatus = "MNO_PENDING"
	TelnyxDownstreamCampaignCampaignStatusMnoAccepted           TelnyxDownstreamCampaignCampaignStatus = "MNO_ACCEPTED"
	TelnyxDownstreamCampaignCampaignStatusMnoRejected           TelnyxDownstreamCampaignCampaignStatus = "MNO_REJECTED"
	TelnyxDownstreamCampaignCampaignStatusMnoProvisioned        TelnyxDownstreamCampaignCampaignStatus = "MNO_PROVISIONED"
	TelnyxDownstreamCampaignCampaignStatusMnoProvisioningFailed TelnyxDownstreamCampaignCampaignStatus = "MNO_PROVISIONING_FAILED"
)

// Campaign is generated by the 10DLC registry once the corresponding campaign
// request is approved. Each campaign is assigned a unique identifier -
// **campaignId**. Once a campaign is activated, limited information is published
// to the NetNumber OSR service for consumption by members of the ecosystem. When a
// campaign is suspended(reversible) or expired(non-reversible), campaign data is
// deleted from the OSR service. Most attributes of campaignare immutable,
// including **usecase**, **vertical**, **brandId** and **cspId**.
type Messaging10dlcPartnerCampaignListSharedByMeResponse struct {
	// Alphanumeric identifier of the brand associated with this campaign.
	BrandID string `json:"brandId,required"`
	// Alphanumeric identifier assigned by the registry for a campaign. This identifier
	// is required by the NetNumber OSR SMS enabling process of 10DLC.
	CampaignID string `json:"campaignId,required"`
	// Campaign usecase. Must be of defined valid types. Use `/10dlc/enum/usecase`
	// operation to retrieve usecases available for given brand.
	Usecase string `json:"usecase,required"`
	// Unix timestamp when campaign was created.
	CreateDate string `json:"createDate"`
	// Current campaign status. Possible values: ACTIVE, EXPIRED. A newly created
	// campaign defaults to ACTIVE status.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID     respjson.Field
		CampaignID  respjson.Field
		Usecase     respjson.Field
		CreateDate  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Messaging10dlcPartnerCampaignListSharedByMeResponse) RawJSON() string { return r.JSON.raw }
func (r *Messaging10dlcPartnerCampaignListSharedByMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcPartnerCampaignGetSharingStatusResponse map[string]CampaignSharingStatus

type Messaging10dlcPartnerCampaignUpdateParams struct {
	// Webhook failover to which campaign status updates are sent.
	WebhookFailoverURL param.Opt[string] `json:"webhookFailoverURL,omitzero"`
	// Webhook to which campaign status updates are sent.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	paramObj
}

func (r Messaging10dlcPartnerCampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow Messaging10dlcPartnerCampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Messaging10dlcPartnerCampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcPartnerCampaignListParams struct {
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
	Sort Messaging10dlcPartnerCampaignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Messaging10dlcPartnerCampaignListParams]'s query parameters
// as `url.Values`.
func (r Messaging10dlcPartnerCampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type Messaging10dlcPartnerCampaignListParamsSort string

const (
	Messaging10dlcPartnerCampaignListParamsSortAssignedPhoneNumbersCount     Messaging10dlcPartnerCampaignListParamsSort = "assignedPhoneNumbersCount"
	Messaging10dlcPartnerCampaignListParamsSortAssignedPhoneNumbersCountDesc Messaging10dlcPartnerCampaignListParamsSort = "-assignedPhoneNumbersCount"
	Messaging10dlcPartnerCampaignListParamsSortBrandDisplayName              Messaging10dlcPartnerCampaignListParamsSort = "brandDisplayName"
	Messaging10dlcPartnerCampaignListParamsSortBrandDisplayNameDesc          Messaging10dlcPartnerCampaignListParamsSort = "-brandDisplayName"
	Messaging10dlcPartnerCampaignListParamsSortTcrBrandID                    Messaging10dlcPartnerCampaignListParamsSort = "tcrBrandId"
	Messaging10dlcPartnerCampaignListParamsSortTcrBranIDDesc                 Messaging10dlcPartnerCampaignListParamsSort = "-tcrBranId"
	Messaging10dlcPartnerCampaignListParamsSortTcrCampaignID                 Messaging10dlcPartnerCampaignListParamsSort = "tcrCampaignId"
	Messaging10dlcPartnerCampaignListParamsSortTcrCampaignIDDesc             Messaging10dlcPartnerCampaignListParamsSort = "-tcrCampaignId"
	Messaging10dlcPartnerCampaignListParamsSortCreatedAt                     Messaging10dlcPartnerCampaignListParamsSort = "createdAt"
	Messaging10dlcPartnerCampaignListParamsSortCreatedAtDesc                 Messaging10dlcPartnerCampaignListParamsSort = "-createdAt"
	Messaging10dlcPartnerCampaignListParamsSortCampaignStatus                Messaging10dlcPartnerCampaignListParamsSort = "campaignStatus"
	Messaging10dlcPartnerCampaignListParamsSortCampaignStatusDesc            Messaging10dlcPartnerCampaignListParamsSort = "-campaignStatus"
)

type Messaging10dlcPartnerCampaignListSharedByMeParams struct {
	// The 1-indexed page number to get. The default value is `1`.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The amount of records per page, limited to between 1 and 500 inclusive. The
	// default value is `10`.
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Messaging10dlcPartnerCampaignListSharedByMeParams]'s query
// parameters as `url.Values`.
func (r Messaging10dlcPartnerCampaignListSharedByMeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
