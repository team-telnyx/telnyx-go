// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Accept and review the Branded Calling and Phone Number Reputation terms of
// service.
//
// TermsOfServiceBrandedCallingService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTermsOfServiceBrandedCallingService] method instead.
type TermsOfServiceBrandedCallingService struct {
	Options []option.RequestOption
}

// NewTermsOfServiceBrandedCallingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTermsOfServiceBrandedCallingService(opts ...option.RequestOption) (r TermsOfServiceBrandedCallingService) {
	r = TermsOfServiceBrandedCallingService{}
	r.Options = opts
	return
}

// Records the authenticated user's agreement to the current Branded Calling ToS.
// No body required. Idempotent - re-calling after agreement is a no-op and returns
// the existing agreement.
//
// This is a prerequisite for activating Branded Calling on any enterprise
// (`POST /enterprises/{id}/branded_calling`); without an agreement, activation
// returns `403`.
func (r *TermsOfServiceBrandedCallingService) Agree(ctx context.Context, opts ...option.RequestOption) (res *TosAgreementWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "terms_of_service/branded_calling/agree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}
