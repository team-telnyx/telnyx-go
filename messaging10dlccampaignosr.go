// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// Messaging10dlcCampaignOsrService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessaging10dlcCampaignOsrService] method instead.
type Messaging10dlcCampaignOsrService struct {
	Options []option.RequestOption
}

// NewMessaging10dlcCampaignOsrService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessaging10dlcCampaignOsrService(opts ...option.RequestOption) (r Messaging10dlcCampaignOsrService) {
	r = Messaging10dlcCampaignOsrService{}
	r.Options = opts
	return
}

// Get OSR campaign attributes
func (r *Messaging10dlcCampaignOsrService) GetAttributes(ctx context.Context, campaignID string, opts ...option.RequestOption) (res *Messaging10dlcCampaignOsrGetAttributesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if campaignID == "" {
		err = errors.New("missing required campaignId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/campaign/%s/osr/attributes", campaignID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Messaging10dlcCampaignOsrGetAttributesResponse map[string]any
