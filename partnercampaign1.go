// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// PartnerCampaignService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPartnerCampaignService] method instead.
type PartnerCampaignService struct {
	Options []option.RequestOption
}

// NewPartnerCampaignService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPartnerCampaignService(opts ...option.RequestOption) (r PartnerCampaignService) {
	r = PartnerCampaignService{}
	r.Options = opts
	return
}
