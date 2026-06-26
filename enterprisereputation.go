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

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// EnterpriseReputationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationService] method instead.
type EnterpriseReputationService struct {
	Options []option.RequestOption
	// Phone-number reputation monitoring (spam-score lookup and tracking).
	Numbers EnterpriseReputationNumberService
	// Phone-number reputation monitoring (spam-score lookup and tracking).
	Loa EnterpriseReputationLoaService
	// Phone-number reputation monitoring (spam-score lookup and tracking).
	Remediation EnterpriseReputationRemediationService
}

// NewEnterpriseReputationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnterpriseReputationService(opts ...option.RequestOption) (r EnterpriseReputationService) {
	r = EnterpriseReputationService{}
	r.Options = opts
	r.Numbers = NewEnterpriseReputationNumberService(opts...)
	r.Loa = NewEnterpriseReputationLoaService(opts...)
	r.Remediation = NewEnterpriseReputationRemediationService(opts...)
	return
}

// Phone Number Reputation tracks how your outbound caller-IDs are perceived (spam
// risk, engagement, etc.) across the call-screening ecosystem. This endpoint reads
// the enterprise-level settings: whether the product is enabled, the refresh
// cadence, and the stored Letter of Authorization document id.
//
// Returns `404` if reputation has never been enabled for this enterprise.
func (r *EnterpriseReputationService) Get(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterpriseReputationPublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disable Phone Number Reputation. All registered numbers are de-registered as a
// cascade. The enterprise itself is unaffected. Returns `204` on success, `404` if
// reputation is not enabled for this enterprise.
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

// Activate Phone Number Reputation for the given enterprise. Requires an uploaded
// Letter of Authorization document (the `loa_document_id` references the Telnyx
// Documents API) and a refresh-frequency selection. After activation, individual
// phone numbers can be registered via `POST .../reputation/numbers`.
//
// **Prerequisite**: the calling user must have agreed to the Phone Number
// Reputation Terms of Service (`POST /terms_of_service/number_reputation/agree`).
//
// Failure modes:
//
// - `403` - Phone Number Reputation Terms of Service not accepted.
// - `404` - enterprise does not exist or does not belong to your account.
// - `400` - reputation already enabled for this enterprise.
// - `422` - `loa_document_id` missing or `check_frequency` invalid.
//
// **Pricing:** This is a billable action. See https://telnyx.com/pricing/numbers
// for current pricing.
func (r *EnterpriseReputationService) Enable(ctx context.Context, enterpriseID string, body EnterpriseReputationEnableParams, opts ...option.RequestOption) (res *EnterpriseReputationPublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update how often Telnyx refreshes the reputation data for this enterprise's
// registered numbers. The new frequency takes effect on the next scheduled
// refresh.
//
// The enterprise's reputation must be in `approved` status. A request made while
// the status is `pending` is rejected with `400 Bad Request`.
func (r *EnterpriseReputationService) UpdateFrequency(ctx context.Context, enterpriseID string, body EnterpriseReputationUpdateFrequencyParams, opts ...option.RequestOption) (res *EnterpriseReputationPublicWrapped, err error) {
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
	// How often Telnyx refreshes the stored reputation data for this enterprise's
	// registered numbers.
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency ReputationCheckFrequency `json:"check_frequency"`
	CreatedAt      time.Time                `json:"created_at" format:"date-time"`
	EnterpriseID   string                   `json:"enterprise_id" format:"uuid"`
	// Id of the signed LOA document.
	LoaDocumentID string `json:"loa_document_id" api:"nullable"`
	// Customer-facing Letter-of-Authorization verification state. `approved` is
	// required (alongside reputation status) before phone numbers can be added.
	//
	// Any of "pending", "approved", "rejected".
	LoaStatus EnterpriseReputationPublicLoaStatus `json:"loa_status"`
	// Populated when `status` is `rejected`.
	RejectionReasons []string `json:"rejection_reasons" api:"nullable"`
	// Lifecycle status of the enterprise's Phone Number Reputation activation.
	//
	// Any of "pending", "approved", "deleted", "rejected".
	Status    EnterpriseReputationPublicStatus `json:"status"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckFrequency   respjson.Field
		CreatedAt        respjson.Field
		EnterpriseID     respjson.Field
		LoaDocumentID    respjson.Field
		LoaStatus        respjson.Field
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

// Customer-facing Letter-of-Authorization verification state. `approved` is
// required (alongside reputation status) before phone numbers can be added.
type EnterpriseReputationPublicLoaStatus string

const (
	EnterpriseReputationPublicLoaStatusPending  EnterpriseReputationPublicLoaStatus = "pending"
	EnterpriseReputationPublicLoaStatusApproved EnterpriseReputationPublicLoaStatus = "approved"
	EnterpriseReputationPublicLoaStatusRejected EnterpriseReputationPublicLoaStatus = "rejected"
)

// Lifecycle status of the enterprise's Phone Number Reputation activation.
type EnterpriseReputationPublicStatus string

const (
	EnterpriseReputationPublicStatusPending  EnterpriseReputationPublicStatus = "pending"
	EnterpriseReputationPublicStatusApproved EnterpriseReputationPublicStatus = "approved"
	EnterpriseReputationPublicStatusDeleted  EnterpriseReputationPublicStatus = "deleted"
	EnterpriseReputationPublicStatusRejected EnterpriseReputationPublicStatus = "rejected"
)

type EnterpriseReputationPublicWrapped struct {
	Data EnterpriseReputationPublic `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationPublicWrapped) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationPublicWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often Telnyx refreshes the stored reputation data for this enterprise's
// registered numbers.
type ReputationCheckFrequency string

const (
	ReputationCheckFrequencyBusinessDaily ReputationCheckFrequency = "business_daily"
	ReputationCheckFrequencyDaily         ReputationCheckFrequency = "daily"
	ReputationCheckFrequencyWeekly        ReputationCheckFrequency = "weekly"
	ReputationCheckFrequencyBiweekly      ReputationCheckFrequency = "biweekly"
	ReputationCheckFrequencyMonthly       ReputationCheckFrequency = "monthly"
	ReputationCheckFrequencyNever         ReputationCheckFrequency = "never"
)

type EnterpriseReputationEnableParams struct {
	// Id of the signed Letter of Authorization document, returned by the Telnyx
	// Documents API after upload (upload via `POST /v2/documents`; see
	// https://developers.telnyx.com/api/documents).
	LoaDocumentID string `json:"loa_document_id" api:"required"`
	// How often Telnyx refreshes the stored reputation data for this enterprise's
	// registered numbers.
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency ReputationCheckFrequency `json:"check_frequency,omitzero"`
	paramObj
}

func (r EnterpriseReputationEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationUpdateFrequencyParams struct {
	// How often Telnyx refreshes the stored reputation data for this enterprise's
	// registered numbers.
	//
	// Any of "business_daily", "daily", "weekly", "biweekly", "monthly", "never".
	CheckFrequency ReputationCheckFrequency `json:"check_frequency,omitzero" api:"required"`
	paramObj
}

func (r EnterpriseReputationUpdateFrequencyParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationUpdateFrequencyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationUpdateFrequencyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
