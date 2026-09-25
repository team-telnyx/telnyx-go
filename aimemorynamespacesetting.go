// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// How a namespace's summaries are written.
//
// AIMemoryNamespaceSettingService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryNamespaceSettingService] method instead.
type AIMemoryNamespaceSettingService struct {
	Options []option.RequestOption
}

// NewAIMemoryNamespaceSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIMemoryNamespaceSettingService(opts ...option.RequestOption) (r AIMemoryNamespaceSettingService) {
	r = AIMemoryNamespaceSettingService{}
	r.Options = opts
	return
}

// What is currently set for this namespace. `instructions: null` means none are
// set and summaries use the neutral default.
func (r *AIMemoryNamespaceSettingService) List(ctx context.Context, namespace string, opts ...option.RequestOption) (res *NamespaceSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/settings", url.PathEscape(namespace))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Only the fields you send are changed; anything omitted is left as it is, so `{}`
// changes nothing. Sending `instructions: null`, or an empty or whitespace-only
// string, clears them and returns summaries to the neutral default.
//
// Instructions are capped at 2000 characters. A longer note is refused rather than
// truncated, because a note cut mid-sentence is a worse steer than none. A change
// reaches each summary the next time that summary is regenerated, not immediately.
func (r *AIMemoryNamespaceSettingService) PatchAll(ctx context.Context, namespace string, body AIMemoryNamespaceSettingPatchAllParams, opts ...option.RequestOption) (res *NamespaceSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/settings", url.PathEscape(namespace))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type NamespaceSettingsResponse struct {
	// A namespace's settings, grouped by what they affect.
	Data NamespaceSettingsResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *NamespaceSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A namespace's settings, grouped by what they affect.
type NamespaceSettingsResponseData struct {
	// Settings that shape this namespace's summaries.
	Summary NamespaceSettingsResponseDataSummary `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceSettingsResponseData) RawJSON() string { return r.JSON.raw }
func (r *NamespaceSettingsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings that shape this namespace's summaries.
type NamespaceSettingsResponseDataSummary struct {
	// Free-form instructions that influence how this namespace's summaries are
	// written, shared by every profile in the namespace. How you use them is up to you
	// -- they steer the outcome, so try a phrasing and see how the summary comes out.
	// Advisory: they steer the summary but never override or deny a profile's own
	// facts, and they do not affect recall. Null or empty means none are set, and
	// summaries use the neutral default. A change reaches each summary the next time
	// it is regenerated.
	Instructions string `json:"instructions" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NamespaceSettingsResponseDataSummary) RawJSON() string { return r.JSON.raw }
func (r *NamespaceSettingsResponseDataSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceSettingPatchAllParams struct {
	// A partial update to a namespace's summary settings.
	//
	// Only the fields present in the request are changed; the rest are left as they
	// are. Sending `instructions: null` (or empty) clears the instructions.
	Summary AIMemoryNamespaceSettingPatchAllParamsSummary `json:"summary,omitzero"`
	paramObj
}

func (r AIMemoryNamespaceSettingPatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMemoryNamespaceSettingPatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMemoryNamespaceSettingPatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A partial update to a namespace's summary settings.
//
// Only the fields present in the request are changed; the rest are left as they
// are. Sending `instructions: null` (or empty) clears the instructions.
type AIMemoryNamespaceSettingPatchAllParamsSummary struct {
	// Replace the namespace's summary instructions. Null or empty clears them and
	// returns to the neutral default. Omit the field to leave the current instructions
	// unchanged.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r AIMemoryNamespaceSettingPatchAllParamsSummary) MarshalJSON() (data []byte, err error) {
	type shadow AIMemoryNamespaceSettingPatchAllParamsSummary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMemoryNamespaceSettingPatchAllParamsSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
