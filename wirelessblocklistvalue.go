// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Wireless Blocklists operations
//
// WirelessBlocklistValueService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWirelessBlocklistValueService] method instead.
type WirelessBlocklistValueService struct {
	Options []option.RequestOption
}

// NewWirelessBlocklistValueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWirelessBlocklistValueService(opts ...option.RequestOption) (r WirelessBlocklistValueService) {
	r = WirelessBlocklistValueService{}
	r.Options = opts
	return
}

// Retrieve all wireless blocklist values for a given blocklist type.
func (r *WirelessBlocklistValueService) List(ctx context.Context, query WirelessBlocklistValueListParams, opts ...option.RequestOption) (res *WirelessBlocklistValueListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "wireless_blocklist_values"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WirelessBlocklistValueListResponse struct {
	Data WirelessBlocklistValueListResponseDataUnion `json:"data"`
	Meta PaginationMeta                              `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistValueListResponse) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistValueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WirelessBlocklistValueListResponseDataUnion contains all possible properties and
// values from [[]WirelessBlocklistValueListResponseDataCountryItem],
// [[]WirelessBlocklistValueListResponseDataMccItem],
// [[]WirelessBlocklistValueListResponseDataPlmnItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCountry OfMcc OfPlmn]
type WirelessBlocklistValueListResponseDataUnion struct {
	// This field will be present if the value is a
	// [[]WirelessBlocklistValueListResponseDataCountryItem] instead of an object.
	OfCountry []WirelessBlocklistValueListResponseDataCountryItem `json:",inline"`
	// This field will be present if the value is a
	// [[]WirelessBlocklistValueListResponseDataMccItem] instead of an object.
	OfMcc []WirelessBlocklistValueListResponseDataMccItem `json:",inline"`
	// This field will be present if the value is a
	// [[]WirelessBlocklistValueListResponseDataPlmnItem] instead of an object.
	OfPlmn []WirelessBlocklistValueListResponseDataPlmnItem `json:",inline"`
	JSON   struct {
		OfCountry respjson.Field
		OfMcc     respjson.Field
		OfPlmn    respjson.Field
		raw       string
	} `json:"-"`
}

func (u WirelessBlocklistValueListResponseDataUnion) AsCountry() (v []WirelessBlocklistValueListResponseDataCountryItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WirelessBlocklistValueListResponseDataUnion) AsMcc() (v []WirelessBlocklistValueListResponseDataMccItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WirelessBlocklistValueListResponseDataUnion) AsPlmn() (v []WirelessBlocklistValueListResponseDataPlmnItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WirelessBlocklistValueListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *WirelessBlocklistValueListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistValueListResponseDataCountryItem struct {
	// ISO 3166-1 Alpha-2 Country Code.
	Code string `json:"code" api:"required"`
	// The name of the country.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistValueListResponseDataCountryItem) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistValueListResponseDataCountryItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistValueListResponseDataMccItem struct {
	// Mobile Country Code.
	Code string `json:"code" api:"required"`
	// The name of the country.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistValueListResponseDataMccItem) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistValueListResponseDataMccItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistValueListResponseDataPlmnItem struct {
	// Public land mobile network code (MCC + MNC).
	Code string `json:"code" api:"required"`
	// The name of the network.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WirelessBlocklistValueListResponseDataPlmnItem) RawJSON() string { return r.JSON.raw }
func (r *WirelessBlocklistValueListResponseDataPlmnItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessBlocklistValueListParams struct {
	// The Wireless Blocklist type for which to list possible values (e.g., `country`,
	// `mcc`, `plmn`).
	//
	// Any of "country", "mcc", "plmn".
	Type WirelessBlocklistValueListParamsType `query:"type,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [WirelessBlocklistValueListParams]'s query parameters as
// `url.Values`.
func (r WirelessBlocklistValueListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The Wireless Blocklist type for which to list possible values (e.g., `country`,
// `mcc`, `plmn`).
type WirelessBlocklistValueListParamsType string

const (
	WirelessBlocklistValueListParamsTypeCountry WirelessBlocklistValueListParamsType = "country"
	WirelessBlocklistValueListParamsTypeMcc     WirelessBlocklistValueListParamsType = "mcc"
	WirelessBlocklistValueListParamsTypePlmn    WirelessBlocklistValueListParamsType = "plmn"
)
