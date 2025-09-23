// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PhoneNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberService] method instead.
type PhoneNumberService struct {
	Options      []option.RequestOption
	Actions      PhoneNumberActionService
	CsvDownloads PhoneNumberCsvDownloadService
	Jobs         PhoneNumberJobService
	Messaging    PhoneNumberMessagingService
	Voice        PhoneNumberVoiceService
	Voicemail    PhoneNumberVoicemailService
}

// NewPhoneNumberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhoneNumberService(opts ...option.RequestOption) (r PhoneNumberService) {
	r = PhoneNumberService{}
	r.Options = opts
	r.Actions = NewPhoneNumberActionService(opts...)
	r.CsvDownloads = NewPhoneNumberCsvDownloadService(opts...)
	r.Jobs = NewPhoneNumberJobService(opts...)
	r.Messaging = NewPhoneNumberMessagingService(opts...)
	r.Voice = NewPhoneNumberVoiceService(opts...)
	r.Voicemail = NewPhoneNumberVoicemailService(opts...)
	return
}

// Retrieve a phone number
func (r *PhoneNumberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a phone number
func (r *PhoneNumberService) Update(ctx context.Context, id string, body PhoneNumberUpdateParams, opts ...option.RequestOption) (res *PhoneNumberUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List phone numbers
func (r *PhoneNumberService) List(ctx context.Context, query PhoneNumberListParams, opts ...option.RequestOption) (res *PhoneNumberListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a phone number
func (r *PhoneNumberService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List phone numbers, This endpoint is a lighter version of the /phone_numbers
// endpoint having higher performance and rate limit.
func (r *PhoneNumberService) SlimList(ctx context.Context, query PhoneNumberSlimListParams, opts ...option.RequestOption) (res *PhoneNumberSlimListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/slim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PhoneNumberDetailed struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Identifies the billing group associated with the phone number.
	BillingGroupID string `json:"billing_group_id"`
	// Indicates if call forwarding will be enabled for this number if forwards_to and
	// forwarding_type are filled in. Defaults to true for backwards compatibility with
	// APIV1 use of numbers endpoints.
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Indicates whether call recording is enabled for this number.
	CallRecordingEnabled bool `json:"call_recording_enabled"`
	// Indicates whether caller ID is enabled for this number.
	CallerIDNameEnabled bool `json:"caller_id_name_enabled"`
	// Indicates whether a CNAM listing is enabled for this number.
	CnamListingEnabled bool `json:"cnam_listing_enabled"`
	// Identifies the connection associated with the phone number.
	ConnectionID string `json:"connection_id"`
	// The user-assigned name of the connection to be associated with this phone
	// number.
	ConnectionName string `json:"connection_name"`
	// The ISO 3166-1 alpha-2 country code of the phone number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Indicates whether deletion lock is enabled for this number. When enabled, this
	// prevents the phone number from being deleted via the API or Telnyx portal.
	DeletionLockEnabled bool `json:"deletion_lock_enabled"`
	// Identifies the emergency address associated with the phone number.
	EmergencyAddressID string `json:"emergency_address_id"`
	// Indicates whether emergency services are enabled for this number.
	EmergencyEnabled bool `json:"emergency_enabled"`
	// Indicates the status of the provisioning of emergency services for the phone
	// number. This field contains information about activity that may be ongoing for a
	// number where it either is being provisioned or deprovisioned but is not yet
	// enabled/disabled.
	//
	// Any of "active", "deprovisioning", "disabled", "provisioning",
	// "provisioning-failed".
	EmergencyStatus PhoneNumberDetailedEmergencyStatus `json:"emergency_status"`
	// If someone attempts to port your phone number away from Telnyx and your phone
	// number has an external PIN set, Telnyx will attempt to verify that you provided
	// the correct external PIN to the winning carrier. Note that not all carriers
	// cooperate with this security mechanism.
	ExternalPin string `json:"external_pin"`
	// The inbound_call_screening setting is a phone number configuration option
	// variable that allows users to configure their settings to block or flag
	// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
	// feature has an additional per-number monthly cost associated with it.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening PhoneNumberDetailedInboundCallScreening `json:"inbound_call_screening"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The name of the messaging profile associated with the phone number.
	MessagingProfileName string `json:"messaging_profile_name"`
	// The +E.164-formatted phone number associated with this record.
	PhoneNumber string `json:"phone_number"`
	// The phone number's type. Note: For numbers purchased prior to July 2023 or when
	// fetching a number's details immediately after a purchase completes, the legacy
	// values `tollfree`, `shortcode` or `longcode` may be returned instead.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost", "landline",
	// "tollfree", "shortcode", "longcode".
	PhoneNumberType PhoneNumberDetailedPhoneNumberType `json:"phone_number_type"`
	// ISO 8601 formatted date indicating when the resource was purchased.
	PurchasedAt string `json:"purchased_at"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates if the phone number was purchased or ported in. For some numbers this
	// information may not be available.
	//
	// Any of "number_order", "port_request".
	SourceType PhoneNumberDetailedSourceType `json:"source_type,nullable"`
	// The phone number's current status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "port-failed",
	// "active", "deleted", "emergency-only", "ported-out", "port-out-pending",
	// "requirement-info-pending", "requirement-info-under-review",
	// "requirement-info-exception", "provision-pending".
	Status PhoneNumberDetailedStatus `json:"status"`
	// Indicates whether T38 Fax Gateway for inbound calls to this number.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled"`
	// A list of user-assigned tags to help manage the phone number.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BillingGroupID        respjson.Field
		CallForwardingEnabled respjson.Field
		CallRecordingEnabled  respjson.Field
		CallerIDNameEnabled   respjson.Field
		CnamListingEnabled    respjson.Field
		ConnectionID          respjson.Field
		ConnectionName        respjson.Field
		CountryISOAlpha2      respjson.Field
		CreatedAt             respjson.Field
		CustomerReference     respjson.Field
		DeletionLockEnabled   respjson.Field
		EmergencyAddressID    respjson.Field
		EmergencyEnabled      respjson.Field
		EmergencyStatus       respjson.Field
		ExternalPin           respjson.Field
		InboundCallScreening  respjson.Field
		MessagingProfileID    respjson.Field
		MessagingProfileName  respjson.Field
		PhoneNumber           respjson.Field
		PhoneNumberType       respjson.Field
		PurchasedAt           respjson.Field
		RecordType            respjson.Field
		SourceType            respjson.Field
		Status                respjson.Field
		T38FaxGatewayEnabled  respjson.Field
		Tags                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberDetailed) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberDetailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the status of the provisioning of emergency services for the phone
// number. This field contains information about activity that may be ongoing for a
// number where it either is being provisioned or deprovisioned but is not yet
// enabled/disabled.
type PhoneNumberDetailedEmergencyStatus string

const (
	PhoneNumberDetailedEmergencyStatusActive             PhoneNumberDetailedEmergencyStatus = "active"
	PhoneNumberDetailedEmergencyStatusDeprovisioning     PhoneNumberDetailedEmergencyStatus = "deprovisioning"
	PhoneNumberDetailedEmergencyStatusDisabled           PhoneNumberDetailedEmergencyStatus = "disabled"
	PhoneNumberDetailedEmergencyStatusProvisioning       PhoneNumberDetailedEmergencyStatus = "provisioning"
	PhoneNumberDetailedEmergencyStatusProvisioningFailed PhoneNumberDetailedEmergencyStatus = "provisioning-failed"
)

// The inbound_call_screening setting is a phone number configuration option
// variable that allows users to configure their settings to block or flag
// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
// feature has an additional per-number monthly cost associated with it.
type PhoneNumberDetailedInboundCallScreening string

const (
	PhoneNumberDetailedInboundCallScreeningDisabled    PhoneNumberDetailedInboundCallScreening = "disabled"
	PhoneNumberDetailedInboundCallScreeningRejectCalls PhoneNumberDetailedInboundCallScreening = "reject_calls"
	PhoneNumberDetailedInboundCallScreeningFlagCalls   PhoneNumberDetailedInboundCallScreening = "flag_calls"
)

// The phone number's type. Note: For numbers purchased prior to July 2023 or when
// fetching a number's details immediately after a purchase completes, the legacy
// values `tollfree`, `shortcode` or `longcode` may be returned instead.
type PhoneNumberDetailedPhoneNumberType string

const (
	PhoneNumberDetailedPhoneNumberTypeLocal      PhoneNumberDetailedPhoneNumberType = "local"
	PhoneNumberDetailedPhoneNumberTypeTollFree   PhoneNumberDetailedPhoneNumberType = "toll_free"
	PhoneNumberDetailedPhoneNumberTypeMobile     PhoneNumberDetailedPhoneNumberType = "mobile"
	PhoneNumberDetailedPhoneNumberTypeNational   PhoneNumberDetailedPhoneNumberType = "national"
	PhoneNumberDetailedPhoneNumberTypeSharedCost PhoneNumberDetailedPhoneNumberType = "shared_cost"
	PhoneNumberDetailedPhoneNumberTypeLandline   PhoneNumberDetailedPhoneNumberType = "landline"
	PhoneNumberDetailedPhoneNumberTypeTollfree   PhoneNumberDetailedPhoneNumberType = "tollfree"
	PhoneNumberDetailedPhoneNumberTypeShortcode  PhoneNumberDetailedPhoneNumberType = "shortcode"
	PhoneNumberDetailedPhoneNumberTypeLongcode   PhoneNumberDetailedPhoneNumberType = "longcode"
)

// Indicates if the phone number was purchased or ported in. For some numbers this
// information may not be available.
type PhoneNumberDetailedSourceType string

const (
	PhoneNumberDetailedSourceTypeNumberOrder PhoneNumberDetailedSourceType = "number_order"
	PhoneNumberDetailedSourceTypePortRequest PhoneNumberDetailedSourceType = "port_request"
)

// The phone number's current status.
type PhoneNumberDetailedStatus string

const (
	PhoneNumberDetailedStatusPurchasePending            PhoneNumberDetailedStatus = "purchase-pending"
	PhoneNumberDetailedStatusPurchaseFailed             PhoneNumberDetailedStatus = "purchase-failed"
	PhoneNumberDetailedStatusPortPending                PhoneNumberDetailedStatus = "port-pending"
	PhoneNumberDetailedStatusPortFailed                 PhoneNumberDetailedStatus = "port-failed"
	PhoneNumberDetailedStatusActive                     PhoneNumberDetailedStatus = "active"
	PhoneNumberDetailedStatusDeleted                    PhoneNumberDetailedStatus = "deleted"
	PhoneNumberDetailedStatusEmergencyOnly              PhoneNumberDetailedStatus = "emergency-only"
	PhoneNumberDetailedStatusPortedOut                  PhoneNumberDetailedStatus = "ported-out"
	PhoneNumberDetailedStatusPortOutPending             PhoneNumberDetailedStatus = "port-out-pending"
	PhoneNumberDetailedStatusRequirementInfoPending     PhoneNumberDetailedStatus = "requirement-info-pending"
	PhoneNumberDetailedStatusRequirementInfoUnderReview PhoneNumberDetailedStatus = "requirement-info-under-review"
	PhoneNumberDetailedStatusRequirementInfoException   PhoneNumberDetailedStatus = "requirement-info-exception"
	PhoneNumberDetailedStatusProvisionPending           PhoneNumberDetailedStatus = "provision-pending"
)

type PhoneNumberGetResponse struct {
	Data PhoneNumberDetailed `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberUpdateResponse struct {
	Data PhoneNumberDetailed `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListResponse struct {
	Data []PhoneNumberDetailed `json:"data"`
	Meta PaginationMeta        `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberDeleteResponse struct {
	Data PhoneNumberDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberDeleteResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Identifies the billing group associated with the phone number.
	BillingGroupID string `json:"billing_group_id"`
	// Indicates if call forwarding will be enabled for this number if forwards_to and
	// forwarding_type are filled in. Defaults to true for backwards compatibility with
	// APIV1 use of numbers endpoints.
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Indicates whether call recording is enabled for this number.
	CallRecordingEnabled bool `json:"call_recording_enabled"`
	// Indicates whether caller ID is enabled for this number.
	CallerIDNameEnabled bool `json:"caller_id_name_enabled"`
	// Indicates whether a CNAM listing is enabled for this number.
	CnamListingEnabled bool `json:"cnam_listing_enabled"`
	// Identifies the connection associated with the phone number.
	ConnectionID string `json:"connection_id"`
	// The user-assigned name of the connection to be associated with this phone
	// number.
	ConnectionName string `json:"connection_name"`
	// ISO 8601 formatted date indicating when the time it took to activate after the
	// purchase.
	CreatedAt string `json:"created_at"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Indicates whether deletion lock is enabled for this number. When enabled, this
	// prevents the phone number from being deleted via the API or Telnyx portal.
	DeletionLockEnabled bool `json:"deletion_lock_enabled"`
	// Identifies the emergency address associated with the phone number.
	EmergencyAddressID string `json:"emergency_address_id"`
	// Indicates whether emergency services are enabled for this number.
	EmergencyEnabled bool `json:"emergency_enabled"`
	// If someone attempts to port your phone number away from Telnyx and your phone
	// number has an external PIN set, Telnyx will attempt to verify that you provided
	// the correct external PIN to the winning carrier. Note that not all carriers
	// cooperate with this security mechanism.
	ExternalPin string `json:"external_pin"`
	// Identifies the messaging profile associated with the phone number.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The name of the messaging profile associated with the phone number.
	MessagingProfileName string `json:"messaging_profile_name"`
	// The +E.164-formatted phone number associated with this record.
	PhoneNumber string `json:"phone_number"`
	// The phone number's type.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost", "landline".
	PhoneNumberType string `json:"phone_number_type"`
	// ISO 8601 formatted date indicating the time the request was made to purchase the
	// number.
	PurchasedAt string `json:"purchased_at"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The phone number's current status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "port-failed",
	// "active", "deleted", "emergency-only", "ported-out", "port-out-pending".
	Status string `json:"status"`
	// Indicates whether T38 Fax Gateway for inbound calls to this number.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled"`
	// A list of user-assigned tags to help manage the phone number.
	Tags []string `json:"tags"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BillingGroupID        respjson.Field
		CallForwardingEnabled respjson.Field
		CallRecordingEnabled  respjson.Field
		CallerIDNameEnabled   respjson.Field
		CnamListingEnabled    respjson.Field
		ConnectionID          respjson.Field
		ConnectionName        respjson.Field
		CreatedAt             respjson.Field
		CustomerReference     respjson.Field
		DeletionLockEnabled   respjson.Field
		EmergencyAddressID    respjson.Field
		EmergencyEnabled      respjson.Field
		ExternalPin           respjson.Field
		MessagingProfileID    respjson.Field
		MessagingProfileName  respjson.Field
		PhoneNumber           respjson.Field
		PhoneNumberType       respjson.Field
		PurchasedAt           respjson.Field
		RecordType            respjson.Field
		Status                respjson.Field
		T38FaxGatewayEnabled  respjson.Field
		Tags                  respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberSlimListResponse struct {
	Data []PhoneNumberSlimListResponseData `json:"data"`
	Meta PaginationMeta                    `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberSlimListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberSlimListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberSlimListResponseData struct {
	// Identifies the resource.
	ID string `json:"id"`
	// Identifies the billing group associated with the phone number.
	BillingGroupID string `json:"billing_group_id"`
	// Indicates if call forwarding will be enabled for this number if forwards_to and
	// forwarding_type are filled in. Defaults to true for backwards compatibility with
	// APIV1 use of numbers endpoints.
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Indicates whether call recording is enabled for this number.
	CallRecordingEnabled bool `json:"call_recording_enabled"`
	// Indicates whether caller ID is enabled for this number.
	CallerIDNameEnabled bool `json:"caller_id_name_enabled"`
	// Indicates whether a CNAM listing is enabled for this number.
	CnamListingEnabled bool `json:"cnam_listing_enabled"`
	// Identifies the connection associated with the phone number.
	ConnectionID string `json:"connection_id"`
	// The ISO 3166-1 alpha-2 country code of the phone number.
	CountryISOAlpha2 string `json:"country_iso_alpha2"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// A customer reference string for customer look ups.
	CustomerReference string `json:"customer_reference"`
	// Identifies the emergency address associated with the phone number.
	EmergencyAddressID string `json:"emergency_address_id"`
	// Indicates whether emergency services are enabled for this number.
	EmergencyEnabled bool `json:"emergency_enabled"`
	// Indicates the status of the provisioning of emergency services for the phone
	// number. This field contains information about activity that may be ongoing for a
	// number where it either is being provisioned or deprovisioned but is not yet
	// enabled/disabled.
	//
	// Any of "active", "deprovisioning", "disabled", "provisioning",
	// "provisioning-failed".
	EmergencyStatus string `json:"emergency_status"`
	// If someone attempts to port your phone number away from Telnyx and your phone
	// number has an external PIN set, Telnyx will attempt to verify that you provided
	// the correct external PIN to the winning carrier. Note that not all carriers
	// cooperate with this security mechanism.
	ExternalPin string `json:"external_pin"`
	// The inbound_call_screening setting is a phone number configuration option
	// variable that allows users to configure their settings to block or flag
	// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
	// feature has an additional per-number monthly cost associated with it.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening string `json:"inbound_call_screening"`
	// The +E.164-formatted phone number associated with this record.
	PhoneNumber string `json:"phone_number"`
	// The phone number's type. Note: For numbers purchased prior to July 2023 or when
	// fetching a number's details immediately after a purchase completes, the legacy
	// values `tollfree`, `shortcode` or `longcode` may be returned instead.
	//
	// Any of "local", "toll_free", "mobile", "national", "shared_cost", "landline",
	// "tollfree", "shortcode", "longcode".
	PhoneNumberType string `json:"phone_number_type"`
	// ISO 8601 formatted date indicating when the resource was purchased.
	PurchasedAt string `json:"purchased_at"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// The phone number's current status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "port-failed",
	// "active", "deleted", "emergency-only", "ported-out", "port-out-pending",
	// "requirement-info-pending", "requirement-info-under-review",
	// "requirement-info-exception", "provision-pending".
	Status string `json:"status"`
	// Indicates whether T38 Fax Gateway for inbound calls to this number.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BillingGroupID        respjson.Field
		CallForwardingEnabled respjson.Field
		CallRecordingEnabled  respjson.Field
		CallerIDNameEnabled   respjson.Field
		CnamListingEnabled    respjson.Field
		ConnectionID          respjson.Field
		CountryISOAlpha2      respjson.Field
		CreatedAt             respjson.Field
		CustomerReference     respjson.Field
		EmergencyAddressID    respjson.Field
		EmergencyEnabled      respjson.Field
		EmergencyStatus       respjson.Field
		ExternalPin           respjson.Field
		InboundCallScreening  respjson.Field
		PhoneNumber           respjson.Field
		PhoneNumberType       respjson.Field
		PurchasedAt           respjson.Field
		RecordType            respjson.Field
		Status                respjson.Field
		T38FaxGatewayEnabled  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberSlimListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberSlimListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberUpdateParams struct {
	// Identifies the billing group associated with the phone number.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// Identifies the connection associated with the phone number.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// If someone attempts to port your phone number away from Telnyx and your phone
	// number has an external PIN set, we will attempt to verify that you provided the
	// correct external PIN to the winning carrier. Note that not all carriers
	// cooperate with this security mechanism.
	ExternalPin param.Opt[string] `json:"external_pin,omitzero"`
	// Indicates whether HD voice is enabled for this number.
	HDVoiceEnabled param.Opt[bool] `json:"hd_voice_enabled,omitzero"`
	// A list of user-assigned tags to help organize phone numbers.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r PhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[tag],
	// filter[phone_number], filter[status], filter[country_iso_alpha2],
	// filter[connection_id], filter[voice.connection_name],
	// filter[voice.usage_payment_method], filter[billing_group_id],
	// filter[emergency_address_id], filter[customer_reference], filter[number_type],
	// filter[source]
	Filter PhoneNumberListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PhoneNumberListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "purchased_at", "phone_number", "connection_name",
	// "usage_payment_method".
	Sort PhoneNumberListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParams]'s query parameters as `url.Values`.
func (r PhoneNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[tag],
// filter[phone_number], filter[status], filter[country_iso_alpha2],
// filter[connection_id], filter[voice.connection_name],
// filter[voice.usage_payment_method], filter[billing_group_id],
// filter[emergency_address_id], filter[customer_reference], filter[number_type],
// filter[source]
type PhoneNumberListParamsFilter struct {
	// Filter by the billing_group_id associated with phone numbers. To filter to only
	// phone numbers that have no billing group associated them, set the value of this
	// filter to the string 'null'.
	BillingGroupID param.Opt[string] `query:"billing_group_id,omitzero" json:"-"`
	// Filter by connection_id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Filter numbers via the customer_reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter by the emergency_address_id associated with phone numbers. To filter only
	// phone numbers that have no emergency address associated with them, set the value
	// of this filter to the string 'null'.
	EmergencyAddressID param.Opt[string] `query:"emergency_address_id,omitzero" json:"-"`
	// Filter by phone number. Requires at least three digits. Non-numerical characters
	// will result in no values being returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by phone number tags.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Filter by phone number country ISO alpha-2 code. Can be a single value or an
	// array of values.
	CountryISOAlpha2 PhoneNumberListParamsFilterCountryISOAlpha2Union `query:"country_iso_alpha2,omitzero" json:"-"`
	// Filter phone numbers by phone number type.
	NumberType PhoneNumberListParamsFilterNumberType `query:"number_type,omitzero" json:"-"`
	// Filter phone numbers by their source. Use 'ported' for numbers ported from other
	// carriers, or 'purchased' for numbers bought directly from Telnyx.
	//
	// Any of "ported", "purchased".
	Source string `query:"source,omitzero" json:"-"`
	// Filter by phone number status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "active",
	// "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by voice connection name pattern matching.
	VoiceConnectionName PhoneNumberListParamsFilterVoiceConnectionName `query:"voice.connection_name,omitzero" json:"-"`
	// Filter by usage_payment_method.
	//
	// Any of "pay-per-minute", "channel".
	VoiceUsagePaymentMethod string `query:"voice.usage_payment_method,omitzero" json:"-"`
	// When set to 'true', filters for phone numbers that do not have any tags applied.
	// All other values are ignored.
	//
	// Any of "true", "false".
	WithoutTags string `query:"without_tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberListParamsFilter](
		"source", "ported", "purchased",
	)
	apijson.RegisterFieldValidator[PhoneNumberListParamsFilter](
		"status", "purchase-pending", "purchase-failed", "port-pending", "active", "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending",
	)
	apijson.RegisterFieldValidator[PhoneNumberListParamsFilter](
		"voice.usage_payment_method", "pay-per-minute", "channel",
	)
	apijson.RegisterFieldValidator[PhoneNumberListParamsFilter](
		"without_tags", "true", "false",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PhoneNumberListParamsFilterCountryISOAlpha2Union struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *PhoneNumberListParamsFilterCountryISOAlpha2Union) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Filter phone numbers by phone number type.
type PhoneNumberListParamsFilterNumberType struct {
	// Filter phone numbers by phone number type.
	//
	// Any of "local", "national", "toll_free", "mobile", "shared_cost".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParamsFilterNumberType]'s query parameters
// as `url.Values`.
func (r PhoneNumberListParamsFilterNumberType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberListParamsFilterNumberType](
		"eq", "local", "national", "toll_free", "mobile", "shared_cost",
	)
}

// Filter by voice connection name pattern matching.
type PhoneNumberListParamsFilterVoiceConnectionName struct {
	// Filter contains connection name. Requires at least three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter ends with connection name. Requires at least three characters.
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter by connection name.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter starts with connection name. Requires at least three characters.
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParamsFilterVoiceConnectionName]'s query
// parameters as `url.Values`.
func (r PhoneNumberListParamsFilterVoiceConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PhoneNumberListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberListParamsPage]'s query parameters as
// `url.Values`.
func (r PhoneNumberListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberListParamsSort string

const (
	PhoneNumberListParamsSortPurchasedAt        PhoneNumberListParamsSort = "purchased_at"
	PhoneNumberListParamsSortPhoneNumber        PhoneNumberListParamsSort = "phone_number"
	PhoneNumberListParamsSortConnectionName     PhoneNumberListParamsSort = "connection_name"
	PhoneNumberListParamsSortUsagePaymentMethod PhoneNumberListParamsSort = "usage_payment_method"
)

type PhoneNumberSlimListParams struct {
	// Include the connection associated with the phone number.
	IncludeConnection param.Opt[bool] `query:"include_connection,omitzero" json:"-"`
	// Include the tags associated with the phone number.
	IncludeTags param.Opt[bool] `query:"include_tags,omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[tag],
	// filter[phone_number], filter[status], filter[country_iso_alpha2],
	// filter[connection_id], filter[voice.connection_name],
	// filter[voice.usage_payment_method], filter[billing_group_id],
	// filter[emergency_address_id], filter[customer_reference], filter[number_type],
	// filter[source]
	Filter PhoneNumberSlimListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PhoneNumberSlimListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "purchased_at", "phone_number", "connection_name",
	// "usage_payment_method".
	Sort PhoneNumberSlimListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSlimListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberSlimListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[tag],
// filter[phone_number], filter[status], filter[country_iso_alpha2],
// filter[connection_id], filter[voice.connection_name],
// filter[voice.usage_payment_method], filter[billing_group_id],
// filter[emergency_address_id], filter[customer_reference], filter[number_type],
// filter[source]
type PhoneNumberSlimListParamsFilter struct {
	// Filter by the billing_group_id associated with phone numbers. To filter to only
	// phone numbers that have no billing group associated them, set the value of this
	// filter to the string 'null'.
	BillingGroupID param.Opt[string] `query:"billing_group_id,omitzero" json:"-"`
	// Filter by connection_id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Filter numbers via the customer_reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter by the emergency_address_id associated with phone numbers. To filter only
	// phone numbers that have no emergency address associated with them, set the value
	// of this filter to the string 'null'.
	EmergencyAddressID param.Opt[string] `query:"emergency_address_id,omitzero" json:"-"`
	// Filter by phone number. Requires at least three digits. Non-numerical characters
	// will result in no values being returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by phone number tags. (This requires the include_tags param)
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Filter by phone number country ISO alpha-2 code. Can be a single value or an
	// array of values.
	CountryISOAlpha2 PhoneNumberSlimListParamsFilterCountryISOAlpha2Union `query:"country_iso_alpha2,omitzero" json:"-"`
	// Filter phone numbers by phone number type.
	NumberType PhoneNumberSlimListParamsFilterNumberType `query:"number_type,omitzero" json:"-"`
	// Filter phone numbers by their source. Use 'ported' for numbers ported from other
	// carriers, or 'purchased' for numbers bought directly from Telnyx.
	//
	// Any of "ported", "purchased".
	Source string `query:"source,omitzero" json:"-"`
	// Filter by phone number status.
	//
	// Any of "purchase-pending", "purchase-failed", "port_pending", "active",
	// "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by voice connection name pattern matching (requires include_connection
	// param).
	VoiceConnectionName PhoneNumberSlimListParamsFilterVoiceConnectionName `query:"voice.connection_name,omitzero" json:"-"`
	// Filter by usage_payment_method.
	//
	// Any of "pay-per-minute", "channel".
	VoiceUsagePaymentMethod string `query:"voice.usage_payment_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSlimListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberSlimListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberSlimListParamsFilter](
		"source", "ported", "purchased",
	)
	apijson.RegisterFieldValidator[PhoneNumberSlimListParamsFilter](
		"status", "purchase-pending", "purchase-failed", "port_pending", "active", "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending",
	)
	apijson.RegisterFieldValidator[PhoneNumberSlimListParamsFilter](
		"voice.usage_payment_method", "pay-per-minute", "channel",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PhoneNumberSlimListParamsFilterCountryISOAlpha2Union struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *PhoneNumberSlimListParamsFilterCountryISOAlpha2Union) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Filter phone numbers by phone number type.
type PhoneNumberSlimListParamsFilterNumberType struct {
	// Filter phone numbers by phone number type.
	//
	// Any of "local", "national", "toll_free", "mobile", "shared_cost".
	Eq string `query:"eq,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSlimListParamsFilterNumberType]'s query
// parameters as `url.Values`.
func (r PhoneNumberSlimListParamsFilterNumberType) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberSlimListParamsFilterNumberType](
		"eq", "local", "national", "toll_free", "mobile", "shared_cost",
	)
}

// Filter by voice connection name pattern matching (requires include_connection
// param).
type PhoneNumberSlimListParamsFilterVoiceConnectionName struct {
	// Filter contains connection name. Requires at least three characters and the
	// include_connection param.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter ends with connection name. Requires at least three characters and the
	// include_connection param.
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter by connection name.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter starts with connection name. Requires at least three characters and the
	// include_connection param.
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSlimListParamsFilterVoiceConnectionName]'s query
// parameters as `url.Values`.
func (r PhoneNumberSlimListParamsFilterVoiceConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PhoneNumberSlimListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberSlimListParamsPage]'s query parameters as
// `url.Values`.
func (r PhoneNumberSlimListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberSlimListParamsSort string

const (
	PhoneNumberSlimListParamsSortPurchasedAt        PhoneNumberSlimListParamsSort = "purchased_at"
	PhoneNumberSlimListParamsSortPhoneNumber        PhoneNumberSlimListParamsSort = "phone_number"
	PhoneNumberSlimListParamsSortConnectionName     PhoneNumberSlimListParamsSort = "connection_name"
	PhoneNumberSlimListParamsSortUsagePaymentMethod PhoneNumberSlimListParamsSort = "usage_payment_method"
)
