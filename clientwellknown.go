// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// ClientWellKnownService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientWellKnownService] method instead.
type ClientWellKnownService struct {
	Options []option.RequestOption
}

// NewClientWellKnownService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientWellKnownService(opts ...option.RequestOption) (r ClientWellKnownService) {
	r = ClientWellKnownService{}
	r.Options = opts
	return
}

// OAuth 2.0 Authorization Server Metadata (RFC 8414)
func (r *ClientWellKnownService) GetAuthorizationServerMetadata(ctx context.Context, opts ...option.RequestOption) (res *ClientWellKnownGetAuthorizationServerMetadataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.telnyx.com/")}, opts...)
	path := ".well-known/oauth-authorization-server"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// OAuth 2.0 Protected Resource Metadata for resource discovery
func (r *ClientWellKnownService) GetProtectedResourceMetadata(ctx context.Context, opts ...option.RequestOption) (res *ClientWellKnownGetProtectedResourceMetadataResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("https://api.telnyx.com/")}, opts...)
	path := ".well-known/oauth-protected-resource"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientWellKnownGetAuthorizationServerMetadataResponse struct {
	// Authorization endpoint URL
	AuthorizationEndpoint string `json:"authorization_endpoint" format:"uri"`
	// Supported PKCE code challenge methods
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported"`
	// Supported grant types
	GrantTypesSupported []string `json:"grant_types_supported"`
	// Token introspection endpoint URL
	IntrospectionEndpoint string `json:"introspection_endpoint" format:"uri"`
	// Authorization server issuer URL
	Issuer string `json:"issuer" format:"uri"`
	// JWK Set endpoint URL
	JwksUri string `json:"jwks_uri" format:"uri"`
	// Dynamic client registration endpoint URL
	RegistrationEndpoint string `json:"registration_endpoint" format:"uri"`
	// Supported response types
	ResponseTypesSupported []string `json:"response_types_supported"`
	// Supported OAuth scopes
	ScopesSupported []string `json:"scopes_supported"`
	// Token endpoint URL
	TokenEndpoint string `json:"token_endpoint" format:"uri"`
	// Supported token endpoint authentication methods
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationEndpoint             respjson.Field
		CodeChallengeMethodsSupported     respjson.Field
		GrantTypesSupported               respjson.Field
		IntrospectionEndpoint             respjson.Field
		Issuer                            respjson.Field
		JwksUri                           respjson.Field
		RegistrationEndpoint              respjson.Field
		ResponseTypesSupported            respjson.Field
		ScopesSupported                   respjson.Field
		TokenEndpoint                     respjson.Field
		TokenEndpointAuthMethodsSupported respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientWellKnownGetAuthorizationServerMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientWellKnownGetAuthorizationServerMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientWellKnownGetProtectedResourceMetadataResponse struct {
	// List of authorization server URLs
	AuthorizationServers []string `json:"authorization_servers" format:"uri"`
	// Protected resource URL
	Resource string `json:"resource" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthorizationServers respjson.Field
		Resource             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientWellKnownGetProtectedResourceMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientWellKnownGetProtectedResourceMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
