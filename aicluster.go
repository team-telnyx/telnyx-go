// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// AIClusterService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIClusterService] method instead.
type AIClusterService struct {
	Options []option.RequestOption
}

// NewAIClusterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIClusterService(opts ...option.RequestOption) (r AIClusterService) {
	r = AIClusterService{}
	r.Options = opts
	return
}

// Fetch a cluster
func (r *AIClusterService) Get(ctx context.Context, taskID string, query AIClusterGetParams, opts ...option.RequestOption) (res *AIClusterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("ai/clusters/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List all clusters
func (r *AIClusterService) List(ctx context.Context, query AIClusterListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIClusterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/clusters"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all clusters
func (r *AIClusterService) ListAutoPaging(ctx context.Context, query AIClusterListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIClusterListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a cluster
func (r *AIClusterService) Delete(ctx context.Context, taskID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("ai/clusters/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Starts a background task to compute how the data in an
// [embedded storage bucket](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding)
// is clustered. This helps identify common themes and patterns in the data.
func (r *AIClusterService) Compute(ctx context.Context, body AIClusterComputeParams, opts ...option.RequestOption) (res *AIClusterComputeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/clusters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a cluster visualization
func (r *AIClusterService) FetchGraph(ctx context.Context, taskID string, query AIClusterFetchGraphParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("ai/clusters/%s/graph", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RecursiveCluster struct {
	ClusterID          string                 `json:"cluster_id,required"`
	ClusterSummary     string                 `json:"cluster_summary,required"`
	TotalNumberOfNodes int64                  `json:"total_number_of_nodes,required"`
	ClusterHeader      string                 `json:"cluster_header"`
	Nodes              []RecursiveClusterNode `json:"nodes"`
	Subclusters        []RecursiveCluster     `json:"subclusters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClusterID          respjson.Field
		ClusterSummary     respjson.Field
		TotalNumberOfNodes respjson.Field
		ClusterHeader      respjson.Field
		Nodes              respjson.Field
		Subclusters        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecursiveCluster) RawJSON() string { return r.JSON.raw }
func (r *RecursiveCluster) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecursiveClusterNode struct {
	// The corresponding source file of your embedded storage bucket that the node is
	// from.
	Filename string `json:"filename,required"`
	// The text of the node.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecursiveClusterNode) RawJSON() string { return r.JSON.raw }
func (r *RecursiveClusterNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterGetResponse struct {
	Data AIClusterGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIClusterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIClusterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterGetResponseData struct {
	Bucket   string             `json:"bucket,required"`
	Clusters []RecursiveCluster `json:"clusters,required"`
	// Any of "pending", "starting", "running", "completed", "failed".
	Status TaskStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Clusters    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIClusterGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIClusterGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterListResponse struct {
	Bucket            string    `json:"bucket,required"`
	CreatedAt         time.Time `json:"created_at,required" format:"date-time"`
	FinishedAt        time.Time `json:"finished_at,required" format:"date-time"`
	MinClusterSize    int64     `json:"min_cluster_size,required"`
	MinSubclusterSize int64     `json:"min_subcluster_size,required"`
	// Any of "pending", "starting", "running", "completed", "failed".
	Status TaskStatus `json:"status,required"`
	TaskID string     `json:"task_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket            respjson.Field
		CreatedAt         respjson.Field
		FinishedAt        respjson.Field
		MinClusterSize    respjson.Field
		MinSubclusterSize respjson.Field
		Status            respjson.Field
		TaskID            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIClusterListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIClusterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterComputeResponse struct {
	Data AIClusterComputeResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIClusterComputeResponse) RawJSON() string { return r.JSON.raw }
func (r *AIClusterComputeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterComputeResponseData struct {
	TaskID string `json:"task_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIClusterComputeResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIClusterComputeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterGetParams struct {
	// Whether or not to include subclusters and their nodes in the response.
	ShowSubclusters param.Opt[bool] `query:"show_subclusters,omitzero" json:"-"`
	// The number of nodes in the cluster to return in the response. Nodes will be
	// sorted by their centrality within the cluster.
	TopNNodes param.Opt[int64] `query:"top_n_nodes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIClusterGetParams]'s query parameters as `url.Values`.
func (r AIClusterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIClusterListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIClusterListParams]'s query parameters as `url.Values`.
func (r AIClusterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIClusterComputeParams struct {
	// The embedded storage bucket to compute the clusters from. The bucket must
	// already be
	// [embedded](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding).
	Bucket string `json:"bucket,required"`
	// Smallest number of related text chunks to qualify as a cluster. Top-level
	// clusters should be thought of as identifying broad themes in your data.
	MinClusterSize param.Opt[int64] `json:"min_cluster_size,omitzero"`
	// Smallest number of related text chunks to qualify as a sub-cluster. Sub-clusters
	// should be thought of as identifying more specific topics within a broader theme.
	MinSubclusterSize param.Opt[int64] `json:"min_subcluster_size,omitzero"`
	// Prefix to filter whcih files in the buckets are included.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Array of files to filter which are included.
	Files []string `json:"files,omitzero"`
	paramObj
}

func (r AIClusterComputeParams) MarshalJSON() (data []byte, err error) {
	type shadow AIClusterComputeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIClusterComputeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIClusterFetchGraphParams struct {
	ClusterID param.Opt[int64] `query:"cluster_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIClusterFetchGraphParams]'s query parameters as
// `url.Values`.
func (r AIClusterFetchGraphParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
