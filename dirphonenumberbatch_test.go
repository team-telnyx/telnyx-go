// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/v4/option"
)

func TestDirPhoneNumberBatchGet(t *testing.T) {
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
	_, err := client.Dir.PhoneNumberBatches.Get(
		context.TODO(),
		"0a4b1f5e-2f12-4c0c-9a98-9b3a7d8b8e62",
		telnyx.DirPhoneNumberBatchGetParams{
			DirID: "16635d38-75a6-4481-82e8-69af60e05011",
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

func TestDirPhoneNumberBatchListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.PhoneNumberBatches.List(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirPhoneNumberBatchListParams{
			FilterStatus: telnyx.DirPhoneNumberBatchListParamsFilterStatusSubmitted,
			PageNumber:   telnyx.Int(1),
			PageSize:     telnyx.Int(20),
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
