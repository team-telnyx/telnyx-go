// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// LedgerBillingGroupReportService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerBillingGroupReportService] method instead.
type LedgerBillingGroupReportService struct {
	Options []option.RequestOption
}

// NewLedgerBillingGroupReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLedgerBillingGroupReportService(opts ...option.RequestOption) (r LedgerBillingGroupReportService) {
	r = LedgerBillingGroupReportService{}
	r.Options = opts
	return
}

// Create a ledger billing group report
func (r *LedgerBillingGroupReportService) New(ctx context.Context, body LedgerBillingGroupReportNewParams, opts ...option.RequestOption) (res *LedgerBillingGroupReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ledger_billing_group_reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a ledger billing group report
func (r *LedgerBillingGroupReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerBillingGroupReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ledger_billing_group_reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerBillingGroupReport struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Uniquely identifies the organization that owns the resource.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Identifies the type of the resource.
	//
	// Any of "ledger_billing_group_report".
	RecordType LedgerBillingGroupReportRecordType `json:"record_type"`
	// External url of the ledger billing group report, if the status is complete
	ReportURL string `json:"report_url" api:"nullable" format:"uri"`
	// Status of the ledger billing group report
	//
	// Any of "pending", "complete", "failed", "deleted".
	Status LedgerBillingGroupReportStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		OrganizationID respjson.Field
		RecordType     respjson.Field
		ReportURL      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LedgerBillingGroupReport) RawJSON() string { return r.JSON.raw }
func (r *LedgerBillingGroupReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type LedgerBillingGroupReportRecordType string

const (
	LedgerBillingGroupReportRecordTypeLedgerBillingGroupReport LedgerBillingGroupReportRecordType = "ledger_billing_group_report"
)

// Status of the ledger billing group report
type LedgerBillingGroupReportStatus string

const (
	LedgerBillingGroupReportStatusPending  LedgerBillingGroupReportStatus = "pending"
	LedgerBillingGroupReportStatusComplete LedgerBillingGroupReportStatus = "complete"
	LedgerBillingGroupReportStatusFailed   LedgerBillingGroupReportStatus = "failed"
	LedgerBillingGroupReportStatusDeleted  LedgerBillingGroupReportStatus = "deleted"
)

type LedgerBillingGroupReportNewResponse struct {
	Data LedgerBillingGroupReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LedgerBillingGroupReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LedgerBillingGroupReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerBillingGroupReportGetResponse struct {
	Data LedgerBillingGroupReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LedgerBillingGroupReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LedgerBillingGroupReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerBillingGroupReportNewParams struct {
	// Month of the ledger billing group report
	Month param.Opt[int64] `json:"month,omitzero"`
	// Year of the ledger billing group report
	Year param.Opt[int64] `json:"year,omitzero"`
	paramObj
}

func (r LedgerBillingGroupReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LedgerBillingGroupReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LedgerBillingGroupReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
