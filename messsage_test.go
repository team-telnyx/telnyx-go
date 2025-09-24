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
									PostbackData:        telnyx.String("postback_data"),
									ShareLocationAction: map[string]interface{}{},
									Text:                telnyx.String("Hello world"),
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
									PostbackData:        telnyx.String("postback_data"),
									ShareLocationAction: map[string]interface{}{},
									Text:                telnyx.String("Hello world"),
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
						PostbackData:        telnyx.String("postback_data"),
						ShareLocationAction: map[string]interface{}{},
						Text:                telnyx.String("Hello world"),
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
