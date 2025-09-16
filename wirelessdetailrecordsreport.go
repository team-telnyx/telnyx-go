// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// WirelessDetailRecordsReportService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWirelessDetailRecordsReportService] method instead.
type WirelessDetailRecordsReportService struct {
	Options []option.RequestOption
}

// NewWirelessDetailRecordsReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWirelessDetailRecordsReportService(opts ...option.RequestOption) (r WirelessDetailRecordsReportService) {
	r = WirelessDetailRecordsReportService{}
	r.Options = opts
	return
}

// Asynchronously create a report containing Wireless Detail Records (WDRs) for the
// SIM cards that consumed wireless data in the given time period.
func (r *WirelessDetailRecordsReportService) New(ctx context.Context, body WirelessDetailRecordsReportNewParams, opts ...option.RequestOption) (res *WirelessDetailRecordsReportNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "wireless/detail_records_reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns one specific WDR report
func (r *WirelessDetailRecordsReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WirelessDetailRecordsReportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireless/detail_records_reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the WDR Reports that match the given parameters.
func (r *WirelessDetailRecordsReportService) List(ctx context.Context, query WirelessDetailRecordsReportListParams, opts ...option.RequestOption) (res *WirelessDetailRecordsReportListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "wireless/detail_records_reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes one specific WDR report.
func (r *WirelessDetailRecordsReportService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WirelessDetailRecordsReportDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("wireless/detail_records_reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WdrReport struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date-time indicating the end time.
	EndTime    string `json:"end_time"`
	RecordType string `json:"record_type"`
	// The URL where the report content, when generated, will be published to.
	ReportURL string `json:"report_url"`
	// ISO 8601 formatted date-time indicating the start time.
	StartTime string `json:"start_time"`
	// Indicates the status of the report, which is updated asynchronously.
	//
	// Any of "pending", "complete", "failed", "deleted".
	Status WdrReportStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EndTime     respjson.Field
		RecordType  respjson.Field
		ReportURL   respjson.Field
		StartTime   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WdrReport) RawJSON() string { return r.JSON.raw }
func (r *WdrReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the status of the report, which is updated asynchronously.
type WdrReportStatus string

const (
	WdrReportStatusPending  WdrReportStatus = "pending"
	WdrReportStatusComplete WdrReportStatus = "complete"
	WdrReportStatusFailed   WdrReportStatus = "failed"
	WdrReportStatusDeleted  WdrReportStatus = "deleted"
)

type WirelessDetailRecordsReportNewResponse struct {
	Data WdrReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessDetailRecordsReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessDetailRecordsReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessDetailRecordsReportGetResponse struct {
	Data WdrReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessDetailRecordsReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessDetailRecordsReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessDetailRecordsReportListResponse struct {
	Data []WdrReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessDetailRecordsReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessDetailRecordsReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessDetailRecordsReportDeleteResponse struct {
	Data WdrReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessDetailRecordsReportDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessDetailRecordsReportDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessDetailRecordsReportNewParams struct {
	// ISO 8601 formatted date-time indicating the end time.
	EndTime param.Opt[string] `json:"end_time,omitzero"`
	// ISO 8601 formatted date-time indicating the start time.
	StartTime param.Opt[string] `json:"start_time,omitzero"`
	paramObj
}

func (r WirelessDetailRecordsReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WirelessDetailRecordsReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WirelessDetailRecordsReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessDetailRecordsReportListParams struct {
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WirelessDetailRecordsReportListParams]'s query parameters
// as `url.Values`.
func (r WirelessDetailRecordsReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
