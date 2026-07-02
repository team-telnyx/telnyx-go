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

func TestMessaging10dlcCampaignBuilderSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging10dlc.CampaignBuilder.Submit(context.TODO(), telnyx.Messaging10dlcCampaignBuilderSubmitParams{
		BrandID:                "brandId",
		Description:            "description",
		Usecase:                "usecase",
		AgeGated:               telnyx.Bool(true),
		AutoRenewal:            telnyx.Bool(true),
		DirectLending:          telnyx.Bool(true),
		EmbeddedLink:           telnyx.Bool(true),
		EmbeddedLinkSample:     telnyx.String("embeddedLinkSample"),
		EmbeddedPhone:          telnyx.Bool(true),
		HelpKeywords:           telnyx.String("helpKeywords"),
		HelpMessage:            telnyx.String("helpMessage"),
		MessageFlow:            telnyx.String("messageFlow"),
		MnoIDs:                 []int64{0},
		NumberPool:             telnyx.Bool(true),
		OptinKeywords:          telnyx.String("optinKeywords"),
		OptinMessage:           telnyx.String("optinMessage"),
		OptoutKeywords:         telnyx.String("optoutKeywords"),
		OptoutMessage:          telnyx.String("optoutMessage"),
		PrivacyPolicyLink:      telnyx.String("privacyPolicyLink"),
		ReferenceID:            telnyx.String("referenceId"),
		ResellerID:             telnyx.String("resellerId"),
		Sample1:                telnyx.String("sample1"),
		Sample2:                telnyx.String("sample2"),
		Sample3:                telnyx.String("sample3"),
		Sample4:                telnyx.String("sample4"),
		Sample5:                telnyx.String("sample5"),
		SubscriberHelp:         telnyx.Bool(true),
		SubscriberOptin:        telnyx.Bool(true),
		SubscriberOptout:       telnyx.Bool(true),
		SubUsecases:            []string{"string"},
		Tag:                    []string{"string"},
		TermsAndConditions:     telnyx.Bool(true),
		TermsAndConditionsLink: telnyx.String("termsAndConditionsLink"),
		WebhookFailoverURL:     telnyx.String("https://webhook.com/93711262-23e5-4048-a966-c0b2a16d5963"),
		WebhookURL:             telnyx.String("https://webhook.com/67ea78a8-9f32-4d04-b62d-f9502e8e5f93"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
