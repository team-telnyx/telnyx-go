// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestEnterpriseDirNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.Dir.New(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseDirNewParams{
			AuthorizerEmail:        "sam@acmeplumbing.example.com",
			AuthorizerName:         "Sam Owner",
			CertifyBrandIsAccurate: true,
			CertifyIPOwnership:     true,
			CertifyNoShaftContent:  true,
			DisplayName:            "Acme Plumbing",
			CallReasons:            []string{"Appointment reminders", "Billing inquiries"},
			Documents: []telnyx.EnterpriseDirNewParamsDocument{{
				DocumentID:   "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
				DocumentType: "business_registration",
				Description:  telnyx.String("Certificate of incorporation."),
			}},
			LogoURL:   telnyx.String("https://acmeplumbing.example.com/logo-256.bmp"),
			Reselling: telnyx.Bool(false),
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

func TestEnterpriseDirListWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.Dir.List(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseDirListParams{
			FilterCallReasonContains:  telnyx.String("filter[call_reason][contains]"),
			FilterDisplayNameContains: telnyx.String("filter[display_name][contains]"),
			FilterExpiringAtGte:       telnyx.Time(time.Now()),
			FilterExpiringAtLte:       telnyx.Time(time.Now()),
			FilterExpiringWithinDays:  telnyx.Int(1),
			FilterStatus:              telnyx.EnterpriseDirListParamsFilterStatusDraft,
			PageNumber:                telnyx.Int(1),
			PageSize:                  telnyx.Int(20),
			Sort:                      telnyx.EnterpriseDirListParamsSortCreatedAt,
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
