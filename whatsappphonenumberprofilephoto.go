// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage Whatsapp phone numbers
//
// WhatsappPhoneNumberProfilePhotoService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappPhoneNumberProfilePhotoService] method instead.
type WhatsappPhoneNumberProfilePhotoService struct {
	Options []option.RequestOption
}

// NewWhatsappPhoneNumberProfilePhotoService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWhatsappPhoneNumberProfilePhotoService(opts ...option.RequestOption) (r WhatsappPhoneNumberProfilePhotoService) {
	r = WhatsappPhoneNumberProfilePhotoService{}
	r.Options = opts
	return
}

// Delete Whatsapp profile photo
func (r *WhatsappPhoneNumberProfilePhotoService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/profile/photo", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Upload Whatsapp profile photo
func (r *WhatsappPhoneNumberProfilePhotoService) Upload(ctx context.Context, phoneNumber string, body WhatsappPhoneNumberProfilePhotoUploadParams, opts ...option.RequestOption) (res *WhatsappPhoneNumberProfilePhotoUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/whatsapp/phone_numbers/%s/profile/photo", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WhatsappPhoneNumberProfilePhotoUploadResponse struct {
	Data WhatsappPhoneNumberProfilePhotoUploadResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappPhoneNumberProfilePhotoUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberProfilePhotoUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberProfilePhotoUploadResponseData struct {
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
func (r WhatsappPhoneNumberProfilePhotoUploadResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappPhoneNumberProfilePhotoUploadResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappPhoneNumberProfilePhotoUploadParams struct {
	// Image file (JPEG recommended, max 10 MB)
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r WhatsappPhoneNumberProfilePhotoUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
