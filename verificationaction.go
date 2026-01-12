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
)

// VerificationActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationActionService] method instead.
type VerificationActionService struct {
	Options []option.RequestOption
}

// NewVerificationActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVerificationActionService(opts ...option.RequestOption) (r VerificationActionService) {
	r = VerificationActionService{}
	r.Options = opts
	return
}

// Verify verification code by ID
func (r *VerificationActionService) Verify(ctx context.Context, verificationID string, body VerificationActionVerifyParams, opts ...option.RequestOption) (res *VerifyVerificationCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verification_id parameter")
		return
	}
	path := fmt.Sprintf("verifications/%s/actions/verify", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VerificationActionVerifyParams struct {
	// This is the code the user submits for verification.
	Code param.Opt[string] `json:"code,omitzero"`
	// Identifies if the verification code has been accepted or rejected. Only
	// permitted if custom_code was used for the verification.
	//
	// Any of "accepted", "rejected".
	Status VerificationActionVerifyParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r VerificationActionVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow VerificationActionVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationActionVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies if the verification code has been accepted or rejected. Only
// permitted if custom_code was used for the verification.
type VerificationActionVerifyParamsStatus string

const (
	VerificationActionVerifyParamsStatusAccepted VerificationActionVerifyParamsStatus = "accepted"
	VerificationActionVerifyParamsStatusRejected VerificationActionVerifyParamsStatus = "rejected"
)
