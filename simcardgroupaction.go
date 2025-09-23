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

// SimCardGroupActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardGroupActionService] method instead.
type SimCardGroupActionService struct {
	Options []option.RequestOption
}

// NewSimCardGroupActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimCardGroupActionService(opts ...option.RequestOption) (r SimCardGroupActionService) {
	r = SimCardGroupActionService{}
	r.Options = opts
	return
}

// This API allows fetching detailed information about a SIM card group action
// resource to make follow-ups in an existing asynchronous operation.
func (r *SimCardGroupActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGroupActionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_group_actions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API allows listing a paginated collection a SIM card group actions. It
// allows to explore a collection of existing asynchronous operation using specific
// filters.
func (r *SimCardGroupActionService) List(ctx context.Context, query SimCardGroupActionListParams, opts ...option.RequestOption) (res *SimCardGroupActionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sim_card_group_actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This action will asynchronously remove an existing Private Wireless Gateway
// definition from a SIM card group. Completing this operation defines that all SIM
// cards in the SIM card group will get their traffic handled by Telnyx's default
// mobile network configuration.
func (r *SimCardGroupActionService) RemovePrivateWirelessGateway(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGroupActionRemovePrivateWirelessGatewayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s/actions/remove_private_wireless_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This action will asynchronously remove an existing Wireless Blocklist to all the
// SIMs in the SIM card group.
func (r *SimCardGroupActionService) RemoveWirelessBlocklist(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardGroupActionRemoveWirelessBlocklistResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s/actions/remove_wireless_blocklist", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This action will asynchronously assign a provisioned Private Wireless Gateway to
// the SIM card group. Completing this operation defines that all SIM cards in the
// SIM card group will get their traffic controlled by the associated Private
// Wireless Gateway. This operation will also imply that new SIM cards assigned to
// a group will inherit its network definitions. If it's moved to a different group
// that doesn't have a Private Wireless Gateway, it'll use Telnyx's default mobile
// network configuration.
func (r *SimCardGroupActionService) SetPrivateWirelessGateway(ctx context.Context, id string, body SimCardGroupActionSetPrivateWirelessGatewayParams, opts ...option.RequestOption) (res *SimCardGroupActionSetPrivateWirelessGatewayResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s/actions/set_private_wireless_gateway", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This action will asynchronously assign a Wireless Blocklist to all the SIMs in
// the SIM card group.
func (r *SimCardGroupActionService) SetWirelessBlocklist(ctx context.Context, id string, body SimCardGroupActionSetWirelessBlocklistParams, opts ...option.RequestOption) (res *SimCardGroupActionSetWirelessBlocklistResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_groups/%s/actions/set_wireless_blocklist", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This object represents a SIM card group action request. It allows tracking the
// current status of an operation that impacts the SIM card group and SIM card in
// it.
type SimCardGroupAction struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// A JSON object representation of the action params.
	Settings SimCardGroupActionSettings `json:"settings"`
	// The SIM card group identification.
	SimCardGroupID string `json:"sim_card_group_id" format:"uuid"`
	// Any of "in-progress", "completed", "failed".
	Status SimCardGroupActionStatus `json:"status"`
	// Represents the type of the operation requested.
	//
	// Any of "set_private_wireless_gateway", "remove_private_wireless_gateway",
	// "set_wireless_blocklist", "remove_wireless_blocklist".
	Type SimCardGroupActionType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		RecordType     respjson.Field
		Settings       respjson.Field
		SimCardGroupID respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupAction) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON object representation of the action params.
type SimCardGroupActionSettings struct {
	// The identification of the related Private Wireless Gateway resource.
	PrivateWirelessGatewayID string `json:"private_wireless_gateway_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrivateWirelessGatewayID respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionSettings) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionStatus string

const (
	SimCardGroupActionStatusInProgress SimCardGroupActionStatus = "in-progress"
	SimCardGroupActionStatusCompleted  SimCardGroupActionStatus = "completed"
	SimCardGroupActionStatusFailed     SimCardGroupActionStatus = "failed"
)

// Represents the type of the operation requested.
type SimCardGroupActionType string

const (
	SimCardGroupActionTypeSetPrivateWirelessGateway    SimCardGroupActionType = "set_private_wireless_gateway"
	SimCardGroupActionTypeRemovePrivateWirelessGateway SimCardGroupActionType = "remove_private_wireless_gateway"
	SimCardGroupActionTypeSetWirelessBlocklist         SimCardGroupActionType = "set_wireless_blocklist"
	SimCardGroupActionTypeRemoveWirelessBlocklist      SimCardGroupActionType = "remove_wireless_blocklist"
)

type SimCardGroupActionGetResponse struct {
	// This object represents a SIM card group action request. It allows tracking the
	// current status of an operation that impacts the SIM card group and SIM card in
	// it.
	Data SimCardGroupAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionListResponse struct {
	Data []SimCardGroupAction `json:"data"`
	Meta PaginationMeta       `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionRemovePrivateWirelessGatewayResponse struct {
	// This object represents a SIM card group action request. It allows tracking the
	// current status of an operation that impacts the SIM card group and SIM card in
	// it.
	Data SimCardGroupAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionRemovePrivateWirelessGatewayResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionRemovePrivateWirelessGatewayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionRemoveWirelessBlocklistResponse struct {
	// This object represents a SIM card group action request. It allows tracking the
	// current status of an operation that impacts the SIM card group and SIM card in
	// it.
	Data SimCardGroupAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionRemoveWirelessBlocklistResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionRemoveWirelessBlocklistResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionSetPrivateWirelessGatewayResponse struct {
	// This object represents a SIM card group action request. It allows tracking the
	// current status of an operation that impacts the SIM card group and SIM card in
	// it.
	Data SimCardGroupAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionSetPrivateWirelessGatewayResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionSetPrivateWirelessGatewayResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionSetWirelessBlocklistResponse struct {
	// This object represents a SIM card group action request. It allows tracking the
	// current status of an operation that impacts the SIM card group and SIM card in
	// it.
	Data SimCardGroupAction `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardGroupActionSetWirelessBlocklistResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardGroupActionSetWirelessBlocklistResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionListParams struct {
	// A valid SIM card group ID.
	FilterSimCardGroupID param.Opt[string] `query:"filter[sim_card_group_id],omitzero" format:"uuid" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by a specific status of the resource's lifecycle.
	//
	// Any of "in-progress", "completed", "failed".
	FilterStatus SimCardGroupActionListParamsFilterStatus `query:"filter[status],omitzero" json:"-"`
	// Filter by action type.
	//
	// Any of "set_private_wireless_gateway", "remove_private_wireless_gateway",
	// "set_wireless_blocklist", "remove_wireless_blocklist".
	FilterType SimCardGroupActionListParamsFilterType `query:"filter[type],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardGroupActionListParams]'s query parameters as
// `url.Values`.
func (r SimCardGroupActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by a specific status of the resource's lifecycle.
type SimCardGroupActionListParamsFilterStatus string

const (
	SimCardGroupActionListParamsFilterStatusInProgress SimCardGroupActionListParamsFilterStatus = "in-progress"
	SimCardGroupActionListParamsFilterStatusCompleted  SimCardGroupActionListParamsFilterStatus = "completed"
	SimCardGroupActionListParamsFilterStatusFailed     SimCardGroupActionListParamsFilterStatus = "failed"
)

// Filter by action type.
type SimCardGroupActionListParamsFilterType string

const (
	SimCardGroupActionListParamsFilterTypeSetPrivateWirelessGateway    SimCardGroupActionListParamsFilterType = "set_private_wireless_gateway"
	SimCardGroupActionListParamsFilterTypeRemovePrivateWirelessGateway SimCardGroupActionListParamsFilterType = "remove_private_wireless_gateway"
	SimCardGroupActionListParamsFilterTypeSetWirelessBlocklist         SimCardGroupActionListParamsFilterType = "set_wireless_blocklist"
	SimCardGroupActionListParamsFilterTypeRemoveWirelessBlocklist      SimCardGroupActionListParamsFilterType = "remove_wireless_blocklist"
)

type SimCardGroupActionSetPrivateWirelessGatewayParams struct {
	// The identification of the related Private Wireless Gateway resource.
	PrivateWirelessGatewayID string `json:"private_wireless_gateway_id,required" format:"uuid"`
	paramObj
}

func (r SimCardGroupActionSetPrivateWirelessGatewayParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupActionSetPrivateWirelessGatewayParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupActionSetPrivateWirelessGatewayParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardGroupActionSetWirelessBlocklistParams struct {
	// The identification of the related Wireless Blocklist resource.
	WirelessBlocklistID string `json:"wireless_blocklist_id,required" format:"uuid"`
	paramObj
}

func (r SimCardGroupActionSetWirelessBlocklistParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardGroupActionSetWirelessBlocklistParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardGroupActionSetWirelessBlocklistParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
