// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.AccessIPAddress.List(context.TODO(), telnyx.AccessIPAddressListParams{
		PageNumber: telnyx.Int(1),
		PageSize:   telnyx.Int(50),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, accessIPAddress := range page.Data {
		t.Logf("%+v\n", accessIPAddress.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, accessIPAddress := range page.Data {
			t.Logf("%+v\n", accessIPAddress.ID)
		}
	}
}
