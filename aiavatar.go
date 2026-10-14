// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Create and manage real-time AI avatars
//
// AIAvatarService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAvatarService] method instead.
type AIAvatarService struct {
	Options []option.RequestOption
}

// NewAIAvatarService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAvatarService(opts ...option.RequestOption) (r AIAvatarService) {
	r = AIAvatarService{}
	r.Options = opts
	return
}

// Uploads an image, validates its file type, size, safe decoding, and resolution,
// stores it in Telnyx-managed private storage, and creates an avatar owned by the
// authenticated organization. Creating an avatar does not update an AI Assistant;
// use the returned ID separately wherever the avatar is consumed. The avatar image
// is stored independently from the rendering model and background, which are
// selected where the avatar is consumed. Supported image types are JPEG, PNG, and
// WebP. Images may be up to 10 MiB, must be between 256 and 8192 pixels on each
// side, and may contain at most 40 megapixels. Each organization may have up to 20
// custom avatars. For best results, use a 16:9 landscape image, preferably 1920 by
// 1080 pixels and at least 1280 by 720 pixels. Show one centered, front-facing
// person from the shoulders up, keep the full head and hair visible with space
// around it, and use a sharp, evenly lit image with unobstructed eyes and mouth.
// Use a transparent PNG or WebP when the avatar will be combined with a selectable
// background. These composition recommendations are not enforced in the first
// release; the first release does not validate the number of faces, pose, or image
// content.
func (r *AIAvatarService) New(ctx context.Context, body AIAvatarNewParams, opts ...option.RequestOption) (res *AvatarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/avatars"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a Telnyx stock avatar or a custom avatar owned by the authenticated
// organization. A custom avatar belonging to another organization is returned as
// not found.
func (r *AIAvatarService) Get(ctx context.Context, avatarID AvatarID, opts ...option.RequestOption) (res *AvatarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/avatars/%v", avatarID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the customer-facing name of a custom avatar owned by the authenticated
// organization. Stock avatars cannot be modified. To replace the image, create a
// new custom avatar so existing references keep their original content.
func (r *AIAvatarService) Update(ctx context.Context, avatarID AvatarID, body AIAvatarUpdateParams, opts ...option.RequestOption) (res *AvatarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/avatars/%v", avatarID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of the Telnyx stock avatars and custom avatars
// available to the authenticated organization. Custom avatars owned by other
// organizations are never returned. The `image_url` value is suitable for
// displaying a thumbnail and may be refreshed by listing or retrieving the avatar
// again.
func (r *AIAvatarService) List(ctx context.Context, query AIAvatarListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Avatar], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/avatars"
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

// Returns a paginated list of the Telnyx stock avatars and custom avatars
// available to the authenticated organization. Custom avatars owned by other
// organizations are never returned. The `image_url` value is suitable for
// displaying a thumbnail and may be refreshed by listing or retrieving the avatar
// again.
func (r *AIAvatarService) ListAutoPaging(ctx context.Context, query AIAvatarListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Avatar] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type Avatar struct {
	// Stable identifier of a Telnyx stock avatar or an organization-owned custom
	// avatar. Custom avatar IDs are returned by the Avatars API and use the `avatar_`
	// prefix.
	ID AvatarID `json:"id" api:"required"`
	// Whether the avatar is provided by Telnyx or was uploaded by the authenticated
	// organization.
	//
	// Any of "stock", "custom".
	AvatarType AvatarType `json:"avatar_type" api:"required"`
	// Indicates that avatar rendering runs on Telnyx-operated infrastructure.
	Hosted bool `json:"hosted" api:"required"`
	// HTTPS image URL suitable for an avatar selection thumbnail. If the URL expires,
	// list or retrieve the avatar again to refresh it.
	ImageURL string `json:"image_url" api:"required" format:"uri"`
	// Customer-facing avatar name.
	Name string `json:"name" api:"required"`
	// Avatar provider. The first release supports Telnyx avatars.
	Provider constant.Telnyx `json:"provider" default:"telnyx"`
	// MIME type of a custom avatar image.
	//
	// Any of "image/jpeg", "image/png", "image/webp".
	ContentType AvatarContentType `json:"content_type"`
	// Time when the custom avatar was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Decoded height of a custom avatar image.
	HeightPx int64 `json:"height_px"`
	// Expiration time of `image_url`. Omitted when the URL does not expire.
	ImageURLExpiresAt time.Time `json:"image_url_expires_at" format:"date-time"`
	// Size of a custom avatar image in bytes.
	SizeBytes int64 `json:"size_bytes"`
	// Time when the custom avatar was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Decoded width of a custom avatar image.
	WidthPx int64 `json:"width_px"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AvatarType        respjson.Field
		Hosted            respjson.Field
		ImageURL          respjson.Field
		Name              respjson.Field
		Provider          respjson.Field
		ContentType       respjson.Field
		CreatedAt         respjson.Field
		HeightPx          respjson.Field
		ImageURLExpiresAt respjson.Field
		SizeBytes         respjson.Field
		UpdatedAt         respjson.Field
		WidthPx           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Avatar) RawJSON() string { return r.JSON.raw }
func (r *Avatar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MIME type of a custom avatar image.
type AvatarContentType string

const (
	AvatarContentTypeImageJpeg AvatarContentType = "image/jpeg"
	AvatarContentTypeImagePng  AvatarContentType = "image/png"
	AvatarContentTypeImageWebp AvatarContentType = "image/webp"
)

// Stable identifier of a Telnyx stock avatar or an organization-owned custom
// avatar. Custom avatar IDs are returned by the Avatars API and use the `avatar_`
// prefix.
type AvatarID string

const (
	AvatarIDElena   AvatarID = "elena"
	AvatarIDBen     AvatarID = "ben"
	AvatarIDJordan  AvatarID = "jordan"
	AvatarIDLin     AvatarID = "lin"
	AvatarIDAvery   AvatarID = "avery"
	AvatarIDRaphael AvatarID = "raphael"
)

type AvatarResponse struct {
	Data Avatar `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvatarResponse) RawJSON() string { return r.JSON.raw }
func (r *AvatarResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the avatar is provided by Telnyx or was uploaded by the authenticated
// organization.
type AvatarType string

const (
	AvatarTypeStock  AvatarType = "stock"
	AvatarTypeCustom AvatarType = "custom"
)

type AIAvatarNewParams struct {
	// JPEG, PNG, or WebP image file. The maximum file size is 10 MiB.
	ImageFile io.Reader `json:"image_file,omitzero" api:"required" format:"binary"`
	// Customer-facing name of the custom avatar.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r AIAvatarNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AIAvatarUpdateParams struct {
	// New customer-facing name of the custom avatar.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r AIAvatarUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAvatarUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAvatarUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAvatarListParams struct {
	// Return avatars whose names contain this value, using case-insensitive matching.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Page number to retrieve, starting at 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of avatars to return per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Return only stock or custom avatars.
	//
	// Any of "stock", "custom".
	FilterAvatarType AvatarType `query:"filter[avatar_type],omitzero" json:"-"`
	// Return the avatar with this exact ID.
	FilterID AvatarID `query:"filter[id],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAvatarListParams]'s query parameters as `url.Values`.
func (r AIAvatarListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
