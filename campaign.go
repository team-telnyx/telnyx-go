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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// CampaignService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignService] method instead.
type CampaignService struct {
	Options []option.RequestOption
	Usecase CampaignUsecaseService
	Osr     CampaignOsrService
}

// NewCampaignService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignService(opts ...option.RequestOption) (r CampaignService) {
	r = CampaignService{}
	r.Options = opts
	r.Usecase = NewCampaignUsecaseService(opts...)
	r.Osr = NewCampaignOsrService(opts...)
	return
}

// Retrieve campaign details by `campaignId`.
func (r *CampaignService) Get(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *TelnyxCampaignCsp, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a campaign's properties by `campaignId`. **Please note:** only sample
// messages are editable.
func (r *CampaignService) Update(ctx context.Context, campaignID string, body CampaignUpdateParams, opts ...option.RequestOption) (res *TelnyxCampaignCsp, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of campaigns associated with a supplied `brandId`.
func (r *CampaignService) List(ctx context.Context, query CampaignListParams, opts ...option.RequestOption) (res *CampaignListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "campaign"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Manually accept a campaign shared with Telnyx
func (r *CampaignService) AcceptSharing(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignAcceptSharingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/acceptSharing/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Terminate a campaign. Note that once deactivated, a campaign cannot be restored.
func (r *CampaignService) Deactivate(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignDeactivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the campaign metadata for each MNO it was submitted to.
func (r *CampaignService) GetMnoMetadata(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignGetMnoMetadataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s/mnoMetadata", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve campaign's operation status at MNO level.
func (r *CampaignService) GetOperationStatus(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignGetOperationStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s/operationStatus", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Sharing Status
func (r *CampaignService) GetSharingStatus(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignGetSharingStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s/sharing", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits an appeal for rejected native campaigns in TELNYX_FAILED or MNO_REJECTED
// status. The appeal is recorded for manual compliance team review and the
// campaign status is reset to TCR_ACCEPTED. Note: Appeal forwarding is handled
// manually to allow proper review before incurring upstream charges.
func (r *CampaignService) SubmitAppeal(ctx context.Context, campaignID string, body CampaignSubmitAppealParams, opts ...option.RequestOption) (res *CampaignSubmitAppealResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s/appeal", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CampaignSharingStatus struct {
	DownstreamCnpID string `json:"downstreamCnpId"`
	SharedDate      string `json:"sharedDate"`
	SharingStatus   string `json:"sharingStatus"`
	StatusDate      string `json:"statusDate"`
	UpstreamCnpID   string `json:"upstreamCnpId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownstreamCnpID respjson.Field
		SharedDate      respjson.Field
		SharingStatus   respjson.Field
		StatusDate      respjson.Field
		UpstreamCnpID   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignSharingStatus) RawJSON() string { return r.JSON.raw }
func (r *CampaignSharingStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign is generated by the 10DLC registry once the corresponding campaign
// request is approved. Each campaign is assigned a unique identifier -
// **campaignId**. Once a campaign is activated, limited information is published
// to the NetNumber OSR service for consumption by members of the ecosystem. When a
// campaign is suspended(reversible) or expired(non-reversible), campaign data is
// deleted from the OSR service. Most attributes of campaignare immutable,
// including **usecase**, **vertical**, **brandId** and **cspId**.
type TelnyxCampaignCsp struct {
	// Unique identifier assigned to the brand.
	BrandID string `json:"brandId,required"`
	// Unique identifier for a campaign.
	CampaignID string `json:"campaignId,required"`
	// Alphanumeric identifier of the CSP associated with this campaign.
	CspID string `json:"cspId,required"`
	// Summary description of this campaign.
	Description string `json:"description,required"`
	// Campaign created from mock brand. Mocked campaign cannot be shared with an
	// upstream CNP.
	Mock bool `json:"mock,required"`
	// Campaign sub-usecases. Must be of defined valid sub-usecase types. Use
	// `/registry/enum/usecase` operation to retrieve list of valid sub-usecases
	SubUsecases []string `json:"subUsecases,required"`
	// Is terms & conditions accepted?
	TermsAndConditions bool `json:"termsAndConditions,required"`
	// Campaign usecase. Must be of defined valid types. Use `/registry/enum/usecase`
	// operation to retrieve usecases available for given brand.
	Usecase string `json:"usecase,required"`
	// Age gated content in campaign.
	AgeGated bool `json:"ageGated"`
	// Campaign subscription auto-renewal status.
	AutoRenewal bool `json:"autoRenewal"`
	// Campaign recent billed date.
	BilledDate string `json:"billedDate"`
	// Display or marketing name of the brand.
	BrandDisplayName string `json:"brandDisplayName"`
	// Campaign status
	//
	// Any of "TCR_PENDING", "TCR_SUSPENDED", "TCR_EXPIRED", "TCR_ACCEPTED",
	// "TCR_FAILED", "TELNYX_ACCEPTED", "TELNYX_FAILED", "MNO_PENDING", "MNO_ACCEPTED",
	// "MNO_REJECTED", "MNO_PROVISIONED", "MNO_PROVISIONING_FAILED".
	CampaignStatus TelnyxCampaignCspCampaignStatus `json:"campaignStatus"`
	// Unix timestamp when campaign was created.
	CreateDate    string `json:"createDate"`
	DirectLending bool   `json:"directLending"`
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
	IsTMobileNumberPoolingEnabled bool `json:"isTMobileNumberPoolingEnabled"`
	// Indicates whether the campaign is registered with T-Mobile.
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// Indicates whether the campaign is suspended with T-Mobile.
	IsTMobileSuspended bool `json:"isTMobileSuspended"`
	// Message flow description.
	MessageFlow string `json:"messageFlow"`
	// When the campaign would be due for its next renew/bill date.
	NextRenewalOrExpirationDate string `json:"nextRenewalOrExpirationDate"`
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
	// Caller supplied campaign reference ID. If supplied, the value must be unique
	// across all submitted campaigns. Can be used to prevent duplicate campaign
	// registrations.
	ReferenceID string `json:"referenceId"`
	// Alphanumeric identifier of the reseller that you want to associate with this
	// campaign.
	ResellerID string `json:"resellerId"`
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
	// Current campaign status. Possible values: ACTIVE, EXPIRED. A newly created
	// campaign defaults to ACTIVE status.
	Status string `json:"status"`
	// Campaign submission status
	//
	// Any of "CREATED", "FAILED", "PENDING".
	SubmissionStatus TelnyxCampaignCspSubmissionStatus `json:"submissionStatus"`
	// Does campaign responds to help keyword(s)?
	SubscriberHelp bool `json:"subscriberHelp"`
	// Does campaign require subscriber to opt-in before SMS is sent to subscriber?
	SubscriberOptin bool `json:"subscriberOptin"`
	// Does campaign support subscriber opt-out keyword(s)?
	SubscriberOptout bool `json:"subscriberOptout"`
	// Unique identifier assigned to the brand by the registry.
	TcrBrandID string `json:"tcrBrandId"`
	// Unique identifier assigned to the campaign by the registry.
	TcrCampaignID string `json:"tcrCampaignId"`
	// Link to the campaign's terms and conditions.
	TermsAndConditionsLink string `json:"termsAndConditionsLink"`
	// Business/industry segment of this campaign (Deprecated). Must be of defined
	// valid types. Use `/registry/enum/vertical` operation to retrieve verticals
	// available for given brand, vertical combination.
	//
	// This field is deprecated.
	//
	// Deprecated: deprecated
	Vertical string `json:"vertical"`
	// Failover webhook to which campaign status updates are sent.
	WebhookFailoverURL string `json:"webhookFailoverURL"`
	// Webhook to which campaign status updates are sent.
	WebhookURL string `json:"webhookURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID                       respjson.Field
		CampaignID                    respjson.Field
		CspID                         respjson.Field
		Description                   respjson.Field
		Mock                          respjson.Field
		SubUsecases                   respjson.Field
		TermsAndConditions            respjson.Field
		Usecase                       respjson.Field
		AgeGated                      respjson.Field
		AutoRenewal                   respjson.Field
		BilledDate                    respjson.Field
		BrandDisplayName              respjson.Field
		CampaignStatus                respjson.Field
		CreateDate                    respjson.Field
		DirectLending                 respjson.Field
		EmbeddedLink                  respjson.Field
		EmbeddedLinkSample            respjson.Field
		EmbeddedPhone                 respjson.Field
		FailureReasons                respjson.Field
		HelpKeywords                  respjson.Field
		HelpMessage                   respjson.Field
		IsTMobileNumberPoolingEnabled respjson.Field
		IsTMobileRegistered           respjson.Field
		IsTMobileSuspended            respjson.Field
		MessageFlow                   respjson.Field
		NextRenewalOrExpirationDate   respjson.Field
		NumberPool                    respjson.Field
		OptinKeywords                 respjson.Field
		OptinMessage                  respjson.Field
		OptoutKeywords                respjson.Field
		OptoutMessage                 respjson.Field
		PrivacyPolicyLink             respjson.Field
		ReferenceID                   respjson.Field
		ResellerID                    respjson.Field
		Sample1                       respjson.Field
		Sample2                       respjson.Field
		Sample3                       respjson.Field
		Sample4                       respjson.Field
		Sample5                       respjson.Field
		Status                        respjson.Field
		SubmissionStatus              respjson.Field
		SubscriberHelp                respjson.Field
		SubscriberOptin               respjson.Field
		SubscriberOptout              respjson.Field
		TcrBrandID                    respjson.Field
		TcrCampaignID                 respjson.Field
		TermsAndConditionsLink        respjson.Field
		Vertical                      respjson.Field
		WebhookFailoverURL            respjson.Field
		WebhookURL                    respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelnyxCampaignCsp) RawJSON() string { return r.JSON.raw }
func (r *TelnyxCampaignCsp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Campaign status
type TelnyxCampaignCspCampaignStatus string

const (
	TelnyxCampaignCspCampaignStatusTcrPending            TelnyxCampaignCspCampaignStatus = "TCR_PENDING"
	TelnyxCampaignCspCampaignStatusTcrSuspended          TelnyxCampaignCspCampaignStatus = "TCR_SUSPENDED"
	TelnyxCampaignCspCampaignStatusTcrExpired            TelnyxCampaignCspCampaignStatus = "TCR_EXPIRED"
	TelnyxCampaignCspCampaignStatusTcrAccepted           TelnyxCampaignCspCampaignStatus = "TCR_ACCEPTED"
	TelnyxCampaignCspCampaignStatusTcrFailed             TelnyxCampaignCspCampaignStatus = "TCR_FAILED"
	TelnyxCampaignCspCampaignStatusTelnyxAccepted        TelnyxCampaignCspCampaignStatus = "TELNYX_ACCEPTED"
	TelnyxCampaignCspCampaignStatusTelnyxFailed          TelnyxCampaignCspCampaignStatus = "TELNYX_FAILED"
	TelnyxCampaignCspCampaignStatusMnoPending            TelnyxCampaignCspCampaignStatus = "MNO_PENDING"
	TelnyxCampaignCspCampaignStatusMnoAccepted           TelnyxCampaignCspCampaignStatus = "MNO_ACCEPTED"
	TelnyxCampaignCspCampaignStatusMnoRejected           TelnyxCampaignCspCampaignStatus = "MNO_REJECTED"
	TelnyxCampaignCspCampaignStatusMnoProvisioned        TelnyxCampaignCspCampaignStatus = "MNO_PROVISIONED"
	TelnyxCampaignCspCampaignStatusMnoProvisioningFailed TelnyxCampaignCspCampaignStatus = "MNO_PROVISIONING_FAILED"
)

// Campaign submission status
type TelnyxCampaignCspSubmissionStatus string

const (
	TelnyxCampaignCspSubmissionStatusCreated TelnyxCampaignCspSubmissionStatus = "CREATED"
	TelnyxCampaignCspSubmissionStatusFailed  TelnyxCampaignCspSubmissionStatus = "FAILED"
	TelnyxCampaignCspSubmissionStatusPending TelnyxCampaignCspSubmissionStatus = "PENDING"
)

type CampaignListResponse struct {
	Page         int64                        `json:"page"`
	Records      []CampaignListResponseRecord `json:"records"`
	TotalRecords int64                        `json:"totalRecords"`
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
func (r CampaignListResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignListResponseRecord struct {
	// Age gated content in campaign.
	AgeGated bool `json:"ageGated"`
	// Number of phone numbers associated with the campaign
	AssignedPhoneNumbersCount float64 `json:"assignedPhoneNumbersCount"`
	// Campaign subscription auto-renewal status.
	AutoRenewal bool `json:"autoRenewal"`
	// Campaign recent billed date.
	BilledDate string `json:"billedDate"`
	// Display or marketing name of the brand.
	BrandDisplayName string `json:"brandDisplayName"`
	// Unique identifier assigned to the brand.
	BrandID string `json:"brandId"`
	// Unique identifier for a campaign.
	CampaignID string `json:"campaignId"`
	// Campaign status
	//
	// Any of "TCR_PENDING", "TCR_SUSPENDED", "TCR_EXPIRED", "TCR_ACCEPTED",
	// "TCR_FAILED", "TELNYX_ACCEPTED", "TELNYX_FAILED", "MNO_PENDING", "MNO_ACCEPTED",
	// "MNO_REJECTED", "MNO_PROVISIONED", "MNO_PROVISIONING_FAILED".
	CampaignStatus string `json:"campaignStatus"`
	// Unix timestamp when campaign was created.
	CreateDate string `json:"createDate"`
	// Alphanumeric identifier of the CSP associated with this campaign.
	CspID string `json:"cspId"`
	// Summary description of this campaign.
	Description   string `json:"description"`
	DirectLending bool   `json:"directLending"`
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
	IsTMobileNumberPoolingEnabled bool `json:"isTMobileNumberPoolingEnabled"`
	// Indicates whether the campaign is registered with T-Mobile.
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// Indicates whether the campaign is suspended with T-Mobile.
	IsTMobileSuspended bool `json:"isTMobileSuspended"`
	// Message flow description.
	MessageFlow string `json:"messageFlow"`
	// Campaign created from mock brand. Mocked campaign cannot be shared with an
	// upstream CNP.
	Mock bool `json:"mock"`
	// When the campaign would be due for its next renew/bill date.
	NextRenewalOrExpirationDate string `json:"nextRenewalOrExpirationDate"`
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
	// Caller supplied campaign reference ID. If supplied, the value must be unique
	// across all submitted campaigns. Can be used to prevent duplicate campaign
	// registrations.
	ReferenceID string `json:"referenceId"`
	// Alphanumeric identifier of the reseller that you want to associate with this
	// campaign.
	ResellerID string `json:"resellerId"`
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
	// Current campaign status. Possible values: ACTIVE, EXPIRED. A newly created
	// campaign defaults to ACTIVE status.
	Status string `json:"status"`
	// Campaign submission status
	//
	// Any of "CREATED", "FAILED", "PENDING".
	SubmissionStatus string `json:"submissionStatus"`
	// Does campaign responds to help keyword(s)?
	SubscriberHelp bool `json:"subscriberHelp"`
	// Does campaign require subscriber to opt-in before SMS is sent to subscriber?
	SubscriberOptin bool `json:"subscriberOptin"`
	// Does campaign support subscriber opt-out keyword(s)?
	SubscriberOptout bool `json:"subscriberOptout"`
	// Campaign sub-usecases. Must be of defined valid sub-usecase types. Use
	// `/registry/enum/usecase` operation to retrieve list of valid sub-usecases
	SubUsecases []string `json:"subUsecases"`
	// Unique identifier assigned to the brand by the registry.
	TcrBrandID string `json:"tcrBrandId"`
	// Unique identifier assigned to the campaign by the registry.
	TcrCampaignID string `json:"tcrCampaignId"`
	// Is terms & conditions accepted?
	TermsAndConditions bool `json:"termsAndConditions"`
	// Link to the campaign's terms and conditions.
	TermsAndConditionsLink string `json:"termsAndConditionsLink"`
	// Campaign usecase. Must be of defined valid types. Use `/registry/enum/usecase`
	// operation to retrieve usecases available for given brand.
	Usecase string `json:"usecase"`
	// Business/industry segment of this campaign (Deprecated). Must be of defined
	// valid types. Use `/registry/enum/vertical` operation to retrieve verticals
	// available for given brand, vertical combination.
	//
	// This field is deprecated.
	//
	// Deprecated: deprecated
	Vertical string `json:"vertical"`
	// Failover webhook to which campaign status updates are sent.
	WebhookFailoverURL string `json:"webhookFailoverURL"`
	// Webhook to which campaign status updates are sent.
	WebhookURL string `json:"webhookURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeGated                      respjson.Field
		AssignedPhoneNumbersCount     respjson.Field
		AutoRenewal                   respjson.Field
		BilledDate                    respjson.Field
		BrandDisplayName              respjson.Field
		BrandID                       respjson.Field
		CampaignID                    respjson.Field
		CampaignStatus                respjson.Field
		CreateDate                    respjson.Field
		CspID                         respjson.Field
		Description                   respjson.Field
		DirectLending                 respjson.Field
		EmbeddedLink                  respjson.Field
		EmbeddedLinkSample            respjson.Field
		EmbeddedPhone                 respjson.Field
		FailureReasons                respjson.Field
		HelpKeywords                  respjson.Field
		HelpMessage                   respjson.Field
		IsTMobileNumberPoolingEnabled respjson.Field
		IsTMobileRegistered           respjson.Field
		IsTMobileSuspended            respjson.Field
		MessageFlow                   respjson.Field
		Mock                          respjson.Field
		NextRenewalOrExpirationDate   respjson.Field
		NumberPool                    respjson.Field
		OptinKeywords                 respjson.Field
		OptinMessage                  respjson.Field
		OptoutKeywords                respjson.Field
		OptoutMessage                 respjson.Field
		PrivacyPolicyLink             respjson.Field
		ReferenceID                   respjson.Field
		ResellerID                    respjson.Field
		Sample1                       respjson.Field
		Sample2                       respjson.Field
		Sample3                       respjson.Field
		Sample4                       respjson.Field
		Sample5                       respjson.Field
		Status                        respjson.Field
		SubmissionStatus              respjson.Field
		SubscriberHelp                respjson.Field
		SubscriberOptin               respjson.Field
		SubscriberOptout              respjson.Field
		SubUsecases                   respjson.Field
		TcrBrandID                    respjson.Field
		TcrCampaignID                 respjson.Field
		TermsAndConditions            respjson.Field
		TermsAndConditionsLink        respjson.Field
		Usecase                       respjson.Field
		Vertical                      respjson.Field
		WebhookFailoverURL            respjson.Field
		WebhookURL                    respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignListResponseRecord) RawJSON() string { return r.JSON.raw }
func (r *CampaignListResponseRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignAcceptSharingResponse = any

type CampaignDeactivateResponse struct {
	Time       float64 `json:"time,required"`
	Message    string  `json:"message"`
	RecordType string  `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Time        respjson.Field
		Message     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignDeactivateResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignDeactivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignGetMnoMetadataResponse struct {
	Number10999 CampaignGetMnoMetadataResponse10999 `json:"10999"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number10999 respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignGetMnoMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignGetMnoMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignGetMnoMetadataResponse10999 struct {
	MinMsgSamples       int64  `json:"minMsgSamples,required"`
	Mno                 string `json:"mno,required"`
	MnoReview           bool   `json:"mnoReview,required"`
	MnoSupport          bool   `json:"mnoSupport,required"`
	NoEmbeddedLink      bool   `json:"noEmbeddedLink,required"`
	NoEmbeddedPhone     bool   `json:"noEmbeddedPhone,required"`
	Qualify             bool   `json:"qualify,required"`
	ReqSubscriberHelp   bool   `json:"reqSubscriberHelp,required"`
	ReqSubscriberOptin  bool   `json:"reqSubscriberOptin,required"`
	ReqSubscriberOptout bool   `json:"reqSubscriberOptout,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MinMsgSamples       respjson.Field
		Mno                 respjson.Field
		MnoReview           respjson.Field
		MnoSupport          respjson.Field
		NoEmbeddedLink      respjson.Field
		NoEmbeddedPhone     respjson.Field
		Qualify             respjson.Field
		ReqSubscriberHelp   respjson.Field
		ReqSubscriberOptin  respjson.Field
		ReqSubscriberOptout respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignGetMnoMetadataResponse10999) RawJSON() string { return r.JSON.raw }
func (r *CampaignGetMnoMetadataResponse10999) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignGetOperationStatusResponse = any

type CampaignGetSharingStatusResponse struct {
	SharedByMe   CampaignSharingStatus `json:"sharedByMe"`
	SharedWithMe CampaignSharingStatus `json:"sharedWithMe"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SharedByMe   respjson.Field
		SharedWithMe respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignGetSharingStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignGetSharingStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignSubmitAppealResponse struct {
	// Timestamp when the appeal was submitted
	AppealedAt time.Time `json:"appealed_at" format:"date-time"`
	// Previous campaign status (currently always null)
	PreviousStatus string `json:"previous_status,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppealedAt     respjson.Field
		PreviousStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignSubmitAppealResponse) RawJSON() string { return r.JSON.raw }
func (r *CampaignSubmitAppealResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignUpdateParams struct {
	// Help message of the campaign.
	AutoRenewal param.Opt[bool] `json:"autoRenewal,omitzero"`
	// Help message of the campaign.
	HelpMessage param.Opt[string] `json:"helpMessage,omitzero"`
	// Message flow description.
	MessageFlow param.Opt[string] `json:"messageFlow,omitzero"`
	// Alphanumeric identifier of the reseller that you want to associate with this
	// campaign.
	ResellerID param.Opt[string] `json:"resellerId,omitzero"`
	// Message sample. Some campaign tiers require 1 or more message samples.
	Sample1 param.Opt[string] `json:"sample1,omitzero"`
	// Message sample. Some campaign tiers require 2 or more message samples.
	Sample2 param.Opt[string] `json:"sample2,omitzero"`
	// Message sample. Some campaign tiers require 3 or more message samples.
	Sample3 param.Opt[string] `json:"sample3,omitzero"`
	// Message sample. Some campaign tiers require 4 or more message samples.
	Sample4 param.Opt[string] `json:"sample4,omitzero"`
	// Message sample. Some campaign tiers require 5 or more message samples.
	Sample5 param.Opt[string] `json:"sample5,omitzero"`
	// Webhook failover to which campaign status updates are sent.
	WebhookFailoverURL param.Opt[string] `json:"webhookFailoverURL,omitzero"`
	// Webhook to which campaign status updates are sent.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	paramObj
}

func (r CampaignUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CampaignUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CampaignUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignListParams struct {
	BrandID string `query:"brandId,required" json:"-"`
	// The 1-indexed page number to get. The default value is `1`.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The amount of records per page, limited to between 1 and 500 inclusive. The
	// default value is `10`.
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// createdAt in descending order.
	//
	// Any of "assignedPhoneNumbersCount", "-assignedPhoneNumbersCount", "campaignId",
	// "-campaignId", "createdAt", "-createdAt", "status", "-status", "tcrCampaignId",
	// "-tcrCampaignId".
	Sort CampaignListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CampaignListParams]'s query parameters as `url.Values`.
func (r CampaignListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type CampaignListParamsSort string

const (
	CampaignListParamsSortAssignedPhoneNumbersCount     CampaignListParamsSort = "assignedPhoneNumbersCount"
	CampaignListParamsSortAssignedPhoneNumbersCountDesc CampaignListParamsSort = "-assignedPhoneNumbersCount"
	CampaignListParamsSortCampaignID                    CampaignListParamsSort = "campaignId"
	CampaignListParamsSortCampaignIDDesc                CampaignListParamsSort = "-campaignId"
	CampaignListParamsSortCreatedAt                     CampaignListParamsSort = "createdAt"
	CampaignListParamsSortCreatedAtDesc                 CampaignListParamsSort = "-createdAt"
	CampaignListParamsSortStatus                        CampaignListParamsSort = "status"
	CampaignListParamsSortStatusDesc                    CampaignListParamsSort = "-status"
	CampaignListParamsSortTcrCampaignID                 CampaignListParamsSort = "tcrCampaignId"
	CampaignListParamsSortTcrCampaignIDDesc             CampaignListParamsSort = "-tcrCampaignId"
)

type CampaignSubmitAppealParams struct {
	// Detailed explanation of why the campaign should be reconsidered and what changes
	// have been made to address the rejection reason.
	AppealReason string `json:"appeal_reason,required"`
	paramObj
}

func (r CampaignSubmitAppealParams) MarshalJSON() (data []byte, err error) {
	type shadow CampaignSubmitAppealParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CampaignSubmitAppealParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
