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

// MCP servers on an instance, served to agents at one MCP endpoint.
//
// AgentGatewayMcpServerService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayMcpServerService] method instead.
type AgentGatewayMcpServerService struct {
	Options []option.RequestOption
}

// NewAgentGatewayMcpServerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentGatewayMcpServerService(opts ...option.RequestOption) (r AgentGatewayMcpServerService) {
	r = AgentGatewayMcpServerService{}
	r.Options = opts
	return
}

// Adds an MCP server. The instance serves its tools at `{gateway_url}/mcp`, each
// named with the server's `name` and a hyphen, to the keys granted the server and,
// with `open_to_every_key`, to every key. The instance must be `ready`; otherwise
// the request answers `409`.
func (r *AgentGatewayMcpServerService) New(ctx context.Context, instanceID string, body AgentGatewayMcpServerNewParams, opts ...option.RequestOption) (res *AgentGatewayMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/mcp_servers", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes an MCP server's settings. A field left out stays as it is. `auth` is
// replaced whole, `tools: []` exposes every tool, and `static_headers: {}` clears
// the headers. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayMcpServerService) Update(ctx context.Context, serverID string, params AgentGatewayMcpServerUpdateParams, opts ...option.RequestOption) (res *AgentGatewayMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/mcp_servers/%s", params.InstanceID, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the MCP servers the instance serves at `{gateway_url}/mcp`, without their
// credentials. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayMcpServerService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayMcpServerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/mcp_servers", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes an MCP server and returns it as it was. Its `id` stays in the tool
// permissions that name it until they are next set. The instance must be `ready`;
// otherwise the request answers `409`.
func (r *AgentGatewayMcpServerService) Delete(ctx context.Context, serverID string, body AgentGatewayMcpServerDeleteParams, opts ...option.RequestOption) (res *AgentGatewayMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if serverID == "" {
		err = errors.New("missing required server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/mcp_servers/%s", body.InstanceID, serverID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The gateway signs each request with AWS Signature Version 4. Without keys or a
// role, it uses the instance's own AWS identity. Keys are write-only.
//
// The properties Region, Type are required.
type AgentGatewayAwsSigV4AuthParam struct {
	// The AWS region to sign for.
	Region      string            `json:"region" api:"required"`
	AccessKeyID param.Opt[string] `json:"access_key_id,omitzero"`
	// The ARN of a role for the gateway to assume.
	RoleName        param.Opt[string] `json:"role_name,omitzero"`
	SecretAccessKey param.Opt[string] `json:"secret_access_key,omitzero"`
	SessionToken    param.Opt[string] `json:"session_token,omitzero"`
	// The AWS service to sign for.
	Service param.Opt[string] `json:"service,omitzero"`
	// This field can be elided, and will marshal its zero value as "aws_sigv4".
	Type constant.AwsSigv4 `json:"type" default:"aws_sigv4"`
	paramObj
}

func (r AgentGatewayAwsSigV4AuthParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayAwsSigV4AuthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayAwsSigV4AuthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayMcpServer struct {
	// Identifies the server.
	ID string `json:"id" api:"required" format:"uuid"`
	// The server's outbound authentication method. Credentials are never returned.
	Auth      AgentGatewayShownAuth `json:"auth" api:"required"`
	CreatedAt time.Time             `json:"created_at" api:"required" format:"date-time"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers" api:"required"`
	// A name for the server, unique on the instance across MCP servers and APIs,
	// regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the
	// server's tools with it and a hyphen.
	Name string `json:"name" api:"required"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey bool `json:"open_to_every_key" api:"required"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers" api:"required"`
	// The tools the server exposes to every key, by name. `[]` exposes every tool.
	Tools []string `json:"tools" api:"required"`
	// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
	//
	// Any of "http", "sse".
	Transport AgentGatewayMcpServerTransport `json:"transport" api:"required"`
	UpdatedAt time.Time                      `json:"updated_at" api:"required" format:"date-time"`
	// The server's MCP endpoint.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Auth             respjson.Field
		CreatedAt        respjson.Field
		ForwardedHeaders respjson.Field
		Name             respjson.Field
		OpenToEveryKey   respjson.Field
		StaticHeaders    respjson.Field
		Tools            respjson.Field
		Transport        respjson.Field
		UpdatedAt        respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayMcpServer) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayMcpServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
type AgentGatewayMcpServerTransport string

const (
	AgentGatewayMcpServerTransportHTTP AgentGatewayMcpServerTransport = "http"
	AgentGatewayMcpServerTransportSse  AgentGatewayMcpServerTransport = "sse"
)

type AgentGatewayMcpServerResponse struct {
	Data AgentGatewayMcpServer `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayMcpServerResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayMcpServerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayMcpServerListResponse struct {
	Data []AgentGatewayMcpServer `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayMcpServerListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayMcpServerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayMcpServerNewParams struct {
	// A name for the server, unique on the instance across MCP servers and APIs,
	// regardless of case. Up to 100 letters, digits, `_`, and `.`. The gateway
	// prefixes the server's tools with it and a hyphen.
	Name string `json:"name" api:"required"`
	// The server's MCP endpoint, `http` or `https`.
	URL string `json:"url" api:"required"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey param.Opt[bool] `json:"open_to_every_key,omitzero"`
	// How the gateway authenticates to the server. Replaced whole on an update;
	// `{"type": "none"}` removes the credential.
	Auth AgentGatewayMcpServerNewParamsAuthUnion `json:"auth,omitzero"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers,omitzero"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers,omitzero"`
	// The tools the server exposes to every key, by name. `[]` exposes every tool.
	Tools []string `json:"tools,omitzero"`
	// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
	//
	// Any of "http", "sse".
	Transport AgentGatewayMcpServerNewParamsTransport `json:"transport,omitzero"`
	paramObj
}

func (r AgentGatewayMcpServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayMcpServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayMcpServerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentGatewayMcpServerNewParamsAuthUnion struct {
	OfNoAuthentication       *AgentGatewayNoAuthParam                 `json:",omitzero,inline"`
	OfAPIKey                 *AgentGatewayStaticAuthParam             `json:",omitzero,inline"`
	OfOAuthClientCredentials *AgentGatewayOAuthClientCredentialsParam `json:",omitzero,inline"`
	OfTokenPassthrough       *AgentGatewayPassthroughAuthParam        `json:",omitzero,inline"`
	OfAwsSignatureVersion4   *AgentGatewayAwsSigV4AuthParam           `json:",omitzero,inline"`
	paramUnion
}

func (u AgentGatewayMcpServerNewParamsAuthUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoAuthentication,
		u.OfAPIKey,
		u.OfOAuthClientCredentials,
		u.OfTokenPassthrough,
		u.OfAwsSignatureVersion4)
}
func (u *AgentGatewayMcpServerNewParamsAuthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentGatewayMcpServerNewParamsAuthUnion) asAny() any {
	if !param.IsOmitted(u.OfNoAuthentication) {
		return u.OfNoAuthentication
	} else if !param.IsOmitted(u.OfAPIKey) {
		return u.OfAPIKey
	} else if !param.IsOmitted(u.OfOAuthClientCredentials) {
		return u.OfOAuthClientCredentials
	} else if !param.IsOmitted(u.OfTokenPassthrough) {
		return u.OfTokenPassthrough
	} else if !param.IsOmitted(u.OfAwsSignatureVersion4) {
		return u.OfAwsSignatureVersion4
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetValue() *string {
	if vt := u.OfAPIKey; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetHeader() *string {
	if vt := u.OfAPIKey; vt != nil && vt.Header.Valid() {
		return &vt.Header.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetClientID() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetClientSecret() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientSecret
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetTokenURL() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.TokenURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetAudience() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Audience.Valid() {
		return &vt.Audience.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetCredentialsIn() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.CredentialsIn)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetResource() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Resource.Valid() {
		return &vt.Resource.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetScopes() []string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return vt.Scopes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetRegion() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil {
		return &vt.Region
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetAccessKeyID() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetRoleName() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.RoleName.Valid() {
		return &vt.RoleName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetSecretAccessKey() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetService() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.Service.Valid() {
		return &vt.Service.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetSessionToken() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.SessionToken.Valid() {
		return &vt.SessionToken.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerNewParamsAuthUnion) GetType() *string {
	if vt := u.OfNoAuthentication; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIKey; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTokenPassthrough; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAwsSignatureVersion4; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AgentGatewayMcpServerNewParamsAuthUnion](
		"",
		apijson.Variant[AgentGatewayNoAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayStaticAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayOAuthClientCredentialsParam](gjson.JSON),
		apijson.Variant[AgentGatewayPassthroughAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayAwsSigV4AuthParam](gjson.JSON),
	)
}

// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
type AgentGatewayMcpServerNewParamsTransport string

const (
	AgentGatewayMcpServerNewParamsTransportHTTP AgentGatewayMcpServerNewParamsTransport = "http"
	AgentGatewayMcpServerNewParamsTransportSse  AgentGatewayMcpServerNewParamsTransport = "sse"
)

type AgentGatewayMcpServerUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// A name for the server, unique on the instance across MCP servers and APIs,
	// regardless of case. Letters, digits, `_`, and `.`. The gateway prefixes the
	// server's tools with it and a hyphen.
	Name param.Opt[string] `json:"name,omitzero"`
	// `true` lets every key on the instance use the server, granted or not.
	OpenToEveryKey param.Opt[bool] `json:"open_to_every_key,omitzero"`
	// The server's MCP endpoint, `http` or `https`.
	URL param.Opt[string] `json:"url,omitzero"`
	// How the gateway authenticates to the server. Replaced whole on an update;
	// `{"type": "none"}` removes the credential.
	Auth AgentGatewayMcpServerUpdateParamsAuthUnion `json:"auth,omitzero"`
	// Headers copied from the agent's request, by name.
	ForwardedHeaders []string `json:"forwarded_headers,omitzero"`
	// Headers sent on every request.
	StaticHeaders map[string]string `json:"static_headers,omitzero"`
	// The tools the server exposes to every key, by name. `[]` exposes every tool.
	Tools []string `json:"tools,omitzero"`
	// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
	//
	// Any of "http", "sse".
	Transport AgentGatewayMcpServerUpdateParamsTransport `json:"transport,omitzero"`
	paramObj
}

func (r AgentGatewayMcpServerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayMcpServerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayMcpServerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentGatewayMcpServerUpdateParamsAuthUnion struct {
	OfNoAuthentication       *AgentGatewayNoAuthParam                 `json:",omitzero,inline"`
	OfAPIKey                 *AgentGatewayStaticAuthParam             `json:",omitzero,inline"`
	OfOAuthClientCredentials *AgentGatewayOAuthClientCredentialsParam `json:",omitzero,inline"`
	OfTokenPassthrough       *AgentGatewayPassthroughAuthParam        `json:",omitzero,inline"`
	OfAwsSignatureVersion4   *AgentGatewayAwsSigV4AuthParam           `json:",omitzero,inline"`
	paramUnion
}

func (u AgentGatewayMcpServerUpdateParamsAuthUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNoAuthentication,
		u.OfAPIKey,
		u.OfOAuthClientCredentials,
		u.OfTokenPassthrough,
		u.OfAwsSignatureVersion4)
}
func (u *AgentGatewayMcpServerUpdateParamsAuthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentGatewayMcpServerUpdateParamsAuthUnion) asAny() any {
	if !param.IsOmitted(u.OfNoAuthentication) {
		return u.OfNoAuthentication
	} else if !param.IsOmitted(u.OfAPIKey) {
		return u.OfAPIKey
	} else if !param.IsOmitted(u.OfOAuthClientCredentials) {
		return u.OfOAuthClientCredentials
	} else if !param.IsOmitted(u.OfTokenPassthrough) {
		return u.OfTokenPassthrough
	} else if !param.IsOmitted(u.OfAwsSignatureVersion4) {
		return u.OfAwsSignatureVersion4
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetValue() *string {
	if vt := u.OfAPIKey; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetHeader() *string {
	if vt := u.OfAPIKey; vt != nil && vt.Header.Valid() {
		return &vt.Header.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetClientID() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetClientSecret() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.ClientSecret
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetTokenURL() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return &vt.TokenURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetAudience() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Audience.Valid() {
		return &vt.Audience.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetCredentialsIn() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.CredentialsIn)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetResource() *string {
	if vt := u.OfOAuthClientCredentials; vt != nil && vt.Resource.Valid() {
		return &vt.Resource.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetScopes() []string {
	if vt := u.OfOAuthClientCredentials; vt != nil {
		return vt.Scopes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetRegion() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil {
		return &vt.Region
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetAccessKeyID() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.AccessKeyID.Valid() {
		return &vt.AccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetRoleName() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.RoleName.Valid() {
		return &vt.RoleName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetSecretAccessKey() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.SecretAccessKey.Valid() {
		return &vt.SecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetService() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.Service.Valid() {
		return &vt.Service.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetSessionToken() *string {
	if vt := u.OfAwsSignatureVersion4; vt != nil && vt.SessionToken.Valid() {
		return &vt.SessionToken.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentGatewayMcpServerUpdateParamsAuthUnion) GetType() *string {
	if vt := u.OfNoAuthentication; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIKey; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOAuthClientCredentials; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTokenPassthrough; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAwsSignatureVersion4; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AgentGatewayMcpServerUpdateParamsAuthUnion](
		"",
		apijson.Variant[AgentGatewayNoAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayStaticAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayOAuthClientCredentialsParam](gjson.JSON),
		apijson.Variant[AgentGatewayPassthroughAuthParam](gjson.JSON),
		apijson.Variant[AgentGatewayAwsSigV4AuthParam](gjson.JSON),
	)
}

// `http` for Streamable HTTP, or `sse` for Server-Sent Events.
type AgentGatewayMcpServerUpdateParamsTransport string

const (
	AgentGatewayMcpServerUpdateParamsTransportHTTP AgentGatewayMcpServerUpdateParamsTransport = "http"
	AgentGatewayMcpServerUpdateParamsTransportSse  AgentGatewayMcpServerUpdateParamsTransport = "sse"
)

type AgentGatewayMcpServerDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
