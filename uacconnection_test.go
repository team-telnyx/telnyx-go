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

func TestUacConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.UacConnections.New(context.TODO(), telnyx.UacConnectionNewParams{
		ConnectionName:                   "my name",
		Active:                           telnyx.Bool(true),
		AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
		AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
		CallCostInWebhooks:               telnyx.Bool(false),
		DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(false),
		DtmfType:                         telnyx.DtmfTypeRfc2833,
		EncodeContactHeaderEnabled:       telnyx.Bool(true),
		EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
		ExternalUacSettings: telnyx.UacExternalSettingsParam{
			AuthUsername:  telnyx.String("auth8492"),
			ExpirationSec: telnyx.Int(600),
			FromUser:      telnyx.String("8492"),
			OutboundProxy: telnyx.String("outbound.sip-pbx.acme.example:5061"),
			Password:      telnyx.String("testtesttest"),
			Proxy:         telnyx.String("sip-pbx.acme.example"),
			Transport:     telnyx.UacExternalSettingsTransportTls,
			Username:      telnyx.String("ext8492"),
		},
		Inbound: telnyx.UacConnectionNewParamsInbound{
			AniNumberFormat:          "+E.164",
			ChannelLimit:             telnyx.Int(10),
			Codecs:                   []string{"G722"},
			DefaultRoutingMethod:     "sequential",
			DnisNumberFormat:         "+e164",
			GenerateRingbackTone:     telnyx.Bool(true),
			IsupHeadersEnabled:       telnyx.Bool(true),
			PrackEnabled:             telnyx.Bool(true),
			ShakenStirEnabled:        telnyx.Bool(true),
			SimultaneousRinging:      "disabled",
			SipCompactHeadersEnabled: telnyx.Bool(true),
			Timeout1xxSecs:           telnyx.Int(10),
			Timeout2xxSecs:           telnyx.Int(20),
		},
		InternalUacSettings: telnyx.UacInternalSettingsParam{
			DestinationUri: telnyx.String("14155550123@acme.sip.telnyx.com"),
		},
		IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
		JitterBuffer: shared.ConnectionJitterBufferParam{
			EnableJitterBuffer:  telnyx.Bool(true),
			JitterbufferMsecMax: telnyx.Int(200),
			JitterbufferMsecMin: telnyx.Int(60),
		},
		NoiseSuppression: telnyx.UacConnectionNewParamsNoiseSuppressionBoth,
		NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
			AttenuationLimit: telnyx.Int(80),
			Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
		},
		OnnetT38PassthroughEnabled: telnyx.Bool(true),
		Outbound: telnyx.UacOutboundParam{
			AniOverride:            telnyx.String("always"),
			AniOverrideType:        telnyx.UacOutboundAniOverrideTypeAlways,
			CallParkingEnabled:     telnyx.Bool(true),
			ChannelLimit:           telnyx.Int(10),
			GenerateRingbackTone:   telnyx.Bool(true),
			InstantRingbackEnabled: telnyx.Bool(true),
			Localization:           telnyx.String("US"),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			T38ReinviteSource:      telnyx.UacOutboundT38ReinviteSourceCustomer,
		},
		Password: telnyx.String("my123secure456password789"),
		RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
			CaptureEnabled:      telnyx.Bool(true),
			Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
			ReportFrequencySecs: telnyx.Int(10),
		},
		SipUriCallingPreference: telnyx.UacConnectionNewParamsSipUriCallingPreferenceDisabled,
		Tags:                    []string{"tag1", "tag2"},
		UserName:                telnyx.String("myusername123"),
		WebhookAPIVersion:       telnyx.UacConnectionNewParamsWebhookAPIVersionV1,
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

