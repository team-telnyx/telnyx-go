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
		return
	}
	path := fmt.Sprintf("number_lookup/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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
	CallerName NumberLookupGetResponseDataCallerName `json:"caller_name"`
	Carrier    NumberLookupGetResponseDataCarrier    `json:"carrier"`
	// Region code that matches the specific country calling code
	CountryCode string `json:"country_code"`
	// Unused
	Fraud string `json:"fraud" api:"nullable"`
	// Hyphen-separated national number, preceded by the national destination code
	// (NDC), with a 0 prefix, if an NDC is found
	NationalFormat string `json:"national_format"`
	// E164-formatted phone number
	PhoneNumber string                                 `json:"phone_number"`
	Portability NumberLookupGetResponseDataPortability `json:"portability"`
	// Identifies the type of record
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallerName     respjson.Field
		Carrier        respjson.Field
		CountryCode    respjson.Field
		Fraud          respjson.Field
		NationalFormat respjson.Field
		PhoneNumber    respjson.Field
		Portability    respjson.Field
		RecordType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberLookupGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *NumberLookupGetResponseData) UnmarshalJSON(data []byte) error {
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
	// Type of number
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
	// Specifies the type of number lookup to be performed
	//
	// Any of "carrier", "caller-name".
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

// Specifies the type of number lookup to be performed
type NumberLookupGetParamsType string

const (
	NumberLookupGetParamsTypeCarrier    NumberLookupGetParamsType = "carrier"
	NumberLookupGetParamsTypeCallerName NumberLookupGetParamsType = "caller-name"
)
