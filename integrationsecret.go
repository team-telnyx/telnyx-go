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

// IntegrationSecretService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationSecretService] method instead.
type IntegrationSecretService struct {
	Options []option.RequestOption
}

// NewIntegrationSecretService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationSecretService(opts ...option.RequestOption) (r IntegrationSecretService) {
	r = IntegrationSecretService{}
	r.Options = opts
	return
}

// Create a new secret with an associated identifier that can be used to securely
// integrate with other services.
func (r *IntegrationSecretService) New(ctx context.Context, body IntegrationSecretNewParams, opts ...option.RequestOption) (res *IntegrationSecretNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integration_secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a list of all integration secrets configured by the user.
func (r *IntegrationSecretService) List(ctx context.Context, query IntegrationSecretListParams, opts ...option.RequestOption) (res *IntegrationSecretListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integration_secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an integration secret given its ID.
func (r *IntegrationSecretService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("integration_secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type IntegrationSecret struct {
	ID         string    `json:"id,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	Identifier string    `json:"identifier,required"`
	RecordType string    `json:"record_type,required"`
	UpdatedAt  time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Identifier  respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSecret) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSecretNewResponse struct {
	Data IntegrationSecret `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSecretNewResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSecretNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSecretListResponse struct {
	Data []IntegrationSecret               `json:"data,required"`
	Meta IntegrationSecretListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSecretListResponseMeta struct {
	PageNumber   int64 `json:"page_number,required"`
	PageSize     int64 `json:"page_size,required"`
	TotalPages   int64 `json:"total_pages,required"`
	TotalResults int64 `json:"total_results,required"`
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
func (r IntegrationSecretListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSecretListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSecretNewParams struct {
	// The unique identifier of the secret.
	Identifier string `json:"identifier,required"`
	// The type of secret.
	//
	// Any of "bearer", "basic".
	Type IntegrationSecretNewParamsType `json:"type,omitzero,required"`
	// The token for the secret. Required for bearer type secrets, ignored otherwise.
	Token param.Opt[string] `json:"token,omitzero"`
	// The password for the secret. Required for basic type secrets, ignored otherwise.
	Password param.Opt[string] `json:"password,omitzero"`
	// The username for the secret. Required for basic type secrets, ignored otherwise.
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r IntegrationSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationSecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of secret.
type IntegrationSecretNewParamsType string

const (
	IntegrationSecretNewParamsTypeBearer IntegrationSecretNewParamsType = "bearer"
	IntegrationSecretNewParamsTypeBasic  IntegrationSecretNewParamsType = "basic"
)

type IntegrationSecretListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[type]
	Filter IntegrationSecretListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page IntegrationSecretListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationSecretListParams]'s query parameters as
// `url.Values`.
func (r IntegrationSecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[type]
type IntegrationSecretListParamsFilter struct {
	// Any of "bearer", "basic".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationSecretListParamsFilter]'s query parameters as
// `url.Values`.
func (r IntegrationSecretListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type IntegrationSecretListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationSecretListParamsPage]'s query parameters as
// `url.Values`.
func (r IntegrationSecretListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
