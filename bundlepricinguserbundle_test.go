// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestBundlePricingUserBundleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.New(context.TODO(), telnyx.BundlePricingUserBundleNewParams{
		IdempotencyKey: telnyx.String("12ade33a-21c0-473b-b055-b3c836e1c292"),
		Items: []telnyx.BundlePricingUserBundleNewParamsItem{{
			BillingBundleID: "12ade33a-21c0-473b-b055-b3c836e1c292",
			Quantity:        0,
		}},
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

func TestBundlePricingUserBundleGetWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.Get(
		context.TODO(),
		"ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		telnyx.BundlePricingUserBundleGetParams{
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

func TestBundlePricingUserBundleListWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.List(context.TODO(), telnyx.BundlePricingUserBundleListParams{
		Filter: telnyx.BundlePricingUserBundleListParamsFilter{
			CountryISO: []string{"US"},
			Resource:   []string{"+15617819942"},
		},
		Page: telnyx.BundlePricingUserBundleListParamsPage{
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

func TestBundlePricingUserBundleDeactivateWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.Deactivate(
		context.TODO(),
		"ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		telnyx.BundlePricingUserBundleDeactivateParams{
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

func TestBundlePricingUserBundleListResourcesWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.ListResources(
		context.TODO(),
		"ca1d2263-d1f1-43ac-ba53-248e7a4bb26a",
		telnyx.BundlePricingUserBundleListResourcesParams{
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

func TestBundlePricingUserBundleListUnusedWithOptionalParams(t *testing.T) {
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
	_, err := client.BundlePricing.UserBundles.ListUnused(context.TODO(), telnyx.BundlePricingUserBundleListUnusedParams{
		Filter: telnyx.BundlePricingUserBundleListUnusedParamsFilter{
			CountryISO: []string{"US"},
			Resource:   []string{"+15617819942"},
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
