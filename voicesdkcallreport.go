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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Retrieve raw Voice SDK call report stats payloads for WebRTC call
// troubleshooting.
//
// VoiceSDKCallReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceSDKCallReportService] method instead.
type VoiceSDKCallReportService struct {
	Options []option.RequestOption
}

// NewVoiceSDKCallReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVoiceSDKCallReportService(opts ...option.RequestOption) (r VoiceSDKCallReportService) {
	r = VoiceSDKCallReportService{}
	r.Options = opts
	return
}

// Returns raw call report stats JSON payloads stored for the authenticated user
// and `call_id`. The user is derived from Telnyx authentication, not from request
// parameters.
func (r *VoiceSDKCallReportService) Get(ctx context.Context, callID string, opts ...option.RequestOption) (res *[]VoiceSDKCallReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callID == "" {
		err = errors.New("missing required call_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("voice_sdk_call_reports/%s", callID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns paginated raw call report stats JSON payloads stored for the
// authenticated user. The user is derived from Telnyx authentication, not from
// request parameters.
func (r *VoiceSDKCallReportService) List(ctx context.Context, query VoiceSDKCallReportListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[VoiceSDKCallReportListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "voice_sdk_call_reports"
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

// Returns paginated raw call report stats JSON payloads stored for the
// authenticated user. The user is derived from Telnyx authentication, not from
// request parameters.
func (r *VoiceSDKCallReportService) ListAutoPaging(ctx context.Context, query VoiceSDKCallReportListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[VoiceSDKCallReportListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// A raw call report stats JSON payload. The schema is intentionally permissive
// because Voice SDK clients can add fields over time.
type VoiceSDKCallReportGetResponse struct {
	// Unique call identifier.
	CallID string `json:"call_id" format:"uuid"`
	// User-scoped storage grouping identifier derived from the authenticated user.
	// This is not a unique per-call report identifier and may be shared by multiple
	// calls for the same user.
	CallReportID string `json:"call_report_id"`
	// Creation timestamp when present.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reason the SDK flushed this stats report segment, for example an intermediate
	// socket-close flush.
	FlushReason map[string]any `json:"flushReason"`
	// Raw logs payload emitted by the Voice SDK and stored without normalization. Live
	// responses commonly return an array of log entries, but object-shaped log
	// payloads are also allowed for compatibility.
	Logs VoiceSDKCallReportGetResponseLogsUnion `json:"logs"`
	// Organization associated with the stored call report when provided by the Voice
	// SDK reporting path.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Zero-based stats segment index when the SDK sends segmented or intermediate
	// reports.
	Segment int64 `json:"segment"`
	// Raw stats payload emitted by the Voice SDK and stored without normalization. The
	// exact shape can vary by SDK platform and version. Live responses commonly return
	// an array of interval snapshots, but object-shaped stats payloads are also
	// allowed for compatibility.
	Stats VoiceSDKCallReportGetResponseStatsUnion `json:"stats"`
	// Time when the call report was stored.
	StoredAt time.Time `json:"stored_at" format:"date-time"`
	// High-level call metadata.
	Summary map[string]any `json:"summary"`
	// Telnyx call leg identifier for correlating the report with call-control, SIP,
	// and media troubleshooting data.
	TelnyxLegID string `json:"telnyx_leg_id"`
	// Telnyx RTC session identifier for correlating the report with Voice SDK
	// signaling and media-session logs.
	TelnyxSessionID string `json:"telnyx_session_id"`
	// Voice SDK user agent string reported by the client. This is the preferred
	// SDK/platform/version dimension when present.
	UserAgent string `json:"user_agent"`
	// Authenticated user that owns the call report.
	UserID string `json:"user_id" format:"uuid"`
	// Legacy SDK version value when the client reports one separately from the user
	// agent.
	Version string `json:"version"`
	// Voice SDK instance identifier.
	VoiceSDKID string `json:"voice_sdk_id"`
	// Decoded Voice SDK identifier metadata emitted by voice-sdk-proxy when available.
	VoiceSDKIDDecoded map[string]any `json:"voice_sdk_id_decoded"`
	// Voice SDK session correlation identifier used to group stats segments for the
	// same SDK session.
	VoiceSDKSessionID string         `json:"voice_sdk_session_id"`
	ExtraFields       map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID            respjson.Field
		CallReportID      respjson.Field
		CreatedAt         respjson.Field
		FlushReason       respjson.Field
		Logs              respjson.Field
		OrganizationID    respjson.Field
		Segment           respjson.Field
		Stats             respjson.Field
		StoredAt          respjson.Field
		Summary           respjson.Field
		TelnyxLegID       respjson.Field
		TelnyxSessionID   respjson.Field
		UserAgent         respjson.Field
		UserID            respjson.Field
		Version           respjson.Field
		VoiceSDKID        respjson.Field
		VoiceSDKIDDecoded respjson.Field
		VoiceSDKSessionID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportGetResponseLogsUnion contains all possible properties and
// values from [[]VoiceSDKCallReportGetResponseLogsArrayItem],
// [VoiceSDKCallReportGetResponseLogsEntries].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfVoiceSDKCallReportGetResponseLogsArray]
type VoiceSDKCallReportGetResponseLogsUnion struct {
	// This field will be present if the value is a
	// [[]VoiceSDKCallReportGetResponseLogsArrayItem] instead of an object.
	OfVoiceSDKCallReportGetResponseLogsArray []VoiceSDKCallReportGetResponseLogsArrayItem `json:",inline"`
	// This field is from variant [VoiceSDKCallReportGetResponseLogsEntries].
	Entries []VoiceSDKCallReportGetResponseLogsEntriesEntry `json:"entries"`
	JSON    struct {
		OfVoiceSDKCallReportGetResponseLogsArray respjson.Field
		Entries                                  respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u VoiceSDKCallReportGetResponseLogsUnion) AsVoiceSDKCallReportGetResponseLogsArray() (v []VoiceSDKCallReportGetResponseLogsArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportGetResponseLogsUnion) AsVoiceSDKCallReportGetResponseLogsEntries() (v VoiceSDKCallReportGetResponseLogsEntries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportGetResponseLogsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportGetResponseLogsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw Voice SDK log entry. Additional SDK-specific fields may be present.
type VoiceSDKCallReportGetResponseLogsArrayItem struct {
	// Raw structured context attached to the log entry.
	Context map[string]any `json:"context"`
	// Log level emitted by the SDK.
	//
	// Any of "debug", "info", "warn", "error".
	Level string `json:"level"`
	// Log message.
	Message string `json:"message"`
	// Time when the log entry was emitted.
	Timestamp   time.Time      `json:"timestamp" format:"date-time"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportGetResponseLogsArrayItem) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportGetResponseLogsArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw logs object emitted by the Voice SDK when logs are grouped under an entries
// field.
type VoiceSDKCallReportGetResponseLogsEntries struct {
	// Raw log entries when the SDK groups logs under an entries field.
	Entries     []VoiceSDKCallReportGetResponseLogsEntriesEntry `json:"entries"`
	ExtraFields map[string]any                                  `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportGetResponseLogsEntries) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportGetResponseLogsEntries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw Voice SDK log entry. Additional SDK-specific fields may be present.
type VoiceSDKCallReportGetResponseLogsEntriesEntry struct {
	// Raw structured context attached to the log entry.
	Context map[string]any `json:"context"`
	// Log level emitted by the SDK.
	//
	// Any of "debug", "info", "warn", "error".
	Level string `json:"level"`
	// Log message.
	Message string `json:"message"`
	// Time when the log entry was emitted.
	Timestamp   time.Time      `json:"timestamp" format:"date-time"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportGetResponseLogsEntriesEntry) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportGetResponseLogsEntriesEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportGetResponseStatsUnion contains all possible properties and
// values from [[]map[string]any], [VoiceSDKCallReportGetResponseStatsObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfAnyMap]
type VoiceSDKCallReportGetResponseStatsUnion struct {
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	// This field is from variant [VoiceSDKCallReportGetResponseStatsObject].
	Audio map[string]any `json:"audio"`
	// This field is from variant [VoiceSDKCallReportGetResponseStatsObject].
	Connection map[string]any `json:"connection"`
	// This field is from variant [VoiceSDKCallReportGetResponseStatsObject].
	Ice map[string]any `json:"ice"`
	// This field is from variant [VoiceSDKCallReportGetResponseStatsObject].
	Transport map[string]any `json:"transport"`
	JSON      struct {
		OfMapOfAnyMap respjson.Field
		Audio         respjson.Field
		Connection    respjson.Field
		Ice           respjson.Field
		Transport     respjson.Field
		raw           string
	} `json:"-"`
}

func (u VoiceSDKCallReportGetResponseStatsUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportGetResponseStatsUnion) AsVoiceSDKCallReportGetResponseStatsObject() (v VoiceSDKCallReportGetResponseStatsObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportGetResponseStatsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportGetResponseStatsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw stats object emitted by the Voice SDK.
type VoiceSDKCallReportGetResponseStatsObject struct {
	// Raw audio stats such as inbound/outbound packet, byte, jitter, packet-loss,
	// bitrate, and audio-level metrics.
	Audio map[string]any `json:"audio"`
	// Raw connection stats such as round-trip time, packets, and bytes sent or
	// received.
	Connection map[string]any `json:"connection"`
	// Raw ICE candidate-pair information, including selected pair, local/remote
	// candidates, state, and nomination data when provided by the SDK.
	Ice map[string]any `json:"ice"`
	// Raw transport stats such as ICE state, DTLS state, SRTP cipher, TLS version, and
	// selected-candidate-pair changes.
	Transport   map[string]any `json:"transport"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio       respjson.Field
		Connection  respjson.Field
		Ice         respjson.Field
		Transport   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportGetResponseStatsObject) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportGetResponseStatsObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw call report stats JSON payload. The schema is intentionally permissive
// because Voice SDK clients can add fields over time.
type VoiceSDKCallReportListResponse struct {
	// Unique call identifier.
	CallID string `json:"call_id" format:"uuid"`
	// User-scoped storage grouping identifier derived from the authenticated user.
	// This is not a unique per-call report identifier and may be shared by multiple
	// calls for the same user.
	CallReportID string `json:"call_report_id"`
	// Creation timestamp when present.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Reason the SDK flushed this stats report segment, for example an intermediate
	// socket-close flush.
	FlushReason map[string]any `json:"flushReason"`
	// Raw logs payload emitted by the Voice SDK and stored without normalization. Live
	// responses commonly return an array of log entries, but object-shaped log
	// payloads are also allowed for compatibility.
	Logs VoiceSDKCallReportListResponseLogsUnion `json:"logs"`
	// Organization associated with the stored call report when provided by the Voice
	// SDK reporting path.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Zero-based stats segment index when the SDK sends segmented or intermediate
	// reports.
	Segment int64 `json:"segment"`
	// Raw stats payload emitted by the Voice SDK and stored without normalization. The
	// exact shape can vary by SDK platform and version. Live responses commonly return
	// an array of interval snapshots, but object-shaped stats payloads are also
	// allowed for compatibility.
	Stats VoiceSDKCallReportListResponseStatsUnion `json:"stats"`
	// Time when the call report was stored.
	StoredAt time.Time `json:"stored_at" format:"date-time"`
	// High-level call metadata.
	Summary map[string]any `json:"summary"`
	// Telnyx call leg identifier for correlating the report with call-control, SIP,
	// and media troubleshooting data.
	TelnyxLegID string `json:"telnyx_leg_id"`
	// Telnyx RTC session identifier for correlating the report with Voice SDK
	// signaling and media-session logs.
	TelnyxSessionID string `json:"telnyx_session_id"`
	// Voice SDK user agent string reported by the client. This is the preferred
	// SDK/platform/version dimension when present.
	UserAgent string `json:"user_agent"`
	// Authenticated user that owns the call report.
	UserID string `json:"user_id" format:"uuid"`
	// Legacy SDK version value when the client reports one separately from the user
	// agent.
	Version string `json:"version"`
	// Voice SDK instance identifier.
	VoiceSDKID string `json:"voice_sdk_id"`
	// Decoded Voice SDK identifier metadata emitted by voice-sdk-proxy when available.
	VoiceSDKIDDecoded map[string]any `json:"voice_sdk_id_decoded"`
	// Voice SDK session correlation identifier used to group stats segments for the
	// same SDK session.
	VoiceSDKSessionID string         `json:"voice_sdk_session_id"`
	ExtraFields       map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID            respjson.Field
		CallReportID      respjson.Field
		CreatedAt         respjson.Field
		FlushReason       respjson.Field
		Logs              respjson.Field
		OrganizationID    respjson.Field
		Segment           respjson.Field
		Stats             respjson.Field
		StoredAt          respjson.Field
		Summary           respjson.Field
		TelnyxLegID       respjson.Field
		TelnyxSessionID   respjson.Field
		UserAgent         respjson.Field
		UserID            respjson.Field
		Version           respjson.Field
		VoiceSDKID        respjson.Field
		VoiceSDKIDDecoded respjson.Field
		VoiceSDKSessionID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportListResponseLogsUnion contains all possible properties and
// values from [[]VoiceSDKCallReportListResponseLogsArrayItem],
// [VoiceSDKCallReportListResponseLogsEntries].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfVoiceSDKCallReportListResponseLogsArray]
type VoiceSDKCallReportListResponseLogsUnion struct {
	// This field will be present if the value is a
	// [[]VoiceSDKCallReportListResponseLogsArrayItem] instead of an object.
	OfVoiceSDKCallReportListResponseLogsArray []VoiceSDKCallReportListResponseLogsArrayItem `json:",inline"`
	// This field is from variant [VoiceSDKCallReportListResponseLogsEntries].
	Entries []VoiceSDKCallReportListResponseLogsEntriesEntry `json:"entries"`
	JSON    struct {
		OfVoiceSDKCallReportListResponseLogsArray respjson.Field
		Entries                                   respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u VoiceSDKCallReportListResponseLogsUnion) AsVoiceSDKCallReportListResponseLogsArray() (v []VoiceSDKCallReportListResponseLogsArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportListResponseLogsUnion) AsVoiceSDKCallReportListResponseLogsEntries() (v VoiceSDKCallReportListResponseLogsEntries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportListResponseLogsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportListResponseLogsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw Voice SDK log entry. Additional SDK-specific fields may be present.
type VoiceSDKCallReportListResponseLogsArrayItem struct {
	// Raw structured context attached to the log entry.
	Context map[string]any `json:"context"`
	// Log level emitted by the SDK.
	//
	// Any of "debug", "info", "warn", "error".
	Level string `json:"level"`
	// Log message.
	Message string `json:"message"`
	// Time when the log entry was emitted.
	Timestamp   time.Time      `json:"timestamp" format:"date-time"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportListResponseLogsArrayItem) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportListResponseLogsArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw logs object emitted by the Voice SDK when logs are grouped under an entries
// field.
type VoiceSDKCallReportListResponseLogsEntries struct {
	// Raw log entries when the SDK groups logs under an entries field.
	Entries     []VoiceSDKCallReportListResponseLogsEntriesEntry `json:"entries"`
	ExtraFields map[string]any                                   `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportListResponseLogsEntries) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportListResponseLogsEntries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw Voice SDK log entry. Additional SDK-specific fields may be present.
type VoiceSDKCallReportListResponseLogsEntriesEntry struct {
	// Raw structured context attached to the log entry.
	Context map[string]any `json:"context"`
	// Log level emitted by the SDK.
	//
	// Any of "debug", "info", "warn", "error".
	Level string `json:"level"`
	// Log message.
	Message string `json:"message"`
	// Time when the log entry was emitted.
	Timestamp   time.Time      `json:"timestamp" format:"date-time"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportListResponseLogsEntriesEntry) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportListResponseLogsEntriesEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportListResponseStatsUnion contains all possible properties and
// values from [[]map[string]any], [VoiceSDKCallReportListResponseStatsObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfAnyMap]
type VoiceSDKCallReportListResponseStatsUnion struct {
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	// This field is from variant [VoiceSDKCallReportListResponseStatsObject].
	Audio map[string]any `json:"audio"`
	// This field is from variant [VoiceSDKCallReportListResponseStatsObject].
	Connection map[string]any `json:"connection"`
	// This field is from variant [VoiceSDKCallReportListResponseStatsObject].
	Ice map[string]any `json:"ice"`
	// This field is from variant [VoiceSDKCallReportListResponseStatsObject].
	Transport map[string]any `json:"transport"`
	JSON      struct {
		OfMapOfAnyMap respjson.Field
		Audio         respjson.Field
		Connection    respjson.Field
		Ice           respjson.Field
		Transport     respjson.Field
		raw           string
	} `json:"-"`
}

func (u VoiceSDKCallReportListResponseStatsUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportListResponseStatsUnion) AsVoiceSDKCallReportListResponseStatsObject() (v VoiceSDKCallReportListResponseStatsObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportListResponseStatsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportListResponseStatsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw stats object emitted by the Voice SDK.
type VoiceSDKCallReportListResponseStatsObject struct {
	// Raw audio stats such as inbound/outbound packet, byte, jitter, packet-loss,
	// bitrate, and audio-level metrics.
	Audio map[string]any `json:"audio"`
	// Raw connection stats such as round-trip time, packets, and bytes sent or
	// received.
	Connection map[string]any `json:"connection"`
	// Raw ICE candidate-pair information, including selected pair, local/remote
	// candidates, state, and nomination data when provided by the SDK.
	Ice map[string]any `json:"ice"`
	// Raw transport stats such as ICE state, DTLS state, SRTP cipher, TLS version, and
	// selected-candidate-pair changes.
	Transport   map[string]any `json:"transport"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio       respjson.Field
		Connection  respjson.Field
		Ice         respjson.Field
		Transport   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportListResponseStatsObject) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportListResponseStatsObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSDKCallReportListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Set the order of the results by creation date. `asc` and `created_at` sort
	// oldest reports first; `desc` and `-created_at` sort newest reports first. If not
	// given, results are sorted by creation date in descending order.
	//
	// Any of "asc", "desc", "created_at", "-created_at".
	Sort VoiceSDKCallReportListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceSDKCallReportListParams]'s query parameters as
// `url.Values`.
func (r VoiceSDKCallReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set the order of the results by creation date. `asc` and `created_at` sort
// oldest reports first; `desc` and `-created_at` sort newest reports first. If not
// given, results are sorted by creation date in descending order.
type VoiceSDKCallReportListParamsSort string

const (
	VoiceSDKCallReportListParamsSortAsc           VoiceSDKCallReportListParamsSort = "asc"
	VoiceSDKCallReportListParamsSortDesc          VoiceSDKCallReportListParamsSort = "desc"
	VoiceSDKCallReportListParamsSortCreatedAt     VoiceSDKCallReportListParamsSort = "created_at"
	VoiceSDKCallReportListParamsSortCreatedAtDesc VoiceSDKCallReportListParamsSort = "-created_at"
)
