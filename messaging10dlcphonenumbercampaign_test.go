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

func TestMessaging10dlcPhoneNumberCampaignNew(t *testing.T) {
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
	_, err := client.Messaging10dlc.PhoneNumberCampaigns.New(context.TODO(), telnyx.Messaging10dlcPhoneNumberCampaignNewParams{
		PhoneNumberCampaignCreate: telnyx.PhoneNumberCampaignCreateParam{
			CampaignID:  "4b300178-131c-d902-d54e-72d90ba1620j",
			PhoneNumber: "+18005550199",
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

func TestMessaging10dlcPhoneNumberCampaignGet(t *testing.T) {
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
	_, err := client.Messaging10dlc.PhoneNumberCampaigns.Get(context.TODO(), "phoneNumber")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcPhoneNumberCampaignUpdate(t *testing.T) {
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
	_, err := client.Messaging10dlc.PhoneNumberCampaigns.Update(
		context.TODO(),
		"phoneNumber",
		telnyx.Messaging10dlcPhoneNumberCampaignUpdateParams{
			PhoneNumberCampaignCreate: telnyx.PhoneNumberCampaignCreateParam{
				CampaignID:  "4b300178-131c-d902-d54e-72d90ba1620j",
				PhoneNumber: "+18005550199",
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

func TestMessaging10dlcPhoneNumberCampaignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging10dlc.PhoneNumberCampaigns.List(context.TODO(), telnyx.Messaging10dlcPhoneNumberCampaignListParams{
		Filter: telnyx.Messaging10dlcPhoneNumberCampaignListParamsFilter{
			TcrBrandID:       telnyx.String("BRANDID"),
			TcrCampaignID:    telnyx.String("CAMPID3"),
			TelnyxBrandID:    telnyx.String("f3575e15-32ce-400e-a4c0-dd78800c20b0"),
			TelnyxCampaignID: telnyx.String("f3575e15-32ce-400e-a4c0-dd78800c20b0"),
		},
		Page:           telnyx.Int(0),
		RecordsPerPage: telnyx.Int(0),
		Sort:           telnyx.Messaging10dlcPhoneNumberCampaignListParamsSortAssignmentStatus,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcPhoneNumberCampaignDelete(t *testing.T) {
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
	_, err := client.Messaging10dlc.PhoneNumberCampaigns.Delete(context.TODO(), "phoneNumber")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
