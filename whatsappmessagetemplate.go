// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// Manage Whatsapp message templates
//
// WhatsappMessageTemplateService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappMessageTemplateService] method instead.
type WhatsappMessageTemplateService struct {
	Options []option.RequestOption
}

// NewWhatsappMessageTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappMessageTemplateService(opts ...option.RequestOption) (r WhatsappMessageTemplateService) {
	r = WhatsappMessageTemplateService{}
	r.Options = opts
	return
}

// Create a Whatsapp message template
func (r *WhatsappMessageTemplateService) New(ctx context.Context, body WhatsappMessageTemplateNewParams, opts ...option.RequestOption) (res *WhatsappMessageTemplateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/whatsapp/message_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Whatsapp message templates
func (r *WhatsappMessageTemplateService) List(ctx context.Context, query WhatsappMessageTemplateListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappMessageTemplateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/whatsapp/message_templates"
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

// List Whatsapp message templates
func (r *WhatsappMessageTemplateService) ListAutoPaging(ctx context.Context, query WhatsappMessageTemplateListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappMessageTemplateListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type WhatsappMessageTemplateNewResponse struct {
	Data WhatsappMessageTemplateNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateNewResponseData struct {
	ID string `json:"id"`
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category string `json:"category"`
	// Whatsapp template components (header, body, footer, buttons)
	Components              []any                                                         `json:"components"`
	CreatedAt               time.Time                                                     `json:"created_at" format:"date-time"`
	Language                string                                                        `json:"language"`
	Name                    string                                                        `json:"name"`
	RecordType              string                                                        `json:"record_type"`
	RejectionReason         string                                                        `json:"rejection_reason"`
	Status                  string                                                        `json:"status"`
	TemplateID              string                                                        `json:"template_id"`
	UpdatedAt               time.Time                                                     `json:"updated_at" format:"date-time"`
	WhatsappBusinessAccount WhatsappMessageTemplateNewResponseDataWhatsappBusinessAccount `json:"whatsapp_business_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Category                respjson.Field
		Components              respjson.Field
		CreatedAt               respjson.Field
		Language                respjson.Field
		Name                    respjson.Field
		RecordType              respjson.Field
		RejectionReason         respjson.Field
		Status                  respjson.Field
		TemplateID              respjson.Field
		UpdatedAt               respjson.Field
		WhatsappBusinessAccount respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateNewResponseDataWhatsappBusinessAccount struct {
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateNewResponseDataWhatsappBusinessAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappMessageTemplateNewResponseDataWhatsappBusinessAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateListResponse struct {
	ID string `json:"id"`
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category WhatsappMessageTemplateListResponseCategory `json:"category"`
	// Whatsapp template components (header, body, footer, buttons)
	Components              []any                                                      `json:"components"`
	CreatedAt               time.Time                                                  `json:"created_at" format:"date-time"`
	Language                string                                                     `json:"language"`
	Name                    string                                                     `json:"name"`
	RecordType              string                                                     `json:"record_type"`
	RejectionReason         string                                                     `json:"rejection_reason"`
	Status                  string                                                     `json:"status"`
	TemplateID              string                                                     `json:"template_id"`
	UpdatedAt               time.Time                                                  `json:"updated_at" format:"date-time"`
	WhatsappBusinessAccount WhatsappMessageTemplateListResponseWhatsappBusinessAccount `json:"whatsapp_business_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Category                respjson.Field
		Components              respjson.Field
		CreatedAt               respjson.Field
		Language                respjson.Field
		Name                    respjson.Field
		RecordType              respjson.Field
		RejectionReason         respjson.Field
		Status                  respjson.Field
		TemplateID              respjson.Field
		UpdatedAt               respjson.Field
		WhatsappBusinessAccount respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateListResponseCategory string

const (
	WhatsappMessageTemplateListResponseCategoryMarketing      WhatsappMessageTemplateListResponseCategory = "MARKETING"
	WhatsappMessageTemplateListResponseCategoryUtility        WhatsappMessageTemplateListResponseCategory = "UTILITY"
	WhatsappMessageTemplateListResponseCategoryAuthentication WhatsappMessageTemplateListResponseCategory = "AUTHENTICATION"
)

type WhatsappMessageTemplateListResponseWhatsappBusinessAccount struct {
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateListResponseWhatsappBusinessAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappMessageTemplateListResponseWhatsappBusinessAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateNewParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category   WhatsappMessageTemplateNewParamsCategory `json:"category,omitzero" api:"required"`
	Components []any                                    `json:"components,omitzero" api:"required"`
	Language   string                                   `json:"language" api:"required"`
	Name       string                                   `json:"name" api:"required"`
	WabaID     string                                   `json:"waba_id" api:"required"`
	paramObj
}

func (r WhatsappMessageTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateNewParamsCategory string

const (
	WhatsappMessageTemplateNewParamsCategoryMarketing      WhatsappMessageTemplateNewParamsCategory = "MARKETING"
	WhatsappMessageTemplateNewParamsCategoryUtility        WhatsappMessageTemplateNewParamsCategory = "UTILITY"
	WhatsappMessageTemplateNewParamsCategoryAuthentication WhatsappMessageTemplateNewParamsCategory = "AUTHENTICATION"
)

type WhatsappMessageTemplateListParams struct {
	// Search templates by name
	FilterSearch param.Opt[string] `query:"filter[search],omitzero" json:"-"`
	// Filter by template status
	FilterStatus param.Opt[string] `query:"filter[status],omitzero" json:"-"`
	// Filter by WABA ID
	FilterWabaID param.Opt[string] `query:"filter[waba_id],omitzero" json:"-"`
	PageNumber   param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize     param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	// Filter by category
	//
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	FilterCategory WhatsappMessageTemplateListParamsFilterCategory `query:"filter[category],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappMessageTemplateListParams]'s query parameters as
// `url.Values`.
func (r WhatsappMessageTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by category
type WhatsappMessageTemplateListParamsFilterCategory string

const (
	WhatsappMessageTemplateListParamsFilterCategoryMarketing      WhatsappMessageTemplateListParamsFilterCategory = "MARKETING"
	WhatsappMessageTemplateListParamsFilterCategoryUtility        WhatsappMessageTemplateListParamsFilterCategory = "UTILITY"
	WhatsappMessageTemplateListParamsFilterCategoryAuthentication WhatsappMessageTemplateListParamsFilterCategory = "AUTHENTICATION"
)
