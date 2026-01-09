// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIAssistantToolService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantToolService] method instead.
type AIAssistantToolService struct {
	Options []option.RequestOption
}

// NewAIAssistantToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAssistantToolService(opts ...option.RequestOption) (r AIAssistantToolService) {
	r = AIAssistantToolService{}
	r.Options = opts
	return
}

// Test a webhook tool for an assistant
func (r *AIAssistantToolService) Test(ctx context.Context, toolID string, params AIAssistantToolTestParams, opts ...option.RequestOption) (res *AIAssistantToolTestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/tools/%s/test", params.AssistantID, toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response model for webhook tool test results
type AIAssistantToolTestResponse struct {
	// Response model for webhook tool test results
	Data AIAssistantToolTestResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantToolTestResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantToolTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for webhook tool test results
type AIAssistantToolTestResponseData struct {
	ContentType string         `json:"content_type,required"`
	Request     map[string]any `json:"request,required"`
	Response    string         `json:"response,required"`
	StatusCode  int64          `json:"status_code,required"`
	Success     bool           `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Request     respjson.Field
		Response    respjson.Field
		StatusCode  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantToolTestResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantToolTestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantToolTestParams struct {
	AssistantID string `path:"assistant_id,required" json:"-"`
	// Key-value arguments to use for the webhook test
	Arguments map[string]any `json:"arguments,omitzero"`
	// Key-value dynamic variables to use for the webhook test
	DynamicVariables map[string]any `json:"dynamic_variables,omitzero"`
	paramObj
}

func (r AIAssistantToolTestParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantToolTestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantToolTestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
