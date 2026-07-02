// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/telnyx-go/v4/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/v4/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/v4/option"
	"github.com/stainless-sdks/telnyx-go/v4/packages/param"
	"github.com/stainless-sdks/telnyx-go/v4/packages/respjson"
)

// Submit and manage the two business references and one financial reference that
// vouch for a DIR. References are contacted to confirm the business identity
// during vetting.
//
// DirReferenceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirReferenceService] method instead.
type DirReferenceService struct {
	Options []option.RequestOption
}

// NewDirReferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDirReferenceService(opts ...option.RequestOption) (r DirReferenceService) {
	r = DirReferenceService{}
	r.Options = opts
	return
}

// Partially update one reference, addressed by the DIR id plus the reference's
// type (business or financial) and slot.
//
// Cosmetic fields (full name, job title, organization, relationship, email) are
// always editable. The phone number and timezone may only be changed while a
// scheduled call has not yet been dialed; if a call is in progress or all attempts
// are complete, those fields are locked. Changing the timezone reschedules any
// pending call into the new local calling window.
func (r *DirReferenceService) Update(ctx context.Context, slot int64, params DirReferenceUpdateParams, opts ...option.RequestOption) (res *DirReferenceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.DirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/references/%v/%v", params.DirID, params.RefType, slot)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List the business and financial references submitted for a DIR.
//
// Returns the two business references (slots 0 and 1) followed by the single
// financial reference. Each entry carries only the customer-supplied details
// (name, title, organization, relationship, phone, email, timezone). Returns an
// empty list when no references were submitted.
func (r *DirReferenceService) List(ctx context.Context, dirID string, opts ...option.RequestOption) (res *DirReferenceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/references", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit the two business references and one financial reference for a DIR.
//
// The DIR's authorizer email must be verified first (see the email-verification
// endpoint). Until it is, this returns `409` and no references are stored.
//
// The request body carries exactly two business references plus one financial
// reference. On success the references are stored and the response echoes them in
// the same shape as the GET. Submitting again converges on the already-stored
// references rather than erroring.
func (r *DirReferenceService) Submit(ctx context.Context, dirID string, body DirReferenceSubmitParams, opts ...option.RequestOption) (res *DirReferenceSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if dirID == "" {
		err = errors.New("missing required dir_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("dir/%s/references", dirID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type DirReferenceUpdateResponse struct {
	// A reference (business or financial) on a DIR, in the customer-facing shape. No
	// internal identifiers are exposed.
	Data DirReferenceUpdateResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference (business or financial) on a DIR, in the customer-facing shape. No
// internal identifiers are exposed.
type DirReferenceUpdateResponseData struct {
	// Full name of the reference contact.
	FullName string `json:"full_name" api:"required"`
	// Reference phone number in E.164 format.
	PhoneE164 string `json:"phone_e164" api:"required"`
	// Always `dir_reference`.
	//
	// Any of "dir_reference".
	RecordType string `json:"record_type" api:"required"`
	// Whether this is a business reference or the financial reference.
	//
	// Any of "business", "financial".
	RefType string `json:"ref_type" api:"required"`
	// Position within the reference type. Business references occupy slots 0 and 1;
	// the financial reference occupies slot 0.
	Slot int64 `json:"slot" api:"required"`
	// IANA timezone id for the reference. Calls are only placed within the reference's
	// local 8am-9pm window.
	Timezone string `json:"timezone" api:"required"`
	// Reference contact email address.
	Email string `json:"email" api:"nullable" format:"email"`
	// Job title of the reference contact.
	JobTitle string `json:"job_title" api:"nullable"`
	// Organization the reference contact belongs to.
	Organization string `json:"organization" api:"nullable"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant string `json:"relationship_to_registrant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName                 respjson.Field
		PhoneE164                respjson.Field
		RecordType               respjson.Field
		RefType                  respjson.Field
		Slot                     respjson.Field
		Timezone                 respjson.Field
		Email                    respjson.Field
		JobTitle                 respjson.Field
		Organization             respjson.Field
		RelationshipToRegistrant respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirReferenceListResponse struct {
	Data []DirReferenceListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceListResponse) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference (business or financial) on a DIR, in the customer-facing shape. No
// internal identifiers are exposed.
type DirReferenceListResponseData struct {
	// Full name of the reference contact.
	FullName string `json:"full_name" api:"required"`
	// Reference phone number in E.164 format.
	PhoneE164 string `json:"phone_e164" api:"required"`
	// Always `dir_reference`.
	//
	// Any of "dir_reference".
	RecordType string `json:"record_type" api:"required"`
	// Whether this is a business reference or the financial reference.
	//
	// Any of "business", "financial".
	RefType string `json:"ref_type" api:"required"`
	// Position within the reference type. Business references occupy slots 0 and 1;
	// the financial reference occupies slot 0.
	Slot int64 `json:"slot" api:"required"`
	// IANA timezone id for the reference. Calls are only placed within the reference's
	// local 8am-9pm window.
	Timezone string `json:"timezone" api:"required"`
	// Reference contact email address.
	Email string `json:"email" api:"nullable" format:"email"`
	// Job title of the reference contact.
	JobTitle string `json:"job_title" api:"nullable"`
	// Organization the reference contact belongs to.
	Organization string `json:"organization" api:"nullable"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant string `json:"relationship_to_registrant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName                 respjson.Field
		PhoneE164                respjson.Field
		RecordType               respjson.Field
		RefType                  respjson.Field
		Slot                     respjson.Field
		Timezone                 respjson.Field
		Email                    respjson.Field
		JobTitle                 respjson.Field
		Organization             respjson.Field
		RelationshipToRegistrant respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceListResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirReferenceSubmitResponse struct {
	Data []DirReferenceSubmitResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference (business or financial) on a DIR, in the customer-facing shape. No
// internal identifiers are exposed.
type DirReferenceSubmitResponseData struct {
	// Full name of the reference contact.
	FullName string `json:"full_name" api:"required"`
	// Reference phone number in E.164 format.
	PhoneE164 string `json:"phone_e164" api:"required"`
	// Always `dir_reference`.
	//
	// Any of "dir_reference".
	RecordType string `json:"record_type" api:"required"`
	// Whether this is a business reference or the financial reference.
	//
	// Any of "business", "financial".
	RefType string `json:"ref_type" api:"required"`
	// Position within the reference type. Business references occupy slots 0 and 1;
	// the financial reference occupies slot 0.
	Slot int64 `json:"slot" api:"required"`
	// IANA timezone id for the reference. Calls are only placed within the reference's
	// local 8am-9pm window.
	Timezone string `json:"timezone" api:"required"`
	// Reference contact email address.
	Email string `json:"email" api:"nullable" format:"email"`
	// Job title of the reference contact.
	JobTitle string `json:"job_title" api:"nullable"`
	// Organization the reference contact belongs to.
	Organization string `json:"organization" api:"nullable"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant string `json:"relationship_to_registrant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FullName                 respjson.Field
		PhoneE164                respjson.Field
		RecordType               respjson.Field
		RefType                  respjson.Field
		Slot                     respjson.Field
		Timezone                 respjson.Field
		Email                    respjson.Field
		JobTitle                 respjson.Field
		Organization             respjson.Field
		RelationshipToRegistrant respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DirReferenceSubmitResponseData) RawJSON() string { return r.JSON.raw }
func (r *DirReferenceSubmitResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirReferenceUpdateParams struct {
	DirID string `path:"dir_id" api:"required" format:"uuid" json:"-"`
	// Any of "business", "financial".
	RefType DirReferenceUpdateParamsRefType `path:"ref_type,omitzero" api:"required" json:"-"`
	// Job title of the reference contact.
	JobTitle param.Opt[string] `json:"job_title,omitzero"`
	// Organization the reference contact belongs to.
	Organization param.Opt[string] `json:"organization,omitzero"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant param.Opt[string] `json:"relationship_to_registrant,omitzero"`
	// Reference contact email address.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Full name of the reference contact.
	FullName param.Opt[string] `json:"full_name,omitzero"`
	// Reference phone number in E.164 format.
	PhoneE164 param.Opt[string] `json:"phone_e164,omitzero"`
	// IANA timezone id for the reference.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r DirReferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DirReferenceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirReferenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DirReferenceUpdateParamsRefType string

const (
	DirReferenceUpdateParamsRefTypeBusiness  DirReferenceUpdateParamsRefType = "business"
	DirReferenceUpdateParamsRefTypeFinancial DirReferenceUpdateParamsRefType = "financial"
)

type DirReferenceSubmitParams struct {
	// Exactly two business references.
	BusinessReferences []DirReferenceSubmitParamsBusinessReference `json:"business_references,omitzero" api:"required"`
	// One reference supplied at submit. The reference type is implied by the field
	// that carries it (business_references vs financial_reference).
	FinancialReference DirReferenceSubmitParamsFinancialReference `json:"financial_reference,omitzero" api:"required"`
	paramObj
}

func (r DirReferenceSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow DirReferenceSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirReferenceSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One reference supplied at submit. The reference type is implied by the field
// that carries it (business_references vs financial_reference).
//
// The properties Email, FullName, PhoneE164, Timezone are required.
type DirReferenceSubmitParamsBusinessReference struct {
	// Reference contact email address. Required: the reference is emailed scheduling
	// and dial-in notices.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the reference contact.
	FullName string `json:"full_name" api:"required"`
	// Reference phone number in E.164 format, e.g. +14155550123.
	PhoneE164 string `json:"phone_e164" api:"required"`
	// IANA timezone id for the reference (e.g. America/New_York). Required: calls are
	// only placed within the reference's local 8am-9pm window.
	Timezone string `json:"timezone" api:"required"`
	// Job title of the reference contact.
	JobTitle param.Opt[string] `json:"job_title,omitzero"`
	// Organization the reference contact belongs to.
	Organization param.Opt[string] `json:"organization,omitzero"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant param.Opt[string] `json:"relationship_to_registrant,omitzero"`
	paramObj
}

func (r DirReferenceSubmitParamsBusinessReference) MarshalJSON() (data []byte, err error) {
	type shadow DirReferenceSubmitParamsBusinessReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirReferenceSubmitParamsBusinessReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One reference supplied at submit. The reference type is implied by the field
// that carries it (business_references vs financial_reference).
//
// The properties Email, FullName, PhoneE164, Timezone are required.
type DirReferenceSubmitParamsFinancialReference struct {
	// Reference contact email address. Required: the reference is emailed scheduling
	// and dial-in notices.
	Email string `json:"email" api:"required" format:"email"`
	// Full name of the reference contact.
	FullName string `json:"full_name" api:"required"`
	// Reference phone number in E.164 format, e.g. +14155550123.
	PhoneE164 string `json:"phone_e164" api:"required"`
	// IANA timezone id for the reference (e.g. America/New_York). Required: calls are
	// only placed within the reference's local 8am-9pm window.
	Timezone string `json:"timezone" api:"required"`
	// Job title of the reference contact.
	JobTitle param.Opt[string] `json:"job_title,omitzero"`
	// Organization the reference contact belongs to.
	Organization param.Opt[string] `json:"organization,omitzero"`
	// How the reference contact is related to the registering business.
	RelationshipToRegistrant param.Opt[string] `json:"relationship_to_registrant,omitzero"`
	paramObj
}

func (r DirReferenceSubmitParamsFinancialReference) MarshalJSON() (data []byte, err error) {
	type shadow DirReferenceSubmitParamsFinancialReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DirReferenceSubmitParamsFinancialReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
