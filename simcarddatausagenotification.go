// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// SimCardDataUsageNotificationService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimCardDataUsageNotificationService] method instead.
type SimCardDataUsageNotificationService struct {
	Options []option.RequestOption
}

// NewSimCardDataUsageNotificationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSimCardDataUsageNotificationService(opts ...option.RequestOption) (r SimCardDataUsageNotificationService) {
	r = SimCardDataUsageNotificationService{}
	r.Options = opts
	return
}

// Creates a new SIM card data usage notification.
func (r *SimCardDataUsageNotificationService) New(ctx context.Context, body SimCardDataUsageNotificationNewParams, opts ...option.RequestOption) (res *SimCardDataUsageNotificationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sim_card_data_usage_notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single SIM Card Data Usage Notification.
func (r *SimCardDataUsageNotificationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardDataUsageNotificationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_data_usage_notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates information for a SIM Card Data Usage Notification.
func (r *SimCardDataUsageNotificationService) Update(ctx context.Context, id string, body SimCardDataUsageNotificationUpdateParams, opts ...option.RequestOption) (res *SimCardDataUsageNotificationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_data_usage_notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists a paginated collection of SIM card data usage notifications. It enables
// exploring the collection using specific filters.
func (r *SimCardDataUsageNotificationService) List(ctx context.Context, query SimCardDataUsageNotificationListParams, opts ...option.RequestOption) (res *SimCardDataUsageNotificationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "sim_card_data_usage_notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete the SIM Card Data Usage Notification.
func (r *SimCardDataUsageNotificationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SimCardDataUsageNotificationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("sim_card_data_usage_notifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The SIM card individual data usage notification information.
type SimCardDataUsageNotification struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  string `json:"created_at"`
	RecordType string `json:"record_type"`
	// The identification UUID of the related SIM card resource.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// Data usage threshold that will trigger the notification.
	Threshold SimCardDataUsageNotificationThreshold `json:"threshold"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		SimCardID   respjson.Field
		Threshold   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotification) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SimCardDataUsageNotification to a
// SimCardDataUsageNotificationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SimCardDataUsageNotificationParam.Overrides()
func (r SimCardDataUsageNotification) ToParam() SimCardDataUsageNotificationParam {
	return param.Override[SimCardDataUsageNotificationParam](json.RawMessage(r.RawJSON()))
}

// Data usage threshold that will trigger the notification.
type SimCardDataUsageNotificationThreshold struct {
	Amount string `json:"amount"`
	// Any of "MB", "GB".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationThreshold) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card individual data usage notification information.
type SimCardDataUsageNotificationParam struct {
	// Identifies the resource.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt  param.Opt[string] `json:"created_at,omitzero"`
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	// The identification UUID of the related SIM card resource.
	SimCardID param.Opt[string] `json:"sim_card_id,omitzero" format:"uuid"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt param.Opt[string] `json:"updated_at,omitzero"`
	// Data usage threshold that will trigger the notification.
	Threshold SimCardDataUsageNotificationThresholdParam `json:"threshold,omitzero"`
	paramObj
}

func (r SimCardDataUsageNotificationParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardDataUsageNotificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardDataUsageNotificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data usage threshold that will trigger the notification.
type SimCardDataUsageNotificationThresholdParam struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Any of "MB", "GB".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r SimCardDataUsageNotificationThresholdParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardDataUsageNotificationThresholdParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardDataUsageNotificationThresholdParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SimCardDataUsageNotificationThresholdParam](
		"unit", "MB", "GB",
	)
}

type SimCardDataUsageNotificationNewResponse struct {
	// The SIM card individual data usage notification information.
	Data SimCardDataUsageNotification `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDataUsageNotificationGetResponse struct {
	// The SIM card individual data usage notification information.
	Data SimCardDataUsageNotification `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDataUsageNotificationUpdateResponse struct {
	// The SIM card individual data usage notification information.
	Data SimCardDataUsageNotification `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDataUsageNotificationListResponse struct {
	Data []SimCardDataUsageNotification `json:"data"`
	Meta PaginationMeta                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationListResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDataUsageNotificationDeleteResponse struct {
	// The SIM card individual data usage notification information.
	Data SimCardDataUsageNotification `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardDataUsageNotificationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SimCardDataUsageNotificationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimCardDataUsageNotificationNewParams struct {
	// The identification UUID of the related SIM card resource.
	SimCardID string `json:"sim_card_id,required" format:"uuid"`
	// Data usage threshold that will trigger the notification.
	Threshold SimCardDataUsageNotificationNewParamsThreshold `json:"threshold,omitzero,required"`
	paramObj
}

func (r SimCardDataUsageNotificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SimCardDataUsageNotificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardDataUsageNotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data usage threshold that will trigger the notification.
type SimCardDataUsageNotificationNewParamsThreshold struct {
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Any of "MB", "GB".
	Unit string `json:"unit,omitzero"`
	paramObj
}

func (r SimCardDataUsageNotificationNewParamsThreshold) MarshalJSON() (data []byte, err error) {
	type shadow SimCardDataUsageNotificationNewParamsThreshold
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardDataUsageNotificationNewParamsThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SimCardDataUsageNotificationNewParamsThreshold](
		"unit", "MB", "GB",
	)
}

type SimCardDataUsageNotificationUpdateParams struct {
	// The SIM card individual data usage notification information.
	SimCardDataUsageNotification SimCardDataUsageNotificationParam
	paramObj
}

func (r SimCardDataUsageNotificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimCardDataUsageNotification)
}
func (r *SimCardDataUsageNotificationUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SimCardDataUsageNotification)
}

type SimCardDataUsageNotificationListParams struct {
	// A valid SIM card ID.
	FilterSimCardID param.Opt[string] `query:"filter[sim_card_id],omitzero" format:"uuid" json:"-"`
	// The page number to load.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SimCardDataUsageNotificationListParams]'s query parameters
// as `url.Values`.
func (r SimCardDataUsageNotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
