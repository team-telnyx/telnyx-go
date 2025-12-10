// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestMessaging10dlcBrandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.New(context.TODO(), telnyx.Messaging10dlcBrandNewParams{
		Country:              "US",
		DisplayName:          "ABC Mobile",
		Email:                "email",
		EntityType:           telnyx.EntityTypePrivateProfit,
		Vertical:             telnyx.VerticalTechnology,
		BusinessContactEmail: telnyx.String("name@example.com"),
		City:                 telnyx.String("New York"),
		CompanyName:          telnyx.String("ABC Inc."),
		Ein:                  telnyx.String("111111111"),
		FirstName:            telnyx.String("John"),
		IPAddress:            telnyx.String("ipAddress"),
		IsReseller:           telnyx.Bool(true),
		LastName:             telnyx.String("Smith"),
		MobilePhone:          telnyx.String("+12024567890"),
		Mock:                 telnyx.Bool(true),
		Phone:                telnyx.String("+12024567890"),
		PostalCode:           telnyx.String("10001"),
		State:                telnyx.String("NY"),
		StockExchange:        telnyx.StockExchangeNasdaq,
		StockSymbol:          telnyx.String("ABC"),
		Street:               telnyx.String("123"),
		WebhookFailoverURL:   telnyx.String("https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6"),
		WebhookURL:           telnyx.String("https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93"),
		Website:              telnyx.String("http://www.abcmobile.com"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandGet(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.Get(context.TODO(), "brandId")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.Update(
		context.TODO(),
		"brandId",
		telnyx.Messaging10dlcBrandUpdateParams{
			Country:              "US",
			DisplayName:          "ABC Mobile",
			Email:                "email",
			EntityType:           telnyx.EntityTypePrivateProfit,
			Vertical:             telnyx.VerticalTechnology,
			AltBusinessID:        telnyx.String("altBusiness_id"),
			AltBusinessIDType:    telnyx.AltBusinessIDTypeNone,
			BusinessContactEmail: telnyx.String("name@example.com"),
			City:                 telnyx.String("New York"),
			CompanyName:          telnyx.String("ABC Inc."),
			Ein:                  telnyx.String("111111111"),
			FirstName:            telnyx.String("John"),
			IdentityStatus:       telnyx.BrandIdentityStatusVerified,
			IPAddress:            telnyx.String("ipAddress"),
			IsReseller:           telnyx.Bool(true),
			LastName:             telnyx.String("Smith"),
			Phone:                telnyx.String("+12024567890"),
			PostalCode:           telnyx.String("10001"),
			State:                telnyx.String("NY"),
			StockExchange:        telnyx.StockExchangeNasdaq,
			StockSymbol:          telnyx.String("ABC"),
			Street:               telnyx.String("123"),
			WebhookFailoverURL:   telnyx.String("https://webhook.com/9010a453-4df8-4be6-a551-1070892888d6"),
			WebhookURL:           telnyx.String("https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93"),
			Website:              telnyx.String("http://www.abcmobile.com"),
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

func TestMessaging10dlcBrandListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.List(context.TODO(), telnyx.Messaging10dlcBrandListParams{
		BrandID:        telnyx.String("826ef77a-348c-445b-81a5-a9b13c68fbfe"),
		Country:        telnyx.String("country"),
		DisplayName:    telnyx.String("displayName"),
		EntityType:     telnyx.String("entityType"),
		Page:           telnyx.Int(1),
		RecordsPerPage: telnyx.Int(0),
		Sort:           telnyx.Messaging10dlcBrandListParamsSortAssignedCampaignsCount,
		State:          telnyx.String("state"),
		TcrBrandID:     telnyx.String("BBAND1"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandDelete(t *testing.T) {
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
	err := client.Messaging10dlc.Brand.Delete(context.TODO(), "brandId")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandGetFeedback(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.GetFeedback(context.TODO(), "brandId")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandResend2faEmail(t *testing.T) {
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
	err := client.Messaging10dlc.Brand.Resend2faEmail(context.TODO(), "brandId")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessaging10dlcBrandRevet(t *testing.T) {
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
	_, err := client.Messaging10dlc.Brand.Revet(context.TODO(), "brandId")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
