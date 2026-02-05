// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type APIError struct {
	Code   string         `json:"code,required" format:"int64"`
	Title  string         `json:"title,required"`
	Detail string         `json:"detail"`
	Meta   map[string]any `json:"meta"`
	Source APIErrorSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Title       respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIError) RawJSON() string { return r.JSON.raw }
func (r *APIError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIErrorSource struct {
	// Indicates which query parameter caused the error.
	Parameter string `json:"parameter"`
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer" format:"json-pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameter   respjson.Field
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIErrorSource) RawJSON() string { return r.JSON.raw }
func (r *APIErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AvailablePhoneNumbersMetadata struct {
	BestEffortResults int64 `json:"best_effort_results"`
	TotalResults      int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BestEffortResults respjson.Field
		TotalResults      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailablePhoneNumbersMetadata) RawJSON() string { return r.JSON.raw }
func (r *AvailablePhoneNumbersMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
// of SIP Trunking calls. The feature is off unless enabled. You may define min and
// max values in msec for customized buffering behaviors. Larger values add latency
// but tolerate more jitter, while smaller values reduce latency but are more
// sensitive to jitter and reordering.
type ConnectionJitterBuffer struct {
	// Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off
	// unless enabled.
	EnableJitterBuffer bool `json:"enable_jitter_buffer"`
	// The maximum jitter buffer size in milliseconds. Must be between 40 and 400. Has
	// no effect if enable_jitter_buffer is not true.
	JitterbufferMsecMax int64 `json:"jitterbuffer_msec_max"`
	// The minimum jitter buffer size in milliseconds. Must be between 40 and 400. Has
	// no effect if enable_jitter_buffer is not true.
	JitterbufferMsecMin int64 `json:"jitterbuffer_msec_min"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableJitterBuffer  respjson.Field
		JitterbufferMsecMax respjson.Field
		JitterbufferMsecMin respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionJitterBuffer) RawJSON() string { return r.JSON.raw }
func (r *ConnectionJitterBuffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConnectionJitterBuffer to a ConnectionJitterBufferParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConnectionJitterBufferParam.Overrides()
func (r ConnectionJitterBuffer) ToParam() ConnectionJitterBufferParam {
	return param.Override[ConnectionJitterBufferParam](json.RawMessage(r.RawJSON()))
}

// Configuration options for Jitter Buffer. Enables Jitter Buffer for RTP streams
// of SIP Trunking calls. The feature is off unless enabled. You may define min and
// max values in msec for customized buffering behaviors. Larger values add latency
// but tolerate more jitter, while smaller values reduce latency but are more
// sensitive to jitter and reordering.
type ConnectionJitterBufferParam struct {
	// Enables Jitter Buffer for RTP streams of SIP Trunking calls. The feature is off
	// unless enabled.
	EnableJitterBuffer param.Opt[bool] `json:"enable_jitter_buffer,omitzero"`
	// The maximum jitter buffer size in milliseconds. Must be between 40 and 400. Has
	// no effect if enable_jitter_buffer is not true.
	JitterbufferMsecMax param.Opt[int64] `json:"jitterbuffer_msec_max,omitzero"`
	// The minimum jitter buffer size in milliseconds. Must be between 40 and 400. Has
	// no effect if enable_jitter_buffer is not true.
	JitterbufferMsecMin param.Opt[int64] `json:"jitterbuffer_msec_min,omitzero"`
	paramObj
}

func (r ConnectionJitterBufferParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectionJitterBufferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectionJitterBufferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for noise suppression. These settings are stored
// regardless of the noise_suppression value, but only take effect when
// noise_suppression is not 'disabled'. If you disable noise suppression and later
// re-enable it, the previously configured settings will be used.
type ConnectionNoiseSuppressionDetails struct {
	// The attenuation limit value for the selected engine. Default values vary by
	// engine: 0 for 'denoiser', 80 for 'deep_filter_net', 'deep_filter_net_large', and
	// all Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite',
	// 'krisp_viva_promodel', 'krisp_viva_ss').
	AttenuationLimit int64 `json:"attenuation_limit"`
	// The noise suppression engine to use. 'denoiser' is the default engine.
	// 'deep_filter_net' and 'deep_filter_net_large' are alternative engines with
	// different performance characteristics. Krisp engines ('krisp_viva_tel',
	// 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss') provide advanced
	// noise suppression capabilities.
	//
	// Any of "denoiser", "deep_filter_net", "deep_filter_net_large", "krisp_viva_tel",
	// "krisp_viva_tel_lite", "krisp_viva_promodel", "krisp_viva_ss".
	Engine ConnectionNoiseSuppressionDetailsEngine `json:"engine"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttenuationLimit respjson.Field
		Engine           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionNoiseSuppressionDetails) RawJSON() string { return r.JSON.raw }
func (r *ConnectionNoiseSuppressionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConnectionNoiseSuppressionDetails to a
// ConnectionNoiseSuppressionDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConnectionNoiseSuppressionDetailsParam.Overrides()
func (r ConnectionNoiseSuppressionDetails) ToParam() ConnectionNoiseSuppressionDetailsParam {
	return param.Override[ConnectionNoiseSuppressionDetailsParam](json.RawMessage(r.RawJSON()))
}

// The noise suppression engine to use. 'denoiser' is the default engine.
// 'deep_filter_net' and 'deep_filter_net_large' are alternative engines with
// different performance characteristics. Krisp engines ('krisp_viva_tel',
// 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss') provide advanced
// noise suppression capabilities.
type ConnectionNoiseSuppressionDetailsEngine string

const (
	ConnectionNoiseSuppressionDetailsEngineDenoiser           ConnectionNoiseSuppressionDetailsEngine = "denoiser"
	ConnectionNoiseSuppressionDetailsEngineDeepFilterNet      ConnectionNoiseSuppressionDetailsEngine = "deep_filter_net"
	ConnectionNoiseSuppressionDetailsEngineDeepFilterNetLarge ConnectionNoiseSuppressionDetailsEngine = "deep_filter_net_large"
	ConnectionNoiseSuppressionDetailsEngineKrispVivaTel       ConnectionNoiseSuppressionDetailsEngine = "krisp_viva_tel"
	ConnectionNoiseSuppressionDetailsEngineKrispVivaTelLite   ConnectionNoiseSuppressionDetailsEngine = "krisp_viva_tel_lite"
	ConnectionNoiseSuppressionDetailsEngineKrispVivaPromodel  ConnectionNoiseSuppressionDetailsEngine = "krisp_viva_promodel"
	ConnectionNoiseSuppressionDetailsEngineKrispVivaSS        ConnectionNoiseSuppressionDetailsEngine = "krisp_viva_ss"
)

// Configuration options for noise suppression. These settings are stored
// regardless of the noise_suppression value, but only take effect when
// noise_suppression is not 'disabled'. If you disable noise suppression and later
// re-enable it, the previously configured settings will be used.
type ConnectionNoiseSuppressionDetailsParam struct {
	// The attenuation limit value for the selected engine. Default values vary by
	// engine: 0 for 'denoiser', 80 for 'deep_filter_net', 'deep_filter_net_large', and
	// all Krisp engines ('krisp_viva_tel', 'krisp_viva_tel_lite',
	// 'krisp_viva_promodel', 'krisp_viva_ss').
	AttenuationLimit param.Opt[int64] `json:"attenuation_limit,omitzero"`
	// The noise suppression engine to use. 'denoiser' is the default engine.
	// 'deep_filter_net' and 'deep_filter_net_large' are alternative engines with
	// different performance characteristics. Krisp engines ('krisp_viva_tel',
	// 'krisp_viva_tel_lite', 'krisp_viva_promodel', 'krisp_viva_ss') provide advanced
	// noise suppression capabilities.
	//
	// Any of "denoiser", "deep_filter_net", "deep_filter_net_large", "krisp_viva_tel",
	// "krisp_viva_tel_lite", "krisp_viva_promodel", "krisp_viva_ss".
	Engine ConnectionNoiseSuppressionDetailsEngine `json:"engine,omitzero"`
	paramObj
}

func (r ConnectionNoiseSuppressionDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectionNoiseSuppressionDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectionNoiseSuppressionDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionsPaginationMeta struct {
	PageNumber   int64 `json:"page_number,required"`
	TotalPages   int64 `json:"total_pages,required"`
	PageSize     int64 `json:"page_size"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionsPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *ConnectionsPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocReqsRequirementType struct {
	// Identifies the associated document
	ID string `json:"id" format:"uuid"`
	// Specifies objective criteria for acceptance
	AcceptanceCriteria DocReqsRequirementTypeAcceptanceCriteria `json:"acceptance_criteria"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Describes the requirement type
	Description string `json:"description"`
	// Provides one or more examples of acceptable documents
	Example string `json:"example"`
	// A short descriptive name for this requirement_type
	Name string `json:"name"`
	// Identifies the type of the resource
	RecordType string `json:"record_type"`
	// Defines the type of this requirement type
	//
	// Any of "document", "address", "textual".
	Type DocReqsRequirementTypeType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was last updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AcceptanceCriteria respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		Example            respjson.Field
		Name               respjson.Field
		RecordType         respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocReqsRequirementType) RawJSON() string { return r.JSON.raw }
func (r *DocReqsRequirementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies objective criteria for acceptance
type DocReqsRequirementTypeAcceptanceCriteria struct {
	// Specifies the allowed characters as a string
	AcceptableCharacters string `json:"acceptable_characters"`
	// Specifies the list of strictly possible values for the requirement. Ignored when
	// empty
	AcceptableValues []string `json:"acceptable_values"`
	// Specifies geography-based acceptance criteria
	LocalityLimit string `json:"locality_limit"`
	// Maximum length allowed for the value
	MaxLength int64 `json:"max_length"`
	// Minimum length allowed for the value
	MinLength int64 `json:"min_length"`
	// Specifies time-based acceptance criteria
	TimeLimit string `json:"time_limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptableCharacters respjson.Field
		AcceptableValues     respjson.Field
		LocalityLimit        respjson.Field
		MaxLength            respjson.Field
		MinLength            respjson.Field
		TimeLimit            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocReqsRequirementTypeAcceptanceCriteria) RawJSON() string { return r.JSON.raw }
func (r *DocReqsRequirementTypeAcceptanceCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the type of this requirement type
type DocReqsRequirementTypeType string

const (
	DocReqsRequirementTypeTypeDocument DocReqsRequirementTypeType = "document"
	DocReqsRequirementTypeTypeAddress  DocReqsRequirementTypeType = "address"
	DocReqsRequirementTypeTypeTextual  DocReqsRequirementTypeType = "textual"
)

type HostedNumber struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The messaging hosted phone number (+E.164 format)
	PhoneNumber string `json:"phone_number" format:"+E.164"`
	RecordType  string `json:"record_type"`
	// Any of "deleted", "failed", "failed_activation", "failed_carrier_rejected",
	// "failed_ineligible_carrier", "failed_number_already_hosted",
	// "failed_number_not_found", "failed_ownership_verification", "failed_timeout",
	// "pending", "provisioning", "successful".
	Status HostedNumberStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HostedNumber) RawJSON() string { return r.JSON.raw }
func (r *HostedNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HostedNumberStatus string

const (
	HostedNumberStatusDeleted                     HostedNumberStatus = "deleted"
	HostedNumberStatusFailed                      HostedNumberStatus = "failed"
	HostedNumberStatusFailedActivation            HostedNumberStatus = "failed_activation"
	HostedNumberStatusFailedCarrierRejected       HostedNumberStatus = "failed_carrier_rejected"
	HostedNumberStatusFailedIneligibleCarrier     HostedNumberStatus = "failed_ineligible_carrier"
	HostedNumberStatusFailedNumberAlreadyHosted   HostedNumberStatus = "failed_number_already_hosted"
	HostedNumberStatusFailedNumberNotFound        HostedNumberStatus = "failed_number_not_found"
	HostedNumberStatusFailedOwnershipVerification HostedNumberStatus = "failed_ownership_verification"
	HostedNumberStatusFailedTimeout               HostedNumberStatus = "failed_timeout"
	HostedNumberStatusPending                     HostedNumberStatus = "pending"
	HostedNumberStatusProvisioning                HostedNumberStatus = "provisioning"
	HostedNumberStatusSuccessful                  HostedNumberStatus = "successful"
)

type InboundMessagePayload struct {
	// Identifies the type of resource.
	ID string                    `json:"id" format:"uuid"`
	Cc []InboundMessagePayloadCc `json:"cc"`
	// Not used for inbound messages.
	CompletedAt time.Time                 `json:"completed_at,nullable" format:"date-time"`
	Cost        InboundMessagePayloadCost `json:"cost,nullable"`
	// Detailed breakdown of the message cost components.
	CostBreakdown InboundMessagePayloadCostBreakdown `json:"cost_breakdown,nullable"`
	// The direction of the message. Inbound messages are sent to you whereas outbound
	// messages are sent from you.
	//
	// Any of "inbound".
	Direction InboundMessagePayloadDirection `json:"direction"`
	// Encoding scheme used for the message body.
	Encoding string `json:"encoding"`
	// These errors may point at addressees when referring to unsuccessful/unconfirmed
	// delivery statuses.
	Errors []telnyx.MessagingError      `json:"errors"`
	From   InboundMessagePayloadFrom    `json:"from"`
	Media  []InboundMessagePayloadMedia `json:"media"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id"`
	// Unique identifier for a messaging profile.
	OrganizationID string `json:"organization_id"`
	// Number of parts into which the message's body must be split.
	Parts int64 `json:"parts"`
	// ISO 8601 formatted date indicating when the message request was received.
	ReceivedAt time.Time `json:"received_at" format:"date-time"`
	// Identifies the type of the resource.
	//
	// Any of "message".
	RecordType InboundMessagePayloadRecordType `json:"record_type"`
	// Not used for inbound messages.
	SentAt time.Time `json:"sent_at,nullable" format:"date-time"`
	// Message subject.
	Subject string `json:"subject,nullable"`
	// Tags associated with the resource.
	Tags []string `json:"tags"`
	// Indicates whether the TCR campaign is billable.
	TcrCampaignBillable bool `json:"tcr_campaign_billable"`
	// The Campaign Registry (TCR) campaign ID associated with the message.
	TcrCampaignID string `json:"tcr_campaign_id,nullable"`
	// The registration status of the TCR campaign.
	TcrCampaignRegistered string `json:"tcr_campaign_registered,nullable"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text string                    `json:"text"`
	To   []InboundMessagePayloadTo `json:"to"`
	// The type of message. This value can be either 'sms' or 'mms'.
	//
	// Any of "SMS", "MMS".
	Type InboundMessagePayloadType `json:"type"`
	// Not used for inbound messages.
	ValidUntil time.Time `json:"valid_until,nullable" format:"date-time"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url,nullable" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL string `json:"webhook_url,nullable" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Cc                    respjson.Field
		CompletedAt           respjson.Field
		Cost                  respjson.Field
		CostBreakdown         respjson.Field
		Direction             respjson.Field
		Encoding              respjson.Field
		Errors                respjson.Field
		From                  respjson.Field
		Media                 respjson.Field
		MessagingProfileID    respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WebhookFailoverURL    respjson.Field
		WebhookURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayload) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (InboundMessagePayload) ImplMessageGetResponseDataUnion() {}

type InboundMessagePayloadCc struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// Any of "queued", "sending", "sent", "delivered", "sending_failed",
	// "delivery_failed", "delivery_unconfirmed".
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
func (r InboundMessagePayloadCc) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessagePayloadCost struct {
	// The amount deducted from your account.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayloadCost) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of the message cost components.
type InboundMessagePayloadCostBreakdown struct {
	CarrierFee InboundMessagePayloadCostBreakdownCarrierFee `json:"carrier_fee"`
	Rate       InboundMessagePayloadCostBreakdownRate       `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierFee  respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayloadCostBreakdown) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessagePayloadCostBreakdownCarrierFee struct {
	// The carrier fee amount.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayloadCostBreakdownCarrierFee) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessagePayloadCostBreakdownRate struct {
	// The rate amount applied.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayloadCostBreakdownRate) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the message. Inbound messages are sent to you whereas outbound
// messages are sent from you.
type InboundMessagePayloadDirection string

const (
	InboundMessagePayloadDirectionInbound InboundMessagePayloadDirection = "inbound"
)

type InboundMessagePayloadFrom struct {
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
func (r InboundMessagePayloadFrom) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessagePayloadMedia struct {
	// The MIME type of the requested media.
	ContentType string `json:"content_type" format:"mime-type"`
	// The SHA256 hash of the requested media.
	HashSha256 string `json:"hash_sha256"`
	// The size of the requested media.
	Size int64 `json:"size"`
	// The url of the media requested to be sent.
	URL string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		HashSha256  respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessagePayloadMedia) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type InboundMessagePayloadRecordType string

const (
	InboundMessagePayloadRecordTypeMessage InboundMessagePayloadRecordType = "message"
)

type InboundMessagePayloadTo struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// Any of "queued", "sending", "sent", "delivered", "sending_failed",
	// "delivery_failed", "delivery_unconfirmed", "webhook_delivered".
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
func (r InboundMessagePayloadTo) RawJSON() string { return r.JSON.raw }
func (r *InboundMessagePayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of message. This value can be either 'sms' or 'mms'.
type InboundMessagePayloadType string

const (
	InboundMessagePayloadTypeSMS InboundMessagePayloadType = "SMS"
	InboundMessagePayloadTypeMms InboundMessagePayloadType = "MMS"
)

// The set of features available for a specific messaging use case (SMS or MMS).
// Features can vary depending on the characteristics the phone number, as well as
// its current product configuration.
type MessagingFeatureSet struct {
	// Send messages to and receive messages from numbers in the same country.
	DomesticTwoWay bool `json:"domestic_two_way,required"`
	// Receive messages from numbers in other countries.
	InternationalInbound bool `json:"international_inbound,required"`
	// Send messages to numbers in other countries.
	InternationalOutbound bool `json:"international_outbound,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DomesticTwoWay        respjson.Field
		InternationalInbound  respjson.Field
		InternationalOutbound respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingFeatureSet) RawJSON() string { return r.JSON.raw }
func (r *MessagingFeatureSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrder struct {
	// Resource unique identifier.
	ID string `json:"id" format:"uuid"`
	// Automatically associate the number with this messaging profile ID when the order
	// is complete.
	MessagingProfileID string         `json:"messaging_profile_id,nullable"`
	PhoneNumbers       []HostedNumber `json:"phone_numbers"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Any of "carrier_rejected", "compliance_review_failed", "deleted", "failed",
	// "incomplete_documentation", "incorrect_billing_information",
	// "ineligible_carrier", "loa_file_invalid", "loa_file_successful", "pending",
	// "provisioning", "successful".
	Status MessagingHostedNumberOrderStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumbers       respjson.Field
		RecordType         respjson.Field
		Status             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberOrder) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberOrderStatus string

const (
	MessagingHostedNumberOrderStatusCarrierRejected             MessagingHostedNumberOrderStatus = "carrier_rejected"
	MessagingHostedNumberOrderStatusComplianceReviewFailed      MessagingHostedNumberOrderStatus = "compliance_review_failed"
	MessagingHostedNumberOrderStatusDeleted                     MessagingHostedNumberOrderStatus = "deleted"
	MessagingHostedNumberOrderStatusFailed                      MessagingHostedNumberOrderStatus = "failed"
	MessagingHostedNumberOrderStatusIncompleteDocumentation     MessagingHostedNumberOrderStatus = "incomplete_documentation"
	MessagingHostedNumberOrderStatusIncorrectBillingInformation MessagingHostedNumberOrderStatus = "incorrect_billing_information"
	MessagingHostedNumberOrderStatusIneligibleCarrier           MessagingHostedNumberOrderStatus = "ineligible_carrier"
	MessagingHostedNumberOrderStatusLoaFileInvalid              MessagingHostedNumberOrderStatus = "loa_file_invalid"
	MessagingHostedNumberOrderStatusLoaFileSuccessful           MessagingHostedNumberOrderStatus = "loa_file_successful"
	MessagingHostedNumberOrderStatusPending                     MessagingHostedNumberOrderStatus = "pending"
	MessagingHostedNumberOrderStatusProvisioning                MessagingHostedNumberOrderStatus = "provisioning"
	MessagingHostedNumberOrderStatusSuccessful                  MessagingHostedNumberOrderStatus = "successful"
)

type MessagingPaginationMeta struct {
	PageNumber   int64 `json:"page_number,required"`
	PageSize     int64 `json:"page_size,required"`
	TotalPages   int64 `json:"total_pages,required"`
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *MessagingPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Metadata struct {
	// Current Page based on pagination settings (included when defaults are used.)
	PageNumber int64 `json:"page_number,required"`
	// Total number of pages based on pagination settings
	TotalPages int64 `json:"total_pages,required"`
	// Number of results to return per page based on pagination settings (included when
	// defaults are used.)
	PageSize int64 `json:"page_size"`
	// Total number of results
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// High level health metrics about the number and it's messaging sending patterns.
type NumberHealthMetrics struct {
	// The ratio of messages received to the number of messages sent.
	InboundOutboundRatio float64 `json:"inbound_outbound_ratio,required"`
	// The number of messages analyzed for the health metrics.
	MessageCount int64 `json:"message_count,required"`
	// The ratio of messages blocked for spam to the number of messages attempted.
	SpamRatio float64 `json:"spam_ratio,required"`
	// The ratio of messages sucessfully delivered to the number of messages attempted.
	SuccessRatio float64 `json:"success_ratio,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundOutboundRatio respjson.Field
		MessageCount         respjson.Field
		SpamRatio            respjson.Field
		SuccessRatio         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberHealthMetrics) RawJSON() string { return r.JSON.raw }
func (r *NumberHealthMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberWithMessagingSettings struct {
	// Identifies the type of resource.
	ID string `json:"id"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The messaging products that this number can be registered to use
	EligibleMessagingProducts []string                                 `json:"eligible_messaging_products"`
	Features                  PhoneNumberWithMessagingSettingsFeatures `json:"features"`
	// High level health metrics about the number and it's messaging sending patterns.
	Health NumberHealthMetrics `json:"health"`
	// The messaging product that the number is registered to use
	MessagingProduct string `json:"messaging_product"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id,nullable"`
	// +E.164 formatted phone number.
	PhoneNumber string `json:"phone_number" format:"e164"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_phone_number", "messaging_settings".
	RecordType PhoneNumberWithMessagingSettingsRecordType `json:"record_type"`
	// The messaging traffic or use case for which the number is currently configured.
	TrafficType string `json:"traffic_type"`
	// The type of the phone number
	//
	// Any of "long-code", "toll-free", "short-code", "longcode", "tollfree",
	// "shortcode".
	Type PhoneNumberWithMessagingSettingsType `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CountryCode               respjson.Field
		CreatedAt                 respjson.Field
		EligibleMessagingProducts respjson.Field
		Features                  respjson.Field
		Health                    respjson.Field
		MessagingProduct          respjson.Field
		MessagingProfileID        respjson.Field
		PhoneNumber               respjson.Field
		RecordType                respjson.Field
		TrafficType               respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberWithMessagingSettings) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberWithMessagingSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberWithMessagingSettingsFeatures struct {
	// The set of features available for a specific messaging use case (SMS or MMS).
	// Features can vary depending on the characteristics the phone number, as well as
	// its current product configuration.
	Mms MessagingFeatureSet `json:"mms,nullable"`
	// The set of features available for a specific messaging use case (SMS or MMS).
	// Features can vary depending on the characteristics the phone number, as well as
	// its current product configuration.
	SMS MessagingFeatureSet `json:"sms,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mms         respjson.Field
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberWithMessagingSettingsFeatures) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberWithMessagingSettingsFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type PhoneNumberWithMessagingSettingsRecordType string

const (
	PhoneNumberWithMessagingSettingsRecordTypeMessagingPhoneNumber PhoneNumberWithMessagingSettingsRecordType = "messaging_phone_number"
	PhoneNumberWithMessagingSettingsRecordTypeMessagingSettings    PhoneNumberWithMessagingSettingsRecordType = "messaging_settings"
)

// The type of the phone number
type PhoneNumberWithMessagingSettingsType string

const (
	PhoneNumberWithMessagingSettingsTypeLongCode  PhoneNumberWithMessagingSettingsType = "long-code"
	PhoneNumberWithMessagingSettingsTypeTollFree  PhoneNumberWithMessagingSettingsType = "toll-free"
	PhoneNumberWithMessagingSettingsTypeShortCode PhoneNumberWithMessagingSettingsType = "short-code"
	PhoneNumberWithMessagingSettingsTypeLongcode  PhoneNumberWithMessagingSettingsType = "longcode"
	PhoneNumberWithMessagingSettingsTypeTollfree  PhoneNumberWithMessagingSettingsType = "tollfree"
	PhoneNumberWithMessagingSettingsTypeShortcode PhoneNumberWithMessagingSettingsType = "shortcode"
)

// Porting order status
type PortingOrderStatus struct {
	// A list of 0 or more details about this porting order's status
	Details []PortingOrdersExceptionType `json:"details"`
	// The current status of the porting order
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "ported", "cancelled", "cancel-pending".
	Value PortingOrderStatusValue `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Details     respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderStatus) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the porting order
type PortingOrderStatusValue string

const (
	PortingOrderStatusValueDraft            PortingOrderStatusValue = "draft"
	PortingOrderStatusValueInProcess        PortingOrderStatusValue = "in-process"
	PortingOrderStatusValueSubmitted        PortingOrderStatusValue = "submitted"
	PortingOrderStatusValueException        PortingOrderStatusValue = "exception"
	PortingOrderStatusValueFocDateConfirmed PortingOrderStatusValue = "foc-date-confirmed"
	PortingOrderStatusValuePorted           PortingOrderStatusValue = "ported"
	PortingOrderStatusValueCancelled        PortingOrderStatusValue = "cancelled"
	PortingOrderStatusValueCancelPending    PortingOrderStatusValue = "cancel-pending"
)

type PortingOrdersExceptionType struct {
	// Identifier of an exception type
	//
	// Any of "ACCOUNT_NUMBER_MISMATCH", "AUTH_PERSON_MISMATCH", "BTN_ATN_MISMATCH",
	// "ENTITY_NAME_MISMATCH", "FOC_EXPIRED", "FOC_REJECTED", "LOCATION_MISMATCH",
	// "LSR_PENDING", "MAIN_BTN_PORTING", "OSP_IRRESPONSIVE", "OTHER",
	// "PASSCODE_PIN_INVALID", "PHONE_NUMBER_HAS_SPECIAL_FEATURE",
	// "PHONE_NUMBER_MISMATCH", "PHONE_NUMBER_NOT_PORTABLE", "PORT_TYPE_INCORRECT",
	// "PORTING_ORDER_SPLIT_REQUIRED", "POSTAL_CODE_MISMATCH",
	// "RATE_CENTER_NOT_PORTABLE", "SV_CONFLICT", "SV_UNKNOWN_FAILURE".
	Code PortingOrdersExceptionTypeCode `json:"code"`
	// Description of an exception type
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrdersExceptionType) RawJSON() string { return r.JSON.raw }
func (r *PortingOrdersExceptionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of an exception type
type PortingOrdersExceptionTypeCode string

const (
	PortingOrdersExceptionTypeCodeAccountNumberMismatch        PortingOrdersExceptionTypeCode = "ACCOUNT_NUMBER_MISMATCH"
	PortingOrdersExceptionTypeCodeAuthPersonMismatch           PortingOrdersExceptionTypeCode = "AUTH_PERSON_MISMATCH"
	PortingOrdersExceptionTypeCodeBtnAtnMismatch               PortingOrdersExceptionTypeCode = "BTN_ATN_MISMATCH"
	PortingOrdersExceptionTypeCodeEntityNameMismatch           PortingOrdersExceptionTypeCode = "ENTITY_NAME_MISMATCH"
	PortingOrdersExceptionTypeCodeFocExpired                   PortingOrdersExceptionTypeCode = "FOC_EXPIRED"
	PortingOrdersExceptionTypeCodeFocRejected                  PortingOrdersExceptionTypeCode = "FOC_REJECTED"
	PortingOrdersExceptionTypeCodeLocationMismatch             PortingOrdersExceptionTypeCode = "LOCATION_MISMATCH"
	PortingOrdersExceptionTypeCodeLsrPending                   PortingOrdersExceptionTypeCode = "LSR_PENDING"
	PortingOrdersExceptionTypeCodeMainBtnPorting               PortingOrdersExceptionTypeCode = "MAIN_BTN_PORTING"
	PortingOrdersExceptionTypeCodeOspIrresponsive              PortingOrdersExceptionTypeCode = "OSP_IRRESPONSIVE"
	PortingOrdersExceptionTypeCodeOther                        PortingOrdersExceptionTypeCode = "OTHER"
	PortingOrdersExceptionTypeCodePasscodePinInvalid           PortingOrdersExceptionTypeCode = "PASSCODE_PIN_INVALID"
	PortingOrdersExceptionTypeCodePhoneNumberHasSpecialFeature PortingOrdersExceptionTypeCode = "PHONE_NUMBER_HAS_SPECIAL_FEATURE"
	PortingOrdersExceptionTypeCodePhoneNumberMismatch          PortingOrdersExceptionTypeCode = "PHONE_NUMBER_MISMATCH"
	PortingOrdersExceptionTypeCodePhoneNumberNotPortable       PortingOrdersExceptionTypeCode = "PHONE_NUMBER_NOT_PORTABLE"
	PortingOrdersExceptionTypeCodePortTypeIncorrect            PortingOrdersExceptionTypeCode = "PORT_TYPE_INCORRECT"
	PortingOrdersExceptionTypeCodePortingOrderSplitRequired    PortingOrdersExceptionTypeCode = "PORTING_ORDER_SPLIT_REQUIRED"
	PortingOrdersExceptionTypeCodePostalCodeMismatch           PortingOrdersExceptionTypeCode = "POSTAL_CODE_MISMATCH"
	PortingOrdersExceptionTypeCodeRateCenterNotPortable        PortingOrdersExceptionTypeCode = "RATE_CENTER_NOT_PORTABLE"
	PortingOrdersExceptionTypeCodeSvConflict                   PortingOrdersExceptionTypeCode = "SV_CONFLICT"
	PortingOrdersExceptionTypeCodeSvUnknownFailure             PortingOrdersExceptionTypeCode = "SV_UNKNOWN_FAILURE"
)

type RoomParticipant struct {
	// A unique identifier for the room participant.
	ID string `json:"id" format:"uuid"`
	// Context provided to the given participant through the client SDK
	Context string `json:"context"`
	// ISO 8601 timestamp when the participant joined the session.
	JoinedAt time.Time `json:"joined_at" format:"date-time"`
	// ISO 8601 timestamp when the participant left the session.
	LeftAt     time.Time `json:"left_at" format:"date-time"`
	RecordType string    `json:"record_type"`
	// Identify the room session that participant is part of.
	SessionID string `json:"session_id" format:"uuid"`
	// ISO 8601 timestamp when the participant was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Context     respjson.Field
		JoinedAt    respjson.Field
		LeftAt      respjson.Field
		RecordType  respjson.Field
		SessionID   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RoomParticipant) RawJSON() string { return r.JSON.raw }
func (r *RoomParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortCode struct {
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the type of the resource.
	//
	// Any of "short_code".
	RecordType ShortCodeRecordType `json:"record_type"`
	// Short digit sequence used to address messages.
	ShortCode string   `json:"short_code"`
	Tags      []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessagingProfileID respjson.Field
		ID                 respjson.Field
		CountryCode        respjson.Field
		CreatedAt          respjson.Field
		RecordType         respjson.Field
		ShortCode          respjson.Field
		Tags               respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortCode) RawJSON() string { return r.JSON.raw }
func (r *ShortCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ShortCodeRecordType string

const (
	ShortCodeRecordTypeShortCode ShortCodeRecordType = "short_code"
)

type SimCardStatus struct {
	// It describes why the SIM card is in the current status.
	Reason string `json:"reason"`
	// The current status of the SIM card. It will be one of the following: <br/>
	//
	// <ul>
	//
	//	<li><code>registering</code> - the card is being registered</li>
	//	<li><code>enabling</code> - the card is being enabled</li>
	//	<li><code>enabled</code> - the card is enabled and ready for use</li>
	//	<li><code>disabling</code> - the card is being disabled</li>
	//	<li><code>disabled</code> - the card has been disabled and cannot be used</li>
	//	<li><code>data_limit_exceeded</code> - the card has exceeded its data consumption limit</li>
	//	<li><code>setting_standby</code> - the process to set the card in stand by is in progress</li>
	//	<li><code>standby</code> - the card is in stand by</li>
	//
	// </ul>
	// Transitioning between the enabled and disabled states may take a period of time.
	//
	// Any of "registering", "enabling", "enabled", "disabling", "disabled",
	// "data_limit_exceeded", "setting_standby", "standby".
	Value SimCardStatusValue `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimCardStatus) RawJSON() string { return r.JSON.raw }
func (r *SimCardStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SimCardStatus to a SimCardStatusParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SimCardStatusParam.Overrides()
func (r SimCardStatus) ToParam() SimCardStatusParam {
	return param.Override[SimCardStatusParam](json.RawMessage(r.RawJSON()))
}

// The current status of the SIM card. It will be one of the following: <br/>
//
// <ul>
//
//	<li><code>registering</code> - the card is being registered</li>
//	<li><code>enabling</code> - the card is being enabled</li>
//	<li><code>enabled</code> - the card is enabled and ready for use</li>
//	<li><code>disabling</code> - the card is being disabled</li>
//	<li><code>disabled</code> - the card has been disabled and cannot be used</li>
//	<li><code>data_limit_exceeded</code> - the card has exceeded its data consumption limit</li>
//	<li><code>setting_standby</code> - the process to set the card in stand by is in progress</li>
//	<li><code>standby</code> - the card is in stand by</li>
//
// </ul>
// Transitioning between the enabled and disabled states may take a period of time.
type SimCardStatusValue string

const (
	SimCardStatusValueRegistering       SimCardStatusValue = "registering"
	SimCardStatusValueEnabling          SimCardStatusValue = "enabling"
	SimCardStatusValueEnabled           SimCardStatusValue = "enabled"
	SimCardStatusValueDisabling         SimCardStatusValue = "disabling"
	SimCardStatusValueDisabled          SimCardStatusValue = "disabled"
	SimCardStatusValueDataLimitExceeded SimCardStatusValue = "data_limit_exceeded"
	SimCardStatusValueSettingStandby    SimCardStatusValue = "setting_standby"
	SimCardStatusValueStandby           SimCardStatusValue = "standby"
)

type SimCardStatusParam struct {
	paramObj
}

func (r SimCardStatusParam) MarshalJSON() (data []byte, err error) {
	type shadow SimCardStatusParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimCardStatusParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimpleSimCard struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// Indicate whether the SIM card has any pending (in-progress) actions.
	ActionsInProgress bool `json:"actions_in_progress"`
	// List of IMEIs authorized to use a given SIM card.
	AuthorizedImeis []string `json:"authorized_imeis,nullable"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The SIM card consumption so far in the current billing cycle.
	CurrentBillingPeriodConsumedData SimpleSimCardCurrentBillingPeriodConsumedData `json:"current_billing_period_consumed_data"`
	// The SIM card individual data limit configuration.
	DataLimit SimpleSimCardDataLimit `json:"data_limit"`
	// The Embedded Identity Document (eID) for eSIM cards.
	Eid string `json:"eid,nullable"`
	// The installation status of the eSIM. Only applicable for eSIM cards.
	//
	// Any of "released", "disabled".
	EsimInstallationStatus SimpleSimCardEsimInstallationStatus `json:"esim_installation_status,nullable"`
	// The ICCID is the identifier of the specific SIM card/chip. Each SIM is
	// internationally identified by its integrated circuit card identifier (ICCID).
	// ICCIDs are stored in the SIM card's memory and are also engraved or printed on
	// the SIM card body during a process called personalization.
	Iccid string `json:"iccid"`
	// SIM cards are identified on their individual network operators by a unique
	// International Mobile Subscriber Identity (IMSI). <br/> Mobile network operators
	// connect mobile phone calls and communicate with their market SIM cards using
	// their IMSIs. The IMSI is stored in the Subscriber Identity Module (SIM) inside
	// the device and is sent by the device to the appropriate network. It is used to
	// acquire the details of the device in the Home Location Register (HLR) or the
	// Visitor Location Register (VLR).
	Imsi string `json:"imsi"`
	// Mobile Station International Subscriber Directory Number (MSISDN) is a number
	// used to identify a mobile phone number internationally. <br/> MSISDN is defined
	// by the E.164 numbering plan. It includes a country code and a National
	// Destination Code which identifies the subscriber's operator.
	Msisdn     string `json:"msisdn"`
	RecordType string `json:"record_type"`
	// List of resources with actions in progress.
	ResourcesWithInProgressActions []map[string]any `json:"resources_with_in_progress_actions"`
	// The group SIMCardGroup identification. This attribute can be <code>null</code>
	// when it's present in an associated resource.
	SimCardGroupID string        `json:"sim_card_group_id" format:"uuid"`
	Status         SimCardStatus `json:"status"`
	// Searchable tags associated with the SIM card
	Tags []string `json:"tags"`
	// The type of SIM card
	//
	// Any of "physical", "esim".
	Type SimpleSimCardType `json:"type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// The version of the SIM card.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		ActionsInProgress                respjson.Field
		AuthorizedImeis                  respjson.Field
		CreatedAt                        respjson.Field
		CurrentBillingPeriodConsumedData respjson.Field
		DataLimit                        respjson.Field
		Eid                              respjson.Field
		EsimInstallationStatus           respjson.Field
		Iccid                            respjson.Field
		Imsi                             respjson.Field
		Msisdn                           respjson.Field
		RecordType                       respjson.Field
		ResourcesWithInProgressActions   respjson.Field
		SimCardGroupID                   respjson.Field
		Status                           respjson.Field
		Tags                             respjson.Field
		Type                             respjson.Field
		UpdatedAt                        respjson.Field
		Version                          respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimpleSimCard) RawJSON() string { return r.JSON.raw }
func (r *SimpleSimCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card consumption so far in the current billing cycle.
type SimpleSimCardCurrentBillingPeriodConsumedData struct {
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimpleSimCardCurrentBillingPeriodConsumedData) RawJSON() string { return r.JSON.raw }
func (r *SimpleSimCardCurrentBillingPeriodConsumedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SIM card individual data limit configuration.
type SimpleSimCardDataLimit struct {
	Amount string `json:"amount"`
	// Any of "MB", "GB".
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimpleSimCardDataLimit) RawJSON() string { return r.JSON.raw }
func (r *SimpleSimCardDataLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The installation status of the eSIM. Only applicable for eSIM cards.
type SimpleSimCardEsimInstallationStatus string

const (
	SimpleSimCardEsimInstallationStatusReleased SimpleSimCardEsimInstallationStatus = "released"
	SimpleSimCardEsimInstallationStatusDisabled SimpleSimCardEsimInstallationStatus = "disabled"
)

// The type of SIM card
type SimpleSimCardType string

const (
	SimpleSimCardTypePhysical SimpleSimCardType = "physical"
	SimpleSimCardTypeEsim     SimpleSimCardType = "esim"
)

type SubNumberOrderRegulatoryRequirementWithValue struct {
	// Any of "textual", "datetime", "address", "document".
	FieldType SubNumberOrderRegulatoryRequirementWithValueFieldType `json:"field_type"`
	// The value of the requirement, this could be an id to a resource or a string
	// value.
	FieldValue string `json:"field_value"`
	RecordType string `json:"record_type"`
	// Unique id for a requirement.
	RequirementID string `json:"requirement_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType     respjson.Field
		FieldValue    respjson.Field
		RecordType    respjson.Field
		RequirementID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubNumberOrderRegulatoryRequirementWithValue) RawJSON() string { return r.JSON.raw }
func (r *SubNumberOrderRegulatoryRequirementWithValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubNumberOrderRegulatoryRequirementWithValueFieldType string

const (
	SubNumberOrderRegulatoryRequirementWithValueFieldTypeTextual  SubNumberOrderRegulatoryRequirementWithValueFieldType = "textual"
	SubNumberOrderRegulatoryRequirementWithValueFieldTypeDatetime SubNumberOrderRegulatoryRequirementWithValueFieldType = "datetime"
	SubNumberOrderRegulatoryRequirementWithValueFieldTypeAddress  SubNumberOrderRegulatoryRequirementWithValueFieldType = "address"
	SubNumberOrderRegulatoryRequirementWithValueFieldTypeDocument SubNumberOrderRegulatoryRequirementWithValueFieldType = "document"
)
