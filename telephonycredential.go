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
)

// TelephonyCredentialService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTelephonyCredentialService] method instead.
type TelephonyCredentialService struct {
	Options []option.RequestOption
}

// NewTelephonyCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTelephonyCredentialService(opts ...option.RequestOption) (r TelephonyCredentialService) {
	r = TelephonyCredentialService{}
	r.Options = opts
	return
}

// Create a credential.
func (r *TelephonyCredentialService) New(ctx context.Context, body TelephonyCredentialNewParams, opts ...option.RequestOption) (res *TelephonyCredentialNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "telephony_credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the details of an existing On-demand Credential.
func (r *TelephonyCredentialService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TelephonyCredentialGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("telephony_credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing credential.
func (r *TelephonyCredentialService) Update(ctx context.Context, id string, body TelephonyCredentialUpdateParams, opts ...option.RequestOption) (res *TelephonyCredentialUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("telephony_credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all On-demand Credentials.
func (r *TelephonyCredentialService) List(ctx context.Context, query TelephonyCredentialListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[TelephonyCredential], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "telephony_credentials"
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

// List all On-demand Credentials.
func (r *TelephonyCredentialService) ListAutoPaging(ctx context.Context, query TelephonyCredentialListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[TelephonyCredential] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete an existing credential.
func (r *TelephonyCredentialService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TelephonyCredentialDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("telephony_credentials/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create an Access Token (JWT) for the credential.
func (r *TelephonyCredentialService) NewToken(ctx context.Context, id string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("telephony_credentials/%s/token", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type TelephonyCredential struct {
	// Identifies the resource.
	ID string `json:"id"`
	// ISO-8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// Defaults to false
	Expired bool `json:"expired"`
	// ISO-8601 formatted date indicating when the resource will expire.
	ExpiresAt string `json:"expires_at"`
	Name      string `json:"name"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Identifies the resource this credential is associated with.
	ResourceID string `json:"resource_id"`
	// The randomly generated SIP password for the credential.
	SipPassword string `json:"sip_password"`
	// The randomly generated SIP username for the credential.
	SipUsername string `json:"sip_username"`
	// ISO-8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Expired     respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		ResourceID  respjson.Field
		SipPassword respjson.Field
		SipUsername respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonyCredential) RawJSON() string { return r.JSON.raw }
func (r *TelephonyCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialNewResponse struct {
	Data TelephonyCredential `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonyCredentialNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TelephonyCredentialNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialGetResponse struct {
	Data TelephonyCredential `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonyCredentialGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TelephonyCredentialGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialUpdateResponse struct {
	Data TelephonyCredential `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonyCredentialUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TelephonyCredentialUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialDeleteResponse struct {
	Data TelephonyCredential `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonyCredentialDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TelephonyCredentialDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialNewParams struct {
	// Identifies the Credential Connection this credential is associated with.
	ConnectionID string `json:"connection_id,required"`
	// ISO-8601 formatted date indicating when the credential will expire.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	// Tags a credential. A single tag can hold at maximum 1000 credentials.
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

func (r TelephonyCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TelephonyCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonyCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialUpdateParams struct {
	// Identifies the Credential Connection this credential is associated with.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// ISO-8601 formatted date indicating when the credential will expire.
	ExpiresAt param.Opt[string] `json:"expires_at,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	// Tags a credential. A single tag can hold at maximum 1000 credentials.
	Tag param.Opt[string] `json:"tag,omitzero"`
	paramObj
}

func (r TelephonyCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TelephonyCredentialUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonyCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyCredentialListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[tag],
	// filter[name], filter[status], filter[resource_id], filter[sip_username]
	Filter TelephonyCredentialListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page TelephonyCredentialListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TelephonyCredentialListParams]'s query parameters as
// `url.Values`.
func (r TelephonyCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[tag],
// filter[name], filter[status], filter[resource_id], filter[sip_username]
type TelephonyCredentialListParamsFilter struct {
	// Filter by name
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter by resource_id
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter by sip_username
	SipUsername param.Opt[string] `query:"sip_username,omitzero" json:"-"`
	// Filter by status
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter by tag
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TelephonyCredentialListParamsFilter]'s query parameters as
// `url.Values`.
func (r TelephonyCredentialListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type TelephonyCredentialListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TelephonyCredentialListParamsPage]'s query parameters as
// `url.Values`.
func (r TelephonyCredentialListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
