// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// SubNumberOrdersReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubNumberOrdersReportService] method instead.
type SubNumberOrdersReportService struct {
	Options []option.RequestOption
}

// NewSubNumberOrdersReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubNumberOrdersReportService(opts ...option.RequestOption) (r SubNumberOrdersReportService) {
	r = SubNumberOrdersReportService{}
	r.Options = opts
	return
}

// Create a CSV report for sub number orders. The report will be generated
// asynchronously and can be downloaded once complete.
func (r *SubNumberOrdersReportService) New(ctx context.Context, body SubNumberOrdersReportNewParams, opts ...option.RequestOption) (res *SubNumberOrdersReportNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sub_number_orders_report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the status and details of a sub number orders report.
func (r *SubNumberOrdersReportService) Get(ctx context.Context, reportID string, opts ...option.RequestOption) (res *SubNumberOrdersReportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders_report/%s", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Download the CSV file for a completed sub number orders report. The report
// status must be 'success' before the file can be downloaded.
func (r *SubNumberOrdersReportService) Download(ctx context.Context, reportID string, opts ...option.RequestOption) (res *io.Reader, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/csv")}, opts...)
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return
	}
	path := fmt.Sprintf("sub_number_orders_report/%s/download", reportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SubNumberOrdersReportNewResponse struct {
	Data SubNumberOrdersReportNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrdersReportNewResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The filters that were applied to generate this report
	Filters SubNumberOrdersReportNewResponseDataFilters `json:"filters"`
	// The type of order report.
	OrderType string `json:"order_type"`
	// Indicates the completion level of the sub number orders report. The report must
	// have a status of 'success' before it can be downloaded.
	//
	// Any of "pending", "success", "failed", "expired".
	Status string `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The ID of the user who created the report.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filters     respjson.Field
		OrderType   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The filters that were applied to generate this report
type SubNumberOrdersReportNewResponseDataFilters struct {
	CountryCode       string    `json:"country_code"`
	CreatedAtGt       time.Time `json:"created_at_gt" format:"date-time"`
	CreatedAtLt       time.Time `json:"created_at_lt" format:"date-time"`
	CustomerReference string    `json:"customer_reference"`
	OrderRequestID    string    `json:"order_request_id" format:"uuid"`
	Status            string    `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode       respjson.Field
		CreatedAtGt       respjson.Field
		CreatedAtLt       respjson.Field
		CustomerReference respjson.Field
		OrderRequestID    respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportNewResponseDataFilters) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportNewResponseDataFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrdersReportGetResponse struct {
	Data SubNumberOrdersReportGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrdersReportGetResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The filters that were applied to generate this report
	Filters SubNumberOrdersReportGetResponseDataFilters `json:"filters"`
	// The type of order report.
	OrderType string `json:"order_type"`
	// Indicates the completion level of the sub number orders report. The report must
	// have a status of 'success' before it can be downloaded.
	//
	// Any of "pending", "success", "failed", "expired".
	Status string `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The ID of the user who created the report.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Filters     respjson.Field
		OrderType   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The filters that were applied to generate this report
type SubNumberOrdersReportGetResponseDataFilters struct {
	CountryCode       string    `json:"country_code"`
	CreatedAtGt       time.Time `json:"created_at_gt" format:"date-time"`
	CreatedAtLt       time.Time `json:"created_at_lt" format:"date-time"`
	CustomerReference string    `json:"customer_reference"`
	OrderRequestID    string    `json:"order_request_id" format:"uuid"`
	Status            string    `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode       respjson.Field
		CreatedAtGt       respjson.Field
		CreatedAtLt       respjson.Field
		CustomerReference respjson.Field
		OrderRequestID    respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrdersReportGetResponseDataFilters) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrdersReportGetResponseDataFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrdersReportNewParams struct {
	// Filter by country code
	CountryCode param.Opt[string] `json:"country_code,omitzero"`
	// Filter for orders created after this date
	CreatedAtGt param.Opt[time.Time] `json:"created_at_gt,omitzero" format:"date-time"`
	// Filter for orders created before this date
	CreatedAtLt param.Opt[time.Time] `json:"created_at_lt,omitzero" format:"date-time"`
	// Filter by customer reference
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Filter by specific order request ID
	OrderRequestID param.Opt[string] `json:"order_request_id,omitzero" format:"uuid"`
	// Filter by order status
	//
	// Any of "pending", "success", "failure".
	Status SubNumberOrdersReportNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r SubNumberOrdersReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubNumberOrdersReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubNumberOrdersReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by order status
type SubNumberOrdersReportNewParamsStatus string

const (
	SubNumberOrdersReportNewParamsStatusPending SubNumberOrdersReportNewParamsStatus = "pending"
	SubNumberOrdersReportNewParamsStatusSuccess SubNumberOrdersReportNewParamsStatus = "success"
	SubNumberOrdersReportNewParamsStatusFailure SubNumberOrdersReportNewParamsStatus = "failure"
)
