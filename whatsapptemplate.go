// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Manage Whatsapp message templates
//
// WhatsappTemplateService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappTemplateService] method instead.
type WhatsappTemplateService struct {
	Options []option.RequestOption
}

// NewWhatsappTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappTemplateService(opts ...option.RequestOption) (r WhatsappTemplateService) {
	r = WhatsappTemplateService{}
	r.Options = opts
	return
}

// Create a Whatsapp message template
func (r *WhatsappTemplateService) New(ctx context.Context, body WhatsappTemplateNewParams, opts ...option.RequestOption) (res *WhatsappTemplateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/whatsapp/message_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Whatsapp message templates
func (r *WhatsappTemplateService) List(ctx context.Context, query WhatsappTemplateListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.WhatsappTemplateData], err error) {
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
func (r *WhatsappTemplateService) ListAutoPaging(ctx context.Context, query WhatsappTemplateListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.WhatsappTemplateData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type WhatsappTemplateNewResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappTemplateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappTemplateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappTemplateNewParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category   WhatsappTemplateNewParamsCategory `json:"category,omitzero" api:"required"`
	Components []any                             `json:"components,omitzero" api:"required"`
	Language   string                            `json:"language" api:"required"`
	Name       string                            `json:"name" api:"required"`
	WabaID     string                            `json:"waba_id" api:"required"`
	paramObj
}

func (r WhatsappTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappTemplateNewParamsCategory string

const (
	WhatsappTemplateNewParamsCategoryMarketing      WhatsappTemplateNewParamsCategory = "MARKETING"
	WhatsappTemplateNewParamsCategoryUtility        WhatsappTemplateNewParamsCategory = "UTILITY"
	WhatsappTemplateNewParamsCategoryAuthentication WhatsappTemplateNewParamsCategory = "AUTHENTICATION"
)

type WhatsappTemplateListParams struct {
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
	FilterCategory WhatsappTemplateListParamsFilterCategory `query:"filter[category],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappTemplateListParams]'s query parameters as
// `url.Values`.
func (r WhatsappTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by category
type WhatsappTemplateListParamsFilterCategory string

const (
	WhatsappTemplateListParamsFilterCategoryMarketing      WhatsappTemplateListParamsFilterCategory = "MARKETING"
	WhatsappTemplateListParamsFilterCategoryUtility        WhatsappTemplateListParamsFilterCategory = "UTILITY"
	WhatsappTemplateListParamsFilterCategoryAuthentication WhatsappTemplateListParamsFilterCategory = "AUTHENTICATION"
)
