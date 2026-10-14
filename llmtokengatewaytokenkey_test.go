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

func TestLlmTokenGatewayTokenKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenKeys.New(context.TODO(), telnyx.LlmTokenGatewayTokenKeyNewParams{
		Name:                "support-service",
		TokenGroupID:        "11111111-1111-4111-8111-111111111111",
		IdempotencyKey:      "Idempotency-Key",
		AllowedModels:       []string{"J!Q0Ok0bzJb7"},
		Blocked:             telnyx.Bool(true),
		BudgetDuration:      telnyx.BudgetDuration1d,
		ExpiresAt:           telnyx.Time(time.Now()),
		MaxBudget:           telnyx.Float(0.000001),
		MaxParallelRequests: telnyx.Int(5),
		RequiredEndUserID:   telnyx.Bool(true),
		RpmLimit:            telnyx.Int(1),
		TokenUserID:         telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		TpmLimit:            telnyx.Int(1),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenKeyGet(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenKeys.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenKeys.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.LlmTokenGatewayTokenKeyUpdateParams{
			IdempotencyKey:      "Idempotency-Key",
			IfMatch:             "If-Match",
			AllowedModels:       []string{"J!Q0Ok0bzJb7"},
			Blocked:             telnyx.Bool(true),
			BudgetDuration:      telnyx.BudgetDuration1d,
			ExpiresAt:           telnyx.Time(time.Now()),
			MaxBudget:           telnyx.Float(0.000001),
			MaxParallelRequests: telnyx.Int(5),
			Name:                telnyx.String("name"),
			RequiredEndUserID:   telnyx.Bool(true),
			RpmLimit:            telnyx.Int(1),
			TpmLimit:            telnyx.Int(1),
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

func TestLlmTokenGatewayTokenKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenKeys.List(context.TODO(), telnyx.LlmTokenGatewayTokenKeyListParams{
		PageNumber:   telnyx.Int(1),
		PageSize:     telnyx.Int(1),
		PageSnapshot: telnyx.String("page[snapshot]"),
		TokenGroupID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		TokenUserID:  telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenKeyDelete(t *testing.T) {
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
	err := client.LlmTokenGateway.TokenKeys.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.LlmTokenGatewayTokenKeyDeleteParams{
			IdempotencyKey: "Idempotency-Key",
			IfMatch:        "If-Match",
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
