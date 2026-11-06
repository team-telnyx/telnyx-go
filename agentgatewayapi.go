// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
	"github.com/tidwall/gjson"
)

// REST APIs on an instance, served to agents as tools.
//
// AgentGatewayAPIService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayAPIService] method instead.
type AgentGatewayAPIService struct {
	Options []option.RequestOption
}

// NewAgentGatewayAPIService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentGatewayAPIService(opts ...option.RequestOption) (r AgentGatewayAPIService) {
	r = AgentGatewayAPIService{}
	r.Options = opts
	return
}

// Adds a REST API by its OpenAPI description. Each operation becomes a tool at
// `{gateway_url}/mcp`, named with the API's `name`, a hyphen, and the lowercased
// `operationId`. A description the gateway cannot fetch yields an API with no
// tools. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayAPIService) New(ctx context.Context, instanceID string, body AgentGatewayAPINewParams, opts ...option.RequestOption) (res *AgentGatewayAPIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/apis", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes an API's settings. A field left out stays as it is, `"base_url": null`
// clears the base URL, and `auth` is replaced whole. The instance must be `ready`;
// otherwise the request answers `409`.
func (r *AgentGatewayAPIService) Update(ctx context.Context, apiID string, params AgentGatewayAPIUpdateParams, opts ...option.RequestOption) (res *AgentGatewayAPIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if apiID == "" {
		err = errors.New("missing required api_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/apis/%s", params.InstanceID, apiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the REST APIs the instance serves as tools, without their credentials. The
// instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayAPIService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayAPIListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/apis", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes an API and returns it as it was. An MCP server's `id` answers `404`. The
// instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayAPIService) Delete(ctx context.Context, apiID string, body AgentGatewayAPIDeleteParams, opts ...option.RequestOption) (res *AgentGatewayAPIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if apiID == "" {
		err = errors.New("missing required api_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/apis/%s", body.InstanceID, apiID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AgentGatewayAPI struct {
	// Identifies the API.
	ID string `json:"id" api:"required" format:"uuid"`
	// The server's outbound authentication method. Credentials are never returned.
	Auth AgentGatewayShownAuth `json:"auth" api:"required"`
	// Where the gateway sends the API's requests. `null` takes the description's first
	// `servers` entry.
	BaseURL   string    `json:"base_url" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers" api:"required"`
	// A name for the API, unique on the instance across MCP servers and APIs,
	// regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the
	// API's tools with it and a hyphen.
	Name string `json:"name" api:"required"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey bool `json:"open_to_every_key" api:"required"`
	// The URL of the API's OpenAPI description.
	SpecURL string `json:"spec_url" api:"required"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers" api:"required"`
	// The operations the API exposes as tools, by tool name. `[]` exposes every
	// operation.
	Tools     []string  `json:"tools" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Auth             respjson.Field
		BaseURL          respjson.Field
		CreatedAt        respjson.Field
		ForwardedHeaders respjson.Field
		Name             respjson.Field
		OpenToEveryKey   respjson.Field
		SpecURL          respjson.Field
		StaticHeaders    respjson.Field
		Tools            respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayAPI) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayAPIResponse struct {
	Data AgentGatewayAPI `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The gateway sends no credential.
type AgentGatewayNoAuthParam struct {
	// Any of "none".
	Type AgentGatewayNoAuthType `json:"type,omitzero"`
	paramObj
}

func (r AgentGatewayNoAuthParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayNoAuthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayNoAuthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayNoAuthType string

const (
	AgentGatewayNoAuthTypeNone AgentGatewayNoAuthType = "none"
)

// The gateway gets its own access token with the `client_credentials` grant and
// sends it as a bearer token on every request, renewing it before it expires. The
// client ID and secret are write-only.
//
// The properties ClientID, ClientSecret, TokenURL, Type are required.
type AgentGatewayOAuthClientCredentialsParam struct {
	ClientID     string `json:"client_id" api:"required"`
	ClientSecret string `json:"client_secret" api:"required"`
	// The OAuth provider's token endpoint.
	TokenURL string `json:"token_url" api:"required"`
	// The token's audience, sent in the token request.
	Audience param.Opt[string] `json:"audience,omitzero"`
	// The resource the token is for, sent in the token request.
	Resource param.Opt[string] `json:"resource,omitzero"`
	// Whether the client ID and secret go in the token request's body or as HTTP Basic
	// authentication.
	//
	// Any of "body", "basic".
	CredentialsIn AgentGatewayOAuthClientCredentialsCredentialsIn `json:"credentials_in,omitzero"`
	// The scopes to request.
	Scopes []string `json:"scopes,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "oauth_client_credentials".
	Type constant.OAuthClientCredentials `json:"type" default:"oauth_client_credentials"`
	paramObj
}

func (r AgentGatewayOAuthClientCredentialsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayOAuthClientCredentialsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayOAuthClientCredentialsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the client ID and secret go in the token request's body or as HTTP Basic
// authentication.
type AgentGatewayOAuthClientCredentialsCredentialsIn string

const (
	AgentGatewayOAuthClientCredentialsCredentialsInBody  AgentGatewayOAuthClientCredentialsCredentialsIn = "body"
	AgentGatewayOAuthClientCredentialsCredentialsInBasic AgentGatewayOAuthClientCredentialsCredentialsIn = "basic"
)

// The gateway forwards the agent's own `Authorization` header to the server
// unchanged.
type AgentGatewayPassthroughAuthParam struct {
	// Any of "passthrough".
	Type AgentGatewayPassthroughAuthType `json:"type,omitzero"`
	paramObj
}

func (r AgentGatewayPassthroughAuthParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayPassthroughAuthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayPassthroughAuthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayPassthroughAuthType string

const (
	AgentGatewayPassthroughAuthTypePassthrough AgentGatewayPassthroughAuthType = "passthrough"
)

// The server's outbound authentication method. Credentials are never returned.
type AgentGatewayShownAuth struct {
	// The method's `type`.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayShownAuth) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayShownAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A static credential the gateway sends on every request. Write-only: responses
// show its `type` only.
//
// The properties Type, Value are required.
type AgentGatewayStaticAuthParam struct {
	// How the credential is sent: `api_key` as `X-API-Key: <value>`, `bearer_token` as
	// `Authorization: Bearer <value>`, `basic` as
	// `Authorization: Basic <base64 of value>` with `value` in the form
	// `user:password`, `token` as `Authorization: token <value>`, and `authorization`
	// as `Authorization: <value>`.
	//
	// Any of "api_key", "bearer_token", "basic", "token", "authorization".
	Type AgentGatewayStaticAuthType `json:"type,omitzero" api:"required"`
	// The credential.
	Value string `json:"value" api:"required"`
	// A header name that replaces `X-API-Key` or `Authorization`.
	Header param.Opt[string] `json:"header,omitzero"`
	paramObj
}

func (r AgentGatewayStaticAuthParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayStaticAuthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayStaticAuthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the credential is sent: `api_key` as `X-API-Key: <value>`, `bearer_token` as
// `Authorization: Bearer <value>`, `basic` as
// `Authorization: Basic <base64 of value>` with `value` in the form
// `user:password`, `token` as `Authorization: token <value>`, and `authorization`
// as `Authorization: <value>`.
type AgentGatewayStaticAuthType string

const (
	AgentGatewayStaticAuthTypeAPIKey        AgentGatewayStaticAuthType = "api_key"
	AgentGatewayStaticAuthTypeBearerToken   AgentGatewayStaticAuthType = "bearer_token"
	AgentGatewayStaticAuthTypeBasic         AgentGatewayStaticAuthType = "basic"
	AgentGatewayStaticAuthTypeToken         AgentGatewayStaticAuthType = "token"
	AgentGatewayStaticAuthTypeAuthorization AgentGatewayStaticAuthType = "authorization"
)

type AgentGatewayAPIListResponse struct {
	Data []AgentGatewayAPI `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayAPIListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayAPIListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayAPINewParams struct {
	// A name for the API, unique on the instance across MCP servers and APIs,
	// regardless of case. Up to 100 letters, digits, `_`, and `.`. The gateway
	// prefixes the API's tools with it and a hyphen.
	Name string `json:"name" api:"required"`
	// The URL of the API's OpenAPI description: OpenAPI 3.0, OpenAPI 3.1, or Swagger
	// 2.0.
	SpecURL string `json:"spec_url" api:"required"`
	// Where the gateway sends the API's requests. Without it, the gateway takes the
	// description's first `servers` entry.
	BaseURL param.Opt[string] `json:"base_url,omitzero"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey param.Opt[bool] `json:"open_to_every_key,omitzero"`
	// How the gateway authenticates to the API. AWS Signature Version 4 is for MCP
	// servers only. Replaced whole on an update.
	Auth AgentGatewayAPINewParamsAuthUnion `json:"auth,omitzero"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers,omitzero"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers,omitzero"`
	// The operations the API exposes as tools, by tool name. `[]` exposes every
	// operation.
	Tools []string `json:"tools,omitzero"`
	paramObj
}

func (r AgentGatewayAPINewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayAPINewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayAPINewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentGatewayAPINewParamsAuthUnion struct {
	OfNoAuthentication       *AgentGatewayNoAuthParam                 `json:",omitzero,inline"`
	OfAPIKey                 *AgentGatewayStaticAuthParam             `json:",omitzero,inline"`
	OfOAuthClientCredentials *AgentGatewayOAuthClientCredentialsParam `json:",omitzero,inline"`
	OfTokenPassthrough       *AgentGatewayPassthroughAuthParam        `json:",omitzero,inline"`
	paramUnion
}

func (u AgentGatewayAPINewParamsAuthUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoAuthentication, u.OfAPIKey, u.OfOAuthClientCredentials, u.OfTokenPassthrough)
}
func (u *AgentGatewayAPINewParamsAuthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentGatewayAPINewParamsAuthUnion) asAny() any {
	if !param.IsOmitted(u.OfNoAuthentication) {
		return u.OfNoAuthentication
	} else if !param.IsOmitted(u.OfAPIKey) {
		return u.OfAPIKey
	} else if !param.IsOmitted(u.OfOAuthClientCredentials) {
		return u.OfOAuthClientCredentials
	} else if !param.IsOmitted(u.OfTokenPassthrough) {
		return u.OfTokenPassthrough
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetValue() *string {
	if vt := u.OfAPIKey; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetHeader() *string {
	if vt := u.OfAPIKey; vt != nil && vt.Header.Valid() {
		return &vt.Header.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetClientID() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetClientSecret() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientSecret
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetTokenURL() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.TokenURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetAudience() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Audience.Valid() {
		return &vt.Audience.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetCredentialsIn() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.CredentialsIn)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetResource() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Resource.Valid() {
		return &vt.Resource.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetScopes() []string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return vt.Scopes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPINewParamsAuthUnion) GetType() *string {
	if vt := u.OfNoAuthentication; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIKey; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTokenPassthrough; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AgentGatewayAPINewParamsAuthUnion](
		"",
		apijson.Variant[AgentGatewayNoAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayStaticAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayOAuthClientCredentialsParam](gjson.JSON),
		apijson.Variant[AgentGatewayPassthroughAuthParam](gjson.JSON),
	)
}

type AgentGatewayAPIUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// Where the gateway sends the API's requests. Without it, the gateway takes the
	// description's first `servers` entry.
	BaseURL param.Opt[string] `json:"base_url,omitzero"`
	// A name for the API, unique on the instance across MCP servers and APIs,
	// regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the
	// API's tools with it and a hyphen.
	Name param.Opt[string] `json:"name,omitzero"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey param.Opt[bool] `json:"open_to_every_key,omitzero"`
	// The URL of the API's OpenAPI description: OpenAPI 3.0, OpenAPI 3.1, or Swagger
	// 2.0.
	SpecURL param.Opt[string] `json:"spec_url,omitzero"`
	// How the gateway authenticates to the API. AWS Signature Version 4 is for MCP
	// servers only. Replaced whole on an update.
	Auth AgentGatewayAPIUpdateParamsAuthUnion `json:"auth,omitzero"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers,omitzero"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers,omitzero"`
	// The operations the API exposes as tools, by tool name. `[]` exposes every
	// operation.
	Tools []string `json:"tools,omitzero"`
	paramObj
}

func (r AgentGatewayAPIUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayAPIUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayAPIUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentGatewayAPIUpdateParamsAuthUnion struct {
	OfNoAuthentication       *AgentGatewayNoAuthParam                 `json:",omitzero,inline"`
	OfAPIKey                 *AgentGatewayStaticAuthParam             `json:",omitzero,inline"`
	OfOAuthClientCredentials *AgentGatewayOAuthClientCredentialsParam `json:",omitzero,inline"`
	OfTokenPassthrough       *AgentGatewayPassthroughAuthParam        `json:",omitzero,inline"`
	paramUnion
}

func (u AgentGatewayAPIUpdateParamsAuthUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoAuthentication, u.OfAPIKey, u.OfOAuthClientCredentials, u.OfTokenPassthrough)
}
func (u *AgentGatewayAPIUpdateParamsAuthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentGatewayAPIUpdateParamsAuthUnion) asAny() any {
	if !param.IsOmitted(u.OfNoAuthentication) {
		return u.OfNoAuthentication
	} else if !param.IsOmitted(u.OfAPIKey) {
		return u.OfAPIKey
	} else if !param.IsOmitted(u.OfOAuthClientCredentials) {
		return u.OfOAuthClientCredentials
	} else if !param.IsOmitted(u.OfTokenPassthrough) {
		return u.OfTokenPassthrough
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetValue() *string {
	if vt := u.OfAPIKey; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetHeader() *string {
	if vt := u.OfAPIKey; vt != nil && vt.Header.Valid() {
		return &vt.Header.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetClientID() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetClientSecret() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientSecret
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetTokenURL() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.TokenURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetAudience() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Audience.Valid() {
		return &vt.Audience.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetCredentialsIn() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.CredentialsIn)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetResource() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Resource.Valid() {
		return &vt.Resource.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetScopes() []string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return vt.Scopes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayAPIUpdateParamsAuthUnion) GetType() *string {
	if vt := u.OfNoAuthentication; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIKey; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTokenPassthrough; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AgentGatewayAPIUpdateParamsAuthUnion](
		"",
		apijson.Variant[AgentGatewayNoAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayStaticAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayOAuthClientCredentialsParam](gjson.JSON),
		apijson.Variant[AgentGatewayPassthroughAuthParam](gjson.JSON),
	)
}

type AgentGatewayAPIDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
