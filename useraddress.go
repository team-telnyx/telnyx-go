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

// UserAddressService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserAddressService] method instead.
type UserAddressService struct {
	Options []option.RequestOption
}

// NewUserAddressService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAddressService(opts ...option.RequestOption) (r UserAddressService) {
	r = UserAddressService{}
	r.Options = opts
	return
}

// Creates a user address.
func (r *UserAddressService) New(ctx context.Context, body UserAddressNewParams, opts ...option.RequestOption) (res *UserAddressNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user_addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing user address.
func (r *UserAddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *UserAddressGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("user_addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your user addresses.
func (r *UserAddressService) List(ctx context.Context, query UserAddressListParams, opts ...option.RequestOption) (res *UserAddressListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user_addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserAddress struct {
	// Uniquely identifies the user address.
	ID string `json:"id" format:"uuid"`
	// The locality of the user address. For US addresses, this corresponds to the
	// state of the address.
	AdministrativeArea string `json:"administrative_area"`
	// The borough of the user address. This field is not used for addresses in the US
	// but is used for some international addresses.
	Borough string `json:"borough"`
	// The business name associated with the user address.
	BusinessName string `json:"business_name"`
	// The two-character (ISO 3166-1 alpha-2) country code of the user address.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Additional street address information about the user address such as, but not
	// limited to, unit number or apartment number.
	ExtendedAddress string `json:"extended_address"`
	// The first name associated with the user address.
	FirstName string `json:"first_name"`
	// The last name associated with the user address.
	LastName string `json:"last_name"`
	// The locality of the user address. For US addresses, this corresponds to the city
	// of the address.
	Locality string `json:"locality"`
	// The neighborhood of the user address. This field is not used for addresses in
	// the US but is used for some international addresses.
	Neighborhood string `json:"neighborhood"`
	// The phone number associated with the user address.
	PhoneNumber string `json:"phone_number"`
	// The postal code of the user address.
	PostalCode string `json:"postal_code"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The primary street address information about the user address.
	StreetAddress string `json:"street_address"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AdministrativeArea respjson.Field
		Borough            respjson.Field
		BusinessName       respjson.Field
		CountryCode        respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		ExtendedAddress    respjson.Field
		FirstName          respjson.Field
		LastName           respjson.Field
		Locality           respjson.Field
		Neighborhood       respjson.Field
		PhoneNumber        respjson.Field
		PostalCode         respjson.Field
		RecordType         respjson.Field
		StreetAddress      respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAddress) RawJSON() string { return r.JSON.raw }
func (r *UserAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAddressNewResponse struct {
	Data UserAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAddressNewResponse) RawJSON() string { return r.JSON.raw }
func (r *UserAddressNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAddressGetResponse struct {
	Data UserAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *UserAddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAddressListResponse struct {
	Data []UserAddress  `json:"data"`
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
func (r UserAddressListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserAddressListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAddressNewParams struct {
	// The business name associated with the user address.
	BusinessName string `json:"business_name,required"`
	// The two-character (ISO 3166-1 alpha-2) country code of the user address.
	CountryCode string `json:"country_code,required"`
	// The first name associated with the user address.
	FirstName string `json:"first_name,required"`
	// The last name associated with the user address.
	LastName string `json:"last_name,required"`
	// The locality of the user address. For US addresses, this corresponds to the city
	// of the address.
	Locality string `json:"locality,required"`
	// The primary street address information about the user address.
	StreetAddress string `json:"street_address,required"`
	// The locality of the user address. For US addresses, this corresponds to the
	// state of the address.
	AdministrativeArea param.Opt[string] `json:"administrative_area,omitzero"`
	// The borough of the user address. This field is not used for addresses in the US
	// but is used for some international addresses.
	Borough param.Opt[string] `json:"borough,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Additional street address information about the user address such as, but not
	// limited to, unit number or apartment number.
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	// The neighborhood of the user address. This field is not used for addresses in
	// the US but is used for some international addresses.
	Neighborhood param.Opt[string] `json:"neighborhood,omitzero"`
	// The phone number associated with the user address.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The postal code of the user address.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// An optional boolean value specifying if verification of the address should be
	// skipped or not. UserAddresses are generally used for shipping addresses, and
	// failure to validate your shipping address will likely result in a failure to
	// deliver SIM cards or other items ordered from Telnyx. Do not use this parameter
	// unless you are sure that the address is correct even though it cannot be
	// validated. If this is set to any value other than true, verification of the
	// address will be attempted, and the user address will not be allowed if
	// verification fails. If verification fails but suggested values are available
	// that might make the address correct, they will be present in the response as
	// well. If this value is set to true, then the verification will not be attempted.
	// Defaults to false (verification will be performed).
	SkipAddressVerification param.Opt[string] `json:"skip_address_verification,omitzero"`
	paramObj
}

func (r UserAddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserAddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserAddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserAddressListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[customer_reference][eq], filter[customer_reference][contains],
	// filter[street_address][contains]
	Filter UserAddressListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page UserAddressListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code> -</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>street_address</code>: sorts the result by the
	//	  <code>street_address</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-street_address</code>: sorts the result by the
	//	  <code>street_address</code> field in descending order.
	//	</li>
	//
	// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "created_at", "first_name", "last_name", "business_name",
	// "street_address".
	Sort UserAddressListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserAddressListParams]'s query parameters as `url.Values`.
func (r UserAddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[customer_reference][eq], filter[customer_reference][contains],
// filter[street_address][contains]
type UserAddressListParamsFilter struct {
	// Filter user addresses via the customer reference. Supports both exact matching
	// (eq) and partial matching (contains). Matching is not case-sensitive.
	CustomerReference UserAddressListParamsFilterCustomerReference `query:"customer_reference,omitzero" json:"-"`
	// Filter user addresses via street address. Supports partial matching (contains).
	// Matching is not case-sensitive.
	StreetAddress UserAddressListParamsFilterStreetAddress `query:"street_address,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserAddressListParamsFilter]'s query parameters as
// `url.Values`.
func (r UserAddressListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter user addresses via the customer reference. Supports both exact matching
// (eq) and partial matching (contains). Matching is not case-sensitive.
type UserAddressListParamsFilterCustomerReference struct {
	// If present, user addresses with <code>customer_reference</code> containing the
	// given value will be returned. Matching is not case-sensitive.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter user addresses via exact customer reference match. Matching is not
	// case-sensitive.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserAddressListParamsFilterCustomerReference]'s query
// parameters as `url.Values`.
func (r UserAddressListParamsFilterCustomerReference) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter user addresses via street address. Supports partial matching (contains).
// Matching is not case-sensitive.
type UserAddressListParamsFilterStreetAddress struct {
	// If present, user addresses with <code>street_address</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserAddressListParamsFilterStreetAddress]'s query
// parameters as `url.Values`.
func (r UserAddressListParamsFilterStreetAddress) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type UserAddressListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserAddressListParamsPage]'s query parameters as
// `url.Values`.
func (r UserAddressListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. By default sorting direction is ascending.
// To have the results sorted in descending order add the <code> -</code>
// prefix.<br/><br/> That is: <ul>
//
//	<li>
//	  <code>street_address</code>: sorts the result by the
//	  <code>street_address</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-street_address</code>: sorts the result by the
//	  <code>street_address</code> field in descending order.
//	</li>
//
// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
type UserAddressListParamsSort string

const (
	UserAddressListParamsSortCreatedAt     UserAddressListParamsSort = "created_at"
	UserAddressListParamsSortFirstName     UserAddressListParamsSort = "first_name"
	UserAddressListParamsSortLastName      UserAddressListParamsSort = "last_name"
	UserAddressListParamsSortBusinessName  UserAddressListParamsSort = "business_name"
	UserAddressListParamsSortStreetAddress UserAddressListParamsSort = "street_address"
)
