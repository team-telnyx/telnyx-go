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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
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

// Creates a new meeting session. When an idempotency_key is supplied in the
// request body, replay lookup is scoped to the authenticated account and compares
// only the key; the request payload is not fingerprinted or compared. If a session
// with that key already exists for the account, the existing session is replayed
// (200); otherwise a new session is created (201). Supports bring-your-own-key
// (BYOK) configuration. The session may enter asynchronous states (e.g. joining,
// waiting_for_admission) before becoming active. Optional `camera_image` input is
// write-only and applies only when no Avatar or Assistant webpage output takes
// precedence. An ignored URL is not fetched. An effective URL source is resolved
// before bot creation; neither the source URL nor image bytes are persisted,
// returned, or logged. Treat signed URLs as credentials.
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
	path := fmt.Sprintf("meeting_sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates mutable properties of a meeting session. Only sessions in the scheduled
// state can be updated; any other state returns 409 with the invalid_state error
// code. All request fields are optional, and an empty object is a valid no-op
// update.
func (r *MeetingSessionService) Update(ctx context.Context, id string, body MeetingSessionUpdateParams, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a list of meeting sessions, optionally filtered by status.
func (r *MeetingSessionService) List(ctx context.Context, query MeetingSessionListParams, opts ...option.RequestOption) (res *MeetingSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "meeting_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Stops a meeting session without deleting its persisted record. Scheduled bots
// are cancelled, while bots that are joining or active are asked to leave. The
// persisted meeting session record remains available.
func (r *MeetingSessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// **Not yet available in production** — this route is not currently routed on
// api.telnyx.com and returns a generic 404; it is documented ahead of rollout.
// Irreversibly requests deletion of provider-hosted aggregate recording media
// under the provider contract. The operation retains the Telnyx-local Meeting
// session, transcript segments, events, artifacts, and usage records. It is
// separate from `DELETE /meeting_sessions/{id}`, which stops or cancels
// participation without deleting the persisted session. A missing/foreign session
// returns 404; provider deletion failures return 502.
func (r *MeetingSessionService) DeleteRecordingMedia(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionDeleteRecordingMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/recording_media", id)
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
	path := fmt.Sprintf("meeting_sessions/%s/events", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns recordings for a meeting session.
func (r *MeetingSessionService) GetRecordings(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionGetRecordingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/recordings", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns transcript segments ordered by ascending `seq`. Default `limit` is 100
// and maximum is 1,000. Continue with `after=meta.next_after`. A long-poll timeout
// returns 200 with empty `data` and `meta.next_after: null`; retain the cursor
// supplied to that request because null is not a replacement cursor.
func (r *MeetingSessionService) GetTranscript(ctx context.Context, id string, query MeetingSessionGetTranscriptParams, opts ...option.RequestOption) (res *MeetingSessionGetTranscriptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/transcript", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a meeting session. All serializer fields are present and required;
// nullable fields use null when absent. No actor, provider-bot, idempotency,
// routing, key, or internal fields are exposed.
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
	// Avatar configuration if an avatar is attached, otherwise null.
	Avatar MeetingSessionAvatar `json:"avatar" api:"required"`
	// Current state of the avatar connection, or null if no avatar is attached.
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
	// Human-readable failure reason if the session failed, or null.
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
	// Provider handling the meeting session.
	Provider string `json:"provider" api:"required"`
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
	// Additional human-readable detail about the status, or null.
	StatusDetail string `json:"status_detail" api:"required"`
	// Timestamp of the last update to the session.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Webhook endpoint for session lifecycle callbacks, or null if not configured.
	WebhookURL string `json:"webhook_url" api:"required" format:"uri"`
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
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSession) RawJSON() string { return r.JSON.raw }
func (r *MeetingSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Assistant configuration if an assistant is attached, otherwise null.
type MeetingSessionAssistant struct {
	// Identifier of the assistant.
	ID string `json:"id" api:"required"`
	// Audio gating strategy for the assistant call leg.
	//
	// Any of "none", "half_duplex".
	AudioGate string `json:"audio_gate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AudioGate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionAssistant) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionAssistant) UnmarshalJSON(data []byte) error {
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

// Avatar configuration if an avatar is attached, otherwise null.
type MeetingSessionAvatar struct {
	// Identifier of the avatar.
	AvatarID string `json:"avatar_id" api:"required"`
	// Avatar provider identifier.
	Provider constant.Anam `json:"provider" default:"anam"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvatarID    respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionAvatar) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current state of the avatar connection, or null if no avatar is attached.
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
	// Text spoken on meeting entry, or null if not set.
	SpeakOnEnter string `json:"speak_on_enter" api:"required"`
	// Whether a summary artifact is generated on session end.
	SummarizeOnEnd bool `json:"summarize_on_end" api:"required"`
	// Configured voice identifier, or null if not set.
	Voice string `json:"voice" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BargeIn        respjson.Field
		SpeakOnEnter   respjson.Field
		SummarizeOnEnd respjson.Field
		Voice          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionConfig) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionConfig) UnmarshalJSON(data []byte) error {
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

type MeetingSessionResponse struct {
	// Represents a meeting session. All serializer fields are present and required;
	// nullable fields use null when absent. No actor, provider-bot, idempotency,
	// routing, key, or internal fields are exposed.
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
	Provider         constant.Recall                 `json:"provider" default:"recall"`
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
	Data []MeetingSessionGetRecordingsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionGetRecordingsResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionGetRecordingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionGetRecordingsResponseData struct {
	// Expiry timestamp when supplied by the provider, or null. The current adapter
	// returns null.
	ExpiresAt string `json:"expires_at" api:"required"`
	Type      string `json:"type" api:"required"`
	// Current provider download URL. The API does not guarantee URL lifetime or
	// refresh behavior.
	URL string `json:"url" api:"required" format:"uri"`
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
	Confidence   float64   `json:"confidence" api:"required"`
	OccurredAt   time.Time `json:"occurred_at" api:"required" format:"date-time"`
	RelativeTs   float64   `json:"relative_ts" api:"required"`
	Seq          int64     `json:"seq" api:"required"`
	SpeakerLabel string    `json:"speaker_label" api:"required"`
	Text         string    `json:"text" api:"required"`
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
	// Client-supplied idempotency key to safely retry creation requests without
	// duplicating sessions. Lookup is scoped to the authenticated account and compares
	// the key only; the request payload is not fingerprinted or compared.
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	// ISO-8601 timestamp in the future at which the bot should join. If omitted, the
	// bot joins immediately.
	JoinAt param.Opt[time.Time] `json:"join_at,omitzero" format:"date-time"`
	// Text the bot speaks when it enters the meeting.
	SpeakOnEnter param.Opt[string] `json:"speak_on_enter,omitzero"`
	// If true, generate a summary artifact when the session ends.
	SummarizeOnEnd param.Opt[bool] `json:"summarize_on_end,omitzero"`
	// Session-default voice identifier used for `speak_on_enter` and ordinary speak
	// actions. A voice supplied on an individual speak action overrides this default
	// for that utterance.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// HTTPS endpoint to receive session lifecycle callbacks. Static validation
	// requires HTTPS, rejects embedded credentials and blocked hosts, and enforces
	// egress policy. Validation makes no network request to the endpoint.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"uri"`
	// Request options for attaching a voice assistant to the session. Routing fields
	// (`call_control_connection_id`, `from`, and `loopback_sip_uri`) are used only to
	// establish the assistant call leg and are omitted from response objects.
	// `audio_gate` is returned with `id` in the assistant response object.
	Assistant MeetingSessionNewParamsAssistant `json:"assistant,omitzero"`
	// Request options for attaching a bring-your-own-key avatar to the session.
	Avatar MeetingSessionNewParamsAvatar `json:"avatar,omitzero"`
	// Write-only static camera-tile image for this session, not a native account or
	// participant profile photo. Supply exactly one JPEG source. When effective, the
	// image is used as the bot's static camera/video output; presentation varies by
	// meeting platform and recording configuration and is not guaranteed in
	// recordings. An effective Avatar or Assistant webpage output takes precedence, so
	// this input is ignored and a URL source is not fetched.
	CameraImage MeetingSessionNewParamsCameraImageUnion `json:"camera_image,omitzero"`
	// Arbitrary key-value metadata attached to the session. The serialized JSON
	// representation must not exceed 16384 characters at runtime.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r MeetingSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request options for attaching a voice assistant to the session. Routing fields
// (`call_control_connection_id`, `from`, and `loopback_sip_uri`) are used only to
// establish the assistant call leg and are omitted from response objects.
// `audio_gate` is returned with `id` in the assistant response object.
//
// The properties ID, CallControlConnectionID, From, LoopbackSipUri are required.
type MeetingSessionNewParamsAssistant struct {
	// Identifier of the assistant to attach.
	ID string `json:"id" api:"required"`
	// Call control connection used to bridge the assistant into the meeting audio.
	CallControlConnectionID string `json:"call_control_connection_id" api:"required"`
	// E.164 calling number used as the originating party for the assistant call leg.
	From string `json:"from" api:"required"`
	// SIP URI to which the assistant media loopback is established.
	LoopbackSipUri string `json:"loopback_sip_uri" api:"required"`
	// Audio gating strategy for the assistant call leg.
	//
	// Any of "none", "half_duplex".
	AudioGate string `json:"audio_gate,omitzero"`
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
		"audio_gate", "none", "half_duplex",
	)
}

// Request options for attaching a bring-your-own-key avatar to the session.
//
// The properties APIKey, AvatarID, Provider are required.
type MeetingSessionNewParamsAvatar struct {
	// Bring-your-own-key API key for the avatar provider. The key is never stored or
	// returned by the API.
	APIKey string `json:"api_key" api:"required"`
	// Identifier of the avatar to use.
	AvatarID string `json:"avatar_id" api:"required"`
	// Avatar provider identifier. Currently only "anam" is supported.
	//
	// This field can be elided, and will marshal its zero value as "anam".
	Provider constant.Anam `json:"provider" default:"anam"`
	paramObj
}

func (r MeetingSessionNewParamsAvatar) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionNewParamsAvatar
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionNewParamsAvatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// The properties Base64Data, Format are required.
type MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source struct {
	// Canonical plain RFC 4648 Base64 for a valid decoded JPEG. Data URIs, whitespace,
	// and the URL-safe alphabet are rejected. The encoded value is limited to
	// 1,835,008 characters and the decoded JPEG to 1,363,148 bytes. The JPEG is
	// limited to 4,096 pixels per dimension, 4 megapixels, and 128 MB of decoder
	// memory. The image bytes are not persisted, returned, or logged.
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
	// service resolves the URL before bot creation and does not persist, return, or
	// log the URL or image bytes.
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
