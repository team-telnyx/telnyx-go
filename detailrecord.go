// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// DetailRecordService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDetailRecordService] method instead.
type DetailRecordService struct {
	Options []option.RequestOption
}

// NewDetailRecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDetailRecordService(opts ...option.RequestOption) (r DetailRecordService) {
	r = DetailRecordService{}
	r.Options = opts
	return
}

// Search for any detail record across the Telnyx Platform
func (r *DetailRecordService) List(ctx context.Context, query DetailRecordListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DetailRecordListResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "detail_records"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Search for any detail record across the Telnyx Platform
func (r *DetailRecordService) ListAutoPaging(ctx context.Context, query DetailRecordListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DetailRecordListResponseUnion] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// DetailRecordListResponseUnion contains all possible properties and values from
// [DetailRecordListResponseMessage], [DetailRecordListResponseConference],
// [DetailRecordListResponseConferenceParticipant], [DetailRecordListResponseAmd],
// [DetailRecordListResponseVerify2Fa], [DetailRecordListResponseSimCardUsage],
// [DetailRecordListResponseMediaStorage].
//
// Use the [DetailRecordListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DetailRecordListResponseUnion struct {
	// Any of nil, nil, nil, nil, nil, nil, nil.
	RecordType string `json:"record_type"`
	// This field is from variant [DetailRecordListResponseMessage].
	Carrier string `json:"carrier"`
	// This field is from variant [DetailRecordListResponseMessage].
	CarrierFee string `json:"carrier_fee"`
	// This field is from variant [DetailRecordListResponseMessage].
	Cld string `json:"cld"`
	// This field is from variant [DetailRecordListResponseMessage].
	Cli string `json:"cli"`
	// This field is from variant [DetailRecordListResponseMessage].
	CompletedAt time.Time `json:"completed_at"`
	Cost        string    `json:"cost"`
	// This field is from variant [DetailRecordListResponseMessage].
	CountryCode    string    `json:"country_code"`
	CreatedAt      time.Time `json:"created_at"`
	Currency       string    `json:"currency"`
	DeliveryStatus string    `json:"delivery_status"`
	// This field is from variant [DetailRecordListResponseMessage].
	DeliveryStatusFailoverURL string `json:"delivery_status_failover_url"`
	// This field is from variant [DetailRecordListResponseMessage].
	DeliveryStatusWebhookURL string `json:"delivery_status_webhook_url"`
	// This field is from variant [DetailRecordListResponseMessage].
	Direction string `json:"direction"`
	// This field is from variant [DetailRecordListResponseMessage].
	Errors []string `json:"errors"`
	// This field is from variant [DetailRecordListResponseMessage].
	Fteu bool   `json:"fteu"`
	Mcc  string `json:"mcc"`
	// This field is from variant [DetailRecordListResponseMessage].
	MessageType string `json:"message_type"`
	Mnc         string `json:"mnc"`
	// This field is from variant [DetailRecordListResponseMessage].
	OnNet bool `json:"on_net"`
	// This field is from variant [DetailRecordListResponseMessage].
	Parts int64 `json:"parts"`
	// This field is from variant [DetailRecordListResponseMessage].
	ProfileID string `json:"profile_id"`
	// This field is from variant [DetailRecordListResponseMessage].
	ProfileName string `json:"profile_name"`
	Rate        string `json:"rate"`
	// This field is from variant [DetailRecordListResponseMessage].
	SentAt time.Time `json:"sent_at"`
	// This field is from variant [DetailRecordListResponseMessage].
	SourceCountryCode string    `json:"source_country_code"`
	Status            string    `json:"status"`
	Tags              string    `json:"tags"`
	UpdatedAt         time.Time `json:"updated_at"`
	UserID            string    `json:"user_id"`
	// This field is from variant [DetailRecordListResponseMessage].
	Uuid          string `json:"uuid"`
	ID            string `json:"id"`
	CallLegID     string `json:"call_leg_id"`
	CallSec       int64  `json:"call_sec"`
	CallSessionID string `json:"call_session_id"`
	ConnectionID  string `json:"connection_id"`
	// This field is from variant [DetailRecordListResponseConference].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [DetailRecordListResponseConference].
	ExpiresAt        time.Time `json:"expires_at"`
	IsTelnyxBillable bool      `json:"is_telnyx_billable"`
	// This field is from variant [DetailRecordListResponseConference].
	Name string `json:"name"`
	// This field is from variant [DetailRecordListResponseConference].
	ParticipantCallSec int64 `json:"participant_call_sec"`
	// This field is from variant [DetailRecordListResponseConference].
	ParticipantCount int64 `json:"participant_count"`
	// This field is from variant [DetailRecordListResponseConference].
	Region string `json:"region"`
	// This field is from variant [DetailRecordListResponseConference].
	StartedAt time.Time `json:"started_at"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	BilledSec int64 `json:"billed_sec"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	ConferenceID string `json:"conference_id"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	DestinationNumber string `json:"destination_number"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	JoinedAt time.Time `json:"joined_at"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	LeftAt time.Time `json:"left_at"`
	// This field is from variant [DetailRecordListResponseConferenceParticipant].
	OriginatingNumber string `json:"originating_number"`
	RateMeasuredIn    string `json:"rate_measured_in"`
	// This field is from variant [DetailRecordListResponseAmd].
	BillingGroupID string `json:"billing_group_id"`
	// This field is from variant [DetailRecordListResponseAmd].
	BillingGroupName string `json:"billing_group_name"`
	// This field is from variant [DetailRecordListResponseAmd].
	ConnectionName string `json:"connection_name"`
	// This field is from variant [DetailRecordListResponseAmd].
	Feature string `json:"feature"`
	// This field is from variant [DetailRecordListResponseAmd].
	InvokedAt time.Time `json:"invoked_at"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	DestinationPhoneNumber string `json:"destination_phone_number"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	VerificationStatus string `json:"verification_status"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	VerifyChannelID string `json:"verify_channel_id"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	VerifyChannelType string `json:"verify_channel_type"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	VerifyProfileID string `json:"verify_profile_id"`
	// This field is from variant [DetailRecordListResponseVerify2Fa].
	VerifyUsageFee string `json:"verify_usage_fee"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	ClosedAt time.Time `json:"closed_at"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	DataCost float64 `json:"data_cost"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	DataRate string `json:"data_rate"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	DataUnit string `json:"data_unit"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	DownlinkData float64 `json:"downlink_data"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	Imsi string `json:"imsi"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	IPAddress string `json:"ip_address"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	SimCardID string `json:"sim_card_id"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	SimCardTags string `json:"sim_card_tags"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	SimGroupID string `json:"sim_group_id"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	SimGroupName string `json:"sim_group_name"`
	// This field is from variant [DetailRecordListResponseSimCardUsage].
	UplinkData float64 `json:"uplink_data"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	ActionType string `json:"action_type"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	AssetID string `json:"asset_id"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	LinkChannelID string `json:"link_channel_id"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	LinkChannelType string `json:"link_channel_type"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	OrgID string `json:"org_id"`
	// This field is from variant [DetailRecordListResponseMediaStorage].
	WebhookID string `json:"webhook_id"`
	JSON      struct {
		RecordType                respjson.Field
		Carrier                   respjson.Field
		CarrierFee                respjson.Field
		Cld                       respjson.Field
		Cli                       respjson.Field
		CompletedAt               respjson.Field
		Cost                      respjson.Field
		CountryCode               respjson.Field
		CreatedAt                 respjson.Field
		Currency                  respjson.Field
		DeliveryStatus            respjson.Field
		DeliveryStatusFailoverURL respjson.Field
		DeliveryStatusWebhookURL  respjson.Field
		Direction                 respjson.Field
		Errors                    respjson.Field
		Fteu                      respjson.Field
		Mcc                       respjson.Field
		MessageType               respjson.Field
		Mnc                       respjson.Field
		OnNet                     respjson.Field
		Parts                     respjson.Field
		ProfileID                 respjson.Field
		ProfileName               respjson.Field
		Rate                      respjson.Field
		SentAt                    respjson.Field
		SourceCountryCode         respjson.Field
		Status                    respjson.Field
		Tags                      respjson.Field
		UpdatedAt                 respjson.Field
		UserID                    respjson.Field
		Uuid                      respjson.Field
		ID                        respjson.Field
		CallLegID                 respjson.Field
		CallSec                   respjson.Field
		CallSessionID             respjson.Field
		ConnectionID              respjson.Field
		EndedAt                   respjson.Field
		ExpiresAt                 respjson.Field
		IsTelnyxBillable          respjson.Field
		Name                      respjson.Field
		ParticipantCallSec        respjson.Field
		ParticipantCount          respjson.Field
		Region                    respjson.Field
		StartedAt                 respjson.Field
		BilledSec                 respjson.Field
		ConferenceID              respjson.Field
		DestinationNumber         respjson.Field
		JoinedAt                  respjson.Field
		LeftAt                    respjson.Field
		OriginatingNumber         respjson.Field
		RateMeasuredIn            respjson.Field
		BillingGroupID            respjson.Field
		BillingGroupName          respjson.Field
		ConnectionName            respjson.Field
		Feature                   respjson.Field
		InvokedAt                 respjson.Field
		DestinationPhoneNumber    respjson.Field
		VerificationStatus        respjson.Field
		VerifyChannelID           respjson.Field
		VerifyChannelType         respjson.Field
		VerifyProfileID           respjson.Field
		VerifyUsageFee            respjson.Field
		ClosedAt                  respjson.Field
		DataCost                  respjson.Field
		DataRate                  respjson.Field
		DataUnit                  respjson.Field
		DownlinkData              respjson.Field
		Imsi                      respjson.Field
		IPAddress                 respjson.Field
		PhoneNumber               respjson.Field
		SimCardID                 respjson.Field
		SimCardTags               respjson.Field
		SimGroupID                respjson.Field
		SimGroupName              respjson.Field
		UplinkData                respjson.Field
		ActionType                respjson.Field
		AssetID                   respjson.Field
		LinkChannelID             respjson.Field
		LinkChannelType           respjson.Field
		OrgID                     respjson.Field
		WebhookID                 respjson.Field
		raw                       string
	} `json:"-"`
}

