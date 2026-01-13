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

// FaxService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFaxService] method instead.
type FaxService struct {
	Options []option.RequestOption
	Actions FaxActionService
}

// NewFaxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFaxService(opts ...option.RequestOption) (r FaxService) {
	r = FaxService{}
	r.Options = opts
	r.Actions = NewFaxActionService(opts...)
	return
}

// Send a fax. Files have size limits and page count limit validations. If a file
// is bigger than 50MB or has more than 350 pages it will fail with
// `file_size_limit_exceeded` and `page_count_limit_exceeded` respectively.
//
// **Expected Webhooks:**
//
// - `fax.queued`
// - `fax.media.processed`
// - `fax.sending.started`
// - `fax.delivered`
// - `fax.failed`
func (r *FaxService) New(ctx context.Context, body FaxNewParams, opts ...option.RequestOption) (res *FaxNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "faxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View a fax
func (r *FaxService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FaxGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("faxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View a list of faxes
func (r *FaxService) List(ctx context.Context, query FaxListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Fax], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "faxes"
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

// View a list of faxes
func (r *FaxService) ListAutoPaging(ctx context.Context, query FaxListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Fax] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a fax
func (r *FaxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("faxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Fax struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 timestamp when resource was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction FaxDirection `json:"direction"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The string used as the caller id name (SIP From Display Name) presented to the
	// destination (`to` number).
	FromDisplayName string `json:"from_display_name"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The URL (or list of URLs) to the PDF used for the fax's media. media_url and
	// media_name/contents can't be submitted together.
	MediaURL string `json:"media_url"`
	// If `store_preview` was set to `true`, this is a link to temporary location. Link
	// expires after 10 minutes.
	PreviewURL string `json:"preview_url"`
	// The quality of the fax. The `ultra` settings provides the highest quality
	// available, but also present longer fax processing times. `ultra_light` is best
	// suited for images, wihle `ultra_dark` is best suited for text.
	//
	// Any of "normal", "high", "very_high", "ultra_light", "ultra_dark".
	Quality FaxQuality `json:"quality"`
	// Identifies the type of the resource.
	//
	// Any of "fax".
	RecordType FaxRecordType `json:"record_type"`
	// Status of the fax
	//
	// Any of "queued", "media.processed", "originated", "sending", "delivered",
	// "failed", "initiated", "receiving", "media.processing", "received".
	Status FaxStatus `json:"status"`
	// Should fax media be stored on temporary URL. It does not support media_name.
	StoreMedia bool `json:"store_media"`
	// If store_media was set to true, this is a link to temporary location. Link
	// expires after 10 minutes.
	StoredMediaURL string `json:"stored_media_url"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// ISO 8601 timestamp when resource was updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Optional failover URL that will receive fax webhooks if webhook_url doesn't
	// return a 2XX response
	WebhookFailoverURL string `json:"webhook_failover_url"`
	// URL that will receive fax webhooks
	WebhookURL string `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ClientState        respjson.Field
		ConnectionID       respjson.Field
		CreatedAt          respjson.Field
		Direction          respjson.Field
		From               respjson.Field
		FromDisplayName    respjson.Field
		MediaName          respjson.Field
		MediaURL           respjson.Field
		PreviewURL         respjson.Field
		Quality            respjson.Field
		RecordType         respjson.Field
		Status             respjson.Field
		StoreMedia         respjson.Field
		StoredMediaURL     respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Fax) RawJSON() string { return r.JSON.raw }
func (r *Fax) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the fax.
type FaxDirection string

const (
	FaxDirectionInbound  FaxDirection = "inbound"
	FaxDirectionOutbound FaxDirection = "outbound"
)

// The quality of the fax. The `ultra` settings provides the highest quality
// available, but also present longer fax processing times. `ultra_light` is best
// suited for images, wihle `ultra_dark` is best suited for text.
type FaxQuality string

const (
	FaxQualityNormal     FaxQuality = "normal"
	FaxQualityHigh       FaxQuality = "high"
	FaxQualityVeryHigh   FaxQuality = "very_high"
	FaxQualityUltraLight FaxQuality = "ultra_light"
	FaxQualityUltraDark  FaxQuality = "ultra_dark"
)

// Identifies the type of the resource.
type FaxRecordType string

const (
	FaxRecordTypeFax FaxRecordType = "fax"
)

// Status of the fax
type FaxStatus string

const (
	FaxStatusQueued          FaxStatus = "queued"
	FaxStatusMediaProcessed  FaxStatus = "media.processed"
	FaxStatusOriginated      FaxStatus = "originated"
	FaxStatusSending         FaxStatus = "sending"
	FaxStatusDelivered       FaxStatus = "delivered"
	FaxStatusFailed          FaxStatus = "failed"
	FaxStatusInitiated       FaxStatus = "initiated"
	FaxStatusReceiving       FaxStatus = "receiving"
	FaxStatusMediaProcessing FaxStatus = "media.processing"
	FaxStatusReceived        FaxStatus = "received"
)

type FaxNewResponse struct {
	Data Fax `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxGetResponse struct {
	Data Fax `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FaxGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxNewParams struct {
	// The connection ID to send the fax with.
	ConnectionID string `json:"connection_id,required"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from,required"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to,required"`
	// The black threshold percentage for monochrome faxes. Only applicable if
	// `monochrome` is set to `true`.
	BlackThreshold param.Opt[int64] `json:"black_threshold,omitzero"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// The `from_display_name` string to be used as the caller id name (SIP From
	// Display Name) presented to the destination (`to` number). The string should have
	// a maximum of 128 characters, containing only letters, numbers, spaces, and
	// -\_~!.+ special characters. If ommited, the display name will be the same as the
	// number in the `from` field.
	FromDisplayName param.Opt[string] `json:"from_display_name,omitzero"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// The URL (or list of URLs) to the PDF used for the fax's media. media_url and
	// media_name/contents can't be submitted together.
	MediaURL param.Opt[string] `json:"media_url,omitzero"`
	// The flag to enable monochrome, true black and white fax results.
	Monochrome param.Opt[bool] `json:"monochrome,omitzero"`
	// Should fax media be stored on temporary URL. It does not support media_name,
	// they can't be submitted together.
	StoreMedia param.Opt[bool] `json:"store_media,omitzero"`
	// Should fax preview be stored on temporary URL.
	StorePreview param.Opt[bool] `json:"store_preview,omitzero"`
	// The flag to disable the T.38 protocol.
	T38Enabled param.Opt[bool] `json:"t38_enabled,omitzero"`
	// Use this field to override the URL to which Telnyx will send subsequent webhooks
	// for this fax.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// The format for the preview file in case the `store_preview` is `true`.
	//
	// Any of "pdf", "tiff".
	PreviewFormat FaxNewParamsPreviewFormat `json:"preview_format,omitzero"`
	// The quality of the fax. The `ultra` settings provides the highest quality
	// available, but also present longer fax processing times. `ultra_light` is best
	// suited for images, wihle `ultra_dark` is best suited for text.
	//
	// Any of "normal", "high", "very_high", "ultra_light", "ultra_dark".
	Quality FaxNewParamsQuality `json:"quality,omitzero"`
	paramObj
}

func (r FaxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FaxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FaxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The format for the preview file in case the `store_preview` is `true`.
type FaxNewParamsPreviewFormat string

const (
	FaxNewParamsPreviewFormatPdf  FaxNewParamsPreviewFormat = "pdf"
	FaxNewParamsPreviewFormatTiff FaxNewParamsPreviewFormat = "tiff"
)

// The quality of the fax. The `ultra` settings provides the highest quality
// available, but also present longer fax processing times. `ultra_light` is best
// suited for images, wihle `ultra_dark` is best suited for text.
type FaxNewParamsQuality string

const (
	FaxNewParamsQualityNormal     FaxNewParamsQuality = "normal"
	FaxNewParamsQualityHigh       FaxNewParamsQuality = "high"
	FaxNewParamsQualityVeryHigh   FaxNewParamsQuality = "very_high"
	FaxNewParamsQualityUltraLight FaxNewParamsQuality = "ultra_light"
	FaxNewParamsQualityUltraDark  FaxNewParamsQuality = "ultra_dark"
)

type FaxListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[created_at][gte], filter[created_at][gt], filter[created_at][lte],
	// filter[created_at][lt], filter[direction][eq], filter[from][eq], filter[to][eq]
	Filter FaxListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParams]'s query parameters as `url.Values`.
func (r FaxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[created_at][gte], filter[created_at][gt], filter[created_at][lte],
// filter[created_at][lt], filter[direction][eq], filter[from][eq], filter[to][eq]
type FaxListParamsFilter struct {
	// Date range filtering operations for fax creation timestamp
	CreatedAt FaxListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	// Direction filtering operations
	Direction FaxListParamsFilterDirection `query:"direction,omitzero" json:"-"`
	// From number filtering operations
	From FaxListParamsFilterFrom `query:"from,omitzero" json:"-"`
	// To number filtering operations
	To FaxListParamsFilterTo `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParamsFilter]'s query parameters as `url.Values`.
func (r FaxListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Date range filtering operations for fax creation timestamp
type FaxListParamsFilterCreatedAt struct {
	// ISO 8601 date time for filtering faxes created after that date
	Gt param.Opt[time.Time] `query:"gt,omitzero" format:"date-time" json:"-"`
	// ISO 8601 date time for filtering faxes created after or on that date
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// ISO 8601 formatted date time for filtering faxes created before that date
	Lt param.Opt[time.Time] `query:"lt,omitzero" format:"date-time" json:"-"`
	// ISO 8601 formatted date time for filtering faxes created on or before that date
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParamsFilterCreatedAt]'s query parameters as
// `url.Values`.
func (r FaxListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction filtering operations
type FaxListParamsFilterDirection struct {
	// The direction, inbound or outbound, for filtering faxes sent from this account
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParamsFilterDirection]'s query parameters as
// `url.Values`.
func (r FaxListParamsFilterDirection) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// From number filtering operations
type FaxListParamsFilterFrom struct {
	// The phone number, in E.164 format for filtering faxes sent from this number
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParamsFilterFrom]'s query parameters as
// `url.Values`.
func (r FaxListParamsFilterFrom) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// To number filtering operations
type FaxListParamsFilterTo struct {
	// The phone number, in E.164 format for filtering faxes sent to this number
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FaxListParamsFilterTo]'s query parameters as `url.Values`.
func (r FaxListParamsFilterTo) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
