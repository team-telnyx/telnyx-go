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

func TestDirPhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.PhoneNumbers.List(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirPhoneNumberListParams{
			PageNumber: telnyx.Int(1),
			PageSize:   telnyx.Int(20),
			Status:     telnyx.DirPhoneNumberListParamsStatusSubmitted,
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

func TestDirPhoneNumberAdd(t *testing.T) {
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
	_, err := client.Dir.PhoneNumbers.Add(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirPhoneNumberAddParams{
			Documents: []telnyx.DirPhoneNumberAddParamsDocument{{
				DocumentID:   "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
				DocumentType: "letter_of_authorization",
				Description:  telnyx.String("LOA authorising Telnyx to register these numbers under the DIR."),
			}},
			PhoneNumbers: []string{"+19493253498", "+12134445566"},
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

func TestDirPhoneNumberRemove(t *testing.T) {
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
	_, err := client.Dir.PhoneNumbers.Remove(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirPhoneNumberRemoveParams{
			PhoneNumbers: []string{"+19493253498"},
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
