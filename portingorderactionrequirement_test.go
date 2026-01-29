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

func TestPortingOrderActionRequirementListWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.ActionRequirements.List(
		context.TODO(),
		"porting_order_id",
		telnyx.PortingOrderActionRequirementListParams{
			Filter: telnyx.PortingOrderActionRequirementListParamsFilter{
				ID:                []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
				ActionType:        "au_id_verification",
				RequirementTypeID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Status:            "created",
			},
			Page: telnyx.PortingOrderActionRequirementListParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
			},
			Sort: telnyx.PortingOrderActionRequirementListParamsSort{
				Value: "created_at",
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

func TestPortingOrderActionRequirementInitiate(t *testing.T) {
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
	_, err := client.PortingOrders.ActionRequirements.Initiate(
		context.TODO(),
		"id",
		telnyx.PortingOrderActionRequirementInitiateParams{
			PortingOrderID: "porting_order_id",
			Params: telnyx.PortingOrderActionRequirementInitiateParamsParams{
				FirstName: "John",
				LastName:  "Doe",
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
