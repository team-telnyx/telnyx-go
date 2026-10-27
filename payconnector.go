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
	"github.com/tidwall/gjson"
)

// Generic payment connector configuration.
//
// PayConnectorService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayConnectorService] method instead.
type PayConnectorService struct {
	Options []option.RequestOption
}

// NewPayConnectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPayConnectorService(opts ...option.RequestOption) (r PayConnectorService) {
	r = PayConnectorService{}
	r.Options = opts
	return
}

// Create a generic payment connector in the authenticated organization. The name
// must be unique in that organization. Authentication supplies ownership; do not
// send user_id or organization_id. Config is encrypted at rest; passwords and
// tokens are omitted from responses. This operation stores configuration and does
// not initiate a payment. See the
// [processor contract](/docs/voice/programmable-voice/pay#request-sent-to-the-payment-processor).
func (r *PayConnectorService) New(ctx context.Context, body PayConnectorNewParams, opts ...option.RequestOption) (res *PayConnectorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pay_connectors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a connector by name within the authenticated organization. Password and
// token are omitted. A missing name or a connector owned by another organization
// returns 404.
func (r *PayConnectorService) Get(ctx context.Context, name string, opts ...option.RequestOption) (res *PayConnectorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("pay_connectors/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a generic connector by name in the authenticated organization. Both test
// and live connectors can be updated through the API. Omitted fields remain
// unchanged. Supplying config replaces the entire config; it does not merge keys.
// Include endpoint_url and the credentials required for the selected auth_type on
// every config replacement. Omitting config preserves existing credentials,
// allowing a mode-only update. The path name cannot be renamed. Public updates use
// PATCH.
func (r *PayConnectorService) Update(ctx context.Context, name string, body PayConnectorUpdateParams, opts ...option.RequestOption) (res *PayConnectorResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("pay_connectors/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List connectors belonging to the authenticated organization, including
// connectors created by other users in that organization. Set filter[type]=generic
// to restrict results to the type covered here. Results default to newest
// created_at first. Use meta.end_cursor as page[after] for the next page, or
// meta.start_cursor as page[before] for the preceding page. Omit both cursors for
// the first page; if both are supplied, page[after] takes precedence. Responses
// omit passwords and tokens.
func (r *PayConnectorService) List(ctx context.Context, query PayConnectorListParams, opts ...option.RequestOption) (res *pagination.EmailBracketCursorPagination[PayConnector], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pay_connectors"
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

// List connectors belonging to the authenticated organization, including
// connectors created by other users in that organization. Set filter[type]=generic
// to restrict results to the type covered here. Results default to newest
// created_at first. Use meta.end_cursor as page[after] for the next page, or
// meta.start_cursor as page[before] for the preceding page. Omit both cursors for
// the first page; if both are supplied, page[after] takes precedence. Responses
// omit passwords and tokens.
func (r *PayConnectorService) ListAutoPaging(ctx context.Context, query PayConnectorListParams, opts ...option.RequestOption) *pagination.EmailBracketCursorPaginationAutoPager[PayConnector] {
	return pagination.NewEmailBracketCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a connector by name in the authenticated organization. Both test and live
// connectors can be deleted. Returns no response body on success; a missing name
// or a connector in another organization returns 404.
func (r *PayConnectorService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return err
	}
	path := fmt.Sprintf("pay_connectors/%s", url.PathEscape(name))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func GenericPayConnectorConfigParamOfBasicAuthenticationAlsoTheValidationDefault(password string, username string) GenericPayConnectorConfigUnionParam {
	var variant GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam
	variant.Password = password
	variant.Username = username
	return GenericPayConnectorConfigUnionParam{OfBasicAuthenticationAlsoTheValidationDefault: &variant}
}

func GenericPayConnectorConfigParamOfBearerAuthentication(token string) GenericPayConnectorConfigUnionParam {
	var variant GenericPayConnectorConfigBearerAuthenticationParam
	variant.Token = token
	return GenericPayConnectorConfigUnionParam{OfBearerAuthentication: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GenericPayConnectorConfigUnionParam struct {
	OfBasicAuthenticationAlsoTheValidationDefault *GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam `json:",omitzero,inline"`
	OfBearerAuthentication                        *GenericPayConnectorConfigBearerAuthenticationParam                        `json:",omitzero,inline"`
	OfNoProcessorAuthentication                   *GenericPayConnectorConfigNoProcessorAuthenticationParam                   `json:",omitzero,inline"`
	paramUnion
}

func (u GenericPayConnectorConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBasicAuthenticationAlsoTheValidationDefault, u.OfBearerAuthentication, u.OfNoProcessorAuthentication)
}
func (u *GenericPayConnectorConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GenericPayConnectorConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBasicAuthenticationAlsoTheValidationDefault) {
		return u.OfBasicAuthenticationAlsoTheValidationDefault
	} else if !param.IsOmitted(u.OfBearerAuthentication) {
		return u.OfBearerAuthentication
	} else if !param.IsOmitted(u.OfNoProcessorAuthentication) {
		return u.OfNoProcessorAuthentication
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenericPayConnectorConfigUnionParam) GetAuthType() *string {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil {
		return (*string)(&vt.AuthType)
	} else if vt := u.OfBearerAuthentication; vt != nil {
		return (*string)(&vt.AuthType)
	} else if vt := u.OfNoProcessorAuthentication; vt != nil {
		return (*string)(&vt.AuthType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenericPayConnectorConfigUnionParam) GetEndpointURL() *string {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil && vt.EndpointURL.Valid() {
		return &vt.EndpointURL.Value
	} else if vt := u.OfBearerAuthentication; vt != nil && vt.EndpointURL.Valid() {
		return &vt.EndpointURL.Value
	} else if vt := u.OfNoProcessorAuthentication; vt != nil && vt.EndpointURL.Valid() {
		return &vt.EndpointURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GenericPayConnectorConfigUnionParam) GetTimeout() *int64 {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfBearerAuthentication; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfNoProcessorAuthentication; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u GenericPayConnectorConfigUnionParam) GetPassword() (res genericPayConnectorConfigUnionParamPassword) {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil {
		res.any = &vt.Password
	} else if vt := u.OfBearerAuthentication; vt != nil {
		res.any = &vt.Password
	} else if vt := u.OfNoProcessorAuthentication; vt != nil {
		res.any = &vt.Password
	}
	return
}

// Can have the runtime types [*string], [*any]
type genericPayConnectorConfigUnionParamPassword struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *any:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u genericPayConnectorConfigUnionParamPassword) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u GenericPayConnectorConfigUnionParam) GetUsername() (res genericPayConnectorConfigUnionParamUsername) {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil {
		res.any = &vt.Username
	} else if vt := u.OfBearerAuthentication; vt != nil {
		res.any = &vt.Username
	} else if vt := u.OfNoProcessorAuthentication; vt != nil {
		res.any = &vt.Username
	}
	return
}

// Can have the runtime types [*string], [*any]
type genericPayConnectorConfigUnionParamUsername struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *any:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u genericPayConnectorConfigUnionParamUsername) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u GenericPayConnectorConfigUnionParam) GetToken() (res genericPayConnectorConfigUnionParamToken) {
	if vt := u.OfBasicAuthenticationAlsoTheValidationDefault; vt != nil {
		res.any = &vt.Token
	} else if vt := u.OfBearerAuthentication; vt != nil {
		res.any = &vt.Token
	} else if vt := u.OfNoProcessorAuthentication; vt != nil {
		res.any = &vt.Token
	}
	return
}

// Can have the runtime types [*any], [*string]
type genericPayConnectorConfigUnionParamToken struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *any:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u genericPayConnectorConfigUnionParamToken) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[GenericPayConnectorConfigUnionParam](
		"",
		apijson.Variant[GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam](gjson.JSON),
		apijson.Variant[GenericPayConnectorConfigBearerAuthenticationParam](gjson.JSON),
		apijson.Variant[GenericPayConnectorConfigNoProcessorAuthenticationParam](gjson.JSON),
	)
}

// The properties Password, Username are required.
type GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam struct {
	// Basic password.
	Password string `json:"password" api:"required"`
	// Basic username.
	Username string `json:"username" api:"required"`
	// Processor HTTP timeout in milliseconds. Omitted or null uses 10000; null is
	// removed before storage. Not the IVR digit-gather timeout.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// HTTP or HTTPS processor URL with a host. Use HTTPS to protect payment data in
	// transit.
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	// Nonblank string required for bearer authentication. Write-only; omitted from
	// responses. Re-send when replacing a bearer-auth config.
	Token any `json:"token,omitzero"`
	// Any of "basic".
	AuthType string `json:"auth_type,omitzero"`
	paramObj
}

func (r GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GenericPayConnectorConfigBasicAuthenticationAlsoTheValidationDefaultParam](
		"auth_type", "basic",
	)
}

// The properties Token, AuthType are required.
type GenericPayConnectorConfigBearerAuthenticationParam struct {
	// Processor bearer token.
	Token string `json:"token" api:"required"`
	// Processor HTTP timeout in milliseconds. Omitted or null uses 10000; null is
	// removed before storage. Not the IVR digit-gather timeout.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// HTTP or HTTPS processor URL with a host. Use HTTPS to protect payment data in
	// transit.
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	// Nonblank string required for basic authentication. Write-only; omitted from
	// responses. Re-send when replacing a basic-auth config.
	Password any `json:"password,omitzero"`
	// Nonblank string required for basic authentication. Returned in connector
	// responses.
	Username any `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "bearer".
	AuthType constant.Bearer `json:"auth_type" default:"bearer"`
	paramObj
}

func (r GenericPayConnectorConfigBearerAuthenticationParam) MarshalJSON() (data []byte, err error) {
	type shadow GenericPayConnectorConfigBearerAuthenticationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenericPayConnectorConfigBearerAuthenticationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AuthType is required.
type GenericPayConnectorConfigNoProcessorAuthenticationParam struct {
	// Processor HTTP timeout in milliseconds. Omitted or null uses 10000; null is
	// removed before storage. Not the IVR digit-gather timeout.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// HTTP or HTTPS processor URL with a host. Use HTTPS to protect payment data in
	// transit.
	EndpointURL param.Opt[string] `json:"endpoint_url,omitzero"`
	// Nonblank string required for bearer authentication. Write-only; omitted from
	// responses. Re-send when replacing a bearer-auth config.
	Token any `json:"token,omitzero"`
	// Nonblank string required for basic authentication. Write-only; omitted from
	// responses. Re-send when replacing a basic-auth config.
	Password any `json:"password,omitzero"`
	// Nonblank string required for basic authentication. Returned in connector
	// responses.
	Username any `json:"username,omitzero"`
	// This field can be elided, and will marshal its zero value as "none".
	AuthType constant.None `json:"auth_type" default:"none"`
	paramObj
}

func (r GenericPayConnectorConfigNoProcessorAuthenticationParam) MarshalJSON() (data []byte, err error) {
	type shadow GenericPayConnectorConfigNoProcessorAuthenticationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenericPayConnectorConfigNoProcessorAuthenticationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization-owned connector. Examples and config properties here describe
// generic connectors; listing is not restricted to that type unless
// filter[type]=generic is supplied. Passwords and bearer tokens are omitted, not
// masked. Defaults for auth_type and timeout need not appear in stored config.
type PayConnector struct {
	// Connector identifier.
	ID string `json:"id" format:"uuid"`
	// Redacted stored configuration. Generic connector fields are shown here. Unknown
	// config keys are discarded. Password and token are never returned.
	Config PayConnectorConfig `json:"config"`
	// Creation time.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Configured mode.
	//
	// Any of "test", "live".
	Mode PayConnectorMode `json:"mode"`
	// Connector name used for retrieval, update, deletion and Pay.
	Name string `json:"name"`
	// Owning organization; assigned from authentication, not request input.
	OrganizationID string `json:"organization_id"`
	// Resource type.
	//
	// Any of "pay_connector".
	RecordType PayConnectorRecordType `json:"record_type"`
	// Connector type; generic is covered by this reference. Other types may appear in
	// unfiltered lists.
	Type string `json:"type"`
	// Last update time.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Creating user; assigned from authentication, not request input.
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Config         respjson.Field
		CreatedAt      respjson.Field
		Mode           respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		RecordType     respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayConnector) RawJSON() string { return r.JSON.raw }
func (r *PayConnector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Redacted stored configuration. Generic connector fields are shown here. Unknown
// config keys are discarded. Password and token are never returned.
type PayConnectorConfig struct {
	// Stored processor authentication selection. Omitted if it was not explicitly
	// configured.
	//
	// Any of "none", "basic", "bearer".
	AuthType string `json:"auth_type"`
	// HTTP or HTTPS processor URL with a host. Use HTTPS to protect payment data in
	// transit.
	EndpointURL string `json:"endpoint_url"`
	// Stored processor HTTP timeout in milliseconds. Omitted when no timeout was
	// configured; the effective runtime default is 10000.
	Timeout int64 `json:"timeout"`
	// Stored username, when supplied. Required to be a nonblank string for basic
	// authentication.
	Username any `json:"username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthType    respjson.Field
		EndpointURL respjson.Field
		Timeout     respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayConnectorConfig) RawJSON() string { return r.JSON.raw }
func (r *PayConnectorConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured mode.
type PayConnectorMode string

const (
	PayConnectorModeTest PayConnectorMode = "test"
	PayConnectorModeLive PayConnectorMode = "live"
)

// Resource type.
type PayConnectorRecordType string

const (
	PayConnectorRecordTypePayConnector PayConnectorRecordType = "pay_connector"
)

type PayConnectorResponse struct {
	// Organization-owned connector. Examples and config properties here describe
	// generic connectors; listing is not restricted to that type unless
	// filter[type]=generic is supplied. Passwords and bearer tokens are omitted, not
	// masked. Defaults for auth_type and timeout need not appear in stored config.
	Data PayConnector `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayConnectorResponse) RawJSON() string { return r.JSON.raw }
func (r *PayConnectorResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PayConnectorNewParams struct {
	// Configuration for a generic connector. A config update replaces the complete
	// object, not individual keys. Unknown config keys are discarded. Authentication
	// credentials must be nonblank strings when required by the selected auth_type.
	Config GenericPayConnectorConfigUnionParam `json:"config,omitzero" api:"required"`
	// Unique within the authenticated organization. Used in resource paths and as
	// connector_name in the Call Control Pay request. Cannot be renamed by update.
	Name string `json:"name" api:"required"`
	// Defaults to test on creation. Test restricts collected card numbers but still
	// contacts the configured processor; use a processor sandbox URL and sandbox
	// credentials. Live removes the test-card restriction. Neither mode locks API
	// updates.
	//
	// Any of "test", "live".
	Mode PayConnectorNewParamsMode `json:"mode,omitzero"`
	// Connector type covered by this reference.
	//
	// Any of "generic".
	Type PayConnectorNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r PayConnectorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PayConnectorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayConnectorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defaults to test on creation. Test restricts collected card numbers but still
// contacts the configured processor; use a processor sandbox URL and sandbox
// credentials. Live removes the test-card restriction. Neither mode locks API
// updates.
type PayConnectorNewParamsMode string

const (
	PayConnectorNewParamsModeTest PayConnectorNewParamsMode = "test"
	PayConnectorNewParamsModeLive PayConnectorNewParamsMode = "live"
)

// Connector type covered by this reference.
type PayConnectorNewParamsType string

const (
	PayConnectorNewParamsTypeGeneric PayConnectorNewParamsType = "generic"
)

type PayConnectorUpdateParams struct {
	// Configuration for a generic connector. A config update replaces the complete
	// object, not individual keys. Unknown config keys are discarded. Authentication
	// credentials must be nonblank strings when required by the selected auth_type.
	Config GenericPayConnectorConfigUnionParam `json:"config,omitzero"`
	// Omit to preserve the existing mode. Use an explicit test or live value.
	//
	// Any of "test", "live".
	Mode PayConnectorUpdateParamsMode `json:"mode,omitzero"`
	// Generic connector type. Omit to preserve the existing value.
	//
	// Any of "generic".
	Type PayConnectorUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r PayConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PayConnectorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayConnectorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Omit to preserve the existing mode. Use an explicit test or live value.
type PayConnectorUpdateParamsMode string

const (
	PayConnectorUpdateParamsModeTest PayConnectorUpdateParamsMode = "test"
	PayConnectorUpdateParamsModeLive PayConnectorUpdateParamsMode = "live"
)

// Generic connector type. Omit to preserve the existing value.
type PayConnectorUpdateParamsType string

const (
	PayConnectorUpdateParamsTypeGeneric PayConnectorUpdateParamsType = "generic"
)

type PayConnectorListParams struct {
	// Equal to this created_at timestamp.
	FilterCreatedAtEq param.Opt[string] `query:"filter[created_at][eq],omitzero" json:"-"`
	// After this created_at timestamp.
	FilterCreatedAtGt param.Opt[string] `query:"filter[created_at][gt],omitzero" json:"-"`
	// At or after this created_at timestamp.
	FilterCreatedAtGte param.Opt[string] `query:"filter[created_at][gte],omitzero" json:"-"`
	// Before this created_at timestamp.
	FilterCreatedAtLt param.Opt[string] `query:"filter[created_at][lt],omitzero" json:"-"`
	// At or before this created_at timestamp.
	FilterCreatedAtLte param.Opt[string] `query:"filter[created_at][lte],omitzero" json:"-"`
	// Exact match on mode.
	FilterMode param.Opt[string] `query:"filter[mode],omitzero" json:"-"`
	// Exact match on name.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Exact match on type.
	FilterType param.Opt[string] `query:"filter[type],omitzero" json:"-"`
	// Equal to this updated_at timestamp.
	FilterUpdatedAtEq param.Opt[string] `query:"filter[updated_at][eq],omitzero" json:"-"`
	// After this updated_at timestamp.
	FilterUpdatedAtGt param.Opt[string] `query:"filter[updated_at][gt],omitzero" json:"-"`
	// At or after this updated_at timestamp.
	FilterUpdatedAtGte param.Opt[string] `query:"filter[updated_at][gte],omitzero" json:"-"`
	// Before this updated_at timestamp.
	FilterUpdatedAtLt param.Opt[string] `query:"filter[updated_at][lt],omitzero" json:"-"`
	// At or before this updated_at timestamp.
	FilterUpdatedAtLte param.Opt[string] `query:"filter[updated_at][lte],omitzero" json:"-"`
	// Opaque cursor from meta.end_cursor.
	PageAfter param.Opt[string] `query:"page[after],omitzero" json:"-"`
	// Opaque cursor from meta.start_cursor.
	PageBefore param.Opt[string] `query:"page[before],omitzero" json:"-"`
	// Number of records per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Comma-separated created_at or updated_at fields, optionally prefixed with - for
	// descending order.
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PayConnectorListParams]'s query parameters as `url.Values`.
func (r PayConnectorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
