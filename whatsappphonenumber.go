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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage Whatsapp phone numbers
//
// WhatsappPhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberService] method instead.
type WhatsappPhoneNumberService struct {
	Options []option.RequestOption
	// Manage Whatsapp phone numbers
	CallingSettings WhatsappPhoneNumberCallingSettingService
	// Manage Whatsapp phone numbers
	Profile WhatsappPhoneNumberProfileService
}

// NewWhatsappPhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberService(opts ...option.RequestOption) (r WhatsappPhoneNumberService) {
	r = WhatsappPhoneNumberService{}
	r.Options = opts
	r.CallingSettings = NewWhatsappPhoneNumberCallingSettingService(opts...)
	r.Profile = NewWhatsappPhoneNumberProfileService(opts...)
	return
}

// List Whatsapp phone numbers
func (r *WhatsappPhoneNumberService) List(ctx context.Context, query WhatsappPhoneNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappPhoneNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/whatsapp/phone_numbers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Whatsapp phone numbers
func (r *WhatsappPhoneNumberService) ListAutoPaging(ctx context.Context, query WhatsappPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappPhoneNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Whatsapp phone number
func (r *WhatsappPhoneNumberService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Resend verification code
func (r *WhatsappPhoneNumberService) ResendVerification(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberResendVerificationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/resend_verification", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Submit verification code for a phone number
func (r *WhatsappPhoneNumberService) Verify(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberVerifyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/verify", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type WhatsappPhoneNumberListResponse struct {
	CallingEnabled bool      `json:"calling_enabled"`
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	DisplayName    string    `json:"display_name"`
	Enabled        bool      `json:"enabled"`
	// Phone number in E164 format
	PhoneNumber string `json:"phone_number"`
	// Whatsapp phone number ID
	PhoneNumberID string `json:"phone_number_id"`
	// Whatsapp quality rating
	QualityRating string `json:"quality_rating"`
	RecordType    string `json:"record_type"`
	Status        string `json:"status"`
	// User ID
	UserID string `json:"user_id"`
	// WABA ID of Whatsapp business account
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallingEnabled respjson.Field
		CreatedAt      respjson.Field
		DisplayName    respjson.Field
		Enabled        respjson.Field
		PhoneNumber    respjson.Field
		PhoneNumberID  respjson.Field
		QualityRating  respjson.Field
		RecordType     respjson.Field
		Status         respjson.Field
		UserID         respjson.Field
		WabaID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappPhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r WhatsappPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappPhoneNumberResendVerificationParams struct {
	// Any of "sms", "voice".
	VerificationMethod WhatsappPhoneNumberResendVerificationParamsVerificationMethod `json:"verification_method,omitzero"`
	paramObj
}

func (r WhatsappPhoneNumberResendVerificationParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberResendVerificationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberResendVerificationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberResendVerificationParamsVerificationMethod string

const (
	WhatsappPhoneNumberResendVerificationParamsVerificationMethodSMS   WhatsappPhoneNumberResendVerificationParamsVerificationMethod = "sms"
	WhatsappPhoneNumberResendVerificationParamsVerificationMethodVoice WhatsappPhoneNumberResendVerificationParamsVerificationMethod = "voice"
)

type WhatsappPhoneNumberVerifyParams struct {
	Code string `json:"code" api:"required"`
	paramObj
}

func (r WhatsappPhoneNumberVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
