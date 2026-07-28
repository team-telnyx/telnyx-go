// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestEmailMessageNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailMessages.New(context.TODO(), telnyx.EmailMessageNewParams{
		From: telnyx.EmailAddressInputUnionParam{
			OfString: telnyx.String("sender@example.com"),
		},
		To: []telnyx.EmailAddressInputUnionParam{{
			OfString: telnyx.String("recipient@example.com"),
		}},
		Attachments: []telnyx.AttachmentRequestParam{{
			Content:     telnyx.String("content"),
			ContentID:   telnyx.String("content_id"),
			ContentType: telnyx.String("content_type"),
			Disposition: telnyx.String("disposition"),
			Filename:    telnyx.String("filename"),
		}},
		Bcc: []telnyx.EmailAddressInputUnionParam{{
			OfString: telnyx.String("string"),
		}},
		Cc: []telnyx.EmailAddressInputUnionParam{{
			OfString: telnyx.String("string"),
		}},
		ForwardOfMessageID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FromName:           telnyx.String("from_name"),
		GroupID:            telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Headers: map[string]string{
			"foo": "string",
		},
		HTMLBody:           telnyx.String("html_body"),
		IgnoreSuppression:  telnyx.Bool(true),
		InReplyToMessageID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InlineCss:          telnyx.Bool(true),
		Metadata: map[string]any{
			"foo": "bar",
		},
		ReplyTo: telnyx.EmailAddressInputUnionParam{
			OfString: telnyx.String("string"),
		},
		ReplyToAll:  telnyx.Bool(true),
		SandboxMode: telnyx.Bool(true),
		ScheduledAt: telnyx.Time(time.Now()),
		SendAt:      telnyx.Time(time.Now()),
		Subject:     telnyx.String("Hello from Telnyx"),
		Tags:        []string{"string"},
		TemplateID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		TemplateVariables: map[string]any{
			"foo": "bar",
		},
		TextBody: telnyx.String("This is a test email."),
		TrackingSettings: telnyx.TrackingSettingsParam{
			ClickTracking: telnyx.Bool(true),
			OpenTracking:  telnyx.Bool(true),
		},
		IdempotencyKey: telnyx.String("8e03978e-40d5-43e8-bc93-6894a57f9326"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageGet(t *testing.T) {
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
	_, err := client.EmailMessages.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailMessages.List(context.TODO(), telnyx.EmailMessageListParams{
		PageCursor: telnyx.String("page_cursor"),
		PageSize:   telnyx.Int(1),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageDelete(t *testing.T) {
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
	err := client.EmailMessages.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailMessages.Batch(context.TODO(), telnyx.EmailMessageBatchParams{
		Messages: []telnyx.EmailMessageBatchParamsMessage{{
			From: telnyx.EmailAddressInputUnionParam{
				OfString: telnyx.String("sender@example.com"),
			},
			To: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("recipient1@example.com"),
			}},
			Attachments: []telnyx.AttachmentRequestParam{{
				Content:     telnyx.String("content"),
				ContentID:   telnyx.String("content_id"),
				ContentType: telnyx.String("content_type"),
				Disposition: telnyx.String("disposition"),
				Filename:    telnyx.String("filename"),
			}},
			Bcc: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("string"),
			}},
			Cc: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("string"),
			}},
			FromName: telnyx.String("from_name"),
			GroupID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Headers: map[string]string{
				"foo": "string",
			},
			HTMLBody:          telnyx.String("html_body"),
			IgnoreSuppression: telnyx.Bool(true),
			InlineCss:         telnyx.Bool(true),
			Metadata: map[string]any{
				"foo": "bar",
			},
			ReplyTo: telnyx.EmailAddressInputUnionParam{
				OfString: telnyx.String("string"),
			},
			SandboxMode: telnyx.Bool(true),
			ScheduledAt: telnyx.Time(time.Now()),
			SendAt:      telnyx.Time(time.Now()),
			Subject:     telnyx.String("Hello 1"),
			Tags:        []string{"string"},
			TemplateID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TemplateVariables: map[string]any{
				"foo": "bar",
			},
			TextBody: telnyx.String("Message 1"),
			TrackingSettings: telnyx.TrackingSettingsParam{
				ClickTracking: telnyx.Bool(true),
				OpenTracking:  telnyx.Bool(true),
			},
		}, {
			From: telnyx.EmailAddressInputUnionParam{
				OfString: telnyx.String("sender@example.com"),
			},
			To: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("recipient2@example.com"),
			}},
			Attachments: []telnyx.AttachmentRequestParam{{
				Content:     telnyx.String("content"),
				ContentID:   telnyx.String("content_id"),
				ContentType: telnyx.String("content_type"),
				Disposition: telnyx.String("disposition"),
				Filename:    telnyx.String("filename"),
			}},
			Bcc: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("string"),
			}},
			Cc: []telnyx.EmailAddressInputUnionParam{{
				OfString: telnyx.String("string"),
			}},
			FromName: telnyx.String("from_name"),
			GroupID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Headers: map[string]string{
				"foo": "string",
			},
			HTMLBody:          telnyx.String("html_body"),
			IgnoreSuppression: telnyx.Bool(true),
			InlineCss:         telnyx.Bool(true),
			Metadata: map[string]any{
				"foo": "bar",
			},
			ReplyTo: telnyx.EmailAddressInputUnionParam{
				OfString: telnyx.String("string"),
			},
			SandboxMode: telnyx.Bool(true),
			ScheduledAt: telnyx.Time(time.Now()),
			SendAt:      telnyx.Time(time.Now()),
			Subject:     telnyx.String("Hello 2"),
			Tags:        []string{"string"},
			TemplateID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TemplateVariables: map[string]any{
				"foo": "bar",
			},
			TextBody: telnyx.String("Message 2"),
			TrackingSettings: telnyx.TrackingSettingsParam{
				ClickTracking: telnyx.Bool(true),
				OpenTracking:  telnyx.Bool(true),
			},
		}},
		SandboxMode:    telnyx.Bool(false),
		IdempotencyKey: telnyx.String("8e03978e-40d5-43e8-bc93-6894a57f9326"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageDeleteAll(t *testing.T) {
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
	err := client.EmailMessages.DeleteAll(context.TODO(), telnyx.EmailMessageDeleteAllParams{
		Address: "dev@stainless.com",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageDeleteSchedule(t *testing.T) {
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
	_, err := client.EmailMessages.DeleteSchedule(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEmailMessageGetEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailMessages.GetEvents(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailMessageGetEventsParams{
			PageCursor: telnyx.String("page_cursor"),
			PageSize:   telnyx.Int(1),
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
