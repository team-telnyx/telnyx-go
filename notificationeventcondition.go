// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// NotificationEventConditionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationEventConditionService] method instead.
type NotificationEventConditionService struct {
	Options []option.RequestOption
}

// NewNotificationEventConditionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNotificationEventConditionService(opts ...option.RequestOption) (r NotificationEventConditionService) {
	r = NotificationEventConditionService{}
	r.Options = opts
	return
}

// Returns a list of your notifications events conditions.
func (r *NotificationEventConditionService) List(ctx context.Context, query NotificationEventConditionListParams, opts ...option.RequestOption) (res *NotificationEventConditionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_event_conditions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NotificationEventConditionListResponse struct {
	Data []NotificationEventConditionListResponseData `json:"data"`
	Meta PaginationMeta                               `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationEventConditionListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationEventConditionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationEventConditionListResponseData struct {
	// A UUID.
	ID string `json:"id"`
	// Dictates whether a notification channel id needs to be provided when creating a
	// notficiation setting.
	AllowMultipleChannels bool `json:"allow_multiple_channels"`
	// Any of "account", "phone_number".
	AssociatedRecordType string `json:"associated_record_type"`
	// Dictates whether a notification setting will take effect immediately.
	Asynchronous bool `json:"asynchronous"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt           time.Time                                             `json:"created_at" format:"date-time"`
	Description         string                                                `json:"description"`
	Enabled             bool                                                  `json:"enabled"`
	Name                string                                                `json:"name"`
	NotificationEventID string                                                `json:"notification_event_id"`
	Parameters          []NotificationEventConditionListResponseDataParameter `json:"parameters"`
	// Dictates the supported notification channel types that can be emitted.
	SupportedChannels []string `json:"supported_channels"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AllowMultipleChannels respjson.Field
		AssociatedRecordType  respjson.Field
		Asynchronous          respjson.Field
		CreatedAt             respjson.Field
		Description           respjson.Field
		Enabled               respjson.Field
		Name                  respjson.Field
		NotificationEventID   respjson.Field
		Parameters            respjson.Field
		SupportedChannels     respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationEventConditionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NotificationEventConditionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationEventConditionListResponseDataParameter struct {
	DataType string `json:"data_type"`
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataType    respjson.Field
		Name        respjson.Field
		Optional    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationEventConditionListResponseDataParameter) RawJSON() string { return r.JSON.raw }
func (r *NotificationEventConditionListResponseDataParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationEventConditionListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[associated_record_type][eq], filter[channel_type_id][eq],
	// filter[notification_profile_id][eq], filter[notification_channel][eq],
	// filter[notification_event_condition_id][eq], filter[status][eq]
	Filter NotificationEventConditionListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NotificationEventConditionListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationEventConditionListParams]'s query parameters as
// `url.Values`.
func (r NotificationEventConditionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[associated_record_type][eq], filter[channel_type_id][eq],
// filter[notification_profile_id][eq], filter[notification_channel][eq],
// filter[notification_event_condition_id][eq], filter[status][eq]
type NotificationEventConditionListParamsFilter struct {
	AssociatedRecordType         NotificationEventConditionListParamsFilterAssociatedRecordType         `query:"associated_record_type,omitzero" json:"-"`
	ChannelTypeID                NotificationEventConditionListParamsFilterChannelTypeID                `query:"channel_type_id,omitzero" json:"-"`
	NotificationChannel          NotificationEventConditionListParamsFilterNotificationChannel          `query:"notification_channel,omitzero" json:"-"`
	NotificationEventConditionID NotificationEventConditionListParamsFilterNotificationEventConditionID `query:"notification_event_condition_id,omitzero" json:"-"`
	NotificationProfileID        NotificationEventConditionListParamsFilterNotificationProfileID        `query:"notification_profile_id,omitzero" json:"-"`
	Status                       NotificationEventConditionListParamsFilterStatus                       `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationEventConditionListParamsFilter]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterAssociatedRecordType struct {
	// Filter by the associated record type
	//
	// Any of "account", "phone_number".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationEventConditionListParamsFilterAssociatedRecordType]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterAssociatedRecordType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterChannelTypeID struct {
	// Filter by the id of a channel type
	//
	// Any of "webhook", "sms", "email", "voice".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationEventConditionListParamsFilterChannelTypeID]'s
// query parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterChannelTypeID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterNotificationChannel struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationEventConditionListParamsFilterNotificationChannel]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterNotificationChannel) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterNotificationEventConditionID struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationEventConditionListParamsFilterNotificationEventConditionID]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterNotificationEventConditionID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterNotificationProfileID struct {
	// Filter by the id of a notification profile
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationEventConditionListParamsFilterNotificationProfileID]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterNotificationProfileID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationEventConditionListParamsFilterStatus struct {
	// The status of a notification setting
	//
	// Any of "enabled", "enable-received", "enable-pending", "enable-submtited",
	// "delete-received", "delete-pending", "delete-submitted", "deleted".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationEventConditionListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NotificationEventConditionListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationEventConditionListParamsPage]'s query
// parameters as `url.Values`.
func (r NotificationEventConditionListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
