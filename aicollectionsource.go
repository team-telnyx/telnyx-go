// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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
// AICollectionSourceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAICollectionSourceService] method instead.
type AICollectionSourceService struct {
	Options []option.RequestOption
}

// NewAICollectionSourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAICollectionSourceService(opts ...option.RequestOption) (r AICollectionSourceService) {
	r = AICollectionSourceService{}
	r.Options = opts
	return
}

// Attaches a new content source to the specified collection and returns the
// created source. The source's content is ingested and embedded so it becomes
// searchable within the collection.
func (r *AICollectionSourceService) New(ctx context.Context, uuid string, body AICollectionSourceNewParams, opts ...option.RequestOption) (res *AICollectionSourceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/sources", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the sources attached to a collection.
func (r *AICollectionSourceService) List(ctx context.Context, uuid string, opts ...option.RequestOption) (res *AICollectionSourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/sources", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes a single source from a collection.
func (r *AICollectionSourceService) Delete(ctx context.Context, sourceID string, body AICollectionSourceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.Uuid == "" {
		err = errors.New("missing required uuid parameter")
		return err
	}
	if sourceID == "" {
		err = errors.New("missing required sourceId parameter")
		return err
	}
	path := fmt.Sprintf("ai/collections/%s/sources/%s", body.Uuid, sourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Replaces the collection's entire source set. The response `meta` reports which
// sources were added, retained, and removed.
func (r *AICollectionSourceService) Replace(ctx context.Context, uuid string, body AICollectionSourceReplaceParams, opts ...option.RequestOption) (res *AICollectionSourceReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/sources", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type Source struct {
	ID string `json:"id"`
	// The Telnyx Storage bucket name. Present only for `bucket` sources.
	BucketID     string `json:"bucket_id"`
	CollectionID string `json:"collection_id" format:"uuid"`
	// Identifies the record type. Always `ai_collection_source`.
	RecordType string `json:"record_type"`
	// The type of Telnyx data attached as a source. `bucket` requires an additional
	// `bucket_id`. Only `voice` is searchable today; `meeting_bot`, `message`, and
	// `bucket` attach but are not yet searchable (Coming soon).
	//
	// Any of "voice", "meeting_bot", "message", "bucket".
	SourceType SourceType `json:"source_type"`
	Status     string     `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		BucketID     respjson.Field
		CollectionID respjson.Field
		RecordType   respjson.Field
		SourceType   respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Source) RawJSON() string { return r.JSON.raw }
func (r *Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SourceType is required.
type SourceRequestParam struct {
	// The type of Telnyx data attached as a source. `bucket` requires an additional
	// `bucket_id`. Only `voice` is searchable today; `meeting_bot`, `message`, and
	// `bucket` attach but are not yet searchable (Coming soon).
	//
	// Any of "voice", "meeting_bot", "message", "bucket".
	SourceType SourceType `json:"source_type,omitzero" api:"required"`
	// The Telnyx Storage bucket name. Required when `source_type` is `bucket`; ignored
	// otherwise.
	BucketID param.Opt[string] `json:"bucket_id,omitzero"`
	paramObj
}

func (r SourceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of Telnyx data attached as a source. `bucket` requires an additional
// `bucket_id`. Only `voice` is searchable today; `meeting_bot`, `message`, and
// `bucket` attach but are not yet searchable (Coming soon).
type SourceType string

const (
	SourceTypeVoice      SourceType = "voice"
	SourceTypeMeetingBot SourceType = "meeting_bot"
	SourceTypeMessage    SourceType = "message"
	SourceTypeBucket     SourceType = "bucket"
)

// Envelope containing a single collection source.
type AICollectionSourceNewResponse struct {
	Data Source `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AICollectionSourceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AICollectionSourceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSourceListResponse struct {
	Data []Source `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AICollectionSourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AICollectionSourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSourceReplaceResponse struct {
	Data []Source `json:"data"`
	// Reports which source IDs were added, retained, and removed by a replace
	// operation.
	Meta AICollectionSourceReplaceResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AICollectionSourceReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *AICollectionSourceReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reports which source IDs were added, retained, and removed by a replace
// operation.
type AICollectionSourceReplaceResponseMeta struct {
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Retained []string `json:"retained"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Removed     respjson.Field
		Retained    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AICollectionSourceReplaceResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AICollectionSourceReplaceResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSourceNewParams struct {
	SourceRequest SourceRequestParam
	paramObj
}

func (r AICollectionSourceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SourceRequest)
}
func (r *AICollectionSourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionSourceDeleteParams struct {
	Uuid string `path:"uuid" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AICollectionSourceReplaceParams struct {
	Sources []SourceRequestParam `json:"sources,omitzero" api:"required"`
	paramObj
}

func (r AICollectionSourceReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow AICollectionSourceReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AICollectionSourceReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
