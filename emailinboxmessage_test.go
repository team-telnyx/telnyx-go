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

func TestEmailInboxMessageUpdate(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageUpdateParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ReadAt: telnyx.EmailInboxMessageUpdateParamsReadAtUnion{
				OfServerReadTime: telnyx.Bool(true),
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

func TestEmailInboxMessageListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageListParams{
			FilterFrom:           telnyx.String("filter[from]"),
			FilterLabel:          telnyx.String("filter[label]"),
			FilterRead:           telnyx.Bool(true),
			FilterReceivedAfter:  telnyx.Time(time.Now()),
			FilterReceivedBefore: telnyx.Time(time.Now()),
			FilterSearch:         telnyx.String("filter[search]"),
			FilterSubject:        telnyx.String("filter[subject]"),
			FilterUnread:         telnyx.Bool(true),
			PageAfter:            telnyx.String("page[after]"),
			PageSize:             telnyx.Int(1),
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

func TestEmailInboxMessageDraftsWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.Drafts(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageDraftsParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			EmailDraftRequest: telnyx.EmailDraftRequestParam{
				Attachments: []map[string]any{{
					"foo": "bar",
				}},
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
				Metadata: map[string]any{
					"foo": "bar",
				},
				ReplyTo:  telnyx.String("reply_to"),
				Subject:  telnyx.String("subject"),
				Tags:     []string{"string"},
				Text:     telnyx.String("text"),
				TextBody: telnyx.String("Thanks for the update — I will review today."),
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
