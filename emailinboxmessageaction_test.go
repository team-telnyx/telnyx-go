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

func TestEmailInboxMessageActionForwardWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.Actions.Forward(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageActionForwardParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			To: telnyx.EmailInboxMessageActionForwardParamsToUnion{
				OfString: telnyx.String("new@example.com"),
			},
			Bcc: telnyx.InboxActionRecipientInputUnionParam{
				OfInboxRecipientList: []telnyx.InboxActionEmailAddressInputUnionParam{{
					OfString: telnyx.String("blind@example.com"),
				}},
			},
			Cc: telnyx.InboxActionRecipientInputUnionParam{
				OfInboxRecipientList: []telnyx.InboxActionEmailAddressInputUnionParam{{
					OfInboxRecipientAddress: &telnyx.InboxActionEmailAddressInputInboxRecipientAddressParam{
						Email: "copy@example.com",
						Name:  telnyx.String("name"),
					},
				}},
			},
			HTML: telnyx.String("html"),
			Text: telnyx.String("FYI"),
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

func TestEmailInboxMessageActionReplyWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.Actions.Reply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageActionReplyParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ReplyEmailInboxMessageRequest: telnyx.ReplyEmailInboxMessageRequestParam{
				HTML: telnyx.String("P"),
				Text: telnyx.String("Thanks for the update."),
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

func TestEmailInboxMessageActionReplyAllWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailInboxes.Messages.Actions.ReplyAll(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.EmailInboxMessageActionReplyAllParams{
			InboxID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ReplyEmailInboxMessageRequest: telnyx.ReplyEmailInboxMessageRequestParam{
				HTML: telnyx.String("P"),
				Text: telnyx.String("Everyone, please review."),
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
