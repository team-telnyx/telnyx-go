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
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// SiprecConnectorService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiprecConnectorService] method instead.
type SiprecConnectorService struct {
	Options []option.RequestOption
}

// NewSiprecConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSiprecConnectorService(opts ...option.RequestOption) (r SiprecConnectorService) {
	r = SiprecConnectorService{}
	r.Options = opts
	return
}

// Creates a new SIPREC connector configuration.
func (r *SiprecConnectorService) New(ctx context.Context, body SiprecConnectorNewParams, opts ...option.RequestOption) (res *SiprecConnectorNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "siprec_connectors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a stored SIPREC connector.
func (r *SiprecConnectorService) Get(ctx context.Context, connectorName string, opts ...option.RequestOption) (res *SiprecConnectorGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectorName == "" {
		err = errors.New("missing required connector_name parameter")
		return
	}
	path := fmt.Sprintf("siprec_connectors/%s", connectorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a stored SIPREC connector configuration.
func (r *SiprecConnectorService) Update(ctx context.Context, connectorName string, body SiprecConnectorUpdateParams, opts ...option.RequestOption) (res *SiprecConnectorUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectorName == "" {
		err = errors.New("missing required connector_name parameter")
		return
	}
	path := fmt.Sprintf("siprec_connectors/%s", connectorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a stored SIPREC connector.
func (r *SiprecConnectorService) Delete(ctx context.Context, connectorName string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if connectorName == "" {
		err = errors.New("missing required connector_name parameter")
		return
	}
	path := fmt.Sprintf("siprec_connectors/%s", connectorName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type SiprecConnectorNewResponse struct {
	Data SiprecConnectorNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorNewResponseData struct {
	// Subdomain to route calls when using Telnyx SRS (optional).
	AppSubdomain string `json:"app_subdomain"`
	// ISO 8601 formatted date/time of creation.
	CreatedAt string `json:"created_at"`
	// Hostname/IPv4 address of the SIPREC SRS.
	Host string `json:"host"`
	// Name for the SIPREC connector resource.
	Name string `json:"name"`
	// Port for the SIPREC SRS.
	Port       int64  `json:"port"`
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date/time of last update.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppSubdomain respjson.Field
		CreatedAt    respjson.Field
		Host         respjson.Field
		Name         respjson.Field
		Port         respjson.Field
		RecordType   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorGetResponse struct {
	Data SiprecConnectorGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorGetResponseData struct {
	// Subdomain to route calls when using Telnyx SRS (optional).
	AppSubdomain string `json:"app_subdomain"`
	// ISO 8601 formatted date/time of creation.
	CreatedAt string `json:"created_at"`
	// Hostname/IPv4 address of the SIPREC SRS.
	Host string `json:"host"`
	// Name for the SIPREC connector resource.
	Name string `json:"name"`
	// Port for the SIPREC SRS.
	Port       int64  `json:"port"`
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date/time of last update.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppSubdomain respjson.Field
		CreatedAt    respjson.Field
		Host         respjson.Field
		Name         respjson.Field
		Port         respjson.Field
		RecordType   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorUpdateResponse struct {
	Data SiprecConnectorUpdateResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorUpdateResponseData struct {
	// Subdomain to route calls when using Telnyx SRS (optional).
	AppSubdomain string `json:"app_subdomain"`
	// ISO 8601 formatted date/time of creation.
	CreatedAt string `json:"created_at"`
	// Hostname/IPv4 address of the SIPREC SRS.
	Host string `json:"host"`
	// Name for the SIPREC connector resource.
	Name string `json:"name"`
	// Port for the SIPREC SRS.
	Port       int64  `json:"port"`
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date/time of last update.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppSubdomain respjson.Field
		CreatedAt    respjson.Field
		Host         respjson.Field
		Name         respjson.Field
		Port         respjson.Field
		RecordType   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SiprecConnectorUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *SiprecConnectorUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorNewParams struct {
	// Hostname/IPv4 address of the SIPREC SRS.
	Host string `json:"host,required"`
	// Name for the SIPREC connector resource.
	Name string `json:"name,required"`
	// Port for the SIPREC SRS.
	Port int64 `json:"port,required"`
	// Subdomain to route the call when using Telnyx SRS (optional for non-Telnyx SRS).
	AppSubdomain param.Opt[string] `json:"app_subdomain,omitzero"`
	paramObj
}

func (r SiprecConnectorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SiprecConnectorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiprecConnectorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SiprecConnectorUpdateParams struct {
	// Hostname/IPv4 address of the SIPREC SRS.
	Host string `json:"host,required"`
	// Name for the SIPREC connector resource.
	Name string `json:"name,required"`
	// Port for the SIPREC SRS.
	Port int64 `json:"port,required"`
	// Subdomain to route the call when using Telnyx SRS (optional for non-Telnyx SRS).
	AppSubdomain param.Opt[string] `json:"app_subdomain,omitzero"`
	paramObj
}

func (r SiprecConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SiprecConnectorUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SiprecConnectorUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
