// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Analyze voice AI sessions, costs, and event hierarchies across Telnyx products.
//
// SessionAnalysisMetadataService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionAnalysisMetadataService] method instead.
type SessionAnalysisMetadataService struct {
	Options []option.RequestOption
}

// NewSessionAnalysisMetadataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSessionAnalysisMetadataService(opts ...option.RequestOption) (r SessionAnalysisMetadataService) {
	r = SessionAnalysisMetadataService{}
	r.Options = opts
	return
}

// Returns all available record types and supported query parameters for session
// analysis.
func (r *SessionAnalysisMetadataService) Get(ctx context.Context, opts ...option.RequestOption) (res *SessionAnalysisMetadataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "session_analysis/metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns detailed metadata for a specific record type, including relationships
// and examples.
func (r *SessionAnalysisMetadataService) GetRecordType(ctx context.Context, recordType string, opts ...option.RequestOption) (res *SessionAnalysisMetadataGetRecordTypeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if recordType == "" {
		err = errors.New("missing required record_type parameter")
		return
	}
	path := fmt.Sprintf("session_analysis/metadata/%s", recordType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ChildRelationshipInfo struct {
	ChildEvent       string               `json:"child_event" api:"required"`
	ChildProduct     string               `json:"child_product" api:"required"`
	ChildRecordType  string               `json:"child_record_type" api:"required"`
	CostRollup       bool                 `json:"cost_rollup" api:"required"`
	Description      string               `json:"description" api:"required"`
	RelationshipType string               `json:"relationship_type" api:"required"`
	TraversalEnabled bool                 `json:"traversal_enabled" api:"required"`
	Via              MetadataFieldMapping `json:"via" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChildEvent       respjson.Field
		ChildProduct     respjson.Field
		ChildRecordType  respjson.Field
		CostRollup       respjson.Field
		Description      respjson.Field
		RelationshipType respjson.Field
		TraversalEnabled respjson.Field
		Via              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChildRelationshipInfo) RawJSON() string { return r.JSON.raw }
func (r *ChildRelationshipInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetadataFieldMapping struct {
	LocalField  string `json:"local_field" api:"required"`
	ParentField string `json:"parent_field" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LocalField  respjson.Field
		ParentField respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataFieldMapping) RawJSON() string { return r.JSON.raw }
func (r *MetadataFieldMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParentRelationshipInfo struct {
	CostRollup       bool                 `json:"cost_rollup" api:"required"`
	Description      string               `json:"description" api:"required"`
	ParentEvent      string               `json:"parent_event" api:"required"`
	ParentProduct    string               `json:"parent_product" api:"required"`
	ParentRecordType string               `json:"parent_record_type" api:"required"`
	RelationshipType string               `json:"relationship_type" api:"required"`
	TraversalEnabled bool                 `json:"traversal_enabled" api:"required"`
	Via              MetadataFieldMapping `json:"via" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CostRollup       respjson.Field
		Description      respjson.Field
		ParentEvent      respjson.Field
		ParentProduct    respjson.Field
		ParentRecordType respjson.Field
		RelationshipType respjson.Field
		TraversalEnabled respjson.Field
		Via              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParentRelationshipInfo) RawJSON() string { return r.JSON.raw }
func (r *ParentRelationshipInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetResponse struct {
	Meta SessionAnalysisMetadataGetResponseMeta `json:"meta" api:"required"`
	// Map of supported query parameter names to their definitions.
	QueryParameters map[string]SessionAnalysisMetadataGetResponseQueryParameter `json:"query_parameters" api:"required"`
	RecordTypes     []SessionAnalysisMetadataGetResponseRecordType              `json:"record_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Meta            respjson.Field
		QueryParameters respjson.Field
		RecordTypes     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetResponseMeta struct {
	LastUpdated      time.Time `json:"last_updated" api:"required" format:"date-time"`
	TotalRecordTypes int64     `json:"total_record_types" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdated      respjson.Field
		TotalRecordTypes respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetResponseQueryParameter struct {
	Default     string   `json:"default" api:"required"`
	Description string   `json:"description" api:"required"`
	Type        string   `json:"type" api:"required"`
	EnumValues  []string `json:"enum_values" api:"nullable"`
	Max         int64    `json:"max" api:"nullable"`
	Min         int64    `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Type        respjson.Field
		EnumValues  respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetResponseQueryParameter) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetResponseQueryParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetResponseRecordType struct {
	Aliases             []string                 `json:"aliases" api:"required"`
	ChildRelationships  []ChildRelationshipInfo  `json:"child_relationships" api:"required"`
	Description         string                   `json:"description" api:"required"`
	Event               string                   `json:"event" api:"required"`
	ParentRelationships []ParentRelationshipInfo `json:"parent_relationships" api:"required"`
	Product             string                   `json:"product" api:"required"`
	RecordType          string                   `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aliases             respjson.Field
		ChildRelationships  respjson.Field
		Description         respjson.Field
		Event               respjson.Field
		ParentRelationships respjson.Field
		Product             respjson.Field
		RecordType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetResponseRecordType) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetResponseRecordType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetRecordTypeResponse struct {
	Aliases            []string                `json:"aliases" api:"required"`
	ChildRelationships []ChildRelationshipInfo `json:"child_relationships" api:"required"`
	Event              string                  `json:"event" api:"required"`
	// Example queries and responses for this record type.
	Examples            map[string]any                                   `json:"examples" api:"required"`
	Meta                SessionAnalysisMetadataGetRecordTypeResponseMeta `json:"meta" api:"required"`
	ParentRelationships []ParentRelationshipInfo                         `json:"parent_relationships" api:"required"`
	Product             string                                           `json:"product" api:"required"`
	RecordType          string                                           `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aliases             respjson.Field
		ChildRelationships  respjson.Field
		Event               respjson.Field
		Examples            respjson.Field
		Meta                respjson.Field
		ParentRelationships respjson.Field
		Product             respjson.Field
		RecordType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetRecordTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetRecordTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionAnalysisMetadataGetRecordTypeResponseMeta struct {
	MaxRecommendedDepth int64 `json:"max_recommended_depth" api:"required"`
	TotalChildren       int64 `json:"total_children" api:"required"`
	TotalParents        int64 `json:"total_parents" api:"required"`
	TotalSiblings       int64 `json:"total_siblings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxRecommendedDepth respjson.Field
		TotalChildren       respjson.Field
		TotalParents        respjson.Field
		TotalSiblings       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionAnalysisMetadataGetRecordTypeResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SessionAnalysisMetadataGetRecordTypeResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
