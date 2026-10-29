// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// WhatsappBusinessAccountPhoneNumberService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappBusinessAccountPhoneNumberService] method instead.
type WhatsappBusinessAccountPhoneNumberService struct {
	Options []option.RequestOption
}

// NewWhatsappBusinessAccountPhoneNumberService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWhatsappBusinessAccountPhoneNumberService(opts ...option.RequestOption) (r WhatsappBusinessAccountPhoneNumberService) {
	r = WhatsappBusinessAccountPhoneNumberService{}
	r.Options = opts
	return
}

// Returns phone numbers registered under the specified WhatsApp Business Account.
func (r *WhatsappBusinessAccountPhoneNumberService) List(ctx context.Context, id string, query WhatsappBusinessAccountPhoneNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappPhoneResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s/phone_numbers", url.PathEscape(id))
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

// Returns phone numbers registered under the specified WhatsApp Business Account.
func (r *WhatsappBusinessAccountPhoneNumberService) ListAutoPaging(ctx context.Context, id string, query WhatsappBusinessAccountPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappPhoneResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Starts verification of a phone number for the specified WhatsApp Business
// Account using the requested verification method.
func (r *WhatsappBusinessAccountPhoneNumberService) InitializeVerification(ctx context.Context, id string, body WhatsappBusinessAccountPhoneNumberInitializeVerificationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s/phone_numbers", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type WhatsappBusinessAccountPhoneNumberListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappBusinessAccountPhoneNumberListParams]'s query
// parameters as `url.Values`.
func (r WhatsappBusinessAccountPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappBusinessAccountPhoneNumberInitializeVerificationParams struct {
	DisplayName string            `json:"display_name" api:"required"`
	PhoneNumber string            `json:"phone_number" api:"required"`
	Language    param.Opt[string] `json:"language,omitzero"`
	// Any of "sms", "voice".
	VerificationMethod WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethod `json:"verification_method,omitzero"`
	paramObj
}

func (r WhatsappBusinessAccountPhoneNumberInitializeVerificationParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappBusinessAccountPhoneNumberInitializeVerificationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappBusinessAccountPhoneNumberInitializeVerificationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethod string

const (
	WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethodSMS   WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethod = "sms"
	WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethodVoice WhatsappBusinessAccountPhoneNumberInitializeVerificationParamsVerificationMethod = "voice"
)
