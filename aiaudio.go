// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// AIAudioService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAudioService] method instead.
type AIAudioService struct {
	Options []option.RequestOption
}

// NewAIAudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIAudioService(opts ...option.RequestOption) (r AIAudioService) {
	r = AIAudioService{}
	r.Options = opts
	return
}

// Transcribe speech to text. This endpoint is consistent with the
// [OpenAI Transcription API](https://platform.openai.com/docs/api-reference/audio/createTranscription)
// and may be used with the OpenAI JS or Python SDK.
func (r *AIAudioService) Transcribe(ctx context.Context, body AIAudioTranscribeParams, opts ...option.RequestOption) (res *AIAudioTranscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/audio/transcriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIAudioTranscribeResponse struct {
	// The transcribed text for the audio file.
	Text string `json:"text,required"`
	// The duration of the audio file in seconds. This is only included if
	// `response_format` is set to `verbose_json`.
	Duration float64 `json:"duration"`
	// Segments of the transcribed text and their corresponding details. This is only
	// included if `response_format` is set to `verbose_json`.
	Segments []AIAudioTranscribeResponseSegment `json:"segments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Duration    respjson.Field
		Segments    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAudioTranscribeResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAudioTranscribeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAudioTranscribeResponseSegment struct {
	// Unique identifier of the segment.
	ID float64 `json:"id,required"`
	// End time of the segment in seconds.
	End float64 `json:"end,required"`
	// Start time of the segment in seconds.
	Start float64 `json:"start,required"`
	// Text content of the segment.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		End         respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAudioTranscribeResponseSegment) RawJSON() string { return r.JSON.raw }
func (r *AIAudioTranscribeResponseSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAudioTranscribeParams struct {
	// ID of the model to use. `distil-whisper/distil-large-v2` is lower latency but
	// English-only. `openai/whisper-large-v3-turbo` is multi-lingual but slightly
	// higher latency.
	//
	// Any of "distil-whisper/distil-large-v2", "openai/whisper-large-v3-turbo".
	Model AIAudioTranscribeParamsModel `json:"model,omitzero,required"`
	// Link to audio file in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a,
	// ogg, wav, or webm. Support for hosted files is limited to 100MB. Cannot be used
	// together with `file`
	FileURL param.Opt[string] `json:"file_url,omitzero"`
	// The audio file object to transcribe, in one of these formats: flac, mp3, mp4,
	// mpeg, mpga, m4a, ogg, wav, or webm. File uploads are limited to 100 MB. Cannot
	// be used together with `file_url`
	File io.Reader `json:"file,omitzero" format:"binary"`
	// The format of the transcript output. Use `verbose_json` to take advantage of
	// timestamps.
	//
	// Any of "json", "verbose_json".
	ResponseFormat AIAudioTranscribeParamsResponseFormat `json:"response_format,omitzero"`
	// The timestamp granularities to populate for this transcription.
	// `response_format` must be set verbose_json to use timestamp granularities.
	// Currently `segment` is supported.
	//
	// Any of "segment".
	TimestampGranularities AIAudioTranscribeParamsTimestampGranularities `json:"timestamp_granularities[],omitzero"`
	paramObj
}

func (r AIAudioTranscribeParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// ID of the model to use. `distil-whisper/distil-large-v2` is lower latency but
// English-only. `openai/whisper-large-v3-turbo` is multi-lingual but slightly
// higher latency.
type AIAudioTranscribeParamsModel string

const (
	AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2 AIAudioTranscribeParamsModel = "distil-whisper/distil-large-v2"
	AIAudioTranscribeParamsModelOpenAIWhisperLargeV3Turbo  AIAudioTranscribeParamsModel = "openai/whisper-large-v3-turbo"
)

// The format of the transcript output. Use `verbose_json` to take advantage of
// timestamps.
type AIAudioTranscribeParamsResponseFormat string

const (
	AIAudioTranscribeParamsResponseFormatJson        AIAudioTranscribeParamsResponseFormat = "json"
	AIAudioTranscribeParamsResponseFormatVerboseJson AIAudioTranscribeParamsResponseFormat = "verbose_json"
)

// The timestamp granularities to populate for this transcription.
// `response_format` must be set verbose_json to use timestamp granularities.
// Currently `segment` is supported.
type AIAudioTranscribeParamsTimestampGranularities string

const (
	AIAudioTranscribeParamsTimestampGranularitiesSegment AIAudioTranscribeParamsTimestampGranularities = "segment"
)
