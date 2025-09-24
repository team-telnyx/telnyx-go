// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// AddressService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
	Actions AddressActionService
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r AddressService) {
	r = AddressService{}
	r.Options = opts
	r.Actions = NewAddressActionService(opts...)
	return
}

// Creates an address.
func (r *AddressService) New(ctx context.Context, body AddressNewParams, opts ...option.RequestOption) (res *AddressNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing address.
func (r *AddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AddressGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your addresses.
func (r *AddressService) List(ctx context.Context, query AddressListParams, opts ...option.RequestOption) (res *AddressListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing address.
func (r *AddressService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AddressDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Address struct {
	// Uniquely identifies the address.
	ID string `json:"id" format:"int64"`
	// Indicates whether or not the address should be considered part of your list of
	// addresses that appear for regular use.
	AddressBook bool `json:"address_book"`
	// The locality of the address. For US addresses, this corresponds to the state of
	// the address.
	AdministrativeArea string `json:"administrative_area"`
	// The borough of the address. This field is not used for addresses in the US but
	// is used for some international addresses.
	Borough string `json:"borough"`
	// The business name associated with the address. An address must have either a
	// first last name or a business name.
	BusinessName string `json:"business_name"`
	// The two-character (ISO 3166-1 alpha-2) country code of the address.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Additional street address information about the address such as, but not limited
	// to, unit number or apartment number.
	ExtendedAddress string `json:"extended_address"`
	// The first name associated with the address. An address must have either a first
	// last name or a business name.
	FirstName string `json:"first_name"`
	// The last name associated with the address. An address must have either a first
	// last name or a business name.
	LastName string `json:"last_name"`
	// The locality of the address. For US addresses, this corresponds to the city of
	// the address.
	Locality string `json:"locality"`
	// The neighborhood of the address. This field is not used for addresses in the US
	// but is used for some international addresses.
	Neighborhood string `json:"neighborhood"`
	// The phone number associated with the address.
	PhoneNumber string `json:"phone_number"`
	// The postal code of the address.
	PostalCode string `json:"postal_code"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The primary street address information about the address.
	StreetAddress string `json:"street_address"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Indicates whether or not the address should be validated for emergency use upon
	// creation or not. This should be left with the default value of `true` unless you
	// have used the `/addresses/actions/validate` endpoint to validate the address
	// separately prior to creation. If an address is not validated for emergency use
	// upon creation and it is not valid, it will not be able to be used for emergency
	// services.
	ValidateAddress bool `json:"validate_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AddressBook        respjson.Field
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
		ValidateAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressNewResponse struct {
	Data Address `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressGetResponse struct {
	Data Address `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressListResponse struct {
	Data []Address      `json:"data"`
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
func (r AddressListResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressDeleteResponse struct {
	Data Address `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressNewParams struct {
	// The business name associated with the address. An address must have either a
	// first last name or a business name.
	BusinessName string `json:"business_name,required"`
	// The two-character (ISO 3166-1 alpha-2) country code of the address.
	CountryCode string `json:"country_code,required"`
	// The first name associated with the address. An address must have either a first
	// last name or a business name.
	FirstName string `json:"first_name,required"`
	// The last name associated with the address. An address must have either a first
	// last name or a business name.
	LastName string `json:"last_name,required"`
	// The locality of the address. For US addresses, this corresponds to the city of
	// the address.
	Locality string `json:"locality,required"`
	// The primary street address information about the address.
	StreetAddress string `json:"street_address,required"`
	// Indicates whether or not the address should be considered part of your list of
	// addresses that appear for regular use.
	AddressBook param.Opt[bool] `json:"address_book,omitzero"`
	// The locality of the address. For US addresses, this corresponds to the state of
	// the address.
	AdministrativeArea param.Opt[string] `json:"administrative_area,omitzero"`
	// The borough of the address. This field is not used for addresses in the US but
	// is used for some international addresses.
	Borough param.Opt[string] `json:"borough,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Additional street address information about the address such as, but not limited
	// to, unit number or apartment number.
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	// The neighborhood of the address. This field is not used for addresses in the US
	// but is used for some international addresses.
	Neighborhood param.Opt[string] `json:"neighborhood,omitzero"`
	// The phone number associated with the address.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// The postal code of the address.
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// Indicates whether or not the address should be validated for emergency use upon
	// creation or not. This should be left with the default value of `true` unless you
	// have used the `/addresses/actions/validate` endpoint to validate the address
	// separately prior to creation. If an address is not validated for emergency use
	// upon creation and it is not valid, it will not be able to be used for emergency
	// services.
	ValidateAddress param.Opt[bool] `json:"validate_address,omitzero"`
	paramObj
}

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[customer_reference][eq], filter[customer_reference][contains],
	// filter[used_as_emergency], filter[street_address][contains],
	// filter[address_book][eq]
	Filter AddressListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page AddressListParamsPage `query:"page,omitzero" json:"-"`
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
	Sort AddressListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParams]'s query parameters as `url.Values`.
func (r AddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[customer_reference][eq], filter[customer_reference][contains],
// filter[used_as_emergency], filter[street_address][contains],
// filter[address_book][eq]
type AddressListParamsFilter struct {
	// If set as 'true', only addresses used as the emergency address for at least one
	// active phone-number will be returned. When set to 'false', the opposite happens:
	// only addresses not used as the emergency address from phone-numbers will be
	// returned.
	UsedAsEmergency param.Opt[string]                  `query:"used_as_emergency,omitzero" json:"-"`
	AddressBook     AddressListParamsFilterAddressBook `query:"address_book,omitzero" json:"-"`
	// If present, addresses with <code>customer_reference</code> containing the given
	// value will be returned. Matching is not case-sensitive.
	CustomerReference AddressListParamsFilterCustomerReferenceUnion `query:"customer_reference,omitzero" json:"-"`
	StreetAddress     AddressListParamsFilterStreetAddress          `query:"street_address,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParamsFilter]'s query parameters as
// `url.Values`.
func (r AddressListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AddressListParamsFilterAddressBook struct {
	// If present, only returns results with the <code>address_book</code> flag equal
	// to the given value.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParamsFilterAddressBook]'s query parameters as
// `url.Values`.
func (r AddressListParamsFilterAddressBook) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AddressListParamsFilterCustomerReferenceUnion struct {
	OfString                                    param.Opt[string]                               `query:",omitzero,inline"`
	OfAddressListsFilterCustomerReferenceObject *AddressListParamsFilterCustomerReferenceObject `query:",omitzero,inline"`
	paramUnion
}

func (u *AddressListParamsFilterCustomerReferenceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAddressListsFilterCustomerReferenceObject) {
		return u.OfAddressListsFilterCustomerReferenceObject
	}
	return nil
}

type AddressListParamsFilterCustomerReferenceObject struct {
	// Partial match for customer_reference. Matching is not case-sensitive.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Exact match for customer_reference.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParamsFilterCustomerReferenceObject]'s query
// parameters as `url.Values`.
func (r AddressListParamsFilterCustomerReferenceObject) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AddressListParamsFilterStreetAddress struct {
	// If present, addresses with <code>street_address</code> containing the given
	// value will be returned. Matching is not case-sensitive. Requires at least three
	// characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParamsFilterStreetAddress]'s query parameters as
// `url.Values`.
func (r AddressListParamsFilterStreetAddress) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type AddressListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AddressListParamsPage]'s query parameters as `url.Values`.
func (r AddressListParamsPage) URLQuery() (v url.Values, err error) {
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
type AddressListParamsSort string

const (
	AddressListParamsSortCreatedAt     AddressListParamsSort = "created_at"
	AddressListParamsSortFirstName     AddressListParamsSort = "first_name"
	AddressListParamsSortLastName      AddressListParamsSort = "last_name"
	AddressListParamsSortBusinessName  AddressListParamsSort = "business_name"
	AddressListParamsSortStreetAddress AddressListParamsSort = "street_address"
)
