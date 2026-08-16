// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Send real-time speech and chat actions to an active meeting session.
//
// MeetingSessionActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingSessionActionService] method instead.
type MeetingSessionActionService struct {
	Options []option.RequestOption
}

// NewMeetingSessionActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMeetingSessionActionService(opts ...option.RequestOption) (r MeetingSessionActionService) {
	r = MeetingSessionActionService{}
	r.Options = opts
	return
}

// Sends a chat message into a meeting session.
func (r *MeetingSessionActionService) SendChat(ctx context.Context, id string, body MeetingSessionActionSendChatParams, opts ...option.RequestOption) (res *ActionAcceptedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/actions/send_chat", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Sends audio / text-to-speech into a meeting session.
func (r *MeetingSessionActionService) Speak(ctx context.Context, id string, body MeetingSessionActionSpeakParams, opts ...option.RequestOption) (res *ActionAcceptedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/actions/speak", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stops any active text-to-speech playback in a meeting session.
func (r *MeetingSessionActionService) StopSpeaking(ctx context.Context, id string, opts ...option.RequestOption) (res *ActionAcceptedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/actions/stop_speaking", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ActionAcceptedResponse struct {
	Data ActionAcceptedResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionAcceptedResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionAcceptedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionAcceptedResponseData struct {
	Accepted bool `json:"accepted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionAcceptedResponseData) RawJSON() string { return r.JSON.raw }
func (r *ActionAcceptedResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionActionSendChatParams struct {
	// Chat message text to send in the meeting.
	Text string `json:"text" api:"required"`
	paramObj
}

func (r MeetingSessionActionSendChatParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionActionSendChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionActionSendChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionActionSpeakParams struct {
	// Text for the bot to speak.
	Text string `json:"text" api:"required"`
	// If true, interrupt any currently playing audio to speak this text immediately.
	Interrupt param.Opt[bool] `json:"interrupt,omitzero"`
	// Voice identifier to use for this utterance. When supplied, it overrides the
	// session-default voice configured at creation; otherwise the speak action uses
	// that session default.
	Voice param.Opt[string] `json:"voice,omitzero"`
	paramObj
}

func (r MeetingSessionActionSpeakParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionActionSpeakParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionActionSpeakParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
