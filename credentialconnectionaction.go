// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// CredentialConnectionActionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialConnectionActionService] method instead.
type CredentialConnectionActionService struct {
	Options []option.RequestOption
}

// NewCredentialConnectionActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCredentialConnectionActionService(opts ...option.RequestOption) (r CredentialConnectionActionService) {
	r = CredentialConnectionActionService{}
	r.Options = opts
	return
}

// Checks the registration_status for a credential connection,
// (`registration_status`) as well as the timestamp for the last SIP registration
// event (`registration_status_updated_at`)
func (r *CredentialConnectionActionService) CheckRegistrationStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *CredentialConnectionActionCheckRegistrationStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("credential_connections/%s/actions/check_registration_status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type CredentialConnectionActionCheckRegistrationStatusResponse struct {
	Data CredentialConnectionActionCheckRegistrationStatusResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionActionCheckRegistrationStatusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialConnectionActionCheckRegistrationStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CredentialConnectionActionCheckRegistrationStatusResponseData struct {
	// The ip used during the SIP connection
	IPAddress string `json:"ip_address"`
	// ISO 8601 formatted date indicating when the resource was last updated.
	LastRegistration string `json:"last_registration"`
	// The port of the SIP connection
	Port int64 `json:"port"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The user name of the SIP connection
	SipUsername string `json:"sip_username"`
	// The current registration status of your SIP connection
	//
	// Any of "Not Applicable", "Not Registered", "Failed", "Expired", "Registered",
	// "Unregistered".
	Status string `json:"status"`
	// The protocol of the SIP connection
	Transport string `json:"transport"`
	// The user agent of the SIP connection
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress        respjson.Field
		LastRegistration respjson.Field
		Port             respjson.Field
		RecordType       respjson.Field
		SipUsername      respjson.Field
		Status           respjson.Field
		Transport        respjson.Field
		UserAgent        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CredentialConnectionActionCheckRegistrationStatusResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *CredentialConnectionActionCheckRegistrationStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
