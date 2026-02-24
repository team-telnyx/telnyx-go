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
)

// AuthenticationProviderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthenticationProviderService] method instead.
type AuthenticationProviderService struct {
	Options []option.RequestOption
}

// NewAuthenticationProviderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAuthenticationProviderService(opts ...option.RequestOption) (r AuthenticationProviderService) {
	r = AuthenticationProviderService{}
	r.Options = opts
	return
}

// Creates an authentication provider.
func (r *AuthenticationProviderService) New(ctx context.Context, body AuthenticationProviderNewParams, opts ...option.RequestOption) (res *AuthenticationProviderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "authentication_providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing authentication provider.
func (r *AuthenticationProviderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthenticationProviderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("authentication_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates settings of an existing authentication provider.
func (r *AuthenticationProviderService) Update(ctx context.Context, id string, body AuthenticationProviderUpdateParams, opts ...option.RequestOption) (res *AuthenticationProviderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("authentication_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your SSO authentication providers.
func (r *AuthenticationProviderService) List(ctx context.Context, query AuthenticationProviderListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AuthenticationProvider], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "authentication_providers"
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

// Returns a list of your SSO authentication providers.
func (r *AuthenticationProviderService) ListAutoPaging(ctx context.Context, query AuthenticationProviderListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AuthenticationProvider] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing authentication provider.
func (r *AuthenticationProviderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthenticationProviderDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("authentication_providers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AuthenticationProvider struct {
	// Uniquely identifies the authentication provider.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the authentication provider was
	// activated.
	ActivatedAt time.Time `json:"activated_at" format:"date-time"`
	// The active status of the authentication provider
	Active bool `json:"active"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The name associated with the authentication provider.
	Name string `json:"name"`
	// The id from the Organization the authentication provider belongs to.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The settings associated with the authentication provider.
	Settings AuthenticationProviderSettings `json:"settings"`
	// The short name associated with the authentication provider. This must be unique
	// and URL-friendly, as it's going to be part of the login URL.
	ShortName string `json:"short_name"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ActivatedAt    respjson.Field
		Active         respjson.Field
		CreatedAt      respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		RecordType     respjson.Field
		Settings       respjson.Field
		ShortName      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProvider) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings associated with the authentication provider.
type AuthenticationProviderSettings struct {
	// The Assertion Consumer Service URL for the service provider (Telnyx).
	AssertionConsumerServiceURL string `json:"assertion_consumer_service_url" format:"uri"`
	// Mapping of SAML attribute names used by the identity provider (IdP).
	IdpAttributeNames map[string]any `json:"idp_attribute_names"`
	// The certificate fingerprint for the identity provider (IdP)
	IdpCertFingerprint string `json:"idp_cert_fingerprint"`
	// The algorithm used to generate the identity provider's (IdP) certificate
	// fingerprint
	//
	// Any of "sha1", "sha256", "sha384", "sha512".
	IdpCertFingerprintAlgorithm string `json:"idp_cert_fingerprint_algorithm"`
	// The full X.509 certificate for the identity provider (IdP).
	IdpCertificate string `json:"idp_certificate"`
	// The Entity ID for the identity provider (IdP).
	IdpEntityID string `json:"idp_entity_id" format:"uri"`
	// The Single Logout (SLO) target URL for the identity provider (IdP).
	IdpSloTargetURL string `json:"idp_slo_target_url" format:"uri"`
	// The SSO target url for the identity provider (IdP).
	IdpSSOTargetURL string `json:"idp_sso_target_url" format:"uri"`
	// The name identifier format associated with the authentication provider. This
	// must be the same for both the Identity Provider (IdP) and the service provider
	// (Telnyx).
	NameIdentifierFormat string `json:"name_identifier_format"`
	// Whether group provisioning is enabled for this authentication provider.
	ProvisionGroups bool `json:"provision_groups"`
	// The Entity ID for the service provider (Telnyx).
	ServiceProviderEntityID string `json:"service_provider_entity_id" format:"uri"`
	// The login URL for the service provider (Telnyx). Users navigate to this URL to
	// initiate SSO login.
	ServiceProviderLoginURL string `json:"service_provider_login_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssertionConsumerServiceURL respjson.Field
		IdpAttributeNames           respjson.Field
		IdpCertFingerprint          respjson.Field
		IdpCertFingerprintAlgorithm respjson.Field
		IdpCertificate              respjson.Field
		IdpEntityID                 respjson.Field
		IdpSloTargetURL             respjson.Field
		IdpSSOTargetURL             respjson.Field
		NameIdentifierFormat        respjson.Field
		ProvisionGroups             respjson.Field
		ServiceProviderEntityID     respjson.Field
		ServiceProviderLoginURL     respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProviderSettings) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProviderSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginationMeta struct {
	PageNumber   int64 `json:"page_number,required"`
	TotalPages   int64 `json:"total_pages,required"`
	PageSize     int64 `json:"page_size"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *PaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings associated with the authentication provider.
//
// The properties IdpCertFingerprint, IdpEntityID, IdpSSOTargetURL are required.
type SettingsParam struct {
	// The certificate fingerprint for the identity provider (IdP)
	IdpCertFingerprint string `json:"idp_cert_fingerprint,required"`
	// The Entity ID for the identity provider (IdP).
	IdpEntityID string `json:"idp_entity_id,required" format:"uri"`
	// The SSO target url for the identity provider (IdP).
	IdpSSOTargetURL string `json:"idp_sso_target_url,required" format:"uri"`
	// The algorithm used to generate the identity provider's (IdP) certificate
	// fingerprint
	//
	// Any of "sha1", "sha256", "sha384", "sha512".
	IdpCertFingerprintAlgorithm SettingsIdpCertFingerprintAlgorithm `json:"idp_cert_fingerprint_algorithm,omitzero"`
	paramObj
}

func (r SettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow SettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The algorithm used to generate the identity provider's (IdP) certificate
// fingerprint
type SettingsIdpCertFingerprintAlgorithm string

const (
	SettingsIdpCertFingerprintAlgorithmSha1   SettingsIdpCertFingerprintAlgorithm = "sha1"
	SettingsIdpCertFingerprintAlgorithmSha256 SettingsIdpCertFingerprintAlgorithm = "sha256"
	SettingsIdpCertFingerprintAlgorithmSha384 SettingsIdpCertFingerprintAlgorithm = "sha384"
	SettingsIdpCertFingerprintAlgorithmSha512 SettingsIdpCertFingerprintAlgorithm = "sha512"
)

type AuthenticationProviderNewResponse struct {
	Data AuthenticationProvider `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProviderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProviderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderGetResponse struct {
	Data AuthenticationProvider `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProviderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProviderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderUpdateResponse struct {
	Data AuthenticationProvider `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProviderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProviderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderDeleteResponse struct {
	Data AuthenticationProvider `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthenticationProviderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthenticationProviderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderNewParams struct {
	// The name associated with the authentication provider.
	Name string `json:"name,required"`
	// The settings associated with the authentication provider.
	Settings SettingsParam `json:"settings,omitzero,required"`
	// The short name associated with the authentication provider. This must be unique
	// and URL-friendly, as it's going to be part of the login URL.
	ShortName string `json:"short_name,required"`
	// The active status of the authentication provider
	Active param.Opt[bool] `json:"active,omitzero"`
	// The URL for the identity provider metadata file to populate the settings
	// automatically. If the settings attribute is provided, that will be used instead.
	SettingsURL param.Opt[string] `json:"settings_url,omitzero" format:"uri"`
	paramObj
}

func (r AuthenticationProviderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthenticationProviderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthenticationProviderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderUpdateParams struct {
	// The active status of the authentication provider
	Active param.Opt[bool] `json:"active,omitzero"`
	// The name associated with the authentication provider.
	Name param.Opt[string] `json:"name,omitzero"`
	// The URL for the identity provider metadata file to populate the settings
	// automatically. If the settings attribute is provided, that will be used instead.
	SettingsURL param.Opt[string] `json:"settings_url,omitzero" format:"uri"`
	// The short name associated with the authentication provider. This must be unique
	// and URL-friendly, as it's going to be part of the login URL.
	ShortName param.Opt[string] `json:"short_name,omitzero"`
	// The settings associated with the authentication provider.
	Settings SettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r AuthenticationProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthenticationProviderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthenticationProviderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthenticationProviderListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code>-</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>name</code>: sorts the result by the
	//	  <code>name</code> field in ascending order.
	//	</li>
	//	<li>
	//	  <code>-name</code>: sorts the result by the
	//	  <code>name</code> field in descending order.
	//	</li>
	//
	// </ul><br/>If not given, results are sorted by <code>created_at</code> in descending order.
	//
	// Any of "name", "-name", "short_name", "-short_name", "active", "-active",
	// "created_at", "-created_at", "updated_at", "-updated_at".
	Sort AuthenticationProviderListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthenticationProviderListParams]'s query parameters as
// `url.Values`.
func (r AuthenticationProviderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. By default sorting direction is ascending.
// To have the results sorted in descending order add the <code>-</code>
// prefix.<br/><br/> That is: <ul>
//
//	<li>
//	  <code>name</code>: sorts the result by the
//	  <code>name</code> field in ascending order.
//	</li>
//	<li>
//	  <code>-name</code>: sorts the result by the
//	  <code>name</code> field in descending order.
//	</li>
//
// </ul><br/>If not given, results are sorted by <code>created_at</code> in descending order.
type AuthenticationProviderListParamsSort string

const (
	AuthenticationProviderListParamsSortName          AuthenticationProviderListParamsSort = "name"
	AuthenticationProviderListParamsSortNameDesc      AuthenticationProviderListParamsSort = "-name"
	AuthenticationProviderListParamsSortShortName     AuthenticationProviderListParamsSort = "short_name"
	AuthenticationProviderListParamsSortShortNameDesc AuthenticationProviderListParamsSort = "-short_name"
	AuthenticationProviderListParamsSortActive        AuthenticationProviderListParamsSort = "active"
	AuthenticationProviderListParamsSortActiveDesc    AuthenticationProviderListParamsSort = "-active"
	AuthenticationProviderListParamsSortCreatedAt     AuthenticationProviderListParamsSort = "created_at"
	AuthenticationProviderListParamsSortCreatedAtDesc AuthenticationProviderListParamsSort = "-created_at"
	AuthenticationProviderListParamsSortUpdatedAt     AuthenticationProviderListParamsSort = "updated_at"
	AuthenticationProviderListParamsSortUpdatedAtDesc AuthenticationProviderListParamsSort = "-updated_at"
)
