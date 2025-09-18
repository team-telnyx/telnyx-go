// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
)

// CampaignOsrService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignOsrService] method instead.
type CampaignOsrService struct {
	Options []option.RequestOption
}

// NewCampaignOsrService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignOsrService(opts ...option.RequestOption) (r CampaignOsrService) {
	r = CampaignOsrService{}
	r.Options = opts
	return
}

// Get My Osr Campaign Attributes
func (r *CampaignOsrService) GetAttributes(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *CampaignOsrGetAttributesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("campaign/%s/osr/attributes", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CampaignOsrGetAttributesResponse = any
