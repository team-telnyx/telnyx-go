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
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ExternalConnectionUploadService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionUploadService] method instead.
type ExternalConnectionUploadService struct {
	Options []option.RequestOption
}

// NewExternalConnectionUploadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalConnectionUploadService(opts ...option.RequestOption) (r ExternalConnectionUploadService) {
	r = ExternalConnectionUploadService{}
	r.Options = opts
	return
}

// Creates a new Upload request to Microsoft teams with the included phone numbers.
// Only one of civic_address_id or location_id must be provided, not both. The
// maximum allowed phone numbers for the numbers_ids array is 1000.
func (r *ExternalConnectionUploadService) New(ctx context.Context, id string, body ExternalConnectionUploadNewParams, opts ...option.RequestOption) (res *ExternalConnectionUploadNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return the details of an Upload request and its phone numbers.
func (r *ExternalConnectionUploadService) Get(ctx context.Context, ticketID string, query ExternalConnectionUploadGetParams, opts ...option.RequestOption) (res *ExternalConnectionUploadGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if ticketID == "" {
		err = errors.New("missing required ticket_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads/%s", query.ID, ticketID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your Upload requests for the given external connection.
func (r *ExternalConnectionUploadService) List(ctx context.Context, id string, query ExternalConnectionUploadListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Upload], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads", id)
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

// Returns a list of your Upload requests for the given external connection.
func (r *ExternalConnectionUploadService) ListAutoPaging(ctx context.Context, id string, query ExternalConnectionUploadListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Upload] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Returns the count of all pending upload requests for the given external
// connection.
func (r *ExternalConnectionUploadService) PendingCount(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionUploadPendingCountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Forces a recheck of the status of all pending Upload requests for the given
// external connection in the background.
func (r *ExternalConnectionUploadService) RefreshStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalConnectionUploadRefreshStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads/refresh", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// If there were any errors during the upload process, this endpoint will retry the
// upload request. In some cases this will reattempt the existing upload request,
// in other cases it may create a new upload request. Please check the ticket_id in
// the response to determine if a new upload request was created.
func (r *ExternalConnectionUploadService) Retry(ctx context.Context, ticketID string, body ExternalConnectionUploadRetryParams, opts ...option.RequestOption) (res *ExternalConnectionUploadRetryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if ticketID == "" {
		err = errors.New("missing required ticket_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/uploads/%s/retry", body.ID, ticketID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TnUploadEntry struct {
	// Identifies the civic address assigned to the phone number entry.
	CivicAddressID string `json:"civic_address_id" format:"uuid"`
	// A code returned by Microsoft Teams if there is an error with the phone number
	// entry upload.
	//
	// Any of "internal_error", "unable_to_retrieve_default_location",
	// "unknown_country_code", "unable_to_retrieve_location",
	// "unable_to_retrieve_partner_info", "unable_to_match_geography_entry".
	ErrorCode TnUploadEntryErrorCode `json:"error_code"`
	// A message returned by Microsoft Teams if there is an error with the upload
	// process.
	ErrorMessage string `json:"error_message"`
	// Represents the status of the phone number entry upload on Telnyx.
	//
	// Any of "pending_assignment", "in_progress", "all_internal_jobs_completed",
	// "release_requested", "release_completed", "error".
	InternalStatus TnUploadEntryInternalStatus `json:"internal_status"`
	// Identifies the location assigned to the phone number entry.
	LocationID string `json:"location_id" format:"uuid"`
	// Uniquely identifies the resource.
	NumberID string `json:"number_id" format:"uuid"`
	// Phone number in E164 format.
	PhoneNumber string `json:"phone_number"`
	// Represents the status of the phone number entry upload on Microsoft Teams.
	//
	// Any of "pending_upload", "pending", "in_progress", "success", "error".
	Status TnUploadEntryStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CivicAddressID respjson.Field
		ErrorCode      respjson.Field
		ErrorMessage   respjson.Field
		InternalStatus respjson.Field
		LocationID     respjson.Field
		NumberID       respjson.Field
		PhoneNumber    respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TnUploadEntry) RawJSON() string { return r.JSON.raw }
func (r *TnUploadEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A code returned by Microsoft Teams if there is an error with the phone number
// entry upload.
type TnUploadEntryErrorCode string

const (
	TnUploadEntryErrorCodeInternalError                   TnUploadEntryErrorCode = "internal_error"
	TnUploadEntryErrorCodeUnableToRetrieveDefaultLocation TnUploadEntryErrorCode = "unable_to_retrieve_default_location"
	TnUploadEntryErrorCodeUnknownCountryCode              TnUploadEntryErrorCode = "unknown_country_code"
	TnUploadEntryErrorCodeUnableToRetrieveLocation        TnUploadEntryErrorCode = "unable_to_retrieve_location"
	TnUploadEntryErrorCodeUnableToRetrievePartnerInfo     TnUploadEntryErrorCode = "unable_to_retrieve_partner_info"
	TnUploadEntryErrorCodeUnableToMatchGeographyEntry     TnUploadEntryErrorCode = "unable_to_match_geography_entry"
)

// Represents the status of the phone number entry upload on Telnyx.
type TnUploadEntryInternalStatus string

const (
	TnUploadEntryInternalStatusPendingAssignment        TnUploadEntryInternalStatus = "pending_assignment"
	TnUploadEntryInternalStatusInProgress               TnUploadEntryInternalStatus = "in_progress"
	TnUploadEntryInternalStatusAllInternalJobsCompleted TnUploadEntryInternalStatus = "all_internal_jobs_completed"
	TnUploadEntryInternalStatusReleaseRequested         TnUploadEntryInternalStatus = "release_requested"
	TnUploadEntryInternalStatusReleaseCompleted         TnUploadEntryInternalStatus = "release_completed"
	TnUploadEntryInternalStatusError                    TnUploadEntryInternalStatus = "error"
)

// Represents the status of the phone number entry upload on Microsoft Teams.
type TnUploadEntryStatus string

const (
	TnUploadEntryStatusPendingUpload TnUploadEntryStatus = "pending_upload"
	TnUploadEntryStatusPending       TnUploadEntryStatus = "pending"
	TnUploadEntryStatusInProgress    TnUploadEntryStatus = "in_progress"
	TnUploadEntryStatusSuccess       TnUploadEntryStatus = "success"
	TnUploadEntryStatusError         TnUploadEntryStatus = "error"
)

type Upload struct {
	// Any of "calling_user_assignment", "first_party_app_assignment".
	AvailableUsages []string `json:"available_usages"`
	// A code returned by Microsoft Teams if there is an error with the upload process.
	ErrorCode string `json:"error_code"`
	// A message set if there is an error with the upload process.
	ErrorMessage string `json:"error_message"`
	LocationID   string `json:"location_id" format:"uuid"`
	// Represents the status of the upload on Microsoft Teams.
	//
	// Any of "pending_upload", "pending", "in_progress", "partial_success", "success",
	// "error".
	Status   UploadStatus `json:"status"`
	TenantID string       `json:"tenant_id" format:"uuid"`
	// Uniquely identifies the resource.
	TicketID        string          `json:"ticket_id" format:"uuid"`
	TnUploadEntries []TnUploadEntry `json:"tn_upload_entries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableUsages respjson.Field
		ErrorCode       respjson.Field
		ErrorMessage    respjson.Field
		LocationID      respjson.Field
		Status          respjson.Field
		TenantID        respjson.Field
		TicketID        respjson.Field
		TnUploadEntries respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Upload) RawJSON() string { return r.JSON.raw }
func (r *Upload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the status of the upload on Microsoft Teams.
type UploadStatus string

const (
	UploadStatusPendingUpload  UploadStatus = "pending_upload"
	UploadStatusPending        UploadStatus = "pending"
	UploadStatusInProgress     UploadStatus = "in_progress"
	UploadStatusPartialSuccess UploadStatus = "partial_success"
	UploadStatusSuccess        UploadStatus = "success"
	UploadStatusError          UploadStatus = "error"
)

type ExternalConnectionUploadNewResponse struct {
	// Describes wether or not the operation was successful
	Success bool `json:"success"`
	// Ticket id of the upload request
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		TicketID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadGetResponse struct {
	Data Upload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadPendingCountResponse struct {
	Data ExternalConnectionUploadPendingCountResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadPendingCountResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadPendingCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadPendingCountResponseData struct {
	// The count of phone numbers that are pending assignment to the external
	// connection.
	PendingNumbersCount int64 `json:"pending_numbers_count"`
	// The count of number uploads that have not yet been uploaded to Microsoft.
	PendingOrdersCount int64 `json:"pending_orders_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PendingNumbersCount respjson.Field
		PendingOrdersCount  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadPendingCountResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadPendingCountResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadRefreshStatusResponse struct {
	// Describes wether or not the operation was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadRefreshStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadRefreshStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadRetryResponse struct {
	Data Upload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionUploadRetryResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionUploadRetryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionUploadNewParams struct {
	NumberIDs []string `json:"number_ids,omitzero" api:"required"`
	// Identifies the civic address to assign all phone numbers to.
	CivicAddressID param.Opt[string] `json:"civic_address_id,omitzero" format:"uuid"`
	// Identifies the location to assign all phone numbers to.
	LocationID param.Opt[string] `json:"location_id,omitzero" format:"uuid"`
	// Any of "calling_user_assignment", "first_party_app_assignment".
	AdditionalUsages []string `json:"additional_usages,omitzero"`
	// The use case of the upload request. NOTE: `calling_user_assignment` is not
	// supported for toll free numbers.
	//
	// Any of "calling_user_assignment", "first_party_app_assignment".
	Usage ExternalConnectionUploadNewParamsUsage `json:"usage,omitzero"`
	paramObj
}

func (r ExternalConnectionUploadNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionUploadNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionUploadNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The use case of the upload request. NOTE: `calling_user_assignment` is not
// supported for toll free numbers.
type ExternalConnectionUploadNewParamsUsage string

const (
	ExternalConnectionUploadNewParamsUsageCallingUserAssignment   ExternalConnectionUploadNewParamsUsage = "calling_user_assignment"
	ExternalConnectionUploadNewParamsUsageFirstPartyAppAssignment ExternalConnectionUploadNewParamsUsage = "first_party_app_assignment"
)

type ExternalConnectionUploadGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type ExternalConnectionUploadListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter parameter for uploads (deepObject style). Supports filtering by status,
	// civic_address_id, location_id, and phone_number with eq/contains operations.
	Filter ExternalConnectionUploadListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParams]'s query parameters as
// `url.Values`.
func (r ExternalConnectionUploadListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for uploads (deepObject style). Supports filtering by status,
// civic_address_id, location_id, and phone_number with eq/contains operations.
type ExternalConnectionUploadListParamsFilter struct {
	CivicAddressID ExternalConnectionUploadListParamsFilterCivicAddressID `query:"civic_address_id,omitzero" json:"-"`
	LocationID     ExternalConnectionUploadListParamsFilterLocationID     `query:"location_id,omitzero" json:"-"`
	PhoneNumber    ExternalConnectionUploadListParamsFilterPhoneNumber    `query:"phone_number,omitzero" json:"-"`
	Status         ExternalConnectionUploadListParamsFilterStatus         `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParamsFilter]'s query
// parameters as `url.Values`.
func (r ExternalConnectionUploadListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUploadListParamsFilterCivicAddressID struct {
	// The civic address ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParamsFilterCivicAddressID]'s
// query parameters as `url.Values`.
func (r ExternalConnectionUploadListParamsFilterCivicAddressID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUploadListParamsFilterLocationID struct {
	// The location ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParamsFilterLocationID]'s query
// parameters as `url.Values`.
func (r ExternalConnectionUploadListParamsFilterLocationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUploadListParamsFilterPhoneNumber struct {
	// The phone number to filter by (partial match)
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// The phone number to filter by (exact match)
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParamsFilterPhoneNumber]'s
// query parameters as `url.Values`.
func (r ExternalConnectionUploadListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUploadListParamsFilterStatus struct {
	// The status of the upload to filter by
	//
	// Any of "pending_upload", "pending", "in_progress", "success", "error".
	Eq []string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionUploadListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r ExternalConnectionUploadListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionUploadRetryParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
