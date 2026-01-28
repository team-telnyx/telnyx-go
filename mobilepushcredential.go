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
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// MobilePushCredentialService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMobilePushCredentialService] method instead.
type MobilePushCredentialService struct {
	Options []option.RequestOption
}

// NewMobilePushCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMobilePushCredentialService(opts ...option.RequestOption) (r MobilePushCredentialService) {
	r = MobilePushCredentialService{}
	r.Options = opts
	return
}

// Creates a new mobile push credential
func (r *MobilePushCredentialService) New(ctx context.Context, body MobilePushCredentialNewParams, opts ...option.RequestOption) (res *PushCredentialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mobile_push_credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves mobile push credential based on the given `push_credential_id`
func (r *MobilePushCredentialService) Get(ctx context.Context, pushCredentialID string, opts ...option.RequestOption) (res *PushCredentialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pushCredentialID == "" {
		err = errors.New("missing required push_credential_id parameter")
		return
	}
	path := fmt.Sprintf("mobile_push_credentials/%s", pushCredentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List mobile push credentials
func (r *MobilePushCredentialService) List(ctx context.Context, query MobilePushCredentialListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PushCredential], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "mobile_push_credentials"
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

// List mobile push credentials
func (r *MobilePushCredentialService) ListAutoPaging(ctx context.Context, query MobilePushCredentialListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PushCredential] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a mobile push credential based on the given `push_credential_id`
func (r *MobilePushCredentialService) Delete(ctx context.Context, pushCredentialID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pushCredentialID == "" {
		err = errors.New("missing required push_credential_id parameter")
		return
	}
	path := fmt.Sprintf("mobile_push_credentials/%s", pushCredentialID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PushCredential struct {
	// Unique identifier of a push credential
	ID string `json:"id,required"`
	// Alias to uniquely identify a credential
	Alias string `json:"alias,required"`
	// Apple certificate for sending push notifications. For iOS only
	Certificate string `json:"certificate,required"`
	// ISO 8601 timestamp when the room was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Apple private key for a given certificate for sending push notifications. For
	// iOS only
	PrivateKey string `json:"private_key,required"`
	// Google server key for sending push notifications. For Android only
	ProjectAccountJsonFile map[string]any `json:"project_account_json_file,required"`
	RecordType             string         `json:"record_type,required"`
	// Type of mobile push credential. Either <code>ios</code> or <code>android</code>
	Type string `json:"type,required"`
	// ISO 8601 timestamp when the room was updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Alias                  respjson.Field
		Certificate            respjson.Field
		CreatedAt              respjson.Field
		PrivateKey             respjson.Field
		ProjectAccountJsonFile respjson.Field
		RecordType             respjson.Field
		Type                   respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PushCredential) RawJSON() string { return r.JSON.raw }
func (r *PushCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Success response with details about a push credential
type PushCredentialResponse struct {
	Data PushCredential `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PushCredentialResponse) RawJSON() string { return r.JSON.raw }
func (r *PushCredentialResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePushCredentialNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfIos *MobilePushCredentialNewParamsCreateMobilePushCredentialRequestIos `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAndroid *MobilePushCredentialNewParamsCreateMobilePushCredentialRequestAndroid `json:",inline"`

	paramObj
}

func (u MobilePushCredentialNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfIos, u.OfAndroid)
}
func (r *MobilePushCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Alias, Certificate, PrivateKey, Type are required.
type MobilePushCredentialNewParamsCreateMobilePushCredentialRequestIos struct {
	// Alias to uniquely identify the credential
	Alias string `json:"alias,required"`
	// Certificate as received from APNs
	Certificate string `json:"certificate,required"`
	// Corresponding private key to the certificate as received from APNs
	PrivateKey string `json:"private_key,required"`
	// Type of mobile push credential. Should be <code>ios</code> here
	//
	// This field can be elided, and will marshal its zero value as "ios".
	Type constant.Ios `json:"type,required"`
	paramObj
}

func (r MobilePushCredentialNewParamsCreateMobilePushCredentialRequestIos) MarshalJSON() (data []byte, err error) {
	type shadow MobilePushCredentialNewParamsCreateMobilePushCredentialRequestIos
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePushCredentialNewParamsCreateMobilePushCredentialRequestIos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Alias, ProjectAccountJsonFile, Type are required.
type MobilePushCredentialNewParamsCreateMobilePushCredentialRequestAndroid struct {
	// Alias to uniquely identify the credential
	Alias string `json:"alias,required"`
	// Private key file in JSON format
	ProjectAccountJsonFile map[string]any `json:"project_account_json_file,omitzero,required"`
	// Type of mobile push credential. Should be <code>android</code> here
	//
	// This field can be elided, and will marshal its zero value as "android".
	Type constant.Android `json:"type,required"`
	paramObj
}

func (r MobilePushCredentialNewParamsCreateMobilePushCredentialRequestAndroid) MarshalJSON() (data []byte, err error) {
	type shadow MobilePushCredentialNewParamsCreateMobilePushCredentialRequestAndroid
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MobilePushCredentialNewParamsCreateMobilePushCredentialRequestAndroid) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePushCredentialListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[type],
	// filter[alias]
	Filter MobilePushCredentialListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobilePushCredentialListParams]'s query parameters as
// `url.Values`.
func (r MobilePushCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[type],
// filter[alias]
type MobilePushCredentialListParamsFilter struct {
	// Unique mobile push credential alias
	Alias param.Opt[string] `query:"alias,omitzero" json:"-"`
	// type of mobile push credentials
	//
	// Any of "ios", "android".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobilePushCredentialListParamsFilter]'s query parameters as
// `url.Values`.
func (r MobilePushCredentialListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
