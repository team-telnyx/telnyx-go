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
	"time"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestPortingOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.New(context.TODO(), telnyx.PortingOrderNewParams{
		PhoneNumbers:      []string{"+13035550000", "+13035550001", "+13035550002"},
		CustomerReference: telnyx.String("Acct 123abc"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingOrderGetWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderGetParams{
			IncludePhoneNumbers: telnyx.Bool(true),
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

func TestPortingOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderUpdateParams{
			ActivationSettings: telnyx.PortingOrderUpdateParamsActivationSettings{
				FocDatetimeRequested: telnyx.Time(time.Now()),
			},
			CustomerReference: telnyx.String("customer_reference"),
			Documents: telnyx.PortingOrderDocumentsParam{
				Invoice: telnyx.String("ce74b771-d23d-4960-81ec-8741b3862146"),
				Loa:     telnyx.String("64ffb720-04c7-455b-92d6-20fcca92e935"),
			},
			EndUser: telnyx.PortingOrderEndUserParam{
				Admin: telnyx.PortingOrderEndUserAdminParam{
					AccountNumber:      telnyx.String("123abc"),
					AuthPersonName:     telnyx.String("Porter McPortersen II"),
					BillingPhoneNumber: telnyx.String("billing_phone_number"),
					BusinessIdentifier: telnyx.String("abc123"),
					EntityName:         telnyx.String("Porter McPortersen"),
					PinPasscode:        telnyx.String("pin_passcode"),
					TaxIdentifier:      telnyx.String("1234abcd"),
				},
				Location: telnyx.PortingOrderEndUserLocationParam{
					AdministrativeArea: telnyx.String("TX"),
					CountryCode:        telnyx.String("US"),
					ExtendedAddress:    telnyx.String("14th Floor"),
					Locality:           telnyx.String("Austin"),
					PostalCode:         telnyx.String("78701"),
					StreetAddress:      telnyx.String("600 Congress Avenue"),
				},
			},
			Messaging: telnyx.PortingOrderUpdateParamsMessaging{
				EnableMessaging: telnyx.Bool(true),
			},
			Misc: telnyx.PortingOrderMiscParam{
				NewBillingPhoneNumber:  telnyx.String("new_billing_phone_number"),
				RemainingNumbersAction: telnyx.PortingOrderMiscRemainingNumbersActionDisconnect,
				Type:                   telnyx.PortingOrderTypeFull,
			},
			PhoneNumberConfiguration: telnyx.PortingOrderPhoneNumberConfigurationParam{
				BillingGroupID:     telnyx.String("f1486bae-f067-460c-ad43-73a92848f902"),
				ConnectionID:       telnyx.String("f1486bae-f067-460c-ad43-73a92848f902"),
				EmergencyAddressID: telnyx.String("f1486bae-f067-460c-ad43-73a92848f902"),
				MessagingProfileID: telnyx.String("f1486bae-f067-460c-ad43-73a92848f901"),
				Tags:               []string{"abc", "123"},
			},
			RequirementGroupID: telnyx.String("DE748D99-06FA-4D90-9F9A-F4B62696BADA"),
			Requirements: []telnyx.PortingOrderUpdateParamsRequirement{{
				FieldValue:        "9787fb5f-cbe5-4de4-b765-3303774ee9fe",
				RequirementTypeID: "59b0762a-b274-4f76-ac32-4d5cf0272e66",
			}},
			UserFeedback: telnyx.PortingOrderUserFeedbackParam{
				UserComment: telnyx.String("I loved my experience porting numbers with Telnyx"),
				UserRating:  telnyx.Int(5),
			},
			WebhookURL: telnyx.String("https://example.com"),
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

func TestPortingOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.List(context.TODO(), telnyx.PortingOrderListParams{
		Filter: telnyx.PortingOrderListParamsFilter{
			ActivationSettingsFastPortEligible: telnyx.Bool(true),
			ActivationSettingsFocDatetimeRequested: telnyx.PortingOrderListParamsFilterActivationSettingsFocDatetimeRequested{
				Gt: telnyx.String("2021-03-25T10:00:00.000Z"),
				Lt: telnyx.String("2021-03-25T10:00:00.000Z"),
			},
			CustomerReference:          telnyx.String("customer_reference"),
			EndUserAdminAuthPersonName: telnyx.String("end_user.admin.auth_person_name"),
			EndUserAdminEntityName:     telnyx.String("end_user.admin.entity_name"),
			MiscType:                   telnyx.PortingOrderTypeFull,
			ParentSupportKey:           telnyx.String("parent_support_key"),
			PhoneNumbersCarrierName:    telnyx.String("phone_numbers.carrier_name"),
			PhoneNumbersCountryCode:    telnyx.String("phone_numbers.country_code"),
			PhoneNumbersPhoneNumber: telnyx.PortingOrderListParamsFilterPhoneNumbersPhoneNumber{
				Contains: telnyx.String("contains"),
			},
		},
		IncludePhoneNumbers: telnyx.Bool(true),
		Page: telnyx.PortingOrderListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.PortingOrderListParamsSort{
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

func TestPortingOrderDelete(t *testing.T) {
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
	err := client.PortingOrders.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingOrderGetAllowedFocWindows(t *testing.T) {
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
	_, err := client.PortingOrders.GetAllowedFocWindows(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingOrderGetExceptionTypes(t *testing.T) {
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
	_, err := client.PortingOrders.GetExceptionTypes(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortingOrderGetLoaTemplateWithOptionalParams(t *testing.T) {
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
	resp, err := client.PortingOrders.GetLoaTemplate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderGetLoaTemplateParams{
			LoaConfigurationID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
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

func TestPortingOrderGetRequirementsWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.GetRequirements(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderGetRequirementsParams{
			Page: telnyx.PortingOrderGetRequirementsParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
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

func TestPortingOrderGetSubRequest(t *testing.T) {
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
	_, err := client.PortingOrders.GetSubRequest(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
