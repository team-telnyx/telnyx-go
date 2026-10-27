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

func TestAITestScheduleV1CronJobGet(t *testing.T) {
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
	_, err := client.AI.Tests.Schedules.V1.CronJobs.Get(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestScheduleV1CronJobUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.Schedules.V1.CronJobs.Update(
		context.TODO(),
		0,
		telnyx.AITestScheduleV1CronJobUpdateParams{
			Agent:             telnyx.Int(0),
			CrontabExpression: telnyx.String("0 4 * * *"),
			ExecutionMode:     telnyx.String("execution_mode"),
			Frequency:         telnyx.Int(0),
			IsActive:          telnyx.Bool(true),
			Name:              telnyx.String("Nightly regression - updated"),
			NotifyOn:          telnyx.AITestScheduleV1CronJobUpdateParamsNotifyOnBoth,
			Personalities:     []int64{0},
			RunAsText:         telnyx.Bool(true),
			Scenarios:         []int64{0},
			Tags:              []string{"string"},
			Timezone:          telnyx.String("timezone"),
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

func TestAITestScheduleV1CronJobDelete(t *testing.T) {
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
	_, err := client.AI.Tests.Schedules.V1.CronJobs.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestScheduleV1CronJobCronJobsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.Schedules.V1.CronJobs.CronJobs(context.TODO(), telnyx.AITestScheduleV1CronJobCronJobsParams{
		CrontabExpression: "0 3 * * *",
		Agent:             telnyx.Int(1259),
		ExecutionMode:     telnyx.String("voice"),
		Frequency:         telnyx.Int(0),
		IsActive:          telnyx.Bool(true),
		Name:              telnyx.String("Nightly regression"),
		NotifyOn:          telnyx.AITestScheduleV1CronJobCronJobsParamsNotifyOnFailure,
		Personalities:     []int64{0},
		RunAsText:         telnyx.Bool(true),
		Scenarios:         []int64{9512},
		Tags:              []string{"string"},
		Timezone:          telnyx.String("America/Chicago"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestScheduleV1CronJobGetCronJobsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.Schedules.V1.CronJobs.GetCronJobs(context.TODO(), telnyx.AITestScheduleV1CronJobGetCronJobsParams{
		AgentID: telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
