// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlAccountTranscriptionJsonService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountTranscriptionJsonService] method instead.
type TexmlAccountTranscriptionJsonService struct {
	Options []option.RequestOption
}

// NewTexmlAccountTranscriptionJsonService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountTranscriptionJsonService(opts ...option.RequestOption) (r TexmlAccountTranscriptionJsonService) {
	r = TexmlAccountTranscriptionJsonService{}
	r.Options = opts
	return
}

// Permanently deletes a recording transcription.
func (r *TexmlAccountTranscriptionJsonService) DeleteRecordingTranscriptionSidJson(ctx context.Context, recordingTranscriptionSid string, body TexmlAccountTranscriptionJsonDeleteRecordingTranscriptionSidJsonParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if recordingTranscriptionSid == "" {
		err = errors.New("missing required recording_transcription_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Transcriptions/%s.json", body.AccountSid, recordingTranscriptionSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns the recording transcription resource identified by its ID.
func (r *TexmlAccountTranscriptionJsonService) GetRecordingTranscriptionSidJson(ctx context.Context, recordingTranscriptionSid string, query TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonParams, opts ...option.RequestOption) (res *TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if recordingTranscriptionSid == "" {
		err = errors.New("missing required recording_transcription_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Transcriptions/%s.json", query.AccountSid, recordingTranscriptionSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponse struct {
	AccountSid string `json:"account_sid"`
	// The version of the API that was used to make the request.
	APIVersion  string    `json:"api_version"`
	CallSid     string    `json:"call_sid"`
	DateCreated time.Time `json:"date_created" format:"date-time"`
	DateUpdated time.Time `json:"date_updated" format:"date-time"`
	// The duration of this recording, given in seconds.
	Duration string `json:"duration,nullable"`
	// Identifier of a resource.
	RecordingSid string `json:"recording_sid"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// The status of the recording transcriptions. The transcription text will be
	// available only when the status is completed.
	//
	// Any of "in-progress", "completed".
	Status TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatus `json:"status"`
	// The recording's transcribed text
	TranscriptionText string `json:"transcription_text"`
	// The relative URI for the recording transcription resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid        respjson.Field
		APIVersion        respjson.Field
		CallSid           respjson.Field
		DateCreated       respjson.Field
		DateUpdated       respjson.Field
		Duration          respjson.Field
		RecordingSid      respjson.Field
		Sid               respjson.Field
		Status            respjson.Field
		TranscriptionText respjson.Field
		Uri               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the recording transcriptions. The transcription text will be
// available only when the status is completed.
type TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatus string

const (
	TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatusInProgress TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatus = "in-progress"
	TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatusCompleted  TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonResponseStatus = "completed"
)

type TexmlAccountTranscriptionJsonDeleteRecordingTranscriptionSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountTranscriptionJsonGetRecordingTranscriptionSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}
