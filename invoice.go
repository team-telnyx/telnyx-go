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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r InvoiceService) {
	r = InvoiceService{}
	r.Options = opts
	return
}

// Retrieve a single invoice by its unique identifier.
func (r *InvoiceService) Get(ctx context.Context, id string, query InvoiceGetParams, opts ...option.RequestOption) (res *InvoiceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of invoices.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *InvoiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InvoiceGetResponse struct {
	Data InvoiceGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetResponseData struct {
	// Present only if the query parameter `action=link` is set.
	DownloadURL string    `json:"download_url" format:"uri"`
	FileID      string    `json:"file_id" format:"uuid"`
	InvoiceID   string    `json:"invoice_id" format:"uuid"`
	Paid        bool      `json:"paid"`
	PeriodEnd   time.Time `json:"period_end" format:"date"`
	PeriodStart time.Time `json:"period_start" format:"date"`
	URL         string    `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL respjson.Field
		FileID      respjson.Field
		InvoiceID   respjson.Field
		Paid        respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvoiceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponse struct {
	Data []InvoiceListResponseData `json:"data"`
	Meta InvoiceListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponseData struct {
	FileID      string    `json:"file_id" format:"uuid"`
	InvoiceID   string    `json:"invoice_id" format:"uuid"`
	Paid        bool      `json:"paid"`
	PeriodEnd   time.Time `json:"period_end" format:"date"`
	PeriodStart time.Time `json:"period_start" format:"date"`
	URL         string    `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		InvoiceID   respjson.Field
		Paid        respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *InvoiceListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceGetParams struct {
	// Invoice action
	//
	// Any of "json", "link".
	Action InvoiceGetParamsAction `query:"action,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceGetParams]'s query parameters as `url.Values`.
func (r InvoiceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Invoice action
type InvoiceGetParamsAction string

const (
	InvoiceGetParamsActionJson InvoiceGetParamsAction = "json"
	InvoiceGetParamsActionLink InvoiceGetParamsAction = "link"
)

type InvoiceListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page InvoiceListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results.
	//
	// Any of "period_start", "-period_start".
	Sort InvoiceListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type InvoiceListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceListParamsPage]'s query parameters as `url.Values`.
func (r InvoiceListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results.
type InvoiceListParamsSort string

const (
	InvoiceListParamsSortPeriodStart     InvoiceListParamsSort = "period_start"
	InvoiceListParamsSortPeriodStartDesc InvoiceListParamsSort = "-period_start"
)
