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

// Manage Whatsapp phone numbers
//
// WhatsappPhoneNumberConversationalComponentService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberConversationalComponentService] method instead.
type WhatsappPhoneNumberConversationalComponentService struct {
	Options []option.RequestOption
}

// NewWhatsappPhoneNumberConversationalComponentService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewWhatsappPhoneNumberConversationalComponentService(opts ...option.RequestOption) (r WhatsappPhoneNumberConversationalComponentService) {
	r = WhatsappPhoneNumberConversationalComponentService{}
	r.Options = opts
	return
}

// Returns the conversational components configured for the specified WhatsApp
// phone number.
func (r *WhatsappPhoneNumberConversationalComponentService) List(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *WhatsappPhoneNumberConversationalComponentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/conversational_components", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the conversational components configured for the specified WhatsApp
// phone number.
func (r *WhatsappPhoneNumberConversationalComponentService) PatchAll(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberConversationalComponentPatchAllParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberConversationalComponentPatchAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/conversational_components", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappPhoneNumberConversationalComponentListResponse struct {
	Data WhatsappPhoneNumberConversationalComponentListResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberConversationalComponentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentListResponseData struct {
	// List of commands
	Commands []WhatsappPhoneNumberConversationalComponentListResponseDataCommand `json:"commands"`
	// List of ice breakers
	IceBreakers []string `json:"ice_breakers"`
	// Phone number in E164 format
	PhoneNumber string `json:"phone_number"`
	RecordType  string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commands    respjson.Field
		IceBreakers respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentListResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberConversationalComponentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentListResponseDataCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentListResponseDataCommand) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberConversationalComponentListResponseDataCommand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentPatchAllResponse struct {
	Data WhatsappPhoneNumberConversationalComponentPatchAllResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentPatchAllResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberConversationalComponentPatchAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentPatchAllResponseData struct {
	// List of commands
	Commands []WhatsappPhoneNumberConversationalComponentPatchAllResponseDataCommand `json:"commands"`
	// List of ice breakers
	IceBreakers []string `json:"ice_breakers"`
	// Phone number in E164 format
	PhoneNumber string `json:"phone_number"`
	RecordType  string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commands    respjson.Field
		IceBreakers respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentPatchAllResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberConversationalComponentPatchAllResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentPatchAllResponseDataCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberConversationalComponentPatchAllResponseDataCommand) RawJSON() string {
	return r.JSON.raw
}
func (r *WhatsappPhoneNumberConversationalComponentPatchAllResponseDataCommand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentPatchAllParams struct {
	// List of commands
	Commands []WhatsappPhoneNumberConversationalComponentPatchAllParamsCommand `json:"commands,omitzero"`
	// List of ice breakers
	IceBreakers []string `json:"ice_breakers,omitzero"`
	paramObj
}

func (r WhatsappPhoneNumberConversationalComponentPatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberConversationalComponentPatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberConversationalComponentPatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberConversationalComponentPatchAllParamsCommand struct {
	Command     param.Opt[string] `json:"command,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r WhatsappPhoneNumberConversationalComponentPatchAllParamsCommand) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberConversationalComponentPatchAllParamsCommand
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberConversationalComponentPatchAllParamsCommand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
