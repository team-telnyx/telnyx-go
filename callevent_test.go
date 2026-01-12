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

func TestCallEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.CallEvents.List(context.TODO(), telnyx.CallEventListParams{
		Filter: telnyx.CallEventListParamsFilter{
			ApplicationName: telnyx.CallEventListParamsFilterApplicationName{
				Contains: telnyx.String("contains"),
			},
			ApplicationSessionID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ConnectionID:         telnyx.String("connection_id"),
			Failed:               telnyx.Bool(false),
			From:                 telnyx.String("+12025550142"),
			LegID:                telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                 telnyx.String("name"),
			OccurredAt: telnyx.CallEventListParamsFilterOccurredAt{
				Eq:  telnyx.String("2019-03-29T11:10:00Z"),
				Gt:  telnyx.String("2019-03-29T11:10:00Z"),
				Gte: telnyx.String("2019-03-29T11:10:00Z"),
				Lt:  telnyx.String("2019-03-29T11:10:00Z"),
				Lte: telnyx.String("2019-03-29T11:10:00Z"),
			},
			OutboundOutboundVoiceProfileID: telnyx.String("outbound.outbound_voice_profile_id"),
			Product:                        "texml",
			Status:                         "init",
			To:                             telnyx.String("+12025550142"),
			Type:                           "webhook",
		},
		Page: telnyx.CallEventListParamsPage{
			After:  telnyx.String("after"),
			Before: telnyx.String("before"),
			Limit:  telnyx.Int(1),
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
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
