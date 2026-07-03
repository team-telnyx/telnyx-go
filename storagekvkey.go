// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Read and write keys within a KV namespace
//
// StorageKvKeyService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageKvKeyService] method instead.
type StorageKvKeyService struct {
	Options []option.RequestOption
}

// NewStorageKvKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageKvKeyService(opts ...option.RequestOption) (r StorageKvKeyService) {
	r = StorageKvKeyService{}
	r.Options = opts
	return
}

// Returns the raw stored value for a key. The response body is the value exactly
// as it was written; the `Content-Type` header echoes the value's stored content
// type (defaults to `application/octet-stream`).
func (r *StorageKvKeyService) Get(ctx context.Context, key string, query StorageKvKeyGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/kvs/%s/keys/%s", query.ID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Creates or replaces the value for a key. The request body is stored verbatim as
// the value — no base64, no JSON envelope — up to 1 MiB. The request's
// `Content-Type` header is stored with the value and echoed back on retrieval.
// Returns `201` when the key is created and `200` when an existing key is updated.
func (r *StorageKvKeyService) Update(ctx context.Context, key string, body io.Reader, params StorageKvKeyUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*"), option.WithRequestBody("application/octet-stream", body)}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return err
	}
	path := fmt.Sprintf("storage/kvs/%s/keys/%s", params.ID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return err
}

// Lists the keys in a namespace. Returns key names and metadata only, never
// values. Results are paginated with `limit` and an opaque `cursor`.
func (r *StorageKvKeyService) List(ctx context.Context, id string, query StorageKvKeyListParams, opts ...option.RequestOption) (res *pagination.CursorFlatPagination[StorageKvKeyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("storage/kvs/%s/keys", id)
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

// Lists the keys in a namespace. Returns key names and metadata only, never
// values. Results are paginated with `limit` and an opaque `cursor`.
func (r *StorageKvKeyService) ListAutoPaging(ctx context.Context, id string, query StorageKvKeyListParams, opts ...option.RequestOption) *pagination.CursorFlatPaginationAutoPager[StorageKvKeyListResponse] {
	return pagination.NewCursorFlatPaginationAutoPager(r.List(ctx, id, query, opts...))
}

// Deletes a key. Idempotent: deleting a key that does not exist still succeeds.
// The namespace itself must exist and be provisioned.
func (r *StorageKvKeyService) Delete(ctx context.Context, key string, body StorageKvKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return err
	}
	path := fmt.Sprintf("storage/kvs/%s/keys/%s", body.ID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type StorageKvKeyListResponse struct {
	Key string `json:"key"`
	// Size of the stored value in bytes.
	SizeBytes int64     `json:"size_bytes"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		SizeBytes   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageKvKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageKvKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageKvKeyGetParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type StorageKvKeyUpdateParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	// Time-to-live in seconds. When set, the key expires and is deleted after this
	// duration. Requires a namespace provisioned with TTL support; namespaces without
	// it return a `409`.
	TtlSecs param.Opt[int64] `query:"ttl_secs,omitzero" json:"-"`
	paramObj
}

func (r StorageKvKeyUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [StorageKvKeyUpdateParams]'s query parameters as
// `url.Values`.
func (r StorageKvKeyUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StorageKvKeyListParams struct {
	// Opaque pagination cursor from a previous response's `meta.cursor`.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of keys to return. Values above 1000 are treated as 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Return only keys that start with this prefix.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageKvKeyListParams]'s query parameters as `url.Values`.
func (r StorageKvKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StorageKvKeyDeleteParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	paramObj
}
