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

// MobilePhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMobilePhoneNumberService] method instead.
type MobilePhoneNumberService struct {
	Options   []option.RequestOption
	Messaging MobilePhoneNumberMessagingService
}

// NewMobilePhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMobilePhoneNumberService(opts ...option.RequestOption) (r MobilePhoneNumberService) {
	r = MobilePhoneNumberService{}
	r.Options = opts
	r.Messaging = NewMobilePhoneNumberMessagingService(opts...)
	return
}

// Retrieve a Mobile Phone Number
func (r *MobilePhoneNumberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MobilePhoneNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/mobile_phone_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Mobile Phone Number
func (r *MobilePhoneNumberService) Update(ctx context.Context, id string, body MobilePhoneNumberUpdateParams, opts ...option.RequestOption) (res *MobilePhoneNumberUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/mobile_phone_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List Mobile Phone Numbers
func (r *MobilePhoneNumberService) List(ctx context.Context, query MobilePhoneNumberListParams, opts ...option.RequestOption) (res *MobilePhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/mobile_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MobilePhoneNumberGetResponse struct {
	Data MobilePhoneNumberGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseData struct {
	// Identifies the resource.
	ID             string                                         `json:"id"`
	CallForwarding MobilePhoneNumberGetResponseDataCallForwarding `json:"call_forwarding"`
	CallRecording  MobilePhoneNumberGetResponseDataCallRecording  `json:"call_recording"`
	// Indicates if caller ID name is enabled.
	CallerIDNameEnabled bool                                        `json:"caller_id_name_enabled"`
	CnamListing         MobilePhoneNumberGetResponseDataCnamListing `json:"cnam_listing"`
	// The ID of the connection associated with this number.
	ConnectionID string `json:"connection_id,nullable" format:"int64"`
	// The name of the connection.
	ConnectionName string `json:"connection_name,nullable"`
	// The type of the connection.
	ConnectionType string `json:"connection_type,nullable"`
	// The ISO 3166-1 alpha-2 country code of the number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string.
	CustomerReference string                                  `json:"customer_reference,nullable"`
	Inbound           MobilePhoneNumberGetResponseDataInbound `json:"inbound"`
	// The inbound call screening setting.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening string `json:"inbound_call_screening,nullable"`
	// Indicates if mobile voice is enabled.
	MobileVoiceEnabled bool `json:"mobile_voice_enabled"`
	// The noise suppression setting.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression string                                   `json:"noise_suppression"`
	Outbound         MobilePhoneNumberGetResponseDataOutbound `json:"outbound"`
	// The +E.164-formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The ID of the SIM card associated with this number.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// The status of the phone number.
	Status string `json:"status"`
	// A list of tags associated with the number.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CallForwarding       respjson.Field
		CallRecording        respjson.Field
		CallerIDNameEnabled  respjson.Field
		CnamListing          respjson.Field
		ConnectionID         respjson.Field
		ConnectionName       respjson.Field
		ConnectionType       respjson.Field
		CountryISOAlpha2     respjson.Field
		CreatedAt            respjson.Field
		CustomerReference    respjson.Field
		Inbound              respjson.Field
		InboundCallScreening respjson.Field
		MobileVoiceEnabled   respjson.Field
		NoiseSuppression     respjson.Field
		Outbound             respjson.Field
		PhoneNumber          respjson.Field
		RecordType           respjson.Field
		SimCardID            respjson.Field
		Status               respjson.Field
		Tags                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseDataCallForwarding struct {
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Any of "always", "on-failure".
	ForwardingType string `json:"forwarding_type,nullable"`
	ForwardsTo     string `json:"forwards_to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallForwardingEnabled respjson.Field
		ForwardingType        respjson.Field
		ForwardsTo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseDataCallForwarding) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseDataCallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseDataCallRecording struct {
	// Any of "single", "dual".
	InboundCallRecordingChannels string `json:"inbound_call_recording_channels"`
	InboundCallRecordingEnabled  bool   `json:"inbound_call_recording_enabled"`
	// Any of "wav", "mp3".
	InboundCallRecordingFormat string `json:"inbound_call_recording_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundCallRecordingChannels respjson.Field
		InboundCallRecordingEnabled  respjson.Field
		InboundCallRecordingFormat   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseDataCallRecording) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseDataCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseDataCnamListing struct {
	CnamListingDetails string `json:"cnam_listing_details,nullable"`
	CnamListingEnabled bool   `json:"cnam_listing_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnamListingDetails respjson.Field
		CnamListingEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseDataCnamListing) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseDataCnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseDataInbound struct {
	// The ID of the app that will intercept inbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept inbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberGetResponseDataOutbound struct {
	// The ID of the app that will intercept outbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept outbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberGetResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberGetResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponse struct {
	Data MobilePhoneNumberUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseData struct {
	// Identifies the resource.
	ID             string                                            `json:"id"`
	CallForwarding MobilePhoneNumberUpdateResponseDataCallForwarding `json:"call_forwarding"`
	CallRecording  MobilePhoneNumberUpdateResponseDataCallRecording  `json:"call_recording"`
	// Indicates if caller ID name is enabled.
	CallerIDNameEnabled bool                                           `json:"caller_id_name_enabled"`
	CnamListing         MobilePhoneNumberUpdateResponseDataCnamListing `json:"cnam_listing"`
	// The ID of the connection associated with this number.
	ConnectionID string `json:"connection_id,nullable" format:"int64"`
	// The name of the connection.
	ConnectionName string `json:"connection_name,nullable"`
	// The type of the connection.
	ConnectionType string `json:"connection_type,nullable"`
	// The ISO 3166-1 alpha-2 country code of the number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string.
	CustomerReference string                                     `json:"customer_reference,nullable"`
	Inbound           MobilePhoneNumberUpdateResponseDataInbound `json:"inbound"`
	// The inbound call screening setting.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening string `json:"inbound_call_screening,nullable"`
	// Indicates if mobile voice is enabled.
	MobileVoiceEnabled bool `json:"mobile_voice_enabled"`
	// The noise suppression setting.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression string                                      `json:"noise_suppression"`
	Outbound         MobilePhoneNumberUpdateResponseDataOutbound `json:"outbound"`
	// The +E.164-formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The ID of the SIM card associated with this number.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// The status of the phone number.
	Status string `json:"status"`
	// A list of tags associated with the number.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CallForwarding       respjson.Field
		CallRecording        respjson.Field
		CallerIDNameEnabled  respjson.Field
		CnamListing          respjson.Field
		ConnectionID         respjson.Field
		ConnectionName       respjson.Field
		ConnectionType       respjson.Field
		CountryISOAlpha2     respjson.Field
		CreatedAt            respjson.Field
		CustomerReference    respjson.Field
		Inbound              respjson.Field
		InboundCallScreening respjson.Field
		MobileVoiceEnabled   respjson.Field
		NoiseSuppression     respjson.Field
		Outbound             respjson.Field
		PhoneNumber          respjson.Field
		RecordType           respjson.Field
		SimCardID            respjson.Field
		Status               respjson.Field
		Tags                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseDataCallForwarding struct {
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Any of "always", "on-failure".
	ForwardingType string `json:"forwarding_type,nullable"`
	ForwardsTo     string `json:"forwards_to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallForwardingEnabled respjson.Field
		ForwardingType        respjson.Field
		ForwardsTo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseDataCallForwarding) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseDataCallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseDataCallRecording struct {
	// Any of "single", "dual".
	InboundCallRecordingChannels string `json:"inbound_call_recording_channels"`
	InboundCallRecordingEnabled  bool   `json:"inbound_call_recording_enabled"`
	// Any of "wav", "mp3".
	InboundCallRecordingFormat string `json:"inbound_call_recording_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundCallRecordingChannels respjson.Field
		InboundCallRecordingEnabled  respjson.Field
		InboundCallRecordingFormat   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseDataCallRecording) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseDataCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseDataCnamListing struct {
	CnamListingDetails string `json:"cnam_listing_details,nullable"`
	CnamListingEnabled bool   `json:"cnam_listing_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnamListingDetails respjson.Field
		CnamListingEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseDataCnamListing) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseDataCnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseDataInbound struct {
	// The ID of the app that will intercept inbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept inbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateResponseDataOutbound struct {
	// The ID of the app that will intercept outbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept outbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberUpdateResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberUpdateResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponse struct {
	Data []MobilePhoneNumberListResponseData `json:"data"`
	Meta MobilePhoneNumberListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseData struct {
	// Identifies the resource.
	ID             string                                          `json:"id"`
	CallForwarding MobilePhoneNumberListResponseDataCallForwarding `json:"call_forwarding"`
	CallRecording  MobilePhoneNumberListResponseDataCallRecording  `json:"call_recording"`
	// Indicates if caller ID name is enabled.
	CallerIDNameEnabled bool                                         `json:"caller_id_name_enabled"`
	CnamListing         MobilePhoneNumberListResponseDataCnamListing `json:"cnam_listing"`
	// The ID of the connection associated with this number.
	ConnectionID string `json:"connection_id,nullable" format:"int64"`
	// The name of the connection.
	ConnectionName string `json:"connection_name,nullable"`
	// The type of the connection.
	ConnectionType string `json:"connection_type,nullable"`
	// The ISO 3166-1 alpha-2 country code of the number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer reference string.
	CustomerReference string                                   `json:"customer_reference,nullable"`
	Inbound           MobilePhoneNumberListResponseDataInbound `json:"inbound"`
	// The inbound call screening setting.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening string `json:"inbound_call_screening,nullable"`
	// Indicates if mobile voice is enabled.
	MobileVoiceEnabled bool `json:"mobile_voice_enabled"`
	// The noise suppression setting.
	//
	// Any of "inbound", "outbound", "both", "disabled".
	NoiseSuppression string                                    `json:"noise_suppression"`
	Outbound         MobilePhoneNumberListResponseDataOutbound `json:"outbound"`
	// The +E.164-formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The ID of the SIM card associated with this number.
	SimCardID string `json:"sim_card_id" format:"uuid"`
	// The status of the phone number.
	Status string `json:"status"`
	// A list of tags associated with the number.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CallForwarding       respjson.Field
		CallRecording        respjson.Field
		CallerIDNameEnabled  respjson.Field
		CnamListing          respjson.Field
		ConnectionID         respjson.Field
		ConnectionName       respjson.Field
		ConnectionType       respjson.Field
		CountryISOAlpha2     respjson.Field
		CreatedAt            respjson.Field
		CustomerReference    respjson.Field
		Inbound              respjson.Field
		InboundCallScreening respjson.Field
		MobileVoiceEnabled   respjson.Field
		NoiseSuppression     respjson.Field
		Outbound             respjson.Field
		PhoneNumber          respjson.Field
		RecordType           respjson.Field
		SimCardID            respjson.Field
		Status               respjson.Field
		Tags                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseDataCallForwarding struct {
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Any of "always", "on-failure".
	ForwardingType string `json:"forwarding_type,nullable"`
	ForwardsTo     string `json:"forwards_to,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallForwardingEnabled respjson.Field
		ForwardingType        respjson.Field
		ForwardsTo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseDataCallForwarding) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseDataCallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseDataCallRecording struct {
	// Any of "single", "dual".
	InboundCallRecordingChannels string `json:"inbound_call_recording_channels"`
	InboundCallRecordingEnabled  bool   `json:"inbound_call_recording_enabled"`
	// Any of "wav", "mp3".
	InboundCallRecordingFormat string `json:"inbound_call_recording_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundCallRecordingChannels respjson.Field
		InboundCallRecordingEnabled  respjson.Field
		InboundCallRecordingFormat   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseDataCallRecording) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseDataCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseDataCnamListing struct {
	CnamListingDetails string `json:"cnam_listing_details,nullable"`
	CnamListingEnabled bool   `json:"cnam_listing_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnamListingDetails respjson.Field
		CnamListingEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseDataCnamListing) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseDataCnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseDataInbound struct {
	// The ID of the app that will intercept inbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept inbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseDataInbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseDataInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseDataOutbound struct {
	// The ID of the app that will intercept outbound calls.
	InterceptionAppID string `json:"interception_app_id,nullable" format:"int64"`
	// The name of the app that will intercept outbound calls.
	InterceptionAppName string `json:"interception_app_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InterceptionAppID   respjson.Field
		InterceptionAppName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberListResponseDataOutbound) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseDataOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
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
func (r MobilePhoneNumberListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateParams struct {
	ConnectionID        param.Opt[string]                           `json:"connection_id,omitzero"`
	CustomerReference   param.Opt[string]                           `json:"customer_reference,omitzero"`
	CallerIDNameEnabled param.Opt[bool]                             `json:"caller_id_name_enabled,omitzero"`
	NoiseSuppression    param.Opt[bool]                             `json:"noise_suppression,omitzero"`
	CallForwarding      MobilePhoneNumberUpdateParamsCallForwarding `json:"call_forwarding,omitzero"`
	CallRecording       MobilePhoneNumberUpdateParamsCallRecording  `json:"call_recording,omitzero"`
	CnamListing         MobilePhoneNumberUpdateParamsCnamListing    `json:"cnam_listing,omitzero"`
	Inbound             MobilePhoneNumberUpdateParamsInbound        `json:"inbound,omitzero"`
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening MobilePhoneNumberUpdateParamsInboundCallScreening `json:"inbound_call_screening,omitzero"`
	Outbound             MobilePhoneNumberUpdateParamsOutbound             `json:"outbound,omitzero"`
	Tags                 []string                                          `json:"tags,omitzero"`
	paramObj
}

func (r MobilePhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateParamsCallForwarding struct {
	ForwardsTo            param.Opt[string] `json:"forwards_to,omitzero"`
	CallForwardingEnabled param.Opt[bool]   `json:"call_forwarding_enabled,omitzero"`
	// Any of "always", "on-failure".
	ForwardingType string `json:"forwarding_type,omitzero"`
	paramObj
}

func (r MobilePhoneNumberUpdateParamsCallForwarding) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParamsCallForwarding
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParamsCallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MobilePhoneNumberUpdateParamsCallForwarding](
		"forwarding_type", "always", "on-failure",
	)
}

type MobilePhoneNumberUpdateParamsCallRecording struct {
	InboundCallRecordingEnabled param.Opt[bool] `json:"inbound_call_recording_enabled,omitzero"`
	// Any of "single", "dual".
	InboundCallRecordingChannels string `json:"inbound_call_recording_channels,omitzero"`
	// Any of "wav", "mp3".
	InboundCallRecordingFormat string `json:"inbound_call_recording_format,omitzero"`
	paramObj
}

func (r MobilePhoneNumberUpdateParamsCallRecording) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParamsCallRecording
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParamsCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MobilePhoneNumberUpdateParamsCallRecording](
		"inbound_call_recording_channels", "single", "dual",
	)
	apijson.RegisterFieldValidator[MobilePhoneNumberUpdateParamsCallRecording](
		"inbound_call_recording_format", "wav", "mp3",
	)
}

type MobilePhoneNumberUpdateParamsCnamListing struct {
	CnamListingDetails param.Opt[string] `json:"cnam_listing_details,omitzero"`
	CnamListingEnabled param.Opt[bool]   `json:"cnam_listing_enabled,omitzero"`
	paramObj
}

func (r MobilePhoneNumberUpdateParamsCnamListing) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParamsCnamListing
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParamsCnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateParamsInbound struct {
	// The ID of the CallControl or TeXML Application that will intercept inbound
	// calls.
	InterceptionAppID param.Opt[string] `json:"interception_app_id,omitzero" format:"int64"`
	paramObj
}

func (r MobilePhoneNumberUpdateParamsInbound) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParamsInbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParamsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberUpdateParamsInboundCallScreening string

const (
	MobilePhoneNumberUpdateParamsInboundCallScreeningDisabled    MobilePhoneNumberUpdateParamsInboundCallScreening = "disabled"
	MobilePhoneNumberUpdateParamsInboundCallScreeningRejectCalls MobilePhoneNumberUpdateParamsInboundCallScreening = "reject_calls"
	MobilePhoneNumberUpdateParamsInboundCallScreeningFlagCalls   MobilePhoneNumberUpdateParamsInboundCallScreening = "flag_calls"
)

type MobilePhoneNumberUpdateParamsOutbound struct {
	// The ID of the CallControl or TeXML Application that will intercept outbound
	// calls.
	InterceptionAppID param.Opt[string] `json:"interception_app_id,omitzero" format:"int64"`
	paramObj
}

func (r MobilePhoneNumberUpdateParamsOutbound) MarshalJSON() (data []byte, err error) {
	type shadow MobilePhoneNumberUpdateParamsOutbound
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePhoneNumberUpdateParamsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberListParams struct {
	// The page number to load
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobilePhoneNumberListParams]'s query parameters as
// `url.Values`.
func (r MobilePhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
