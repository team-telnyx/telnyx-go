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

// Configure AI assistant specifications
//
// AIAssistantTagService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantTagService] method instead.
type AIAssistantTagService struct {
	Options []option.RequestOption
}

// NewAIAssistantTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAssistantTagService(opts ...option.RequestOption) (r AIAssistantTagService) {
	r = AIAssistantTagService{}
	r.Options = opts
	return
}

// Add Assistant Tag
func (r *AIAssistantTagService) New(ctx context.Context, assistantID string, body AIAssistantTagNewParams, opts ...option.RequestOption) (res *AIAssistantTagNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/tags", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get All Tags
func (r *AIAssistantTagService) List(ctx context.Context, opts ...option.RequestOption) (res *AIAssistantTagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Remove Assistant Tag
func (r *AIAssistantTagService) Delete(ctx context.Context, tag string, body AIAssistantTagDeleteParams, opts ...option.RequestOption) (res *AIAssistantTagDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/tags/%s", body.AssistantID, tag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AIAssistantTagNewResponse struct {
	Tags []string `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantTagNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantTagNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTagListResponse struct {
	Tags []string `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantTagListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantTagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTagDeleteResponse struct {
	Tags []string `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantTagDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantTagDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTagNewParams struct {
	Tag string `json:"tag" api:"required"`
	paramObj
}

func (r AIAssistantTagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTagDeleteParams struct {
	AssistantID string `path:"assistant_id" api:"required" json:"-"`
	paramObj
}
