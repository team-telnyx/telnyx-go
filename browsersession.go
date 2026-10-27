// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// BrowserSessionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserSessionService] method instead.
type BrowserSessionService struct {
	Options []option.RequestOption
}

// NewBrowserSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrowserSessionService(opts ...option.RequestOption) (r BrowserSessionService) {
	r = BrowserSessionService{}
	r.Options = opts
	return
}
