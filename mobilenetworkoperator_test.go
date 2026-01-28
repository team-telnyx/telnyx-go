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

func TestMobileNetworkOperatorListWithOptionalParams(t *testing.T) {
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
	_, err := client.MobileNetworkOperators.List(context.TODO(), telnyx.MobileNetworkOperatorListParams{
		Filter: telnyx.MobileNetworkOperatorListParamsFilter{
			CountryCode: telnyx.String("US"),
			Mcc:         telnyx.String("310"),
			Mnc:         telnyx.String("410"),
			Name: telnyx.MobileNetworkOperatorListParamsFilterName{
				Contains:   telnyx.String("T&T"),
				EndsWith:   telnyx.String("T"),
				StartsWith: telnyx.String("AT"),
			},
			NetworkPreferencesEnabled: telnyx.Bool(true),
			Tadig:                     telnyx.String("USACG"),
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