func (u DetailRecordListResponseUnion) AsMessage() (v DetailRecordListResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsConference() (v DetailRecordListResponseConference) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsConferenceParticipant() (v DetailRecordListResponseConferenceParticipant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsAmd() (v DetailRecordListResponseAmd) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsVerify2Fa() (v DetailRecordListResponseVerify2Fa) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsSimCardUsage() (v DetailRecordListResponseSimCardUsage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DetailRecordListResponseUnion) AsMediaStorage() (v DetailRecordListResponseMediaStorage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DetailRecordListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *DetailRecordListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseMessage struct {
	// Identifies the record schema
	RecordType string `json:"record_type,required"`
	// Country-specific carrier used to send or receive the message
	Carrier string `json:"carrier"`
	// Fee charged by certain carriers in order to deliver certain message types.
	// Telnyx passes this fee on to the customer according to our pricing table
	CarrierFee string `json:"carrier_fee"`
	// The recipient of the message (to parameter in the Messaging API)
	Cld string `json:"cld"`
	// The sender of the message (from parameter in the Messaging API). For
	// Alphanumeric ID messages, this is the sender ID value
	Cli string `json:"cli"`
	// Message completion time
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// Amount, in the user currency, for the Telnyx billing cost
	Cost string `json:"cost"`
	// Two-letter representation of the country of the cld property using the ISO
	// 3166-1 alpha-2 format
	CountryCode string `json:"country_code"`
	// Message creation time
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency string `json:"currency"`
	// Final webhook delivery status
	DeliveryStatus string `json:"delivery_status"`
	// Failover customer-provided URL which Telnyx posts delivery status webhooks to
	DeliveryStatusFailoverURL string `json:"delivery_status_failover_url"`
	// Primary customer-provided URL which Telnyx posts delivery status webhooks to
	DeliveryStatusWebhookURL string `json:"delivery_status_webhook_url"`
	// Logical direction of the message from the Telnyx customer's perspective. It's
	// inbound when the Telnyx customer receives the message, or outbound otherwise
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Telnyx API error codes returned by the Telnyx gateway
	Errors []string `json:"errors"`
	// Indicates whether this is a Free-To-End-User (FTEU) short code message
	Fteu bool `json:"fteu"`
	// Mobile country code. Only available for certain products, such as Global
	// Outbound-Only from Alphanumeric Sender ID
	Mcc string `json:"mcc"`
	// Describes the Messaging service used to send the message. Available services
	// are: Short Message Service (SMS), Multimedia Messaging Service (MMS), and Rich
	// Communication Services (RCS)
	//
	// Any of "SMS", "MMS", "RCS".
	MessageType string `json:"message_type"`
	// Mobile network code. Only available for certain products, such as Global
	// Outbound-Only from Alphanumeric Sender ID
	Mnc string `json:"mnc"`
	// Indicates whether both sender and recipient numbers are Telnyx-managed
	OnNet bool `json:"on_net"`
	// Number of message parts. The message is broken down in multiple parts when its
	// length surpasses the limit of 160 characters
	Parts int64 `json:"parts"`
	// Unique identifier of the Messaging Profile used to send or receive the message
	ProfileID string `json:"profile_id"`
	// Name of the Messaging Profile used to send or receive the message
	ProfileName string `json:"profile_name"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	Rate string `json:"rate"`
	// Time when the message was sent
	SentAt time.Time `json:"sent_at" format:"date-time"`
	// Two-letter representation of the country of the cli property using the ISO
	// 3166-1 alpha-2 format
	SourceCountryCode string `json:"source_country_code"`
	// Final status of the message after the delivery attempt
	//
	// Any of "gw_timeout", "delivered", "dlr_unconfirmed", "dlr_timeout", "received",
	// "gw_reject", "failed".
	Status string `json:"status"`
	// Comma-separated tags assigned to the Telnyx number associated with the message
	Tags string `json:"tags"`
	// Message updated time
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Identifier of the Telnyx account who owns the message
	UserID string `json:"user_id"`
	// Unique identifier of the message
	Uuid string `json:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType                respjson.Field
		Carrier                   respjson.Field
		CarrierFee                respjson.Field
		Cld                       respjson.Field
		Cli                       respjson.Field
		CompletedAt               respjson.Field
		Cost                      respjson.Field
		CountryCode               respjson.Field
		CreatedAt                 respjson.Field
		Currency                  respjson.Field
		DeliveryStatus            respjson.Field
		DeliveryStatusFailoverURL respjson.Field
		DeliveryStatusWebhookURL  respjson.Field
		Direction                 respjson.Field
		Errors                    respjson.Field
		Fteu                      respjson.Field
		Mcc                       respjson.Field
		MessageType               respjson.Field
		Mnc                       respjson.Field
		OnNet                     respjson.Field
		Parts                     respjson.Field
		ProfileID                 respjson.Field
		ProfileName               respjson.Field
		Rate                      respjson.Field
		SentAt                    respjson.Field
		SourceCountryCode         respjson.Field
		Status                    respjson.Field
		Tags                      respjson.Field
		UpdatedAt                 respjson.Field
		UserID                    respjson.Field
		Uuid                      respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseConference struct {
	RecordType string `json:"record_type,required"`
	// Conference id
	ID string `json:"id"`
	// Telnyx UUID that identifies the conference call leg
	CallLegID string `json:"call_leg_id"`
	// Duration of the conference call in seconds
	CallSec int64 `json:"call_sec"`
	// Telnyx UUID that identifies with conference call session
	CallSessionID string `json:"call_session_id"`
	// Connection id
	ConnectionID string `json:"connection_id"`
	// Conference end time
	EndedAt time.Time `json:"ended_at" format:"date-time"`
	// Conference expiry time
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Indicates whether Telnyx billing charges might be applicable
	IsTelnyxBillable bool `json:"is_telnyx_billable"`
	// Conference name
	Name string `json:"name"`
	// Sum of the conference call duration for all participants in seconds
	ParticipantCallSec int64 `json:"participant_call_sec"`
	// Number of participants that joined the conference call
	ParticipantCount int64 `json:"participant_count"`
	// Region where the conference is hosted
	Region string `json:"region"`
	// Conference start time
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// User id
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType         respjson.Field
		ID                 respjson.Field
		CallLegID          respjson.Field
		CallSec            respjson.Field
		CallSessionID      respjson.Field
		ConnectionID       respjson.Field
		EndedAt            respjson.Field
		ExpiresAt          respjson.Field
		IsTelnyxBillable   respjson.Field
		Name               respjson.Field
		ParticipantCallSec respjson.Field
		ParticipantCount   respjson.Field
		Region             respjson.Field
		StartedAt          respjson.Field
		UserID             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseConference) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseConference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseConferenceParticipant struct {
	RecordType string `json:"record_type,required"`
	// Participant id
	ID string `json:"id"`
	// Duration of the conference call for billing purposes
	BilledSec int64 `json:"billed_sec"`
	// Telnyx UUID that identifies the conference call leg
	CallLegID string `json:"call_leg_id"`
	// Duration of the conference call in seconds
	CallSec int64 `json:"call_sec"`
	// Telnyx UUID that identifies with conference call session
	CallSessionID string `json:"call_session_id"`
	// Conference id
	ConferenceID string `json:"conference_id"`
	// Currency amount for Telnyx billing cost
	Cost string `json:"cost"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency string `json:"currency"`
	// Number called by the participant to join the conference
	DestinationNumber string `json:"destination_number"`
	// Indicates whether Telnyx billing charges might be applicable
	IsTelnyxBillable bool `json:"is_telnyx_billable"`
	// Participant join time
	JoinedAt time.Time `json:"joined_at" format:"date-time"`
	// Participant leave time
	LeftAt time.Time `json:"left_at" format:"date-time"`
	// Participant origin number used in the conference call
	OriginatingNumber string `json:"originating_number"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	Rate string `json:"rate"`
	// Billing unit used to calculate the Telnyx billing cost
	RateMeasuredIn string `json:"rate_measured_in"`
	// User id
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType        respjson.Field
		ID                respjson.Field
		BilledSec         respjson.Field
		CallLegID         respjson.Field
		CallSec           respjson.Field
		CallSessionID     respjson.Field
		ConferenceID      respjson.Field
		Cost              respjson.Field
		Currency          respjson.Field
		DestinationNumber respjson.Field
		IsTelnyxBillable  respjson.Field
		JoinedAt          respjson.Field
		LeftAt            respjson.Field
		OriginatingNumber respjson.Field
		Rate              respjson.Field
		RateMeasuredIn    respjson.Field
		UserID            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseConferenceParticipant) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseConferenceParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseAmd struct {
	RecordType string `json:"record_type,required"`
	// Feature invocation id
	ID string `json:"id"`
	// Billing Group id
	BillingGroupID string `json:"billing_group_id"`
	// Name of the Billing Group specified in billing_group_id
	BillingGroupName string `json:"billing_group_name"`
	// Telnyx UUID that identifies the related call leg
	CallLegID string `json:"call_leg_id"`
	// Telnyx UUID that identifies the related call session
	CallSessionID string `json:"call_session_id"`
	// Connection id
	ConnectionID string `json:"connection_id"`
	// Connection name
	ConnectionName string `json:"connection_name"`
	// Currency amount for Telnyx billing cost
	Cost string `json:"cost"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency string `json:"currency"`
	// Feature name
	//
	// Any of "PREMIUM".
	Feature string `json:"feature"`
	// Feature invocation time
	InvokedAt time.Time `json:"invoked_at" format:"date-time"`
	// Indicates whether Telnyx billing charges might be applicable
	IsTelnyxBillable bool `json:"is_telnyx_billable"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	Rate string `json:"rate"`
	// Billing unit used to calculate the Telnyx billing cost
	RateMeasuredIn string `json:"rate_measured_in"`
	// User-provided tags
	Tags string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType       respjson.Field
		ID               respjson.Field
		BillingGroupID   respjson.Field
		BillingGroupName respjson.Field
		CallLegID        respjson.Field
		CallSessionID    respjson.Field
		ConnectionID     respjson.Field
		ConnectionName   respjson.Field
		Cost             respjson.Field
		Currency         respjson.Field
		Feature          respjson.Field
		InvokedAt        respjson.Field
		IsTelnyxBillable respjson.Field
		Rate             respjson.Field
		RateMeasuredIn   respjson.Field
		Tags             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseAmd) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseAmd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseVerify2Fa struct {
	RecordType string `json:"record_type,required"`
	// Unique ID of the verification
	ID        string    `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx account currency used to describe monetary values, including billing
	// costs
	Currency       string `json:"currency"`
	DeliveryStatus string `json:"delivery_status"`
	// E.164 formatted phone number
	DestinationPhoneNumber string `json:"destination_phone_number"`
	// Currency amount per billing unit used to calculate the Telnyx billing costs
	Rate string `json:"rate"`
	// Billing unit used to calculate the Telnyx billing costs
	RateMeasuredIn     string    `json:"rate_measured_in"`
	UpdatedAt          time.Time `json:"updated_at" format:"date-time"`
	VerificationStatus string    `json:"verification_status"`
	VerifyChannelID    string    `json:"verify_channel_id" format:"uuid"`
	// Depending on the type of verification, the `verify_channel_id` points to one of
	// the following channel ids;
	//
	// ---
	//
	// | verify_channel_type | verify_channel_id |
	// | ------------------- | ----------------- |
	// | sms, psd2           | messaging_id      |
	// | call, flashcall     | call_control_id   |
	//
	// ---
	//
	// Any of "sms", "psd2", "call", "flashcall".
	VerifyChannelType string `json:"verify_channel_type"`
	VerifyProfileID   string `json:"verify_profile_id" format:"uuid"`
	// Currency amount for Verify Usage Fee
	VerifyUsageFee string `json:"verify_usage_fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType             respjson.Field
		ID                     respjson.Field
		CreatedAt              respjson.Field
		Currency               respjson.Field
		DeliveryStatus         respjson.Field
		DestinationPhoneNumber respjson.Field
		Rate                   respjson.Field
		RateMeasuredIn         respjson.Field
		UpdatedAt              respjson.Field
		VerificationStatus     respjson.Field
		VerifyChannelID        respjson.Field
		VerifyChannelType      respjson.Field
		VerifyProfileID        respjson.Field
		VerifyUsageFee         respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseVerify2Fa) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseVerify2Fa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseSimCardUsage struct {
	RecordType string `json:"record_type,required"`
	// Unique identifier for this SIM Card Usage
	ID string `json:"id" format:"uuid"`
	// Event close time
	ClosedAt time.Time `json:"closed_at" format:"date-time"`
	// Event creation time
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency string `json:"currency"`
	// Data cost
	DataCost float64 `json:"data_cost"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	DataRate string `json:"data_rate"`
	// Unit of wireless link consumption
	DataUnit string `json:"data_unit"`
	// Number of megabytes downloaded
	DownlinkData float64 `json:"downlink_data"`
	// International Mobile Subscriber Identity
	Imsi string `json:"imsi"`
	// Ip address that generated the event
	IPAddress string `json:"ip_address"`
	// Mobile country code
	Mcc string `json:"mcc"`
	// Mobile network code
	Mnc string `json:"mnc"`
	// Telephone number associated to SIM card
	PhoneNumber string `json:"phone_number"`
	// Unique identifier for SIM card
	SimCardID string `json:"sim_card_id"`
	// User-provided tags
	SimCardTags string `json:"sim_card_tags"`
	// Unique identifier for SIM group
	SimGroupID string `json:"sim_group_id"`
	// Sim group name for sim card
	SimGroupName string `json:"sim_group_name"`
	// Number of megabytes uploaded
	UplinkData float64 `json:"uplink_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType   respjson.Field
		ID           respjson.Field
		ClosedAt     respjson.Field
		CreatedAt    respjson.Field
		Currency     respjson.Field
		DataCost     respjson.Field
		DataRate     respjson.Field
		DataUnit     respjson.Field
		DownlinkData respjson.Field
		Imsi         respjson.Field
		IPAddress    respjson.Field
		Mcc          respjson.Field
		Mnc          respjson.Field
		PhoneNumber  respjson.Field
		SimCardID    respjson.Field
		SimCardTags  respjson.Field
		SimGroupID   respjson.Field
		SimGroupName respjson.Field
		UplinkData   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseSimCardUsage) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseSimCardUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListResponseMediaStorage struct {
	RecordType string `json:"record_type,required"`
	// Unique identifier for the Media Storage Event
	ID string `json:"id"`
	// Type of action performed against the Media Storage API
	ActionType string `json:"action_type"`
	// Asset id
	AssetID string `json:"asset_id"`
	// Currency amount for Telnyx billing cost
	Cost string `json:"cost"`
	// Event creation time
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Telnyx account currency used to describe monetary values, including billing cost
	Currency string `json:"currency"`
	// Link channel id
	LinkChannelID string `json:"link_channel_id"`
	// Link channel type
	LinkChannelType string `json:"link_channel_type"`
	// Organization owner id
	OrgID string `json:"org_id"`
	// Currency amount per billing unit used to calculate the Telnyx billing cost
	Rate string `json:"rate"`
	// Billing unit used to calculate the Telnyx billing cost
	RateMeasuredIn string `json:"rate_measured_in"`
	// Request status
	Status string `json:"status"`
	// User id
	UserID string `json:"user_id"`
	// Webhook id
	WebhookID string `json:"webhook_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType      respjson.Field
		ID              respjson.Field
		ActionType      respjson.Field
		AssetID         respjson.Field
		Cost            respjson.Field
		CreatedAt       respjson.Field
		Currency        respjson.Field
		LinkChannelID   respjson.Field
		LinkChannelType respjson.Field
		OrgID           respjson.Field
		Rate            respjson.Field
		RateMeasuredIn  respjson.Field
		Status          respjson.Field
		UserID          respjson.Field
		WebhookID       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DetailRecordListResponseMediaStorage) RawJSON() string { return r.JSON.raw }
func (r *DetailRecordListResponseMediaStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DetailRecordListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter records on a given record attribute and value. <br/>Example:
	// filter[status]=delivered. <br/>Required: filter[record_type] must be specified.
	Filter DetailRecordListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. <br/>Example: sort=-created_at
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DetailRecordListParams]'s query parameters as `url.Values`.
func (r DetailRecordListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter records on a given record attribute and value. <br/>Example:
// filter[status]=delivered. <br/>Required: filter[record_type] must be specified.
//
// The property RecordType is required.
type DetailRecordListParamsFilter struct {
	// Filter by the given record type.
	//
	// Any of "ai-voice-assistant", "amd", "call-control", "conference",
	// "conference-participant", "embedding", "fax", "inference",
	// "inference-speech-to-text", "media_storage", "media-streaming", "messaging",
	// "noise-suppression", "recording", "sip-trunking", "siprec-client", "stt", "tts",
	// "verify", "webrtc", "wireless".
	RecordType string `query:"record_type,omitzero,required" json:"-"`
	// Filter by the given user-friendly date range. You can specify one of the
	// following enum values, or a dynamic one using this format: last_N_days.
	//
	// Any of "yesterday", "today", "tomorrow", "last_week", "this_week", "next_week",
	// "last_month", "this_month", "next_month".
	DateRange   string         `query:"date_range,omitzero" json:"-"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

// URLQuery serializes [DetailRecordListParamsFilter]'s query parameters as
// `url.Values`.
func (r DetailRecordListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
