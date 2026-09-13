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
	"github.com/team-telnyx/telnyx-go/v4/shared"
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
	// Manage Whatsapp phone numbers
	ConversationalComponents WhatsappPhoneNumberConversationalComponentService
}

// NewWhatsappPhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberService(opts ...option.RequestOption) (r WhatsappPhoneNumberService) {
	r = WhatsappPhoneNumberService{}
	r.Options = opts
	r.CallingSettings = NewWhatsappPhoneNumberCallingSettingService(opts...)
	r.Profile = NewWhatsappPhoneNumberProfileService(opts...)
	r.ConversationalComponents = NewWhatsappPhoneNumberConversationalComponentService(opts...)
	return
}

// Returns WhatsApp phone numbers linked to the authenticated Telnyx account.
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

// Returns WhatsApp phone numbers linked to the authenticated Telnyx account.
func (r *WhatsappPhoneNumberService) ListAutoPaging(ctx context.Context, query WhatsappPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappPhoneNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Removes the specified phone number from Telnyx WhatsApp management.
func (r *WhatsappPhoneNumberService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a list of the phone numbers registered for WhatsApp on your account.
func (r *WhatsappPhoneNumberService) Get(ctx context.Context, query WhatsappPhoneNumberGetParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "whatsapp/phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Requests a new verification code for the specified WhatsApp phone number.
func (r *WhatsappPhoneNumberService) ResendVerification(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberResendVerificationParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/resend_verification", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Returns whether the 24-hour conversation window is currently open for a given
// source/destination pair. If window_active is false, only template messages may
// be sent.
func (r *WhatsappPhoneNumberService) GetConversationWindow(ctx context.Context, phoneNumber string, query WhatsappPhoneNumberGetConversationWindowParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberGetConversationWindowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/conversation_window", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns one WhatsApp phone number linked to the authenticated Telnyx account.
// For a coexistence number in the `syncing` state, the response includes
// `sync_progress`.
func (r *WhatsappPhoneNumberService) GetPhoneNumber(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *WhatsappPhoneNumberGetPhoneNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("whatsapp/phone_numbers/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submits the verification code received for the specified WhatsApp phone number.
func (r *WhatsappPhoneNumberService) Verify(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberVerifyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/verify", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type WhatsappPhoneNumberListResponse struct {
	CallingEnabled bool `json:"calling_enabled"`
	// Current lifecycle state for a coexistence number. This is null for a standard
	// Cloud API number.
	//
	// Any of "pending_onboarding", "sync_pending", "syncing", "sync_complete",
	// "active", "history_declined", "sync_deadline_expired", "offboarded",
	// "disconnected".
	CoexistenceState WhatsappPhoneNumberListResponseCoexistenceState `json:"coexistence_state" api:"nullable"`
	CreatedAt        time.Time                                       `json:"created_at" format:"date-time"`
	DisplayName      string                                          `json:"display_name"`
	Enabled          bool                                            `json:"enabled"`
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
	SyncProgress WhatsappPhoneNumberListResponseSyncProgress `json:"sync_progress" api:"nullable"`
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
func (r WhatsappPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current lifecycle state for a coexistence number. This is null for a standard
// Cloud API number.
type WhatsappPhoneNumberListResponseCoexistenceState string

const (
	WhatsappPhoneNumberListResponseCoexistenceStatePendingOnboarding   WhatsappPhoneNumberListResponseCoexistenceState = "pending_onboarding"
	WhatsappPhoneNumberListResponseCoexistenceStateSyncPending         WhatsappPhoneNumberListResponseCoexistenceState = "sync_pending"
	WhatsappPhoneNumberListResponseCoexistenceStateSyncing             WhatsappPhoneNumberListResponseCoexistenceState = "syncing"
	WhatsappPhoneNumberListResponseCoexistenceStateSyncComplete        WhatsappPhoneNumberListResponseCoexistenceState = "sync_complete"
	WhatsappPhoneNumberListResponseCoexistenceStateActive              WhatsappPhoneNumberListResponseCoexistenceState = "active"
	WhatsappPhoneNumberListResponseCoexistenceStateHistoryDeclined     WhatsappPhoneNumberListResponseCoexistenceState = "history_declined"
	WhatsappPhoneNumberListResponseCoexistenceStateSyncDeadlineExpired WhatsappPhoneNumberListResponseCoexistenceState = "sync_deadline_expired"
	WhatsappPhoneNumberListResponseCoexistenceStateOffboarded          WhatsappPhoneNumberListResponseCoexistenceState = "offboarded"
	WhatsappPhoneNumberListResponseCoexistenceStateDisconnected        WhatsappPhoneNumberListResponseCoexistenceState = "disconnected"
)

// Synchronization progress. This object is returned only while a coexistence
// number is synchronizing.
type WhatsappPhoneNumberListResponseSyncProgress struct {
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
func (r WhatsappPhoneNumberListResponseSyncProgress) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberListResponseSyncProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetResponse struct {
	Data []WhatsappPhoneNumberGetResponseData `json:"data"`
	Meta shared.MessagingPaginationMeta       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetResponseData struct {
	CallingEnabled bool `json:"calling_enabled"`
	// Current lifecycle state for a coexistence number. This is null for a standard
	// Cloud API number.
	//
	// Any of "pending_onboarding", "sync_pending", "syncing", "sync_complete",
	// "active", "history_declined", "sync_deadline_expired", "offboarded",
	// "disconnected".
	CoexistenceState string    `json:"coexistence_state" api:"nullable"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	DisplayName      string    `json:"display_name"`
	Enabled          bool      `json:"enabled"`
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
	SyncProgress WhatsappPhoneNumberGetResponseDataSyncProgress `json:"sync_progress" api:"nullable"`
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
func (r WhatsappPhoneNumberGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Synchronization progress. This object is returned only while a coexistence
// number is synchronizing.
type WhatsappPhoneNumberGetResponseDataSyncProgress struct {
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
func (r WhatsappPhoneNumberGetResponseDataSyncProgress) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetResponseDataSyncProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetConversationWindowResponse struct {
	Data WhatsappPhoneNumberGetConversationWindowResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberGetConversationWindowResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetConversationWindowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetConversationWindowResponseData struct {
	// Timestamp of the last inbound message that opened the window
	LastUserMessageAt time.Time `json:"last_user_message_at" format:"date-time"`
	// Whether the 24-hour conversation window is currently open
	WindowActive bool `json:"window_active"`
	// When the window closes. Null if no active window.
	WindowExpiresAt time.Time `json:"window_expires_at" api:"nullable" format:"date-time"`
	// Window type. Currently always 24h when present.
	WindowType string `json:"window_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUserMessageAt respjson.Field
		WindowActive      respjson.Field
		WindowExpiresAt   respjson.Field
		WindowType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberGetConversationWindowResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetConversationWindowResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetPhoneNumberResponse struct {
	Data WhatsappPhoneNumberGetPhoneNumberResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberGetPhoneNumberResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetPhoneNumberResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberGetPhoneNumberResponseData struct {
	CallingEnabled bool `json:"calling_enabled"`
	// Current lifecycle state for a coexistence number. This is null for a standard
	// Cloud API number.
	//
	// Any of "pending_onboarding", "sync_pending", "syncing", "sync_complete",
	// "active", "history_declined", "sync_deadline_expired", "offboarded",
	// "disconnected".
	CoexistenceState string    `json:"coexistence_state" api:"nullable"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	DisplayName      string    `json:"display_name"`
	Enabled          bool      `json:"enabled"`
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
	SyncProgress WhatsappPhoneNumberGetPhoneNumberResponseDataSyncProgress `json:"sync_progress" api:"nullable"`
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
func (r WhatsappPhoneNumberGetPhoneNumberResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberGetPhoneNumberResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Synchronization progress. This object is returned only while a coexistence
// number is synchronizing.
type WhatsappPhoneNumberGetPhoneNumberResponseDataSyncProgress struct {
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
func (r WhatsappPhoneNumberGetPhoneNumberResponseDataSyncProgress) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberGetPhoneNumberResponseDataSyncProgress) UnmarshalJSON(data []byte) error {
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

type WhatsappPhoneNumberGetParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappPhoneNumberGetParams]'s query parameters as
// `url.Values`.
func (r WhatsappPhoneNumberGetParams) URLQuery() (v url.Values, err error) {
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

type WhatsappPhoneNumberGetConversationWindowParams struct {
	// Destination phone number in E.164 format
	DestinationNumber string `query:"destination_number" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappPhoneNumberGetConversationWindowParams]'s query
// parameters as `url.Values`.
func (r WhatsappPhoneNumberGetConversationWindowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
