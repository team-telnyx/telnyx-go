// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// TeXML REST Commands
//
// TexmlAccountCallRecordingsJsonService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallRecordingsJsonService] method instead.
type TexmlAccountCallRecordingsJsonService struct {
	Options []option.RequestOption
}

// NewTexmlAccountCallRecordingsJsonService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountCallRecordingsJsonService(opts ...option.RequestOption) (r TexmlAccountCallRecordingsJsonService) {
	r = TexmlAccountCallRecordingsJsonService{}
	r.Options = opts
	return
}

// Starts recording with specified parameters for call idientified by call_sid.
func (r *TexmlAccountCallRecordingsJsonService) RecordingsJson(ctx context.Context, callSid string, params TexmlAccountCallRecordingsJsonRecordingsJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallRecordingsJsonRecordingsJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Recordings.json", params.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns recordings for a call identified by call_sid.
func (r *TexmlAccountCallRecordingsJsonService) GetRecordingsJson(ctx context.Context, callSid string, query TexmlAccountCallRecordingsJsonGetRecordingsJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallRecordingsJsonGetRecordingsJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Recordings.json", query.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TexmlAccountCallRecordingsJsonRecordingsJsonResponse struct {
	AccountSid string `json:"account_sid"`
	CallSid    string `json:"call_sid"`
	// Any of 1, 2.
	Channels      int64     `json:"channels"`
	ConferenceSid string    `json:"conference_sid" api:"nullable" format:"uuid"`
	DateCreated   time.Time `json:"date_created" format:"date-time"`
	DateUpdated   time.Time `json:"date_updated" format:"date-time"`
	// The duration of this recording, given in seconds.
	Duration  string `json:"duration" api:"nullable"`
	ErrorCode string `json:"error_code" api:"nullable"`
	// The price of this recording, the currency is specified in the price_unit field.
	Price string `json:"price" api:"nullable"`
	// The unit in which the price is given.
	PriceUnit string `json:"price_unit" api:"nullable"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// Defines how the recording was created.
	//
	// Any of "StartCallRecordingAPI", "StartConferenceRecordingAPI", "OutboundAPI",
	// "DialVerb", "Conference", "RecordVerb", "Trunking".
	Source    TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource `json:"source"`
	StartTime time.Time                                                  `json:"start_time" format:"date-time"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	Track TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrack `json:"track"`
	// The relative URI for this recording resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid    respjson.Field
		CallSid       respjson.Field
		Channels      respjson.Field
		ConferenceSid respjson.Field
		DateCreated   respjson.Field
		DateUpdated   respjson.Field
		Duration      respjson.Field
		ErrorCode     respjson.Field
		Price         respjson.Field
		PriceUnit     respjson.Field
		Sid           respjson.Field
		Source        respjson.Field
		StartTime     respjson.Field
		Track         respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallRecordingsJsonRecordingsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallRecordingsJsonRecordingsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines how the recording was created.
type TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource string

const (
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceStartCallRecordingAPI       TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "StartCallRecordingAPI"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceStartConferenceRecordingAPI TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "StartConferenceRecordingAPI"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceOutboundAPI                 TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "OutboundAPI"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceDialVerb                    TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "DialVerb"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceConference                  TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "Conference"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceRecordVerb                  TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "RecordVerb"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseSourceTrunking                    TexmlAccountCallRecordingsJsonRecordingsJsonResponseSource = "Trunking"
)

// The audio track to record for the call. The default is `both`.
type TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrack string

const (
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrackInbound  TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrack = "inbound"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrackOutbound TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrack = "outbound"
	TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrackBoth     TexmlAccountCallRecordingsJsonRecordingsJsonResponseTrack = "both"
)

type TexmlAccountCallRecordingsJsonGetRecordingsJsonResponse struct {
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
func (r TexmlAccountCallRecordingsJsonGetRecordingsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallRecordingsJsonGetRecordingsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallRecordingsJsonRecordingsJsonParams struct {
	AccountSid string `path:"account_sid" api:"required" json:"-"`
	// Whether to play a beep when recording is started.
	PlayBeep param.Opt[bool] `json:"PlayBeep,omitzero"`
	// Url where status callbacks will be sent.
	RecordingStatusCallback param.Opt[string] `json:"RecordingStatusCallback,omitzero" format:"uri"`
	// The changes to the recording's state that should generate a call to
	// `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`.
	RecordingStatusCallbackEvent param.Opt[string] `json:"RecordingStatusCallbackEvent,omitzero"`
	// Whether to send RecordingUrl in webhooks.
	SendRecordingURL param.Opt[bool] `json:"SendRecordingUrl,omitzero"`
	// When `dual`, final audio file has the first leg on channel A, and the rest on
	// channel B. `single` mixes both tracks into a single channel.
	//
	// Any of "single", "dual".
	RecordingChannels TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannels `json:"RecordingChannels,omitzero"`
	// HTTP method used to send status callbacks.
	//
	// Any of "GET", "POST".
	RecordingStatusCallbackMethod TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethod `json:"RecordingStatusCallbackMethod,omitzero"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	RecordingTrack TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrack `json:"RecordingTrack,omitzero"`
	paramObj
}

func (r TexmlAccountCallRecordingsJsonRecordingsJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallRecordingsJsonRecordingsJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallRecordingsJsonRecordingsJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When `dual`, final audio file has the first leg on channel A, and the rest on
// channel B. `single` mixes both tracks into a single channel.
type TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannels string

const (
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannelsSingle TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannels = "single"
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannelsDual   TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannels = "dual"
)

// HTTP method used to send status callbacks.
type TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethod string

const (
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethodGet  TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethod = "GET"
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethodPost TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethod = "POST"
)

// The audio track to record for the call. The default is `both`.
type TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrack string

const (
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrackInbound  TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrack = "inbound"
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrackOutbound TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrack = "outbound"
	TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrackBoth     TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrack = "both"
)

type TexmlAccountCallRecordingsJsonGetRecordingsJsonParams struct {
	AccountSid string `path:"account_sid" api:"required" json:"-"`
	paramObj
}
