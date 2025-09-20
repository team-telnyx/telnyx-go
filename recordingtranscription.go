// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// RecordingTranscriptionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordingTranscriptionService] method instead.
type RecordingTranscriptionService struct {
	Options []option.RequestOption
}

// NewRecordingTranscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecordingTranscriptionService(opts ...option.RequestOption) (r RecordingTranscriptionService) {
	r = RecordingTranscriptionService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing recording transcription.
func (r *RecordingTranscriptionService) Get(ctx context.Context, recordingTranscriptionID string, opts ...option.RequestOption) (res *RecordingTranscriptionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingTranscriptionID == "" {
		err = errors.New("missing required recording_transcription_id parameter")
		return
	}
	path := fmt.Sprintf("recording_transcriptions/%s", recordingTranscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your recording transcriptions.
func (r *RecordingTranscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *RecordingTranscriptionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "recording_transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a recording transcription.
func (r *RecordingTranscriptionService) Delete(ctx context.Context, recordingTranscriptionID string, opts ...option.RequestOption) (res *RecordingTranscriptionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordingTranscriptionID == "" {
		err = errors.New("missing required recording_transcription_id parameter")
		return
	}
	path := fmt.Sprintf("recording_transcriptions/%s", recordingTranscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RecordingTranscription struct {
	// Uniquely identifies the recording transcription.
	ID string `json:"id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The duration of the recording transcription in milliseconds.
	DurationMillis int64 `json:"duration_millis"`
	// Any of "recording_transcription".
	RecordType RecordingTranscriptionRecordType `json:"record_type"`
	// Uniquely identifies the recording associated with this transcription.
	RecordingID string `json:"recording_id"`
	// The status of the recording transcription. Only `completed` has transcription
	// text available.
	//
	// Any of "in-progress", "completed".
	Status RecordingTranscriptionStatus `json:"status"`
	// The recording's transcribed text.
	TranscriptionText string `json:"transcription_text"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		DurationMillis    respjson.Field
		RecordType        respjson.Field
		RecordingID       respjson.Field
		Status            respjson.Field
		TranscriptionText respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscription) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionRecordType string

const (
	RecordingTranscriptionRecordTypeRecordingTranscription RecordingTranscriptionRecordType = "recording_transcription"
)

// The status of the recording transcription. Only `completed` has transcription
// text available.
type RecordingTranscriptionStatus string

const (
	RecordingTranscriptionStatusInProgress RecordingTranscriptionStatus = "in-progress"
	RecordingTranscriptionStatusCompleted  RecordingTranscriptionStatus = "completed"
)

type RecordingTranscriptionGetResponse struct {
	Data RecordingTranscription `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionListResponse struct {
	Data []RecordingTranscription               `json:"data"`
	Meta RecordingTranscriptionListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionListResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionListResponseMeta struct {
	Cursors RecordingTranscriptionListResponseMetaCursors `json:"cursors"`
	// Path to next page.
	Next string `json:"next"`
	// Path to previous page.
	Previous string `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursors     respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionListResponseMetaCursors struct {
	// Opaque identifier of next page.
	After string `json:"after"`
	// Opaque identifier of previous page.
	Before string `json:"before"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionListResponseMetaCursors) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionListResponseMetaCursors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingTranscriptionDeleteResponse struct {
	Data RecordingTranscription `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingTranscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingTranscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
