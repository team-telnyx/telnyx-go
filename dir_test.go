// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestDirGet(t *testing.T) {
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
	_, err := client.Dir.Get(context.TODO(), "16635d38-75a6-4481-82e8-69af60e05011")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.Update(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirUpdateParams{
			AuthorizerEmail: telnyx.String("dev@stainless.com"),
			AuthorizerName:  telnyx.String("authorizer_name"),
			CallReasons:     []string{"Appointment reminders", "Billing inquiries", "Lab results"},
			DisplayName:     telnyx.String("Acme Plumbing & Wellness"),
			LogoURL:         telnyx.String("https://acmeplumbing.example.com/logo-v2-256.bmp"),
			Reselling:       telnyx.Bool(true),
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

func TestDirListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.List(context.TODO(), telnyx.DirListParams{
		FilterCallReasonContains:  telnyx.String("filter[call_reason][contains]"),
		FilterDisplayNameContains: telnyx.String("filter[display_name][contains]"),
		FilterEnterpriseID:        telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FilterExpiringAtGte:       telnyx.Time(time.Now()),
		FilterExpiringAtLte:       telnyx.Time(time.Now()),
		FilterStatus:              telnyx.DirListParamsFilterStatusDraft,
		PageNumber:                telnyx.Int(1),
		PageSize:                  telnyx.Int(20),
		Sort:                      telnyx.DirListParamsSortCreatedAt,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirDelete(t *testing.T) {
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
	err := client.Dir.Delete(context.TODO(), "16635d38-75a6-4481-82e8-69af60e05011")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirListDocumentTypes(t *testing.T) {
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
	_, err := client.Dir.ListDocumentTypes(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirListInfringementClaimsWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.ListInfringementClaims(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirListInfringementClaimsParams{
			PageNumber: telnyx.Int(1),
			PageSize:   telnyx.Int(20),
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

func TestDirSubmit(t *testing.T) {
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
	_, err := client.Dir.Submit(context.TODO(), "16635d38-75a6-4481-82e8-69af60e05011")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDirUpdateInfringementWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.UpdateInfringement(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirUpdateInfringementParams{
			CertifyBrandIsAccurate:      true,
			CertifyIPOwnership:          true,
			CertifyNoInfringement:       true,
			CertifyNoShaftContent:       true,
			InfringementResolutionNotes: "Updated the display name to remove the disputed mark and re-uploaded the authorization.",
			CallReasons:                 []string{"string"},
			DisplayName:                 telnyx.String("x"),
			Documents: []telnyx.DirUpdateInfringementParamsDocument{{
				DocumentID:   "2a7e8337-e803-4057-a4ae-26c40eb0bc6c",
				DocumentType: "business_registration",
				Description:  telnyx.String("Certificate of incorporation."),
			}},
			LogoURL: telnyx.String("logo_url"),
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
