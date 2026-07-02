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

// Manage Whatsapp business accounts
//
// WhatsappBusinessAccountService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappBusinessAccountService] method instead.
type WhatsappBusinessAccountService struct {
	Options      []option.RequestOption
	PhoneNumbers WhatsappBusinessAccountPhoneNumberService
	// Manage Whatsapp business accounts
	Settings WhatsappBusinessAccountSettingService
}

// NewWhatsappBusinessAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappBusinessAccountService(opts ...option.RequestOption) (r WhatsappBusinessAccountService) {
	r = WhatsappBusinessAccountService{}
	r.Options = opts
	r.PhoneNumbers = NewWhatsappBusinessAccountPhoneNumberService(opts...)
	r.Settings = NewWhatsappBusinessAccountSettingService(opts...)
	return
}

// Get a single Whatsapp Business Account
func (r *WhatsappBusinessAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappBusinessAccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Whatsapp Business Accounts
func (r *WhatsappBusinessAccountService) List(ctx context.Context, query WhatsappBusinessAccountListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappBusinessAccountListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/whatsapp/business_accounts"
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

// List Whatsapp Business Accounts
func (r *WhatsappBusinessAccountService) ListAutoPaging(ctx context.Context, query WhatsappBusinessAccountListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappBusinessAccountListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Whatsapp Business Account
func (r *WhatsappBusinessAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/business_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WhatsappBusinessAccountGetResponse struct {
	Data WhatsappBusinessAccountGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountGetResponseData struct {
	// Internal ID of Whatsapp business account
	ID string `json:"id" format:"uuid"`
	// Account review status of Whatsapp business account
	AccountReviewStatus string `json:"account_review_status"`
	// Business verification status of Whatsapp business account
	BusinessVerificationStatus string `json:"business_verification_status"`
	// Country associated with Whatsapp business account
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of Whatsapp business account
	Name string `json:"name"`
	// Count of phone numbers associated with Whatsapp business account
	PhoneNumbersCount int64  `json:"phone_numbers_count"`
	RecordType        string `json:"record_type"`
	// Status of Whatsapp business account
	Status string `json:"status"`
	// WABA ID of Whatsapp business account
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AccountReviewStatus        respjson.Field
		BusinessVerificationStatus respjson.Field
		Country                    respjson.Field
		CreatedAt                  respjson.Field
		Name                       respjson.Field
		PhoneNumbersCount          respjson.Field
		RecordType                 respjson.Field
		Status                     respjson.Field
		WabaID                     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountListResponse struct {
	// Internal ID of Whatsapp business account
	ID string `json:"id" format:"uuid"`
	// Account review status of Whatsapp business account
	AccountReviewStatus string `json:"account_review_status"`
	// Business verification status of Whatsapp business account
	BusinessVerificationStatus string `json:"business_verification_status"`
	// Country associated with Whatsapp business account
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Name of Whatsapp business account
	Name string `json:"name"`
	// Count of phone numbers associated with Whatsapp business account
	PhoneNumbersCount int64  `json:"phone_numbers_count"`
	RecordType        string `json:"record_type"`
	// Status of Whatsapp business account
	Status string `json:"status"`
	// WABA ID of Whatsapp business account
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AccountReviewStatus        respjson.Field
		BusinessVerificationStatus respjson.Field
		Country                    respjson.Field
		CreatedAt                  respjson.Field
		Name                       respjson.Field
		PhoneNumbersCount          respjson.Field
		RecordType                 respjson.Field
		Status                     respjson.Field
		WabaID                     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappBusinessAccountListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappBusinessAccountListParams]'s query parameters as
// `url.Values`.
func (r WhatsappBusinessAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
