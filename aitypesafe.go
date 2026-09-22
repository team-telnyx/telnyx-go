// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AITypesafeService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITypesafeService] method instead.
type AITypesafeService struct {
	Options []option.RequestOption
	// Beta API for evaluating shared context with typed questions and structured
	// answers. Telnyx manages model selection.
	V1 AITypesafeV1Service
}

// NewAITypesafeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAITypesafeService(opts ...option.RequestOption) (r AITypesafeService) {
	r = AITypesafeService{}
	r.Options = opts
	r.V1 = NewAITypesafeV1Service(opts...)
	return
}
