// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// MessagingURLDomainService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingURLDomainService] method instead.
type MessagingURLDomainService struct {
	Options []option.RequestOption
}

// NewMessagingURLDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingURLDomainService(opts ...option.RequestOption) (r MessagingURLDomainService) {
	r = MessagingURLDomainService{}
	r.Options = opts
	return
}

// List messaging URL domains
func (r *MessagingURLDomainService) List(ctx context.Context, query MessagingURLDomainListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[MessagingURLDomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging_url_domains"
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

// List messaging URL domains
func (r *MessagingURLDomainService) ListAutoPaging(ctx context.Context, query MessagingURLDomainListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[MessagingURLDomainListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

type MessagingURLDomainListResponse struct {
	ID         string `json:"id"`
	RecordType string `json:"record_type"`
	URLDomain  string `json:"url_domain"`
	UseCase    string `json:"use_case"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RecordType  respjson.Field
		URLDomain   respjson.Field
		UseCase     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingURLDomainListResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingURLDomainListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingURLDomainListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MessagingURLDomainListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingURLDomainListParams]'s query parameters as
// `url.Values`.
func (r MessagingURLDomainListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MessagingURLDomainListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingURLDomainListParamsPage]'s query parameters as
// `url.Values`.
func (r MessagingURLDomainListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
