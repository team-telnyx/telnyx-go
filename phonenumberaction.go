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

// Configure your phone numbers
//
// PhoneNumberActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberActionService] method instead.
type PhoneNumberActionService struct {
	Options []option.RequestOption
}

// NewPhoneNumberActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberActionService(opts ...option.RequestOption) (r PhoneNumberActionService) {
	r = PhoneNumberActionService{}
	r.Options = opts
	return
}

// Change the bundle status for a phone number (set to being in a bundle or remove
// from a bundle)
func (r *PhoneNumberActionService) ChangeBundleStatus(ctx context.Context, id string, body PhoneNumberActionChangeBundleStatusParams, opts ...option.RequestOption) (res *PhoneNumberActionChangeBundleStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/actions/bundle_status_change", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enable emergency for a phone number
func (r *PhoneNumberActionService) EnableEmergency(ctx context.Context, id string, body PhoneNumberActionEnableEmergencyParams, opts ...option.RequestOption) (res *PhoneNumberActionEnableEmergencyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/actions/enable_emergency", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Verifies ownership of the provided phone numbers and returns a mapping of
// numbers to their IDs, plus a list of numbers not found in the account.
func (r *PhoneNumberActionService) VerifyOwnership(ctx context.Context, body PhoneNumberActionVerifyOwnershipParams, opts ...option.RequestOption) (res *PhoneNumberActionVerifyOwnershipResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/actions/verify_ownership"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PhoneNumberWithVoiceSettings struct {
	// Identifies the type of resource.
	ID string `json:"id"`
	// The call forwarding settings for a phone number.
	CallForwarding CallForwarding `json:"call_forwarding"`
	// The call recording settings for a phone number.
	CallRecording CallRecording `json:"call_recording"`
	// The CNAM listing settings for a phone number.
	CnamListing CnamListing `json:"cnam_listing"`
	// Identifies the connection associated with this phone number.
	ConnectionID string `json:"connection_id"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// The emergency services settings for a phone number.
	Emergency PhoneNumberWithVoiceSettingsEmergency `json:"emergency"`
	// The inbound_call_screening setting is a phone number configuration option
	// variable that allows users to configure their settings to block or flag
	// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
	// feature has an additional per-number monthly cost associated with it.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening PhoneNumberWithVoiceSettingsInboundCallScreening `json:"inbound_call_screening"`
	// The media features settings for a phone number.
	MediaFeatures MediaFeatures `json:"media_features"`
	// The phone number in +E164 format.
	PhoneNumber string `json:"phone_number"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Controls whether a tech prefix is enabled for this phone number.
	TechPrefixEnabled bool `json:"tech_prefix_enabled"`
	// This field allows you to rewrite the destination number of an inbound call
	// before the call is routed to you. The value of this field may be any
	// alphanumeric value, and the value will replace the number originally dialed.
	TranslatedNumber string `json:"translated_number"`
	// Controls whether a number is billed per minute or uses your concurrent channels.
	//
	// Any of "pay-per-minute", "channel".
	UsagePaymentMethod PhoneNumberWithVoiceSettingsUsagePaymentMethod `json:"usage_payment_method"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CallForwarding       respjson.Field
		CallRecording        respjson.Field
		CnamListing          respjson.Field
		ConnectionID         respjson.Field
		CustomerReference    respjson.Field
		Emergency            respjson.Field
		InboundCallScreening respjson.Field
		MediaFeatures        respjson.Field
		PhoneNumber          respjson.Field
		RecordType           respjson.Field
		TechPrefixEnabled    respjson.Field
		TranslatedNumber     respjson.Field
		UsagePaymentMethod   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberWithVoiceSettings) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberWithVoiceSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The emergency services settings for a phone number.
type PhoneNumberWithVoiceSettingsEmergency struct {
	// Identifies the address to be used with emergency services.
	EmergencyAddressID string `json:"emergency_address_id"`
	// Allows you to enable or disable emergency services on the phone number. In order
	// to enable emergency services, you must also set an emergency_address_id.
	EmergencyEnabled bool `json:"emergency_enabled"`
	// Represents the state of the number regarding emergency activation.
	//
	// Any of "disabled", "active", "provisioning", "deprovisioning",
	// "provisioning-failed".
	EmergencyStatus string `json:"emergency_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmergencyAddressID respjson.Field
		EmergencyEnabled   respjson.Field
		EmergencyStatus    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberWithVoiceSettingsEmergency) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberWithVoiceSettingsEmergency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The inbound_call_screening setting is a phone number configuration option
// variable that allows users to configure their settings to block or flag
// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
// feature has an additional per-number monthly cost associated with it.
type PhoneNumberWithVoiceSettingsInboundCallScreening string

const (
	PhoneNumberWithVoiceSettingsInboundCallScreeningDisabled    PhoneNumberWithVoiceSettingsInboundCallScreening = "disabled"
	PhoneNumberWithVoiceSettingsInboundCallScreeningRejectCalls PhoneNumberWithVoiceSettingsInboundCallScreening = "reject_calls"
	PhoneNumberWithVoiceSettingsInboundCallScreeningFlagCalls   PhoneNumberWithVoiceSettingsInboundCallScreening = "flag_calls"
)

// Controls whether a number is billed per minute or uses your concurrent channels.
type PhoneNumberWithVoiceSettingsUsagePaymentMethod string

const (
	PhoneNumberWithVoiceSettingsUsagePaymentMethodPayPerMinute PhoneNumberWithVoiceSettingsUsagePaymentMethod = "pay-per-minute"
	PhoneNumberWithVoiceSettingsUsagePaymentMethodChannel      PhoneNumberWithVoiceSettingsUsagePaymentMethod = "channel"
)

type PhoneNumberActionChangeBundleStatusResponse struct {
	Data PhoneNumberWithVoiceSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberActionChangeBundleStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberActionChangeBundleStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionEnableEmergencyResponse struct {
	Data PhoneNumberWithVoiceSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberActionEnableEmergencyResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberActionEnableEmergencyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionVerifyOwnershipResponse struct {
	Data PhoneNumberActionVerifyOwnershipResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberActionVerifyOwnershipResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberActionVerifyOwnershipResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionVerifyOwnershipResponseData struct {
	// The list of phone numbers which you own and are in an editable state
	Found []PhoneNumberActionVerifyOwnershipResponseDataFound `json:"found"`
	// Phone numbers that are not found in the account
	NotFound []string `json:"not_found"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Found       respjson.Field
		NotFound    respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberActionVerifyOwnershipResponseData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberActionVerifyOwnershipResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionVerifyOwnershipResponseDataFound struct {
	// Identifies the resource.
	ID string `json:"id"`
	// The phone number in E.164 format
	NumberValE164 string `json:"number_val_e164"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		NumberValE164 respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberActionVerifyOwnershipResponseDataFound) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberActionVerifyOwnershipResponseDataFound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionChangeBundleStatusParams struct {
	// The new bundle_id setting for the number. If you are assigning the number to a
	// bundle, this is the unique ID of the bundle you wish to use. If you are removing
	// the number from a bundle, this must be null. You cannot assign a number from one
	// bundle to another directly. You must first remove it from a bundle, and then
	// assign it to a new bundle.
	BundleID string `json:"bundle_id" api:"required"`
	paramObj
}

func (r PhoneNumberActionChangeBundleStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberActionChangeBundleStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberActionChangeBundleStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionEnableEmergencyParams struct {
	// Identifies the address to be used with emergency services.
	EmergencyAddressID string `json:"emergency_address_id" api:"required"`
	// Indicates whether to enable emergency services on this number.
	EmergencyEnabled bool `json:"emergency_enabled" api:"required"`
	paramObj
}

func (r PhoneNumberActionEnableEmergencyParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberActionEnableEmergencyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberActionEnableEmergencyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberActionVerifyOwnershipParams struct {
	// Array of phone numbers to verify ownership for
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r PhoneNumberActionVerifyOwnershipParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberActionVerifyOwnershipParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberActionVerifyOwnershipParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
