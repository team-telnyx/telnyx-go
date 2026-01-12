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

func TestMobileVoiceConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MobileVoiceConnections.New(context.TODO(), telnyx.MobileVoiceConnectionNewParams{
		Active:         telnyx.Bool(true),
		ConnectionName: telnyx.String("connection_name"),
		Inbound: telnyx.MobileVoiceConnectionNewParamsInbound{
			ChannelLimit: telnyx.Int(0),
		},
		Outbound: telnyx.MobileVoiceConnectionNewParamsOutbound{
			ChannelLimit:           telnyx.Int(0),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		Tags:                    []string{"string"},
		WebhookAPIVersion:       telnyx.MobileVoiceConnectionNewParamsWebhookAPIVersionV1,
		WebhookEventFailoverURL: telnyx.String("webhook_event_failover_url"),
		WebhookEventURL:         telnyx.String("webhook_event_url"),
		WebhookTimeoutSecs:      telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMobileVoiceConnectionGet(t *testing.T) {
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
	_, err := client.MobileVoiceConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMobileVoiceConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MobileVoiceConnections.Update(
		context.TODO(),
		"id",
		telnyx.MobileVoiceConnectionUpdateParams{
			Active:         telnyx.Bool(true),
			ConnectionName: telnyx.String("connection_name"),
			Inbound: telnyx.MobileVoiceConnectionUpdateParamsInbound{
				ChannelLimit: telnyx.Int(0),
			},
			Outbound: telnyx.MobileVoiceConnectionUpdateParamsOutbound{
				ChannelLimit:           telnyx.Int(0),
				OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
			},
			Tags:                    []string{"string"},
			WebhookAPIVersion:       telnyx.MobileVoiceConnectionUpdateParamsWebhookAPIVersionV1,
			WebhookEventFailoverURL: telnyx.String("webhook_event_failover_url"),
			WebhookEventURL:         telnyx.String("webhook_event_url"),
			WebhookTimeoutSecs:      telnyx.Int(0),
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

func TestMobileVoiceConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.MobileVoiceConnections.List(context.TODO(), telnyx.MobileVoiceConnectionListParams{
		FilterConnectionNameContains: telnyx.String("filter[connection_name][contains]"),
		PageNumber:                   telnyx.Int(0),
		PageSize:                     telnyx.Int(0),
		Sort:                         telnyx.String("sort"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMobileVoiceConnectionDelete(t *testing.T) {
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
	_, err := client.MobileVoiceConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
