// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
)

// Configure AI assistant specifications
//
// AIAssistantInstructionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantInstructionService] method instead.
type AIAssistantInstructionService struct {
	Options []option.RequestOption
}

// NewAIAssistantInstructionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantInstructionService(opts ...option.RequestOption) (r AIAssistantInstructionService) {
	r = AIAssistantInstructionService{}
	r.Options = opts
	return
}

// Enhance an assistant's instructions using an LLM. The endpoint reads the
// assistant's current instructions and tools, then streams back improved
// instructions as they are generated.
//
// Optionally provide an `enhancement_prompt` to steer the changes (for example,
// "make the instructions more concise" or "add error handling guidance"). When
// omitted, the assistant's existing instructions are used as the basis for the
// enhancement.
//
// The enhancement focuses on tool-calling reliability, clarity and precision,
// completeness and error handling, tool schema alignment, and conversation flow
// structure.
//
// The response is streamed as `text/plain` using chunked transfer encoding;
// consume the body incrementally to render the enhanced instructions as they
// arrive.
func (r *AIAssistantInstructionService) Enhance(ctx context.Context, assistantID string, body AIAssistantInstructionEnhanceParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/instructions/enhance", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AIAssistantInstructionEnhanceParams struct {
	// Optional guidance describing how the instructions should be enhanced. When
	// provided, the LLM applies these requested changes in addition to fixing any
	// identified issues.
	EnhancementPrompt param.Opt[string] `json:"enhancement_prompt,omitzero"`
	// The instructions to enhance. When omitted, the assistant's existing instructions
	// are used.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r AIAssistantInstructionEnhanceParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantInstructionEnhanceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantInstructionEnhanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
