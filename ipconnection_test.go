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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

func TestIPConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IPConnections.New(context.TODO(), telnyx.IPConnectionNewParams{
		Active:                           telnyx.Bool(true),
		AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
		AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
		CallCostInWebhooks:               telnyx.Bool(false),
		ConnectionName:                   telnyx.String("string"),
		DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
		DtmfType:                         telnyx.DtmfTypeRfc2833,
		EncodeContactHeaderEnabled:       telnyx.Bool(true),
		EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
		Inbound: telnyx.IPConnectionNewParamsInbound{
			AniNumberFormat:             "+E.164",
			ChannelLimit:                telnyx.Int(10),
			Codecs:                      []string{"G722"},
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
		IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
		JitterBuffer: shared.ConnectionJitterBufferParam{
			EnableJitterBuffer:  telnyx.Bool(true),
			JitterbufferMsecMax: telnyx.Int(200),
			JitterbufferMsecMin: telnyx.Int(60),
		},
		NoiseSuppression: telnyx.IPConnectionNewParamsNoiseSuppressionBoth,
		NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
			AttenuationLimit: telnyx.Int(80),
			Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
		},
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
		WebhookAPIVersion:       telnyx.IPConnectionNewParamsWebhookAPIVersionV1,
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
	_, err := client.IPConnections.Update(
		context.TODO(),
		"id",
		telnyx.IPConnectionUpdateParams{
			Active:                           telnyx.Bool(true),
			AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
			AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
			CallCostInWebhooks:               telnyx.Bool(false),
			ConnectionName:                   telnyx.String("string"),
			DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
			DtmfType:                         telnyx.DtmfTypeRfc2833,
			EncodeContactHeaderEnabled:       telnyx.Bool(true),
			EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
			Inbound: telnyx.InboundIPParam{
				AniNumberFormat:             telnyx.InboundIPAniNumberFormatPlusE164,
				ChannelLimit:                telnyx.Int(10),
				Codecs:                      []string{"G722"},
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
			IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
			JitterBuffer: shared.ConnectionJitterBufferParam{
				EnableJitterBuffer:  telnyx.Bool(true),
				JitterbufferMsecMax: telnyx.Int(200),
				JitterbufferMsecMin: telnyx.Int(60),
			},
			NoiseSuppression: telnyx.IPConnectionUpdateParamsNoiseSuppressionBoth,
			NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
				AttenuationLimit: telnyx.Int(80),
				Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
			},
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
			WebhookAPIVersion:       telnyx.IPConnectionUpdateParamsWebhookAPIVersionV1,
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
	_, err := client.IPConnections.List(context.TODO(), telnyx.IPConnectionListParams{
		Filter: telnyx.IPConnectionListParamsFilter{
			ConnectionName: telnyx.IPConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("contains"),
			},
			Fqdn:                   telnyx.String("fqdn"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.IPConnectionListParamsSortConnectionName,
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
	_, err := client.IPConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
