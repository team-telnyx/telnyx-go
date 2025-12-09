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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcPhoneNumberAssignmentByProfileService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcPhoneNumberAssignmentByProfileService] method instead.
type Number10dlcPhoneNumberAssignmentByProfileService struct {
	Options []option.RequestOption
}

// NewNumber10dlcPhoneNumberAssignmentByProfileService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewNumber10dlcPhoneNumberAssignmentByProfileService(opts ...option.RequestOption) (r Number10dlcPhoneNumberAssignmentByProfileService) {
	r = Number10dlcPhoneNumberAssignmentByProfileService{}
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
func (r *Number10dlcPhoneNumberAssignmentByProfileService) Assign(ctx context.Context, body Number10dlcPhoneNumberAssignmentByProfileAssignParams, opts ...option.RequestOption) (res *Number10dlcPhoneNumberAssignmentByProfileAssignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/phoneNumberAssignmentByProfile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Check the status of the individual phone number/campaign assignments associated
// with the supplied `taskId`.
func (r *Number10dlcPhoneNumberAssignmentByProfileService) GetPhoneNumberStatus(ctx context.Context, taskID string, query Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusParams, opts ...option.RequestOption) (res *Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/phoneNumberAssignmentByProfile/%s/phoneNumbers", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Check the status of the task associated with assigning all phone numbers on a
// messaging profile to a campaign by `taskId`.
func (r *Number10dlcPhoneNumberAssignmentByProfileService) GetTaskStatus(ctx context.Context, taskID string, opts ...option.RequestOption) (res *Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/phoneNumberAssignmentByProfile/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Number10dlcPhoneNumberAssignmentByProfileAssignResponse struct {
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
func (r Number10dlcPhoneNumberAssignmentByProfileAssignResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcPhoneNumberAssignmentByProfileAssignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse struct {
	Records []Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord `json:"records,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Records     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord struct {
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
func (r Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord) RawJSON() string {
	return r.JSON.raw
}
func (r *Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusResponseRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponse struct {
	// An enumeration.
	//
	// Any of "pending", "processing", "completed", "failed".
	Status    Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus `json:"status,required"`
	TaskID    string                                                               `json:"taskId,required"`
	CreatedAt time.Time                                                            `json:"createdAt" format:"date-time"`
	UpdatedAt time.Time                                                            `json:"updatedAt" format:"date-time"`
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
func (r Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An enumeration.
type Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus string

const (
	Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatusPending    Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus = "pending"
	Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatusProcessing Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus = "processing"
	Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatusCompleted  Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus = "completed"
	Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatusFailed     Number10dlcPhoneNumberAssignmentByProfileGetTaskStatusResponseStatus = "failed"
)

type Number10dlcPhoneNumberAssignmentByProfileAssignParams struct {
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

func (r Number10dlcPhoneNumberAssignmentByProfileAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcPhoneNumberAssignmentByProfileAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcPhoneNumberAssignmentByProfileAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusParams struct {
	Page           param.Opt[int64] `query:"page,omitzero" json:"-"`
	RecordsPerPage param.Opt[int64] `query:"recordsPerPage,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusParams]'s query
// parameters as `url.Values`.
func (r Number10dlcPhoneNumberAssignmentByProfileGetPhoneNumberStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
