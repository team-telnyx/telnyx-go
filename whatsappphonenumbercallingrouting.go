// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage Whatsapp phone numbers
//
// WhatsappPhoneNumberCallingRoutingService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberCallingRoutingService] method instead.
type WhatsappPhoneNumberCallingRoutingService struct {
	Options []option.RequestOption
}

// NewWhatsappPhoneNumberCallingRoutingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberCallingRoutingService(opts ...option.RequestOption) (r WhatsappPhoneNumberCallingRoutingService) {
	r = WhatsappPhoneNumberCallingRoutingService{}
	r.Options = opts
	return
}

// Retrieve the routing connection currently stored for a BYON (Bring Your Own
// Number) phone number: the connection that inbound WhatsApp calls to the number
// are delivered to.
//
// Use it to check the result of
// `PATCH /whatsapp/phone_numbers/{id}/calling_routing`. A read made immediately
// after an update can still return the previous value.
//
// Sub-users need read permission on connections.
func (r *WhatsappPhoneNumberCallingRoutingService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappPhoneNumberCallingRoutingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("whatsapp/phone_numbers/%s/calling_routing", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set or clear the connection that inbound WhatsApp calls to a BYON (Bring Your
// Own Number) phone number are delivered to.
//
// The update is processed asynchronously. A `202` response means the request was
// accepted, not that the routing changed. Check the result with
// `GET /whatsapp/phone_numbers/{id}/calling_routing`, which can return the
// previous value immediately after an update. An update for a number that is not a
// WhatsApp Calling number in the account returns 404.
//
// The connection must belong to the same account and must not be a WhatsApp
// connection. Send `connection_id: null` to clear the routing; omitting
// `connection_id` is rejected. Numbers active on Telnyx are rejected, because they
// route through their own connection assignment.
//
// Sub-users need update permission on connections, and read permission to check
// the result with `GET`.
func (r *WhatsappPhoneNumberCallingRoutingService) PatchAll(ctx context.Context, id string, body WhatsappPhoneNumberCallingRoutingPatchAllParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberCallingRoutingPatchAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("whatsapp/phone_numbers/%s/calling_routing", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappCallingRoutingData struct {
	// ID of the routing connection, or `null` when none is set.
	ConnectionID string `json:"connection_id" api:"required"`
	// Phone number in E.164 format, with a leading `+`.
	PhoneNumber string `json:"phone_number" api:"required"`
	// Identifies the type of the resource.
	//
	// Any of "whatsapp_calling_routing".
	RecordType WhatsappCallingRoutingDataRecordType `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID respjson.Field
		PhoneNumber  respjson.Field
		RecordType   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappCallingRoutingData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappCallingRoutingData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type WhatsappCallingRoutingDataRecordType string

const (
	WhatsappCallingRoutingDataRecordTypeWhatsappCallingRouting WhatsappCallingRoutingDataRecordType = "whatsapp_calling_routing"
)

type WhatsappPhoneNumberCallingRoutingListResponse struct {
	Data WhatsappCallingRoutingData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingRoutingListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingRoutingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingRoutingPatchAllResponse struct {
	Data WhatsappCallingRoutingData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingRoutingPatchAllResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingRoutingPatchAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingRoutingPatchAllParams struct {
	// ID of the connection to deliver inbound WhatsApp calls to: a positive integer up
	// to 9223372036854775807, sent as a decimal string or an integer. Send a string to
	// keep large IDs exact. Non-null values are returned as strings. `null` clears the
	// routing.
	ConnectionID WhatsappPhoneNumberCallingRoutingPatchAllParamsConnectionIDUnion `json:"connection_id,omitzero" api:"required"`
	paramObj
}

func (r WhatsappPhoneNumberCallingRoutingPatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberCallingRoutingPatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberCallingRoutingPatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WhatsappPhoneNumberCallingRoutingPatchAllParamsConnectionIDUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	paramUnion
}

func (u WhatsappPhoneNumberCallingRoutingPatchAllParamsConnectionIDUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *WhatsappPhoneNumberCallingRoutingPatchAllParamsConnectionIDUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WhatsappPhoneNumberCallingRoutingPatchAllParamsConnectionIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}
