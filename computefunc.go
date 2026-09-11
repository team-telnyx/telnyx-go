// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ComputeFuncService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFuncService] method instead.
type ComputeFuncService struct {
	Options []option.RequestOption
}

// NewComputeFuncService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeFuncService(opts ...option.RequestOption) (r ComputeFuncService) {
	r = ComputeFuncService{}
	r.Options = opts
	return
}

// Returns logs oldest first. `type=runtime` (default) returns function
// stdout/stderr. `type=invocations` returns one platform-generated record per HTTP
// request served.
func (r *ComputeFuncService) GetLogs(ctx context.Context, id string, query ComputeFuncGetLogsParams, opts ...option.RequestOption) (res *ComputeFuncGetLogsResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/logs", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns aggregate request, latency, CPU, memory, and resource-limit metrics for
// a function over the requested window.
func (r *ComputeFuncService) GetMetricAggregates(ctx context.Context, id string, query ComputeFuncGetMetricAggregatesParams, opts ...option.RequestOption) (res *ComputeFuncGetMetricAggregatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/metric_aggregates", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists a function's ship history newest first, including per-ship failure stage
// and reason when recorded.
func (r *ComputeFuncService) GetRevisions(ctx context.Context, id string, query ComputeFuncGetRevisionsParams, opts ...option.RequestOption) (res *ComputeFuncGetRevisionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/revisions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the latest ship outcome. The stage is `none` on success, `pending` while
// building, or a failure stage such as `build`, `platform`, `pre_build`, `deploy`,
// or `security_review`. This stage-neutral customer-facing path is an alias over
// the same inspection resource as `build_log_inspection`.
func (r *ComputeFuncService) GetShipInspection(ctx context.Context, id string, opts ...option.RequestOption) (res *ComputeFuncGetShipInspectionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/ship_inspection", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type FunctionsObservabilityPaginationMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionsObservabilityPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *FunctionsObservabilityPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogsMeta struct {
	HasMore bool `json:"has_more"`
	Partial bool `json:"partial"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Partial     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogsMeta) RawJSON() string { return r.JSON.raw }
func (r *LogsMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFuncGetLogsResponseUnion contains all possible properties and values from
// [ComputeFuncGetLogsResponseFuncRuntimeLogsResponse],
// [ComputeFuncGetLogsResponseFuncInvocationLogsResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ComputeFuncGetLogsResponseUnion struct {
	// This field is a union of
	// [[]ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData],
	// [[]ComputeFuncGetLogsResponseFuncInvocationLogsResponseData]
	Data ComputeFuncGetLogsResponseUnionData `json:"data"`
	// This field is from variant [ComputeFuncGetLogsResponseFuncRuntimeLogsResponse].
	Meta LogsMeta `json:"meta"`
	JSON struct {
		Data respjson.Field
		Meta respjson.Field
		raw  string
	} `json:"-"`
}

func (u ComputeFuncGetLogsResponseUnion) AsComputeFuncGetLogsResponseFuncRuntimeLogsResponse() (v ComputeFuncGetLogsResponseFuncRuntimeLogsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ComputeFuncGetLogsResponseUnion) AsComputeFuncGetLogsResponseFuncInvocationLogsResponse() (v ComputeFuncGetLogsResponseFuncInvocationLogsResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ComputeFuncGetLogsResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ComputeFuncGetLogsResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ComputeFuncGetLogsResponseUnionData is an implicit subunion of
// [ComputeFuncGetLogsResponseUnion]. ComputeFuncGetLogsResponseUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ComputeFuncGetLogsResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfComputeFuncGetLogsResponseFuncRuntimeLogsResponseData
// OfComputeFuncGetLogsResponseFuncInvocationLogsResponseData]
type ComputeFuncGetLogsResponseUnionData struct {
	// This field will be present if the value is a
	// [[]ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData] instead of an object.
	OfComputeFuncGetLogsResponseFuncRuntimeLogsResponseData []ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData `json:",inline"`
	// This field will be present if the value is a
	// [[]ComputeFuncGetLogsResponseFuncInvocationLogsResponseData] instead of an
	// object.
	OfComputeFuncGetLogsResponseFuncInvocationLogsResponseData []ComputeFuncGetLogsResponseFuncInvocationLogsResponseData `json:",inline"`
	JSON                                                       struct {
		OfComputeFuncGetLogsResponseFuncRuntimeLogsResponseData    respjson.Field
		OfComputeFuncGetLogsResponseFuncInvocationLogsResponseData respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (r *ComputeFuncGetLogsResponseUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetLogsResponseFuncRuntimeLogsResponse struct {
	Data []ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData `json:"data"`
	Meta LogsMeta                                                `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetLogsResponseFuncRuntimeLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetLogsResponseFuncRuntimeLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	// Any of "compute_func_runtime_log".
	RecordType string    `json:"record_type"`
	Timestamp  time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Message     respjson.Field
		RecordType  respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetLogsResponseFuncRuntimeLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetLogsResponseFuncInvocationLogsResponse struct {
	Data []ComputeFuncGetLogsResponseFuncInvocationLogsResponseData `json:"data"`
	Meta LogsMeta                                                   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetLogsResponseFuncInvocationLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetLogsResponseFuncInvocationLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetLogsResponseFuncInvocationLogsResponseData struct {
	DurationMs float64 `json:"duration_ms"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	// Any of "compute_func_invocation_log".
	RecordType        string    `json:"record_type"`
	Region            string    `json:"region"`
	RequestSizeBytes  int64     `json:"request_size_bytes"`
	ResponseSizeBytes int64     `json:"response_size_bytes"`
	StatusCode        int64     `json:"status_code"`
	Timestamp         time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs        respjson.Field
		Method            respjson.Field
		Path              respjson.Field
		RecordType        respjson.Field
		Region            respjson.Field
		RequestSizeBytes  respjson.Field
		ResponseSizeBytes respjson.Field
		StatusCode        respjson.Field
		Timestamp         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetLogsResponseFuncInvocationLogsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetLogsResponseFuncInvocationLogsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetMetricAggregatesResponse struct {
	Data []ComputeFuncGetMetricAggregatesResponseData `json:"data"`
	Meta FunctionsObservabilityPaginationMeta         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetMetricAggregatesResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetMetricAggregatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetMetricAggregatesResponseData struct {
	CPUUsedCoresAvg        float64   `json:"cpu_used_cores_avg" api:"nullable"`
	CPUUsedCoresMax        float64   `json:"cpu_used_cores_max" api:"nullable"`
	EndTime                time.Time `json:"end_time" format:"date-time"`
	FunctionID             string    `json:"function_id"`
	FunctionName           string    `json:"function_name"`
	MemoryUsedBytesAvg     float64   `json:"memory_used_bytes_avg" api:"nullable"`
	MemoryUsedBytesMax     float64   `json:"memory_used_bytes_max" api:"nullable"`
	Product                string    `json:"product"`
	RecordType             string    `json:"record_type"`
	RequestClientErrorRate float64   `json:"request_client_error_rate" api:"nullable"`
	RequestCount           float64   `json:"request_count" api:"nullable"`
	RequestErrorRate       float64   `json:"request_error_rate" api:"nullable"`
	RequestLatencyAvgMs    float64   `json:"request_latency_avg_ms" api:"nullable"`
	RequestLatencyP50Ms    float64   `json:"request_latency_p50_ms" api:"nullable"`
	RequestLatencyP95Ms    float64   `json:"request_latency_p95_ms" api:"nullable"`
	RequestLatencyP99Ms    float64   `json:"request_latency_p99_ms" api:"nullable"`
	RequestSuccessRate     float64   `json:"request_success_rate" api:"nullable"`
	StartTime              time.Time `json:"start_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CPUUsedCoresAvg        respjson.Field
		CPUUsedCoresMax        respjson.Field
		EndTime                respjson.Field
		FunctionID             respjson.Field
		FunctionName           respjson.Field
		MemoryUsedBytesAvg     respjson.Field
		MemoryUsedBytesMax     respjson.Field
		Product                respjson.Field
		RecordType             respjson.Field
		RequestClientErrorRate respjson.Field
		RequestCount           respjson.Field
		RequestErrorRate       respjson.Field
		RequestLatencyAvgMs    respjson.Field
		RequestLatencyP50Ms    respjson.Field
		RequestLatencyP95Ms    respjson.Field
		RequestLatencyP99Ms    respjson.Field
		RequestSuccessRate     respjson.Field
		StartTime              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetMetricAggregatesResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetMetricAggregatesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetRevisionsResponse struct {
	Data []ComputeFuncGetRevisionsResponseData `json:"data"`
	Meta FunctionsObservabilityPaginationMeta  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetRevisionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetRevisionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetRevisionsResponseData struct {
	Active        bool      `json:"active"`
	BuildOkAt     time.Time `json:"build_ok_at" format:"date-time"`
	BuildStatus   string    `json:"build_status"`
	CommitSha     string    `json:"commit_sha"`
	DeployStatus  string    `json:"deploy_status"`
	FailureReason string    `json:"failure_reason"`
	FailureStage  string    `json:"failure_stage"`
	Image         string    `json:"image"`
	RecordType    string    `json:"record_type"`
	RevisionID    string    `json:"revision_id"`
	ShippedAt     time.Time `json:"shipped_at" format:"date-time"`
	ShippedBy     string    `json:"shipped_by"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active        respjson.Field
		BuildOkAt     respjson.Field
		BuildStatus   respjson.Field
		CommitSha     respjson.Field
		DeployStatus  respjson.Field
		FailureReason respjson.Field
		FailureStage  respjson.Field
		Image         respjson.Field
		RecordType    respjson.Field
		RevisionID    respjson.Field
		ShippedAt     respjson.Field
		ShippedBy     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetRevisionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetRevisionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetShipInspectionResponse struct {
	Data ComputeFuncGetShipInspectionResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetShipInspectionResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetShipInspectionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetShipInspectionResponseData struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Reason    string    `json:"reason"`
	// Stable record type retained by both inspection path aliases.
	//
	// Any of "build_log_inspection".
	RecordType string `json:"record_type"`
	Runtime    string `json:"runtime"`
	Snippet    string `json:"snippet"`
	// Any of "build", "platform", "pre_build", "deploy", "security_review", "none",
	// "pending".
	Stage string `json:"stage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Reason      respjson.Field
		RecordType  respjson.Field
		Runtime     respjson.Field
		Snippet     respjson.Field
		Stage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeFuncGetShipInspectionResponseData) RawJSON() string { return r.JSON.raw }
func (r *ComputeFuncGetShipInspectionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncGetLogsParams struct {
	// Return records at or before this RFC 3339 timestamp.
	EndTime param.Opt[time.Time] `query:"end_time,omitzero" format:"date-time" json:"-"`
	// Maximum records to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Return records at or after this RFC 3339 timestamp.
	StartTime param.Opt[time.Time] `query:"start_time,omitzero" format:"date-time" json:"-"`
	// Log stream to return.
	//
	// Any of "runtime", "invocations".
	Type ComputeFuncGetLogsParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFuncGetLogsParams]'s query parameters as
// `url.Values`.
func (r ComputeFuncGetLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Log stream to return.
type ComputeFuncGetLogsParamsType string

const (
	ComputeFuncGetLogsParamsTypeRuntime     ComputeFuncGetLogsParamsType = "runtime"
	ComputeFuncGetLogsParamsTypeInvocations ComputeFuncGetLogsParamsType = "invocations"
)

type ComputeFuncGetMetricAggregatesParams struct {
	// Exclusive window end, UTC ISO 8601 with milliseconds
	EndTime time.Time `query:"end_time" api:"required" format:"date-time" json:"-"`
	// Inclusive window start, UTC ISO 8601 with milliseconds
	StartTime time.Time `query:"start_time" api:"required" format:"date-time" json:"-"`
	// Edge site filter
	FilterEdgeSite param.Opt[string] `query:"filter[edge_site],omitzero" json:"-"`
	// Kubernetes namespace filter
	FilterNamespace param.Opt[string] `query:"filter[namespace],omitzero" json:"-"`
	PageNumber      param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize        param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFuncGetMetricAggregatesParams]'s query parameters as
// `url.Values`.
func (r ComputeFuncGetMetricAggregatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeFuncGetRevisionsParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeFuncGetRevisionsParams]'s query parameters as
// `url.Values`.
func (r ComputeFuncGetRevisionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
