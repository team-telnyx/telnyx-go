// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// OrganizationUserActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserActionService] method instead.
type OrganizationUserActionService struct {
	Options []option.RequestOption
}

// NewOrganizationUserActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserActionService(opts ...option.RequestOption) (r OrganizationUserActionService) {
	r = OrganizationUserActionService{}
	r.Options = opts
	return
}

// Deletes a user in your organization.
func (r *OrganizationUserActionService) Remove(ctx context.Context, id string, opts ...option.RequestOption) (res *OrganizationUserActionRemoveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/users/%s/actions/remove", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type OrganizationUserActionRemoveResponse struct {
	Data OrganizationUserActionRemoveResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationUserActionRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserActionRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUserActionRemoveResponseData struct {
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
func (r OrganizationUserActionRemoveResponseData) RawJSON() string { return r.JSON.raw }
func (r *OrganizationUserActionRemoveResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
