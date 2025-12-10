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

func TestFqdnConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FqdnConnections.New(context.TODO(), telnyx.FqdnConnectionNewParams{
		ConnectionName:                   "string",
		Active:                           telnyx.Bool(true),
		AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
		AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
		CallCostInWebhooks:               telnyx.Bool(false),
		DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
		DtmfType:                         telnyx.DtmfTypeRfc2833,
		EncodeContactHeaderEnabled:       telnyx.Bool(true),
		EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
		Inbound: telnyx.InboundFqdnParam{
			AniNumberFormat:             telnyx.InboundFqdnAniNumberFormatPlusE164,
			ChannelLimit:                telnyx.Int(10),
			Codecs:                      []string{"G722"},
			DefaultPrimaryFqdnID:        telnyx.String("default_primary_fqdn_id"),
			DefaultRoutingMethod:        telnyx.InboundFqdnDefaultRoutingMethodSequential,
			DefaultSecondaryFqdnID:      telnyx.String("default_secondary_fqdn_id"),
			DefaultTertiaryFqdnID:       telnyx.String("default_tertiary_fqdn_id"),
			DnisNumberFormat:            telnyx.InboundFqdnDnisNumberFormatPlusE164,
			GenerateRingbackTone:        telnyx.Bool(true),
			IsupHeadersEnabled:          telnyx.Bool(true),
			PrackEnabled:                telnyx.Bool(true),
			ShakenStirEnabled:           telnyx.Bool(true),
			SipCompactHeadersEnabled:    telnyx.Bool(true),
			SipRegion:                   telnyx.InboundFqdnSipRegionUs,
			SipSubdomain:                telnyx.String("string"),
			SipSubdomainReceiveSettings: telnyx.InboundFqdnSipSubdomainReceiveSettingsOnlyMyConnections,
			Timeout1xxSecs:              telnyx.Int(10),
			Timeout2xxSecs:              telnyx.Int(10),
		},
		IosPushCredentialID:        telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
		MicrosoftTeamsSbc:          telnyx.Bool(true),
		OnnetT38PassthroughEnabled: telnyx.Bool(true),
		Outbound: telnyx.OutboundFqdnParam{
			AniOverride:            telnyx.String("+1234567890"),
			AniOverrideType:        telnyx.OutboundFqdnAniOverrideTypeAlways,
			CallParkingEnabled:     telnyx.Bool(true),
			ChannelLimit:           telnyx.Int(10),
			EncryptedMedia:         telnyx.EncryptedMediaSrtp,
			GenerateRingbackTone:   telnyx.Bool(true),
			InstantRingbackEnabled: telnyx.Bool(true),
			IPAuthenticationMethod: telnyx.OutboundFqdnIPAuthenticationMethodCredentialAuthentication,
			IPAuthenticationToken:  telnyx.String("aBcD1234"),
			Localization:           telnyx.String("string"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
			T38ReinviteSource:      telnyx.OutboundFqdnT38ReinviteSourceCustomer,
			TechPrefix:             telnyx.String("0123"),
			Timeout1xxSecs:         telnyx.Int(10),
			Timeout2xxSecs:         telnyx.Int(10),
		},
		RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
			CaptureEnabled:      telnyx.Bool(true),
			Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
			ReportFrequencySecs: telnyx.Int(10),
		},
		Tags:                    []string{"tag1", "tag2"},
		TransportProtocol:       telnyx.TransportProtocolUdp,
		WebhookAPIVersion:       telnyx.WebhookAPIVersion1,
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

func TestFqdnConnectionGet(t *testing.T) {
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
	_, err := client.FqdnConnections.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFqdnConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.FqdnConnections.Update(
		context.TODO(),
		"id",
		telnyx.FqdnConnectionUpdateParams{
			Active:                           telnyx.Bool(true),
			AnchorsiteOverride:               telnyx.AnchorsiteOverrideLatency,
			AndroidPushCredentialID:          telnyx.String("06b09dfd-7154-4980-8b75-cebf7a9d4f8e"),
			CallCostInWebhooks:               telnyx.Bool(true),
			ConnectionName:                   telnyx.String("string"),
			DefaultOnHoldComfortNoiseEnabled: telnyx.Bool(true),
			DtmfType:                         telnyx.DtmfTypeRfc2833,
			EncodeContactHeaderEnabled:       telnyx.Bool(true),
			EncryptedMedia:                   telnyx.EncryptedMediaSrtp,
			Inbound: telnyx.InboundFqdnParam{
				AniNumberFormat:             telnyx.InboundFqdnAniNumberFormatPlusE164,
				ChannelLimit:                telnyx.Int(10),
				Codecs:                      []string{"G722"},
				DefaultPrimaryFqdnID:        telnyx.String("default_primary_fqdn_id"),
				DefaultRoutingMethod:        telnyx.InboundFqdnDefaultRoutingMethodSequential,
				DefaultSecondaryFqdnID:      telnyx.String("default_secondary_fqdn_id"),
				DefaultTertiaryFqdnID:       telnyx.String("default_tertiary_fqdn_id"),
				DnisNumberFormat:            telnyx.InboundFqdnDnisNumberFormatPlusE164,
				GenerateRingbackTone:        telnyx.Bool(true),
				IsupHeadersEnabled:          telnyx.Bool(true),
				PrackEnabled:                telnyx.Bool(true),
				ShakenStirEnabled:           telnyx.Bool(true),
				SipCompactHeadersEnabled:    telnyx.Bool(true),
				SipRegion:                   telnyx.InboundFqdnSipRegionUs,
				SipSubdomain:                telnyx.String("string"),
				SipSubdomainReceiveSettings: telnyx.InboundFqdnSipSubdomainReceiveSettingsOnlyMyConnections,
				Timeout1xxSecs:              telnyx.Int(10),
				Timeout2xxSecs:              telnyx.Int(10),
			},
			IosPushCredentialID:        telnyx.String("ec0c8e5d-439e-4620-a0c1-9d9c8d02a836"),
			OnnetT38PassthroughEnabled: telnyx.Bool(true),
			Outbound: telnyx.OutboundFqdnParam{
				AniOverride:            telnyx.String("ani_override"),
				AniOverrideType:        telnyx.OutboundFqdnAniOverrideTypeNormal,
				CallParkingEnabled:     telnyx.Bool(true),
				ChannelLimit:           telnyx.Int(0),
				EncryptedMedia:         telnyx.EncryptedMediaSrtp,
				GenerateRingbackTone:   telnyx.Bool(true),
				InstantRingbackEnabled: telnyx.Bool(true),
				IPAuthenticationMethod: telnyx.OutboundFqdnIPAuthenticationMethodCredentialAuthentication,
				IPAuthenticationToken:  telnyx.String("ip_authentication_token"),
				Localization:           telnyx.String("US"),
				OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
				T38ReinviteSource:      telnyx.OutboundFqdnT38ReinviteSourceTelnyx,
				TechPrefix:             telnyx.String("tech_prefix"),
				Timeout1xxSecs:         telnyx.Int(1),
				Timeout2xxSecs:         telnyx.Int(1),
			},
			RtcpSettings: telnyx.ConnectionRtcpSettingsParam{
				CaptureEnabled:      telnyx.Bool(true),
				Port:                telnyx.ConnectionRtcpSettingsPortRtcpMux,
				ReportFrequencySecs: telnyx.Int(10),
			},
			Tags:                    []string{"tag1", "tag2"},
			TransportProtocol:       telnyx.TransportProtocolUdp,
			WebhookAPIVersion:       telnyx.WebhookAPIVersion1,
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

func TestFqdnConnectionListWithOptionalParams(t *testing.T) {
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
	_, err := client.FqdnConnections.List(context.TODO(), telnyx.FqdnConnectionListParams{
		Filter: telnyx.FqdnConnectionListParamsFilter{
			ConnectionName: telnyx.FqdnConnectionListParamsFilterConnectionName{
				Contains: telnyx.String("contains"),
			},
			Fqdn:                   telnyx.String("fqdn"),
			OutboundVoiceProfileID: telnyx.String("outbound_voice_profile_id"),
		},
		Page: telnyx.FqdnConnectionListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.FqdnConnectionListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFqdnConnectionDelete(t *testing.T) {
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
	_, err := client.FqdnConnections.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
