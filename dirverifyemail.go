// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Verify ownership of a DIR's authorizer email. A short code is emailed and
// confirmed; the email must be verified before references can be submitted.
//
// DirVerifyEmailService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirVerifyEmailService] method instead.
type DirVerifyEmailService struct {
	Options []option.RequestOption
}

// NewDirVerifyEmailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirVerifyEmailService(opts ...option.RequestOption) (r DirVerifyEmailService) {
	r = DirVerifyEmailService{}
	r.Options = opts
	return
}

// Email a 6-digit code to the DIR's authorizer email to confirm ownership of that
// address.
//
// The code expires in 15 minutes. Requesting a new code invalidates any previous
// one. Resends are rate limited (a short cooldown plus a daily cap). Submit the
// code to `POST /dir/{dir_id}/verify_email/confirm`.
func (r *DirVerifyEmailService) New(ctx context.Context, dirID string, opts ...option.RequestOption) (res *EmailVerificationStatusWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/verify_email", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Whether the DIR's current authorizer email has been verified.
func (r *DirVerifyEmailService) List(ctx context.Context, dirID string, opts ...option.RequestOption) (res *EmailVerificationStatusWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/verify_email", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit the 6-digit code that was emailed to the DIR's authorizer email. On
// success the authorizer email is marked verified.
//
// For security, any failure (wrong, expired, already-used, or too many attempts)
// returns the same generic message.
func (r *DirVerifyEmailService) Confirm(ctx context.Context, dirID string, body DirVerifyEmailConfirmParams, opts ...option.RequestOption) (res *EmailVerificationStatusWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/verify_email/confirm", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EmailVerificationStatusWrapped struct {
	// Verification state for a DIR's authorizer email.
	Data EmailVerificationStatusWrappedData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailVerificationStatusWrapped) RawJSON() string { return r.JSON.raw }
func (r *EmailVerificationStatusWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification state for a DIR's authorizer email.
type EmailVerificationStatusWrappedData struct {
	// Whether the DIR's authorizer email has been confirmed.
	EmailVerified bool `json:"email_verified" api:"required"`
	// Always `email_verification`.
	//
	// Any of "email_verification".
	RecordType string `json:"record_type" api:"required"`
	// `sent` after a code is emailed; `verified` after a successful confirm;
	// `unverified` when no verification is in progress.
	//
	// Any of "sent", "verified", "unverified".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailVerified respjson.Field
		RecordType    respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailVerificationStatusWrappedData) RawJSON() string { return r.JSON.raw }
func (r *EmailVerificationStatusWrappedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirVerifyEmailConfirmParams struct {
	// The 6-digit code sent to the authorizer email.
	Code string `json:"code" api:"required"`
	paramObj
}

func (r DirVerifyEmailConfirmParams) MarshalJSON() (data []byte, err error) {
	type shadow DirVerifyEmailConfirmParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirVerifyEmailConfirmParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
