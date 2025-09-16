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

// ExternalConnectionPhoneNumberService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionPhoneNumberService] method instead.
type ExternalConnectionPhoneNumberService struct {
	Options []option.RequestOption
}

// NewExternalConnectionPhoneNumberService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalConnectionPhoneNumberService(opts ...option.RequestOption) (r ExternalConnectionPhoneNumberService) {
	r = ExternalConnectionPhoneNumberService{}
	r.Options = opts
	return
}

// Return the details of a phone number associated with the given external
// connection.
func (r *ExternalConnectionPhoneNumberService) Get(ctx context.Context, phoneNumberID string, query ExternalConnectionPhoneNumberGetParams, opts ...option.RequestOption) (res *ExternalConnectionPhoneNumberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/phone_numbers/%s", query.ID, phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Asynchronously update settings of the phone number associated with the given
// external connection.
func (r *ExternalConnectionPhoneNumberService) Update(ctx context.Context, phoneNumberID string, params ExternalConnectionPhoneNumberUpdateParams, opts ...option.RequestOption) (res *ExternalConnectionPhoneNumberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/phone_numbers/%s", params.ID, phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of all active phone numbers associated with the given external
// connection.
func (r *ExternalConnectionPhoneNumberService) List(ctx context.Context, id string, query ExternalConnectionPhoneNumberListParams, opts ...option.RequestOption) (res *ExternalConnectionPhoneNumberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/phone_numbers", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ExternalConnectionPhoneNumber struct {
	// Any of "FirstPartyAppAssignment", "InboundCalling", "Office365",
	// "OutboundCalling", "UserAssignment".
	AcquiredCapabilities []string `json:"acquired_capabilities"`
	// Identifies the civic address assigned to the phone number.
	CivicAddressID string `json:"civic_address_id" format:"uuid"`
	// The iso country code that will be displayed to the user when they receive a call
	// from this phone number.
	DisplayedCountryCode string `json:"displayed_country_code"`
	// Identifies the location assigned to the phone number.
	LocationID string `json:"location_id" format:"uuid"`
	// Phone number ID from the Telnyx API.
	NumberID string `json:"number_id"`
	// Phone number in E164 format.
	TelephoneNumber string `json:"telephone_number"`
	// Uniquely identifies the resource.
	TicketID string `json:"ticket_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcquiredCapabilities respjson.Field
		CivicAddressID       respjson.Field
		DisplayedCountryCode respjson.Field
		LocationID           respjson.Field
		NumberID             respjson.Field
		TelephoneNumber      respjson.Field
		TicketID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionPhoneNumberGetResponse struct {
	Data ExternalConnectionPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionPhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionPhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionPhoneNumberUpdateResponse struct {
	Data ExternalConnectionPhoneNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionPhoneNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionPhoneNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionPhoneNumberListResponse struct {
	Data []ExternalConnectionPhoneNumber         `json:"data"`
	Meta ExternalVoiceIntegrationsPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionPhoneNumberGetParams struct {
	ID string `path:"id,required" format:"int64" json:"-"`
	paramObj
}

type ExternalConnectionPhoneNumberUpdateParams struct {
	ID string `path:"id,required" format:"int64" json:"-"`
	// Identifies the location to assign the phone number to.
	LocationID param.Opt[string] `json:"location_id,omitzero" format:"uuid"`
	paramObj
}

func (r ExternalConnectionPhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalConnectionPhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalConnectionPhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionPhoneNumberListParams struct {
	// Filter parameter for phone numbers (deepObject style). Supports filtering by
	// phone_number, civic_address_id, and location_id with eq/contains operations.
	Filter ExternalConnectionPhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page ExternalConnectionPhoneNumberListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionPhoneNumberListParams]'s query parameters
// as `url.Values`.
func (r ExternalConnectionPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for phone numbers (deepObject style). Supports filtering by
// phone_number, civic_address_id, and location_id with eq/contains operations.
type ExternalConnectionPhoneNumberListParamsFilter struct {
	CivicAddressID ExternalConnectionPhoneNumberListParamsFilterCivicAddressID `query:"civic_address_id,omitzero" json:"-"`
	LocationID     ExternalConnectionPhoneNumberListParamsFilterLocationID     `query:"location_id,omitzero" json:"-"`
	PhoneNumber    ExternalConnectionPhoneNumberListParamsFilterPhoneNumber    `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionPhoneNumberListParamsFilter]'s query
// parameters as `url.Values`.
func (r ExternalConnectionPhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionPhoneNumberListParamsFilterCivicAddressID struct {
	// The civic address ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [ExternalConnectionPhoneNumberListParamsFilterCivicAddressID]'s query parameters
// as `url.Values`.
func (r ExternalConnectionPhoneNumberListParamsFilterCivicAddressID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionPhoneNumberListParamsFilterLocationID struct {
	// The location ID to filter by
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionPhoneNumberListParamsFilterLocationID]'s
// query parameters as `url.Values`.
func (r ExternalConnectionPhoneNumberListParamsFilterLocationID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalConnectionPhoneNumberListParamsFilterPhoneNumber struct {
	// The phone number to filter by (partial match)
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// The phone number to filter by (exact match)
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionPhoneNumberListParamsFilterPhoneNumber]'s
// query parameters as `url.Values`.
func (r ExternalConnectionPhoneNumberListParamsFilterPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type ExternalConnectionPhoneNumberListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionPhoneNumberListParamsPage]'s query
// parameters as `url.Values`.
func (r ExternalConnectionPhoneNumberListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
