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

func TestMessageGet(t *testing.T) {
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
	_, err := client.Messages.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageCancelScheduled(t *testing.T) {
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
	_, err := client.Messages.CancelScheduled(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageGetGroupMessages(t *testing.T) {
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
	_, err := client.Messages.GetGroupMessages(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageScheduleWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Schedule(context.TODO(), telnyx.MessageScheduleParams{
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		From:               telnyx.String("+18445550001"),
		MediaURLs:          []string{"string"},
		MessagingProfileID: telnyx.String("abc85f64-5717-4562-b3fc-2c9600000000"),
		SendAt:             telnyx.Time(time.Now()),
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageScheduleParamsTypeSMS,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.Send(context.TODO(), telnyx.MessageSendParams{
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		Encoding:           telnyx.MessageSendParamsEncodingAuto,
		From:               telnyx.String("+18445550001"),
		MediaURLs:          []string{"http://example.com"},
		MessagingProfileID: telnyx.String("abc85f64-5717-4562-b3fc-2c9600000000"),
		SendAt:             telnyx.Time(time.Now()),
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendGroupMmsWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendGroupMms(context.TODO(), telnyx.MessageSendGroupMmsParams{
		From:               "+13125551234",
		To:                 []string{"+18655551234", "+14155551234"},
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendLongCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendLongCode(context.TODO(), telnyx.MessageSendLongCodeParams{
		From:               "+18445550001",
		To:                 "+13125550002",
		AutoDetect:         telnyx.Bool(true),
		Encoding:           telnyx.MessageSendLongCodeParamsEncodingAuto,
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendLongCodeParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendNumberPoolWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendNumberPool(context.TODO(), telnyx.MessageSendNumberPoolParams{
		MessagingProfileID: "abc85f64-5717-4562-b3fc-2c9600000000",
		To:                 "+13125550002",
		AutoDetect:         telnyx.Bool(true),
		Encoding:           telnyx.MessageSendNumberPoolParamsEncodingAuto,
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendNumberPoolParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendShortCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendShortCode(context.TODO(), telnyx.MessageSendShortCodeParams{
		From:               "+18445550001",
		To:                 "+18445550001",
		AutoDetect:         telnyx.Bool(true),
		Encoding:           telnyx.MessageSendShortCodeParamsEncodingAuto,
		MediaURLs:          []string{"http://example.com"},
		Subject:            telnyx.String("From Telnyx!"),
		Text:               telnyx.String("Hello, World!"),
		Type:               telnyx.MessageSendShortCodeParamsTypeMms,
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("https://backup.example.com/hooks"),
		WebhookURL:         telnyx.String("http://example.com/webhooks"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendWhatsappWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendWhatsapp(context.TODO(), telnyx.MessageSendWhatsappParams{
		From: "+13125551234",
		To:   "+13125551234",
		WhatsappMessage: telnyx.WhatsappMessageContentParam{
			Audio: telnyx.WhatsappMediaParam{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			BizOpaqueCallbackData: telnyx.String("biz_opaque_callback_data"),
			Contacts: []telnyx.WhatsappContactParam{{
				Addresses: []telnyx.WhatsappContactAddressParam{{
					City:        telnyx.String("city"),
					Country:     telnyx.String("country"),
					CountryCode: telnyx.String("country_code"),
					State:       telnyx.String("state"),
					Street:      telnyx.String("street"),
					Type:        telnyx.String("type"),
					Zip:         telnyx.String("zip"),
				}},
				Birthday: telnyx.String("birthday"),
				Emails: []telnyx.WhatsappContactEmailParam{{
					Email: telnyx.String("email"),
					Type:  telnyx.String("type"),
				}},
				Name: telnyx.String("name"),
				Org: telnyx.WhatsappContactOrgParam{
					Company:    telnyx.String("company"),
					Department: telnyx.String("department"),
					Title:      telnyx.String("title"),
				},
				Phones: []telnyx.WhatsappContactPhoneParam{{
					Phone: telnyx.String("phone"),
					Type:  telnyx.String("type"),
					WaID:  telnyx.String("wa_id"),
				}},
				URLs: []telnyx.WhatsappContactURLParam{{
					Type: telnyx.String("type"),
					URL:  telnyx.String("url"),
				}},
			}},
			Document: telnyx.WhatsappMediaParam{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Image: telnyx.WhatsappMediaParam{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Interactive: telnyx.WhatsappInteractiveParam{
				Action: telnyx.WhatsappInteractiveActionParam{
					Button: telnyx.String("button"),
					Buttons: []telnyx.WhatsappInteractiveActionButtonParam{{
						Reply: telnyx.WhatsappInteractiveActionButtonReplyParam{
							ID:    telnyx.String("id"),
							Title: telnyx.String("title"),
						},
						Type: "reply",
					}},
					Cards: []telnyx.WhatsappInteractiveActionCardParam{{
						Action: telnyx.WhatsappInteractiveActionCardActionParam{
							CatalogID:         telnyx.String("catalog_id"),
							ProductRetailerID: telnyx.String("product_retailer_id"),
						},
						Body: telnyx.WhatsappInteractiveActionCardBodyParam{
							Text: telnyx.String("text"),
						},
						CardIndex: telnyx.Int(0),
						Header: telnyx.WhatsappInteractiveActionCardHeaderParam{
							Image: telnyx.WhatsappMediaParam{
								Caption:  telnyx.String("caption"),
								Filename: telnyx.String("filename"),
								Link:     telnyx.String("http://example.com/media.jpg"),
								Voice:    telnyx.Bool(true),
							},
							Type: "image",
							Video: telnyx.WhatsappMediaParam{
								Caption:  telnyx.String("caption"),
								Filename: telnyx.String("filename"),
								Link:     telnyx.String("http://example.com/media.jpg"),
								Voice:    telnyx.Bool(true),
							},
						},
						Type: "cta_url",
					}},
					CatalogID: telnyx.String("catalog_id"),
					Mode:      telnyx.String("mode"),
					Name:      telnyx.String("name"),
					Parameters: telnyx.WhatsappInteractiveActionParametersParam{
						DisplayText: telnyx.String("display_text"),
						URL:         telnyx.String("url"),
					},
					ProductRetailerID: telnyx.String("product_retailer_id"),
					Sections: []telnyx.WhatsappInteractiveActionSectionParam{{
						ProductItems: []telnyx.WhatsappInteractiveActionSectionProductItemParam{{
							ProductRetailerID: telnyx.String("product_retailer_id"),
						}},
						Rows: []telnyx.WhatsappInteractiveActionSectionRowParam{{
							ID:          telnyx.String("id"),
							Description: telnyx.String("description"),
							Title:       telnyx.String("title"),
						}},
						Title: telnyx.String("title"),
					}},
				},
				Body: telnyx.WhatsappInteractiveBodyParam{
					Text: telnyx.String("text"),
				},
				Footer: telnyx.WhatsappInteractiveFooterParam{
					Text: telnyx.String("text"),
				},
				Header: telnyx.WhatsappInteractiveHeaderParam{
					Document: telnyx.WhatsappMediaParam{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
					Image: telnyx.WhatsappMediaParam{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
					SubText: telnyx.String("sub_text"),
					Text:    telnyx.String("text"),
					Video: telnyx.WhatsappMediaParam{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
				},
				Type: telnyx.WhatsappInteractiveTypeCtaURL,
			},
			Location: telnyx.WhatsappLocationParam{
				Address:   telnyx.String("address"),
				Latitude:  telnyx.String("latitude"),
				Longitude: telnyx.String("longitude"),
				Name:      telnyx.String("name"),
			},
			Reaction: telnyx.WhatsappReactionParam{
				Emoji:     telnyx.String("emoji"),
				MessageID: telnyx.String("message_id"),
			},
			Sticker: telnyx.WhatsappMediaParam{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Type: telnyx.WhatsappMessageContentTypeAudio,
			Video: telnyx.WhatsappMediaParam{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
		},
		Type:       telnyx.MessageSendWhatsappParamsTypeWhatsapp,
		WebhookURL: telnyx.String("webhook_url"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageSendWithAlphanumericSenderWithOptionalParams(t *testing.T) {
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
	_, err := client.Messages.SendWithAlphanumericSender(context.TODO(), telnyx.MessageSendWithAlphanumericSenderParams{
		From:               "MyCompany",
		MessagingProfileID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Text:               "text",
		To:                 "+E.164",
		UseProfileWebhooks: telnyx.Bool(true),
		WebhookFailoverURL: telnyx.String("webhook_failover_url"),
		WebhookURL:         telnyx.String("webhook_url"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
