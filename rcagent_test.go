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

func TestRcAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Agents.New(context.TODO(), telnyx.RcAgentNewParams{
		BrandID: "11111111-1111-4111-8111-111111111111",
		Configuration: telnyx.AgentConfigurationParam{
			Basics: telnyx.AgentConfigurationBasicsUnionParam{
				OfAgentConfigurationBasicsUnionMember0: &telnyx.AgentConfigurationBasicsUnionMember0Param{
					Email: telnyx.AgentEmailContactParam{
						Address: "support@example.com",
						Label:   "Support",
					},
					BrandColor:  telnyx.String("#123456"),
					Description: telnyx.String("Order confirmations and delivery updates"),
					HeroURL:     telnyx.String("https://www.example.com/rcs/hero.png"),
					LogoURL:     telnyx.String("https://www.example.com/rcs/logo.png"),
					PhoneNumber: telnyx.AgentPhoneContactParam{
						Label:  "x",
						Number: "+49605132",
					},
					PrivacyPolicyURL:      telnyx.String("https://www.example.com/privacy"),
					TermsAndConditionsURL: telnyx.String("https://www.example.com/terms"),
					Website: telnyx.AgentWebsiteContactParam{
						Label: "x",
						URL:   "https://example.com",
					},
				},
			},
			Campaign: telnyx.AgentCampaignConfigurationParam{
				CompanyOverview:       "x",
				AdditionalInformation: telnyx.String("x"),
				AgentOverview:         telnyx.String("x"),
				ConsentSettings: telnyx.AgentConsentConfigurationParam{
					CallToAction: "x",
					DoubleOptIn:  true,
					HelpResponse: "x",
					OptInMessage: "x",
					OptInMethods: []telnyx.AgentConsentConfigurationOptInMethodParam{{
						MethodType:  "SMS",
						Description: telnyx.String("x"),
					}},
					OptOutResponse:       "x",
					CallToActionMediaURL: telnyx.String("https://example.com"),
					CallToActionURL:      telnyx.String("https://example.com"),
					DoubleOptInMessage:   telnyx.String("x"),
				},
				Interactions: []telnyx.AgentInteractionParam{{
					InteractionType: telnyx.AgentInteractionInteractionTypeTransactionalUpdates,
					Description:     telnyx.String("x"),
				}},
				MessageExamples: []string{"x"},
			},
			Testing: telnyx.AgentTestingConfigurationParam{
				TestURL:               "https://example.com",
				AdditionalInformation: telnyx.String("x"),
				MessageID:             telnyx.String("x"),
			},
		},
		DisplayName:    "Acme Order Updates",
		UseCase:        telnyx.AgentUseCaseTransactional,
		IdempotencyKey: "Idempotency-Key",
		HostingRegion:  telnyx.String("hosting_region"),
		ProfileID:      telnyx.String("profile_id"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcAgentGet(t *testing.T) {
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
	_, err := client.Rcs.Agents.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Agents.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.RcAgentUpdateParams{
			Configuration: telnyx.AgentConfigurationParam{
				Basics: telnyx.AgentConfigurationBasicsUnionParam{
					OfAgentConfigurationBasicsUnionMember0: &telnyx.AgentConfigurationBasicsUnionMember0Param{
						PhoneNumber: telnyx.AgentPhoneContactParam{
							Label:  "x",
							Number: "+49605132",
						},
						BrandColor:  telnyx.String("#2FDCd1"),
						Description: telnyx.String("x"),
						Email: telnyx.AgentEmailContactParam{
							Address: "dev@stainless.com",
							Label:   "x",
						},
						HeroURL:               telnyx.String("https://example.com"),
						LogoURL:               telnyx.String("https://example.com"),
						PrivacyPolicyURL:      telnyx.String("https://example.com"),
						TermsAndConditionsURL: telnyx.String("https://example.com"),
						Website: telnyx.AgentWebsiteContactParam{
							Label: "x",
							URL:   "https://example.com",
						},
					},
				},
				Campaign: telnyx.AgentCampaignConfigurationParam{
					CompanyOverview:       "x",
					AdditionalInformation: telnyx.String("x"),
					AgentOverview:         telnyx.String("x"),
					ConsentSettings: telnyx.AgentConsentConfigurationParam{
						CallToAction: "x",
						DoubleOptIn:  true,
						HelpResponse: "x",
						OptInMessage: "x",
						OptInMethods: []telnyx.AgentConsentConfigurationOptInMethodParam{{
							MethodType:  "SMS",
							Description: telnyx.String("x"),
						}},
						OptOutResponse:       "x",
						CallToActionMediaURL: telnyx.String("https://example.com"),
						CallToActionURL:      telnyx.String("https://example.com"),
						DoubleOptInMessage:   telnyx.String("x"),
					},
					Interactions: []telnyx.AgentInteractionParam{{
						InteractionType: telnyx.AgentInteractionInteractionTypeTransactionalUpdates,
						Description:     telnyx.String("x"),
					}},
					MessageExamples: []string{"x"},
				},
				Testing: telnyx.AgentTestingConfigurationParam{
					TestURL:               "https://example.com",
					AdditionalInformation: telnyx.String("x"),
					MessageID:             telnyx.String("x"),
				},
			},
			DisplayName:   telnyx.String("Acme Delivery Updates"),
			HostingRegion: telnyx.String("hosting_region"),
			ProfileID:     telnyx.String("profile_id"),
			UseCase:       telnyx.AgentUseCaseMultiUse,
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

func TestRcAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Agents.List(context.TODO(), telnyx.RcAgentListParams{
		BrandID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcAgentLaunchWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Agents.Launch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.RcAgentLaunchParams{
			Campaign: telnyx.RcAgentLaunchParamsCampaign{
				AgentCampaignConfigurationParam: telnyx.AgentCampaignConfigurationParam{
					CompanyOverview:       "Acme provides online retail services.",
					AdditionalInformation: telnyx.String("x"),
					AgentOverview:         telnyx.String("The agent sends order confirmations and delivery updates."),
					ConsentSettings: telnyx.AgentConsentConfigurationParam{
						CallToAction: "Select RCS updates during checkout.",
						DoubleOptIn:  false,
						HelpResponse: "Contact support@example.com for help.",
						OptInMessage: "You are subscribed to Acme order updates.",
						OptInMethods: []telnyx.AgentConsentConfigurationOptInMethodParam{{
							MethodType:  "WEBSITE",
							Description: telnyx.String("x"),
						}},
						OptOutResponse:       "You will receive no more messages.",
						CallToActionMediaURL: telnyx.String("https://www.example.com/rcs/opt-in.png"),
						CallToActionURL:      telnyx.String("https://www.example.com/checkout"),
						DoubleOptInMessage:   telnyx.String("x"),
					},
					Interactions: []telnyx.AgentInteractionParam{{
						InteractionType: telnyx.AgentInteractionInteractionTypeTransactionalUpdates,
						Description:     telnyx.String("x"),
					}},
					MessageExamples: []string{"Your Acme order is confirmed.", "Your Acme order has shipped.", "Your Acme order was delivered."},
				},
				AgentOverview: "The agent sends order confirmations and delivery updates.",
				ConsentSettings: telnyx.AgentConsentConfigurationParam{
					CallToAction: "Select RCS updates during checkout.",
					DoubleOptIn:  false,
					HelpResponse: "Contact support@example.com for help.",
					OptInMessage: "You are subscribed to Acme order updates.",
					OptInMethods: []telnyx.AgentConsentConfigurationOptInMethodParam{{
						MethodType:  "WEBSITE",
						Description: telnyx.String("x"),
					}},
					OptOutResponse:       "You will receive no more messages.",
					CallToActionMediaURL: telnyx.String("https://www.example.com/rcs/opt-in.png"),
					CallToActionURL:      telnyx.String("https://www.example.com/checkout"),
					DoubleOptInMessage:   telnyx.String("x"),
				},
				Interactions: []telnyx.AgentInteractionParam{{
					InteractionType: telnyx.AgentInteractionInteractionTypeTransactionalUpdates,
					Description:     telnyx.String("x"),
				}},
				MessageExamples: []string{"Your Acme order is confirmed.", "Your Acme order has shipped.", "Your Acme order was delivered."},
			},
			Testing: telnyx.AgentTestingConfigurationParam{
				TestURL:               "https://www.example.com/rcs/test-video",
				AdditionalInformation: telnyx.String("Demonstrates START, STOP, HELP, and an order-status interaction."),
				MessageID:             telnyx.String("x"),
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

func TestRcAgentGetCarrierApprovals(t *testing.T) {
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
	_, err := client.Rcs.Agents.GetCarrierApprovals(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcAgentSubmit(t *testing.T) {
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
	_, err := client.Rcs.Agents.Submit(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
