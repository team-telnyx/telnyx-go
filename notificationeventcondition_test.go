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

func TestNotificationEventConditionListWithOptionalParams(t *testing.T) {
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
	_, err := client.NotificationEventConditions.List(context.TODO(), telnyx.NotificationEventConditionListParams{
		Filter: telnyx.NotificationEventConditionListParamsFilter{
			AssociatedRecordType: telnyx.NotificationEventConditionListParamsFilterAssociatedRecordType{
				Eq: "phone_number",
			},
			ChannelTypeID: telnyx.NotificationEventConditionListParamsFilterChannelTypeID{
				Eq: "webhook",
			},
			NotificationChannel: telnyx.NotificationEventConditionListParamsFilterNotificationChannel{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			NotificationEventConditionID: telnyx.NotificationEventConditionListParamsFilterNotificationEventConditionID{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			NotificationProfileID: telnyx.NotificationEventConditionListParamsFilterNotificationProfileID{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			Status: telnyx.NotificationEventConditionListParamsFilterStatus{
				Eq: "enable-received",
			},
		},
		Page: telnyx.NotificationEventConditionListParamsPage{
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
