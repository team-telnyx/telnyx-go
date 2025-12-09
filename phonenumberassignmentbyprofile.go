// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// PhoneNumberAssignmentByProfileService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberAssignmentByProfileService] method instead.
type PhoneNumberAssignmentByProfileService struct {
	Options []option.RequestOption
}

// NewPhoneNumberAssignmentByProfileService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPhoneNumberAssignmentByProfileService(opts ...option.RequestOption) (r PhoneNumberAssignmentByProfileService) {
	r = PhoneNumberAssignmentByProfileService{}
	r.Options = opts
	return
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusStarting  TaskStatus = "starting"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)
