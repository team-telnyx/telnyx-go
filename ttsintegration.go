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
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
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
// TtsIntegrationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTtsIntegrationService] method instead.
type TtsIntegrationService struct {
	Options []option.RequestOption
	// Register and manage text-to-speech (TTS) engines you publish, and find the
	// public engines other organizations published, so an AI Assistant, TeXML or Call
	// Control can speak with one of their voices, using your own key for it, instead
	// of a Telnyx-integrated vendor.
	Actions TtsIntegrationActionService
}

// NewTtsIntegrationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTtsIntegrationService(opts ...option.RequestOption) (r TtsIntegrationService) {
	r = TtsIntegrationService{}
	r.Options = opts
	r.Actions = NewTtsIntegrationActionService(opts...)
	return
}

// Register a WebSocket text-to-speech engine. Telnyx reads the engine's
// `voices_url` while it registers the engine and refuses the registration with a
// `422` if the list is unreachable or invalid. The engine starts at
// `status: draft` and cannot serve other organizations' traffic until Telnyx
// approves it (your own organization can try it with its own key meanwhile). It is
// `visibility: private` unless you set `public`: a private engine can be used only
// by your organization, a public one by any organization once it is `active`. See
// the [Community TTS engines guide](/docs/voice/tts/community-engines) for the
// full review flow.
//
// `websocket_url` must be `wss://` and `voices_url` must be `https://`, both on a
// public host; a plain `ws://` URL is rejected, and so is a credential embedded in
// either. Callers bring their own key for your engine: Telnyx stores none, so
// there is nothing to register here for synthesis. `voices_api_key_ref` is
// optional and used only to read your voice list.
func (r *TtsIntegrationService) New(ctx context.Context, body TtsIntegrationNewParams, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tts_integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve one of your own community TTS engines by ID.
func (r *TtsIntegrationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Same as `PUT`: only the fields you send change.
func (r *TtsIntegrationService) Update(ctx context.Context, id string, body TtsIntegrationUpdateParams, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List every community TTS engine your organization has registered, at any status
// or visibility. Use `GET /tts_integrations/discoverable` instead to also see
// other organizations' public, active engines.
func (r *TtsIntegrationService) List(ctx context.Context, query TtsIntegrationListParams, opts ...option.RequestOption) (res *TtsIntegrationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tts_integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete one of your own community TTS engines. The engine disappears from every
// listing and lookup and can no longer serve traffic, but its `name` stays
// reserved for your organization: no other organization can register it, and you
// can register it again later. An assistant or call still configured with one of
// its voices cannot use it until an engine with that name exists again.
func (r *TtsIntegrationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("tts_integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Change one of your own community TTS engines. Every field is optional: omit a
// field to leave it unchanged. `name` cannot be changed. Telnyx reads the
// `voices_url` again and refuses an update that leaves the engine with an
// unreadable list. Changing `websocket_url`, `voices_url`,
// `supported_audio_formats` or `capabilities`, or making a `private` engine
// `public`, sends the engine back to `status: draft` and clears its review. See
// the [Community TTS engines guide](/docs/voice/tts/community-engines) for the
// full review flow.
func (r *TtsIntegrationService) Replace(ctx context.Context, id string, body TtsIntegrationReplaceParams, opts ...option.RequestOption) (res *CommunityTtsEngine, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// The recorded runs of `test_connection` for one of your engines, newest first.
// Telnyx keeps the latest 50.
func (r *TtsIntegrationService) GetConformanceRuns(ctx context.Context, id string, query TtsIntegrationGetConformanceRunsParams, opts ...option.RequestOption) (res *TtsIntegrationGetConformanceRunsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/conformance_runs", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the engines you can use: your own, at any status, and other organizations'
// engines that are `public` and `active`. Each entry carries what a caller needs
// to choose an engine: its `name` (the first part of its voices,
// `<name>.<voice_id>`), `voices_count`, audio formats, capabilities, and
// `documentation_url`. It never carries the publisher's contact details, review
// notes, or `voices_api_key_ref`. Read an engine's voices with
// `GET /tts_integrations/{id}/voices`.
func (r *TtsIntegrationService) GetDiscoverable(ctx context.Context, query TtsIntegrationGetDiscoverableParams, opts ...option.RequestOption) (res *TtsIntegrationGetDiscoverableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "tts_integrations/discoverable"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// The voices Telnyx last read from the engine's `voices_url`, in the shape the
// engine published. A caller selects one as `<engine name>.<voice_id>`. The list
// is replaced only when a new read succeeds, so an engine whose `voices_url` is
// down keeps serving its last list. `[]` until the first read succeeds.
func (r *TtsIntegrationService) GetVoices(ctx context.Context, id string, opts ...option.RequestOption) (res *TtsIntegrationGetVoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tts_integrations/%s/voices", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An audio format an engine can produce. Offer `linear16`: it serves every
// surface. Telnyx converts `linear16` to MP3 for the commands that play MP3.
type CommunityTtsAudioFormat string

const (
	CommunityTtsAudioFormatMP3       CommunityTtsAudioFormat = "mp3"
	CommunityTtsAudioFormatLinear16  CommunityTtsAudioFormat = "linear16"
	CommunityTtsAudioFormatWav       CommunityTtsAudioFormat = "wav"
	CommunityTtsAudioFormatMulaw     CommunityTtsAudioFormat = "mulaw"
	CommunityTtsAudioFormatAlaw      CommunityTtsAudioFormat = "alaw"
	CommunityTtsAudioFormatOggVorbis CommunityTtsAudioFormat = "ogg_vorbis"
)

// The full record of a community TTS engine you own. Select its voices by the
// engine's `name`, as `<name>.<voice_id>`, with your own key for it
// (`api_key_ref`).
type CommunityTtsEngine struct {
	ID string `json:"id" api:"required"`
	// What the engine supports. An unknown key is rejected.
	Capabilities CommunityTtsEngineCapabilities `json:"capabilities" api:"required"`
	// The engine's unique name: the first part of its voices.
	Name string `json:"name" api:"required"`
	// Whether the engine can serve traffic right now, independent of `visibility` and
	// `review_status`. `draft`: not yet approved (new and awaiting its first review,
	// rejected and awaiting a fix, or approved earlier and awaiting re-review after a
	// review-sensitive change). `active`: serving traffic. `disabled`: owner-paused,
	// reversible via `enable`. `suspended`: paused by Telnyx (repeated failures of a
	// public engine, or an abuse action); the owner cannot clear this by calling
	// `enable`, and Telnyx can reinstate it.
	//
	// Any of "draft", "active", "disabled", "suspended".
	Status                CommunityTtsEngineStatus  `json:"status" api:"required"`
	SupportedAudioFormats []CommunityTtsAudioFormat `json:"supported_audio_formats" api:"required"`
	// `private`: only your organization can use the engine. `public`: any organization
	// can, once it is `active`.
	//
	// Any of "private", "public".
	Visibility             CommunityTtsEngineVisibility `json:"visibility" api:"required"`
	VoicesURL              string                       `json:"voices_url" api:"required"`
	WebsocketURL           string                       `json:"websocket_url" api:"required"`
	ConnectTimeoutMs       int64                        `json:"connect_timeout_ms"`
	ContactEmail           string                       `json:"contact_email" api:"nullable"`
	CreatedAt              time.Time                    `json:"created_at" format:"date-time"`
	DataRetentionPolicyURL string                       `json:"data_retention_policy_url" api:"nullable"`
	DocumentationURL       string                       `json:"documentation_url" api:"nullable"`
	// The reviewer's notes on a rejection.
	ReviewNotes string `json:"review_notes" api:"nullable"`
	// `null`: never submitted, or an approved engine whose review-sensitive field
	// changed and awaits re-review. `pending`: submitted. `approved`: approved.
	// `rejected`: rejected, with `review_notes`.
	//
	// Any of "pending", "approved", "rejected".
	ReviewStatus    CommunityTtsEngineReviewStatus `json:"review_status" api:"nullable"`
	ReviewedAt      time.Time                      `json:"reviewed_at" api:"nullable" format:"date-time"`
	UpdatedAt       time.Time                      `json:"updated_at" api:"nullable" format:"date-time"`
	VoicesAPIKeyRef string                         `json:"voices_api_key_ref" api:"nullable"`
	// The number of voices in the stored list.
	VoicesCount int64 `json:"voices_count"`
	// `true` after three failed reads in a row, until a read succeeds. A degraded
	// engine keeps serving from its last list.
	VoicesDegraded bool `json:"voices_degraded"`
	// When the stored list was last read successfully.
	VoicesFetchedAt time.Time `json:"voices_fetched_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Capabilities           respjson.Field
		Name                   respjson.Field
		Status                 respjson.Field
		SupportedAudioFormats  respjson.Field
		Visibility             respjson.Field
		VoicesURL              respjson.Field
		WebsocketURL           respjson.Field
		ConnectTimeoutMs       respjson.Field
		ContactEmail           respjson.Field
		CreatedAt              respjson.Field
		DataRetentionPolicyURL respjson.Field
		DocumentationURL       respjson.Field
		ReviewNotes            respjson.Field
		ReviewStatus           respjson.Field
		ReviewedAt             respjson.Field
		UpdatedAt              respjson.Field
		VoicesAPIKeyRef        respjson.Field
		VoicesCount            respjson.Field
		VoicesDegraded         respjson.Field
		VoicesFetchedAt        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommunityTtsEngine) RawJSON() string { return r.JSON.raw }
func (r *CommunityTtsEngine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `null`: never submitted, or an approved engine whose review-sensitive field
// changed and awaits re-review. `pending`: submitted. `approved`: approved.
// `rejected`: rejected, with `review_notes`.
type CommunityTtsEngineReviewStatus string

const (
	CommunityTtsEngineReviewStatusPending  CommunityTtsEngineReviewStatus = "pending"
	CommunityTtsEngineReviewStatusApproved CommunityTtsEngineReviewStatus = "approved"
	CommunityTtsEngineReviewStatusRejected CommunityTtsEngineReviewStatus = "rejected"
)

// What the engine supports. An unknown key is rejected.
type CommunityTtsEngineCapabilities struct {
	// Required. `true` when the engine sends audio in small frames as it synthesizes;
	// `false` when it sends one frame per sentence.
	StreamedAudio bool `json:"streamed_audio" api:"required"`
	// The engine speaks the buffered text when a frame carries `"flush": true`.
	FlushSupport bool `json:"flush_support"`
	// The engine supports a restart of the current speech (barge-in).
	ForceInterruptSupport bool `json:"force_interrupt_support"`
	// The engine accepts SSML in the text.
	SsmlSupport bool `json:"ssml_support"`
	// The settings callers can pass to the engine in `voice_settings`, one entry per
	// setting: `{"type": "number" | "integer", "min", "max"}`, `{"type": "boolean"}`,
	// `{"type": "string"}` (at most 1000 characters) or
	// `{"type": "enum", "values": [...]}`, each with an optional `default`. A caller's
	// settings are checked against it before Telnyx connects to the engine: an
	// undeclared setting, a value of the wrong type, or a value out of bounds is
	// refused. An engine that declares none receives the settings as sent.
	VoiceSettingsSchema map[string]any `json:"voice_settings_schema" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StreamedAudio         respjson.Field
		FlushSupport          respjson.Field
		ForceInterruptSupport respjson.Field
		SsmlSupport           respjson.Field
		VoiceSettingsSchema   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommunityTtsEngineCapabilities) RawJSON() string { return r.JSON.raw }
func (r *CommunityTtsEngineCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CommunityTtsEngineCapabilities to a
// CommunityTtsEngineCapabilitiesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CommunityTtsEngineCapabilitiesParam.Overrides()
func (r CommunityTtsEngineCapabilities) ToParam() CommunityTtsEngineCapabilitiesParam {
	return param.Override[CommunityTtsEngineCapabilitiesParam](json.RawMessage(r.RawJSON()))
}

// What the engine supports. An unknown key is rejected.
//
// The property StreamedAudio is required.
type CommunityTtsEngineCapabilitiesParam struct {
	// Required. `true` when the engine sends audio in small frames as it synthesizes;
	// `false` when it sends one frame per sentence.
	StreamedAudio bool `json:"streamed_audio" api:"required"`
	// The engine speaks the buffered text when a frame carries `"flush": true`.
	FlushSupport param.Opt[bool] `json:"flush_support,omitzero"`
	// The engine supports a restart of the current speech (barge-in).
	ForceInterruptSupport param.Opt[bool] `json:"force_interrupt_support,omitzero"`
	// The engine accepts SSML in the text.
	SsmlSupport param.Opt[bool] `json:"ssml_support,omitzero"`
	// The settings callers can pass to the engine in `voice_settings`, one entry per
	// setting: `{"type": "number" | "integer", "min", "max"}`, `{"type": "boolean"}`,
	// `{"type": "string"}` (at most 1000 characters) or
	// `{"type": "enum", "values": [...]}`, each with an optional `default`. A caller's
	// settings are checked against it before Telnyx connects to the engine: an
	// undeclared setting, a value of the wrong type, or a value out of bounds is
	// refused. An engine that declares none receives the settings as sent.
	VoiceSettingsSchema map[string]any `json:"voice_settings_schema,omitzero"`
	paramObj
}

func (r CommunityTtsEngineCapabilitiesParam) MarshalJSON() (data []byte, err error) {
	type shadow CommunityTtsEngineCapabilitiesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommunityTtsEngineCapabilitiesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommunityTtsEngineListMeta struct {
	PageNumber   int64 `json:"page_number" api:"required"`
	PageSize     int64 `json:"page_size" api:"required"`
	TotalPages   int64 `json:"total_pages" api:"required"`
	TotalResults int64 `json:"total_results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommunityTtsEngineListMeta) RawJSON() string { return r.JSON.raw }
func (r *CommunityTtsEngineListMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the engine can serve traffic right now, independent of `visibility` and
// `review_status`. `draft`: not yet approved (new and awaiting its first review,
// rejected and awaiting a fix, or approved earlier and awaiting re-review after a
// review-sensitive change). `active`: serving traffic. `disabled`: owner-paused,
// reversible via `enable`. `suspended`: paused by Telnyx (repeated failures of a
// public engine, or an abuse action); the owner cannot clear this by calling
// `enable`, and Telnyx can reinstate it.
type CommunityTtsEngineStatus string

const (
	CommunityTtsEngineStatusDraft     CommunityTtsEngineStatus = "draft"
	CommunityTtsEngineStatusActive    CommunityTtsEngineStatus = "active"
	CommunityTtsEngineStatusDisabled  CommunityTtsEngineStatus = "disabled"
	CommunityTtsEngineStatusSuspended CommunityTtsEngineStatus = "suspended"
)

// `private`: only your organization can use the engine. `public`: any organization
// can, once it is `active`.
type CommunityTtsEngineVisibility string

const (
	CommunityTtsEngineVisibilityPrivate CommunityTtsEngineVisibility = "private"
	CommunityTtsEngineVisibilityPublic  CommunityTtsEngineVisibility = "public"
)

// Only the fields you send change. `name` cannot be changed. Sending `null` for
// `websocket_url`, `voices_url`, `supported_audio_formats`, `capabilities` or
// `connect_timeout_ms` is refused. Changing `websocket_url`, `voices_url`,
// `supported_audio_formats` or `capabilities`, or making a `private` engine
// `public`, sends the engine back to `status: draft` and clears its review.
type UpdateCommunityTtsEngineRequestParam struct {
	ContactEmail           param.Opt[string] `json:"contact_email,omitzero"`
	DataRetentionPolicyURL param.Opt[string] `json:"data_retention_policy_url,omitzero"`
	DocumentationURL       param.Opt[string] `json:"documentation_url,omitzero"`
	VoicesAPIKeyRef        param.Opt[string] `json:"voices_api_key_ref,omitzero"`
	ConnectTimeoutMs       param.Opt[int64]  `json:"connect_timeout_ms,omitzero"`
	// See the create request.
	VoicesURL param.Opt[string] `json:"voices_url,omitzero"`
	// See the create request.
	WebsocketURL param.Opt[string] `json:"websocket_url,omitzero"`
	// What the engine supports. An unknown key is rejected.
	Capabilities          CommunityTtsEngineCapabilitiesParam `json:"capabilities,omitzero"`
	SupportedAudioFormats []CommunityTtsAudioFormat           `json:"supported_audio_formats,omitzero"`
	// `private`: only your organization can use the engine. `public`: any organization
	// can, once it is `active`.
	//
	// Any of "private", "public".
	Visibility CommunityTtsEngineVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r UpdateCommunityTtsEngineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCommunityTtsEngineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCommunityTtsEngineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationListResponse struct {
	Data []CommunityTtsEngine       `json:"data" api:"required"`
	Meta CommunityTtsEngineListMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationGetConformanceRunsResponse struct {
	Data []CommunityTtsConformanceRun `json:"data" api:"required"`
	Meta CommunityTtsEngineListMeta   `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationGetConformanceRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationGetConformanceRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationGetDiscoverableResponse struct {
	Data []TtsIntegrationGetDiscoverableResponseData `json:"data" api:"required"`
	Meta CommunityTtsEngineListMeta                  `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationGetDiscoverableResponse) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationGetDiscoverableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a caller needs to choose an engine. It carries no publisher contact
// details, review notes or credential references.
type TtsIntegrationGetDiscoverableResponseData struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Whether the engine can serve traffic right now, independent of `visibility` and
	// `review_status`. `draft`: not yet approved (new and awaiting its first review,
	// rejected and awaiting a fix, or approved earlier and awaiting re-review after a
	// review-sensitive change). `active`: serving traffic. `disabled`: owner-paused,
	// reversible via `enable`. `suspended`: paused by Telnyx (repeated failures of a
	// public engine, or an abuse action); the owner cannot clear this by calling
	// `enable`, and Telnyx can reinstate it.
	//
	// Any of "draft", "active", "disabled", "suspended".
	Status CommunityTtsEngineStatus `json:"status" api:"required"`
	// `private`: only your organization can use the engine. `public`: any organization
	// can, once it is `active`.
	//
	// Any of "private", "public".
	Visibility CommunityTtsEngineVisibility `json:"visibility" api:"required"`
	// What the engine supports. An unknown key is rejected.
	Capabilities           CommunityTtsEngineCapabilities `json:"capabilities"`
	ConnectTimeoutMs       int64                          `json:"connect_timeout_ms"`
	CreatedAt              time.Time                      `json:"created_at" format:"date-time"`
	DataRetentionPolicyURL string                         `json:"data_retention_policy_url" api:"nullable"`
	DocumentationURL       string                         `json:"documentation_url" api:"nullable"`
	SupportedAudioFormats  []CommunityTtsAudioFormat      `json:"supported_audio_formats"`
	VoicesCount            int64                          `json:"voices_count"`
	WebsocketURL           string                         `json:"websocket_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Name                   respjson.Field
		Status                 respjson.Field
		Visibility             respjson.Field
		Capabilities           respjson.Field
		ConnectTimeoutMs       respjson.Field
		CreatedAt              respjson.Field
		DataRetentionPolicyURL respjson.Field
		DocumentationURL       respjson.Field
		SupportedAudioFormats  respjson.Field
		VoicesCount            respjson.Field
		WebsocketURL           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationGetDiscoverableResponseData) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationGetDiscoverableResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationGetVoicesResponse struct {
	Voices []TtsIntegrationGetVoicesResponseVoice `json:"voices" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationGetVoicesResponse) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationGetVoicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationGetVoicesResponseVoice struct {
	// Unique in the list; no dots. A caller selects the voice as
	// `<engine name>.<voice_id>`.
	VoiceID string `json:"voice_id" api:"required"`
	// A display string, for example `male`, `female` or `neutral`.
	Gender string `json:"gender"`
	// A BCP-47 tag, at most 35 characters.
	Language string `json:"language"`
	// A display name.
	Name string `json:"name"`
	// A display string.
	Provider string `json:"provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VoiceID     respjson.Field
		Gender      respjson.Field
		Language    respjson.Field
		Name        respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsIntegrationGetVoicesResponseVoice) RawJSON() string { return r.JSON.raw }
func (r *TtsIntegrationGetVoicesResponseVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationNewParams struct {
	// What the engine supports. An unknown key is rejected.
	Capabilities CommunityTtsEngineCapabilitiesParam `json:"capabilities,omitzero" api:"required"`
	// The engine's unique name, and the first part of its voices
	// (`<name>.<voice_id>`). 1-64 characters: lowercase letters, digits, hyphens and
	// underscores, starting with a letter or digit. Built-in vendor names and the
	// generic voice names `male`, `female`, `man`, `woman`, `alice` and `custom` are
	// reserved. Unique across all Telnyx accounts and cannot be changed later. A
	// deleted engine's name stays reserved for the organization that owned it.
	Name string `json:"name" api:"required"`
	// The audio formats the engine can produce, at least one, without duplicates.
	SupportedAudioFormats []CommunityTtsAudioFormat `json:"supported_audio_formats,omitzero" api:"required"`
	// The `https://` URL of the engine's voice listing, with the same host rules. It
	// answers `GET` with
	// `{"voices": [{"provider", "name", "voice_id", "language", "gender"}]}`: at most
	// 1000 voices, each `voice_id` unique and matching
	// `^[A-Za-z0-9][A-Za-z0-9_-]{0,127}$`, each `language` at most 35 characters.
	// Telnyx reads it when the engine is registered or updated, every 15 minutes, and
	// on `refresh_voices`.
	VoicesURL string `json:"voices_url" api:"required"`
	// The `wss://` URL Telnyx connects to. It must be reachable on the public
	// internet: plain `ws://`, private, loopback, link-local, shared, documentation,
	// multicast and reserved addresses, and internal hostnames are rejected, as are
	// credentials embedded in the URL or in credential-looking query parameters. At
	// most 2048 characters.
	WebsocketURL string `json:"websocket_url" api:"required"`
	// A contact address stored for the Telnyx platform team. Telnyx does not email it
	// automatically. At most 254 characters.
	ContactEmail param.Opt[string] `json:"contact_email,omitzero"`
	// URL of the publisher's data retention policy, at most 2048 characters. It is
	// returned in the discoverable listing.
	DataRetentionPolicyURL param.Opt[string] `json:"data_retention_policy_url,omitzero"`
	// An `https://` page where callers learn how to get a key for the engine and how
	// to use it, with a host and no credentials in it, at most 2048 characters. It is
	// returned in the discoverable listing.
	DocumentationURL param.Opt[string] `json:"documentation_url,omitzero"`
	// The identifier of an Integration Secret in your organization. When set, Telnyx
	// sends its value as a bearer token when it reads `voices_url`. It is never used
	// for synthesis. Send `null` on an update to clear it. At most 255 characters.
	VoicesAPIKeyRef param.Opt[string] `json:"voices_api_key_ref,omitzero"`
	// How long Telnyx waits for the WebSocket connection to open, from 500 to 10000.
	// Default: 3000.
	ConnectTimeoutMs param.Opt[int64] `json:"connect_timeout_ms,omitzero"`
	// `private`: only your organization can use the engine. `public`: any organization
	// can, once it is `active`.
	//
	// Any of "private", "public".
	Visibility CommunityTtsEngineVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r TtsIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TtsIntegrationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TtsIntegrationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationUpdateParams struct {
	// Only the fields you send change. `name` cannot be changed. Sending `null` for
	// `websocket_url`, `voices_url`, `supported_audio_formats`, `capabilities` or
	// `connect_timeout_ms` is refused. Changing `websocket_url`, `voices_url`,
	// `supported_audio_formats` or `capabilities`, or making a `private` engine
	// `public`, sends the engine back to `status: draft` and clears its review.
	UpdateCommunityTtsEngineRequest UpdateCommunityTtsEngineRequestParam
	paramObj
}

func (r TtsIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCommunityTtsEngineRequest)
}
func (r *TtsIntegrationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page TtsIntegrationListParamsPage `query:"page,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "draft", "active", "disabled", "suspended".
	Status CommunityTtsEngineStatus `query:"status,omitzero" json:"-"`
	// Filter by visibility.
	//
	// Any of "private", "public".
	Visibility CommunityTtsEngineVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationListParams]'s query parameters as
// `url.Values`.
func (r TtsIntegrationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type TtsIntegrationListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Clamped to 1-100.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationListParamsPage]'s query parameters as
// `url.Values`.
func (r TtsIntegrationListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TtsIntegrationReplaceParams struct {
	// Only the fields you send change. `name` cannot be changed. Sending `null` for
	// `websocket_url`, `voices_url`, `supported_audio_formats`, `capabilities` or
	// `connect_timeout_ms` is refused. Changing `websocket_url`, `voices_url`,
	// `supported_audio_formats` or `capabilities`, or making a `private` engine
	// `public`, sends the engine back to `status: draft` and clears its review.
	UpdateCommunityTtsEngineRequest UpdateCommunityTtsEngineRequestParam
	paramObj
}

func (r TtsIntegrationReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCommunityTtsEngineRequest)
}
func (r *TtsIntegrationReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsIntegrationGetConformanceRunsParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page TtsIntegrationGetConformanceRunsParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationGetConformanceRunsParams]'s query parameters
// as `url.Values`.
func (r TtsIntegrationGetConformanceRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type TtsIntegrationGetConformanceRunsParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Clamped to 1-100.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationGetConformanceRunsParamsPage]'s query
// parameters as `url.Values`.
func (r TtsIntegrationGetConformanceRunsParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TtsIntegrationGetDiscoverableParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page TtsIntegrationGetDiscoverableParamsPage `query:"page,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "draft", "active", "disabled", "suspended".
	Status CommunityTtsEngineStatus `query:"status,omitzero" json:"-"`
	// Filter by visibility.
	//
	// Any of "private", "public".
	Visibility CommunityTtsEngineVisibility `query:"visibility,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationGetDiscoverableParams]'s query parameters as
// `url.Values`.
func (r TtsIntegrationGetDiscoverableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type TtsIntegrationGetDiscoverableParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Clamped to 1-100.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TtsIntegrationGetDiscoverableParamsPage]'s query parameters
// as `url.Values`.
func (r TtsIntegrationGetDiscoverableParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
