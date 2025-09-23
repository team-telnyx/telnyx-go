// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// TexmlAccountCallRecordingService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallRecordingService] method instead.
type TexmlAccountCallRecordingService struct {
	Options []option.RequestOption
}

// NewTexmlAccountCallRecordingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountCallRecordingService(opts ...option.RequestOption) (r TexmlAccountCallRecordingService) {
	r = TexmlAccountCallRecordingService{}
	r.Options = opts
	return
}

// Updates recording resource for particular call.
func (r *TexmlAccountCallRecordingService) RecordingSidJson(ctx context.Context, recordingSid string, params TexmlAccountCallRecordingRecordingSidJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallRecordingRecordingSidJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if params.CallSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	if recordingSid == "" {
		err = errors.New("missing required recording_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Recordings/%s.json", params.AccountSid, params.CallSid, recordingSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TexmlAccountCallRecordingRecordingSidJsonResponse struct {
	AccountSid string `json:"account_sid"`
	CallSid    string `json:"call_sid"`
	// Any of 1, 2.
	Channels      int64     `json:"channels"`
	ConferenceSid string    `json:"conference_sid,nullable" format:"uuid"`
	DateCreated   time.Time `json:"date_created" format:"date-time"`
	DateUpdated   time.Time `json:"date_updated" format:"date-time"`
	// The duration of this recording, given in seconds.
	Duration  string `json:"duration,nullable"`
	ErrorCode string `json:"error_code,nullable"`
	// The price of this recording, the currency is specified in the price_unit field.
	Price string `json:"price,nullable"`
	// The unit in which the price is given.
	PriceUnit string `json:"price_unit,nullable"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// Defines how the recording was created.
	//
	// Any of "StartCallRecordingAPI", "StartConferenceRecordingAPI", "OutboundAPI",
	// "DialVerb", "Conference", "RecordVerb", "Trunking".
	Source    TexmlAccountCallRecordingRecordingSidJsonResponseSource `json:"source"`
	StartTime time.Time                                               `json:"start_time" format:"date-time"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	Track TexmlAccountCallRecordingRecordingSidJsonResponseTrack `json:"track"`
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
func (r TexmlAccountCallRecordingRecordingSidJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallRecordingRecordingSidJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines how the recording was created.
type TexmlAccountCallRecordingRecordingSidJsonResponseSource string

const (
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceStartCallRecordingAPI       TexmlAccountCallRecordingRecordingSidJsonResponseSource = "StartCallRecordingAPI"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceStartConferenceRecordingAPI TexmlAccountCallRecordingRecordingSidJsonResponseSource = "StartConferenceRecordingAPI"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceOutboundAPI                 TexmlAccountCallRecordingRecordingSidJsonResponseSource = "OutboundAPI"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceDialVerb                    TexmlAccountCallRecordingRecordingSidJsonResponseSource = "DialVerb"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceConference                  TexmlAccountCallRecordingRecordingSidJsonResponseSource = "Conference"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceRecordVerb                  TexmlAccountCallRecordingRecordingSidJsonResponseSource = "RecordVerb"
	TexmlAccountCallRecordingRecordingSidJsonResponseSourceTrunking                    TexmlAccountCallRecordingRecordingSidJsonResponseSource = "Trunking"
)

// The audio track to record for the call. The default is `both`.
type TexmlAccountCallRecordingRecordingSidJsonResponseTrack string

const (
	TexmlAccountCallRecordingRecordingSidJsonResponseTrackInbound  TexmlAccountCallRecordingRecordingSidJsonResponseTrack = "inbound"
	TexmlAccountCallRecordingRecordingSidJsonResponseTrackOutbound TexmlAccountCallRecordingRecordingSidJsonResponseTrack = "outbound"
	TexmlAccountCallRecordingRecordingSidJsonResponseTrackBoth     TexmlAccountCallRecordingRecordingSidJsonResponseTrack = "both"
)

type TexmlAccountCallRecordingRecordingSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	CallSid    string `path:"call_sid,required" json:"-"`
	// Any of "in-progress", "paused", "stopped".
	Status TexmlAccountCallRecordingRecordingSidJsonParamsStatus `json:"Status,omitzero"`
	paramObj
}

func (r TexmlAccountCallRecordingRecordingSidJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallRecordingRecordingSidJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallRecordingRecordingSidJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallRecordingRecordingSidJsonParamsStatus string

const (
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusInProgress TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "in-progress"
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusPaused     TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "paused"
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusStopped    TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "stopped"
)
