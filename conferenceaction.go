// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// ConferenceActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConferenceActionService] method instead.
type ConferenceActionService struct {
	Options []option.RequestOption
}

// NewConferenceActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConferenceActionService(opts ...option.RequestOption) (r ConferenceActionService) {
	r = ConferenceActionService{}
	r.Options = opts
	return
}

// Update conference participant supervisor_role
func (r *ConferenceActionService) Update(ctx context.Context, id string, body ConferenceActionUpdateParams, opts ...option.RequestOption) (res *ConferenceActionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/update", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Hold a list of participants in a conference call
func (r *ConferenceActionService) Hold(ctx context.Context, id string, body ConferenceActionHoldParams, opts ...option.RequestOption) (res *ConferenceActionHoldResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/hold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Join an existing call leg to a conference. Issue the Join Conference command
// with the conference ID in the path and the `call_control_id` of the leg you wish
// to join to the conference as an attribute. The conference can have up to a
// certain amount of active participants, as set by the `max_participants`
// parameter in conference creation request.
//
// **Expected Webhooks:**
//
// - `conference.participant.joined`
// - `conference.participant.left`
func (r *ConferenceActionService) Join(ctx context.Context, id string, body ConferenceActionJoinParams, opts ...option.RequestOption) (res *ConferenceActionJoinResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/join", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a call leg from a conference and moves it back to parked state.
//
// **Expected Webhooks:**
//
// - `conference.participant.left`
func (r *ConferenceActionService) Leave(ctx context.Context, id string, body ConferenceActionLeaveParams, opts ...option.RequestOption) (res *ConferenceActionLeaveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/leave", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Mute a list of participants in a conference call
func (r *ConferenceActionService) Mute(ctx context.Context, id string, body ConferenceActionMuteParams, opts ...option.RequestOption) (res *ConferenceActionMuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/mute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Play audio to all or some participants on a conference call.
func (r *ConferenceActionService) Play(ctx context.Context, id string, body ConferenceActionPlayParams, opts ...option.RequestOption) (res *ConferenceActionPlayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/play", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Pause conference recording.
func (r *ConferenceActionService) RecordPause(ctx context.Context, id string, body ConferenceActionRecordPauseParams, opts ...option.RequestOption) (res *ConferenceActionRecordPauseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/record_pause", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Resume conference recording.
func (r *ConferenceActionService) RecordResume(ctx context.Context, id string, body ConferenceActionRecordResumeParams, opts ...option.RequestOption) (res *ConferenceActionRecordResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/record_resume", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Start recording the conference. Recording will stop on conference end, or via
// the Stop Recording command.
//
// **Expected Webhooks:**
//
// - `conference.recording.saved`
func (r *ConferenceActionService) RecordStart(ctx context.Context, id string, body ConferenceActionRecordStartParams, opts ...option.RequestOption) (res *ConferenceActionRecordStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/record_start", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop recording the conference.
//
// **Expected Webhooks:**
//
// - `conference.recording.saved`
func (r *ConferenceActionService) RecordStop(ctx context.Context, id string, body ConferenceActionRecordStopParams, opts ...option.RequestOption) (res *ConferenceActionRecordStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/record_stop", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Convert text to speech and play it to all or some participants.
func (r *ConferenceActionService) Speak(ctx context.Context, id string, body ConferenceActionSpeakParams, opts ...option.RequestOption) (res *ConferenceActionSpeakResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/speak", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Stop audio being played to all or some participants on a conference call.
func (r *ConferenceActionService) Stop(ctx context.Context, id string, body ConferenceActionStopParams, opts ...option.RequestOption) (res *ConferenceActionStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/stop", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unhold a list of participants in a conference call
func (r *ConferenceActionService) Unhold(ctx context.Context, id string, body ConferenceActionUnholdParams, opts ...option.RequestOption) (res *ConferenceActionUnholdResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/unhold", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unmute a list of participants in a conference call
func (r *ConferenceActionService) Unmute(ctx context.Context, id string, body ConferenceActionUnmuteParams, opts ...option.RequestOption) (res *ConferenceActionUnmuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/actions/unmute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConferenceCommandResult struct {
	Result string `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceCommandResult) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCommandResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallControlID, SupervisorRole are required.
type UpdateConferenceParam struct {
	// Unique identifier and token for controlling the call
	CallControlID string `json:"call_control_id,required"`
	// Sets the participant as a supervisor for the conference. A conference can have
	// multiple supervisors. "barge" means the supervisor enters the conference as a
	// normal participant. This is the same as "none". "monitor" means the supervisor
	// is muted but can hear all participants. "whisper" means that only the specified
	// "whisper_call_control_ids" can hear the supervisor. Defaults to "none".
	//
	// Any of "barge", "monitor", "none", "whisper".
	SupervisorRole UpdateConferenceSupervisorRole `json:"supervisor_role,omitzero,required"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region UpdateConferenceRegion `json:"region,omitzero"`
	// Array of unique call_control_ids the supervisor can whisper to. If none
	// provided, the supervisor will join the conference as a monitoring participant
	// only.
	WhisperCallControlIDs []string `json:"whisper_call_control_ids,omitzero"`
	paramObj
}

func (r UpdateConferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateConferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateConferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the participant as a supervisor for the conference. A conference can have
// multiple supervisors. "barge" means the supervisor enters the conference as a
// normal participant. This is the same as "none". "monitor" means the supervisor
// is muted but can hear all participants. "whisper" means that only the specified
// "whisper_call_control_ids" can hear the supervisor. Defaults to "none".
type UpdateConferenceSupervisorRole string

const (
	UpdateConferenceSupervisorRoleBarge   UpdateConferenceSupervisorRole = "barge"
	UpdateConferenceSupervisorRoleMonitor UpdateConferenceSupervisorRole = "monitor"
	UpdateConferenceSupervisorRoleNone    UpdateConferenceSupervisorRole = "none"
	UpdateConferenceSupervisorRoleWhisper UpdateConferenceSupervisorRole = "whisper"
)

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type UpdateConferenceRegion string

const (
	UpdateConferenceRegionAustralia  UpdateConferenceRegion = "Australia"
	UpdateConferenceRegionEurope     UpdateConferenceRegion = "Europe"
	UpdateConferenceRegionMiddleEast UpdateConferenceRegion = "Middle East"
	UpdateConferenceRegionUs         UpdateConferenceRegion = "US"
)

type ConferenceActionUpdateResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionHoldResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionHoldResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionHoldResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionJoinResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionJoinResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionJoinResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionLeaveResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionLeaveResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionLeaveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionMuteResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionMuteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionMuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionPlayResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionPlayResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionPlayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionRecordPauseResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionRecordPauseResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionRecordPauseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionRecordResumeResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionRecordResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionRecordResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionRecordStartResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionRecordStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionRecordStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionRecordStopResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionRecordStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionRecordStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionSpeakResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionSpeakResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionSpeakResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionStopResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionUnholdResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionUnholdResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionUnholdResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionUnmuteResponse struct {
	Data ConferenceCommandResult `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceActionUnmuteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceActionUnmuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceActionUpdateParams struct {
	UpdateConference UpdateConferenceParam
	paramObj
}

func (r ConferenceActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateConference)
}
func (r *ConferenceActionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateConference)
}

type ConferenceActionHoldParams struct {
	// The URL of a file to be played to the participants when they are put on hold.
	// media_name and audio_url cannot be used together in one request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// The media_name of a file to be played to the participants when they are put on
	// hold. The media_name must point to a file previously uploaded to
	// api.telnyx.com/v2/media by the same user/organization. The file must either be a
	// WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// List of unique identifiers and tokens for controlling the call. When empty all
	// participants will be placed on hold.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionHoldParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionHoldParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionHoldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionHoldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionHoldParamsRegion string

const (
	ConferenceActionHoldParamsRegionAustralia  ConferenceActionHoldParamsRegion = "Australia"
	ConferenceActionHoldParamsRegionEurope     ConferenceActionHoldParamsRegion = "Europe"
	ConferenceActionHoldParamsRegionMiddleEast ConferenceActionHoldParamsRegion = "Middle East"
	ConferenceActionHoldParamsRegionUs         ConferenceActionHoldParamsRegion = "US"
)

type ConferenceActionJoinParams struct {
	// Unique identifier and token for controlling the call
	CallControlID string `json:"call_control_id,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string. Please note that the client_state will be updated for
	// the participient call leg and the change will not affect conferencing webhooks
	// unless the participient is the owner of the conference.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Whether the conference should end and all remaining participants be hung up
	// after the participant leaves the conference. Defaults to "false".
	EndConferenceOnExit param.Opt[bool] `json:"end_conference_on_exit,omitzero"`
	// Whether the participant should be put on hold immediately after joining the
	// conference. Defaults to "false".
	Hold param.Opt[bool] `json:"hold,omitzero"`
	// The URL of a file to be played to the participant when they are put on hold
	// after joining the conference. hold_media_name and hold_audio_url cannot be used
	// together in one request. Takes effect only when "start_conference_on_create" is
	// set to "false". This property takes effect only if "hold" is set to "true".
	HoldAudioURL param.Opt[string] `json:"hold_audio_url,omitzero"`
	// The media_name of a file to be played to the participant when they are put on
	// hold after joining the conference. The media_name must point to a file
	// previously uploaded to api.telnyx.com/v2/media by the same user/organization.
	// The file must either be a WAV or MP3 file. Takes effect only when
	// "start_conference_on_create" is set to "false". This property takes effect only
	// if "hold" is set to "true".
	HoldMediaName param.Opt[string] `json:"hold_media_name,omitzero"`
	// Whether the participant should be muted immediately after joining the
	// conference. Defaults to "false".
	Mute param.Opt[bool] `json:"mute,omitzero"`
	// Whether the conference should end after the participant leaves the conference.
	// NOTE this doesn't hang up the other participants. Defaults to "false".
	SoftEndConferenceOnExit param.Opt[bool] `json:"soft_end_conference_on_exit,omitzero"`
	// Whether the conference should be started after the participant joins the
	// conference. Defaults to "false".
	StartConferenceOnEnter param.Opt[bool] `json:"start_conference_on_enter,omitzero"`
	// Whether a beep sound should be played when the participant joins and/or leaves
	// the conference. Can be used to override the conference-level setting.
	//
	// Any of "always", "never", "on_enter", "on_exit".
	BeepEnabled ConferenceActionJoinParamsBeepEnabled `json:"beep_enabled,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionJoinParamsRegion `json:"region,omitzero"`
	// Sets the joining participant as a supervisor for the conference. A conference
	// can have multiple supervisors. "barge" means the supervisor enters the
	// conference as a normal participant. This is the same as "none". "monitor" means
	// the supervisor is muted but can hear all participants. "whisper" means that only
	// the specified "whisper_call_control_ids" can hear the supervisor. Defaults to
	// "none".
	//
	// Any of "barge", "monitor", "none", "whisper".
	SupervisorRole ConferenceActionJoinParamsSupervisorRole `json:"supervisor_role,omitzero"`
	// Array of unique call_control_ids the joining supervisor can whisper to. If none
	// provided, the supervisor will join the conference as a monitoring participant
	// only.
	WhisperCallControlIDs []string `json:"whisper_call_control_ids,omitzero"`
	paramObj
}

func (r ConferenceActionJoinParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionJoinParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionJoinParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a beep sound should be played when the participant joins and/or leaves
// the conference. Can be used to override the conference-level setting.
type ConferenceActionJoinParamsBeepEnabled string

const (
	ConferenceActionJoinParamsBeepEnabledAlways  ConferenceActionJoinParamsBeepEnabled = "always"
	ConferenceActionJoinParamsBeepEnabledNever   ConferenceActionJoinParamsBeepEnabled = "never"
	ConferenceActionJoinParamsBeepEnabledOnEnter ConferenceActionJoinParamsBeepEnabled = "on_enter"
	ConferenceActionJoinParamsBeepEnabledOnExit  ConferenceActionJoinParamsBeepEnabled = "on_exit"
)

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionJoinParamsRegion string

const (
	ConferenceActionJoinParamsRegionAustralia  ConferenceActionJoinParamsRegion = "Australia"
	ConferenceActionJoinParamsRegionEurope     ConferenceActionJoinParamsRegion = "Europe"
	ConferenceActionJoinParamsRegionMiddleEast ConferenceActionJoinParamsRegion = "Middle East"
	ConferenceActionJoinParamsRegionUs         ConferenceActionJoinParamsRegion = "US"
)

// Sets the joining participant as a supervisor for the conference. A conference
// can have multiple supervisors. "barge" means the supervisor enters the
// conference as a normal participant. This is the same as "none". "monitor" means
// the supervisor is muted but can hear all participants. "whisper" means that only
// the specified "whisper_call_control_ids" can hear the supervisor. Defaults to
// "none".
type ConferenceActionJoinParamsSupervisorRole string

const (
	ConferenceActionJoinParamsSupervisorRoleBarge   ConferenceActionJoinParamsSupervisorRole = "barge"
	ConferenceActionJoinParamsSupervisorRoleMonitor ConferenceActionJoinParamsSupervisorRole = "monitor"
	ConferenceActionJoinParamsSupervisorRoleNone    ConferenceActionJoinParamsSupervisorRole = "none"
	ConferenceActionJoinParamsSupervisorRoleWhisper ConferenceActionJoinParamsSupervisorRole = "whisper"
)

type ConferenceActionLeaveParams struct {
	// Unique identifier and token for controlling the call
	CallControlID string `json:"call_control_id,required"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Whether a beep sound should be played when the participant leaves the
	// conference. Can be used to override the conference-level setting.
	//
	// Any of "always", "never", "on_enter", "on_exit".
	BeepEnabled ConferenceActionLeaveParamsBeepEnabled `json:"beep_enabled,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionLeaveParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionLeaveParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionLeaveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionLeaveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a beep sound should be played when the participant leaves the
// conference. Can be used to override the conference-level setting.
type ConferenceActionLeaveParamsBeepEnabled string

const (
	ConferenceActionLeaveParamsBeepEnabledAlways  ConferenceActionLeaveParamsBeepEnabled = "always"
	ConferenceActionLeaveParamsBeepEnabledNever   ConferenceActionLeaveParamsBeepEnabled = "never"
	ConferenceActionLeaveParamsBeepEnabledOnEnter ConferenceActionLeaveParamsBeepEnabled = "on_enter"
	ConferenceActionLeaveParamsBeepEnabledOnExit  ConferenceActionLeaveParamsBeepEnabled = "on_exit"
)

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionLeaveParamsRegion string

const (
	ConferenceActionLeaveParamsRegionAustralia  ConferenceActionLeaveParamsRegion = "Australia"
	ConferenceActionLeaveParamsRegionEurope     ConferenceActionLeaveParamsRegion = "Europe"
	ConferenceActionLeaveParamsRegionMiddleEast ConferenceActionLeaveParamsRegion = "Middle East"
	ConferenceActionLeaveParamsRegionUs         ConferenceActionLeaveParamsRegion = "US"
)

type ConferenceActionMuteParams struct {
	// Array of unique identifiers and tokens for controlling the call. When empty all
	// participants will be muted.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionMuteParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionMuteParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionMuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionMuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionMuteParamsRegion string

const (
	ConferenceActionMuteParamsRegionAustralia  ConferenceActionMuteParamsRegion = "Australia"
	ConferenceActionMuteParamsRegionEurope     ConferenceActionMuteParamsRegion = "Europe"
	ConferenceActionMuteParamsRegionMiddleEast ConferenceActionMuteParamsRegion = "Middle East"
	ConferenceActionMuteParamsRegionUs         ConferenceActionMuteParamsRegion = "US"
)

type ConferenceActionPlayParams struct {
	// The URL of a file to be played back in the conference. media_name and audio_url
	// cannot be used together in one request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// The media_name of a file to be played back in the conference. The media_name
	// must point to a file previously uploaded to api.telnyx.com/v2/media by the same
	// user/organization. The file must either be a WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// List of call control ids identifying participants the audio file should be
	// played to. If not given, the audio file will be played to the entire conference.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// The number of times the audio file should be played. If supplied, the value must
	// be an integer between 1 and 100, or the special string `infinity` for an endless
	// loop.
	Loop LoopcountUnionParam `json:"loop,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionPlayParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionPlayParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionPlayParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionPlayParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionPlayParamsRegion string

const (
	ConferenceActionPlayParamsRegionAustralia  ConferenceActionPlayParamsRegion = "Australia"
	ConferenceActionPlayParamsRegionEurope     ConferenceActionPlayParamsRegion = "Europe"
	ConferenceActionPlayParamsRegionMiddleEast ConferenceActionPlayParamsRegion = "Middle East"
	ConferenceActionPlayParamsRegionUs         ConferenceActionPlayParamsRegion = "US"
)

type ConferenceActionRecordPauseParams struct {
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Use this field to pause specific recording.
	RecordingID param.Opt[string] `json:"recording_id,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionRecordPauseParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionRecordPauseParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionRecordPauseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionRecordPauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionRecordPauseParamsRegion string

const (
	ConferenceActionRecordPauseParamsRegionAustralia  ConferenceActionRecordPauseParamsRegion = "Australia"
	ConferenceActionRecordPauseParamsRegionEurope     ConferenceActionRecordPauseParamsRegion = "Europe"
	ConferenceActionRecordPauseParamsRegionMiddleEast ConferenceActionRecordPauseParamsRegion = "Middle East"
	ConferenceActionRecordPauseParamsRegionUs         ConferenceActionRecordPauseParamsRegion = "US"
)

type ConferenceActionRecordResumeParams struct {
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Use this field to resume specific recording.
	RecordingID param.Opt[string] `json:"recording_id,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionRecordResumeParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionRecordResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionRecordResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionRecordResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionRecordResumeParamsRegion string

const (
	ConferenceActionRecordResumeParamsRegionAustralia  ConferenceActionRecordResumeParamsRegion = "Australia"
	ConferenceActionRecordResumeParamsRegionEurope     ConferenceActionRecordResumeParamsRegion = "Europe"
	ConferenceActionRecordResumeParamsRegionMiddleEast ConferenceActionRecordResumeParamsRegion = "Middle East"
	ConferenceActionRecordResumeParamsRegionUs         ConferenceActionRecordResumeParamsRegion = "US"
)

type ConferenceActionRecordStartParams struct {
	// The audio file format used when storing the conference recording. Can be either
	// `mp3` or `wav`.
	//
	// Any of "wav", "mp3".
	Format ConferenceActionRecordStartParamsFormat `json:"format,omitzero,required"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `conference_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// The custom recording file name to be used instead of the default `call_leg_id`.
	// Telnyx will still add a Unix timestamp suffix.
	CustomFileName param.Opt[string] `json:"custom_file_name,omitzero"`
	// If enabled, a beep sound will be played at the start of a recording.
	PlayBeep param.Opt[bool] `json:"play_beep,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionRecordStartParamsRegion `json:"region,omitzero"`
	// When set to `trim-silence`, silence will be removed from the beginning and end
	// of the recording.
	//
	// Any of "trim-silence".
	Trim ConferenceActionRecordStartParamsTrim `json:"trim,omitzero"`
	paramObj
}

func (r ConferenceActionRecordStartParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionRecordStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionRecordStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The audio file format used when storing the conference recording. Can be either
// `mp3` or `wav`.
type ConferenceActionRecordStartParamsFormat string

const (
	ConferenceActionRecordStartParamsFormatWav ConferenceActionRecordStartParamsFormat = "wav"
	ConferenceActionRecordStartParamsFormatMP3 ConferenceActionRecordStartParamsFormat = "mp3"
)

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionRecordStartParamsRegion string

const (
	ConferenceActionRecordStartParamsRegionAustralia  ConferenceActionRecordStartParamsRegion = "Australia"
	ConferenceActionRecordStartParamsRegionEurope     ConferenceActionRecordStartParamsRegion = "Europe"
	ConferenceActionRecordStartParamsRegionMiddleEast ConferenceActionRecordStartParamsRegion = "Middle East"
	ConferenceActionRecordStartParamsRegionUs         ConferenceActionRecordStartParamsRegion = "US"
)

// When set to `trim-silence`, silence will be removed from the beginning and end
// of the recording.
type ConferenceActionRecordStartParamsTrim string

const (
	ConferenceActionRecordStartParamsTrimTrimSilence ConferenceActionRecordStartParamsTrim = "trim-silence"
)

type ConferenceActionRecordStopParams struct {
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with
	// the same `command_id` for the same `call_control_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Uniquely identifies the resource.
	RecordingID param.Opt[string] `json:"recording_id,omitzero" format:"uuid"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionRecordStopParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionRecordStopParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionRecordStopParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionRecordStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionRecordStopParamsRegion string

const (
	ConferenceActionRecordStopParamsRegionAustralia  ConferenceActionRecordStopParamsRegion = "Australia"
	ConferenceActionRecordStopParamsRegionEurope     ConferenceActionRecordStopParamsRegion = "Europe"
	ConferenceActionRecordStopParamsRegionMiddleEast ConferenceActionRecordStopParamsRegion = "Middle East"
	ConferenceActionRecordStopParamsRegionUs         ConferenceActionRecordStopParamsRegion = "US"
)

type ConferenceActionSpeakParams struct {
	// The text or SSML to be converted into speech. There is a 3,000 character limit.
	Payload string `json:"payload,required"`
	// Specifies the voice used in speech synthesis.
	//
	//   - Define voices using the format `<Provider>.<Model>.<VoiceId>`. Specifying only
	//     the provider will give default values for voice_id and model_id.
	//
	//     **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM`). The `ModelId` part
	//     is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an
	//     integration identifier secret in
	//     `"voice_settings": {"api_key_ref": "<secret_identifier>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	//
	// For service_level basic, you may define the gender of the speaker (male or
	// female).
	Voice string `json:"voice,required"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Call Control IDs of participants who will hear the spoken text. When empty all
	// participants will hear the spoken text.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// The language you want spoken. This parameter is ignored when a `Polly.*` voice
	// is specified.
	//
	// Any of "arb", "cmn-CN", "cy-GB", "da-DK", "de-DE", "en-AU", "en-GB",
	// "en-GB-WLS", "en-IN", "en-US", "es-ES", "es-MX", "es-US", "fr-CA", "fr-FR",
	// "hi-IN", "is-IS", "it-IT", "ja-JP", "ko-KR", "nb-NO", "nl-NL", "pl-PL", "pt-BR",
	// "pt-PT", "ro-RO", "ru-RU", "sv-SE", "tr-TR".
	Language ConferenceActionSpeakParamsLanguage `json:"language,omitzero"`
	// The type of the provided payload. The payload can either be plain text, or
	// Speech Synthesis Markup Language (SSML).
	//
	// Any of "text", "ssml".
	PayloadType ConferenceActionSpeakParamsPayloadType `json:"payload_type,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionSpeakParamsRegion `json:"region,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings ConferenceActionSpeakParamsVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r ConferenceActionSpeakParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionSpeakParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionSpeakParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The language you want spoken. This parameter is ignored when a `Polly.*` voice
// is specified.
type ConferenceActionSpeakParamsLanguage string

const (
	ConferenceActionSpeakParamsLanguageArb     ConferenceActionSpeakParamsLanguage = "arb"
	ConferenceActionSpeakParamsLanguageCmnCn   ConferenceActionSpeakParamsLanguage = "cmn-CN"
	ConferenceActionSpeakParamsLanguageCyGB    ConferenceActionSpeakParamsLanguage = "cy-GB"
	ConferenceActionSpeakParamsLanguageDaDk    ConferenceActionSpeakParamsLanguage = "da-DK"
	ConferenceActionSpeakParamsLanguageDeDe    ConferenceActionSpeakParamsLanguage = "de-DE"
	ConferenceActionSpeakParamsLanguageEnAu    ConferenceActionSpeakParamsLanguage = "en-AU"
	ConferenceActionSpeakParamsLanguageEnGB    ConferenceActionSpeakParamsLanguage = "en-GB"
	ConferenceActionSpeakParamsLanguageEnGBWls ConferenceActionSpeakParamsLanguage = "en-GB-WLS"
	ConferenceActionSpeakParamsLanguageEnIn    ConferenceActionSpeakParamsLanguage = "en-IN"
	ConferenceActionSpeakParamsLanguageEnUs    ConferenceActionSpeakParamsLanguage = "en-US"
	ConferenceActionSpeakParamsLanguageEsEs    ConferenceActionSpeakParamsLanguage = "es-ES"
	ConferenceActionSpeakParamsLanguageEsMx    ConferenceActionSpeakParamsLanguage = "es-MX"
	ConferenceActionSpeakParamsLanguageEsUs    ConferenceActionSpeakParamsLanguage = "es-US"
	ConferenceActionSpeakParamsLanguageFrCa    ConferenceActionSpeakParamsLanguage = "fr-CA"
	ConferenceActionSpeakParamsLanguageFrFr    ConferenceActionSpeakParamsLanguage = "fr-FR"
	ConferenceActionSpeakParamsLanguageHiIn    ConferenceActionSpeakParamsLanguage = "hi-IN"
	ConferenceActionSpeakParamsLanguageIsIs    ConferenceActionSpeakParamsLanguage = "is-IS"
	ConferenceActionSpeakParamsLanguageItIt    ConferenceActionSpeakParamsLanguage = "it-IT"
	ConferenceActionSpeakParamsLanguageJaJp    ConferenceActionSpeakParamsLanguage = "ja-JP"
	ConferenceActionSpeakParamsLanguageKoKr    ConferenceActionSpeakParamsLanguage = "ko-KR"
	ConferenceActionSpeakParamsLanguageNbNo    ConferenceActionSpeakParamsLanguage = "nb-NO"
	ConferenceActionSpeakParamsLanguageNlNl    ConferenceActionSpeakParamsLanguage = "nl-NL"
	ConferenceActionSpeakParamsLanguagePlPl    ConferenceActionSpeakParamsLanguage = "pl-PL"
	ConferenceActionSpeakParamsLanguagePtBr    ConferenceActionSpeakParamsLanguage = "pt-BR"
	ConferenceActionSpeakParamsLanguagePtPt    ConferenceActionSpeakParamsLanguage = "pt-PT"
	ConferenceActionSpeakParamsLanguageRoRo    ConferenceActionSpeakParamsLanguage = "ro-RO"
	ConferenceActionSpeakParamsLanguageRuRu    ConferenceActionSpeakParamsLanguage = "ru-RU"
	ConferenceActionSpeakParamsLanguageSvSe    ConferenceActionSpeakParamsLanguage = "sv-SE"
	ConferenceActionSpeakParamsLanguageTrTr    ConferenceActionSpeakParamsLanguage = "tr-TR"
)

// The type of the provided payload. The payload can either be plain text, or
// Speech Synthesis Markup Language (SSML).
type ConferenceActionSpeakParamsPayloadType string

const (
	ConferenceActionSpeakParamsPayloadTypeText ConferenceActionSpeakParamsPayloadType = "text"
	ConferenceActionSpeakParamsPayloadTypeSsml ConferenceActionSpeakParamsPayloadType = "ssml"
)

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionSpeakParamsRegion string

const (
	ConferenceActionSpeakParamsRegionAustralia  ConferenceActionSpeakParamsRegion = "Australia"
	ConferenceActionSpeakParamsRegionEurope     ConferenceActionSpeakParamsRegion = "Europe"
	ConferenceActionSpeakParamsRegionMiddleEast ConferenceActionSpeakParamsRegion = "Middle East"
	ConferenceActionSpeakParamsRegionUs         ConferenceActionSpeakParamsRegion = "US"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConferenceActionSpeakParamsVoiceSettingsUnion struct {
	OfElevenLabsVoiceSettings *ElevenLabsVoiceSettingsParam `json:",omitzero,inline"`
	OfTelnyxVoiceSettings     *TelnyxVoiceSettingsParam     `json:",omitzero,inline"`
	paramUnion
}

func (u ConferenceActionSpeakParamsVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenLabsVoiceSettings, u.OfTelnyxVoiceSettings)
}
func (u *ConferenceActionSpeakParamsVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConferenceActionSpeakParamsVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenLabsVoiceSettings) {
		return u.OfElevenLabsVoiceSettings
	} else if !param.IsOmitted(u.OfTelnyxVoiceSettings) {
		return u.OfTelnyxVoiceSettings
	}
	return nil
}

type ConferenceActionStopParams struct {
	// List of call control ids identifying participants the audio file should stop be
	// played to. If not given, the audio will be stoped to the entire conference.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionStopParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionStopParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionStopParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionStopParamsRegion string

const (
	ConferenceActionStopParamsRegionAustralia  ConferenceActionStopParamsRegion = "Australia"
	ConferenceActionStopParamsRegionEurope     ConferenceActionStopParamsRegion = "Europe"
	ConferenceActionStopParamsRegionMiddleEast ConferenceActionStopParamsRegion = "Middle East"
	ConferenceActionStopParamsRegionUs         ConferenceActionStopParamsRegion = "US"
)

type ConferenceActionUnholdParams struct {
	// List of unique identifiers and tokens for controlling the call. Enter each call
	// control ID to be unheld.
	CallControlIDs []string `json:"call_control_ids,omitzero,required"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionUnholdParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionUnholdParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionUnholdParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionUnholdParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionUnholdParamsRegion string

const (
	ConferenceActionUnholdParamsRegionAustralia  ConferenceActionUnholdParamsRegion = "Australia"
	ConferenceActionUnholdParamsRegionEurope     ConferenceActionUnholdParamsRegion = "Europe"
	ConferenceActionUnholdParamsRegionMiddleEast ConferenceActionUnholdParamsRegion = "Middle East"
	ConferenceActionUnholdParamsRegionUs         ConferenceActionUnholdParamsRegion = "US"
)

type ConferenceActionUnmuteParams struct {
	// List of unique identifiers and tokens for controlling the call. Enter each call
	// control ID to be unmuted. When empty all participants will be unmuted.
	CallControlIDs []string `json:"call_control_ids,omitzero"`
	// Region where the conference data is located. Defaults to the region defined in
	// user's data locality settings (Europe or US).
	//
	// Any of "Australia", "Europe", "Middle East", "US".
	Region ConferenceActionUnmuteParamsRegion `json:"region,omitzero"`
	paramObj
}

func (r ConferenceActionUnmuteParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceActionUnmuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceActionUnmuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region where the conference data is located. Defaults to the region defined in
// user's data locality settings (Europe or US).
type ConferenceActionUnmuteParamsRegion string

const (
	ConferenceActionUnmuteParamsRegionAustralia  ConferenceActionUnmuteParamsRegion = "Australia"
	ConferenceActionUnmuteParamsRegionEurope     ConferenceActionUnmuteParamsRegion = "Europe"
	ConferenceActionUnmuteParamsRegionMiddleEast ConferenceActionUnmuteParamsRegion = "Middle East"
	ConferenceActionUnmuteParamsRegionUs         ConferenceActionUnmuteParamsRegion = "US"
)
