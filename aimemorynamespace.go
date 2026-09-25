// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Whether a write has finished.
//
// AIMemoryNamespaceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryNamespaceService] method instead.
type AIMemoryNamespaceService struct {
	Options  []option.RequestOption
	Profiles AIMemoryNamespaceProfileService
	// How a namespace's summaries are written.
	Settings AIMemoryNamespaceSettingService
}

// NewAIMemoryNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMemoryNamespaceService(opts ...option.RequestOption) (r AIMemoryNamespaceService) {
	r = AIMemoryNamespaceService{}
	r.Options = opts
	r.Profiles = NewAIMemoryNamespaceProfileService(opts...)
	r.Settings = NewAIMemoryNamespaceSettingService(opts...)
	return
}

// Whether a write has finished. Both `ingest` and `remember` return an
// `operation_id`, and a memory is not recallable until its operation completes —
// extraction, embedding and consolidation all run first.
func (r *AIMemoryNamespaceService) Get(ctx context.Context, operationID string, query AIMemoryNamespaceGetParams, opts ...option.RequestOption) (res *AIMemoryNamespaceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/operations/%s", url.PathEscape(query.Namespace), url.PathEscape(operationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AIMemoryNamespaceGetResponse struct {
	Data AIMemoryNamespaceGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceGetResponseData struct {
	OperationID string `json:"operation_id" api:"required"`
	// Where the write is. `completed`, `failed` and `cancelled` are terminal: stop
	// polling at any of them, and treat `failed` and `cancelled` as writes that did
	// not happen.
	//
	// Any of "pending", "processing", "completed", "failed", "cancelled".
	Status      string `json:"status" api:"required"`
	CompletedAt string `json:"completed_at" api:"nullable"`
	CreatedAt   string `json:"created_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OperationID respjson.Field
		Status      respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceGetParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`
	paramObj
}
