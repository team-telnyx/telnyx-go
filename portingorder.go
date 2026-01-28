// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// PortingOrderService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderService] method instead.
type PortingOrderService struct {
	Options                   []option.RequestOption
	PhoneNumberConfigurations PortingOrderPhoneNumberConfigurationService
	Actions                   PortingOrderActionService
	ActivationJobs            PortingOrderActivationJobService
	AdditionalDocuments       PortingOrderAdditionalDocumentService
	Comments                  PortingOrderCommentService
	VerificationCodes         PortingOrderVerificationCodeService
	ActionRequirements        PortingOrderActionRequirementService
	AssociatedPhoneNumbers    PortingOrderAssociatedPhoneNumberService
	PhoneNumberBlocks         PortingOrderPhoneNumberBlockService
	PhoneNumberExtensions     PortingOrderPhoneNumberExtensionService
}

// NewPortingOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortingOrderService(opts ...option.RequestOption) (r PortingOrderService) {
	r = PortingOrderService{}
	r.Options = opts
	r.PhoneNumberConfigurations = NewPortingOrderPhoneNumberConfigurationService(opts...)
	r.Actions = NewPortingOrderActionService(opts...)
	r.ActivationJobs = NewPortingOrderActivationJobService(opts...)
	r.AdditionalDocuments = NewPortingOrderAdditionalDocumentService(opts...)
	r.Comments = NewPortingOrderCommentService(opts...)
	r.VerificationCodes = NewPortingOrderVerificationCodeService(opts...)
	r.ActionRequirements = NewPortingOrderActionRequirementService(opts...)
	r.AssociatedPhoneNumbers = NewPortingOrderAssociatedPhoneNumberService(opts...)
	r.PhoneNumberBlocks = NewPortingOrderPhoneNumberBlockService(opts...)
	r.PhoneNumberExtensions = NewPortingOrderPhoneNumberExtensionService(opts...)
	return
}

