// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// Voicemail API
//
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

// Create voicemail settings for a phone number. You can also configure a custom
// greeting by setting the `greeting` object: use `mode` `custom_greeting` together
// with a `media_name` that points to an audio file uploaded through the Media
// Storage API, or `mode` `default` to use the standard system greeting.
func (r *PhoneNumberVoicemailService) New(ctx context.Context, phoneNumberID string, body PhoneNumberVoicemailNewParams, opts ...option.RequestOption) (res *PhoneNumberVoicemailNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the voicemail settings for a phone number
func (r *PhoneNumberVoicemailService) Get(ctx context.Context, phoneNumberID string, opts ...option.RequestOption) (res *PhoneNumberVoicemailGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update voicemail settings for a phone number. You can also configure a custom
// greeting by setting the `greeting` object: use `mode` `custom_greeting` together
// with a `media_name` that points to an audio file uploaded through the Media
// Storage API, or `mode` `default` to use the standard system greeting.
func (r *PhoneNumberVoicemailService) Update(ctx context.Context, phoneNumberID string, body PhoneNumberVoicemailUpdateParams, opts ...option.RequestOption) (res *PhoneNumberVoicemailUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("phone_numbers/%s/voicemail", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type VoicemailPrefResponse struct {
	// Whether voicemail is enabled.
	Enabled bool `json:"enabled"`
	// Controls the greeting a caller hears before leaving a voicemail. Set `mode` to
	// `default` to play the standard system greeting, or to `custom_greeting` to play
	// your own audio. When `mode` is `custom_greeting`, `media_name` is required and
	// must reference an audio file already uploaded to your account through the Media
	// Storage API.
	Greeting VoicemailPrefResponseGreeting `json:"greeting"`
	// The pin used for the voicemail.
	Pin string `json:"pin"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Greeting    respjson.Field
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

// Controls the greeting a caller hears before leaving a voicemail. Set `mode` to
// `default` to play the standard system greeting, or to `custom_greeting` to play
// your own audio. When `mode` is `custom_greeting`, `media_name` is required and
// must reference an audio file already uploaded to your account through the Media
// Storage API.
type VoicemailPrefResponseGreeting struct {
	// The name of the media file to play as the greeting. Required when `mode` is
	// `custom_greeting`; ignored when `mode` is `default`. The value must match the
	// `media_name` of a file you previously uploaded with the Media Storage API
	// (`POST /v2/media`).
	MediaName string `json:"media_name" api:"nullable"`
	// The greeting mode. `default` plays the standard system greeting.
	// `custom_greeting` plays the audio referenced by `media_name`.
	//
	// Any of "default", "custom_greeting".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaName   respjson.Field
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicemailPrefResponseGreeting) RawJSON() string { return r.JSON.raw }
func (r *VoicemailPrefResponseGreeting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicemailRequestParam struct {
	// Whether voicemail is enabled.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The pin used for voicemail
	Pin param.Opt[string] `json:"pin,omitzero"`
	// Controls the greeting a caller hears before leaving a voicemail. Set `mode` to
	// `default` to play the standard system greeting, or to `custom_greeting` to play
	// your own audio. When `mode` is `custom_greeting`, `media_name` is required and
	// must reference an audio file already uploaded to your account through the Media
	// Storage API.
	Greeting VoicemailRequestGreetingParam `json:"greeting,omitzero"`
	paramObj
}

func (r VoicemailRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow VoicemailRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicemailRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls the greeting a caller hears before leaving a voicemail. Set `mode` to
// `default` to play the standard system greeting, or to `custom_greeting` to play
// your own audio. When `mode` is `custom_greeting`, `media_name` is required and
// must reference an audio file already uploaded to your account through the Media
// Storage API.
type VoicemailRequestGreetingParam struct {
	// The name of the media file to play as the greeting. Required when `mode` is
	// `custom_greeting`; ignored when `mode` is `default`. The value must match the
	// `media_name` of a file you previously uploaded with the Media Storage API
	// (`POST /v2/media`).
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// The greeting mode. `default` plays the standard system greeting.
	// `custom_greeting` plays the audio referenced by `media_name`.
	//
	// Any of "default", "custom_greeting".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r VoicemailRequestGreetingParam) MarshalJSON() (data []byte, err error) {
	type shadow VoicemailRequestGreetingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicemailRequestGreetingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoicemailRequestGreetingParam](
		"mode", "default", "custom_greeting",
	)
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
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoicemailUpdateParams struct {
	VoicemailRequest VoicemailRequestParam
	paramObj
}

func (r PhoneNumberVoicemailUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VoicemailRequest)
}
func (r *PhoneNumberVoicemailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
