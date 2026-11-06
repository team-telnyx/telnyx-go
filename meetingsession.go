// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
	"github.com/tidwall/gjson"
)

// MeetingSessionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingSessionService] method instead.
type MeetingSessionService struct {
	Options []option.RequestOption
	// Send real-time speech and chat actions to an active meeting session.
	Actions MeetingSessionActionService
	// Create and retrieve asynchronous summaries and action-item artifacts.
	Artifacts MeetingSessionArtifactService
}

// NewMeetingSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeetingSessionService(opts ...option.RequestOption) (r MeetingSessionService) {
	r = MeetingSessionService{}
	r.Options = opts
	r.Actions = NewMeetingSessionActionService(opts...)
	r.Artifacts = NewMeetingSessionArtifactService(opts...)
	return
}

// Creates a meeting session and returns 201, or 200 for a replay. Use a unique
// `idempotency_key` per intended join and reuse the same key and request on
// retries. Reusing a key with a different join configuration or meeting URL can
// return 409 invalid_state. Always retry with the original request. The session
// can pass through joining and waiting_for_admission before active. Optional
// `camera_image` input is write-only and applies when no assistant avatar takes
// precedence. An ignored URL is not fetched. An effective URL source is resolved
// before bot creation; the source URL and image bytes are not returned in the
// session response. Treat signed URLs as credentials. With
// `assistant.avatar: true`, the assistant is read before the bot is created: an
// assistant without a Telnyx avatar configured is rejected with 400
// `invalid_request` (message prefix `assistant_avatar_not_configured:`), an
// unknown assistant with 400 `invalid_request` (prefix `assistant_not_found:`),
// and a failed lookup with 503 `assistant_lookup_unavailable`; no bot is created
// in any of these cases.
func (r *MeetingSessionService) New(ctx context.Context, body MeetingSessionNewParams, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "meeting_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single meeting session by ID. A session that does not exist or that
// belongs to a different account both return 404.
func (r *MeetingSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable properties of a meeting session. Only sessions in the scheduled
// state can be updated; any other state returns 409 with the invalid_state error
// code. All request fields are optional, and an empty object is a valid no-op
// update. Calendar-controlled Google Meet and Teams sessions return 409
// calendar_managed; change the event instead. A conflicting pending edit returns
// 409 edit_in_progress. On a timeout or 503, read the session and retry the same
// changes.
func (r *MeetingSessionService) Update(ctx context.Context, id string, body MeetingSessionUpdateParams, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of meeting sessions, optionally filtered by status. List items
// omit assistant.dynamic_variables and recording_export. Some items include
// remote_state_observed_at; retrieve the individual session for current state.
func (r *MeetingSessionService) List(ctx context.Context, query MeetingSessionListParams, opts ...option.RequestOption) (res *MeetingSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "meeting_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stops a meeting session without deleting its persisted record. Scheduled bots
// are cancelled, while bots that are joining or active are asked to leave. The
// persisted meeting session record remains available. A successful response can
// report leaving while departure is pending. Poll until ended, failed, or
// admission_denied. Exported Storage objects are unaffected.
func (r *MeetingSessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Irreversibly requests deletion of the aggregate recording media for the session.
// The operation retains the Telnyx-local Meeting session, transcript segments,
// events, artifacts, and usage records. It is separate from
// `DELETE /meeting_sessions/{id}`, which stops or cancels participation without
// deleting the persisted session. A missing/foreign session returns 404; a
// deletion request whose outcome is unknown returns 502. Customer Storage exports
// do not support this operation (422 unsupported_capability); delete the exported
// object through Storage instead. Stopping a session does not delete an exported
// object.
func (r *MeetingSessionService) DeleteRecordingMedia(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionDeleteRecordingMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/recording_media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns stored events ordered by ascending `seq`. To continue, pass the last
// returned item's `seq` as `after`. An empty page means no later stored events
// existed at read time; this operation returns no separate next-page cursor.
// Default `limit` is 100 and maximum is 1,000.
func (r *MeetingSessionService) GetEvents(ctx context.Context, id string, query MeetingSessionGetEventsParams, opts ...option.RequestOption) (res *MeetingSessionGetEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/events", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve recording export status. Customer Storage exports include
// recording_export with the bucket, region, key and pending/exported/failed
// status, or null when no reference is available. The data array is empty for
// these exports. Download or share the object using the Storage API once its
// status is exported. Zoom recordings and existing Google Meet or Teams recordings
// held by the Meeting service continue returning temporary download links in data
// through this same endpoint and session ID. Those recordings do not require
// Storage setup or migration. Refresh this endpoint when a link expires. Select
// the download method from the returned fields, not the platform alone.
func (r *MeetingSessionService) GetRecordings(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionGetRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/recordings", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns transcript segments ordered by ascending seq. Default limit is 100,
// maximum 1000. Continue with after=meta.next_after. On an empty page,
// meta.next_after preserves a positive supplied after cursor, or is null when no
// positive cursor was supplied. wait_seconds holds an empty read for up to 25
// seconds; a timeout still returns 200.
func (r *MeetingSessionService) GetTranscript(ctx context.Context, id string, query MeetingSessionGetTranscriptParams, opts ...option.RequestOption) (res *MeetingSessionGetTranscriptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/transcript", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a meeting session. Core serializer fields are required;
// platform-specific fields are optional. Nullable fields use null when absent. No
// actor, provider-bot, idempotency, routing, key, or internal fields are exposed.
type MeetingSession struct {
	// Unique identifier for the meeting session.
	ID string `json:"id" api:"required"`
	// Identifier of the owning account.
	AccountID string `json:"account_id" api:"required"`
	// Assistant configuration if an assistant is attached, otherwise null.
	Assistant MeetingSessionAssistant `json:"assistant" api:"required"`
	// Current state of the assistant, or null if no assistant is attached.
	//
	// Any of "starting", "connected", "failed", "ended".
	AssistantState MeetingSessionAssistantState `json:"assistant_state" api:"required"`
	// Timestamp of the last assistant state change, or null.
	AssistantStateChangedAt time.Time `json:"assistant_state_changed_at" api:"required" format:"date-time"`
	// The assistant's avatar, captured from its `avatar_settings` when the session was
	// created. For display only; later changes to the assistant do not update it.
	Avatar MeetingSessionAvatar `json:"avatar" api:"required"`
	// Current state of the avatar, or null if the session has no avatar. `starting`
	// until the avatar's video is playing, then `connected`; `degraded` when the
	// session has fallen back to voice-only (the `avatar.degraded` event carries the
	// reason); `disconnected` when the avatar has ended.
	//
	// Any of "starting", "connected", "degraded", "disconnected".
	AvatarState MeetingSessionAvatarState `json:"avatar_state" api:"required"`
	// Timestamp of the last avatar state change, or null.
	AvatarStateChangedAt time.Time `json:"avatar_state_changed_at" api:"required" format:"date-time"`
	// Display name of the bot in the meeting.
	BotName string               `json:"bot_name" api:"required"`
	Config  MeetingSessionConfig `json:"config" api:"required"`
	// Timestamp when the session was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Timestamp when the session ended, or null if ongoing.
	EndedAt time.Time `json:"ended_at" api:"required" format:"date-time"`
	// A machine-readable lifecycle reason, or null. Values include left_on_request,
	// meeting_ended, everyone_left, no_one_joined, silence_timeout,
	// max_duration_reached, admission_timeout, admission_denied, removed_from_meeting,
	// recording_permission_denied, meeting_not_found, meeting_not_started,
	// meeting_access_denied, meeting_full, cancelled, cancelled_calendar_disconnected,
	// cancelled_calendar_event_deleted, cancelled_calendar_superseded,
	// cancelled_policy, internal_error, and other. Treat this vocabulary as
	// extensible.
	FailureReason string `json:"failure_reason" api:"required"`
	// Scheduled join time, or null for immediate join.
	JoinAt time.Time `json:"join_at" api:"required" format:"date-time"`
	// Timestamp when the session first became `active`, or null if it never became
	// active. This remains positive admission evidence after terminal transitions.
	JoinedAt time.Time `json:"joined_at" api:"required" format:"date-time"`
	// The meeting URL the bot joins.
	MeetingURL string `json:"meeting_url" api:"required" format:"uri"`
	// Arbitrary key-value metadata attached to the session.
	Metadata map[string]any `json:"metadata" api:"required"`
	// Detected meeting platform.
	//
	// Any of "zoom", "google_meet", "teams", "webex", "unknown".
	Platform MeetingSessionPlatform `json:"platform" api:"required"`
	// Always telnyx.
	Provider constant.Telnyx `json:"provider" default:"telnyx"`
	// Whether the session is being recorded.
	Recording bool `json:"recording" api:"required"`
	// Lifecycle status. `waiting_for_admission` means the bot reached the meeting
	// lobby and may require host approval. `active` means the bot entered the
	// meeting/media path. `ended` alone does not prove attendance; use non-null
	// `joined_at` as positive evidence that the session became active.
	// `admission_denied` is reserved for an explicit provider denial, while
	// cancellation or another termination can end a never-admitted session as `ended`.
	//
	// Any of "scheduled", "joining", "waiting_for_admission", "active", "leaving",
	// "ended", "failed", "admission_denied".
	Status MeetingSessionStatus `json:"status" api:"required"`
	// A machine-readable lifecycle reason, or null. Values include left_on_request,
	// meeting_ended, everyone_left, no_one_joined, silence_timeout,
	// max_duration_reached, admission_timeout, admission_denied, removed_from_meeting,
	// recording_permission_denied, meeting_not_found, meeting_not_started,
	// meeting_access_denied, meeting_full, cancelled, cancelled_calendar_disconnected,
	// cancelled_calendar_event_deleted, cancelled_calendar_superseded,
	// cancelled_policy, internal_error, and other. Treat this vocabulary as
	// extensible.
	StatusDetail string `json:"status_detail" api:"required"`
	// Timestamp of the last update to the session.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Webhook endpoint for session lifecycle callbacks, or null if not configured.
	WebhookURL string `json:"webhook_url" api:"required" format:"uri"`
	// Customer Storage recording reference. Pending is not downloadable; exported
	// means upload was verified; failed is terminal. Download with ordinary Storage
	// authorization. No presigned URL is returned.
	RecordingExport MeetingSessionRecordingExport `json:"recording_export" api:"nullable"`
	// Optional observation time on session list items. Retrieve the individual session
	// when current state is required.
	RemoteStateObservedAt time.Time `json:"remote_state_observed_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		Assistant               respjson.Field
		AssistantState          respjson.Field
		AssistantStateChangedAt respjson.Field
		Avatar                  respjson.Field
		AvatarState             respjson.Field
		AvatarStateChangedAt    respjson.Field
		BotName                 respjson.Field
		Config                  respjson.Field
		CreatedAt               respjson.Field
		EndedAt                 respjson.Field
		FailureReason           respjson.Field
		JoinAt                  respjson.Field
		JoinedAt                respjson.Field
		MeetingURL              respjson.Field
		Metadata                respjson.Field
		Platform                respjson.Field
		Provider                respjson.Field
		Recording               respjson.Field
		Status                  respjson.Field
		StatusDetail            respjson.Field
		UpdatedAt               respjson.Field
		WebhookURL              respjson.Field
		RecordingExport         respjson.Field
		RemoteStateObservedAt   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSession) RawJSON() string { return r.JSON.raw }
func (r *MeetingSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the assistant, or null if no assistant is attached.
type MeetingSessionAssistantState string

const (
	MeetingSessionAssistantStateStarting  MeetingSessionAssistantState = "starting"
	MeetingSessionAssistantStateConnected MeetingSessionAssistantState = "connected"
	MeetingSessionAssistantStateFailed    MeetingSessionAssistantState = "failed"
	MeetingSessionAssistantStateEnded     MeetingSessionAssistantState = "ended"
)

// The assistant's avatar, captured from its `avatar_settings` when the session was
// created. For display only; later changes to the assistant do not update it.
type MeetingSessionAvatar struct {
	// Identifier of the avatar.
	AvatarID string `json:"avatar_id" api:"required"`
	// Identifier of the avatar's background.
	BackgroundID string `json:"background_id" api:"required"`
	// Avatar rendering model, for example `foxtrot-1`.
	Model string `json:"model" api:"required"`
	// Avatar provider identifier.
	Provider constant.Telnyx `json:"provider" default:"telnyx"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarID     respjson.Field
		BackgroundID respjson.Field
		Model        respjson.Field
		Provider     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionAvatar) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the avatar, or null if the session has no avatar. `starting`
// until the avatar's video is playing, then `connected`; `degraded` when the
// session has fallen back to voice-only (the `avatar.degraded` event carries the
// reason); `disconnected` when the avatar has ended.
type MeetingSessionAvatarState string

const (
	MeetingSessionAvatarStateStarting     MeetingSessionAvatarState = "starting"
	MeetingSessionAvatarStateConnected    MeetingSessionAvatarState = "connected"
	MeetingSessionAvatarStateDegraded     MeetingSessionAvatarState = "degraded"
	MeetingSessionAvatarStateDisconnected MeetingSessionAvatarState = "disconnected"
)

type MeetingSessionConfig struct {
	// When enabled, a human participant `speech_on` event interrupts and stops the
	// current bot audio; it does not bypass admission or initiate speech. Assistant
	// sessions reject `barge_in: true`.
	BargeIn bool `json:"barge_in" api:"required"`
	// The message posted to chat on join, or null when unset.
	ChatOnEnter string `json:"chat_on_enter" api:"required"`
	// Detection settings. Enabling audio or video detection is currently unsupported.
	DeepfakeDetection MeetingSessionConfigDeepfakeDetection `json:"deepfake_detection" api:"required"`
	// Text spoken on meeting entry, or null if not set.
	SpeakOnEnter string `json:"speak_on_enter" api:"required"`
	// Whether a summary artifact is generated on session end.
	SummarizeOnEnd bool `json:"summarize_on_end" api:"required"`
	// Configured voice identifier, or null if not set.
	Voice string `json:"voice" api:"required"`
	// Automatic departure settings for Google Meet and Teams. Omitted fields use the
	// default behavior. Set at creation; cannot be changed later. Zoom ignores this
	// policy.
	LeavePolicy MeetingSessionLeavePolicy `json:"leave_policy"`
	// Recording intent set at creation for Google Meet and Teams. Distinct from the
	// top-level recording field, which reports current activity.
	Recording bool `json:"recording"`
	// Non-default recording layout. Omitted for the default speaker_view layout.
	//
	// Any of "gallery_view", "speaker_view_no_sidebar".
	RecordingView string `json:"recording_view"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BargeIn           respjson.Field
		ChatOnEnter       respjson.Field
		DeepfakeDetection respjson.Field
		SpeakOnEnter      respjson.Field
		SummarizeOnEnd    respjson.Field
		Voice             respjson.Field
		LeavePolicy       respjson.Field
		Recording         respjson.Field
		RecordingView     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionConfig) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detection settings. Enabling audio or video detection is currently unsupported.
type MeetingSessionConfigDeepfakeDetection struct {
	Audio bool `json:"audio" api:"required"`
	Video bool `json:"video" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio       respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionConfigDeepfakeDetection) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionConfigDeepfakeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detected meeting platform.
type MeetingSessionPlatform string

const (
	MeetingSessionPlatformZoom       MeetingSessionPlatform = "zoom"
	MeetingSessionPlatformGoogleMeet MeetingSessionPlatform = "google_meet"
	MeetingSessionPlatformTeams      MeetingSessionPlatform = "teams"
	MeetingSessionPlatformWebex      MeetingSessionPlatform = "webex"
	MeetingSessionPlatformUnknown    MeetingSessionPlatform = "unknown"
)

// Lifecycle status. `waiting_for_admission` means the bot reached the meeting
// lobby and may require host approval. `active` means the bot entered the
// meeting/media path. `ended` alone does not prove attendance; use non-null
// `joined_at` as positive evidence that the session became active.
// `admission_denied` is reserved for an explicit provider denial, while
// cancellation or another termination can end a never-admitted session as `ended`.
type MeetingSessionStatus string

const (
	MeetingSessionStatusScheduled           MeetingSessionStatus = "scheduled"
	MeetingSessionStatusJoining             MeetingSessionStatus = "joining"
	MeetingSessionStatusWaitingForAdmission MeetingSessionStatus = "waiting_for_admission"
	MeetingSessionStatusActive              MeetingSessionStatus = "active"
	MeetingSessionStatusLeaving             MeetingSessionStatus = "leaving"
	MeetingSessionStatusEnded               MeetingSessionStatus = "ended"
	MeetingSessionStatusFailed              MeetingSessionStatus = "failed"
	MeetingSessionStatusAdmissionDenied     MeetingSessionStatus = "admission_denied"
)

// Automatic departure settings for Google Meet and Teams. Omitted fields use the
// default behavior. Set at creation; cannot be changed later. Zoom ignores this
// policy.
type MeetingSessionLeavePolicy struct {
	// Seconds alone after another participant has attended and left. Does not start
	// just because the bot arrives first.
	AloneSeconds int64 `json:"alone_seconds"`
	// Case-insensitive names or phrases identifying other bots to exclude when
	// deciding whether this bot is alone. Matching splits names on spaces, hyphens and
	// underscores. Null or an empty array disables filtering.
	BotNameKeywords []string `json:"bot_name_keywords" api:"nullable"`
	// Seconds to wait for the host to start the meeting.
	HostWaitSeconds int64 `json:"host_wait_seconds"`
	// Seconds to wait for lobby admission.
	LobbyWaitSeconds int64 `json:"lobby_wait_seconds"`
	// Maximum seconds after joining. Null means no per-join duration limit; service
	// safety limits still apply.
	MaxDurationSeconds int64 `json:"max_duration_seconds" api:"nullable"`
	// Seconds after joining before silence detection starts a fresh inactivity timer.
	SilenceGraceSeconds int64 `json:"silence_grace_seconds"`
	// Continuous audio inactivity before departure after the silence grace period.
	SilenceSeconds int64 `json:"silence_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AloneSeconds        respjson.Field
		BotNameKeywords     respjson.Field
		HostWaitSeconds     respjson.Field
		LobbyWaitSeconds    respjson.Field
		MaxDurationSeconds  respjson.Field
		SilenceGraceSeconds respjson.Field
		SilenceSeconds      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionLeavePolicy) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionLeavePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MeetingSessionLeavePolicy to a
// MeetingSessionLeavePolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MeetingSessionLeavePolicyParam.Overrides()
func (r MeetingSessionLeavePolicy) ToParam() MeetingSessionLeavePolicyParam {
	return param.Override[MeetingSessionLeavePolicyParam](json.RawMessage(r.RawJSON()))
}

// Automatic departure settings for Google Meet and Teams. Omitted fields use the
// default behavior. Set at creation; cannot be changed later. Zoom ignores this
// policy.
type MeetingSessionLeavePolicyParam struct {
	// Maximum seconds after joining. Null means no per-join duration limit; service
	// safety limits still apply.
	MaxDurationSeconds param.Opt[int64] `json:"max_duration_seconds,omitzero"`
	// Seconds alone after another participant has attended and left. Does not start
	// just because the bot arrives first.
	AloneSeconds param.Opt[int64] `json:"alone_seconds,omitzero"`
	// Seconds to wait for the host to start the meeting.
	HostWaitSeconds param.Opt[int64] `json:"host_wait_seconds,omitzero"`
	// Seconds to wait for lobby admission.
	LobbyWaitSeconds param.Opt[int64] `json:"lobby_wait_seconds,omitzero"`
	// Seconds after joining before silence detection starts a fresh inactivity timer.
	SilenceGraceSeconds param.Opt[int64] `json:"silence_grace_seconds,omitzero"`
	// Continuous audio inactivity before departure after the silence grace period.
	SilenceSeconds param.Opt[int64] `json:"silence_seconds,omitzero"`
	// Case-insensitive names or phrases identifying other bots to exclude when
	// deciding whether this bot is alone. Matching splits names on spaces, hyphens and
	// underscores. Null or an empty array disables filtering.
	BotNameKeywords []string `json:"bot_name_keywords,omitzero"`
	paramObj
}

func (r MeetingSessionLeavePolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionLeavePolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionLeavePolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer-owned Telnyx Storage destination for recording exports. Uses the
// authenticated organization and user; do not supply credentials, region,
// organization ID or object key.
type MeetingSessionRecordingDestination struct {
	// Existing accessible bucket. Consecutive dots and IPv4-address names are
	// rejected.
	Bucket string `json:"bucket" api:"required"`
	// Optional object-key prefix. One trailing slash is normalized. Empty path
	// segments, . and .. segments are rejected.
	Prefix string `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Prefix      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionRecordingDestination) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionRecordingDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MeetingSessionRecordingDestination to a
// MeetingSessionRecordingDestinationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MeetingSessionRecordingDestinationParam.Overrides()
func (r MeetingSessionRecordingDestination) ToParam() MeetingSessionRecordingDestinationParam {
	return param.Override[MeetingSessionRecordingDestinationParam](json.RawMessage(r.RawJSON()))
}

// Customer-owned Telnyx Storage destination for recording exports. Uses the
// authenticated organization and user; do not supply credentials, region,
// organization ID or object key.
//
// The property Bucket is required.
type MeetingSessionRecordingDestinationParam struct {
	// Existing accessible bucket. Consecutive dots and IPv4-address names are
	// rejected.
	Bucket string `json:"bucket" api:"required"`
	// Optional object-key prefix. One trailing slash is normalized. Empty path
	// segments, . and .. segments are rejected.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	paramObj
}

func (r MeetingSessionRecordingDestinationParam) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionRecordingDestinationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionRecordingDestinationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer Storage recording reference. Pending is not downloadable; exported
// means upload was verified; failed is terminal. Download with ordinary Storage
// authorization. No presigned URL is returned.
type MeetingSessionRecordingExport struct {
	Bucket     string    `json:"bucket" api:"required"`
	Bytes      int64     `json:"bytes" api:"required"`
	ExportedAt time.Time `json:"exported_at" api:"required" format:"date-time"`
	Key        string    `json:"key" api:"required"`
	Region     string    `json:"region" api:"required"`
	// Any of "pending", "exported", "failed".
	Status MeetingSessionRecordingExportStatus `json:"status" api:"required"`
	Type   constant.Video                      `json:"type" default:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Bytes       respjson.Field
		ExportedAt  respjson.Field
		Key         respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionRecordingExport) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionRecordingExport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionRecordingExportStatus string

const (
	MeetingSessionRecordingExportStatusPending  MeetingSessionRecordingExportStatus = "pending"
	MeetingSessionRecordingExportStatusExported MeetingSessionRecordingExportStatus = "exported"
	MeetingSessionRecordingExportStatusFailed   MeetingSessionRecordingExportStatus = "failed"
)

type MeetingSessionResponse struct {
	// Represents a meeting session. Core serializer fields are required;
	// platform-specific fields are optional. Nullable fields use null when absent. No
	// actor, provider-bot, idempotency, routing, key, or internal fields are exposed.
	Data MeetingSession `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionListResponse struct {
	Data []MeetingSession `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionDeleteRecordingMediaResponse struct {
	Data MeetingSessionDeleteRecordingMediaResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionDeleteRecordingMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionDeleteRecordingMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionDeleteRecordingMediaResponseData struct {
	// Any of "requested", "already_in_progress".
	DeletionStatus string `json:"deletion_status" api:"required"`
	// The account-scoped Meeting Session identifier.
	MeetingSessionID string                          `json:"meeting_session_id" api:"required"`
	Provider         constant.Telnyx                 `json:"provider" default:"telnyx"`
	Scope            constant.ProviderRecordingMedia `json:"scope" default:"provider_recording_media"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletionStatus   respjson.Field
		MeetingSessionID respjson.Field
		Provider         respjson.Field
		Scope            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionDeleteRecordingMediaResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionDeleteRecordingMediaResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetEventsResponse struct {
	Data []MeetingSessionGetEventsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetEventsResponseData struct {
	OccurredAt time.Time      `json:"occurred_at" api:"required" format:"date-time"`
	Payload    map[string]any `json:"payload" api:"required"`
	Seq        int64          `json:"seq" api:"required"`
	Type       string         `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OccurredAt  respjson.Field
		Payload     respjson.Field
		Seq         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetEventsResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetEventsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetRecordingsResponse struct {
	// Temporary download URLs for recordings held by the Meeting service, including
	// Zoom and existing Google Meet and Teams recordings. Empty for customer Storage
	// exports; use recording_export instead. Choose the download method from these
	// fields rather than the platform alone.
	Data []MeetingSessionGetRecordingsResponseData `json:"data" api:"required"`
	// Customer Storage recording reference. Pending is not downloadable; exported
	// means upload was verified; failed is terminal. Download with ordinary Storage
	// authorization. No presigned URL is returned.
	RecordingExport MeetingSessionRecordingExport `json:"recording_export" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		RecordingExport respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetRecordingsResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetRecordingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetRecordingsResponseData struct {
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	Type      string    `json:"type" api:"required"`
	URL       string    `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetRecordingsResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetRecordingsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetTranscriptResponse struct {
	Data []MeetingSessionGetTranscriptResponseData `json:"data" api:"required"`
	Meta MeetingSessionGetTranscriptResponseMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetTranscriptResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetTranscriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetTranscriptResponseData struct {
	Confidence float64   `json:"confidence" api:"required"`
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	RelativeTs float64   `json:"relative_ts" api:"required"`
	// Monotonic transcript append cursor. This is separate from the event-history
	// cursor and does not establish conversational turn order.
	Seq int64 `json:"seq" api:"required"`
	// Recorded participant or bot label, or null when unavailable. Labels are
	// attribution, not verified identity. Segments from concurrent speakers can
	// finalize in a different order from the conversation.
	SpeakerLabel string `json:"speaker_label" api:"required"`
	Text         string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence   respjson.Field
		OccurredAt   respjson.Field
		RelativeTs   respjson.Field
		Seq          respjson.Field
		SpeakerLabel respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetTranscriptResponseData) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetTranscriptResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetTranscriptResponseMeta struct {
	// Cursor to pass as `after` on the next request, or null when the response
	// contains no segments.
	NextAfter int64 `json:"next_after" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextAfter   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetTranscriptResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetTranscriptResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionNewParams struct {
	// The meeting URL the bot should join.
	MeetingURL string `json:"meeting_url" api:"required" format:"uri"`
	// When enabled, a human participant `speech_on` event interrupts and stops the
	// current bot audio; it does not bypass admission or initiate speech. Assistant
	// sessions reject `barge_in: true`.
	BargeIn param.Opt[bool] `json:"barge_in,omitzero"`
	// Display name for the bot in the meeting. Defaults to "Meeting Bot".
	BotName param.Opt[string] `json:"bot_name,omitzero"`
	// A message the bot posts to the meeting's chat as soon as it becomes active —
	// typically a recording disclosure. Delivered at most once. Independent of
	// `speak_on_enter`: both may be set, and the chat message posts first because it
	// does not wait for text-to-speech or avatar startup. Rejected with 422
	// `unsupported_capability` on platforms without meeting chat. For Google Meet and
	// Teams, text must be nonblank, at most 4096 UTF-8 bytes, with no control
	// characters except tabs and line breaks. Counts toward the 1000-chat-request
	// session limit.
	ChatOnEnter param.Opt[string] `json:"chat_on_enter,omitzero"`
	// Use a unique key per intended join and reuse the same key and request on
	// retries. Reusing a key with a different join configuration or meeting URL can
	// return 409 invalid_state. Always retry with the original request.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// Future ISO-8601 join timestamp including timezone. Google Meet and Teams accept
	// dates at most three years ahead. Omit for an immediate join. Assistant sessions
	// created through this endpoint must join immediately.
	JoinAt param.Opt[time.Time] `json:"join_at,omitzero" format:"date-time"`
	// Google Meet and Teams only. Defaults to true when recording_export is supplied
	// and false otherwise. True requires recording_export.bucket; false cannot be
	// combined with recording_export. Fixed at creation, with no pause/resume action.
	// Zoom rejects recording, recording_view, and recording_export with 422
	// unsupported_capability; omit these fields for Zoom.
	Recording param.Opt[bool] `json:"recording,omitzero"`
	// Text the bot speaks after joining. Available for sessions without a digital
	// avatar. With an AI Assistant, the greeting waits until the assistant is
	// connected. Rejected with 400 invalid_request when assistant.avatar is true.
	SpeakOnEnter param.Opt[string] `json:"speak_on_enter,omitzero"`
	// If true, generate a summary artifact when the session ends.
	SummarizeOnEnd param.Opt[bool] `json:"summarize_on_end,omitzero"`
	// Session-default voice identifier used for `speak_on_enter` and ordinary speak
	// actions. A voice supplied on an individual speak action overrides this default
	// for that utterance. With an assistant attached, its conversational voice remains
	// configured on the assistant.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// HTTPS endpoint to receive session lifecycle callbacks. Static validation
	// requires HTTPS, rejects embedded credentials and blocked hosts, and enforces
	// egress policy. Validation makes no network request to the endpoint.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"uri"`
	// Attach a Telnyx AI Assistant owned by the authenticated account. Supply its ID
	// and optional conversation settings. Unknown fields are rejected.
	Assistant MeetingSessionNewParamsAssistant `json:"assistant,omitzero"`
	// Not supported for Google Meet and Teams joins (422 unsupported_capability).
	// Write-only static camera-tile image for this session, not an account or
	// participant profile photo. Supply exactly one JPEG source. When effective, the
	// image is used as the bot's static camera/video output; presentation varies by
	// meeting platform and recording configuration and is not guaranteed in
	// recordings. An attached assistant or avatar takes precedence, so this input is
	// ignored and a URL source is not fetched.
	CameraImage MeetingSessionNewParamsCameraImageUnion `json:"camera_image,omitzero"`
	// Automatic departure settings for Google Meet and Teams. Omitted fields use the
	// default behavior. Set at creation; cannot be changed later. Zoom ignores this
	// policy.
	LeavePolicy MeetingSessionLeavePolicyParam `json:"leave_policy,omitzero"`
	// Arbitrary key-value metadata attached to the session. The serialized JSON
	// representation must not exceed 16384 characters at runtime.
	Metadata map[string]any `json:"metadata,omitzero"`
	// Customer-owned Telnyx Storage destination for recording exports. Uses the
	// authenticated organization and user; do not supply credentials, region,
	// organization ID or object key.
	RecordingExport MeetingSessionRecordingDestinationParam `json:"recording_export,omitzero"`
	// Recording layout for Google Meet and Teams. Defaults to speaker_view. Accepted
	// at creation only; does not enable recording by itself.
	//
	// Any of "speaker_view", "gallery_view", "speaker_view_no_sidebar".
	RecordingView MeetingSessionNewParamsRecordingView `json:"recording_view,omitzero"`
	paramObj
}

func (r MeetingSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attach a Telnyx AI Assistant owned by the authenticated account. Supply its ID
// and optional conversation settings. Unknown fields are rejected.
//
// The property ID is required.
type MeetingSessionNewParamsAssistant struct {
	// Identifier of the assistant to attach.
	ID string `json:"id" api:"required"`
	// Beta. Show the avatar configured on the assistant (its `avatar_settings`; see
	// [AI Assistant avatars](/docs/inference/ai-assistants/avatars)) as the bot's
	// camera, lip-synced to the assistant's own speech. The request carries no avatar
	// ID or key: the assistant is read at creation and the request is rejected with
	// 400 `invalid_request` if it has no Telnyx avatar configured. Available for
	// Google Meet, Zoom, and Teams sessions. REST-only. Not echoed back on
	// `assistant`; the session's top-level `avatar` and `avatar_state` fields report
	// it instead. If the avatar fails, the session continues voice-only and
	// `avatar_state` becomes `degraded`.
	Avatar param.Opt[bool] `json:"avatar,omitzero"`
	// Leave when the assistant conversation ends or fails. Defaults to false. Other
	// leave-policy rules and explicit leave requests still apply.
	LeaveOnEnd param.Opt[bool] `json:"leave_on_end,omitzero"`
	// Use half_duplex (default) to pause listening while the assistant speaks, or
	// full_duplex to allow participants to interrupt. Full-duplex pricing depends on
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

func (r MeetingSessionNewParamsAssistant) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParamsAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParamsAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MeetingSessionNewParamsAssistant](
		"audio_gate", "half_duplex", "full_duplex",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MeetingSessionNewParamsCameraImageUnion struct {
	OfMeetingSessionCameraImageBase64Source *MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source `json:",omitzero,inline"`
	OfMeetingSessionCameraImageURLSource    *MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource    `json:",omitzero,inline"`
	paramUnion
}

func (u MeetingSessionNewParamsCameraImageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMeetingSessionCameraImageBase64Source, u.OfMeetingSessionCameraImageURLSource)
}
func (u *MeetingSessionNewParamsCameraImageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MeetingSessionNewParamsCameraImageUnion) asAny() any {
	if !param.IsOmitted(u.OfMeetingSessionCameraImageBase64Source) {
		return u.OfMeetingSessionCameraImageBase64Source
	} else if !param.IsOmitted(u.OfMeetingSessionCameraImageURLSource) {
		return u.OfMeetingSessionCameraImageURLSource
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MeetingSessionNewParamsCameraImageUnion) GetBase64Data() *string {
	if vt := u.OfMeetingSessionCameraImageBase64Source; vt != nil {
		return &vt.Base64Data
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MeetingSessionNewParamsCameraImageUnion) GetURL() *string {
	if vt := u.OfMeetingSessionCameraImageURLSource; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MeetingSessionNewParamsCameraImageUnion) GetFormat() *string {
	if vt := u.OfMeetingSessionCameraImageBase64Source; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfMeetingSessionCameraImageURLSource; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[MeetingSessionNewParamsCameraImageUnion](
		"",
		apijson.Variant[MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source](gjson.JSON),
		apijson.Variant[MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource](gjson.JSON),
	)
}

// The properties Base64Data, Format are required.
type MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source struct {
	// Canonical plain RFC 4648 Base64 for a valid decoded JPEG. Data URIs, whitespace,
	// and the URL-safe alphabet are rejected. The encoded value is limited to
	// 1,835,008 characters and the decoded JPEG to 1,363,148 bytes. The JPEG is
	// limited to 4,096 pixels per dimension, 4 megapixels, and 128 MB of decoder
	// memory. The image bytes are not returned in the session response.
	Base64Data string `json:"base64_data" api:"required"`
	// Only JPEG images are accepted.
	//
	// This field can be elided, and will marshal its zero value as "jpeg".
	Format constant.Jpeg `json:"format" default:"jpeg"`
	paramObj
}

func (r MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Format, URL are required.
type MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource struct {
	// Public HTTPS JPEG URL with at most 2,048 characters and no credentials,
	// fragment, surrounding whitespace, raw control characters, or explicit
	// non-default port. Signed queries are allowed but must be treated as credentials.
	// Fetching is limited to public network destinations, a five-second timeout, no
	// redirects, a 2xx image/jpeg response with identity or no content encoding, and a
	// 1,363,148-byte limit enforced against both declared and streamed content. The
	// service resolves the URL before bot creation and does not return the URL or
	// image bytes in the session response.
	URL string `json:"url" api:"required" format:"uri"`
	// Only JPEG images are accepted.
	//
	// This field can be elided, and will marshal its zero value as "jpeg".
	Format constant.Jpeg `json:"format" default:"jpeg"`
	paramObj
}

func (r MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParamsCameraImageMeetingSessionCameraImageURLSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording layout for Google Meet and Teams. Defaults to speaker_view. Accepted
// at creation only; does not enable recording by itself.
type MeetingSessionNewParamsRecordingView string

const (
	MeetingSessionNewParamsRecordingViewSpeakerView          MeetingSessionNewParamsRecordingView = "speaker_view"
	MeetingSessionNewParamsRecordingViewGalleryView          MeetingSessionNewParamsRecordingView = "gallery_view"
	MeetingSessionNewParamsRecordingViewSpeakerViewNoSidebar MeetingSessionNewParamsRecordingView = "speaker_view_no_sidebar"
)

type MeetingSessionUpdateParams struct {
	// Updated display name for the bot.
	BotName param.Opt[string] `json:"bot_name,omitzero"`
	// ISO-8601 timestamp for the bot to join. May be updated to reschedule.
	JoinAt param.Opt[time.Time] `json:"join_at,omitzero" format:"date-time"`
	paramObj
}

func (r MeetingSessionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionListParams struct {
	// Filter meeting sessions by current status.
	//
	// Any of "scheduled", "joining", "waiting_for_admission", "active", "leaving",
	// "ended", "failed", "admission_denied".
	Status MeetingSessionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingSessionListParams]'s query parameters as
// `url.Values`.
func (r MeetingSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter meeting sessions by current status.
type MeetingSessionListParamsStatus string

const (
	MeetingSessionListParamsStatusScheduled           MeetingSessionListParamsStatus = "scheduled"
	MeetingSessionListParamsStatusJoining             MeetingSessionListParamsStatus = "joining"
	MeetingSessionListParamsStatusWaitingForAdmission MeetingSessionListParamsStatus = "waiting_for_admission"
	MeetingSessionListParamsStatusActive              MeetingSessionListParamsStatus = "active"
	MeetingSessionListParamsStatusLeaving             MeetingSessionListParamsStatus = "leaving"
	MeetingSessionListParamsStatusEnded               MeetingSessionListParamsStatus = "ended"
	MeetingSessionListParamsStatusFailed              MeetingSessionListParamsStatus = "failed"
	MeetingSessionListParamsStatusAdmissionDenied     MeetingSessionListParamsStatus = "admission_denied"
)

type MeetingSessionGetEventsParams struct {
	// Return results with a cursor position after this value.
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// Maximum number of results to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingSessionGetEventsParams]'s query parameters as
// `url.Values`.
func (r MeetingSessionGetEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeetingSessionGetTranscriptParams struct {
	// Return results with a cursor position after this value.
	After param.Opt[int64] `query:"after,omitzero" json:"-"`
	// Maximum number of results to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Long-poll duration in seconds. The server holds the connection open for up to
	// this many seconds, waiting for new or updated results before returning an empty
	// response. Set to 0 for an immediate response.
	WaitSeconds param.Opt[int64] `query:"wait_seconds,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingSessionGetTranscriptParams]'s query parameters as
// `url.Values`.
func (r MeetingSessionGetTranscriptParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
