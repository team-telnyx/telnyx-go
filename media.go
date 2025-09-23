// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// MediaService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaService] method instead.
type MediaService struct {
	Options []option.RequestOption
}

// NewMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaService(opts ...option.RequestOption) (r MediaService) {
	r = MediaService{}
	r.Options = opts
	return
}

// Returns the information about a stored media file.
func (r *MediaService) Get(ctx context.Context, mediaName string, opts ...option.RequestOption) (res *MediaGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if mediaName == "" {
		err = errors.New("missing required media_name parameter")
		return
	}
	path := fmt.Sprintf("media/%s", mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a stored media file.
func (r *MediaService) Update(ctx context.Context, mediaName string, body MediaUpdateParams, opts ...option.RequestOption) (res *MediaUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if mediaName == "" {
		err = errors.New("missing required media_name parameter")
		return
	}
	path := fmt.Sprintf("media/%s", mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of stored media files.
func (r *MediaService) List(ctx context.Context, query MediaListParams, opts ...option.RequestOption) (res *MediaListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a stored media file.
func (r *MediaService) Delete(ctx context.Context, mediaName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if mediaName == "" {
		err = errors.New("missing required media_name parameter")
		return
	}
	path := fmt.Sprintf("media/%s", mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Downloads a stored media file.
func (r *MediaService) Download(ctx context.Context, mediaName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if mediaName == "" {
		err = errors.New("missing required media_name parameter")
		return
	}
	path := fmt.Sprintf("media/%s/download", mediaName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload media file to Telnyx so it can be used with other Telnyx services
func (r *MediaService) Upload(ctx context.Context, body MediaUploadParams, opts ...option.RequestOption) (res *MediaUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MediaResource struct {
	// Content type of the file
	ContentType string `json:"content_type"`
	// ISO 8601 formatted date of when the media resource was created
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date of when the media resource will expire and be deleted.
	ExpiresAt string `json:"expires_at"`
	// Uniquely identifies a media resource.
	MediaName string `json:"media_name"`
	// ISO 8601 formatted date of when the media resource was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		MediaName   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaResource) RawJSON() string { return r.JSON.raw }
func (r *MediaResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaGetResponse struct {
	Data MediaResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaUpdateResponse struct {
	Data MediaResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaListResponse struct {
	Data []MediaResource `json:"data"`
	Meta PaginationMeta  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaListResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaUploadResponse struct {
	Data MediaResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaUpdateParams struct {
	// The URL where the media to be stored in Telnyx network is currently hosted. The
	// maximum allowed size is 20 MB.
	MediaURL param.Opt[string] `json:"media_url,omitzero"`
	// The number of seconds after which the media resource will be deleted, defaults
	// to 2 days. The maximum allowed vale is 630720000, which translates to 20 years.
	TtlSecs param.Opt[int64] `json:"ttl_secs,omitzero"`
	paramObj
}

func (r MediaUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[content_type][]
	Filter MediaListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaListParams]'s query parameters as `url.Values`.
func (r MediaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[content_type][]
type MediaListParamsFilter struct {
	// Filters files by given content types
	ContentType []string `query:"content_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MediaListParamsFilter]'s query parameters as `url.Values`.
func (r MediaListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MediaUploadParams struct {
	// The URL where the media to be stored in Telnyx network is currently hosted. The
	// maximum allowed size is 20 MB.
	MediaURL string `json:"media_url,required"`
	// The unique identifier of a file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// The number of seconds after which the media resource will be deleted, defaults
	// to 2 days. The maximum allowed vale is 630720000, which translates to 20 years.
	TtlSecs param.Opt[int64] `json:"ttl_secs,omitzero"`
	paramObj
}

func (r MediaUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
