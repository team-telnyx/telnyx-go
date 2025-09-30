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

func TestFaxApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FaxApplications.New(context.TODO(), telnyx.FaxApplicationNewParams{
		ApplicationName:    "fax-router",
		WebhookEventURL:    "https://example.com",
		Active:             telnyx.Bool(false),
		AnchorsiteOverride: telnyx.AnchorsiteOverrideAmsterdamNetherlands,
		Inbound: telnyx.FaxApplicationNewParamsInbound{
			ChannelLimit:                telnyx.Int(10),
			SipSubdomain:                telnyx.String("example"),
			SipSubdomainReceiveSettings: "only_my_connections",
		},
		Outbound: telnyx.FaxApplicationNewParamsOutbound{
			ChannelLimit:           telnyx.Int(10),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		Tags:                    []string{"tag1", "tag2"},
		WebhookEventFailoverURL: telnyx.String("https://failover.example.com"),
		WebhookTimeoutSecs:      telnyx.Int(25),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxApplicationGet(t *testing.T) {
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
	_, err := client.FaxApplications.Get(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.FaxApplications.Update(
		context.TODO(),
		"1293384261075731499",
		telnyx.FaxApplicationUpdateParams{
			ApplicationName:    "fax-router",
			WebhookEventURL:    "https://example.com",
			Active:             telnyx.Bool(false),
			AnchorsiteOverride: telnyx.AnchorsiteOverrideAmsterdamNetherlands,
			FaxEmailRecipient:  telnyx.String("user@example.com"),
			Inbound: telnyx.FaxApplicationUpdateParamsInbound{
				ChannelLimit:                telnyx.Int(10),
				SipSubdomain:                telnyx.String("example"),
				SipSubdomainReceiveSettings: "only_my_connections",
			},
			Outbound: telnyx.FaxApplicationUpdateParamsOutbound{
				ChannelLimit:           telnyx.Int(10),
				OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			},
			Tags:                    []string{"tag1", "tag2"},
			WebhookEventFailoverURL: telnyx.String("https://failover.example.com"),
			WebhookTimeoutSecs:      telnyx.Int(25),
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

func TestFaxApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.FaxApplications.List(context.TODO(), telnyx.FaxApplicationListParams{
		Filter: telnyx.FaxApplicationListParamsFilter{
			ApplicationName: telnyx.FaxApplicationListParamsFilterApplicationName{
				Contains: telnyx.String("fax-app"),
			},
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		Page: telnyx.FaxApplicationListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.FaxApplicationListParamsSortApplicationName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxApplicationDelete(t *testing.T) {
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
	_, err := client.FaxApplications.Delete(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
