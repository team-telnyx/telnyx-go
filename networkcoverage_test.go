// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestNetworkCoverageListWithOptionalParams(t *testing.T) {
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
	_, err := client.NetworkCoverage.List(context.TODO(), telnyx.NetworkCoverageListParams{
		Filter: telnyx.NetworkCoverageListParamsFilter{
			LocationCode:   telnyx.String("silicon_valley-ca"),
			LocationPop:    telnyx.String("SV1"),
			LocationRegion: telnyx.String("AMER"),
			LocationSite:   telnyx.String("SJC"),
		},
		Filters: telnyx.NetworkCoverageListParamsFilters{
			AvailableServices: telnyx.NetworkCoverageListParamsFiltersAvailableServicesUnion{
				OfAvailableService: telnyx.Opt(telnyx.AvailableServiceCloudVpn),
			},
		},
		Page: telnyx.NetworkCoverageListParamsPage{
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
