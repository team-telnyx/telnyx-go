// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestEnterpriseReputationLoaUpdate(t *testing.T) {
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
	_, err := client.Enterprises.Reputation.Loa.Update(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseReputationLoaUpdateParams{
			LoaDocumentID: "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
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

func TestEnterpriseReputationLoaRenderWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Enterprises.Reputation.Loa.Render(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseReputationLoaRenderParams{
			Agent: telnyx.AgentInputParam{
				AdministrativeArea: "administrative_area",
				City:               "city",
				ContactEmail:       "dev@stainless.com",
				ContactName:        "contact_name",
				ContactPhone:       "+13125550000",
				ContactTitle:       "contact_title",
				Country:            "US",
				LegalName:          "legal_name",
				PostalCode:         "postal_code",
				StreetAddress:      "street_address",
				Dba:                telnyx.String("dba"),
				ExtendedAddress:    telnyx.String("extended_address"),
			},
			Signature: telnyx.EnterpriseReputationLoaRenderParamsSignature{
				ImageBase64: "image_base64",
				SignerName:  telnyx.String("signer_name"),
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
