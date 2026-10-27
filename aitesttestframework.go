// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AITestTestFrameworkService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkService] method instead.
type AITestTestFrameworkService struct {
	Options []option.RequestOption
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	V1 AITestTestFrameworkV1Service
}

// NewAITestTestFrameworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkService(opts ...option.RequestOption) (r AITestTestFrameworkService) {
	r = AITestTestFrameworkService{}
	r.Options = opts
	r.V1 = NewAITestTestFrameworkV1Service(opts...)
	return
}
