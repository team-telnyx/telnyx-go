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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Look up phone number data
//
// NumberLookupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberLookupService] method instead.
type NumberLookupService struct {
	Options []option.RequestOption
}

// NewNumberLookupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumberLookupService(opts ...option.RequestOption) (r NumberLookupService) {
	r = NumberLookupService{}
	r.Options = opts
	return
}

// Returns information about the provided phone number.
func (r *NumberLookupService) Get(ctx context.Context, phoneNumber string, query NumberLookupGetParams, opts ...option.RequestOption) (res *NumberLookupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("number_lookup/%s", url.PathEscape(phoneNumber))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type NumberLookupGetResponse struct {
	Data NumberLookupGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetResponseData struct {
	// Identity information for the phone number's owner, including name and postal
	// address. Populated when `type=caller-identity` is requested.
	CallerIdentity NumberLookupGetResponseDataCallerIdentity `json:"caller_identity"`
	CallerName     NumberLookupGetResponseDataCallerName     `json:"caller_name"`
	Carrier        NumberLookupGetResponseDataCarrier        `json:"carrier"`
	// Region code that matches the specific country calling code
	CountryCode string `json:"country_code"`
	// Enhanced caller name for the phone number's owner. Populated when
	// `type=enhanced-caller-name` is requested.
	EnhancedCallerName NumberLookupGetResponseDataEnhancedCallerName `json:"enhanced_caller_name"`
	// Unused
	Fraud string `json:"fraud" api:"nullable"`
	// Hyphen-separated national number, preceded by the national destination code
	// (NDC), with a 0 prefix, if an NDC is found
	NationalFormat string `json:"national_format"`
	// Network number ID (NNID) override for the phone number, when one applies
	// (available only to accounts enabled for NNID); otherwise null.
	NnidOverride NumberLookupGetResponseDataNnidOverride `json:"nnid_override" api:"nullable"`
	// E164-formatted phone number
	PhoneNumber string                                 `json:"phone_number"`
	Portability NumberLookupGetResponseDataPortability `json:"portability"`
	// Identifies the type of record
	RecordType string `json:"record_type"`
	// Indicates whether the phone number is valid.
	ValidNumber bool `json:"valid_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerIdentity     respjson.Field
		CallerName         respjson.Field
		Carrier            respjson.Field
		CountryCode        respjson.Field
		EnhancedCallerName respjson.Field
		Fraud              respjson.Field
		NationalFormat     respjson.Field
		NnidOverride       respjson.Field
		PhoneNumber        respjson.Field
		Portability        respjson.Field
		RecordType         respjson.Field
		ValidNumber        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identity information for the phone number's owner, including name and postal
// address. Populated when `type=caller-identity` is requested.
type NumberLookupGetResponseDataCallerIdentity struct {
	// Postal address associated with the phone number's owner.
	Address NumberLookupGetResponseDataCallerIdentityAddress `json:"address"`
	// How strongly the phone number is linked to the returned name, on a scale of 50
	// to 90 (higher means more supporting signals). Reflects signal strength, not
	// probability: 70 indicates stronger, more consistent evidence than 60, not a 70%
	// chance of being correct. As a guide, 70+ is high confidence, 60 moderate, and 50
	// limited. Below 50 no name is returned, so scores under 50 are never present.
	ConfidenceScore float64 `json:"confidence_score" api:"nullable"`
	// A lookup-specific error code if the identity could not be resolved; otherwise
	// null.
	ErrorCode string `json:"error_code" api:"nullable"`
	// The name associated with the phone number.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address         respjson.Field
		ConfidenceScore respjson.Field
		ErrorCode       respjson.Field
		Name            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataCallerIdentity) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataCallerIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Postal address associated with the phone number's owner.
type NumberLookupGetResponseDataCallerIdentityAddress struct {
	// City.
	City string `json:"city" api:"nullable"`
	// Two-letter ISO 3166-1 country code.
	CountryCode string `json:"country_code" api:"nullable"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"nullable"`
	// Two-letter state or province code.
	StateCode string `json:"state_code" api:"nullable"`
	// First line of the street address.
	StreetLine1 string `json:"street_line_1" api:"nullable"`
	// Second line of the street address, if any.
	StreetLine2 string `json:"street_line_2" api:"nullable"`
	// ZIP+4 extension, if available.
	Zip4 string `json:"zip4" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		CountryCode respjson.Field
		PostalCode  respjson.Field
		StateCode   respjson.Field
		StreetLine1 respjson.Field
		StreetLine2 respjson.Field
		Zip4        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataCallerIdentityAddress) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataCallerIdentityAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetResponseDataCallerName struct {
	// The name of the requested phone number's owner as per the CNAM database
	CallerName string `json:"caller_name"`
	// A caller-name lookup specific error code, expressed as a stringified 5-digit
	// integer
	ErrorCode string `json:"error_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerName  respjson.Field
		ErrorCode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataCallerName) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataCallerName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetResponseDataCarrier struct {
	// Unused
	ErrorCode string `json:"error_code" api:"nullable"`
	// Region code that matches the specific country calling code if the requested
	// phone number type is mobile
	MobileCountryCode string `json:"mobile_country_code"`
	// National destination code (NDC), with a 0 prefix, if an NDC is found and the
	// requested phone number type is mobile
	MobileNetworkCode string `json:"mobile_network_code"`
	// SPID (Service Provider ID) name, if the requested phone number has been ported;
	// otherwise, the name of carrier who owns the phone number block
	Name string `json:"name"`
	// If known to Telnyx and applicable, the primary network carrier.
	NormalizedCarrier string `json:"normalized_carrier"`
	// A phone number type that identifies the type of service associated with the
	// requested phone number
	//
	// Any of "fixed line", "mobile", "voip", "fixed line or mobile", "toll free",
	// "premium rate", "shared cost", "personal number", "pager", "uan", "voicemail",
	// "unknown".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCode         respjson.Field
		MobileCountryCode respjson.Field
		MobileNetworkCode respjson.Field
		Name              respjson.Field
		NormalizedCarrier respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataCarrier) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enhanced caller name for the phone number's owner. Populated when
// `type=enhanced-caller-name` is requested.
type NumberLookupGetResponseDataEnhancedCallerName struct {
	// The name associated with the phone number.
	CallerName string `json:"caller_name" api:"nullable"`
	// A lookup-specific error code if the enhanced caller name could not be resolved;
	// otherwise null.
	ErrorCode string `json:"error_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerName  respjson.Field
		ErrorCode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataEnhancedCallerName) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataEnhancedCallerName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Network number ID (NNID) override for the phone number, when one applies
// (available only to accounts enabled for NNID); otherwise null.
type NumberLookupGetResponseDataNnidOverride struct {
	// Name associated with the NNID.
	Name string `json:"name"`
	// The network number ID.
	Nnid string `json:"nnid"`
	// The NNID type.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Nnid        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataNnidOverride) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataNnidOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetResponseDataPortability struct {
	// Alternative SPID (Service Provider ID). Often used when a carrier is using a
	// number from another carrier
	Altspid string `json:"altspid"`
	// Alternative service provider name
	AltspidCarrierName string `json:"altspid_carrier_name"`
	// Alternative service provider type
	AltspidCarrierType string `json:"altspid_carrier_type"`
	// City name extracted from the locality in the Local Exchange Routing Guide (LERG)
	// database
	City string `json:"city"`
	// The line type of the phone number. Uses the same value set as carrier.type; see
	// the API reference for the full list.
	//
	// Any of "fixed line", "mobile", "voip", "fixed line or mobile", "toll free",
	// "premium rate", "shared cost", "personal number", "pager", "uan", "voicemail",
	// "unknown".
	LineType string `json:"line_type"`
	// Local Routing Number, if assigned to the requested phone number
	Lrn string `json:"lrn"`
	// Operating Company Name (OCN) as per the Local Exchange Routing Guide (LERG)
	// database
	Ocn string `json:"ocn"`
	// ISO-formatted date when the requested phone number has been ported
	PortedDate string `json:"ported_date"`
	// Indicates whether or not the requested phone number has been ported
	//
	// Any of "Y", "N", "".
	PortedStatus string `json:"ported_status"`
	// SPID (Service Provider ID)
	Spid string `json:"spid"`
	// Service provider name
	SpidCarrierName string `json:"spid_carrier_name"`
	// Service provider type
	SpidCarrierType string `json:"spid_carrier_type"`
	State           string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Altspid            respjson.Field
		AltspidCarrierName respjson.Field
		AltspidCarrierType respjson.Field
		City               respjson.Field
		LineType           respjson.Field
		Lrn                respjson.Field
		Ocn                respjson.Field
		PortedDate         respjson.Field
		PortedStatus       respjson.Field
		Spid               respjson.Field
		SpidCarrierName    respjson.Field
		SpidCarrierType    respjson.Field
		State              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseDataPortability) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseDataPortability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberLookupGetParams struct {
	// Specifies the type of number lookup to be performed. May be supplied more than
	// once to request several data sources in a single call (for example,
	// `?type=carrier&type=enhanced-caller-name`). Each requested type populates the
	// corresponding object in the response.
	//
	// Any of "carrier", "caller-name", "caller-identity", "enhanced-caller-name".
	Type NumberLookupGetParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NumberLookupGetParams]'s query parameters as `url.Values`.
func (r NumberLookupGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the type of number lookup to be performed. May be supplied more than
// once to request several data sources in a single call (for example,
// `?type=carrier&type=enhanced-caller-name`). Each requested type populates the
// corresponding object in the response.
type NumberLookupGetParamsType string

const (
	NumberLookupGetParamsTypeCarrier            NumberLookupGetParamsType = "carrier"
	NumberLookupGetParamsTypeCallerName         NumberLookupGetParamsType = "caller-name"
	NumberLookupGetParamsTypeCallerIdentity     NumberLookupGetParamsType = "caller-identity"
	NumberLookupGetParamsTypeEnhancedCallerName NumberLookupGetParamsType = "enhanced-caller-name"
)
