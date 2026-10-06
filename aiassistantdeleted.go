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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Configure AI assistant specifications
//
// AIAssistantDeletedService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantDeletedService] method instead.
type AIAssistantDeletedService struct {
	Options []option.RequestOption
}

// NewAIAssistantDeletedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantDeletedService(opts ...option.RequestOption) (r AIAssistantDeletedService) {
	r = AIAssistantDeletedService{}
	r.Options = opts
	return
}

// List the organization's soft-deleted assistants in the Recently Deleted list.
//
// Each entry includes `deleted_at` and `permanently_deleted_at`, the point after
// which the assistant is erased automatically and can no longer be restored.
func (r *AIAssistantDeletedService) List(ctx context.Context, query AIAssistantDeletedListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[DeletedAssistant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/assistants/deleted"
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

// List the organization's soft-deleted assistants in the Recently Deleted list.
//
// Each entry includes `deleted_at` and `permanently_deleted_at`, the point after
// which the assistant is erased automatically and can no longer be restored.
func (r *AIAssistantDeletedService) ListAutoPaging(ctx context.Context, query AIAssistantDeletedListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[DeletedAssistant] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a soft-deleted assistant from the Recently Deleted list by
// `assistant_id`, including its `deleted_at` and `permanently_deleted_at`
// timestamps. This is a read-only view; the assistant cannot be modified while it
// remains deleted.
func (r *AIAssistantDeletedService) Get(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *DeletedAssistant, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/deleted", url.PathEscape(assistantID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A soft-deleted assistant in the Recently Deleted list: the full assistant
// configuration plus deletion metadata.
type DeletedAssistant struct {
	// Timestamp of the soft delete.
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// Point after which the assistant is permanently deleted automatically and can no
	// longer be restored.
	PermanentlyDeletedAt time.Time `json:"permanently_deleted_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletedAt            respjson.Field
		PermanentlyDeletedAt respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
	InferenceEmbedding
}

// Returns the unmodified JSON received from the API
func (r DeletedAssistant) RawJSON() string { return r.JSON.raw }
func (r *DeletedAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantDeletedListParams struct {
	// Page number to retrieve (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items to return per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantDeletedListParams]'s query parameters as
// `url.Values`.
func (r AIAssistantDeletedListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
