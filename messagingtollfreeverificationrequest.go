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
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// MessagingTollfreeVerificationRequestService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingTollfreeVerificationRequestService] method instead.
type MessagingTollfreeVerificationRequestService struct {
	Options []option.RequestOption
}

// NewMessagingTollfreeVerificationRequestService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewMessagingTollfreeVerificationRequestService(opts ...option.RequestOption) (r MessagingTollfreeVerificationRequestService) {
	r = MessagingTollfreeVerificationRequestService{}
	r.Options = opts
	return
}

// Submit a new tollfree verification request
func (r *MessagingTollfreeVerificationRequestService) New(ctx context.Context, body MessagingTollfreeVerificationRequestNewParams, opts ...option.RequestOption) (res *VerificationRequestEgress, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging_tollfree/verification/requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single verification request by its ID.
func (r *MessagingTollfreeVerificationRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VerificationRequestStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_tollfree/verification/requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing tollfree verification request. This is particularly useful
// when there are pending customer actions to be taken.
func (r *MessagingTollfreeVerificationRequestService) Update(ctx context.Context, id string, body MessagingTollfreeVerificationRequestUpdateParams, opts ...option.RequestOption) (res *VerificationRequestEgress, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_tollfree/verification/requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of previously-submitted tollfree verification requests
func (r *MessagingTollfreeVerificationRequestService) List(ctx context.Context, query MessagingTollfreeVerificationRequestListParams, opts ...option.RequestOption) (res *pagination.DefaultPaginationForMessagingTollfree[VerificationRequestStatus], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging_tollfree/verification/requests"
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

// Get a list of previously-submitted tollfree verification requests
func (r *MessagingTollfreeVerificationRequestService) ListAutoPaging(ctx context.Context, query MessagingTollfreeVerificationRequestListParams, opts ...option.RequestOption) *pagination.DefaultPaginationForMessagingTollfreeAutoPager[VerificationRequestStatus] {
	return pagination.NewDefaultPaginationForMessagingTollfreeAutoPager(r.List(ctx, query, opts...))
}

// Delete a verification request
//
// A request may only be deleted when when the request is in the "rejected" state.
//
// - `HTTP 200`: request successfully deleted
// - `HTTP 400`: request exists but can't be deleted (i.e. not rejected)
// - `HTTP 404`: request unknown or already deleted
func (r *MessagingTollfreeVerificationRequestService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_tollfree/verification/requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A phone number
type TfPhoneNumber struct {
	PhoneNumber string `json:"phoneNumber,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TfPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *TfPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TfPhoneNumber to a TfPhoneNumberParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TfPhoneNumberParam.Overrides()
func (r TfPhoneNumber) ToParam() TfPhoneNumberParam {
	return param.Override[TfPhoneNumberParam](json.RawMessage(r.RawJSON()))
}

// A phone number
//
// The property PhoneNumber is required.
type TfPhoneNumberParam struct {
	PhoneNumber string `json:"phoneNumber,required"`
	paramObj
}

func (r TfPhoneNumberParam) MarshalJSON() (data []byte, err error) {
	type shadow TfPhoneNumberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TfPhoneNumberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The body of a tollfree verification request
//
// The properties AdditionalInformation, BusinessAddr1, BusinessCity,
// BusinessContactEmail, BusinessContactFirstName, BusinessContactLastName,
// BusinessContactPhone, BusinessName, BusinessState, BusinessZip,
// CorporateWebsite, IsvReseller, MessageVolume, OptInWorkflow,
// OptInWorkflowImageURLs, PhoneNumbers, ProductionMessageContent, UseCase,
// UseCaseSummary are required.
type TfVerificationRequestParam struct {
	// Any additional information
	AdditionalInformation string `json:"additionalInformation,required"`
	// Line 1 of the business address
	BusinessAddr1 string `json:"businessAddr1,required"`
	// The city of the business address; the first letter should be capitalized
	BusinessCity string `json:"businessCity,required"`
	// The email address of the business contact
	BusinessContactEmail string `json:"businessContactEmail,required"`
	// First name of the business contact; there are no specific requirements on
	// formatting
	BusinessContactFirstName string `json:"businessContactFirstName,required"`
	// Last name of the business contact; there are no specific requirements on
	// formatting
	BusinessContactLastName string `json:"businessContactLastName,required"`
	// The phone number of the business contact in E.164 format
	BusinessContactPhone string `json:"businessContactPhone,required"`
	// Name of the business; there are no specific formatting requirements
	BusinessName string `json:"businessName,required"`
	// The full name of the state (not the 2 letter code) of the business address; the
	// first letter should be capitalized
	BusinessState string `json:"businessState,required"`
	// The ZIP code of the business address
	BusinessZip string `json:"businessZip,required"`
	// A URL, including the scheme, pointing to the corporate website
	CorporateWebsite string `json:"corporateWebsite,required"`
	// ISV name
	IsvReseller string `json:"isvReseller,required"`
	// Message Volume Enums
	//
	// Any of "10", "100", "1,000", "10,000", "100,000", "250,000", "500,000",
	// "750,000", "1,000,000", "5,000,000", "10,000,000+".
	MessageVolume Volume `json:"messageVolume,omitzero,required"`
	// Human-readable description of how end users will opt into receiving messages
	// from the given phone numbers
	OptInWorkflow string `json:"optInWorkflow,required"`
	// Images showing the opt-in workflow
	OptInWorkflowImageURLs []URLParam `json:"optInWorkflowImageURLs,omitzero,required"`
	// The phone numbers to request the verification of
	PhoneNumbers []TfPhoneNumberParam `json:"phoneNumbers,omitzero,required"`
	// An example of a message that will be sent from the given phone numbers
	ProductionMessageContent string `json:"productionMessageContent,required"`
	// Tollfree usecase categories
	//
	// Any of "2FA", "App Notifications", "Appointments", "Auctions", "Auto Repair
	// Services", "Bank Transfers", "Billing", "Booking Confirmations", "Business
	// Updates", "COVID-19 Alerts", "Career Training", "Chatbot", "Conversational /
	// Alerts", "Courier Services & Deliveries", "Emergency Alerts", "Events &
	// Planning", "Financial Services", "Fraud Alerts", "Fundraising", "General
	// Marketing", "General School Updates", "HR / Staffing", "Healthcare Alerts",
	// "Housing Community Updates", "Insurance Services", "Job Dispatch", "Legal
	// Services", "Mixed", "Motivational Reminders", "Notary Notifications", "Order
	// Notifications", "Political", "Public Works", "Real Estate Services", "Religious
	// Services", "Repair and Diagnostics Alerts", "Rewards Program", "Surveys",
	// "System Alerts", "Voting Reminders", "Waitlist Alerts", "Webinar Reminders",
	// "Workshop Alerts".
	UseCase UseCaseCategories `json:"useCase,omitzero,required"`
	// Human-readable summary of the desired use-case
	UseCaseSummary string `json:"useCaseSummary,required"`
	// ISO 3166-1 alpha-2 country code of the issuing business authority. Must be
	// exactly 2 letters. Automatically converted to uppercase. Required from
	// January 2026.
	BusinessRegistrationCountry param.Opt[string] `json:"businessRegistrationCountry,omitzero"`
	// Official business registration number (e.g., Employer Identification Number
	// (EIN) in the U.S.). Required from January 2026.
	BusinessRegistrationNumber param.Opt[string] `json:"businessRegistrationNumber,omitzero"`
	// Type of business registration being provided. Required from January 2026.
	BusinessRegistrationType param.Opt[string] `json:"businessRegistrationType,omitzero"`
	// Campaign Verify Authorization Token required for Political use case submissions
	// starting February 17, 2026. This token is validated by Zipwhip and must be
	// provided for all Political use case verifications after the deadline.
	CampaignVerifyAuthorizationToken param.Opt[string] `json:"campaignVerifyAuthorizationToken,omitzero"`
	// Doing Business As (DBA) name if different from legal name
	DoingBusinessAs param.Opt[string] `json:"doingBusinessAs,omitzero"`
	// The message returned when users text 'HELP'
	HelpMessageResponse param.Opt[string] `json:"helpMessageResponse,omitzero"`
	// Message sent to users confirming their opt-in to receive messages
	OptInConfirmationResponse param.Opt[string] `json:"optInConfirmationResponse,omitzero"`
	// Keywords used to collect and process consumer opt-ins
	OptInKeywords param.Opt[string] `json:"optInKeywords,omitzero"`
	// URL pointing to the business's privacy policy. Plain string, no URL format
	// validation.
	PrivacyPolicyURL param.Opt[string] `json:"privacyPolicyURL,omitzero"`
	// URL pointing to the business's terms and conditions. Plain string, no URL format
	// validation.
	TermsAndConditionURL param.Opt[string] `json:"termsAndConditionURL,omitzero"`
	// Indicates if messaging content requires age gating (e.g., 18+). Defaults to
	// false if not provided.
	AgeGatedContent param.Opt[bool] `json:"ageGatedContent,omitzero"`
	// Line 2 of the business address
	BusinessAddr2 param.Opt[string] `json:"businessAddr2,omitzero"`
	// URL that should receive webhooks relating to this verification request
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Business entity classification
	//
	// Any of "SOLE_PROPRIETOR", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT",
	// "GOVERNMENT".
	EntityType TollFreeVerificationEntityType `json:"entityType,omitzero"`
	paramObj
}

func (r TfVerificationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TfVerificationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TfVerificationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tollfree verification status
type TfVerificationStatus string

const (
	TfVerificationStatusVerified           TfVerificationStatus = "Verified"
	TfVerificationStatusRejected           TfVerificationStatus = "Rejected"
	TfVerificationStatusWaitingForVendor   TfVerificationStatus = "Waiting For Vendor"
	TfVerificationStatusWaitingForCustomer TfVerificationStatus = "Waiting For Customer"
	TfVerificationStatusWaitingForTelnyx   TfVerificationStatus = "Waiting For Telnyx"
	TfVerificationStatusInProgress         TfVerificationStatus = "In Progress"
)

// Business entity classification
type TollFreeVerificationEntityType string

const (
	TollFreeVerificationEntityTypeSoleProprietor TollFreeVerificationEntityType = "SOLE_PROPRIETOR"
	TollFreeVerificationEntityTypePrivateProfit  TollFreeVerificationEntityType = "PRIVATE_PROFIT"
	TollFreeVerificationEntityTypePublicProfit   TollFreeVerificationEntityType = "PUBLIC_PROFIT"
	TollFreeVerificationEntityTypeNonProfit      TollFreeVerificationEntityType = "NON_PROFIT"
	TollFreeVerificationEntityTypeGovernment     TollFreeVerificationEntityType = "GOVERNMENT"
)

type URL struct {
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r URL) RawJSON() string { return r.JSON.raw }
func (r *URL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this URL to a URLParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// URLParam.Overrides()
func (r URL) ToParam() URLParam {
	return param.Override[URLParam](json.RawMessage(r.RawJSON()))
}

// The property URL is required.
type URLParam struct {
	URL string `json:"url,required" format:"uri"`
	paramObj
}

func (r URLParam) MarshalJSON() (data []byte, err error) {
	type shadow URLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *URLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tollfree usecase categories
type UseCaseCategories string

const (
	UseCaseCategoriesTwoFa                      UseCaseCategories = "2FA"
	UseCaseCategoriesAppNotifications           UseCaseCategories = "App Notifications"
	UseCaseCategoriesAppointments               UseCaseCategories = "Appointments"
	UseCaseCategoriesAuctions                   UseCaseCategories = "Auctions"
	UseCaseCategoriesAutoRepairServices         UseCaseCategories = "Auto Repair Services"
	UseCaseCategoriesBankTransfers              UseCaseCategories = "Bank Transfers"
	UseCaseCategoriesBilling                    UseCaseCategories = "Billing"
	UseCaseCategoriesBookingConfirmations       UseCaseCategories = "Booking Confirmations"
	UseCaseCategoriesBusinessUpdates            UseCaseCategories = "Business Updates"
	UseCaseCategoriesCovid19Alerts              UseCaseCategories = "COVID-19 Alerts"
	UseCaseCategoriesCareerTraining             UseCaseCategories = "Career Training"
	UseCaseCategoriesChatbot                    UseCaseCategories = "Chatbot"
	UseCaseCategoriesConversationalAlerts       UseCaseCategories = "Conversational / Alerts"
	UseCaseCategoriesCourierServicesDeliveries  UseCaseCategories = "Courier Services & Deliveries"
	UseCaseCategoriesEmergencyAlerts            UseCaseCategories = "Emergency Alerts"
	UseCaseCategoriesEventsPlanning             UseCaseCategories = "Events & Planning"
	UseCaseCategoriesFinancialServices          UseCaseCategories = "Financial Services"
	UseCaseCategoriesFraudAlerts                UseCaseCategories = "Fraud Alerts"
	UseCaseCategoriesFundraising                UseCaseCategories = "Fundraising"
	UseCaseCategoriesGeneralMarketing           UseCaseCategories = "General Marketing"
	UseCaseCategoriesGeneralSchoolUpdates       UseCaseCategories = "General School Updates"
	UseCaseCategoriesHrStaffing                 UseCaseCategories = "HR / Staffing"
	UseCaseCategoriesHealthcareAlerts           UseCaseCategories = "Healthcare Alerts"
	UseCaseCategoriesHousingCommunityUpdates    UseCaseCategories = "Housing Community Updates"
	UseCaseCategoriesInsuranceServices          UseCaseCategories = "Insurance Services"
	UseCaseCategoriesJobDispatch                UseCaseCategories = "Job Dispatch"
	UseCaseCategoriesLegalServices              UseCaseCategories = "Legal Services"
	UseCaseCategoriesMixed                      UseCaseCategories = "Mixed"
	UseCaseCategoriesMotivationalReminders      UseCaseCategories = "Motivational Reminders"
	UseCaseCategoriesNotaryNotifications        UseCaseCategories = "Notary Notifications"
	UseCaseCategoriesOrderNotifications         UseCaseCategories = "Order Notifications"
	UseCaseCategoriesPolitical                  UseCaseCategories = "Political"
	UseCaseCategoriesPublicWorks                UseCaseCategories = "Public Works"
	UseCaseCategoriesRealEstateServices         UseCaseCategories = "Real Estate Services"
	UseCaseCategoriesReligiousServices          UseCaseCategories = "Religious Services"
	UseCaseCategoriesRepairAndDiagnosticsAlerts UseCaseCategories = "Repair and Diagnostics Alerts"
	UseCaseCategoriesRewardsProgram             UseCaseCategories = "Rewards Program"
	UseCaseCategoriesSurveys                    UseCaseCategories = "Surveys"
	UseCaseCategoriesSystemAlerts               UseCaseCategories = "System Alerts"
	UseCaseCategoriesVotingReminders            UseCaseCategories = "Voting Reminders"
	UseCaseCategoriesWaitlistAlerts             UseCaseCategories = "Waitlist Alerts"
	UseCaseCategoriesWebinarReminders           UseCaseCategories = "Webinar Reminders"
	UseCaseCategoriesWorkshopAlerts             UseCaseCategories = "Workshop Alerts"
)

// A verification request as it comes out of the database
type VerificationRequestEgress struct {
	ID                       string `json:"id,required" format:"uuid"`
	AdditionalInformation    string `json:"additionalInformation,required"`
	BusinessAddr1            string `json:"businessAddr1,required"`
	BusinessCity             string `json:"businessCity,required"`
	BusinessContactEmail     string `json:"businessContactEmail,required"`
	BusinessContactFirstName string `json:"businessContactFirstName,required"`
	BusinessContactLastName  string `json:"businessContactLastName,required"`
	BusinessContactPhone     string `json:"businessContactPhone,required"`
	BusinessName             string `json:"businessName,required"`
	BusinessState            string `json:"businessState,required"`
	BusinessZip              string `json:"businessZip,required"`
	CorporateWebsite         string `json:"corporateWebsite,required"`
	IsvReseller              string `json:"isvReseller,required"`
	// Message Volume Enums
	//
	// Any of "10", "100", "1,000", "10,000", "100,000", "250,000", "500,000",
	// "750,000", "1,000,000", "5,000,000", "10,000,000+".
	MessageVolume            Volume          `json:"messageVolume,required"`
	OptInWorkflow            string          `json:"optInWorkflow,required"`
	OptInWorkflowImageURLs   []URL           `json:"optInWorkflowImageURLs,required"`
	PhoneNumbers             []TfPhoneNumber `json:"phoneNumbers,required"`
	ProductionMessageContent string          `json:"productionMessageContent,required"`
	// Tollfree usecase categories
	//
	// Any of "2FA", "App Notifications", "Appointments", "Auctions", "Auto Repair
	// Services", "Bank Transfers", "Billing", "Booking Confirmations", "Business
	// Updates", "COVID-19 Alerts", "Career Training", "Chatbot", "Conversational /
	// Alerts", "Courier Services & Deliveries", "Emergency Alerts", "Events &
	// Planning", "Financial Services", "Fraud Alerts", "Fundraising", "General
	// Marketing", "General School Updates", "HR / Staffing", "Healthcare Alerts",
	// "Housing Community Updates", "Insurance Services", "Job Dispatch", "Legal
	// Services", "Mixed", "Motivational Reminders", "Notary Notifications", "Order
	// Notifications", "Political", "Public Works", "Real Estate Services", "Religious
	// Services", "Repair and Diagnostics Alerts", "Rewards Program", "Surveys",
	// "System Alerts", "Voting Reminders", "Waitlist Alerts", "Webinar Reminders",
	// "Workshop Alerts".
	UseCase                     UseCaseCategories `json:"useCase,required"`
	UseCaseSummary              string            `json:"useCaseSummary,required"`
	VerificationRequestID       string            `json:"verificationRequestId,required"`
	AgeGatedContent             bool              `json:"ageGatedContent"`
	BusinessAddr2               string            `json:"businessAddr2"`
	BusinessRegistrationCountry string            `json:"businessRegistrationCountry"`
	BusinessRegistrationNumber  string            `json:"businessRegistrationNumber"`
	BusinessRegistrationType    string            `json:"businessRegistrationType"`
	// Campaign Verify Authorization Token required for Political use case submissions
	// starting February 17, 2026
	CampaignVerifyAuthorizationToken string `json:"campaignVerifyAuthorizationToken,nullable"`
	DoingBusinessAs                  string `json:"doingBusinessAs"`
	// Business entity classification
	//
	// Any of "SOLE_PROPRIETOR", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT",
	// "GOVERNMENT".
	EntityType                TollFreeVerificationEntityType `json:"entityType"`
	HelpMessageResponse       string                         `json:"helpMessageResponse"`
	OptInConfirmationResponse string                         `json:"optInConfirmationResponse"`
	OptInKeywords             string                         `json:"optInKeywords"`
	PrivacyPolicyURL          string                         `json:"privacyPolicyURL"`
	TermsAndConditionURL      string                         `json:"termsAndConditionURL"`
	// Tollfree verification status
	//
	// Any of "Verified", "Rejected", "Waiting For Vendor", "Waiting For Customer",
	// "Waiting For Telnyx", "In Progress".
	VerificationStatus TfVerificationStatus `json:"verificationStatus"`
	WebhookURL         string               `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		AdditionalInformation            respjson.Field
		BusinessAddr1                    respjson.Field
		BusinessCity                     respjson.Field
		BusinessContactEmail             respjson.Field
		BusinessContactFirstName         respjson.Field
		BusinessContactLastName          respjson.Field
		BusinessContactPhone             respjson.Field
		BusinessName                     respjson.Field
		BusinessState                    respjson.Field
		BusinessZip                      respjson.Field
		CorporateWebsite                 respjson.Field
		IsvReseller                      respjson.Field
		MessageVolume                    respjson.Field
		OptInWorkflow                    respjson.Field
		OptInWorkflowImageURLs           respjson.Field
		PhoneNumbers                     respjson.Field
		ProductionMessageContent         respjson.Field
		UseCase                          respjson.Field
		UseCaseSummary                   respjson.Field
		VerificationRequestID            respjson.Field
		AgeGatedContent                  respjson.Field
		BusinessAddr2                    respjson.Field
		BusinessRegistrationCountry      respjson.Field
		BusinessRegistrationNumber       respjson.Field
		BusinessRegistrationType         respjson.Field
		CampaignVerifyAuthorizationToken respjson.Field
		DoingBusinessAs                  respjson.Field
		EntityType                       respjson.Field
		HelpMessageResponse              respjson.Field
		OptInConfirmationResponse        respjson.Field
		OptInKeywords                    respjson.Field
		PrivacyPolicyURL                 respjson.Field
		TermsAndConditionURL             respjson.Field
		VerificationStatus               respjson.Field
		WebhookURL                       respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationRequestEgress) RawJSON() string { return r.JSON.raw }
func (r *VerificationRequestEgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A verification request and its status, suitable for returning to users
type VerificationRequestStatus struct {
	ID                       string `json:"id,required" format:"uuid"`
	AdditionalInformation    string `json:"additionalInformation,required"`
	BusinessAddr1            string `json:"businessAddr1,required"`
	BusinessCity             string `json:"businessCity,required"`
	BusinessContactEmail     string `json:"businessContactEmail,required"`
	BusinessContactFirstName string `json:"businessContactFirstName,required"`
	BusinessContactLastName  string `json:"businessContactLastName,required"`
	BusinessContactPhone     string `json:"businessContactPhone,required"`
	BusinessName             string `json:"businessName,required"`
	BusinessState            string `json:"businessState,required"`
	BusinessZip              string `json:"businessZip,required"`
	CorporateWebsite         string `json:"corporateWebsite,required"`
	IsvReseller              string `json:"isvReseller,required"`
	// Message Volume Enums
	//
	// Any of "10", "100", "1,000", "10,000", "100,000", "250,000", "500,000",
	// "750,000", "1,000,000", "5,000,000", "10,000,000+".
	MessageVolume            Volume          `json:"messageVolume,required"`
	OptInWorkflow            string          `json:"optInWorkflow,required"`
	OptInWorkflowImageURLs   []URL           `json:"optInWorkflowImageURLs,required"`
	PhoneNumbers             []TfPhoneNumber `json:"phoneNumbers,required"`
	ProductionMessageContent string          `json:"productionMessageContent,required"`
	// Tollfree usecase categories
	//
	// Any of "2FA", "App Notifications", "Appointments", "Auctions", "Auto Repair
	// Services", "Bank Transfers", "Billing", "Booking Confirmations", "Business
	// Updates", "COVID-19 Alerts", "Career Training", "Chatbot", "Conversational /
	// Alerts", "Courier Services & Deliveries", "Emergency Alerts", "Events &
	// Planning", "Financial Services", "Fraud Alerts", "Fundraising", "General
	// Marketing", "General School Updates", "HR / Staffing", "Healthcare Alerts",
	// "Housing Community Updates", "Insurance Services", "Job Dispatch", "Legal
	// Services", "Mixed", "Motivational Reminders", "Notary Notifications", "Order
	// Notifications", "Political", "Public Works", "Real Estate Services", "Religious
	// Services", "Repair and Diagnostics Alerts", "Rewards Program", "Surveys",
	// "System Alerts", "Voting Reminders", "Waitlist Alerts", "Webinar Reminders",
	// "Workshop Alerts".
	UseCase        UseCaseCategories `json:"useCase,required"`
	UseCaseSummary string            `json:"useCaseSummary,required"`
	// Tollfree verification status
	//
	// Any of "Verified", "Rejected", "Waiting For Vendor", "Waiting For Customer",
	// "Waiting For Telnyx", "In Progress".
	VerificationStatus          TfVerificationStatus `json:"verificationStatus,required"`
	AgeGatedContent             bool                 `json:"ageGatedContent"`
	BusinessAddr2               string               `json:"businessAddr2"`
	BusinessRegistrationCountry string               `json:"businessRegistrationCountry"`
	BusinessRegistrationNumber  string               `json:"businessRegistrationNumber"`
	BusinessRegistrationType    string               `json:"businessRegistrationType"`
	// Campaign Verify Authorization Token required for Political use case submissions
	// starting February 17, 2026
	CampaignVerifyAuthorizationToken string    `json:"campaignVerifyAuthorizationToken,nullable"`
	CreatedAt                        time.Time `json:"createdAt" format:"date-time"`
	DoingBusinessAs                  string    `json:"doingBusinessAs"`
	// Business entity classification
	//
	// Any of "SOLE_PROPRIETOR", "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT",
	// "GOVERNMENT".
	EntityType                TollFreeVerificationEntityType `json:"entityType"`
	HelpMessageResponse       string                         `json:"helpMessageResponse"`
	OptInConfirmationResponse string                         `json:"optInConfirmationResponse"`
	OptInKeywords             string                         `json:"optInKeywords"`
	PrivacyPolicyURL          string                         `json:"privacyPolicyURL"`
	Reason                    string                         `json:"reason"`
	TermsAndConditionURL      string                         `json:"termsAndConditionURL"`
	UpdatedAt                 time.Time                      `json:"updatedAt" format:"date-time"`
	WebhookURL                string                         `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		AdditionalInformation            respjson.Field
		BusinessAddr1                    respjson.Field
		BusinessCity                     respjson.Field
		BusinessContactEmail             respjson.Field
		BusinessContactFirstName         respjson.Field
		BusinessContactLastName          respjson.Field
		BusinessContactPhone             respjson.Field
		BusinessName                     respjson.Field
		BusinessState                    respjson.Field
		BusinessZip                      respjson.Field
		CorporateWebsite                 respjson.Field
		IsvReseller                      respjson.Field
		MessageVolume                    respjson.Field
		OptInWorkflow                    respjson.Field
		OptInWorkflowImageURLs           respjson.Field
		PhoneNumbers                     respjson.Field
		ProductionMessageContent         respjson.Field
		UseCase                          respjson.Field
		UseCaseSummary                   respjson.Field
		VerificationStatus               respjson.Field
		AgeGatedContent                  respjson.Field
		BusinessAddr2                    respjson.Field
		BusinessRegistrationCountry      respjson.Field
		BusinessRegistrationNumber       respjson.Field
		BusinessRegistrationType         respjson.Field
		CampaignVerifyAuthorizationToken respjson.Field
		CreatedAt                        respjson.Field
		DoingBusinessAs                  respjson.Field
		EntityType                       respjson.Field
		HelpMessageResponse              respjson.Field
		OptInConfirmationResponse        respjson.Field
		OptInKeywords                    respjson.Field
		PrivacyPolicyURL                 respjson.Field
		Reason                           respjson.Field
		TermsAndConditionURL             respjson.Field
		UpdatedAt                        respjson.Field
		WebhookURL                       respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationRequestStatus) RawJSON() string { return r.JSON.raw }
func (r *VerificationRequestStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message Volume Enums
type Volume string

const (
	VolumeV10           Volume = "10"
	VolumeV100          Volume = "100"
	VolumeV1000         Volume = "1,000"
	VolumeV10000        Volume = "10,000"
	VolumeV100000       Volume = "100,000"
	VolumeV250000       Volume = "250,000"
	VolumeV500000       Volume = "500,000"
	VolumeV750000       Volume = "750,000"
	VolumeV1000000      Volume = "1,000,000"
	VolumeV5000000      Volume = "5,000,000"
	VolumeV10000000Plus Volume = "10,000,000+"
)

type MessagingTollfreeVerificationRequestNewParams struct {
	// The body of a tollfree verification request
	TfVerificationRequest TfVerificationRequestParam
	paramObj
}

func (r MessagingTollfreeVerificationRequestNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TfVerificationRequest)
}
func (r *MessagingTollfreeVerificationRequestNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TfVerificationRequest)
}

type MessagingTollfreeVerificationRequestUpdateParams struct {
	// The body of a tollfree verification request
	TfVerificationRequest TfVerificationRequestParam
	paramObj
}

func (r MessagingTollfreeVerificationRequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TfVerificationRequest)
}
func (r *MessagingTollfreeVerificationRequestUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TfVerificationRequest)
}

type MessagingTollfreeVerificationRequestListParams struct {
	Page int64 `query:"page,required" json:"-"`
	// Request this many records per page
	//
	//	This value is automatically clamped if the provided value is too large.
	PageSize    int64                `query:"page_size,required" json:"-"`
	DateEnd     param.Opt[time.Time] `query:"date_end,omitzero" format:"date-time" json:"-"`
	DateStart   param.Opt[time.Time] `query:"date_start,omitzero" format:"date-time" json:"-"`
	PhoneNumber param.Opt[string]    `query:"phone_number,omitzero" json:"-"`
	// Tollfree verification status
	//
	// Any of "Verified", "Rejected", "Waiting For Vendor", "Waiting For Customer",
	// "Waiting For Telnyx", "In Progress".
	Status TfVerificationStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingTollfreeVerificationRequestListParams]'s query
// parameters as `url.Values`.
func (r MessagingTollfreeVerificationRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
