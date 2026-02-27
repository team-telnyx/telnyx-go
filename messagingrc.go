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

// Send RCS messages
//
// MessagingRcService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingRcService] method instead.
type MessagingRcService struct {
	Options []option.RequestOption
	// Send RCS messages
	Agents MessagingRcAgentService
}

// NewMessagingRcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingRcService(opts ...option.RequestOption) (r MessagingRcService) {
	r = MessagingRcService{}
	r.Options = opts
	r.Agents = NewMessagingRcAgentService(opts...)
	return
}

// Adds a test phone number to an RCS agent for testing purposes.
func (r *MessagingRcService) InviteTestNumber(ctx context.Context, phoneNumber string, body MessagingRcInviteTestNumberParams, opts ...option.RequestOption) (res *MessagingRcInviteTestNumberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("messaging/rcs/test_number_invite/%s/%s", body.ID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Check RCS capabilities (batch)
func (r *MessagingRcService) ListBulkCapabilities(ctx context.Context, body MessagingRcListBulkCapabilitiesParams, opts ...option.RequestOption) (res *MessagingRcListBulkCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messaging/rcs/bulk_capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check RCS capabilities
func (r *MessagingRcService) GetCapabilities(ctx context.Context, phoneNumber string, query MessagingRcGetCapabilitiesParams, opts ...option.RequestOption) (res *MessagingRcGetCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("messaging/rcs/capabilities/%s/%s", query.AgentID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RcsCapabilities struct {
	// RCS agent ID
	AgentID string `json:"agent_id"`
	// RCS agent name
	AgentName string `json:"agent_name"`
	// List of RCS capabilities
	Features []string `json:"features"`
	// Phone number
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource
	//
	// Any of "rcs.capabilities".
	RecordType RcsCapabilitiesRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		AgentName   respjson.Field
		Features    respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsCapabilities) RawJSON() string { return r.JSON.raw }
func (r *RcsCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource
type RcsCapabilitiesRecordType string

const (
	RcsCapabilitiesRecordTypeRcsCapabilities RcsCapabilitiesRecordType = "rcs.capabilities"
)

type MessagingRcInviteTestNumberResponse struct {
	Data MessagingRcInviteTestNumberResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingRcInviteTestNumberResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingRcInviteTestNumberResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcInviteTestNumberResponseData struct {
	// RCS agent ID
	AgentID string `json:"agent_id"`
	// Phone number that was invited for testing
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource
	//
	// Any of "rcs.test_number_invite".
	RecordType string `json:"record_type"`
	// Status of the test number invitation
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingRcInviteTestNumberResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessagingRcInviteTestNumberResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcListBulkCapabilitiesResponse struct {
	Data []RcsCapabilities `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingRcListBulkCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingRcListBulkCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcGetCapabilitiesResponse struct {
	Data RcsCapabilities `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingRcGetCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingRcGetCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcInviteTestNumberParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type MessagingRcListBulkCapabilitiesParams struct {
	// RCS Agent ID
	AgentID string `json:"agent_id" api:"required"`
	// List of phone numbers to check
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r MessagingRcListBulkCapabilitiesParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingRcListBulkCapabilitiesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingRcListBulkCapabilitiesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingRcGetCapabilitiesParams struct {
	AgentID string `path:"agent_id" api:"required" json:"-"`
	paramObj
}
