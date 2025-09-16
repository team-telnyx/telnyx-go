// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestVirtualCrossConnectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualCrossConnects.New(context.TODO(), telnyx.VirtualCrossConnectNewParams{
		BgpAsn:                  1234,
		CloudProvider:           telnyx.VirtualCrossConnectNewParamsCloudProviderAws,
		CloudProviderRegion:     "us-east-1",
		NetworkID:               "6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		PrimaryCloudAccountID:   "123456789012",
		RegionCode:              "ashburn-va",
		BandwidthMbps:           telnyx.Float(50),
		Name:                    telnyx.String("test interface"),
		PrimaryBgpKey:           telnyx.String("yFV4wEPtPVPfDUGLWiyQzwga"),
		PrimaryCloudIP:          telnyx.String("169.254.0.2"),
		PrimaryTelnyxIP:         telnyx.String("169.254.0.1"),
		SecondaryBgpKey:         telnyx.String("ge1lONeK9RcA83uuWaw9DvZy"),
		SecondaryCloudAccountID: telnyx.String(""),
		SecondaryCloudIP:        telnyx.String("169.254.0.4"),
		SecondaryTelnyxIP:       telnyx.String("169.254.0.3"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualCrossConnectGet(t *testing.T) {
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
	_, err := client.VirtualCrossConnects.Get(context.TODO(), "6a09cdc3-8948-47f0-aa62-74ac943d6c58")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualCrossConnectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualCrossConnects.Update(
		context.TODO(),
		"6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		telnyx.VirtualCrossConnectUpdateParams{
			PrimaryCloudIP:               telnyx.String("169.254.0.2"),
			PrimaryEnabled:               telnyx.Bool(true),
			PrimaryRoutingAnnouncement:   telnyx.Bool(false),
			SecondaryCloudIP:             telnyx.String("169.254.0.4"),
			SecondaryEnabled:             telnyx.Bool(true),
			SecondaryRoutingAnnouncement: telnyx.Bool(false),
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

func TestVirtualCrossConnectListWithOptionalParams(t *testing.T) {
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
	_, err := client.VirtualCrossConnects.List(context.TODO(), telnyx.VirtualCrossConnectListParams{
		Filter: telnyx.VirtualCrossConnectListParamsFilter{
			NetworkID: telnyx.String("6a09cdc3-8948-47f0-aa62-74ac943d6c58"),
		},
		Page: telnyx.VirtualCrossConnectListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVirtualCrossConnectDelete(t *testing.T) {
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
	_, err := client.VirtualCrossConnects.Delete(context.TODO(), "6a09cdc3-8948-47f0-aa62-74ac943d6c58")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
