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

func TestMessagingTollfreeVerificationRequestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingTollfree.Verification.Requests.New(context.TODO(), telnyx.MessagingTollfreeVerificationRequestNewParams{
		TfVerificationRequest: telnyx.TfVerificationRequestParam{
			AdditionalInformation:    "additionalInformation",
			BusinessAddr1:            "600 Congress Avenue",
			BusinessCity:             "Austin",
			BusinessContactEmail:     "email@example.com",
			BusinessContactFirstName: "John",
			BusinessContactLastName:  "Doe",
			BusinessContactPhone:     "+18005550100",
			BusinessName:             "Telnyx LLC",
			BusinessState:            "Texas",
			BusinessZip:              "78701",
			CorporateWebsite:         "http://example.com",
			IsvReseller:              "isvReseller",
			MessageVolume:            telnyx.Volume100_000,
			OptInWorkflow:            "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
			OptInWorkflowImageURLs: []telnyx.URLParam{{
				URL: "https://telnyx.com/sign-up",
			}, {
				URL: "https://telnyx.com/company/data-privacy",
			}},
			PhoneNumbers: []telnyx.TfPhoneNumberParam{{
				PhoneNumber: "+18773554398",
			}, {
				PhoneNumber: "+18773554399",
			}},
			ProductionMessageContent:    "Your Telnyx OTP is XXXX",
			UseCase:                     telnyx.UseCaseCategories2Fa,
			UseCaseSummary:              "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
			AgeGatedContent:             telnyx.Bool(true),
			BusinessAddr2:               telnyx.String("14th Floor"),
			BusinessRegistrationCountry: telnyx.String("US"),
			BusinessRegistrationNumber:  telnyx.String("12-3456789"),
			BusinessRegistrationType:    telnyx.String("EIN"),
			DoingBusinessAs:             telnyx.String("Acme Services"),
			EntityType:                  telnyx.TollFreeVerificationEntityTypeSoleProprietor,
			HelpMessageResponse:         telnyx.String("Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com"),
			OptInConfirmationResponse:   telnyx.String("You have successfully opted in to receive messages from Acme Corp"),
			OptInKeywords:               telnyx.String("START, YES, SUBSCRIBE"),
			PrivacyPolicyURL:            telnyx.String("https://example.com/privacy"),
			TermsAndConditionURL:        telnyx.String("https://example.com/terms"),
			WebhookURL:                  telnyx.String("http://example-webhook.com"),
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

func TestMessagingTollfreeVerificationRequestGet(t *testing.T) {
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
	_, err := client.MessagingTollfree.Verification.Requests.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingTollfreeVerificationRequestUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingTollfree.Verification.Requests.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.MessagingTollfreeVerificationRequestUpdateParams{
			TfVerificationRequest: telnyx.TfVerificationRequestParam{
				AdditionalInformation:    "additionalInformation",
				BusinessAddr1:            "600 Congress Avenue",
				BusinessCity:             "Austin",
				BusinessContactEmail:     "email@example.com",
				BusinessContactFirstName: "John",
				BusinessContactLastName:  "Doe",
				BusinessContactPhone:     "+18005550100",
				BusinessName:             "Telnyx LLC",
				BusinessState:            "Texas",
				BusinessZip:              "78701",
				CorporateWebsite:         "http://example.com",
				IsvReseller:              "isvReseller",
				MessageVolume:            telnyx.Volume100_000,
				OptInWorkflow:            "User signs into the Telnyx portal, enters a number and is prompted to select whether they want to use 2FA verification for security purposes. If they've opted in a confirmation message is sent out to the handset",
				OptInWorkflowImageURLs: []telnyx.URLParam{{
					URL: "https://telnyx.com/sign-up",
				}, {
					URL: "https://telnyx.com/company/data-privacy",
				}},
				PhoneNumbers: []telnyx.TfPhoneNumberParam{{
					PhoneNumber: "+18773554398",
				}, {
					PhoneNumber: "+18773554399",
				}},
				ProductionMessageContent:    "Your Telnyx OTP is XXXX",
				UseCase:                     telnyx.UseCaseCategories2Fa,
				UseCaseSummary:              "This is a use case where Telnyx sends out 2FA codes to portal users to verify their identity in order to sign into the portal",
				AgeGatedContent:             telnyx.Bool(true),
				BusinessAddr2:               telnyx.String("14th Floor"),
				BusinessRegistrationCountry: telnyx.String("US"),
				BusinessRegistrationNumber:  telnyx.String("12-3456789"),
				BusinessRegistrationType:    telnyx.String("EIN"),
				DoingBusinessAs:             telnyx.String("Acme Services"),
				EntityType:                  telnyx.TollFreeVerificationEntityTypeSoleProprietor,
				HelpMessageResponse:         telnyx.String("Reply HELP for assistance or STOP to unsubscribe. Contact: support@example.com"),
				OptInConfirmationResponse:   telnyx.String("You have successfully opted in to receive messages from Acme Corp"),
				OptInKeywords:               telnyx.String("START, YES, SUBSCRIBE"),
				PrivacyPolicyURL:            telnyx.String("https://example.com/privacy"),
				TermsAndConditionURL:        telnyx.String("https://example.com/terms"),
				WebhookURL:                  telnyx.String("http://example-webhook.com"),
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

func TestMessagingTollfreeVerificationRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.MessagingTollfree.Verification.Requests.List(context.TODO(), telnyx.MessagingTollfreeVerificationRequestListParams{
		Page:        1,
		PageSize:    1,
		DateEnd:     telnyx.Time(time.Now()),
		DateStart:   telnyx.Time(time.Now()),
		PhoneNumber: telnyx.String("phone_number"),
		Status:      telnyx.TfVerificationStatusVerified,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingTollfreeVerificationRequestDelete(t *testing.T) {
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
	err := client.MessagingTollfree.Verification.Requests.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
