// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingPhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingPhoneNumberService] method instead.
type PortingPhoneNumberService struct {
	Options []option.RequestOption
}

// NewPortingPhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPortingPhoneNumberService(opts ...option.RequestOption) (r PortingPhoneNumberService) {
	r = PortingPhoneNumberService{}
	r.Options = opts
	return
}

// Returns a list of your porting phone numbers.
func (r *PortingPhoneNumberService) List(ctx context.Context, query PortingPhoneNumberListParams, opts ...option.RequestOption) (res *PortingPhoneNumberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "porting_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PortingPhoneNumberListResponse struct {
	Data []PortingPhoneNumberListResponseData `json:"data"`
	Meta PaginationMeta                       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingPhoneNumberListResponseData struct {
	// Activation status
	//
	// Any of "New", "Pending", "Conflict", "Cancel Pending", "Failed", "Concurred",
	// "Activate RDY", "Disconnect Pending", "Concurrence Sent", "Old", "Sending",
	// "Active", "Cancelled".
	ActivationStatus string `json:"activation_status"`
	// E164 formatted phone number
	PhoneNumber string `json:"phone_number"`
	// The type of the phone number
	//
	// Any of "landline", "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType string `json:"phone_number_type"`
	// Specifies whether Telnyx is able to confirm portability this number in the
	// United States & Canada. International phone numbers are provisional by default.
	//
	// Any of "pending", "confirmed", "provisional".
	PortabilityStatus string `json:"portability_status"`
	// Identifies the associated port request
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// The current status of the porting order
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "cancel-pending", "ported", "cancelled".
	PortingOrderStatus string `json:"porting_order_status"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the requirements in a INTL porting order
	//
	// Any of "requirement-info-pending", "requirement-info-under-review",
	// "requirement-info-exception", "approved".
	RequirementsStatus string `json:"requirements_status"`
	// A key to reference this porting order when contacting Telnyx customer support
	SupportKey string `json:"support_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivationStatus   respjson.Field
		PhoneNumber        respjson.Field
		PhoneNumberType    respjson.Field
		PortabilityStatus  respjson.Field
		PortingOrderID     respjson.Field
		PortingOrderStatus respjson.Field
		RecordType         respjson.Field
		RequirementsStatus respjson.Field
		SupportKey         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingPhoneNumberListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[porting_order_status]
	Filter PortingPhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingPhoneNumberListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingPhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r PortingPhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[porting_order_status]
type PortingPhoneNumberListParamsFilter struct {
	// Filter results by porting order status
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "cancel-pending", "ported", "cancelled".
	PortingOrderStatus string `query:"porting_order_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingPhoneNumberListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortingPhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PortingPhoneNumberListParamsFilter](
		"porting_order_status", "draft", "in-process", "submitted", "exception", "foc-date-confirmed", "cancel-pending", "ported", "cancelled",
	)
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingPhoneNumberListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingPhoneNumberListParamsPage]'s query parameters as
// `url.Values`.
func (r PortingPhoneNumberListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
