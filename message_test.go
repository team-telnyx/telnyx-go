// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestMessageGet(t *testing.T) {
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
	_, err := client.Messages.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageCancelScheduled(t *testing.T) {
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
	_, err := client.Messages.CancelScheduled(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Schedule(context.TODO(), telnyx.MessageScheduleParams{
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		From:               telnyx.String("+18445550001"),
		MediaURLs:          []string{"string"},
		MessagingProfileID: telnyx.String("abc85f64-5717-4562-b3fc-2c9600000000"),
		SendAt:             telnyx.Time(time.Now()),
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageScheduleParamsTypeSMS,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Send(context.TODO(), telnyx.MessageSendParams{
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		From:               telnyx.String("+18445550001"),
		MediaURLs:          []string{"http://example.com"},
		MessagingProfileID: telnyx.String("abc85f64-5717-4562-b3fc-2c9600000000"),
		SendAt:             telnyx.Time(time.Now()),
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendGroupMmsWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendGroupMms(context.TODO(), telnyx.MessageSendGroupMmsParams{
		From:               "+13125551234",
		To:                 []string{"+18655551234", "+14155551234"},
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendLongCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendLongCode(context.TODO(), telnyx.MessageSendLongCodeParams{
		From:               "+18445550001",
		To:                 "+13125550002",
		AutoDetect:         telnyx.Bool(true),
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendLongCodeParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendNumberPoolWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendNumberPool(context.TODO(), telnyx.MessageSendNumberPoolParams{
		MessagingProfileID: "abc85f64-5717-4562-b3fc-2c9600000000",
		To:                 "+13125550002",
		AutoDetect:         telnyx.Bool(true),
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendNumberPoolParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendShortCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendShortCode(context.TODO(), telnyx.MessageSendShortCodeParams{
		From:               "+18445550001",
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendShortCodeParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
