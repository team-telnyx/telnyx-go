// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// RcsAgentService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRcsAgentService] method instead.
type RcsAgentService struct {
	Options []option.RequestOption
}

// NewRcsAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRcsAgentService(opts ...option.RequestOption) (r RcsAgentService) {
	r = RcsAgentService{}
	r.Options = opts
	return
}

type RcsAgent struct {
	// RCS Agent ID
	AgentID string `json:"agent_id"`
	// Human readable agent name
	AgentName string `json:"agent_name"`
	// Date and time when the resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specifies whether the agent is enabled
	Enabled bool `json:"enabled"`
	// Messaging profile ID associated with the RCS Agent
	ProfileID string `json:"profile_id,nullable" format:"uuid"`
	// Date and time when the resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// User ID associated with the RCS Agent
	UserID string `json:"user_id"`
	// Failover URL to receive RCS events
	WebhookFailoverURL string `json:"webhook_failover_url,nullable" format:"url"`
	// URL to receive RCS events
	WebhookURL string `json:"webhook_url,nullable" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID            respjson.Field
		AgentName          respjson.Field
		CreatedAt          respjson.Field
		Enabled            respjson.Field
		ProfileID          respjson.Field
		UpdatedAt          respjson.Field
		UserID             respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgent) RawJSON() string { return r.JSON.raw }
func (r *RcsAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentResponse struct {
	Data RcsAgent `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentResponse) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
