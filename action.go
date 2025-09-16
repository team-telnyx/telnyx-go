// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/stainless-sdks/telnyx-go/option"
)

// ActionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	Options  []option.RequestOption
	Purchase ActionPurchaseService
	Register ActionRegisterService
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.Options = opts
	r.Purchase = NewActionPurchaseService(opts...)
	r.Register = NewActionRegisterService(opts...)
	return
}
