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

// Inspect Stateful Actor types
//
// ComputeService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeService] method instead.
type ComputeService struct {
	Options []option.RequestOption
	// Manage Edge Compute functions
	Funcs ComputeFuncService
}

// NewComputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewComputeService(opts ...option.RequestOption) (r ComputeService) {
	r = ComputeService{}
	r.Options = opts
	r.Funcs = NewComputeFuncService(opts...)
	return
}

// Retrieve one account-scoped Stateful Actor type. owner_function_id identifies
// the function that owns and ships the actor class; other functions may reference
// the same actor type through actor bindings.
func (r *ComputeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/actors/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ComputeGetResponse struct {
	Data ComputeGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGetResponseData struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	// ID of the function that owns and ships this actor type. Empty when the type is
	// unowned.
	OwnerFunctionID string `json:"owner_function_id"`
	// Last deployment fan-out report, not a live probe. Omitted on actor types without
	// a placement-aware deployment.
	Placement  ComputeGetResponseDataPlacement `json:"placement"`
	RecordType string                          `json:"record_type"`
	Status     string                          `json:"status"`
	Type       string                          `json:"type"`
	UpdatedAt  string                          `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		OwnerFunctionID respjson.Field
		Placement       respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Last deployment fan-out report, not a live probe. Omitted on actor types without
// a placement-aware deployment.
type ComputeGetResponseDataPlacement struct {
	ClusterResults []ComputeGetResponseDataPlacementClusterResult `json:"cluster_results"`
	// Public regions with successful deployment
	Clusters         []string `json:"clusters"`
	ExcludedClusters []string `json:"excluded_clusters"`
	// Optional ISO-8601 millisecond timestamp
	ObservedAt string `json:"observed_at"`
	Partial    bool   `json:"partial"`
	// Any of "everywhere", "single_site".
	Policy            string   `json:"policy"`
	UncheckedClusters []string `json:"unchecked_clusters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterResults    respjson.Field
		Clusters          respjson.Field
		ExcludedClusters  respjson.Field
		ObservedAt        respjson.Field
		Partial           respjson.Field
		Policy            respjson.Field
		UncheckedClusters respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGetResponseDataPlacement) RawJSON() string { return r.JSON.raw }
func (r *ComputeGetResponseDataPlacement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGetResponseDataPlacementClusterResult struct {
	// Public region, or unknown when no mapping exists
	ClusterID string `json:"cluster_id"`
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
	Status    string `json:"status"`
	Success   bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID   respjson.Field
		Error       respjson.Field
		ErrorCode   respjson.Field
		Status      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGetResponseDataPlacementClusterResult) RawJSON() string { return r.JSON.raw }
func (r *ComputeGetResponseDataPlacementClusterResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
