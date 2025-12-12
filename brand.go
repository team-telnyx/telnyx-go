// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// BrandService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options         []option.RequestOption
	ExternalVetting BrandExternalVettingService
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.Options = opts
	r.ExternalVetting = NewBrandExternalVettingService(opts...)
	return
}
