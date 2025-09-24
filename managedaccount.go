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

// ManagedAccountService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedAccountService] method instead.
type ManagedAccountService struct {
	Options []option.RequestOption
	Actions ManagedAccountActionService
}

// NewManagedAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewManagedAccountService(opts ...option.RequestOption) (r ManagedAccountService) {
	r = ManagedAccountService{}
	r.Options = opts
	r.Actions = NewManagedAccountActionService(opts...)
	return
}

// Create a new managed account owned by the authenticated user. You need to be
// explictly approved by Telnyx in order to become a manager account.
func (r *ManagedAccountService) New(ctx context.Context, body ManagedAccountNewParams, opts ...option.RequestOption) (res *ManagedAccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "managed_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of a single managed account.
func (r *ManagedAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ManagedAccountGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("managed_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a single managed account.
func (r *ManagedAccountService) Update(ctx context.Context, id string, body ManagedAccountUpdateParams, opts ...option.RequestOption) (res *ManagedAccountUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("managed_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists the accounts managed by the current user. Users need to be explictly
// approved by Telnyx in order to become manager accounts.
func (r *ManagedAccountService) List(ctx context.Context, query ManagedAccountListParams, opts ...option.RequestOption) (res *ManagedAccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "managed_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Display information about allocatable global outbound channels for the current
// user. Only usable by account managers.
func (r *ManagedAccountService) GetAllocatableGlobalOutboundChannels(ctx context.Context, opts ...option.RequestOption) (res *ManagedAccountGetAllocatableGlobalOutboundChannelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "managed_accounts/allocatable_global_outbound_channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the amount of allocatable global outbound channels allocated to a
// specific managed account.
func (r *ManagedAccountService) UpdateGlobalChannelLimit(ctx context.Context, id string, body ManagedAccountUpdateGlobalChannelLimitParams, opts ...option.RequestOption) (res *ManagedAccountUpdateGlobalChannelLimitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("managed_accounts/%s/update_global_channel_limit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ManagedAccount struct {
	// Uniquely identifies the managed account.
	ID string `json:"id,required" format:"uuid"`
	// The managed account's V2 API access key
	APIKey string `json:"api_key,required"`
	// The managed account's V1 API token
	APIToken string `json:"api_token,required"`
	// The manager account's email, which serves as the V1 API user identifier
	APIUser string `json:"api_user,required"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,required"`
	// The managed account's email.
	Email string `json:"email,required" format:"email"`
	// The ID of the manager account associated with the managed account.
	ManagerAccountID string `json:"manager_account_id,required"`
	// Identifies the type of the resource.
	//
	// Any of "managed_account".
	RecordType ManagedAccountRecordType `json:"record_type,required"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string                `json:"updated_at,required"`
	Balance   ManagedAccountBalance `json:"balance"`
	// Boolean value that indicates if the managed account is able to have custom
	// pricing set for it or not. If false, uses the pricing of the manager account.
	// Defaults to false. There may be time lag between when the value is changed and
	// pricing changes take effect.
	ManagedAccountAllowCustomPricing bool `json:"managed_account_allow_custom_pricing"`
	// The organization the managed account is associated with.
	OrganizationName string `json:"organization_name"`
	// Boolean value that indicates if the billing information and charges to the
	// managed account "roll up" to the manager account. If true, the managed account
	// will not have its own balance and will use the shared balance with the manager
	// account. This value cannot be changed after account creation without going
	// through Telnyx support as changes require manual updates to the account ledger.
	// Defaults to false.
	RollupBilling bool `json:"rollup_billing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		APIKey                           respjson.Field
		APIToken                         respjson.Field
		APIUser                          respjson.Field
		CreatedAt                        respjson.Field
		Email                            respjson.Field
		ManagerAccountID                 respjson.Field
		RecordType                       respjson.Field
		UpdatedAt                        respjson.Field
		Balance                          respjson.Field
		ManagedAccountAllowCustomPricing respjson.Field
		OrganizationName                 respjson.Field
		RollupBilling                    respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccount) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ManagedAccountRecordType string

const (
	ManagedAccountRecordTypeManagedAccount ManagedAccountRecordType = "managed_account"
)

type ManagedAccountBalance struct {
	// Available amount to spend (balance + credit limit)
	AvailableCredit string `json:"available_credit" format:"decimal"`
	// The account's current balance.
	Balance string `json:"balance" format:"decimal"`
	// The account's credit limit.
	CreditLimit string `json:"credit_limit" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// Identifies the type of the resource.
	//
	// Any of "balance".
	RecordType ManagedAccountBalanceRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableCredit respjson.Field
		Balance         respjson.Field
		CreditLimit     respjson.Field
		Currency        respjson.Field
		RecordType      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountBalance) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ManagedAccountBalanceRecordType string

const (
	ManagedAccountBalanceRecordTypeBalance ManagedAccountBalanceRecordType = "balance"
)

type ManagedAccountNewResponse struct {
	Data ManagedAccount `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountGetResponse struct {
	Data ManagedAccount `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountUpdateResponse struct {
	Data ManagedAccount `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountListResponse struct {
	Data []ManagedAccountListResponseData `json:"data"`
	Meta PaginationMeta                   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountListResponseData struct {
	// Uniquely identifies the managed account.
	ID string `json:"id,required" format:"uuid"`
	// The manager account's email, which serves as the V1 API user identifier
	APIUser string `json:"api_user,required"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,required"`
	// The managed account's email.
	Email string `json:"email,required" format:"email"`
	// The ID of the manager account associated with the managed account.
	ManagerAccountID string `json:"manager_account_id,required"`
	// Identifies the type of the resource.
	//
	// Any of "managed_account".
	RecordType string `json:"record_type,required"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at,required"`
	// Boolean value that indicates if the managed account is able to have custom
	// pricing set for it or not. If false, uses the pricing of the manager account.
	// Defaults to false. There may be time lag between when the value is changed and
	// pricing changes take effect.
	ManagedAccountAllowCustomPricing bool `json:"managed_account_allow_custom_pricing"`
	// The organization the managed account is associated with.
	OrganizationName string `json:"organization_name"`
	// Boolean value that indicates if the billing information and charges to the
	// managed account "roll up" to the manager account. If true, the managed account
	// will not have its own balance and will use the shared balance with the manager
	// account. This value cannot be changed after account creation without going
	// through Telnyx support as changes require manual updates to the account ledger.
	// Defaults to false.
	RollupBilling bool `json:"rollup_billing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		APIUser                          respjson.Field
		CreatedAt                        respjson.Field
		Email                            respjson.Field
		ManagerAccountID                 respjson.Field
		RecordType                       respjson.Field
		UpdatedAt                        respjson.Field
		ManagedAccountAllowCustomPricing respjson.Field
		OrganizationName                 respjson.Field
		RollupBilling                    respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountListResponseData) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountGetAllocatableGlobalOutboundChannelsResponse struct {
	Data ManagedAccountGetAllocatableGlobalOutboundChannelsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountGetAllocatableGlobalOutboundChannelsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ManagedAccountGetAllocatableGlobalOutboundChannelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountGetAllocatableGlobalOutboundChannelsResponseData struct {
	// The total amount of allocatable global outbound channels available to the
	// authenticated manager. Will be 0 if the feature is not enabled for their
	// account.
	AllocatableGlobalOutboundChannels int64 `json:"allocatable_global_outbound_channels"`
	// Boolean value that indicates if the managed account is able to have custom
	// pricing set for it or not. If false, uses the pricing of the manager account.
	// Defaults to false. This value may be changed, but there may be time lag between
	// when the value is changed and pricing changes take effect.
	ManagedAccountAllowCustomPricing bool `json:"managed_account_allow_custom_pricing"`
	// The type of the data contained in this record.
	RecordType string `json:"record_type"`
	// The total number of allocatable global outbound channels currently allocated
	// across all managed accounts for the authenticated user. This includes any amount
	// of channels allocated by default at managed account creation time. Will be 0 if
	// the feature is not enabled for their account.
	TotalGlobalChannelsAllocated int64 `json:"total_global_channels_allocated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllocatableGlobalOutboundChannels respjson.Field
		ManagedAccountAllowCustomPricing  respjson.Field
		RecordType                        respjson.Field
		TotalGlobalChannelsAllocated      respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountGetAllocatableGlobalOutboundChannelsResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *ManagedAccountGetAllocatableGlobalOutboundChannelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountUpdateGlobalChannelLimitResponse struct {
	Data ManagedAccountUpdateGlobalChannelLimitResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountUpdateGlobalChannelLimitResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountUpdateGlobalChannelLimitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountUpdateGlobalChannelLimitResponseData struct {
	// The user ID of the managed account.
	ID string `json:"id"`
	// Integer value that indicates the number of allocatable global outbound channels
	// that are allocated to the managed account. If the value is 0 then the account
	// will have no usable channels and will not be able to perform outbound calling.
	ChannelLimit int64 `json:"channel_limit"`
	// The email of the managed account.
	Email string `json:"email"`
	// The user ID of the manager of the account.
	ManagerAccountID string `json:"manager_account_id"`
	// The name of the type of data in the response.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ChannelLimit     respjson.Field
		Email            respjson.Field
		ManagerAccountID respjson.Field
		RecordType       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountUpdateGlobalChannelLimitResponseData) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountUpdateGlobalChannelLimitResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountNewParams struct {
	// The name of the business for which the new managed account is being created,
	// that will be used as the managed accounts's organization's name.
	BusinessName string `json:"business_name,required"`
	// The email address for the managed account. If not provided, the email address
	// will be generated based on the email address of the manager account.
	Email param.Opt[string] `json:"email,omitzero"`
	// Boolean value that indicates if the managed account is able to have custom
	// pricing set for it or not. If false, uses the pricing of the manager account.
	// Defaults to false. This value may be changed after creation, but there may be
	// time lag between when the value is changed and pricing changes take effect.
	ManagedAccountAllowCustomPricing param.Opt[bool] `json:"managed_account_allow_custom_pricing,omitzero"`
	// Password for the managed account. If a password is not supplied, the account
	// will not be able to be signed into directly. (A password reset may still be
	// performed later to enable sign-in via password.)
	Password param.Opt[string] `json:"password,omitzero"`
	// Boolean value that indicates if the billing information and charges to the
	// managed account "roll up" to the manager account. If true, the managed account
	// will not have its own balance and will use the shared balance with the manager
	// account. This value cannot be changed after account creation without going
	// through Telnyx support as changes require manual updates to the account ledger.
	// Defaults to false.
	RollupBilling param.Opt[bool] `json:"rollup_billing,omitzero"`
	paramObj
}

func (r ManagedAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ManagedAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManagedAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountUpdateParams struct {
	// Boolean value that indicates if the managed account is able to have custom
	// pricing set for it or not. If false, uses the pricing of the manager account.
	// Defaults to false. This value may be changed, but there may be time lag between
	// when the value is changed and pricing changes take effect.
	ManagedAccountAllowCustomPricing param.Opt[bool] `json:"managed_account_allow_custom_pricing,omitzero"`
	paramObj
}

func (r ManagedAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ManagedAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManagedAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountListParams struct {
	// Specifies if cancelled accounts should be included in the results.
	IncludeCancelledAccounts param.Opt[bool] `query:"include_cancelled_accounts,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[email][contains], filter[email][eq], filter[organization_name][contains],
	// filter[organization_name][eq]
	Filter ManagedAccountListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page ManagedAccountListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code> -</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>email</code>: sorts the result by the
	//	  <code>email</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-email</code>: sorts the result by the
	//	  <code>email</code> field in descending order.
	//	</li>
	//
	// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "created_at", "email".
	Sort ManagedAccountListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManagedAccountListParams]'s query parameters as
// `url.Values`.
func (r ManagedAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[email][contains], filter[email][eq], filter[organization_name][contains],
// filter[organization_name][eq]
type ManagedAccountListParamsFilter struct {
	Email            ManagedAccountListParamsFilterEmail            `query:"email,omitzero" json:"-"`
	OrganizationName ManagedAccountListParamsFilterOrganizationName `query:"organization_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManagedAccountListParamsFilter]'s query parameters as
// `url.Values`.
func (r ManagedAccountListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManagedAccountListParamsFilterEmail struct {
	// If present, email containing the given value will be returned. Matching is not
	// case-sensitive. Requires at least three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// If present, only returns results with the <code>email</code> matching exactly
	// the value given.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManagedAccountListParamsFilterEmail]'s query parameters as
// `url.Values`.
func (r ManagedAccountListParamsFilterEmail) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ManagedAccountListParamsFilterOrganizationName struct {
	// If present, only returns results with the <code>organization_name</code>
	// containing the given value. Matching is not case-sensitive. Requires at least
	// three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// If present, only returns results with the <code>organization_name</code>
	// matching exactly the value given.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManagedAccountListParamsFilterOrganizationName]'s query
// parameters as `url.Values`.
func (r ManagedAccountListParamsFilterOrganizationName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type ManagedAccountListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ManagedAccountListParamsPage]'s query parameters as
// `url.Values`.
func (r ManagedAccountListParamsPage) URLQuery() (v url.Values, err error) {
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
//	  <code>email</code>: sorts the result by the
//	  <code>email</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-email</code>: sorts the result by the
//	  <code>email</code> field in descending order.
//	</li>
//
// </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order.
type ManagedAccountListParamsSort string

const (
	ManagedAccountListParamsSortCreatedAt ManagedAccountListParamsSort = "created_at"
	ManagedAccountListParamsSortEmail     ManagedAccountListParamsSort = "email"
)

type ManagedAccountUpdateGlobalChannelLimitParams struct {
	// Integer value that indicates the number of allocatable global outbound channels
	// that should be allocated to the managed account. Must be 0 or more. If the value
	// is 0 then the account will have no usable channels and will not be able to
	// perform outbound calling.
	ChannelLimit param.Opt[int64] `json:"channel_limit,omitzero"`
	paramObj
}

func (r ManagedAccountUpdateGlobalChannelLimitParams) MarshalJSON() (data []byte, err error) {
	type shadow ManagedAccountUpdateGlobalChannelLimitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManagedAccountUpdateGlobalChannelLimitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
