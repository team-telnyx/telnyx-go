// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIOpenAIChatService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIChatService] method instead.
type AIOpenAIChatService struct {
	Options []option.RequestOption
}

// NewAIOpenAIChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIChatService(opts ...option.RequestOption) (r AIOpenAIChatService) {
	r = AIOpenAIChatService{}
	r.Options = opts
	return
}

// Chat with a language model. This endpoint is consistent with the
// [OpenAI Chat Completions API](https://platform.openai.com/docs/api-reference/chat)
// and may be used with the OpenAI JS or Python SDK by setting the base URL to
// `https://api.telnyx.com/v2/ai/openai`.
func (r *AIOpenAIChatService) NewCompletion(ctx context.Context, body AIOpenAIChatNewCompletionParams, opts ...option.RequestOption) (res *AIOpenAIChatNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FunctionDefinition struct {
	Name        string         `json:"name" api:"required"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunctionDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionDefinition to a FunctionDefinitionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionDefinitionParam.Overrides()
func (r FunctionDefinition) ToParam() FunctionDefinitionParam {
	return param.Override[FunctionDefinitionParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type FunctionDefinitionParam struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Parameters  map[string]any    `json:"parameters,omitzero"`
	paramObj
}

func (r FunctionDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAIChatNewCompletionResponse map[string]any

type AIOpenAIChatNewCompletionParams struct {
	ChatCompletionRequest ChatCompletionRequestParam
	paramObj
}

func (r AIOpenAIChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChatCompletionRequest)
}
func (r *AIOpenAIChatNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
