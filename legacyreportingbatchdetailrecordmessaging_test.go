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

func TestLegacyReportingBatchDetailRecordMessagingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Legacy.Reporting.BatchDetailRecords.Messaging.New(context.TODO(), telnyx.LegacyReportingBatchDetailRecordMessagingNewParams{
		EndTime:     time.Now(),
		StartTime:   time.Now(),
		Connections: []int64{123, 456},
		Directions:  []int64{1, 2},
		Filters: []telnyx.FilterParam{{
			BillingGroup: telnyx.String("adfaa016-f921-4b6c-97bb-e4c1dad231c5"),
			Cld:          telnyx.String("+13129457420"),
			CldFilter:    telnyx.FilterCldFilterContains,
			Cli:          telnyx.String("+13129457420"),
			CliFilter:    telnyx.FilterCliFilterContains,
			FilterType:   telnyx.FilterFilterTypeAnd,
			TagsList:     telnyx.String("tag1"),
		}},
		IncludeMessageBody:       telnyx.Bool(true),
		ManagedAccounts:          []string{"f47ac10b-58cc-4372-a567-0e02b2c3d479", "6ba7b810-9dad-11d1-80b4-00c04fd430c8"},
		Profiles:                 []string{"3fa85f64-5717-4562-b3fc-2c963f66afa6", "7d4e3f8a-9b2c-4e1d-8f5a-1a2b3c4d5e6f"},
		RecordTypes:              []int64{1, 2},
		ReportName:               telnyx.String("My MDR Report"),
		SelectAllManagedAccounts: telnyx.Bool(false),
		Timezone:                 telnyx.String("UTC"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegacyReportingBatchDetailRecordMessagingGet(t *testing.T) {
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
	_, err := client.Legacy.Reporting.BatchDetailRecords.Messaging.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegacyReportingBatchDetailRecordMessagingList(t *testing.T) {
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
	_, err := client.Legacy.Reporting.BatchDetailRecords.Messaging.List(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegacyReportingBatchDetailRecordMessagingDelete(t *testing.T) {
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
	_, err := client.Legacy.Reporting.BatchDetailRecords.Messaging.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
