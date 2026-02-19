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

func TestCallControlApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CallControlApplications.New(context.TODO(), telnyx.CallControlApplicationNewParams{
		ApplicationName:         "call-router",
		WebhookEventURL:         "https://example.com",
		Active:                  telnyx.Bool(false),
		AnchorsiteOverride:      telnyx.CallControlApplicationNewParamsAnchorsiteOverrideLatency,
		CallCostInWebhooks:      telnyx.Bool(true),
		DtmfType:                telnyx.CallControlApplicationNewParamsDtmfTypeInband,
		FirstCommandTimeout:     telnyx.Bool(true),
		FirstCommandTimeoutSecs: telnyx.Int(10),
		Inbound: telnyx.CallControlApplicationInboundParam{
			ChannelLimit:                telnyx.Int(10),
			ShakenStirEnabled:           telnyx.Bool(true),
			SipSubdomain:                telnyx.String("example"),
			SipSubdomainReceiveSettings: telnyx.CallControlApplicationInboundSipSubdomainReceiveSettingsOnlyMyConnections,
		},
		Outbound: telnyx.CallControlApplicationOutboundParam{
			ChannelLimit:           telnyx.Int(10),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		RedactDtmfDebugLogging:  telnyx.Bool(true),
		WebhookAPIVersion:       telnyx.CallControlApplicationNewParamsWebhookAPIVersionV1,
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

func TestCallControlApplicationGet(t *testing.T) {
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
	_, err := client.CallControlApplications.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCallControlApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CallControlApplications.Update(
		context.TODO(),
		"id",
		telnyx.CallControlApplicationUpdateParams{
			ApplicationName:         "call-router",
			WebhookEventURL:         "https://example.com",
			Active:                  telnyx.Bool(false),
			AnchorsiteOverride:      telnyx.CallControlApplicationUpdateParamsAnchorsiteOverrideLatency,
			CallCostInWebhooks:      telnyx.Bool(true),
			DtmfType:                telnyx.CallControlApplicationUpdateParamsDtmfTypeInband,
			FirstCommandTimeout:     telnyx.Bool(true),
			FirstCommandTimeoutSecs: telnyx.Int(10),
			Inbound: telnyx.CallControlApplicationInboundParam{
				ChannelLimit:                telnyx.Int(10),
				ShakenStirEnabled:           telnyx.Bool(true),
				SipSubdomain:                telnyx.String("example"),
				SipSubdomainReceiveSettings: telnyx.CallControlApplicationInboundSipSubdomainReceiveSettingsOnlyMyConnections,
			},
			Outbound: telnyx.CallControlApplicationOutboundParam{
				ChannelLimit:           telnyx.Int(10),
				OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			},
			RedactDtmfDebugLogging:  telnyx.Bool(true),
			Tags:                    []string{"tag1", "tag2"},
			WebhookAPIVersion:       telnyx.CallControlApplicationUpdateParamsWebhookAPIVersionV1,
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

func TestCallControlApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.CallControlApplications.List(context.TODO(), telnyx.CallControlApplicationListParams{
		Filter: telnyx.CallControlApplicationListParamsFilter{
			ApplicationName: telnyx.CallControlApplicationListParamsFilterApplicationName{
				Contains: telnyx.String("contains"),
			},
			ApplicationSessionID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ConnectionID:         telnyx.String("connection_id"),
			Failed:               telnyx.Bool(false),
			From:                 telnyx.String("+12025550142"),
			LegID:                telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                 telnyx.String("name"),
			OccurredAt: telnyx.CallControlApplicationListParamsFilterOccurredAt{
				Eq:  telnyx.String("2019-03-29T11:10:00Z"),
				Gt:  telnyx.String("2019-03-29T11:10:00Z"),
				Gte: telnyx.String("2019-03-29T11:10:00Z"),
				Lt:  telnyx.String("2019-03-29T11:10:00Z"),
				Lte: telnyx.String("2019-03-29T11:10:00Z"),
			},
			OutboundOutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			Product:                        "texml",
			Status:                         "init",
			To:                             telnyx.String("+12025550142"),
			Type:                           "webhook",
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.CallControlApplicationListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCallControlApplicationDelete(t *testing.T) {
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
	_, err := client.CallControlApplications.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
