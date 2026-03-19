// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
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

// Get a Whatsapp message template by ID
func (r *WhatsappMessageTemplateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappMessageTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Whatsapp message template
func (r *WhatsappMessageTemplateService) Update(ctx context.Context, id string, body WhatsappMessageTemplateUpdateParams, opts ...option.RequestOption) (res *WhatsappMessageTemplateUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Whatsapp message templates
func (r *WhatsappMessageTemplateService) List(ctx context.Context, query WhatsappMessageTemplateListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.WhatsappTemplateData], err error) {
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
func (r *WhatsappMessageTemplateService) ListAutoPaging(ctx context.Context, query WhatsappMessageTemplateListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.WhatsappTemplateData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a Whatsapp message template
func (r *WhatsappMessageTemplateService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp_message_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WhatsappMessageTemplateNewResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
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

type WhatsappMessageTemplateGetResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateResponse struct {
	Data shared.WhatsappTemplateData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMessageTemplateUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMessageTemplateUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateNewParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category   WhatsappMessageTemplateNewParamsCategory `json:"category,omitzero" api:"required"`
	Components []map[string]any                         `json:"components,omitzero" api:"required"`
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

type WhatsappMessageTemplateUpdateParams struct {
	// Any of "MARKETING", "UTILITY", "AUTHENTICATION".
	Category   WhatsappMessageTemplateUpdateParamsCategory `json:"category,omitzero"`
	Components []map[string]any                            `json:"components,omitzero"`
	paramObj
}

func (r WhatsappMessageTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMessageTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMessageTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappMessageTemplateUpdateParamsCategory string

const (
	WhatsappMessageTemplateUpdateParamsCategoryMarketing      WhatsappMessageTemplateUpdateParamsCategory = "MARKETING"
	WhatsappMessageTemplateUpdateParamsCategoryUtility        WhatsappMessageTemplateUpdateParamsCategory = "UTILITY"
	WhatsappMessageTemplateUpdateParamsCategoryAuthentication WhatsappMessageTemplateUpdateParamsCategory = "AUTHENTICATION"
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
