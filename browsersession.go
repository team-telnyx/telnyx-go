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

// BrowserSessionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserSessionService] method instead.
type BrowserSessionService struct {
	Options []option.RequestOption
	// Execute browser actions: navigate, observe, click, fill, select, press, scroll,
	// wait_for, screenshot, batch.
	Actions BrowserSessionActionService
}

// NewBrowserSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrowserSessionService(opts ...option.RequestOption) (r BrowserSessionService) {
	r = BrowserSessionService{}
	r.Options = opts
	r.Actions = NewBrowserSessionActionService(opts...)
	return
}

// Creates an isolated browser session with network access restricted to the
// domains you allow.
func (r *BrowserSessionService) New(ctx context.Context, body BrowserSessionNewParams, opts ...option.RequestOption) (res *BrowserSessionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browser_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the current status of a browser session.
func (r *BrowserSessionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BrowserSessionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("browser_sessions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists sessions for the authenticated org. Optional status filter.
func (r *BrowserSessionService) List(ctx context.Context, query BrowserSessionListParams, opts ...option.RequestOption) (res *BrowserSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "browser_sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Closes a ready session and releases its resources. An already closed, failed, or
// expired session is returned unchanged; a provisioning session returns 409
// session_not_ready.
func (r *BrowserSessionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BrowserSessionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("browser_sessions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns artifact metadata and a signed download URL. The download URL requires
// Bearer auth.
func (r *BrowserSessionService) GetArtifact(ctx context.Context, artifactID string, query BrowserSessionGetArtifactParams, opts ...option.RequestOption) (res *BrowserSessionGetArtifactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifactId parameter")
		return nil, err
	}
	path := fmt.Sprintf("browser_sessions/%s/artifacts/%s", url.PathEscape(query.ID), url.PathEscape(artifactID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Session struct {
	ID             string        `json:"id" api:"required"`
	ClosedAt       time.Time     `json:"closed_at" api:"required" format:"date-time"`
	CreatedAt      time.Time     `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt      time.Time     `json:"expires_at" api:"required" format:"date-time"`
	FailureReason  string        `json:"failure_reason" api:"required"`
	LastActivityAt time.Time     `json:"last_activity_at" api:"required" format:"date-time"`
	Limits         SessionLimits `json:"limits" api:"required"`
	Metadata       any           `json:"metadata" api:"required"`
	Policy         SessionPolicy `json:"policy" api:"required"`
	ReadyAt        time.Time     `json:"ready_at" api:"required" format:"date-time"`
	// Any of "provisioning", "ready", "closing", "closed", "failed", "expired".
	Status SessionStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ClosedAt       respjson.Field
		CreatedAt      respjson.Field
		ExpiresAt      respjson.Field
		FailureReason  respjson.Field
		LastActivityAt respjson.Field
		Limits         respjson.Field
		Metadata       respjson.Field
		Policy         respjson.Field
		ReadyAt        respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Session) RawJSON() string { return r.JSON.raw }
func (r *Session) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionLimits struct {
	IdleTimeoutS int64 `json:"idle_timeout_s" api:"required"`
	LifetimeS    int64 `json:"lifetime_s" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IdleTimeoutS respjson.Field
		LifetimeS    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionLimits) RawJSON() string { return r.JSON.raw }
func (r *SessionLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionPolicy struct {
	AllowedDomains []string `json:"allowed_domains" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedDomains respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionPolicy) RawJSON() string { return r.JSON.raw }
func (r *SessionPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionStatus string

const (
	SessionStatusProvisioning SessionStatus = "provisioning"
	SessionStatusReady        SessionStatus = "ready"
	SessionStatusClosing      SessionStatus = "closing"
	SessionStatusClosed       SessionStatus = "closed"
	SessionStatusFailed       SessionStatus = "failed"
	SessionStatusExpired      SessionStatus = "expired"
)

type BrowserSessionNewResponse struct {
	Data Session `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionGetResponse struct {
	Data Session `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionListResponse struct {
	Data []Session `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionDeleteResponse struct {
	Data Session `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionGetArtifactResponse struct {
	Data BrowserSessionGetArtifactResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionGetArtifactResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionGetArtifactResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionGetArtifactResponseData struct {
	ID          string    `json:"id" api:"required"`
	ActionID    string    `json:"action_id" api:"required"`
	Checksum    string    `json:"checksum" api:"required"`
	ContentType string    `json:"content_type" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// Signed relative path (starts with /v2/artifact_downloads/). Prepend your API
	// host, e.g. https://api.telnyx.com, before calling it.
	DownloadURL string    `json:"download_url" api:"required"`
	ExpiresAt   time.Time `json:"expires_at" api:"required" format:"date-time"`
	SessionID   string    `json:"session_id" api:"required"`
	SizeBytes   int64     `json:"size_bytes" api:"required"`
	Type        string    `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ActionID    respjson.Field
		Checksum    respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		DownloadURL respjson.Field
		ExpiresAt   respjson.Field
		SessionID   respjson.Field
		SizeBytes   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionGetArtifactResponseData) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionGetArtifactResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionNewParams struct {
	Policy BrowserSessionNewParamsPolicy `json:"policy,omitzero" api:"required"`
	// Scoped to the organization. The same key and identical normalized policy,
	// limits, and metadata replay the existing session (HTTP 201); a different request
	// returns 409 idempotency_conflict. Replays do not restart terminal sessions.
	IdempotencyKey param.Opt[string]             `json:"idempotency_key,omitzero"`
	Limits         BrowserSessionNewParamsLimits `json:"limits,omitzero"`
	// Arbitrary key-value metadata.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r BrowserSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BrowserSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AllowedDomains is required.
type BrowserSessionNewParamsPolicy struct {
	// Domain patterns the browser can navigate to. Exact match (example.com) or
	// wildcard (\*.example.com). IP literals not allowed.
	AllowedDomains []string `json:"allowed_domains,omitzero" api:"required"`
	paramObj
}

func (r BrowserSessionNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	type shadow BrowserSessionNewParamsPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserSessionNewParamsPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionNewParamsLimits struct {
	// Idle timeout in seconds. Must not exceed lifetime_s.
	IdleTimeoutS param.Opt[int64] `json:"idle_timeout_s,omitzero"`
	// Max session duration in seconds (60-1800).
	LifetimeS param.Opt[int64] `json:"lifetime_s,omitzero"`
	paramObj
}

func (r BrowserSessionNewParamsLimits) MarshalJSON() (data []byte, err error) {
	type shadow BrowserSessionNewParamsLimits
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserSessionNewParamsLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionListParams struct {
	// Any of "provisioning", "ready", "closing", "closed", "failed", "expired".
	Status BrowserSessionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrowserSessionListParams]'s query parameters as
// `url.Values`.
func (r BrowserSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrowserSessionListParamsStatus string

const (
	BrowserSessionListParamsStatusProvisioning BrowserSessionListParamsStatus = "provisioning"
	BrowserSessionListParamsStatusReady        BrowserSessionListParamsStatus = "ready"
	BrowserSessionListParamsStatusClosing      BrowserSessionListParamsStatus = "closing"
	BrowserSessionListParamsStatusClosed       BrowserSessionListParamsStatus = "closed"
	BrowserSessionListParamsStatusFailed       BrowserSessionListParamsStatus = "failed"
	BrowserSessionListParamsStatusExpired      BrowserSessionListParamsStatus = "expired"
)

type BrowserSessionGetArtifactParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
