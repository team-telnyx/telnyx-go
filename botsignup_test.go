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

func TestBotSignupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.BotSignup.New(context.TODO(), telnyx.BotSignupNewParams{
		BotChallengeAnswer:      "35",
		BotChallengeNonce:       "c6feda4e-6501-4db9-a21f-665e5b4ce2ba",
		PrivacyPolicyURL:        "https://telnyx.com/privacy-policy",
		TermsAndConditionsURL:   "https://telnyx.com/terms-and-conditions-of-service",
		TermsOfService:          true,
		Email:                   telnyx.String("agent-owner@example.com"),
		TermsAndConditionsEuURL: telnyx.String("https://telnyx.com/terms-and-conditions-of-service-eu"),
		TermsOfServiceEu:        true,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBotSignupResendMagicLink(t *testing.T) {
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
	_, err := client.BotSignup.ResendMagicLink(context.TODO(), telnyx.BotSignupResendMagicLinkParams{
		Email: "agent-owner@example.com",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
