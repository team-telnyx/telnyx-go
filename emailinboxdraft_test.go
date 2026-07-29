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

func TestEmailInboxDraftNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftNewParams{
			EmailDraftRequest: telnyx.EmailDraftRequestParam{
				Attachments: []any{map[string]any{}},
				Bcc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				Cc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				FromEmail: telnyx.String("from_email"),
				FromName:  telnyx.String("from_name"),
				Headers: map[string]string{
					"foo": "string",
				},
				HTML:     telnyx.String("html"),
				HTMLBody: telnyx.String("html_body"),
				Labels:   []string{"important"},
				Metadata: map[string]any{},
				ReplyTo:  telnyx.String("reply_to"),
				Subject:  telnyx.String("Quarterly update"),
				Tags:     []string{"string"},
				Text:     telnyx.String("text"),
				TextBody: telnyx.String("Here is the update."),
				To: []telnyx.EmailAddressInputUnionParam{{
					OfEmailAddress: &telnyx.EmailAddressParam{
						Email: "recipient@example.com",
						Name:  telnyx.String("Recipient"),
					},
				}},
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

func TestEmailInboxDraftGet(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftGetParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestEmailInboxDraftUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftUpdateParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			EmailDraftRequest: telnyx.EmailDraftRequestParam{
				Attachments: []any{map[string]any{}},
				Bcc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				Cc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				FromEmail: telnyx.String("from_email"),
				FromName:  telnyx.String("from_name"),
				Headers: map[string]string{
					"foo": "string",
				},
				HTML:     telnyx.String("html"),
				HTMLBody: telnyx.String("html_body"),
				Labels:   []string{"string"},
				Metadata: map[string]any{},
				ReplyTo:  telnyx.String("reply_to"),
				Subject:  telnyx.String("Quarterly update (revised)"),
				Tags:     []string{"string"},
				Text:     telnyx.String("text"),
				TextBody: telnyx.String("Updated body."),
				To: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
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

func TestEmailInboxDraftListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftListParams{
			FilterStatus: telnyx.EmailInboxDraftListParamsFilterStatusDraft,
			PageAfter:    telnyx.String("page[after]"),
			PageSize:     telnyx.Int(1),
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

func TestEmailInboxDraftDelete(t *testing.T) {
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
	err := client.EmailInboxes.Drafts.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftDeleteParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestEmailInboxDraftPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.Patch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftPatchParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			EmailDraftRequest: telnyx.EmailDraftRequestParam{
				Attachments: []any{map[string]any{}},
				Bcc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				Cc: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
				FromEmail: telnyx.String("from_email"),
				FromName:  telnyx.String("from_name"),
				Headers: map[string]string{
					"foo": "string",
				},
				HTML:     telnyx.String("html"),
				HTMLBody: telnyx.String("html_body"),
				Labels:   []string{"string"},
				Metadata: map[string]any{},
				ReplyTo:  telnyx.String("reply_to"),
				Subject:  telnyx.String("subject"),
				Tags:     []string{"string"},
				Text:     telnyx.String("text"),
				TextBody: telnyx.String("text_body"),
				To: []telnyx.EmailAddressInputUnionParam{{
					OfString: telnyx.String("string"),
				}},
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

func TestEmailInboxDraftSend(t *testing.T) {
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
	_, err := client.EmailInboxes.Drafts.Send(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxDraftSendParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
