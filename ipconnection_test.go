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

func TestIPConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IPConnections.New(context.TODO(), telnyx.IPConnectionNewParams{
		Active:                           telnyx.Bool(true),
		AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
		AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
		ConnectionName:                   telnyx.String("string"),
		DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
		DtmfType:                         telnyx.DtmfTypeRfc2833,
		EncodeContactHeaderEnabled:       telnyx.Bool(true),
		EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
		Inbound: telnyx.IPConnectionNewParamsInbound{
			AniNumberFormat:             "+E.164",
			ChannelLimit:                telnyx.Int(10),
			Codecs:                      []string{"string"},
			DefaultRoutingMethod:        "sequential",
			DnisNumberFormat:            "+e164",
			GenerateRingbackTone:        telnyx.Bool(true),
			IsupHeadersEnabled:          telnyx.Bool(true),
			PrackEnabled:                telnyx.Bool(true),
			ShakenStirEnabled:           telnyx.Bool(true),
			SipCompactHeadersEnabled:    telnyx.Bool(true),
			SipRegion:                   "US",
			SipSubdomain:                telnyx.String("test"),
			SipSubdomainReceiveSettings: "only_my_connections",
			Timeout1xxSecs:              telnyx.Int(10),
			Timeout2xxSecs:              telnyx.Int(20),
		},
		IosPushCredentialID:        telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
		OnnetT38PassthroughEnabled: telnyx.Bool(false),
		Outbound: telnyx.OutboundIPParam{
			AniOverride:            telnyx.String("string"),
			AniOverrideType:        telnyx.OutboundIPAniOverrideTypeAlways,
			CallParkingEnabled:     telnyx.Bool(true),
			ChannelLimit:           telnyx.Int(10),
			GenerateRingbackTone:   telnyx.Bool(true),
			InstantRingbackEnabled: telnyx.Bool(true),
			IPAuthenticationMethod: telnyx.OutboundIPIPAuthenticationMethodToken,
			IPAuthenticationToken:  telnyx.String("string"),
			Localization:           telnyx.String("string"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
			T38ReinviteSource:      telnyx.OutboundIPT38ReinviteSourceCustomer,
			TechPrefix:             telnyx.String("string"),
		},
		RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
			CaptureEnabled:      telnyx.Bool(true),
			Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
			ReportFrequencySecs: telnyx.Int(10),
		},
		Tags:                    []string{"tag1", "tag2"},
		TransportProtocol:       telnyx.IPConnectionNewParamsTransportProtocolUdp,
		WebhookAPIVersion:       telnyx.IPConnectionNewParamsWebhookAPIVersion1,
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

func TestIPConnectionGet(t *testing.T) {
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
	_, err := client.IPConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IPConnections.Update(
		context.TODO(),
		"id",
		telnyx.IPConnectionUpdateParams{
			Active:                           telnyx.Bool(true),
			AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
			AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
			ConnectionName:                   telnyx.String("string"),
			DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
			DtmfType:                         telnyx.DtmfTypeRfc2833,
			EncodeContactHeaderEnabled:       telnyx.Bool(true),
			EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
			Inbound: telnyx.InboundIPParam{
				AniNumberFormat:             telnyx.InboundIPAniNumberFormatPlusE164,
				ChannelLimit:                telnyx.Int(10),
				Codecs:                      []string{"string"},
				DefaultPrimaryIPID:          telnyx.String("192.168.0.0"),
				DefaultRoutingMethod:        telnyx.InboundIPDefaultRoutingMethodSequential,
				DefaultSecondaryIPID:        telnyx.String("192.168.0.0"),
				DefaultTertiaryIPID:         telnyx.String("192.168.0.0"),
				DnisNumberFormat:            telnyx.InboundIPDnisNumberFormatPlusE164,
				GenerateRingbackTone:        telnyx.Bool(true),
				IsupHeadersEnabled:          telnyx.Bool(true),
				PrackEnabled:                telnyx.Bool(true),
				ShakenStirEnabled:           telnyx.Bool(true),
				SipCompactHeadersEnabled:    telnyx.Bool(true),
				SipRegion:                   telnyx.InboundIPSipRegionUs,
				SipSubdomain:                telnyx.String("test"),
				SipSubdomainReceiveSettings: telnyx.InboundIPSipSubdomainReceiveSettingsOnlyMyConnections,
				Timeout1xxSecs:              telnyx.Int(10),
				Timeout2xxSecs:              telnyx.Int(20),
			},
			IosPushCredentialID:        telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
			OnnetT38PassthroughEnabled: telnyx.Bool(false),
			Outbound: telnyx.OutboundIPParam{
				AniOverride:            telnyx.String("string"),
				AniOverrideType:        telnyx.OutboundIPAniOverrideTypeAlways,
				CallParkingEnabled:     telnyx.Bool(true),
				ChannelLimit:           telnyx.Int(10),
				GenerateRingbackTone:   telnyx.Bool(true),
				InstantRingbackEnabled: telnyx.Bool(true),
				IPAuthenticationMethod: telnyx.OutboundIPIPAuthenticationMethodToken,
				IPAuthenticationToken:  telnyx.String("string"),
				Localization:           telnyx.String("string"),
				OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
				T38ReinviteSource:      telnyx.OutboundIPT38ReinviteSourceCustomer,
				TechPrefix:             telnyx.String("string"),
			},
			RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
				CaptureEnabled:      telnyx.Bool(true),
				Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
				ReportFrequencySecs: telnyx.Int(10),
			},
			Tags:                    []string{"tag1", "tag2"},
			TransportProtocol:       telnyx.IPConnectionUpdateParamsTransportProtocolUdp,
			WebhookAPIVersion:       telnyx.IPConnectionUpdateParamsWebhookAPIVersion1,
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

func TestIPConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.IPConnections.List(context.TODO(), telnyx.IPConnectionListParams{
		Filter: telnyx.IPConnectionListParamsFilter{
			ConnectionName: telnyx.IPConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("contains"),
			},
			Fqdn:                   telnyx.String("fqdn"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		Page: telnyx.IPConnectionListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.IPConnectionListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIPConnectionDelete(t *testing.T) {
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
	_, err := client.IPConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
