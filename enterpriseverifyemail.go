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

// Verify ownership of a DIR's authorizer email. A short code is emailed and
// confirmed; the email must be verified before references can be submitted.
//
// EnterpriseVerifyEmailService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseVerifyEmailService] method instead.
type EnterpriseVerifyEmailService struct {
	Options []option.RequestOption
}

// NewEnterpriseVerifyEmailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnterpriseVerifyEmailService(opts ...option.RequestOption) (r EnterpriseVerifyEmailService) {
	r = EnterpriseVerifyEmailService{}
	r.Options = opts
	return
}

// Email a 6-digit code to the enterprise account's contact email to confirm
// ownership of that address.
//
// A BPO (Business Process Outsourcer) account has no DIR, so it proves ownership
// of its own contact email here rather than through a DIR. A BPO account cannot be
// approved for use until this contact email is verified.
//
// The code expires in 15 minutes. Requesting a new code invalidates any previous
// one. Resends are rate limited (a short cooldown plus a daily cap). Submit the
// code to `POST /enterprises/{enterprise_id}/verify_email/confirm`.
func (r *EnterpriseVerifyEmailService) New(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterpriseEmailVerificationStatusWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/verify_email", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Submit the 6-digit code that was emailed to the enterprise account's contact
// email. On success the contact email is marked verified.
//
// For security, any failure (wrong, expired, already-used, or too many attempts)
// returns the same generic message.
func (r *EnterpriseVerifyEmailService) Confirm(ctx context.Context, enterpriseID string, body EnterpriseVerifyEmailConfirmParams, opts ...option.RequestOption) (res *EnterpriseEmailVerificationStatusWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/verify_email/confirm", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EnterpriseEmailVerificationStatusWrapped struct {
	// Verification state for an enterprise account's contact email.
	Data EnterpriseEmailVerificationStatusWrappedData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseEmailVerificationStatusWrapped) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseEmailVerificationStatusWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification state for an enterprise account's contact email.
type EnterpriseEmailVerificationStatusWrappedData struct {
	// Whether the enterprise account's contact email has been confirmed.
	EmailVerified bool `json:"email_verified" api:"required"`
	// Always `email_verification`.
	//
	// Any of "email_verification".
	RecordType string `json:"record_type" api:"required"`
	// `sent` after a code is emailed; `verified` after a successful confirm.
	//
	// Any of "sent", "verified".
	Status string `json:"status" api:"required"`
	// When the code just sent stops being accepted. Present on a send response; null
	// on a confirm response.
	ExpiresAt time.Time `json:"expires_at" api:"nullable" format:"date-time"`
	// How many more codes may be requested for this enterprise account today. Present
	// on a send response; null on a confirm response.
	SendsRemainingToday int64 `json:"sends_remaining_today" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailVerified       respjson.Field
		RecordType          respjson.Field
		Status              respjson.Field
		ExpiresAt           respjson.Field
		SendsRemainingToday respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseEmailVerificationStatusWrappedData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseEmailVerificationStatusWrappedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseVerifyEmailConfirmParams struct {
	// The 6-digit code sent to the enterprise account's contact email.
	Code string `json:"code" api:"required"`
	paramObj
}

func (r EnterpriseVerifyEmailConfirmParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseVerifyEmailConfirmParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseVerifyEmailConfirmParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
