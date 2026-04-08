// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage Number Reputation enrollment and check frequency settings for an
// enterprise
//
// EnterpriseReputationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationService] method instead.
type EnterpriseReputationService struct {
	Options []option.RequestOption
	// Associate phone numbers with an enterprise for reputation monitoring and
	// retrieve reputation scores
	Numbers EnterpriseReputationNumberService
}

// NewEnterpriseReputationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnterpriseReputationService(opts ...option.RequestOption) (r EnterpriseReputationService) {
	r = EnterpriseReputationService{}
	r.Options = opts
	r.Numbers = NewEnterpriseReputationNumberService(opts...)
	return
}

// Retrieve the current Number Reputation settings for an enterprise.
//
// Returns the enrollment status (`pending`, `approved`, `rejected`, `deleted`),
// check frequency, and any rejection reasons.
//
// Returns `404` if reputation has not been enabled for this enterprise.
func (r *EnterpriseReputationService) Get(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterpriseReputationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disable Number Reputation for an enterprise.
//
// This will:
//
// - Delete the reputation settings record
// - Log the deletion for audit purposes
// - Stop all future automated reputation checks
//
// **Note:** Can only be performed on `approved` reputation settings.
func (r *EnterpriseReputationService) Disable(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return err
	}
	path := fmt.Sprintf("enterprises/%s/reputation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Enable Number Reputation service for an enterprise.
//
// **Requirements:**
//
// - Signed LOA (Letter of Authorization) document ID
// - Reputation check frequency (defaults to `business_daily`)
// - Number Reputation Terms of Service must be accepted
//
// **Flow:**
//
// 1. Registers the enterprise for reputation monitoring
// 2. Creates reputation settings with `pending` status
// 3. Awaits admin approval before monitoring begins
//
// **Resubmission After Rejection:** If a previously rejected record exists, this
// endpoint will delete it and create a new `pending` record.
//
// **Available Frequencies:**
//
// - `business_daily` — Monday–Friday
// - `daily` — Every day
// - `weekly` — Once per week
// - `biweekly` — Once every two weeks
// - `monthly` — Once per month
// - `never` — Manual refresh only
func (r *EnterpriseReputationService) Enable(ctx context.Context, enterpriseID string, body EnterpriseReputationEnableParams, opts ...option.RequestOption) (res *EnterpriseReputationEnableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update how often reputation data is automatically refreshed.
//
// **Note:** The enterprise must have `approved` reputation settings. Updating
// frequency on `pending` or `rejected` settings will return an error.
//
// **Available Frequencies:**
//
// - `business_daily` — Monday–Friday
// - `daily` — Every day including weekends
// - `weekly` — Once per week
// - `biweekly` — Once every two weeks
// - `monthly` — Once per month
// - `never` — Manual refresh only (no automatic checks)
func (r *EnterpriseReputationService) UpdateFrequency(ctx context.Context, enterpriseID string, body EnterpriseReputationUpdateFrequencyParams, opts ...option.RequestOption) (res *EnterpriseReputationUpdateFrequencyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/frequency", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type EnterpriseReputationPublic struct {
	// Frequency for refreshing reputation data
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency EnterpriseReputationPublicCheckFrequency `json:"check_frequency"`
	// When the reputation settings were created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the associated enterprise
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// ID of the signed LOA document
	LoaDocumentID string `json:"loa_document_id" api:"nullable"`
	// Reasons for rejection (present when status is rejected)
	RejectionReasons []string `json:"rejection_reasons" api:"nullable"`
	// Current enrollment status
	//
	// Any of "pending", "approved", "rejected", "deleted".
	Status EnterpriseReputationPublicStatus `json:"status"`
	// When the reputation settings were last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckFrequency   respjson.Field
		CreatedAt        respjson.Field
		EnterpriseID     respjson.Field
		LoaDocumentID    respjson.Field
		RejectionReasons respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationPublic) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationPublic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Frequency for refreshing reputation data
type EnterpriseReputationPublicCheckFrequency string

const (
	EnterpriseReputationPublicCheckFrequencyBusinessDaily EnterpriseReputationPublicCheckFrequency = "business_daily"
	EnterpriseReputationPublicCheckFrequencyDaily         EnterpriseReputationPublicCheckFrequency = "daily"
	EnterpriseReputationPublicCheckFrequencyWeekly        EnterpriseReputationPublicCheckFrequency = "weekly"
	EnterpriseReputationPublicCheckFrequencyBiweekly      EnterpriseReputationPublicCheckFrequency = "biweekly"
	EnterpriseReputationPublicCheckFrequencyMonthly       EnterpriseReputationPublicCheckFrequency = "monthly"
	EnterpriseReputationPublicCheckFrequencyNever         EnterpriseReputationPublicCheckFrequency = "never"
)

// Current enrollment status
type EnterpriseReputationPublicStatus string

const (
	EnterpriseReputationPublicStatusPending  EnterpriseReputationPublicStatus = "pending"
	EnterpriseReputationPublicStatusApproved EnterpriseReputationPublicStatus = "approved"
	EnterpriseReputationPublicStatusRejected EnterpriseReputationPublicStatus = "rejected"
	EnterpriseReputationPublicStatusDeleted  EnterpriseReputationPublicStatus = "deleted"
)

type EnterpriseReputationGetResponse struct {
	Data EnterpriseReputationPublic `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationEnableResponse struct {
	Data EnterpriseReputationPublic `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationEnableResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationEnableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationUpdateFrequencyResponse struct {
	Data EnterpriseReputationPublic `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationUpdateFrequencyResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationUpdateFrequencyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationEnableParams struct {
	// ID of the signed Letter of Authorization (LOA) document uploaded to the document
	// service
	LoaDocumentID string `json:"loa_document_id" api:"required"`
	// Frequency for automatically refreshing reputation data
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency EnterpriseReputationEnableParamsCheckFrequency `json:"check_frequency,omitzero"`
	paramObj
}

func (r EnterpriseReputationEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Frequency for automatically refreshing reputation data
type EnterpriseReputationEnableParamsCheckFrequency string

const (
	EnterpriseReputationEnableParamsCheckFrequencyBusinessDaily EnterpriseReputationEnableParamsCheckFrequency = "business_daily"
	EnterpriseReputationEnableParamsCheckFrequencyDaily         EnterpriseReputationEnableParamsCheckFrequency = "daily"
	EnterpriseReputationEnableParamsCheckFrequencyWeekly        EnterpriseReputationEnableParamsCheckFrequency = "weekly"
	EnterpriseReputationEnableParamsCheckFrequencyBiweekly      EnterpriseReputationEnableParamsCheckFrequency = "biweekly"
	EnterpriseReputationEnableParamsCheckFrequencyMonthly       EnterpriseReputationEnableParamsCheckFrequency = "monthly"
	EnterpriseReputationEnableParamsCheckFrequencyNever         EnterpriseReputationEnableParamsCheckFrequency = "never"
)

type EnterpriseReputationUpdateFrequencyParams struct {
	// New frequency for refreshing reputation data
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency EnterpriseReputationUpdateFrequencyParamsCheckFrequency `json:"check_frequency,omitzero" api:"required"`
	paramObj
}

func (r EnterpriseReputationUpdateFrequencyParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationUpdateFrequencyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationUpdateFrequencyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New frequency for refreshing reputation data
type EnterpriseReputationUpdateFrequencyParamsCheckFrequency string

const (
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyBusinessDaily EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "business_daily"
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyDaily         EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "daily"
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyWeekly        EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "weekly"
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyBiweekly      EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "biweekly"
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyMonthly       EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "monthly"
	EnterpriseReputationUpdateFrequencyParamsCheckFrequencyNever         EnterpriseReputationUpdateFrequencyParamsCheckFrequency = "never"
)
