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

func TestReportListMdrsWithOptionalParams(t *testing.T) {
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
	_, err := client.Reports.ListMdrs(context.TODO(), telnyx.ReportListMdrsParams{
		ID:          telnyx.String("e093fbe0-5bde-11eb-ae93-0242ac130002"),
		Cld:         telnyx.String("+15551237654"),
		Cli:         telnyx.String("+15551237654"),
		Direction:   telnyx.ReportListMdrsParamsDirectionInbound,
		EndDate:     telnyx.String("end_date"),
		MessageType: telnyx.ReportListMdrsParamsMessageTypeSMS,
		Profile:     telnyx.String("My profile"),
		StartDate:   telnyx.String("start_date"),
		Status:      telnyx.ReportListMdrsParamsStatusDelivered,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReportListWdrsWithOptionalParams(t *testing.T) {
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
	_, err := client.Reports.ListWdrs(context.TODO(), telnyx.ReportListWdrsParams{
		ID:      telnyx.String("e093fbe0-5bde-11eb-ae93-0242ac130002"),
		EndDate: telnyx.String("2021-06-01T00:00:00Z"),
		Imsi:    telnyx.String("123456"),
		Mcc:     telnyx.String("204"),
		Mnc:     telnyx.String("01"),
		Page: telnyx.ReportListWdrsParamsPage{
			Number: telnyx.Int(0),
			Size:   telnyx.Int(0),
		},
		PhoneNumber:  telnyx.String("+12345678910"),
		SimCardID:    telnyx.String("877f80a6-e5b2-4687-9a04-88076265720f"),
		SimGroupID:   telnyx.String("f05a189f-7c46-4531-ac56-1460dc465a42"),
		SimGroupName: telnyx.String("sim name"),
		Sort:         []string{"created_at"},
		StartDate:    telnyx.String("2021-05-01T00:00:00Z"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
