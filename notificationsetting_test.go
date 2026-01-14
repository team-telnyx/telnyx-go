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

func TestNotificationSettingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.NotificationSettings.New(context.TODO(), telnyx.NotificationSettingNewParams{
		NotificationSetting: telnyx.NotificationSettingParam{
			NotificationChannelID:        telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			NotificationEventConditionID: telnyx.String("70c7c5cb-dce2-4124-accb-870d39dbe852"),
			NotificationProfileID:        telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			Parameters: []telnyx.NotificationSettingParameterParam{{
				Name:  telnyx.String("phone_number"),
				Value: telnyx.String("+13125550000"),
			}},
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

func TestNotificationSettingGet(t *testing.T) {
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
	_, err := client.NotificationSettings.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationSettingListWithOptionalParams(t *testing.T) {
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
	_, err := client.NotificationSettings.List(context.TODO(), telnyx.NotificationSettingListParams{
		Filter: telnyx.NotificationSettingListParamsFilter{
			AssociatedRecordType: telnyx.NotificationSettingListParamsFilterAssociatedRecordType{
				Eq: "phone_number",
			},
			ChannelTypeID: telnyx.NotificationSettingListParamsFilterChannelTypeID{
				Eq: "webhook",
			},
			NotificationChannel: telnyx.NotificationSettingListParamsFilterNotificationChannel{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			NotificationEventConditionID: telnyx.NotificationSettingListParamsFilterNotificationEventConditionID{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			NotificationProfileID: telnyx.NotificationSettingListParamsFilterNotificationProfileID{
				Eq: telnyx.String("12455643-3cf1-4683-ad23-1cd32f7d5e0a"),
			},
			Status: telnyx.NotificationSettingListParamsFilterStatus{
				Eq: "enable-received",
			},
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationSettingDelete(t *testing.T) {
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
	_, err := client.NotificationSettings.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
