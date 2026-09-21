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
func (r *WhatsappBusinessAccountPhoneNumberService) List(ctx context.Context, id string, query WhatsappBusinessAccountPhoneNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappBusinessAccountPhoneNumberListResponse], err error) {
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
func (r *WhatsappBusinessAccountPhoneNumberService) ListAutoPaging(ctx context.Context, id string, query WhatsappBusinessAccountPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappBusinessAccountPhoneNumberListResponse] {
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

type WhatsappBusinessAccountPhoneNumberListResponse struct {
	CallingEnabled bool `json:"calling_enabled"`
	// Current lifecycle state for a coexistence number. This is null for a standard
	// Cloud API number.
	//
	// Any of "pending_onboarding", "sync_pending", "syncing", "sync_complete",
	// "active", "history_declined", "sync_deadline_expired", "offboarded",
	// "disconnected".
	CoexistenceState WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState `json:"coexistence_state" api:"nullable"`
	CreatedAt        time.Time                                                      `json:"created_at" format:"date-time"`
	DisplayName      string                                                         `json:"display_name"`
	Enabled          bool                                                           `json:"enabled"`
	// Indicates whether the number is connected to both the WhatsApp Business app and
	// Cloud API through WhatsApp Coexistence.
	IsOnBizApp bool `json:"is_on_biz_app"`
	// Phone number in E164 format
	PhoneNumber string `json:"phone_number"`
	// Whatsapp phone number ID
	PhoneNumberID string `json:"phone_number_id"`
	// Whatsapp quality rating
	QualityRating string `json:"quality_rating"`
	RecordType    string `json:"record_type"`
	Status        string `json:"status"`
	// Deadline for initiating the current coexistence synchronization cycle. This is
	// null when no deadline applies.
	SyncDeadline time.Time `json:"sync_deadline" api:"nullable" format:"date-time"`
	// Synchronization progress. This object is returned only while a coexistence
	// number is synchronizing.
	SyncProgress WhatsappBusinessAccountPhoneNumberListResponseSyncProgress `json:"sync_progress" api:"nullable"`
	// User ID
	UserID string `json:"user_id"`
	// WABA ID of Whatsapp business account
	WabaID string `json:"waba_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallingEnabled   respjson.Field
		CoexistenceState respjson.Field
		CreatedAt        respjson.Field
		DisplayName      respjson.Field
		Enabled          respjson.Field
		IsOnBizApp       respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		QualityRating    respjson.Field
		RecordType       respjson.Field
		Status           respjson.Field
		SyncDeadline     respjson.Field
		SyncProgress     respjson.Field
		UserID           respjson.Field
		WabaID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappBusinessAccountPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state for a coexistence number. This is null for a standard
// Cloud API number.
type WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState string

const (
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStatePendingOnboarding   WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "pending_onboarding"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateSyncPending         WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "sync_pending"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateSyncing             WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "syncing"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateSyncComplete        WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "sync_complete"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateActive              WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "active"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateHistoryDeclined     WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "history_declined"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateSyncDeadlineExpired WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "sync_deadline_expired"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateOffboarded          WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "offboarded"
	WhatsappBusinessAccountPhoneNumberListResponseCoexistenceStateDisconnected        WhatsappBusinessAccountPhoneNumberListResponseCoexistenceState = "disconnected"
)

// Synchronization progress. This object is returned only while a coexistence
// number is synchronizing.
type WhatsappBusinessAccountPhoneNumberListResponseSyncProgress struct {
	ContactsStatus    string `json:"contacts_status"`
	HistoryChunkOrder int64  `json:"history_chunk_order" api:"nullable"`
	HistoryPhase      int64  `json:"history_phase" api:"nullable"`
	HistoryProgress   int64  `json:"history_progress" api:"nullable"`
	HistoryStatus     string `json:"history_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactsStatus    respjson.Field
		HistoryChunkOrder respjson.Field
		HistoryPhase      respjson.Field
		HistoryProgress   respjson.Field
		HistoryStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappBusinessAccountPhoneNumberListResponseSyncProgress) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappBusinessAccountPhoneNumberListResponseSyncProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
