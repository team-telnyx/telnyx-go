// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// MesssageService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMesssageService] method instead.
type MesssageService struct {
	Options []option.RequestOption
}

// NewMesssageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMesssageService(opts ...option.RequestOption) (r MesssageService) {
	r = MesssageService{}
	r.Options = opts
	return
}

// Send an RCS message
func (r *MesssageService) Rcs(ctx context.Context, body MesssageRcsParams, opts ...option.RequestOption) (res *MesssageRcsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messsages/rcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a Whatsapp message
func (r *MesssageService) Whatsapp(ctx context.Context, body MesssageWhatsappParams, opts ...option.RequestOption) (res *MesssageWhatsappResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messsages/whatsapp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RcsAgentMessage struct {
	ContentMessage RcsAgentMessageContentMessage `json:"content_message"`
	// RCS Event to send to the recipient
	Event RcsAgentMessageEvent `json:"event"`
	// Timestamp in UTC of when this message is considered expired
	ExpireTime time.Time `json:"expire_time" format:"date-time"`
	// Duration in seconds ending with 's'
	Ttl string `json:"ttl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentMessage respjson.Field
		Event          respjson.Field
		ExpireTime     respjson.Field
		Ttl            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessage) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsAgentMessage to a RcsAgentMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsAgentMessageParam.Overrides()
func (r RcsAgentMessage) ToParam() RcsAgentMessageParam {
	return param.Override[RcsAgentMessageParam](json.RawMessage(r.RawJSON()))
}

type RcsAgentMessageContentMessage struct {
	ContentInfo RcsContentInfo                        `json:"content_info"`
	RichCard    RcsAgentMessageContentMessageRichCard `json:"rich_card"`
	// List of suggested actions and replies
	Suggestions []RcsSuggestion `json:"suggestions"`
	// Text (maximum 3072 characters)
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentInfo respjson.Field
		RichCard    respjson.Field
		Suggestions respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessage) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageRichCard struct {
	// Carousel of cards.
	CarouselCard RcsAgentMessageContentMessageRichCardCarouselCard `json:"carousel_card"`
	// Standalone card
	StandaloneCard RcsAgentMessageContentMessageRichCardStandaloneCard `json:"standalone_card"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarouselCard   respjson.Field
		StandaloneCard respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carousel of cards.
type RcsAgentMessageContentMessageRichCardCarouselCard struct {
	// The list of contents for each card in the carousel. A carousel can have a
	// minimum of 2 cards and a maximum 10 cards.
	CardContents []RcsCardContent `json:"card_contents,required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardContents respjson.Field
		CardWidth    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCardCarouselCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCardCarouselCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standalone card
type RcsAgentMessageContentMessageRichCardStandaloneCard struct {
	CardContent RcsCardContent `json:"card_content,required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation,required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardContent             respjson.Field
		CardOrientation         respjson.Field
		ThumbnailImageAlignment respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCardStandaloneCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCardStandaloneCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RCS Event to send to the recipient
type RcsAgentMessageEvent struct {
	// Any of "TYPE_UNSPECIFIED", "IS_TYPING", "READ".
	EventType string `json:"event_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageParam struct {
	// Timestamp in UTC of when this message is considered expired
	ExpireTime param.Opt[time.Time] `json:"expire_time,omitzero" format:"date-time"`
	// Duration in seconds ending with 's'
	Ttl            param.Opt[string]                  `json:"ttl,omitzero"`
	ContentMessage RcsAgentMessageContentMessageParam `json:"content_message,omitzero"`
	// RCS Event to send to the recipient
	Event RcsAgentMessageEventParam `json:"event,omitzero"`
	paramObj
}

func (r RcsAgentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageParam struct {
	// Text (maximum 3072 characters)
	Text        param.Opt[string]                          `json:"text,omitzero"`
	ContentInfo RcsContentInfoParam                        `json:"content_info,omitzero"`
	RichCard    RcsAgentMessageContentMessageRichCardParam `json:"rich_card,omitzero"`
	// List of suggested actions and replies
	Suggestions []RcsSuggestionParam `json:"suggestions,omitzero"`
	paramObj
}

func (r RcsAgentMessageContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageRichCardParam struct {
	// Carousel of cards.
	CarouselCard RcsAgentMessageContentMessageRichCardCarouselCardParam `json:"carousel_card,omitzero"`
	// Standalone card
	StandaloneCard RcsAgentMessageContentMessageRichCardStandaloneCardParam `json:"standalone_card,omitzero"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carousel of cards.
//
// The properties CardContents, CardWidth are required.
type RcsAgentMessageContentMessageRichCardCarouselCardParam struct {
	// The list of contents for each card in the carousel. A carousel can have a
	// minimum of 2 cards and a maximum 10 cards.
	CardContents []RcsCardContentParam `json:"card_contents,omitzero,required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width,omitzero,required"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardCarouselCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardCarouselCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardCarouselCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardCarouselCardParam](
		"card_width", "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM",
	)
}

// Standalone card
//
// The properties CardContent, CardOrientation, ThumbnailImageAlignment are
// required.
type RcsAgentMessageContentMessageRichCardStandaloneCardParam struct {
	CardContent RcsCardContentParam `json:"card_content,omitzero,required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation,omitzero,required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment,omitzero,required"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardStandaloneCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardStandaloneCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardStandaloneCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardStandaloneCardParam](
		"card_orientation", "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL",
	)
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardStandaloneCardParam](
		"thumbnail_image_alignment", "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT",
	)
}

// RCS Event to send to the recipient
type RcsAgentMessageEventParam struct {
	// Any of "TYPE_UNSPECIFIED", "IS_TYPING", "READ".
	EventType string `json:"event_type,omitzero"`
	paramObj
}

func (r RcsAgentMessageEventParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageEventParam](
		"event_type", "TYPE_UNSPECIFIED", "IS_TYPING", "READ",
	)
}

type RcsCardContent struct {
	// Description of the card (at most 2000 characters)
	Description string `json:"description"`
	// A media file within a rich card.
	Media RcsCardContentMedia `json:"media"`
	// List of suggestions to include in the card. Maximum 10 suggestions.
	Suggestions []RcsSuggestion `json:"suggestions"`
	// Title of the card (at most 200 characters)
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Media       respjson.Field
		Suggestions respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsCardContent) RawJSON() string { return r.JSON.raw }
func (r *RcsCardContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsCardContent to a RcsCardContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsCardContentParam.Overrides()
func (r RcsCardContent) ToParam() RcsCardContentParam {
	return param.Override[RcsCardContentParam](json.RawMessage(r.RawJSON()))
}

// A media file within a rich card.
type RcsCardContentMedia struct {
	ContentInfo RcsContentInfo `json:"content_info"`
	// The height of the media within a rich card with a vertical layout. For a
	// standalone card with horizontal layout, height is not customizable, and this
	// field is ignored.
	//
	// Any of "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL".
	Height string `json:"height"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentInfo respjson.Field
		Height      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsCardContentMedia) RawJSON() string { return r.JSON.raw }
func (r *RcsCardContentMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsCardContentParam struct {
	// Description of the card (at most 2000 characters)
	Description param.Opt[string] `json:"description,omitzero"`
	// Title of the card (at most 200 characters)
	Title param.Opt[string] `json:"title,omitzero"`
	// A media file within a rich card.
	Media RcsCardContentMediaParam `json:"media,omitzero"`
	// List of suggestions to include in the card. Maximum 10 suggestions.
	Suggestions []RcsSuggestionParam `json:"suggestions,omitzero"`
	paramObj
}

func (r RcsCardContentParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsCardContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsCardContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media file within a rich card.
type RcsCardContentMediaParam struct {
	ContentInfo RcsContentInfoParam `json:"content_info,omitzero"`
	// The height of the media within a rich card with a vertical layout. For a
	// standalone card with horizontal layout, height is not customizable, and this
	// field is ignored.
	//
	// Any of "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL".
	Height string `json:"height,omitzero"`
	paramObj
}

func (r RcsCardContentMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsCardContentMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsCardContentMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsCardContentMediaParam](
		"height", "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL",
	)
}

type RcsContentInfo struct {
	// Publicly reachable URL of the file.
	FileURL string `json:"file_url,required" format:"url"`
	// If set the URL content will not be cached.
	ForceRefresh bool `json:"force_refresh"`
	// Publicly reachable URL of the thumbnail. Maximum size of 100 kB.
	ThumbnailURL string `json:"thumbnail_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileURL      respjson.Field
		ForceRefresh respjson.Field
		ThumbnailURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsContentInfo) RawJSON() string { return r.JSON.raw }
func (r *RcsContentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsContentInfo to a RcsContentInfoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsContentInfoParam.Overrides()
func (r RcsContentInfo) ToParam() RcsContentInfoParam {
	return param.Override[RcsContentInfoParam](json.RawMessage(r.RawJSON()))
}

// The property FileURL is required.
type RcsContentInfoParam struct {
	// Publicly reachable URL of the file.
	FileURL string `json:"file_url,required" format:"url"`
	// If set the URL content will not be cached.
	ForceRefresh param.Opt[bool] `json:"force_refresh,omitzero"`
	// Publicly reachable URL of the thumbnail. Maximum size of 100 kB.
	ThumbnailURL param.Opt[string] `json:"thumbnail_url,omitzero" format:"url"`
	paramObj
}

func (r RcsContentInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsContentInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsContentInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestion struct {
	// When tapped, initiates the corresponding native action on the device.
	Action RcsSuggestionAction `json:"action"`
	Reply  RcsSuggestionReply  `json:"reply"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestion) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsSuggestion to a RcsSuggestionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsSuggestionParam.Overrides()
func (r RcsSuggestion) ToParam() RcsSuggestionParam {
	return param.Override[RcsSuggestionParam](json.RawMessage(r.RawJSON()))
}

// When tapped, initiates the corresponding native action on the device.
type RcsSuggestionAction struct {
	// Opens the user's default calendar app and starts the new calendar event flow
	// with the agent-specified event data pre-filled.
	CreateCalendarEventAction RcsSuggestionActionCreateCalendarEventAction `json:"create_calendar_event_action"`
	// Opens the user's default dialer app with the agent-specified phone number filled
	// in.
	DialAction RcsSuggestionActionDialAction `json:"dial_action"`
	// Fallback URL to use if a client doesn't support a suggested action. Fallback
	// URLs open in new browser windows. Maximum 2048 characters.
	FallbackURL string `json:"fallback_url" format:"url"`
	// Opens the user's default web browser app to the specified URL.
	OpenURLAction RcsSuggestionActionOpenURLAction `json:"open_url_action"`
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData string `json:"postback_data"`
	// Opens the RCS app's location chooser so the user can pick a location to send
	// back to the agent.
	ShareLocationAction map[string]any `json:"share_location_action"`
	// Text that is shown in the suggested action. Maximum 25 characters.
	Text string `json:"text"`
	// Opens the user's default map app and selects the agent-specified location.
	ViewLocationAction RcsSuggestionActionViewLocationAction `json:"view_location_action"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateCalendarEventAction respjson.Field
		DialAction                respjson.Field
		FallbackURL               respjson.Field
		OpenURLAction             respjson.Field
		PostbackData              respjson.Field
		ShareLocationAction       respjson.Field
		Text                      respjson.Field
		ViewLocationAction        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default calendar app and starts the new calendar event flow
// with the agent-specified event data pre-filled.
type RcsSuggestionActionCreateCalendarEventAction struct {
	// Event description. Maximum 500 characters.
	Description string    `json:"description"`
	EndTime     time.Time `json:"end_time" format:"date-time"`
	StartTime   time.Time `json:"start_time" format:"date-time"`
	// Event title. Maximum 100 characters.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		EndTime     respjson.Field
		StartTime   respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionCreateCalendarEventAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionCreateCalendarEventAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default dialer app with the agent-specified phone number filled
// in.
type RcsSuggestionActionDialAction struct {
	// Phone number in +E.164 format
	PhoneNumber string `json:"phone_number,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionDialAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionDialAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default web browser app to the specified URL.
type RcsSuggestionActionOpenURLAction struct {
	// URL open application, browser or webview.
	//
	// Any of "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW".
	Application string `json:"application,required"`
	URL         string `json:"url,required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode,required"`
	// Accessbility description for webview.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Application     respjson.Field
		URL             respjson.Field
		WebviewViewMode respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionOpenURLAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionOpenURLAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default map app and selects the agent-specified location.
type RcsSuggestionActionViewLocationAction struct {
	// The label of the pin dropped
	Label   string                                       `json:"label"`
	LatLong RcsSuggestionActionViewLocationActionLatLong `json:"lat_long"`
	// query string (Android only)
	Query string `json:"query"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		LatLong     respjson.Field
		Query       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionViewLocationAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionViewLocationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionActionViewLocationActionLatLong struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `json:"latitude,required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionViewLocationActionLatLong) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionViewLocationActionLatLong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionReply struct {
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData string `json:"postback_data"`
	// Text that is shown in the suggested reply (maximum 25 characters)
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostbackData respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionReply) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionParam struct {
	// When tapped, initiates the corresponding native action on the device.
	Action RcsSuggestionActionParam `json:"action,omitzero"`
	Reply  RcsSuggestionReplyParam  `json:"reply,omitzero"`
	paramObj
}

func (r RcsSuggestionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When tapped, initiates the corresponding native action on the device.
type RcsSuggestionActionParam struct {
	// Fallback URL to use if a client doesn't support a suggested action. Fallback
	// URLs open in new browser windows. Maximum 2048 characters.
	FallbackURL param.Opt[string] `json:"fallback_url,omitzero" format:"url"`
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData param.Opt[string] `json:"postback_data,omitzero"`
	// Text that is shown in the suggested action. Maximum 25 characters.
	Text param.Opt[string] `json:"text,omitzero"`
	// Opens the user's default calendar app and starts the new calendar event flow
	// with the agent-specified event data pre-filled.
	CreateCalendarEventAction RcsSuggestionActionCreateCalendarEventActionParam `json:"create_calendar_event_action,omitzero"`
	// Opens the user's default dialer app with the agent-specified phone number filled
	// in.
	DialAction RcsSuggestionActionDialActionParam `json:"dial_action,omitzero"`
	// Opens the user's default web browser app to the specified URL.
	OpenURLAction RcsSuggestionActionOpenURLActionParam `json:"open_url_action,omitzero"`
	// Opens the RCS app's location chooser so the user can pick a location to send
	// back to the agent.
	ShareLocationAction map[string]any `json:"share_location_action,omitzero"`
	// Opens the user's default map app and selects the agent-specified location.
	ViewLocationAction RcsSuggestionActionViewLocationActionParam `json:"view_location_action,omitzero"`
	paramObj
}

func (r RcsSuggestionActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default calendar app and starts the new calendar event flow
// with the agent-specified event data pre-filled.
type RcsSuggestionActionCreateCalendarEventActionParam struct {
	// Event description. Maximum 500 characters.
	Description param.Opt[string]    `json:"description,omitzero"`
	EndTime     param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	StartTime   param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	// Event title. Maximum 100 characters.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r RcsSuggestionActionCreateCalendarEventActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionCreateCalendarEventActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionCreateCalendarEventActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default dialer app with the agent-specified phone number filled
// in.
//
// The property PhoneNumber is required.
type RcsSuggestionActionDialActionParam struct {
	// Phone number in +E.164 format
	PhoneNumber string `json:"phone_number,required"`
	paramObj
}

func (r RcsSuggestionActionDialActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionDialActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionDialActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default web browser app to the specified URL.
//
// The properties Application, URL, WebviewViewMode are required.
type RcsSuggestionActionOpenURLActionParam struct {
	// URL open application, browser or webview.
	//
	// Any of "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW".
	Application string `json:"application,omitzero,required"`
	URL         string `json:"url,required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode,omitzero,required"`
	// Accessbility description for webview.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r RcsSuggestionActionOpenURLActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionOpenURLActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionOpenURLActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsSuggestionActionOpenURLActionParam](
		"application", "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW",
	)
	apijson.RegisterFieldValidator[RcsSuggestionActionOpenURLActionParam](
		"webview_view_mode", "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL",
	)
}

// Opens the user's default map app and selects the agent-specified location.
type RcsSuggestionActionViewLocationActionParam struct {
	// The label of the pin dropped
	Label param.Opt[string] `json:"label,omitzero"`
	// query string (Android only)
	Query   param.Opt[string]                                 `json:"query,omitzero"`
	LatLong RcsSuggestionActionViewLocationActionLatLongParam `json:"lat_long,omitzero"`
	paramObj
}

func (r RcsSuggestionActionViewLocationActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionViewLocationActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionViewLocationActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type RcsSuggestionActionViewLocationActionLatLongParam struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `json:"latitude,required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude,required"`
	paramObj
}

func (r RcsSuggestionActionViewLocationActionLatLongParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionViewLocationActionLatLongParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionViewLocationActionLatLongParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionReplyParam struct {
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData param.Opt[string] `json:"postback_data,omitzero"`
	// Text that is shown in the suggested reply (maximum 25 characters)
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r RcsSuggestionReplyParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionReplyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionReplyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsResponse struct {
	Data MesssageRcsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageRcsResponse) RawJSON() string { return r.JSON.raw }
func (r *MesssageRcsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsResponseData struct {
	// message ID
	ID                 string                      `json:"id"`
	Body               RcsAgentMessage             `json:"body"`
	Direction          string                      `json:"direction"`
	Encoding           string                      `json:"encoding"`
	From               MesssageRcsResponseDataFrom `json:"from"`
	MessagingProfileID string                      `json:"messaging_profile_id"`
	OrganizationID     string                      `json:"organization_id"`
	ReceivedAt         time.Time                   `json:"received_at" format:"date-time"`
	RecordType         string                      `json:"record_type"`
	To                 []MesssageRcsResponseDataTo `json:"to"`
	Type               string                      `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Body               respjson.Field
		Direction          respjson.Field
		Encoding           respjson.Field
		From               respjson.Field
		MessagingProfileID respjson.Field
		OrganizationID     respjson.Field
		ReceivedAt         respjson.Field
		RecordType         respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageRcsResponseData) RawJSON() string { return r.JSON.raw }
func (r *MesssageRcsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsResponseDataFrom struct {
	// agent ID
	AgentID   string `json:"agent_id"`
	AgentName string `json:"agent_name"`
	Carrier   string `json:"carrier"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		AgentName   respjson.Field
		Carrier     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageRcsResponseDataFrom) RawJSON() string { return r.JSON.raw }
func (r *MesssageRcsResponseDataFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsResponseDataTo struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageRcsResponseDataTo) RawJSON() string { return r.JSON.raw }
func (r *MesssageRcsResponseDataTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponse struct {
	Data MesssageWhatsappResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponse) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseData struct {
	// message ID
	ID                 string                           `json:"id"`
	Body               MesssageWhatsappResponseDataBody `json:"body"`
	Direction          string                           `json:"direction"`
	Encoding           string                           `json:"encoding"`
	From               MesssageWhatsappResponseDataFrom `json:"from"`
	MessagingProfileID string                           `json:"messaging_profile_id"`
	OrganizationID     string                           `json:"organization_id"`
	ReceivedAt         time.Time                        `json:"received_at" format:"date-time"`
	RecordType         string                           `json:"record_type"`
	To                 []MesssageWhatsappResponseDataTo `json:"to"`
	Type               string                           `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Body               respjson.Field
		Direction          respjson.Field
		Encoding           respjson.Field
		From               respjson.Field
		MessagingProfileID respjson.Field
		OrganizationID     respjson.Field
		ReceivedAt         respjson.Field
		RecordType         respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseData) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBody struct {
	Audio MesssageWhatsappResponseDataBodyAudio `json:"audio"`
	// custom data to return with status update
	BizOpaqueCallbackData string                                      `json:"biz_opaque_callback_data"`
	Contacts              []MesssageWhatsappResponseDataBodyContact   `json:"contacts"`
	Document              MesssageWhatsappResponseDataBodyDocument    `json:"document"`
	Image                 MesssageWhatsappResponseDataBodyImage       `json:"image"`
	Interactive           MesssageWhatsappResponseDataBodyInteractive `json:"interactive"`
	Location              MesssageWhatsappResponseDataBodyLocation    `json:"location"`
	Reaction              MesssageWhatsappResponseDataBodyReaction    `json:"reaction"`
	Sticker               MesssageWhatsappResponseDataBodySticker     `json:"sticker"`
	// Any of "audio", "document", "image", "sticker", "video", "interactive",
	// "location", "template", "reaction", "contacts".
	Type  string                                `json:"type"`
	Video MesssageWhatsappResponseDataBodyVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio                 respjson.Field
		BizOpaqueCallbackData respjson.Field
		Contacts              respjson.Field
		Document              respjson.Field
		Image                 respjson.Field
		Interactive           respjson.Field
		Location              respjson.Field
		Reaction              respjson.Field
		Sticker               respjson.Field
		Type                  respjson.Field
		Video                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBody) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyAudio struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyAudio) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContact struct {
	Addresses []MesssageWhatsappResponseDataBodyContactAddress `json:"addresses"`
	Birthday  string                                           `json:"birthday"`
	Emails    []MesssageWhatsappResponseDataBodyContactEmail   `json:"emails"`
	Name      string                                           `json:"name"`
	Org       MesssageWhatsappResponseDataBodyContactOrg       `json:"org"`
	Phones    []MesssageWhatsappResponseDataBodyContactPhone   `json:"phones"`
	URLs      []MesssageWhatsappResponseDataBodyContactURL     `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses   respjson.Field
		Birthday    respjson.Field
		Emails      respjson.Field
		Name        respjson.Field
		Org         respjson.Field
		Phones      respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContact) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContactAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"`
	Zip         string `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Type        respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContactAddress) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContactAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContactEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContactEmail) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContactEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContactOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Department  respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContactOrg) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContactOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContactPhone struct {
	Phone string `json:"phone" format:"E+164"`
	Type  string `json:"type"`
	WaID  string `json:"wa_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phone       respjson.Field
		Type        respjson.Field
		WaID        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContactPhone) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContactPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyContactURL struct {
	Type string `json:"type"`
	URL  string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyContactURL) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyContactURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyDocument struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyDocument) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyImage struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyImage) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractive struct {
	Action MesssageWhatsappResponseDataBodyInteractiveAction `json:"action"`
	Body   MesssageWhatsappResponseDataBodyInteractiveBody   `json:"body"`
	Footer MesssageWhatsappResponseDataBodyInteractiveFooter `json:"footer"`
	Header MesssageWhatsappResponseDataBodyInteractiveHeader `json:"header"`
	// Any of "cta_url", "list", "carousel", "button", "location_request_message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Body        respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractive) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveAction struct {
	Button            string                                                      `json:"button"`
	Buttons           []MesssageWhatsappResponseDataBodyInteractiveActionButton   `json:"buttons"`
	Cards             []MesssageWhatsappResponseDataBodyInteractiveActionCard     `json:"cards"`
	CatalogID         string                                                      `json:"catalog_id"`
	Mode              string                                                      `json:"mode"`
	Name              string                                                      `json:"name"`
	Parameters        MesssageWhatsappResponseDataBodyInteractiveActionParameters `json:"parameters"`
	ProductRetailerID string                                                      `json:"product_retailer_id"`
	Sections          []MesssageWhatsappResponseDataBodyInteractiveActionSection  `json:"sections"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Button            respjson.Field
		Buttons           respjson.Field
		Cards             respjson.Field
		CatalogID         respjson.Field
		Mode              respjson.Field
		Name              respjson.Field
		Parameters        respjson.Field
		ProductRetailerID respjson.Field
		Sections          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveAction) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionButton struct {
	Reply MesssageWhatsappResponseDataBodyInteractiveActionButtonReply `json:"reply"`
	// Any of "reply".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reply       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionButton) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionButtonReply struct {
	// unique identifier for each button, 256 character maximum
	ID string `json:"id"`
	// button label, 20 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionButtonReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionButtonReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCard struct {
	Action MesssageWhatsappResponseDataBodyInteractiveActionCardAction `json:"action"`
	Body   MesssageWhatsappResponseDataBodyInteractiveActionCardBody   `json:"body"`
	// unique index for each card (0-9)
	CardIndex int64                                                       `json:"card_index"`
	Header    MesssageWhatsappResponseDataBodyInteractiveActionCardHeader `json:"header"`
	// Any of "cta_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Body        respjson.Field
		CardIndex   respjson.Field
		Header      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCard) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCardAction struct {
	// the unique ID of the catalog
	CatalogID string `json:"catalog_id"`
	// the unique retailer ID of the product
	ProductRetailerID string `json:"product_retailer_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CatalogID         respjson.Field
		ProductRetailerID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCardAction) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCardAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCardBody struct {
	// 160 character maximum, up to 2 line breaks
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCardBody) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCardBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCardHeader struct {
	Image MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderImage `json:"image"`
	// Any of "image", "video".
	Type  string                                                           `json:"type"`
	Video MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCardHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCardHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderImage struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderImage) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderVideo struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderVideo) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionCardHeaderVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionParameters struct {
	// button label text, 20 character maximum
	DisplayText string `json:"display_text"`
	// button URL to load when tapped by the user
	URL string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayText respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionSection struct {
	ProductItems []MesssageWhatsappResponseDataBodyInteractiveActionSectionProductItem `json:"product_items"`
	Rows         []MesssageWhatsappResponseDataBodyInteractiveActionSectionRow         `json:"rows"`
	// section title, 24 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductItems respjson.Field
		Rows         respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionSection) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveActionSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionSectionProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductRetailerID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionSectionProductItem) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionSectionProductItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveActionSectionRow struct {
	// arbitrary string identifying the row, 200 character maximum
	ID string `json:"id"`
	// row description, 72 character maximum
	Description string `json:"description"`
	// row title, 24 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveActionSectionRow) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveActionSectionRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveBody struct {
	// body text, 1024 character maximum
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveBody) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveFooter struct {
	// footer text, 60 character maximum
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveFooter) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveHeader struct {
	Document MesssageWhatsappResponseDataBodyInteractiveHeaderDocument `json:"document"`
	Image    MesssageWhatsappResponseDataBodyInteractiveHeaderImage    `json:"image"`
	SubText  string                                                    `json:"sub_text"`
	// header text, 60 character maximum
	Text  string                                                 `json:"text"`
	Video MesssageWhatsappResponseDataBodyInteractiveHeaderVideo `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		Image       respjson.Field
		SubText     respjson.Field
		Text        respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveHeader) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveHeaderDocument struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveHeaderDocument) RawJSON() string {
	return r.JSON.raw
}
func (r *MesssageWhatsappResponseDataBodyInteractiveHeaderDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveHeaderImage struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveHeaderImage) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveHeaderImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyInteractiveHeaderVideo struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyInteractiveHeaderVideo) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyInteractiveHeaderVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyLocation struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyLocation) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyReaction struct {
	Empji     string `json:"empji"`
	MessageID string `json:"message_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Empji       respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyReaction) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodySticker struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodySticker) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodySticker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataBodyVideo struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataBodyVideo) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataBodyVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataFrom struct {
	// The carrier of the sender.
	Carrier string `json:"carrier"`
	// The line-type of the sender.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// Any of "received", "delivered".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataFrom) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappResponseDataTo struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MesssageWhatsappResponseDataTo) RawJSON() string { return r.JSON.raw }
func (r *MesssageWhatsappResponseDataTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsParams struct {
	// RCS Agent ID
	AgentID      string               `json:"agent_id,required"`
	AgentMessage RcsAgentMessageParam `json:"agent_message,omitzero,required"`
	// A valid messaging profile ID
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// Phone number in +E.164 format
	To string `json:"to,required"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL  param.Opt[string]            `json:"webhook_url,omitzero" format:"url"`
	MmsFallback MesssageRcsParamsMmsFallback `json:"mms_fallback,omitzero"`
	SMSFallback MesssageRcsParamsSMSFallback `json:"sms_fallback,omitzero"`
	// Message type - must be set to "RCS"
	//
	// Any of "RCS".
	Type MesssageRcsParamsType `json:"type,omitzero"`
	paramObj
}

func (r MesssageRcsParams) MarshalJSON() (data []byte, err error) {
	type shadow MesssageRcsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageRcsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsParamsMmsFallback struct {
	// Phone number in +E.164 format
	From param.Opt[string] `json:"from,omitzero"`
	// Subject of the message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Text
	Text param.Opt[string] `json:"text,omitzero"`
	// List of media URLs
	MediaURLs []string `json:"media_urls,omitzero"`
	paramObj
}

func (r MesssageRcsParamsMmsFallback) MarshalJSON() (data []byte, err error) {
	type shadow MesssageRcsParamsMmsFallback
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageRcsParamsMmsFallback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageRcsParamsSMSFallback struct {
	// Phone number in +E.164 format
	From param.Opt[string] `json:"from,omitzero"`
	// Text (maximum 3072 characters)
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MesssageRcsParamsSMSFallback) MarshalJSON() (data []byte, err error) {
	type shadow MesssageRcsParamsSMSFallback
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageRcsParamsSMSFallback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message type - must be set to "RCS"
type MesssageRcsParamsType string

const (
	MesssageRcsParamsTypeRcs MesssageRcsParamsType = "RCS"
)

type MesssageWhatsappParams struct {
	// Phone number in +E.164 format associated with Whatsapp account
	From string `json:"from,required"`
	// Phone number in +E.164 format
	To              string                                `json:"to,required"`
	WhatsappMessage MesssageWhatsappParamsWhatsappMessage `json:"whatsapp_message,omitzero,required"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// Message type - must be set to "WHATSAPP"
	//
	// Any of "WHATSAPP".
	Type MesssageWhatsappParamsType `json:"type,omitzero"`
	paramObj
}

func (r MesssageWhatsappParams) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessage struct {
	// custom data to return with status update
	BizOpaqueCallbackData param.Opt[string]                                `json:"biz_opaque_callback_data,omitzero"`
	Audio                 MesssageWhatsappParamsWhatsappMessageAudio       `json:"audio,omitzero"`
	Contacts              []MesssageWhatsappParamsWhatsappMessageContact   `json:"contacts,omitzero"`
	Document              MesssageWhatsappParamsWhatsappMessageDocument    `json:"document,omitzero"`
	Image                 MesssageWhatsappParamsWhatsappMessageImage       `json:"image,omitzero"`
	Interactive           MesssageWhatsappParamsWhatsappMessageInteractive `json:"interactive,omitzero"`
	Location              MesssageWhatsappParamsWhatsappMessageLocation    `json:"location,omitzero"`
	Reaction              MesssageWhatsappParamsWhatsappMessageReaction    `json:"reaction,omitzero"`
	Sticker               MesssageWhatsappParamsWhatsappMessageSticker     `json:"sticker,omitzero"`
	// Any of "audio", "document", "image", "sticker", "video", "interactive",
	// "location", "template", "reaction", "contacts".
	Type  string                                     `json:"type,omitzero"`
	Video MesssageWhatsappParamsWhatsappMessageVideo `json:"video,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessage) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MesssageWhatsappParamsWhatsappMessage](
		"type", "audio", "document", "image", "sticker", "video", "interactive", "location", "template", "reaction", "contacts",
	)
}

type MesssageWhatsappParamsWhatsappMessageAudio struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageAudio) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageAudio
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageAudio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContact struct {
	Birthday  param.Opt[string]                                     `json:"birthday,omitzero"`
	Name      param.Opt[string]                                     `json:"name,omitzero"`
	Addresses []MesssageWhatsappParamsWhatsappMessageContactAddress `json:"addresses,omitzero"`
	Emails    []MesssageWhatsappParamsWhatsappMessageContactEmail   `json:"emails,omitzero"`
	Org       MesssageWhatsappParamsWhatsappMessageContactOrg       `json:"org,omitzero"`
	Phones    []MesssageWhatsappParamsWhatsappMessageContactPhone   `json:"phones,omitzero"`
	URLs      []MesssageWhatsappParamsWhatsappMessageContactURL     `json:"urls,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContact) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContactAddress struct {
	City        param.Opt[string] `json:"city,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	CountryCode param.Opt[string] `json:"country_code,omitzero"`
	State       param.Opt[string] `json:"state,omitzero"`
	Street      param.Opt[string] `json:"street,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	Zip         param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContactAddress) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContactAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContactAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContactEmail struct {
	Email param.Opt[string] `json:"email,omitzero"`
	Type  param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContactEmail) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContactEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContactEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContactOrg struct {
	Company    param.Opt[string] `json:"company,omitzero"`
	Department param.Opt[string] `json:"department,omitzero"`
	Title      param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContactOrg) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContactOrg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContactOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContactPhone struct {
	Phone param.Opt[string] `json:"phone,omitzero" format:"E+164"`
	Type  param.Opt[string] `json:"type,omitzero"`
	WaID  param.Opt[string] `json:"wa_id,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContactPhone) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContactPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContactPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageContactURL struct {
	Type param.Opt[string] `json:"type,omitzero"`
	URL  param.Opt[string] `json:"url,omitzero" format:"url"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageContactURL) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageContactURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageContactURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageDocument struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageDocument) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageImage struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageImage) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractive struct {
	Action MesssageWhatsappParamsWhatsappMessageInteractiveAction `json:"action,omitzero"`
	Body   MesssageWhatsappParamsWhatsappMessageInteractiveBody   `json:"body,omitzero"`
	Footer MesssageWhatsappParamsWhatsappMessageInteractiveFooter `json:"footer,omitzero"`
	Header MesssageWhatsappParamsWhatsappMessageInteractiveHeader `json:"header,omitzero"`
	// Any of "cta_url", "list", "carousel", "button", "location_request_message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractive) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractive
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MesssageWhatsappParamsWhatsappMessageInteractive](
		"type", "cta_url", "list", "carousel", "button", "location_request_message",
	)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveAction struct {
	Button            param.Opt[string]                                                `json:"button,omitzero"`
	CatalogID         param.Opt[string]                                                `json:"catalog_id,omitzero"`
	Mode              param.Opt[string]                                                `json:"mode,omitzero"`
	Name              param.Opt[string]                                                `json:"name,omitzero"`
	ProductRetailerID param.Opt[string]                                                `json:"product_retailer_id,omitzero"`
	Buttons           []MesssageWhatsappParamsWhatsappMessageInteractiveActionButton   `json:"buttons,omitzero"`
	Cards             []MesssageWhatsappParamsWhatsappMessageInteractiveActionCard     `json:"cards,omitzero"`
	Parameters        MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters `json:"parameters,omitzero"`
	Sections          []MesssageWhatsappParamsWhatsappMessageInteractiveActionSection  `json:"sections,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveAction) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionButton struct {
	Reply MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply `json:"reply,omitzero"`
	// Any of "reply".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionButton) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MesssageWhatsappParamsWhatsappMessageInteractiveActionButton](
		"type", "reply",
	)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply struct {
	// unique identifier for each button, 256 character maximum
	ID param.Opt[string] `json:"id,omitzero"`
	// button label, 20 character maximum
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionButtonReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCard struct {
	// unique index for each card (0-9)
	CardIndex param.Opt[int64]                                                 `json:"card_index,omitzero"`
	Action    MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction `json:"action,omitzero"`
	Body      MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody   `json:"body,omitzero"`
	Header    MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader `json:"header,omitzero"`
	// Any of "cta_url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCard) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MesssageWhatsappParamsWhatsappMessageInteractiveActionCard](
		"type", "cta_url",
	)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction struct {
	// the unique ID of the catalog
	CatalogID param.Opt[string] `json:"catalog_id,omitzero"`
	// the unique retailer ID of the product
	ProductRetailerID param.Opt[string] `json:"product_retailer_id,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCardAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody struct {
	// 160 character maximum, up to 2 line breaks
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCardBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader struct {
	Image MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage `json:"image,omitzero"`
	// Any of "image", "video".
	Type  string                                                                `json:"type,omitzero"`
	Video MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo `json:"video,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeader](
		"type", "image", "video",
	)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionCardHeaderVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters struct {
	// button label text, 20 character maximum
	DisplayText param.Opt[string] `json:"display_text,omitzero"`
	// button URL to load when tapped by the user
	URL param.Opt[string] `json:"url,omitzero" format:"url"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionSection struct {
	// section title, 24 character maximum
	Title        param.Opt[string]                                                          `json:"title,omitzero"`
	ProductItems []MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem `json:"product_items,omitzero"`
	Rows         []MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow         `json:"rows,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionSection) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem struct {
	ProductRetailerID param.Opt[string] `json:"product_retailer_id,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow struct {
	// arbitrary string identifying the row, 200 character maximum
	ID param.Opt[string] `json:"id,omitzero"`
	// row description, 72 character maximum
	Description param.Opt[string] `json:"description,omitzero"`
	// row title, 24 character maximum
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveActionSectionRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveBody struct {
	// body text, 1024 character maximum
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveBody) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveFooter struct {
	// footer text, 60 character maximum
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveFooter) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveHeader struct {
	SubText param.Opt[string] `json:"sub_text,omitzero"`
	// header text, 60 character maximum
	Text     param.Opt[string]                                              `json:"text,omitzero"`
	Document MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument `json:"document,omitzero"`
	Image    MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage    `json:"image,omitzero"`
	Video    MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo    `json:"video,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveHeader) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveHeaderDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveHeaderImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageInteractiveHeaderVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageLocation struct {
	Address   param.Opt[string] `json:"address,omitzero"`
	Latitude  param.Opt[string] `json:"latitude,omitzero"`
	Longitude param.Opt[string] `json:"longitude,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageLocation) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageReaction struct {
	Empji     param.Opt[string] `json:"empji,omitzero"`
	MessageID param.Opt[string] `json:"message_id,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageReaction) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageReaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageSticker struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageSticker) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageSticker
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageSticker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MesssageWhatsappParamsWhatsappMessageVideo struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r MesssageWhatsappParamsWhatsappMessageVideo) MarshalJSON() (data []byte, err error) {
	type shadow MesssageWhatsappParamsWhatsappMessageVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MesssageWhatsappParamsWhatsappMessageVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message type - must be set to "WHATSAPP"
type MesssageWhatsappParamsType string

const (
	MesssageWhatsappParamsTypeWhatsapp MesssageWhatsappParamsType = "WHATSAPP"
)
