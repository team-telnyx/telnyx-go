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

// AIIntegrationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIntegrationService] method instead.
type AIIntegrationService struct {
	Options     []option.RequestOption
	Connections AIIntegrationConnectionService
}

// NewAIIntegrationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIIntegrationService(opts ...option.RequestOption) (r AIIntegrationService) {
	r = AIIntegrationService{}
	r.Options = opts
	r.Connections = NewAIIntegrationConnectionService(opts...)
	return
}

// Retrieve integration details
func (r *AIIntegrationService) Get(ctx context.Context, integrationID string, opts ...option.RequestOption) (res *AIIntegrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return
	}
	path := fmt.Sprintf("ai/integrations/%s", integrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available integrations.
func (r *AIIntegrationService) List(ctx context.Context, opts ...option.RequestOption) (res *AIIntegrationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIIntegrationGetResponse struct {
	ID             string   `json:"id" api:"required"`
	AvailableTools []string `json:"available_tools" api:"required"`
	Description    string   `json:"description" api:"required"`
	DisplayName    string   `json:"display_name" api:"required"`
	LogoURL        string   `json:"logo_url" api:"required"`
	Name           string   `json:"name" api:"required"`
	// Any of "disconnected", "connected".
	Status AIIntegrationGetResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AvailableTools respjson.Field
		Description    respjson.Field
		DisplayName    respjson.Field
		LogoURL        respjson.Field
		Name           respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIIntegrationGetResponseStatus string

const (
	AIIntegrationGetResponseStatusDisconnected AIIntegrationGetResponseStatus = "disconnected"
	AIIntegrationGetResponseStatusConnected    AIIntegrationGetResponseStatus = "connected"
)

type AIIntegrationListResponse struct {
	Data []AIIntegrationListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIIntegrationListResponseData struct {
	ID             string   `json:"id" api:"required"`
	AvailableTools []string `json:"available_tools" api:"required"`
	Description    string   `json:"description" api:"required"`
	DisplayName    string   `json:"display_name" api:"required"`
	LogoURL        string   `json:"logo_url" api:"required"`
	Name           string   `json:"name" api:"required"`
	// Any of "disconnected", "connected".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AvailableTools respjson.Field
		Description    respjson.Field
		DisplayName    respjson.Field
		LogoURL        respjson.Field
		Name           respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIIntegrationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIIntegrationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
