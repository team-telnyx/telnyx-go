// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// CampaignService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignService] method instead.
type CampaignService struct {
	Options []option.RequestOption
	Usecase CampaignUsecaseService
	Osr     CampaignOsrService
}

// NewCampaignService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignService(opts ...option.RequestOption) (r CampaignService) {
	r = CampaignService{}
	r.Options = opts
	r.Usecase = NewCampaignUsecaseService(opts...)
	r.Osr = NewCampaignOsrService(opts...)
	return
}
