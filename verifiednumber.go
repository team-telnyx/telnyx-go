// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// VerifiedNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifiedNumberService] method instead.
type VerifiedNumberService struct {
	Options []option.RequestOption
	Actions VerifiedNumberActionService
}

// NewVerifiedNumberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifiedNumberService(opts ...option.RequestOption) (r VerifiedNumberService) {
	r = VerifiedNumberService{}
	r.Options = opts
	r.Actions = NewVerifiedNumberActionService(opts...)
	return
}

// Initiates phone number verification procedure.
func (r *VerifiedNumberService) New(ctx context.Context, body VerifiedNumberNewParams, opts ...option.RequestOption) (res *VerifiedNumberNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "verified_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a verified number
func (r *VerifiedNumberService) Get(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *VerifiedNumberDataWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("verified_numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a paginated list of Verified Numbers.
func (r *VerifiedNumberService) List(ctx context.Context, query VerifiedNumberListParams, opts ...option.RequestOption) (res *VerifiedNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "verified_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a verified number
func (r *VerifiedNumberService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *VerifiedNumberDataWrapper, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("verified_numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type VerifiedNumber struct {
	PhoneNumber string `json:"phone_number"`
	// The possible verified numbers record types.
	//
	// Any of "verified_number".
	RecordType VerifiedNumberRecordType `json:"record_type"`
	VerifiedAt string                   `json:"verified_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedNumber) RawJSON() string { return r.JSON.raw }
func (r *VerifiedNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The possible verified numbers record types.
type VerifiedNumberRecordType string

const (
	VerifiedNumberRecordTypeVerifiedNumber VerifiedNumberRecordType = "verified_number"
)

type VerifiedNumberDataWrapper struct {
	Data VerifiedNumber `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedNumberDataWrapper) RawJSON() string { return r.JSON.raw }
func (r *VerifiedNumberDataWrapper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifiedNumberNewResponse struct {
	PhoneNumber        string `json:"phone_number"`
	VerificationMethod string `json:"verification_method"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber        respjson.Field
		VerificationMethod respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedNumberNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifiedNumberNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of Verified Numbers
type VerifiedNumberListResponse struct {
	Data []VerifiedNumber               `json:"data,required"`
	Meta VerifiedNumberListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifiedNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifiedNumberListResponseMeta struct {
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
func (r VerifiedNumberListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *VerifiedNumberListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifiedNumberNewParams struct {
	PhoneNumber string `json:"phone_number,required"`
	// Verification method.
	//
	// Any of "sms", "call".
	VerificationMethod VerifiedNumberNewParamsVerificationMethod `json:"verification_method,omitzero,required"`
	paramObj
}

func (r VerifiedNumberNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VerifiedNumberNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerifiedNumberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification method.
type VerifiedNumberNewParamsVerificationMethod string

const (
	VerifiedNumberNewParamsVerificationMethodSMS  VerifiedNumberNewParamsVerificationMethod = "sms"
	VerifiedNumberNewParamsVerificationMethodCall VerifiedNumberNewParamsVerificationMethod = "call"
)

type VerifiedNumberListParams struct {
	// Consolidated page parameter (deepObject style). Use page[size] and page[number]
	// in the query string. Originally: page[size], page[number]
	Page VerifiedNumberListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VerifiedNumberListParams]'s query parameters as
// `url.Values`.
func (r VerifiedNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Use page[size] and page[number]
// in the query string. Originally: page[size], page[number]
type VerifiedNumberListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VerifiedNumberListParamsPage]'s query parameters as
// `url.Values`.
func (r VerifiedNumberListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
