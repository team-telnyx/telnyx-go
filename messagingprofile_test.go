// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestMessagingProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingProfiles.New(context.TODO(), telnyx.MessagingProfileNewParams{
		Name:                    "My name",
		WhitelistedDestinations: []string{"US"},
		AlphaSender:             telnyx.String("sqF"),
		DailySpendLimit:         telnyx.String("269125115713"),
		DailySpendLimitEnabled:  telnyx.Bool(true),
		Enabled:                 telnyx.Bool(true),
		MmsFallBackToSMS:        telnyx.Bool(true),
		MmsTranscoding:          telnyx.Bool(true),
		NumberPoolSettings: telnyx.NumberPoolSettingsParam{
			LongCodeWeight: 1,
			SkipUnhealthy:  true,
			TollFreeWeight: 10,
			Geomatch:       telnyx.Bool(false),
			StickySender:   telnyx.Bool(false),
		},
		URLShortenerSettings: telnyx.URLShortenerSettingsParam{
			Domain:               "example.ex",
			Prefix:               telnyx.String(""),
			ReplaceBlacklistOnly: telnyx.Bool(true),
			SendWebhooks:         telnyx.Bool(false),
		},
		WebhookAPIVersion:  telnyx.MessagingProfileNewParamsWebhookAPIVersion2,
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("https://www.example.com/hooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingProfileGet(t *testing.T) {
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
	_, err := client.MessagingProfiles.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingProfiles.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.MessagingProfileUpdateParams{
			AlphaSender:            telnyx.String("sqF"),
			DailySpendLimit:        telnyx.String("269125115713"),
			DailySpendLimitEnabled: telnyx.Bool(true),
			Enabled:                telnyx.Bool(true),
			MmsFallBackToSMS:       telnyx.Bool(true),
			MmsTranscoding:         telnyx.Bool(true),
			Name:                   telnyx.String("Updated Profile for Messages"),
			NumberPoolSettings: telnyx.NumberPoolSettingsParam{
				LongCodeWeight: 2,
				SkipUnhealthy:  false,
				TollFreeWeight: 10,
				Geomatch:       telnyx.Bool(false),
				StickySender:   telnyx.Bool(true),
			},
			URLShortenerSettings: telnyx.URLShortenerSettingsParam{
				Domain:               "example.ex",
				Prefix:               telnyx.String("cmpny"),
				ReplaceBlacklistOnly: telnyx.Bool(true),
				SendWebhooks:         telnyx.Bool(false),
			},
			V1Secret:                telnyx.String("rP1VamejkU2v0qIUxntqLW2c"),
			WebhookAPIVersion:       telnyx.MessagingProfileUpdateParamsWebhookAPIVersion2,
			WebhookFailoverURL:      telnyx.String("https://backup.example.com/hooks"),
			WebhookURL:              telnyx.String("https://www.example.com/hooks"),
			WhitelistedDestinations: []string{"US"},
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

func TestMessagingProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingProfiles.List(context.TODO(), telnyx.MessagingProfileListParams{
		Filter: telnyx.MessagingProfileListParamsFilter{
			Name: telnyx.String("name"),
		},
		Page: telnyx.MessagingProfileListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingProfileDelete(t *testing.T) {
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
	_, err := client.MessagingProfiles.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingProfileListPhoneNumbersWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingProfiles.ListPhoneNumbers(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.MessagingProfileListPhoneNumbersParams{
			Page: telnyx.MessagingProfileListPhoneNumbersParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
			},
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

func TestMessagingProfileListShortCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingProfiles.ListShortCodes(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.MessagingProfileListShortCodesParams{
			Page: telnyx.MessagingProfileListShortCodesParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
			},
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
