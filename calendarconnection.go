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
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Connect calendars and manage automatic meeting joins.
//
// CalendarConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCalendarConnectionService] method instead.
type CalendarConnectionService struct {
	Options []option.RequestOption
	// Connect calendars and manage automatic meeting joins.
	Meetings CalendarConnectionMeetingService
}

// NewCalendarConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCalendarConnectionService(opts ...option.RequestOption) (r CalendarConnectionService) {
	r = CalendarConnectionService{}
	r.Options = opts
	r.Meetings = NewCalendarConnectionMeetingService(opts...)
	return
}

// Start Google Calendar or Microsoft Outlook OAuth. Open oauth_url in a browser.
// No connection exists until authorization succeeds. The resulting connection is
// owned by the connecting user.
func (r *CalendarConnectionService) New(ctx context.Context, body CalendarConnectionNewParams, opts ...option.RequestOption) (res *CalendarConnectionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calendar_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a calendar connection owned by the authenticated connecting user.
// Missing and foreign connections both return 404. Signing secrets are omitted.
func (r *CalendarConnectionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CalendarConnectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calendar_connections/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update policy or defaults for future sessions. Supply at least one field.
// Existing sessions retain their configuration. Null clears nullable overrides.
// Assistant, recording, and webhook defaults apply to Google Meet and Teams.
// Assistant defaults require default_barge_in to be false or null and
// default_voice to be null. Recording=true requires a destination; false rejects
// one. A non-null default_webhook_url update returns webhook_secret once; save it
// to verify future session deliveries.
func (r *CalendarConnectionService) Update(ctx context.Context, id string, body CalendarConnectionUpdateParams, opts ...option.RequestOption) (res *CalendarConnectionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calendar_connections/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List only the authenticated connecting user’s calendar connections. Session
// resources created by those connections are shared within the organization.
func (r *CalendarConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *CalendarConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calendar_connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disconnect the calendar, cancel future joins, and clear retained calendar
// identity fields. Previously created session records, transcripts, artifacts, and
// Storage objects remain. The response is the updated connection.
func (r *CalendarConnectionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CalendarConnectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calendar_connections/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type CalendarConnection struct {
	// Calendar connection ID.
	ID string `json:"id" api:"required"`
	// all_meetings joins eligible events; external_organizer_excluded requires the
	// connected user to be the organizer; per_meeting requires an explicit opt-in.
	//
	// Any of "all_meetings", "external_organizer_excluded", "per_meeting".
	AutoJoinPolicy CalendarConnectionAutoJoinPolicy `json:"auto_join_policy" api:"required"`
	// Bot display name; null clears the override.
	BotName string `json:"bot_name" api:"required"`
	// Creation time.
	CreatedAt        time.Time               `json:"created_at" api:"required" format:"date-time"`
	DefaultAssistant MeetingSessionAssistant `json:"default_assistant" api:"required"`
	// Default manual speech interruption. Cannot be true with default_assistant.
	DefaultBargeIn bool `json:"default_barge_in" api:"required"`
	// Google Meet and Teams recording intent. When null, recording is enabled only if
	// a destination is set. True requires default_recording_export; false cannot be
	// combined with a destination.
	DefaultRecording bool `json:"default_recording" api:"required"`
	// Customer-owned Telnyx Storage destination for recording exports. Uses the
	// authenticated organization and user; do not supply credentials, region,
	// organization ID or object key.
	DefaultRecordingExport MeetingSessionRecordingDestination `json:"default_recording_export" api:"required"`
	// Generate a summary when each future session transcript finalizes.
	DefaultSummarizeOnEnd bool `json:"default_summarize_on_end" api:"required"`
	// Manual-speech voice; must be null when default_assistant is set.
	DefaultVoice string `json:"default_voice" api:"required"`
	// HTTPS destination for future session webhooks. Null clears the default. A
	// non-null update returns webhook_secret once; save it to verify signatures for
	// sessions created with this default. Existing sessions retain their original
	// signing secret.
	DefaultWebhookURL string `json:"default_webhook_url" api:"required" format:"uri"`
	// Calendar provider: Google Calendar or Microsoft Outlook.
	//
	// Any of "google_calendar", "microsoft_outlook".
	Platform CalendarConnectionPlatform `json:"platform" api:"required"`
	// Connected calendar email address.
	PlatformEmail string `json:"platform_email" api:"required"`
	// Connection status.
	//
	// Any of "connected", "disconnected", "error".
	Status CalendarConnectionStatus `json:"status" api:"required"`
	// Additional connection status detail.
	StatusDetail string `json:"status_detail" api:"required"`
	// Last update time.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AutoJoinPolicy         respjson.Field
		BotName                respjson.Field
		CreatedAt              respjson.Field
		DefaultAssistant       respjson.Field
		DefaultBargeIn         respjson.Field
		DefaultRecording       respjson.Field
		DefaultRecordingExport respjson.Field
		DefaultSummarizeOnEnd  respjson.Field
		DefaultVoice           respjson.Field
		DefaultWebhookURL      respjson.Field
		Platform               respjson.Field
		PlatformEmail          respjson.Field
		Status                 respjson.Field
		StatusDetail           respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnection) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// all_meetings joins eligible events; external_organizer_excluded requires the
// connected user to be the organizer; per_meeting requires an explicit opt-in.
type CalendarConnectionAutoJoinPolicy string

const (
	CalendarConnectionAutoJoinPolicyAllMeetings               CalendarConnectionAutoJoinPolicy = "all_meetings"
	CalendarConnectionAutoJoinPolicyExternalOrganizerExcluded CalendarConnectionAutoJoinPolicy = "external_organizer_excluded"
	CalendarConnectionAutoJoinPolicyPerMeeting                CalendarConnectionAutoJoinPolicy = "per_meeting"
)

// Calendar provider: Google Calendar or Microsoft Outlook.
type CalendarConnectionPlatform string

const (
	CalendarConnectionPlatformGoogleCalendar   CalendarConnectionPlatform = "google_calendar"
	CalendarConnectionPlatformMicrosoftOutlook CalendarConnectionPlatform = "microsoft_outlook"
)

// Connection status.
type CalendarConnectionStatus string

const (
	CalendarConnectionStatusConnected    CalendarConnectionStatus = "connected"
	CalendarConnectionStatusDisconnected CalendarConnectionStatus = "disconnected"
	CalendarConnectionStatusError        CalendarConnectionStatus = "error"
)

type CalendarConnectionResponse struct {
	Data CalendarConnection `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionAssistant struct {
	// Identifier of the assistant.
	ID string `json:"id" api:"required"`
	// Audio gating strategy in force for the assistant call leg.
	//
	// Any of "half_duplex", "full_duplex".
	AudioGate MeetingSessionAssistantAudioGate `json:"audio_gate" api:"required"`
	// Whether the bot leaves when the Assistant's conversation ends or fails.
	LeaveOnEnd bool `json:"leave_on_end" api:"required"`
	// Dynamic variables in individual-session responses; null when none were supplied.
	// Omitted from session list items.
	DynamicVariables map[string]string `json:"dynamic_variables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AudioGate        respjson.Field
		LeaveOnEnd       respjson.Field
		DynamicVariables respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionAssistant) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Audio gating strategy in force for the assistant call leg.
type MeetingSessionAssistantAudioGate string

const (
	MeetingSessionAssistantAudioGateHalfDuplex MeetingSessionAssistantAudioGate = "half_duplex"
	MeetingSessionAssistantAudioGateFullDuplex MeetingSessionAssistantAudioGate = "full_duplex"
)

type CalendarConnectionNewResponse struct {
	Data CalendarConnectionNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionNewResponseData struct {
	// Open in a browser to authorize calendar access.
	OAuthURL string `json:"oauth_url" api:"required" format:"uri"`
	// Calendar provider: Google Calendar or Microsoft Outlook.
	//
	// Any of "google_calendar", "microsoft_outlook".
	Platform string `json:"platform" api:"required"`
	// Authorization has not completed.
	Status constant.PendingOAuth `json:"status" default:"pending_oauth"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OAuthURL    respjson.Field
		Platform    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionUpdateResponse struct {
	Data CalendarConnectionUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionUpdateResponseData struct {
	// Calendar connection ID.
	ID string `json:"id" api:"required"`
	// all_meetings joins eligible events; external_organizer_excluded requires the
	// connected user to be the organizer; per_meeting requires an explicit opt-in.
	//
	// Any of "all_meetings", "external_organizer_excluded", "per_meeting".
	AutoJoinPolicy string `json:"auto_join_policy" api:"required"`
	// Bot display name; null clears the override.
	BotName string `json:"bot_name" api:"required"`
	// Creation time.
	CreatedAt        time.Time               `json:"created_at" api:"required" format:"date-time"`
	DefaultAssistant MeetingSessionAssistant `json:"default_assistant" api:"required"`
	// Default manual speech interruption. Cannot be true with default_assistant.
	DefaultBargeIn bool `json:"default_barge_in" api:"required"`
	// Google Meet and Teams recording intent. When null, recording is enabled only if
	// a destination is set. True requires default_recording_export; false cannot be
	// combined with a destination.
	DefaultRecording bool `json:"default_recording" api:"required"`
	// Customer-owned Telnyx Storage destination for recording exports. Uses the
	// authenticated organization and user; do not supply credentials, region,
	// organization ID or object key.
	DefaultRecordingExport MeetingSessionRecordingDestination `json:"default_recording_export" api:"required"`
	// Generate a summary when each future session transcript finalizes.
	DefaultSummarizeOnEnd bool `json:"default_summarize_on_end" api:"required"`
	// Manual-speech voice; must be null when default_assistant is set.
	DefaultVoice string `json:"default_voice" api:"required"`
	// HTTPS destination for future session webhooks. Null clears the default. A
	// non-null update returns webhook_secret once; save it to verify signatures for
	// sessions created with this default. Existing sessions retain their original
	// signing secret.
	DefaultWebhookURL string `json:"default_webhook_url" api:"required" format:"uri"`
	// Calendar provider: Google Calendar or Microsoft Outlook.
	//
	// Any of "google_calendar", "microsoft_outlook".
	Platform string `json:"platform" api:"required"`
	// Connected calendar email address.
	PlatformEmail string `json:"platform_email" api:"required"`
	// Connection status.
	//
	// Any of "connected", "disconnected", "error".
	Status string `json:"status" api:"required"`
	// Additional connection status detail.
	StatusDetail string `json:"status_detail" api:"required"`
	// Last update time.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Returned only when PATCH supplies a non-null default_webhook_url. Not returned
	// by list or retrieve. Save this secret to verify webhooks for future sessions.
	WebhookSecret string `json:"webhook_secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AutoJoinPolicy         respjson.Field
		BotName                respjson.Field
		CreatedAt              respjson.Field
		DefaultAssistant       respjson.Field
		DefaultBargeIn         respjson.Field
		DefaultRecording       respjson.Field
		DefaultRecordingExport respjson.Field
		DefaultSummarizeOnEnd  respjson.Field
		DefaultVoice           respjson.Field
		DefaultWebhookURL      respjson.Field
		Platform               respjson.Field
		PlatformEmail          respjson.Field
		Status                 respjson.Field
		StatusDetail           respjson.Field
		UpdatedAt              respjson.Field
		WebhookSecret          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionListResponse struct {
	Data []CalendarConnection `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionNewParams struct {
	// Calendar provider: Google Calendar or Microsoft Outlook.
	//
	// Any of "google_calendar", "microsoft_outlook".
	Platform CalendarConnectionNewParamsPlatform `json:"platform,omitzero" api:"required"`
	paramObj
}

func (r CalendarConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CalendarConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CalendarConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Calendar provider: Google Calendar or Microsoft Outlook.
type CalendarConnectionNewParamsPlatform string

const (
	CalendarConnectionNewParamsPlatformGoogleCalendar   CalendarConnectionNewParamsPlatform = "google_calendar"
	CalendarConnectionNewParamsPlatformMicrosoftOutlook CalendarConnectionNewParamsPlatform = "microsoft_outlook"
)

type CalendarConnectionUpdateParams struct {
	// Bot display name; null clears the override.
	BotName param.Opt[string] `json:"bot_name,omitzero"`
	// Default manual speech interruption. Cannot be true with default_assistant.
	DefaultBargeIn param.Opt[bool] `json:"default_barge_in,omitzero"`
	// Google Meet and Teams recording intent. When null, recording is enabled only if
	// a destination is set. True requires default_recording_export; false cannot be
	// combined with a destination.
	DefaultRecording param.Opt[bool] `json:"default_recording,omitzero"`
	// Generate a summary when each future session transcript finalizes.
	DefaultSummarizeOnEnd param.Opt[bool] `json:"default_summarize_on_end,omitzero"`
	// Manual-speech voice; must be null when default_assistant is set.
	DefaultVoice param.Opt[string] `json:"default_voice,omitzero"`
	// HTTPS destination for future session webhooks. Null clears the default. A
	// non-null update returns webhook_secret once; save it to verify signatures for
	// sessions created with this default. Existing sessions retain their original
	// signing secret.
	DefaultWebhookURL param.Opt[string] `json:"default_webhook_url,omitzero" format:"uri"`
	// Default assistant for future calendar-scheduled Google Meet and Teams sessions.
	// The saved configuration must fit within 8 KiB. Use id, audio_gate,
	// dynamic_variables, and leave_on_end; avatar is not accepted. audio_gate defaults
	// to full_duplex.
	DefaultAssistant CalendarConnectionUpdateParamsDefaultAssistant `json:"default_assistant,omitzero"`
	// all_meetings joins eligible events; external_organizer_excluded requires the
	// connected user to be the organizer; per_meeting requires an explicit opt-in.
	//
	// Any of "all_meetings", "external_organizer_excluded", "per_meeting".
	AutoJoinPolicy CalendarConnectionUpdateParamsAutoJoinPolicy `json:"auto_join_policy,omitzero"`
	// Customer-owned Telnyx Storage destination for recording exports. Uses the
	// authenticated organization and user; do not supply credentials, region,
	// organization ID or object key.
	DefaultRecordingExport MeetingSessionRecordingDestinationParam `json:"default_recording_export,omitzero"`
	paramObj
}

func (r CalendarConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CalendarConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CalendarConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// all_meetings joins eligible events; external_organizer_excluded requires the
// connected user to be the organizer; per_meeting requires an explicit opt-in.
type CalendarConnectionUpdateParamsAutoJoinPolicy string

const (
	CalendarConnectionUpdateParamsAutoJoinPolicyAllMeetings               CalendarConnectionUpdateParamsAutoJoinPolicy = "all_meetings"
	CalendarConnectionUpdateParamsAutoJoinPolicyExternalOrganizerExcluded CalendarConnectionUpdateParamsAutoJoinPolicy = "external_organizer_excluded"
	CalendarConnectionUpdateParamsAutoJoinPolicyPerMeeting                CalendarConnectionUpdateParamsAutoJoinPolicy = "per_meeting"
)

// Default assistant for future calendar-scheduled Google Meet and Teams sessions.
// The saved configuration must fit within 8 KiB. Use id, audio_gate,
// dynamic_variables, and leave_on_end; avatar is not accepted. audio_gate defaults
// to full_duplex.
//
// The property ID is required.
type CalendarConnectionUpdateParamsDefaultAssistant struct {
	// Identifier of the assistant to attach.
	ID string `json:"id" api:"required"`
	// Leave when the assistant conversation ends or fails. Defaults to false. Other
	// leave-policy rules and explicit leave requests still apply.
	LeaveOnEnd param.Opt[bool] `json:"leave_on_end,omitzero"`
	// Use full_duplex (default) to allow participants to interrupt, or half_duplex to
	// pause listening while the assistant speaks. Full-duplex pricing depends on
	// participant audio usage; see the Meeting pricing guide.
	//
	// Any of "half_duplex", "full_duplex".
	AudioGate string `json:"audio_gate,omitzero"`
	// Per-conversation values for the
	// [dynamic variables](/docs/inference/ai-assistants/dynamic-variables) used in the
	// Assistant's instructions, greeting, or tools. Delivered before the Assistant's
	// first utterance, so they resolve for the opening line as well as the rest of the
	// conversation. At most 63 entries; keys 1-128 characters; values must be strings.
	// The map is budgeted in aggregate at 1,047,552 bytes (1023 KiB) rather than
	// capped per value. `streaming_audio`, `ai_assistant_streaming_audio` and
	// `meeting_session_id` are reserved and rejected with `400 invalid_request`.
	DynamicVariables map[string]string `json:"dynamic_variables,omitzero"`
	paramObj
}

func (r CalendarConnectionUpdateParamsDefaultAssistant) MarshalJSON() (data []byte, err error) {
	type shadow CalendarConnectionUpdateParamsDefaultAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CalendarConnectionUpdateParamsDefaultAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CalendarConnectionUpdateParamsDefaultAssistant](
		"audio_gate", "half_duplex", "full_duplex",
	)
}
