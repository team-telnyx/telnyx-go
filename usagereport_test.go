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

func TestUsageReportListWithOptionalParams(t *testing.T) {
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
	_, err := client.UsageReports.List(context.TODO(), telnyx.UsageReportListParams{
		Dimensions:      []string{"string"},
		Metrics:         []string{"string"},
		Product:         "product",
		DateRange:       telnyx.String("date_range"),
		EndDate:         telnyx.String("end_date"),
		Filter:          telnyx.String("filter"),
		Format:          telnyx.UsageReportListParamsFormatCsv,
		ManagedAccounts: telnyx.Bool(true),
		Page: telnyx.UsageReportListParamsPage{
			Number: telnyx.Int(2),
			Size:   telnyx.Int(5000),
		},
		Sort:                []string{"string"},
		StartDate:           telnyx.String("start_date"),
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

func TestUsageReportGetOptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.UsageReports.GetOptions(context.TODO(), telnyx.UsageReportGetOptionsParams{
		Product:             telnyx.String("product"),
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
