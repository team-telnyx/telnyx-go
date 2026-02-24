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

func TestVirtualCrossConnectsCoverageListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.VirtualCrossConnectsCoverage.List(context.TODO(), telnyx.VirtualCrossConnectsCoverageListParams{
		Filter: telnyx.VirtualCrossConnectsCoverageListParamsFilter{
			CloudProvider:       "aws",
			CloudProviderRegion: telnyx.String("us-east-1"),
			LocationCode:        telnyx.String("silicon_valley-ca"),
			LocationPop:         telnyx.String("SV1"),
			LocationRegion:      telnyx.String("AMER"),
			LocationSite:        telnyx.String("SJC"),
		},
		Filters: telnyx.VirtualCrossConnectsCoverageListParamsFilters{
			AvailableBandwidth: telnyx.VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthUnion{
				OfInt: telnyx.Int(0),
			},
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
