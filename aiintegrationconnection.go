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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIIntegrationConnectionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIntegrationConnectionService] method instead.
type AIIntegrationConnectionService struct {
	Options []option.RequestOption
}

// NewAIIntegrationConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIntegrationConnectionService(opts ...option.RequestOption) (r AIIntegrationConnectionService) {
	r = AIIntegrationConnectionService{}
	r.Options = opts
	return
}

// Get user setup integrations
func (r *AIIntegrationConnectionService) Get(ctx context.Context, userConnectionID string, opts ...option.RequestOption) (res *AIIntegrationConnectionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userConnectionID == "" {
		err = errors.New("missing required user_connection_id parameter")
		return
	}
	path := fmt.Sprintf("ai/integrations/connections/%s", userConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List user setup integrations
func (r *AIIntegrationConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *AIIntegrationConnectionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/integrations/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific integration connection.
func (r *AIIntegrationConnectionService) Delete(ctx context.Context, userConnectionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userConnectionID == "" {
		err = errors.New("missing required user_connection_id parameter")
		return
	}
	path := fmt.Sprintf("ai/integrations/connections/%s", userConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AIIntegrationConnectionGetResponse struct {
	Data AIIntegrationConnectionGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIIntegrationConnectionGetResponseData struct {
	ID            string   `json:"id" api:"required"`
	AllowedTools  []string `json:"allowed_tools" api:"required"`
	IntegrationID string   `json:"integration_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AllowedTools  respjson.Field
		IntegrationID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationConnectionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationConnectionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIIntegrationConnectionListResponse struct {
	Data []AIIntegrationConnectionListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIIntegrationConnectionListResponseData struct {
	ID            string   `json:"id" api:"required"`
	AllowedTools  []string `json:"allowed_tools" api:"required"`
	IntegrationID string   `json:"integration_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AllowedTools  respjson.Field
		IntegrationID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationConnectionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationConnectionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
