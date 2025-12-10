// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// BrandExternalVettingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandExternalVettingService] method instead.
type BrandExternalVettingService struct {
	Options []option.RequestOption
}

// NewBrandExternalVettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrandExternalVettingService(opts ...option.RequestOption) (r BrandExternalVettingService) {
	r = BrandExternalVettingService{}
	r.Options = opts
	return
}
