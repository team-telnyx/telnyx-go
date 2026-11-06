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

func TestAIToolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tools.New(context.TODO(), telnyx.AIToolNewParams{
		DisplayName: "display_name",
		Type:        "type",
		ClientSideTool: map[string]any{
			"foo": "bar",
		},
		Function: map[string]any{
			"foo": "bar",
		},
		Handoff: map[string]any{
			"foo": "bar",
		},
		Invite: map[string]any{
			"foo": "bar",
		},
		Pay: telnyx.PayToolParams{
			ConnectorName: "connector_name",
			ConnectorParameters: map[string]any{
				"campaignid":   "bar",
				"customer_ref": "bar",
			},
			Currency:                telnyx.String("currency"),
			Description:             telnyx.String("description"),
			InterDigitTimeoutMillis: telnyx.Int(1),
			MaxAttempts:             telnyx.Int(1),
			MinPostalCodeLength:     telnyx.Int(5),
			PaymentMethod:           telnyx.String("payment_method"),
			PostalCode: telnyx.PayToolParamsPostalCodeUnion{
				OfBool: telnyx.Bool(false),
			},
			Prompts: telnyx.PayToolParamsPrompts{
				BankAccountNumber: telnyx.PayToolParamsPromptsBankAccountNumberUnion{
					OfString: telnyx.String("x"),
				},
				BankRoutingNumber: telnyx.PayToolParamsPromptsBankRoutingNumberUnion{
					OfString: telnyx.String("x"),
				},
				ExpirationDate: telnyx.PayToolParamsPromptsExpirationDateUnion{
					OfString: telnyx.String("x"),
				},
				PaymentCardNumber: telnyx.PayToolParamsPromptsPaymentCardNumberUnion{
					OfString: telnyx.String("x"),
				},
				PostalCode: telnyx.PayToolParamsPromptsPostalCodeUnion{
					OfString: telnyx.String("x"),
				},
				SecurityCode: telnyx.PayToolParamsPromptsSecurityCodeUnion{
					OfString: telnyx.String("x"),
				},
			},
			TimeoutMillis:  telnyx.Int(12000),
			ValidCardTypes: []string{"visa", "mastercard"},
		},
		Retrieval: map[string]any{
			"foo": "bar",
		},
		TimeoutMs: telnyx.Int(0),
		UpdateDynamicVariables: telnyx.UpdateDynamicVariablesToolParams{
			Description: "Collect caller details into conversation variables.",
			Name:        "collect_details",
			UpdatableVariables: []telnyx.UpdateDynamicVariablesToolParamsUpdatableVariable{{
				Name:        "customer_name",
				Description: telnyx.String("The caller's full name."),
				Type:        telnyx.String("string"),
			}},
		},
		Vision: telnyx.VisionToolParams{
			MaxFrameAgeSecs: telnyx.Int(1),
		},
		Webhook: map[string]any{
			"foo": "bar",
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

func TestAIToolGet(t *testing.T) {
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
	_, err := client.AI.Tools.Get(context.TODO(), "tool_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIToolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tools.Update(
		context.TODO(),
		"tool_id",
		telnyx.AIToolUpdateParams{
			ClientSideTool: map[string]any{
				"foo": "bar",
			},
			DisplayName: telnyx.String("display_name"),
			Function: map[string]any{
				"foo": "bar",
			},
			Handoff: map[string]any{
				"foo": "bar",
			},
			Invite: map[string]any{
				"foo": "bar",
			},
			Pay: telnyx.PayToolParams{
				ConnectorName: "connector_name",
				ConnectorParameters: map[string]any{
					"campaignid":   "bar",
					"customer_ref": "bar",
				},
				Currency:                telnyx.String("currency"),
				Description:             telnyx.String("description"),
				InterDigitTimeoutMillis: telnyx.Int(1),
				MaxAttempts:             telnyx.Int(1),
				MinPostalCodeLength:     telnyx.Int(5),
				PaymentMethod:           telnyx.String("payment_method"),
				PostalCode: telnyx.PayToolParamsPostalCodeUnion{
					OfBool: telnyx.Bool(false),
				},
				Prompts: telnyx.PayToolParamsPrompts{
					BankAccountNumber: telnyx.PayToolParamsPromptsBankAccountNumberUnion{
						OfString: telnyx.String("x"),
					},
					BankRoutingNumber: telnyx.PayToolParamsPromptsBankRoutingNumberUnion{
						OfString: telnyx.String("x"),
					},
					ExpirationDate: telnyx.PayToolParamsPromptsExpirationDateUnion{
						OfString: telnyx.String("x"),
					},
					PaymentCardNumber: telnyx.PayToolParamsPromptsPaymentCardNumberUnion{
						OfString: telnyx.String("x"),
					},
					PostalCode: telnyx.PayToolParamsPromptsPostalCodeUnion{
						OfString: telnyx.String("x"),
					},
					SecurityCode: telnyx.PayToolParamsPromptsSecurityCodeUnion{
						OfString: telnyx.String("x"),
					},
				},
				TimeoutMillis:  telnyx.Int(12000),
				ValidCardTypes: []string{"visa", "mastercard"},
			},
			Retrieval: map[string]any{
				"foo": "bar",
			},
			TimeoutMs: telnyx.Int(0),
			Type:      telnyx.String("type"),
			UpdateDynamicVariables: telnyx.UpdateDynamicVariablesToolParams{
				Description: "Collect caller details into conversation variables.",
				Name:        "collect_details",
				UpdatableVariables: []telnyx.UpdateDynamicVariablesToolParamsUpdatableVariable{{
					Name:        "customer_name",
					Description: telnyx.String("The caller's full name."),
					Type:        telnyx.String("string"),
				}},
			},
			Vision: telnyx.VisionToolParams{
				MaxFrameAgeSecs: telnyx.Int(1),
			},
			Webhook: map[string]any{
				"foo": "bar",
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

func TestAIToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tools.List(context.TODO(), telnyx.AIToolListParams{
		FilterName: telnyx.String("filter[name]"),
		FilterType: telnyx.String("filter[type]"),
		PageNumber: telnyx.Int(1),
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

func TestAIToolDelete(t *testing.T) {
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
	err := client.AI.Tools.Delete(context.TODO(), "tool_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
