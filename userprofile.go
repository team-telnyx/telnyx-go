// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Read and update the authenticated user's profile and data locality.
//
// UserProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserProfileService] method instead.
type UserProfileService struct {
	Options []option.RequestOption
}

// NewUserProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserProfileService(opts ...option.RequestOption) (r UserProfileService) {
	r = UserProfileService{}
	r.Options = opts
	return
}

// Retrieve the authenticated user's profile, including data locality and migration
// status.
//
// **Beta:** this API should be considered a beta integration target because it may
// change if user account settings require changes
func (r *UserProfileService) List(ctx context.Context, opts ...option.RequestOption) (res *UserProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user_profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the authenticated user's profile using top-level JSON fields. Returns the
// updated profile.
//
// **Beta:** this API should be considered a beta integration target because it may
// change if user account settings require changes
func (r *UserProfileService) PatchAll(ctx context.Context, body UserProfilePatchAllParams, opts ...option.RequestOption) (res *UserProfilePatchAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user_profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// The authenticated user's profile.
type UserProfile struct {
	// Identifier of the user profile, distinct from the user ID.
	ID int64 `json:"id"`
	// First line of the profile address.
	Address1 string `json:"address_1" api:"nullable"`
	// Second line of the profile address.
	Address2 string `json:"address_2" api:"nullable"`
	// Business name associated with the profile.
	BusinessName string `json:"business_name" api:"nullable"`
	// City in the profile address.
	City string `json:"city" api:"nullable"`
	// Country code in the profile address.
	Country string `json:"country" api:"nullable"`
	// The user's data locality: USA (United States), DEU (Germany), AUS (Australia),
	// UAE (United Arab Emirates), or CAN (Canada).
	//
	// Any of "USA", "DEU", "AUS", "UAE", "CAN".
	DataLocality UserProfileDataLocality `json:"data_locality" api:"nullable"`
	// Status of the user's data locality migration, or null when no migration status
	// is recorded.
	//
	// Any of "not_started", "in_progress", "finished".
	DataLocalityMigrationStatus UserProfileDataLocalityMigrationStatus `json:"data_locality_migration_status" api:"nullable"`
	// Timestamp when the user's data locality was last changed, or null if it has not
	// been changed.
	DataLocalityUpdatedAt time.Time `json:"data_locality_updated_at" api:"nullable" format:"date-time"`
	// The user's first name.
	FirstName string `json:"first_name" api:"nullable"`
	// The user's last name.
	LastName string `json:"last_name" api:"nullable"`
	// Contact phone number associated with the profile.
	PhoneNumber string `json:"phone_number" api:"nullable"`
	// Identifies the type of resource.
	//
	// Any of "user_profile".
	RecordType UserProfileRecordType `json:"record_type"`
	// State or region in the profile address.
	State string `json:"state" api:"nullable"`
	// Timestamp of the most recent profile update.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Identifier of the user who owns this profile.
	UserID string `json:"user_id"`
	// Postal code in the profile address.
	Zip         string         `json:"zip" api:"nullable"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		Address1                    respjson.Field
		Address2                    respjson.Field
		BusinessName                respjson.Field
		City                        respjson.Field
		Country                     respjson.Field
		DataLocality                respjson.Field
		DataLocalityMigrationStatus respjson.Field
		DataLocalityUpdatedAt       respjson.Field
		FirstName                   respjson.Field
		LastName                    respjson.Field
		PhoneNumber                 respjson.Field
		RecordType                  respjson.Field
		State                       respjson.Field
		UpdatedAt                   respjson.Field
		UserID                      respjson.Field
		Zip                         respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfile) RawJSON() string { return r.JSON.raw }
func (r *UserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The user's data locality: USA (United States), DEU (Germany), AUS (Australia),
// UAE (United Arab Emirates), or CAN (Canada).
type UserProfileDataLocality string

const (
	UserProfileDataLocalityUsa UserProfileDataLocality = "USA"
	UserProfileDataLocalityDeu UserProfileDataLocality = "DEU"
	UserProfileDataLocalityAus UserProfileDataLocality = "AUS"
	UserProfileDataLocalityUae UserProfileDataLocality = "UAE"
	UserProfileDataLocalityCan UserProfileDataLocality = "CAN"
)

// Status of the user's data locality migration, or null when no migration status
// is recorded.
type UserProfileDataLocalityMigrationStatus string

const (
	UserProfileDataLocalityMigrationStatusNotStarted UserProfileDataLocalityMigrationStatus = "not_started"
	UserProfileDataLocalityMigrationStatusInProgress UserProfileDataLocalityMigrationStatus = "in_progress"
	UserProfileDataLocalityMigrationStatusFinished   UserProfileDataLocalityMigrationStatus = "finished"
)

// Identifies the type of resource.
type UserProfileRecordType string

const (
	UserProfileRecordTypeUserProfile UserProfileRecordType = "user_profile"
)

type UserProfileListResponse struct {
	// The authenticated user's profile.
	Data UserProfile `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserProfilePatchAllResponse struct {
	// The authenticated user's profile.
	Data UserProfile `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserProfilePatchAllResponse) RawJSON() string { return r.JSON.raw }
func (r *UserProfilePatchAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserProfilePatchAllParams struct {
	// First line of the profile address.
	Address1 param.Opt[string] `json:"address_1,omitzero"`
	// Second line of the profile address.
	Address2 param.Opt[string] `json:"address_2,omitzero"`
	// Business name associated with the profile.
	BusinessName param.Opt[string] `json:"business_name,omitzero"`
	// City in the profile address.
	City param.Opt[string] `json:"city,omitzero"`
	// Country code in the profile address.
	Country param.Opt[string] `json:"country,omitzero"`
	// The user's first name.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// The user's last name.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// Low-balance notification threshold in the billing currency. The resulting
	// profile must have both a threshold and notification frequency, or neither.
	LowBalanceThresholdAmount param.Opt[float64] `json:"low_balance_threshold_amount,omitzero"`
	// Contact phone number associated with the profile.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// State or region in the profile address.
	State param.Opt[string] `json:"state,omitzero"`
	// Postal code in the profile address.
	Zip param.Opt[string] `json:"zip,omitzero"`
	// Company size reported by the user.
	CompanySize param.Opt[string] `json:"company_size,omitzero"`
	// Default custom-pricing setting for managed accounts.
	DefaultManagedAccountAllowCustomPricing param.Opt[bool] `json:"default_managed_account_allow_custom_pricing,omitzero"`
	// Default emergency terms-of-service inheritance setting for managed accounts,
	// where supported.
	DefaultManagedAccountInheritEmergencyTermsOfService param.Opt[bool] `json:"default_managed_account_inherit_emergency_terms_of_service,omitzero"`
	// Default rollup-billing setting for managed accounts.
	DefaultManagedAccountRollupBilling param.Opt[bool] `json:"default_managed_account_rollup_billing,omitzero"`
	// Default setting for notifying the manager account.
	DefaultNotifyManagerAccount param.Opt[bool] `json:"default_notify_manager_account,omitzero"`
	// Default alphanumeric external PIN for phone numbers.
	DefaultNumbersExternalPin param.Opt[string] `json:"default_numbers_external_pin,omitzero"`
	// Whether the user received a free number at signup.
	GotFreeNumberAtSignup param.Opt[bool] `json:"got_free_number_at_signup,omitzero"`
	// Whether to notify the user when an invoice is available.
	NotifyOnInvoiceAvailable param.Opt[bool] `json:"notify_on_invoice_available,omitzero"`
	// Whether to notify the user when an invoice is paid.
	NotifyOnInvoicePayment param.Opt[bool] `json:"notify_on_invoice_payment,omitzero"`
	// Whether the user is an NPAC user.
	NpacUser param.Opt[bool] `json:"npac_user,omitzero"`
	// Whether port-out PIN validation is enabled.
	PortOutPinValidation param.Opt[bool] `json:"port_out_pin_validation,omitzero"`
	// Primary use case for the account.
	PrimaryUseCase param.Opt[string] `json:"primary_use_case,omitzero"`
	// The user's role.
	Role param.Opt[string] `json:"role,omitzero"`
	// Whether the user opts in to subscriptions.
	UserSubscriptionOptIn param.Opt[bool] `json:"user_subscription_opt_in,omitzero"`
	// Low-balance notification frequency. Organization sub-users cannot set a non-null
	// frequency. The resulting profile must have both a threshold and frequency, or
	// neither.
	//
	// Any of 2, 4, 12, 24.
	LowBalanceNotificationFrequency int64 `json:"low_balance_notification_frequency,omitzero"`
	// Data locality: USA (United States), DEU (Germany), AUS (Australia), UAE (United
	// Arab Emirates), or CAN (Canada). A change is permitted only if data locality has
	// not previously been changed. A successful change sets data_locality_updated_at
	// and initializes data_locality_migration_status to not_started.
	//
	// Any of "USA", "DEU", "AUS", "UAE", "CAN".
	DataLocality UserProfilePatchAllParamsDataLocality `json:"data_locality,omitzero"`
	// Build methods reported during onboarding.
	OnboardingBuildMethod []string `json:"onboarding_build_method,omitzero"`
	// Industries reported during onboarding.
	OnboardingIndustries []string `json:"onboarding_industries,omitzero"`
	// Monthly budget ranges reported during onboarding.
	OnboardingMonthlyBudget []string `json:"onboarding_monthly_budget,omitzero"`
	// Products of interest reported during onboarding.
	OnboardingProductInterest []string `json:"onboarding_product_interest,omitzero"`
	// Onboarding project types.
	OnboardingProjectTypes []string `json:"onboarding_project_types,omitzero"`
	// Reasons for considering Telnyx.
	OnboardingReasonsForConsideration []string `json:"onboarding_reasons_for_consideration,omitzero"`
	// Referral sources reported during onboarding.
	OnboardingReferrer []string `json:"onboarding_referrer,omitzero"`
	// Other onboarding use cases.
	OnboardingUseCasesOther []string `json:"onboarding_use_cases_other,omitzero"`
	// User interface theme: 0 for dark or 1 for light.
	//
	// Any of 0, 1.
	UiTheme int64 `json:"ui_theme,omitzero"`
	// Use cases reported by the user.
	UseCases []string `json:"use_cases,omitzero"`
	paramObj
}

func (r UserProfilePatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow UserProfilePatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserProfilePatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data locality: USA (United States), DEU (Germany), AUS (Australia), UAE (United
// Arab Emirates), or CAN (Canada). A change is permitted only if data locality has
// not previously been changed. A successful change sets data_locality_updated_at
// and initializes data_locality_migration_status to not_started.
type UserProfilePatchAllParamsDataLocality string

const (
	UserProfilePatchAllParamsDataLocalityUsa UserProfilePatchAllParamsDataLocality = "USA"
	UserProfilePatchAllParamsDataLocalityDeu UserProfilePatchAllParamsDataLocality = "DEU"
	UserProfilePatchAllParamsDataLocalityAus UserProfilePatchAllParamsDataLocality = "AUS"
	UserProfilePatchAllParamsDataLocalityUae UserProfilePatchAllParamsDataLocality = "UAE"
	UserProfilePatchAllParamsDataLocalityCan UserProfilePatchAllParamsDataLocality = "CAN"
)
