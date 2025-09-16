// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// VerifyProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifyProfileService] method instead.
type VerifyProfileService struct {
	Options []option.RequestOption
}

// NewVerifyProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifyProfileService(opts ...option.RequestOption) (r VerifyProfileService) {
	r = VerifyProfileService{}
	r.Options = opts
	return
}

// Creates a new Verify profile to associate verifications with.
func (r *VerifyProfileService) New(ctx context.Context, body VerifyProfileNewParams, opts ...option.RequestOption) (res *VerifyProfileData, err error) {
	opts = append(r.Options[:], opts...)
	path := "verify_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a single Verify profile.
func (r *VerifyProfileService) Get(ctx context.Context, verifyProfileID string, opts ...option.RequestOption) (res *VerifyProfileData, err error) {
	opts = append(r.Options[:], opts...)
	if verifyProfileID == "" {
		err = errors.New("missing required verify_profile_id parameter")
		return
	}
	path := fmt.Sprintf("verify_profiles/%s", verifyProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Verify profile
func (r *VerifyProfileService) Update(ctx context.Context, verifyProfileID string, body VerifyProfileUpdateParams, opts ...option.RequestOption) (res *VerifyProfileData, err error) {
	opts = append(r.Options[:], opts...)
	if verifyProfileID == "" {
		err = errors.New("missing required verify_profile_id parameter")
		return
	}
	path := fmt.Sprintf("verify_profiles/%s", verifyProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets a paginated list of Verify profiles.
func (r *VerifyProfileService) List(ctx context.Context, query VerifyProfileListParams, opts ...option.RequestOption) (res *VerifyProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "verify_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Verify profile
func (r *VerifyProfileService) Delete(ctx context.Context, verifyProfileID string, opts ...option.RequestOption) (res *VerifyProfileData, err error) {
	opts = append(r.Options[:], opts...)
	if verifyProfileID == "" {
		err = errors.New("missing required verify_profile_id parameter")
		return
	}
	path := fmt.Sprintf("verify_profiles/%s", verifyProfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all Verify profile message templates.
func (r *VerifyProfileService) GetTemplates(ctx context.Context, opts ...option.RequestOption) (res *VerifyProfileGetTemplatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "verify_profiles/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VerifyProfile struct {
	ID        string                 `json:"id" format:"uuid"`
	Call      VerifyProfileCall      `json:"call"`
	CreatedAt string                 `json:"created_at"`
	Flashcall VerifyProfileFlashcall `json:"flashcall"`
	Language  string                 `json:"language"`
	Name      string                 `json:"name"`
	// The possible verification profile record types.
	//
	// Any of "verification_profile".
	RecordType         VerifyProfileRecordType `json:"record_type"`
	SMS                VerifyProfileSMS        `json:"sms"`
	UpdatedAt          string                  `json:"updated_at"`
	WebhookFailoverURL string                  `json:"webhook_failover_url"`
	WebhookURL         string                  `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Call               respjson.Field
		CreatedAt          respjson.Field
		Flashcall          respjson.Field
		Language           respjson.Field
		Name               respjson.Field
		RecordType         respjson.Field
		SMS                respjson.Field
		UpdatedAt          respjson.Field
		WebhookFailoverURL respjson.Field
		WebhookURL         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfile) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileCall struct {
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName string `json:"app_name"`
	// The length of the verify code to generate.
	CodeLength int64 `json:"code_length"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs int64 `json:"default_verification_timeout_secs"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID string `json:"messaging_template_id" format:"uuid"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations"`
	ExtraFields             map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppName                        respjson.Field
		CodeLength                     respjson.Field
		DefaultVerificationTimeoutSecs respjson.Field
		MessagingTemplateID            respjson.Field
		WhitelistedDestinations        respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileCall) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileFlashcall struct {
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs int64          `json:"default_verification_timeout_secs"`
	ExtraFields                    map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultVerificationTimeoutSecs respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileFlashcall) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileFlashcall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The possible verification profile record types.
type VerifyProfileRecordType string

const (
	VerifyProfileRecordTypeVerificationProfile VerifyProfileRecordType = "verification_profile"
)

type VerifyProfileSMS struct {
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender string `json:"alpha_sender,nullable"`
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName string `json:"app_name"`
	// The length of the verify code to generate.
	CodeLength int64 `json:"code_length"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs int64 `json:"default_verification_timeout_secs"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID string `json:"messaging_template_id" format:"uuid"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations"`
	ExtraFields             map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlphaSender                    respjson.Field
		AppName                        respjson.Field
		CodeLength                     respjson.Field
		DefaultVerificationTimeoutSecs respjson.Field
		MessagingTemplateID            respjson.Field
		WhitelistedDestinations        respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileSMS) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileData struct {
	Data VerifyProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileData) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of Verify profiles
type VerifyProfileListResponse struct {
	Data []VerifyProfile               `json:"data,required"`
	Meta VerifyProfileListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of Verify profile message templates
type VerifyProfileGetTemplatesResponse struct {
	Data []VerifyProfileGetTemplatesResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileGetTemplatesResponse) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileGetTemplatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileGetTemplatesResponseData struct {
	ID   string `json:"id" format:"uuid"`
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifyProfileGetTemplatesResponseData) RawJSON() string { return r.JSON.raw }
func (r *VerifyProfileGetTemplatesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileNewParams struct {
	Name               string                          `json:"name,required"`
	Language           param.Opt[string]               `json:"language,omitzero"`
	WebhookFailoverURL param.Opt[string]               `json:"webhook_failover_url,omitzero"`
	WebhookURL         param.Opt[string]               `json:"webhook_url,omitzero"`
	Call               VerifyProfileNewParamsCall      `json:"call,omitzero"`
	Flashcall          VerifyProfileNewParamsFlashcall `json:"flashcall,omitzero"`
	SMS                VerifyProfileNewParamsSMS       `json:"sms,omitzero"`
	paramObj
}

func (r VerifyProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerifyProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileNewParamsCall struct {
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName param.Opt[string] `json:"app_name,omitzero"`
	// The length of the verify code to generate.
	CodeLength param.Opt[int64] `json:"code_length,omitzero"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID param.Opt[string] `json:"messaging_template_id,omitzero" format:"uuid"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations,omitzero"`
	ExtraFields             map[string]any `json:"-"`
	paramObj
}

func (r VerifyProfileNewParamsCall) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileNewParamsCall
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileNewParamsCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileNewParamsFlashcall struct {
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations,omitzero"`
	ExtraFields             map[string]any `json:"-"`
	paramObj
}

func (r VerifyProfileNewParamsFlashcall) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileNewParamsFlashcall
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileNewParamsFlashcall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property WhitelistedDestinations is required.
type VerifyProfileNewParamsSMS struct {
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitzero,required"`
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender param.Opt[string] `json:"alpha_sender,omitzero"`
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName param.Opt[string] `json:"app_name,omitzero"`
	// The length of the verify code to generate.
	CodeLength param.Opt[int64] `json:"code_length,omitzero"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID param.Opt[string] `json:"messaging_template_id,omitzero" format:"uuid"`
	ExtraFields         map[string]any    `json:"-"`
	paramObj
}

func (r VerifyProfileNewParamsSMS) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileNewParamsSMS
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileNewParamsSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileUpdateParams struct {
	Language           param.Opt[string]                  `json:"language,omitzero"`
	Name               param.Opt[string]                  `json:"name,omitzero"`
	WebhookFailoverURL param.Opt[string]                  `json:"webhook_failover_url,omitzero"`
	WebhookURL         param.Opt[string]                  `json:"webhook_url,omitzero"`
	Call               VerifyProfileUpdateParamsCall      `json:"call,omitzero"`
	Flashcall          VerifyProfileUpdateParamsFlashcall `json:"flashcall,omitzero"`
	SMS                VerifyProfileUpdateParamsSMS       `json:"sms,omitzero"`
	paramObj
}

func (r VerifyProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VerifyProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileUpdateParamsCall struct {
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName param.Opt[string] `json:"app_name,omitzero"`
	// The length of the verify code to generate.
	CodeLength param.Opt[int64] `json:"code_length,omitzero"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID param.Opt[string] `json:"messaging_template_id,omitzero" format:"uuid"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations,omitzero"`
	ExtraFields             map[string]any `json:"-"`
	paramObj
}

func (r VerifyProfileUpdateParamsCall) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileUpdateParamsCall
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileUpdateParamsCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileUpdateParamsFlashcall struct {
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations,omitzero"`
	ExtraFields             map[string]any `json:"-"`
	paramObj
}

func (r VerifyProfileUpdateParamsFlashcall) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileUpdateParamsFlashcall
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileUpdateParamsFlashcall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileUpdateParamsSMS struct {
	// The alphanumeric sender ID to use when sending to destinations that require an
	// alphanumeric sender ID.
	AlphaSender param.Opt[string] `json:"alpha_sender,omitzero"`
	// The name that identifies the application requesting 2fa in the verification
	// message.
	AppName param.Opt[string] `json:"app_name,omitzero"`
	// The length of the verify code to generate.
	CodeLength param.Opt[int64] `json:"code_length,omitzero"`
	// For every request that is initiated via this Verify profile, this sets the
	// number of seconds before a verification request code expires. Once the
	// verification request expires, the user cannot use the code to verify their
	// identity.
	DefaultVerificationTimeoutSecs param.Opt[int64] `json:"default_verification_timeout_secs,omitzero"`
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateID param.Opt[string] `json:"messaging_template_id,omitzero" format:"uuid"`
	// Enabled country destinations to send verification codes. The elements in the
	// list must be valid ISO 3166-1 alpha-2 country codes. If set to `["*"]`, all
	// destinations will be allowed.
	WhitelistedDestinations []string       `json:"whitelisted_destinations,omitzero"`
	ExtraFields             map[string]any `json:"-"`
	paramObj
}

func (r VerifyProfileUpdateParamsSMS) MarshalJSON() (data []byte, err error) {
	type shadow VerifyProfileUpdateParamsSMS
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *VerifyProfileUpdateParamsSMS) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifyProfileListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[name]
	Filter VerifyProfileListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page VerifyProfileListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VerifyProfileListParams]'s query parameters as
// `url.Values`.
func (r VerifyProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[name]
type VerifyProfileListParamsFilter struct {
	// Optional filter for profile names.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VerifyProfileListParamsFilter]'s query parameters as
// `url.Values`.
func (r VerifyProfileListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type VerifyProfileListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VerifyProfileListParamsPage]'s query parameters as
// `url.Values`.
func (r VerifyProfileListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
