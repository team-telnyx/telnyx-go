// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Requirement Groups
//
// RequirementGroupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequirementGroupService] method instead.
type RequirementGroupService struct {
	Options []option.RequestOption
}

// NewRequirementGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRequirementGroupService(opts ...option.RequestOption) (r RequirementGroupService) {
	r = RequirementGroupService{}
	r.Options = opts
	return
}

// Create a new requirement group
func (r *RequirementGroupService) New(ctx context.Context, body RequirementGroupNewParams, opts ...option.RequestOption) (res *RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "requirement_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single requirement group by ID
func (r *RequirementGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirement_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update requirement values in requirement group
func (r *RequirementGroupService) Update(ctx context.Context, id string, body RequirementGroupUpdateParams, opts ...option.RequestOption) (res *RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirement_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List requirement groups
func (r *RequirementGroupService) List(ctx context.Context, query RequirementGroupListParams, opts ...option.RequestOption) (res *[]RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "requirement_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a requirement group by ID
func (r *RequirementGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirement_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Submit a Requirement Group for Approval
func (r *RequirementGroupService) SubmitForApproval(ctx context.Context, id string, opts ...option.RequestOption) (res *RequirementGroup, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("requirement_groups/%s/submit_for_approval", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type RequirementGroup struct {
	ID                     string            `json:"id"`
	Action                 string            `json:"action"`
	CountryCode            string            `json:"country_code"`
	CreatedAt              time.Time         `json:"created_at" format:"date-time"`
	CustomerReference      string            `json:"customer_reference"`
	PhoneNumberType        string            `json:"phone_number_type"`
	RecordType             string            `json:"record_type"`
	RegulatoryRequirements []UserRequirement `json:"regulatory_requirements"`
	// Any of "approved", "unapproved", "pending-approval", "declined", "expired".
	Status    RequirementGroupStatus `json:"status"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Action                 respjson.Field
		CountryCode            respjson.Field
		CreatedAt              respjson.Field
		CustomerReference      respjson.Field
		PhoneNumberType        respjson.Field
		RecordType             respjson.Field
		RegulatoryRequirements respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequirementGroup) RawJSON() string { return r.JSON.raw }
func (r *RequirementGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGroupStatus string

const (
	RequirementGroupStatusApproved        RequirementGroupStatus = "approved"
	RequirementGroupStatusUnapproved      RequirementGroupStatus = "unapproved"
	RequirementGroupStatusPendingApproval RequirementGroupStatus = "pending-approval"
	RequirementGroupStatusDeclined        RequirementGroupStatus = "declined"
	RequirementGroupStatusExpired         RequirementGroupStatus = "expired"
)

type UserRequirement struct {
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	ExpiresAt     time.Time `json:"expires_at" format:"date-time"`
	FieldType     string    `json:"field_type"`
	FieldValue    string    `json:"field_value"`
	RequirementID string    `json:"requirement_id"`
	// Any of "approved", "unapproved", "pending-approval", "declined", "expired".
	Status    UserRequirementStatus `json:"status"`
	UpdatedAt time.Time             `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		FieldType     respjson.Field
		FieldValue    respjson.Field
		RequirementID respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserRequirement) RawJSON() string { return r.JSON.raw }
func (r *UserRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRequirementStatus string

const (
	UserRequirementStatusApproved        UserRequirementStatus = "approved"
	UserRequirementStatusUnapproved      UserRequirementStatus = "unapproved"
	UserRequirementStatusPendingApproval UserRequirementStatus = "pending-approval"
	UserRequirementStatusDeclined        UserRequirementStatus = "declined"
	UserRequirementStatusExpired         UserRequirementStatus = "expired"
)

type RequirementGroupNewParams struct {
	// Any of "ordering", "porting".
	Action RequirementGroupNewParamsAction `json:"action,omitzero" api:"required"`
	// ISO alpha 2 country code
	CountryCode string `json:"country_code" api:"required"`
	// Any of "local", "toll_free", "mobile", "national", "shared_cost".
	PhoneNumberType        RequirementGroupNewParamsPhoneNumberType         `json:"phone_number_type,omitzero" api:"required"`
	CustomerReference      param.Opt[string]                                `json:"customer_reference,omitzero"`
	RegulatoryRequirements []RequirementGroupNewParamsRegulatoryRequirement `json:"regulatory_requirements,omitzero"`
	paramObj
}

func (r RequirementGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RequirementGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequirementGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGroupNewParamsAction string

const (
	RequirementGroupNewParamsActionOrdering RequirementGroupNewParamsAction = "ordering"
	RequirementGroupNewParamsActionPorting  RequirementGroupNewParamsAction = "porting"
)

type RequirementGroupNewParamsPhoneNumberType string

const (
	RequirementGroupNewParamsPhoneNumberTypeLocal      RequirementGroupNewParamsPhoneNumberType = "local"
	RequirementGroupNewParamsPhoneNumberTypeTollFree   RequirementGroupNewParamsPhoneNumberType = "toll_free"
	RequirementGroupNewParamsPhoneNumberTypeMobile     RequirementGroupNewParamsPhoneNumberType = "mobile"
	RequirementGroupNewParamsPhoneNumberTypeNational   RequirementGroupNewParamsPhoneNumberType = "national"
	RequirementGroupNewParamsPhoneNumberTypeSharedCost RequirementGroupNewParamsPhoneNumberType = "shared_cost"
)

type RequirementGroupNewParamsRegulatoryRequirement struct {
	FieldValue    param.Opt[string] `json:"field_value,omitzero"`
	RequirementID param.Opt[string] `json:"requirement_id,omitzero"`
	paramObj
}

func (r RequirementGroupNewParamsRegulatoryRequirement) MarshalJSON() (data []byte, err error) {
	type shadow RequirementGroupNewParamsRegulatoryRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequirementGroupNewParamsRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGroupUpdateParams struct {
	// Reference for the customer
	CustomerReference      param.Opt[string]                                   `json:"customer_reference,omitzero"`
	RegulatoryRequirements []RequirementGroupUpdateParamsRegulatoryRequirement `json:"regulatory_requirements,omitzero"`
	paramObj
}

func (r RequirementGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RequirementGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequirementGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGroupUpdateParamsRegulatoryRequirement struct {
	// New value for the regulatory requirement
	FieldValue param.Opt[string] `json:"field_value,omitzero"`
	// Unique identifier for the regulatory requirement
	RequirementID param.Opt[string] `json:"requirement_id,omitzero"`
	paramObj
}

func (r RequirementGroupUpdateParamsRegulatoryRequirement) MarshalJSON() (data []byte, err error) {
	type shadow RequirementGroupUpdateParamsRegulatoryRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequirementGroupUpdateParamsRegulatoryRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequirementGroupListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[country_code], filter[phone_number_type], filter[action], filter[status],
	// filter[customer_reference]
	Filter RequirementGroupListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementGroupListParams]'s query parameters as
// `url.Values`.
func (r RequirementGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[country_code], filter[phone_number_type], filter[action], filter[status],
// filter[customer_reference]
type RequirementGroupListParamsFilter struct {
	// Filter requirement groups by country code (iso alpha 2)
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter requirement groups by customer reference
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter requirement groups by action type
	//
	// Any of "ordering", "porting", "action".
	Action string `query:"action,omitzero" json:"-"`
	// Filter requirement groups by phone number type.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost".
	PhoneNumberType string `query:"phone_number_type,omitzero" json:"-"`
	// Filter requirement groups by status
	//
	// Any of "approved", "unapproved", "pending-approval", "declined", "expired".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequirementGroupListParamsFilter]'s query parameters as
// `url.Values`.
func (r RequirementGroupListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
