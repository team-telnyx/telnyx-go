// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Validate email addresses synchronously or in asynchronous batches.
//
// EmailValidationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailValidationService] method instead.
type EmailValidationService struct {
	Options []option.RequestOption
	// Validate email addresses synchronously or in asynchronous batches.
	Batch EmailValidationBatchService
}

// NewEmailValidationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailValidationService(opts ...option.RequestOption) (r EmailValidationService) {
	r = EmailValidationService{}
	r.Options = opts
	r.Batch = NewEmailValidationBatchService(opts...)
	return
}

// Validates a single email address and returns deliverability checks.
func (r *EmailValidationService) New(ctx context.Context, params EmailValidationNewParams, opts ...option.RequestOption) (res *EmailValidationNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email_validations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type EmailValidationCheck struct {
	Pass bool `json:"pass" api:"required"`
	// Human-readable check detail. Omitted when nil.
	Details string `json:"details"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pass        respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationCheck) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationNewResponse struct {
	Data EmailValidationNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationNewResponseData struct {
	Checks EmailValidationNewResponseDataChecks `json:"checks" api:"required"`
	Email  string                               `json:"email" api:"required"`
	// Any of "email_validation".
	RecordType string  `json:"record_type" api:"required"`
	RiskScore  float64 `json:"risk_score" api:"required"`
	Valid      bool    `json:"valid" api:"required"`
	// Suggested correction for typo. Omitted when nil.
	DidYouMean string `json:"did_you_mean"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		Email       respjson.Field
		RecordType  respjson.Field
		RiskScore   respjson.Field
		Valid       respjson.Field
		DidYouMean  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationNewResponseDataChecks struct {
	Disposable EmailValidationCheck                     `json:"disposable" api:"required"`
	Mx         EmailValidationCheck                     `json:"mx" api:"required"`
	RoleBased  EmailValidationCheck                     `json:"role_based" api:"required"`
	Syntax     EmailValidationCheck                     `json:"syntax" api:"required"`
	Typo       EmailValidationNewResponseDataChecksTypo `json:"typo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Disposable  respjson.Field
		Mx          respjson.Field
		RoleBased   respjson.Field
		Syntax      respjson.Field
		Typo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationNewResponseDataChecks) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationNewResponseDataChecks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationNewResponseDataChecksTypo struct {
	// Suggested correction for common typos. Omitted when nil.
	Suggestion string `json:"suggestion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Suggestion  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EmailValidationCheck
}

// Returns the unmodified JSON received from the API
func (r EmailValidationNewResponseDataChecksTypo) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationNewResponseDataChecksTypo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationNewParams struct {
	// Email address to validate. Any non-empty string is accepted; invalid syntax
	// returns valid=false rather than a request error.
	Email          string            `json:"email" api:"required"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r EmailValidationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailValidationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailValidationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
