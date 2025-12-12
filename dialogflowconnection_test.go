// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestDialogflowConnectionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DialogflowConnections.New(
		context.TODO(),
		"connection_id",
		telnyx.DialogflowConnectionNewParams{
			ServiceAccount: map[string]any{
				"type":                        "bar",
				"project_id":                  "bar",
				"private_key_id":              "bar",
				"private_key":                 "bar",
				"client_email":                "bar",
				"client_id":                   "bar",
				"auth_uri":                    "bar",
				"token_uri":                   "bar",
				"auth_provider_x509_cert_url": "bar",
				"client_x509_cert_url":        "bar",
			},
			ConversationProfileID: telnyx.String("a-VMHLWzTmKjiJw5S6O0-w"),
			DialogflowAPI:         telnyx.DialogflowConnectionNewParamsDialogflowAPICx,
			Environment:           telnyx.String("development"),
			Location:              telnyx.String("global"),
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

func TestDialogflowConnectionGet(t *testing.T) {
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
	_, err := client.DialogflowConnections.Get(context.TODO(), "connection_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDialogflowConnectionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.DialogflowConnections.Update(
		context.TODO(),
		"connection_id",
		telnyx.DialogflowConnectionUpdateParams{
			ServiceAccount: map[string]any{
				"type":                        "bar",
				"project_id":                  "bar",
				"private_key_id":              "bar",
				"private_key":                 "bar",
				"client_email":                "bar",
				"client_id":                   "bar",
				"auth_uri":                    "bar",
				"token_uri":                   "bar",
				"auth_provider_x509_cert_url": "bar",
				"client_x509_cert_url":        "bar",
			},
			ConversationProfileID: telnyx.String("a-VMHLWzTmKjiJw5S6O0-w"),
			DialogflowAPI:         telnyx.DialogflowConnectionUpdateParamsDialogflowAPICx,
			Environment:           telnyx.String("development"),
			Location:              telnyx.String("global"),
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

func TestDialogflowConnectionDelete(t *testing.T) {
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
	err := client.DialogflowConnections.Delete(context.TODO(), "connection_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
