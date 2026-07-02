// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// X402Service contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewX402Service] method instead.
type X402Service struct {
	Options []option.RequestOption
	// Operations for x402 cryptocurrency payment transactions. Fund your Telnyx
	// account using USDC stablecoin payments via the x402 protocol.
	CreditAccount X402CreditAccountService
}

// NewX402Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewX402Service(opts ...option.RequestOption) (r X402Service) {
	r = X402Service{}
	r.Options = opts
	r.CreditAccount = NewX402CreditAccountService(opts...)
	return
}
