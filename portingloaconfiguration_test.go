// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestPortingLoaConfigurationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Porting.LoaConfigurations.New(context.TODO(), telnyx.PortingLoaConfigurationNewParams{
		Address: telnyx.PortingLoaConfigurationNewParamsAddress{
			City:            "Austin",
			CountryCode:     "US",
			State:           "TX",
			StreetAddress:   "600 Congress Avenue",
			ZipCode:         "78701",
			ExtendedAddress: telnyx.String("14th Floor"),
		},
		CompanyName: "Telnyx",
		Contact: telnyx.PortingLoaConfigurationNewParamsContact{
			Email:       "testing@telnyx.com",
			PhoneNumber: "+12003270001",
		},
		Logo: telnyx.PortingLoaConfigurationNewParamsLogo{
			DocumentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
		Name: "My LOA Configuration",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingLoaConfigurationGet(t *testing.T) {
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
	_, err := client.Porting.LoaConfigurations.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingLoaConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Porting.LoaConfigurations.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingLoaConfigurationUpdateParams{
			Address: telnyx.PortingLoaConfigurationUpdateParamsAddress{
				City:            "Austin",
				CountryCode:     "US",
				State:           "TX",
				StreetAddress:   "600 Congress Avenue",
				ZipCode:         "78701",
				ExtendedAddress: telnyx.String("14th Floor"),
			},
			CompanyName: "Telnyx",
			Contact: telnyx.PortingLoaConfigurationUpdateParamsContact{
				Email:       "testing@telnyx.com",
				PhoneNumber: "+12003270001",
			},
			Logo: telnyx.PortingLoaConfigurationUpdateParamsLogo{
				DocumentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			},
			Name: "My LOA Configuration",
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

func TestPortingLoaConfigurationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Porting.LoaConfigurations.List(context.TODO(), telnyx.PortingLoaConfigurationListParams{
		Page: telnyx.PortingLoaConfigurationListParamsPage{
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

func TestPortingLoaConfigurationDelete(t *testing.T) {
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
	err := client.Porting.LoaConfigurations.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingLoaConfigurationPreview0WithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Porting.LoaConfigurations.Preview0(context.TODO(), telnyx.PortingLoaConfigurationPreview0Params{
		Address: telnyx.PortingLoaConfigurationPreview0ParamsAddress{
			City:            "Austin",
			CountryCode:     "US",
			State:           "TX",
			StreetAddress:   "600 Congress Avenue",
			ZipCode:         "78701",
			ExtendedAddress: telnyx.String("14th Floor"),
		},
		CompanyName: "Telnyx",
		Contact: telnyx.PortingLoaConfigurationPreview0ParamsContact{
			Email:       "testing@telnyx.com",
			PhoneNumber: "+12003270001",
		},
		Logo: telnyx.PortingLoaConfigurationPreview0ParamsLogo{
			DocumentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		},
		Name: "My LOA Configuration",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestPortingLoaConfigurationPreview1(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Porting.LoaConfigurations.Preview1(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
