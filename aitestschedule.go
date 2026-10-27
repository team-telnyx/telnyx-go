// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AITestScheduleService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestScheduleService] method instead.
type AITestScheduleService struct {
	Options []option.RequestOption
	V1      AITestScheduleV1Service
}

// NewAITestScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAITestScheduleService(opts ...option.RequestOption) (r AITestScheduleService) {
	r = AITestScheduleService{}
	r.Options = opts
	r.V1 = NewAITestScheduleV1Service(opts...)
	return
}
