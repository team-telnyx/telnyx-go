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

// BulkSimCardActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkSimCardActionService] method instead.
type BulkSimCardActionService struct {
	Options []option.RequestOption
}

// NewBulkSimCardActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBulkSimCardActionService(opts ...option.RequestOption) (r BulkSimCardActionService) {
	r = BulkSimCardActionService{}
	r.Options = opts
	return
}

// This API fetches information about a bulk SIM card action. A bulk SIM card
// action contains details about a collection of individual SIM card actions.
func (r *BulkSimCardActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BulkSimCardActionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("bulk_sim_card_actions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API lists a paginated collection of bulk SIM card actions. A bulk SIM card
// action contains details about a collection of individual SIM card actions.
func (r *BulkSimCardActionService) List(ctx context.Context, query BulkSimCardActionListParams, opts ...option.RequestOption) (res *BulkSimCardActionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "bulk_sim_card_actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BulkSimCardActionGetResponse struct {
	Data BulkSimCardActionGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionGetResponseData struct {
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
	Settings              map[string]any                                          `json:"settings"`
	SimCardActionsSummary []BulkSimCardActionGetResponseDataSimCardActionsSummary `json:"sim_card_actions_summary"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ActionType            respjson.Field
		CreatedAt             respjson.Field
		RecordType            respjson.Field
		Settings              respjson.Field
		SimCardActionsSummary respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionGetResponseDataSimCardActionsSummary struct {
	Count int64 `json:"count"`
	// Any of "in-progress", "completed", "failed", "interrupted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionGetResponseDataSimCardActionsSummary) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionGetResponseDataSimCardActionsSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionListResponse struct {
	Data []BulkSimCardActionListResponseData `json:"data"`
	Meta PaginationMeta                      `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionListResponseData struct {
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
	Settings              map[string]any                                           `json:"settings"`
	SimCardActionsSummary []BulkSimCardActionListResponseDataSimCardActionsSummary `json:"sim_card_actions_summary"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ActionType            respjson.Field
		CreatedAt             respjson.Field
		RecordType            respjson.Field
		Settings              respjson.Field
		SimCardActionsSummary respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionListResponseDataSimCardActionsSummary struct {
	Count int64 `json:"count"`
	// Any of "in-progress", "completed", "failed", "interrupted".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkSimCardActionListResponseDataSimCardActionsSummary) RawJSON() string { return r.JSON.raw }
func (r *BulkSimCardActionListResponseDataSimCardActionsSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkSimCardActionListParams struct {
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by action type.
	//
	// Any of "bulk_set_public_ips".
	FilterActionType BulkSimCardActionListParamsFilterActionType `query:"filter[action_type],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BulkSimCardActionListParams]'s query parameters as
// `url.Values`.
func (r BulkSimCardActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by action type.
type BulkSimCardActionListParamsFilterActionType string

const (
	BulkSimCardActionListParamsFilterActionTypeBulkSetPublicIPs BulkSimCardActionListParamsFilterActionType = "bulk_set_public_ips"
)
