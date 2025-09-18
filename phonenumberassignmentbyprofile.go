// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PhoneNumberAssignmentByProfileService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberAssignmentByProfileService] method instead.
type PhoneNumberAssignmentByProfileService struct {
	Options []option.RequestOption
}

// NewPhoneNumberAssignmentByProfileService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPhoneNumberAssignmentByProfileService(opts ...option.RequestOption) (r PhoneNumberAssignmentByProfileService) {
	r = PhoneNumberAssignmentByProfileService{}
	r.Options = opts
	return
}

// This endpoint allows you to link all phone numbers associated with a Messaging
// Profile to a campaign. **Please note:** if you want to assign phone numbers to a
// campaign that you did not create with Telnyx 10DLC services, this endpoint
// allows that provided that you've shared the campaign with Telnyx. In this case,
// only provide the parameter, `tcrCampaignId`, and not `campaignId`. In all other
// cases (where the campaign you're assigning was created with Telnyx 10DLC
// services), only provide `campaignId`, not `tcrCampaignId`.
func (r *PhoneNumberAssignmentByProfileService) Assign(ctx context.Context, body PhoneNumberAssignmentByProfileAssignParams, opts ...option.RequestOption) (res *PhoneNumberAssignmentByProfileAssignResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "phoneNumberAssignmentByProfile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the status of the individual phone number/campaign assignments associated
// with the supplied `taskId`.
func (r *PhoneNumberAssignmentByProfileService) GetPhoneNumberStatus(ctx context.Context, taskID string, query PhoneNumberAssignmentByProfileGetPhoneNumberStatusParams, opts ...option.RequestOption) (res *PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("phoneNumberAssignmentByProfile/%s/phoneNumbers", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check the status of the task associated with assigning all phone numbers on a
// messaging profile to a campaign by `taskId`.
func (r *PhoneNumberAssignmentByProfileService) GetStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *PhoneNumberAssignmentByProfileGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("phoneNumberAssignmentByProfile/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusStarting  TaskStatus = "starting"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

// PhoneNumberAssignmentByProfileAssignResponseUnion contains all possible
// properties and values from
// [PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse],
// [PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PhoneNumberAssignmentByProfileAssignResponseUnion struct {
	// This field is from variant
	// [PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse].
	MessagingProfileID string `json:"messagingProfileId"`
	// This field is from variant
	// [PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse].
	TaskID string `json:"taskId"`
	// This field is from variant
	// [PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse].
	CampaignID string `json:"campaignId"`
	// This field is from variant
	// [PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse].
	TcrCampaignID string `json:"tcrCampaignId"`
	// This field is from variant
	// [PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage].
	Message string `json:"message"`
	JSON    struct {
		MessagingProfileID respjson.Field
		TaskID             respjson.Field
		CampaignID         respjson.Field
		TcrCampaignID      respjson.Field
		Message            respjson.Field
		raw                string
	} `json:"-"`
}

func (u PhoneNumberAssignmentByProfileAssignResponseUnion) AsAssignProfileToCampaignResponse() (v PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PhoneNumberAssignmentByProfileAssignResponseUnion) AsSettingsDataErrorMessage() (v PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PhoneNumberAssignmentByProfileAssignResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *PhoneNumberAssignmentByProfileAssignResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse struct {
	// The ID of the messaging profile that you want to link to the specified campaign.
	MessagingProfileID string `json:"messagingProfileId,required"`
	// The ID of the task associated with assigning a messaging profile to a campaign.
	TaskID string `json:"taskId,required"`
	// The ID of the campaign you want to link to the specified messaging profile. If
	// you supply this ID in the request, do not also include a tcrCampaignId.
	CampaignID string `json:"campaignId"`
	// The TCR ID of the shared campaign you want to link to the specified messaging
	// profile (for campaigns not created using Telnyx 10DLC services only). If you
	// supply this ID in the request, do not also include a campaignId.
	TcrCampaignID string `json:"tcrCampaignId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessagingProfileID respjson.Field
		TaskID             respjson.Field
		CampaignID         respjson.Field
		TcrCampaignID      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberAssignmentByProfileAssignResponseAssignProfileToCampaignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberAssignmentByProfileAssignResponseSettingsDataErrorMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse struct {
	Records []PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord `json:"records,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Records     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord struct {
	// The phone number that the status is being checked for.
	PhoneNumber string `json:"phoneNumber,required"`
	// The status of the associated phone number assignment.
	Status string `json:"status,required"`
	// The ID of the task associated with the phone number.
	TaskID string `json:"taskId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileGetStatusResponse struct {
	// An enumeration.
	//
	// Any of "pending", "processing", "completed", "failed".
	Status    PhoneNumberAssignmentByProfileGetStatusResponseStatus `json:"status,required"`
	TaskID    string                                                `json:"taskId,required"`
	CreatedAt time.Time                                             `json:"createdAt" format:"date-time"`
	UpdatedAt time.Time                                             `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAssignmentByProfileGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberAssignmentByProfileGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An enumeration.
type PhoneNumberAssignmentByProfileGetStatusResponseStatus string

const (
	PhoneNumberAssignmentByProfileGetStatusResponseStatusPending    PhoneNumberAssignmentByProfileGetStatusResponseStatus = "pending"
	PhoneNumberAssignmentByProfileGetStatusResponseStatusProcessing PhoneNumberAssignmentByProfileGetStatusResponseStatus = "processing"
	PhoneNumberAssignmentByProfileGetStatusResponseStatusCompleted  PhoneNumberAssignmentByProfileGetStatusResponseStatus = "completed"
	PhoneNumberAssignmentByProfileGetStatusResponseStatusFailed     PhoneNumberAssignmentByProfileGetStatusResponseStatus = "failed"
)

type PhoneNumberAssignmentByProfileAssignParams struct {
	// The ID of the messaging profile that you want to link to the specified campaign.
	MessagingProfileID string `json:"messagingProfileId,required"`
	// The ID of the campaign you want to link to the specified messaging profile. If
	// you supply this ID in the request, do not also include a tcrCampaignId.
	CampaignID param.Opt[string] `json:"campaignId,omitzero"`
	// The TCR ID of the shared campaign you want to link to the specified messaging
	// profile (for campaigns not created using Telnyx 10DLC services only). If you
	// supply this ID in the request, do not also include a campaignId.
	TcrCampaignID param.Opt[string] `json:"tcrCampaignId,omitzero"`
	paramObj
}

func (r PhoneNumberAssignmentByProfileAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberAssignmentByProfileAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberAssignmentByProfileAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAssignmentByProfileGetPhoneNumberStatusParams struct {
	Page           param.Opt[int64] `query:"page,omitzero" json:"-"`
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberAssignmentByProfileGetPhoneNumberStatusParams]'s
// query parameters as `url.Values`.
func (r PhoneNumberAssignmentByProfileGetPhoneNumberStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
