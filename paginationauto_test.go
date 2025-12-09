// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.AccessIPAddress.ListAutoPaging(context.TODO(), telnyx.AccessIPAddressListParams{
		PageNumber: telnyx.Int(1),
		PageSize:   telnyx.Int(50),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		accessIPAddress := iter.Current()
		t.Logf("%+v\n", accessIPAddress.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
