// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AIAnthropicService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAnthropicService] method instead.
type AIAnthropicService struct {
	Options []option.RequestOption
	V1      AIAnthropicV1Service
}

// NewAIAnthropicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAnthropicService(opts ...option.RequestOption) (r AIAnthropicService) {
	r = AIAnthropicService{}
	r.Options = opts
	r.V1 = NewAIAnthropicV1Service(opts...)
	return
}
