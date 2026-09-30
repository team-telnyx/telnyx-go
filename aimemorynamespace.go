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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

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

// Create a namespace. An organization can have at most five, `default` among them
// — a sixth returns `403`.
func (r *AIMemoryNamespaceService) New(ctx context.Context, body AIMemoryNamespaceNewParams, opts ...option.RequestOption) (res *AIMemoryNamespaceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/memory/namespaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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

// Every namespace in your organization, `default` among them.
func (r *AIMemoryNamespaceService) List(ctx context.Context, opts ...option.RequestOption) (res *AIMemoryNamespaceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/memory/namespaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a namespace and every profile and memory in it. `default` cannot be
// deleted. This cannot be undone.
func (r *AIMemoryNamespaceService) Delete(ctx context.Context, namespace string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s", url.PathEscape(namespace))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// An isolated memory store within your organization.
type Namespace struct {
	// The namespace's unique identifier.
	ID string `json:"id" api:"required"`
	// The namespace's name, used in the path. `default` exists for every organization.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Namespace) RawJSON() string { return r.JSON.raw }
func (r *Namespace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceNewResponse struct {
	// An isolated memory store within your organization.
	Data Namespace `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type AIMemoryNamespaceListResponse struct {
	Data []Namespace `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceNewParams struct {
	// A name for the new namespace, unique within your organization.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r AIMemoryNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMemoryNamespaceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMemoryNamespaceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceGetParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`
	paramObj
}
