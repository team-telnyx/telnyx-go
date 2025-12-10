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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// InexplicitNumberOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInexplicitNumberOrderService] method instead.
type InexplicitNumberOrderService struct {
	Options []option.RequestOption
}

// NewInexplicitNumberOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInexplicitNumberOrderService(opts ...option.RequestOption) (r InexplicitNumberOrderService) {
	r = InexplicitNumberOrderService{}
	r.Options = opts
	return
}

// Create an inexplicit number order to programmatically purchase phone numbers
// without specifying exact numbers.
func (r *InexplicitNumberOrderService) New(ctx context.Context, body InexplicitNumberOrderNewParams, opts ...option.RequestOption) (res *InexplicitNumberOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inexplicit_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing inexplicit number order by ID.
func (r *InexplicitNumberOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *InexplicitNumberOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("inexplicit_number_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a paginated list of inexplicit number orders.
func (r *InexplicitNumberOrderService) List(ctx context.Context, query InexplicitNumberOrderListParams, opts ...option.RequestOption) (res *InexplicitNumberOrderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "inexplicit_number_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InexplicitNumberOrderNewResponse struct {
	Data InexplicitNumberOrderNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderNewResponseData struct {
	// Unique identifier for the inexplicit number order
	ID string `json:"id"`
	// Billing group id to apply to phone numbers that are purchased
	BillingGroupID string `json:"billing_group_id"`
	// Connection id to apply to phone numbers that are purchased
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reference label for the customer
	CustomerReference string `json:"customer_reference"`
	// Messaging profile id to apply to phone numbers that are purchased
	MessagingProfileID string                                              `json:"messaging_profile_id"`
	OrderingGroups     []InexplicitNumberOrderNewResponseDataOrderingGroup `json:"ordering_groups"`
	// ISO 8601 formatted date indicating when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		OrderingGroups     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderNewResponseDataOrderingGroup struct {
	// Filter for phone numbers in a given state / province
	AdministrativeArea string `json:"administrative_area"`
	// Quantity of phone numbers allocated
	CountAllocated int64 `json:"count_allocated"`
	// Quantity of phone numbers requested
	CountRequested int64 `json:"count_requested"`
	// Country where you would like to purchase phone numbers
	CountryISO string `json:"country_iso"`
	// ISO 8601 formatted date indicating when the ordering group was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Error reason if applicable
	ErrorReason string `json:"error_reason"`
	// Filter to exclude phone numbers that are currently on hold/reserved for your
	// account.
	ExcludeHeldNumbers bool `json:"exclude_held_numbers"`
	// Filter by area code
	NationalDestinationCode string `json:"national_destination_code"`
	// Array of orders created to fulfill the inexplicit order
	Orders []InexplicitNumberOrderNewResponseDataOrderingGroupOrder `json:"orders"`
	// Number type
	PhoneNumberType string `json:"phone_number_type"`
	// Filter for phone numbers that contain the digits specified
	PhoneNumberContains string `json:"phone_number[contains]"`
	// Filter by the ending digits of the phone number
	PhoneNumberEndsWith string `json:"phone_number[ends_with]"`
	// Filter by the starting digits of the phone number
	PhoneNumberStartsWith string `json:"phone_number[starts_with]"`
	// Filter to exclude phone numbers that need additional time after to purchase to
	// activate. Only applicable for +1 toll_free numbers.
	Quickship bool `json:"quickship"`
	// Status of the ordering group
	//
	// Any of "pending", "processing", "failed", "success", "partial_success".
	Status string `json:"status"`
	// Ordering strategy used
	//
	// Any of "always", "never".
	Strategy string `json:"strategy"`
	// ISO 8601 formatted date indicating when the ordering group was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea      respjson.Field
		CountAllocated          respjson.Field
		CountRequested          respjson.Field
		CountryISO              respjson.Field
		CreatedAt               respjson.Field
		ErrorReason             respjson.Field
		ExcludeHeldNumbers      respjson.Field
		NationalDestinationCode respjson.Field
		Orders                  respjson.Field
		PhoneNumberType         respjson.Field
		PhoneNumberContains     respjson.Field
		PhoneNumberEndsWith     respjson.Field
		PhoneNumberStartsWith   respjson.Field
		Quickship               respjson.Field
		Status                  respjson.Field
		Strategy                respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderNewResponseDataOrderingGroup) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderNewResponseDataOrderingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderNewResponseDataOrderingGroupOrder struct {
	// ID of the main number order
	NumberOrderID string `json:"number_order_id,required" format:"uuid"`
	// Array of sub number order IDs
	SubNumberOrderIDs []string `json:"sub_number_order_ids,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOrderID     respjson.Field
		SubNumberOrderIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderNewResponseDataOrderingGroupOrder) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderNewResponseDataOrderingGroupOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderGetResponse struct {
	Data InexplicitNumberOrderGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderGetResponseData struct {
	// Unique identifier for the inexplicit number order
	ID string `json:"id"`
	// Billing group id to apply to phone numbers that are purchased
	BillingGroupID string `json:"billing_group_id"`
	// Connection id to apply to phone numbers that are purchased
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reference label for the customer
	CustomerReference string `json:"customer_reference"`
	// Messaging profile id to apply to phone numbers that are purchased
	MessagingProfileID string                                              `json:"messaging_profile_id"`
	OrderingGroups     []InexplicitNumberOrderGetResponseDataOrderingGroup `json:"ordering_groups"`
	// ISO 8601 formatted date indicating when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		OrderingGroups     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderGetResponseDataOrderingGroup struct {
	// Filter for phone numbers in a given state / province
	AdministrativeArea string `json:"administrative_area"`
	// Quantity of phone numbers allocated
	CountAllocated int64 `json:"count_allocated"`
	// Quantity of phone numbers requested
	CountRequested int64 `json:"count_requested"`
	// Country where you would like to purchase phone numbers
	CountryISO string `json:"country_iso"`
	// ISO 8601 formatted date indicating when the ordering group was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Error reason if applicable
	ErrorReason string `json:"error_reason"`
	// Filter to exclude phone numbers that are currently on hold/reserved for your
	// account.
	ExcludeHeldNumbers bool `json:"exclude_held_numbers"`
	// Filter by area code
	NationalDestinationCode string `json:"national_destination_code"`
	// Array of orders created to fulfill the inexplicit order
	Orders []InexplicitNumberOrderGetResponseDataOrderingGroupOrder `json:"orders"`
	// Number type
	PhoneNumberType string `json:"phone_number_type"`
	// Filter for phone numbers that contain the digits specified
	PhoneNumberContains string `json:"phone_number[contains]"`
	// Filter by the ending digits of the phone number
	PhoneNumberEndsWith string `json:"phone_number[ends_with]"`
	// Filter by the starting digits of the phone number
	PhoneNumberStartsWith string `json:"phone_number[starts_with]"`
	// Filter to exclude phone numbers that need additional time after to purchase to
	// activate. Only applicable for +1 toll_free numbers.
	Quickship bool `json:"quickship"`
	// Status of the ordering group
	//
	// Any of "pending", "processing", "failed", "success", "partial_success".
	Status string `json:"status"`
	// Ordering strategy used
	//
	// Any of "always", "never".
	Strategy string `json:"strategy"`
	// ISO 8601 formatted date indicating when the ordering group was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea      respjson.Field
		CountAllocated          respjson.Field
		CountRequested          respjson.Field
		CountryISO              respjson.Field
		CreatedAt               respjson.Field
		ErrorReason             respjson.Field
		ExcludeHeldNumbers      respjson.Field
		NationalDestinationCode respjson.Field
		Orders                  respjson.Field
		PhoneNumberType         respjson.Field
		PhoneNumberContains     respjson.Field
		PhoneNumberEndsWith     respjson.Field
		PhoneNumberStartsWith   respjson.Field
		Quickship               respjson.Field
		Status                  respjson.Field
		Strategy                respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderGetResponseDataOrderingGroup) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderGetResponseDataOrderingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderGetResponseDataOrderingGroupOrder struct {
	// ID of the main number order
	NumberOrderID string `json:"number_order_id,required" format:"uuid"`
	// Array of sub number order IDs
	SubNumberOrderIDs []string `json:"sub_number_order_ids,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOrderID     respjson.Field
		SubNumberOrderIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderGetResponseDataOrderingGroupOrder) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderGetResponseDataOrderingGroupOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderListResponse struct {
	Data []InexplicitNumberOrderListResponseData `json:"data"`
	Meta PaginationMeta                          `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderListResponseData struct {
	// Unique identifier for the inexplicit number order
	ID string `json:"id"`
	// Billing group id to apply to phone numbers that are purchased
	BillingGroupID string `json:"billing_group_id"`
	// Connection id to apply to phone numbers that are purchased
	ConnectionID string `json:"connection_id"`
	// ISO 8601 formatted date indicating when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reference label for the customer
	CustomerReference string `json:"customer_reference"`
	// Messaging profile id to apply to phone numbers that are purchased
	MessagingProfileID string                                               `json:"messaging_profile_id"`
	OrderingGroups     []InexplicitNumberOrderListResponseDataOrderingGroup `json:"ordering_groups"`
	// ISO 8601 formatted date indicating when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		CustomerReference  respjson.Field
		MessagingProfileID respjson.Field
		OrderingGroups     respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderListResponseDataOrderingGroup struct {
	// Filter for phone numbers in a given state / province
	AdministrativeArea string `json:"administrative_area"`
	// Quantity of phone numbers allocated
	CountAllocated int64 `json:"count_allocated"`
	// Quantity of phone numbers requested
	CountRequested int64 `json:"count_requested"`
	// Country where you would like to purchase phone numbers
	CountryISO string `json:"country_iso"`
	// ISO 8601 formatted date indicating when the ordering group was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Error reason if applicable
	ErrorReason string `json:"error_reason"`
	// Filter to exclude phone numbers that are currently on hold/reserved for your
	// account.
	ExcludeHeldNumbers bool `json:"exclude_held_numbers"`
	// Filter by area code
	NationalDestinationCode string `json:"national_destination_code"`
	// Array of orders created to fulfill the inexplicit order
	Orders []InexplicitNumberOrderListResponseDataOrderingGroupOrder `json:"orders"`
	// Number type
	PhoneNumberType string `json:"phone_number_type"`
	// Filter for phone numbers that contain the digits specified
	PhoneNumberContains string `json:"phone_number[contains]"`
	// Filter by the ending digits of the phone number
	PhoneNumberEndsWith string `json:"phone_number[ends_with]"`
	// Filter by the starting digits of the phone number
	PhoneNumberStartsWith string `json:"phone_number[starts_with]"`
	// Filter to exclude phone numbers that need additional time after to purchase to
	// activate. Only applicable for +1 toll_free numbers.
	Quickship bool `json:"quickship"`
	// Status of the ordering group
	//
	// Any of "pending", "processing", "failed", "success", "partial_success".
	Status string `json:"status"`
	// Ordering strategy used
	//
	// Any of "always", "never".
	Strategy string `json:"strategy"`
	// ISO 8601 formatted date indicating when the ordering group was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea      respjson.Field
		CountAllocated          respjson.Field
		CountRequested          respjson.Field
		CountryISO              respjson.Field
		CreatedAt               respjson.Field
		ErrorReason             respjson.Field
		ExcludeHeldNumbers      respjson.Field
		NationalDestinationCode respjson.Field
		Orders                  respjson.Field
		PhoneNumberType         respjson.Field
		PhoneNumberContains     respjson.Field
		PhoneNumberEndsWith     respjson.Field
		PhoneNumberStartsWith   respjson.Field
		Quickship               respjson.Field
		Status                  respjson.Field
		Strategy                respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderListResponseDataOrderingGroup) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderListResponseDataOrderingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderListResponseDataOrderingGroupOrder struct {
	// ID of the main number order
	NumberOrderID string `json:"number_order_id,required" format:"uuid"`
	// Array of sub number order IDs
	SubNumberOrderIDs []string `json:"sub_number_order_ids,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOrderID     respjson.Field
		SubNumberOrderIDs respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InexplicitNumberOrderListResponseDataOrderingGroupOrder) RawJSON() string { return r.JSON.raw }
func (r *InexplicitNumberOrderListResponseDataOrderingGroupOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderNewParams struct {
	// Group(s) of numbers to order. You can have multiple ordering_groups objects
	// added to a single request.
	OrderingGroups []InexplicitNumberOrderNewParamsOrderingGroup `json:"ordering_groups,omitzero,required"`
	// Billing group id to apply to phone numbers that are purchased
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// Connection id to apply to phone numbers that are purchased
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// Reference label for the customer
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Messaging profile id to apply to phone numbers that are purchased
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	paramObj
}

func (r InexplicitNumberOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InexplicitNumberOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InexplicitNumberOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CountRequested, CountryISO, PhoneNumberType are required.
type InexplicitNumberOrderNewParamsOrderingGroup struct {
	// Quantity of phone numbers to order
	CountRequested string `json:"count_requested,required"`
	// Country where you would like to purchase phone numbers. Allowable values: US, CA
	//
	// Any of "US", "CA".
	CountryISO string `json:"country_iso,omitzero,required"`
	// Number type (local, toll-free, etc.)
	PhoneNumberType string `json:"phone_number_type,required"`
	// Filter for phone numbers in a given state / province
	AdministrativeArea param.Opt[string] `json:"administrative_area,omitzero"`
	// Filter to exclude phone numbers that are currently on hold/reserved for your
	// account.
	ExcludeHeldNumbers param.Opt[bool] `json:"exclude_held_numbers,omitzero"`
	// Filter for phone numbers in a given city / region / rate center
	Locality param.Opt[string] `json:"locality,omitzero"`
	// Filter by area code
	NationalDestinationCode param.Opt[string] `json:"national_destination_code,omitzero"`
	// Filter to exclude phone numbers that need additional time after to purchase to
	// activate. Only applicable for +1 toll_free numbers.
	Quickship param.Opt[bool] `json:"quickship,omitzero"`
	// Filter for phone numbers that have the features to satisfy your use case (e.g.,
	// ["voice"])
	Features []string `json:"features,omitzero"`
	// Phone number search criteria
	PhoneNumber InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber `json:"phone_number,omitzero"`
	// Ordering strategy. Define what action should be taken if we don't have enough
	// phone numbers to fulfill your request. Allowable values are: always = proceed
	// with ordering phone numbers, regardless of current inventory levels; never = do
	// not place any orders unless there are enough phone numbers to satisfy the
	// request. If not specified, the always strategy will be enforced.
	//
	// Any of "always", "never".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r InexplicitNumberOrderNewParamsOrderingGroup) MarshalJSON() (data []byte, err error) {
	type shadow InexplicitNumberOrderNewParamsOrderingGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InexplicitNumberOrderNewParamsOrderingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InexplicitNumberOrderNewParamsOrderingGroup](
		"country_iso", "US", "CA",
	)
	apijson.RegisterFieldValidator[InexplicitNumberOrderNewParamsOrderingGroup](
		"strategy", "always", "never",
	)
}

// Phone number search criteria
type InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber struct {
	// Filter for phone numbers that contain the digits specified
	Contains param.Opt[string] `json:"contains,omitzero"`
	// Filter by the ending digits of the phone number
	EndsWith param.Opt[string] `json:"ends_with,omitzero"`
	// Filter by the starting digits of the phone number
	StartsWith param.Opt[string] `json:"starts_with,omitzero"`
	paramObj
}

func (r InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber) MarshalJSON() (data []byte, err error) {
	type shadow InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InexplicitNumberOrderListParams struct {
	// The page number to load
	PageNumber param.Opt[int64] `query:"page_number,omitzero" json:"-"`
	// The size of the page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InexplicitNumberOrderListParams]'s query parameters as
// `url.Values`.
func (r InexplicitNumberOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
