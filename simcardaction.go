// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// SimCardActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardActionService] method instead.
type SimCardActionService struct {
	Options []option.RequestOption
}

// NewSimCardActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimCardActionService(opts ...option.RequestOption) (r SimCardActionService) {
	r = SimCardActionService{}
	r.Options = opts
	return
}

// This API fetches detailed information about a SIM card action to follow-up on an
// existing asynchronous operation.
func (r *SimCardActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardActionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_actions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API lists a paginated collection of SIM card actions. It enables exploring
// a collection of existing asynchronous operations using specific filters.
func (r *SimCardActionService) List(ctx context.Context, query SimCardActionListParams, opts ...option.RequestOption) (res *SimCardActionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_card_actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This API triggers an asynchronous operation to set a public IP for each of the
// specified SIM cards.<br/> For each SIM Card a SIM Card Action will be generated.
// The status of the SIM Card Action can be followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API.
func (r *SimCardActionService) BulkSetPublicIPs(ctx context.Context, body SimCardActionBulkSetPublicIPsParams, opts ...option.RequestOption) (res *SimCardActionBulkSetPublicIPsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_cards/actions/bulk_set_public_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API disables a SIM card, disconnecting it from the network and making it
// impossible to consume data.<br/> The API will trigger an asynchronous operation
// called a SIM Card Action. Transitioning to the disabled state may take a period
// of time. The status of the SIM Card Action can be followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API.
func (r *SimCardActionService) Disable(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardActionDisableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/actions/disable", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This API enables a SIM card, connecting it to the network and making it possible
// to consume data.<br/> To enable a SIM card, it must be associated with a SIM
// card group.<br/> The API will trigger an asynchronous operation called a SIM
// Card Action. Transitioning to the enabled state may take a period of time. The
// status of the SIM Card Action can be followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API.
func (r *SimCardActionService) Enable(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardActionEnableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/actions/enable", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This API removes an existing public IP from a SIM card. <br/><br/> The API will
// trigger an asynchronous operation called a SIM Card Action. The status of the
// SIM Card Action can be followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API.
func (r *SimCardActionService) RemovePublicIP(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardActionRemovePublicIPResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/actions/remove_public_ip", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This API makes a SIM card reachable on the public internet by mapping a random
// public IP to the SIM card. <br/><br/> The API will trigger an asynchronous
// operation called a SIM Card Action. The status of the SIM Card Action can be
// followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API. <br/><br/> Setting a Public IP to a SIM Card incurs a charge and will only
// succeed if the account has sufficient funds.
func (r *SimCardActionService) SetPublicIP(ctx context.Context, id string, body SimCardActionSetPublicIPParams, opts ...option.RequestOption) (res *SimCardActionSetPublicIPResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/actions/set_public_ip", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The SIM card will be able to connect to the network once the process to set it
// to standby has been completed, thus making it possible to consume data.<br/> To
// set a SIM card to standby, it must be associated with SIM card group.<br/> The
// API will trigger an asynchronous operation called a SIM Card Action.
// Transitioning to the standby state may take a period of time. The status of the
// SIM Card Action can be followed through the
// [List SIM Card Action](https://developersdev.telnyx.com/docs/api/v2/wireless/SIM-Card-Actions#ListSIMCardActions)
// API.
func (r *SimCardActionService) SetStandby(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardActionSetStandbyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_cards/%s/actions/set_standby", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// It validates whether SIM card registration codes are valid or not.
func (r *SimCardActionService) ValidateRegistrationCodes(ctx context.Context, body SimCardActionValidateRegistrationCodesParams, opts ...option.RequestOption) (res *SimCardActionValidateRegistrationCodesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_cards/actions/validate_registration_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This object represents a SIM card action. It allows tracking the current status
// of an operation that impacts the SIM card.
type SimCardAction struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The operation type. It can be one of the following: <br/>
	//
	// <ul>
	//
	//	<li><code>enable</code> - move the SIM card to the <code>enabled</code> status</li>
	//	<li><code>enable_standby_sim_card</code> - move a SIM card previously on the <code>standby</code> status to the <code>enabled</code> status after it consumes data.</li>
	//	<li><code>disable</code> - move the SIM card to the <code>disabled</code> status</li>
	//	<li><code>set_standby</code> - move the SIM card to the <code>standby</code> status</li>
	//	</ul>
	//
	// Any of "enable", "enable_standby_sim_card", "disable", "set_standby".
	ActionType SimCardActionActionType `json:"action_type"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// A JSON object representation of the action params.
	Settings map[string]any `json:"settings,nullable"`
	// The related SIM card identifier.
	SimCardID string              `json:"sim_card_id" format:"uuid"`
	Status    SimCardActionStatus `json:"status"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ActionType  respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		Settings    respjson.Field
		SimCardID   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardAction) RawJSON() string { return r.JSON.raw }
func (r *SimCardAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The operation type. It can be one of the following: <br/>
//
// <ul>
//
//	<li><code>enable</code> - move the SIM card to the <code>enabled</code> status</li>
//	<li><code>enable_standby_sim_card</code> - move a SIM card previously on the <code>standby</code> status to the <code>enabled</code> status after it consumes data.</li>
//	<li><code>disable</code> - move the SIM card to the <code>disabled</code> status</li>
//	<li><code>set_standby</code> - move the SIM card to the <code>standby</code> status</li>
//	</ul>
type SimCardActionActionType string

const (
	SimCardActionActionTypeEnable               SimCardActionActionType = "enable"
	SimCardActionActionTypeEnableStandbySimCard SimCardActionActionType = "enable_standby_sim_card"
	SimCardActionActionTypeDisable              SimCardActionActionType = "disable"
	SimCardActionActionTypeSetStandby           SimCardActionActionType = "set_standby"
)

type SimCardActionStatus struct {
	// It describes why the SIM card action is in the current status. This will be
	// <code>null</code> for self-explanatory statuses, such as
	// <code>in-progress</code> and <code>completed</code> but will include further
	// information on statuses like <code>interrupted</code> and <code>failed</code>.
	Reason string `json:"reason"`
	// The current status of the SIM card action.
	//
	// Any of "in-progress", "completed", "failed", "interrupted".
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionStatus) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionGetResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionListResponse struct {
	Data []SimCardAction `json:"data"`
	Meta PaginationMeta  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionBulkSetPublicIPsResponse struct {
	// This object represents a bulk SIM card action. It groups SIM card actions
	// created through a bulk endpoint under a single resource for further lookup.
	Data SimCardActionBulkSetPublicIPsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionBulkSetPublicIPsResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionBulkSetPublicIPsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This object represents a bulk SIM card action. It groups SIM card actions
// created through a bulk endpoint under a single resource for further lookup.
type SimCardActionBulkSetPublicIPsResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The operation type. It can be one of the following: <br/>
	//
	// <ul>
	// <li><code>bulk_set_public_ips</code> - set a public IP for each specified SIM card.</li>
	// </ul>
	//
	// Any of "bulk_set_public_ips".
	ActionType string `json:"action_type"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// A JSON object representation of the bulk action payload.
	Settings map[string]any `json:"settings"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ActionType  respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		Settings    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionBulkSetPublicIPsResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionBulkSetPublicIPsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionDisableResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionDisableResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionDisableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionEnableResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionEnableResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionEnableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionRemovePublicIPResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionRemovePublicIPResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionRemovePublicIPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionSetPublicIPResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionSetPublicIPResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionSetPublicIPResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionSetStandbyResponse struct {
	// This object represents a SIM card action. It allows tracking the current status
	// of an operation that impacts the SIM card.
	Data SimCardAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionSetStandbyResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionSetStandbyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionValidateRegistrationCodesResponse struct {
	Data []SimCardActionValidateRegistrationCodesResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionValidateRegistrationCodesResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionValidateRegistrationCodesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionValidateRegistrationCodesResponseData struct {
	// The validation message
	InvalidDetail string `json:"invalid_detail,nullable"`
	RecordType    string `json:"record_type"`
	// The 10-digit SIM card registration code
	RegistrationCode string `json:"registration_code"`
	// The attribute that denotes whether the code is valid or not
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvalidDetail    respjson.Field
		RecordType       respjson.Field
		RegistrationCode respjson.Field
		Valid            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardActionValidateRegistrationCodesResponseData) RawJSON() string { return r.JSON.raw }
func (r *SimCardActionValidateRegistrationCodesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionListParams struct {
	// Consolidated filter parameter for SIM card actions (deepObject style).
	// Originally: filter[sim_card_id], filter[status],
	// filter[bulk_sim_card_action_id], filter[action_type]
	Filter SimCardActionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated pagination parameter (deepObject style). Originally: page[number],
	// page[size]
	Page SimCardActionListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardActionListParams]'s query parameters as
// `url.Values`.
func (r SimCardActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for SIM card actions (deepObject style).
// Originally: filter[sim_card_id], filter[status],
// filter[bulk_sim_card_action_id], filter[action_type]
type SimCardActionListParamsFilter struct {
	// Filter by a bulk SIM card action ID.
	BulkSimCardActionID param.Opt[string] `query:"bulk_sim_card_action_id,omitzero" format:"uuid" json:"-"`
	// A valid SIM card ID.
	SimCardID param.Opt[string] `query:"sim_card_id,omitzero" format:"uuid" json:"-"`
	// Filter by action type.
	//
	// Any of "enable", "enable_standby_sim_card", "disable", "set_standby",
	// "remove_public_ip", "set_public_ip".
	ActionType string `query:"action_type,omitzero" json:"-"`
	// Filter by a specific status of the resource's lifecycle.
	//
	// Any of "in-progress", "completed", "failed".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardActionListParamsFilter]'s query parameters as
// `url.Values`.
func (r SimCardActionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated pagination parameter (deepObject style). Originally: page[number],
// page[size]
type SimCardActionListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardActionListParamsPage]'s query parameters as
// `url.Values`.
func (r SimCardActionListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SimCardActionBulkSetPublicIPsParams struct {
	SimCardIDs []string `json:"sim_card_ids,omitzero,required" format:"uuid"`
	paramObj
}

func (r SimCardActionBulkSetPublicIPsParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardActionBulkSetPublicIPsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardActionBulkSetPublicIPsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardActionSetPublicIPParams struct {
	// The code of the region where the public IP should be assigned. A list of
	// available regions can be found at the regions endpoint
	RegionCode param.Opt[string] `query:"region_code,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardActionSetPublicIPParams]'s query parameters as
// `url.Values`.
func (r SimCardActionSetPublicIPParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SimCardActionValidateRegistrationCodesParams struct {
	RegistrationCodes []string `json:"registration_codes,omitzero"`
	paramObj
}

func (r SimCardActionValidateRegistrationCodesParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardActionValidateRegistrationCodesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardActionValidateRegistrationCodesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
