// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
)

// Connect calendars and manage automatic meeting joins.
//
// MeetingBotCalendarService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingBotCalendarService] method instead.
type MeetingBotCalendarService struct {
	Options []option.RequestOption
}

// NewMeetingBotCalendarService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMeetingBotCalendarService(opts ...option.RequestOption) (r MeetingBotCalendarService) {
	r = MeetingBotCalendarService{}
	r.Options = opts
	return
}

// Browser redirect after calendar consent. Use the returned oauth_url rather than
// constructing this request. Verified state identifies the connecting user.
// Returns HTML for success, cancellation, or failure; no JSON error envelope or
// bearer token is used.
func (r *MeetingBotCalendarService) GetOAuthCallback(ctx context.Context, query MeetingBotCalendarGetOAuthCallbackParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "meeting_bot/calendar/oauth_callback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MeetingBotCalendarGetOAuthCallbackParams struct {
	// Authorization code supplied by the calendar provider.
	Code param.Opt[string] `query:"code,omitzero" json:"-"`
	// Authorization error supplied by the calendar provider, such as access_denied.
	Error param.Opt[string] `query:"error,omitzero" json:"-"`
	// Unmodified signed state from the authorization URL.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MeetingBotCalendarGetOAuthCallbackParams]'s query
// parameters as `url.Values`.
func (r MeetingBotCalendarGetOAuthCallbackParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
