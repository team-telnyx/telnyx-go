// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// FQDN connection operations
//
// FqdnConnectionFqdnAuthenticationService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFqdnConnectionFqdnAuthenticationService] method instead.
type FqdnConnectionFqdnAuthenticationService struct {
	Options []option.RequestOption
}

// NewFqdnConnectionFqdnAuthenticationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFqdnConnectionFqdnAuthenticationService(opts ...option.RequestOption) (r FqdnConnectionFqdnAuthenticationService) {
	r = FqdnConnectionFqdnAuthenticationService{}
	r.Options = opts
	return
}

// Retrieves the details of an existing FQDN authentication strategy for a specific
// FQDN connection.
func (r *FqdnConnectionFqdnAuthenticationService) List(ctx context.Context, fqdnConnectionID string, opts ...option.RequestOption) (res *FqdnConnectionFqdnAuthenticationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fqdnConnectionID == "" {
		err = errors.New("missing required fqdn_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("fqdn_connections/%s/fqdn_authentication", fqdnConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the FQDN authentication strategy for a specific FQDN connection.
func (r *FqdnConnectionFqdnAuthenticationService) PatchAll(ctx context.Context, fqdnConnectionID string, body FqdnConnectionFqdnAuthenticationPatchAllParams, opts ...option.RequestOption) (res *FqdnConnectionFqdnAuthenticationPatchAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if fqdnConnectionID == "" {
		err = errors.New("missing required fqdn_connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("fqdn_connections/%s/fqdn_authentication", fqdnConnectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type FqdnAuthentication struct {
	// Identifies the resource.
	ID string `json:"id"`
	// The ID of the FQDN connection this authentication strategy belongs to.
	ConnectionID string `json:"connection_id"`
	// The failover webhook URL.
	FailoverURL string `json:"failover_url" format:"uri"`
	// The outbound authentication type.
	//
	// Any of "ip-authentication", "credential-authentication".
	FqdnOutboundAuthentication FqdnAuthenticationFqdnOutboundAuthentication `json:"fqdn_outbound_authentication"`
	// The IP authentication method.
	//
	// Any of "token", "p-charge-info".
	IPAuthenticationMethod FqdnAuthenticationIPAuthenticationMethod `json:"ip_authentication_method"`
	// Whether the connection is a Microsoft Teams SBC.
	MicrosoftTeamsSbc bool `json:"microsoft_teams_sbc"`
	// The password for authentication.
	Password string `json:"password"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The TXT record name for Microsoft Teams SBC DNS verification.
	TxtName string `json:"txt_name"`
	// The TTL for the TXT record.
	TxtTtl int64 `json:"txt_ttl"`
	// The TXT record value for Microsoft Teams SBC DNS verification.
	TxtValue string `json:"txt_value"`
	// The username for authentication.
	UserName string `json:"user_name"`
	// The webhook URL for authentication events.
	WebhookURL string `json:"webhook_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		ConnectionID               respjson.Field
		FailoverURL                respjson.Field
		FqdnOutboundAuthentication respjson.Field
		IPAuthenticationMethod     respjson.Field
		MicrosoftTeamsSbc          respjson.Field
		Password                   respjson.Field
		RecordType                 respjson.Field
		TxtName                    respjson.Field
		TxtTtl                     respjson.Field
		TxtValue                   respjson.Field
		UserName                   respjson.Field
		WebhookURL                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnAuthentication) RawJSON() string { return r.JSON.raw }
func (r *FqdnAuthentication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outbound authentication type.
type FqdnAuthenticationFqdnOutboundAuthentication string

const (
	FqdnAuthenticationFqdnOutboundAuthenticationIPAuthentication         FqdnAuthenticationFqdnOutboundAuthentication = "ip-authentication"
	FqdnAuthenticationFqdnOutboundAuthenticationCredentialAuthentication FqdnAuthenticationFqdnOutboundAuthentication = "credential-authentication"
)

// The IP authentication method.
type FqdnAuthenticationIPAuthenticationMethod string

const (
	FqdnAuthenticationIPAuthenticationMethodToken       FqdnAuthenticationIPAuthenticationMethod = "token"
	FqdnAuthenticationIPAuthenticationMethodPChargeInfo FqdnAuthenticationIPAuthenticationMethod = "p-charge-info"
)

type FqdnConnectionFqdnAuthenticationListResponse struct {
	Data FqdnAuthentication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionFqdnAuthenticationListResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionFqdnAuthenticationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionFqdnAuthenticationPatchAllResponse struct {
	Data FqdnAuthentication `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FqdnConnectionFqdnAuthenticationPatchAllResponse) RawJSON() string { return r.JSON.raw }
func (r *FqdnConnectionFqdnAuthenticationPatchAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FqdnConnectionFqdnAuthenticationPatchAllParams struct {
	// The failover webhook URL.
	FailoverURL param.Opt[string] `json:"failover_url,omitzero" format:"uri"`
	// The password for authentication.
	Password param.Opt[string] `json:"password,omitzero"`
	// The TXT record name for Microsoft Teams SBC DNS verification.
	TxtName param.Opt[string] `json:"txt_name,omitzero"`
	// The TTL for the TXT record.
	TxtTtl param.Opt[int64] `json:"txt_ttl,omitzero"`
	// The TXT record value for Microsoft Teams SBC DNS verification.
	TxtValue param.Opt[string] `json:"txt_value,omitzero"`
	// The username for authentication.
	UserName param.Opt[string] `json:"user_name,omitzero"`
	// The webhook URL for authentication events.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"uri"`
	// The outbound authentication type.
	//
	// Any of "ip-authentication", "credential-authentication".
	FqdnOutboundAuthentication FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthentication `json:"fqdn_outbound_authentication,omitzero"`
	// The IP authentication method.
	//
	// Any of "token", "p-charge-info".
	IPAuthenticationMethod FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethod `json:"ip_authentication_method,omitzero"`
	paramObj
}

func (r FqdnConnectionFqdnAuthenticationPatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow FqdnConnectionFqdnAuthenticationPatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FqdnConnectionFqdnAuthenticationPatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outbound authentication type.
type FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthentication string

const (
	FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthenticationIPAuthentication         FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthentication = "ip-authentication"
	FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthenticationCredentialAuthentication FqdnConnectionFqdnAuthenticationPatchAllParamsFqdnOutboundAuthentication = "credential-authentication"
)

// The IP authentication method.
type FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethod string

const (
	FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethodToken       FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethod = "token"
	FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethodPChargeInfo FqdnConnectionFqdnAuthenticationPatchAllParamsIPAuthenticationMethod = "p-charge-info"
)
