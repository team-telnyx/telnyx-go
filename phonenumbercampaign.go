// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PhoneNumberCampaignService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberCampaignService] method instead.
type PhoneNumberCampaignService struct {
	Options []option.RequestOption
}

// NewPhoneNumberCampaignService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberCampaignService(opts ...option.RequestOption) (r PhoneNumberCampaignService) {
	r = PhoneNumberCampaignService{}
	r.Options = opts
	return
}

type PhoneNumberCampaign struct {
	// For shared campaigns, this is the TCR campaign ID, otherwise it is the campaign
	// ID
	CampaignID  string `json:"campaignId,required"`
	CreatedAt   string `json:"createdAt,required"`
	PhoneNumber string `json:"phoneNumber,required"`
	UpdatedAt   string `json:"updatedAt,required"`
	// The assignment status of the number.
	//
	// Any of "FAILED_ASSIGNMENT", "PENDING_ASSIGNMENT", "ASSIGNED",
	// "PENDING_UNASSIGNMENT", "FAILED_UNASSIGNMENT".
	AssignmentStatus PhoneNumberCampaignAssignmentStatus `json:"assignmentStatus"`
	// Brand ID. Empty if the number is associated to a shared campaign.
	BrandID string `json:"brandId"`
	// Extra info about a failure to assign/unassign a number. Relevant only if the
	// assignmentStatus is either FAILED_ASSIGNMENT or FAILED_UNASSIGNMENT
	FailureReasons string `json:"failureReasons"`
	// TCR's alphanumeric ID for the brand.
	TcrBrandID string `json:"tcrBrandId"`
	// TCR's alphanumeric ID for the campaign.
	TcrCampaignID string `json:"tcrCampaignId"`
	// Campaign ID. Empty if the number is associated to a shared campaign.
	TelnyxCampaignID string `json:"telnyxCampaignId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CampaignID       respjson.Field
		CreatedAt        respjson.Field
		PhoneNumber      respjson.Field
		UpdatedAt        respjson.Field
		AssignmentStatus respjson.Field
		BrandID          respjson.Field
		FailureReasons   respjson.Field
		TcrBrandID       respjson.Field
		TcrCampaignID    respjson.Field
		TelnyxCampaignID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberCampaign) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The assignment status of the number.
type PhoneNumberCampaignAssignmentStatus string

const (
	PhoneNumberCampaignAssignmentStatusFailedAssignment    PhoneNumberCampaignAssignmentStatus = "FAILED_ASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusPendingAssignment   PhoneNumberCampaignAssignmentStatus = "PENDING_ASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusAssigned            PhoneNumberCampaignAssignmentStatus = "ASSIGNED"
	PhoneNumberCampaignAssignmentStatusPendingUnassignment PhoneNumberCampaignAssignmentStatus = "PENDING_UNASSIGNMENT"
	PhoneNumberCampaignAssignmentStatusFailedUnassignment  PhoneNumberCampaignAssignmentStatus = "FAILED_UNASSIGNMENT"
)

// The properties CampaignID, PhoneNumber are required.
type PhoneNumberCampaignCreateParam struct {
	// The ID of the campaign you want to link to the specified phone number.
	CampaignID string `json:"campaignId,required"`
	// The phone number you want to link to a specified campaign.
	PhoneNumber string `json:"phoneNumber,required"`
	paramObj
}

func (r PhoneNumberCampaignCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberCampaignCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberCampaignCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
