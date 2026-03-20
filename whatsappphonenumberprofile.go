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
// WhatsappPhoneNumberProfileService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberProfileService] method instead.
type WhatsappPhoneNumberProfileService struct {
	Options []option.RequestOption
	// Manage Whatsapp phone numbers
	Photo WhatsappPhoneNumberProfilePhotoService
}

// NewWhatsappPhoneNumberProfileService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberProfileService(opts ...option.RequestOption) (r WhatsappPhoneNumberProfileService) {
	r = WhatsappPhoneNumberProfileService{}
	r.Options = opts
	r.Photo = NewWhatsappPhoneNumberProfilePhotoService(opts...)
	return
}

// Get phone number business profile
func (r *WhatsappPhoneNumberProfileService) Get(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *WhatsappPhoneNumberProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/profile", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update phone number business profile
func (r *WhatsappPhoneNumberProfileService) Update(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberProfileUpdateParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/profile", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type WhatsappProfileData struct {
	ID          string    `json:"id"`
	About       string    `json:"about"`
	Address     string    `json:"address"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Description string    `json:"description"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	// Whatsapp phone number ID
	PhoneNumberID   string    `json:"phone_number_id"`
	ProfilePhotoURL string    `json:"profile_photo_url"`
	RecordType      string    `json:"record_type"`
	UpdatedAt       time.Time `json:"updated_at" format:"date-time"`
	Website         string    `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		About           respjson.Field
		Address         respjson.Field
		Category        respjson.Field
		CreatedAt       respjson.Field
		Description     respjson.Field
		DisplayName     respjson.Field
		Email           respjson.Field
		PhoneNumberID   respjson.Field
		ProfilePhotoURL respjson.Field
		RecordType      respjson.Field
		UpdatedAt       respjson.Field
		Website         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappProfileData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappProfileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberProfileGetResponse struct {
	Data WhatsappProfileData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberProfileUpdateResponse struct {
	Data WhatsappProfileData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberProfileUpdateParams struct {
	About       param.Opt[string] `json:"about,omitzero"`
	Address     param.Opt[string] `json:"address,omitzero"`
	Category    param.Opt[string] `json:"category,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	Email       param.Opt[string] `json:"email,omitzero"`
	Website     param.Opt[string] `json:"website,omitzero"`
	paramObj
}

func (r WhatsappPhoneNumberProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappPhoneNumberProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappPhoneNumberProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
