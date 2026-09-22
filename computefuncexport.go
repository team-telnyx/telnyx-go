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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ComputeFuncExportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeFuncExportService] method instead.
type ComputeFuncExportService struct {
	Options []option.RequestOption
}

// NewComputeFuncExportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewComputeFuncExportService(opts ...option.RequestOption) (r ComputeFuncExportService) {
	r = ComputeFuncExportService{}
	r.Options = opts
	return
}

// Configures the external OTLP endpoint a function's runtime and/or invocation
// logs are pushed to as they happen. This operation is a **full replace, not a
// patch**: `endpoint`, `headers`, `runtime_export_enabled`, and
// `invocation_export_enabled` are all required on every call — omitting any of
// them is a 422, not "keep the current value". Headers are encrypted at rest and
// never returned in any response.
//
// The endpoint must be an HTTPS URL. When export is configured, new log records
// are converted to OTLP log records and delivered continuously; export never
// bypasses platform log storage, and delivery retries with a bounded policy while
// the destination is unreachable. Only logs generated after configuration are
// exported — there is no historical replay.
func (r *ComputeFuncExportService) New(ctx context.Context, id string, body ComputeFuncExportNewParams, opts ...option.RequestOption) (res *FuncLogExportConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/logs/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the function's configured log export destination and which log types are
// exported. Headers are never returned. Returns 404 (error code 10005) when no
// destination is configured for the function.
func (r *ComputeFuncExportService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *FuncLogExportConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/funcs/%s/logs/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stops exporting a function's logs and removes its destination configuration.
// Idempotent: deleting when nothing is configured succeeds.
func (r *ComputeFuncExportService) DeleteAll(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("compute/funcs/%s/logs/export", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type FuncLogExportConfigResponse struct {
	// Metadata-only view of a function's log export destination. Header values are
	// write-only (encrypted server-side) and never appear in any response.
	Data FuncLogExportConfigResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FuncLogExportConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *FuncLogExportConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata-only view of a function's log export destination. Header values are
// write-only (encrypted server-side) and never appear in any response.
type FuncLogExportConfigResponseData struct {
	// Configuration record ID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether export is enabled for this function
	Enabled bool `json:"enabled"`
	// HTTPS OTLP endpoint URL logs are pushed to
	Endpoint string `json:"endpoint" format:"uri"`
	// Function ID this configuration belongs to
	FuncID string `json:"func_id"`
	// Whether invocation records (one per HTTP request) are exported
	InvocationExportEnabled bool `json:"invocation_export_enabled"`
	// Any of "compute_func_log_export_config".
	RecordType string `json:"record_type"`
	// Whether runtime logs (function stdout/stderr) are exported
	RuntimeExportEnabled bool      `json:"runtime_export_enabled"`
	UpdatedAt            time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Enabled                 respjson.Field
		Endpoint                respjson.Field
		FuncID                  respjson.Field
		InvocationExportEnabled respjson.Field
		RecordType              respjson.Field
		RuntimeExportEnabled    respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FuncLogExportConfigResponseData) RawJSON() string { return r.JSON.raw }
func (r *FuncLogExportConfigResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeFuncExportNewParams struct {
	// HTTPS URL to push logs to
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// Headers attached to every export push, as key-value pairs (e.g. an auth token
	// the collector expects). Required even when empty — {} means "no headers".
	// Encrypted at rest; never returned.
	Headers map[string]string `json:"headers,omitzero" api:"required"`
	// Export invocation records (one per HTTP request) to this destination
	InvocationExportEnabled bool `json:"invocation_export_enabled" api:"required"`
	// Export runtime logs (function stdout/stderr) to this destination
	RuntimeExportEnabled bool `json:"runtime_export_enabled" api:"required"`
	paramObj
}

func (r ComputeFuncExportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeFuncExportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeFuncExportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
