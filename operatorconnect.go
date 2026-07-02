// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// OperatorConnectService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperatorConnectService] method instead.
type OperatorConnectService struct {
	Options []option.RequestOption
	// External Connections operations
	Actions OperatorConnectActionService
}

// NewOperatorConnectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperatorConnectService(opts ...option.RequestOption) (r OperatorConnectService) {
	r = OperatorConnectService{}
	r.Options = opts
	r.Actions = NewOperatorConnectActionService(opts...)
	return
}
