// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingOrderPhoneNumberConfigurationService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderPhoneNumberConfigurationService] method instead.
type PortingOrderPhoneNumberConfigurationService struct {
	Options []option.RequestOption
}

// NewPortingOrderPhoneNumberConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewPortingOrderPhoneNumberConfigurationService(opts ...option.RequestOption) (r PortingOrderPhoneNumberConfigurationService) {
	r = PortingOrderPhoneNumberConfigurationService{}
	r.Options = opts
	return
}

// Creates a list of phone number configurations.
func (r *PortingOrderPhoneNumberConfigurationService) New(ctx context.Context, body PortingOrderPhoneNumberConfigurationNewParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberConfigurationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting_orders/phone_number_configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of phone number configurations paginated.
func (r *PortingOrderPhoneNumberConfigurationService) List(ctx context.Context, query PortingOrderPhoneNumberConfigurationListParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberConfigurationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting_orders/phone_number_configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PortingOrderPhoneNumberConfigurationNewResponse struct {
	Data []PortingOrderPhoneNumberConfigurationNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberConfigurationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberConfigurationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfigurationNewResponseData struct {
	// Uniquely identifies this phone number configuration
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the associated porting phone number
	PortingPhoneNumberID string `json:"porting_phone_number_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Identifies the associated user bundle
	UserBundleID string `json:"user_bundle_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		PortingPhoneNumberID respjson.Field
		RecordType           respjson.Field
		UpdatedAt            respjson.Field
		UserBundleID         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberConfigurationNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberConfigurationNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfigurationListResponse struct {
	Data []PortingOrderPhoneNumberConfigurationListResponseData `json:"data"`
	Meta PaginationMeta                                         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberConfigurationListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberConfigurationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfigurationListResponseData struct {
	// Uniquely identifies this phone number configuration
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the associated porting phone number
	PortingPhoneNumberID string `json:"porting_phone_number_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Identifies the associated user bundle
	UserBundleID string `json:"user_bundle_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		PortingPhoneNumberID respjson.Field
		RecordType           respjson.Field
		UpdatedAt            respjson.Field
		UserBundleID         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberConfigurationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberConfigurationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfigurationNewParams struct {
	PhoneNumberConfigurations []PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration `json:"phone_number_configurations,omitzero"`
	paramObj
}

func (r PortingOrderPhoneNumberConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberConfigurationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberConfigurationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PortingPhoneNumberID, UserBundleID are required.
type PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration struct {
	// Identifies the porting phone number to be configured.
	PortingPhoneNumberID string `json:"porting_phone_number_id,required" format:"uuid"`
	// Identifies the user bundle to be associated with the porting phone number.
	UserBundleID string `json:"user_bundle_id,required" format:"uuid"`
	paramObj
}

func (r PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfigurationListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[porting_order.status][in][], filter[porting_phone_number][in][],
	// filter[user_bundle_id][in][]
	Filter PortingOrderPhoneNumberConfigurationListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingOrderPhoneNumberConfigurationListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderPhoneNumberConfigurationListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberConfigurationListParams]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberConfigurationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[porting_order.status][in][], filter[porting_phone_number][in][],
// filter[user_bundle_id][in][]
type PortingOrderPhoneNumberConfigurationListParamsFilter struct {
	PortingOrder PortingOrderPhoneNumberConfigurationListParamsFilterPortingOrder `query:"porting_order,omitzero" json:"-"`
	// Filter results by a list of porting phone number IDs
	PortingPhoneNumber []string `query:"porting_phone_number,omitzero" format:"uuid" json:"-"`
	// Filter results by a list of user bundle IDs
	UserBundleID []string `query:"user_bundle_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberConfigurationListParamsFilter]'s
// query parameters as `url.Values`.
func (r PortingOrderPhoneNumberConfigurationListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderPhoneNumberConfigurationListParamsFilterPortingOrder struct {
	// Filter results by specific porting order statuses
	//
	// Any of "activation-in-progress", "cancel-pending", "cancelled", "draft",
	// "exception", "foc-date-confirmed", "in-process", "ported", "submitted".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [PortingOrderPhoneNumberConfigurationListParamsFilterPortingOrder]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberConfigurationListParamsFilterPortingOrder) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingOrderPhoneNumberConfigurationListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberConfigurationListParamsPage]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberConfigurationListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderPhoneNumberConfigurationListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberConfigurationListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberConfigurationListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingOrderPhoneNumberConfigurationListParamsSort](
		"value", "created_at", "-created_at",
	)
}
