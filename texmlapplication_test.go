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

func TestTexmlApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TexmlApplications.New(context.TODO(), telnyx.TexmlApplicationNewParams{
		FriendlyName:            "call-router",
		VoiceURL:                "https://example.com",
		Active:                  telnyx.Bool(false),
		AnchorsiteOverride:      telnyx.AnchorsiteOverrideAmsterdamNetherlands,
		CallCostInWebhooks:      telnyx.Bool(false),
		DtmfType:                telnyx.DtmfTypeInband,
		FirstCommandTimeout:     telnyx.Bool(true),
		FirstCommandTimeoutSecs: telnyx.Int(10),
		Inbound: telnyx.TexmlApplicationNewParamsInbound{
			ChannelLimit:                telnyx.Int(10),
			ShakenStirEnabled:           telnyx.Bool(true),
			SipSubdomain:                telnyx.String("example"),
			SipSubdomainReceiveSettings: "only_my_connections",
		},
		Outbound: telnyx.TexmlApplicationNewParamsOutbound{
			ChannelLimit:           telnyx.Int(10),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		StatusCallback:       telnyx.String("https://example.com"),
		StatusCallbackMethod: telnyx.TexmlApplicationNewParamsStatusCallbackMethodGet,
		Tags:                 []string{"tag1", "tag2"},
		VoiceFallbackURL:     telnyx.String("https://fallback.example.com"),
		VoiceMethod:          telnyx.TexmlApplicationNewParamsVoiceMethodGet,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTexmlApplicationGet(t *testing.T) {
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
	_, err := client.TexmlApplications.Get(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTexmlApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.TexmlApplications.Update(
		context.TODO(),
		"1293384261075731499",
		telnyx.TexmlApplicationUpdateParams{
			FriendlyName:            "call-router",
			VoiceURL:                "https://example.com",
			Active:                  telnyx.Bool(false),
			AnchorsiteOverride:      telnyx.AnchorsiteOverrideAmsterdamNetherlands,
			CallCostInWebhooks:      telnyx.Bool(false),
			DtmfType:                telnyx.DtmfTypeInband,
			FirstCommandTimeout:     telnyx.Bool(true),
			FirstCommandTimeoutSecs: telnyx.Int(10),
			Inbound: telnyx.TexmlApplicationUpdateParamsInbound{
				ChannelLimit:                telnyx.Int(10),
				ShakenStirEnabled:           telnyx.Bool(true),
				SipSubdomain:                telnyx.String("example"),
				SipSubdomainReceiveSettings: "only_my_connections",
			},
			Outbound: telnyx.TexmlApplicationUpdateParamsOutbound{
				ChannelLimit:           telnyx.Int(10),
				OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			},
			StatusCallback:       telnyx.String("https://example.com"),
			StatusCallbackMethod: telnyx.TexmlApplicationUpdateParamsStatusCallbackMethodGet,
			Tags:                 []string{"tag1", "tag2"},
			VoiceFallbackURL:     telnyx.String("https://fallback.example.com"),
			VoiceMethod:          telnyx.TexmlApplicationUpdateParamsVoiceMethodGet,
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

func TestTexmlApplicationListWithOptionalParams(t *testing.T) {
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
	_, err := client.TexmlApplications.List(context.TODO(), telnyx.TexmlApplicationListParams{
		Filter: telnyx.TexmlApplicationListParamsFilter{
			FriendlyName:           telnyx.String("friendly_name"),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		Page: telnyx.TexmlApplicationListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.TexmlApplicationListParamsSortFriendlyName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTexmlApplicationDelete(t *testing.T) {
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
	_, err := client.TexmlApplications.Delete(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
