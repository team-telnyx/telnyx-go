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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// MessagingHostedNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingHostedNumberService] method instead.
type MessagingHostedNumberService struct {
	Options []option.RequestOption
}

// NewMessagingHostedNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingHostedNumberService(opts ...option.RequestOption) (r MessagingHostedNumberService) {
	r = MessagingHostedNumberService{}
	r.Options = opts
	return
}

// Delete a messaging hosted number
func (r *MessagingHostedNumberService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MessagingHostedNumberDeleteResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
