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
)

// Register and manage text-to-speech (TTS) engines you publish, and find the
// public engines other organizations published, so an AI Assistant, TeXML or Call
// Control can speak with one of their voices, using your own key for it, instead
// of a Telnyx-integrated vendor.
//
// TtsIntegrationActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTtsIntegrationActionService] method instead.
type TtsIntegrationActionService struct {
	Options []option.RequestOption
}

// NewTtsIntegrationActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTtsIntegrationActionService(opts ...option.RequestOption) (r TtsIntegrationActionService) {
	r = TtsIntegrationActionService{}
	r.Options = opts
	return
}

// Pause an `active` engine so it stops serving traffic, without deleting it.
// Owner-initiated and owner-reversible via `enable`. Only an `active` engine can
// be disabled. While an engine is disabled, no organization can use it, including
// other organizations if it is `public`.
func (r *TtsIntegrationActionService) Disable(ctx context.Context, id string, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/actions/disable", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Resume a `disabled` engine. Only a `disabled` engine can be enabled; a
// `suspended` engine can only be reinstated by Telnyx.
func (r *TtsIntegrationActionService) Enable(ctx context.Context, id string, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/actions/enable", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Make Telnyx read the engine's `voices_url` now instead of at the next scheduled
// read (every 15 minutes). On success the stored list is replaced and
// `voices_degraded` is cleared. The previous list is kept when the read fails.
func (r *TtsIntegrationActionService) RefreshVoices(ctx context.Context, id string, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/actions/refresh_voices", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Ask Telnyx to review a `draft` engine so it can serve other organizations. Only
// a `draft` engine with no pending or approved review can be submitted (a rejected
// one can be submitted again). The engine moves to `review_status: pending`;
// Telnyx approves it (`status: active`, `review_status: approved`) or rejects it
// with `review_notes`. See the
// [Community TTS engines guide](/docs/voice/tts/community-engines) for the full
// review flow.
func (r *TtsIntegrationActionService) RequestReview(ctx context.Context, id string, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/actions/request_review", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Check an engine end to end with your own key, without a call. Telnyx reads the
// voice list, connects to the engine's `websocket_url` with the Integration Secret
// named by `api_key_ref`, sends a short text, and checks the answer: the
// handshake, the credentials, the audio, the final frame and a clean close. The
// result is recorded (see `conformance_runs`) and returned. A test within 10
// seconds of the previous one is refused with `429`, and a `suspended` engine
// cannot be tested. See
// [Try your engine before approval](/docs/voice/tts/community-engines#try-your-engine-before-approval).
func (r *TtsIntegrationActionService) TestConnection(ctx context.Context, id string, body TtsIntegrationActionTestConnectionParams, opts ...option.RequestOption) (res *CommunityTtsConformanceRun, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/actions/test_connection", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// One recorded run of `test_connection`.
type CommunityTtsConformanceRun struct {
	ID string `json:"id" api:"required"`
	// Each check is `ok`, `failed` or `skipped`: `voices_url_reachable`,
	// `voices_schema_valid`, `handshake`, `auth`, `init_accepted`, `audio_received`,
	// `final_frame_received`, `clean_close`.
	//
	// Any of "ok", "failed", "skipped".
	Checks map[string]string `json:"checks" api:"required"`
	// `test_connection`: run by the owner. `conformance`: run by Telnyx during review.
	//
	// Any of "test_connection", "conformance".
	Kind CommunityTtsConformanceRunKind `json:"kind" api:"required"`
	// `true` when every check is `ok`.
	Success bool `json:"success" api:"required"`
	// The first failure, in words.
	Error string `json:"error" api:"nullable"`
	// How long the whole run took.
	LatencyMs          int64     `json:"latency_ms" api:"nullable"`
	StartedAt          time.Time `json:"started_at" format:"date-time"`
	TimeToFirstAudioMs int64     `json:"time_to_first_audio_ms" api:"nullable"`
	VoiceID            string    `json:"voice_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Checks             respjson.Field
		Kind               respjson.Field
		Success            respjson.Field
		Error              respjson.Field
		LatencyMs          respjson.Field
		StartedAt          respjson.Field
		TimeToFirstAudioMs respjson.Field
		VoiceID            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommunityTtsConformanceRun) RawJSON() string { return r.JSON.raw }
func (r *CommunityTtsConformanceRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `test_connection`: run by the owner. `conformance`: run by Telnyx during review.
type CommunityTtsConformanceRunKind string

const (
	CommunityTtsConformanceRunKindTestConnection CommunityTtsConformanceRunKind = "test_connection"
	CommunityTtsConformanceRunKindConformance    CommunityTtsConformanceRunKind = "conformance"
)

type TtsIntegrationActionTestConnectionParams struct {
	// The identifier of the Integration Secret in your organization that Telnyx
	// connects to the engine with, as a caller would.
	APIKeyRef string `json:"api_key_ref" api:"required"`
	// The voice to test. Default: the first voice in the engine's list. It must be in
	// the list.
	VoiceID param.Opt[string] `json:"voice_id,omitzero"`
	paramObj
}

func (r TtsIntegrationActionTestConnectionParams) MarshalJSON() (data []byte, err error) {
	type shadow TtsIntegrationActionTestConnectionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TtsIntegrationActionTestConnectionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
