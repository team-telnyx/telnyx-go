// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIAssistantCanaryDeployService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantCanaryDeployService] method instead.
type AIAssistantCanaryDeployService struct {
	Options []option.RequestOption
}

// NewAIAssistantCanaryDeployService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantCanaryDeployService(opts ...option.RequestOption) (r AIAssistantCanaryDeployService) {
	r = AIAssistantCanaryDeployService{}
	r.Options = opts
	return
}

// Endpoint to create a canary deploy configuration for an assistant.
//
// Creates a new canary deploy configuration with multiple version IDs and their
// traffic percentages for A/B testing or gradual rollouts of assistant versions.
func (r *AIAssistantCanaryDeployService) New(ctx context.Context, assistantID string, body AIAssistantCanaryDeployNewParams, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Endpoint to get a canary deploy configuration for an assistant.
//
// Retrieves the current canary deploy configuration with all version IDs and their
// traffic percentages for the specified assistant.
func (r *AIAssistantCanaryDeployService) Get(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Endpoint to update a canary deploy configuration for an assistant.
//
// Updates the existing canary deploy configuration with new version IDs and
// percentages. All old versions and percentages are replaces by new ones from this
// request.
func (r *AIAssistantCanaryDeployService) Update(ctx context.Context, assistantID string, body AIAssistantCanaryDeployUpdateParams, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Endpoint to delete a canary deploy configuration for an assistant.
//
// Removes all canary deploy configurations for the specified assistant.
func (r *AIAssistantCanaryDeployService) Delete(ctx context.Context, assistantID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Request model for creating or updating canary deploys.
//
// The property Versions is required.
type CanaryDeployParam struct {
	// List of version configurations
	Versions []VersionConfigParam `json:"versions,omitzero,required"`
	paramObj
}

func (r CanaryDeployParam) MarshalJSON() (data []byte, err error) {
	type shadow CanaryDeployParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CanaryDeployParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model for canary deploy operations.
type CanaryDeployResponse struct {
	AssistantID string          `json:"assistant_id,required"`
	CreatedAt   time.Time       `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time       `json:"updated_at,required" format:"date-time"`
	Versions    []VersionConfig `json:"versions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CanaryDeployResponse) RawJSON() string { return r.JSON.raw }
func (r *CanaryDeployResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a single version in canary deploy.
type VersionConfig struct {
	// Percentage of traffic for this version [1-99]
	Percentage float64 `json:"percentage,required"`
	// Version ID string that references assistant_versions.version_id
	VersionID string `json:"version_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		VersionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionConfig) RawJSON() string { return r.JSON.raw }
func (r *VersionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VersionConfig to a VersionConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VersionConfigParam.Overrides()
func (r VersionConfig) ToParam() VersionConfigParam {
	return param.Override[VersionConfigParam](json.RawMessage(r.RawJSON()))
}

// Configuration for a single version in canary deploy.
//
// The properties Percentage, VersionID are required.
type VersionConfigParam struct {
	// Percentage of traffic for this version [1-99]
	Percentage float64 `json:"percentage,required"`
	// Version ID string that references assistant_versions.version_id
	VersionID string `json:"version_id,required"`
	paramObj
}

func (r VersionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow VersionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantCanaryDeployNewParams struct {
	// Request model for creating or updating canary deploys.
	CanaryDeploy CanaryDeployParam
	paramObj
}

func (r AIAssistantCanaryDeployNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CanaryDeploy)
}
func (r *AIAssistantCanaryDeployNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CanaryDeploy)
}

type AIAssistantCanaryDeployUpdateParams struct {
	// Request model for creating or updating canary deploys.
	CanaryDeploy CanaryDeployParam
	paramObj
}

func (r AIAssistantCanaryDeployUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CanaryDeploy)
}
func (r *AIAssistantCanaryDeployUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CanaryDeploy)
}
