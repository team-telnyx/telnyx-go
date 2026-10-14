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

// WhatsappSignupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappSignupService] method instead.
type WhatsappSignupService struct {
	Options []option.RequestOption
}

// NewWhatsappSignupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWhatsappSignupService(opts ...option.RequestOption) (r WhatsappSignupService) {
	r = WhatsappSignupService{}
	r.Options = opts
	return
}

// List finished WhatsApp Embedded Signup sessions, with status `SUCCESS` or
// `FAILURE`, newest first. Use the `errors` of a failed session to see why
// provisioning stopped.
func (r *WhatsappSignupService) List(ctx context.Context, query WhatsappSignupListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappSignupListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "whatsapp/signup"
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

// List finished WhatsApp Embedded Signup sessions, with status `SUCCESS` or
// `FAILURE`, newest first. Use the `errors` of a failed session to see why
// provisioning stopped.
func (r *WhatsappSignupService) ListAutoPaging(ctx context.Context, query WhatsappSignupListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappSignupListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Retrieve the current state and status of a WhatsApp Embedded Signup session.
func (r *WhatsappSignupService) GetStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappSignupGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("whatsapp/signup/%s/status", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WhatsappSignupListResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Tech provider customer reference submitted with the signup
	CustomerID string `json:"customer_id"`
	// Errors recorded during provisioning
	Errors []string `json:"errors"`
	// Any of "whatsapp_session_id".
	RecordType WhatsappSignupListResponseRecordType `json:"record_type"`
	// Signup session ID
	SessionID string `json:"session_id"`
	// Current provisioning step, or the step that failed. `REGISTER_PHONE_NUMBER` is
	// skipped for coexistence signups.
	//
	// Any of "SUBSCRIBE_WEBHOOKS", "ASSIGNED_USERS", "SHARE_CREDIT_LINE",
	// "REGISTER_PHONE_NUMBER", "ADD_WABA_TO_DB", "ADD_PHONE_NUMBERS_TO_DB",
	// "COMPLETE".
	State WhatsappSignupListResponseState `json:"state"`
	// `INCOMPLETE` waits for the Meta signup code, `PENDING` is in progress, `SUCCESS`
	// finished provisioning, and `FAILURE` stopped at `state`.
	//
	// Any of "INCOMPLETE", "PENDING", "SUCCESS", "FAILURE".
	Status    WhatsappSignupListResponseStatus `json:"status"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	// Telnyx user ID that owns the signup session
	UserID string `json:"user_id"`
	// Meta WhatsApp Business Account ID
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		CustomerID  respjson.Field
		Errors      respjson.Field
		RecordType  respjson.Field
		SessionID   respjson.Field
		State       respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		WabaID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappSignupListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappSignupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappSignupListResponseRecordType string

const (
	WhatsappSignupListResponseRecordTypeWhatsappSessionID WhatsappSignupListResponseRecordType = "whatsapp_session_id"
)

// Current provisioning step, or the step that failed. `REGISTER_PHONE_NUMBER` is
// skipped for coexistence signups.
type WhatsappSignupListResponseState string

const (
	WhatsappSignupListResponseStateSubscribeWebhooks   WhatsappSignupListResponseState = "SUBSCRIBE_WEBHOOKS"
	WhatsappSignupListResponseStateAssignedUsers       WhatsappSignupListResponseState = "ASSIGNED_USERS"
	WhatsappSignupListResponseStateShareCreditLine     WhatsappSignupListResponseState = "SHARE_CREDIT_LINE"
	WhatsappSignupListResponseStateRegisterPhoneNumber WhatsappSignupListResponseState = "REGISTER_PHONE_NUMBER"
	WhatsappSignupListResponseStateAddWabaToDB         WhatsappSignupListResponseState = "ADD_WABA_TO_DB"
	WhatsappSignupListResponseStateAddPhoneNumbersToDB WhatsappSignupListResponseState = "ADD_PHONE_NUMBERS_TO_DB"
	WhatsappSignupListResponseStateComplete            WhatsappSignupListResponseState = "COMPLETE"
)

// `INCOMPLETE` waits for the Meta signup code, `PENDING` is in progress, `SUCCESS`
// finished provisioning, and `FAILURE` stopped at `state`.
type WhatsappSignupListResponseStatus string

const (
	WhatsappSignupListResponseStatusIncomplete WhatsappSignupListResponseStatus = "INCOMPLETE"
	WhatsappSignupListResponseStatusPending    WhatsappSignupListResponseStatus = "PENDING"
	WhatsappSignupListResponseStatusSuccess    WhatsappSignupListResponseStatus = "SUCCESS"
	WhatsappSignupListResponseStatusFailure    WhatsappSignupListResponseStatus = "FAILURE"
)

type WhatsappSignupGetStatusResponse struct {
	Data WhatsappSignupGetStatusResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappSignupGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappSignupGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappSignupGetStatusResponseData struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Tech provider customer reference submitted with the signup
	CustomerID string `json:"customer_id"`
	// Errors recorded during provisioning
	Errors []string `json:"errors"`
	// Any of "whatsapp_session_id".
	RecordType string `json:"record_type"`
	// Signup session ID
	SessionID string `json:"session_id"`
	// Current provisioning step, or the step that failed. `REGISTER_PHONE_NUMBER` is
	// skipped for coexistence signups.
	//
	// Any of "SUBSCRIBE_WEBHOOKS", "ASSIGNED_USERS", "SHARE_CREDIT_LINE",
	// "REGISTER_PHONE_NUMBER", "ADD_WABA_TO_DB", "ADD_PHONE_NUMBERS_TO_DB",
	// "COMPLETE".
	State string `json:"state"`
	// `INCOMPLETE` waits for the Meta signup code, `PENDING` is in progress, `SUCCESS`
	// finished provisioning, and `FAILURE` stopped at `state`.
	//
	// Any of "INCOMPLETE", "PENDING", "SUCCESS", "FAILURE".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Telnyx user ID that owns the signup session
	UserID string `json:"user_id"`
	// Meta WhatsApp Business Account ID
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		CustomerID  respjson.Field
		Errors      respjson.Field
		RecordType  respjson.Field
		SessionID   respjson.Field
		State       respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		WabaID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappSignupGetStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappSignupGetStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappSignupListParams struct {
	// Return sessions created at or after this time (RFC 3339)
	FilterCreatedAtGte param.Opt[time.Time] `query:"filter[created_at][gte],omitzero" format:"date-time" json:"-"`
	// Return sessions created at or before this time (RFC 3339)
	FilterCreatedAtLte param.Opt[time.Time] `query:"filter[created_at][lte],omitzero" format:"date-time" json:"-"`
	// Filter by the tech provider customer reference
	FilterCustomerID param.Opt[string] `query:"filter[customer_id],omitzero" json:"-"`
	// Filter by Meta WABA ID
	FilterWabaID param.Opt[string] `query:"filter[waba_id],omitzero" json:"-"`
	PageNumber   param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize     param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappSignupListParams]'s query parameters as
// `url.Values`.
func (r WhatsappSignupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
