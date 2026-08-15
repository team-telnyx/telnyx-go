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

func TestMeetingSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MeetingSessions.New(context.TODO(), telnyx.MeetingSessionNewParams{
		MeetingURL: "https://zoom.us/j/1234567890",
		Assistant: telnyx.MeetingSessionNewParamsAssistant{
			ID:                      "asst_fake-uuid-1234",
			CallControlConnectionID: "conn-fake-abcdef",
			From:                    "+12025550199",
			LoopbackSipUri:          "sip:loopback@example.invalid",
			AudioGate:               "half_duplex",
		},
		Avatar: telnyx.MeetingSessionNewParamsAvatar{
			APIKey:   "fake_avatar_api_key_do_not_use",
			AvatarID: "avatar_fake-001",
		},
		BargeIn: telnyx.Bool(true),
		BotName: telnyx.String("Notetaker"),
		CameraImage: telnyx.MeetingSessionNewParamsCameraImageUnion{
			OfMeetingSessionCameraImageBase64Source: &telnyx.MeetingSessionNewParamsCameraImageMeetingSessionCameraImageBase64Source{
				Base64Data: "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/wAALCAACAAIBAREA/8QAFAABAAAAAAAAAAAAAAAAAAAACP/EAB4QAAAEBwAAAAAAAAAAAAAAAAAEBgcCFic1RVNi/9oACAEBAAA/AH8hGJbWR09TxKW4vhC2qHgf/9k=",
			},
		},
		IdempotencyKey: telnyx.String("x"),
		JoinAt:         telnyx.Time(time.Now()),
		Metadata: map[string]any{
			"foo": "bar",
		},
		SpeakOnEnter:   telnyx.String("x"),
		SummarizeOnEnd: telnyx.Bool(true),
		Voice:          telnyx.String("x"),
		WebhookURL:     telnyx.String("https://example.com"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionGet(t *testing.T) {
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
	_, err := client.MeetingSessions.Get(context.TODO(), "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MeetingSessions.Update(
		context.TODO(),
		"mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		telnyx.MeetingSessionUpdateParams{
			BotName: telnyx.String("x"),
			JoinAt:  telnyx.Time(time.Now()),
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

func TestMeetingSessionListWithOptionalParams(t *testing.T) {
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
	_, err := client.MeetingSessions.List(context.TODO(), telnyx.MeetingSessionListParams{
		Status: telnyx.MeetingSessionListParamsStatusScheduled,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionDelete(t *testing.T) {
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
	_, err := client.MeetingSessions.Delete(context.TODO(), "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionDeleteRecordingMedia(t *testing.T) {
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
	_, err := client.MeetingSessions.DeleteRecordingMedia(context.TODO(), "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionGetEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.MeetingSessions.GetEvents(
		context.TODO(),
		"mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		telnyx.MeetingSessionGetEventsParams{
			After: telnyx.Int(0),
			Limit: telnyx.Int(1),
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

func TestMeetingSessionGetRecordings(t *testing.T) {
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
	_, err := client.MeetingSessions.GetRecordings(context.TODO(), "mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeetingSessionGetTranscriptWithOptionalParams(t *testing.T) {
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
	_, err := client.MeetingSessions.GetTranscript(
		context.TODO(),
		"mtgsess_a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		telnyx.MeetingSessionGetTranscriptParams{
			After:       telnyx.Int(0),
			Limit:       telnyx.Int(1),
			WaitSeconds: telnyx.Int(0),
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
