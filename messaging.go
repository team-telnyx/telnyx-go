// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// MessagingService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingService] method instead.
type MessagingService struct {
	Options []option.RequestOption
	Rcs     MessagingRcService
}

// NewMessagingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingService(opts ...option.RequestOption) (r MessagingService) {
	r = MessagingService{}
	r.Options = opts
	r.Rcs = NewMessagingRcService(opts...)
	return
}
