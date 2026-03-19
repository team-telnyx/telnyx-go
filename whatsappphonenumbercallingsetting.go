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
)

// Manage Whatsapp phone numbers
//
// WhatsappPhoneNumberCallingSettingService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberCallingSettingService] method instead.
type WhatsappPhoneNumberCallingSettingService struct {
	Options []option.RequestOption
}

// NewWhatsappPhoneNumberCallingSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberCallingSettingService(opts ...option.RequestOption) (r WhatsappPhoneNumberCallingSettingService) {
	r = WhatsappPhoneNumberCallingSettingService{}
	r.Options = opts
	return
}

// Get calling settings for a phone number
func (r *WhatsappPhoneNumberCallingSettingService) Get(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *WhatsappPhoneNumberCallingSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/calling_settings", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Enable or disable Whatsapp calling for a phone number
func (r *WhatsappPhoneNumberCallingSettingService) Update(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberCallingSettingUpdateParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberCallingSettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/calling_settings", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappPhoneNumberCallingSettingGetResponse struct {
	Data WhatsappPhoneNumberCallingSettingGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingSettingGetResponseData struct {
	// True if calling is enabled on the phone
	Enabled bool `json:"enabled"`
	// Phone number in E164 format
	PhoneNumber string    `json:"phone_number"`
	RecordType  string    `json:"record_type"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingSettingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingSettingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingSettingUpdateResponse struct {
	Data WhatsappPhoneNumberCallingSettingUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingSettingUpdateResponseData struct {
	// True if calling is enabled on the phone
	Enabled bool `json:"enabled"`
	// Phone number in E164 format
	PhoneNumber string    `json:"phone_number"`
	RecordType  string    `json:"record_type"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		PhoneNumber respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberCallingSettingUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberCallingSettingUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberCallingSettingUpdateParams struct {
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r WhatsappPhoneNumberCallingSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberCallingSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberCallingSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
