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

func TestDirReferenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.References.New(
		context.TODO(),
		"16635d38-75a6-4481-82e8-69af60e05011",
		telnyx.DirReferenceNewParams{
			BusinessReferences: []telnyx.ReferenceInputParam{{
				Email:                    "dana.reyes@example.com",
				FullName:                 "Dana Reyes",
				PhoneE164:                "+14155550123",
				Timezone:                 "America/New_York",
				JobTitle:                 telnyx.String("VP of Operations"),
				Organization:             telnyx.String("Acme Logistics"),
				RelationshipToRegistrant: telnyx.String("Supplier"),
			}, {
				Email:                    "dana.reyes@example.com",
				FullName:                 "Dana Reyes",
				PhoneE164:                "+14155550123",
				Timezone:                 "America/New_York",
				JobTitle:                 telnyx.String("VP of Operations"),
				Organization:             telnyx.String("Acme Logistics"),
				RelationshipToRegistrant: telnyx.String("Supplier"),
			}},
			FinancialReference: telnyx.ReferenceInputParam{
				Email:                    "dana.reyes@example.com",
				FullName:                 "Dana Reyes",
				PhoneE164:                "+14155550123",
				Timezone:                 "America/New_York",
				JobTitle:                 telnyx.String("VP of Operations"),
				Organization:             telnyx.String("Acme Logistics"),
				RelationshipToRegistrant: telnyx.String("Supplier"),
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

func TestDirReferenceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Dir.References.Update(
		context.TODO(),
		0,
		telnyx.DirReferenceUpdateParams{
			DirID:                    "16635d38-75a6-4481-82e8-69af60e05011",
			RefType:                  telnyx.DirReferenceUpdateParamsRefTypeBusiness,
			Email:                    telnyx.String("dana.reyes@example.com"),
			FullName:                 telnyx.String("Dana Reyes"),
			JobTitle:                 telnyx.String("VP of Operations"),
			Organization:             telnyx.String("Acme Logistics"),
			PhoneE164:                telnyx.String("+14155550123"),
			RelationshipToRegistrant: telnyx.String("Supplier"),
			Timezone:                 telnyx.String("America/New_York"),
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

func TestDirReferenceList(t *testing.T) {
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
	_, err := client.Dir.References.List(context.TODO(), "16635d38-75a6-4481-82e8-69af60e05011")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
