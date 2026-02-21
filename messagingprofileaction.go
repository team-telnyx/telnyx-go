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

// MessagingProfileActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingProfileActionService] method instead.
type MessagingProfileActionService struct {
	Options []option.RequestOption
}

// NewMessagingProfileActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingProfileActionService(opts ...option.RequestOption) (r MessagingProfileActionService) {
	r = MessagingProfileActionService{}
	r.Options = opts
	return
}

// Regenerate the v1 secret for a messaging profile.
func (r *MessagingProfileActionService) RegenerateSecret(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingProfileActionRegenerateSecretResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_profiles/%s/actions/regenerate_secret", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type MessagingProfileActionRegenerateSecretResponse struct {
	Data MessagingProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingProfileActionRegenerateSecretResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingProfileActionRegenerateSecretResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
