// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
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

// Submit the 6-digit code that was emailed to the DIR's authorizer email. On
// success the authorizer email is marked verified.
//
// For security, any failure (wrong, expired, already-used, or too many attempts)
// returns the same generic message.
func (r *DirVerifyEmailService) ConfirmCode(ctx context.Context, dirID string, body DirVerifyEmailConfirmCodeParams, opts ...option.RequestOption) (res *DirVerifyEmailConfirmCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/verify_email/confirm", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Email a 6-digit code to the DIR's authorizer email to confirm ownership of that
// address.
//
// The code expires in 15 minutes. Requesting a new code invalidates any previous
// one. Resends are rate limited (a short cooldown plus a daily cap). Submit the
// code to `POST /dir/{dir_id}/verify_email/confirm`.
func (r *DirVerifyEmailService) SendCode(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirVerifyEmailSendCodeResponse, err error) {
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
func (r *DirVerifyEmailService) Status(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirVerifyEmailStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/verify_email", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DirVerifyEmailConfirmCodeResponse struct {
	// Verification state for a DIR's authorizer email.
	Data DirVerifyEmailConfirmCodeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirVerifyEmailConfirmCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailConfirmCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification state for a DIR's authorizer email.
type DirVerifyEmailConfirmCodeResponseData struct {
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
func (r DirVerifyEmailConfirmCodeResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailConfirmCodeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirVerifyEmailSendCodeResponse struct {
	// Verification state for a DIR's authorizer email.
	Data DirVerifyEmailSendCodeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirVerifyEmailSendCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailSendCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification state for a DIR's authorizer email.
type DirVerifyEmailSendCodeResponseData struct {
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
func (r DirVerifyEmailSendCodeResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailSendCodeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirVerifyEmailStatusResponse struct {
	// Verification state for a DIR's authorizer email.
	Data DirVerifyEmailStatusResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirVerifyEmailStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification state for a DIR's authorizer email.
type DirVerifyEmailStatusResponseData struct {
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
func (r DirVerifyEmailStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirVerifyEmailStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirVerifyEmailConfirmCodeParams struct {
	// The 6-digit code sent to the authorizer email.
	Code string `json:"code" api:"required"`
	paramObj
}

func (r DirVerifyEmailConfirmCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow DirVerifyEmailConfirmCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirVerifyEmailConfirmCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
