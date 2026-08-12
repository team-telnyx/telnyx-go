// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and manage logical collections of your Telnyx data, tune retrieval
// settings, manage sources, and run collection-scoped semantic search.
//
// AICollectionSettingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAICollectionSettingService] method instead.
type AICollectionSettingService struct {
	Options []option.RequestOption
}

// NewAICollectionSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAICollectionSettingService(opts ...option.RequestOption) (r AICollectionSettingService) {
	r = AICollectionSettingService{}
	r.Options = opts
	return
}

// Replaces the collection's retrieval settings.
func (r *AICollectionSettingService) New(ctx context.Context, uuid string, body AICollectionSettingNewParams, opts ...option.RequestOption) (res *SettingsEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/settings", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the retrieval settings for a collection.
func (r *AICollectionSettingService) List(ctx context.Context, uuid string, opts ...option.RequestOption) (res *SettingsEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/settings", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially updates the collection's retrieval settings.
func (r *AICollectionSettingService) PatchAll(ctx context.Context, uuid string, body AICollectionSettingPatchAllParams, opts ...option.RequestOption) (res *SettingsEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/settings", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// How documents are retrieved when searching the collection.
type RetrievalSettings struct {
	// Retrieval strategy. `vector` runs semantic similarity search; `hybrid` combines
	// vector similarity with keyword matching; `keyword` runs lexical (BM25) matching.
	//
	// Any of "vector", "hybrid", "keyword".
	RetrievalType RetrievalSettingsRetrievalType `json:"retrieval_type"`
	// Number of top results to retrieve (1–50).
	TopK int64 `json:"top_k"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RetrievalType respjson.Field
		TopK          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalSettings) RawJSON() string { return r.JSON.raw }
func (r *RetrievalSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RetrievalSettings to a RetrievalSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RetrievalSettingsParam.Overrides()
func (r RetrievalSettings) ToParam() RetrievalSettingsParam {
	return param.Override[RetrievalSettingsParam](json.RawMessage(r.RawJSON()))
}

// Retrieval strategy. `vector` runs semantic similarity search; `hybrid` combines
// vector similarity with keyword matching; `keyword` runs lexical (BM25) matching.
type RetrievalSettingsRetrievalType string

const (
	RetrievalSettingsRetrievalTypeVector  RetrievalSettingsRetrievalType = "vector"
	RetrievalSettingsRetrievalTypeHybrid  RetrievalSettingsRetrievalType = "hybrid"
	RetrievalSettingsRetrievalTypeKeyword RetrievalSettingsRetrievalType = "keyword"
)

// How documents are retrieved when searching the collection.
type RetrievalSettingsParam struct {
	// Number of top results to retrieve (1–50).
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Retrieval strategy. `vector` runs semantic similarity search; `hybrid` combines
	// vector similarity with keyword matching; `keyword` runs lexical (BM25) matching.
	//
	// Any of "vector", "hybrid", "keyword".
	RetrievalType RetrievalSettingsRetrievalType `json:"retrieval_type,omitzero"`
	paramObj
}

func (r RetrievalSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RetrievalSettingsWrapper struct {
	// Identifies the record type. Always `ai_collection_settings`.
	RecordType string `json:"record_type"`
	// How documents are retrieved when searching the collection.
	Retrieval RetrievalSettings `json:"retrieval"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordType  respjson.Field
		Retrieval   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalSettingsWrapper) RawJSON() string { return r.JSON.raw }
func (r *RetrievalSettingsWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RetrievalSettingsWrapper to a
// RetrievalSettingsWrapperParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RetrievalSettingsWrapperParam.Overrides()
func (r RetrievalSettingsWrapper) ToParam() RetrievalSettingsWrapperParam {
	return param.Override[RetrievalSettingsWrapperParam](json.RawMessage(r.RawJSON()))
}

type RetrievalSettingsWrapperParam struct {
	// How documents are retrieved when searching the collection.
	Retrieval RetrievalSettingsParam `json:"retrieval,omitzero"`
	paramObj
}

func (r RetrievalSettingsWrapperParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalSettingsWrapperParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalSettingsWrapperParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingsEnvelope struct {
	Data RetrievalSettingsWrapper `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SettingsEnvelope) RawJSON() string { return r.JSON.raw }
func (r *SettingsEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingsRequestParam struct {
	// How documents are retrieved when searching the collection.
	Retrieval RetrievalSettingsParam `json:"retrieval,omitzero"`
	paramObj
}

func (r SettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSettingNewParams struct {
	SettingsRequest SettingsRequestParam
	paramObj
}

func (r AICollectionSettingNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsRequest)
}
func (r *AICollectionSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSettingPatchAllParams struct {
	SettingsRequest SettingsRequestParam
	paramObj
}

func (r AICollectionSettingPatchAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SettingsRequest)
}
func (r *AICollectionSettingPatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
