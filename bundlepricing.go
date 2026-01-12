// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// BundlePricingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBundlePricingService] method instead.
type BundlePricingService struct {
	Options        []option.RequestOption
	BillingBundles BundlePricingBillingBundleService
	UserBundles    BundlePricingUserBundleService
}

// NewBundlePricingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBundlePricingService(opts ...option.RequestOption) (r BundlePricingService) {
	r = BundlePricingService{}
	r.Options = opts
	r.BillingBundles = NewBundlePricingBillingBundleService(opts...)
	r.UserBundles = NewBundlePricingUserBundleService(opts...)
	return
}
