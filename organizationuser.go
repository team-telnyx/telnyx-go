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

// OrganizationUserService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserService] method instead.
type OrganizationUserService struct {
	Options []option.RequestOption
	Actions OrganizationUserActionService
}

// NewOrganizationUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserService(opts ...option.RequestOption) (r OrganizationUserService) {
	r = OrganizationUserService{}
	r.Options = opts
	r.Actions = NewOrganizationUserActionService(opts...)
	return
}

// Returns a user in your organization.
func (r *OrganizationUserService) Get(ctx context.Context, id string, query OrganizationUserGetParams, opts ...option.RequestOption) (res *OrganizationUserGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of the users in your organization.
func (r *OrganizationUserService) List(ctx context.Context, query OrganizationUserListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[OrganizationUserListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "organizations/users"
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

// Returns a list of the users in your organization.
func (r *OrganizationUserService) ListAutoPaging(ctx context.Context, query OrganizationUserListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[OrganizationUserListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Returns a report of all users in your organization with their group memberships.
// This endpoint returns all users without pagination and always includes group
// information. The report can be retrieved in JSON or CSV format by sending
// specific content-type headers.
func (r *OrganizationUserService) GetGroupsReport(ctx context.Context, query OrganizationUserGetGroupsReportParams, opts ...option.RequestOption) (res *OrganizationUserGetGroupsReportResponse, err error) {
	if !param.IsOmitted(query.Accept) {
		opts = append(opts, option.WithHeader("Accept", fmt.Sprintf("%s", query.Accept)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "organizations/users/users_groups_report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A reference to a group that a user belongs to.
type UserGroupReference struct {
	// The unique identifier of the group.
	ID string `json:"id,required"`
	// The name of the group.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGroupReference) RawJSON() string { return r.JSON.raw }
func (r *UserGroupReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetResponse struct {
	Data OrganizationUserGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetResponseData struct {
	// Identifies the specific resource.
	ID string `json:"id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The email address of the user.
	Email string `json:"email" format:"email"`
	// The groups the user belongs to. Only included when include_groups parameter is
	// true.
	Groups []UserGroupReference `json:"groups"`
	// ISO 8601 formatted date indicating when the resource last signed into the
	// portal. Null if the user has never signed in.
	LastSignInAt string `json:"last_sign_in_at,nullable"`
	// Indicates whether this user is allowed to bypass SSO and use password
	// authentication.
	OrganizationUserBypassesSSO bool `json:"organization_user_bypasses_sso"`
	// Identifies the type of the resource. Can be 'organization_owner' or
	// 'organization_sub_user'.
	RecordType string `json:"record_type"`
	// The status of the account.
	//
	// Any of "enabled", "disabled", "blocked".
	UserStatus string `json:"user_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		CreatedAt                   respjson.Field
		Email                       respjson.Field
		Groups                      respjson.Field
		LastSignInAt                respjson.Field
		OrganizationUserBypassesSSO respjson.Field
		RecordType                  respjson.Field
		UserStatus                  respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserListResponse struct {
	// Identifies the specific resource.
	ID string `json:"id"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The email address of the user.
	Email string `json:"email" format:"email"`
	// The groups the user belongs to. Only included when include_groups parameter is
	// true.
	Groups []UserGroupReference `json:"groups"`
	// ISO 8601 formatted date indicating when the resource last signed into the
	// portal. Null if the user has never signed in.
	LastSignInAt string `json:"last_sign_in_at,nullable"`
	// Indicates whether this user is allowed to bypass SSO and use password
	// authentication.
	OrganizationUserBypassesSSO bool `json:"organization_user_bypasses_sso"`
	// Identifies the type of the resource. Can be 'organization_owner' or
	// 'organization_sub_user'.
	RecordType string `json:"record_type"`
	// The status of the account.
	//
	// Any of "enabled", "disabled", "blocked".
	UserStatus OrganizationUserListResponseUserStatus `json:"user_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		CreatedAt                   respjson.Field
		Email                       respjson.Field
		Groups                      respjson.Field
		LastSignInAt                respjson.Field
		OrganizationUserBypassesSSO respjson.Field
		RecordType                  respjson.Field
		UserStatus                  respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the account.
type OrganizationUserListResponseUserStatus string

const (
	OrganizationUserListResponseUserStatusEnabled  OrganizationUserListResponseUserStatus = "enabled"
	OrganizationUserListResponseUserStatusDisabled OrganizationUserListResponseUserStatus = "disabled"
	OrganizationUserListResponseUserStatusBlocked  OrganizationUserListResponseUserStatus = "blocked"
)

type OrganizationUserGetGroupsReportResponse struct {
	Data []OrganizationUserGetGroupsReportResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserGetGroupsReportResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserGetGroupsReportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An organization user with their group memberships always included.
type OrganizationUserGetGroupsReportResponseData struct {
	// Identifies the specific resource.
	ID string `json:"id,required"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at,required"`
	// The email address of the user.
	Email string `json:"email,required" format:"email"`
	// The groups the user belongs to.
	Groups []UserGroupReference `json:"groups,required"`
	// Identifies the type of the resource. Can be 'organization_owner' or
	// 'organization_sub_user'.
	RecordType string `json:"record_type,required"`
	// The status of the account.
	//
	// Any of "enabled", "disabled", "blocked".
	UserStatus string `json:"user_status,required"`
	// ISO 8601 formatted date indicating when the resource last signed into the
	// portal. Null if the user has never signed in.
	LastSignInAt string `json:"last_sign_in_at,nullable"`
	// Indicates whether this user is allowed to bypass SSO and use password
	// authentication.
	OrganizationUserBypassesSSO bool `json:"organization_user_bypasses_sso"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		CreatedAt                   respjson.Field
		Email                       respjson.Field
		Groups                      respjson.Field
		RecordType                  respjson.Field
		UserStatus                  respjson.Field
		LastSignInAt                respjson.Field
		OrganizationUserBypassesSSO respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserGetGroupsReportResponseData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserGetGroupsReportResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserGetParams struct {
	// When set to true, includes the groups array for each user in the response. The
	// groups array contains objects with id and name for each group the user belongs
	// to.
	IncludeGroups param.Opt[bool] `query:"include_groups,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationUserListParams struct {
	// Filter by email address (partial match)
	FilterEmail param.Opt[string] `query:"filter[email],omitzero" json:"-"`
	// When set to true, includes the groups array for each user in the response. The
	// groups array contains objects with id and name for each group the user belongs
	// to.
	IncludeGroups param.Opt[bool] `query:"include_groups,omitzero" json:"-"`
	// The page number to load
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by user status
	//
	// Any of "enabled", "disabled", "blocked".
	FilterUserStatus OrganizationUserListParamsFilterUserStatus `query:"filter[user_status],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationUserListParams]'s query parameters as
// `url.Values`.
func (r OrganizationUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by user status
type OrganizationUserListParamsFilterUserStatus string

const (
	OrganizationUserListParamsFilterUserStatusEnabled  OrganizationUserListParamsFilterUserStatus = "enabled"
	OrganizationUserListParamsFilterUserStatusDisabled OrganizationUserListParamsFilterUserStatus = "disabled"
	OrganizationUserListParamsFilterUserStatusBlocked  OrganizationUserListParamsFilterUserStatus = "blocked"
)

type OrganizationUserGetGroupsReportParams struct {
	// Any of "application/json", "text/csv".
	Accept OrganizationUserGetGroupsReportParamsAccept `header:"Accept,omitzero" json:"-"`
	paramObj
}

type OrganizationUserGetGroupsReportParamsAccept string

const (
	OrganizationUserGetGroupsReportParamsAcceptApplicationJson OrganizationUserGetGroupsReportParamsAccept = "application/json"
	OrganizationUserGetGroupsReportParamsAcceptTextCsv         OrganizationUserGetGroupsReportParamsAccept = "text/csv"
)
