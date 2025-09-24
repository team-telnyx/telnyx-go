// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PhoneNumberCsvDownloadService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberCsvDownloadService] method instead.
type PhoneNumberCsvDownloadService struct {
	Options []option.RequestOption
}

// NewPhoneNumberCsvDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberCsvDownloadService(opts ...option.RequestOption) (r PhoneNumberCsvDownloadService) {
	r = PhoneNumberCsvDownloadService{}
	r.Options = opts
	return
}

// Create a CSV download
func (r *PhoneNumberCsvDownloadService) New(ctx context.Context, body PhoneNumberCsvDownloadNewParams, opts ...option.RequestOption) (res *PhoneNumberCsvDownloadNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/csv_downloads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a CSV download
func (r *PhoneNumberCsvDownloadService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberCsvDownloadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/csv_downloads/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List CSV downloads
func (r *PhoneNumberCsvDownloadService) List(ctx context.Context, query PhoneNumberCsvDownloadListParams, opts ...option.RequestOption) (res *PhoneNumberCsvDownloadListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/csv_downloads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CsvDownload struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates the completion level of the CSV report. Only complete CSV download
	// requests will be able to be retrieved.
	//
	// Any of "pending", "complete", "failed", "expired".
	Status CsvDownloadStatus `json:"status"`
	// The URL at which the CSV file can be retrieved.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CsvDownload) RawJSON() string { return r.JSON.raw }
func (r *CsvDownload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the completion level of the CSV report. Only complete CSV download
// requests will be able to be retrieved.
type CsvDownloadStatus string

const (
	CsvDownloadStatusPending  CsvDownloadStatus = "pending"
	CsvDownloadStatusComplete CsvDownloadStatus = "complete"
	CsvDownloadStatusFailed   CsvDownloadStatus = "failed"
	CsvDownloadStatusExpired  CsvDownloadStatus = "expired"
)

type PhoneNumberCsvDownloadNewResponse struct {
	Data []CsvDownload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCsvDownloadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCsvDownloadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCsvDownloadGetResponse struct {
	Data []CsvDownload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCsvDownloadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCsvDownloadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCsvDownloadListResponse struct {
	Data []CsvDownload  `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCsvDownloadListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCsvDownloadListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberCsvDownloadNewParams struct {
	// Which format to use when generating the CSV file. The default for backwards
	// compatibility is 'V1'
	//
	// Any of "V1", "V2".
	CsvFormat PhoneNumberCsvDownloadNewParamsCsvFormat `query:"csv_format,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number],
	// filter[status], filter[voice.connection_name],
	// filter[voice.usage_payment_method], filter[billing_group_id],
	// filter[emergency_address_id], filter[customer_reference]
	Filter PhoneNumberCsvDownloadNewParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCsvDownloadNewParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberCsvDownloadNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which format to use when generating the CSV file. The default for backwards
// compatibility is 'V1'
type PhoneNumberCsvDownloadNewParamsCsvFormat string

const (
	PhoneNumberCsvDownloadNewParamsCsvFormatV1 PhoneNumberCsvDownloadNewParamsCsvFormat = "V1"
	PhoneNumberCsvDownloadNewParamsCsvFormatV2 PhoneNumberCsvDownloadNewParamsCsvFormat = "V2"
)

// Consolidated filter parameter (deepObject style). Originally:
// filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number],
// filter[status], filter[voice.connection_name],
// filter[voice.usage_payment_method], filter[billing_group_id],
// filter[emergency_address_id], filter[customer_reference]
type PhoneNumberCsvDownloadNewParamsFilter struct {
	// Filter by the billing_group_id associated with phone numbers. To filter to only
	// phone numbers that have no billing group associated them, set the value of this
	// filter to the string 'null'.
	BillingGroupID param.Opt[string] `query:"billing_group_id,omitzero" json:"-"`
	// Filter by connection_id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Filter numbers via the customer_reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter by the emergency_address_id associated with phone numbers. To filter only
	// phone numbers that have no emergency address associated with them, set the value
	// of this filter to the string 'null'.
	EmergencyAddressID param.Opt[string] `query:"emergency_address_id,omitzero" json:"-"`
	// Filter by phone number that have bundles.
	HasBundle param.Opt[string] `query:"has_bundle,omitzero" json:"-"`
	// Filter by phone number. Requires at least three digits. Non-numerical characters
	// will result in no values being returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by phone number tags.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Filter by phone number status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "active",
	// "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by voice connection name pattern matching.
	VoiceConnectionName PhoneNumberCsvDownloadNewParamsFilterVoiceConnectionName `query:"voice.connection_name,omitzero" json:"-"`
	// Filter by usage_payment_method.
	//
	// Any of "pay-per-minute", "channel".
	VoiceUsagePaymentMethod string `query:"voice.usage_payment_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCsvDownloadNewParamsFilter]'s query parameters
// as `url.Values`.
func (r PhoneNumberCsvDownloadNewParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by voice connection name pattern matching.
type PhoneNumberCsvDownloadNewParamsFilterVoiceConnectionName struct {
	// Filter contains connection name. Requires at least three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter ends with connection name. Requires at least three characters.
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter by connection name.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter starts with connection name. Requires at least three characters.
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCsvDownloadNewParamsFilterVoiceConnectionName]'s
// query parameters as `url.Values`.
func (r PhoneNumberCsvDownloadNewParamsFilterVoiceConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PhoneNumberCsvDownloadListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PhoneNumberCsvDownloadListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCsvDownloadListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberCsvDownloadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PhoneNumberCsvDownloadListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberCsvDownloadListParamsPage]'s query parameters as
// `url.Values`.
func (r PhoneNumberCsvDownloadListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
