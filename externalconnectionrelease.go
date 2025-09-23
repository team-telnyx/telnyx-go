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

// ExternalConnectionReleaseService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionReleaseService] method instead.
type ExternalConnectionReleaseService struct {
	Options []option.RequestOption
}

// NewExternalConnectionReleaseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalConnectionReleaseService(opts ...option.RequestOption) (r ExternalConnectionReleaseService) {
	r = ExternalConnectionReleaseService{}
	r.Options = opts
	return
}

// Return the details of a Release request and its phone numbers.
func (r *ExternalConnectionReleaseService) Get(ctx context.Context, releaseID string, query ExternalConnectionReleaseGetParams, opts ...option.RequestOption) (res *ExternalConnectionReleaseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if releaseID == "" {
		err = errors.New("missing required release_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/releases/%s", query.ID, releaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your Releases for the given external connection. These are
// automatically created when you change the `connection_id` of a phone number that
// is currently on Microsoft Teams.
func (r *ExternalConnectionReleaseService) List(ctx context.Context, id string, query ExternalConnectionReleaseListParams, opts ...option.RequestOption) (res *ExternalConnectionReleaseListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/releases", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExternalConnectionReleaseGetResponse struct {
	Data ExternalConnectionReleaseGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseGetResponseData struct {
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A message set if there is an error with the upload process.
	ErrorMessage string `json:"error_message"`
	// Represents the status of the release on Microsoft Teams.
	//
	// Any of "pending_upload", "pending", "in_progress", "complete", "failed",
	// "expired", "unknown".
	Status           string                                                    `json:"status"`
	TelephoneNumbers []ExternalConnectionReleaseGetResponseDataTelephoneNumber `json:"telephone_numbers"`
	TenantID         string                                                    `json:"tenant_id" format:"uuid"`
	// Uniquely identifies the resource.
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		ErrorMessage     respjson.Field
		Status           respjson.Field
		TelephoneNumbers respjson.Field
		TenantID         respjson.Field
		TicketID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseGetResponseDataTelephoneNumber struct {
	// Phone number ID from the Telnyx API.
	NumberID string `json:"number_id"`
	// Phone number in E164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberID    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseGetResponseDataTelephoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseGetResponseDataTelephoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseListResponse struct {
	Data []ExternalConnectionReleaseListResponseData `json:"data"`
	Meta ExternalVoiceIntegrationsPaginationMeta     `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseListResponseData struct {
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A message set if there is an error with the upload process.
	ErrorMessage string `json:"error_message"`
	// Represents the status of the release on Microsoft Teams.
	//
	// Any of "pending_upload", "pending", "in_progress", "complete", "failed",
	// "expired", "unknown".
	Status           string                                                     `json:"status"`
	TelephoneNumbers []ExternalConnectionReleaseListResponseDataTelephoneNumber `json:"telephone_numbers"`
	TenantID         string                                                     `json:"tenant_id" format:"uuid"`
	// Uniquely identifies the resource.
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		ErrorMessage     respjson.Field
		Status           respjson.Field
		TelephoneNumbers respjson.Field
		TenantID         respjson.Field
		TicketID         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseListResponseDataTelephoneNumber struct {
	// Phone number ID from the Telnyx API.
	NumberID string `json:"number_id"`
	// Phone number in E164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberID    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionReleaseListResponseDataTelephoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionReleaseListResponseDataTelephoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionReleaseGetParams struct {
	ID string `path:"id,required" format:"int64" json:"-"`
	paramObj
}

type ExternalConnectionReleaseListParams struct {
	// Filter parameter for releases (deepObject style). Supports filtering by status,
	// civic_address_id, location_id, and phone_number with eq/contains operations.
	Filter ExternalConnectionReleaseListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page ExternalConnectionReleaseListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParams]'s query parameters as
// `url.Values`.
func (r ExternalConnectionReleaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for releases (deepObject style). Supports filtering by status,
// civic_address_id, location_id, and phone_number with eq/contains operations.
type ExternalConnectionReleaseListParamsFilter struct {
	CivicAddressID ExternalConnectionReleaseListParamsFilterCivicAddressID `query:"civic_address_id,omitzero" json:"-"`
	LocationID     ExternalConnectionReleaseListParamsFilterLocationID     `query:"location_id,omitzero" json:"-"`
	// Phone number filter operations. Use 'eq' for exact matches or 'contains' for
	// partial matches.
	PhoneNumber ExternalConnectionReleaseListParamsFilterPhoneNumber `query:"phone_number,omitzero" json:"-"`
	Status      ExternalConnectionReleaseListParamsFilterStatus      `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsFilter]'s query
// parameters as `url.Values`.
func (r ExternalConnectionReleaseListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionReleaseListParamsFilterCivicAddressID struct {
	// The civic address ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsFilterCivicAddressID]'s
// query parameters as `url.Values`.
func (r ExternalConnectionReleaseListParamsFilterCivicAddressID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionReleaseListParamsFilterLocationID struct {
	// The location ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsFilterLocationID]'s
// query parameters as `url.Values`.
func (r ExternalConnectionReleaseListParamsFilterLocationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Phone number filter operations. Use 'eq' for exact matches or 'contains' for
// partial matches.
type ExternalConnectionReleaseListParamsFilterPhoneNumber struct {
	// The partial phone number to filter by. Requires 3-15 digits.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// The phone number to filter by
	Eq param.Opt[string] `query:"eq,omitzero" format:"E164" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsFilterPhoneNumber]'s
// query parameters as `url.Values`.
func (r ExternalConnectionReleaseListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionReleaseListParamsFilterStatus struct {
	// The status of the release to filter by
	//
	// Any of "pending_upload", "pending", "in_progress", "complete", "failed",
	// "expired", "unknown".
	Eq []string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r ExternalConnectionReleaseListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type ExternalConnectionReleaseListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionReleaseListParamsPage]'s query parameters
// as `url.Values`.
func (r ExternalConnectionReleaseListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
