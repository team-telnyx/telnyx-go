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

// RequirementService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequirementService] method instead.
type RequirementService struct {
	Options []option.RequestOption
}

// NewRequirementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequirementService(opts ...option.RequestOption) (r RequirementService) {
	r = RequirementService{}
	r.Options = opts
	return
}

// Retrieve a document requirement record
func (r *RequirementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RequirementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all requirements with filtering, sorting, and pagination
func (r *RequirementService) List(ctx context.Context, query RequirementListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[RequirementListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "requirements"
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

// List all requirements with filtering, sorting, and pagination
func (r *RequirementService) ListAutoPaging(ctx context.Context, query RequirementListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[RequirementListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

type RequirementGetResponse struct {
	Data RequirementGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RequirementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGetResponseData struct {
	// Identifies the associated document
	ID string `json:"id" format:"uuid"`
	// Indicates whether this requirement applies to branded_calling, ordering,
	// porting, or both ordering and porting
	//
	// Any of "both", "branded_calling", "ordering", "porting".
	Action string `json:"action"`
	// The 2-character (ISO 3166-1 alpha-2) country code where this requirement applies
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The locality where this requirement applies
	Locality string `json:"locality"`
	// Indicates the phone_number_type this requirement applies to. Leave blank if this
	// requirement applies to all number_types.
	//
	// Any of "local", "national", "toll_free".
	PhoneNumberType string `json:"phone_number_type"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Lists the requirement types necessary to fulfill this requirement
	RequirementsTypes []shared.DocReqsRequirementType `json:"requirements_types"`
	// ISO 8601 formatted date-time indicating when the resource was last updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Action            respjson.Field
		CountryCode       respjson.Field
		CreatedAt         respjson.Field
		Locality          respjson.Field
		PhoneNumberType   respjson.Field
		RecordType        respjson.Field
		RequirementsTypes respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *RequirementGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementListResponse struct {
	// Identifies the associated document
	ID string `json:"id" format:"uuid"`
	// Indicates whether this requirement applies to branded_calling, ordering,
	// porting, or both ordering and porting
	//
	// Any of "both", "branded_calling", "ordering", "porting".
	Action RequirementListResponseAction `json:"action"`
	// The 2-character (ISO 3166-1 alpha-2) country code where this requirement applies
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The locality where this requirement applies
	Locality string `json:"locality"`
	// Indicates the phone_number_type this requirement applies to. Leave blank if this
	// requirement applies to all number_types.
	//
	// Any of "local", "national", "toll_free".
	PhoneNumberType RequirementListResponsePhoneNumberType `json:"phone_number_type"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Lists the requirement types necessary to fulfill this requirement
	RequirementsTypes []shared.DocReqsRequirementType `json:"requirements_types"`
	// ISO 8601 formatted date-time indicating when the resource was last updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Action            respjson.Field
		CountryCode       respjson.Field
		CreatedAt         respjson.Field
		Locality          respjson.Field
		PhoneNumberType   respjson.Field
		RecordType        respjson.Field
		RequirementsTypes respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementListResponse) RawJSON() string { return r.JSON.raw }
func (r *RequirementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether this requirement applies to branded_calling, ordering,
// porting, or both ordering and porting
type RequirementListResponseAction string

const (
	RequirementListResponseActionBoth           RequirementListResponseAction = "both"
	RequirementListResponseActionBrandedCalling RequirementListResponseAction = "branded_calling"
	RequirementListResponseActionOrdering       RequirementListResponseAction = "ordering"
	RequirementListResponseActionPorting        RequirementListResponseAction = "porting"
)

// Indicates the phone_number_type this requirement applies to. Leave blank if this
// requirement applies to all number_types.
type RequirementListResponsePhoneNumberType string

const (
	RequirementListResponsePhoneNumberTypeLocal    RequirementListResponsePhoneNumberType = "local"
	RequirementListResponsePhoneNumberTypeNational RequirementListResponsePhoneNumberType = "national"
	RequirementListResponsePhoneNumberTypeTollFree RequirementListResponsePhoneNumberType = "toll_free"
)

type RequirementListParams struct {
	// Consolidated filter parameter for requirements (deepObject style). Originally:
	// filter[country_code], filter[phone_number_type], filter[action]
	Filter RequirementListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page RequirementListParamsPage `query:"page,omitzero" json:"-"`
	// Consolidated sort parameter for requirements (deepObject style). Originally:
	// sort[]
	//
	// Any of "created_at", "updated_at", "country_code", "phone_number_type",
	// "-created_at", "-updated_at", "-country_code", "-phone_number_type".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementListParams]'s query parameters as `url.Values`.
func (r RequirementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for requirements (deepObject style). Originally:
// filter[country_code], filter[phone_number_type], filter[action]
type RequirementListParamsFilter struct {
	// Filters results to those applying to a 2-character (ISO 3166-1 alpha-2) country
	// code
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filters requirements to those applying to a specific action.
	//
	// Any of "branded_calling", "ordering", "porting".
	Action string `query:"action,omitzero" json:"-"`
	// Filters results to those applying to a specific phone_number_type
	//
	// Any of "local", "national", "toll_free".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementListParamsFilter]'s query parameters as
// `url.Values`.
func (r RequirementListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type RequirementListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementListParamsPage]'s query parameters as
// `url.Values`.
func (r RequirementListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
