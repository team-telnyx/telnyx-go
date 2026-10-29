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

func TestPayConnectorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PayConnectors.New(context.TODO(), telnyx.PayConnectorNewParams{
		Config: telnyx.GenericPayConnectorConfigUnionParam{
			OfNoProcessorAuthentication: &telnyx.GenericPayConnectorConfigNoProcessorAuthenticationParam{
				Token:       telnyx.String("token"),
				EndpointURL: telnyx.String("https://processor.example.com/pay"),
				Password:    telnyx.String("password"),
				Timeout:     telnyx.Int(1),
				Username:    telnyx.String("username"),
			},
		},
		Name: "my-payment-connector",
		Mode: telnyx.PayConnectorNewParamsModeTest,
		Type: telnyx.PayConnectorNewParamsTypeGeneric,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPayConnectorGet(t *testing.T) {
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
	_, err := client.PayConnectors.Get(context.TODO(), "my-payment-connector")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPayConnectorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PayConnectors.Update(
		context.TODO(),
		"my-payment-connector",
		telnyx.PayConnectorUpdateParams{
			Config: telnyx.GenericPayConnectorConfigUnionParam{
				OfBasicAuthenticationAlsoTheValidationDefault: &telnyx.GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam{
					Password:    "P",
					Username:    "P",
					Token:       telnyx.String("token"),
					AuthType:    "basic",
					EndpointURL: telnyx.String("https://processor.example.com/pay"),
					Timeout:     telnyx.Int(1),
				},
			},
			Mode: telnyx.PayConnectorUpdateParamsModeLive,
			Type: telnyx.PayConnectorUpdateParamsTypeGeneric,
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

func TestPayConnectorListWithOptionalParams(t *testing.T) {
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
	_, err := client.PayConnectors.List(context.TODO(), telnyx.PayConnectorListParams{
		FilterCreatedAtEq:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterCreatedAtGt:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterCreatedAtGte: telnyx.String("2026-10-01T00:00:00Z"),
		FilterCreatedAtLt:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterCreatedAtLte: telnyx.String("2026-10-01T00:00:00Z"),
		FilterMode:         telnyx.String("test"),
		FilterName:         telnyx.String("my-payment-connector"),
		FilterType:         telnyx.String("generic"),
		FilterUpdatedAtEq:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterUpdatedAtGt:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterUpdatedAtGte: telnyx.String("2026-10-01T00:00:00Z"),
		FilterUpdatedAtLt:  telnyx.String("2026-10-01T00:00:00Z"),
		FilterUpdatedAtLte: telnyx.String("2026-10-01T00:00:00Z"),
		PageAfter:          telnyx.String("page[after]"),
		PageBefore:         telnyx.String("page[before]"),
		PageSize:           telnyx.Int(1),
		Sort:               telnyx.String("-created_at"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPayConnectorDelete(t *testing.T) {
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
	err := client.PayConnectors.Delete(context.TODO(), "my-payment-connector")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
