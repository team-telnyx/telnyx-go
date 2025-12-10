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
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// NotificationProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationProfileService] method instead.
type NotificationProfileService struct {
	Options []option.RequestOption
}

// NewNotificationProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNotificationProfileService(opts ...option.RequestOption) (r NotificationProfileService) {
	r = NotificationProfileService{}
	r.Options = opts
	return
}

// Create a notification profile.
func (r *NotificationProfileService) New(ctx context.Context, body NotificationProfileNewParams, opts ...option.RequestOption) (res *NotificationProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "notification_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a notification profile.
func (r *NotificationProfileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a notification profile.
func (r *NotificationProfileService) Update(ctx context.Context, notificationProfileID string, body NotificationProfileUpdateParams, opts ...option.RequestOption) (res *NotificationProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if notificationProfileID == "" {
		err = errors.New("missing required notification_profile_id parameter")
		return
	}
	path := fmt.Sprintf("notification_profiles/%s", notificationProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your notifications profiles.
func (r *NotificationProfileService) List(ctx context.Context, query NotificationProfileListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[NotificationProfile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "notification_profiles"
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

// Returns a list of your notifications profiles.
func (r *NotificationProfileService) ListAutoPaging(ctx context.Context, query NotificationProfileListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[NotificationProfile] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a notification profile.
func (r *NotificationProfileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *NotificationProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("notification_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A Collection of Notification Channels
type NotificationProfile struct {
	// A UUID.
	ID string `json:"id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A human readable name.
	Name string `json:"name"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationProfile) RawJSON() string { return r.JSON.raw }
func (r *NotificationProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NotificationProfile to a NotificationProfileParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NotificationProfileParam.Overrides()
func (r NotificationProfile) ToParam() NotificationProfileParam {
	return param.Override[NotificationProfileParam](json.RawMessage(r.RawJSON()))
}

// A Collection of Notification Channels
type NotificationProfileParam struct {
	// A human readable name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r NotificationProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationProfileNewResponse struct {
	// A Collection of Notification Channels
	Data NotificationProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationProfileGetResponse struct {
	// A Collection of Notification Channels
	Data NotificationProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationProfileUpdateResponse struct {
	// A Collection of Notification Channels
	Data NotificationProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationProfileDeleteResponse struct {
	// A Collection of Notification Channels
	Data NotificationProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *NotificationProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationProfileNewParams struct {
	// A Collection of Notification Channels
	NotificationProfile NotificationProfileParam
	paramObj
}

func (r NotificationProfileNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationProfile)
}
func (r *NotificationProfileNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NotificationProfile)
}

type NotificationProfileUpdateParams struct {
	// A Collection of Notification Channels
	NotificationProfile NotificationProfileParam
	paramObj
}

func (r NotificationProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NotificationProfile)
}
func (r *NotificationProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NotificationProfile)
}

type NotificationProfileListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NotificationProfileListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationProfileListParams]'s query parameters as
// `url.Values`.
func (r NotificationProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NotificationProfileListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NotificationProfileListParamsPage]'s query parameters as
// `url.Values`.
func (r NotificationProfileListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
