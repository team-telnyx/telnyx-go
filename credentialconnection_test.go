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

func TestCredentialConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CredentialConnections.New(context.TODO(), telnyx.CredentialConnectionNewParams{
		ConnectionName:                   "my name",
		Password:                         "my123secure456password789",
		UserName:                         "myusername123",
		Active:                           telnyx.Bool(true),
		AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
		AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
		CallCostInWebhooks:               telnyx.Bool(false),
		DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(false),
		DtmfType:                         telnyx.DtmfTypeRfc2833,
		EncodeContactHeaderEnabled:       telnyx.Bool(true),
		EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
		Inbound: telnyx.CredentialInboundParam{
			AniNumberFormat:          telnyx.CredentialInboundAniNumberFormatPlusE164,
			ChannelLimit:             telnyx.Int(10),
			Codecs:                   []string{"G722"},
			DnisNumberFormat:         telnyx.CredentialInboundDnisNumberFormatPlusE164,
			GenerateRingbackTone:     telnyx.Bool(true),
			IsupHeadersEnabled:       telnyx.Bool(true),
			PrackEnabled:             telnyx.Bool(true),
			ShakenStirEnabled:        telnyx.Bool(true),
			SimultaneousRinging:      telnyx.CredentialInboundSimultaneousRingingDisabled,
			SipCompactHeadersEnabled: telnyx.Bool(true),
			Timeout1xxSecs:           telnyx.Int(10),
			Timeout2xxSecs:           telnyx.Int(20),
		},
		IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
		NoiseSuppression:    telnyx.CredentialConnectionNewParamsNoiseSuppressionBoth,
		NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
			AttenuationLimit: telnyx.Int(80),
			Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
		},
		OnnetT38PassthroughEnabled: telnyx.Bool(true),
		Outbound: telnyx.CredentialOutboundParam{
			AniOverride:            telnyx.String("always"),
			AniOverrideType:        telnyx.CredentialOutboundAniOverrideTypeAlways,
			CallParkingEnabled:     telnyx.Bool(true),
			ChannelLimit:           telnyx.Int(10),
			GenerateRingbackTone:   telnyx.Bool(true),
			InstantRingbackEnabled: telnyx.Bool(true),
			Localization:           telnyx.String("US"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
			T38ReinviteSource:      telnyx.CredentialOutboundT38ReinviteSourceCustomer,
		},
		RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
			CaptureEnabled:      telnyx.Bool(true),
			Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
			ReportFrequencySecs: telnyx.Int(10),
		},
		SipUriCallingPreference: telnyx.CredentialConnectionNewParamsSipUriCallingPreferenceDisabled,
		Tags:                    []string{"tag1", "tag2"},
		WebhookAPIVersion:       telnyx.CredentialConnectionNewParamsWebhookAPIVersionV1,
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

func TestCredentialConnectionGet(t *testing.T) {
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
	_, err := client.CredentialConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCredentialConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CredentialConnections.Update(
		context.TODO(),
		"id",
		telnyx.CredentialConnectionUpdateParams{
			Active:                           telnyx.Bool(true),
			AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
			AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
			CallCostInWebhooks:               telnyx.Bool(false),
			ConnectionName:                   telnyx.String("my name"),
			DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(false),
			DtmfType:                         telnyx.DtmfTypeRfc2833,
			EncodeContactHeaderEnabled:       telnyx.Bool(true),
			EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
			Inbound: telnyx.CredentialInboundParam{
				AniNumberFormat:          telnyx.CredentialInboundAniNumberFormatPlusE164,
				ChannelLimit:             telnyx.Int(10),
				Codecs:                   []string{"G722"},
				DnisNumberFormat:         telnyx.CredentialInboundDnisNumberFormatPlusE164,
				GenerateRingbackTone:     telnyx.Bool(true),
				IsupHeadersEnabled:       telnyx.Bool(true),
				PrackEnabled:             telnyx.Bool(true),
				ShakenStirEnabled:        telnyx.Bool(true),
				SimultaneousRinging:      telnyx.CredentialInboundSimultaneousRingingDisabled,
				SipCompactHeadersEnabled: telnyx.Bool(true),
				Timeout1xxSecs:           telnyx.Int(10),
				Timeout2xxSecs:           telnyx.Int(20),
			},
			IosPushCredentialID: telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
			NoiseSuppression:    telnyx.CredentialConnectionUpdateParamsNoiseSuppressionBoth,
			NoiseSuppressionDetails: shared.ConnectionNoiseSuppressionDetailsParam{
				AttenuationLimit: telnyx.Int(80),
				Engine:           shared.ConnectionNoiseSuppressionDetailsEngineDeepFilterNet,
			},
			OnnetT38PassthroughEnabled: telnyx.Bool(true),
			Outbound: telnyx.CredentialOutboundParam{
				AniOverride:            telnyx.String("always"),
				AniOverrideType:        telnyx.CredentialOutboundAniOverrideTypeAlways,
				CallParkingEnabled:     telnyx.Bool(true),
				ChannelLimit:           telnyx.Int(10),
				GenerateRingbackTone:   telnyx.Bool(true),
				InstantRingbackEnabled: telnyx.Bool(true),
				Localization:           telnyx.String("US"),
				OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
				T38ReinviteSource:      telnyx.CredentialOutboundT38ReinviteSourceCustomer,
			},
			Password: telnyx.String("my123secure456password789"),
			RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
				CaptureEnabled:      telnyx.Bool(true),
				Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
				ReportFrequencySecs: telnyx.Int(10),
			},
			SipUriCallingPreference: telnyx.CredentialConnectionUpdateParamsSipUriCallingPreferenceDisabled,
			Tags:                    []string{"tag1", "tag2"},
			UserName:                telnyx.String("myusername123"),
			WebhookAPIVersion:       telnyx.CredentialConnectionUpdateParamsWebhookAPIVersionV1,
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

func TestCredentialConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.CredentialConnections.List(context.TODO(), telnyx.CredentialConnectionListParams{
		Filter: telnyx.CredentialConnectionListParamsFilter{
			ConnectionName: telnyx.CredentialConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("contains"),
			},
			Fqdn:                   telnyx.String("fqdn"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.CredentialConnectionListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCredentialConnectionDelete(t *testing.T) {
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
	_, err := client.CredentialConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
