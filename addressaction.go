// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// AddressActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressActionService] method instead.
type AddressActionService struct {
	Options []option.RequestOption
}

// NewAddressActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressActionService(opts ...option.RequestOption) (r AddressActionService) {
	r = AddressActionService{}
	r.Options = opts
	return
}

// Accepts this address suggestion as a new emergency address for Operator Connect
// and finishes the uploads of the numbers associated with it to Microsoft.
func (r *AddressActionService) AcceptSuggestions(ctx context.Context, addressUuid string, body AddressActionAcceptSuggestionsParams, opts ...option.RequestOption) (res *AddressActionAcceptSuggestionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if addressUuid == "" {
		err = errors.New("missing required address_uuid parameter")
		return
	}
	path := fmt.Sprintf("addresses/%s/actions/accept_suggestions", addressUuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Validates an address for emergency services.
func (r *AddressActionService) Validate(ctx context.Context, body AddressActionValidateParams, opts ...option.RequestOption) (res *AddressActionValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "addresses/actions/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AddressActionAcceptSuggestionsResponse struct {
	Data AddressActionAcceptSuggestionsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressActionAcceptSuggestionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressActionAcceptSuggestionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressActionAcceptSuggestionsResponseData struct {
	// The UUID of the location.
	ID string `json:"id" format:"uuid"`
	// Indicates if the address suggestions are accepted.
	Accepted bool `json:"accepted"`
	// Any of "address_suggestion".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Accepted    respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressActionAcceptSuggestionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AddressActionAcceptSuggestionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressActionValidateResponse struct {
	Data AddressActionValidateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressActionValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *AddressActionValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressActionValidateResponseData struct {
	// Indicates whether an address is valid or invalid.
	//
	// Any of "valid", "invalid".
	Result string `json:"result,required"`
	// Provides normalized address when available.
	Suggested AddressActionValidateResponseDataSuggested `json:"suggested,required"`
	Errors    []shared.APIError                          `json:"errors"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Suggested   respjson.Field
		Errors      respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressActionValidateResponseData) RawJSON() string { return r.JSON.raw }
func (r *AddressActionValidateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provides normalized address when available.
type AddressActionValidateResponseDataSuggested struct {
	// The locality of the address. For US addresses, this corresponds to the state of
	// the address.
	AdministrativeArea string `json:"administrative_area"`
	// The two-character (ISO 3166-1 alpha-2) country code of the address.
	CountryCode string `json:"country_code"`
	// Additional street address information about the address such as, but not limited
	// to, unit number or apartment number.
	ExtendedAddress string `json:"extended_address"`
	// The locality of the address. For US addresses, this corresponds to the city of
	// the address.
	Locality string `json:"locality"`
	// The postal code of the address.
	PostalCode string `json:"postal_code"`
	// The primary street address information about the address.
	StreetAddress string         `json:"street_address"`
	ExtraFields   map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		CountryCode        respjson.Field
		ExtendedAddress    respjson.Field
		Locality           respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressActionValidateResponseDataSuggested) RawJSON() string { return r.JSON.raw }
func (r *AddressActionValidateResponseDataSuggested) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressActionAcceptSuggestionsParams struct {
	// The ID of the address.
	ID param.Opt[string] `json:"id,omitzero"`
	paramObj
}

func (r AddressActionAcceptSuggestionsParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressActionAcceptSuggestionsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressActionAcceptSuggestionsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressActionValidateParams struct {
	// The two-character (ISO 3166-1 alpha-2) country code of the address.
	CountryCode string `json:"country_code,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// The primary street address information about the address.
	StreetAddress string `json:"street_address,required"`
	// The locality of the address. For US addresses, this corresponds to the state of
	// the address.
	AdministrativeArea param.Opt[string] `json:"administrative_area,omitzero"`
	// Additional street address information about the address such as, but not limited
	// to, unit number or apartment number.
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	// The locality of the address. For US addresses, this corresponds to the city of
	// the address.
	Locality param.Opt[string] `json:"locality,omitzero"`
	paramObj
}

func (r AddressActionValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow AddressActionValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressActionValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