// Creates a new porting order object.
func (r *PortingOrderService) New(ctx context.Context, body PortingOrderNewParams, opts ...option.RequestOption) (res *PortingOrderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing porting order.
func (r *PortingOrderService) Get(ctx context.Context, id string, query PortingOrderGetParams, opts ...option.RequestOption) (res *PortingOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edits the details of an existing porting order.
//
// Any or all of a porting orders attributes may be included in the resource object
// included in a PATCH request.
//
// If a request does not include all of the attributes for a resource, the system
// will interpret the missing attributes as if they were included with their
// current values. To explicitly set something to null, it must be included in the
// request with a null value.
func (r *PortingOrderService) Update(ctx context.Context, id string, body PortingOrderUpdateParams, opts ...option.RequestOption) (res *PortingOrderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Returns a list of your porting order.
func (r *PortingOrderService) List(ctx context.Context, query PortingOrderListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortingOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "porting_orders"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of your porting order.
func (r *PortingOrderService) ListAutoPaging(ctx context.Context, query PortingOrderListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortingOrder] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing porting order. This operation is restrict to porting orders
// in draft state.
func (r *PortingOrderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns a list of allowed FOC dates for a porting order.
func (r *PortingOrderService) GetAllowedFocWindows(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingOrderGetAllowedFocWindowsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/allowed_foc_windows", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all possible exception types for a porting order.
func (r *PortingOrderService) GetExceptionTypes(ctx context.Context, opts ...option.RequestOption) (res *PortingOrderGetExceptionTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting_orders/exception_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Download a porting order loa template
func (r *PortingOrderService) GetLoaTemplate(ctx context.Context, id string, query PortingOrderGetLoaTemplateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/loa_template", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of all requirements based on country/number type for this porting
// order.
func (r *PortingOrderService) GetRequirements(ctx context.Context, id string, query PortingOrderGetRequirementsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortingOrderGetRequirementsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/requirements", id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of all requirements based on country/number type for this porting
// order.
func (r *PortingOrderService) GetRequirementsAutoPaging(ctx context.Context, id string, query PortingOrderGetRequirementsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortingOrderGetRequirementsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.GetRequirements(ctx, id, query, opts...))
}

// Retrieve the associated V1 sub_request_id and port_request_id
func (r *PortingOrderService) GetSubRequest(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingOrderGetSubRequestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/sub_request", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PortingOrder struct {
	// Uniquely identifies this porting order
	ID                 string                         `json:"id" format:"uuid"`
	ActivationSettings PortingOrderActivationSettings `json:"activation_settings"`
	// For specific porting orders, we may require additional steps to be taken before
	// submitting the porting order.
	//
	// Any of "associated_phone_numbers", "phone_number_verification_codes".
	AdditionalSteps []string `json:"additional_steps"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A customer-specified group reference for customer bookkeeping purposes
	CustomerGroupReference string `json:"customer_group_reference,nullable"`
	// A customer-specified reference number for customer bookkeeping purposes
	CustomerReference string `json:"customer_reference,nullable"`
	// A description of the porting order
	Description string `json:"description"`
	// Can be specified directly or via the `requirement_group_id` parameter.
	Documents PortingOrderDocuments `json:"documents"`
	EndUser   PortingOrderEndUser   `json:"end_user"`
	// Information about messaging porting process.
	Messaging PortingOrderMessaging `json:"messaging"`
	Misc      PortingOrderMisc      `json:"misc,nullable"`
	// Identifies the old service provider
	OldServiceProviderOcn string `json:"old_service_provider_ocn"`
	// A key to reference for the porting order group when contacting Telnyx customer
	// support. This information is not available for porting orders in `draft` state
	ParentSupportKey         string                               `json:"parent_support_key,nullable"`
	PhoneNumberConfiguration PortingOrderPhoneNumberConfiguration `json:"phone_number_configuration"`
	// The type of the phone number
	//
	// Any of "landline", "local", "mobile", "national", "shared_cost", "toll_free".
	PhoneNumberType PortingOrderPhoneNumberType `json:"phone_number_type"`
	// Count of phone numbers associated with this porting order
	PortingPhoneNumbersCount int64 `json:"porting_phone_numbers_count"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// List of documentation requirements for porting numbers. Can be set directly or
	// via the `requirement_group_id` parameter.
	Requirements []PortingOrderRequirement `json:"requirements"`
	// Is true when the required documentation is met
	RequirementsMet bool `json:"requirements_met"`
	// Porting order status
	Status shared.PortingOrderStatus `json:"status"`
	// A key to reference this porting order when contacting Telnyx customer support.
	// This information is not available in draft porting orders.
	SupportKey string `json:"support_key,nullable"`
	// ISO 8601 formatted date indicating when the resource was created.
	UpdatedAt    time.Time                `json:"updated_at" format:"date-time"`
	UserFeedback PortingOrderUserFeedback `json:"user_feedback"`
	// Identifies the user (or organization) who requested the porting order
	UserID     string `json:"user_id" format:"uuid"`
	WebhookURL string `json:"webhook_url,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ActivationSettings       respjson.Field
		AdditionalSteps          respjson.Field
		CreatedAt                respjson.Field
		CustomerGroupReference   respjson.Field
		CustomerReference        respjson.Field
		Description              respjson.Field
		Documents                respjson.Field
		EndUser                  respjson.Field
		Messaging                respjson.Field
		Misc                     respjson.Field
		OldServiceProviderOcn    respjson.Field
		ParentSupportKey         respjson.Field
		PhoneNumberConfiguration respjson.Field
		PhoneNumberType          respjson.Field
		PortingPhoneNumbersCount respjson.Field
		RecordType               respjson.Field
		Requirements             respjson.Field
		RequirementsMet          respjson.Field
		Status                   respjson.Field
		SupportKey               respjson.Field
		UpdatedAt                respjson.Field
		UserFeedback             respjson.Field
		UserID                   respjson.Field
		WebhookURL               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrder) RawJSON() string { return r.JSON.raw }
func (r *PortingOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the phone number
type PortingOrderPhoneNumberType string

const (
	PortingOrderPhoneNumberTypeLandline   PortingOrderPhoneNumberType = "landline"
	PortingOrderPhoneNumberTypeLocal      PortingOrderPhoneNumberType = "local"
	PortingOrderPhoneNumberTypeMobile     PortingOrderPhoneNumberType = "mobile"
	PortingOrderPhoneNumberTypeNational   PortingOrderPhoneNumberType = "national"
	PortingOrderPhoneNumberTypeSharedCost PortingOrderPhoneNumberType = "shared_cost"
	PortingOrderPhoneNumberTypeTollFree   PortingOrderPhoneNumberType = "toll_free"
)

type PortingOrderActivationSettings struct {
	// Activation status
	//
	// Any of "New", "Pending", "Conflict", "Cancel Pending", "Failed", "Concurred",
	// "Activate RDY", "Disconnect Pending", "Concurrence Sent", "Old", "Sending",
	// "Active", "Cancelled".
	ActivationStatus PortingOrderActivationSettingsActivationStatus `json:"activation_status,nullable"`
	// Indicates whether this porting order is eligible for FastPort
	FastPortEligible bool `json:"fast_port_eligible"`
	// ISO 8601 formatted Date/Time of the FOC date
	FocDatetimeActual time.Time `json:"foc_datetime_actual,nullable" format:"date-time"`
	// ISO 8601 formatted Date/Time requested for the FOC date
	FocDatetimeRequested time.Time `json:"foc_datetime_requested,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivationStatus     respjson.Field
		FastPortEligible     respjson.Field
		FocDatetimeActual    respjson.Field
		FocDatetimeRequested respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActivationSettings) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActivationSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activation status
type PortingOrderActivationSettingsActivationStatus string

const (
	PortingOrderActivationSettingsActivationStatusNew               PortingOrderActivationSettingsActivationStatus = "New"
	PortingOrderActivationSettingsActivationStatusPending           PortingOrderActivationSettingsActivationStatus = "Pending"
	PortingOrderActivationSettingsActivationStatusConflict          PortingOrderActivationSettingsActivationStatus = "Conflict"
	PortingOrderActivationSettingsActivationStatusCancelPending     PortingOrderActivationSettingsActivationStatus = "Cancel Pending"
	PortingOrderActivationSettingsActivationStatusFailed            PortingOrderActivationSettingsActivationStatus = "Failed"
	PortingOrderActivationSettingsActivationStatusConcurred         PortingOrderActivationSettingsActivationStatus = "Concurred"
	PortingOrderActivationSettingsActivationStatusActivateRdy       PortingOrderActivationSettingsActivationStatus = "Activate RDY"
	PortingOrderActivationSettingsActivationStatusDisconnectPending PortingOrderActivationSettingsActivationStatus = "Disconnect Pending"
	PortingOrderActivationSettingsActivationStatusConcurrenceSent   PortingOrderActivationSettingsActivationStatus = "Concurrence Sent"
	PortingOrderActivationSettingsActivationStatusOld               PortingOrderActivationSettingsActivationStatus = "Old"
	PortingOrderActivationSettingsActivationStatusSending           PortingOrderActivationSettingsActivationStatus = "Sending"
	PortingOrderActivationSettingsActivationStatusActive            PortingOrderActivationSettingsActivationStatus = "Active"
	PortingOrderActivationSettingsActivationStatusCancelled         PortingOrderActivationSettingsActivationStatus = "Cancelled"
)

// Can be specified directly or via the `requirement_group_id` parameter.
type PortingOrderDocuments struct {
	// Returned ID of the submitted Invoice via the Documents endpoint
	Invoice string `json:"invoice,nullable" format:"uuid"`
	// Returned ID of the submitted LOA via the Documents endpoint
	Loa string `json:"loa,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoice     respjson.Field
		Loa         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderDocuments) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderDocuments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderDocuments to a PortingOrderDocumentsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderDocumentsParam.Overrides()
func (r PortingOrderDocuments) ToParam() PortingOrderDocumentsParam {
	return param.Override[PortingOrderDocumentsParam](json.RawMessage(r.RawJSON()))
}

// Can be specified directly or via the `requirement_group_id` parameter.
type PortingOrderDocumentsParam struct {
	// Returned ID of the submitted Invoice via the Documents endpoint
	Invoice param.Opt[string] `json:"invoice,omitzero" format:"uuid"`
	// Returned ID of the submitted LOA via the Documents endpoint
	Loa param.Opt[string] `json:"loa,omitzero" format:"uuid"`
	paramObj
}

func (r PortingOrderDocumentsParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderDocumentsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderDocumentsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderEndUser struct {
	Admin    PortingOrderEndUserAdmin    `json:"admin"`
	Location PortingOrderEndUserLocation `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Admin       respjson.Field
		Location    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderEndUser) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderEndUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderEndUser to a PortingOrderEndUserParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderEndUserParam.Overrides()
func (r PortingOrderEndUser) ToParam() PortingOrderEndUserParam {
	return param.Override[PortingOrderEndUserParam](json.RawMessage(r.RawJSON()))
}

type PortingOrderEndUserParam struct {
	Admin    PortingOrderEndUserAdminParam    `json:"admin,omitzero"`
	Location PortingOrderEndUserLocationParam `json:"location,omitzero"`
	paramObj
}

func (r PortingOrderEndUserParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderEndUserParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderEndUserParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderEndUserAdmin struct {
	// The authorized person's account number with the current service provider
	AccountNumber string `json:"account_number,nullable"`
	// Name of person authorizing the porting order
	AuthPersonName string `json:"auth_person_name,nullable"`
	// Billing phone number associated with these phone numbers
	BillingPhoneNumber string `json:"billing_phone_number,nullable"`
	// European business identification number. Applicable only in the European Union
	BusinessIdentifier string `json:"business_identifier,nullable"`
	// Person Name or Company name requesting the port
	EntityName string `json:"entity_name,nullable"`
	// PIN/passcode possibly required by the old service provider for extra
	// verification
	PinPasscode string `json:"pin_passcode,nullable"`
	// European tax identification number. Applicable only in the European Union
	TaxIdentifier string `json:"tax_identifier,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber      respjson.Field
		AuthPersonName     respjson.Field
		BillingPhoneNumber respjson.Field
		BusinessIdentifier respjson.Field
		EntityName         respjson.Field
		PinPasscode        respjson.Field
		TaxIdentifier      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderEndUserAdmin) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderEndUserAdmin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderEndUserAdmin to a
// PortingOrderEndUserAdminParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderEndUserAdminParam.Overrides()
func (r PortingOrderEndUserAdmin) ToParam() PortingOrderEndUserAdminParam {
	return param.Override[PortingOrderEndUserAdminParam](json.RawMessage(r.RawJSON()))
}

type PortingOrderEndUserAdminParam struct {
	// The authorized person's account number with the current service provider
	AccountNumber param.Opt[string] `json:"account_number,omitzero"`
	// Name of person authorizing the porting order
	AuthPersonName param.Opt[string] `json:"auth_person_name,omitzero"`
	// Billing phone number associated with these phone numbers
	BillingPhoneNumber param.Opt[string] `json:"billing_phone_number,omitzero"`
	// European business identification number. Applicable only in the European Union
	BusinessIdentifier param.Opt[string] `json:"business_identifier,omitzero"`
	// Person Name or Company name requesting the port
	EntityName param.Opt[string] `json:"entity_name,omitzero"`
	// PIN/passcode possibly required by the old service provider for extra
	// verification
	PinPasscode param.Opt[string] `json:"pin_passcode,omitzero"`
	// European tax identification number. Applicable only in the European Union
	TaxIdentifier param.Opt[string] `json:"tax_identifier,omitzero"`
	paramObj
}

func (r PortingOrderEndUserAdminParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderEndUserAdminParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderEndUserAdminParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderEndUserLocation struct {
	// State, province, or similar of billing address
	AdministrativeArea string `json:"administrative_area,nullable"`
	// ISO3166-1 alpha-2 country code of billing address
	CountryCode string `json:"country_code,nullable"`
	// Second line of billing address
	ExtendedAddress string `json:"extended_address,nullable"`
	// City or municipality of billing address
	Locality string `json:"locality,nullable"`
	// Postal Code of billing address
	PostalCode string `json:"postal_code,nullable"`
	// First line of billing address
	StreetAddress string `json:"street_address,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		CountryCode        respjson.Field
		ExtendedAddress    respjson.Field
		Locality           respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderEndUserLocation) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderEndUserLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderEndUserLocation to a
// PortingOrderEndUserLocationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderEndUserLocationParam.Overrides()
func (r PortingOrderEndUserLocation) ToParam() PortingOrderEndUserLocationParam {
	return param.Override[PortingOrderEndUserLocationParam](json.RawMessage(r.RawJSON()))
}

type PortingOrderEndUserLocationParam struct {
	// State, province, or similar of billing address
	AdministrativeArea param.Opt[string] `json:"administrative_area,omitzero"`
	// ISO3166-1 alpha-2 country code of billing address
	CountryCode param.Opt[string] `json:"country_code,omitzero"`
	// Second line of billing address
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	// City or municipality of billing address
	Locality param.Opt[string] `json:"locality,omitzero"`
	// Postal Code of billing address
	PostalCode param.Opt[string] `json:"postal_code,omitzero"`
	// First line of billing address
	StreetAddress param.Opt[string] `json:"street_address,omitzero"`
	paramObj
}

func (r PortingOrderEndUserLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderEndUserLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderEndUserLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about messaging porting process.
type PortingOrderMessaging struct {
	// Indicates whether Telnyx will port messaging capabilities from the losing
	// carrier. If false, any messaging capabilities will stay with their current
	// provider.
	EnableMessaging bool `json:"enable_messaging"`
	// Indicates whether the porting order can also port messaging capabilities.
	MessagingCapable bool `json:"messaging_capable"`
	// Indicates whether the messaging porting has been completed.
	MessagingPortCompleted bool `json:"messaging_port_completed"`
	// The current status of the messaging porting.
	//
	// Any of "not_applicable", "pending", "activating", "exception", "canceled",
	// "partial_port_complete", "ported".
	MessagingPortStatus PortingOrderMessagingMessagingPortStatus `json:"messaging_port_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableMessaging        respjson.Field
		MessagingCapable       respjson.Field
		MessagingPortCompleted respjson.Field
		MessagingPortStatus    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderMessaging) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the messaging porting.
type PortingOrderMessagingMessagingPortStatus string

const (
	PortingOrderMessagingMessagingPortStatusNotApplicable       PortingOrderMessagingMessagingPortStatus = "not_applicable"
	PortingOrderMessagingMessagingPortStatusPending             PortingOrderMessagingMessagingPortStatus = "pending"
	PortingOrderMessagingMessagingPortStatusActivating          PortingOrderMessagingMessagingPortStatus = "activating"
	PortingOrderMessagingMessagingPortStatusException           PortingOrderMessagingMessagingPortStatus = "exception"
	PortingOrderMessagingMessagingPortStatusCanceled            PortingOrderMessagingMessagingPortStatus = "canceled"
	PortingOrderMessagingMessagingPortStatusPartialPortComplete PortingOrderMessagingMessagingPortStatus = "partial_port_complete"
	PortingOrderMessagingMessagingPortStatusPorted              PortingOrderMessagingMessagingPortStatus = "ported"
)

type PortingOrderMisc struct {
	// New billing phone number for the remaining numbers. Used in case the current
	// billing phone number is being ported to Telnyx. This will be set on your account
	// with your current service provider and should be one of the numbers remaining on
	// that account.
	NewBillingPhoneNumber string `json:"new_billing_phone_number,nullable"`
	// Remaining numbers can be either kept with their current service provider or
	// disconnected. 'new_billing_telephone_number' is required when
	// 'remaining_numbers_action' is 'keep'.
	//
	// Any of "keep", "disconnect".
	RemainingNumbersAction PortingOrderMiscRemainingNumbersAction `json:"remaining_numbers_action,nullable"`
	// A port can be either 'full' or 'partial'. When type is 'full' the other
	// attributes should be omitted.
	//
	// Any of "full", "partial".
	Type PortingOrderType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewBillingPhoneNumber  respjson.Field
		RemainingNumbersAction respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderMisc) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderMisc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderMisc to a PortingOrderMiscParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderMiscParam.Overrides()
func (r PortingOrderMisc) ToParam() PortingOrderMiscParam {
	return param.Override[PortingOrderMiscParam](json.RawMessage(r.RawJSON()))
}

// Remaining numbers can be either kept with their current service provider or
// disconnected. 'new_billing_telephone_number' is required when
// 'remaining_numbers_action' is 'keep'.
type PortingOrderMiscRemainingNumbersAction string

const (
	PortingOrderMiscRemainingNumbersActionKeep       PortingOrderMiscRemainingNumbersAction = "keep"
	PortingOrderMiscRemainingNumbersActionDisconnect PortingOrderMiscRemainingNumbersAction = "disconnect"
)

type PortingOrderMiscParam struct {
	// New billing phone number for the remaining numbers. Used in case the current
	// billing phone number is being ported to Telnyx. This will be set on your account
	// with your current service provider and should be one of the numbers remaining on
	// that account.
	NewBillingPhoneNumber param.Opt[string] `json:"new_billing_phone_number,omitzero"`
	// Remaining numbers can be either kept with their current service provider or
	// disconnected. 'new_billing_telephone_number' is required when
	// 'remaining_numbers_action' is 'keep'.
	//
	// Any of "keep", "disconnect".
	RemainingNumbersAction PortingOrderMiscRemainingNumbersAction `json:"remaining_numbers_action,omitzero"`
	// A port can be either 'full' or 'partial'. When type is 'full' the other
	// attributes should be omitted.
	//
	// Any of "full", "partial".
	Type PortingOrderType `json:"type,omitzero"`
	paramObj
}

func (r PortingOrderMiscParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderMiscParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderMiscParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderPhoneNumberConfiguration struct {
	// identifies the billing group to set on the numbers when ported
	BillingGroupID string `json:"billing_group_id,nullable"`
	// identifies the connection to set on the numbers when ported
	ConnectionID string `json:"connection_id,nullable"`
	// identifies the emergency address to set on the numbers when ported
	EmergencyAddressID string `json:"emergency_address_id,nullable"`
	// identifies the messaging profile to set on the numbers when ported
	MessagingProfileID string   `json:"messaging_profile_id,nullable"`
	Tags               []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroupID     respjson.Field
		ConnectionID       respjson.Field
		EmergencyAddressID respjson.Field
		MessagingProfileID respjson.Field
		Tags               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderPhoneNumberConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderPhoneNumberConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderPhoneNumberConfiguration to a
// PortingOrderPhoneNumberConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderPhoneNumberConfigurationParam.Overrides()
func (r PortingOrderPhoneNumberConfiguration) ToParam() PortingOrderPhoneNumberConfigurationParam {
	return param.Override[PortingOrderPhoneNumberConfigurationParam](json.RawMessage(r.RawJSON()))
}

type PortingOrderPhoneNumberConfigurationParam struct {
	// identifies the billing group to set on the numbers when ported
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// identifies the connection to set on the numbers when ported
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// identifies the emergency address to set on the numbers when ported
	EmergencyAddressID param.Opt[string] `json:"emergency_address_id,omitzero"`
	// identifies the messaging profile to set on the numbers when ported
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	Tags               []string          `json:"tags,omitzero"`
	paramObj
}

func (r PortingOrderPhoneNumberConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderPhoneNumberConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderPhoneNumberConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderRequirement struct {
	// Type of value expected on field_value field
	//
	// Any of "document".
	FieldType PortingOrderRequirementFieldType `json:"field_type"`
	// identifies the document that satisfies this requirement
	FieldValue string `json:"field_value"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Identifies the requirement type that meets this requirement
	RequirementTypeID string `json:"requirement_type_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType         respjson.Field
		FieldValue        respjson.Field
		RecordType        respjson.Field
		RequirementTypeID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderRequirement) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of value expected on field_value field
type PortingOrderRequirementFieldType string

const (
	PortingOrderRequirementFieldTypeDocument PortingOrderRequirementFieldType = "document"
)

// A port can be either 'full' or 'partial'. When type is 'full' the other
// attributes should be omitted.
type PortingOrderType string

const (
	PortingOrderTypeFull    PortingOrderType = "full"
	PortingOrderTypePartial PortingOrderType = "partial"
)

type PortingOrderUserFeedback struct {
	// A comment related to the customer rating.
	UserComment string `json:"user_comment,nullable"`
	// Once an order is ported, cancellation is requested or the request is cancelled,
	// the user may rate their experience
	UserRating int64 `json:"user_rating,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserComment respjson.Field
		UserRating  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderUserFeedback) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderUserFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PortingOrderUserFeedback to a
// PortingOrderUserFeedbackParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortingOrderUserFeedbackParam.Overrides()
func (r PortingOrderUserFeedback) ToParam() PortingOrderUserFeedbackParam {
	return param.Override[PortingOrderUserFeedbackParam](json.RawMessage(r.RawJSON()))
}

type PortingOrderUserFeedbackParam struct {
	// A comment related to the customer rating.
	UserComment param.Opt[string] `json:"user_comment,omitzero"`
	// Once an order is ported, cancellation is requested or the request is cancelled,
	// the user may rate their experience
	UserRating param.Opt[int64] `json:"user_rating,omitzero"`
	paramObj
}

func (r PortingOrderUserFeedbackParam) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderUserFeedbackParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderUserFeedbackParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrdersActivationJob struct {
	// Uniquely identifies this activation job
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the activation job should be executed.
	// This time should be between some activation window.
	ActivateAt time.Time `json:"activate_at" format:"date-time"`
	// Specifies the type of this activation job
	//
	// Any of "scheduled", "on-demand".
	ActivationType PortingOrdersActivationJobActivationType `json:"activation_type"`
	// List of allowed activation windows for this activation job
	ActivationWindows []PortingOrdersActivationJobActivationWindow `json:"activation_windows"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Specifies the status of this activation job
	//
	// Any of "created", "in-process", "completed", "failed".
	Status PortingOrdersActivationJobStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was created.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ActivateAt        respjson.Field
		ActivationType    respjson.Field
		ActivationWindows respjson.Field
		CreatedAt         respjson.Field
		RecordType        respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrdersActivationJob) RawJSON() string { return r.JSON.raw }
func (r *PortingOrdersActivationJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of this activation job
type PortingOrdersActivationJobActivationType string

const (
	PortingOrdersActivationJobActivationTypeScheduled PortingOrdersActivationJobActivationType = "scheduled"
	PortingOrdersActivationJobActivationTypeOnDemand  PortingOrdersActivationJobActivationType = "on-demand"
)

type PortingOrdersActivationJobActivationWindow struct {
	// ISO 8601 formatted date indicating when the activation window ends
	EndAt time.Time `json:"end_at" format:"date-time"`
	// ISO 8601 formatted date indicating when the activation window starts
	StartAt time.Time `json:"start_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndAt       respjson.Field
		StartAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrdersActivationJobActivationWindow) RawJSON() string { return r.JSON.raw }
func (r *PortingOrdersActivationJobActivationWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the status of this activation job
type PortingOrdersActivationJobStatus string

const (
	PortingOrdersActivationJobStatusCreated   PortingOrdersActivationJobStatus = "created"
	PortingOrdersActivationJobStatusInProcess PortingOrdersActivationJobStatus = "in-process"
	PortingOrdersActivationJobStatusCompleted PortingOrdersActivationJobStatus = "completed"
	PortingOrdersActivationJobStatusFailed    PortingOrdersActivationJobStatus = "failed"
)

type PortingOrderNewResponse struct {
	Data []PortingOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetResponse struct {
	Data PortingOrder                `json:"data"`
	Meta PortingOrderGetResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetResponseMeta struct {
	// Link to list all phone numbers
	PhoneNumbersURL string `json:"phone_numbers_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbersURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderUpdateResponse struct {
	Data PortingOrder                   `json:"data"`
	Meta PortingOrderUpdateResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderUpdateResponseMeta struct {
	// Link to list all phone numbers
	PhoneNumbersURL string `json:"phone_numbers_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbersURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderUpdateResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderUpdateResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetAllowedFocWindowsResponse struct {
	Data []PortingOrderGetAllowedFocWindowsResponseData `json:"data"`
	Meta PaginationMeta                                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetAllowedFocWindowsResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetAllowedFocWindowsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetAllowedFocWindowsResponseData struct {
	// ISO 8601 formatted date indicating the end of the range of foc window
	EndedAt time.Time `json:"ended_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating the start of the range of foc window.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndedAt     respjson.Field
		RecordType  respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetAllowedFocWindowsResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetAllowedFocWindowsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetExceptionTypesResponse struct {
	Data []shared.PortingOrdersExceptionType `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetExceptionTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetExceptionTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetRequirementsResponse struct {
	// Type of value expected on field_value field
	//
	// Any of "document", "textual".
	FieldType PortingOrderGetRequirementsResponseFieldType `json:"field_type"`
	// Identifies the document that satisfies this requirement
	FieldValue string `json:"field_value"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Status of the requirement
	RequirementStatus string `json:"requirement_status"`
	// Identifies the requirement type that meets this requirement
	RequirementType PortingOrderGetRequirementsResponseRequirementType `json:"requirement_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldType         respjson.Field
		FieldValue        respjson.Field
		RecordType        respjson.Field
		RequirementStatus respjson.Field
		RequirementType   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetRequirementsResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetRequirementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of value expected on field_value field
type PortingOrderGetRequirementsResponseFieldType string

const (
	PortingOrderGetRequirementsResponseFieldTypeDocument PortingOrderGetRequirementsResponseFieldType = "document"
	PortingOrderGetRequirementsResponseFieldTypeTextual  PortingOrderGetRequirementsResponseFieldType = "textual"
)

// Identifies the requirement type that meets this requirement
type PortingOrderGetRequirementsResponseRequirementType struct {
	// Identifies the requirement type
	ID string `json:"id"`
	// The acceptance criteria for the requirement type
	AcceptanceCriteria map[string]any `json:"acceptance_criteria"`
	// A description of the requirement type
	Description string `json:"description"`
	// An example of the requirement type
	Example string `json:"example"`
	// The name of the requirement type
	Name string `json:"name"`
	// The type of the requirement type
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AcceptanceCriteria respjson.Field
		Description        respjson.Field
		Example            respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetRequirementsResponseRequirementType) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetRequirementsResponseRequirementType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetSubRequestResponse struct {
	Data PortingOrderGetSubRequestResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetSubRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetSubRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetSubRequestResponseData struct {
	// Identifies the Port Request associated with the Porting Order
	PortRequestID string `json:"port_request_id"`
	// Identifies the Sub Request associated with the Porting Order
	SubRequestID string `json:"sub_request_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PortRequestID respjson.Field
		SubRequestID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderGetSubRequestResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderGetSubRequestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderNewParams struct {
	// The list of +E.164 formatted phone numbers
	PhoneNumbers []string `json:"phone_numbers,omitzero,required"`
	// A customer-specified reference number for customer bookkeeping purposes
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// A customer-specified group reference for customer bookkeeping purposes
	CustomerGroupReference param.Opt[string] `json:"customer_group_reference,omitzero"`
	paramObj
}

func (r PortingOrderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderGetParams struct {
	// Include the first 50 phone number objects in the results
	IncludePhoneNumbers param.Opt[bool] `query:"include_phone_numbers,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderGetParams]'s query parameters as `url.Values`.
func (r PortingOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderUpdateParams struct {
	CustomerGroupReference param.Opt[string] `json:"customer_group_reference,omitzero"`
	CustomerReference      param.Opt[string] `json:"customer_reference,omitzero"`
	// If present, we will read the current values from the specified Requirement Group
	// into the Documents and Requirements for this Porting Order. Note that any future
	// changes in the Requirement Group would have no impact on this Porting Order. We
	// will return an error if a specified Requirement Group conflicts with documents
	// or requirements in the same request.
	RequirementGroupID param.Opt[string]                          `json:"requirement_group_id,omitzero" format:"uuid"`
	WebhookURL         param.Opt[string]                          `json:"webhook_url,omitzero" format:"uri"`
	ActivationSettings PortingOrderUpdateParamsActivationSettings `json:"activation_settings,omitzero"`
	// Can be specified directly or via the `requirement_group_id` parameter.
	Documents                PortingOrderDocumentsParam                `json:"documents,omitzero"`
	EndUser                  PortingOrderEndUserParam                  `json:"end_user,omitzero"`
	Messaging                PortingOrderUpdateParamsMessaging         `json:"messaging,omitzero"`
	Misc                     PortingOrderMiscParam                     `json:"misc,omitzero"`
	PhoneNumberConfiguration PortingOrderPhoneNumberConfigurationParam `json:"phone_number_configuration,omitzero"`
	// List of requirements for porting numbers.
	Requirements []PortingOrderUpdateParamsRequirement `json:"requirements,omitzero"`
	UserFeedback PortingOrderUserFeedbackParam         `json:"user_feedback,omitzero"`
	paramObj
}

func (r PortingOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderUpdateParamsActivationSettings struct {
	// ISO 8601 formatted Date/Time requested for the FOC date
	FocDatetimeRequested param.Opt[time.Time] `json:"foc_datetime_requested,omitzero" format:"date-time"`
	paramObj
}

func (r PortingOrderUpdateParamsActivationSettings) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderUpdateParamsActivationSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderUpdateParamsActivationSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderUpdateParamsMessaging struct {
	// Indicates whether Telnyx will port messaging capabilities from the losing
	// carrier. If false, any messaging capabilities will stay with their current
	// provider.
	EnableMessaging param.Opt[bool] `json:"enable_messaging,omitzero"`
	paramObj
}

func (r PortingOrderUpdateParamsMessaging) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderUpdateParamsMessaging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderUpdateParamsMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies a value for a requirement on the Porting Order.
//
// The properties FieldValue, RequirementTypeID are required.
type PortingOrderUpdateParamsRequirement struct {
	// identifies the document or provides the text value that satisfies this
	// requirement
	FieldValue string `json:"field_value,required"`
	// Identifies the requirement type that the `field_value` fulfills
	RequirementTypeID string `json:"requirement_type_id,required"`
	paramObj
}

func (r PortingOrderUpdateParamsRequirement) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderUpdateParamsRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderUpdateParamsRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderListParams struct {
	// Include the first 50 phone number objects in the results
	IncludePhoneNumbers param.Opt[bool]  `query:"include_phone_numbers,omitzero" json:"-"`
	PageNumber          param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize            param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[customer_reference], filter[customer_group_reference],
	// filter[parent_support_key], filter[phone_numbers.country_code],
	// filter[phone_numbers.carrier_name], filter[misc.type],
	// filter[end_user.admin.entity_name], filter[end_user.admin.auth_person_name],
	// filter[activation_settings.fast_port_eligible],
	// filter[activation_settings.foc_datetime_requested][gt],
	// filter[activation_settings.foc_datetime_requested][lt],
	// filter[phone_numbers.phone_number][contains]
	Filter PortingOrderListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated sort parameter (deepObject style). Originally: sort[value]
	Sort PortingOrderListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParams]'s query parameters as `url.Values`.
func (r PortingOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[customer_reference], filter[customer_group_reference],
// filter[parent_support_key], filter[phone_numbers.country_code],
// filter[phone_numbers.carrier_name], filter[misc.type],
// filter[end_user.admin.entity_name], filter[end_user.admin.auth_person_name],
// filter[activation_settings.fast_port_eligible],
// filter[activation_settings.foc_datetime_requested][gt],
// filter[activation_settings.foc_datetime_requested][lt],
// filter[phone_numbers.phone_number][contains]
type PortingOrderListParamsFilter struct {
	// Filter results by customer_group_reference
	CustomerGroupReference param.Opt[string] `query:"customer_group_reference,omitzero" json:"-"`
	// Filter results by customer_reference
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter results by parent_support_key
	ParentSupportKey   param.Opt[string]                              `query:"parent_support_key,omitzero" json:"-"`
	ActivationSettings PortingOrderListParamsFilterActivationSettings `query:"activation_settings,omitzero" json:"-"`
	EndUser            PortingOrderListParamsFilterEndUser            `query:"end_user,omitzero" json:"-"`
	Misc               PortingOrderListParamsFilterMisc               `query:"misc,omitzero" json:"-"`
	PhoneNumbers       PortingOrderListParamsFilterPhoneNumbers       `query:"phone_numbers,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortingOrderListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderListParamsFilterActivationSettings struct {
	// Filter results by fast port eligible
	FastPortEligible param.Opt[bool] `query:"fast_port_eligible,omitzero" json:"-"`
	// FOC datetime range filtering operations
	FocDatetimeRequested PortingOrderListParamsFilterActivationSettingsFocDatetimeRequested `query:"foc_datetime_requested,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterActivationSettings]'s query
// parameters as `url.Values`.
func (r PortingOrderListParamsFilterActivationSettings) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// FOC datetime range filtering operations
type PortingOrderListParamsFilterActivationSettingsFocDatetimeRequested struct {
	// Filter results by foc date later than this value
	Gt param.Opt[string] `query:"gt,omitzero" json:"-"`
	// Filter results by foc date earlier than this value
	Lt param.Opt[string] `query:"lt,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [PortingOrderListParamsFilterActivationSettingsFocDatetimeRequested]'s query
// parameters as `url.Values`.
func (r PortingOrderListParamsFilterActivationSettingsFocDatetimeRequested) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderListParamsFilterEndUser struct {
	Admin PortingOrderListParamsFilterEndUserAdmin `query:"admin,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterEndUser]'s query parameters as
// `url.Values`.
func (r PortingOrderListParamsFilterEndUser) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderListParamsFilterEndUserAdmin struct {
	// Filter results by authorized person
	AuthPersonName param.Opt[string] `query:"auth_person_name,omitzero" json:"-"`
	// Filter results by person or company name
	EntityName param.Opt[string] `query:"entity_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterEndUserAdmin]'s query
// parameters as `url.Values`.
func (r PortingOrderListParamsFilterEndUserAdmin) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderListParamsFilterMisc struct {
	// Filter results by porting order type
	//
	// Any of "full", "partial".
	Type PortingOrderType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterMisc]'s query parameters as
// `url.Values`.
func (r PortingOrderListParamsFilterMisc) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderListParamsFilterPhoneNumbers struct {
	// Filter results by old service provider
	CarrierName param.Opt[string] `query:"carrier_name,omitzero" json:"-"`
	// Filter results by country ISO 3166-1 alpha-2 code
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Phone number pattern filtering operations
	PhoneNumber PortingOrderListParamsFilterPhoneNumbersPhoneNumber `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterPhoneNumbers]'s query
// parameters as `url.Values`.
func (r PortingOrderListParamsFilterPhoneNumbers) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Phone number pattern filtering operations
type PortingOrderListParamsFilterPhoneNumbersPhoneNumber struct {
	// Filter results by full or partial phone_number
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsFilterPhoneNumbersPhoneNumber]'s
// query parameters as `url.Values`.
func (r PortingOrderListParamsFilterPhoneNumbersPhoneNumber) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated sort parameter (deepObject style). Originally: sort[value]
type PortingOrderListParamsSort struct {
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at", "-created_at",
	// "activation_settings.foc_datetime_requested",
	// "-activation_settings.foc_datetime_requested".
	Value string `query:"value,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderListParamsSort]'s query parameters as
// `url.Values`.
func (r PortingOrderListParamsSort) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderGetLoaTemplateParams struct {
	// The identifier of the LOA configuration to use for the template. If not
	// provided, the default LOA configuration will be used.
	LoaConfigurationID param.Opt[string] `query:"loa_configuration_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderGetLoaTemplateParams]'s query parameters as
// `url.Values`.
func (r PortingOrderGetLoaTemplateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PortingOrderGetRequirementsParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingOrderGetRequirementsParams]'s query parameters as
// `url.Values`.
func (r PortingOrderGetRequirementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
