// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// PhoneNumberBlockService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberBlockService] method instead.
type PhoneNumberBlockService struct {
	Options []option.RequestOption
	Jobs    PhoneNumberBlockJobService
}

// NewPhoneNumberBlockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberBlockService(opts ...option.RequestOption) (r PhoneNumberBlockService) {
	r = PhoneNumberBlockService{}
	r.Options = opts
	r.Jobs = NewPhoneNumberBlockJobService(opts...)
	return
}
