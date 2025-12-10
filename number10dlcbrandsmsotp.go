// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcBrandSMSOtpService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcBrandSMSOtpService] method instead.
type Number10dlcBrandSMSOtpService struct {
	Options []option.RequestOption
}

// NewNumber10dlcBrandSMSOtpService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumber10dlcBrandSMSOtpService(opts ...option.RequestOption) (r Number10dlcBrandSMSOtpService) {
	r = Number10dlcBrandSMSOtpService{}
	r.Options = opts
	return
}

// Query the status of an SMS OTP (One-Time Password) for Sole Proprietor brand
// verification.
//
// This endpoint allows you to check the delivery and verification status of an OTP
// sent during the Sole Proprietor brand verification process. You can query by
// either:
//
// - `referenceId` - The reference ID returned when the OTP was initially triggered
// - `brandId` - Query parameter for portal users to look up OTP status by Brand ID
//
// The response includes delivery status, verification dates, and detailed delivery
// information.
func (r *Number10dlcBrandSMSOtpService) GetStatus(ctx context.Context, referenceID string, query Number10dlcBrandSMSOtpGetStatusParams, opts ...option.RequestOption) (res *Number10dlcBrandSMSOtpGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if referenceID == "" {
		err = errors.New("missing required referenceId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/smsOtp/%s", referenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Trigger or re-trigger an SMS OTP (One-Time Password) for Sole Proprietor brand
// verification.
//
// **Important Notes:**
//
//   - Only allowed for Sole Proprietor (`SOLE_PROPRIETOR`) brands
//   - Triggers generation of a one-time password sent to the `mobilePhone` number in
//     the brand's profile
//   - Campaigns cannot be created until OTP verification is complete
//   - US/CA numbers only for real OTPs; mock brands can use non-US/CA numbers for
//     testing
//   - Returns a `referenceId` that can be used to check OTP status via the GET
//     `/10dlc/brand/smsOtp/{referenceId}` endpoint
//
// **Use Cases:**
//
// - Initial OTP trigger after Sole Proprietor brand creation
// - Re-triggering OTP if the user didn't receive or needs a new code
func (r *Number10dlcBrandSMSOtpService) Trigger(ctx context.Context, brandID string, body Number10dlcBrandSMSOtpTriggerParams, opts ...option.RequestOption) (res *Number10dlcBrandSMSOtpTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/smsOtp", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Verify the SMS OTP (One-Time Password) for Sole Proprietor brand verification.
//
// **Verification Flow:**
//
// 1. User receives OTP via SMS after triggering
// 2. User submits the OTP pin through this endpoint
// 3. Upon successful verification:
//   - A `BRAND_OTP_VERIFIED` webhook event is sent to the CSP
//   - The brand's `identityStatus` changes to `VERIFIED`
//   - Campaigns can now be created for this brand
//
// **Error Handling:**
//
// Provides proper error responses for:
//
// - Invalid OTP pins
// - Expired OTPs
// - OTP verification failures
func (r *Number10dlcBrandSMSOtpService) Verify(ctx context.Context, brandID string, body Number10dlcBrandSMSOtpVerifyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/smsOtp", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Status information for an SMS OTP sent during Sole Proprietor brand verification
type Number10dlcBrandSMSOtpGetStatusResponse struct {
	// The Brand ID associated with this OTP request
	BrandID string `json:"brandId,required"`
	// The current delivery status of the OTP SMS message. Common values include:
	// `DELIVERED_HANDSET`, `PENDING`, `FAILED`, `EXPIRED`
	DeliveryStatus string `json:"deliveryStatus,required"`
	// The mobile phone number where the OTP was sent, in E.164 format
	MobilePhone string `json:"mobilePhone,required"`
	// The reference ID for this OTP request, used for status queries
	ReferenceID string `json:"referenceId,required"`
	// The timestamp when the OTP request was initiated
	RequestDate time.Time `json:"requestDate,required" format:"date-time"`
	// The timestamp when the delivery status was last updated
	DeliveryStatusDate time.Time `json:"deliveryStatusDate" format:"date-time"`
	// Additional details about the delivery status
	DeliveryStatusDetails string `json:"deliveryStatusDetails"`
	// The timestamp when the OTP was successfully verified (if applicable)
	VerifyDate time.Time `json:"verifyDate" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID               respjson.Field
		DeliveryStatus        respjson.Field
		MobilePhone           respjson.Field
		ReferenceID           respjson.Field
		RequestDate           respjson.Field
		DeliveryStatusDate    respjson.Field
		DeliveryStatusDetails respjson.Field
		VerifyDate            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandSMSOtpGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandSMSOtpGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response after successfully triggering a Brand SMS OTP
type Number10dlcBrandSMSOtpTriggerResponse struct {
	// The Brand ID for which the OTP was triggered
	BrandID string `json:"brandId,required"`
	// The reference ID that can be used to check OTP status
	ReferenceID string `json:"referenceId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID     respjson.Field
		ReferenceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandSMSOtpTriggerResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandSMSOtpTriggerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandSMSOtpGetStatusParams struct {
	// Filter by Brand ID for easier lookup in portal applications
	BrandID param.Opt[string] `query:"brandId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcBrandSMSOtpGetStatusParams]'s query parameters
// as `url.Values`.
func (r Number10dlcBrandSMSOtpGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type Number10dlcBrandSMSOtpTriggerParams struct {
	// SMS message template to send the OTP. Must include `@OTP_PIN@` placeholder which
	// will be replaced with the actual PIN
	PinSMS string `json:"pinSms,required"`
	// SMS message to send upon successful OTP verification
	SuccessSMS string `json:"successSms,required"`
	paramObj
}

func (r Number10dlcBrandSMSOtpTriggerParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandSMSOtpTriggerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandSMSOtpTriggerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandSMSOtpVerifyParams struct {
	// The OTP PIN received via SMS
	OtpPin string `json:"otpPin,required"`
	paramObj
}

func (r Number10dlcBrandSMSOtpVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandSMSOtpVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandSMSOtpVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
