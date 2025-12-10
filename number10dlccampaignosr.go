// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// Number10dlcCampaignOsrService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcCampaignOsrService] method instead.
type Number10dlcCampaignOsrService struct {
	Options []option.RequestOption
}

// NewNumber10dlcCampaignOsrService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumber10dlcCampaignOsrService(opts ...option.RequestOption) (r Number10dlcCampaignOsrService) {
	r = Number10dlcCampaignOsrService{}
	r.Options = opts
	return
}

// Get My Osr Campaign Attributes
func (r *Number10dlcCampaignOsrService) GetAttributes(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Number10dlcCampaignOsrGetAttributesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/campaign/%s/osr/attributes", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Number10dlcCampaignOsrGetAttributesResponse map[string]any
