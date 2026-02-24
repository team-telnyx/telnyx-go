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

// VerifiedNumberActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifiedNumberActionService] method instead.
type VerifiedNumberActionService struct {
	Options []option.RequestOption
}

// NewVerifiedNumberActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVerifiedNumberActionService(opts ...option.RequestOption) (r VerifiedNumberActionService) {
	r = VerifiedNumberActionService{}
	r.Options = opts
	return
}

// Submit verification code
func (r *VerifiedNumberActionService) SubmitVerificationCode(ctx context.Context, phoneNumber string, body VerifiedNumberActionSubmitVerificationCodeParams, opts ...option.RequestOption) (res *VerifiedNumberDataWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("verified_numbers/%s/actions/verify", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VerifiedNumberActionSubmitVerificationCodeParams struct {
	VerificationCode string `json:"verification_code" api:"required"`
	paramObj
}

func (r VerifiedNumberActionSubmitVerificationCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow VerifiedNumberActionSubmitVerificationCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerifiedNumberActionSubmitVerificationCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
