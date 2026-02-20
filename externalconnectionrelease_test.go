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

func TestExternalConnectionReleaseGet(t *testing.T) {
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
	_, err := client.ExternalConnections.Releases.Get(
		context.TODO(),
		"7b6a6449-b055-45a6-81f6-f6f0dffa4cc6",
		telnyx.ExternalConnectionReleaseGetParams{
			ID: "1293384261075731499",
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

func TestExternalConnectionReleaseListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.Releases.List(
		context.TODO(),
		"1293384261075731499",
		telnyx.ExternalConnectionReleaseListParams{
			Filter: telnyx.ExternalConnectionReleaseListParamsFilter{
				CivicAddressID: telnyx.ExternalConnectionReleaseListParamsFilterCivicAddressID{
					Eq: telnyx.String("19990261512338516954"),
				},
				LocationID: telnyx.ExternalConnectionReleaseListParamsFilterLocationID{
					Eq: telnyx.String("19995665508264022121"),
				},
				PhoneNumber: telnyx.ExternalConnectionReleaseListParamsFilterPhoneNumber{
					Contains: telnyx.String("+123"),
					Eq:       telnyx.String("+1234567890"),
				},
				Status: telnyx.ExternalConnectionReleaseListParamsFilterStatus{
					Eq: []string{"pending", "in_progress"},
				},
			},
			PageNumber: telnyx.Int(0),
			PageSize:   telnyx.Int(0),
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
