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
func (r *VoiceSDKCallReportService) Get(ctx context.Context, callID string, opts ...option.RequestOption) (res *[]VoiceSDKCallReport, err error) {
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
func (r *VoiceSDKCallReportService) List(ctx context.Context, query VoiceSDKCallReportListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[VoiceSDKCallReport], err error) {
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
func (r *VoiceSDKCallReportService) ListAutoPaging(ctx context.Context, query VoiceSDKCallReportListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[VoiceSDKCallReport] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// A raw call report stats JSON payload. The schema is intentionally permissive
// because Voice SDK clients can add fields over time.
type VoiceSDKCallReport struct {
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
	Logs VoiceSDKCallReportLogsUnion `json:"logs"`
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
	Stats VoiceSDKCallReportStatsUnion `json:"stats"`
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
func (r VoiceSDKCallReport) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportLogsUnion contains all possible properties and values from
// [[]VoiceSDKCallReportLogEntry], [VoiceSDKCallReportLogsEntries].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfVoiceSDKCallReportLogEntryArray]
type VoiceSDKCallReportLogsUnion struct {
	// This field will be present if the value is a [[]VoiceSDKCallReportLogEntry]
	// instead of an object.
	OfVoiceSDKCallReportLogEntryArray []VoiceSDKCallReportLogEntry `json:",inline"`
	// This field is from variant [VoiceSDKCallReportLogsEntries].
	Entries []VoiceSDKCallReportLogEntry `json:"entries"`
	JSON    struct {
		OfVoiceSDKCallReportLogEntryArray respjson.Field
		Entries                           respjson.Field
		raw                               string
	} `json:"-"`
}

func (u VoiceSDKCallReportLogsUnion) AsVoiceSDKCallReportLogEntryArray() (v []VoiceSDKCallReportLogEntry) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportLogsUnion) AsVoiceSDKCallReportLogsEntries() (v VoiceSDKCallReportLogsEntries) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportLogsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportLogsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw logs object emitted by the Voice SDK when logs are grouped under an entries
// field.
type VoiceSDKCallReportLogsEntries struct {
	// Raw log entries when the SDK groups logs under an entries field.
	Entries     []VoiceSDKCallReportLogEntry `json:"entries"`
	ExtraFields map[string]any               `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSDKCallReportLogsEntries) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportLogsEntries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VoiceSDKCallReportStatsUnion contains all possible properties and values from
// [[]map[string]any], [VoiceSDKCallReportStatsUnionMember1].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfAnyMap]
type VoiceSDKCallReportStatsUnion struct {
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	// This field is from variant [VoiceSDKCallReportStatsUnionMember1].
	Audio map[string]any `json:"audio"`
	// This field is from variant [VoiceSDKCallReportStatsUnionMember1].
	Connection map[string]any `json:"connection"`
	// This field is from variant [VoiceSDKCallReportStatsUnionMember1].
	Ice map[string]any `json:"ice"`
	// This field is from variant [VoiceSDKCallReportStatsUnionMember1].
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

func (u VoiceSDKCallReportStatsUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSDKCallReportStatsUnion) AsVoiceSDKCallReportStatsUnionMember1() (v VoiceSDKCallReportStatsUnionMember1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSDKCallReportStatsUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSDKCallReportStatsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Raw stats object emitted by the Voice SDK.
type VoiceSDKCallReportStatsUnionMember1 struct {
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
func (r VoiceSDKCallReportStatsUnionMember1) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportStatsUnionMember1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A raw Voice SDK log entry. Additional SDK-specific fields may be present.
type VoiceSDKCallReportLogEntry struct {
	// Raw structured context attached to the log entry.
	Context map[string]any `json:"context"`
	// Log level emitted by the SDK.
	//
	// Any of "debug", "info", "warn", "error".
	Level VoiceSDKCallReportLogEntryLevel `json:"level"`
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
func (r VoiceSDKCallReportLogEntry) RawJSON() string { return r.JSON.raw }
func (r *VoiceSDKCallReportLogEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log level emitted by the SDK.
type VoiceSDKCallReportLogEntryLevel string

const (
	VoiceSDKCallReportLogEntryLevelDebug VoiceSDKCallReportLogEntryLevel = "debug"
	VoiceSDKCallReportLogEntryLevelInfo  VoiceSDKCallReportLogEntryLevel = "info"
	VoiceSDKCallReportLogEntryLevelWarn  VoiceSDKCallReportLogEntryLevel = "warn"
	VoiceSDKCallReportLogEntryLevelError VoiceSDKCallReportLogEntryLevel = "error"
)

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
