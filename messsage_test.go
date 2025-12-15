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

func TestMesssageRcsWithOptionalParams(t *testing.T) {
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
	_, err := client.Messsages.Rcs(context.TODO(), telnyx.MesssageRcsParams{
		AgentID: "Agent007",
		AgentMessage: telnyx.RcsAgentMessageParam{
			ContentMessage: telnyx.RcsAgentMessageContentMessageParam{
				ContentInfo: telnyx.RcsContentInfoParam{
					FileURL:      "https://example.com/elephant.jpg",
					ForceRefresh: telnyx.Bool(true),
					ThumbnailURL: telnyx.String("thumbnail_url"),
				},
				RichCard: telnyx.RcsAgentMessageContentMessageRichCardParam{
					CarouselCard: telnyx.RcsAgentMessageContentMessageRichCardCarouselCardParam{
						CardContents: []telnyx.RcsCardContentParam{{
							Description: telnyx.String("description"),
							Media: telnyx.RcsCardContentMediaParam{
								ContentInfo: telnyx.RcsContentInfoParam{
									FileURL:      "https://example.com/elephant.jpg",
									ForceRefresh: telnyx.Bool(true),
									ThumbnailURL: telnyx.String("thumbnail_url"),
								},
								Height: "MEDIUM",
							},
							Suggestions: []telnyx.RcsSuggestionParam{{
								Action: telnyx.RcsSuggestionActionParam{
									CreateCalendarEventAction: telnyx.RcsSuggestionActionCreateCalendarEventActionParam{
										Description: telnyx.String("description"),
										EndTime:     telnyx.Time(time.Now()),
										StartTime:   telnyx.Time(time.Now()),
										Title:       telnyx.String("title"),
									},
									DialAction: telnyx.RcsSuggestionActionDialActionParam{
										PhoneNumber: "+13125551234",
									},
									FallbackURL: telnyx.String("fallback_url"),
									OpenURLAction: telnyx.RcsSuggestionActionOpenURLActionParam{
										Application:     "BROWSER",
										URL:             "http://example.com",
										WebviewViewMode: "HALF",
										Description:     telnyx.String("description"),
									},
									PostbackData: telnyx.String("postback_data"),
									ShareLocationAction: map[string]any{
										"foo": "bar",
									},
									Text: telnyx.String("Hello world"),
									ViewLocationAction: telnyx.RcsSuggestionActionViewLocationActionParam{
										Label: telnyx.String("label"),
										LatLong: telnyx.RcsSuggestionActionViewLocationActionLatLongParam{
											Latitude:  41.8,
											Longitude: -87.6,
										},
										Query: telnyx.String("query"),
									},
								},
								Reply: telnyx.RcsSuggestionReplyParam{
									PostbackData: telnyx.String("postback_data"),
									Text:         telnyx.String("text"),
								},
							}},
							Title: telnyx.String("Elephant"),
						}},
						CardWidth: "SMALL",
					},
					StandaloneCard: telnyx.RcsAgentMessageContentMessageRichCardStandaloneCardParam{
						CardContent: telnyx.RcsCardContentParam{
							Description: telnyx.String("description"),
							Media: telnyx.RcsCardContentMediaParam{
								ContentInfo: telnyx.RcsContentInfoParam{
									FileURL:      "https://example.com/elephant.jpg",
									ForceRefresh: telnyx.Bool(true),
									ThumbnailURL: telnyx.String("thumbnail_url"),
								},
								Height: "MEDIUM",
							},
							Suggestions: []telnyx.RcsSuggestionParam{{
								Action: telnyx.RcsSuggestionActionParam{
									CreateCalendarEventAction: telnyx.RcsSuggestionActionCreateCalendarEventActionParam{
										Description: telnyx.String("description"),
										EndTime:     telnyx.Time(time.Now()),
										StartTime:   telnyx.Time(time.Now()),
										Title:       telnyx.String("title"),
									},
									DialAction: telnyx.RcsSuggestionActionDialActionParam{
										PhoneNumber: "+13125551234",
									},
									FallbackURL: telnyx.String("fallback_url"),
									OpenURLAction: telnyx.RcsSuggestionActionOpenURLActionParam{
										Application:     "BROWSER",
										URL:             "http://example.com",
										WebviewViewMode: "HALF",
										Description:     telnyx.String("description"),
									},
									PostbackData: telnyx.String("postback_data"),
									ShareLocationAction: map[string]any{
										"foo": "bar",
									},
									Text: telnyx.String("Hello world"),
									ViewLocationAction: telnyx.RcsSuggestionActionViewLocationActionParam{
										Label: telnyx.String("label"),
										LatLong: telnyx.RcsSuggestionActionViewLocationActionLatLongParam{
											Latitude:  41.8,
											Longitude: -87.6,
										},
										Query: telnyx.String("query"),
									},
								},
								Reply: telnyx.RcsSuggestionReplyParam{
									PostbackData: telnyx.String("postback_data"),
									Text:         telnyx.String("text"),
								},
							}},
							Title: telnyx.String("Elephant"),
						},
						CardOrientation:         "HORIZONTAL",
						ThumbnailImageAlignment: "LEFT",
					},
				},
				Suggestions: []telnyx.RcsSuggestionParam{{
					Action: telnyx.RcsSuggestionActionParam{
						CreateCalendarEventAction: telnyx.RcsSuggestionActionCreateCalendarEventActionParam{
							Description: telnyx.String("description"),
							EndTime:     telnyx.Time(time.Now()),
							StartTime:   telnyx.Time(time.Now()),
							Title:       telnyx.String("title"),
						},
						DialAction: telnyx.RcsSuggestionActionDialActionParam{
							PhoneNumber: "+13125551234",
						},
						FallbackURL: telnyx.String("fallback_url"),
						OpenURLAction: telnyx.RcsSuggestionActionOpenURLActionParam{
							Application:     "BROWSER",
							URL:             "http://example.com",
							WebviewViewMode: "HALF",
							Description:     telnyx.String("description"),
						},
						PostbackData: telnyx.String("postback_data"),
						ShareLocationAction: map[string]any{
							"foo": "bar",
						},
						Text: telnyx.String("Hello world"),
						ViewLocationAction: telnyx.RcsSuggestionActionViewLocationActionParam{
							Label: telnyx.String("label"),
							LatLong: telnyx.RcsSuggestionActionViewLocationActionLatLongParam{
								Latitude:  41.8,
								Longitude: -87.6,
							},
							Query: telnyx.String("query"),
						},
					},
					Reply: telnyx.RcsSuggestionReplyParam{
						PostbackData: telnyx.String("postback_data"),
						Text:         telnyx.String("text"),
					},
				}},
				Text: telnyx.String("Hello world!"),
			},
			Event: telnyx.RcsAgentMessageEventParam{
				EventType: "IS_TYPING",
			},
			ExpireTime: telnyx.Time(time.Now()),
			Ttl:        telnyx.String("10.5s"),
		},
		MessagingProfileID: "messaging_profile_id",
		To:                 "+13125551234",
		MmsFallback: telnyx.MesssageRcsParamsMmsFallback{
			From:      telnyx.String("+13125551234"),
			MediaURLs: []string{"string"},
			Subject:   telnyx.String("Test Message"),
			Text:      telnyx.String("Hello world!"),
		},
		SMSFallback: telnyx.MesssageRcsParamsSMSFallback{
			From: telnyx.String("+13125551234"),
			Text: telnyx.String("Hello world!"),
		},
		Type:       telnyx.MesssageRcsParamsTypeRcs,
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

func TestMesssageWhatsappWithOptionalParams(t *testing.T) {
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
	_, err := client.Messsages.Whatsapp(context.TODO(), telnyx.MesssageWhatsappParams{
		From: "+13125551234",
		To:   "+13125551234",
		WhatsappMessage: telnyx.MesssageWhatsappParamsWhatsappMessage{
			Audio: telnyx.MesssageWhatsappParamsWhatsappMessageAudio{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			BizOpaqueCallbackData: telnyx.String("biz_opaque_callback_data"),
			Contacts: []telnyx.MesssageWhatsappParamsWhatsappMessageContact{{
				Addresses: []telnyx.MesssageWhatsappParamsWhatsappMessageContactAddress{{
					City:        telnyx.String("city"),
					Country:     telnyx.String("country"),
					CountryCode: telnyx.String("country_code"),
					State:       telnyx.String("state"),
					Street:      telnyx.String("street"),
					Type:        telnyx.String("type"),
					Zip:         telnyx.String("zip"),
				}},
				Birthday: telnyx.String("birthday"),
				Emails: []telnyx.MesssageWhatsappParamsWhatsappMessageContactEmail{{
					Email: telnyx.String("email"),
					Type:  telnyx.String("type"),
				}},
				Name: telnyx.String("name"),
				Org: telnyx.MesssageWhatsappParamsWhatsappMessageContactOrg{
					Company:    telnyx.String("company"),
					Department: telnyx.String("department"),
					Title:      telnyx.String("title"),
				},
				Phones: []telnyx.MesssageWhatsappParamsWhatsappMessageContactPhone{{
					Phone: telnyx.String("phone"),
					Type:  telnyx.String("type"),
					WaID:  telnyx.String("wa_id"),
				}},
				URLs: []telnyx.MesssageWhatsappParamsWhatsappMessageContactURL{{
					Type: telnyx.String("type"),
					URL:  telnyx.String("url"),
				}},
			}},
			Document: telnyx.MesssageWhatsappParamsWhatsappMessageDocument{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Image: telnyx.MesssageWhatsappParamsWhatsappMessageImage{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Interactive: telnyx.MesssageWhatsappParamsWhatsappMessageInteractive{
				Action: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveAction{
					Button: telnyx.String("button"),
					Buttons: []telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionButton{{
						Reply: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply{
							ID:    telnyx.String("id"),
							Title: telnyx.String("title"),
						},
						Type: "reply",
					}},
					Cards: []telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCard{{
						Action: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction{
							CatalogID:         telnyx.String("catalog_id"),
							ProductRetailerID: telnyx.String("product_retailer_id"),
						},
						Body: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody{
							Text: telnyx.String("text"),
						},
						CardIndex: telnyx.Int(0),
						Header: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader{
							Image: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage{
								Caption:  telnyx.String("caption"),
								Filename: telnyx.String("filename"),
								Link:     telnyx.String("http://example.com/media.jpg"),
								Voice:    telnyx.Bool(true),
							},
							Type: "image",
							Video: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo{
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
					Parameters: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters{
						DisplayText: telnyx.String("display_text"),
						URL:         telnyx.String("url"),
					},
					ProductRetailerID: telnyx.String("product_retailer_id"),
					Sections: []telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionSection{{
						ProductItems: []telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem{{
							ProductRetailerID: telnyx.String("product_retailer_id"),
						}},
						Rows: []telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow{{
							ID:          telnyx.String("id"),
							Description: telnyx.String("description"),
							Title:       telnyx.String("title"),
						}},
						Title: telnyx.String("title"),
					}},
				},
				Body: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveBody{
					Text: telnyx.String("text"),
				},
				Footer: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveFooter{
					Text: telnyx.String("text"),
				},
				Header: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveHeader{
					Document: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
					Image: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
					SubText: telnyx.String("sub_text"),
					Text:    telnyx.String("text"),
					Video: telnyx.MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo{
						Caption:  telnyx.String("caption"),
						Filename: telnyx.String("filename"),
						Link:     telnyx.String("http://example.com/media.jpg"),
						Voice:    telnyx.Bool(true),
					},
				},
				Type: "cta_url",
			},
			Location: telnyx.MesssageWhatsappParamsWhatsappMessageLocation{
				Address:   telnyx.String("address"),
				Latitude:  telnyx.String("latitude"),
				Longitude: telnyx.String("longitude"),
				Name:      telnyx.String("name"),
			},
			Reaction: telnyx.MesssageWhatsappParamsWhatsappMessageReaction{
				Empji:     telnyx.String("empji"),
				MessageID: telnyx.String("message_id"),
			},
			Sticker: telnyx.MesssageWhatsappParamsWhatsappMessageSticker{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
			Type: "audio",
			Video: telnyx.MesssageWhatsappParamsWhatsappMessageVideo{
				Caption:  telnyx.String("caption"),
				Filename: telnyx.String("filename"),
				Link:     telnyx.String("http://example.com/media.jpg"),
				Voice:    telnyx.Bool(true),
			},
		},
		Type:       telnyx.MesssageWhatsappParamsTypeWhatsapp,
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