func TestUacConnectionGet(t *testing.T) {
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
	_, err := client.UacConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUacConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.UacConnections.Update(
		context.TODO(),
		"id",
		telnyx.UacConnectionUpdateParams{
			Active:                           telnyx.Bool(true),
			AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
			AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
			CallCostInWebhooks:               telnyx.Bool(false),
			ConnectionName:                   telnyx.String("my name"),
			DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(false),
			DtmfType:                         telnyx.DtmfTypeRfc2833,
			EncodeContactHeaderEnabled:       telnyx.Bool(true),
			EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
			ExternalUacSettings: telnyx.UacExternalSettingsParam{
				AuthUsername:  telnyx.String("auth8492"),
				ExpirationSec: telnyx.Int(600),
				FromUser:      telnyx.String("8492"),
				OutboundProxy: telnyx.String("outbound.sip-pbx.acme.example:5061"),
				Password:      telnyx.String("testtesttest"),
				Proxy:         telnyx.String("sip-pbx.acme.example"),
				Transport:     telnyx.UacExternalSettingsTransportTls,
				Username:      telnyx.String("ext8492"),
			},
			Inbound: telnyx.UacConnectionUpdateParamsInbound{
				AniNumberFormat:          "+E.164",
				ChannelLimit:             telnyx.Int(10),
				Codecs:                   []string{"G722"},
				DefaultRoutingMethod:     "sequential",
				DnisNumberFormat:         "+e164",
				GenerateRingbackTone:     telnyx.Bool(true),
				IsupHeadersEnabled:       telnyx.Bool(true),
				PrackEnabled:             telnyx.Bool(true),
				ShakenStirEnabled:        telnyx.Bool(true),
				SimultaneousRinging:      "disabled",
				SipCompactHeadersEnabled: telnyx.Bool(true),
				Timeout1xxSecs:           telnyx.Int(10),
				Timeout2xxSecs:           telnyx.Int(20),
			},
			InternalUacSettings: telnyx.UacInternalSettingsParam{
				DestinationUri: telnyx.String("14155550123@acme.sip.telnyx.com"),
			},
			IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
			JitterBuffer: shared.ConnectionJitterBufferParam{
				EnableJitterBuffer:  telnyx.Bool(true),
				JitterbufferMsecMax: telnyx.Int(200),
				JitterbufferMsecMin: telnyx.Int(60),
			},
			NoiseSuppression: telnyx.UacConnectionUpdateParamsNoiseSuppressionBoth,
			NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
				AttenuationLimit: telnyx.Int(80),
				Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
			},
			OnnetT38PassthroughEnabled: telnyx.Bool(true),
			Outbound: telnyx.UacOutboundParam{
				AniOverride:            telnyx.String("always"),
				AniOverrideType:        telnyx.UacOutboundAniOverrideTypeAlways,
				CallParkingEnabled:     telnyx.Bool(true),
				ChannelLimit:           telnyx.Int(10),
				GenerateRingbackTone:   telnyx.Bool(true),
				InstantRingbackEnabled: telnyx.Bool(true),
				Localization:           telnyx.String("US"),
				OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
				T38ReinviteSource:      telnyx.UacOutboundT38ReinviteSourceCustomer,
			},
			Password: telnyx.String("my123secure456password789"),
			RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
				CaptureEnabled:      telnyx.Bool(true),
				Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
				ReportFrequencySecs: telnyx.Int(10),
			},
			SipUriCallingPreference: telnyx.UacConnectionUpdateParamsSipUriCallingPreferenceDisabled,
			Tags:                    []string{"tag1", "tag2"},
			UserName:                telnyx.String("myusername123"),
			WebhookAPIVersion:       telnyx.UacConnectionUpdateParamsWebhookAPIVersionV1,
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

func TestUacConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.UacConnections.List(context.TODO(), telnyx.UacConnectionListParams{
		Filter: telnyx.UacConnectionListParamsFilter{
			ConnectionName: telnyx.UacConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("contains"),
			},
			Fqdn:                   telnyx.String("fqdn"),
			OutboundVoiceProfileID: telnyx.String("1293384261075731499"),
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.UacConnectionListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUacConnectionDelete(t *testing.T) {
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
	_, err := client.UacConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
