// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/internal/testutil"
	"github.com/team-telnyx/telnyx-go/option"
)

func TestVerifyProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VerifyProfiles.New(context.TODO(), telnyx.VerifyProfileNewParams{
		Name: "Test Profile",
		Call: telnyx.VerifyProfileNewParamsCall{
			AppName:                        telnyx.String("Example Secure App"),
			CodeLength:                     telnyx.Int(6),
			DefaultVerificationTimeoutSecs: telnyx.Int(300),
			MessagingTemplateID:            telnyx.String("0abb5b4f-459f-445a-bfcd-488998b7572d"),
			WhitelistedDestinations:        []string{"US", "CA"},
		},
		Flashcall: telnyx.VerifyProfileNewParamsFlashcall{
			DefaultVerificationTimeoutSecs: telnyx.Int(300),
			WhitelistedDestinations:        []string{"US", "CA"},
		},
		Language: telnyx.String("en-US"),
		SMS: telnyx.VerifyProfileNewParamsSMS{
			WhitelistedDestinations:        []string{"US", "CA"},
			AlphaSender:                    telnyx.String("sqF"),
			AppName:                        telnyx.String("Example Secure App"),
			CodeLength:                     telnyx.Int(6),
			DefaultVerificationTimeoutSecs: telnyx.Int(300),
			MessagingTemplateID:            telnyx.String("0abb5b4f-459f-445a-bfcd-488998b7572d"),
		},
		WebhookFailoverURL: telnyx.String("http://example.com/webhook/failover"),
		WebhookURL:         telnyx.String("http://example.com/webhook"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerifyProfileGet(t *testing.T) {
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
	_, err := client.VerifyProfiles.Get(context.TODO(), "12ade33a-21c0-473b-b055-b3c836e1c292")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerifyProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.VerifyProfiles.Update(
		context.TODO(),
		"12ade33a-21c0-473b-b055-b3c836e1c292",
		telnyx.VerifyProfileUpdateParams{
			Call: telnyx.VerifyProfileUpdateParamsCall{
				AppName:                        telnyx.String("Example Secure App"),
				CodeLength:                     telnyx.Int(6),
				DefaultVerificationTimeoutSecs: telnyx.Int(300),
				MessagingTemplateID:            telnyx.String("0abb5b4f-459f-445a-bfcd-488998b7572d"),
				WhitelistedDestinations:        []string{"US", "CA"},
			},
			Flashcall: telnyx.VerifyProfileUpdateParamsFlashcall{
				DefaultVerificationTimeoutSecs: telnyx.Int(300),
				WhitelistedDestinations:        []string{"US", "CA"},
			},
			Language: telnyx.String("en-US"),
			Name:     telnyx.String("Test Profile"),
			SMS: telnyx.VerifyProfileUpdateParamsSMS{
				AlphaSender:                    telnyx.String("sqF"),
				AppName:                        telnyx.String("Example Secure App"),
				CodeLength:                     telnyx.Int(6),
				DefaultVerificationTimeoutSecs: telnyx.Int(300),
				MessagingTemplateID:            telnyx.String("0abb5b4f-459f-445a-bfcd-488998b7572d"),
				WhitelistedDestinations:        []string{"US", "CA"},
			},
			WebhookFailoverURL: telnyx.String("http://example.com/webhook/failover"),
			WebhookURL:         telnyx.String("http://example.com/webhook"),
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

func TestVerifyProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.VerifyProfiles.List(context.TODO(), telnyx.VerifyProfileListParams{
		Filter: telnyx.VerifyProfileListParamsFilter{
			Name: telnyx.String("name"),
		},
		Page: telnyx.VerifyProfileListParamsPage{
			Number: telnyx.Int(0),
			Size:   telnyx.Int(0),
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

func TestVerifyProfileDelete(t *testing.T) {
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
	_, err := client.VerifyProfiles.Delete(context.TODO(), "12ade33a-21c0-473b-b055-b3c836e1c292")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerifyProfileNewTemplate(t *testing.T) {
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
	_, err := client.VerifyProfiles.NewTemplate(context.TODO(), telnyx.VerifyProfileNewTemplateParams{
		Text: "Your {{app_name}} verification code is: {{code}}.",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerifyProfileGetTemplates(t *testing.T) {
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
	_, err := client.VerifyProfiles.GetTemplates(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVerifyProfileUpdateTemplate(t *testing.T) {
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
	_, err := client.VerifyProfiles.UpdateTemplate(
		context.TODO(),
		"12ade33a-21c0-473b-b055-b3c836e1c292",
		telnyx.VerifyProfileUpdateTemplateParams{
			Text: "Your {{app_name}} verification code is: {{code}}.",
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
