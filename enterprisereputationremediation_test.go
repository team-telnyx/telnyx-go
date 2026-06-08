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

func TestEnterpriseReputationRemediationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.Reputation.Remediation.New(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseReputationRemediationNewParams{
			CallPurpose:  "Appointment reminders for our dental clinic.",
			ContactEmail: "ops@example.com",
			PhoneNumbers: []string{"+19493253498", "+12134445566"},
			WebhookURL:   telnyx.String("https://example.com/webhooks/remediation"),
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

func TestEnterpriseReputationRemediationGet(t *testing.T) {
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
	_, err := client.Enterprises.Reputation.Remediation.Get(
		context.TODO(),
		"b7c1f1c0-7a9d-4f0a-9d3e-2f6a1c4b8e21",
		telnyx.EnterpriseReputationRemediationGetParams{
			EnterpriseID: "4a6192a4-573d-446d-b3ce-aff9117272a6",
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

func TestEnterpriseReputationRemediationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.Reputation.Remediation.List(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseReputationRemediationListParams{
			FilterCreatedAtGte: telnyx.Time(time.Now()),
			FilterCreatedAtLte: telnyx.Time(time.Now()),
			FilterStatus:       telnyx.EnterpriseReputationRemediationListParamsFilterStatusInProgress,
			PageNumber:         telnyx.Int(1),
			PageSize:           telnyx.Int(20),
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
