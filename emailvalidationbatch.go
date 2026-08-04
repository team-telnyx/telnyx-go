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

// Validate email addresses synchronously or in asynchronous batches.
//
// EmailValidationBatchService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailValidationBatchService] method instead.
type EmailValidationBatchService struct {
	Options []option.RequestOption
}

// NewEmailValidationBatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailValidationBatchService(opts ...option.RequestOption) (r EmailValidationBatchService) {
	r = EmailValidationBatchService{}
	r.Options = opts
	return
}

// Creates an asynchronous batch validation job for up to 1,000 email addresses.
func (r *EmailValidationBatchService) New(ctx context.Context, params EmailValidationBatchNewParams, opts ...option.RequestOption) (res *EmailValidationBatchNewResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email_validations/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves the current status and, once completed, validation results for a batch
// job.
func (r *EmailValidationBatchService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailValidationBatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_validations/batch/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type EmailValidationBatchStatus string

const (
	EmailValidationBatchStatusPending    EmailValidationBatchStatus = "pending"
	EmailValidationBatchStatusProcessing EmailValidationBatchStatus = "processing"
	EmailValidationBatchStatusCompleted  EmailValidationBatchStatus = "completed"
	EmailValidationBatchStatusFailed     EmailValidationBatchStatus = "failed"
)

type EmailValidationBatchNewResponse struct {
	// Shape returned by the create endpoint. Includes duplicates_removed.
	Data EmailValidationBatchNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationBatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shape returned by the create endpoint. Includes duplicates_removed.
type EmailValidationBatchNewResponseData struct {
	ID                string `json:"id" api:"required" format:"uuid"`
	DuplicatesRemoved int64  `json:"duplicates_removed" api:"required"`
	// Any of "email_validation_batch".
	RecordType string `json:"record_type" api:"required"`
	// Any of "pending", "processing", "completed", "failed".
	Status     EmailValidationBatchStatus `json:"status" api:"required"`
	Total      int64                      `json:"total" api:"required"`
	WebhookURL string                     `json:"webhook_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DuplicatesRemoved respjson.Field
		RecordType        respjson.Field
		Status            respjson.Field
		Total             respjson.Field
		WebhookURL        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationBatchNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationBatchGetResponse struct {
	// Shape returned by the GET endpoint. Does not include duplicates_removed.
	Data EmailValidationBatchGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationBatchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Shape returned by the GET endpoint. Does not include duplicates_removed.
type EmailValidationBatchGetResponseData struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "email_validation_batch".
	RecordType string `json:"record_type" api:"required"`
	// Any of "pending", "processing", "completed", "failed".
	Status      EmailValidationBatchStatus `json:"status" api:"required"`
	Total       int64                      `json:"total" api:"required"`
	CompletedAt time.Time                  `json:"completed_at" format:"date-time"`
	// Map keyed by original email address. Present only when the batch is completed.
	Results    map[string]EmailValidationBatchGetResponseDataResult `json:"results"`
	WebhookURL string                                               `json:"webhook_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		Total       respjson.Field
		CompletedAt respjson.Field
		Results     respjson.Field
		WebhookURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationBatchGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationBatchGetResponseDataResult struct {
	Checks    EmailValidationBatchGetResponseDataResultsChecks `json:"checks" api:"required"`
	Email     string                                           `json:"email" api:"required"`
	RiskScore float64                                          `json:"risk_score" api:"required"`
	Valid     bool                                             `json:"valid" api:"required"`
	// Suggested correction for typo. Omitted when nil.
	DidYouMean string `json:"did_you_mean"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Checks      respjson.Field
		Email       respjson.Field
		RiskScore   respjson.Field
		Valid       respjson.Field
		DidYouMean  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailValidationBatchGetResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchGetResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationBatchGetResponseDataResultsChecks struct {
	Disposable EmailValidationCheck                                 `json:"disposable" api:"required"`
	Mx         EmailValidationCheck                                 `json:"mx" api:"required"`
	RoleBased  EmailValidationCheck                                 `json:"role_based" api:"required"`
	Syntax     EmailValidationCheck                                 `json:"syntax" api:"required"`
	Typo       EmailValidationBatchGetResponseDataResultsChecksTypo `json:"typo" api:"required"`
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
func (r EmailValidationBatchGetResponseDataResultsChecks) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchGetResponseDataResultsChecks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationBatchGetResponseDataResultsChecksTypo struct {
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
func (r EmailValidationBatchGetResponseDataResultsChecksTypo) RawJSON() string { return r.JSON.raw }
func (r *EmailValidationBatchGetResponseDataResultsChecksTypo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailValidationBatchNewParams struct {
	Emails []string `json:"emails,omitzero" api:"required"`
	// URL for batch completion webhook. Empty string is treated as omitted.
	// SSRF-protected; private/reserved IPs and internal hostnames are rejected.
	WebhookURL     param.Opt[string] `json:"webhook_url,omitzero" format:"uri"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r EmailValidationBatchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailValidationBatchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailValidationBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
