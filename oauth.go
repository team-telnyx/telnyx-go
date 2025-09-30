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

// OAuthService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthService] method instead.
type OAuthService struct {
	Options []option.RequestOption
}

// NewOAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOAuthService(opts ...option.RequestOption) (r OAuthService) {
	r = OAuthService{}
	r.Options = opts
	return
}

// OAuth 2.0 authorization endpoint for the authorization code flow
func (r *OAuthService) Authorize(ctx context.Context, query OAuthAuthorizeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "oauth/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Create an OAuth authorization grant
func (r *OAuthService) NewGrant(ctx context.Context, body OAuthNewGrantParams, opts ...option.RequestOption) (res *OAuthNewGrantResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/grants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Exchange authorization code, client credentials, or refresh token for access
// token
func (r *OAuthService) ExchangeToken(ctx context.Context, body OAuthExchangeTokenParams, opts ...option.RequestOption) (res *OAuthExchangeTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Introspect an OAuth access token to check its validity and metadata
func (r *OAuthService) IntrospectToken(ctx context.Context, body OAuthIntrospectTokenParams, opts ...option.RequestOption) (res *OAuthIntrospectTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/introspect"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Register a new OAuth client dynamically (RFC 7591)
func (r *OAuthService) RegisterClient(ctx context.Context, body OAuthRegisterClientParams, opts ...option.RequestOption) (res *OAuthRegisterClientResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details about an OAuth consent token
func (r *OAuthService) GetConsent(ctx context.Context, consentToken string, opts ...option.RequestOption) (res *OAuthGetConsentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if consentToken == "" {
		err = errors.New("missing required consent_token parameter")
		return
	}
	path := fmt.Sprintf("oauth/consent/%s", consentToken)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the JSON Web Key Set for token verification
func (r *OAuthService) GetJwks(ctx context.Context, opts ...option.RequestOption) (res *OAuthGetJwksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/jwks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OAuthNewGrantResponse struct {
	// Redirect URI with authorization code or error
	RedirectUri string `json:"redirect_uri,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RedirectUri respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthNewGrantResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthNewGrantResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthExchangeTokenResponse struct {
	// The access token
	AccessToken string `json:"access_token,required"`
	// Token lifetime in seconds
	ExpiresIn int64 `json:"expires_in,required"`
	// Token type
	//
	// Any of "Bearer".
	TokenType OAuthExchangeTokenResponseTokenType `json:"token_type,required"`
	// Refresh token (if applicable)
	RefreshToken string `json:"refresh_token"`
	// Space-separated list of granted scopes
	Scope string `json:"scope"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken  respjson.Field
		ExpiresIn    respjson.Field
		TokenType    respjson.Field
		RefreshToken respjson.Field
		Scope        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthExchangeTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthExchangeTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token type
type OAuthExchangeTokenResponseTokenType string

const (
	OAuthExchangeTokenResponseTokenTypeBearer OAuthExchangeTokenResponseTokenType = "Bearer"
)

type OAuthIntrospectTokenResponse struct {
	// Whether the token is active
	Active bool `json:"active,required"`
	// Audience
	Aud string `json:"aud"`
	// Client identifier
	ClientID string `json:"client_id"`
	// Expiration timestamp
	Exp int64 `json:"exp"`
	// Issued at timestamp
	Iat int64 `json:"iat"`
	// Issuer
	Iss string `json:"iss"`
	// Space-separated list of scopes
	Scope string `json:"scope"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Aud         respjson.Field
		ClientID    respjson.Field
		Exp         respjson.Field
		Iat         respjson.Field
		Iss         respjson.Field
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthIntrospectTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthIntrospectTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthRegisterClientResponse struct {
	// Unique client identifier
	ClientID string `json:"client_id,required"`
	// Unix timestamp of when the client ID was issued
	ClientIDIssuedAt int64 `json:"client_id_issued_at,required"`
	// Human-readable client name
	ClientName string `json:"client_name"`
	// Client secret (only for confidential clients)
	ClientSecret string `json:"client_secret"`
	// Array of allowed grant types
	GrantTypes []string `json:"grant_types"`
	// URL of the client logo
	LogoUri string `json:"logo_uri" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri" format:"uri"`
	// Array of redirection URIs
	RedirectUris []string `json:"redirect_uris" format:"uri"`
	// Array of allowed response types
	ResponseTypes []string `json:"response_types"`
	// Space-separated scope values
	Scope string `json:"scope"`
	// Token endpoint authentication method
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID                respjson.Field
		ClientIDIssuedAt        respjson.Field
		ClientName              respjson.Field
		ClientSecret            respjson.Field
		GrantTypes              respjson.Field
		LogoUri                 respjson.Field
		PolicyUri               respjson.Field
		RedirectUris            respjson.Field
		ResponseTypes           respjson.Field
		Scope                   respjson.Field
		TokenEndpointAuthMethod respjson.Field
		TosUri                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthRegisterClientResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthRegisterClientResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGetConsentResponse struct {
	Data OAuthGetConsentResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGetConsentResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthGetConsentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGetConsentResponseData struct {
	// Client ID
	ClientID string `json:"client_id"`
	// URL of the client logo
	LogoUri string `json:"logo_uri,nullable" format:"uri"`
	// Client name
	Name string `json:"name"`
	// URL of the client's privacy policy
	PolicyUri string `json:"policy_uri,nullable" format:"uri"`
	// The redirect URI for this authorization
	RedirectUri     string                                      `json:"redirect_uri" format:"uri"`
	RequestedScopes []OAuthGetConsentResponseDataRequestedScope `json:"requested_scopes"`
	// URL of the client's terms of service
	TosUri string `json:"tos_uri,nullable" format:"uri"`
	// Whether the client is verified
	Verified bool `json:"verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID        respjson.Field
		LogoUri         respjson.Field
		Name            respjson.Field
		PolicyUri       respjson.Field
		RedirectUri     respjson.Field
		RequestedScopes respjson.Field
		TosUri          respjson.Field
		Verified        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGetConsentResponseData) RawJSON() string { return r.JSON.raw }
func (r *OAuthGetConsentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGetConsentResponseDataRequestedScope struct {
	// Scope ID
	ID string `json:"id"`
	// Scope description
	Description string `json:"description"`
	// Scope name
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGetConsentResponseDataRequestedScope) RawJSON() string { return r.JSON.raw }
func (r *OAuthGetConsentResponseDataRequestedScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGetJwksResponse struct {
	Keys []OAuthGetJwksResponseKey `json:"keys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGetJwksResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthGetJwksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGetJwksResponseKey struct {
	// Algorithm
	Alg string `json:"alg"`
	// Key ID
	Kid string `json:"kid"`
	// Key type
	Kty string `json:"kty"`
	// Key use
	Use string `json:"use"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alg         respjson.Field
		Kid         respjson.Field
		Kty         respjson.Field
		Use         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGetJwksResponseKey) RawJSON() string { return r.JSON.raw }
func (r *OAuthGetJwksResponseKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthAuthorizeParams struct {
	// OAuth client identifier
	ClientID string `query:"client_id,required" json:"-"`
	// Redirect URI
	RedirectUri string `query:"redirect_uri,required" format:"uri" json:"-"`
	// OAuth response type
	//
	// Any of "code".
	ResponseType OAuthAuthorizeParamsResponseType `query:"response_type,omitzero,required" json:"-"`
	// PKCE code challenge
	CodeChallenge param.Opt[string] `query:"code_challenge,omitzero" json:"-"`
	// Space-separated list of requested scopes
	Scope param.Opt[string] `query:"scope,omitzero" json:"-"`
	// State parameter for CSRF protection
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// PKCE code challenge method
	//
	// Any of "plain", "S256".
	CodeChallengeMethod OAuthAuthorizeParamsCodeChallengeMethod `query:"code_challenge_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OAuthAuthorizeParams]'s query parameters as `url.Values`.
func (r OAuthAuthorizeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// OAuth response type
type OAuthAuthorizeParamsResponseType string

const (
	OAuthAuthorizeParamsResponseTypeCode OAuthAuthorizeParamsResponseType = "code"
)

// PKCE code challenge method
type OAuthAuthorizeParamsCodeChallengeMethod string

const (
	OAuthAuthorizeParamsCodeChallengeMethodPlain OAuthAuthorizeParamsCodeChallengeMethod = "plain"
	OAuthAuthorizeParamsCodeChallengeMethodS256  OAuthAuthorizeParamsCodeChallengeMethod = "S256"
)

type OAuthNewGrantParams struct {
	// Whether the grant is allowed
	Allowed bool `json:"allowed,required"`
	// Consent token
	ConsentToken string `json:"consent_token,required"`
	paramObj
}

func (r OAuthNewGrantParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthNewGrantParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthNewGrantParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthExchangeTokenParams struct {
	// OAuth 2.0 grant type
	//
	// Any of "client_credentials", "authorization_code", "refresh_token".
	GrantType OAuthExchangeTokenParamsGrantType `json:"grant_type,omitzero,required"`
	// OAuth client ID (if not using HTTP Basic auth)
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// OAuth client secret (if not using HTTP Basic auth)
	ClientSecret param.Opt[string] `json:"client_secret,omitzero"`
	// Authorization code (for authorization_code flow)
	Code param.Opt[string] `json:"code,omitzero"`
	// PKCE code verifier (for authorization_code flow)
	CodeVerifier param.Opt[string] `json:"code_verifier,omitzero"`
	// Redirect URI (for authorization_code flow)
	RedirectUri param.Opt[string] `json:"redirect_uri,omitzero" format:"uri"`
	// Refresh token (for refresh_token flow)
	RefreshToken param.Opt[string] `json:"refresh_token,omitzero"`
	// Space-separated list of requested scopes (for client_credentials)
	Scope param.Opt[string] `json:"scope,omitzero"`
	paramObj
}

func (r OAuthExchangeTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthExchangeTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthExchangeTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth 2.0 grant type
type OAuthExchangeTokenParamsGrantType string

const (
	OAuthExchangeTokenParamsGrantTypeClientCredentials OAuthExchangeTokenParamsGrantType = "client_credentials"
	OAuthExchangeTokenParamsGrantTypeAuthorizationCode OAuthExchangeTokenParamsGrantType = "authorization_code"
	OAuthExchangeTokenParamsGrantTypeRefreshToken      OAuthExchangeTokenParamsGrantType = "refresh_token"
)

type OAuthIntrospectTokenParams struct {
	// The token to introspect
	Token string `json:"token,required"`
	paramObj
}

func (r OAuthIntrospectTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthIntrospectTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthIntrospectTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthRegisterClientParams struct {
	// Human-readable string name of the client to be presented to the end-user
	ClientName param.Opt[string] `json:"client_name,omitzero"`
	// URL of the client logo
	LogoUri param.Opt[string] `json:"logo_uri,omitzero" format:"uri"`
	// URL of the client's privacy policy
	PolicyUri param.Opt[string] `json:"policy_uri,omitzero" format:"uri"`
	// Space-separated string of scope values that the client may use
	Scope param.Opt[string] `json:"scope,omitzero"`
	// URL of the client's terms of service
	TosUri param.Opt[string] `json:"tos_uri,omitzero" format:"uri"`
	// Array of OAuth 2.0 grant type strings that the client may use
	//
	// Any of "authorization_code", "client_credentials", "refresh_token".
	GrantTypes []string `json:"grant_types,omitzero"`
	// Array of redirection URI strings for use in redirect-based flows
	RedirectUris []string `json:"redirect_uris,omitzero" format:"uri"`
	// Array of the OAuth 2.0 response type strings that the client may use
	ResponseTypes []string `json:"response_types,omitzero"`
	// Authentication method for the token endpoint
	//
	// Any of "none", "client_secret_basic", "client_secret_post".
	TokenEndpointAuthMethod OAuthRegisterClientParamsTokenEndpointAuthMethod `json:"token_endpoint_auth_method,omitzero"`
	paramObj
}

func (r OAuthRegisterClientParams) MarshalJSON() (data []byte, err error) {
	type shadow OAuthRegisterClientParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OAuthRegisterClientParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authentication method for the token endpoint
type OAuthRegisterClientParamsTokenEndpointAuthMethod string

const (
	OAuthRegisterClientParamsTokenEndpointAuthMethodNone              OAuthRegisterClientParamsTokenEndpointAuthMethod = "none"
	OAuthRegisterClientParamsTokenEndpointAuthMethodClientSecretBasic OAuthRegisterClientParamsTokenEndpointAuthMethod = "client_secret_basic"
	OAuthRegisterClientParamsTokenEndpointAuthMethodClientSecretPost  OAuthRegisterClientParamsTokenEndpointAuthMethod = "client_secret_post"
)
