// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlAccountCallService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallService] method instead.
type TexmlAccountCallService struct {
	Options        []option.RequestOption
	RecordingsJson TexmlAccountCallRecordingsJsonService
	Recordings     TexmlAccountCallRecordingService
	Siprec         TexmlAccountCallSiprecService
	Streams        TexmlAccountCallStreamService
}

// NewTexmlAccountCallService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlAccountCallService(opts ...option.RequestOption) (r TexmlAccountCallService) {
	r = TexmlAccountCallService{}
	r.Options = opts
	r.RecordingsJson = NewTexmlAccountCallRecordingsJsonService(opts...)
	r.Recordings = NewTexmlAccountCallRecordingService(opts...)
	r.Siprec = NewTexmlAccountCallSiprecService(opts...)
	r.Streams = NewTexmlAccountCallStreamService(opts...)
	return
}

// Returns an individual call identified by its CallSid. This endpoint is
// eventually consistent.
func (r *TexmlAccountCallService) Get(ctx context.Context, callSid string, query TexmlAccountCallGetParams, opts ...option.RequestOption) (res *TexmlAccountCallGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s", query.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update TeXML call. Please note that the keys present in the payload MUST BE
// formatted in CamelCase as specified in the example.
func (r *TexmlAccountCallService) Update(ctx context.Context, callSid string, params TexmlAccountCallUpdateParams, opts ...option.RequestOption) (res *TexmlAccountCallUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s", params.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Initiate an outbound TeXML call. Telnyx will request TeXML from the XML Request
// URL configured for the connection in the Mission Control Portal.
func (r *TexmlAccountCallService) Calls(ctx context.Context, accountSid string, body TexmlAccountCallCallsParams, opts ...option.RequestOption) (res *TexmlAccountCallCallsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns multiple call resouces for an account. This endpoint is eventually
// consistent.
func (r *TexmlAccountCallService) GetCalls(ctx context.Context, accountSid string, query TexmlAccountCallGetCallsParams, opts ...option.RequestOption) (res *TexmlAccountCallGetCallsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Starts siprec session with specified parameters for call idientified by
// call_sid.
func (r *TexmlAccountCallService) SiprecJson(ctx context.Context, callSid string, params TexmlAccountCallSiprecJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallSiprecJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Siprec.json", params.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Starts streaming media from a call to a specific WebSocket address.
func (r *TexmlAccountCallService) StreamsJson(ctx context.Context, callSid string, params TexmlAccountCallStreamsJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallStreamsJsonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Streams.json", params.AccountSid, callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type UpdateCallParam struct {
	// A failover URL for which Telnyx will retrieve the TeXML call instructions if the
	// Url is not responding.
	FallbackURL param.Opt[string] `json:"FallbackUrl,omitzero"`
	// The value to set the call status to. Setting the status to completed ends the
	// call.
	Status param.Opt[string] `json:"Status,omitzero"`
	// URL destination for Telnyx to send status callback events to for the call.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// TeXML to replace the current one with.
	Texml param.Opt[string] `json:"Texml,omitzero"`
	// The URL where TeXML will make a request to retrieve a new set of TeXML
	// instructions to continue the call flow.
	URL param.Opt[string] `json:"Url,omitzero"`
	// HTTP request type used for `FallbackUrl`.
	//
	// Any of "GET", "POST".
	FallbackMethod UpdateCallFallbackMethod `json:"FallbackMethod,omitzero"`
	// HTTP request type used for `Url`.
	//
	// Any of "GET", "POST".
	Method UpdateCallMethod `json:"Method,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod UpdateCallStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	paramObj
}

func (r UpdateCallParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `FallbackUrl`.
type UpdateCallFallbackMethod string

const (
	UpdateCallFallbackMethodGet  UpdateCallFallbackMethod = "GET"
	UpdateCallFallbackMethodPost UpdateCallFallbackMethod = "POST"
)

// HTTP request type used for `Url`.
type UpdateCallMethod string

const (
	UpdateCallMethodGet  UpdateCallMethod = "GET"
	UpdateCallMethodPost UpdateCallMethod = "POST"
)

// HTTP request type used for `StatusCallback`.
type UpdateCallStatusCallbackMethod string

const (
	UpdateCallStatusCallbackMethodGet  UpdateCallStatusCallbackMethod = "GET"
	UpdateCallStatusCallbackMethodPost UpdateCallStatusCallbackMethod = "POST"
)

type TexmlAccountCallGetResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The value of the answering machine detection result, if this feature was enabled
	// for the call.
	//
	// Any of "human", "machine", "not_sure".
	AnsweredBy TexmlAccountCallGetResponseAnsweredBy `json:"answered_by"`
	// Caller ID, if present.
	CallerName string `json:"caller_name"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The direction of this call.
	//
	// Any of "inbound", "outbound".
	Direction TexmlAccountCallGetResponseDirection `json:"direction"`
	// The duration of this call, given in seconds.
	Duration string `json:"duration"`
	// The end time of this call.
	EndTime string `json:"end_time"`
	// The phone number or SIP address that made this call.
	From string `json:"from"`
	// The from number formatted for display.
	FromFormatted string `json:"from_formatted"`
	// The price of this call, the currency is specified in the price_unit field. Only
	// populated when the call cost feature is enabled for the account.
	Price string `json:"price"`
	// The unit in which the price is given.
	PriceUnit string `json:"price_unit"`
	// The identifier of this call.
	Sid string `json:"sid"`
	// The start time of this call.
	StartTime string `json:"start_time"`
	// The status of this call.
	//
	// Any of "ringing", "in-progress", "canceled", "completed", "failed", "busy",
	// "no-answer".
	Status TexmlAccountCallGetResponseStatus `json:"status"`
	// The phone number or SIP address that received this call.
	To string `json:"to"`
	// The to number formatted for display.
	ToFormatted string `json:"to_formatted"`
	// The relative URI for this call.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid    respjson.Field
		AnsweredBy    respjson.Field
		CallerName    respjson.Field
		DateCreated   respjson.Field
		DateUpdated   respjson.Field
		Direction     respjson.Field
		Duration      respjson.Field
		EndTime       respjson.Field
		From          respjson.Field
		FromFormatted respjson.Field
		Price         respjson.Field
		PriceUnit     respjson.Field
		Sid           respjson.Field
		StartTime     respjson.Field
		Status        respjson.Field
		To            respjson.Field
		ToFormatted   respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The value of the answering machine detection result, if this feature was enabled
// for the call.
type TexmlAccountCallGetResponseAnsweredBy string

const (
	TexmlAccountCallGetResponseAnsweredByHuman   TexmlAccountCallGetResponseAnsweredBy = "human"
	TexmlAccountCallGetResponseAnsweredByMachine TexmlAccountCallGetResponseAnsweredBy = "machine"
	TexmlAccountCallGetResponseAnsweredByNotSure TexmlAccountCallGetResponseAnsweredBy = "not_sure"
)

// The direction of this call.
type TexmlAccountCallGetResponseDirection string

const (
	TexmlAccountCallGetResponseDirectionInbound  TexmlAccountCallGetResponseDirection = "inbound"
	TexmlAccountCallGetResponseDirectionOutbound TexmlAccountCallGetResponseDirection = "outbound"
)

// The status of this call.
type TexmlAccountCallGetResponseStatus string

const (
	TexmlAccountCallGetResponseStatusRinging    TexmlAccountCallGetResponseStatus = "ringing"
	TexmlAccountCallGetResponseStatusInProgress TexmlAccountCallGetResponseStatus = "in-progress"
	TexmlAccountCallGetResponseStatusCanceled   TexmlAccountCallGetResponseStatus = "canceled"
	TexmlAccountCallGetResponseStatusCompleted  TexmlAccountCallGetResponseStatus = "completed"
	TexmlAccountCallGetResponseStatusFailed     TexmlAccountCallGetResponseStatus = "failed"
	TexmlAccountCallGetResponseStatusBusy       TexmlAccountCallGetResponseStatus = "busy"
	TexmlAccountCallGetResponseStatusNoAnswer   TexmlAccountCallGetResponseStatus = "no-answer"
)

type TexmlAccountCallUpdateResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The value of the answering machine detection result, if this feature was enabled
	// for the call.
	//
	// Any of "human", "machine", "not_sure".
	AnsweredBy TexmlAccountCallUpdateResponseAnsweredBy `json:"answered_by"`
	// Caller ID, if present.
	CallerName string `json:"caller_name"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The direction of this call.
	//
	// Any of "inbound", "outbound".
	Direction TexmlAccountCallUpdateResponseDirection `json:"direction"`
	// The duration of this call, given in seconds.
	Duration string `json:"duration"`
	// The end time of this call.
	EndTime string `json:"end_time"`
	// The phone number or SIP address that made this call.
	From string `json:"from"`
	// The from number formatted for display.
	FromFormatted string `json:"from_formatted"`
	// The price of this call, the currency is specified in the price_unit field. Only
	// populated when the call cost feature is enabled for the account.
	Price string `json:"price"`
	// The unit in which the price is given.
	PriceUnit string `json:"price_unit"`
	// The identifier of this call.
	Sid string `json:"sid"`
	// The start time of this call.
	StartTime string `json:"start_time"`
	// The status of this call.
	//
	// Any of "ringing", "in-progress", "canceled", "completed", "failed", "busy",
	// "no-answer".
	Status TexmlAccountCallUpdateResponseStatus `json:"status"`
	// The phone number or SIP address that received this call.
	To string `json:"to"`
	// The to number formatted for display.
	ToFormatted string `json:"to_formatted"`
	// The relative URI for this call.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid    respjson.Field
		AnsweredBy    respjson.Field
		CallerName    respjson.Field
		DateCreated   respjson.Field
		DateUpdated   respjson.Field
		Direction     respjson.Field
		Duration      respjson.Field
		EndTime       respjson.Field
		From          respjson.Field
		FromFormatted respjson.Field
		Price         respjson.Field
		PriceUnit     respjson.Field
		Sid           respjson.Field
		StartTime     respjson.Field
		Status        respjson.Field
		To            respjson.Field
		ToFormatted   respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The value of the answering machine detection result, if this feature was enabled
// for the call.
type TexmlAccountCallUpdateResponseAnsweredBy string

const (
	TexmlAccountCallUpdateResponseAnsweredByHuman   TexmlAccountCallUpdateResponseAnsweredBy = "human"
	TexmlAccountCallUpdateResponseAnsweredByMachine TexmlAccountCallUpdateResponseAnsweredBy = "machine"
	TexmlAccountCallUpdateResponseAnsweredByNotSure TexmlAccountCallUpdateResponseAnsweredBy = "not_sure"
)

// The direction of this call.
type TexmlAccountCallUpdateResponseDirection string

const (
	TexmlAccountCallUpdateResponseDirectionInbound  TexmlAccountCallUpdateResponseDirection = "inbound"
	TexmlAccountCallUpdateResponseDirectionOutbound TexmlAccountCallUpdateResponseDirection = "outbound"
)

// The status of this call.
type TexmlAccountCallUpdateResponseStatus string

const (
	TexmlAccountCallUpdateResponseStatusRinging    TexmlAccountCallUpdateResponseStatus = "ringing"
	TexmlAccountCallUpdateResponseStatusInProgress TexmlAccountCallUpdateResponseStatus = "in-progress"
	TexmlAccountCallUpdateResponseStatusCanceled   TexmlAccountCallUpdateResponseStatus = "canceled"
	TexmlAccountCallUpdateResponseStatusCompleted  TexmlAccountCallUpdateResponseStatus = "completed"
	TexmlAccountCallUpdateResponseStatusFailed     TexmlAccountCallUpdateResponseStatus = "failed"
	TexmlAccountCallUpdateResponseStatusBusy       TexmlAccountCallUpdateResponseStatus = "busy"
	TexmlAccountCallUpdateResponseStatusNoAnswer   TexmlAccountCallUpdateResponseStatus = "no-answer"
)

type TexmlAccountCallCallsResponse struct {
	From   string `json:"from"`
	Status string `json:"status"`
	To     string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallCallsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallCallsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallGetCallsResponse struct {
	Calls []TexmlAccountCallGetCallsResponseCall `json:"calls"`
	// The number of the last element on the page, zero-indexed.
	End int64 `json:"end"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Calls.json?Page=0&PageSize=1
	FirstPageUri string `json:"first_page_uri"`
	// /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Calls.json?Page=1&PageSize=1&PageToken=MTY4AjgyNDkwNzIxMQ
	NextPageUri string `json:"next_page_uri"`
	// Current page number, zero-indexed.
	Page int64 `json:"page"`
	// The number of items on the page
	PageSize int64 `json:"page_size"`
	// The number of the first element on the page, zero-indexed.
	Start int64 `json:"start"`
	// The URI of the current page.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calls        respjson.Field
		End          respjson.Field
		FirstPageUri respjson.Field
		NextPageUri  respjson.Field
		Page         respjson.Field
		PageSize     respjson.Field
		Start        respjson.Field
		Uri          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallGetCallsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallGetCallsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallGetCallsResponseCall struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The value of the answering machine detection result, if this feature was enabled
	// for the call.
	//
	// Any of "human", "machine", "not_sure".
	AnsweredBy string `json:"answered_by"`
	// Caller ID, if present.
	CallerName string `json:"caller_name"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The direction of this call.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// The duration of this call, given in seconds.
	Duration string `json:"duration"`
	// The end time of this call.
	EndTime string `json:"end_time"`
	// The phone number or SIP address that made this call.
	From string `json:"from"`
	// The from number formatted for display.
	FromFormatted string `json:"from_formatted"`
	// The price of this call, the currency is specified in the price_unit field. Only
	// populated when the call cost feature is enabled for the account.
	Price string `json:"price"`
	// The unit in which the price is given.
	PriceUnit string `json:"price_unit"`
	// The identifier of this call.
	Sid string `json:"sid"`
	// The start time of this call.
	StartTime string `json:"start_time"`
	// The status of this call.
	//
	// Any of "ringing", "in-progress", "canceled", "completed", "failed", "busy",
	// "no-answer".
	Status string `json:"status"`
	// The phone number or SIP address that received this call.
	To string `json:"to"`
	// The to number formatted for display.
	ToFormatted string `json:"to_formatted"`
	// The relative URI for this call.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid    respjson.Field
		AnsweredBy    respjson.Field
		CallerName    respjson.Field
		DateCreated   respjson.Field
		DateUpdated   respjson.Field
		Direction     respjson.Field
		Duration      respjson.Field
		EndTime       respjson.Field
		From          respjson.Field
		FromFormatted respjson.Field
		Price         respjson.Field
		PriceUnit     respjson.Field
		Sid           respjson.Field
		StartTime     respjson.Field
		Status        respjson.Field
		To            respjson.Field
		ToFormatted   respjson.Field
		Uri           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallGetCallsResponseCall) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallGetCallsResponseCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallSiprecJsonResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The id of the call the resource belongs to.
	CallSid string `json:"call_sid"`
	// The date and time the siprec session was created.
	DateCreated string `json:"date_created"`
	// The date and time the siprec session was last updated.
	DateUpdated string `json:"date_updated"`
	// The error code of the siprec session.
	ErrorCode string `json:"error_code"`
	// The SID of the siprec session.
	Sid string `json:"sid"`
	// The date and time the siprec session was started.
	StartTime string `json:"start_time"`
	// The status of the siprec session.
	//
	// Any of "in-progress", "stopped".
	Status TexmlAccountCallSiprecJsonResponseStatus `json:"status"`
	// The track used for the siprec session.
	//
	// Any of "both_tracks", "inbound_track", "outbound_track".
	Track TexmlAccountCallSiprecJsonResponseTrack `json:"track"`
	// The URI of the siprec session.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid  respjson.Field
		CallSid     respjson.Field
		DateCreated respjson.Field
		DateUpdated respjson.Field
		ErrorCode   respjson.Field
		Sid         respjson.Field
		StartTime   respjson.Field
		Status      respjson.Field
		Track       respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallSiprecJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallSiprecJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the siprec session.
type TexmlAccountCallSiprecJsonResponseStatus string

const (
	TexmlAccountCallSiprecJsonResponseStatusInProgress TexmlAccountCallSiprecJsonResponseStatus = "in-progress"
	TexmlAccountCallSiprecJsonResponseStatusStopped    TexmlAccountCallSiprecJsonResponseStatus = "stopped"
)

// The track used for the siprec session.
type TexmlAccountCallSiprecJsonResponseTrack string

const (
	TexmlAccountCallSiprecJsonResponseTrackBothTracks    TexmlAccountCallSiprecJsonResponseTrack = "both_tracks"
	TexmlAccountCallSiprecJsonResponseTrackInboundTrack  TexmlAccountCallSiprecJsonResponseTrack = "inbound_track"
	TexmlAccountCallSiprecJsonResponseTrackOutboundTrack TexmlAccountCallSiprecJsonResponseTrack = "outbound_track"
)

type TexmlAccountCallStreamsJsonResponse struct {
	AccountSid  string    `json:"account_sid"`
	CallSid     string    `json:"call_sid"`
	DateUpdated time.Time `json:"date_updated" format:"date-time"`
	// The user specified name of Stream.
	Name string `json:"name"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// The status of the streaming.
	//
	// Any of "in-progress".
	Status TexmlAccountCallStreamsJsonResponseStatus `json:"status"`
	// The relative URI for this streaming resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid  respjson.Field
		CallSid     respjson.Field
		DateUpdated respjson.Field
		Name        respjson.Field
		Sid         respjson.Field
		Status      respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallStreamsJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallStreamsJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the streaming.
type TexmlAccountCallStreamsJsonResponseStatus string

const (
	TexmlAccountCallStreamsJsonResponseStatusInProgress TexmlAccountCallStreamsJsonResponseStatus = "in-progress"
)

type TexmlAccountCallGetParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountCallUpdateParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	UpdateCall UpdateCallParam
	paramObj
}

func (r TexmlAccountCallUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCall)
}
func (r *TexmlAccountCallUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateCall)
}

type TexmlAccountCallCallsParams struct {
	// The ID of the TeXML Application.
	ApplicationSid string `json:"ApplicationSid,required"`
	// The phone number of the party that initiated the call. Phone numbers are
	// formatted with a `+` and country code.
	From string `json:"From,required"`
	// The phone number of the called party. Phone numbers are formatted with a `+` and
	// country code.
	To string `json:"To,required"`
	// Select whether to perform answering machine detection in the background. By
	// default execution is blocked until Answering Machine Detection is completed.
	AsyncAmd param.Opt[bool] `json:"AsyncAmd,omitzero"`
	// URL destination for Telnyx to send AMD callback events to for the call.
	AsyncAmdStatusCallback param.Opt[string] `json:"AsyncAmdStatusCallback,omitzero"`
	// To be used as the caller id name (SIP From Display Name) presented to the
	// destination (`To` number). The string should have a maximum of 128 characters,
	// containing only letters, numbers, spaces, and `-_~!.+` special characters. If
	// ommited, the display name will be the same as the number in the `From` field.
	CallerID param.Opt[string] `json:"CallerId,omitzero"`
	// Whether to cancel ongoing playback on `greeting ended` detection. Defaults to
	// `true`.
	CancelPlaybackOnDetectMessageEnd param.Opt[bool] `json:"CancelPlaybackOnDetectMessageEnd,omitzero"`
	// Whether to cancel ongoing playback on `machine` detection. Defaults to `true`.
	CancelPlaybackOnMachineDetection param.Opt[bool] `json:"CancelPlaybackOnMachineDetection,omitzero"`
	// A failover URL for which Telnyx will retrieve the TeXML call instructions if the
	// `Url` is not responding.
	FallbackURL param.Opt[string] `json:"FallbackUrl,omitzero"`
	// If initial silence duration is greater than this value, consider it a machine.
	// Ignored when `premium` detection is used.
	MachineDetectionSilenceTimeout param.Opt[int64] `json:"MachineDetectionSilenceTimeout,omitzero"`
	// Silence duration threshold after a greeting message or voice for it be
	// considered human. Ignored when `premium` detection is used.
	MachineDetectionSpeechEndThreshold param.Opt[int64] `json:"MachineDetectionSpeechEndThreshold,omitzero"`
	// Maximum threshold of a human greeting. If greeting longer than this value,
	// considered machine. Ignored when `premium` detection is used.
	MachineDetectionSpeechThreshold param.Opt[int64] `json:"MachineDetectionSpeechThreshold,omitzero"`
	// Maximum timeout threshold in milliseconds for overall detection.
	MachineDetectionTimeout param.Opt[int64] `json:"MachineDetectionTimeout,omitzero"`
	// The list of comma-separated codecs to be offered on a call.
	PreferredCodecs param.Opt[string] `json:"PreferredCodecs,omitzero"`
	// Whether to record the entire participant's call leg. Defaults to `false`.
	Record param.Opt[bool] `json:"Record,omitzero"`
	// The URL the recording callbacks will be sent to.
	RecordingStatusCallback param.Opt[string] `json:"RecordingStatusCallback,omitzero"`
	// The changes to the recording's state that should generate a call to
	// `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`.
	RecordingStatusCallbackEvent param.Opt[string] `json:"RecordingStatusCallbackEvent,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected. The timer only starts when the speech is detected. Please
	// note that the transcription is used to detect silence and the related charge
	// will be applied. The minimum value is 0. The default value is 0 (infinite)
	RecordingTimeout param.Opt[int64] `json:"RecordingTimeout,omitzero"`
	// Whether to send RecordingUrl in webhooks.
	SendRecordingURL param.Opt[bool] `json:"SendRecordingUrl,omitzero"`
	// The password to use for SIP authentication.
	SipAuthPassword param.Opt[string] `json:"SipAuthPassword,omitzero"`
	// The username to use for SIP authentication.
	SipAuthUsername param.Opt[string] `json:"SipAuthUsername,omitzero"`
	// URL destination for Telnyx to send status callback events to for the call.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// The URL from which Telnyx will retrieve the TeXML call instructions.
	URL param.Opt[string] `json:"Url,omitzero"`
	// HTTP request type used for `AsyncAmdStatusCallback`. The default value is
	// inherited from TeXML Application setting.
	//
	// Any of "GET", "POST".
	AsyncAmdStatusCallbackMethod TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethod `json:"AsyncAmdStatusCallbackMethod,omitzero"`
	// Custom HTTP headers to be sent with the call. Each header should be an object
	// with 'name' and 'value' properties.
	CustomHeaders []TexmlAccountCallCallsParamsCustomHeader `json:"CustomHeaders,omitzero"`
	// Allows you to chose between Premium and Standard detections.
	//
	// Any of "Premium", "Regular".
	DetectionMode TexmlAccountCallCallsParamsDetectionMode `json:"DetectionMode,omitzero"`
	// Enables Answering Machine Detection.
	//
	// Any of "Enable", "Disable", "DetectMessageEnd".
	MachineDetection TexmlAccountCallCallsParamsMachineDetection `json:"MachineDetection,omitzero"`
	// The number of channels in the final recording. Defaults to `mono`.
	//
	// Any of "mono", "dual".
	RecordingChannels TexmlAccountCallCallsParamsRecordingChannels `json:"RecordingChannels,omitzero"`
	// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	RecordingStatusCallbackMethod TexmlAccountCallCallsParamsRecordingStatusCallbackMethod `json:"RecordingStatusCallbackMethod,omitzero"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	RecordingTrack TexmlAccountCallCallsParamsRecordingTrack `json:"RecordingTrack,omitzero"`
	// The call events for which Telnyx should send a webhook. Multiple events can be
	// defined when separated by a space.
	//
	// Any of "initiated", "ringing", "answered", "completed".
	StatusCallbackEvent TexmlAccountCallCallsParamsStatusCallbackEvent `json:"StatusCallbackEvent,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlAccountCallCallsParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// Whether to trim any leading and trailing silence from the recording. Defaults to
	// `trim-silence`.
	//
	// Any of "trim-silence", "do-not-trim".
	Trim TexmlAccountCallCallsParamsTrim `json:"Trim,omitzero"`
	// HTTP request type used for `Url`. The default value is inherited from TeXML
	// Application setting.
	//
	// Any of "GET", "POST".
	URLMethod TexmlAccountCallCallsParamsURLMethod `json:"UrlMethod,omitzero"`
	paramObj
}

func (r TexmlAccountCallCallsParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallCallsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallCallsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `AsyncAmdStatusCallback`. The default value is
// inherited from TeXML Application setting.
type TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethod string

const (
	TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethodGet  TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethod = "GET"
	TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethodPost TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethod = "POST"
)

// The properties Name, Value are required.
type TexmlAccountCallCallsParamsCustomHeader struct {
	// The name of the custom header
	Name string `json:"name,required"`
	// The value of the custom header
	Value string `json:"value,required"`
	paramObj
}

func (r TexmlAccountCallCallsParamsCustomHeader) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallCallsParamsCustomHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallCallsParamsCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to chose between Premium and Standard detections.
type TexmlAccountCallCallsParamsDetectionMode string

const (
	TexmlAccountCallCallsParamsDetectionModePremium TexmlAccountCallCallsParamsDetectionMode = "Premium"
	TexmlAccountCallCallsParamsDetectionModeRegular TexmlAccountCallCallsParamsDetectionMode = "Regular"
)

// Enables Answering Machine Detection.
type TexmlAccountCallCallsParamsMachineDetection string

const (
	TexmlAccountCallCallsParamsMachineDetectionEnable           TexmlAccountCallCallsParamsMachineDetection = "Enable"
	TexmlAccountCallCallsParamsMachineDetectionDisable          TexmlAccountCallCallsParamsMachineDetection = "Disable"
	TexmlAccountCallCallsParamsMachineDetectionDetectMessageEnd TexmlAccountCallCallsParamsMachineDetection = "DetectMessageEnd"
)

// The number of channels in the final recording. Defaults to `mono`.
type TexmlAccountCallCallsParamsRecordingChannels string

const (
	TexmlAccountCallCallsParamsRecordingChannelsMono TexmlAccountCallCallsParamsRecordingChannels = "mono"
	TexmlAccountCallCallsParamsRecordingChannelsDual TexmlAccountCallCallsParamsRecordingChannels = "dual"
)

// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
type TexmlAccountCallCallsParamsRecordingStatusCallbackMethod string

const (
	TexmlAccountCallCallsParamsRecordingStatusCallbackMethodGet  TexmlAccountCallCallsParamsRecordingStatusCallbackMethod = "GET"
	TexmlAccountCallCallsParamsRecordingStatusCallbackMethodPost TexmlAccountCallCallsParamsRecordingStatusCallbackMethod = "POST"
)

// The audio track to record for the call. The default is `both`.
type TexmlAccountCallCallsParamsRecordingTrack string

const (
	TexmlAccountCallCallsParamsRecordingTrackInbound  TexmlAccountCallCallsParamsRecordingTrack = "inbound"
	TexmlAccountCallCallsParamsRecordingTrackOutbound TexmlAccountCallCallsParamsRecordingTrack = "outbound"
	TexmlAccountCallCallsParamsRecordingTrackBoth     TexmlAccountCallCallsParamsRecordingTrack = "both"
)

// The call events for which Telnyx should send a webhook. Multiple events can be
// defined when separated by a space.
type TexmlAccountCallCallsParamsStatusCallbackEvent string

const (
	TexmlAccountCallCallsParamsStatusCallbackEventInitiated TexmlAccountCallCallsParamsStatusCallbackEvent = "initiated"
	TexmlAccountCallCallsParamsStatusCallbackEventRinging   TexmlAccountCallCallsParamsStatusCallbackEvent = "ringing"
	TexmlAccountCallCallsParamsStatusCallbackEventAnswered  TexmlAccountCallCallsParamsStatusCallbackEvent = "answered"
	TexmlAccountCallCallsParamsStatusCallbackEventCompleted TexmlAccountCallCallsParamsStatusCallbackEvent = "completed"
)

// HTTP request type used for `StatusCallback`.
type TexmlAccountCallCallsParamsStatusCallbackMethod string

const (
	TexmlAccountCallCallsParamsStatusCallbackMethodGet  TexmlAccountCallCallsParamsStatusCallbackMethod = "GET"
	TexmlAccountCallCallsParamsStatusCallbackMethodPost TexmlAccountCallCallsParamsStatusCallbackMethod = "POST"
)

// Whether to trim any leading and trailing silence from the recording. Defaults to
// `trim-silence`.
type TexmlAccountCallCallsParamsTrim string

const (
	TexmlAccountCallCallsParamsTrimTrimSilence TexmlAccountCallCallsParamsTrim = "trim-silence"
	TexmlAccountCallCallsParamsTrimDoNotTrim   TexmlAccountCallCallsParamsTrim = "do-not-trim"
)

// HTTP request type used for `Url`. The default value is inherited from TeXML
// Application setting.
type TexmlAccountCallCallsParamsURLMethod string

const (
	TexmlAccountCallCallsParamsURLMethodGet  TexmlAccountCallCallsParamsURLMethod = "GET"
	TexmlAccountCallCallsParamsURLMethodPost TexmlAccountCallCallsParamsURLMethod = "POST"
)

type TexmlAccountCallGetCallsParams struct {
	// Filters calls by their end date. Expected format is YYYY-MM-DD
	EndTime param.Opt[string] `query:"EndTime,omitzero" json:"-"`
	// Filters calls by their end date (after). Expected format is YYYY-MM-DD
	EndTimeGt param.Opt[string] `query:"EndTime_gt,omitzero" json:"-"`
	// Filters calls by their end date (before). Expected format is YYYY-MM-DD
	EndTimeLt param.Opt[string] `query:"EndTime_lt,omitzero" json:"-"`
	// Filters calls by the from number.
	From param.Opt[string] `query:"From,omitzero" json:"-"`
	// The number of the page to be displayed, zero-indexed, should be used in
	// conjuction with PageToken.
	Page param.Opt[int64] `query:"Page,omitzero" json:"-"`
	// The number of records to be displayed on a page
	PageSize param.Opt[int64] `query:"PageSize,omitzero" json:"-"`
	// Used to request the next page of results.
	PageToken param.Opt[string] `query:"PageToken,omitzero" json:"-"`
	// Filters calls by their start date. Expected format is YYYY-MM-DD.
	StartTime param.Opt[string] `query:"StartTime,omitzero" json:"-"`
	// Filters calls by their start date (after). Expected format is YYYY-MM-DD
	StartTimeGt param.Opt[string] `query:"StartTime_gt,omitzero" json:"-"`
	// Filters calls by their start date (before). Expected format is YYYY-MM-DD
	StartTimeLt param.Opt[string] `query:"StartTime_lt,omitzero" json:"-"`
	// Filters calls by the to number.
	To param.Opt[string] `query:"To,omitzero" json:"-"`
	// Filters calls by status.
	//
	// Any of "canceled", "completed", "failed", "busy", "no-answer".
	Status TexmlAccountCallGetCallsParamsStatus `query:"Status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TexmlAccountCallGetCallsParams]'s query parameters as
// `url.Values`.
func (r TexmlAccountCallGetCallsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters calls by status.
type TexmlAccountCallGetCallsParamsStatus string

const (
	TexmlAccountCallGetCallsParamsStatusCanceled  TexmlAccountCallGetCallsParamsStatus = "canceled"
	TexmlAccountCallGetCallsParamsStatusCompleted TexmlAccountCallGetCallsParamsStatus = "completed"
	TexmlAccountCallGetCallsParamsStatusFailed    TexmlAccountCallGetCallsParamsStatus = "failed"
	TexmlAccountCallGetCallsParamsStatusBusy      TexmlAccountCallGetCallsParamsStatus = "busy"
	TexmlAccountCallGetCallsParamsStatusNoAnswer  TexmlAccountCallGetCallsParamsStatus = "no-answer"
)

type TexmlAccountCallSiprecJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	// The name of the connector to use for the SIPREC session.
	ConnectorName param.Opt[string] `json:"ConnectorName,omitzero"`
	// Name of the SIPREC session. May be used to stop the SIPREC session from TeXML
	// instruction.
	Name param.Opt[string] `json:"Name,omitzero"`
	// Sets `Session-Expires` header to the INVITE. A reinvite is sent every half the
	// value set. Usefull for session keep alive. Minimum value is 90, set to 0 to
	// disable.
	SessionTimeoutSecs param.Opt[int64] `json:"SessionTimeoutSecs,omitzero"`
	// URL destination for Telnyx to send status callback events to for the siprec
	// session.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// When set, custom parameters will be added as metadata
	// (recording.session.ExtensionParameters). Otherwise, they’ll be added to sip
	// headers.
	//
	// Any of true, false.
	IncludeMetadataCustomHeaders bool `json:"IncludeMetadataCustomHeaders,omitzero"`
	// Controls whether to encrypt media sent to your SRS using SRTP and TLS. When set
	// you need to configure SRS port in your connector to 5061.
	//
	// Any of true, false.
	Secure bool `json:"Secure,omitzero"`
	// Specifies SIP transport protocol.
	//
	// Any of "udp", "tcp", "tls".
	SipTransport TexmlAccountCallSiprecJsonParamsSipTransport `json:"SipTransport,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlAccountCallSiprecJsonParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// The track to be used for siprec session. Can be `both_tracks`, `inbound_track`
	// or `outbound_track`. Defaults to `both_tracks`.
	//
	// Any of "both_tracks", "inbound_track", "outbound_track".
	Track TexmlAccountCallSiprecJsonParamsTrack `json:"Track,omitzero"`
	paramObj
}

func (r TexmlAccountCallSiprecJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallSiprecJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallSiprecJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies SIP transport protocol.
type TexmlAccountCallSiprecJsonParamsSipTransport string

const (
	TexmlAccountCallSiprecJsonParamsSipTransportUdp TexmlAccountCallSiprecJsonParamsSipTransport = "udp"
	TexmlAccountCallSiprecJsonParamsSipTransportTcp TexmlAccountCallSiprecJsonParamsSipTransport = "tcp"
	TexmlAccountCallSiprecJsonParamsSipTransportTls TexmlAccountCallSiprecJsonParamsSipTransport = "tls"
)

// HTTP request type used for `StatusCallback`.
type TexmlAccountCallSiprecJsonParamsStatusCallbackMethod string

const (
	TexmlAccountCallSiprecJsonParamsStatusCallbackMethodGet  TexmlAccountCallSiprecJsonParamsStatusCallbackMethod = "GET"
	TexmlAccountCallSiprecJsonParamsStatusCallbackMethodPost TexmlAccountCallSiprecJsonParamsStatusCallbackMethod = "POST"
)

// The track to be used for siprec session. Can be `both_tracks`, `inbound_track`
// or `outbound_track`. Defaults to `both_tracks`.
type TexmlAccountCallSiprecJsonParamsTrack string

const (
	TexmlAccountCallSiprecJsonParamsTrackBothTracks    TexmlAccountCallSiprecJsonParamsTrack = "both_tracks"
	TexmlAccountCallSiprecJsonParamsTrackInboundTrack  TexmlAccountCallSiprecJsonParamsTrack = "inbound_track"
	TexmlAccountCallSiprecJsonParamsTrackOutboundTrack TexmlAccountCallSiprecJsonParamsTrack = "outbound_track"
)

type TexmlAccountCallStreamsJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	// The user specified name of Stream.
	Name param.Opt[string] `json:"Name,omitzero"`
	// Url where status callbacks will be sent.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero" format:"uri"`
	// The destination WebSocket address where the stream is going to be delivered.
	URL param.Opt[string] `json:"Url,omitzero"`
	// Indicates codec for bidirectional streaming RTP payloads. Used only with
	// stream_bidirectional_mode=rtp. Case sensitive.
	//
	// Any of "PCMU", "PCMA", "G722".
	BidirectionalCodec TexmlAccountCallStreamsJsonParamsBidirectionalCodec `json:"BidirectionalCodec,omitzero"`
	// Configures method of bidirectional streaming (mp3, rtp).
	//
	// Any of "mp3", "rtp".
	BidirectionalMode TexmlAccountCallStreamsJsonParamsBidirectionalMode `json:"BidirectionalMode,omitzero"`
	// HTTP method used to send status callbacks.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlAccountCallStreamsJsonParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// Tracks to be included in the stream
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	Track TexmlAccountCallStreamsJsonParamsTrack `json:"Track,omitzero"`
	paramObj
}

func (r TexmlAccountCallStreamsJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallStreamsJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallStreamsJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates codec for bidirectional streaming RTP payloads. Used only with
// stream_bidirectional_mode=rtp. Case sensitive.
type TexmlAccountCallStreamsJsonParamsBidirectionalCodec string

const (
	TexmlAccountCallStreamsJsonParamsBidirectionalCodecPcmu TexmlAccountCallStreamsJsonParamsBidirectionalCodec = "PCMU"
	TexmlAccountCallStreamsJsonParamsBidirectionalCodecPcma TexmlAccountCallStreamsJsonParamsBidirectionalCodec = "PCMA"
	TexmlAccountCallStreamsJsonParamsBidirectionalCodecG722 TexmlAccountCallStreamsJsonParamsBidirectionalCodec = "G722"
)

// Configures method of bidirectional streaming (mp3, rtp).
type TexmlAccountCallStreamsJsonParamsBidirectionalMode string

const (
	TexmlAccountCallStreamsJsonParamsBidirectionalModeMP3 TexmlAccountCallStreamsJsonParamsBidirectionalMode = "mp3"
	TexmlAccountCallStreamsJsonParamsBidirectionalModeRtp TexmlAccountCallStreamsJsonParamsBidirectionalMode = "rtp"
)

// HTTP method used to send status callbacks.
type TexmlAccountCallStreamsJsonParamsStatusCallbackMethod string

const (
	TexmlAccountCallStreamsJsonParamsStatusCallbackMethodGet  TexmlAccountCallStreamsJsonParamsStatusCallbackMethod = "GET"
	TexmlAccountCallStreamsJsonParamsStatusCallbackMethodPost TexmlAccountCallStreamsJsonParamsStatusCallbackMethod = "POST"
)

// Tracks to be included in the stream
type TexmlAccountCallStreamsJsonParamsTrack string

const (
	TexmlAccountCallStreamsJsonParamsTrackInboundTrack  TexmlAccountCallStreamsJsonParamsTrack = "inbound_track"
	TexmlAccountCallStreamsJsonParamsTrackOutboundTrack TexmlAccountCallStreamsJsonParamsTrack = "outbound_track"
	TexmlAccountCallStreamsJsonParamsTrackBothTracks    TexmlAccountCallStreamsJsonParamsTrack = "both_tracks"
)
