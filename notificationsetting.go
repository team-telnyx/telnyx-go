// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// NotificationSettingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationSettingService] method instead.
type NotificationSettingService struct {
	Options []option.RequestOption
}

// NewNotificationSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationSettingService(opts ...option.RequestOption) (r NotificationSettingService) {
	r = NotificationSettingService{}
	r.Options = opts
	return
}

// Add a notification setting.
func (r *NotificationSettingService) New(ctx context.Context, body NotificationSettingNewParams, opts ...option.RequestOption) (res *NotificationSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a notification setting.
func (r *NotificationSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_settings/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List notification settings.
func (r *NotificationSettingService) List(ctx context.Context, query NotificationSettingListParams, opts ...option.RequestOption) (res *NotificationSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a notification setting.
func (r *NotificationSettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationSettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_settings/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type NotificationSetting struct {
	// A UUID.
	ID                        string `json:"id"`
	AssociatedRecordType      string `json:"associated_record_type"`
	AssociatedRecordTypeValue string `json:"associated_record_type_value"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A UUID reference to the associated Notification Channel.
	NotificationChannelID string `json:"notification_channel_id"`
	// A UUID reference to the associated Notification Event Condition.
	NotificationEventConditionID string `json:"notification_event_condition_id"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileID string                         `json:"notification_profile_id"`
	Parameters            []NotificationSettingParameter `json:"parameters"`
	// Most preferences apply immediately; however, other may needs to propagate.
	//
	// Any of "enabled", "enable-received", "enable-pending", "enable-submtited",
	// "delete-received", "delete-pending", "delete-submitted", "deleted".
	Status NotificationSettingStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AssociatedRecordType         respjson.Field
		AssociatedRecordTypeValue    respjson.Field
		CreatedAt                    respjson.Field
		NotificationChannelID        respjson.Field
		NotificationEventConditionID respjson.Field
		NotificationProfileID        respjson.Field
		Parameters                   respjson.Field
		Status                       respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSetting) RawJSON() string { return r.JSON.raw }
func (r *NotificationSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationSetting to a NotificationSettingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationSettingParam.Overrides()
func (r NotificationSetting) ToParam() NotificationSettingParam {
	return param.Override[NotificationSettingParam](json.RawMessage(r.RawJSON()))
}

type NotificationSettingParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSettingParameter) RawJSON() string { return r.JSON.raw }
func (r *NotificationSettingParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Most preferences apply immediately; however, other may needs to propagate.
type NotificationSettingStatus string

const (
	NotificationSettingStatusEnabled         NotificationSettingStatus = "enabled"
	NotificationSettingStatusEnableReceived  NotificationSettingStatus = "enable-received"
	NotificationSettingStatusEnablePending   NotificationSettingStatus = "enable-pending"
	NotificationSettingStatusEnableSubmtited NotificationSettingStatus = "enable-submtited"
	NotificationSettingStatusDeleteReceived  NotificationSettingStatus = "delete-received"
	NotificationSettingStatusDeletePending   NotificationSettingStatus = "delete-pending"
	NotificationSettingStatusDeleteSubmitted NotificationSettingStatus = "delete-submitted"
	NotificationSettingStatusDeleted         NotificationSettingStatus = "deleted"
)

type NotificationSettingParam struct {
	// A UUID.
	ID                        param.Opt[string] `json:"id,omitzero"`
	AssociatedRecordType      param.Opt[string] `json:"associated_record_type,omitzero"`
	AssociatedRecordTypeValue param.Opt[string] `json:"associated_record_type_value,omitzero"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	// A UUID reference to the associated Notification Channel.
	NotificationChannelID param.Opt[string] `json:"notification_channel_id,omitzero"`
	// A UUID reference to the associated Notification Event Condition.
	NotificationEventConditionID param.Opt[string] `json:"notification_event_condition_id,omitzero"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileID param.Opt[string] `json:"notification_profile_id,omitzero"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt  param.Opt[time.Time]                `json:"updated_at,omitzero" format:"date-time"`
	Parameters []NotificationSettingParameterParam `json:"parameters,omitzero"`
	// Most preferences apply immediately; however, other may needs to propagate.
	//
	// Any of "enabled", "enable-received", "enable-pending", "enable-submtited",
	// "delete-received", "delete-pending", "delete-submitted", "deleted".
	Status NotificationSettingStatus `json:"status,omitzero"`
	paramObj
}

func (r NotificationSettingParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationSettingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationSettingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingParameterParam struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r NotificationSettingParameterParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationSettingParameterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationSettingParameterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingNewResponse struct {
	Data NotificationSetting `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingGetResponse struct {
	Data NotificationSetting `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingListResponse struct {
	Data []NotificationSetting `json:"data"`
	Meta PaginationMeta        `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingDeleteResponse struct {
	Data NotificationSetting `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationSettingNewParams struct {
	NotificationSetting NotificationSettingParam
	paramObj
}

func (r NotificationSettingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationSetting)
}
func (r *NotificationSettingNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NotificationSetting)
}

type NotificationSettingListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[associated_record_type][eq], filter[channel_type_id][eq],
	// filter[notification_profile_id][eq], filter[notification_channel][eq],
	// filter[notification_event_condition_id][eq], filter[status][eq]
	Filter NotificationSettingListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NotificationSettingListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParams]'s query parameters as
// `url.Values`.
func (r NotificationSettingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[associated_record_type][eq], filter[channel_type_id][eq],
// filter[notification_profile_id][eq], filter[notification_channel][eq],
// filter[notification_event_condition_id][eq], filter[status][eq]
type NotificationSettingListParamsFilter struct {
	AssociatedRecordType         NotificationSettingListParamsFilterAssociatedRecordType         `query:"associated_record_type,omitzero" json:"-"`
	ChannelTypeID                NotificationSettingListParamsFilterChannelTypeID                `query:"channel_type_id,omitzero" json:"-"`
	NotificationChannel          NotificationSettingListParamsFilterNotificationChannel          `query:"notification_channel,omitzero" json:"-"`
	NotificationEventConditionID NotificationSettingListParamsFilterNotificationEventConditionID `query:"notification_event_condition_id,omitzero" json:"-"`
	NotificationProfileID        NotificationSettingListParamsFilterNotificationProfileID        `query:"notification_profile_id,omitzero" json:"-"`
	Status                       NotificationSettingListParamsFilterStatus                       `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilter]'s query parameters as
// `url.Values`.
func (r NotificationSettingListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterAssociatedRecordType struct {
	// Filter by the associated record type
	//
	// Any of "account", "phone_number".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilterAssociatedRecordType]'s
// query parameters as `url.Values`.
func (r NotificationSettingListParamsFilterAssociatedRecordType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterChannelTypeID struct {
	// Filter by the id of a channel type
	//
	// Any of "webhook", "sms", "email", "voice".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilterChannelTypeID]'s query
// parameters as `url.Values`.
func (r NotificationSettingListParamsFilterChannelTypeID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterNotificationChannel struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilterNotificationChannel]'s
// query parameters as `url.Values`.
func (r NotificationSettingListParamsFilterNotificationChannel) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterNotificationEventConditionID struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationSettingListParamsFilterNotificationEventConditionID]'s query
// parameters as `url.Values`.
func (r NotificationSettingListParamsFilterNotificationEventConditionID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterNotificationProfileID struct {
	// Filter by the id of a notification profile
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilterNotificationProfileID]'s
// query parameters as `url.Values`.
func (r NotificationSettingListParamsFilterNotificationProfileID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationSettingListParamsFilterStatus struct {
	// The status of a notification setting
	//
	// Any of "enabled", "enable-received", "enable-pending", "enable-submtited",
	// "delete-received", "delete-pending", "delete-submitted", "deleted".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r NotificationSettingListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NotificationSettingListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationSettingListParamsPage]'s query parameters as
// `url.Values`.
func (r NotificationSettingListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
