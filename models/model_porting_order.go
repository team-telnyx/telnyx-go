/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"encoding/json"
	"time"
)

// checks if the PortingOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortingOrder{}

// PortingOrder struct for PortingOrder
type PortingOrder struct {
	// Uniquely identifies this porting order
	Id *string `json:"id,omitempty"`
	// A customer-specified reference number for customer bookkeeping purposes
	CustomerReference *string `json:"customer_reference,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Status *PortingOrderStatus `json:"status,omitempty"`
	// A key to reference this porting order when contacting Telnyx customer support. This information is not available in draft porting orders.
	SupportKey *string `json:"support_key,omitempty"`
	// A key to reference for the porting order group when contacting Telnyx customer support. This information is not available for porting orders in `draft` state
	ParentSupportKey *string `json:"parent_support_key,omitempty"`
	// Count of phone numbers associated with this porting order
	PortingPhoneNumbersCount *int32 `json:"porting_phone_numbers_count,omitempty"`
	// Identifies the old service provider
	OldServiceProviderOcn *string `json:"old_service_provider_ocn,omitempty"`
	Documents *PortingOrderDocuments `json:"documents,omitempty"`
	Misc *PortingOrderMisc `json:"misc,omitempty"`
	EndUser *PortingOrderEndUser `json:"end_user,omitempty"`
	ActivationSettings *PortingOrderActivationSettings `json:"activation_settings,omitempty"`
	PhoneNumberConfiguration *PortingOrderPhoneNumberConfiguration `json:"phone_number_configuration,omitempty"`
	// The type of the phone number
	PhoneNumberType *string `json:"phone_number_type,omitempty"`
	// A description of the porting order
	Description *string `json:"description,omitempty"`
	// List of documentation requirements for porting numbers. Can be set directly or via the `requirement_group_id` parameter.
	Requirements []PortingOrderRequirement `json:"requirements,omitempty"`
	// Is true when the required documentation is met
	RequirementsMet *bool `json:"requirements_met,omitempty"`
	UserFeedback *PortingOrderUserFeedback `json:"user_feedback,omitempty"`
	// Identifies the user (or organization) who requested the porting order
	UserId *string `json:"user_id,omitempty"`
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	Messaging *PortingOrderMessaging `json:"messaging,omitempty"`
}

// NewPortingOrder instantiates a new PortingOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortingOrder() *PortingOrder {
	this := PortingOrder{}
	return &this
}

// NewPortingOrderWithDefaults instantiates a new PortingOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortingOrderWithDefaults() *PortingOrder {
	this := PortingOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortingOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortingOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortingOrder) SetId(v string) {
	o.Id = &v
}

// GetCustomerReference returns the CustomerReference field value if set, zero value otherwise.
func (o *PortingOrder) GetCustomerReference() string {
	if o == nil || IsNil(o.CustomerReference) {
		var ret string
		return ret
	}
	return *o.CustomerReference
}

// GetCustomerReferenceOk returns a tuple with the CustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetCustomerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerReference) {
		return nil, false
	}
	return o.CustomerReference, true
}

// HasCustomerReference returns a boolean if a field has been set.
func (o *PortingOrder) HasCustomerReference() bool {
	if o != nil && !IsNil(o.CustomerReference) {
		return true
	}

	return false
}

// SetCustomerReference gets a reference to the given string and assigns it to the CustomerReference field.
func (o *PortingOrder) SetCustomerReference(v string) {
	o.CustomerReference = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortingOrder) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortingOrder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortingOrder) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortingOrder) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortingOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortingOrder) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PortingOrder) GetStatus() PortingOrderStatus {
	if o == nil || IsNil(o.Status) {
		var ret PortingOrderStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetStatusOk() (*PortingOrderStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PortingOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PortingOrderStatus and assigns it to the Status field.
func (o *PortingOrder) SetStatus(v PortingOrderStatus) {
	o.Status = &v
}

// GetSupportKey returns the SupportKey field value if set, zero value otherwise.
func (o *PortingOrder) GetSupportKey() string {
	if o == nil || IsNil(o.SupportKey) {
		var ret string
		return ret
	}
	return *o.SupportKey
}

// GetSupportKeyOk returns a tuple with the SupportKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetSupportKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SupportKey) {
		return nil, false
	}
	return o.SupportKey, true
}

// HasSupportKey returns a boolean if a field has been set.
func (o *PortingOrder) HasSupportKey() bool {
	if o != nil && !IsNil(o.SupportKey) {
		return true
	}

	return false
}

// SetSupportKey gets a reference to the given string and assigns it to the SupportKey field.
func (o *PortingOrder) SetSupportKey(v string) {
	o.SupportKey = &v
}

// GetParentSupportKey returns the ParentSupportKey field value if set, zero value otherwise.
func (o *PortingOrder) GetParentSupportKey() string {
	if o == nil || IsNil(o.ParentSupportKey) {
		var ret string
		return ret
	}
	return *o.ParentSupportKey
}

// GetParentSupportKeyOk returns a tuple with the ParentSupportKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetParentSupportKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ParentSupportKey) {
		return nil, false
	}
	return o.ParentSupportKey, true
}

// HasParentSupportKey returns a boolean if a field has been set.
func (o *PortingOrder) HasParentSupportKey() bool {
	if o != nil && !IsNil(o.ParentSupportKey) {
		return true
	}

	return false
}

// SetParentSupportKey gets a reference to the given string and assigns it to the ParentSupportKey field.
func (o *PortingOrder) SetParentSupportKey(v string) {
	o.ParentSupportKey = &v
}

// GetPortingPhoneNumbersCount returns the PortingPhoneNumbersCount field value if set, zero value otherwise.
func (o *PortingOrder) GetPortingPhoneNumbersCount() int32 {
	if o == nil || IsNil(o.PortingPhoneNumbersCount) {
		var ret int32
		return ret
	}
	return *o.PortingPhoneNumbersCount
}

// GetPortingPhoneNumbersCountOk returns a tuple with the PortingPhoneNumbersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetPortingPhoneNumbersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PortingPhoneNumbersCount) {
		return nil, false
	}
	return o.PortingPhoneNumbersCount, true
}

// HasPortingPhoneNumbersCount returns a boolean if a field has been set.
func (o *PortingOrder) HasPortingPhoneNumbersCount() bool {
	if o != nil && !IsNil(o.PortingPhoneNumbersCount) {
		return true
	}

	return false
}

// SetPortingPhoneNumbersCount gets a reference to the given int32 and assigns it to the PortingPhoneNumbersCount field.
func (o *PortingOrder) SetPortingPhoneNumbersCount(v int32) {
	o.PortingPhoneNumbersCount = &v
}

// GetOldServiceProviderOcn returns the OldServiceProviderOcn field value if set, zero value otherwise.
func (o *PortingOrder) GetOldServiceProviderOcn() string {
	if o == nil || IsNil(o.OldServiceProviderOcn) {
		var ret string
		return ret
	}
	return *o.OldServiceProviderOcn
}

// GetOldServiceProviderOcnOk returns a tuple with the OldServiceProviderOcn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetOldServiceProviderOcnOk() (*string, bool) {
	if o == nil || IsNil(o.OldServiceProviderOcn) {
		return nil, false
	}
	return o.OldServiceProviderOcn, true
}

// HasOldServiceProviderOcn returns a boolean if a field has been set.
func (o *PortingOrder) HasOldServiceProviderOcn() bool {
	if o != nil && !IsNil(o.OldServiceProviderOcn) {
		return true
	}

	return false
}

// SetOldServiceProviderOcn gets a reference to the given string and assigns it to the OldServiceProviderOcn field.
func (o *PortingOrder) SetOldServiceProviderOcn(v string) {
	o.OldServiceProviderOcn = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *PortingOrder) GetDocuments() PortingOrderDocuments {
	if o == nil || IsNil(o.Documents) {
		var ret PortingOrderDocuments
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetDocumentsOk() (*PortingOrderDocuments, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *PortingOrder) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given PortingOrderDocuments and assigns it to the Documents field.
func (o *PortingOrder) SetDocuments(v PortingOrderDocuments) {
	o.Documents = &v
}

// GetMisc returns the Misc field value if set, zero value otherwise.
func (o *PortingOrder) GetMisc() PortingOrderMisc {
	if o == nil || IsNil(o.Misc) {
		var ret PortingOrderMisc
		return ret
	}
	return *o.Misc
}

// GetMiscOk returns a tuple with the Misc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetMiscOk() (*PortingOrderMisc, bool) {
	if o == nil || IsNil(o.Misc) {
		return nil, false
	}
	return o.Misc, true
}

// HasMisc returns a boolean if a field has been set.
func (o *PortingOrder) HasMisc() bool {
	if o != nil && !IsNil(o.Misc) {
		return true
	}

	return false
}

// SetMisc gets a reference to the given PortingOrderMisc and assigns it to the Misc field.
func (o *PortingOrder) SetMisc(v PortingOrderMisc) {
	o.Misc = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *PortingOrder) GetEndUser() PortingOrderEndUser {
	if o == nil || IsNil(o.EndUser) {
		var ret PortingOrderEndUser
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetEndUserOk() (*PortingOrderEndUser, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *PortingOrder) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given PortingOrderEndUser and assigns it to the EndUser field.
func (o *PortingOrder) SetEndUser(v PortingOrderEndUser) {
	o.EndUser = &v
}

// GetActivationSettings returns the ActivationSettings field value if set, zero value otherwise.
func (o *PortingOrder) GetActivationSettings() PortingOrderActivationSettings {
	if o == nil || IsNil(o.ActivationSettings) {
		var ret PortingOrderActivationSettings
		return ret
	}
	return *o.ActivationSettings
}

// GetActivationSettingsOk returns a tuple with the ActivationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetActivationSettingsOk() (*PortingOrderActivationSettings, bool) {
	if o == nil || IsNil(o.ActivationSettings) {
		return nil, false
	}
	return o.ActivationSettings, true
}

// HasActivationSettings returns a boolean if a field has been set.
func (o *PortingOrder) HasActivationSettings() bool {
	if o != nil && !IsNil(o.ActivationSettings) {
		return true
	}

	return false
}

// SetActivationSettings gets a reference to the given PortingOrderActivationSettings and assigns it to the ActivationSettings field.
func (o *PortingOrder) SetActivationSettings(v PortingOrderActivationSettings) {
	o.ActivationSettings = &v
}

// GetPhoneNumberConfiguration returns the PhoneNumberConfiguration field value if set, zero value otherwise.
func (o *PortingOrder) GetPhoneNumberConfiguration() PortingOrderPhoneNumberConfiguration {
	if o == nil || IsNil(o.PhoneNumberConfiguration) {
		var ret PortingOrderPhoneNumberConfiguration
		return ret
	}
	return *o.PhoneNumberConfiguration
}

// GetPhoneNumberConfigurationOk returns a tuple with the PhoneNumberConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetPhoneNumberConfigurationOk() (*PortingOrderPhoneNumberConfiguration, bool) {
	if o == nil || IsNil(o.PhoneNumberConfiguration) {
		return nil, false
	}
	return o.PhoneNumberConfiguration, true
}

// HasPhoneNumberConfiguration returns a boolean if a field has been set.
func (o *PortingOrder) HasPhoneNumberConfiguration() bool {
	if o != nil && !IsNil(o.PhoneNumberConfiguration) {
		return true
	}

	return false
}

// SetPhoneNumberConfiguration gets a reference to the given PortingOrderPhoneNumberConfiguration and assigns it to the PhoneNumberConfiguration field.
func (o *PortingOrder) SetPhoneNumberConfiguration(v PortingOrderPhoneNumberConfiguration) {
	o.PhoneNumberConfiguration = &v
}

// GetPhoneNumberType returns the PhoneNumberType field value if set, zero value otherwise.
func (o *PortingOrder) GetPhoneNumberType() string {
	if o == nil || IsNil(o.PhoneNumberType) {
		var ret string
		return ret
	}
	return *o.PhoneNumberType
}

// GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetPhoneNumberTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumberType) {
		return nil, false
	}
	return o.PhoneNumberType, true
}

// HasPhoneNumberType returns a boolean if a field has been set.
func (o *PortingOrder) HasPhoneNumberType() bool {
	if o != nil && !IsNil(o.PhoneNumberType) {
		return true
	}

	return false
}

// SetPhoneNumberType gets a reference to the given string and assigns it to the PhoneNumberType field.
func (o *PortingOrder) SetPhoneNumberType(v string) {
	o.PhoneNumberType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PortingOrder) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PortingOrder) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PortingOrder) SetDescription(v string) {
	o.Description = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *PortingOrder) GetRequirements() []PortingOrderRequirement {
	if o == nil || IsNil(o.Requirements) {
		var ret []PortingOrderRequirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetRequirementsOk() ([]PortingOrderRequirement, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *PortingOrder) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []PortingOrderRequirement and assigns it to the Requirements field.
func (o *PortingOrder) SetRequirements(v []PortingOrderRequirement) {
	o.Requirements = v
}

// GetRequirementsMet returns the RequirementsMet field value if set, zero value otherwise.
func (o *PortingOrder) GetRequirementsMet() bool {
	if o == nil || IsNil(o.RequirementsMet) {
		var ret bool
		return ret
	}
	return *o.RequirementsMet
}

// GetRequirementsMetOk returns a tuple with the RequirementsMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetRequirementsMetOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirementsMet) {
		return nil, false
	}
	return o.RequirementsMet, true
}

// HasRequirementsMet returns a boolean if a field has been set.
func (o *PortingOrder) HasRequirementsMet() bool {
	if o != nil && !IsNil(o.RequirementsMet) {
		return true
	}

	return false
}

// SetRequirementsMet gets a reference to the given bool and assigns it to the RequirementsMet field.
func (o *PortingOrder) SetRequirementsMet(v bool) {
	o.RequirementsMet = &v
}

// GetUserFeedback returns the UserFeedback field value if set, zero value otherwise.
func (o *PortingOrder) GetUserFeedback() PortingOrderUserFeedback {
	if o == nil || IsNil(o.UserFeedback) {
		var ret PortingOrderUserFeedback
		return ret
	}
	return *o.UserFeedback
}

// GetUserFeedbackOk returns a tuple with the UserFeedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetUserFeedbackOk() (*PortingOrderUserFeedback, bool) {
	if o == nil || IsNil(o.UserFeedback) {
		return nil, false
	}
	return o.UserFeedback, true
}

// HasUserFeedback returns a boolean if a field has been set.
func (o *PortingOrder) HasUserFeedback() bool {
	if o != nil && !IsNil(o.UserFeedback) {
		return true
	}

	return false
}

// SetUserFeedback gets a reference to the given PortingOrderUserFeedback and assigns it to the UserFeedback field.
func (o *PortingOrder) SetUserFeedback(v PortingOrderUserFeedback) {
	o.UserFeedback = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PortingOrder) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PortingOrder) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PortingOrder) SetUserId(v string) {
	o.UserId = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *PortingOrder) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *PortingOrder) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *PortingOrder) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *PortingOrder) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *PortingOrder) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *PortingOrder) SetRecordType(v string) {
	o.RecordType = &v
}

// GetMessaging returns the Messaging field value if set, zero value otherwise.
func (o *PortingOrder) GetMessaging() PortingOrderMessaging {
	if o == nil || IsNil(o.Messaging) {
		var ret PortingOrderMessaging
		return ret
	}
	return *o.Messaging
}

// GetMessagingOk returns a tuple with the Messaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingOrder) GetMessagingOk() (*PortingOrderMessaging, bool) {
	if o == nil || IsNil(o.Messaging) {
		return nil, false
	}
	return o.Messaging, true
}

// HasMessaging returns a boolean if a field has been set.
func (o *PortingOrder) HasMessaging() bool {
	if o != nil && !IsNil(o.Messaging) {
		return true
	}

	return false
}

// SetMessaging gets a reference to the given PortingOrderMessaging and assigns it to the Messaging field.
func (o *PortingOrder) SetMessaging(v PortingOrderMessaging) {
	o.Messaging = &v
}

func (o PortingOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortingOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CustomerReference) {
		toSerialize["customer_reference"] = o.CustomerReference
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SupportKey) {
		toSerialize["support_key"] = o.SupportKey
	}
	if !IsNil(o.ParentSupportKey) {
		toSerialize["parent_support_key"] = o.ParentSupportKey
	}
	if !IsNil(o.PortingPhoneNumbersCount) {
		toSerialize["porting_phone_numbers_count"] = o.PortingPhoneNumbersCount
	}
	if !IsNil(o.OldServiceProviderOcn) {
		toSerialize["old_service_provider_ocn"] = o.OldServiceProviderOcn
	}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Misc) {
		toSerialize["misc"] = o.Misc
	}
	if !IsNil(o.EndUser) {
		toSerialize["end_user"] = o.EndUser
	}
	if !IsNil(o.ActivationSettings) {
		toSerialize["activation_settings"] = o.ActivationSettings
	}
	if !IsNil(o.PhoneNumberConfiguration) {
		toSerialize["phone_number_configuration"] = o.PhoneNumberConfiguration
	}
	if !IsNil(o.PhoneNumberType) {
		toSerialize["phone_number_type"] = o.PhoneNumberType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	if !IsNil(o.RequirementsMet) {
		toSerialize["requirements_met"] = o.RequirementsMet
	}
	if !IsNil(o.UserFeedback) {
		toSerialize["user_feedback"] = o.UserFeedback
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.Messaging) {
		toSerialize["messaging"] = o.Messaging
	}
	return toSerialize, nil
}

type NullablePortingOrder struct {
	value *PortingOrder
	isSet bool
}

func (v NullablePortingOrder) Get() *PortingOrder {
	return v.value
}

func (v *NullablePortingOrder) Set(val *PortingOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePortingOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePortingOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortingOrder(val *PortingOrder) *NullablePortingOrder {
	return &NullablePortingOrder{value: val, isSet: true}
}

func (v NullablePortingOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortingOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


