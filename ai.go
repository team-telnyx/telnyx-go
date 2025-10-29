// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// AIService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options       []option.RequestOption
	Assistants    AIAssistantService
	Audio         AIAudioService
	Chat          AIChatService
	Clusters      AIClusterService
	Conversations AIConversationService
	Embeddings    AIEmbeddingService
	FineTuning    AIFineTuningService
	Integrations  AIIntegrationService
	McpServers    AIMcpServerService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.Options = opts
	r.Assistants = NewAIAssistantService(opts...)
	r.Audio = NewAIAudioService(opts...)
	r.Chat = NewAIChatService(opts...)
	r.Clusters = NewAIClusterService(opts...)
	r.Conversations = NewAIConversationService(opts...)
	r.Embeddings = NewAIEmbeddingService(opts...)
	r.FineTuning = NewAIFineTuningService(opts...)
	r.Integrations = NewAIIntegrationService(opts...)
	r.McpServers = NewAIMcpServerService(opts...)
	return
}

// This endpoint returns a list of Open Source and OpenAI models that are available
// for use. <br /><br /> **Note**: Model `id`'s will be in the form
// `{source}/{model_name}`. For example `openai/gpt-4` or
// `mistralai/Mistral-7B-Instruct-v0.1` consistent with HuggingFace naming
// conventions.
func (r *AIService) GetModels(ctx context.Context, opts ...option.RequestOption) (res *AIGetModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Generate a summary of a file's contents.
//
// Supports the following text formats:
//
// - PDF, HTML, txt, json, csv
//
// Supports the following media formats (billed for both the transcription and
// summary):
//
// - flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm
// - Up to 100 MB
func (r *AIService) Summarize(ctx context.Context, body AISummarizeParams, opts ...option.RequestOption) (res *AISummarizeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/summarize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIGetModelsResponse struct {
	Data   []AIGetModelsResponseData `json:"data,required"`
	Object string                    `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIGetModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIGetModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIGetModelsResponseData struct {
	ID      string `json:"id,required"`
	Created int64  `json:"created,required"`
	OwnedBy string `json:"owned_by,required"`
	Object  string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		OwnedBy     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIGetModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIGetModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeResponse struct {
	Data AISummarizeResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISummarizeResponse) RawJSON() string { return r.JSON.raw }
func (r *AISummarizeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeResponseData struct {
	Summary string `json:"summary,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISummarizeResponseData) RawJSON() string { return r.JSON.raw }
func (r *AISummarizeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeParams struct {
	// The name of the bucket that contains the file to be summarized.
	Bucket string `json:"bucket,required"`
	// The name of the file to be summarized.
	Filename string `json:"filename,required"`
	// A system prompt to guide the summary generation.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	paramObj
}

func (r AISummarizeParams) MarshalJSON() (data []byte, err error) {
	type shadow AISummarizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AISummarizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
