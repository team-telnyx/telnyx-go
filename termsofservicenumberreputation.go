// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Terms of Service agreement endpoints
//
// TermsOfServiceNumberReputationService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTermsOfServiceNumberReputationService] method instead.
type TermsOfServiceNumberReputationService struct {
	Options []option.RequestOption
}

// NewTermsOfServiceNumberReputationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTermsOfServiceNumberReputationService(opts ...option.RequestOption) (r TermsOfServiceNumberReputationService) {
	r = TermsOfServiceNumberReputationService{}
	r.Options = opts
	return
}

// Accept the Terms of Service for the Number Reputation product. Must be called
// before using Number Reputation endpoints.
//
// Returns `400` with error code `10015` if the user has already agreed to the
// current ToS version.
func (r *TermsOfServiceNumberReputationService) Agree(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "terms_of_service/number_reputation/agree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}
