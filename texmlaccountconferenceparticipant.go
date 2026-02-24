// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// TexmlAccountConferenceParticipantService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountConferenceParticipantService] method instead.
type TexmlAccountConferenceParticipantService struct {
	Options []option.RequestOption
}

// NewTexmlAccountConferenceParticipantService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountConferenceParticipantService(opts ...option.RequestOption) (r TexmlAccountConferenceParticipantService) {
	r = TexmlAccountConferenceParticipantService{}
	r.Options = opts
	return
}

// Gets conference participant resource
func (r *TexmlAccountConferenceParticipantService) Get(ctx context.Context, callSidOrParticipantLabel string, query TexmlAccountConferenceParticipantGetParams, opts ...option.RequestOption) (res *TexmlAccountConferenceParticipantGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if query.ConferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	if callSidOrParticipantLabel == "" {
		err = errors.New("missing required call_sid_or_participant_label parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Participants/%s", query.AccountSid, query.ConferenceSid, callSidOrParticipantLabel)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a conference participant
func (r *TexmlAccountConferenceParticipantService) Update(ctx context.Context, callSidOrParticipantLabel string, params TexmlAccountConferenceParticipantUpdateParams, opts ...option.RequestOption) (res *TexmlAccountConferenceParticipantUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if params.ConferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	if callSidOrParticipantLabel == "" {
		err = errors.New("missing required call_sid_or_participant_label parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Participants/%s", params.AccountSid, params.ConferenceSid, callSidOrParticipantLabel)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Deletes a conference participant
func (r *TexmlAccountConferenceParticipantService) Delete(ctx context.Context, callSidOrParticipantLabel string, body TexmlAccountConferenceParticipantDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if body.ConferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	if callSidOrParticipantLabel == "" {
		err = errors.New("missing required call_sid_or_participant_label parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Participants/%s", body.AccountSid, body.ConferenceSid, callSidOrParticipantLabel)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Dials a new conference participant
func (r *TexmlAccountConferenceParticipantService) Participants(ctx context.Context, conferenceSid string, params TexmlAccountConferenceParticipantParticipantsParams, opts ...option.RequestOption) (res *TexmlAccountConferenceParticipantParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Participants", params.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists conference participants
func (r *TexmlAccountConferenceParticipantService) GetParticipants(ctx context.Context, conferenceSid string, query TexmlAccountConferenceParticipantGetParticipantsParams, opts ...option.RequestOption) (res *TexmlAccountConferenceParticipantGetParticipantsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Participants", query.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TexmlAccountConferenceParticipantGetResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// The identifier of this participant's call.
	CallSid string `json:"call_sid"`
	// The identifier of this participant's call.
	CallSidLegacy string `json:"call_sid_legacy"`
	// Whether the participant is coaching another call.
	Coaching bool `json:"coaching"`
	// The identifier of the coached participant's call.
	CoachingCallSid string `json:"coaching_call_sid"`
	// The identifier of the coached participant's call.
	CoachingCallSidLegacy string `json:"coaching_call_sid_legacy"`
	// The unique identifier for the conference.
	ConferenceSid string `json:"conference_sid" format:"uuid"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// Whether the conference ends when the participant leaves.
	EndConferenceOnExit bool `json:"end_conference_on_exit"`
	// Whether the participant is on hold.
	Hold bool `json:"hold"`
	// Whether the participant is muted.
	Muted bool `json:"muted"`
	// The status of the participant's call in the conference.
	//
	// Any of "connecting", "connected", "completed".
	Status TexmlAccountConferenceParticipantGetResponseStatus `json:"status"`
	// The relative URI for this participant.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid            respjson.Field
		APIVersion            respjson.Field
		CallSid               respjson.Field
		CallSidLegacy         respjson.Field
		Coaching              respjson.Field
		CoachingCallSid       respjson.Field
		CoachingCallSidLegacy respjson.Field
		ConferenceSid         respjson.Field
		DateCreated           respjson.Field
		DateUpdated           respjson.Field
		EndConferenceOnExit   respjson.Field
		Hold                  respjson.Field
		Muted                 respjson.Field
		Status                respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceParticipantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceParticipantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the participant's call in the conference.
type TexmlAccountConferenceParticipantGetResponseStatus string

const (
	TexmlAccountConferenceParticipantGetResponseStatusConnecting TexmlAccountConferenceParticipantGetResponseStatus = "connecting"
	TexmlAccountConferenceParticipantGetResponseStatusConnected  TexmlAccountConferenceParticipantGetResponseStatus = "connected"
	TexmlAccountConferenceParticipantGetResponseStatusCompleted  TexmlAccountConferenceParticipantGetResponseStatus = "completed"
)

type TexmlAccountConferenceParticipantUpdateResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// The identifier of this participant's call.
	CallSid string `json:"call_sid"`
	// The identifier of this participant's call.
	CallSidLegacy string `json:"call_sid_legacy"`
	// Whether the participant is coaching another call.
	Coaching bool `json:"coaching"`
	// The identifier of the coached participant's call.
	CoachingCallSid string `json:"coaching_call_sid"`
	// The identifier of the coached participant's call.
	CoachingCallSidLegacy string `json:"coaching_call_sid_legacy"`
	// The unique identifier for the conference.
	ConferenceSid string `json:"conference_sid" format:"uuid"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// Whether the conference ends when the participant leaves.
	EndConferenceOnExit bool `json:"end_conference_on_exit"`
	// Whether the participant is on hold.
	Hold bool `json:"hold"`
	// Whether the participant is muted.
	Muted bool `json:"muted"`
	// The status of the participant's call in the conference.
	//
	// Any of "connecting", "connected", "completed".
	Status TexmlAccountConferenceParticipantUpdateResponseStatus `json:"status"`
	// The relative URI for this participant.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid            respjson.Field
		APIVersion            respjson.Field
		CallSid               respjson.Field
		CallSidLegacy         respjson.Field
		Coaching              respjson.Field
		CoachingCallSid       respjson.Field
		CoachingCallSidLegacy respjson.Field
		ConferenceSid         respjson.Field
		DateCreated           respjson.Field
		DateUpdated           respjson.Field
		EndConferenceOnExit   respjson.Field
		Hold                  respjson.Field
		Muted                 respjson.Field
		Status                respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceParticipantUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceParticipantUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the participant's call in the conference.
type TexmlAccountConferenceParticipantUpdateResponseStatus string

const (
	TexmlAccountConferenceParticipantUpdateResponseStatusConnecting TexmlAccountConferenceParticipantUpdateResponseStatus = "connecting"
	TexmlAccountConferenceParticipantUpdateResponseStatusConnected  TexmlAccountConferenceParticipantUpdateResponseStatus = "connected"
	TexmlAccountConferenceParticipantUpdateResponseStatusCompleted  TexmlAccountConferenceParticipantUpdateResponseStatus = "completed"
)

type TexmlAccountConferenceParticipantParticipantsResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The identifier of this participant's call.
	CallSid string `json:"call_sid"`
	// Whether the participant is coaching another call.
	Coaching bool `json:"coaching"`
	// The identifier of the coached participant's call.
	CoachingCallSid string `json:"coaching_call_sid"`
	// The unique identifier for the conference.
	ConferenceSid string `json:"conference_sid" format:"uuid"`
	// Whether the conference ends when the participant leaves.
	EndConferenceOnExit bool `json:"end_conference_on_exit"`
	// Whether the participant is on hold.
	Hold bool `json:"hold"`
	// Whether the participant is muted.
	Muted bool `json:"muted"`
	// The status of the participant's call in the conference.
	//
	// Any of "connecting", "connected", "completed".
	Status TexmlAccountConferenceParticipantParticipantsResponseStatus `json:"status"`
	// The relative URI for this participant.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid          respjson.Field
		CallSid             respjson.Field
		Coaching            respjson.Field
		CoachingCallSid     respjson.Field
		ConferenceSid       respjson.Field
		EndConferenceOnExit respjson.Field
		Hold                respjson.Field
		Muted               respjson.Field
		Status              respjson.Field
		Uri                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceParticipantParticipantsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceParticipantParticipantsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the participant's call in the conference.
type TexmlAccountConferenceParticipantParticipantsResponseStatus string

const (
	TexmlAccountConferenceParticipantParticipantsResponseStatusConnecting TexmlAccountConferenceParticipantParticipantsResponseStatus = "connecting"
	TexmlAccountConferenceParticipantParticipantsResponseStatusConnected  TexmlAccountConferenceParticipantParticipantsResponseStatus = "connected"
	TexmlAccountConferenceParticipantParticipantsResponseStatusCompleted  TexmlAccountConferenceParticipantParticipantsResponseStatus = "completed"
)

type TexmlAccountConferenceParticipantGetParticipantsResponse struct {
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Participants.json?page=0&pagesize=20
	FirstPageUri string `json:"first_page_uri"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Participants.json?Page=1&PageSize=1&PageToken=MTY4AjgyNDkwNzIxMQ
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize     int64                                                                 `json:"page_size"`
	Participants []TexmlAccountConferenceParticipantGetParticipantsResponseParticipant `json:"participants"`
	// The number of the first element on the page, zero-indexed.
	Start int64 `json:"start"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End          respjson.Field
		FirstPageUri respjson.Field
		NextPageUri  respjson.Field
		Page         respjson.Field
		PageSize     respjson.Field
		Participants respjson.Field
		Start        respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceParticipantGetParticipantsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceParticipantGetParticipantsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceParticipantGetParticipantsResponseParticipant struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// The identifier of this participant's call.
	CallSid string `json:"call_sid"`
	// The identifier of this participant's call.
	CallSidLegacy string `json:"call_sid_legacy"`
	// Whether the participant is coaching another call.
	Coaching bool `json:"coaching"`
	// The identifier of the coached participant's call.
	CoachingCallSid string `json:"coaching_call_sid"`
	// The identifier of the coached participant's call.
	CoachingCallSidLegacy string `json:"coaching_call_sid_legacy"`
	// The unique identifier for the conference.
	ConferenceSid string `json:"conference_sid" format:"uuid"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// Whether the conference ends when the participant leaves.
	EndConferenceOnExit bool `json:"end_conference_on_exit"`
	// Whether the participant is on hold.
	Hold bool `json:"hold"`
	// Whether the participant is muted.
	Muted bool `json:"muted"`
	// The status of the participant's call in the conference.
	//
	// Any of "connecting", "connected", "completed".
	Status string `json:"status"`
	// The relative URI for this participant.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid            respjson.Field
		APIVersion            respjson.Field
		CallSid               respjson.Field
		CallSidLegacy         respjson.Field
		Coaching              respjson.Field
		CoachingCallSid       respjson.Field
		CoachingCallSidLegacy respjson.Field
		ConferenceSid         respjson.Field
		DateCreated           respjson.Field
		DateUpdated           respjson.Field
		EndConferenceOnExit   respjson.Field
		Hold                  respjson.Field
		Muted                 respjson.Field
		Status                respjson.Field
		Uri                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceParticipantGetParticipantsResponseParticipant) RawJSON() string {
	return r.JSON.raw
}
func (r *TexmlAccountConferenceParticipantGetParticipantsResponseParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceParticipantGetParams struct {
	AccountSid    string `path:"account_sid,required" json:"-"`
	ConferenceSid string `path:"conference_sid,required" json:"-"`
	paramObj
}

type TexmlAccountConferenceParticipantUpdateParams struct {
	AccountSid    string `path:"account_sid,required" json:"-"`
	ConferenceSid string `path:"conference_sid,required" json:"-"`
	// The URL to call to announce something to the participant. The URL may return an
	// MP3 fileo a WAV file, or a TwiML document that contains `<Play>`, `<Say>`,
	// `<Pause>`, or `<Redirect>` verbs.
	AnnounceURL param.Opt[string] `json:"AnnounceUrl,omitzero"`
	// Whether to play a notification beep to the conference when the participant
	// exits.
	BeepOnExit param.Opt[bool] `json:"BeepOnExit,omitzero"`
	// The SID of the participant who is being coached. The participant being coached
	// is the only participant who can hear the participant who is coaching.
	CallSidToCoach param.Opt[string] `json:"CallSidToCoach,omitzero"`
	// Whether the participant is coaching another call. When `true`, `CallSidToCoach`
	// has to be given.
	Coaching param.Opt[bool] `json:"Coaching,omitzero"`
	// Whether to end the conference when the participant leaves.
	EndConferenceOnExit param.Opt[bool] `json:"EndConferenceOnExit,omitzero"`
	// Whether the participant should be on hold.
	Hold param.Opt[bool] `json:"Hold,omitzero"`
	// The URL to be called using the `HoldMethod` for music that plays when the
	// participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML
	// document that contains `<Play>`, `<Say>`, `<Pause>`, or `<Redirect>` verbs.
	HoldURL param.Opt[string] `json:"HoldUrl,omitzero"`
	// Whether the participant should be muted.
	Muted param.Opt[bool] `json:"Muted,omitzero"`
	// The URL to call for an audio file to play while the participant is waiting for
	// the conference to start.
	WaitURL param.Opt[string] `json:"WaitUrl,omitzero"`
	// The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	AnnounceMethod TexmlAccountConferenceParticipantUpdateParamsAnnounceMethod `json:"AnnounceMethod,omitzero"`
	// The HTTP method to use when calling the `HoldUrl`.
	//
	// Any of "GET", "POST".
	HoldMethod TexmlAccountConferenceParticipantUpdateParamsHoldMethod `json:"HoldMethod,omitzero"`
	paramObj
}

func (r TexmlAccountConferenceParticipantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountConferenceParticipantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountConferenceParticipantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.
type TexmlAccountConferenceParticipantUpdateParamsAnnounceMethod string

const (
	TexmlAccountConferenceParticipantUpdateParamsAnnounceMethodGet  TexmlAccountConferenceParticipantUpdateParamsAnnounceMethod = "GET"
	TexmlAccountConferenceParticipantUpdateParamsAnnounceMethodPost TexmlAccountConferenceParticipantUpdateParamsAnnounceMethod = "POST"
)

// The HTTP method to use when calling the `HoldUrl`.
type TexmlAccountConferenceParticipantUpdateParamsHoldMethod string

const (
	TexmlAccountConferenceParticipantUpdateParamsHoldMethodGet  TexmlAccountConferenceParticipantUpdateParamsHoldMethod = "GET"
	TexmlAccountConferenceParticipantUpdateParamsHoldMethodPost TexmlAccountConferenceParticipantUpdateParamsHoldMethod = "POST"
)

type TexmlAccountConferenceParticipantDeleteParams struct {
	AccountSid    string `path:"account_sid,required" json:"-"`
	ConferenceSid string `path:"conference_sid,required" json:"-"`
	paramObj
}

type TexmlAccountConferenceParticipantParticipantsParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	// The URL the result of answering machine detection will be sent to.
	AmdStatusCallback param.Opt[string] `json:"AmdStatusCallback,omitzero"`
	// The SID of the TeXML application that will handle the new participant's call.
	// Required unless joining an existing conference by its ConferenceSid.
	ApplicationSid param.Opt[string] `json:"ApplicationSid,omitzero"`
	// To be used as the caller id name (SIP From Display Name) presented to the
	// destination (`To` number). The string should have a maximum of 128 characters,
	// containing only letters, numbers, spaces, and `-_~!.+` special characters. If
	// ommited, the display name will be the same as the number in the `From` field.
	CallerID param.Opt[string] `json:"CallerId,omitzero"`
	// The SID of the participant who is being coached. The participant being coached
	// is the only participant who can hear the participant who is coaching.
	CallSidToCoach param.Opt[string] `json:"CallSidToCoach,omitzero"`
	// Whether to cancel ongoing playback on `greeting ended` detection. Defaults to
	// `true`.
	CancelPlaybackOnDetectMessageEnd param.Opt[bool] `json:"CancelPlaybackOnDetectMessageEnd,omitzero"`
	// Whether to cancel ongoing playback on `machine` detection. Defaults to `true`.
	CancelPlaybackOnMachineDetection param.Opt[bool] `json:"CancelPlaybackOnMachineDetection,omitzero"`
	// Whether the participant is coaching another call. When `true`, `CallSidToCoach`
	// has to be given.
	Coaching param.Opt[bool] `json:"Coaching,omitzero"`
	// The URL the conference recording callbacks will be sent to.
	ConferenceRecordingStatusCallback param.Opt[string] `json:"ConferenceRecordingStatusCallback,omitzero"`
	// The changes to the conference recording's state that should generate a call to
	// `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`. `failed` and
	// `absent` are synonymous.
	ConferenceRecordingStatusCallbackEvent param.Opt[string] `json:"ConferenceRecordingStatusCallbackEvent,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected. The timer only starts when the speech is detected. Please
	// note that the transcription is used to detect silence and the related charge
	// will be applied. The minimum value is 0. The default value is 0 (infinite)
	ConferenceRecordingTimeout param.Opt[int64] `json:"ConferenceRecordingTimeout,omitzero"`
	// The URL the conference callbacks will be sent to.
	ConferenceStatusCallback param.Opt[string] `json:"ConferenceStatusCallback,omitzero"`
	// The changes to the conference's state that should generate a call to
	// `ConferenceStatusCallback`. Can be: `start`, `end`, `join` and `leave`. Separate
	// multiple values with a space. By default no callbacks are sent.
	ConferenceStatusCallbackEvent param.Opt[string] `json:"ConferenceStatusCallbackEvent,omitzero"`
	// Whether participant shall be bridged to conference before the participant
	// answers (from early media if available). Defaults to `false`.
	EarlyMedia param.Opt[bool] `json:"EarlyMedia,omitzero"`
	// Whether to end the conference when the participant leaves. Defaults to `false`.
	EndConferenceOnExit param.Opt[bool] `json:"EndConferenceOnExit,omitzero"`
	// The phone number of the party that initiated the call. Phone numbers are
	// formatted with a `+` and country code.
	From param.Opt[string] `json:"From,omitzero"`
	// A unique label for the participant that will be added to the conference. The
	// label can be used to reference the participant for updates via the TeXML REST
	// API.
	Label param.Opt[string] `json:"Label,omitzero"`
	// If initial silence duration is greater than this value, consider it a machine.
	// Ignored when `premium` detection is used.
	MachineDetectionSilenceTimeout param.Opt[int64] `json:"MachineDetectionSilenceTimeout,omitzero"`
	// Silence duration threshold after a greeting message or voice for it be
	// considered human. Ignored when `premium` detection is used.
	MachineDetectionSpeechEndThreshold param.Opt[int64] `json:"MachineDetectionSpeechEndThreshold,omitzero"`
	// Maximum threshold of a human greeting. If greeting longer than this value,
	// considered machine. Ignored when `premium` detection is used.
	MachineDetectionSpeechThreshold param.Opt[int64] `json:"MachineDetectionSpeechThreshold,omitzero"`
	// How long answering machine detection should go on for before sending an
	// `Unknown` result. Given in milliseconds.
	MachineDetectionTimeout param.Opt[int64] `json:"MachineDetectionTimeout,omitzero"`
	// The maximum number of participants in the conference. Can be a positive integer
	// from 2 to 800. The default value is 250.
	MaxParticipants param.Opt[int64] `json:"MaxParticipants,omitzero"`
	// Whether the participant should be muted.
	Muted param.Opt[bool] `json:"Muted,omitzero"`
	// The list of comma-separated codecs to be offered on a call.
	PreferredCodecs param.Opt[string] `json:"PreferredCodecs,omitzero"`
	// Whether to record the entire participant's call leg. Defaults to `false`.
	Record param.Opt[bool] `json:"Record,omitzero"`
	// The URL the recording callbacks will be sent to.
	RecordingStatusCallback param.Opt[string] `json:"RecordingStatusCallback,omitzero"`
	// The changes to the recording's state that should generate a call to
	// `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`.
	RecordingStatusCallbackEvent param.Opt[string] `json:"RecordingStatusCallbackEvent,omitzero"`
	// The password to use for SIP authentication.
	SipAuthPassword param.Opt[string] `json:"SipAuthPassword,omitzero"`
	// The username to use for SIP authentication.
	SipAuthUsername param.Opt[string] `json:"SipAuthUsername,omitzero"`
	// Whether to start the conference when the participant enters. Defaults to `true`.
	StartConferenceOnEnter param.Opt[bool] `json:"StartConferenceOnEnter,omitzero"`
	// URL destination for Telnyx to send status callback events to for the call.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// The changes to the call's state that should generate a call to `StatusCallback`.
	// Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple
	// values with a space. The default value is `completed`.
	StatusCallbackEvent param.Opt[string] `json:"StatusCallbackEvent,omitzero"`
	// The maximum duration of the call in seconds.
	TimeLimit param.Opt[int64] `json:"TimeLimit,omitzero"`
	// The number of seconds that we should allow the phone to ring before assuming
	// there is no answer. Can be an integer between 5 and 120, inclusive. The default
	// value is 30.
	TimeoutSeconds param.Opt[int64] `json:"Timeout,omitzero"`
	// The phone number of the called party. Phone numbers are formatted with a `+` and
	// country code.
	To param.Opt[string] `json:"To,omitzero"`
	// The URL to call for an audio file to play while the participant is waiting for
	// the conference to start.
	WaitURL param.Opt[string] `json:"WaitUrl,omitzero"`
	// HTTP request type used for `AmdStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	AmdStatusCallbackMethod TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethod `json:"AmdStatusCallbackMethod,omitzero"`
	// Whether to play a notification beep to the conference when the participant
	// enters and exits.
	//
	// Any of "true", "false", "onEnter", "onExit".
	Beep TexmlAccountConferenceParticipantParticipantsParamsBeep `json:"Beep,omitzero"`
	// Whether to record the conference the participant is joining. Defualts to
	// `do-not-record`. The boolean values `true` and `false` are synonymous with
	// `record-from-start` and `do-not-record` respectively.
	//
	// Any of "true", "false", "record-from-start", "do-not-record".
	ConferenceRecord TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord `json:"ConferenceRecord,omitzero"`
	// HTTP request type used for `ConferenceRecordingStatusCallback`. Defaults to
	// `POST`.
	//
	// Any of "GET", "POST".
	ConferenceRecordingStatusCallbackMethod TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethod `json:"ConferenceRecordingStatusCallbackMethod,omitzero"`
	// HTTP request type used for `ConferenceStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	ConferenceStatusCallbackMethod TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethod `json:"ConferenceStatusCallbackMethod,omitzero"`
	// Whether to trim any leading and trailing silence from the conference recording.
	// Defaults to `trim-silence`.
	//
	// Any of "trim-silence", "do-not-trim".
	ConferenceTrim TexmlAccountConferenceParticipantParticipantsParamsConferenceTrim `json:"ConferenceTrim,omitzero"`
	// Custom HTTP headers to be sent with the call. Each header should be an object
	// with 'name' and 'value' properties.
	CustomHeaders []TexmlAccountConferenceParticipantParticipantsParamsCustomHeader `json:"CustomHeaders,omitzero"`
	// Whether to detect if a human or an answering machine picked up the call. Use
	// `Enable` if you would like to ne notified as soon as the called party is
	// identified. Use `DetectMessageEnd`, if you would like to leave a message on an
	// answering machine.
	//
	// Any of "Enable", "DetectMessageEnd".
	MachineDetection TexmlAccountConferenceParticipantParticipantsParamsMachineDetection `json:"MachineDetection,omitzero"`
	// The number of channels in the final recording. Defaults to `mono`.
	//
	// Any of "mono", "dual".
	RecordingChannels TexmlAccountConferenceParticipantParticipantsParamsRecordingChannels `json:"RecordingChannels,omitzero"`
	// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	RecordingStatusCallbackMethod TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethod `json:"RecordingStatusCallbackMethod,omitzero"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	RecordingTrack TexmlAccountConferenceParticipantParticipantsParamsRecordingTrack `json:"RecordingTrack,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// Whether to trim any leading and trailing silence from the recording. Defaults to
	// `trim-silence`.
	//
	// Any of "trim-silence", "do-not-trim".
	Trim TexmlAccountConferenceParticipantParticipantsParamsTrim `json:"Trim,omitzero"`
	paramObj
}

func (r TexmlAccountConferenceParticipantParticipantsParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountConferenceParticipantParticipantsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountConferenceParticipantParticipantsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `AmdStatusCallback`. Defaults to `POST`.
type TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethod string

const (
	TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethodGet  TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethod = "GET"
	TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethodPost TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethod = "POST"
)

// Whether to play a notification beep to the conference when the participant
// enters and exits.
type TexmlAccountConferenceParticipantParticipantsParamsBeep string

const (
	TexmlAccountConferenceParticipantParticipantsParamsBeepTrue    TexmlAccountConferenceParticipantParticipantsParamsBeep = "true"
	TexmlAccountConferenceParticipantParticipantsParamsBeepFalse   TexmlAccountConferenceParticipantParticipantsParamsBeep = "false"
	TexmlAccountConferenceParticipantParticipantsParamsBeepOnEnter TexmlAccountConferenceParticipantParticipantsParamsBeep = "onEnter"
	TexmlAccountConferenceParticipantParticipantsParamsBeepOnExit  TexmlAccountConferenceParticipantParticipantsParamsBeep = "onExit"
)

// Whether to record the conference the participant is joining. Defualts to
// `do-not-record`. The boolean values `true` and `false` are synonymous with
// `record-from-start` and `do-not-record` respectively.
type TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord string

const (
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordTrue            TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord = "true"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordFalse           TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord = "false"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordRecordFromStart TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord = "record-from-start"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordDoNotRecord     TexmlAccountConferenceParticipantParticipantsParamsConferenceRecord = "do-not-record"
)

// HTTP request type used for `ConferenceRecordingStatusCallback`. Defaults to
// `POST`.
type TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethod string

const (
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethodGet  TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethod = "GET"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethodPost TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethod = "POST"
)

// HTTP request type used for `ConferenceStatusCallback`. Defaults to `POST`.
type TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethod string

const (
	TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethodGet  TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethod = "GET"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethodPost TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethod = "POST"
)

// Whether to trim any leading and trailing silence from the conference recording.
// Defaults to `trim-silence`.
type TexmlAccountConferenceParticipantParticipantsParamsConferenceTrim string

const (
	TexmlAccountConferenceParticipantParticipantsParamsConferenceTrimTrimSilence TexmlAccountConferenceParticipantParticipantsParamsConferenceTrim = "trim-silence"
	TexmlAccountConferenceParticipantParticipantsParamsConferenceTrimDoNotTrim   TexmlAccountConferenceParticipantParticipantsParamsConferenceTrim = "do-not-trim"
)

// The properties Name, Value are required.
type TexmlAccountConferenceParticipantParticipantsParamsCustomHeader struct {
	// The name of the custom header
	Name string `json:"name,required"`
	// The value of the custom header
	Value string `json:"value,required"`
	paramObj
}

func (r TexmlAccountConferenceParticipantParticipantsParamsCustomHeader) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountConferenceParticipantParticipantsParamsCustomHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountConferenceParticipantParticipantsParamsCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether to detect if a human or an answering machine picked up the call. Use
// `Enable` if you would like to ne notified as soon as the called party is
// identified. Use `DetectMessageEnd`, if you would like to leave a message on an
// answering machine.
type TexmlAccountConferenceParticipantParticipantsParamsMachineDetection string

const (
	TexmlAccountConferenceParticipantParticipantsParamsMachineDetectionEnable           TexmlAccountConferenceParticipantParticipantsParamsMachineDetection = "Enable"
	TexmlAccountConferenceParticipantParticipantsParamsMachineDetectionDetectMessageEnd TexmlAccountConferenceParticipantParticipantsParamsMachineDetection = "DetectMessageEnd"
)

// The number of channels in the final recording. Defaults to `mono`.
type TexmlAccountConferenceParticipantParticipantsParamsRecordingChannels string

const (
	TexmlAccountConferenceParticipantParticipantsParamsRecordingChannelsMono TexmlAccountConferenceParticipantParticipantsParamsRecordingChannels = "mono"
	TexmlAccountConferenceParticipantParticipantsParamsRecordingChannelsDual TexmlAccountConferenceParticipantParticipantsParamsRecordingChannels = "dual"
)

// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
type TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethod string

const (
	TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethodGet  TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethod = "GET"
	TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethodPost TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethod = "POST"
)

// The audio track to record for the call. The default is `both`.
type TexmlAccountConferenceParticipantParticipantsParamsRecordingTrack string

const (
	TexmlAccountConferenceParticipantParticipantsParamsRecordingTrackInbound  TexmlAccountConferenceParticipantParticipantsParamsRecordingTrack = "inbound"
	TexmlAccountConferenceParticipantParticipantsParamsRecordingTrackOutbound TexmlAccountConferenceParticipantParticipantsParamsRecordingTrack = "outbound"
	TexmlAccountConferenceParticipantParticipantsParamsRecordingTrackBoth     TexmlAccountConferenceParticipantParticipantsParamsRecordingTrack = "both"
)

// HTTP request type used for `StatusCallback`.
type TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethod string

const (
	TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethodGet  TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethod = "GET"
	TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethodPost TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethod = "POST"
)

// Whether to trim any leading and trailing silence from the recording. Defaults to
// `trim-silence`.
type TexmlAccountConferenceParticipantParticipantsParamsTrim string

const (
	TexmlAccountConferenceParticipantParticipantsParamsTrimTrimSilence TexmlAccountConferenceParticipantParticipantsParamsTrim = "trim-silence"
	TexmlAccountConferenceParticipantParticipantsParamsTrimDoNotTrim   TexmlAccountConferenceParticipantParticipantsParamsTrim = "do-not-trim"
)

type TexmlAccountConferenceParticipantGetParticipantsParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}
