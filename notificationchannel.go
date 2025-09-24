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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// NotificationChannelService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationChannelService] method instead.
type NotificationChannelService struct {
	Options []option.RequestOption
}

// NewNotificationChannelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationChannelService(opts ...option.RequestOption) (r NotificationChannelService) {
	r = NotificationChannelService{}
	r.Options = opts
	return
}

// Create a notification channel.
func (r *NotificationChannelService) New(ctx context.Context, body NotificationChannelNewParams, opts ...option.RequestOption) (res *NotificationChannelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a notification channel.
func (r *NotificationChannelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationChannelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a notification channel.
func (r *NotificationChannelService) Update(ctx context.Context, id string, body NotificationChannelUpdateParams, opts ...option.RequestOption) (res *NotificationChannelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List notification channels.
func (r *NotificationChannelService) List(ctx context.Context, query NotificationChannelListParams, opts ...option.RequestOption) (res *NotificationChannelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a notification channel.
func (r *NotificationChannelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationChannelDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_channels/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A Notification Channel
type NotificationChannel struct {
	// A UUID.
	ID string `json:"id"`
	// The destination associated with the channel type.
	ChannelDestination string `json:"channel_destination"`
	// A Channel Type ID
	//
	// Any of "sms", "voice", "email", "webhook".
	ChannelTypeID NotificationChannelChannelTypeID `json:"channel_type_id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileID string `json:"notification_profile_id"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ChannelDestination    respjson.Field
		ChannelTypeID         respjson.Field
		CreatedAt             respjson.Field
		NotificationProfileID respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationChannel) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationChannel to a NotificationChannelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationChannelParam.Overrides()
func (r NotificationChannel) ToParam() NotificationChannelParam {
	return param.Override[NotificationChannelParam](json.RawMessage(r.RawJSON()))
}

// A Channel Type ID
type NotificationChannelChannelTypeID string

const (
	NotificationChannelChannelTypeIDSMS     NotificationChannelChannelTypeID = "sms"
	NotificationChannelChannelTypeIDVoice   NotificationChannelChannelTypeID = "voice"
	NotificationChannelChannelTypeIDEmail   NotificationChannelChannelTypeID = "email"
	NotificationChannelChannelTypeIDWebhook NotificationChannelChannelTypeID = "webhook"
)

// A Notification Channel
type NotificationChannelParam struct {
	// A UUID.
	ID param.Opt[string] `json:"id,omitzero"`
	// The destination associated with the channel type.
	ChannelDestination param.Opt[string] `json:"channel_destination,omitzero"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt param.Opt[time.Time] `json:"created_at,omitzero" format:"date-time"`
	// A UUID reference to the associated Notification Profile.
	NotificationProfileID param.Opt[string] `json:"notification_profile_id,omitzero"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt param.Opt[time.Time] `json:"updated_at,omitzero" format:"date-time"`
	// A Channel Type ID
	//
	// Any of "sms", "voice", "email", "webhook".
	ChannelTypeID NotificationChannelChannelTypeID `json:"channel_type_id,omitzero"`
	paramObj
}

func (r NotificationChannelParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationChannelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationChannelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelNewResponse struct {
	// A Notification Channel
	Data NotificationChannel `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationChannelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelGetResponse struct {
	// A Notification Channel
	Data NotificationChannel `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationChannelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelUpdateResponse struct {
	// A Notification Channel
	Data NotificationChannel `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationChannelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelListResponse struct {
	Data []NotificationChannel `json:"data"`
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
func (r NotificationChannelListResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelDeleteResponse struct {
	// A Notification Channel
	Data NotificationChannel `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationChannelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationChannelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationChannelNewParams struct {
	// A Notification Channel
	NotificationChannel NotificationChannelParam
	paramObj
}

func (r NotificationChannelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationChannel)
}
func (r *NotificationChannelNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NotificationChannel)
}

type NotificationChannelUpdateParams struct {
	// A Notification Channel
	NotificationChannel NotificationChannelParam
	paramObj
}

func (r NotificationChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationChannel)
}
func (r *NotificationChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NotificationChannel)
}

type NotificationChannelListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[associated_record_type][eq], filter[channel_type_id][eq],
	// filter[notification_profile_id][eq], filter[notification_channel][eq],
	// filter[notification_event_condition_id][eq], filter[status][eq]
	Filter NotificationChannelListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NotificationChannelListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParams]'s query parameters as
// `url.Values`.
func (r NotificationChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[associated_record_type][eq], filter[channel_type_id][eq],
// filter[notification_profile_id][eq], filter[notification_channel][eq],
// filter[notification_event_condition_id][eq], filter[status][eq]
type NotificationChannelListParamsFilter struct {
	AssociatedRecordType         NotificationChannelListParamsFilterAssociatedRecordType         `query:"associated_record_type,omitzero" json:"-"`
	ChannelTypeID                NotificationChannelListParamsFilterChannelTypeID                `query:"channel_type_id,omitzero" json:"-"`
	NotificationChannel          NotificationChannelListParamsFilterNotificationChannel          `query:"notification_channel,omitzero" json:"-"`
	NotificationEventConditionID NotificationChannelListParamsFilterNotificationEventConditionID `query:"notification_event_condition_id,omitzero" json:"-"`
	NotificationProfileID        NotificationChannelListParamsFilterNotificationProfileID        `query:"notification_profile_id,omitzero" json:"-"`
	Status                       NotificationChannelListParamsFilterStatus                       `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilter]'s query parameters as
// `url.Values`.
func (r NotificationChannelListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterAssociatedRecordType struct {
	// Filter by the associated record type
	//
	// Any of "account", "phone_number".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilterAssociatedRecordType]'s
// query parameters as `url.Values`.
func (r NotificationChannelListParamsFilterAssociatedRecordType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterChannelTypeID struct {
	// Filter by the id of a channel type
	//
	// Any of "webhook", "sms", "email", "voice".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilterChannelTypeID]'s query
// parameters as `url.Values`.
func (r NotificationChannelListParamsFilterChannelTypeID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterNotificationChannel struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilterNotificationChannel]'s
// query parameters as `url.Values`.
func (r NotificationChannelListParamsFilterNotificationChannel) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterNotificationEventConditionID struct {
	// Filter by the id of a notification channel
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NotificationChannelListParamsFilterNotificationEventConditionID]'s query
// parameters as `url.Values`.
func (r NotificationChannelListParamsFilterNotificationEventConditionID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterNotificationProfileID struct {
	// Filter by the id of a notification profile
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilterNotificationProfileID]'s
// query parameters as `url.Values`.
func (r NotificationChannelListParamsFilterNotificationProfileID) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationChannelListParamsFilterStatus struct {
	// The status of a notification setting
	//
	// Any of "enabled", "enable-received", "enable-pending", "enable-submtited",
	// "delete-received", "delete-pending", "delete-submitted", "deleted".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsFilterStatus]'s query
// parameters as `url.Values`.
func (r NotificationChannelListParamsFilterStatus) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NotificationChannelListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationChannelListParamsPage]'s query parameters as
// `url.Values`.
func (r NotificationChannelListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
