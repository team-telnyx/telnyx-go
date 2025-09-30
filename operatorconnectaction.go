// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// OperatorConnectActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperatorConnectActionService] method instead.
type OperatorConnectActionService struct {
	Options []option.RequestOption
}

// NewOperatorConnectActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperatorConnectActionService(opts ...option.RequestOption) (r OperatorConnectActionService) {
	r = OperatorConnectActionService{}
	r.Options = opts
	return
}

// This endpoint will make an asynchronous request to refresh the Operator Connect
// integration with Microsoft Teams for the current user. This will create new
// external connections on the user's account if needed, and/or report the
// integration results as
// [log messages](https://developers.telnyx.com/api/external-voice-integrations/list-external-connection-log-messages).
func (r *OperatorConnectActionService) Refresh(ctx context.Context, opts ...option.RequestOption) (res *OperatorConnectActionRefreshResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "operator_connect/actions/refresh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type OperatorConnectActionRefreshResponse struct {
	// A message describing the result of the operation
	Message string `json:"message"`
	// Describes wether or not the operation was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperatorConnectActionRefreshResponse) RawJSON() string { return r.JSON.raw }
func (r *OperatorConnectActionRefreshResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
