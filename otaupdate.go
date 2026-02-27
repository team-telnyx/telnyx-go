// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
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

// OTA updates operations
//
// OtaUpdateService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOtaUpdateService] method instead.
type OtaUpdateService struct {
	Options []option.RequestOption
}

// NewOtaUpdateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOtaUpdateService(opts ...option.RequestOption) (r OtaUpdateService) {
	r = OtaUpdateService{}
	r.Options = opts
	return
}

// This API returns the details of an Over the Air (OTA) update.
func (r *OtaUpdateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OtaUpdateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("ota_updates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List OTA updates
func (r *OtaUpdateService) List(ctx context.Context, query OtaUpdateListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[OtaUpdateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ota_updates"
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

// List OTA updates
func (r *OtaUpdateService) ListAutoPaging(ctx context.Context, query OtaUpdateListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[OtaUpdateListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type OtaUpdateGetResponse struct {
	// This object represents an Over the Air (OTA) update request. It allows tracking
	// the current status of a operation that apply settings in a particular SIM card.
	// <br/><br/>
	Data OtaUpdateGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtaUpdateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OtaUpdateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This object represents an Over the Air (OTA) update request. It allows tracking
// the current status of a operation that apply settings in a particular SIM card.
// <br/><br/>
type OtaUpdateGetResponseData struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// A JSON object representation of the operation. The information present here will
	// relate directly to the source of the OTA request.
	Settings OtaUpdateGetResponseDataSettings `json:"settings"`
	// The identification UUID of the related SIM card resource.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// Any of "in-progress", "completed", "failed".
	Status string `json:"status"`
	// Represents the type of the operation requested. This will relate directly to the
	// source of the request.
	//
	// Any of "sim_card_network_preferences".
	Type string `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		Settings    respjson.Field
		SimCardID   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtaUpdateGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *OtaUpdateGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON object representation of the operation. The information present here will
// relate directly to the source of the OTA request.
type OtaUpdateGetResponseDataSettings struct {
	// A list of mobile network operators and the priority that should be applied when
	// the SIM is connecting to the network.
	MobileNetworkOperatorsPreferences []OtaUpdateGetResponseDataSettingsMobileNetworkOperatorsPreference `json:"mobile_network_operators_preferences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MobileNetworkOperatorsPreferences respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtaUpdateGetResponseDataSettings) RawJSON() string { return r.JSON.raw }
func (r *OtaUpdateGetResponseDataSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtaUpdateGetResponseDataSettingsMobileNetworkOperatorsPreference struct {
	// The mobile network operator resource identification UUID.
	MobileNetworkOperatorID string `json:"mobile_network_operator_id" format:"uuid"`
	// The mobile network operator resource name.
	MobileNetworkOperatorName string `json:"mobile_network_operator_name"`
	// It determines what is the priority of a specific network operator that should be
	// assumed by a SIM card when connecting to a network. The highest priority is 0,
	// the second highest is 1 and so on.
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MobileNetworkOperatorID   respjson.Field
		MobileNetworkOperatorName respjson.Field
		Priority                  respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtaUpdateGetResponseDataSettingsMobileNetworkOperatorsPreference) RawJSON() string {
	return r.JSON.raw
}
func (r *OtaUpdateGetResponseDataSettingsMobileNetworkOperatorsPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This object represents an Over the Air (OTA) update request. It allows tracking
// the current status of a operation that apply settings in a particular SIM card.
// <br/><br/>
type OtaUpdateListResponse struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// The identification UUID of the related SIM card resource.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// Any of "in-progress", "completed", "failed".
	Status OtaUpdateListResponseStatus `json:"status"`
	// Represents the type of the operation requested. This will relate directly to the
	// source of the request.
	//
	// Any of "sim_card_network_preferences".
	Type OtaUpdateListResponseType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		SimCardID   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtaUpdateListResponse) RawJSON() string { return r.JSON.raw }
func (r *OtaUpdateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtaUpdateListResponseStatus string

const (
	OtaUpdateListResponseStatusInProgress OtaUpdateListResponseStatus = "in-progress"
	OtaUpdateListResponseStatusCompleted  OtaUpdateListResponseStatus = "completed"
	OtaUpdateListResponseStatusFailed     OtaUpdateListResponseStatus = "failed"
)

// Represents the type of the operation requested. This will relate directly to the
// source of the request.
type OtaUpdateListResponseType string

const (
	OtaUpdateListResponseTypeSimCardNetworkPreferences OtaUpdateListResponseType = "sim_card_network_preferences"
)

type OtaUpdateListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter for OTA updates (deepObject style). Originally:
	// filter[status], filter[sim_card_id], filter[type]
	Filter OtaUpdateListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OtaUpdateListParams]'s query parameters as `url.Values`.
func (r OtaUpdateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for OTA updates (deepObject style). Originally:
// filter[status], filter[sim_card_id], filter[type]
type OtaUpdateListParamsFilter struct {
	// The SIM card identification UUID.
	SimCardID param.Opt[string] `query:"sim_card_id,omitzero" json:"-"`
	// Filter by a specific status of the resource's lifecycle.
	//
	// Any of "in-progress", "completed", "failed".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by type.
	//
	// Any of "sim_card_network_preferences".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OtaUpdateListParamsFilter]'s query parameters as
// `url.Values`.
func (r OtaUpdateListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
