// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// ReportService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportService] method instead.
type ReportService struct {
	Options         []option.RequestOption
	CdrUsageReports ReportCdrUsageReportService
	MdrUsageReports ReportMdrUsageReportService
}

// NewReportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewReportService(opts ...option.RequestOption) (r ReportService) {
	r = ReportService{}
	r.Options = opts
	r.CdrUsageReports = NewReportCdrUsageReportService(opts...)
	r.MdrUsageReports = NewReportMdrUsageReportService(opts...)
	return
}

// Fetch all Mdr records
func (r *ReportService) ListMdrs(ctx context.Context, query ReportListMdrsParams, opts ...option.RequestOption) (res *ReportListMdrsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reports/mdrs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetch all Wdr records
func (r *ReportService) ListWdrs(ctx context.Context, query ReportListWdrsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ReportListWdrsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "reports/wdrs"
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

// Fetch all Wdr records
func (r *ReportService) ListWdrsAutoPaging(ctx context.Context, query ReportListWdrsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ReportListWdrsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListWdrs(ctx, query, opts...))
}

type ReportListMdrsResponse struct {
	Data []ReportListMdrsResponseData `json:"data"`
	Meta PaginationMetaReporting      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListMdrsResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportListMdrsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListMdrsResponseData struct {
	// Id of message detail record
	ID string `json:"id"`
	// The destination number for a call, or the callee
	Cld string `json:"cld"`
	// The number associated with the person initiating the call, or the caller
	Cli string `json:"cli"`
	// Final cost. Cost is calculated as rate \* parts
	Cost string `json:"cost"`
	// Message sent time
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Currency of the rate and cost
	//
	// Any of "AUD", "CAD", "EUR", "GBP", "USD".
	Currency string `json:"currency"`
	// Direction of message - inbound or outbound.
	Direction string `json:"direction"`
	// Type of message
	//
	// Any of "SMS", "MMS".
	MessageType string `json:"message_type"`
	// Number of parts this message has. Max number of character is 160. If message
	// contains more characters then that it will be broken down in multiple parts
	Parts float64 `json:"parts"`
	// Configured profile name. New profiles can be created and configured on Telnyx
	// portal
	ProfileName string `json:"profile_name"`
	// Rate applied to the message
	Rate       string `json:"rate"`
	RecordType string `json:"record_type"`
	// Message status
	//
	// Any of "GW_TIMEOUT", "DELIVERED", "DLR_UNCONFIRMED", "DLR_TIMEOUT", "RECEIVED",
	// "GW_REJECT", "FAILED".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cld         respjson.Field
		Cli         respjson.Field
		Cost        respjson.Field
		CreatedAt   respjson.Field
		Currency    respjson.Field
		Direction   respjson.Field
		MessageType respjson.Field
		Parts       respjson.Field
		ProfileName respjson.Field
		Rate        respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListMdrsResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReportListMdrsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListWdrsResponse struct {
	// WDR id
	ID   string                     `json:"id"`
	Cost ReportListWdrsResponseCost `json:"cost"`
	// Record created time
	CreatedAt    time.Time                          `json:"created_at" format:"date-time"`
	DownlinkData ReportListWdrsResponseDownlinkData `json:"downlink_data"`
	// Session duration in seconds.
	DurationSeconds float64 `json:"duration_seconds"`
	// International mobile subscriber identity.
	Imsi string `json:"imsi"`
	// Mobile country code.
	Mcc string `json:"mcc"`
	// Mobile network code.
	Mnc string `json:"mnc"`
	// Phone number
	PhoneNumber string                     `json:"phone_number"`
	Rate        ReportListWdrsResponseRate `json:"rate"`
	RecordType  string                     `json:"record_type"`
	// Sim card unique identifier
	SimCardID string `json:"sim_card_id"`
	// Sim group unique identifier
	SimGroupID string `json:"sim_group_id"`
	// Defined sim group name
	SimGroupName string                           `json:"sim_group_name"`
	UplinkData   ReportListWdrsResponseUplinkData `json:"uplink_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Cost            respjson.Field
		CreatedAt       respjson.Field
		DownlinkData    respjson.Field
		DurationSeconds respjson.Field
		Imsi            respjson.Field
		Mcc             respjson.Field
		Mnc             respjson.Field
		PhoneNumber     respjson.Field
		Rate            respjson.Field
		RecordType      respjson.Field
		SimCardID       respjson.Field
		SimGroupID      respjson.Field
		SimGroupName    respjson.Field
		UplinkData      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListWdrsResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportListWdrsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListWdrsResponseCost struct {
	// Final cost. Cost is calculated as rate \* unit
	Amount string `json:"amount"`
	// Currency of the rate and cost
	//
	// Any of "AUD", "CAD", "EUR", "GBP", "USD".
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListWdrsResponseCost) RawJSON() string { return r.JSON.raw }
func (r *ReportListWdrsResponseCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListWdrsResponseDownlinkData struct {
	// Downlink data
	Amount float64 `json:"amount"`
	// Transmission unit
	//
	// Any of "B", "KB", "MB".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListWdrsResponseDownlinkData) RawJSON() string { return r.JSON.raw }
func (r *ReportListWdrsResponseDownlinkData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListWdrsResponseRate struct {
	// Rate from which cost is calculated
	Amount string `json:"amount"`
	// Currency of the rate and cost
	//
	// Any of "AUD", "CAD", "EUR", "GBP", "USD".
	Currency string `json:"currency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListWdrsResponseRate) RawJSON() string { return r.JSON.raw }
func (r *ReportListWdrsResponseRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListWdrsResponseUplinkData struct {
	// Uplink data
	Amount float64 `json:"amount"`
	// Transmission unit
	//
	// Any of "B", "KB", "MB".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportListWdrsResponseUplinkData) RawJSON() string { return r.JSON.raw }
func (r *ReportListWdrsResponseUplinkData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportListMdrsParams struct {
	// Message uuid
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Destination number
	Cld param.Opt[string] `query:"cld,omitzero" json:"-"`
	// Origination number
	Cli param.Opt[string] `query:"cli,omitzero" json:"-"`
	// Pagination end date
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Name of the profile
	Profile param.Opt[string] `query:"profile,omitzero" json:"-"`
	// Pagination start date
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Direction (inbound or outbound)
	//
	// Any of "INBOUND", "OUTBOUND".
	Direction ReportListMdrsParamsDirection `query:"direction,omitzero" json:"-"`
	// Type of message
	//
	// Any of "SMS", "MMS".
	MessageType ReportListMdrsParamsMessageType `query:"message_type,omitzero" json:"-"`
	// Message status
	//
	// Any of "GW_TIMEOUT", "DELIVERED", "DLR_UNCONFIRMED", "DLR_TIMEOUT", "RECEIVED",
	// "GW_REJECT", "FAILED".
	Status ReportListMdrsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportListMdrsParams]'s query parameters as `url.Values`.
func (r ReportListMdrsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction (inbound or outbound)
type ReportListMdrsParamsDirection string

const (
	ReportListMdrsParamsDirectionInbound  ReportListMdrsParamsDirection = "INBOUND"
	ReportListMdrsParamsDirectionOutbound ReportListMdrsParamsDirection = "OUTBOUND"
)

// Type of message
type ReportListMdrsParamsMessageType string

const (
	ReportListMdrsParamsMessageTypeSMS ReportListMdrsParamsMessageType = "SMS"
	ReportListMdrsParamsMessageTypeMms ReportListMdrsParamsMessageType = "MMS"
)

// Message status
type ReportListMdrsParamsStatus string

const (
	ReportListMdrsParamsStatusGwTimeout      ReportListMdrsParamsStatus = "GW_TIMEOUT"
	ReportListMdrsParamsStatusDelivered      ReportListMdrsParamsStatus = "DELIVERED"
	ReportListMdrsParamsStatusDlrUnconfirmed ReportListMdrsParamsStatus = "DLR_UNCONFIRMED"
	ReportListMdrsParamsStatusDlrTimeout     ReportListMdrsParamsStatus = "DLR_TIMEOUT"
	ReportListMdrsParamsStatusReceived       ReportListMdrsParamsStatus = "RECEIVED"
	ReportListMdrsParamsStatusGwReject       ReportListMdrsParamsStatus = "GW_REJECT"
	ReportListMdrsParamsStatusFailed         ReportListMdrsParamsStatus = "FAILED"
)

type ReportListWdrsParams struct {
	// WDR uuid
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// End date
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// International mobile subscriber identity
	Imsi param.Opt[string] `query:"imsi,omitzero" json:"-"`
	// Mobile country code
	Mcc param.Opt[string] `query:"mcc,omitzero" json:"-"`
	// Mobile network code
	Mnc        param.Opt[string] `query:"mnc,omitzero" json:"-"`
	PageNumber param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	// Phone number
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Sim card unique identifier
	SimCardID param.Opt[string] `query:"sim_card_id,omitzero" json:"-"`
	// Sim group unique identifier
	SimGroupID param.Opt[string] `query:"sim_group_id,omitzero" json:"-"`
	// Sim group name
	SimGroupName param.Opt[string] `query:"sim_group_name,omitzero" json:"-"`
	// Start date
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Field used to order the data. If no field is specified, default value is
	// 'created_at'
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportListWdrsParams]'s query parameters as `url.Values`.
func (r ReportListWdrsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
