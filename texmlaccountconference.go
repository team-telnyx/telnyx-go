// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// TexmlAccountConferenceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountConferenceService] method instead.
type TexmlAccountConferenceService struct {
	Options      []option.RequestOption
	Participants TexmlAccountConferenceParticipantService
}

// NewTexmlAccountConferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlAccountConferenceService(opts ...option.RequestOption) (r TexmlAccountConferenceService) {
	r = TexmlAccountConferenceService{}
	r.Options = opts
	r.Participants = NewTexmlAccountConferenceParticipantService(opts...)
	return
}

// Returns a conference resource.
func (r *TexmlAccountConferenceService) Get(ctx context.Context, conferenceSid string, query TexmlAccountConferenceGetParams, opts ...option.RequestOption) (res *TexmlAccountConferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s", query.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a conference resource.
func (r *TexmlAccountConferenceService) Update(ctx context.Context, conferenceSid string, params TexmlAccountConferenceUpdateParams, opts ...option.RequestOption) (res *TexmlAccountConferenceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s", params.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists conference resources.
func (r *TexmlAccountConferenceService) GetConferences(ctx context.Context, accountSid string, query TexmlAccountConferenceGetConferencesParams, opts ...option.RequestOption) (res *TexmlAccountConferenceGetConferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Lists conference recordings
func (r *TexmlAccountConferenceService) GetRecordings(ctx context.Context, conferenceSid string, query TexmlAccountConferenceGetRecordingsParams, opts ...option.RequestOption) (res *TexmlAccountConferenceGetRecordingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Recordings", query.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns recordings for a conference identified by conference_sid.
func (r *TexmlAccountConferenceService) GetRecordingsJson(ctx context.Context, conferenceSid string, query TexmlAccountConferenceGetRecordingsJsonParams, opts ...option.RequestOption) (res *TexmlAccountConferenceGetRecordingsJsonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if conferenceSid == "" {
		err = errors.New("missing required conference_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Conferences/%s/Recordings.json", query.AccountSid, conferenceSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TexmlAccountConferenceGetResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// Caller ID, if present.
	CallSidEndingConference string `json:"call_sid_ending_conference"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// A string that you assigned to describe this conference room.
	FriendlyName string `json:"friendly_name"`
	// The reason why a conference ended. When a conference is in progress, will be
	// null.
	//
	// Any of "participant-with-end-conference-on-exit-left", "last-participant-left",
	// "conference-ended-via-api", "time-exceeded".
	ReasonConferenceEnded TexmlAccountConferenceGetResponseReasonConferenceEnded `json:"reason_conference_ended"`
	// A string representing the region where the conference is hosted.
	Region string `json:"region"`
	// The unique identifier of the conference.
	Sid string `json:"sid"`
	// The status of this conference.
	//
	// Any of "init", "in-progress", "completed".
	Status TexmlAccountConferenceGetResponseStatus `json:"status"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this conference.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid              respjson.Field
		APIVersion              respjson.Field
		CallSidEndingConference respjson.Field
		DateCreated             respjson.Field
		DateUpdated             respjson.Field
		FriendlyName            respjson.Field
		ReasonConferenceEnded   respjson.Field
		Region                  respjson.Field
		Sid                     respjson.Field
		Status                  respjson.Field
		SubresourceUris         respjson.Field
		Uri                     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why a conference ended. When a conference is in progress, will be
// null.
type TexmlAccountConferenceGetResponseReasonConferenceEnded string

const (
	TexmlAccountConferenceGetResponseReasonConferenceEndedParticipantWithEndConferenceOnExitLeft TexmlAccountConferenceGetResponseReasonConferenceEnded = "participant-with-end-conference-on-exit-left"
	TexmlAccountConferenceGetResponseReasonConferenceEndedLastParticipantLeft                    TexmlAccountConferenceGetResponseReasonConferenceEnded = "last-participant-left"
	TexmlAccountConferenceGetResponseReasonConferenceEndedConferenceEndedViaAPI                  TexmlAccountConferenceGetResponseReasonConferenceEnded = "conference-ended-via-api"
	TexmlAccountConferenceGetResponseReasonConferenceEndedTimeExceeded                           TexmlAccountConferenceGetResponseReasonConferenceEnded = "time-exceeded"
)

// The status of this conference.
type TexmlAccountConferenceGetResponseStatus string

const (
	TexmlAccountConferenceGetResponseStatusInit       TexmlAccountConferenceGetResponseStatus = "init"
	TexmlAccountConferenceGetResponseStatusInProgress TexmlAccountConferenceGetResponseStatus = "in-progress"
	TexmlAccountConferenceGetResponseStatusCompleted  TexmlAccountConferenceGetResponseStatus = "completed"
)

type TexmlAccountConferenceUpdateResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// Caller ID, if present.
	CallSidEndingConference string `json:"call_sid_ending_conference"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// A string that you assigned to describe this conference room.
	FriendlyName string `json:"friendly_name"`
	// The reason why a conference ended. When a conference is in progress, will be
	// null.
	//
	// Any of "participant-with-end-conference-on-exit-left", "last-participant-left",
	// "conference-ended-via-api", "time-exceeded".
	ReasonConferenceEnded TexmlAccountConferenceUpdateResponseReasonConferenceEnded `json:"reason_conference_ended"`
	// A string representing the region where the conference is hosted.
	Region string `json:"region"`
	// The unique identifier of the conference.
	Sid string `json:"sid"`
	// The status of this conference.
	//
	// Any of "init", "in-progress", "completed".
	Status TexmlAccountConferenceUpdateResponseStatus `json:"status"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this conference.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid              respjson.Field
		APIVersion              respjson.Field
		CallSidEndingConference respjson.Field
		DateCreated             respjson.Field
		DateUpdated             respjson.Field
		FriendlyName            respjson.Field
		ReasonConferenceEnded   respjson.Field
		Region                  respjson.Field
		Sid                     respjson.Field
		Status                  respjson.Field
		SubresourceUris         respjson.Field
		Uri                     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reason why a conference ended. When a conference is in progress, will be
// null.
type TexmlAccountConferenceUpdateResponseReasonConferenceEnded string

const (
	TexmlAccountConferenceUpdateResponseReasonConferenceEndedParticipantWithEndConferenceOnExitLeft TexmlAccountConferenceUpdateResponseReasonConferenceEnded = "participant-with-end-conference-on-exit-left"
	TexmlAccountConferenceUpdateResponseReasonConferenceEndedLastParticipantLeft                    TexmlAccountConferenceUpdateResponseReasonConferenceEnded = "last-participant-left"
	TexmlAccountConferenceUpdateResponseReasonConferenceEndedConferenceEndedViaAPI                  TexmlAccountConferenceUpdateResponseReasonConferenceEnded = "conference-ended-via-api"
	TexmlAccountConferenceUpdateResponseReasonConferenceEndedTimeExceeded                           TexmlAccountConferenceUpdateResponseReasonConferenceEnded = "time-exceeded"
)

// The status of this conference.
type TexmlAccountConferenceUpdateResponseStatus string

const (
	TexmlAccountConferenceUpdateResponseStatusInit       TexmlAccountConferenceUpdateResponseStatus = "init"
	TexmlAccountConferenceUpdateResponseStatusInProgress TexmlAccountConferenceUpdateResponseStatus = "in-progress"
	TexmlAccountConferenceUpdateResponseStatusCompleted  TexmlAccountConferenceUpdateResponseStatus = "completed"
)

type TexmlAccountConferenceGetConferencesResponse struct {
	Conferences []TexmlAccountConferenceGetConferencesResponseConference `json:"conferences"`
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences.json?Page=0&PageSize=1
	FirstPageUri string `json:"first_page_uri"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences.json?Page=1&PageSize=1&PageToken=MTY4AjgyNDkwNzIxMQ
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize int64 `json:"page_size"`
	// The number of the first element on the page, zero-indexed.
	Start int64 `json:"start"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conferences  respjson.Field
		End          respjson.Field
		FirstPageUri respjson.Field
		NextPageUri  respjson.Field
		Page         respjson.Field
		PageSize     respjson.Field
		Start        respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetConferencesResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetConferencesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceGetConferencesResponseConference struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion string `json:"api_version"`
	// Caller ID, if present.
	CallSidEndingConference string `json:"call_sid_ending_conference"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// A string that you assigned to describe this conference room.
	FriendlyName string `json:"friendly_name"`
	// The reason why a conference ended. When a conference is in progress, will be
	// null.
	//
	// Any of "participant-with-end-conference-on-exit-left", "last-participant-left",
	// "conference-ended-via-api", "time-exceeded".
	ReasonConferenceEnded string `json:"reason_conference_ended"`
	// A string representing the region where the conference is hosted.
	Region string `json:"region"`
	// The unique identifier of the conference.
	Sid string `json:"sid"`
	// The status of this conference.
	//
	// Any of "init", "in-progress", "completed".
	Status string `json:"status"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this conference.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid              respjson.Field
		APIVersion              respjson.Field
		CallSidEndingConference respjson.Field
		DateCreated             respjson.Field
		DateUpdated             respjson.Field
		FriendlyName            respjson.Field
		ReasonConferenceEnded   respjson.Field
		Region                  respjson.Field
		Sid                     respjson.Field
		Status                  respjson.Field
		SubresourceUris         respjson.Field
		Uri                     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetConferencesResponseConference) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetConferencesResponseConference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceGetRecordingsResponse struct {
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Recordings.json?page=0&pagesize=20
	FirstPageUri string `json:"first_page_uri"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Recordings.json?Page=1&PageSize=1&PageToken=MTY4AjgyNDkwNzIxMQ
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize   int64                                                  `json:"page_size"`
	Recordings []TexmlAccountConferenceGetRecordingsResponseRecording `json:"recordings"`
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
		Recordings   respjson.Field
		Start        respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetRecordingsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetRecordingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceGetRecordingsResponseRecording struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The identifier of the related participant's call.
	CallSid string `json:"call_sid"`
	// The number of channels in the recording.
	Channels int64 `json:"channels"`
	// The identifier of the related conference.
	ConferenceSid string `json:"conference_sid"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// Duratin of the recording in seconds.
	Duration int64 `json:"duration"`
	// The recording error, if any.
	ErrorCode string `json:"error_code"`
	// The URL to use to download the recording.
	MediaURL string `json:"media_url"`
	// The unique identifier of the recording.
	Sid string `json:"sid"`
	// How the recording was started.
	//
	// Any of "DialVerb", "Conference", "OutboundAPI", "Trunking", "RecordVerb",
	// "StartCallRecordingAPI", "StartConferenceRecordingAPI".
	Source string `json:"source"`
	// The timestamp of when the recording was started.
	StartTime string `json:"start_time"`
	// The status of the recording.
	//
	// Any of "processing", "absent", "completed", "deleted".
	Status string `json:"status"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this recording.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid      respjson.Field
		CallSid         respjson.Field
		Channels        respjson.Field
		ConferenceSid   respjson.Field
		DateCreated     respjson.Field
		DateUpdated     respjson.Field
		Duration        respjson.Field
		ErrorCode       respjson.Field
		MediaURL        respjson.Field
		Sid             respjson.Field
		Source          respjson.Field
		StartTime       respjson.Field
		Status          respjson.Field
		SubresourceUris respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetRecordingsResponseRecording) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetRecordingsResponseRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceGetRecordingsJsonResponse struct {
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// Relative uri to the first page of the query results
	FirstPageUri string `json:"first_page_uri" format:"uri"`
	// Relative uri to the next page of the query results
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize int64 `json:"page_size"`
	// Relative uri to the previous page of the query results
	PreviousPageUri string                              `json:"previous_page_uri" format:"uri"`
	Recordings      []TexmlGetCallRecordingResponseBody `json:"recordings"`
	// The number of the first element on the page, zero-indexed.
	Start int64 `json:"start"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End             respjson.Field
		FirstPageUri    respjson.Field
		NextPageUri     respjson.Field
		Page            respjson.Field
		PageSize        respjson.Field
		PreviousPageUri respjson.Field
		Recordings      respjson.Field
		Start           respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountConferenceGetRecordingsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountConferenceGetRecordingsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountConferenceGetParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountConferenceUpdateParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	// The URL we should call to announce something into the conference. The URL may
	// return an MP3 file, a WAV file, or a TwiML document that contains `<Play>`,
	// `<Say>`, `<Pause>`, or `<Redirect>` verbs.
	AnnounceURL param.Opt[string] `json:"AnnounceUrl,omitzero"`
	// The new status of the resource. Specifying `completed` will end the conference
	// and hang up all participants.
	Status param.Opt[string] `json:"Status,omitzero"`
	// The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	AnnounceMethod TexmlAccountConferenceUpdateParamsAnnounceMethod `json:"AnnounceMethod,omitzero"`
	paramObj
}

func (r TexmlAccountConferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountConferenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountConferenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method used to call the `AnnounceUrl`. Defaults to `POST`.
type TexmlAccountConferenceUpdateParamsAnnounceMethod string

const (
	TexmlAccountConferenceUpdateParamsAnnounceMethodGet  TexmlAccountConferenceUpdateParamsAnnounceMethod = "GET"
	TexmlAccountConferenceUpdateParamsAnnounceMethodPost TexmlAccountConferenceUpdateParamsAnnounceMethod = "POST"
)

type TexmlAccountConferenceGetConferencesParams struct {
	// Filters conferences by the creation date. Expected format is YYYY-MM-DD. Also
	// accepts inequality operators, e.g. DateCreated>=2023-05-22.
	DateCreated param.Opt[string] `query:"DateCreated,omitzero" json:"-"`
	// Filters conferences by the time they were last updated. Expected format is
	// YYYY-MM-DD. Also accepts inequality operators, e.g. DateUpdated>=2023-05-22.
	DateUpdated param.Opt[string] `query:"DateUpdated,omitzero" json:"-"`
	// Filters conferences by their friendly name.
	FriendlyName param.Opt[string] `query:"FriendlyName,omitzero" json:"-"`
	// The number of the page to be displayed, zero-indexed, should be used in
	// conjuction with PageToken.
	Page param.Opt[int64] `query:"Page,omitzero" json:"-"`
	// The number of records to be displayed on a page
	PageSize param.Opt[int64] `query:"PageSize,omitzero" json:"-"`
	// Used to request the next page of results.
	PageToken param.Opt[string] `query:"PageToken,omitzero" json:"-"`
	// Filters conferences by status.
	//
	// Any of "init", "in-progress", "completed".
	Status TexmlAccountConferenceGetConferencesParamsStatus `query:"Status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlAccountConferenceGetConferencesParams]'s query
// parameters as `url.Values`.
func (r TexmlAccountConferenceGetConferencesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters conferences by status.
type TexmlAccountConferenceGetConferencesParamsStatus string

const (
	TexmlAccountConferenceGetConferencesParamsStatusInit       TexmlAccountConferenceGetConferencesParamsStatus = "init"
	TexmlAccountConferenceGetConferencesParamsStatusInProgress TexmlAccountConferenceGetConferencesParamsStatus = "in-progress"
	TexmlAccountConferenceGetConferencesParamsStatusCompleted  TexmlAccountConferenceGetConferencesParamsStatus = "completed"
)

type TexmlAccountConferenceGetRecordingsParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountConferenceGetRecordingsJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}
