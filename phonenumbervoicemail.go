// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// PhoneNumberVoicemailService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberVoicemailService] method instead.
type PhoneNumberVoicemailService struct {
	Options []option.RequestOption
}

// NewPhoneNumberVoicemailService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberVoicemailService(opts ...option.RequestOption) (r PhoneNumberVoicemailService) {
	r = PhoneNumberVoicemailService{}
	r.Options = opts
	return
}

// Create voicemail settings for a phone number
func (r *PhoneNumberVoicemailService) New(ctx context.Context, phoneNumberID string, body PhoneNumberVoicemailNewParams, opts ...option.RequestOption) (res *PhoneNumberVoicemailNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the voicemail settings for a phone number
func (r *PhoneNumberVoicemailService) Get(ctx context.Context, phoneNumberID string, opts ...option.RequestOption) (res *PhoneNumberVoicemailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update voicemail settings for a phone number
func (r *PhoneNumberVoicemailService) Update(ctx context.Context, phoneNumberID string, body PhoneNumberVoicemailUpdateParams, opts ...option.RequestOption) (res *PhoneNumberVoicemailUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type VoicemailPrefResponse struct {
	// Whether voicemail is enabled.
	Enabled bool `json:"enabled"`
	// The pin used for the voicemail.
	Pin string `json:"pin"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Pin         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicemailPrefResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicemailPrefResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicemailRequestParam struct {
	// Whether voicemail is enabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The pin used for voicemail
	Pin param.Opt[string] `json:"pin,omitzero"`
	paramObj
}

func (r VoicemailRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow VoicemailRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicemailRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoicemailNewResponse struct {
	Data VoicemailPrefResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberVoicemailNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberVoicemailNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoicemailGetResponse struct {
	Data VoicemailPrefResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberVoicemailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberVoicemailGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoicemailUpdateResponse struct {
	Data VoicemailPrefResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberVoicemailUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberVoicemailUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoicemailNewParams struct {
	VoicemailRequest VoicemailRequestParam
	paramObj
}

func (r PhoneNumberVoicemailNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VoicemailRequest)
}
func (r *PhoneNumberVoicemailNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.VoicemailRequest)
}

type PhoneNumberVoicemailUpdateParams struct {
	VoicemailRequest VoicemailRequestParam
	paramObj
}

func (r PhoneNumberVoicemailUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VoicemailRequest)
}
func (r *PhoneNumberVoicemailUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.VoicemailRequest)
}
