// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Connect calendars and manage automatic meeting joins.
//
// CalendarConnectionMeetingService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCalendarConnectionMeetingService] method instead.
type CalendarConnectionMeetingService struct {
	Options []option.RequestOption
}

// NewCalendarConnectionMeetingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCalendarConnectionMeetingService(opts ...option.RequestOption) (r CalendarConnectionMeetingService) {
	r = CalendarConnectionMeetingService{}
	r.Options = opts
	return
}

// Set an individual meeting override: true opts in, false opts out, null follows
// the connection policy. A supported meeting link is required. Changes re-evaluate
// the meeting immediately; poll the linked session for its final lifecycle
// outcome.
func (r *CalendarConnectionMeetingService) Update(ctx context.Context, meetingID string, params CalendarConnectionMeetingUpdateParams, opts ...option.RequestOption) (res *CalendarConnectionMeetingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if meetingID == "" {
		err = errors.New("missing required meeting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calendar_connections/%s/meetings/%s", url.PathEscape(params.ID), url.PathEscape(meetingID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List tracked meetings in a window. Defaults from now through seven days from
// now. from must not exceed to and the window must not exceed 31 days.
// bot_will_join is the effective policy decision; check session_id and the linked
// session status for actual attendance.
func (r *CalendarConnectionMeetingService) List(ctx context.Context, id string, query CalendarConnectionMeetingListParams, opts ...option.RequestOption) (res *CalendarConnectionMeetingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calendar_connections/%s/meetings", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CalendarMeeting struct {
	// Tracked calendar meeting ID.
	ID string `json:"id" api:"required"`
	// Calendar attendees; identity fields are cleared after the retention window.
	Attendees []CalendarMeetingAttendee `json:"attendees" api:"required"`
	// Effective policy decision, not proof that the bot has joined.
	BotWillJoin bool `json:"bot_will_join" api:"required"`
	// Explicit override; null follows the connection policy.
	BotWillJoinOverride bool `json:"bot_will_join_override" api:"required"`
	// Scheduling decision.
	//
	// Any of "pending", "scheduled", "skipped_policy", "skipped_no_url",
	// "skipped_duplicate", "skipped_unsupported", "cancelled".
	Decision CalendarMeetingDecision `json:"decision" api:"required"`
	// Event end time.
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// Whether the connected calendar user is not the organizer.
	IsExternalOrganizer bool                     `json:"is_external_organizer" api:"required"`
	Organizer           CalendarMeetingOrganizer `json:"organizer" api:"required"`
	// Detected meeting platform.
	//
	// Any of "zoom", "google_meet", "teams", "webex", "unknown".
	Platform CalendarMeetingPlatform `json:"platform" api:"required"`
	// Linked meeting session ID, or null until one exists.
	SessionID string `json:"session_id" api:"required"`
	// Event start time.
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Attendees           respjson.Field
		BotWillJoin         respjson.Field
		BotWillJoinOverride respjson.Field
		Decision            respjson.Field
		EndTime             respjson.Field
		IsExternalOrganizer respjson.Field
		Organizer           respjson.Field
		Platform            respjson.Field
		SessionID           respjson.Field
		StartTime           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarMeeting) RawJSON() string { return r.JSON.raw }
func (r *CalendarMeeting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarMeetingAttendee struct {
	// Display name.
	DisplayName string `json:"display_name" api:"required"`
	// Email address.
	Email string `json:"email" api:"required"`
	// Whether attendance is optional.
	Optional bool `json:"optional" api:"required"`
	// Calendar RSVP status.
	ResponseStatus string `json:"response_status" api:"required"`
	// Whether this is the connected calendar user.
	Self bool `json:"self" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName    respjson.Field
		Email          respjson.Field
		Optional       respjson.Field
		ResponseStatus respjson.Field
		Self           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarMeetingAttendee) RawJSON() string { return r.JSON.raw }
func (r *CalendarMeetingAttendee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scheduling decision.
type CalendarMeetingDecision string

const (
	CalendarMeetingDecisionPending            CalendarMeetingDecision = "pending"
	CalendarMeetingDecisionScheduled          CalendarMeetingDecision = "scheduled"
	CalendarMeetingDecisionSkippedPolicy      CalendarMeetingDecision = "skipped_policy"
	CalendarMeetingDecisionSkippedNoURL       CalendarMeetingDecision = "skipped_no_url"
	CalendarMeetingDecisionSkippedDuplicate   CalendarMeetingDecision = "skipped_duplicate"
	CalendarMeetingDecisionSkippedUnsupported CalendarMeetingDecision = "skipped_unsupported"
	CalendarMeetingDecisionCancelled          CalendarMeetingDecision = "cancelled"
)

type CalendarMeetingOrganizer struct {
	// Display name.
	DisplayName string `json:"display_name" api:"required"`
	// Email address.
	Email string `json:"email" api:"required"`
	// Whether this is the connected calendar user.
	Self bool `json:"self" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Email       respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarMeetingOrganizer) RawJSON() string { return r.JSON.raw }
func (r *CalendarMeetingOrganizer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detected meeting platform.
type CalendarMeetingPlatform string

const (
	CalendarMeetingPlatformZoom       CalendarMeetingPlatform = "zoom"
	CalendarMeetingPlatformGoogleMeet CalendarMeetingPlatform = "google_meet"
	CalendarMeetingPlatformTeams      CalendarMeetingPlatform = "teams"
	CalendarMeetingPlatformWebex      CalendarMeetingPlatform = "webex"
	CalendarMeetingPlatformUnknown    CalendarMeetingPlatform = "unknown"
)

type CalendarConnectionMeetingUpdateResponse struct {
	Data CalendarMeeting `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionMeetingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionMeetingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionMeetingListResponse struct {
	Data []CalendarMeeting `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarConnectionMeetingListResponse) RawJSON() string { return r.JSON.raw }
func (r *CalendarConnectionMeetingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionMeetingUpdateParams struct {
	// True opts in, false opts out, null follows connection policy. A supported
	// meeting URL is still required.
	BotWillJoin param.Opt[bool] `json:"bot_will_join,omitzero" api:"required"`
	ID          string          `path:"id" api:"required" json:"-"`
	paramObj
}

func (r CalendarConnectionMeetingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CalendarConnectionMeetingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CalendarConnectionMeetingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarConnectionMeetingListParams struct {
	// Window start; defaults to now.
	From param.Opt[time.Time] `query:"from,omitzero" format:"date-time" json:"-"`
	// Window end; defaults to seven days from now.
	To param.Opt[time.Time] `query:"to,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [CalendarConnectionMeetingListParams]'s query parameters as
// `url.Values`.
func (r CalendarConnectionMeetingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
