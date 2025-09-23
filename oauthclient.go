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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// OAuthClientService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthClientService] method instead.
type OAuthClientService struct {
	Options []option.RequestOption
}

// NewOAuthClientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthClientService(opts ...option.RequestOption) (r OAuthClientService) {
	r = OAuthClientService{}
	r.Options = opts
	return
}

// Create a new OAuth client
func (r *OAuthClientService) New(ctx context.Context, body OAuthClientNewParams, opts ...option.RequestOption) (res *OAuthClientNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth_clients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a single OAuth client by ID
func (r *OAuthClientService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OAuthClientGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("oauth_clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing OAuth client
func (r *OAuthClientService) Update(ctx context.Context, id string, body OAuthClientUpdateParams, opts ...option.RequestOption) (res *OAuthClientUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("oauth_clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of OAuth clients for the authenticated user
func (r *OAuthClientService) List(ctx context.Context, query OAuthClientListParams, opts ...option.RequestOption) (res *OAuthClientListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth_clients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an OAuth client
func (r *OAuthClientService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("oauth_clients/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type OAuthClientNewResponse struct {
	Data OAuthClientNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientNewResponseData struct {
	// OAuth client identifier
	ClientID string `json:"client_id,required"`
	// OAuth client type
	//
	// Any of "public", "confidential".
	ClientType string `json:"client_type,required"`
	// Timestamp when the client was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the OAuth client
	Name string `json:"name,required"`
	// Organization ID that owns this OAuth client
	OrgID string `json:"org_id,required"`
	// Record type identifier
	//
	// Any of "oauth_client".
	RecordType string `json:"record_type,required"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce bool `json:"require_pkce,required"`
	// Timestamp when the client was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID that created this OAuth client
	UserID string `json:"user_id,required"`
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes"`
	// Client secret (only included when available, for confidential clients)
	ClientSecret string `json:"client_secret,nullable"`
	// URL of the client logo
	LogoUri string `json:"logo_uri,nullable" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri,nullable" format:"uri"`
	// List of allowed redirect URIs
	RedirectUris []string `json:"redirect_uris" format:"uri"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID          respjson.Field
		ClientType        respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		OrgID             respjson.Field
		RecordType        respjson.Field
		RequirePkce       respjson.Field
		UpdatedAt         respjson.Field
		UserID            respjson.Field
		AllowedGrantTypes respjson.Field
		AllowedScopes     respjson.Field
		ClientSecret      respjson.Field
		LogoUri           respjson.Field
		PolicyUri         respjson.Field
		RedirectUris      respjson.Field
		TosUri            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientGetResponse struct {
	Data OAuthClientGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientGetResponseData struct {
	// OAuth client identifier
	ClientID string `json:"client_id,required"`
	// OAuth client type
	//
	// Any of "public", "confidential".
	ClientType string `json:"client_type,required"`
	// Timestamp when the client was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the OAuth client
	Name string `json:"name,required"`
	// Organization ID that owns this OAuth client
	OrgID string `json:"org_id,required"`
	// Record type identifier
	//
	// Any of "oauth_client".
	RecordType string `json:"record_type,required"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce bool `json:"require_pkce,required"`
	// Timestamp when the client was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID that created this OAuth client
	UserID string `json:"user_id,required"`
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes"`
	// Client secret (only included when available, for confidential clients)
	ClientSecret string `json:"client_secret,nullable"`
	// URL of the client logo
	LogoUri string `json:"logo_uri,nullable" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri,nullable" format:"uri"`
	// List of allowed redirect URIs
	RedirectUris []string `json:"redirect_uris" format:"uri"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID          respjson.Field
		ClientType        respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		OrgID             respjson.Field
		RecordType        respjson.Field
		RequirePkce       respjson.Field
		UpdatedAt         respjson.Field
		UserID            respjson.Field
		AllowedGrantTypes respjson.Field
		AllowedScopes     respjson.Field
		ClientSecret      respjson.Field
		LogoUri           respjson.Field
		PolicyUri         respjson.Field
		RedirectUris      respjson.Field
		TosUri            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientUpdateResponse struct {
	Data OAuthClientUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientUpdateResponseData struct {
	// OAuth client identifier
	ClientID string `json:"client_id,required"`
	// OAuth client type
	//
	// Any of "public", "confidential".
	ClientType string `json:"client_type,required"`
	// Timestamp when the client was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the OAuth client
	Name string `json:"name,required"`
	// Organization ID that owns this OAuth client
	OrgID string `json:"org_id,required"`
	// Record type identifier
	//
	// Any of "oauth_client".
	RecordType string `json:"record_type,required"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce bool `json:"require_pkce,required"`
	// Timestamp when the client was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID that created this OAuth client
	UserID string `json:"user_id,required"`
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes"`
	// Client secret (only included when available, for confidential clients)
	ClientSecret string `json:"client_secret,nullable"`
	// URL of the client logo
	LogoUri string `json:"logo_uri,nullable" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri,nullable" format:"uri"`
	// List of allowed redirect URIs
	RedirectUris []string `json:"redirect_uris" format:"uri"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID          respjson.Field
		ClientType        respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		OrgID             respjson.Field
		RecordType        respjson.Field
		RequirePkce       respjson.Field
		UpdatedAt         respjson.Field
		UserID            respjson.Field
		AllowedGrantTypes respjson.Field
		AllowedScopes     respjson.Field
		ClientSecret      respjson.Field
		LogoUri           respjson.Field
		PolicyUri         respjson.Field
		RedirectUris      respjson.Field
		TosUri            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientListResponse struct {
	Data []OAuthClientListResponseData `json:"data"`
	Meta OAuthClientListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientListResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientListResponseData struct {
	// OAuth client identifier
	ClientID string `json:"client_id,required"`
	// OAuth client type
	//
	// Any of "public", "confidential".
	ClientType string `json:"client_type,required"`
	// Timestamp when the client was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name for the OAuth client
	Name string `json:"name,required"`
	// Organization ID that owns this OAuth client
	OrgID string `json:"org_id,required"`
	// Record type identifier
	//
	// Any of "oauth_client".
	RecordType string `json:"record_type,required"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce bool `json:"require_pkce,required"`
	// Timestamp when the client was last updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// User ID that created this OAuth client
	UserID string `json:"user_id,required"`
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes"`
	// Client secret (only included when available, for confidential clients)
	ClientSecret string `json:"client_secret,nullable"`
	// URL of the client logo
	LogoUri string `json:"logo_uri,nullable" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri,nullable" format:"uri"`
	// List of allowed redirect URIs
	RedirectUris []string `json:"redirect_uris" format:"uri"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID          respjson.Field
		ClientType        respjson.Field
		CreatedAt         respjson.Field
		Name              respjson.Field
		OrgID             respjson.Field
		RecordType        respjson.Field
		RequirePkce       respjson.Field
		UpdatedAt         respjson.Field
		UserID            respjson.Field
		AllowedGrantTypes respjson.Field
		AllowedScopes     respjson.Field
		ClientSecret      respjson.Field
		LogoUri           respjson.Field
		PolicyUri         respjson.Field
		RedirectUris      respjson.Field
		TosUri            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthClientListResponseData) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientListResponseMeta struct {
	// Current page number
	PageNumber int64 `json:"page_number"`
	// Number of items per page
	PageSize int64 `json:"page_size"`
	// Total number of pages
	TotalPages int64 `json:"total_pages"`
	// Total number of results
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
func (r OAuthClientListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *OAuthClientListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientNewParams struct {
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types,omitzero,required"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes,omitzero,required"`
	// OAuth client type
	//
	// Any of "public", "confidential".
	ClientType OAuthClientNewParamsClientType `json:"client_type,omitzero,required"`
	// The name of the OAuth client
	Name string `json:"name,required"`
	// URL of the client logo
	LogoUri param.Opt[string] `json:"logo_uri,omitzero" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri param.Opt[string] `json:"policy_uri,omitzero" format:"uri"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce param.Opt[bool] `json:"require_pkce,omitzero"`
	// URL of the client's terms of service
	TosUri param.Opt[string] `json:"tos_uri,omitzero" format:"uri"`
	// List of redirect URIs (required for authorization_code flow)
	RedirectUris []string `json:"redirect_uris,omitzero" format:"uri"`
	paramObj
}

func (r OAuthClientNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthClientNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthClientNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth client type
type OAuthClientNewParamsClientType string

const (
	OAuthClientNewParamsClientTypePublic       OAuthClientNewParamsClientType = "public"
	OAuthClientNewParamsClientTypeConfidential OAuthClientNewParamsClientType = "confidential"
)

type OAuthClientUpdateParams struct {
	// URL of the client logo
	LogoUri param.Opt[string] `json:"logo_uri,omitzero" format:"uri"`
	// The name of the OAuth client
	Name param.Opt[string] `json:"name,omitzero"`
	// URL of the client's privacy policy
	PolicyUri param.Opt[string] `json:"policy_uri,omitzero" format:"uri"`
	// Whether PKCE (Proof Key for Code Exchange) is required for this client
	RequirePkce param.Opt[bool] `json:"require_pkce,omitzero"`
	// URL of the client's terms of service
	TosUri param.Opt[string] `json:"tos_uri,omitzero" format:"uri"`
	// List of allowed OAuth grant types
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	AllowedGrantTypes []string `json:"allowed_grant_types,omitzero"`
	// List of allowed OAuth scopes
	AllowedScopes []string `json:"allowed_scopes,omitzero"`
	// List of redirect URIs
	RedirectUris []string `json:"redirect_uris,omitzero" format:"uri"`
	paramObj
}

func (r OAuthClientUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthClientUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthClientUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthClientListParams struct {
	// Filter by client ID
	FilterClientID param.Opt[string] `query:"filter[client_id],omitzero" json:"-"`
	// Filter by exact client name
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Filter by client name containing text
	FilterNameContains param.Opt[string] `query:"filter[name][contains],omitzero" json:"-"`
	// Filter by verification status
	FilterVerified param.Opt[bool] `query:"filter[verified],omitzero" json:"-"`
	// Page number
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by allowed grant type
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	FilterAllowedGrantTypesContains OAuthClientListParamsFilterAllowedGrantTypesContains `query:"filter[allowed_grant_types][contains],omitzero" json:"-"`
	// Filter by client type
	//
	// Any of "confidential", "public".
	FilterClientType OAuthClientListParamsFilterClientType `query:"filter[client_type],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OAuthClientListParams]'s query parameters as `url.Values`.
func (r OAuthClientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by allowed grant type
type OAuthClientListParamsFilterAllowedGrantTypesContains string

const (
	OAuthClientListParamsFilterAllowedGrantTypesContainsClientCredentials OAuthClientListParamsFilterAllowedGrantTypesContains = "client_credentials"
	OAuthClientListParamsFilterAllowedGrantTypesContainsAuthorizationCode OAuthClientListParamsFilterAllowedGrantTypesContains = "authorization_code"
	OAuthClientListParamsFilterAllowedGrantTypesContainsRefreshToken      OAuthClientListParamsFilterAllowedGrantTypesContains = "refresh_token"
)

// Filter by client type
type OAuthClientListParamsFilterClientType string

const (
	OAuthClientListParamsFilterClientTypeConfidential OAuthClientListParamsFilterClientType = "confidential"
	OAuthClientListParamsFilterClientTypePublic       OAuthClientListParamsFilterClientType = "public"
)
