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

func TestUserProfileList(t *testing.T) {
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
	_, err := client.UserProfile.List(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserProfilePatchAllWithOptionalParams(t *testing.T) {
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
	_, err := client.UserProfile.PatchAll(context.TODO(), telnyx.UserProfilePatchAllParams{
		Address1:                                telnyx.String("123 Example Street"),
		Address2:                                telnyx.String("Suite 100"),
		BusinessName:                            telnyx.String("Example Company"),
		City:                                    telnyx.String("Chicago"),
		CompanySize:                             telnyx.String("11-50"),
		Country:                                 telnyx.String("US"),
		DataLocality:                            telnyx.UserProfilePatchAllParamsDataLocalityDeu,
		DefaultManagedAccountAllowCustomPricing: telnyx.Bool(false),
		DefaultManagedAccountInheritEmergencyTermsOfService: telnyx.Bool(false),
		DefaultManagedAccountRollupBilling:                  telnyx.Bool(false),
		DefaultNotifyManagerAccount:                         telnyx.Bool(true),
		DefaultNumbersExternalPin:                           telnyx.String("1234"),
		FirstName:                                           telnyx.String("Alex"),
		GotFreeNumberAtSignup:                               telnyx.Bool(false),
		LastName:                                            telnyx.String("Taylor"),
		LowBalanceNotificationFrequency:                     12,
		LowBalanceThresholdAmount:                           telnyx.Float(10),
		NotifyOnInvoiceAvailable:                            telnyx.Bool(true),
		NotifyOnInvoicePayment:                              telnyx.Bool(false),
		NpacUser:                                            telnyx.Bool(false),
		OnboardingBuildMethod:                               []string{"API"},
		OnboardingIndustries:                                []string{"software"},
		OnboardingMonthlyBudget:                             []string{"1000-5000"},
		OnboardingProductInterest:                           []string{"voice"},
		OnboardingProjectTypes:                              []string{"contact center"},
		OnboardingReasonsForConsideration:                   []string{"API integration"},
		OnboardingReferrer:                                  []string{"search engine"},
		OnboardingUseCasesOther:                             []string{"customer support"},
		PhoneNumber:                                         telnyx.String("+13125550100"),
		PortOutPinValidation:                                telnyx.Bool(false),
		PrimaryUseCase:                                      telnyx.String("configure sip trunks"),
		Role:                                                telnyx.String("account manager"),
		State:                                               telnyx.String("IL"),
		UiTheme:                                             1,
		UseCases:                                            []string{"voice"},
		UserSubscriptionOptIn:                               telnyx.Bool(false),
		Zip:                                                 telnyx.String("60601"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
