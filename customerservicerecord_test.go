// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestCustomerServiceRecordNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomerServiceRecords.New(context.TODO(), telnyx.CustomerServiceRecordNewParams{
		PhoneNumber: "+13035553000",
		AdditionalData: telnyx.CustomerServiceRecordNewParamsAdditionalData{
			AccountNumber:        telnyx.String("123456789"),
			AddressLine1:         telnyx.String("123 Main St"),
			AuthorizedPersonName: telnyx.String("John Doe"),
			BillingPhoneNumber:   telnyx.String("+12065551212"),
			City:                 telnyx.String("New York"),
			CustomerCode:         telnyx.String("123456789"),
			Name:                 telnyx.String("Entity Inc."),
			Pin:                  telnyx.String("1234"),
			State:                telnyx.String("NY"),
			ZipCode:              telnyx.String("10001"),
		},
		WebhookURL: telnyx.String("https://example.com/webhook"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerServiceRecordGet(t *testing.T) {
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
	_, err := client.CustomerServiceRecords.Get(context.TODO(), "customer_service_record_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomerServiceRecordListWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomerServiceRecords.List(context.TODO(), telnyx.CustomerServiceRecordListParams{
		Filter: telnyx.CustomerServiceRecordListParamsFilter{
			CreatedAt: telnyx.CustomerServiceRecordListParamsFilterCreatedAt{
				Gt: telnyx.Time(time.Now()),
				Lt: telnyx.Time(time.Now()),
			},
			PhoneNumber: telnyx.CustomerServiceRecordListParamsFilterPhoneNumber{
				Eq: telnyx.String("+12441239999"),
				In: []string{"+12441239999"},
			},
			Status: telnyx.CustomerServiceRecordListParamsFilterStatus{
				Eq: "pending",
				In: []string{"pending"},
			},
		},
		Page: telnyx.CustomerServiceRecordListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.CustomerServiceRecordListParamsSort{
			Value: "created_at",
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

func TestCustomerServiceRecordVerifyPhoneNumberCoverage(t *testing.T) {
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
	_, err := client.CustomerServiceRecords.VerifyPhoneNumberCoverage(context.TODO(), telnyx.CustomerServiceRecordVerifyPhoneNumberCoverageParams{
		PhoneNumbers: []string{"+13035553000"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
