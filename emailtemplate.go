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
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create, list, retrieve, update, delete, and render Liquid email templates.
//
// EmailTemplateService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailTemplateService] method instead.
type EmailTemplateService struct {
	Options []option.RequestOption
}

// NewEmailTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailTemplateService(opts ...option.RequestOption) (r EmailTemplateService) {
	r = EmailTemplateService{}
	r.Options = opts
	return
}

// Creates a Liquid email template. Variables are auto-extracted when omitted.
func (r *EmailTemplateService) New(ctx context.Context, params EmailTemplateNewParams, opts ...option.RequestOption) (res *EmailTemplateResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "email_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the account-owned template identified by ID, including its Liquid
// subject and bodies, declared variables, and timestamps.
func (r *EmailTemplateService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailTemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates one or more fields of the specified email template and returns the
// updated template.
func (r *EmailTemplateService) Update(ctx context.Context, id string, body EmailTemplateUpdateParams, opts ...option.RequestOption) (res *EmailTemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists templates sorted newest first by `created_at desc, id desc`.
func (r *EmailTemplateService) List(ctx context.Context, query EmailTemplateListParams, opts ...option.RequestOption) (res *pagination.EmailCursorPagination[EmailTemplate], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "email_templates"
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

// Lists templates sorted newest first by `created_at desc, id desc`.
func (r *EmailTemplateService) ListAutoPaging(ctx context.Context, query EmailTemplateListParams, opts ...option.RequestOption) *pagination.EmailCursorPaginationAutoPager[EmailTemplate] {
	return pagination.NewEmailCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes the account-owned template. The operation returns `204` with no body and
// prevents future sends or renders from using the deleted template ID.
func (r *EmailTemplateService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("email_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Renders a template using the provided Liquid variables. Missing
// `template_variables` defaults to `{}`.
//
// When the template has `strict_variables` enabled and a required variable (per
// `variable_schema`) is missing, returns 422 naming the variable. When the
// template has `autoescape` enabled, the rendered `html_body` expression output is
// HTML-escaped at the output boundary; `subject` and `text_body` are not
// autoescaped.
func (r *EmailTemplateService) Render(ctx context.Context, id string, body EmailTemplateRenderParams, opts ...option.RequestOption) (res *EmailTemplateRenderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_templates/%s/render", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Replaces template fields. Behaves identically to PATCH; provided for
// compatibility with Phoenix resource routes.
func (r *EmailTemplateService) Replace(ctx context.Context, id string, body EmailTemplateReplaceParams, opts ...option.RequestOption) (res *EmailTemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_templates/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type EmailTemplate struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether HTML autoescaping is enabled for this template. When `true`, only
	// rendered `html_body` expression output is HTML-escaped at the output boundary;
	// `subject` and `text_body` are never autoescaped.
	Autoescape bool      `json:"autoescape" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	HTMLBody   string    `json:"html_body" api:"required"`
	Name       string    `json:"name" api:"required"`
	// Any of "email_template".
	RecordType EmailTemplateRecordType `json:"record_type" api:"required"`
	// Whether strict variable validation is enabled for this template. When `true`,
	// sends and renders that are missing a variable marked `required: true` in
	// `variable_schema` fail with 422 naming the variable.
	StrictVariables bool      `json:"strict_variables" api:"required"`
	Subject         string    `json:"subject" api:"required"`
	TextBody        string    `json:"text_body" api:"required"`
	UpdatedAt       time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Structured variable requirements, or `null` when the template uses only the
	// legacy `variables` array.
	VariableSchema map[string]EmailTemplateVariableSchema `json:"variable_schema" api:"required"`
	// Legacy unstructured variable names. This path remains supported unchanged.
	Variables []string `json:"variables" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Autoescape      respjson.Field
		CreatedAt       respjson.Field
		HTMLBody        respjson.Field
		Name            respjson.Field
		RecordType      respjson.Field
		StrictVariables respjson.Field
		Subject         respjson.Field
		TextBody        respjson.Field
		UpdatedAt       respjson.Field
		VariableSchema  respjson.Field
		Variables       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplate) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateRecordType string

const (
	EmailTemplateRecordTypeEmailTemplate EmailTemplateRecordType = "email_template"
)

type EmailTemplateVariableSchema struct {
	// Whether the variable must be supplied when strict variable validation is
	// enabled.
	Required bool `json:"required" api:"required"`
	// Default value for an optional variable. Rejected when `required` is `true`.
	Default string `json:"default"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Required    respjson.Field
		Default     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVariableSchema) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVariableSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateResponse struct {
	Data EmailTemplate `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateEmailTemplateRequestParam struct {
	// Liquid template HTML body.
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	// Liquid template subject.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Liquid template text body.
	TextBody param.Opt[string] `json:"text_body,omitzero"`
	// Per-template HTML autoescaping setting.
	Autoescape param.Opt[bool]   `json:"autoescape,omitzero"`
	Name       param.Opt[string] `json:"name,omitzero"`
	// Per-template strict variable-validation setting.
	StrictVariables param.Opt[bool] `json:"strict_variables,omitzero"`
	// Structured variable requirements. Required variables cannot define defaults;
	// invalid combinations return 422. Set to `null` to clear the schema.
	VariableSchema map[string]UpdateEmailTemplateRequestVariableSchemaParam `json:"variable_schema,omitzero"`
	Variables      []string                                                 `json:"variables,omitzero"`
	paramObj
}

func (r UpdateEmailTemplateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateEmailTemplateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateEmailTemplateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Required is required.
type UpdateEmailTemplateRequestVariableSchemaParam struct {
	// Whether the variable must be supplied when strict variable validation is
	// enabled.
	Required bool `json:"required" api:"required"`
	// Default value for an optional variable. Rejected when `required` is `true`.
	Default param.Opt[string] `json:"default,omitzero"`
	paramObj
}

func (r UpdateEmailTemplateRequestVariableSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateEmailTemplateRequestVariableSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateEmailTemplateRequestVariableSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateRenderResponse struct {
	// Template object with `subject`, `html_body`, and `text_body` replaced by their
	// Liquid-rendered values. All other template fields (id, name, variables, etc.)
	// remain unchanged.
	Data EmailTemplateRenderResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateRenderResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateRenderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Template object with `subject`, `html_body`, and `text_body` replaced by their
// Liquid-rendered values. All other template fields (id, name, variables, etc.)
// remain unchanged.
type EmailTemplateRenderResponseData struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	EmailTemplate
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateRenderResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateRenderResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateNewParams struct {
	// Letters, numbers, spaces, hyphens, and underscores only.
	Name string `json:"name" api:"required"`
	// Liquid template HTML body.
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	// Liquid template subject.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Liquid template text body.
	TextBody param.Opt[string] `json:"text_body,omitzero"`
	// Per-template HTML autoescaping setting. Defaults to `false` for backward
	// compatibility. When `true`, the rendered `html_body` HTML-escapes each Liquid
	// expression's output at the output boundary (after its filters run, before
	// concatenation with literal template markup). Input values are never mutated and
	// `subject`/`text_body` are never autoescaped. The boundary escape is idempotent:
	// HTML entities already present in the output (e.g. from an explicit `escape`
	// filter) are preserved, so an explicit `escape`/`escape_once` is never
	// double-escaped, and markup introduced by any later filter in the chain is still
	// escaped.
	Autoescape param.Opt[bool] `json:"autoescape,omitzero"`
	// Per-template strict variable-validation setting. Defaults to `false` for
	// backward compatibility. When `true`, a send or render that is missing a variable
	// marked `required: true` in `variable_schema` fails with 422 naming the variable.
	// Missing optional variables never fail; their schema `default` (when set) is
	// applied to the render.
	StrictVariables param.Opt[bool]   `json:"strict_variables,omitzero"`
	IdempotencyKey  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	// Structured variable requirements. Required variables cannot define defaults;
	// invalid combinations return 422. This is independent of the legacy `variables`
	// array. On render with `strict_variables` enabled: `required` variables must be
	// supplied as non-empty values — absent, `null`, empty string, empty object `{}`,
	// and empty array `[]` all fail with 422 naming the variable, while present values
	// such as `false` and `0` pass (they are present, not empty). Optional variables
	// fall back to their `default` when absent.
	VariableSchema map[string]EmailTemplateNewParamsVariableSchema `json:"variable_schema,omitzero"`
	// Template variables. Auto-extracted from subject/body fields when absent.
	Variables []string `json:"variables,omitzero"`
	paramObj
}

func (r EmailTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Required is required.
type EmailTemplateNewParamsVariableSchema struct {
	// Whether the variable must be supplied when strict variable validation is
	// enabled.
	Required bool `json:"required" api:"required"`
	// Default value for an optional variable. Rejected when `required` is `true`.
	Default param.Opt[string] `json:"default,omitzero"`
	paramObj
}

func (r EmailTemplateNewParamsVariableSchema) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateNewParamsVariableSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateNewParamsVariableSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateUpdateParams struct {
	UpdateEmailTemplateRequest UpdateEmailTemplateRequestParam
	paramObj
}

func (r EmailTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateEmailTemplateRequest)
}
func (r *EmailTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateListParams struct {
	// Opaque URL-safe Base64 cursor returned by a previous list response.
	PageCursor param.Opt[string] `query:"page_cursor,omitzero" json:"-"`
	// Number of results to return. Defaults to 25; maximum is 100. Invalid values are
	// clamped to the valid range.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailTemplateListParams]'s query parameters as
// `url.Values`.
func (r EmailTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailTemplateRenderParams struct {
	// Variables for Liquid template rendering. Non-object values are silently treated
	// as an empty object.
	TemplateVariables map[string]any `json:"template_variables,omitzero"`
	paramObj
}

func (r EmailTemplateRenderParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateRenderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateRenderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailTemplateReplaceParams struct {
	UpdateEmailTemplateRequest UpdateEmailTemplateRequestParam
	paramObj
}

func (r EmailTemplateReplaceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateEmailTemplateRequest)
}
func (r *EmailTemplateReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
