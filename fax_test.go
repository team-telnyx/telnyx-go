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

func TestFaxNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Faxes.New(context.TODO(), telnyx.FaxNewParams{
		ConnectionID:    "234423",
		From:            "+13125790015",
		To:              "+13127367276",
		ClientState:     telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
		FromDisplayName: telnyx.String("Company Name"),
		MediaName:       telnyx.String("my_media_uploaded_to_media_storage_api"),
		MediaURL:        telnyx.String("https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"),
		Monochrome:      telnyx.Bool(true),
		PreviewFormat:   telnyx.FaxNewParamsPreviewFormatPdf,
		Quality:         telnyx.FaxNewParamsQualityHigh,
		StoreMedia:      telnyx.Bool(true),
		StorePreview:    telnyx.Bool(true),
		T38Enabled:      telnyx.Bool(true),
		WebhookURL:      telnyx.String("https://www.example.com/server-b/"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxGet(t *testing.T) {
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
	_, err := client.Faxes.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFaxListWithOptionalParams(t *testing.T) {
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
	_, err := client.Faxes.List(context.TODO(), telnyx.FaxListParams{
		Filter: telnyx.FaxListParamsFilter{
			CreatedAt: telnyx.FaxListParamsFilterCreatedAt{
				Gt:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lt:  telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			Direction: telnyx.FaxListParamsFilterDirection{
				Eq: telnyx.String("inbound"),
			},
			From: telnyx.FaxListParamsFilterFrom{
				Eq: telnyx.String("+13127367276"),
			},
			To: telnyx.FaxListParamsFilterTo{
				Eq: telnyx.String("+13127367276"),
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

func TestFaxDelete(t *testing.T) {
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
	err := client.Faxes.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
