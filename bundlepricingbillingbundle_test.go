// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/internal/testutil"
	"github.com/team-telnyx/telnyx-go/option"
)

func TestBundlePricingBillingBundleGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.BundlePricing.BillingBundles.Get(
		context.TODO(),
		"8661948c-a386-4385-837f-af00f40f111a",
		telnyx.BundlePricingBillingBundleGetParams{
			AuthorizationBearer: telnyx.String("authorization_bearer"),
		},
	)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBundlePricingBillingBundleListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.BundlePricing.BillingBundles.List(context.TODO(), telnyx.BundlePricingBillingBundleListParams{
		Filter: telnyx.BundlePricingBillingBundleListParamsFilter{
			CountryISO: []string{"US"},
			Resource:   []string{"+15617819942"},
		},
		Page: telnyx.BundlePricingBillingBundleListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		AuthorizationBearer: telnyx.String("authorization_bearer"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
