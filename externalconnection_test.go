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

func TestExternalConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.New(context.TODO(), telnyx.ExternalConnectionNewParams{
		ExternalSipConnection: telnyx.ExternalConnectionNewParamsExternalSipConnectionZoom,
		Outbound: telnyx.ExternalConnectionNewParamsOutbound{
			ChannelLimit:           telnyx.Int(10),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		Active: telnyx.Bool(false),
		Inbound: telnyx.ExternalConnectionNewParamsInbound{
			OutboundVoiceProfileID: "12345678-1234-1234-1234-123456789012",
			ChannelLimit:           telnyx.Int(10),
		},
		Tags:                    []string{"tag1", "tag2"},
		WebhookEventFailoverURL: telnyx.String("https://failover.example.com"),
		WebhookEventURL:         telnyx.String("https://example.com"),
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

func TestExternalConnectionGet(t *testing.T) {
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
	_, err := client.ExternalConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.Update(
		context.TODO(),
		"id",
		telnyx.ExternalConnectionUpdateParams{
			Outbound: telnyx.ExternalConnectionUpdateParamsOutbound{
				OutboundVoiceProfileID: "outbound_voice_profile_id",
				ChannelLimit:           telnyx.Int(10),
			},
			Active: telnyx.Bool(false),
			Inbound: telnyx.ExternalConnectionUpdateParamsInbound{
				ChannelLimit: telnyx.Int(10),
			},
			Tags:                    []string{"tag1", "tag2"},
			WebhookEventFailoverURL: telnyx.String("https://failover.example.com"),
			WebhookEventURL:         telnyx.String("https://example.com"),
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

func TestExternalConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.List(context.TODO(), telnyx.ExternalConnectionListParams{
		Filter: telnyx.ExternalConnectionListParamsFilter{
			ID: telnyx.String("1930241863466354012"),
			ConnectionName: telnyx.ExternalConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("My Connection"),
			},
			CreatedAt:             telnyx.String("2022-12-31"),
			ExternalSipConnection: "zoom",
			PhoneNumber: telnyx.ExternalConnectionListParamsFilterPhoneNumber{
				Contains: telnyx.String("+15555555555"),
			},
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

func TestExternalConnectionDelete(t *testing.T) {
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
	_, err := client.ExternalConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExternalConnectionUpdateLocation(t *testing.T) {
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
	_, err := client.ExternalConnections.UpdateLocation(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.ExternalConnectionUpdateLocationParams{
			ID:                       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			StaticEmergencyAddressID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
