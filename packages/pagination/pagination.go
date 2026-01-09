// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type DefaultPaginationMeta struct {
	PageNumber int64 `json:"page_number,required"`
	TotalPages int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *DefaultPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPagination[T any] struct {
	Data []T                   `json:"data"`
	Meta DefaultPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultPagination[T]) GetNextPage() (res *DefaultPagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	currentPage := r.Meta.PageNumber
	if currentPage >= r.Meta.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultPaginationAutoPager[T any] struct {
	page *DefaultPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultPaginationAutoPager[T any](page *DefaultPagination[T], err error) *DefaultPaginationAutoPager[T] {
	return &DefaultPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultPaginationAutoPager[T]) Index() int {
	return r.run
}

type DefaultFlatPaginationMeta struct {
	PageNumber int64 `json:"page_number,required"`
	TotalPages int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultFlatPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *DefaultFlatPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultFlatPagination[T any] struct {
	Data []T                       `json:"data"`
	Meta DefaultFlatPaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultFlatPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultFlatPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultFlatPagination[T]) GetNextPage() (res *DefaultFlatPagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	currentPage := r.Meta.PageNumber
	if currentPage >= r.Meta.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page[number]", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultFlatPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultFlatPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultFlatPaginationAutoPager[T any] struct {
	page *DefaultFlatPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultFlatPaginationAutoPager[T any](page *DefaultFlatPagination[T], err error) *DefaultFlatPaginationAutoPager[T] {
	return &DefaultFlatPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultFlatPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultFlatPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultFlatPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultFlatPaginationAutoPager[T]) Index() int {
	return r.run
}

type DefaultFlatPaginationTopLevelArray[T any] struct {
	Items []T `json:",inline"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultFlatPaginationTopLevelArray[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultFlatPaginationTopLevelArray[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultFlatPaginationTopLevelArray[T]) GetNextPage() (res *DefaultFlatPaginationTopLevelArray[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page[number]"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page[number]", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultFlatPaginationTopLevelArray[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultFlatPaginationTopLevelArray[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultFlatPaginationTopLevelArrayAutoPager[T any] struct {
	page *DefaultFlatPaginationTopLevelArray[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultFlatPaginationTopLevelArrayAutoPager[T any](page *DefaultFlatPaginationTopLevelArray[T], err error) *DefaultFlatPaginationTopLevelArrayAutoPager[T] {
	return &DefaultFlatPaginationTopLevelArrayAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultFlatPaginationTopLevelArrayAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultFlatPaginationTopLevelArrayAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultFlatPaginationTopLevelArrayAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultFlatPaginationTopLevelArrayAutoPager[T]) Index() int {
	return r.run
}

type DefaultPaginationForLogMessagesMeta struct {
	PageNumber int64 `json:"page_number,required"`
	TotalPages int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultPaginationForLogMessagesMeta) RawJSON() string { return r.JSON.raw }
func (r *DefaultPaginationForLogMessagesMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultPaginationForLogMessages[T any] struct {
	LogMessages []T                                 `json:"log_messages"`
	Meta        DefaultPaginationForLogMessagesMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogMessages respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultPaginationForLogMessages[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultPaginationForLogMessages[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultPaginationForLogMessages[T]) GetNextPage() (res *DefaultPaginationForLogMessages[T], err error) {
	if len(r.LogMessages) == 0 {
		return nil, nil
	}
	currentPage := r.Meta.PageNumber
	if currentPage >= r.Meta.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultPaginationForLogMessages[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultPaginationForLogMessages[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultPaginationForLogMessagesAutoPager[T any] struct {
	page *DefaultPaginationForLogMessages[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultPaginationForLogMessagesAutoPager[T any](page *DefaultPaginationForLogMessages[T], err error) *DefaultPaginationForLogMessagesAutoPager[T] {
	return &DefaultPaginationForLogMessagesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultPaginationForLogMessagesAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.LogMessages) == 0 {
		return false
	}
	if r.idx >= len(r.page.LogMessages) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.LogMessages) == 0 {
			return false
		}
	}
	r.cur = r.page.LogMessages[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultPaginationForLogMessagesAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultPaginationForLogMessagesAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultPaginationForLogMessagesAutoPager[T]) Index() int {
	return r.run
}

type DefaultPaginationForMessagingTollfree[T any] struct {
	Records      []T   `json:"records"`
	TotalRecords int64 `json:"total_records"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Records      respjson.Field
		TotalRecords respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultPaginationForMessagingTollfree[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultPaginationForMessagingTollfree[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultPaginationForMessagingTollfree[T]) GetNextPage() (res *DefaultPaginationForMessagingTollfree[T], err error) {
	if len(r.Records) == 0 {
		return nil, nil
	}
	u := r.cfg.Request.URL
	currentPage, err := strconv.ParseInt(u.Query().Get("page"), 10, 64)
	if err != nil {
		currentPage = 1
	}
	if currentPage >= r.TotalRecords {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultPaginationForMessagingTollfree[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultPaginationForMessagingTollfree[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultPaginationForMessagingTollfreeAutoPager[T any] struct {
	page *DefaultPaginationForMessagingTollfree[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultPaginationForMessagingTollfreeAutoPager[T any](page *DefaultPaginationForMessagingTollfree[T], err error) *DefaultPaginationForMessagingTollfreeAutoPager[T] {
	return &DefaultPaginationForMessagingTollfreeAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultPaginationForMessagingTollfreeAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Records) == 0 {
		return false
	}
	if r.idx >= len(r.page.Records) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Records) == 0 {
			return false
		}
	}
	r.cur = r.page.Records[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultPaginationForMessagingTollfreeAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultPaginationForMessagingTollfreeAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultPaginationForMessagingTollfreeAutoPager[T]) Index() int {
	return r.run
}

type DefaultFlatPaginationForInexplicitNumberOrdersMeta struct {
	PageNumber int64 `json:"page_number,required"`
	TotalPages int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultFlatPaginationForInexplicitNumberOrdersMeta) RawJSON() string { return r.JSON.raw }
func (r *DefaultFlatPaginationForInexplicitNumberOrdersMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultFlatPaginationForInexplicitNumberOrders[T any] struct {
	Data []T                                                `json:"data"`
	Meta DefaultFlatPaginationForInexplicitNumberOrdersMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DefaultFlatPaginationForInexplicitNumberOrders[T]) RawJSON() string { return r.JSON.raw }
func (r *DefaultFlatPaginationForInexplicitNumberOrders[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DefaultFlatPaginationForInexplicitNumberOrders[T]) GetNextPage() (res *DefaultFlatPaginationForInexplicitNumberOrders[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	currentPage := r.Meta.PageNumber
	if currentPage >= r.Meta.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page_number", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DefaultFlatPaginationForInexplicitNumberOrders[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DefaultFlatPaginationForInexplicitNumberOrders[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T any] struct {
	page *DefaultFlatPaginationForInexplicitNumberOrders[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T any](page *DefaultFlatPaginationForInexplicitNumberOrders[T], err error) *DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T] {
	return &DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T]) Current() T {
	return r.cur
}

func (r *DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T]) Err() error {
	return r.err
}

func (r *DefaultFlatPaginationForInexplicitNumberOrdersAutoPager[T]) Index() int {
	return r.run
}

type PerPagePaginationMeta struct {
	PageNumber int64 `json:"page_number,required"`
	TotalPages int64 `json:"total_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PerPagePaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PerPagePaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PerPagePagination[T any] struct {
	Data []T                   `json:"data"`
	Meta PerPagePaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PerPagePagination[T]) RawJSON() string { return r.JSON.raw }
func (r *PerPagePagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PerPagePagination[T]) GetNextPage() (res *PerPagePagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	currentPage := r.Meta.PageNumber
	if currentPage >= r.Meta.TotalPages {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PerPagePagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PerPagePagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PerPagePaginationAutoPager[T any] struct {
	page *PerPagePagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPerPagePaginationAutoPager[T any](page *PerPagePagination[T], err error) *PerPagePaginationAutoPager[T] {
	return &PerPagePaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PerPagePaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PerPagePaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *PerPagePaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *PerPagePaginationAutoPager[T]) Index() int {
	return r.run
}

type PerPagePaginationV2[T any] struct {
	Records      []T   `json:"records"`
	Page         int64 `json:"page,required"`
	TotalRecords int64 `json:"totalRecords,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Records      respjson.Field
		Page         respjson.Field
		TotalRecords respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PerPagePaginationV2[T]) RawJSON() string { return r.JSON.raw }
func (r *PerPagePaginationV2[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PerPagePaginationV2[T]) GetNextPage() (res *PerPagePaginationV2[T], err error) {
	if len(r.Records) == 0 {
		return nil, nil
	}
	currentPage := r.Page
	if currentPage >= r.TotalRecords {
		return nil, nil
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PerPagePaginationV2[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PerPagePaginationV2[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PerPagePaginationV2AutoPager[T any] struct {
	page *PerPagePaginationV2[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPerPagePaginationV2AutoPager[T any](page *PerPagePaginationV2[T], err error) *PerPagePaginationV2AutoPager[T] {
	return &PerPagePaginationV2AutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PerPagePaginationV2AutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Records) == 0 {
		return false
	}
	if r.idx >= len(r.page.Records) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Records) == 0 {
			return false
		}
	}
	r.cur = r.page.Records[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PerPagePaginationV2AutoPager[T]) Current() T {
	return r.cur
}

func (r *PerPagePaginationV2AutoPager[T]) Err() error {
	return r.err
}

func (r *PerPagePaginationV2AutoPager[T]) Index() int {
	return r.run
}
