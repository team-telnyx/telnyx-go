// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
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
func (r *PortingPhoneNumberService) List(ctx context.Context, query PortingPhoneNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortingPhoneNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "porting_phone_numbers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of your porting phone numbers.
func (r *PortingPhoneNumberService) ListAutoPaging(ctx context.Context, query PortingPhoneNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortingPhoneNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type PortingPhoneNumberListResponse struct {
	// Activation status
	//
	// Any of "New", "Pending", "Conflict", "Cancel Pending", "Failed", "Concurred",
	// "Activate RDY", "Disconnect Pending", "Concurrence Sent", "Old", "Sending",
	// "Active", "Cancelled".
	ActivationStatus PortingPhoneNumberListResponseActivationStatus `json:"activation_status"`
	// E164 formatted phone number
	PhoneNumber string `json:"phone_number"`
	// The type of the phone number
	//
	// Any of "landline", "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType PortingPhoneNumberListResponsePhoneNumberType `json:"phone_number_type"`
	// Specifies whether Telnyx is able to confirm portability this number in the
	// United States & Canada. International phone numbers are provisional by default.
	//
	// Any of "pending", "confirmed", "provisional".
	PortabilityStatus PortingPhoneNumberListResponsePortabilityStatus `json:"portability_status"`
	// Identifies the associated port request
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// The current status of the porting order
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "cancel-pending", "ported", "cancelled".
	PortingOrderStatus PortingPhoneNumberListResponsePortingOrderStatus `json:"porting_order_status"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The current status of the requirements in a INTL porting order
	//
	// Any of "requirement-info-pending", "requirement-info-under-review",
	// "requirement-info-exception", "approved".
	RequirementsStatus PortingPhoneNumberListResponseRequirementsStatus `json:"requirements_status"`
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
func (r PortingPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activation status
type PortingPhoneNumberListResponseActivationStatus string

const (
	PortingPhoneNumberListResponseActivationStatusNew               PortingPhoneNumberListResponseActivationStatus = "New"
	PortingPhoneNumberListResponseActivationStatusPending           PortingPhoneNumberListResponseActivationStatus = "Pending"
	PortingPhoneNumberListResponseActivationStatusConflict          PortingPhoneNumberListResponseActivationStatus = "Conflict"
	PortingPhoneNumberListResponseActivationStatusCancelPending     PortingPhoneNumberListResponseActivationStatus = "Cancel Pending"
	PortingPhoneNumberListResponseActivationStatusFailed            PortingPhoneNumberListResponseActivationStatus = "Failed"
	PortingPhoneNumberListResponseActivationStatusConcurred         PortingPhoneNumberListResponseActivationStatus = "Concurred"
	PortingPhoneNumberListResponseActivationStatusActivateRdy       PortingPhoneNumberListResponseActivationStatus = "Activate RDY"
	PortingPhoneNumberListResponseActivationStatusDisconnectPending PortingPhoneNumberListResponseActivationStatus = "Disconnect Pending"
	PortingPhoneNumberListResponseActivationStatusConcurrenceSent   PortingPhoneNumberListResponseActivationStatus = "Concurrence Sent"
	PortingPhoneNumberListResponseActivationStatusOld               PortingPhoneNumberListResponseActivationStatus = "Old"
	PortingPhoneNumberListResponseActivationStatusSending           PortingPhoneNumberListResponseActivationStatus = "Sending"
	PortingPhoneNumberListResponseActivationStatusActive            PortingPhoneNumberListResponseActivationStatus = "Active"
	PortingPhoneNumberListResponseActivationStatusCancelled         PortingPhoneNumberListResponseActivationStatus = "Cancelled"
)

// The type of the phone number
type PortingPhoneNumberListResponsePhoneNumberType string

const (
	PortingPhoneNumberListResponsePhoneNumberTypeLandline   PortingPhoneNumberListResponsePhoneNumberType = "landline"
	PortingPhoneNumberListResponsePhoneNumberTypeLocal      PortingPhoneNumberListResponsePhoneNumberType = "local"
	PortingPhoneNumberListResponsePhoneNumberTypeMobile     PortingPhoneNumberListResponsePhoneNumberType = "mobile"
	PortingPhoneNumberListResponsePhoneNumberTypeNational   PortingPhoneNumberListResponsePhoneNumberType = "national"
	PortingPhoneNumberListResponsePhoneNumberTypeSharedCost PortingPhoneNumberListResponsePhoneNumberType = "shared_cost"
	PortingPhoneNumberListResponsePhoneNumberTypeTollFree   PortingPhoneNumberListResponsePhoneNumberType = "toll_free"
)

// Specifies whether Telnyx is able to confirm portability this number in the
// United States & Canada. International phone numbers are provisional by default.
type PortingPhoneNumberListResponsePortabilityStatus string

const (
	PortingPhoneNumberListResponsePortabilityStatusPending     PortingPhoneNumberListResponsePortabilityStatus = "pending"
	PortingPhoneNumberListResponsePortabilityStatusConfirmed   PortingPhoneNumberListResponsePortabilityStatus = "confirmed"
	PortingPhoneNumberListResponsePortabilityStatusProvisional PortingPhoneNumberListResponsePortabilityStatus = "provisional"
)

// The current status of the porting order
type PortingPhoneNumberListResponsePortingOrderStatus string

const (
	PortingPhoneNumberListResponsePortingOrderStatusDraft            PortingPhoneNumberListResponsePortingOrderStatus = "draft"
	PortingPhoneNumberListResponsePortingOrderStatusInProcess        PortingPhoneNumberListResponsePortingOrderStatus = "in-process"
	PortingPhoneNumberListResponsePortingOrderStatusSubmitted        PortingPhoneNumberListResponsePortingOrderStatus = "submitted"
	PortingPhoneNumberListResponsePortingOrderStatusException        PortingPhoneNumberListResponsePortingOrderStatus = "exception"
	PortingPhoneNumberListResponsePortingOrderStatusFocDateConfirmed PortingPhoneNumberListResponsePortingOrderStatus = "foc-date-confirmed"
	PortingPhoneNumberListResponsePortingOrderStatusCancelPending    PortingPhoneNumberListResponsePortingOrderStatus = "cancel-pending"
	PortingPhoneNumberListResponsePortingOrderStatusPorted           PortingPhoneNumberListResponsePortingOrderStatus = "ported"
	PortingPhoneNumberListResponsePortingOrderStatusCancelled        PortingPhoneNumberListResponsePortingOrderStatus = "cancelled"
)

// The current status of the requirements in a INTL porting order
type PortingPhoneNumberListResponseRequirementsStatus string

const (
	PortingPhoneNumberListResponseRequirementsStatusRequirementInfoPending     PortingPhoneNumberListResponseRequirementsStatus = "requirement-info-pending"
	PortingPhoneNumberListResponseRequirementsStatusRequirementInfoUnderReview PortingPhoneNumberListResponseRequirementsStatus = "requirement-info-under-review"
	PortingPhoneNumberListResponseRequirementsStatusRequirementInfoException   PortingPhoneNumberListResponseRequirementsStatus = "requirement-info-exception"
	PortingPhoneNumberListResponseRequirementsStatusApproved                   PortingPhoneNumberListResponseRequirementsStatus = "approved"
)

type PortingPhoneNumberListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[porting_order_status]
	Filter PortingPhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
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
