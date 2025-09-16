// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// ConferenceService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConferenceService] method instead.
type ConferenceService struct {
	Options []option.RequestOption
	Actions ConferenceActionService
}

// NewConferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConferenceService(opts ...option.RequestOption) (r ConferenceService) {
	r = ConferenceService{}
	r.Options = opts
	r.Actions = NewConferenceActionService(opts...)
	return
}

// Create a conference from an existing call leg using a `call_control_id` and a
// conference name. Upon creating the conference, the call will be automatically
// bridged to the conference. Conferences will expire after all participants have
// left the conference or after 4 hours regardless of the number of active
// participants.
//
// **Expected Webhooks (see
// [callback schema](https://developers.telnyx.com/api/call-control/create-conference#callbacks)
// below):**
//
// - `conference.created`
// - `conference.participant.joined`
// - `conference.participant.left`
// - `conference.ended`
// - `conference.recording.saved`
// - `conference.floor.changed`
func (r *ConferenceService) New(ctx context.Context, body ConferenceNewParams, opts ...option.RequestOption) (res *ConferenceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "conferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an existing conference
func (r *ConferenceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists conferences. Conferences are created on demand, and will expire after all
// participants have left the conference or after 4 hours regardless of the number
// of active participants. Conferences are listed in descending order by
// `expires_at`.
func (r *ConferenceService) List(ctx context.Context, query ConferenceListParams, opts ...option.RequestOption) (res *ConferenceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "conferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Lists conference participants
func (r *ConferenceService) ListParticipants(ctx context.Context, conferenceID string, query ConferenceListParticipantsParams, opts ...option.RequestOption) (res *ConferenceListParticipantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if conferenceID == "" {
		err = errors.New("missing required conference_id parameter")
		return
	}
	path := fmt.Sprintf("conferences/%s/participants", conferenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Conference struct {
	// Uniquely identifies the conference
	ID string `json:"id,required"`
	// ISO 8601 formatted date of when the conference was created
	CreatedAt string `json:"created_at,required"`
	// ISO 8601 formatted date of when the conference will expire
	ExpiresAt string `json:"expires_at,required"`
	// Name of the conference
	Name string `json:"name,required"`
	// Any of "conference".
	RecordType ConferenceRecordType `json:"record_type,required"`
	// Identifies the connection associated with the conference
	ConnectionID string `json:"connection_id"`
	// Reason why the conference ended
	//
	// Any of "all_left", "ended_via_api", "host_left", "time_exceeded".
	EndReason ConferenceEndReason `json:"end_reason"`
	// IDs related to who ended the conference. It is expected for them to all be there
	// or all be null
	EndedBy ConferenceEndedBy `json:"ended_by"`
	// Region where the conference is hosted
	Region string `json:"region"`
	// Status of the conference
	//
	// Any of "init", "in_progress", "completed".
	Status ConferenceStatus `json:"status"`
	// ISO 8601 formatted date of when the conference was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		ExpiresAt    respjson.Field
		Name         respjson.Field
		RecordType   respjson.Field
		ConnectionID respjson.Field
		EndReason    respjson.Field
		EndedBy      respjson.Field
		Region       respjson.Field
		Status       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conference) RawJSON() string { return r.JSON.raw }
func (r *Conference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceRecordType string

const (
	ConferenceRecordTypeConference ConferenceRecordType = "conference"
)

// Reason why the conference ended
type ConferenceEndReason string

const (
	ConferenceEndReasonAllLeft      ConferenceEndReason = "all_left"
	ConferenceEndReasonEndedViaAPI  ConferenceEndReason = "ended_via_api"
	ConferenceEndReasonHostLeft     ConferenceEndReason = "host_left"
	ConferenceEndReasonTimeExceeded ConferenceEndReason = "time_exceeded"
)

// IDs related to who ended the conference. It is expected for them to all be there
// or all be null
type ConferenceEndedBy struct {
	// Call Control ID which ended the conference
	CallControlID string `json:"call_control_id"`
	// Call Session ID which ended the conference
	CallSessionID string `json:"call_session_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallSessionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceEndedBy) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEndedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the conference
type ConferenceStatus string

const (
	ConferenceStatusInit       ConferenceStatus = "init"
	ConferenceStatusInProgress ConferenceStatus = "in_progress"
	ConferenceStatusCompleted  ConferenceStatus = "completed"
)

type ConferenceNewResponse struct {
	Data Conference `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceGetResponse struct {
	Data Conference `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceListResponse struct {
	Data []Conference   `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceListParticipantsResponse struct {
	Data []ConferenceListParticipantsResponseData `json:"data"`
	Meta PaginationMeta                           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceListParticipantsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConferenceListParticipantsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceListParticipantsResponseData struct {
	// Uniquely identifies the participant
	ID string `json:"id,required"`
	// Call Control ID associated with the partiipant of the conference
	CallControlID string `json:"call_control_id,required"`
	// Uniquely identifies the call leg associated with the participant
	CallLegID string `json:"call_leg_id,required"`
	// Info about the conference that the participant is in
	Conference ConferenceListParticipantsResponseDataConference `json:"conference,required"`
	// ISO 8601 formatted date of when the participant was created
	CreatedAt string `json:"created_at,required"`
	// Whether the conference will end and all remaining participants be hung up after
	// the participant leaves the conference.
	EndConferenceOnExit bool `json:"end_conference_on_exit,required"`
	// Whether the participant is muted.
	Muted bool `json:"muted,required"`
	// Whether the participant is put on_hold.
	OnHold bool `json:"on_hold,required"`
	// Any of "participant".
	RecordType string `json:"record_type,required"`
	// Whether the conference will end after the participant leaves the conference.
	SoftEndConferenceOnExit bool `json:"soft_end_conference_on_exit,required"`
	// The status of the participant with respect to the lifecycle within the
	// conference
	//
	// Any of "joining", "joined", "left".
	Status string `json:"status,required"`
	// ISO 8601 formatted date of when the participant was last updated
	UpdatedAt string `json:"updated_at,required"`
	// Array of unique call_control_ids the participant can whisper to..
	WhisperCallControlIDs []string `json:"whisper_call_control_ids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CallControlID           respjson.Field
		CallLegID               respjson.Field
		Conference              respjson.Field
		CreatedAt               respjson.Field
		EndConferenceOnExit     respjson.Field
		Muted                   respjson.Field
		OnHold                  respjson.Field
		RecordType              respjson.Field
		SoftEndConferenceOnExit respjson.Field
		Status                  respjson.Field
		UpdatedAt               respjson.Field
		WhisperCallControlIDs   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceListParticipantsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceListParticipantsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Info about the conference that the participant is in
type ConferenceListParticipantsResponseDataConference struct {
	// Uniquely identifies the conference
	ID string `json:"id"`
	// Name of the conference
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceListParticipantsResponseDataConference) RawJSON() string { return r.JSON.raw }
func (r *ConferenceListParticipantsResponseDataConference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceNewParams struct {
	// Unique identifier and token for controlling the call
	CallControlID string `json:"call_control_id,required"`
	// Name of the conference
	Name string `json:"name,required"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string. The client_state will be updated for the creator call
	// leg and will be used for all webhooks related to the created conference.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Toggle background comfort noise.
	ComfortNoise param.Opt[bool] `json:"comfort_noise,omitzero"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore
	// subsequent commands with the same `command_id` as one that has already been
	// executed.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Time length (minutes) after which the conference will end.
	DurationMinutes param.Opt[int64] `json:"duration_minutes,omitzero"`
	// The URL of a file to be played to participants joining the conference. The URL
	// can point to either a WAV or MP3 file. hold_media_name and hold_audio_url cannot
	// be used together in one request. Takes effect only when
	// "start_conference_on_create" is set to "false".
	HoldAudioURL param.Opt[string] `json:"hold_audio_url,omitzero"`
	// The media_name of a file to be played to participants joining the conference.
	// The media_name must point to a file previously uploaded to
	// api.telnyx.com/v2/media by the same user/organization. The file must either be a
	// WAV or MP3 file. Takes effect only when "start_conference_on_create" is set to
	// "false".
	HoldMediaName param.Opt[string] `json:"hold_media_name,omitzero"`
	// The maximum number of active conference participants to allow. Must be between 2
	// and 800. Defaults to 250
	MaxParticipants param.Opt[int64] `json:"max_participants,omitzero"`
	// Whether the conference should be started on creation. If the conference isn't
	// started all participants that join are automatically put on hold. Defaults to
	// "true".
	StartConferenceOnCreate param.Opt[bool] `json:"start_conference_on_create,omitzero"`
	// Whether a beep sound should be played when participants join and/or leave the
	// conference.
	//
	// Any of "always", "never", "on_enter", "on_exit".
	BeepEnabled ConferenceNewParamsBeepEnabled `json:"beep_enabled,omitzero"`
	paramObj
}

func (r ConferenceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConferenceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConferenceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a beep sound should be played when participants join and/or leave the
// conference.
type ConferenceNewParamsBeepEnabled string

const (
	ConferenceNewParamsBeepEnabledAlways  ConferenceNewParamsBeepEnabled = "always"
	ConferenceNewParamsBeepEnabledNever   ConferenceNewParamsBeepEnabled = "never"
	ConferenceNewParamsBeepEnabledOnEnter ConferenceNewParamsBeepEnabled = "on_enter"
	ConferenceNewParamsBeepEnabledOnExit  ConferenceNewParamsBeepEnabled = "on_exit"
)

type ConferenceListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[application_name][contains], filter[outbound.outbound_voice_profile_id],
	// filter[leg_id], filter[application_session_id], filter[connection_id],
	// filter[product], filter[failed], filter[from], filter[to], filter[name],
	// filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]
	Filter ConferenceListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[before], page[limit], page[size], page[number]
	Page ConferenceListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParams]'s query parameters as `url.Values`.
func (r ConferenceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[application_name][contains], filter[outbound.outbound_voice_profile_id],
// filter[leg_id], filter[application_session_id], filter[connection_id],
// filter[product], filter[failed], filter[from], filter[to], filter[name],
// filter[type], filter[occurred_at][eq/gt/gte/lt/lte], filter[status]
type ConferenceListParamsFilter struct {
	// The unique identifier of the call session. A session may include multiple call
	// leg events.
	ApplicationSessionID param.Opt[string] `query:"application_session_id,omitzero" format:"uuid" json:"-"`
	// The unique identifier of the conection.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Delivery failed or not.
	Failed param.Opt[bool] `query:"failed,omitzero" json:"-"`
	// Filter by From number.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// The unique identifier of an individual call leg.
	LegID param.Opt[string] `query:"leg_id,omitzero" format:"uuid" json:"-"`
	// If present, conferences will be filtered to those with a matching `name`
	// attribute. Matching is case-sensitive
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Identifies the associated outbound voice profile.
	OutboundOutboundVoiceProfileID param.Opt[string] `query:"outbound.outbound_voice_profile_id,omitzero" format:"int64" json:"-"`
	// Filter by To number.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	// Application name filters
	ApplicationName ConferenceListParamsFilterApplicationName `query:"application_name,omitzero" json:"-"`
	// Event occurred_at filters
	OccurredAt ConferenceListParamsFilterOccurredAt `query:"occurred_at,omitzero" json:"-"`
	// Filter by product.
	//
	// Any of "call_control", "fax", "texml".
	Product string `query:"product,omitzero" json:"-"`
	// If present, conferences will be filtered by status.
	//
	// Any of "init", "in_progress", "completed".
	Status string `query:"status,omitzero" json:"-"`
	// Event type
	//
	// Any of "command", "webhook".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParamsFilter]'s query parameters as
// `url.Values`.
func (r ConferenceListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[ConferenceListParamsFilter](
		"product", "call_control", "fax", "texml",
	)
	apijson.RegisterFieldValidator[ConferenceListParamsFilter](
		"status", "init", "in_progress", "completed",
	)
	apijson.RegisterFieldValidator[ConferenceListParamsFilter](
		"type", "command", "webhook",
	)
}

// Application name filters
type ConferenceListParamsFilterApplicationName struct {
	// If present, applications with <code>application_name</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParamsFilterApplicationName]'s query
// parameters as `url.Values`.
func (r ConferenceListParamsFilterApplicationName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Event occurred_at filters
type ConferenceListParamsFilterOccurredAt struct {
	// Event occurred_at: equal
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Event occurred_at: greater than
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Event occurred_at: greater than or equal
	Gte param.Opt[string] `query:"gte,omitzero" json:"-"`
	// Event occurred_at: lower than
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	// Event occurred_at: lower than or equal
	Lte param.Opt[string] `query:"lte,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParamsFilterOccurredAt]'s query parameters as
// `url.Values`.
func (r ConferenceListParamsFilterOccurredAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[before], page[limit], page[size], page[number]
type ConferenceListParamsPage struct {
	// Opaque identifier of next page
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Opaque identifier of previous page
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Limit of records per single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParamsPage]'s query parameters as
// `url.Values`.
func (r ConferenceListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConferenceListParticipantsParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[muted],
	// filter[on_hold], filter[whispering]
	Filter ConferenceListParticipantsParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[after],
	// page[before], page[limit], page[size], page[number]
	Page ConferenceListParticipantsParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParticipantsParams]'s query parameters as
// `url.Values`.
func (r ConferenceListParticipantsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[muted],
// filter[on_hold], filter[whispering]
type ConferenceListParticipantsParamsFilter struct {
	// If present, participants will be filtered to those who are/are not muted
	Muted param.Opt[bool] `query:"muted,omitzero" json:"-"`
	// If present, participants will be filtered to those who are/are not put on hold
	OnHold param.Opt[bool] `query:"on_hold,omitzero" json:"-"`
	// If present, participants will be filtered to those who are whispering or are not
	Whispering param.Opt[bool] `query:"whispering,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParticipantsParamsFilter]'s query parameters
// as `url.Values`.
func (r ConferenceListParticipantsParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[after],
// page[before], page[limit], page[size], page[number]
type ConferenceListParticipantsParamsPage struct {
	// Opaque identifier of next page
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Opaque identifier of previous page
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Limit of records per single page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConferenceListParticipantsParamsPage]'s query parameters as
// `url.Values`.
func (r ConferenceListParticipantsParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
