// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// VerificationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationService] method instead.
type VerificationService struct {
	Options       []option.RequestOption
	ByPhoneNumber VerificationByPhoneNumberService
	Actions       VerificationActionService
}

// NewVerificationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerificationService(opts ...option.RequestOption) (r VerificationService) {
	r = VerificationService{}
	r.Options = opts
	r.ByPhoneNumber = NewVerificationByPhoneNumberService(opts...)
	r.Actions = NewVerificationActionService(opts...)
	return
}

// Retrieve verification
func (r *VerificationService) Get(ctx context.Context, verificationID string, opts ...option.RequestOption) (res *VerificationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if verificationID == "" {
		err = errors.New("missing required verification_id parameter")
		return
	}
	path := fmt.Sprintf("verifications/%s", verificationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Trigger Call verification
func (r *VerificationService) TriggerCall(ctx context.Context, body VerificationTriggerCallParams, opts ...option.RequestOption) (res *CreateVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "verifications/call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Trigger Flash call verification
func (r *VerificationService) TriggerFlashcall(ctx context.Context, body VerificationTriggerFlashcallParams, opts ...option.RequestOption) (res *CreateVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "verifications/flashcall"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Trigger SMS verification
func (r *VerificationService) TriggerSMS(ctx context.Context, body VerificationTriggerSMSParams, opts ...option.RequestOption) (res *CreateVerificationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "verifications/sms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateVerificationResponse struct {
	Data Verification `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateVerificationResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateVerificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Verification struct {
	ID        string `json:"id" format:"uuid"`
	CreatedAt string `json:"created_at"`
	// Send a self-generated numeric code to the end-user
	CustomCode string `json:"custom_code,nullable"`
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// The possible verification record types.
	//
	// Any of "verification".
	RecordType VerificationRecordType `json:"record_type"`
	// The possible statuses of the verification request.
	//
	// Any of "pending", "accepted", "invalid", "expired", "error".
	Status VerificationStatus `json:"status"`
	// This is the number of seconds before the code of the request is expired. Once
	// this request has expired, the code will no longer verify the user. Note: this
	// will override the `default_verification_timeout_secs` on the Verify profile.
	TimeoutSecs int64 `json:"timeout_secs"`
	// The possible types of verification.
	//
	// Any of "sms", "call", "flashcall".
	Type      VerificationType `json:"type"`
	UpdatedAt string           `json:"updated_at"`
	// The identifier of the associated Verify profile.
	VerifyProfileID string `json:"verify_profile_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CustomCode      respjson.Field
		PhoneNumber     respjson.Field
		RecordType      respjson.Field
		Status          respjson.Field
		TimeoutSecs     respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		VerifyProfileID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Verification) RawJSON() string { return r.JSON.raw }
func (r *Verification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The possible verification record types.
type VerificationRecordType string

const (
	VerificationRecordTypeVerification VerificationRecordType = "verification"
)

// The possible statuses of the verification request.
type VerificationStatus string

const (
	VerificationStatusPending  VerificationStatus = "pending"
	VerificationStatusAccepted VerificationStatus = "accepted"
	VerificationStatusInvalid  VerificationStatus = "invalid"
	VerificationStatusExpired  VerificationStatus = "expired"
	VerificationStatusError    VerificationStatus = "error"
)

// The possible types of verification.
type VerificationType string

const (
	VerificationTypeSMS       VerificationType = "sms"
	VerificationTypeCall      VerificationType = "call"
	VerificationTypeFlashcall VerificationType = "flashcall"
)

type VerificationGetResponse struct {
	Data Verification `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VerificationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerificationTriggerCallParams struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number,required"`
	// The identifier of the associated Verify profile.
	VerifyProfileID string `json:"verify_profile_id,required" format:"uuid"`
	// Send a self-generated numeric code to the end-user
	CustomCode param.Opt[string] `json:"custom_code,omitzero"`
	// The number of seconds the verification code is valid for.
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	paramObj
}

func (r VerificationTriggerCallParams) MarshalJSON() (data []byte, err error) {
	type shadow VerificationTriggerCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationTriggerCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerificationTriggerFlashcallParams struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number,required"`
	// The identifier of the associated Verify profile.
	VerifyProfileID string `json:"verify_profile_id,required" format:"uuid"`
	// The number of seconds the verification code is valid for.
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	paramObj
}

func (r VerificationTriggerFlashcallParams) MarshalJSON() (data []byte, err error) {
	type shadow VerificationTriggerFlashcallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationTriggerFlashcallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerificationTriggerSMSParams struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number,required"`
	// The identifier of the associated Verify profile.
	VerifyProfileID string `json:"verify_profile_id,required" format:"uuid"`
	// Send a self-generated numeric code to the end-user
	CustomCode param.Opt[string] `json:"custom_code,omitzero"`
	// The number of seconds the verification code is valid for.
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	paramObj
}

func (r VerificationTriggerSMSParams) MarshalJSON() (data []byte, err error) {
	type shadow VerificationTriggerSMSParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerificationTriggerSMSParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
