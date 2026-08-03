// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// PricingService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingService] method instead.
type PricingService struct {
	Options []option.RequestOption
	// Public pricing operations
	Products PricingProductService
}

// NewPricingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPricingService(opts ...option.RequestOption) (r PricingService) {
	r = PricingService{}
	r.Options = opts
	r.Products = NewPricingProductService(opts...)
	return
}
