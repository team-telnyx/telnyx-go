// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// PortingOrderPhoneNumberExtensionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderPhoneNumberExtensionService] method instead.
type PortingOrderPhoneNumberExtensionService struct {
	Options []option.RequestOption
}

// NewPortingOrderPhoneNumberExtensionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPortingOrderPhoneNumberExtensionService(opts ...option.RequestOption) (r PortingOrderPhoneNumberExtensionService) {
	r = PortingOrderPhoneNumberExtensionService{}
	r.Options = opts
	return
}

// Creates a new phone number extension.
func (r *PortingOrderPhoneNumberExtensionService) New(ctx context.Context, portingOrderID string, body PortingOrderPhoneNumberExtensionNewParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberExtensionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_extensions", portingOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of all phone number extensions of a porting order.
func (r *PortingOrderPhoneNumberExtensionService) List(ctx context.Context, portingOrderID string, query PortingOrderPhoneNumberExtensionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortingPhoneNumberExtension], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if portingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_extensions", portingOrderID)
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

// Returns a list of all phone number extensions of a porting order.
func (r *PortingOrderPhoneNumberExtensionService) ListAutoPaging(ctx context.Context, portingOrderID string, query PortingOrderPhoneNumberExtensionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortingPhoneNumberExtension] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, portingOrderID, query, opts...))
}

// Deletes a phone number extension.
func (r *PortingOrderPhoneNumberExtensionService) Delete(ctx context.Context, id string, body PortingOrderPhoneNumberExtensionDeleteParams, opts ...option.RequestOption) (res *PortingOrderPhoneNumberExtensionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.PortingOrderID == "" {
		err = errors.New("missing required porting_order_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/phone_number_extensions/%s", body.PortingOrderID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PortingPhoneNumberExtension struct {
	// Uniquely identifies this porting phone number extension.
	ID string `json:"id" format:"uuid"`
	// Specifies the activation ranges for this porting phone number extension. The
	// activation range must be within the extension range and should not overlap with
	// other activation ranges.
	ActivationRanges []PortingPhoneNumberExtensionActivationRange `json:"activation_ranges"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies the extension range for this porting phone number extension.
	ExtensionRange PortingPhoneNumberExtensionExtensionRange `json:"extension_range"`
	// Identifies the porting phone number associated with this porting phone number
	// extension.
	PortingPhoneNumberID string `json:"porting_phone_number_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ActivationRanges     respjson.Field
		CreatedAt            respjson.Field
		ExtensionRange       respjson.Field
		PortingPhoneNumberID respjson.Field
		RecordType           respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberExtension) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingPhoneNumberExtensionActivationRange struct {
	// Specifies the end of the activation range. It must be no more than the end of
	// the extension range.
	EndAt int64 `json:"end_at"`
	// Specifies the start of the activation range. Must be greater or equal the start
	// of the extension range.
	StartAt int64 `json:"start_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberExtensionActivationRange) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberExtensionActivationRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the extension range for this porting phone number extension.
type PortingPhoneNumberExtensionExtensionRange struct {
	// Specifies the end of the extension range for this porting phone number
	// extension.
	EndAt int64 `json:"end_at"`
	// Specifies the start of the extension range for this porting phone number
	// extension.
	StartAt int64 `json:"start_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingPhoneNumberExtensionExtensionRange) RawJSON() string { return r.JSON.raw }
func (r *PortingPhoneNumberExtensionExtensionRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberExtensionNewResponse struct {
	Data PortingPhoneNumberExtension `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberExtensionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberExtensionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberExtensionDeleteResponse struct {
	Data PortingPhoneNumberExtension `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberExtensionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberExtensionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberExtensionNewParams struct {
	// Specifies the activation ranges for this porting phone number extension. The
	// activation range must be within the extension range and should not overlap with
	// other activation ranges.
	ActivationRanges []PortingOrderPhoneNumberExtensionNewParamsActivationRange `json:"activation_ranges,omitzero,required"`
	ExtensionRange   PortingOrderPhoneNumberExtensionNewParamsExtensionRange    `json:"extension_range,omitzero,required"`
	// Identifies the porting phone number associated with this porting phone number
	// extension.
	PortingPhoneNumberID string `json:"porting_phone_number_id,required" format:"uuid"`
	paramObj
}

func (r PortingOrderPhoneNumberExtensionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberExtensionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberExtensionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndAt, StartAt are required.
type PortingOrderPhoneNumberExtensionNewParamsActivationRange struct {
	// Specifies the end of the activation range. It must be no more than the end of
	// the extension range.
	EndAt int64 `json:"end_at,required"`
	// Specifies the start of the activation range. Must be greater or equal the start
	// of the extension range.
	StartAt int64 `json:"start_at,required"`
	paramObj
}

func (r PortingOrderPhoneNumberExtensionNewParamsActivationRange) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberExtensionNewParamsActivationRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberExtensionNewParamsActivationRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EndAt, StartAt are required.
type PortingOrderPhoneNumberExtensionNewParamsExtensionRange struct {
	// Specifies the end of the extension range for this porting phone number
	// extension.
	EndAt int64 `json:"end_at,required"`
	// Specifies the start of the extension range for this porting phone number
	// extension.
	StartAt int64 `json:"start_at,required"`
	paramObj
}

func (r PortingOrderPhoneNumberExtensionNewParamsExtensionRange) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberExtensionNewParamsExtensionRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberExtensionNewParamsExtensionRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberExtensionListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[porting_phone_number_id]
	Filter PortingOrderPhoneNumberExtensionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderPhoneNumberExtensionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberExtensionListParams]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberExtensionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[porting_phone_number_id]
type PortingOrderPhoneNumberExtensionListParamsFilter struct {
	// Filter results by porting phone number id
	PortingPhoneNumberID param.Opt[string] `query:"porting_phone_number_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberExtensionListParamsFilter]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberExtensionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderPhoneNumberExtensionListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order
	//
	// Any of "-created_at", "created_at".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderPhoneNumberExtensionListParamsSort]'s query
// parameters as `url.Values`.
func (r PortingOrderPhoneNumberExtensionListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderPhoneNumberExtensionDeleteParams struct {
	PortingOrderID string `path:"porting_order_id,required" format:"uuid" json:"-"`
	paramObj
}
