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

// Two factor authentication API
//
// VerificationByPhoneNumberActionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationByPhoneNumberActionService] method instead.
type VerificationByPhoneNumberActionService struct {
	Options []option.RequestOption
}

// NewVerificationByPhoneNumberActionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVerificationByPhoneNumberActionService(opts ...option.RequestOption) (r VerificationByPhoneNumberActionService) {
	r = VerificationByPhoneNumberActionService{}
	r.Options = opts
	return
}

// Verify verification code by phone number
func (r *VerificationByPhoneNumberActionService) Verify(ctx context.Context, phoneNumber string, body VerificationByPhoneNumberActionVerifyParams, opts ...option.RequestOption) (res *VerifyVerificationCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("verifications/by_phone_number/%s/actions/verify", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VerifyVerificationCodeResponse struct {
	Data VerifyVerificationCodeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyVerificationCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifyVerificationCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyVerificationCodeResponseData struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number" api:"required"`
	// Identifies if the verification code has been accepted or rejected.
	//
	// Any of "accepted", "rejected".
	ResponseCode string `json:"response_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber  respjson.Field
		ResponseCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyVerificationCodeResponseData) RawJSON() string { return r.JSON.raw }
func (r *VerifyVerificationCodeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerificationByPhoneNumberActionVerifyParams struct {
	// This is the code the user submits for verification.
	Code string `json:"code" api:"required"`
	// The identifier of the associated Verify profile.
	VerifyProfileID string `json:"verify_profile_id" api:"required" format:"uuid"`
	paramObj
}

func (r VerificationByPhoneNumberActionVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow VerificationByPhoneNumberActionVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationByPhoneNumberActionVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
