// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/stainless-sdks/telnyx-go/v4/option"
)

// ReputationService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReputationService] method instead.
type ReputationService struct {
	Options []option.RequestOption
	// Phone-number reputation monitoring (spam-score lookup and tracking).
	Numbers ReputationNumberService
}

// NewReputationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReputationService(opts ...option.RequestOption) (r ReputationService) {
	r = ReputationService{}
	r.Options = opts
	r.Numbers = NewReputationNumberService(opts...)
	return
}
