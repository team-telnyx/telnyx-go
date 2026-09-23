// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
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
	return res, err
}

// Word-level timing detail. Only present when using a `deepgram/*` model with
// `model_config` options that enable word timestamps.
type AudioTranscriptionResponseWord struct {
	// End time of the word in seconds.
	End float64 `json:"end" api:"required"`
	// Start time of the word in seconds.
	Start float64 `json:"start" api:"required"`
	// The transcribed word.
	Word string `json:"word" api:"required"`
	// Confidence score for the word (0.0 to 1.0).
	Confidence float64 `json:"confidence"`
	// The transcribed word with punctuation and capitalisation applied. Only present
	// when `punctuate` or `smart_format` is enabled via `model_config`.
	PunctuatedWord string `json:"punctuated_word"`
	// Speaker index. Only present when diarization is enabled via `model_config`.
	Speaker int64 `json:"speaker"`
	// Confidence score for the speaker assignment (0.0 to 1.0). Only present when
	// diarization is enabled via `model_config`.
	SpeakerConfidence float64 `json:"speaker_confidence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End               respjson.Field
		Start             respjson.Field
		Word              respjson.Field
		Confidence        respjson.Field
		PunctuatedWord    respjson.Field
		Speaker           respjson.Field
		SpeakerConfidence respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioTranscriptionResponseWord) RawJSON() string { return r.JSON.raw }
func (r *AudioTranscriptionResponseWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response fields vary by model. `distil-whisper/distil-large-v2` returns `text`,
// `duration`, and `segments` in `verbose_json` mode.
// `openai/whisper-large-v3-turbo` returns `text` only. The `deepgram/*` models
// return `text` and, depending on `model_config`, may include `words` with
// per-word timestamps and speaker labels. The Parakeet models
// (`nvidia/parakeet-v3`, `omi-health/omi-med-stt-v1`) return `text` only.
type AIAudioTranscribeResponse struct {
	// The transcribed text for the audio file.
	Text string `json:"text" api:"required"`
	// The duration of the audio file in seconds. Returned by
	// `distil-whisper/distil-large-v2` and the `deepgram/*` models when
	// `response_format` is `verbose_json`. Not returned by
	// `openai/whisper-large-v3-turbo`.
	Duration float64 `json:"duration"`
	// Segments of the transcribed text and their corresponding details. Returned by
	// `distil-whisper/distil-large-v2` and the `deepgram/*` models when
	// `response_format` is `verbose_json`; Deepgram segments also carry nested `words`
	// and `speakers`. Not returned by `openai/whisper-large-v3-turbo`.
	Segments []AIAudioTranscribeResponseSegment `json:"segments"`
	// Word-level timestamps and optional speaker labels. Only returned by the
	// `deepgram/*` models when word-level output is enabled via `model_config`.
	Words []AudioTranscriptionResponseWord `json:"words"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Duration    respjson.Field
		Segments    respjson.Field
		Words       respjson.Field
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
	ID float64 `json:"id" api:"required"`
	// End time of the segment in seconds.
	End float64 `json:"end" api:"required"`
	// Start time of the segment in seconds.
	Start float64 `json:"start" api:"required"`
	// Text content of the segment.
	Text string `json:"text" api:"required"`
	// Speaker indices heard in this segment. Returned by the `deepgram/*` models when
	// `diarize` is enabled via `model_config`.
	Speakers []int64 `json:"speakers"`
	// Word-level timing detail for this segment. Returned by the `deepgram/*` models
	// when word-level output is enabled via `model_config`.
	Words []AudioTranscriptionResponseWord `json:"words"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		End         respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		Speakers    respjson.Field
		Words       respjson.Field
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
	// higher latency. The `deepgram/*` models only accept mp3/wav files:
	// `deepgram/nova-3` covers ~49 languages plus `multi` and `deepgram/nova-2` covers
	// ~33, while the `-medical` variants are tuned for clinical vocabulary and accept
	// English only (`en` and its regional variants, e.g. `en-US`, `en-GB`).
	// `nvidia/parakeet-v3` is multilingual with automatic language detection;
	// `omi-health/omi-med-stt-v1` is a medical model, English only.
	//
	// Any of "distil-whisper/distil-large-v2", "openai/whisper-large-v3-turbo",
	// "deepgram/nova-2", "deepgram/nova-2-medical", "deepgram/nova-3",
	// "deepgram/nova-3-medical", "nvidia/parakeet-v3", "omi-health/omi-med-stt-v1".
	Model AIAudioTranscribeParamsModel `json:"model,omitzero" api:"required"`
	// Link to audio file in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a,
	// ogg, wav, or webm. Support for hosted files is limited to 100MB. Cannot be used
	// together with `file`. Note: the `deepgram/*` models only support mp3 and wav
	// formats.
	FileURL param.Opt[string] `json:"file_url,omitzero"`
	// The language of the audio to be transcribed. `deepgram/nova-3` supports ~49
	// languages plus `multi`, and `deepgram/nova-2` supports ~33 plus `multi`; the
	// `-medical` variants are English only (`en` and its regional variants, e.g.
	// `en-US`, `en-GB`). Deepgram models validate on the base language and forward the
	// full tag, so regional variants such as `de-CH` and `pt-BR` are accepted where
	// the base language is supported; an unsupported language returns a 400. For
	// `openai/whisper-large-v3-turbo`, supports multiple languages.
	// `distil-whisper/distil-large-v2` does not support language parameter.
	// `nvidia/parakeet-v3` detects the language automatically;
	// `omi-health/omi-med-stt-v1` is English only.
	Language param.Opt[string] `json:"language,omitzero"`
	// The audio file object to transcribe, in one of these formats: flac, mp3, mp4,
	// mpeg, mpga, m4a, ogg, wav, or webm. File uploads are limited to 100 MB. Cannot
	// be used together with `file_url`. Note: the `deepgram/*` models only support mp3
	// and wav formats.
	File io.Reader `json:"file,omitzero" format:"binary"`
	// Additional model-specific configuration parameters. Only allowed with the
	// `deepgram/*` models. Can include Deepgram-specific options such as
	// `smart_format`, `punctuate`, `diarize`, `utterance`, `numerals`, and `language`.
	// If `language` is provided both as a top-level parameter and in `model_config`,
	// the top-level parameter takes precedence.
	ModelConfig map[string]any `json:"model_config,omitzero"`
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
// higher latency. The `deepgram/*` models only accept mp3/wav files:
// `deepgram/nova-3` covers ~49 languages plus `multi` and `deepgram/nova-2` covers
// ~33, while the `-medical` variants are tuned for clinical vocabulary and accept
// English only (`en` and its regional variants, e.g. `en-US`, `en-GB`).
// `nvidia/parakeet-v3` is multilingual with automatic language detection;
// `omi-health/omi-med-stt-v1` is a medical model, English only.
type AIAudioTranscribeParamsModel string

const (
	AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2 AIAudioTranscribeParamsModel = "distil-whisper/distil-large-v2"
	AIAudioTranscribeParamsModelOpenAIWhisperLargeV3Turbo  AIAudioTranscribeParamsModel = "openai/whisper-large-v3-turbo"
	AIAudioTranscribeParamsModelDeepgramNova2              AIAudioTranscribeParamsModel = "deepgram/nova-2"
	AIAudioTranscribeParamsModelDeepgramNova2Medical       AIAudioTranscribeParamsModel = "deepgram/nova-2-medical"
	AIAudioTranscribeParamsModelDeepgramNova3              AIAudioTranscribeParamsModel = "deepgram/nova-3"
	AIAudioTranscribeParamsModelDeepgramNova3Medical       AIAudioTranscribeParamsModel = "deepgram/nova-3-medical"
	AIAudioTranscribeParamsModelNvidiaParakeetV3           AIAudioTranscribeParamsModel = "nvidia/parakeet-v3"
	AIAudioTranscribeParamsModelOmiHealthOmiMedSttV1       AIAudioTranscribeParamsModel = "omi-health/omi-med-stt-v1"
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
