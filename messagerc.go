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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// MessageRcService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageRcService] method instead.
type MessageRcService struct {
	Options []option.RequestOption
}

// NewMessageRcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageRcService(opts ...option.RequestOption) (r MessageRcService) {
	r = MessageRcService{}
	r.Options = opts
	return
}

// Generate a deeplink URL that can be used to start an RCS conversation with a
// specific agent.
func (r *MessageRcService) GenerateDeeplink(ctx context.Context, agentID string, query MessageRcGenerateDeeplinkParams, opts ...option.RequestOption) (res *MessageRcGenerateDeeplinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("messages/rcs/deeplinks/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Send an RCS message
func (r *MessageRcService) Send(ctx context.Context, body MessageRcSendParams, opts ...option.RequestOption) (res *MessageRcSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/rcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessageRcGenerateDeeplinkResponse struct {
	Data MessageRcGenerateDeeplinkResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRcGenerateDeeplinkResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageRcGenerateDeeplinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcGenerateDeeplinkResponseData struct {
	// The generated deeplink URL
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRcGenerateDeeplinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageRcGenerateDeeplinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendResponse struct {
	Data MessageRcSendResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageRcSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageRcSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendResponseData struct {
	// message ID
	ID                 string                        `json:"id"`
	Body               RcsAgentMessage               `json:"body"`
	Direction          string                        `json:"direction"`
	Encoding           string                        `json:"encoding"`
	From               MessageRcSendResponseDataFrom `json:"from"`
	MessagingProfileID string                        `json:"messaging_profile_id"`
	OrganizationID     string                        `json:"organization_id"`
	ReceivedAt         time.Time                     `json:"received_at" format:"date-time"`
	RecordType         string                        `json:"record_type"`
	To                 []MessageRcSendResponseDataTo `json:"to"`
	Type               string                        `json:"type"`
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
func (r MessageRcSendResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageRcSendResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendResponseDataFrom struct {
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
func (r MessageRcSendResponseDataFrom) RawJSON() string { return r.JSON.raw }
func (r *MessageRcSendResponseDataFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendResponseDataTo struct {
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
func (r MessageRcSendResponseDataTo) RawJSON() string { return r.JSON.raw }
func (r *MessageRcSendResponseDataTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcGenerateDeeplinkParams struct {
	// Pre-filled message body (URL encoded)
	Body param.Opt[string] `query:"body,omitzero" json:"-"`
	// Phone number in E164 format (URL encoded)
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageRcGenerateDeeplinkParams]'s query parameters as
// `url.Values`.
func (r MessageRcGenerateDeeplinkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageRcSendParams struct {
	// RCS Agent ID
	AgentID      string               `json:"agent_id,required"`
	AgentMessage RcsAgentMessageParam `json:"agent_message,omitzero,required"`
	// A valid messaging profile ID
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// Phone number in +E.164 format
	To string `json:"to,required"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL  param.Opt[string]              `json:"webhook_url,omitzero" format:"url"`
	MmsFallback MessageRcSendParamsMmsFallback `json:"mms_fallback,omitzero"`
	SMSFallback MessageRcSendParamsSMSFallback `json:"sms_fallback,omitzero"`
	// Message type - must be set to "RCS"
	//
	// Any of "RCS".
	Type MessageRcSendParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageRcSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageRcSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageRcSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendParamsMmsFallback struct {
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

func (r MessageRcSendParamsMmsFallback) MarshalJSON() (data []byte, err error) {
	type shadow MessageRcSendParamsMmsFallback
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageRcSendParamsMmsFallback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageRcSendParamsSMSFallback struct {
	// Phone number in +E.164 format
	From param.Opt[string] `json:"from,omitzero"`
	// Text (maximum 3072 characters)
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MessageRcSendParamsSMSFallback) MarshalJSON() (data []byte, err error) {
	type shadow MessageRcSendParamsSMSFallback
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageRcSendParamsSMSFallback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message type - must be set to "RCS"
type MessageRcSendParamsType string

const (
	MessageRcSendParamsTypeRcs MessageRcSendParamsType = "RCS"
)
