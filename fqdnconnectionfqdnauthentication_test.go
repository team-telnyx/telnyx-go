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

func TestFqdnConnectionFqdnAuthenticationList(t *testing.T) {
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
	_, err := client.FqdnConnections.FqdnAuthentication.List(context.TODO(), "fqdn_connection_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFqdnConnectionFqdnAuthenticationPatchAllWithOptionalParams(t *testing.T) {
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
	_, err := client.FqdnConnections.FqdnAuthentication.PatchAll(
		context.TODO(),
		"fqdn_connection_id",
		telnyx.FqdnConnectionFqdnAuthenticationPatchAllParams{
			FailoverURL:                telnyx.String("https://failover.example.com"),
			FqdnOutboundAuthentication: telnyx.FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthenticationIPAuthentication,
			IPAuthenticationMethod:     telnyx.FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethodPChargeInfo,
			Password:                   telnyx.String("new_password"),
			TxtName:                    telnyx.String("new_txt_name"),
			TxtTtl:                     telnyx.Int(300),
			TxtValue:                   telnyx.String("new_txt_value"),
			UserName:                   telnyx.String("newusername"),
			WebhookURL:                 telnyx.String("https://example.com"),
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
