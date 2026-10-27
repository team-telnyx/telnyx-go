// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AITestScheduleV1Service contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestScheduleV1Service] method instead.
type AITestScheduleV1Service struct {
	Options []option.RequestOption
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	CronJobs AITestScheduleV1CronJobService
}

// NewAITestScheduleV1Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITestScheduleV1Service(opts ...option.RequestOption) (r AITestScheduleV1Service) {
	r = AITestScheduleV1Service{}
	r.Options = opts
	r.CronJobs = NewAITestScheduleV1CronJobService(opts...)
	return
}
