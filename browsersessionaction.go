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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
	"github.com/tidwall/gjson"
)

// Execute browser actions: navigate, observe, click, fill, select, press, scroll,
// wait_for, screenshot, batch.
//
// BrowserSessionActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowserSessionActionService] method instead.
type BrowserSessionActionService struct {
	Options []option.RequestOption
}

// NewBrowserSessionActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrowserSessionActionService(opts ...option.RequestOption) (r BrowserSessionActionService) {
	r = BrowserSessionActionService{}
	r.Options = opts
	return
}

// Executes a browser action on the session. Returns 200 if completed within
// wait_seconds, 202 if still in-flight.
//
// Action types: navigate, observe, click, fill, select, press, scroll, wait_for,
// screenshot, batch.
func (r *BrowserSessionActionService) New(ctx context.Context, id string, params BrowserSessionActionNewParams, opts ...option.RequestOption) (res *BrowserSessionActionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("browser_sessions/%s/actions", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the stored status and result of an action. Use this to poll actions that
// returned 202; HTTP 200 can contain any action status. Stored observations
// contain metadata with empty tree and refs, not the full snapshot; submit a fresh
// observe action to obtain current refs.
func (r *BrowserSessionActionService) Get(ctx context.Context, actionID string, query BrowserSessionActionGetParams, opts ...option.RequestOption) (res *BrowserSessionActionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if actionID == "" {
		err = errors.New("missing required actionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("browser_sessions/%s/actions/%s", url.PathEscape(query.ID), url.PathEscape(actionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ActionResponse struct {
	ID          string    `json:"id" api:"required"`
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// Null unless an execution failure is recorded. Action failures (policy_denied,
	// stale_observation, stale_ref, navigation_blocked, timeout, execution_lost,
	// result_too_large) are returned on the action resource with HTTP 200. An unknown
	// action with execution_lost has retryable=false; inspect session state and
	// reconcile possible effects before any new action. Separately, a request may
	// return HTTP 503 with a top-level execution_lost error and retryable=true; that
	// is not proof of safe re-execution.
	Error ActionResponseError `json:"error" api:"required"`
	// Normalized request metadata; defaults may be materialized and fill values are
	// redacted as {redacted: true, length: number}, including inside batch actions.
	Request   any                  `json:"request" api:"required"`
	Result    ActionResponseResult `json:"result" api:"required"`
	Seq       int64                `json:"seq" api:"required"`
	SessionID string               `json:"session_id" api:"required"`
	StartedAt time.Time            `json:"started_at" api:"required" format:"date-time"`
	// POST returns 202 for queued or running actions. Terminal states are succeeded,
	// failed, and unknown (the execution outcome could not be determined, with
	// error.code=execution_lost). GET returns 200 for both nonterminal and terminal
	// states.
	//
	// Any of "queued", "running", "succeeded", "failed", "unknown".
	Status ActionResponseStatus `json:"status" api:"required"`
	Type   string               `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Error       respjson.Field
		Request     respjson.Field
		Result      respjson.Field
		Seq         respjson.Field
		SessionID   respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Null unless an execution failure is recorded. Action failures (policy_denied,
// stale_observation, stale_ref, navigation_blocked, timeout, execution_lost,
// result_too_large) are returned on the action resource with HTTP 200. An unknown
// action with execution_lost has retryable=false; inspect session state and
// reconcile possible effects before any new action. Separately, a request may
// return HTTP 503 with a top-level execution_lost error and retryable=true; that
// is not proof of safe re-execution.
type ActionResponseError struct {
	Code      string `json:"code" api:"required"`
	Retryable bool   `json:"retryable" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Retryable   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseError) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionResponseResult struct {
	Artifact ActionResponseResultArtifact `json:"artifact"`
	// Full tree and refs are delivered only in the original successful POST response
	// when completed within wait_seconds. GET polling and idempotent replays return
	// metadata with an empty tree and refs plus ref_count. Submit a fresh observe
	// action to obtain usable refs.
	Observation ActionResponseResultObservation `json:"observation"`
	// Present only for batch actions.
	Results []any  `json:"results"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Artifact    respjson.Field
		Observation respjson.Field
		Results     respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionResponseResultArtifact struct {
	ArtifactID  string `json:"artifact_id"`
	ContentType string `json:"content_type"`
	SizeBytes   int64  `json:"size_bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArtifactID  respjson.Field
		ContentType respjson.Field
		SizeBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseResultArtifact) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseResultArtifact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full tree and refs are delivered only in the original successful POST response
// when completed within wait_seconds. GET polling and idempotent replays return
// metadata with an empty tree and refs plus ref_count. Submit a fresh observe
// action to obtain usable refs.
type ActionResponseResultObservation struct {
	ObservationID string `json:"observation_id"`
	// Number of refs in the original observation, present on stored observation
	// metadata.
	RefCount  int64  `json:"ref_count"`
	Refs      any    `json:"refs"`
	Title     string `json:"title"`
	Tree      string `json:"tree"`
	Truncated bool   `json:"truncated"`
	URL       string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObservationID respjson.Field
		RefCount      respjson.Field
		Refs          respjson.Field
		Title         respjson.Field
		Tree          respjson.Field
		Truncated     respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseResultObservation) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseResultObservation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// POST returns 202 for queued or running actions. Terminal states are succeeded,
// failed, and unknown (the execution outcome could not be determined, with
// error.code=execution_lost). GET returns 200 for both nonterminal and terminal
// states.
type ActionResponseStatus string

const (
	ActionResponseStatusQueued    ActionResponseStatus = "queued"
	ActionResponseStatusRunning   ActionResponseStatus = "running"
	ActionResponseStatusSucceeded ActionResponseStatus = "succeeded"
	ActionResponseStatusFailed    ActionResponseStatus = "failed"
	ActionResponseStatusUnknown   ActionResponseStatus = "unknown"
)

// Click an element identified by observation_id + ref.
//
// The properties ObservationID, Ref, Type are required.
type ClickActionParam struct {
	ObservationID      string            `json:"observation_id" api:"required"`
	Ref                string            `json:"ref" api:"required"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	TimeoutMs          param.Opt[int64]  `json:"timeout_ms,omitzero"`
	// Any of "no_wait", "dom_stable", "url_change".
	Wait ClickActionWait `json:"wait,omitzero"`
	// This field can be elided, and will marshal its zero value as "click".
	Type constant.Click `json:"type" default:"click"`
	paramObj
}

func (r ClickActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ClickActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClickActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClickActionWait string

const (
	ClickActionWaitNoWait    ClickActionWait = "no_wait"
	ClickActionWaitDomStable ClickActionWait = "dom_stable"
	ClickActionWaitURLChange ClickActionWait = "url_change"
)

// Type text into the element identified by observation_id + ref.
//
// The properties ObservationID, Ref, Type, Value are required.
type FillActionParam struct {
	ObservationID      string            `json:"observation_id" api:"required"`
	Ref                string            `json:"ref" api:"required"`
	Value              string            `json:"value" api:"required"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "fill".
	Type constant.Fill `json:"type" default:"fill"`
	paramObj
}

func (r FillActionParam) MarshalJSON() (data []byte, err error) {
	type shadow FillActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FillActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Navigate to a URL.
//
// The properties Type, URL are required.
type NavigateActionParam struct {
	URL                string            `json:"url" api:"required" format:"uri"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// Any of "load", "dom_stable", "none".
	WaitUntil NavigateActionWaitUntil `json:"wait_until,omitzero"`
	// This field can be elided, and will marshal its zero value as "navigate".
	Type constant.Navigate `json:"type" default:"navigate"`
	paramObj
}

func (r NavigateActionParam) MarshalJSON() (data []byte, err error) {
	type shadow NavigateActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NavigateActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NavigateActionWaitUntil string

const (
	NavigateActionWaitUntilLoad      NavigateActionWaitUntil = "load"
	NavigateActionWaitUntilDomStable NavigateActionWaitUntil = "dom_stable"
	NavigateActionWaitUntilNone      NavigateActionWaitUntil = "none"
)

// Capture an observation (DOM snapshot) of the current page.
//
// The property Type is required.
type ObserveActionParam struct {
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "observe".
	Type constant.Observe `json:"type" default:"observe"`
	paramObj
}

func (r ObserveActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ObserveActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObserveActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Press a keyboard key.
//
// The properties Key, Type are required.
type PressActionParam struct {
	// Any of "Enter", "Tab", "Escape", "Backspace", "ArrowUp", "ArrowDown",
	// "ArrowLeft", "ArrowRight", "PageUp", "PageDown", "Home", "End".
	Key                PressActionKey    `json:"key,omitzero" api:"required"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "press".
	Type constant.Press `json:"type" default:"press"`
	paramObj
}

func (r PressActionParam) MarshalJSON() (data []byte, err error) {
	type shadow PressActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PressActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PressActionKey string

const (
	PressActionKeyEnter      PressActionKey = "Enter"
	PressActionKeyTab        PressActionKey = "Tab"
	PressActionKeyEscape     PressActionKey = "Escape"
	PressActionKeyBackspace  PressActionKey = "Backspace"
	PressActionKeyArrowUp    PressActionKey = "ArrowUp"
	PressActionKeyArrowDown  PressActionKey = "ArrowDown"
	PressActionKeyArrowLeft  PressActionKey = "ArrowLeft"
	PressActionKeyArrowRight PressActionKey = "ArrowRight"
	PressActionKeyPageUp     PressActionKey = "PageUp"
	PressActionKeyPageDown   PressActionKey = "PageDown"
	PressActionKeyHome       PressActionKey = "Home"
	PressActionKeyEnd        PressActionKey = "End"
)

// Capture a screenshot of the current page.
//
// The property Type is required.
type ScreenshotActionParam struct {
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "screenshot".
	Type constant.Screenshot `json:"type" default:"screenshot"`
	paramObj
}

func (r ScreenshotActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ScreenshotActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScreenshotActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scroll the page.
//
// The property Type is required.
type ScrollActionParam struct {
	DeltaX             param.Opt[int64]  `json:"delta_x,omitzero"`
	DeltaY             param.Opt[int64]  `json:"delta_y,omitzero"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "scroll".
	Type constant.Scroll `json:"type" default:"scroll"`
	paramObj
}

func (r ScrollActionParam) MarshalJSON() (data []byte, err error) {
	type shadow ScrollActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScrollActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Select an option on the element identified by observation_id + ref.
//
// The properties ObservationID, Ref, Type, Value are required.
type SelectActionParam struct {
	ObservationID      string            `json:"observation_id" api:"required"`
	Ref                string            `json:"ref" api:"required"`
	Value              string            `json:"value" api:"required"`
	IdempotencyKey     param.Opt[string] `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]   `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "select".
	Type constant.Select `json:"type" default:"select"`
	paramObj
}

func (r SelectActionParam) MarshalJSON() (data []byte, err error) {
	type shadow SelectActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SelectActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wait for a condition before continuing.
//
// The properties Condition, Type are required.
type WaitForActionParam struct {
	Condition          WaitForActionConditionUnionParam `json:"condition,omitzero" api:"required"`
	IdempotencyKey     param.Opt[string]                `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]                  `json:"include_observation,omitzero"`
	TimeoutMs          param.Opt[int64]                 `json:"timeout_ms,omitzero"`
	// This field can be elided, and will marshal its zero value as "wait_for".
	Type constant.WaitFor `json:"type" default:"wait_for"`
	paramObj
}

func (r WaitForActionParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WaitForActionConditionUnionParam struct {
	OfUnionMember0               *WaitForActionConditionUnionMember0Param `json:",omitzero,inline"`
	OfUnionMember1               *WaitForActionConditionUnionMember1Param `json:",omitzero,inline"`
	OfUnionMember2               *WaitForActionConditionUnionMember2Param `json:",omitzero,inline"`
	OfWaitForActionConditionKind *WaitForActionConditionKindParam         `json:",omitzero,inline"`
	paramUnion
}

func (u WaitForActionConditionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUnionMember0, u.OfUnionMember1, u.OfUnionMember2, u.OfWaitForActionConditionKind)
}
func (u *WaitForActionConditionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WaitForActionConditionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUnionMember0) {
		return u.OfUnionMember0
	} else if !param.IsOmitted(u.OfUnionMember1) {
		return u.OfUnionMember1
	} else if !param.IsOmitted(u.OfUnionMember2) {
		return u.OfUnionMember2
	} else if !param.IsOmitted(u.OfWaitForActionConditionKind) {
		return u.OfWaitForActionConditionKind
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WaitForActionConditionUnionParam) GetPattern() *string {
	if vt := u.OfUnionMember0; vt != nil {
		return &vt.Pattern
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WaitForActionConditionUnionParam) GetObservationID() *string {
	if vt := u.OfUnionMember1; vt != nil {
		return &vt.ObservationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WaitForActionConditionUnionParam) GetRef() *string {
	if vt := u.OfUnionMember1; vt != nil {
		return &vt.Ref
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WaitForActionConditionUnionParam) GetText() *string {
	if vt := u.OfUnionMember2; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WaitForActionConditionUnionParam) GetKind() *string {
	if vt := u.OfUnionMember0; vt != nil {
		return (*string)(&vt.Kind)
	} else if vt := u.OfUnionMember1; vt != nil {
		return (*string)(&vt.Kind)
	} else if vt := u.OfUnionMember2; vt != nil {
		return (*string)(&vt.Kind)
	} else if vt := u.OfWaitForActionConditionKind; vt != nil {
		return (*string)(&vt.Kind)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[WaitForActionConditionUnionParam](
		"",
		apijson.Variant[WaitForActionConditionUnionMember0Param](gjson.JSON),
		apijson.Variant[WaitForActionConditionUnionMember1Param](gjson.JSON),
		apijson.Variant[WaitForActionConditionUnionMember2Param](gjson.JSON),
		apijson.Variant[WaitForActionConditionKindParam](gjson.JSON),
	)
}

// The properties Kind, Pattern are required.
type WaitForActionConditionUnionMember0Param struct {
	Pattern string `json:"pattern" api:"required"`
	// This field can be elided, and will marshal its zero value as "url_pattern".
	Kind constant.URLPattern `json:"kind" default:"url_pattern"`
	paramObj
}

func (r WaitForActionConditionUnionMember0Param) MarshalJSON() (data []byte, err error) {
	type shadow WaitForActionConditionUnionMember0Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForActionConditionUnionMember0Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Kind, ObservationID, Ref are required.
type WaitForActionConditionUnionMember1Param struct {
	ObservationID string `json:"observation_id" api:"required"`
	Ref           string `json:"ref" api:"required"`
	// This field can be elided, and will marshal its zero value as "ref".
	Kind constant.Ref `json:"kind" default:"ref"`
	paramObj
}

func (r WaitForActionConditionUnionMember1Param) MarshalJSON() (data []byte, err error) {
	type shadow WaitForActionConditionUnionMember1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForActionConditionUnionMember1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Kind, Text are required.
type WaitForActionConditionUnionMember2Param struct {
	Text string `json:"text" api:"required"`
	// This field can be elided, and will marshal its zero value as "text".
	Kind constant.Text `json:"kind" default:"text"`
	paramObj
}

func (r WaitForActionConditionUnionMember2Param) MarshalJSON() (data []byte, err error) {
	type shadow WaitForActionConditionUnionMember2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForActionConditionUnionMember2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewWaitForActionConditionKindParam() WaitForActionConditionKindParam {
	return WaitForActionConditionKindParam{
		Kind: "dom_stable",
	}
}

// This struct has a constant value, construct it with
// [NewWaitForActionConditionKindParam].
type WaitForActionConditionKindParam struct {
	Kind constant.DomStable `json:"kind" default:"dom_stable"`
	paramObj
}

func (r WaitForActionConditionKindParam) MarshalJSON() (data []byte, err error) {
	type shadow WaitForActionConditionKindParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WaitForActionConditionKindParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionActionNewResponse struct {
	Data ActionResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionActionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionActionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionActionGetResponse struct {
	Data ActionResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrowserSessionActionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BrowserSessionActionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrowserSessionActionNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	// Navigate to a URL.
	OfNavigateAction *NavigateActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Capture
	// an observation (DOM snapshot) of the current page.
	OfObserveAction *ObserveActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Click
	// an element identified by observation_id + ref.
	OfClickAction *ClickActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Type
	// text into the element identified by observation_id + ref.
	OfFillAction *FillActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Select
	// an option on the element identified by observation_id + ref.
	OfSelectAction *SelectActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Press a
	// keyboard key.
	OfPressAction *PressActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Scroll
	// the page.
	OfScrollAction *ScrollActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Capture
	// a screenshot of the current page.
	OfScreenshotAction *ScreenshotActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Wait
	// for a condition before continuing.
	OfWaitForAction *WaitForActionParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Execute
	// 1-50 sub-actions sequentially. Non-recursive: a batch inside a batch is rejected
	// (400). An empty array is rejected (400). A failed sub-action stops the batch;
	// earlier browser mutations are not rolled back, but partial sub-action results
	// are not returned.
	OfBatchAction *BrowserSessionActionNewParamsBodyBatchAction `json:",inline"`

	// Maximum seconds the submission request waits for completion, not an execution
	// timeout. If still queued or running, returns 202 without cancelling execution;
	// poll the returned action ID.
	WaitSeconds param.Opt[int64] `query:"wait_seconds,omitzero" json:"-"`
	paramObj
}

func (u BrowserSessionActionNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNavigateAction,
		u.OfObserveAction,
		u.OfClickAction,
		u.OfFillAction,
		u.OfSelectAction,
		u.OfPressAction,
		u.OfScrollAction,
		u.OfScreenshotAction,
		u.OfWaitForAction,
		u.OfBatchAction)
}
func (r *BrowserSessionActionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BrowserSessionActionNewParams]'s query parameters as
// `url.Values`.
func (r BrowserSessionActionNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Execute 1-50 sub-actions sequentially. Non-recursive: a batch inside a batch is
// rejected (400). An empty array is rejected (400). A failed sub-action stops the
// batch; earlier browser mutations are not rolled back, but partial sub-action
// results are not returned.
//
// The properties Actions, Type are required.
type BrowserSessionActionNewParamsBodyBatchAction struct {
	Actions            []BrowserSessionActionNewParamsBodyBatchActionActionUnion `json:"actions,omitzero" api:"required"`
	IdempotencyKey     param.Opt[string]                                         `json:"idempotency_key,omitzero"`
	IncludeObservation param.Opt[bool]                                           `json:"include_observation,omitzero"`
	// This field can be elided, and will marshal its zero value as "batch".
	Type constant.Batch `json:"type" default:"batch"`
	paramObj
}

func (r BrowserSessionActionNewParamsBodyBatchAction) MarshalJSON() (data []byte, err error) {
	type shadow BrowserSessionActionNewParamsBodyBatchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowserSessionActionNewParamsBodyBatchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BrowserSessionActionNewParamsBodyBatchActionActionUnion struct {
	OfNavigateAction   *NavigateActionParam   `json:",omitzero,inline"`
	OfObserveAction    *ObserveActionParam    `json:",omitzero,inline"`
	OfClickAction      *ClickActionParam      `json:",omitzero,inline"`
	OfFillAction       *FillActionParam       `json:",omitzero,inline"`
	OfSelectAction     *SelectActionParam     `json:",omitzero,inline"`
	OfPressAction      *PressActionParam      `json:",omitzero,inline"`
	OfScrollAction     *ScrollActionParam     `json:",omitzero,inline"`
	OfScreenshotAction *ScreenshotActionParam `json:",omitzero,inline"`
	OfWaitForAction    *WaitForActionParam    `json:",omitzero,inline"`
	paramUnion
}

func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNavigateAction,
		u.OfObserveAction,
		u.OfClickAction,
		u.OfFillAction,
		u.OfSelectAction,
		u.OfPressAction,
		u.OfScrollAction,
		u.OfScreenshotAction,
		u.OfWaitForAction)
}
func (u *BrowserSessionActionNewParamsBodyBatchActionActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BrowserSessionActionNewParamsBodyBatchActionActionUnion) asAny() any {
	if !param.IsOmitted(u.OfNavigateAction) {
		return u.OfNavigateAction
	} else if !param.IsOmitted(u.OfObserveAction) {
		return u.OfObserveAction
	} else if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfFillAction) {
		return u.OfFillAction
	} else if !param.IsOmitted(u.OfSelectAction) {
		return u.OfSelectAction
	} else if !param.IsOmitted(u.OfPressAction) {
		return u.OfPressAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	} else if !param.IsOmitted(u.OfScreenshotAction) {
		return u.OfScreenshotAction
	} else if !param.IsOmitted(u.OfWaitForAction) {
		return u.OfWaitForAction
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetURL() *string {
	if vt := u.OfNavigateAction; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetWaitUntil() *string {
	if vt := u.OfNavigateAction; vt != nil {
		return (*string)(&vt.WaitUntil)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetWait() *string {
	if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.Wait)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetKey() *string {
	if vt := u.OfPressAction; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetDeltaX() *int64 {
	if vt := u.OfScrollAction; vt != nil && vt.DeltaX.Valid() {
		return &vt.DeltaX.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetDeltaY() *int64 {
	if vt := u.OfScrollAction; vt != nil && vt.DeltaY.Valid() {
		return &vt.DeltaY.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetCondition() *WaitForActionConditionUnionParam {
	if vt := u.OfWaitForAction; vt != nil {
		return &vt.Condition
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetType() *string {
	if vt := u.OfNavigateAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObserveAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFillAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSelectAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPressAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfScrollAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfScreenshotAction; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWaitForAction; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetIdempotencyKey() *string {
	if vt := u.OfNavigateAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfObserveAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfClickAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfFillAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfSelectAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfPressAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfScrollAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfScreenshotAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	} else if vt := u.OfWaitForAction; vt != nil && vt.IdempotencyKey.Valid() {
		return &vt.IdempotencyKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetIncludeObservation() *bool {
	if vt := u.OfNavigateAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfObserveAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfClickAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfFillAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfSelectAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfPressAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfScrollAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfScreenshotAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	} else if vt := u.OfWaitForAction; vt != nil && vt.IncludeObservation.Valid() {
		return &vt.IncludeObservation.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetObservationID() *string {
	if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.ObservationID)
	} else if vt := u.OfFillAction; vt != nil {
		return (*string)(&vt.ObservationID)
	} else if vt := u.OfSelectAction; vt != nil {
		return (*string)(&vt.ObservationID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetRef() *string {
	if vt := u.OfClickAction; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfFillAction; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfSelectAction; vt != nil {
		return (*string)(&vt.Ref)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetTimeoutMs() *int64 {
	if vt := u.OfClickAction; vt != nil && vt.TimeoutMs.Valid() {
		return &vt.TimeoutMs.Value
	} else if vt := u.OfWaitForAction; vt != nil && vt.TimeoutMs.Valid() {
		return &vt.TimeoutMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BrowserSessionActionNewParamsBodyBatchActionActionUnion) GetValue() *string {
	if vt := u.OfFillAction; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfSelectAction; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BrowserSessionActionNewParamsBodyBatchActionActionUnion](
		"",
		apijson.Variant[NavigateActionParam](gjson.JSON),
		apijson.Variant[ObserveActionParam](gjson.JSON),
		apijson.Variant[ClickActionParam](gjson.JSON),
		apijson.Variant[FillActionParam](gjson.JSON),
		apijson.Variant[SelectActionParam](gjson.JSON),
		apijson.Variant[PressActionParam](gjson.JSON),
		apijson.Variant[ScrollActionParam](gjson.JSON),
		apijson.Variant[ScreenshotActionParam](gjson.JSON),
		apijson.Variant[WaitForActionParam](gjson.JSON),
	)
}

type BrowserSessionActionGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
