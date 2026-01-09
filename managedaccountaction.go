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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ManagedAccountActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedAccountActionService] method instead.
type ManagedAccountActionService struct {
	Options []option.RequestOption
}

// NewManagedAccountActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManagedAccountActionService(opts ...option.RequestOption) (r ManagedAccountActionService) {
	r = ManagedAccountActionService{}
	r.Options = opts
	return
}

// Disables a managed account, forbidding it to use Telnyx services, including
// sending or receiving phone calls and SMS messages. Ongoing phone calls will not
// be affected. The managed account and its sub-users will no longer be able to log
// in via the mission control portal.
func (r *ManagedAccountActionService) Disable(ctx context.Context, id string, opts ...option.RequestOption) (res *ManagedAccountActionDisableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("managed_accounts/%s/actions/disable", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enables a managed account and its sub-users to use Telnyx services.
func (r *ManagedAccountActionService) Enable(ctx context.Context, id string, body ManagedAccountActionEnableParams, opts ...option.RequestOption) (res *ManagedAccountActionEnableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("managed_accounts/%s/actions/enable", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ManagedAccountActionDisableResponse struct {
	Data ManagedAccount `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountActionDisableResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountActionDisableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountActionEnableResponse struct {
	Data ManagedAccount `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ManagedAccountActionEnableResponse) RawJSON() string { return r.JSON.raw }
func (r *ManagedAccountActionEnableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ManagedAccountActionEnableParams struct {
	// When true, all connections owned by this managed account will automatically be
	// re-enabled. Note: Any connections that do not pass validations will not be
	// re-enabled.
	ReenableAllConnections param.Opt[bool] `json:"reenable_all_connections,omitzero"`
	paramObj
}

func (r ManagedAccountActionEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow ManagedAccountActionEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ManagedAccountActionEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
